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

// AddDNSZoneForSubnetsV1Response add Dns zone for subnets v1 response
// swagger:model AddDnsZoneForSubnetsV1Response
type AddDNSZoneForSubnetsV1Response struct {

	// failed
	Failed map[string]string `json:"failed,omitempty"`

	// success
	// Unique: true
	Success []string `json:"success"`
}

// Validate validates this add Dns zone for subnets v1 response
func (m *AddDNSZoneForSubnetsV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSuccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddDNSZoneForSubnetsV1Response) validateSuccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Success) { // not required
		return nil
	}

	if err := validate.UniqueItems("success", "body", m.Success); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddDNSZoneForSubnetsV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddDNSZoneForSubnetsV1Response) UnmarshalBinary(b []byte) error {
	var res AddDNSZoneForSubnetsV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
