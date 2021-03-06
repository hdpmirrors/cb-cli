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

// PlatformSecurityGroupsResponse platform security groups response
// swagger:model PlatformSecurityGroupsResponse
type PlatformSecurityGroupsResponse struct {

	// security groups
	SecurityGroups map[string][]PlatformSecurityGroupResponse `json:"securityGroups,omitempty"`
}

// Validate validates this platform security groups response
func (m *PlatformSecurityGroupsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSecurityGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlatformSecurityGroupsResponse) validateSecurityGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityGroups) { // not required
		return nil
	}

	for k := range m.SecurityGroups {

		if err := validate.Required("securityGroups"+"."+k, "body", m.SecurityGroups[k]); err != nil {
			return err
		}

		if err := validate.UniqueItems("securityGroups"+"."+k, "body", m.SecurityGroups[k]); err != nil {
			return err
		}

		for i := 0; i < len(m.SecurityGroups[k]); i++ {

			if err := m.SecurityGroups[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("securityGroups" + "." + k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlatformSecurityGroupsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlatformSecurityGroupsResponse) UnmarshalBinary(b []byte) error {
	var res PlatformSecurityGroupsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
