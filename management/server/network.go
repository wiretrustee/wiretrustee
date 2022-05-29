package server

import (
	"fmt"
	"github.com/rs/xid"
	"math/rand"
	"net"
	"sync"
	"time"
)

type NetworkMap struct {
	Peers   []*Peer
	Network *Network
}

type Network struct {
	Id  string
	Net net.IPNet
	Dns string
	// Serial is an ID that increments by 1 when any change to the network happened (e.g. new peer has been added).
	// Used to synchronize state to the client apps.
	Serial uint64

	mu sync.Mutex `json:"-"`
}

// NewNetwork creates a new Network initializing it with a Serial=0
func NewNetwork() *Network {
	return &Network{
		Id:     xid.New().String(),
		Net:    net.IPNet{IP: net.ParseIP("100.64.0.0"), Mask: net.IPMask{255, 255, 0, 0}},
		Dns:    "",
		Serial: 0}
}

// IncSerial increments Serial by 1 reflecting that the network state has been changed
func (n *Network) IncSerial() {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Serial = n.Serial + 1
}

// CurrentSerial returns the Network.Serial of the network (latest state id)
func (n *Network) CurrentSerial() uint64 {
	n.mu.Lock()
	defer n.mu.Unlock()
	return n.Serial
}

func (n *Network) Copy() *Network {
	return &Network{
		Id:     n.Id,
		Net:    n.Net,
		Dns:    n.Dns,
		Serial: n.Serial,
	}
}

// AllocatePeerIP pics an available IP from an net.IPNet.
// This method considers already taken IPs and reuses IPs if there are gaps in takenIps
// E.g. if ipNet=100.30.0.0/16 and takenIps=[100.30.0.1, 100.30.0.4] then the result would be 100.30.0.2 or 100.30.0.3
func AllocatePeerIP(ipNet net.IPNet, takenIps []net.IP) (net.IP, error) {
	takenIpMap := make(map[string]struct{})
	takenIpMap[ipNet.IP.String()] = struct{}{}
	for _, ip := range takenIps {
		takenIpMap[ip.String()] = struct{}{}
	}

	ips, _, err := generateIPs(&ipNet, takenIpMap)
	if err != nil {
		return nil, fmt.Errorf("failed allocating new IP for the ipNet %s and takenIps %s", ipNet.String(), takenIps)
	}

	if len(ips) == 0 {
		return nil, fmt.Errorf("failed allocating new IP for the ipNet %s - network is out of IPs", ipNet.String())
	}

	// pick a random IP
	s := rand.NewSource(time.Now().Unix())
	r := rand.New(s)
	intn := r.Intn(len(ips))

	return ips[intn], nil
}

// generateIPs generates a list of all possible IPs of the given network excluding IPs specified in the exclusion list
func generateIPs(ipNet *net.IPNet, exclusions map[string]struct{}) ([]net.IP, int, error) {

	var ips []net.IP
	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); incIP(ip) {
		if _, ok := exclusions[ip.String()]; !ok && ip[3] != 0 {
			ips = append(ips, copyIP(ip))
		}
	}

	// remove network address and broadcast address
	lenIPs := len(ips)
	switch {
	case lenIPs < 2:
		return ips, lenIPs, nil

	default:
		return ips[1 : len(ips)-1], lenIPs - 2, nil
	}
}

func copyIP(ip net.IP) net.IP {
	dup := make(net.IP, len(ip))
	copy(dup, ip)
	return dup
}

func incIP(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
