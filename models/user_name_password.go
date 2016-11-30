package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*UserNamePassword user name password

swagger:model UserNamePassword
*/
type UserNamePassword struct {

	/* old password in ambari

	Required: true
	Max Length: 2147483647
	Min Length: 1
	*/
	OldPassword string `json:"oldPassword"`

	/* new password in ambari

	Required: true
	Max Length: 2147483647
	Min Length: 1
	*/
	Password string `json:"password"`

	/* new user name in ambari

	Required: true
	Max Length: 2147483647
	Min Length: 1
	*/
	UserName string `json:"userName"`
}

// Validate validates this user name password
func (m *UserNamePassword) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOldPassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePassword(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserNamePassword) validateOldPassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("oldPassword", "body", string(m.OldPassword)); err != nil {
		return err
	}

	if err := validate.MinLength("oldPassword", "body", string(m.OldPassword), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("oldPassword", "body", string(m.OldPassword), 2147483647); err != nil {
		return err
	}

	return nil
}

func (m *UserNamePassword) validatePassword(formats strfmt.Registry) error {

	if err := validate.RequiredString("password", "body", string(m.Password)); err != nil {
		return err
	}

	if err := validate.MinLength("password", "body", string(m.Password), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("password", "body", string(m.Password), 2147483647); err != nil {
		return err
	}

	return nil
}

func (m *UserNamePassword) validateUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("userName", "body", string(m.UserName)); err != nil {
		return err
	}

	if err := validate.MinLength("userName", "body", string(m.UserName), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("userName", "body", string(m.UserName), 2147483647); err != nil {
		return err
	}

	return nil
}
