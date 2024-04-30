//go:build !ios && !android

package networkwatcher

import (
	"context"
	"errors"
	"net"
	"net/netip"
	"os"
	"runtime/debug"

	"github.com/cenkalti/backoff/v4"
	log "github.com/sirupsen/logrus"

	"github.com/netbirdio/netbird/client/internal/routemanager"
)

// Start begins watching for network changes and calls the callback function and stops when a change is detected.
func (nw *NetworkWatcher) Start(ctx context.Context, callback func()) {
	if IsDisabled() {
		log.Info("Network watcher: disabled, not starting")
		return
	}

	if nw.cancel != nil {
		log.Warn("Network watcher: already running, stopping previous watcher")
		nw.Stop()
	}

	if ctx.Err() != nil {
		log.Info("Network watcher: not starting, context is already cancelled")
		return
	}

	ctx, nw.cancel = context.WithCancel(ctx)
	defer nw.Stop()

	var nexthop4, nexthop6 netip.Addr
	var intf4, intf6 *net.Interface

	operation := func() error {
		var errv4, errv6 error
		nexthop4, intf4, errv4 = routemanager.GetNextHop(netip.IPv4Unspecified())
		nexthop6, intf6, errv6 = routemanager.GetNextHop(netip.IPv6Unspecified())

		if errv4 != nil && errv6 != nil {
			return errors.New("failed to get default next hops")
		}

		if errv4 == nil {
			log.Debugf("Network watcher: IPv4 default route: %s, interface: %s", nexthop4, intf4.Name)
		}
		if errv6 == nil {
			log.Debugf("Network watcher: IPv6 default route: %s, interface: %s", nexthop6, intf6.Name)
		}

		// continue if either route was found
		return nil
	}

	expBackOff := backoff.WithContext(backoff.NewExponentialBackOff(), ctx)

	if err := backoff.Retry(operation, expBackOff); err != nil {
		log.Errorf("Network watcher: failed to get default next hops: %v", err)
		return
	}

	// recover in case sys ops panic
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Network watcher: panic occurred: %v, stack trace: %s", r, string(debug.Stack()))
		}
	}()

	if err := checkChange(ctx, nexthop4, intf4, nexthop6, intf6, callback); err != nil && !errors.Is(err, context.Canceled) {
		log.Errorf("Network watcher: failed to start: %v", err)
	}
}

// Stop stops the network watcher.
func (nw *NetworkWatcher) Stop() {
	if nw.cancel != nil {
		nw.cancel()
		nw.cancel = nil
		log.Info("Network watcher: stopped")
	}
}

func IsDisabled() bool {
	return os.Getenv("NB_DISABLE_NETWORK_WATCHER") == "true"
}