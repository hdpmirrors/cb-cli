// Code generated by go-swagger; DO NOT EDIT.

package model

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

// ClusterResponse cluster response
// swagger:model ClusterResponse

type ClusterResponse struct {

	// [DEPRECATED] use RdsConfig instead! details of the external Ambari database
	AmbariDatabaseDetails *AmbariDatabaseDetails `json:"ambariDatabaseDetails,omitempty"`

	// details of the Ambari package repository
	AmbariRepoDetailsJSON *AmbariRepoDetails `json:"ambariRepoDetailsJson,omitempty"`

	// public ambari ip of the stack
	AmbariServerIP string `json:"ambariServerIp,omitempty"`

	// public ambari url
	AmbariServerURL string `json:"ambariServerUrl,omitempty"`

	// details of the Ambari stack
	AmbariStackDetails *AmbariStackDetailsResponse `json:"ambariStackDetails,omitempty"`

	// Additional information for ambari cluster
	Attributes map[string]interface{} `json:"attributes,omitempty"`

	// blueprint for the cluster
	Blueprint *BlueprintResponse `json:"blueprint,omitempty"`

	// blueprint custom properties
	BlueprintCustomProperties string `json:"blueprintCustomProperties,omitempty"`

	// blueprint id for the cluster
	BlueprintID int64 `json:"blueprintId,omitempty"`

	// blueprint inputs in the cluster
	// Unique: true
	BlueprintInputs []*BlueprintInput `json:"blueprintInputs"`

	// name of the cluster
	Cluster string `json:"cluster,omitempty"`

	// cluster exposed services for topologies
	ClusterExposedServicesForTopologies map[string][]ClusterExposedServiceResponse `json:"clusterExposedServicesForTopologies,omitempty"`

	// config recommendation strategy
	ConfigStrategy string `json:"configStrategy,omitempty"`

	// Epoch time of cluster creation finish
	CreationFinished int64 `json:"creationFinished,omitempty"`

	// custom containers
	CustomContainers *CustomContainerResponse `json:"customContainers,omitempty"`

	// custom queue for yarn orchestrator
	CustomQueue string `json:"customQueue,omitempty"`

	// description of the resource
	Description string `json:"description,omitempty"`

	// executor type of cluster
	ExecutorType string `json:"executorType,omitempty"`

	// ambari blueprint JSON, set this or the url field
	ExtendedBlueprintText string `json:"extendedBlueprintText,omitempty"`

	// filesystem for a specific stack
	FileSystemResponse *FileSystemResponse `json:"fileSystemResponse,omitempty"`

	// gateway
	Gateway *GatewayJSON `json:"gateway,omitempty"`

	// collection of hostgroups
	// Unique: true
	HostGroups []*HostGroupResponse `json:"hostGroups"`

	// duration - how long the cluster is running in hours
	HoursUp int32 `json:"hoursUp,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// kerberos response
	KerberosResponse *KerberosResponse `json:"kerberosResponse,omitempty"`

	// LDAP config for the cluster
	LdapConfig *LdapConfigResponse `json:"ldapConfig,omitempty"`

	// LDAP config id for the cluster
	LdapConfigID int64 `json:"ldapConfigId,omitempty"`

	// duration - how long the cluster is running in minutes (minus hours)
	MinutesUp int32 `json:"minutesUp,omitempty"`

	// name of the resource
	Name string `json:"name,omitempty"`

	// proxy configuration name for the cluster
	ProxyName string `json:"proxyName,omitempty"`

	// RDS configuration names for the cluster
	// Unique: true
	RdsConfigIds []int64 `json:"rdsConfigIds"`

	// RDS configurations for the cluster
	// Unique: true
	RdsConfigs []*RDSConfigResponse `json:"rdsConfigs"`

	// tells wether the cluster is secured or not
	Secure *bool `json:"secure,omitempty"`

	// shared service for a specific stack
	SharedServiceResponse *SharedServiceResponse `json:"sharedServiceResponse,omitempty"`

	// status of the cluster
	Status string `json:"status,omitempty"`

	// status message of the cluster
	StatusReason string `json:"statusReason,omitempty"`

	// duration - how long the cluster is running in milliseconds
	Uptime int64 `json:"uptime,omitempty"`

	// ambari username
	UserName string `json:"userName,omitempty"`

	// workspace of the resource
	Workspace *WorkspaceResourceResponse `json:"workspace,omitempty"`
}

/* polymorph ClusterResponse ambariDatabaseDetails false */

/* polymorph ClusterResponse ambariRepoDetailsJson false */

/* polymorph ClusterResponse ambariServerIp false */

/* polymorph ClusterResponse ambariServerUrl false */

/* polymorph ClusterResponse ambariStackDetails false */

/* polymorph ClusterResponse attributes false */

/* polymorph ClusterResponse blueprint false */

/* polymorph ClusterResponse blueprintCustomProperties false */

/* polymorph ClusterResponse blueprintId false */

/* polymorph ClusterResponse blueprintInputs false */

/* polymorph ClusterResponse cluster false */

/* polymorph ClusterResponse clusterExposedServicesForTopologies false */

/* polymorph ClusterResponse configStrategy false */

/* polymorph ClusterResponse creationFinished false */

/* polymorph ClusterResponse customContainers false */

/* polymorph ClusterResponse customQueue false */

/* polymorph ClusterResponse description false */

/* polymorph ClusterResponse executorType false */

/* polymorph ClusterResponse extendedBlueprintText false */

/* polymorph ClusterResponse fileSystemResponse false */

/* polymorph ClusterResponse gateway false */

/* polymorph ClusterResponse hostGroups false */

/* polymorph ClusterResponse hoursUp false */

/* polymorph ClusterResponse id false */

/* polymorph ClusterResponse kerberosResponse false */

/* polymorph ClusterResponse ldapConfig false */

/* polymorph ClusterResponse ldapConfigId false */

/* polymorph ClusterResponse minutesUp false */

/* polymorph ClusterResponse name false */

/* polymorph ClusterResponse proxyName false */

/* polymorph ClusterResponse rdsConfigIds false */

/* polymorph ClusterResponse rdsConfigs false */

/* polymorph ClusterResponse secure false */

/* polymorph ClusterResponse sharedServiceResponse false */

/* polymorph ClusterResponse status false */

/* polymorph ClusterResponse statusReason false */

/* polymorph ClusterResponse uptime false */

/* polymorph ClusterResponse userName false */

/* polymorph ClusterResponse workspace false */

// Validate validates this cluster response
func (m *ClusterResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAmbariDatabaseDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAmbariRepoDetailsJSON(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateAmbariStackDetails(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBlueprint(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateBlueprintInputs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateClusterExposedServicesForTopologies(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateConfigStrategy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCustomContainers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateExecutorType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFileSystemResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateGateway(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateHostGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateKerberosResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateLdapConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigIds(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateRdsConfigs(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSharedServiceResponse(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateWorkspace(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterResponse) validateAmbariDatabaseDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariDatabaseDetails) { // not required
		return nil
	}

	if m.AmbariDatabaseDetails != nil {

		if err := m.AmbariDatabaseDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariDatabaseDetails")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateAmbariRepoDetailsJSON(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariRepoDetailsJSON) { // not required
		return nil
	}

	if m.AmbariRepoDetailsJSON != nil {

		if err := m.AmbariRepoDetailsJSON.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariRepoDetailsJson")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateAmbariStackDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.AmbariStackDetails) { // not required
		return nil
	}

	if m.AmbariStackDetails != nil {

		if err := m.AmbariStackDetails.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ambariStackDetails")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateBlueprint(formats strfmt.Registry) error {

	if swag.IsZero(m.Blueprint) { // not required
		return nil
	}

	if m.Blueprint != nil {

		if err := m.Blueprint.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("blueprint")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateBlueprintInputs(formats strfmt.Registry) error {

	if swag.IsZero(m.BlueprintInputs) { // not required
		return nil
	}

	if err := validate.UniqueItems("blueprintInputs", "body", m.BlueprintInputs); err != nil {
		return err
	}

	for i := 0; i < len(m.BlueprintInputs); i++ {

		if swag.IsZero(m.BlueprintInputs[i]) { // not required
			continue
		}

		if m.BlueprintInputs[i] != nil {

			if err := m.BlueprintInputs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("blueprintInputs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterResponse) validateClusterExposedServicesForTopologies(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterExposedServicesForTopologies) { // not required
		return nil
	}

	if err := validate.Required("clusterExposedServicesForTopologies", "body", m.ClusterExposedServicesForTopologies); err != nil {
		return err
	}

	return nil
}

var clusterResponseTypeConfigStrategyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NEVER_APPLY","ONLY_STACK_DEFAULTS_APPLY","ALWAYS_APPLY","ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterResponseTypeConfigStrategyPropEnum = append(clusterResponseTypeConfigStrategyPropEnum, v)
	}
}

const (
	// ClusterResponseConfigStrategyNEVERAPPLY captures enum value "NEVER_APPLY"
	ClusterResponseConfigStrategyNEVERAPPLY string = "NEVER_APPLY"
	// ClusterResponseConfigStrategyONLYSTACKDEFAULTSAPPLY captures enum value "ONLY_STACK_DEFAULTS_APPLY"
	ClusterResponseConfigStrategyONLYSTACKDEFAULTSAPPLY string = "ONLY_STACK_DEFAULTS_APPLY"
	// ClusterResponseConfigStrategyALWAYSAPPLY captures enum value "ALWAYS_APPLY"
	ClusterResponseConfigStrategyALWAYSAPPLY string = "ALWAYS_APPLY"
	// ClusterResponseConfigStrategyALWAYSAPPLYDONTOVERRIDECUSTOMVALUES captures enum value "ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"
	ClusterResponseConfigStrategyALWAYSAPPLYDONTOVERRIDECUSTOMVALUES string = "ALWAYS_APPLY_DONT_OVERRIDE_CUSTOM_VALUES"
)

// prop value enum
func (m *ClusterResponse) validateConfigStrategyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterResponseTypeConfigStrategyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterResponse) validateConfigStrategy(formats strfmt.Registry) error {

	if swag.IsZero(m.ConfigStrategy) { // not required
		return nil
	}

	// value enum
	if err := m.validateConfigStrategyEnum("configStrategy", "body", m.ConfigStrategy); err != nil {
		return err
	}

	return nil
}

func (m *ClusterResponse) validateCustomContainers(formats strfmt.Registry) error {

	if swag.IsZero(m.CustomContainers) { // not required
		return nil
	}

	if m.CustomContainers != nil {

		if err := m.CustomContainers.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("customContainers")
			}
			return err
		}
	}

	return nil
}

var clusterResponseTypeExecutorTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CONTAINER","DEFAULT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterResponseTypeExecutorTypePropEnum = append(clusterResponseTypeExecutorTypePropEnum, v)
	}
}

const (
	// ClusterResponseExecutorTypeCONTAINER captures enum value "CONTAINER"
	ClusterResponseExecutorTypeCONTAINER string = "CONTAINER"
	// ClusterResponseExecutorTypeDEFAULT captures enum value "DEFAULT"
	ClusterResponseExecutorTypeDEFAULT string = "DEFAULT"
)

// prop value enum
func (m *ClusterResponse) validateExecutorTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterResponseTypeExecutorTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterResponse) validateExecutorType(formats strfmt.Registry) error {

	if swag.IsZero(m.ExecutorType) { // not required
		return nil
	}

	// value enum
	if err := m.validateExecutorTypeEnum("executorType", "body", m.ExecutorType); err != nil {
		return err
	}

	return nil
}

func (m *ClusterResponse) validateFileSystemResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.FileSystemResponse) { // not required
		return nil
	}

	if m.FileSystemResponse != nil {

		if err := m.FileSystemResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fileSystemResponse")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateGateway(formats strfmt.Registry) error {

	if swag.IsZero(m.Gateway) { // not required
		return nil
	}

	if m.Gateway != nil {

		if err := m.Gateway.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gateway")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateHostGroups(formats strfmt.Registry) error {

	if swag.IsZero(m.HostGroups) { // not required
		return nil
	}

	if err := validate.UniqueItems("hostGroups", "body", m.HostGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.HostGroups); i++ {

		if swag.IsZero(m.HostGroups[i]) { // not required
			continue
		}

		if m.HostGroups[i] != nil {

			if err := m.HostGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("hostGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterResponse) validateKerberosResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.KerberosResponse) { // not required
		return nil
	}

	if m.KerberosResponse != nil {

		if err := m.KerberosResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kerberosResponse")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateLdapConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.LdapConfig) { // not required
		return nil
	}

	if m.LdapConfig != nil {

		if err := m.LdapConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ldapConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ClusterResponse) validateRdsConfigIds(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsConfigIds) { // not required
		return nil
	}

	if err := validate.UniqueItems("rdsConfigIds", "body", m.RdsConfigIds); err != nil {
		return err
	}

	return nil
}

func (m *ClusterResponse) validateRdsConfigs(formats strfmt.Registry) error {

	if swag.IsZero(m.RdsConfigs) { // not required
		return nil
	}

	if err := validate.UniqueItems("rdsConfigs", "body", m.RdsConfigs); err != nil {
		return err
	}

	for i := 0; i < len(m.RdsConfigs); i++ {

		if swag.IsZero(m.RdsConfigs[i]) { // not required
			continue
		}

		if m.RdsConfigs[i] != nil {

			if err := m.RdsConfigs[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("rdsConfigs" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ClusterResponse) validateSharedServiceResponse(formats strfmt.Registry) error {

	if swag.IsZero(m.SharedServiceResponse) { // not required
		return nil
	}

	if m.SharedServiceResponse != nil {

		if err := m.SharedServiceResponse.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sharedServiceResponse")
			}
			return err
		}
	}

	return nil
}

var clusterResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["REQUESTED","CREATE_IN_PROGRESS","AVAILABLE","UPDATE_IN_PROGRESS","UPDATE_REQUESTED","UPDATE_FAILED","CREATE_FAILED","ENABLE_SECURITY_FAILED","PRE_DELETE_IN_PROGRESS","DELETE_IN_PROGRESS","DELETE_FAILED","DELETE_COMPLETED","STOPPED","STOP_REQUESTED","START_REQUESTED","STOP_IN_PROGRESS","START_IN_PROGRESS","START_FAILED","STOP_FAILED","WAIT_FOR_SYNC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterResponseTypeStatusPropEnum = append(clusterResponseTypeStatusPropEnum, v)
	}
}

const (
	// ClusterResponseStatusREQUESTED captures enum value "REQUESTED"
	ClusterResponseStatusREQUESTED string = "REQUESTED"
	// ClusterResponseStatusCREATEINPROGRESS captures enum value "CREATE_IN_PROGRESS"
	ClusterResponseStatusCREATEINPROGRESS string = "CREATE_IN_PROGRESS"
	// ClusterResponseStatusAVAILABLE captures enum value "AVAILABLE"
	ClusterResponseStatusAVAILABLE string = "AVAILABLE"
	// ClusterResponseStatusUPDATEINPROGRESS captures enum value "UPDATE_IN_PROGRESS"
	ClusterResponseStatusUPDATEINPROGRESS string = "UPDATE_IN_PROGRESS"
	// ClusterResponseStatusUPDATEREQUESTED captures enum value "UPDATE_REQUESTED"
	ClusterResponseStatusUPDATEREQUESTED string = "UPDATE_REQUESTED"
	// ClusterResponseStatusUPDATEFAILED captures enum value "UPDATE_FAILED"
	ClusterResponseStatusUPDATEFAILED string = "UPDATE_FAILED"
	// ClusterResponseStatusCREATEFAILED captures enum value "CREATE_FAILED"
	ClusterResponseStatusCREATEFAILED string = "CREATE_FAILED"
	// ClusterResponseStatusENABLESECURITYFAILED captures enum value "ENABLE_SECURITY_FAILED"
	ClusterResponseStatusENABLESECURITYFAILED string = "ENABLE_SECURITY_FAILED"
	// ClusterResponseStatusPREDELETEINPROGRESS captures enum value "PRE_DELETE_IN_PROGRESS"
	ClusterResponseStatusPREDELETEINPROGRESS string = "PRE_DELETE_IN_PROGRESS"
	// ClusterResponseStatusDELETEINPROGRESS captures enum value "DELETE_IN_PROGRESS"
	ClusterResponseStatusDELETEINPROGRESS string = "DELETE_IN_PROGRESS"
	// ClusterResponseStatusDELETEFAILED captures enum value "DELETE_FAILED"
	ClusterResponseStatusDELETEFAILED string = "DELETE_FAILED"
	// ClusterResponseStatusDELETECOMPLETED captures enum value "DELETE_COMPLETED"
	ClusterResponseStatusDELETECOMPLETED string = "DELETE_COMPLETED"
	// ClusterResponseStatusSTOPPED captures enum value "STOPPED"
	ClusterResponseStatusSTOPPED string = "STOPPED"
	// ClusterResponseStatusSTOPREQUESTED captures enum value "STOP_REQUESTED"
	ClusterResponseStatusSTOPREQUESTED string = "STOP_REQUESTED"
	// ClusterResponseStatusSTARTREQUESTED captures enum value "START_REQUESTED"
	ClusterResponseStatusSTARTREQUESTED string = "START_REQUESTED"
	// ClusterResponseStatusSTOPINPROGRESS captures enum value "STOP_IN_PROGRESS"
	ClusterResponseStatusSTOPINPROGRESS string = "STOP_IN_PROGRESS"
	// ClusterResponseStatusSTARTINPROGRESS captures enum value "START_IN_PROGRESS"
	ClusterResponseStatusSTARTINPROGRESS string = "START_IN_PROGRESS"
	// ClusterResponseStatusSTARTFAILED captures enum value "START_FAILED"
	ClusterResponseStatusSTARTFAILED string = "START_FAILED"
	// ClusterResponseStatusSTOPFAILED captures enum value "STOP_FAILED"
	ClusterResponseStatusSTOPFAILED string = "STOP_FAILED"
	// ClusterResponseStatusWAITFORSYNC captures enum value "WAIT_FOR_SYNC"
	ClusterResponseStatusWAITFORSYNC string = "WAIT_FOR_SYNC"
)

// prop value enum
func (m *ClusterResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterResponse) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

func (m *ClusterResponse) validateWorkspace(formats strfmt.Registry) error {

	if swag.IsZero(m.Workspace) { // not required
		return nil
	}

	if m.Workspace != nil {

		if err := m.Workspace.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("workspace")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ClusterResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterResponse) UnmarshalBinary(b []byte) error {
	var res ClusterResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
