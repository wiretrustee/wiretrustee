package server

import (
	"fmt"
	"io"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/netbirdio/netbird/relay/messages"
	"github.com/netbirdio/netbird/relay/server/listener"
	"github.com/netbirdio/netbird/relay/server/listener/udp"
)

// Server
// todo:
// authentication: provide JWT token via RPC call. The MGM server can forward the token to the agents.
// connection timeout handling
type Server struct {
	store *Store

	listener listener.Listener
}

func NewServer() *Server {
	return &Server{
		store: NewStore(),
	}
}

func (r *Server) Listen(address string) error {
	r.listener = udp.NewListener(address)
	return r.listener.Listen(r.accept)
}

func (r *Server) Close() error {
	if r.listener == nil {
		return nil
	}
	return r.listener.Close()
}

func (r *Server) accept(conn net.Conn) {
	peer, err := handShake(conn)
	if err != nil {
		log.Errorf("failed to handshake wiht %s: %s", conn.RemoteAddr(), err)
		cErr := conn.Close()
		if cErr != nil {
			log.Errorf("failed to close connection, %s: %s", conn.RemoteAddr(), cErr)
		}
		return
	}
	peer.Log.Debugf("peer connected from: %s", conn.RemoteAddr())

	r.store.AddPeer(peer)
	defer func() {
		peer.Log.Debugf("teardown connection")
		r.store.DeletePeer(peer)
	}()

	for {
		buf := make([]byte, 1500) // todo: optimize buffer size
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				peer.Log.Errorf("failed to read message: %s", err)
			}
			return
		}

		msgType, err := messages.DetermineClientMsgType(buf[:n])
		if err != nil {
			log.Errorf("failed to determine message type: %s", err)
			return
		}
		switch msgType {
		case messages.MsgTypeTransport:
			msg := buf[:n]
			peerID, err := messages.UnmarshalTransportID(msg)
			if err != nil {
				peer.Log.Errorf("failed to unmarshal transport message: %s", err)
				continue
			}
			go func() {
				stringPeerID := messages.HashIDToString(peerID)
				dp, ok := r.store.Peer(stringPeerID)
				if !ok {
					peer.Log.Errorf("peer not found: %s", stringPeerID)
					return
				}
				err := messages.UpdateTransportMsg(msg, peer.ID())
				if err != nil {
					peer.Log.Errorf("failed to update transport message: %s", err)
					return
				}
				_, err = dp.conn.Write(msg)
				if err != nil {
					peer.Log.Errorf("failed to write transport message to: %s", dp.String())
				}
				return
			}()
		}
	}
}

func handShake(conn net.Conn) (*Peer, error) {
	buf := make([]byte, 1500)
	n, err := conn.Read(buf)
	if err != nil {
		log.Errorf("failed to read message: %s", err)
		return nil, err
	}
	msgType, err := messages.DetermineClientMsgType(buf[:n])
	if err != nil {
		return nil, err
	}
	if msgType != messages.MsgTypeHello {
		tErr := fmt.Errorf("invalid message type")
		log.Errorf("failed to handshake: %s", tErr)
		return nil, tErr
	}
	peerId, err := messages.UnmarshalHelloMsg(buf[:n])
	if err != nil {
		log.Errorf("failed to handshake: %s", err)
		return nil, err
	}
	p := NewPeer(peerId, conn)

	msg := messages.MarshalHelloResponse()
	_, err = conn.Write(msg)
	return p, err
}
