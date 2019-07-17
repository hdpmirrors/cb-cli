// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EnvironmentV1Request environment v1 request
// swagger:model EnvironmentV1Request
type EnvironmentV1Request struct {

	// SSH key for accessing cluster node instances.
	Authentication *EnvironmentAuthenticationV1Request `json:"authentication,omitempty"`

	// Cloud platform of the environment.
	// Max Length: 100
	// Min Length: 0
	CloudPlatform *string `json:"cloudPlatform,omitempty"`

	// Name of the credential of the environment. If the name is given, the detailed credential is ignored in the request.
	CredentialName string `json:"credentialName,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// Properties for FreeIpa which can be attached to the given environment
	FreeIpa *AttachedFreeIpaRequest `json:"freeIpa,omitempty"`

	// Location of the environment.
	// Required: true
	Location *LocationV1Request `json:"location"`

	// name of the resource
	// Required: true
	// Max Length: 255
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Network related specifics of the environment.
	Network *EnvironmentNetworkV1Request `json:"network,omitempty"`

	// Regions of the environment.
	// Unique: true
	Regions []string `json:"regions"`

	// Security control for FreeIPA and Datalake deployment.
	SecurityAccess *SecurityAccessV1Request `json:"securityAccess,omitempty"`

	// Telemetry related specifics of the environment.
	Telemetry *TelemetryRequest `json:"telemetry,omitempty"`

	// Configuration that the connection going directly or with ccm.
	// Enum: [DIRECT CCM]
	Tunnel string `json:"tunnel,omitempty"`
}

// Validate validates this environment v1 request
func (m *EnvironmentV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCloudPlatform(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFreeIpa(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecurityAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTelemetry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTunnel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentV1Request) validateAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.Authentication) { // not required
		return nil
	}

	if m.Authentication != nil {
		if err := m.Authentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("authentication")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateCloudPlatform(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudPlatform) { // not required
		return nil
	}

	if err := validate.MinLength("cloudPlatform", "body", string(*m.CloudPlatform), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("cloudPlatform", "body", string(*m.CloudPlatform), 100); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateFreeIpa(formats strfmt.Registry) error {

	if swag.IsZero(m.FreeIpa) { // not required
		return nil
	}

	if m.FreeIpa != nil {
		if err := m.FreeIpa.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("freeIpa")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 255); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {
		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV1Request) validateSecurityAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.SecurityAccess) { // not required
		return nil
	}

	if m.SecurityAccess != nil {
		if err := m.SecurityAccess.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securityAccess")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV1Request) validateTelemetry(formats strfmt.Registry) error {

	if swag.IsZero(m.Telemetry) { // not required
		return nil
	}

	if m.Telemetry != nil {
		if err := m.Telemetry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("telemetry")
			}
			return err
		}
	}

	return nil
}

var environmentV1RequestTypeTunnelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DIRECT","CCM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentV1RequestTypeTunnelPropEnum = append(environmentV1RequestTypeTunnelPropEnum, v)
	}
}

const (

	// EnvironmentV1RequestTunnelDIRECT captures enum value "DIRECT"
	EnvironmentV1RequestTunnelDIRECT string = "DIRECT"

	// EnvironmentV1RequestTunnelCCM captures enum value "CCM"
	EnvironmentV1RequestTunnelCCM string = "CCM"
)

// prop value enum
func (m *EnvironmentV1Request) validateTunnelEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentV1RequestTypeTunnelPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentV1Request) validateTunnel(formats strfmt.Registry) error {

	if swag.IsZero(m.Tunnel) { // not required
		return nil
	}

	// value enum
	if err := m.validateTunnelEnum("tunnel", "body", m.Tunnel); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentV1Request) UnmarshalBinary(b []byte) error {
	var res EnvironmentV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
