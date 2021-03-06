// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectScanner project scanner
//
// swagger:model ProjectScanner
type ProjectScanner struct {

	// The identifier of the scanner registration
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this project scanner
func (m *ProjectScanner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ProjectScanner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectScanner) UnmarshalBinary(b []byte) error {
	var res ProjectScanner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
