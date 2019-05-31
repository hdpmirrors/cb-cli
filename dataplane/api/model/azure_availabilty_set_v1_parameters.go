// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// AzureAvailabiltySetV1Parameters azure availabilty set v1 parameters
// swagger:model AzureAvailabiltySetV1Parameters
type AzureAvailabiltySetV1Parameters struct {

	// fault domain count
	FaultDomainCount int32 `json:"faultDomainCount,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// update domain count
	UpdateDomainCount int32 `json:"updateDomainCount,omitempty"`
}

// Validate validates this azure availabilty set v1 parameters
func (m *AzureAvailabiltySetV1Parameters) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AzureAvailabiltySetV1Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureAvailabiltySetV1Parameters) UnmarshalBinary(b []byte) error {
	var res AzureAvailabiltySetV1Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}