// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SupportedVersionsV4Response supported versions v4 response
// swagger:model SupportedVersionsV4Response
type SupportedVersionsV4Response struct {

	// supported versions
	// Unique: true
	SupportedVersions []*SupportedVersionV4Response `json:"supportedVersions"`
}

// Validate validates this supported versions v4 response
func (m *SupportedVersionsV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSupportedVersions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SupportedVersionsV4Response) validateSupportedVersions(formats strfmt.Registry) error {

	if swag.IsZero(m.SupportedVersions) { // not required
		return nil
	}

	if err := validate.UniqueItems("supportedVersions", "body", m.SupportedVersions); err != nil {
		return err
	}

	for i := 0; i < len(m.SupportedVersions); i++ {
		if swag.IsZero(m.SupportedVersions[i]) { // not required
			continue
		}

		if m.SupportedVersions[i] != nil {
			if err := m.SupportedVersions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("supportedVersions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *SupportedVersionsV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SupportedVersionsV4Response) UnmarshalBinary(b []byte) error {
	var res SupportedVersionsV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
