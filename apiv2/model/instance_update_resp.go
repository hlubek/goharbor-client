// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InstanceUpdateResp instance update resp
//
// swagger:model InstanceUpdateResp
type InstanceUpdateResp struct {

	// ID of instance updated
	Updated int64 `json:"updated,omitempty"`
}

// Validate validates this instance update resp
func (m *InstanceUpdateResp) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *InstanceUpdateResp) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstanceUpdateResp) UnmarshalBinary(b []byte) error {
	var res InstanceUpdateResp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
