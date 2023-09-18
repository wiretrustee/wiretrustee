package server

import (
	"testing"
	"time"

	"github.com/netbirdio/netbird/management/proto"
)

//var peersUpdater *PeersUpdateManager

func TestCreateChannel(t *testing.T) {
	peer := "test-create"
	peersUpdater := NewPeersUpdateManager()
	defer peersUpdater.CloseChannel(peer)

	if peersUpdater.Len() != 0 {
		t.Error("peersUpdated should not have any channels yet")
	}

	if ch := peersUpdater.GetChannel(peer); ch != nil {
		t.Errorf("We should not have channel for %s yet", peer)
	}

	_ = peersUpdater.CreateChannel(peer)
	if ch := peersUpdater.GetChannel(peer); ch == nil {
		t.Error("Error creating the channel")
	}

	if peersUpdater.Len() != 1 {
		t.Error("peersUpdated should have 1 channel")
	}
}

func TestSendUpdate(t *testing.T) {
	peer := "test-sendupdate"
	peersUpdater := NewPeersUpdateManager()
	update1 := &UpdateMessage{Update: &proto.SyncResponse{
		NetworkMap: &proto.NetworkMap{
			Serial: 0,
		},
	}}
	_ = peersUpdater.CreateChannel(peer)
	if ch := peersUpdater.GetChannel(peer); ch == nil {
		t.Error("Error creating the channel")
	}
	err := peersUpdater.SendUpdate(peer, update1)
	if err != nil {
		t.Error("Error sending update: ", err)
	}
	select {
	case <-peersUpdater.GetChannel(peer):
	default:
		t.Error("Update wasn't send")
	}

	for range [channelBufferSize]int{} {
		err = peersUpdater.SendUpdate(peer, update1)
		if err != nil {
			t.Errorf("got an early error sending update: %v ", err)
		}
	}

	update2 := &UpdateMessage{Update: &proto.SyncResponse{
		NetworkMap: &proto.NetworkMap{
			Serial: 10,
		},
	}}

	err = peersUpdater.SendUpdate(peer, update2)
	if err != nil {
		t.Error("update shouldn't return an error when channel buffer is full")
	}
	timeout := time.After(5 * time.Second)
	for range [channelBufferSize]int{} {
		select {
		case <-timeout:
			t.Error("timed out reading previously sent updates")
		case updateReader := <-peersUpdater.GetChannel(peer):
			if updateReader.Update.NetworkMap.Serial == update2.Update.NetworkMap.Serial {
				t.Error("got the update that shouldn't have been sent")
			}
		}
	}

}

func TestCloseChannel(t *testing.T) {
	peer := "test-close"
	peersUpdater := NewPeersUpdateManager()
	_ = peersUpdater.CreateChannel(peer)
	if ch := peersUpdater.GetChannel(peer); ch == nil {
		t.Error("Error creating the channel")
	}
	peersUpdater.CloseChannel(peer)
	if ch := peersUpdater.GetChannel(peer); ch != nil {
		t.Error("Error closing the channel")
	}
}
