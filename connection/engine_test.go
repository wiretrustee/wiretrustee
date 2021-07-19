package connection

import (
	"bytes"
	"context"
	ice "github.com/pion/ice/v2"
	log "github.com/sirupsen/logrus"
	"github.com/wiretrustee/wiretrustee/iface"
	sig "github.com/wiretrustee/wiretrustee/signal"
	"golang.zx2c4.com/wireguard/wgctrl"
	"golang.zx2c4.com/wireguard/wgctrl/wgtypes"
	"io/ioutil"
	"strings"
	"testing"
	"time"
)

var engine *Engine
var testKey wgtypes.Key
var testPeer Peer

const ifaceName = "utun9991"

func Test_Start(t *testing.T) {
	level, _ := log.ParseLevel("Debug")
	log.SetLevel(level)
	b := bytes.NewBufferString("")
	log.SetOutput(b)

	var err error
	testKey, err = wgtypes.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()

	iceUrl, err := ice.ParseURL("stun:stun.wiretrustee.com:3468")
	if err != nil {
		t.Fatal(err)
	}
	var stunURLs = []*ice.URL{iceUrl}

	iFaceBlackList := make(map[string]struct{})

	signalClient, err := sig.NewClient(ctx, "signal.wiretrustee.com:10000", testKey)
	if err != nil {
		t.Fatal(err)
	}

	engine = NewEngine(signalClient, stunURLs, ifaceName, "10.99.91.1/24", iFaceBlackList)

	var emptyPeer []Peer
	err = engine.Start(testKey, emptyPeer)
	if err != nil {
		t.Fatal(err)
	}
	wg, err := wgctrl.New()
	if err != nil {
		t.Fatal(err)
	}
	defer wg.Close()

	_, err = wg.Device(ifaceName)
	if err != nil {
		t.Fatal(err)
	}

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	expectedMSG := "connected to the Signal Exchange Stream"
	if !strings.Contains(string(out), expectedMSG) {
		t.Fatalf("expected \"%s\" got \"%s\"", expectedMSG, string(out))
	}
}

func TestEngine_InitializePeerWithoutRemote(t *testing.T) {
	b := bytes.NewBufferString("")
	log.SetOutput(b)
	tmpKey, err := wgtypes.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	testPeer = Peer{
		tmpKey.PublicKey().String(),
		"10.99.91.2/32",
	}
	go engine.InitializePeer(iface.WgPort, testKey, testPeer)
	// Let the connections initialize
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for {
		status := engine.GetPeerConnectionStatus(testPeer.WgPubKey)
		err = ctx.Err()
		if (status != nil && *status == StatusConnecting) || err != nil {
			if err != nil {
				t.Fatal(err)
			}
			//success
			break
		}
	}
}

func TestEngine_Initialize2PeersWithoutRemote(t *testing.T) {
	b := bytes.NewBufferString("")
	log.SetOutput(b)
	tmpKey1, err := wgtypes.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	tmpKey2, err := wgtypes.GenerateKey()
	if err != nil {
		t.Fatal(err)
	}
	testPeer1 := Peer{
		tmpKey1.PublicKey().String(),
		"10.99.91.2/32",
	}
	testPeer2 := Peer{
		tmpKey2.PublicKey().String(),
		"10.99.91.3/32",
	}
	go engine.InitializePeer(iface.WgPort, testKey, testPeer1)
	go engine.InitializePeer(iface.WgPort, testKey, testPeer2)
	// Let the connections initialize
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	for {
		status1 := engine.GetPeerConnectionStatus(testPeer1.WgPubKey)
		status2 := engine.GetPeerConnectionStatus(testPeer2.WgPubKey)
		err = ctx.Err()
		if (status1 != nil && status2 != nil) || err != nil {
			if err != nil {
				t.Fatal(err)
			}
			if *status1 == StatusConnecting && *status2 == StatusConnecting {
				//success
				break
			}
		}
	}
}

func TestEngine_RemovePeerConnectionWithoutRemote(t *testing.T) {
	b := bytes.NewBufferString("")
	log.SetOutput(b)

	err := engine.RemovePeerConnection(testPeer)
	if err != nil {
		t.Fatal(err)
	}
	// Let the connections initialize
	if _, exists := engine.conns[testPeer.WgPubKey]; exists {
		t.Fatal("couldn't remove peer")
	}
	// Let the connections close
	time.Sleep(100 * time.Millisecond)
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	expectedMSG := "removing connection attempt with Peer: " + testPeer.WgPubKey + ", not retrying"
	if !strings.Contains(string(out), expectedMSG) {
		t.Fatalf("expected \"%s\" got \"%s\"", expectedMSG, string(out))
	}
}

func Test_CloseInterface(t *testing.T) {
	err := iface.Close()
	if err != nil {
		t.Fatal(err)
	}
}
