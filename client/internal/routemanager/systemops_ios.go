//go:build ios

package routemanager

import (
	"net/netip"
)

func addToRouteTableIfNoExists(prefix netip.Prefix, addr string, devName string) error {
	return nil
}

func removeFromRouteTableIfNonSystem(prefix netip.Prefix, addr string, devName string) error {
	return nil
}
