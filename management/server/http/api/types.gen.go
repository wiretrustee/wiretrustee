// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.1-0.20220912230023-4a1477f6a8ba DO NOT EDIT.
package api

import (
	"time"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for GroupPatchOperationOp.
const (
	GroupPatchOperationOpAdd     GroupPatchOperationOp = "add"
	GroupPatchOperationOpRemove  GroupPatchOperationOp = "remove"
	GroupPatchOperationOpReplace GroupPatchOperationOp = "replace"
)

// Defines values for GroupPatchOperationPath.
const (
	GroupPatchOperationPathName  GroupPatchOperationPath = "name"
	GroupPatchOperationPathPeers GroupPatchOperationPath = "peers"
)

// Defines values for NameserverNsType.
const (
	NameserverNsTypeUdp NameserverNsType = "udp"
)

// Defines values for NameserverGroupPatchOperationOp.
const (
	NameserverGroupPatchOperationOpAdd     NameserverGroupPatchOperationOp = "add"
	NameserverGroupPatchOperationOpRemove  NameserverGroupPatchOperationOp = "remove"
	NameserverGroupPatchOperationOpReplace NameserverGroupPatchOperationOp = "replace"
)

// Defines values for NameserverGroupPatchOperationPath.
const (
	NameserverGroupPatchOperationPathDescription NameserverGroupPatchOperationPath = "description"
	NameserverGroupPatchOperationPathEnabled     NameserverGroupPatchOperationPath = "enabled"
	NameserverGroupPatchOperationPathGroups      NameserverGroupPatchOperationPath = "groups"
	NameserverGroupPatchOperationPathName        NameserverGroupPatchOperationPath = "name"
	NameserverGroupPatchOperationPathNameservers NameserverGroupPatchOperationPath = "nameservers"
)

// Defines values for PatchMinimumOp.
const (
	PatchMinimumOpAdd     PatchMinimumOp = "add"
	PatchMinimumOpRemove  PatchMinimumOp = "remove"
	PatchMinimumOpReplace PatchMinimumOp = "replace"
)

// Defines values for RoutePatchOperationOp.
const (
	RoutePatchOperationOpAdd     RoutePatchOperationOp = "add"
	RoutePatchOperationOpRemove  RoutePatchOperationOp = "remove"
	RoutePatchOperationOpReplace RoutePatchOperationOp = "replace"
)

// Defines values for RoutePatchOperationPath.
const (
	RoutePatchOperationPathDescription RoutePatchOperationPath = "description"
	RoutePatchOperationPathEnabled     RoutePatchOperationPath = "enabled"
	RoutePatchOperationPathMasquerade  RoutePatchOperationPath = "masquerade"
	RoutePatchOperationPathMetric      RoutePatchOperationPath = "metric"
	RoutePatchOperationPathNetwork     RoutePatchOperationPath = "network"
	RoutePatchOperationPathNetworkId   RoutePatchOperationPath = "network_id"
	RoutePatchOperationPathPeer        RoutePatchOperationPath = "peer"
)

// Defines values for RulePatchOperationOp.
const (
	RulePatchOperationOpAdd     RulePatchOperationOp = "add"
	RulePatchOperationOpRemove  RulePatchOperationOp = "remove"
	RulePatchOperationOpReplace RulePatchOperationOp = "replace"
)

// Defines values for RulePatchOperationPath.
const (
	RulePatchOperationPathDescription  RulePatchOperationPath = "description"
	RulePatchOperationPathDestinations RulePatchOperationPath = "destinations"
	RulePatchOperationPathDisabled     RulePatchOperationPath = "disabled"
	RulePatchOperationPathFlow         RulePatchOperationPath = "flow"
	RulePatchOperationPathName         RulePatchOperationPath = "name"
	RulePatchOperationPathSources      RulePatchOperationPath = "sources"
)

// Group defines model for Group.
type Group struct {
	// Id Group ID
	Id string `json:"id"`

	// Name Group Name identifier
	Name string `json:"name"`

	// Peers List of peers object
	Peers []PeerMinimum `json:"peers"`

	// PeersCount Count of peers associated to the group
	PeersCount int `json:"peers_count"`
}

// GroupMinimum defines model for GroupMinimum.
type GroupMinimum struct {
	// Id Group ID
	Id string `json:"id"`

	// Name Group Name identifier
	Name string `json:"name"`

	// PeersCount Count of peers associated to the group
	PeersCount int `json:"peers_count"`
}

// GroupPatchOperation defines model for GroupPatchOperation.
type GroupPatchOperation struct {
	// Op Patch operation type
	Op GroupPatchOperationOp `json:"op"`

	// Path Group field to update in form /<field>
	Path GroupPatchOperationPath `json:"path"`

	// Value Values to be applied
	Value []string `json:"value"`
}

// GroupPatchOperationOp Patch operation type
type GroupPatchOperationOp string

// GroupPatchOperationPath Group field to update in form /<field>
type GroupPatchOperationPath string

// Nameserver defines model for Nameserver.
type Nameserver struct {
	// Ip Nameserver IP
	Ip string `json:"ip"`

	// NsType Nameserver Type
	NsType NameserverNsType `json:"ns_type"`

	// Port Nameserver Port
	Port int `json:"port"`
}

// NameserverNsType Nameserver Type
type NameserverNsType string

// NameserverGroup defines model for NameserverGroup.
type NameserverGroup struct {
	// Description Nameserver group  description
	Description string `json:"description"`

	// Enabled Nameserver group status
	Enabled bool `json:"enabled"`

	// Groups Nameserver group tag groups
	Groups []string `json:"groups"`

	// Id Nameserver group ID
	Id string `json:"id"`

	// Name Nameserver group name
	Name string `json:"name"`

	// Nameservers Nameserver group
	Nameservers []Nameserver `json:"nameservers"`
}

// NameserverGroupPatchOperation defines model for NameserverGroupPatchOperation.
type NameserverGroupPatchOperation struct {
	// Op Patch operation type
	Op NameserverGroupPatchOperationOp `json:"op"`

	// Path Nameserver group field to update in form /<field>
	Path NameserverGroupPatchOperationPath `json:"path"`

	// Value Values to be applied
	Value []string `json:"value"`
}

// NameserverGroupPatchOperationOp Patch operation type
type NameserverGroupPatchOperationOp string

// NameserverGroupPatchOperationPath Nameserver group field to update in form /<field>
type NameserverGroupPatchOperationPath string

// NameserverGroupRequest defines model for NameserverGroupRequest.
type NameserverGroupRequest struct {
	// Description Nameserver group  description
	Description string `json:"description"`

	// Enabled Nameserver group status
	Enabled bool `json:"enabled"`

	// Groups Nameserver group tag groups
	Groups []string `json:"groups"`

	// Name Nameserver group name
	Name string `json:"name"`

	// Nameservers Nameserver group
	Nameservers []Nameserver `json:"nameservers"`
}

// PatchMinimum defines model for PatchMinimum.
type PatchMinimum struct {
	// Op Patch operation type
	Op PatchMinimumOp `json:"op"`

	// Value Values to be applied
	Value []string `json:"value"`
}

// PatchMinimumOp Patch operation type
type PatchMinimumOp string

// Peer defines model for Peer.
type Peer struct {
	// Connected Peer to Management connection status
	Connected bool `json:"connected"`

	// Groups Groups that the peer belongs to
	Groups []GroupMinimum `json:"groups"`

	// Hostname Hostname of the machine
	Hostname string `json:"hostname"`

	// Id Peer ID
	Id string `json:"id"`

	// Ip Peer's IP address
	Ip string `json:"ip"`

	// LastSeen Last time peer connected to Netbird's management service
	LastSeen time.Time `json:"last_seen"`

	// Name Peer's hostname
	Name string `json:"name"`

	// Os Peer's operating system and version
	Os string `json:"os"`

	// SshEnabled Indicates whether SSH server is enabled on this peer
	SshEnabled bool `json:"ssh_enabled"`

	// UiVersion Peer's desktop UI version
	UiVersion *string `json:"ui_version,omitempty"`

	// UserId User ID of the user that enrolled this peer
	UserId *string `json:"user_id,omitempty"`

	// Version Peer's daemon or cli version
	Version string `json:"version"`
}

// PeerMinimum defines model for PeerMinimum.
type PeerMinimum struct {
	// Id Peer ID
	Id string `json:"id"`

	// Name Peer's hostname
	Name string `json:"name"`
}

// Route defines model for Route.
type Route struct {
	// Description Route description
	Description string `json:"description"`

	// Enabled Route status
	Enabled bool `json:"enabled"`

	// Id Route Id
	Id string `json:"id"`

	// Masquerade Indicate if peer should masquerade traffic to this route's prefix
	Masquerade bool `json:"masquerade"`

	// Metric Route metric number. Lowest number has higher priority
	Metric int `json:"metric"`

	// Network Network range in CIDR format
	Network string `json:"network"`

	// NetworkId Route network identifier, to group HA routes
	NetworkId string `json:"network_id"`

	// NetworkType Network type indicating if it is IPv4 or IPv6
	NetworkType string `json:"network_type"`

	// Peer Peer Identifier associated with route
	Peer string `json:"peer"`
}

// RoutePatchOperation defines model for RoutePatchOperation.
type RoutePatchOperation struct {
	// Op Patch operation type
	Op RoutePatchOperationOp `json:"op"`

	// Path Route field to update in form /<field>
	Path RoutePatchOperationPath `json:"path"`

	// Value Values to be applied
	Value []string `json:"value"`
}

// RoutePatchOperationOp Patch operation type
type RoutePatchOperationOp string

// RoutePatchOperationPath Route field to update in form /<field>
type RoutePatchOperationPath string

// RouteRequest defines model for RouteRequest.
type RouteRequest struct {
	// Description Route description
	Description string `json:"description"`

	// Enabled Route status
	Enabled bool `json:"enabled"`

	// Masquerade Indicate if peer should masquerade traffic to this route's prefix
	Masquerade bool `json:"masquerade"`

	// Metric Route metric number. Lowest number has higher priority
	Metric int `json:"metric"`

	// Network Network range in CIDR format
	Network string `json:"network"`

	// NetworkId Route network identifier, to group HA routes
	NetworkId string `json:"network_id"`

	// Peer Peer Identifier associated with route
	Peer string `json:"peer"`
}

// Rule defines model for Rule.
type Rule struct {
	// Description Rule friendly description
	Description string `json:"description"`

	// Destinations Rule destination groups
	Destinations []GroupMinimum `json:"destinations"`

	// Disabled Rules status
	Disabled bool `json:"disabled"`

	// Flow Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Id Rule ID
	Id string `json:"id"`

	// Name Rule name identifier
	Name string `json:"name"`

	// Sources Rule source groups
	Sources []GroupMinimum `json:"sources"`
}

// RuleMinimum defines model for RuleMinimum.
type RuleMinimum struct {
	// Description Rule friendly description
	Description string `json:"description"`

	// Disabled Rules status
	Disabled bool `json:"disabled"`

	// Flow Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Name Rule name identifier
	Name string `json:"name"`
}

// RulePatchOperation defines model for RulePatchOperation.
type RulePatchOperation struct {
	// Op Patch operation type
	Op RulePatchOperationOp `json:"op"`

	// Path Rule field to update in form /<field>
	Path RulePatchOperationPath `json:"path"`

	// Value Values to be applied
	Value []string `json:"value"`
}

// RulePatchOperationOp Patch operation type
type RulePatchOperationOp string

// RulePatchOperationPath Rule field to update in form /<field>
type RulePatchOperationPath string

// SetupKey defines model for SetupKey.
type SetupKey struct {
	// AutoGroups Setup key groups to auto-assign to peers registered with this key
	AutoGroups []string `json:"auto_groups"`

	// Expires Setup Key expiration date
	Expires time.Time `json:"expires"`

	// Id Setup Key ID
	Id string `json:"id"`

	// Key Setup Key value
	Key string `json:"key"`

	// LastUsed Setup key last usage date
	LastUsed time.Time `json:"last_used"`

	// Name Setup key name identifier
	Name string `json:"name"`

	// Revoked Setup key revocation status
	Revoked bool `json:"revoked"`

	// State Setup key status, "valid", "overused","expired" or "revoked"
	State string `json:"state"`

	// Type Setup key type, one-off for single time usage and reusable
	Type string `json:"type"`

	// UpdatedAt Setup key last update date
	UpdatedAt time.Time `json:"updated_at"`

	// UsedTimes Usage count of setup key
	UsedTimes int `json:"used_times"`

	// Valid Setup key validity status
	Valid bool `json:"valid"`
}

// SetupKeyRequest defines model for SetupKeyRequest.
type SetupKeyRequest struct {
	// AutoGroups Setup key groups to auto-assign to peers registered with this key
	AutoGroups []string `json:"auto_groups"`

	// ExpiresIn Expiration time in seconds
	ExpiresIn int `json:"expires_in"`

	// Name Setup Key name
	Name string `json:"name"`

	// Revoked Setup key revocation status
	Revoked bool `json:"revoked"`

	// Type Setup key type, one-off for single time usage and reusable
	Type string `json:"type"`
}

// User defines model for User.
type User struct {
	// AutoGroups Groups to auto-assign to peers registered by this user
	AutoGroups []string `json:"auto_groups"`

	// Email User's email address
	Email string `json:"email"`

	// Id User ID
	Id string `json:"id"`

	// Name User's name from idp provider
	Name string `json:"name"`

	// Role User's NetBird account role
	Role string `json:"role"`
}

// UserRequest defines model for UserRequest.
type UserRequest struct {
	// AutoGroups Groups to auto-assign to peers registered by this user
	AutoGroups []string `json:"auto_groups"`

	// Role User's NetBird account role
	Role string `json:"role"`
}

// PostApiGroupsJSONBody defines parameters for PostApiGroups.
type PostApiGroupsJSONBody struct {
	Name  string    `json:"name"`
	Peers *[]string `json:"peers,omitempty"`
}

// PatchApiGroupsIdJSONBody defines parameters for PatchApiGroupsId.
type PatchApiGroupsIdJSONBody = []GroupPatchOperation

// PutApiGroupsIdJSONBody defines parameters for PutApiGroupsId.
type PutApiGroupsIdJSONBody struct {
	Name  *string   `json:"Name,omitempty"`
	Peers *[]string `json:"Peers,omitempty"`
}

// PatchApiNameserversIdJSONBody defines parameters for PatchApiNameserversId.
type PatchApiNameserversIdJSONBody = []NameserverGroupPatchOperation

// PutApiPeersIdJSONBody defines parameters for PutApiPeersId.
type PutApiPeersIdJSONBody struct {
	Name       string `json:"name"`
	SshEnabled bool   `json:"ssh_enabled"`
}

// PatchApiRoutesIdJSONBody defines parameters for PatchApiRoutesId.
type PatchApiRoutesIdJSONBody = []RoutePatchOperation

// PostApiRulesJSONBody defines parameters for PostApiRules.
type PostApiRulesJSONBody struct {
	// Description Rule friendly description
	Description  string    `json:"description"`
	Destinations *[]string `json:"destinations,omitempty"`

	// Disabled Rules status
	Disabled bool `json:"disabled"`

	// Flow Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Name Rule name identifier
	Name    string    `json:"name"`
	Sources *[]string `json:"sources,omitempty"`
}

// PatchApiRulesIdJSONBody defines parameters for PatchApiRulesId.
type PatchApiRulesIdJSONBody = []RulePatchOperation

// PutApiRulesIdJSONBody defines parameters for PutApiRulesId.
type PutApiRulesIdJSONBody struct {
	// Description Rule friendly description
	Description  string    `json:"description"`
	Destinations *[]string `json:"destinations,omitempty"`

	// Disabled Rules status
	Disabled bool `json:"disabled"`

	// Flow Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Name Rule name identifier
	Name    string    `json:"name"`
	Sources *[]string `json:"sources,omitempty"`
}

// PostApiGroupsJSONRequestBody defines body for PostApiGroups for application/json ContentType.
type PostApiGroupsJSONRequestBody PostApiGroupsJSONBody

// PatchApiGroupsIdJSONRequestBody defines body for PatchApiGroupsId for application/json ContentType.
type PatchApiGroupsIdJSONRequestBody = PatchApiGroupsIdJSONBody

// PutApiGroupsIdJSONRequestBody defines body for PutApiGroupsId for application/json ContentType.
type PutApiGroupsIdJSONRequestBody PutApiGroupsIdJSONBody

// PostApiNameserversJSONRequestBody defines body for PostApiNameservers for application/json ContentType.
type PostApiNameserversJSONRequestBody = NameserverGroupRequest

// PatchApiNameserversIdJSONRequestBody defines body for PatchApiNameserversId for application/json ContentType.
type PatchApiNameserversIdJSONRequestBody = PatchApiNameserversIdJSONBody

// PutApiNameserversIdJSONRequestBody defines body for PutApiNameserversId for application/json ContentType.
type PutApiNameserversIdJSONRequestBody = NameserverGroupRequest

// PutApiPeersIdJSONRequestBody defines body for PutApiPeersId for application/json ContentType.
type PutApiPeersIdJSONRequestBody PutApiPeersIdJSONBody

// PostApiRoutesJSONRequestBody defines body for PostApiRoutes for application/json ContentType.
type PostApiRoutesJSONRequestBody = RouteRequest

// PatchApiRoutesIdJSONRequestBody defines body for PatchApiRoutesId for application/json ContentType.
type PatchApiRoutesIdJSONRequestBody = PatchApiRoutesIdJSONBody

// PutApiRoutesIdJSONRequestBody defines body for PutApiRoutesId for application/json ContentType.
type PutApiRoutesIdJSONRequestBody = RouteRequest

// PostApiRulesJSONRequestBody defines body for PostApiRules for application/json ContentType.
type PostApiRulesJSONRequestBody PostApiRulesJSONBody

// PatchApiRulesIdJSONRequestBody defines body for PatchApiRulesId for application/json ContentType.
type PatchApiRulesIdJSONRequestBody = PatchApiRulesIdJSONBody

// PutApiRulesIdJSONRequestBody defines body for PutApiRulesId for application/json ContentType.
type PutApiRulesIdJSONRequestBody PutApiRulesIdJSONBody

// PostApiSetupKeysJSONRequestBody defines body for PostApiSetupKeys for application/json ContentType.
type PostApiSetupKeysJSONRequestBody = SetupKeyRequest

// PutApiSetupKeysIdJSONRequestBody defines body for PutApiSetupKeysId for application/json ContentType.
type PutApiSetupKeysIdJSONRequestBody = SetupKeyRequest

// PutApiUsersIdJSONRequestBody defines body for PutApiUsersId for application/json ContentType.
type PutApiUsersIdJSONRequestBody = UserRequest
