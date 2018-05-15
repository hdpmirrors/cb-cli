// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackResponse stack response
// swagger:model StackResponse

type StackResponse struct {

	// account id of the resource owner that is provided by OAuth provider
	Account string `json:"account,omitempty"`

	// specific version of ambari
	AmbariVersion string `json:"ambariVersion,omitempty"`

	// stack related application tags
	ApplicationTags map[string]string `json:"applicationTags,omitempty"`

	// availability zone of the stack
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// type of cloud provider
	// Read Only: true
	CloudPlatform string `json:"cloudPlatform,omitempty"`

	// details of the Cloudbreak that provisioned the stack
	CloudbreakDetails *CloudbreakDetailsJSON `json:"cloudbreakDetails,omitempty"`

	// related events for a cloudbreak stack
	CloudbreakEvents []*CloudbreakEvent `json:"cloudbreakEvents"`

	// usage information for a specific stack
	CloudbreakUsages []*CloudbreakUsage `json:"cloudbreakUsages"`

	// cluster
	Cluster *ClusterResponse `json:"cluster,omitempty"`

	// using the cluster name to create subdomain
	ClusterNameAsSubdomain *bool `json:"clusterNameAsSubdomain,omitempty"`

	// creation time of the stack in long
	Created int64 `json:"created,omitempty"`

	// stack related credential
	Credential *CredentialResponse `json:"credential,omitempty"`

	// credential resource id for the stack
	CredentialID int64 `json:"credentialId,omitempty"`

	// custom domain name for the nodes in the stack
	CustomDomain string `json:"customDomain,omitempty"`

	// custom hostname for nodes in the stack
	CustomHostname string `json:"customHostname,omitempty"`

	// Custom parameters as a json
	CustomInputs map[string]interface{} `json:"customInputs,omitempty"`

	// stack related default tags
	DefaultTags map[string]string `json:"defaultTags,omitempty"`

	// failure policy in case of failures
	FailurePolicy *FailurePolicyResponse `json:"failurePolicy,omitempty"`

	// the related flex subscription
	FlexSubscription *FlexSubscriptionResponse `json:"flexSubscription,omitempty"`

	// port of the gateway secured proxy
	GatewayPort int32 `json:"gatewayPort,omitempty"`

	// hardware information where pairing hostmetadata with instancemetadata
	// Unique: true
	HardwareInfos []*HardwareInfoResponse `json:"hardwareInfos"`

	// specific version of HDP
	HdpVersion string `json:"hdpVersion,omitempty"`

	// using the hostgroup names to create hostnames
	HostgroupNameAsHostname *bool `json:"hostgroupNameAsHostname,omitempty"`

	// id of the stack
	ID int64 `json:"id,omitempty"`

	// image of the stack
	Image *ImageJSON `json:"image,omitempty"`

	// instance groups
	InstanceGroups []*InstanceGroupResponse `json:"instanceGroups"`

	// name of the stack
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// stack related network
	Network *NetworkResponse `json:"network,omitempty"`

	// network resource id for the stack
	NetworkID int64 `json:"networkId,omitempty"`

	// node count of the stack
	NodeCount int32 `json:"nodeCount,omitempty"`

	// action on failure
	OnFailureAction string `json:"onFailureAction,omitempty"`

	// the details of the container orchestrator api to use
	Orchestrator *OrchestratorResponse `json:"orchestrator,omitempty"`

	// id of the resource owner that is provided by OAuth provider
	Owner string `json:"owner,omitempty"`

	// additional cloud specific parameters for stack
	Parameters map[string]string `json:"parameters,omitempty"`

	// cloud provider api variant
	PlatformVariant string `json:"platformVariant,omitempty"`

	// resource is visible in account
	Public *bool `json:"public,omitempty"`

	// region of the stack
	Region string `json:"region,omitempty"`

	// stack related authentication
	StackAuthentication *StackAuthenticationResponse `json:"stackAuthentication,omitempty"`

	// status of the stack
	Status string `json:"status,omitempty"`

	// status message of the stack
	StatusReason string `json:"statusReason,omitempty"`

	// stack related userdefined tags
	UserDefinedTags map[string]string `json:"userDefinedTags,omitempty"`
}

/* polymorph StackResponse account false */

/* polymorph StackResponse ambariVersion false */

/* polymorph StackResponse applicationTags false */

/* polymorph StackResponse availabilityZone false */

/* polymorph StackResponse cloudPlatform false */

/* polymorph StackResponse cloudbreakDetails false */

/* polymorph StackResponse cloudbreakEvents false */

/* polymorph StackResponse cloudbreakUsages false */

/* polymorph StackResponse cluster false */

/* polymorph StackResponse clusterNameAsSubdomain false */

/* polymorph StackResponse created false */

/* polymorph StackResponse credential false */

/* polymorph StackResponse credentialId false */

/* polymorph StackResponse customDomain false */

/* polymorph StackResponse customHostname false */

/* polymorph StackResponse customInputs false */

/* polymorph StackResponse defaultTags false */

/* polymorph StackResponse failurePolicy false */

/* polymorph StackResponse flexSubscription false */

/* polymorph StackResponse gatewayPort false */

/* polymorph StackResponse hardwareInfos false */

/* polymorph StackResponse hdpVersion false */

/* polymorph StackResponse hostgroupNameAsHostname false */

/* polymorph StackResponse id false */

/* polymorph StackResponse image false */

/* polymorph StackResponse instanceGroups false */

/* polymorph StackResponse name false */

/* polymorph StackResponse network false */

/* polymorph StackResponse networkId false */

/* polymorph StackResponse nodeCount false */

/* polymorph StackResponse onFailureAction false */

/* polymorph StackResponse orchestrator false */

/* polymorph StackResponse owner false */

/* polymorph StackResponse parameters false */

/* polymorph StackResponse platformVariant false */

/* polymorph StackResponse public false */

/* polymorph StackResponse region false */

/* polymorph StackResponse stackAuthentication false */

/* polymorph StackResponse status false */

/* polymorph StackResponse statusReason false */

/* polymorph StackResponse userDefinedTags false */

// Validate validates this stack response
func (m *StackResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCloudbreakDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCloudbreakEvents(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCloudbreakUsages(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCluster(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCredential(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFailurePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFlexSubscription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHardwareInfos(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateImage(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOnFailureAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrchestrator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStackAuthentication(formats); err != nil {
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

func (m *StackResponse) validateCloudbreakDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudbreakDetails) { // not required
		return nil
	}

	if m.CloudbreakDetails != nil {

		if err := m.CloudbreakDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloudbreakDetails")
			}
			return err
		}
	}

	return nil
}

func (m *StackResponse) validateCloudbreakEvents(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudbreakEvents) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudbreakEvents); i++ {

		if swag.IsZero(m.CloudbreakEvents[i]) { // not required
			continue
		}

		if m.CloudbreakEvents[i] != nil {

			if err := m.CloudbreakEvents[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudbreakEvents" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackResponse) validateCloudbreakUsages(formats strfmt.Registry) error {

	if swag.IsZero(m.CloudbreakUsages) { // not required
		return nil
	}

	for i := 0; i < len(m.CloudbreakUsages); i++ {

		if swag.IsZero(m.CloudbreakUsages[i]) { // not required
			continue
		}

		if m.CloudbreakUsages[i] != nil {

			if err := m.CloudbreakUsages[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("cloudbreakUsages" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackResponse) validateCluster(formats strfmt.Registry) error {

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

func (m *StackResponse) validateCredential(formats strfmt.Registry) error {

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

func (m *StackResponse) validateFailurePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.FailurePolicy) { // not required
		return nil
	}

	if m.FailurePolicy != nil {

		if err := m.FailurePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failurePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *StackResponse) validateFlexSubscription(formats strfmt.Registry) error {

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

func (m *StackResponse) validateHardwareInfos(formats strfmt.Registry) error {

	if swag.IsZero(m.HardwareInfos) { // not required
		return nil
	}

	if err := validate.UniqueItems("hardwareInfos", "body", m.HardwareInfos); err != nil {
		return err
	}

	for i := 0; i < len(m.HardwareInfos); i++ {

		if swag.IsZero(m.HardwareInfos[i]) { // not required
			continue
		}

		if m.HardwareInfos[i] != nil {

			if err := m.HardwareInfos[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hardwareInfos" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackResponse) validateImage(formats strfmt.Registry) error {

	if swag.IsZero(m.Image) { // not required
		return nil
	}

	if m.Image != nil {

		if err := m.Image.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("image")
			}
			return err
		}
	}

	return nil
}

func (m *StackResponse) validateInstanceGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.InstanceGroups) { // not required
		return nil
	}

	for i := 0; i < len(m.InstanceGroups); i++ {

		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {

			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackResponse) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *StackResponse) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {

		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

var stackResponseTypeOnFailureActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ROLLBACK","DO_NOTHING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackResponseTypeOnFailureActionPropEnum = append(stackResponseTypeOnFailureActionPropEnum, v)
	}
}

const (
	// StackResponseOnFailureActionROLLBACK captures enum value "ROLLBACK"
	StackResponseOnFailureActionROLLBACK string = "ROLLBACK"
	// StackResponseOnFailureActionDONOTHING captures enum value "DO_NOTHING"
	StackResponseOnFailureActionDONOTHING string = "DO_NOTHING"
)

// prop value enum
func (m *StackResponse) validateOnFailureActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackResponseTypeOnFailureActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackResponse) validateOnFailureAction(formats strfmt.Registry) error {

	if swag.IsZero(m.OnFailureAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnFailureActionEnum("onFailureAction", "body", m.OnFailureAction); err != nil {
		return err
	}

	return nil
}

func (m *StackResponse) validateOrchestrator(formats strfmt.Registry) error {

	if swag.IsZero(m.Orchestrator) { // not required
		return nil
	}

	if m.Orchestrator != nil {

		if err := m.Orchestrator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orchestrator")
			}
			return err
		}
	}

	return nil
}

func (m *StackResponse) validateStackAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.StackAuthentication) { // not required
		return nil
	}

	if m.StackAuthentication != nil {

		if err := m.StackAuthentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackAuthentication")
			}
			return err
		}
	}

	return nil
}

var stackResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackResponseTypeStatusPropEnum = append(stackResponseTypeStatusPropEnum, v)
	}
}

const (
	// StackResponseStatusREQUESTED captures enum value "REQUESTED"
	StackResponseStatusREQUESTED string = "REQUESTED"
	// StackResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	StackResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"
	// StackResponseStatusAVAILABLE captures enum value "AVAILABLE"
	StackResponseStatusAVAILABLE string = "AVAILABLE"
	// StackResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	StackResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"
	// StackResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	StackResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"
	// StackResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	StackResponseStatusUPDATEFAILED string = "UPDATE_FAILED"
	// StackResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	StackResponseStatusCREATEFAILED string = "CREATE_FAILED"
	// StackResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	StackResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"
	// StackResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	StackResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"
	// StackResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	StackResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"
	// StackResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	StackResponseStatusDELETEFAILED string = "DELETE_FAILED"
	// StackResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	StackResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"
	// StackResponseStatusSTOPPED captures enum value "STOPPED"
	StackResponseStatusSTOPPED string = "STOPPED"
	// StackResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	StackResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"
	// StackResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	StackResponseStatusSTARTREQUESTED string = "START_REQUESTED"
	// StackResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	StackResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"
	// StackResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	StackResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"
	// StackResponseStatusSTARTFAILED captures enum value "START_FAILED"
	StackResponseStatusSTARTFAILED string = "START_FAILED"
	// StackResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	StackResponseStatusSTOPFAILED string = "STOP_FAILED"
	// StackResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	StackResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"
)

// prop value enum
func (m *StackResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *StackResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackResponse) UnmarshalBinary(b []byte) error {
	var res StackResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
