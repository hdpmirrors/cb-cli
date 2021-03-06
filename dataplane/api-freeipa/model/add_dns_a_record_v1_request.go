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

// AddDNSARecordV1Request add Dns a record v1 request
// swagger:model AddDnsARecordV1Request
type AddDNSARecordV1Request struct {

	// Tries to create a reverse pointer for the record (PTR). Only if reverse zone already exists
	CreateReverse bool `json:"createReverse,omitempty"`

	// It's the domain. Eg if your FQDN is ipaserver.cloudera.site, it's 'cloudera.site'. '168.192.in-addr.arpa' for a reverse record like '5.1.168.192.in-addr.arpa'
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9\.]+$
	DNSZone string `json:"dnsZone,omitempty"`

	// CRN of the environment
	// Required: true
	EnvironmentCrn *string `json:"environmentCrn"`

	// Hostname name without the domain. eg. 'ipaserver' from 'ipaserver.cloudera.site'
	// Required: true
	// Pattern: ^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9]+$
	Hostname *string `json:"hostname"`

	// The IP address of the host the A record should point to. Only IPv4
	// Required: true
	// Pattern: ^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$
	IP *string `json:"ip"`
}

// Validate validates this add Dns a record v1 request
func (m *AddDNSARecordV1Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDNSZone(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEnvironmentCrn(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHostname(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AddDNSARecordV1Request) validateDNSZone(formats strfmt.Registry) error {

	if swag.IsZero(m.DNSZone) { // not required
		return nil
	}

	if err := validate.Pattern("dnsZone", "body", string(m.DNSZone), `^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9\.]+$`); err != nil {
		return err
	}

	return nil
}

func (m *AddDNSARecordV1Request) validateEnvironmentCrn(formats strfmt.Registry) error {

	if err := validate.Required("environmentCrn", "body", m.EnvironmentCrn); err != nil {
		return err
	}

	return nil
}

func (m *AddDNSARecordV1Request) validateHostname(formats strfmt.Registry) error {

	if err := validate.Required("hostname", "body", m.Hostname); err != nil {
		return err
	}

	if err := validate.Pattern("hostname", "body", string(*m.Hostname), `^[a-zA-Z0-9]+[a-zA-Z0-9-\.]*[a-zA-Z0-9]+$`); err != nil {
		return err
	}

	return nil
}

func (m *AddDNSARecordV1Request) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	if err := validate.Pattern("ip", "body", string(*m.IP), `^((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AddDNSARecordV1Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AddDNSARecordV1Request) UnmarshalBinary(b []byte) error {
	var res AddDNSARecordV1Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
