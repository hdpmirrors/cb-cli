// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// BaseImageV4Response base image v4 response
// swagger:model BaseImageV4Response
type BaseImageV4Response struct {

	// ambari repo
	AmbariRepo *AmbariRepositoryV4Response `json:"ambariRepo,omitempty"`

	// cdh stacks
	CdhStacks []*ClouderaManagerStackDetailsV4Response `json:"cdhStacks"`

	// cloudera manager repo
	ClouderaManagerRepo *ClouderaManagerRepositoryV4Response `json:"clouderaManagerRepo,omitempty"`

	// cm build number
	CmBuildNumber string `json:"cmBuildNumber,omitempty"`

	// created
	Created int64 `json:"created,omitempty"`

	// date
	Date string `json:"date,omitempty"`

	// default image
	DefaultImage bool `json:"defaultImage,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// hdf stacks
	HdfStacks []*AmbariStackDetailsV4Response `json:"hdfStacks"`

	// hdp stacks
	HdpStacks []*AmbariStackDetailsV4Response `json:"hdpStacks"`

	// images
	Images map[string]map[string]string `json:"images,omitempty"`

	// os
	Os string `json:"os,omitempty"`

	// os type
	OsType string `json:"osType,omitempty"`

	// package versions
	PackageVersions map[string]string `json:"packageVersions,omitempty"`

	// repository
	Repository map[string]string `json:"repository,omitempty"`

	// stack details
	StackDetails *BaseStackDetailsV4Response `json:"stackDetails,omitempty"`

	// uuid
	UUID string `json:"uuid,omitempty"`

	// version
	Version string `json:"version,omitempty"`
}

// Validate validates this base image v4 response
func (m *BaseImageV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbariRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCdhStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateClouderaManagerRepo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdfStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHdpStacks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStackDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BaseImageV4Response) validateAmbariRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariRepo) { // not required
		return nil
	}

	if m.AmbariRepo != nil {
		if err := m.AmbariRepo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariRepo")
			}
			return err
		}
	}

	return nil
}

func (m *BaseImageV4Response) validateCdhStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.CdhStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.CdhStacks); i++ {
		if swag.IsZero(m.CdhStacks[i]) { // not required
			continue
		}

		if m.CdhStacks[i] != nil {
			if err := m.CdhStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cdhStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseImageV4Response) validateClouderaManagerRepo(formats strfmt.Registry) error {

	if swag.IsZero(m.ClouderaManagerRepo) { // not required
		return nil
	}

	if m.ClouderaManagerRepo != nil {
		if err := m.ClouderaManagerRepo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clouderaManagerRepo")
			}
			return err
		}
	}

	return nil
}

func (m *BaseImageV4Response) validateHdfStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.HdfStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.HdfStacks); i++ {
		if swag.IsZero(m.HdfStacks[i]) { // not required
			continue
		}

		if m.HdfStacks[i] != nil {
			if err := m.HdfStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hdfStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseImageV4Response) validateHdpStacks(formats strfmt.Registry) error {

	if swag.IsZero(m.HdpStacks) { // not required
		return nil
	}

	for i := 0; i < len(m.HdpStacks); i++ {
		if swag.IsZero(m.HdpStacks[i]) { // not required
			continue
		}

		if m.HdpStacks[i] != nil {
			if err := m.HdpStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hdpStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *BaseImageV4Response) validateStackDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.StackDetails) { // not required
		return nil
	}

	if m.StackDetails != nil {
		if err := m.StackDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackDetails")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BaseImageV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BaseImageV4Response) UnmarshalBinary(b []byte) error {
	var res BaseImageV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
