// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/unkeyed/unkey-go/internal/utils"
	"github.com/unkeyed/unkey-go/models/components"
)

// UpdateKeyType - Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
// Deprecated, use 'async' instead
//
// https://unkey.dev/docs/features/ratelimiting - Learn more
//
// Deprecated type: This will be removed in a future release, please migrate away from it as soon as possible.
type UpdateKeyType string

const (
	UpdateKeyTypeFast       UpdateKeyType = "fast"
	UpdateKeyTypeConsistent UpdateKeyType = "consistent"
)

func (e UpdateKeyType) ToPointer() *UpdateKeyType {
	return &e
}
func (e *UpdateKeyType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "fast":
		fallthrough
	case "consistent":
		*e = UpdateKeyType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateKeyType: %v", v)
	}
}

// UpdateKeyRatelimit - Unkey comes with per-key ratelimiting out of the box. Set `null` to disable.
type UpdateKeyRatelimit struct {
	// Fast ratelimiting doesn't add latency, while consistent ratelimiting is more accurate.
	// Deprecated, use 'async' instead
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	Type *UpdateKeyType `json:"type,omitempty"`
	// Asnyc ratelimiting doesn't add latency, while sync ratelimiting is slightly more accurate.
	Async *bool `default:"false" json:"async"`
	// The total amount of requests allowed in a single window.
	Limit int64 `json:"limit"`
	// How many tokens to refill during each refillInterval.
	// Deprecated, use 'limit' instead.
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RefillRate *int64 `json:"refillRate,omitempty"`
	// Determines the speed at which tokens are refilled, in milliseconds.
	// Deprecated, use 'duration'
	//
	// Deprecated field: This will be removed in a future release, please migrate away from it as soon as possible.
	RefillInterval *int64 `json:"refillInterval,omitempty"`
	// The duration of each ratelimit window, in milliseconds.
	// This field will become required in a future version.
	Duration *int64 `json:"duration,omitempty"`
}

func (u UpdateKeyRatelimit) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(u, "", false)
}

func (u *UpdateKeyRatelimit) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &u, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *UpdateKeyRatelimit) GetType() *UpdateKeyType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UpdateKeyRatelimit) GetAsync() *bool {
	if o == nil {
		return nil
	}
	return o.Async
}

func (o *UpdateKeyRatelimit) GetLimit() int64 {
	if o == nil {
		return 0
	}
	return o.Limit
}

func (o *UpdateKeyRatelimit) GetRefillRate() *int64 {
	if o == nil {
		return nil
	}
	return o.RefillRate
}

func (o *UpdateKeyRatelimit) GetRefillInterval() *int64 {
	if o == nil {
		return nil
	}
	return o.RefillInterval
}

func (o *UpdateKeyRatelimit) GetDuration() *int64 {
	if o == nil {
		return nil
	}
	return o.Duration
}

// UpdateKeyInterval - Unkey will automatically refill verifications at the set interval. If null is used the refill functionality will be removed from the key.
type UpdateKeyInterval string

const (
	UpdateKeyIntervalDaily   UpdateKeyInterval = "daily"
	UpdateKeyIntervalMonthly UpdateKeyInterval = "monthly"
)

func (e UpdateKeyInterval) ToPointer() *UpdateKeyInterval {
	return &e
}
func (e *UpdateKeyInterval) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "daily":
		fallthrough
	case "monthly":
		*e = UpdateKeyInterval(v)
		return nil
	default:
		return fmt.Errorf("invalid value for UpdateKeyInterval: %v", v)
	}
}

// UpdateKeyRefill - Unkey enables you to refill verifications for each key at regular intervals.
type UpdateKeyRefill struct {
	// Unkey will automatically refill verifications at the set interval. If null is used the refill functionality will be removed from the key.
	Interval UpdateKeyInterval `json:"interval"`
	// The amount of verifications to refill for each occurrence is determined individually for each key.
	Amount int64 `json:"amount"`
}

func (o *UpdateKeyRefill) GetInterval() UpdateKeyInterval {
	if o == nil {
		return UpdateKeyInterval("")
	}
	return o.Interval
}

func (o *UpdateKeyRefill) GetAmount() int64 {
	if o == nil {
		return 0
	}
	return o.Amount
}

type Roles struct {
	// The id of the role. Provide either `id` or `name`. If both are provided `id` is used.
	ID *string `json:"id,omitempty"`
	// Identify the role via its name. Provide either `id` or `name`. If both are provided `id` is used.
	Name *string `json:"name,omitempty"`
	// Set to true to automatically create the permissions they do not exist yet. Only works when specifying `name`.
	//                     Autocreating roles requires your root key to have the `rbac.*.create_role` permission, otherwise the request will get rejected
	Create *bool `json:"create,omitempty"`
}

func (o *Roles) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Roles) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Roles) GetCreate() *bool {
	if o == nil {
		return nil
	}
	return o.Create
}

type Permissions struct {
	// The id of the permission. Provide either `id` or `name`. If both are provided `id` is used.
	ID *string `json:"id,omitempty"`
	// Identify the permission via its name. Provide either `id` or `name`. If both are provided `id` is used.
	Name *string `json:"name,omitempty"`
	// Set to true to automatically create the permissions they do not exist yet. Only works when specifying `name`.
	//                     Autocreating permissions requires your root key to have the `rbac.*.create_permission` permission, otherwise the request will get rejected
	Create *bool `json:"create,omitempty"`
}

func (o *Permissions) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Permissions) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *Permissions) GetCreate() *bool {
	if o == nil {
		return nil
	}
	return o.Create
}

// UpdateKeyRequestBody - Update a key's configuration.
//
//	The `apis.<API_ID>.update_key` permission is required.
type UpdateKeyRequestBody struct {
	// The id of the key you want to modify
	KeyID string `json:"keyId"`
	// The name of the key
	Name *string `json:"name,omitempty"`
	// The id of the tenant associated with this key. Use whatever reference you have in your system to identify the tenant. When verifying the key, we will send this field back to you, so you know who is accessing your API.
	OwnerID *string `json:"ownerId,omitempty"`
	// Any additional metadata you want to store with the key
	Meta map[string]any `json:"meta,omitempty"`
	// The unix timestamp in milliseconds when the key will expire. If this field is null or undefined, the key is not expiring.
	Expires *int64 `json:"expires,omitempty"`
	// Unkey comes with per-key ratelimiting out of the box. Set `null` to disable.
	Ratelimit *UpdateKeyRatelimit `json:"ratelimit,omitempty"`
	// The number of requests that can be made with this key before it becomes invalid. Set `null` to disable.
	Remaining *int64 `json:"remaining,omitempty"`
	// Unkey enables you to refill verifications for each key at regular intervals.
	Refill *UpdateKeyRefill `json:"refill,omitempty"`
	// Set if key is enabled or disabled. If disabled, the key cannot be used to verify.
	Enabled *bool `json:"enabled,omitempty"`
	// The roles you want to set for this key. This overwrites all existing roles.
	//                 Setting roles requires the `rbac.*.add_role_to_key` permission.
	Roles []Roles `json:"roles,omitempty"`
	// The permissions you want to set for this key. This overwrites all existing permissions.
	//                 Setting permissions requires the `rbac.*.add_permission_to_key` permission.
	Permissions []Permissions `json:"permissions,omitempty"`
}

func (o *UpdateKeyRequestBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *UpdateKeyRequestBody) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *UpdateKeyRequestBody) GetOwnerID() *string {
	if o == nil {
		return nil
	}
	return o.OwnerID
}

func (o *UpdateKeyRequestBody) GetMeta() map[string]any {
	if o == nil {
		return nil
	}
	return o.Meta
}

func (o *UpdateKeyRequestBody) GetExpires() *int64 {
	if o == nil {
		return nil
	}
	return o.Expires
}

func (o *UpdateKeyRequestBody) GetRatelimit() *UpdateKeyRatelimit {
	if o == nil {
		return nil
	}
	return o.Ratelimit
}

func (o *UpdateKeyRequestBody) GetRemaining() *int64 {
	if o == nil {
		return nil
	}
	return o.Remaining
}

func (o *UpdateKeyRequestBody) GetRefill() *UpdateKeyRefill {
	if o == nil {
		return nil
	}
	return o.Refill
}

func (o *UpdateKeyRequestBody) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *UpdateKeyRequestBody) GetRoles() []Roles {
	if o == nil {
		return nil
	}
	return o.Roles
}

func (o *UpdateKeyRequestBody) GetPermissions() []Permissions {
	if o == nil {
		return nil
	}
	return o.Permissions
}

// UpdateKeyResponseBody - The key was successfully updated, it may take up to 30s for this to take effect in all regions
type UpdateKeyResponseBody struct {
}

type UpdateKeyResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The key was successfully updated, it may take up to 30s for this to take effect in all regions
	Object *UpdateKeyResponseBody
}

func (o *UpdateKeyResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *UpdateKeyResponse) GetObject() *UpdateKeyResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
