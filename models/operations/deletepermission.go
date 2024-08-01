// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type DeletePermissionRequestBody struct {
	// The id of the permission you want to delete.
	PermissionID string `json:"permissionId"`
}

func (o *DeletePermissionRequestBody) GetPermissionID() string {
	if o == nil {
		return ""
	}
	return o.PermissionID
}

// DeletePermissionResponseBody - Sucessfully deleted a permission
type DeletePermissionResponseBody struct {
}

type DeletePermissionResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// Sucessfully deleted a permission
	Object *DeletePermissionResponseBody
}

func (o *DeletePermissionResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeletePermissionResponse) GetObject() *DeletePermissionResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}
