package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/netbirdio/netbird/relay/healthcheck"
	"github.com/netbirdio/netbird/relay/messages"
)

const (
	bufferSize = 8820
)

type Peer struct {
	log    *log.Entry
	idS    string
	idB    []byte
	conn   net.Conn
	connMu sync.RWMutex
	store  *Store
}

func NewPeer(id []byte, conn net.Conn, store *Store) *Peer {
	stringID := messages.HashIDToString(id)
	return &Peer{
		log:   log.WithField("peer_id", stringID),
		idS:   stringID,
		idB:   id,
		conn:  conn,
		store: store,
	}
}

func (p *Peer) Work() {
	ctx, cancel := context.WithCancel(context.Background())
	hc := healthcheck.NewSender(ctx)
	go p.healthcheck(ctx, hc)
	defer cancel()

	buf := make([]byte, bufferSize)
	for {
		n, err := p.conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				p.log.Errorf("failed to read message: %s", err)
			}
			return
		}

		msg := buf[:n]

		msgType, err := messages.DetermineClientMsgType(msg)
		if err != nil {
			p.log.Errorf("failed to determine message type: %s", err)
			return
		}
		switch msgType {
		case messages.MsgTypeHealthCheck:
			hc.OnHCResponse()
		case messages.MsgTypeTransport:
			peerID, err := messages.UnmarshalTransportID(msg)
			if err != nil {
				p.log.Errorf("failed to unmarshal transport message: %s", err)
				continue
			}
			stringPeerID := messages.HashIDToString(peerID)
			dp, ok := p.store.Peer(stringPeerID)
			if !ok {
				p.log.Errorf("peer not found: %s", stringPeerID)
				continue
			}
			err = messages.UpdateTransportMsg(msg, p.idB)
			if err != nil {
				p.log.Errorf("failed to update transport message: %s", err)
				continue
			}
			_, err = dp.Write(msg)
			if err != nil {
				p.log.Errorf("failed to write transport message to: %s", dp.String())
			}
		case messages.MsgTypeClose:
			p.log.Infof("peer exited gracefully")
			_ = p.conn.Close()
			return
		}
	}
}

// Write writes data to the connection
// it has been called by the remote peer
func (p *Peer) Write(b []byte) (int, error) {
	p.connMu.RLock()
	defer p.connMu.RUnlock()
	return p.conn.Write(b)
}

func (p *Peer) CloseGracefully(ctx context.Context) {
	p.connMu.Lock()
	_, err := p.writeWithTimeout(ctx, messages.MarshalCloseMsg())
	if err != nil {
		log.Errorf("failed to send close message to peer: %s", p.String())
	}

	err = p.conn.Close()
	if err != nil {
		log.Errorf("failed to close connection to peer: %s", err)
	}

	defer p.connMu.Unlock()
}

func (p *Peer) String() string {
	return p.idS
}

func (p *Peer) writeWithTimeout(ctx context.Context, buf []byte) (int, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	writeDone := make(chan struct{})
	var (
		n   int
		err error
	)

	go func() {
		_, err = p.conn.Write(buf)
		close(writeDone)
	}()

	select {
	case <-ctx.Done():
		return 0, fmt.Errorf("write operation timed out")
	case <-writeDone:
		return n, err
	}
}

func (p *Peer) healthcheck(ctx context.Context, hc *healthcheck.Sender) {
	for {
		select {
		case <-hc.HealthCheck:
			p.log.Debugf("sending healthcheck message")
			_, err := p.Write(messages.MarshalHealthcheck())
			if err != nil {
				p.log.Errorf("failed to send healthcheck message: %s", err)
				return
			}
		case <-hc.Timeout:
			p.log.Errorf("peer healthcheck timeout")
			_ = p.conn.Close()
			return
		case <-ctx.Done():
			return
		}
	}
}