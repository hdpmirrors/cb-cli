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

// AddDNSCnameRecordV1Request add Dns cname record v1 request
// swagger:model AddDnsCnameRecordV1Request
type AddDNSCnameRecordV1Request struct {

	// DNS name without the domain. eg. 'ipaserver' from 'ipaserver.cloudera.site'
	// Required: true
	// Pattern: ^(\*\.)?[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9]+$
	Cname *string `json:"cname"`

	// It's the domain. Eg if your FQDN is ipaserver.cloudera.site, it's 'cloudera.site'. '168.192.in-addr.arpa' for a reverse record like '5.1.168.192.in-addr.arpa'
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9\.]+$
	DNSZone string `json:"dnsZone,omitempty"`

	// CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// The fully qualified domain name of the host the CNAME should point to.
	// Required: true
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9]+$
	TargetFqdn *string `json:"targetFqdn"`
}

// Validate validates this add Dns cname record v1 request
func (m *AddDNSCnameRecordV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDNSZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTargetFqdn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddDNSCnameRecordV1Request) validateCname(formats strfmt.Registry) error {

	if err := validate.Required("cname", "body", m.Cname); err != nil {
		return err
	}

	if err := validate.Pattern("cname", "body", string(*m.Cname), `^(\*\.)?[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *AddDNSCnameRecordV1Request) validateDNSZone(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSZone) { // not required
		return nil
	}

	if err := validate.Pattern("dnsZone", "body", string(m.DNSZone), `^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9\.]+$`); err != nil {
		return err
	}

	return nil
}

func (m *AddDNSCnameRecordV1Request) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *AddDNSCnameRecordV1Request) validateTargetFqdn(formats strfmt.Registry) error {

	if err := validate.Required("targetFqdn", "body", m.TargetFqdn); err != nil {
		return err
	}

	if err := validate.Pattern("targetFqdn", "body", string(*m.TargetFqdn), `^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9]+$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddDNSCnameRecordV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddDNSCnameRecordV1Request) UnmarshalBinary(b []byte) error {
	var res AddDNSCnameRecordV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}