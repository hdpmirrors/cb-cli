// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// TagsV4Request tags v4 request
// swagger:model TagsV4Request
type TagsV4Request struct {

	// stack related application tags
	Application map[string]string `json:"application,omitempty"`

	// stack related default tags
	Defaults map[string]string `json:"defaults,omitempty"`

	// stack related userdefined tags
	UserDefined map[string]string `json:"userDefined,omitempty"`
}

// Validate validates this tags v4 request
func (m *TagsV4Request) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TagsV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TagsV4Request) UnmarshalBinary(b []byte) error {
	var res TagsV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}