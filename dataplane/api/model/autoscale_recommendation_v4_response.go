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

// AutoscaleRecommendationV4Response autoscale recommendation v4 response
// swagger:model AutoscaleRecommendationV4Response
type AutoscaleRecommendationV4Response struct {

	// load based host groups
	// Unique: true
	LoadBasedHostGroups []string `json:"loadBasedHostGroups"`

	// time based host groups
	// Unique: true
	TimeBasedHostGroups []string `json:"timeBasedHostGroups"`
}

// Validate validates this autoscale recommendation v4 response
func (m *AutoscaleRecommendationV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLoadBasedHostGroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeBasedHostGroups(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AutoscaleRecommendationV4Response) validateLoadBasedHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.LoadBasedHostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("loadBasedHostGroups", "body", m.LoadBasedHostGroups); err != nil {
		return err
	}

	return nil
}

func (m *AutoscaleRecommendationV4Response) validateTimeBasedHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeBasedHostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("timeBasedHostGroups", "body", m.TimeBasedHostGroups); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AutoscaleRecommendationV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AutoscaleRecommendationV4Response) UnmarshalBinary(b []byte) error {
	var res AutoscaleRecommendationV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}