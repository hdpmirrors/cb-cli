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

// EnvironmentNetworkV1Response environment network v1 response
// swagger:model EnvironmentNetworkV1Response
type EnvironmentNetworkV1Response struct {

	// Subnet ids of the specified networks
	Aws *EnvironmentNetworkAwsV1Params `json:"aws,omitempty"`

	// Subnet ids of the specified networks
	Azure *EnvironmentNetworkAzureV1Params `json:"azure,omitempty"`

	// id of the resource
	Crn string `json:"crn,omitempty"`

	// name of the resource
	// Required: true
	Name *string `json:"name"`

	// network cidr
	// Max Length: 255
	// Min Length: 0
	NetworkCidr *string `json:"networkCidr,omitempty"`

	// A flag to enable or disable the private sutbet creation.
	// Enum: [ENABLED DISABLED]
	PrivateSubnetCreation string `json:"privateSubnetCreation,omitempty"`

	// Subnet ids of the specified networks
	// Unique: true
	SubnetIds []string `json:"subnetIds"`

	// Subnet metadata of the specified networks
	SubnetMetas map[string]CloudSubnet `json:"subnetMetas,omitempty"`

	// Yarn parameters
	Yarn *EnvironmentNetworkYarnV1Params `json:"yarn,omitempty"`
}

// Validate validates this environment network v1 response
func (m *EnvironmentNetworkV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNetworkCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrivateSubnetCreation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubnetMetas(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateYarn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentNetworkV1Response) validateAws(formats strfmt.Registry) error {

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

func (m *EnvironmentNetworkV1Response) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentNetworkV1Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentNetworkV1Response) validateNetworkCidr(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkCidr) { // not required
		return nil
	}

	if err := validate.MinLength("networkCidr", "body", string(*m.NetworkCidr), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("networkCidr", "body", string(*m.NetworkCidr), 255); err != nil {
		return err
	}

	return nil
}

var environmentNetworkV1ResponseTypePrivateSubnetCreationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		environmentNetworkV1ResponseTypePrivateSubnetCreationPropEnum = append(environmentNetworkV1ResponseTypePrivateSubnetCreationPropEnum, v)
	}
}

const (

	// EnvironmentNetworkV1ResponsePrivateSubnetCreationENABLED captures enum value "ENABLED"
	EnvironmentNetworkV1ResponsePrivateSubnetCreationENABLED string = "ENABLED"

	// EnvironmentNetworkV1ResponsePrivateSubnetCreationDISABLED captures enum value "DISABLED"
	EnvironmentNetworkV1ResponsePrivateSubnetCreationDISABLED string = "DISABLED"
)

// prop value enum
func (m *EnvironmentNetworkV1Response) validatePrivateSubnetCreationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, environmentNetworkV1ResponseTypePrivateSubnetCreationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EnvironmentNetworkV1Response) validatePrivateSubnetCreation(formats strfmt.Registry) error {

	if swag.IsZero(m.PrivateSubnetCreation) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrivateSubnetCreationEnum("privateSubnetCreation", "body", m.PrivateSubnetCreation); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentNetworkV1Response) validateSubnetIds(formats strfmt.Registry) error {

	if swag.IsZero(m.SubnetIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("subnetIds", "body", m.SubnetIds); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentNetworkV1Response) validateSubnetMetas(formats strfmt.Registry) error {

	if swag.IsZero(m.SubnetMetas) { // not required
		return nil
	}

	for k := range m.SubnetMetas {

		if err := validate.Required("subnetMetas"+"."+k, "body", m.SubnetMetas[k]); err != nil {
			return err
		}
		if val, ok := m.SubnetMetas[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *EnvironmentNetworkV1Response) validateYarn(formats strfmt.Registry) error {

	if swag.IsZero(m.Yarn) { // not required
		return nil
	}

	if m.Yarn != nil {
		if err := m.Yarn.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("yarn")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentNetworkV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentNetworkV1Response) UnmarshalBinary(b []byte) error {
	var res EnvironmentNetworkV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
