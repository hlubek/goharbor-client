// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Platform platform
//
// swagger:model Platform
type Platform struct {

	// The features of the OS that the artifact applys to
	OsFeatures []string `json:"'os.features'"`

	// The version of the OS that the artifact applys to
	OsVersion string `json:"'os.version',omitempty"`

	// The architecture that the artifact applys to
	Architecture string `json:"architecture,omitempty"`

	// The OS that the artifact applys to
	Os string `json:"os,omitempty"`

	// The variant of the CPU
	Variant string `json:"variant,omitempty"`
}

// Validate validates this platform
func (m *Platform) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Platform) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Platform) UnmarshalBinary(b []byte) error {
	var res Platform
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}