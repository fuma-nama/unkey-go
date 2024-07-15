// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"errors"
	"fmt"
	"github.com/unkeyed/unkey-go/internal/utils"
	"github.com/unkeyed/unkey-go/models/components"
)

type SetRolesRoles2 struct {
	// The name of the role
	Name string `json:"name"`
	// Set to true to automatically create the role if it does not yet exist.
	Create *bool `json:"create,omitempty"`
}

func (o *SetRolesRoles2) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *SetRolesRoles2) GetCreate() *bool {
	if o == nil {
		return nil
	}
	return o.Create
}

type SetRolesRoles1 struct {
	// The id of the role.
	ID string `json:"id"`
}

func (o *SetRolesRoles1) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type SetRolesRolesType string

const (
	SetRolesRolesTypeSetRolesRoles1 SetRolesRolesType = "setRoles_roles_1"
	SetRolesRolesTypeSetRolesRoles2 SetRolesRolesType = "setRoles_roles_2"
)

type SetRolesRoles struct {
	SetRolesRoles1 *SetRolesRoles1
	SetRolesRoles2 *SetRolesRoles2

	Type SetRolesRolesType
}

func CreateSetRolesRolesSetRolesRoles1(setRolesRoles1 SetRolesRoles1) SetRolesRoles {
	typ := SetRolesRolesTypeSetRolesRoles1

	return SetRolesRoles{
		SetRolesRoles1: &setRolesRoles1,
		Type:           typ,
	}
}

func CreateSetRolesRolesSetRolesRoles2(setRolesRoles2 SetRolesRoles2) SetRolesRoles {
	typ := SetRolesRolesTypeSetRolesRoles2

	return SetRolesRoles{
		SetRolesRoles2: &setRolesRoles2,
		Type:           typ,
	}
}

func (u *SetRolesRoles) UnmarshalJSON(data []byte) error {

	var setRolesRoles1 SetRolesRoles1 = SetRolesRoles1{}
	if err := utils.UnmarshalJSON(data, &setRolesRoles1, "", true, true); err == nil {
		u.SetRolesRoles1 = &setRolesRoles1
		u.Type = SetRolesRolesTypeSetRolesRoles1
		return nil
	}

	var setRolesRoles2 SetRolesRoles2 = SetRolesRoles2{}
	if err := utils.UnmarshalJSON(data, &setRolesRoles2, "", true, true); err == nil {
		u.SetRolesRoles2 = &setRolesRoles2
		u.Type = SetRolesRolesTypeSetRolesRoles2
		return nil
	}

	return fmt.Errorf("could not unmarshal `%s` into any supported union types for SetRolesRoles", string(data))
}

func (u SetRolesRoles) MarshalJSON() ([]byte, error) {
	if u.SetRolesRoles1 != nil {
		return utils.MarshalJSON(u.SetRolesRoles1, "", true)
	}

	if u.SetRolesRoles2 != nil {
		return utils.MarshalJSON(u.SetRolesRoles2, "", true)
	}

	return nil, errors.New("could not marshal union type SetRolesRoles: all fields are null")
}

type SetRolesRequestBody struct {
	// The id of the key.
	KeyID string `json:"keyId"`
	// The roles you want to add to this key
	Roles []SetRolesRoles `json:"roles"`
}

func (o *SetRolesRequestBody) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *SetRolesRequestBody) GetRoles() []SetRolesRoles {
	if o == nil {
		return []SetRolesRoles{}
	}
	return o.Roles
}

type SetRolesResponseBody struct {
	// The id of the role. This is used internally
	ID string `json:"id"`
	// The name of the role
	Name string `json:"name"`
}

func (o *SetRolesResponseBody) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

func (o *SetRolesResponseBody) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

type SetRolesResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// All currently connected roles
	ResponseBodies []SetRolesResponseBody
}

func (o *SetRolesResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *SetRolesResponse) GetResponseBodies() []SetRolesResponseBody {
	if o == nil {
		return nil
	}
	return o.ResponseBodies
}
