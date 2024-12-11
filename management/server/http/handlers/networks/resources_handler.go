package networks

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/netbirdio/netbird/management/server/http/api"
	"github.com/netbirdio/netbird/management/server/http/configs"
	"github.com/netbirdio/netbird/management/server/http/util"
	"github.com/netbirdio/netbird/management/server/jwtclaims"
	"github.com/netbirdio/netbird/management/server/networks/resources"
	"github.com/netbirdio/netbird/management/server/networks/resources/types"
)

type resourceHandler struct {
	resourceManager  resources.Manager
	extractFromToken func(ctx context.Context, claims jwtclaims.AuthorizationClaims) (string, string, error)
	claimsExtractor  *jwtclaims.ClaimsExtractor
}

func addResourceEndpoints(resourcesManager resources.Manager, extractFromToken func(ctx context.Context, claims jwtclaims.AuthorizationClaims) (string, string, error), authCfg configs.AuthCfg, router *mux.Router) {
	resourceHandler := newResourceHandler(resourcesManager, extractFromToken, authCfg)
	router.HandleFunc("/networks/{networkId}/resources", resourceHandler.getAllResources).Methods("GET", "OPTIONS")
	router.HandleFunc("/networks/{networkId}/resources", resourceHandler.createResource).Methods("POST", "OPTIONS")
	router.HandleFunc("/networks/{networkId}/resources/{resourceId}", resourceHandler.getResource).Methods("GET", "OPTIONS")
	router.HandleFunc("/networks/{networkId}/resources/{resourceId}", resourceHandler.updateResource).Methods("PUT", "OPTIONS")
	router.HandleFunc("/networks/{networkId}/resources/{resourceId}", resourceHandler.deleteResource).Methods("DELETE", "OPTIONS")
}

func newResourceHandler(resourceManager resources.Manager, extractFromToken func(ctx context.Context, claims jwtclaims.AuthorizationClaims) (string, string, error), authCfg configs.AuthCfg) *resourceHandler {
	return &resourceHandler{
		resourceManager:  resourceManager,
		extractFromToken: extractFromToken,
		claimsExtractor: jwtclaims.NewClaimsExtractor(
			jwtclaims.WithAudience(authCfg.Audience),
			jwtclaims.WithUserIDClaim(authCfg.UserIDClaim),
		),
	}
}

func (h *resourceHandler) getAllResources(w http.ResponseWriter, r *http.Request) {
	claims := h.claimsExtractor.FromRequestContext(r)
	accountID, userID, err := h.extractFromToken(r.Context(), claims)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	networkID := mux.Vars(r)["networkId"]
	resources, err := h.resourceManager.GetAllResources(r.Context(), accountID, userID, networkID)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	var resourcesResponse []*api.NetworkResource
	for _, resource := range resources {
		resourcesResponse = append(resourcesResponse, resource.ToAPIResponse())
	}

	util.WriteJSONObject(r.Context(), w, resourcesResponse)
}

func (h *resourceHandler) createResource(w http.ResponseWriter, r *http.Request) {
	claims := h.claimsExtractor.FromRequestContext(r)
	accountID, userID, err := h.extractFromToken(r.Context(), claims)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	var req api.NetworkResourceRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.WriteErrorResponse("couldn't parse JSON request", http.StatusBadRequest, w)
		return
	}

	resource := &types.NetworkResource{}
	resource.FromAPIRequest(&req)

	resource.NetworkID = mux.Vars(r)["networkId"]
	resource.AccountID = accountID
	resource, err = h.resourceManager.CreateResource(r.Context(), userID, resource)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	util.WriteJSONObject(r.Context(), w, resource.ToAPIResponse())
}

func (h *resourceHandler) getResource(w http.ResponseWriter, r *http.Request) {
	claims := h.claimsExtractor.FromRequestContext(r)
	accountID, userID, err := h.extractFromToken(r.Context(), claims)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	networkID := mux.Vars(r)["networkId"]
	resourceID := mux.Vars(r)["resourceId"]
	resource, err := h.resourceManager.GetResource(r.Context(), accountID, userID, networkID, resourceID)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	util.WriteJSONObject(r.Context(), w, resource.ToAPIResponse())
}

func (h *resourceHandler) updateResource(w http.ResponseWriter, r *http.Request) {
	claims := h.claimsExtractor.FromRequestContext(r)
	accountID, userID, err := h.extractFromToken(r.Context(), claims)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	var req api.NetworkResourceRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		util.WriteErrorResponse("couldn't parse JSON request", http.StatusBadRequest, w)
		return
	}

	resource := &types.NetworkResource{}
	resource.FromAPIRequest(&req)

	resource.ID = mux.Vars(r)["resourceId"]
	resource.NetworkID = mux.Vars(r)["networkId"]
	resource.AccountID = accountID
	resource, err = h.resourceManager.UpdateResource(r.Context(), userID, resource)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	util.WriteJSONObject(r.Context(), w, resource.ToAPIResponse())
}

func (h *resourceHandler) deleteResource(w http.ResponseWriter, r *http.Request) {
	claims := h.claimsExtractor.FromRequestContext(r)
	accountID, userID, err := h.extractFromToken(r.Context(), claims)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	networkID := mux.Vars(r)["networkId"]
	resourceID := mux.Vars(r)["resourceId"]
	err = h.resourceManager.DeleteResource(r.Context(), accountID, userID, networkID, resourceID)
	if err != nil {
		util.WriteError(r.Context(), err, w)
		return
	}

	util.WriteJSONObject(r.Context(), w, util.EmptyObject{})
}
