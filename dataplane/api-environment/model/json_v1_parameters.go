// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JSONV1Parameters Json v1 parameters
// swagger:model JsonV1Parameters
type JSONV1Parameters struct {

	// credential Json
	// Required: true
	CredentialJSON *string `json:"credentialJson"`
}

// Validate validates this Json v1 parameters
func (m *JSONV1Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialJSON(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JSONV1Parameters) validateCredentialJSON(formats strfmt.Registry) error {

	if err := validate.Required("credentialJson", "body", m.CredentialJSON); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JSONV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JSONV1Parameters) UnmarshalBinary(b []byte) error {
	var res JSONV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}