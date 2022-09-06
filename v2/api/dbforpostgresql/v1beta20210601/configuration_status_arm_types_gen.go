// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

type Configuration_STATUSARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties_STATUSARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ConfigurationProperties_STATUSARM struct {
	// AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	// DataType: Data type of the configuration.
	DataType *ConfigurationProperties_STATUS_DataType `json:"dataType,omitempty"`

	// DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	// DocumentationLink: Configuration documentation link.
	DocumentationLink *string `json:"documentationLink,omitempty"`

	// IsConfigPendingRestart: Configuration is pending restart or not.
	IsConfigPendingRestart *bool `json:"isConfigPendingRestart,omitempty"`

	// IsDynamicConfig: Configuration dynamic or static.
	IsDynamicConfig *bool `json:"isDynamicConfig,omitempty"`

	// IsReadOnly: Configuration read-only or not.
	IsReadOnly *bool `json:"isReadOnly,omitempty"`

	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Unit: Configuration unit.
	Unit *string `json:"unit,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

type SystemData_STATUSARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_STATUS_CreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_STATUS_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type ConfigurationProperties_STATUS_DataType string

const (
	ConfigurationProperties_STATUS_DataType_Boolean     = ConfigurationProperties_STATUS_DataType("Boolean")
	ConfigurationProperties_STATUS_DataType_Enumeration = ConfigurationProperties_STATUS_DataType("Enumeration")
	ConfigurationProperties_STATUS_DataType_Integer     = ConfigurationProperties_STATUS_DataType("Integer")
	ConfigurationProperties_STATUS_DataType_Numeric     = ConfigurationProperties_STATUS_DataType("Numeric")
)

type SystemData_STATUS_CreatedByType string

const (
	SystemData_STATUS_CreatedByType_Application     = SystemData_STATUS_CreatedByType("Application")
	SystemData_STATUS_CreatedByType_Key             = SystemData_STATUS_CreatedByType("Key")
	SystemData_STATUS_CreatedByType_ManagedIdentity = SystemData_STATUS_CreatedByType("ManagedIdentity")
	SystemData_STATUS_CreatedByType_User            = SystemData_STATUS_CreatedByType("User")
)

type SystemData_STATUS_LastModifiedByType string

const (
	SystemData_STATUS_LastModifiedByType_Application     = SystemData_STATUS_LastModifiedByType("Application")
	SystemData_STATUS_LastModifiedByType_Key             = SystemData_STATUS_LastModifiedByType("Key")
	SystemData_STATUS_LastModifiedByType_ManagedIdentity = SystemData_STATUS_LastModifiedByType("ManagedIdentity")
	SystemData_STATUS_LastModifiedByType_User            = SystemData_STATUS_LastModifiedByType("User")
)
