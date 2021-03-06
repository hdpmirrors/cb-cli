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

// AccountTagRequest account tag request
// swagger:model AccountTagRequest
type AccountTagRequest struct {

	// key of the tag configurations
	// Required: true
	// Max Length: 127
	// Min Length: 3
	Key *string `json:"key"`

	// value of the tag configurations
	// Required: true
	// Max Length: 255
	// Min Length: 3
	Value *string `json:"value"`
}

// Validate validates this account tag request
func (m *AccountTagRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AccountTagRequest) validateKey(formats strfmt.Registry) error {

	if err := validate.Required("key", "body", m.Key); err != nil {
		return err
	}

	if err := validate.MinLength("key", "body", string(*m.Key), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("key", "body", string(*m.Key), 127); err != nil {
		return err
	}

	return nil
}

func (m *AccountTagRequest) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	if err := validate.MinLength("value", "body", string(*m.Value), 3); err != nil {
		return err
	}

	if err := validate.MaxLength("value", "body", string(*m.Value), 255); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AccountTagRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AccountTagRequest) UnmarshalBinary(b []byte) error {
	var res AccountTagRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
