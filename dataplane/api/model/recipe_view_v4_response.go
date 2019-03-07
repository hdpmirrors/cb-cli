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

// RecipeViewV4Response recipe view v4 response
// swagger:model RecipeViewV4Response
type RecipeViewV4Response struct {

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// type of recipe [PRE_AMBARI_START,PRE_TERMINATION,POST_AMBARI_START,POST_CLUSTER_INSTALL]. The default is PRE_AMBARI_START
	// Required: true
	// Enum: [PRE_AMBARI_START PRE_TERMINATION POST_AMBARI_START POST_CLUSTER_INSTALL]
	Type *string `json:"type"`
}

// Validate validates this recipe view v4 response
func (m *RecipeViewV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *RecipeViewV4Response) validateDescription(formats strfmt.Registry) error {

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

func (m *RecipeViewV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

var recipeViewV4ResponseTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PRE_AMBARI_START","PRE_TERMINATION","POST_AMBARI_START","POST_CLUSTER_INSTALL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		recipeViewV4ResponseTypeTypePropEnum = append(recipeViewV4ResponseTypeTypePropEnum, v)
	}
}

const (

	// RecipeViewV4ResponseTypePREAMBARISTART captures enum value "PRE_AMBARI_START"
	RecipeViewV4ResponseTypePREAMBARISTART string = "PRE_AMBARI_START"

	// RecipeViewV4ResponseTypePRETERMINATION captures enum value "PRE_TERMINATION"
	RecipeViewV4ResponseTypePRETERMINATION string = "PRE_TERMINATION"

	// RecipeViewV4ResponseTypePOSTAMBARISTART captures enum value "POST_AMBARI_START"
	RecipeViewV4ResponseTypePOSTAMBARISTART string = "POST_AMBARI_START"

	// RecipeViewV4ResponseTypePOSTCLUSTERINSTALL captures enum value "POST_CLUSTER_INSTALL"
	RecipeViewV4ResponseTypePOSTCLUSTERINSTALL string = "POST_CLUSTER_INSTALL"
)

// prop value enum
func (m *RecipeViewV4Response) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, recipeViewV4ResponseTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *RecipeViewV4Response) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *RecipeViewV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *RecipeViewV4Response) UnmarshalBinary(b []byte) error {
	var res RecipeViewV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}