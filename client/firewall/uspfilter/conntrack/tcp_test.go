package conntrack

import (
	"net"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTCPStateMachine(t *testing.T) {
	tracker := NewTCPTracker(DefaultTCPTimeout)
	defer tracker.Close()

	srcIP := net.ParseIP("100.64.0.1")
	dstIP := net.ParseIP("100.64.0.2")
	srcPort := uint16(12345)
	dstPort := uint16(80)

	t.Run("Security Tests", func(t *testing.T) {
		tests := []struct {
			name     string
			flags    uint8
			wantDrop bool
			desc     string
		}{
			{
				name:     "Block unsolicited SYN-ACK",
				flags:    TCPSyn | TCPAck,
				wantDrop: true,
				desc:     "Should block SYN-ACK without prior SYN",
			},
			{
				name:     "Block invalid SYN-FIN",
				flags:    TCPSyn | TCPFin,
				wantDrop: true,
				desc:     "Should block invalid SYN-FIN combination",
			},
			{
				name:     "Block unsolicited RST",
				flags:    TCPRst,
				wantDrop: true,
				desc:     "Should block RST without connection",
			},
			{
				name:     "Block unsolicited ACK",
				flags:    TCPAck,
				wantDrop: true,
				desc:     "Should block ACK without connection",
			},
			{
				name:     "Block data without connection",
				flags:    TCPAck | TCPPush,
				wantDrop: true,
				desc:     "Should block data without established connection",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				isValid := tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, tt.flags)
				require.Equal(t, !tt.wantDrop, isValid, tt.desc)
			})
		}
	})

	t.Run("Connection Flow Tests", func(t *testing.T) {
		tests := []struct {
			name string
			test func(*testing.T)
			desc string
		}{
			{
				name: "Normal Handshake",
				test: func(t *testing.T) {
					// Send initial SYN
					tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPSyn)

					// Receive SYN-ACK
					valid := tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPSyn|TCPAck)
					require.True(t, valid, "SYN-ACK should be allowed")

					// Send ACK
					tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPAck)

					// Test data transfer
					valid = tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPPush|TCPAck)
					require.True(t, valid, "Data should be allowed after handshake")
				},
			},
			{
				name: "Normal Close",
				test: func(t *testing.T) {
					// First establish connection
					establishConnection(t, tracker, srcIP, dstIP, srcPort, dstPort)

					// Send FIN
					tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPFin|TCPAck)

					// Receive ACK for FIN
					valid := tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPAck)
					require.True(t, valid, "ACK for FIN should be allowed")

					// Receive FIN from other side
					valid = tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPFin|TCPAck)
					require.True(t, valid, "FIN should be allowed")

					// Send final ACK
					tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPAck)
				},
			},
			{
				name: "RST During Connection",
				test: func(t *testing.T) {
					// First establish connection
					establishConnection(t, tracker, srcIP, dstIP, srcPort, dstPort)

					// Receive RST
					valid := tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPRst)
					require.True(t, valid, "RST should be allowed for established connection")

					// Verify connection is closed
					valid = tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPPush|TCPAck)
					require.False(t, valid, "Data should be blocked after RST")
				},
			},
			{
				name: "Simultaneous Close",
				test: func(t *testing.T) {
					// First establish connection
					establishConnection(t, tracker, srcIP, dstIP, srcPort, dstPort)

					// Both sides send FIN+ACK
					tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPFin|TCPAck)
					valid := tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPFin|TCPAck)
					require.True(t, valid, "Simultaneous FIN should be allowed")

					// Both sides send final ACK
					tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPAck)
					valid = tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPAck)
					require.True(t, valid, "Final ACKs should be allowed")
				},
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				tracker = NewTCPTracker(DefaultTCPTimeout)
				tt.test(t)
			})
		}
	})
}

// Helper to establish a TCP connection
func establishConnection(t *testing.T, tracker *TCPTracker, srcIP, dstIP net.IP, srcPort, dstPort uint16) {
	tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPSyn)

	valid := tracker.IsValidInbound(dstIP, srcIP, dstPort, srcPort, TCPSyn|TCPAck)
	require.True(t, valid, "SYN-ACK should be allowed")

	tracker.TrackOutbound(srcIP, dstIP, srcPort, dstPort, TCPAck)
}
