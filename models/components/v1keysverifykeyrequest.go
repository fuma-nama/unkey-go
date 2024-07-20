// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package components

import (
	"github.com/unkeyed/unkey-go/internal/utils"
)

// Authorization - Perform RBAC checks
type Authorization struct {
	// A query for which permissions you require
	Permissions *PermissionQuery `json:"permissions,omitempty"`
}

func (o *Authorization) GetPermissions() *PermissionQuery {
	if o == nil {
		return nil
	}
	return o.Permissions
}

// V1KeysVerifyKeyRequestRatelimit - Use 'ratelimits' with `[{ name: "default", cost: 2}]`
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type V1KeysVerifyKeyRequestRatelimit struct {
	// Override how many tokens are deducted during the ratelimit operation.
	Cost *int64 `default:"1" json:"cost"`
}

func (v V1KeysVerifyKeyRequestRatelimit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *V1KeysVerifyKeyRequestRatelimit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *V1KeysVerifyKeyRequestRatelimit) GetCost() *int64 {
	if o == nil {
		return nil
	}
	return o.Cost
}

type Ratelimits struct {
	// The name of the ratelimit.
	Name string `json:"name"`
	// Optionally override how expensive this operation is and how many tokens are deducted from the current limit.
	Cost *int64 `default:"1" json:"cost"`
	// Optionally override the limit.
	Limit *int64 `json:"limit,omitempty"`
	// Optionally override the ratelimit window duration.
	Duration *int64 `json:"duration,omitempty"`
}

func (r Ratelimits) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Ratelimits) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Ratelimits) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *Ratelimits) GetCost() *int64 {
	if o == nil {
		return nil
	}
	return o.Cost
}

func (o *Ratelimits) GetLimit() *int64 {
	if o == nil {
		return nil
	}
	return o.Limit
}

func (o *Ratelimits) GetDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

type V1KeysVerifyKeyRequest struct {
	// The id of the api where the key belongs to. This is optional for now but will be required soon.
	// The key will be verified against the api's configuration. If the key does not belong to the api, the verification will fail.
	APIID *string `json:"apiId,omitempty"`
	// The key to verify
	Key string `json:"key"`
	// Perform RBAC checks
	Authorization *Authorization `json:"authorization,omitempty"`
	// Use 'ratelimits' with `[{ name: "default", cost: 2}]`
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Ratelimit *V1KeysVerifyKeyRequestRatelimit `json:"ratelimit,omitempty"`
	// You can check against multiple ratelimits when verifying a key. Let's say you are building an app that uses AI under the hood and you want to limit your customers to 500 requests per hour, but also ensure they use up less than 20k tokens per day.
	//
	Ratelimits []Ratelimits `json:"ratelimits,omitempty"`
}

func (o *V1KeysVerifyKeyRequest) GetAPIID() *string {
	if o == nil {
		return nil
	}
	return o.APIID
}

func (o *V1KeysVerifyKeyRequest) GetKey() string {
	if o == nil {
		return ""
	}
	return o.Key
}

func (o *V1KeysVerifyKeyRequest) GetAuthorization() *Authorization {
	if o == nil {
		return nil
	}
	return o.Authorization
}

func (o *V1KeysVerifyKeyRequest) GetRatelimit() *V1KeysVerifyKeyRequestRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

func (o *V1KeysVerifyKeyRequest) GetRatelimits() []Ratelimits {
	if o == nil {
		return nil
	}
	return o.Ratelimits
}
