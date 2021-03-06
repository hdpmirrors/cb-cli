// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CloudStorageResponse cloud storage response
// swagger:model CloudStorageResponse
type CloudStorageResponse struct {

	// cloud storage account mapping
	AccountMapping *AccountMappingBase `json:"accountMapping,omitempty"`

	// aws
	Aws *AwsStorageParameters `json:"aws,omitempty"`

	// identities
	Identities []*StorageIdentityBase `json:"identities"`

	// locations
	Locations []*StorageLocationBase `json:"locations"`
}

// Validate validates this cloud storage response
func (m *CloudStorageResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountMapping(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentities(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudStorageResponse) validateAccountMapping(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountMapping) { // not required
		return nil
	}

	if m.AccountMapping != nil {
		if err := m.AccountMapping.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("accountMapping")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStorageResponse) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *CloudStorageResponse) validateIdentities(formats strfmt.Registry) error {

	if swag.IsZero(m.Identities) { // not required
		return nil
	}

	for i := 0; i < len(m.Identities); i++ {
		if swag.IsZero(m.Identities[i]) { // not required
			continue
		}

		if m.Identities[i] != nil {
			if err := m.Identities[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("identities" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CloudStorageResponse) validateLocations(formats strfmt.Registry) error {

	if swag.IsZero(m.Locations) { // not required
		return nil
	}

	for i := 0; i < len(m.Locations); i++ {
		if swag.IsZero(m.Locations[i]) { // not required
			continue
		}

		if m.Locations[i] != nil {
			if err := m.Locations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("locations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudStorageResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudStorageResponse) UnmarshalBinary(b []byte) error {
	var res CloudStorageResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
