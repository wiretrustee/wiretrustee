package client

import (
	"context"
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	auth "github.com/netbirdio/netbird/relay/auth/hmac"
)

const (
	maxConcurrentServers = 7
)

type connResult struct {
	RelayClient *Client
	Url         string
	Err         error
}

type ServerPicker struct {
	TokenStore *auth.TokenStore
	PeerID     string
}

func (sp *ServerPicker) PickServer(ctx context.Context, urls []string) (*Client, error) {
	totalServers := len(urls)

	connResultChan := make(chan connResult, totalServers)
	successChan := make(chan connResult, 1)

	concurrentLimiter := make(chan struct{}, maxConcurrentServers)
	for _, url := range urls {
		concurrentLimiter <- struct{}{}
		go func(url string) {
			defer func() { <-concurrentLimiter }()
			sp.startConnection(ctx, connResultChan, url)
		}(url)
	}

	go sp.processConnResults(connResultChan, successChan)

	select {
	case cr, ok := <-successChan:
		if !ok {
			return nil, errors.New("failed to connect to any relay server: all attempts failed")
		}
		log.Infof("chosen home Relay server: %s", cr.Url)
		return cr.RelayClient, nil
	case <-ctx.Done():
		return nil, fmt.Errorf("failed to connect to any relay server: %w", ctx.Err())
	}
}

func (sp *ServerPicker) startConnection(ctx context.Context, resultChan chan connResult, url string) {
	relayClient := NewClient(ctx, url, sp.TokenStore, sp.PeerID)
	err := relayClient.Connect()
	resultChan <- connResult{
		RelayClient: relayClient,
		Url:         url,
		Err:         err,
	}
}

func (sp *ServerPicker) processConnResults(resultChan chan connResult, successChan chan connResult) {
	var hasSuccess bool
	for cr := range resultChan {
		if cr.Err != nil {
			continue
		}
		log.Infof("connected to relay server: %s", cr.Url)

		if hasSuccess {
			log.Infof("closing unnecessary connection to %s", cr.Url)
			if err := cr.RelayClient.Close(); err != nil {
				log.Errorf("failed to close connection to %s: %v", cr.Url, err)
			}
			continue
		}

		hasSuccess = true
		successChan <- cr
	}
	close(successChan)
}