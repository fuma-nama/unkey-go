// Code generated by Speakeasy (https://speakeasyapi.com). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// ErrConflictCode - A machine readable error code.
type ErrConflictCode string

const (
	ErrConflictCodeConflict ErrConflictCode = "CONFLICT"
)

func (e ErrConflictCode) ToPointer() *ErrConflictCode {
	return &e
}
func (e *ErrConflictCode) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "CONFLICT":
		*e = ErrConflictCode(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ErrConflictCode: %v", v)
	}
}

type ErrConflictError struct {
	// A machine readable error code.
	Code ErrConflictCode `json:"code"`
	// A link to our documentation with more details about this error code
	Docs string `json:"docs"`
	// A human readable explanation of what went wrong
	Message string `json:"message"`
	// Please always include the requestId in your error report
	RequestID string `json:"requestId"`
}

func (o *ErrConflictError) GetCode() ErrConflictCode {
	if o == nil {
		return ErrConflictCode("")
	}
	return o.Code
}

func (o *ErrConflictError) GetDocs() string {
	if o == nil {
		return ""
	}
	return o.Docs
}

func (o *ErrConflictError) GetMessage() string {
	if o == nil {
		return ""
	}
	return o.Message
}

func (o *ErrConflictError) GetRequestID() string {
	if o == nil {
		return ""
	}
	return o.RequestID
}

// ErrConflict - This response is sent when a request conflicts with the current state of the server.
type ErrConflict struct {
	Error_ ErrConflictError `json:"error"`
}

var _ error = &ErrConflict{}

func (e *ErrConflict) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
