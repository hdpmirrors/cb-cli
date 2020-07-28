// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// PlatformResourceGroupResponse platform resource group response
// swagger:model PlatformResourceGroupResponse
type PlatformResourceGroupResponse struct {

	// name
	Name string `json:"name,omitempty"`
}

// Validate validates this platform resource group response
func (m *PlatformResourceGroupResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PlatformResourceGroupResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformResourceGroupResponse) UnmarshalBinary(b []byte) error {
	var res PlatformResourceGroupResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}