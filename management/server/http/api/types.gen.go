// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"time"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Group defines model for Group.
type Group struct {
	// Group ID
	ID string `json:"ID"`

	// Group Name identifier
	Name string `json:"Name"`

	// List of peers object
	Peers []PeerMinimum `json:"Peers"`

	// Count of peers associated to the group
	PeersCount int `json:"PeersCount"`
}

// GroupMinimum defines model for GroupMinimum.
type GroupMinimum struct {
	// Group ID
	ID string `json:"ID"`

	// Group Name identifier
	Name string `json:"Name"`

	// Count of peers associated to the group
	PeersCount int `json:"PeersCount"`
}

// Peer defines model for Peer.
type Peer struct {
	// Provides information of who activated the Peer. User or Setup Key
	ActivatedBy struct {
		Type  string `json:"Type"`
		Value string `json:"Value"`
	} `json:"ActivatedBy"`

	// Peer to Management connection status
	Connected bool `json:"Connected"`

	// Groups that the peer belongs to
	Groups []GroupMinimum `json:"Groups"`

	// Peer ID
	ID string `json:"ID"`

	// Peer's IP address
	IP string `json:"IP"`

	// Last time peer connected to Netbird's management service
	LastSeen time.Time `json:"LastSeen"`

	// Peer's hostname
	Name string `json:"Name"`

	// Peer's operating system and version
	OS string `json:"OS"`

	// Peer's daemon or cli version
	Version string `json:"Version"`
}

// PeerMinimum defines model for PeerMinimum.
type PeerMinimum struct {
	// Peer ID
	ID string `json:"ID"`

	// Peer's hostname
	Name string `json:"Name"`
}

// Rule defines model for Rule.
type Rule struct {
	// Rule friendly description
	Description string `json:"Description"`

	// Rule destination groups
	Destination []GroupMinimum `json:"Destination"`

	// Rules status
	Disabled bool `json:"Disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"Flow"`

	// Rule ID
	ID string `json:"ID"`

	// Rule name identifier
	Name string `json:"Name"`

	// Rule source groups
	Source []GroupMinimum `json:"Source"`
}

// RuleMinimum defines model for RuleMinimum.
type RuleMinimum struct {
	// Rule friendly description
	Description string `json:"Description"`

	// Rules status
	Disabled bool `json:"Disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"Flow"`

	// Rule name identifier
	Name string `json:"Name"`
}

// SetupKey defines model for SetupKey.
type SetupKey struct {
	// Setup Key expiration date
	Expires time.Time `json:"Expires"`

	// Setup Key ID
	Id string `json:"Id"`

	// Setup Key value
	Key string `json:"Key"`

	// Setup key last usage date
	LastUsed time.Time `json:"LastUsed"`

	// Setup key name identifier
	Name string `json:"Name"`

	// Setup key revocation status
	Revoked bool `json:"Revoked"`

	// Setup key status, "valid", "overused","expired" or "revoked"
	State string `json:"State"`

	// Setup key type, one-off for single time usage and reusable
	Type string `json:"Type"`

	// Usage count of setup key
	UsedTimes int `json:"UsedTimes"`

	// Setup key validity status
	Valid bool `json:"Valid"`
}

// SetupKeyRequest defines model for SetupKeyRequest.
type SetupKeyRequest struct {
	// Expiration time in seconds
	ExpiresIn int `json:"ExpiresIn"`

	// Setup Key name
	Name string `json:"Name"`

	// Setup key revocation status
	Revoked bool `json:"Revoked"`

	// Setup key type, one-off for single time usage and reusable
	Type string `json:"Type"`
}

// User defines model for User.
type User struct {
	// User's email address
	Email string `json:"Email"`

	// User ID
	ID string `json:"ID"`

	// User's name from idp provider
	Name string `json:"Name"`

	// User's Netbird account role
	Role string `json:"Role"`
}

// PostApiGroupsJSONBody defines parameters for PostApiGroups.
type PostApiGroupsJSONBody struct {
	Name  string    `json:"Name"`
	Peers *[]string `json:"Peers,omitempty"`
}

// PutApiGroupsIdJSONBody defines parameters for PutApiGroupsId.
type PutApiGroupsIdJSONBody struct {
	Name  *string   `json:"Name,omitempty"`
	Peers *[]string `json:"Peers,omitempty"`
}

// PutApiPeersIdJSONBody defines parameters for PutApiPeersId.
type PutApiPeersIdJSONBody struct {
	Name string `json:"Name"`
}

// PostApiRulesJSONBody defines parameters for PostApiRules.
type PostApiRulesJSONBody struct {
	// Rule friendly description
	Description string    `json:"Description"`
	Destination *[]string `json:"Destination,omitempty"`

	// Rules status
	Disabled bool `json:"Disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"Flow"`

	// Rule name identifier
	Name   string    `json:"Name"`
	Source *[]string `json:"Source,omitempty"`
}

// PutApiRulesIdJSONBody defines parameters for PutApiRulesId.
type PutApiRulesIdJSONBody struct {
	// Rule friendly description
	Description string    `json:"Description"`
	Destination *[]string `json:"Destination,omitempty"`

	// Rules status
	Disabled bool `json:"Disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"Flow"`

	// Rule name identifier
	Name   string    `json:"Name"`
	Source *[]string `json:"Source,omitempty"`
}

// PostApiSetupKeysJSONBody defines parameters for PostApiSetupKeys.
type PostApiSetupKeysJSONBody = SetupKeyRequest

// PutApiSetupKeysIdJSONBody defines parameters for PutApiSetupKeysId.
type PutApiSetupKeysIdJSONBody = SetupKeyRequest

// PostApiGroupsJSONRequestBody defines body for PostApiGroups for application/json ContentType.
type PostApiGroupsJSONRequestBody PostApiGroupsJSONBody

// PutApiGroupsIdJSONRequestBody defines body for PutApiGroupsId for application/json ContentType.
type PutApiGroupsIdJSONRequestBody PutApiGroupsIdJSONBody

// PutApiPeersIdJSONRequestBody defines body for PutApiPeersId for application/json ContentType.
type PutApiPeersIdJSONRequestBody PutApiPeersIdJSONBody

// PostApiRulesJSONRequestBody defines body for PostApiRules for application/json ContentType.
type PostApiRulesJSONRequestBody PostApiRulesJSONBody

// PutApiRulesIdJSONRequestBody defines body for PutApiRulesId for application/json ContentType.
type PutApiRulesIdJSONRequestBody PutApiRulesIdJSONBody

// PostApiSetupKeysJSONRequestBody defines body for PostApiSetupKeys for application/json ContentType.
type PostApiSetupKeysJSONRequestBody = PostApiSetupKeysJSONBody

// PutApiSetupKeysIdJSONRequestBody defines body for PutApiSetupKeysId for application/json ContentType.
type PutApiSetupKeysIdJSONRequestBody = PutApiSetupKeysIdJSONBody
