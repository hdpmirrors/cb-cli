// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// RdsConfig rds config
// swagger:model RdsConfig

type RdsConfig struct {

	// Name of the JDBC connection driver (for example: 'org.postgresql.Driver')
	// Required: true
	ConnectionDriver *string `json:"connectionDriver"`

	// Password to use for the jdbc connection
	// Required: true
	ConnectionPassword *string `json:"connectionPassword"`

	// JDBC connection URL in the form of jdbc:<db-type>://<address>:<port>/<db>
	// Required: true
	// Pattern: ^jdbc:postgresql://[-\w\.]*:\d{1,5}/?\w*
	ConnectionURL *string `json:"connectionURL"`

	// Username to use for the jdbc connection
	// Required: true
	ConnectionUserName *string `json:"connectionUserName"`

	// Name of the external database engine (MYSQL, POSTGRES...)
	// Required: true
	DatabaseEngine *string `json:"databaseEngine"`

	// Name of the RDS configuration resource
	// Required: true
	// Max Length: 50
	// Min Length: 4
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Type of RDS, aka the service name that will use the RDS like HIVE, DRUID, SUPERSET, RANGER, etc.
	// Required: true
	// Max Length: 12
	// Min Length: 4
	// Pattern: (^[a-zA-Z][-a-zA-Z0-9]*[a-zA-Z0-9]$)
	Type *string `json:"type"`

	// If true, then the RDS configuration will be validated
	Validated *bool `json:"validated,omitempty"`
}

/* polymorph RdsConfig connectionDriver false */

/* polymorph RdsConfig connectionPassword false */

/* polymorph RdsConfig connectionURL false */

/* polymorph RdsConfig connectionUserName false */

/* polymorph RdsConfig databaseEngine false */

/* polymorph RdsConfig name false */

/* polymorph RdsConfig type false */

/* polymorph RdsConfig validated false */

// Validate validates this rds config
func (m *RdsConfig) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionDriver(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionPassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionURL(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConnectionUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDatabaseEngine(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RdsConfig) validateConnectionDriver(formats strfmt.Registry) error {

	if err := validate.Required("connectionDriver", "body", m.ConnectionDriver); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateConnectionPassword(formats strfmt.Registry) error {

	if err := validate.Required("connectionPassword", "body", m.ConnectionPassword); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateConnectionURL(formats strfmt.Registry) error {

	if err := validate.Required("connectionURL", "body", m.ConnectionURL); err != nil {
		return err
	}

	if err := validate.Pattern("connectionURL", "body", string(*m.ConnectionURL), `^jdbc:postgresql://[-\w\.]*:\d{1,5}/?\w*`); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateConnectionUserName(formats strfmt.Registry) error {

	if err := validate.Required("connectionUserName", "body", m.ConnectionUserName); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateDatabaseEngine(formats strfmt.Registry) error {

	if err := validate.Required("databaseEngine", "body", m.DatabaseEngine); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 50); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *RdsConfig) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	if err := validate.MinLength("type", "body", string(*m.Type), 4); err != nil {
		return err
	}

	if err := validate.MaxLength("type", "body", string(*m.Type), 12); err != nil {
		return err
	}

	if err := validate.Pattern("type", "body", string(*m.Type), `(^[a-zA-Z][-a-zA-Z0-9]*[a-zA-Z0-9]$)`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RdsConfig) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RdsConfig) UnmarshalBinary(b []byte) error {
	var res RdsConfig
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
