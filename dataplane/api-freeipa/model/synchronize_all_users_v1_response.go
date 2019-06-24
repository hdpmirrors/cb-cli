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

// SynchronizeAllUsersV1Response synchronize all users v1 response
// swagger:model SynchronizeAllUsersV1Response
type SynchronizeAllUsersV1Response struct {

	// User synchronization operation end time
	EndTime string `json:"endTime,omitempty"`

	// User synchronization operation id
	// Required: true
	ID *string `json:"id"`

	// User synchronization operation start time
	StartTime string `json:"startTime,omitempty"`

	// User synchronization operation status
	// Required: true
	// Enum: [RUNNING COMPLETED FAILED]
	Status *string `json:"status"`
}

// Validate validates this synchronize all users v1 response
func (m *SynchronizeAllUsersV1Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SynchronizeAllUsersV1Response) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

var synchronizeAllUsersV1ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["RUNNING","COMPLETED","FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		synchronizeAllUsersV1ResponseTypeStatusPropEnum = append(synchronizeAllUsersV1ResponseTypeStatusPropEnum, v)
	}
}

const (

	// SynchronizeAllUsersV1ResponseStatusRUNNING captures enum value "RUNNING"
	SynchronizeAllUsersV1ResponseStatusRUNNING string = "RUNNING"

	// SynchronizeAllUsersV1ResponseStatusCOMPLETED captures enum value "COMPLETED"
	SynchronizeAllUsersV1ResponseStatusCOMPLETED string = "COMPLETED"

	// SynchronizeAllUsersV1ResponseStatusFAILED captures enum value "FAILED"
	SynchronizeAllUsersV1ResponseStatusFAILED string = "FAILED"
)

// prop value enum
func (m *SynchronizeAllUsersV1Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, synchronizeAllUsersV1ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *SynchronizeAllUsersV1Response) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SynchronizeAllUsersV1Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SynchronizeAllUsersV1Response) UnmarshalBinary(b []byte) error {
	var res SynchronizeAllUsersV1Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}