// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type ServersAdvancedThreatProtectionSetting_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *AdvancedThreatProtectionProperties_STATUS `json:"properties,omitempty"`

	// SystemData: SystemData of AdvancedThreatProtectionResource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of an Advanced Threat Protection state.
type AdvancedThreatProtectionProperties_STATUS struct {
	// CreationTime: Specifies the UTC creation time of the policy.
	CreationTime *string `json:"creationTime,omitempty"`

	// State: Specifies the state of the Advanced Threat Protection, whether it is enabled or disabled or a state has not been
	// applied yet on the specific database or server.
	State *AdvancedThreatProtectionProperties_State_STATUS `json:"state,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

type AdvancedThreatProtectionProperties_State_STATUS string

const (
	AdvancedThreatProtectionProperties_State_STATUS_Disabled = AdvancedThreatProtectionProperties_State_STATUS("Disabled")
	AdvancedThreatProtectionProperties_State_STATUS_Enabled  = AdvancedThreatProtectionProperties_State_STATUS("Enabled")
	AdvancedThreatProtectionProperties_State_STATUS_New      = AdvancedThreatProtectionProperties_State_STATUS("New")
)

// Mapping from string to AdvancedThreatProtectionProperties_State_STATUS
var advancedThreatProtectionProperties_State_STATUS_Values = map[string]AdvancedThreatProtectionProperties_State_STATUS{
	"disabled": AdvancedThreatProtectionProperties_State_STATUS_Disabled,
	"enabled":  AdvancedThreatProtectionProperties_State_STATUS_Enabled,
	"new":      AdvancedThreatProtectionProperties_State_STATUS_New,
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}