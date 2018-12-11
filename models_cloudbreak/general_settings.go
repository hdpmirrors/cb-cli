// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GeneralSettings general settings
// swagger:model GeneralSettings

type GeneralSettings struct {

	// credential resource name for the stack
	// Required: true
	CredentialName *string `json:"credentialName"`

	// name of the stack
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`
}

/* polymorph GeneralSettings credentialName false */

/* polymorph GeneralSettings name false */

// Validate validates this general settings
func (m *GeneralSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredentialName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GeneralSettings) validateCredentialName(formats strfmt.Registry) error {

	if err := validate.Required("credentialName", "body", m.CredentialName); err != nil {
		return err
	}

	return nil
}

func (m *GeneralSettings) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GeneralSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneralSettings) UnmarshalBinary(b []byte) error {
	var res GeneralSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}