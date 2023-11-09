package http

import (
	"bytes"
	"encoding/json"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"

	"github.com/netbirdio/netbird/management/server/http/api"

	"github.com/netbirdio/netbird/management/server/jwtclaims"

	"github.com/magiconair/properties/assert"

	"github.com/netbirdio/netbird/management/server"
	"github.com/netbirdio/netbird/management/server/mock_server"
)

const testPeerID = "test_peer"
const noUpdateChannelTestPeerID = "no-update-channel"

func initTestMetaData(peers ...*server.Peer) *PeersHandler {
	return &PeersHandler{
		accountManager: &mock_server.MockAccountManager{
			UpdatePeerFunc: func(accountID, userID string, update *server.Peer) (*server.Peer, error) {
				var p *server.Peer
				for _, peer := range peers {
					if update.ID == peer.ID {
						p = peer.Copy()
						break
					}
				}
				p.SSHEnabled = update.SSHEnabled
				p.LoginExpirationEnabled = update.LoginExpirationEnabled
				p.Name = update.Name
				return p, nil
			},
			GetPeerFunc: func(accountID, peerID, userID string) (*server.Peer, error) {
				var p *server.Peer
				for _, peer := range peers {
					if peerID == peer.ID {
						p = peer.Copy()
						break
					}
				}
				return p, nil
			},
			GetPeersFunc: func(accountID, userID string) ([]*server.Peer, error) {
				return peers, nil
			},
			GetAccountFromTokenFunc: func(claims jwtclaims.AuthorizationClaims) (*server.Account, *server.User, error) {
				user := server.NewAdminUser("test_user")
				return &server.Account{
					Id:     claims.AccountId,
					Domain: "hotmail.com",
					Peers: map[string]*server.Peer{
						peers[0].ID: peers[0],
					},
					Users: map[string]*server.User{
						"test_user": user,
					},
					Settings: &server.Settings{
						PeerLoginExpirationEnabled: true,
						PeerLoginExpiration:        time.Hour,
					},
					Network: &server.Network{
						Identifier: "ciclqisab2ss43jdn8q0",
						Net: net.IPNet{
							IP:   net.ParseIP("100.67.0.0"),
							Mask: net.IPv4Mask(255, 255, 0, 0),
						},
						Serial: 51,
					},
				}, user, nil
			},

			GetAllConnectedPeersFunc: func() (map[string]struct{}, error) {
				statuses := make(map[string]struct{})
				for _, peer := range peers {
					if peer.ID == noUpdateChannelTestPeerID {
						break
					}
					statuses[peer.ID] = struct{}{}
				}
				return statuses, nil
			},
		},
		claimsExtractor: jwtclaims.NewClaimsExtractor(
			jwtclaims.WithFromRequestContext(func(r *http.Request) jwtclaims.AuthorizationClaims {
				return jwtclaims.AuthorizationClaims{
					UserId:    "test_user",
					Domain:    "hotmail.com",
					AccountId: "test_id",
				}
			}),
		),
	}
}

// Tests the GetAllPeers endpoint reachable in the route /api/peers
// Use the metadata generated by initTestMetaData() to check for values
func TestGetPeers(t *testing.T) {

	peer := &server.Peer{
		ID:                     testPeerID,
		Key:                    "key",
		SetupKey:               "setupkey",
		IP:                     net.ParseIP("100.64.0.1"),
		Status:                 &server.PeerStatus{Connected: true},
		Name:                   "PeerName",
		LoginExpirationEnabled: false,
		Meta: server.PeerSystemMeta{
			Hostname:  "hostname",
			GoOS:      "GoOS",
			Kernel:    "kernel",
			Core:      "core",
			Platform:  "platform",
			OS:        "OS",
			WtVersion: "development",
		},
	}

	peer1 := peer.Copy()
	peer1.ID = noUpdateChannelTestPeerID

	expectedUpdatedPeer := peer.Copy()
	expectedUpdatedPeer.LoginExpirationEnabled = true
	expectedUpdatedPeer.SSHEnabled = true
	expectedUpdatedPeer.Name = "New Name"

	expectedPeer1 := peer1.Copy()
	expectedPeer1.Status.Connected = false

	tt := []struct {
		name           string
		expectedStatus int
		requestType    string
		requestPath    string
		requestBody    io.Reader
		expectedArray  bool
		expectedPeer   *server.Peer
	}{
		{
			name:           "GetPeersMetaData",
			requestType:    http.MethodGet,
			requestPath:    "/api/peers/",
			expectedStatus: http.StatusOK,
			expectedArray:  true,
			expectedPeer:   peer,
		},
		{
			name:           "GetPeer with update channel",
			requestType:    http.MethodGet,
			requestPath:    "/api/peers/" + testPeerID,
			expectedStatus: http.StatusOK,
			expectedArray:  false,
			expectedPeer:   peer,
		},
		{
			name:           "GetPeer with no update channel",
			requestType:    http.MethodGet,
			requestPath:    "/api/peers/" + peer1.ID,
			expectedStatus: http.StatusOK,
			expectedArray:  false,
			expectedPeer:   expectedPeer1,
		},
		{
			name:           "PutPeer",
			requestType:    http.MethodPut,
			requestPath:    "/api/peers/" + testPeerID,
			expectedStatus: http.StatusOK,
			expectedArray:  false,
			requestBody:    bytes.NewBufferString("{\"login_expiration_enabled\":true,\"name\":\"New Name\",\"ssh_enabled\":true}"),
			expectedPeer:   expectedUpdatedPeer,
		},
	}

	rr := httptest.NewRecorder()

	p := initTestMetaData(peer, peer1)

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			recorder := httptest.NewRecorder()
			req := httptest.NewRequest(tc.requestType, tc.requestPath, tc.requestBody)

			router := mux.NewRouter()
			router.HandleFunc("/api/peers/", p.GetAllPeers).Methods("GET")
			router.HandleFunc("/api/peers/{peerId}", p.HandlePeer).Methods("GET")
			router.HandleFunc("/api/peers/{peerId}", p.HandlePeer).Methods("PUT")
			router.ServeHTTP(recorder, req)

			res := recorder.Result()
			defer res.Body.Close()

			if status := rr.Code; status != tc.expectedStatus {
				t.Fatalf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			content, err := io.ReadAll(res.Body)
			if err != nil {
				t.Fatalf("I don't know what I expected; %v", err)
			}

			var got *api.Peer
			if tc.expectedArray {
				respBody := []*api.Peer{}
				err = json.Unmarshal(content, &respBody)
				if err != nil {
					t.Fatalf("Sent content is not in correct json format; %v", err)
				}

				// hardcode this check for now as we only have two peers in this suite
				assert.Equal(t, len(respBody), 2)
				assert.Equal(t, respBody[1].Connected, false)

				got = respBody[0]
			} else {
				got = &api.Peer{}
				err = json.Unmarshal(content, got)
				if err != nil {
					t.Fatalf("Sent content is not in correct json format; %v", err)
				}
			}

			t.Log(got)

			assert.Equal(t, got.Name, tc.expectedPeer.Name)
			assert.Equal(t, got.Version, tc.expectedPeer.Meta.WtVersion)
			assert.Equal(t, got.Ip, tc.expectedPeer.IP.String())
			assert.Equal(t, got.Os, "OS core")
			assert.Equal(t, got.LoginExpirationEnabled, tc.expectedPeer.LoginExpirationEnabled)
			assert.Equal(t, got.SshEnabled, tc.expectedPeer.SSHEnabled)
			assert.Equal(t, got.Connected, tc.expectedPeer.Status.Connected)
		})
	}
}
