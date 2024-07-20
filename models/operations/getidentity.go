// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type GetIdentityRequest struct {
	IdentityID *string `queryParam:"style=form,explode=true,name=identityId"`
	ExternalID *string `queryParam:"style=form,explode=true,name=externalId"`
}

func (o *GetIdentityRequest) GetIdentityID() *string {
	if o == nil {
		return nil
	}
	return o.IdentityID
}

func (o *GetIdentityRequest) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

type GetIdentityRatelimits struct {
	// The name of this limit. You will need to use this again when verifying a key.
	Name string `json:"name"`
	// How many requests may pass within a given window before requests are rejected.
	Limit int64 `json:"limit"`
	// The duration for each ratelimit window in milliseconds.
	Duration int64 `json:"duration"`
}

func (o *GetIdentityRatelimits) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GetIdentityRatelimits) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *GetIdentityRatelimits) GetDuration() int64 {
	if o == nil {
		return 0
	}
	return o.Duration
}

// GetIdentityResponseBody - The configuration for an api
type GetIdentityResponseBody struct {
	// The id of this identity. Used to interact with unkey's API
	ID string `json:"id"`
	// The id in your system
	ExternalID string `json:"externalId"`
	// The meta object defined for this identity.
	Meta map[string]any `json:"meta"`
	// When verifying keys, you can specify which limits you want to use and all keys attached to this identity, will share the limits.
	Ratelimits []GetIdentityRatelimits `json:"ratelimits"`
}

func (o *GetIdentityResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *GetIdentityResponseBody) GetExternalID() string {
	if o == nil {
		return ""
	}
	return o.ExternalID
}

func (o *GetIdentityResponseBody) GetMeta() map[string]any {
	if o == nil {
		return map[string]any{}
	}
	return o.Meta
}

func (o *GetIdentityResponseBody) GetRatelimits() []GetIdentityRatelimits {
	if o == nil {
		return []GetIdentityRatelimits{}
	}
	return o.Ratelimits
}

type GetIdentityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The configuration for an api
	Object *GetIdentityResponseBody
}

func (o *GetIdentityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *GetIdentityResponse) GetObject() *GetIdentityResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
