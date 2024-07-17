// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/internal/utils"
	"github.com/unkeyed/unkey-go/models/components"
)

type UpdateIdentityRatelimits struct {
	// The name of this limit. You will need to use this again when verifying a key.
	Name string `json:"name"`
	// How many requests may pass within a given window before requests are rejected.
	Limit float64 `json:"limit"`
	// The duration for each ratelimit window in milliseconds.
	Duration float64 `json:"duration"`
}

func (o *UpdateIdentityRatelimits) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *UpdateIdentityRatelimits) GetLimit() float64 {
	if o == nil {
		return 0.0
	}
	return o.Limit
}

func (o *UpdateIdentityRatelimits) GetDuration() float64 {
	if o == nil {
		return 0.0
	}
	return o.Duration
}

type UpdateIdentityRequestBody struct {
	// The id of the identity to update, use either `identityId` or `externalId`, if both are provided, `identityId` takes precedence.
	IdentityID *string `json:"identityId,omitempty"`
	// The externalId of the identity to update, use either `identityId` or `externalId`, if both are provided, `identityId` takes precedence.
	ExternalID *string `json:"externalId,omitempty"`
	// This is not yet used but here for future compatibility.
	Environment *string `default:"default" json:"environment"`
	// Attach metadata to this identity that you need to have access to when verifying a key.
	//
	// Set to `{}` to clear.
	//
	// This will be returned as part of the `verifyKey` response.
	//
	Meta map[string]any `json:"meta,omitempty"`
	// Attach ratelimits to this identity.
	//
	// This overwrites all existing ratelimits on this identity.
	// Setting an empty array will delete all existing ratelimits.
	//
	// When verifying keys, you can specify which limits you want to use and all keys attached to this identity, will share the limits.
	Ratelimits []UpdateIdentityRatelimits `json:"ratelimits,omitempty"`
}

func (u UpdateIdentityRequestBody) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateIdentityRequestBody) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateIdentityRequestBody) GetIdentityID() *string {
	if o == nil {
		return nil
	}
	return o.IdentityID
}

func (o *UpdateIdentityRequestBody) GetExternalID() *string {
	if o == nil {
		return nil
	}
	return o.ExternalID
}

func (o *UpdateIdentityRequestBody) GetEnvironment() *string {
	if o == nil {
		return nil
	}
	return o.Environment
}

func (o *UpdateIdentityRequestBody) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *UpdateIdentityRequestBody) GetRatelimits() []UpdateIdentityRatelimits {
	if o == nil {
		return nil
	}
	return o.Ratelimits
}

type UpdateIdentityResponseBody struct {
	// The id of the permission. This is used internally
	ID string `json:"id"`
	// The name of the permission
	Name string `json:"name"`
}

func (o *UpdateIdentityResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *UpdateIdentityResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type UpdateIdentityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The identity after the update.
	ResponseBodies []UpdateIdentityResponseBody
}

func (o *UpdateIdentityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateIdentityResponse) GetResponseBodies() []UpdateIdentityResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
