package jwtclaims

import (
	"net/http"

	"github.com/golang-jwt/jwt"
)

const (
	TokenUserProperty    = "user"
	AccountIDSuffix      = "wt_account_id"
	DomainIDSuffix       = "wt_account_domain"
	DomainCategorySuffix = "wt_account_domain_category"
	UserIDClaim          = "sub"
)

// Extract function type
type ExtractClaims func(r *http.Request) AuthorizationClaims

// ClaimsExtractor struct that holds the extract function
type ClaimsExtractor struct {
	authAudience string
	userIDClaim  string

	FromRequestContext ExtractClaims
}

// ClaimsExtractorOption is a function that configures the ClaimsExtractor
type ClaimsExtractorOption func(*ClaimsExtractor)

// WithAudience sets the audience for the extractor
func WithAudience(audience string) ClaimsExtractorOption {
	return func(c *ClaimsExtractor) {
		c.authAudience = audience
	}
}

// WithUserIDClaim sets the user id claim for the extractor
func WithUserIDClaim(userIDClaim string) ClaimsExtractorOption {
	return func(c *ClaimsExtractor) {
		c.userIDClaim = userIDClaim
	}
}

// WithFromRequestContext sets the function that extracts claims from the request context
func WithFromRequestContext(ec ExtractClaims) ClaimsExtractorOption {
	return func(c *ClaimsExtractor) {
		c.FromRequestContext = ec
	}
}

// NewClaimsExtractor returns an extractor, and if provided with a function with ExtractClaims signature,
// then it will use that logic. Uses ExtractClaimsFromRequestContext by default
func NewClaimsExtractor(options ...ClaimsExtractorOption) *ClaimsExtractor {
	ce := &ClaimsExtractor{}
	for _, option := range options {
		option(ce)
	}
	if ce.FromRequestContext == nil {
		ce.FromRequestContext = ce.fromRequestContext
	}
	if ce.userIDClaim == "" {
		ce.userIDClaim = UserIDClaim
	}
	return ce
}

// FromToken extracts claims from the token (after auth)
func (c *ClaimsExtractor) FromToken(token *jwt.Token) AuthorizationClaims {
	claims := token.Claims.(jwt.MapClaims)
	jwtClaims := AuthorizationClaims{}
	jwtClaims.UserId = claims[c.userIDClaim].(string)
	accountIdClaim, ok := claims[c.authAudience+AccountIDSuffix]
	if ok {
		jwtClaims.AccountId = accountIdClaim.(string)
	}
	domainClaim, ok := claims[c.authAudience+DomainIDSuffix]
	if ok {
		jwtClaims.Domain = domainClaim.(string)
	}
	domainCategoryClaim, ok := claims[c.authAudience+DomainCategorySuffix]
	if ok {
		jwtClaims.DomainCategory = domainCategoryClaim.(string)
	}
	return jwtClaims
}

// fromRequestContext extracts claims from the request context previously filled by the JWT token (after auth)
func (c *ClaimsExtractor) fromRequestContext(r *http.Request) AuthorizationClaims {
	if r.Context().Value(TokenUserProperty) == nil {
		return AuthorizationClaims{}
	}
	token := r.Context().Value(TokenUserProperty).(*jwt.Token)
	return c.FromToken(token)
}
