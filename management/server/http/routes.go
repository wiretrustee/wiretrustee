package http

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/netbirdio/netbird/management/server"
	"github.com/netbirdio/netbird/management/server/http/api"
	"github.com/netbirdio/netbird/management/server/jwtclaims"
	"github.com/netbirdio/netbird/route"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

// Routes is the routes handler of the account
type Routes struct {
	jwtExtractor   jwtclaims.ClaimsExtractor
	accountManager server.AccountManager
	authAudience   string
}

// NewRoutes returns a new instance of Routes handler
func NewRoutes(accountManager server.AccountManager, authAudience string) *Routes {
	return &Routes{
		accountManager: accountManager,
		authAudience:   authAudience,
		jwtExtractor:   *jwtclaims.NewClaimsExtractor(nil),
	}
}

// GetAllRoutesHandler returns the list of routes for the account
func (h *Routes) GetAllRoutesHandler(w http.ResponseWriter, r *http.Request) {
	account, err := getJWTAccount(h.accountManager, h.jwtExtractor, h.authAudience, r)
	if err != nil {
		log.Error(err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	routes, err := h.accountManager.ListRoutes(account.Id)
	if err != nil {
		log.Error(err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}
	apiRoutes := make([]*api.Route, 0)
	for _, r := range routes {
		apiRoutes = append(apiRoutes, toRouteResponse(account, r))
	}

	writeJSONObject(w, apiRoutes)
}

// CreateRouteHandler handles route creation request
func (h *Routes) CreateRouteHandler(w http.ResponseWriter, r *http.Request) {
	account, err := getJWTAccount(h.accountManager, h.jwtExtractor, h.authAudience, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	var req api.PostApiRoutesJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	peerKey := req.Peer
	if req.Peer != "" {
		peer, err := h.accountManager.GetPeerByIP(account.Id, req.Peer)
		if err != nil {
			log.Error(err)
			http.Redirect(w, r, "/", http.StatusUnprocessableEntity)
			return
		}
		peerKey = peer.Key
	}

	_, newPrefix, err := route.ParsePrefix(req.Prefix)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't parse update prefix %s", req.Prefix), http.StatusBadRequest)
		return
	}

	newRoute, err := h.accountManager.CreateRoute(account.Id, newPrefix.String(), peerKey, req.Description, req.Masquerade, req.Metric, req.Enabled)
	if err != nil {
		log.Error(err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	resp := toRouteResponse(account, newRoute)

	writeJSONObject(w, &resp)
}

// UpdateRouteHandler handles update to a route identified by a given ID
func (h *Routes) UpdateRouteHandler(w http.ResponseWriter, r *http.Request) {
	account, err := getJWTAccount(h.accountManager, h.jwtExtractor, h.authAudience, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	routeID := vars["id"]
	if len(routeID) == 0 {
		http.Error(w, "invalid route Id", http.StatusBadRequest)
		return
	}

	oldroute, err := h.accountManager.GetRoute(account.Id, routeID)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't find route id %s", routeID), http.StatusNotFound)
		return
	}

	var req api.PutApiRoutesIdJSONBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	newRoute := oldroute.Copy()

	prefixType, newPrefix, err := route.ParsePrefix(req.Prefix)
	if err != nil {
		http.Error(w, fmt.Sprintf("couldn't parse update prefix %s for route id %s", req.Prefix, routeID), http.StatusBadRequest)
		return
	}

	peerKey := req.Peer
	if req.Peer != "" {
		peer, err := h.accountManager.GetPeerByIP(account.Id, req.Peer)
		if err != nil {
			log.Error(err)
			http.Redirect(w, r, "/", http.StatusUnprocessableEntity)
			return
		}
		peerKey = peer.Key
	}

	newRoute.Prefix = newPrefix
	newRoute.PrefixType = prefixType
	newRoute.Masquerade = req.Masquerade
	newRoute.Peer = peerKey
	newRoute.Metric = req.Metric
	newRoute.Description = req.Description
	newRoute.Enabled = req.Enabled

	if err := h.accountManager.SaveRoute(account.Id, newRoute); err != nil {
		log.Errorf("failed updating route \"%s\" under account %s %v", routeID, account.Id, err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	resp := toRouteResponse(account, newRoute)

	writeJSONObject(w, &resp)
}

// PatchRouteHandler handles patch updates to a route identified by a given ID
func (h *Routes) PatchRouteHandler(w http.ResponseWriter, r *http.Request) {
	account, err := getJWTAccount(h.accountManager, h.jwtExtractor, h.authAudience, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	routeID := vars["id"]
	if len(routeID) == 0 {
		http.Error(w, "invalid route Id", http.StatusBadRequest)
		return
	}

	_, err = h.accountManager.GetRoute(account.Id, routeID)
	if err != nil {
		log.Error(err)
		http.Error(w, fmt.Sprintf("couldn't find route id %s", routeID), http.StatusNotFound)
		return
	}

	var req api.PatchApiRoutesIdJSONRequestBody
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(req) == 0 {
		http.Error(w, "no patch instruction received", http.StatusBadRequest)
		return
	}

	var operations []server.RouteUpdateOperation

	for _, patch := range req {
		switch patch.Path {
		case api.RoutePatchOperationPathPrefix:
			if patch.Op != api.RoutePatchOperationOpReplace {
				http.Error(w, fmt.Sprintf("Prefix field only accepts replace operation, got %s", patch.Op),
					http.StatusBadRequest)
				return
			}
			operations = append(operations, server.RouteUpdateOperation{
				Type:   server.UpdateRoutePrefix,
				Values: patch.Value,
			})
		case api.RoutePatchOperationPathDescription:
			if patch.Op != api.RoutePatchOperationOpReplace {
				http.Error(w, fmt.Sprintf("Description field only accepts replace operation, got %s", patch.Op),
					http.StatusBadRequest)
				return
			}
			operations = append(operations, server.RouteUpdateOperation{
				Type:   server.UpdateRouteDescription,
				Values: patch.Value,
			})
		case api.RoutePatchOperationPathPeer:
			if patch.Op != api.RoutePatchOperationOpReplace {
				http.Error(w, fmt.Sprintf("Peer field only accepts replace operation, got %s", patch.Op),
					http.StatusBadRequest)
				return
			}
			if len(patch.Value) > 1 {
				http.Error(w, fmt.Sprintf("Value field only accepts 1 value, got %d", len(patch.Value)),
					http.StatusBadRequest)
				return
			}
			peerValue := patch.Value
			if patch.Value[0] != "" {
				peer, err := h.accountManager.GetPeerByIP(account.Id, patch.Value[0])
				if err != nil {
					log.Error(err)
					http.Redirect(w, r, "/", http.StatusUnprocessableEntity)
					return
				}
				peerValue = []string{peer.Key}
			}
			operations = append(operations, server.RouteUpdateOperation{
				Type:   server.UpdateRoutePeer,
				Values: peerValue,
			})
		case api.RoutePatchOperationPathMetric:
			if patch.Op != api.RoutePatchOperationOpReplace {
				http.Error(w, fmt.Sprintf("Metric field only accepts replace operation, got %s", patch.Op),
					http.StatusBadRequest)
				return
			}
			operations = append(operations, server.RouteUpdateOperation{
				Type:   server.UpdateRouteMetric,
				Values: patch.Value,
			})
		case api.RoutePatchOperationPathMasquerade:
			if patch.Op != api.RoutePatchOperationOpReplace {
				http.Error(w, fmt.Sprintf("Masquerade field only accepts replace operation, got %s", patch.Op),
					http.StatusBadRequest)
				return
			}
			operations = append(operations, server.RouteUpdateOperation{
				Type:   server.UpdateRouteMasquerade,
				Values: patch.Value,
			})
		case api.RoutePatchOperationPathEnabled:
			if patch.Op != api.RoutePatchOperationOpReplace {
				http.Error(w, fmt.Sprintf("Enabled field only accepts replace operation, got %s", patch.Op),
					http.StatusBadRequest)
				return
			}
			operations = append(operations, server.RouteUpdateOperation{
				Type:   server.UpdateRouteEnabled,
				Values: patch.Value,
			})
		default:
			http.Error(w, "invalid patch path", http.StatusBadRequest)
			return
		}
	}

	route, err := h.accountManager.UpdateRoute(account.Id, routeID, operations)

	if err != nil {
		errStatus, ok := status.FromError(err)
		if ok && errStatus.Code() == codes.Internal {
			http.Error(w, errStatus.String(), http.StatusInternalServerError)
			return
		}

		if ok && errStatus.Code() == codes.NotFound {
			http.Error(w, errStatus.String(), http.StatusNotFound)
			return
		}

		if ok && errStatus.Code() == codes.InvalidArgument {
			http.Error(w, errStatus.String(), http.StatusBadRequest)
			return
		}

		log.Errorf("failed updating route %s under account %s %v", routeID, account.Id, err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	resp := toRouteResponse(account, route)

	writeJSONObject(w, &resp)
}

// DeleteRouteHandler handles route deletion request
func (h *Routes) DeleteRouteHandler(w http.ResponseWriter, r *http.Request) {
	account, err := getJWTAccount(h.accountManager, h.jwtExtractor, h.authAudience, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	routeID := mux.Vars(r)["id"]
	if len(routeID) == 0 {
		http.Error(w, "invalid route ID", http.StatusBadRequest)
		return
	}

	err = h.accountManager.DeleteRoute(account.Id, routeID)
	if err != nil {
		log.Errorf("failed delete route %s under account %s %v", routeID, account.Id, err)
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	writeJSONObject(w, "")
}

// GetRouteHandler handles a route Get request identified by ID
func (h *Routes) GetRouteHandler(w http.ResponseWriter, r *http.Request) {
	account, err := getJWTAccount(h.accountManager, h.jwtExtractor, h.authAudience, r)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusInternalServerError)
		return
	}

	routeID := mux.Vars(r)["id"]
	if len(routeID) == 0 {
		http.Error(w, "invalid route ID", http.StatusBadRequest)
		return
	}

	foundRoute, err := h.accountManager.GetRoute(account.Id, routeID)
	if err != nil {
		http.Error(w, "route not found", http.StatusNotFound)
		return
	}

	writeJSONObject(w, toRouteResponse(account, foundRoute))
}

func toRouteResponse(account *server.Account, serverRoute *route.Route) *api.Route {
	var peerIP string
	if serverRoute.Peer != "" {
		peer, found := account.Peers[serverRoute.Peer]
		if !found {
			panic("peer id not found")
		}
		peerIP = peer.IP.String()
	}

	return &api.Route{
		Id:          serverRoute.ID,
		Description: serverRoute.Description,
		Enabled:     serverRoute.Enabled,
		Peer:        peerIP,
		Prefix:      serverRoute.Prefix.String(),
		PrefixType:  serverRoute.PrefixType.String(),
		Masquerade:  serverRoute.Masquerade,
		Metric:      serverRoute.Metric,
	}
}
