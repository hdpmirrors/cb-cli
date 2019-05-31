// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DatabaseServerTestV4Request database server test v4 request
// swagger:model DatabaseServerTestV4Request
type DatabaseServerTestV4Request struct {

	// Unsaved database server config to be tested for connectivity
	DatabaseServer *DatabaseServerV4Request `json:"databaseServer,omitempty"`

	// Identifiers of saved database server config to be tested for connectivity
	ExistingDatabaseServer *DatabaseServerV4Identifiers `json:"existingDatabaseServer,omitempty"`
}

// Validate validates this database server test v4 request
func (m *DatabaseServerTestV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatabaseServer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExistingDatabaseServer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatabaseServerTestV4Request) validateDatabaseServer(formats strfmt.Registry) error {

	if swag.IsZero(m.DatabaseServer) { // not required
		return nil
	}

	if m.DatabaseServer != nil {
		if err := m.DatabaseServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("databaseServer")
			}
			return err
		}
	}

	return nil
}

func (m *DatabaseServerTestV4Request) validateExistingDatabaseServer(formats strfmt.Registry) error {

	if swag.IsZero(m.ExistingDatabaseServer) { // not required
		return nil
	}

	if m.ExistingDatabaseServer != nil {
		if err := m.ExistingDatabaseServer.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("existingDatabaseServer")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatabaseServerTestV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatabaseServerTestV4Request) UnmarshalBinary(b []byte) error {
	var res DatabaseServerTestV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}