// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package operations

import (
	"github.com/unkeyed/unkey-go/models/components"
)

type DeleteIdentityRequestBody struct {
	// The id of the identity to delete
	IdentityID string `json:"identityId"`
}

func (o *DeleteIdentityRequestBody) GetIdentityID() string {
	if o == nil {
		return ""
	}
	return o.IdentityID
}

// DeleteIdentityResponseBody - The identity was successfully deleted, it may take up to 30s for this to take effect in all regions
type DeleteIdentityResponseBody struct {
}

type DeleteIdentityResponse struct {
	HTTPMeta components.HTTPMetadata `json:"-"`
	// The identity was successfully deleted, it may take up to 30s for this to take effect in all regions
	Object *DeleteIdentityResponseBody
}

func (o *DeleteIdentityResponse) GetHTTPMeta() components.HTTPMetadata {
	if o == nil {
		return components.HTTPMetadata{}
	}
	return o.HTTPMeta
}

func (o *DeleteIdentityResponse) GetObject() *DeleteIdentityResponseBody {
	if o == nil {
		return nil
	}
	return o.Object
}