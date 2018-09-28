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

// StackViewResponse stack view response
// swagger:model StackViewResponse

type StackViewResponse struct {

	// type of cloud provider
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// cluster object on stack
	Cluster *ClusterViewResponse `json:"cluster,omitempty"`

	// creation time of the stack in long
	Created int64 `json:"created,omitempty"`

	// stack related credential
	Credential *CredentialResponse `json:"credential,omitempty"`

	// the related flex subscription
	FlexSubscription *FlexSubscriptionResponse `json:"flexSubscription,omitempty"`

	// specific version of HDP
	HdpVersion string `json:"hdpVersion,omitempty"`

	// id of the stack
	ID int64 `json:"id,omitempty"`

	// name of the stack
	// Required: true
	Name *string `json:"name"`

	// node count of the stack
	NodeCount int32 `json:"nodeCount,omitempty"`

	// id of the resource owner that is provided by OAuth provider
	Owner string `json:"owner,omitempty"`

	// additional cloud specific parameters for stack
	Parameters map[string]string `json:"parameters,omitempty"`

	// cloud provider api variant
	PlatformVariant string `json:"platformVariant,omitempty"`

	// status of the stack
	Status string `json:"status,omitempty"`
}

/* polymorph StackViewResponse cloudPlatform false */

/* polymorph StackViewResponse cluster false */

/* polymorph StackViewResponse created false */

/* polymorph StackViewResponse credential false */

/* polymorph StackViewResponse flexSubscription false */

/* polymorph StackViewResponse hdpVersion false */

/* polymorph StackViewResponse id false */

/* polymorph StackViewResponse name false */

/* polymorph StackViewResponse nodeCount false */

/* polymorph StackViewResponse owner false */

/* polymorph StackViewResponse parameters false */

/* polymorph StackViewResponse platformVariant false */

/* polymorph StackViewResponse status false */

// Validate validates this stack view response
func (m *StackViewResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlexSubscription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackViewResponse) validateCluster(formats strfmt.Registry) error {

	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {

		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *StackViewResponse) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {

		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

func (m *StackViewResponse) validateFlexSubscription(formats strfmt.Registry) error {

	if swag.IsZero(m.FlexSubscription) { // not required
		return nil
	}

	if m.FlexSubscription != nil {

		if err := m.FlexSubscription.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flexSubscription")
			}
			return err
		}
	}

	return nil
}

func (m *StackViewResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var stackViewResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackViewResponseTypeStatusPropEnum = append(stackViewResponseTypeStatusPropEnum, v)
	}
}

const (
	// StackViewResponseStatusREQUESTED captures enum value "REQUESTED"
	StackViewResponseStatusREQUESTED string = "REQUESTED"
	// StackViewResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackViewResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"
	// StackViewResponseStatusAVAILABLE captures enum value "AVAILABLE"
	StackViewResponseStatusAVAILABLE string = "AVAILABLE"
	// StackViewResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackViewResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"
	// StackViewResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackViewResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"
	// StackViewResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackViewResponseStatusUPDATEFAILED string = "UPDATE_FAILED"
	// StackViewResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackViewResponseStatusCREATEFAILED string = "CREATE_FAILED"
	// StackViewResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackViewResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"
	// StackViewResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackViewResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"
	// StackViewResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackViewResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"
	// StackViewResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackViewResponseStatusDELETEFAILED string = "DELETE_FAILED"
	// StackViewResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackViewResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"
	// StackViewResponseStatusSTOPPED captures enum value "STOPPED"
	StackViewResponseStatusSTOPPED string = "STOPPED"
	// StackViewResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackViewResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"
	// StackViewResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackViewResponseStatusSTARTREQUESTED string = "START_REQUESTED"
	// StackViewResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackViewResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"
	// StackViewResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackViewResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"
	// StackViewResponseStatusSTARTFAILED captures enum value "START_FAILED"
	StackViewResponseStatusSTARTFAILED string = "START_FAILED"
	// StackViewResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackViewResponseStatusSTOPFAILED string = "STOP_FAILED"
	// StackViewResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackViewResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"
)

// prop value enum
func (m *StackViewResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackViewResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackViewResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackViewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackViewResponse) UnmarshalBinary(b []byte) error {
	var res StackViewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
