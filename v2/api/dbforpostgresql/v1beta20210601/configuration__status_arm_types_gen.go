// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601

type Configuration_StatusARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties_StatusARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ConfigurationProperties_StatusARM struct {
	// AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	// DataType: Data type of the configuration.
	DataType *ConfigurationPropertiesStatusDataType `json:"dataType,omitempty"`

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

type SystemData_StatusARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemDataStatusCreatedByType `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemDataStatusLastModifiedByType `json:"lastModifiedByType,omitempty"`
}

type ConfigurationPropertiesStatusDataType string

const (
	ConfigurationPropertiesStatusDataTypeBoolean     = ConfigurationPropertiesStatusDataType("Boolean")
	ConfigurationPropertiesStatusDataTypeEnumeration = ConfigurationPropertiesStatusDataType("Enumeration")
	ConfigurationPropertiesStatusDataTypeInteger     = ConfigurationPropertiesStatusDataType("Integer")
	ConfigurationPropertiesStatusDataTypeNumeric     = ConfigurationPropertiesStatusDataType("Numeric")
)

type SystemDataStatusCreatedByType string

const (
	SystemDataStatusCreatedByTypeApplication     = SystemDataStatusCreatedByType("Application")
	SystemDataStatusCreatedByTypeKey             = SystemDataStatusCreatedByType("Key")
	SystemDataStatusCreatedByTypeManagedIdentity = SystemDataStatusCreatedByType("ManagedIdentity")
	SystemDataStatusCreatedByTypeUser            = SystemDataStatusCreatedByType("User")
)

type SystemDataStatusLastModifiedByType string

const (
	SystemDataStatusLastModifiedByTypeApplication     = SystemDataStatusLastModifiedByType("Application")
	SystemDataStatusLastModifiedByTypeKey             = SystemDataStatusLastModifiedByType("Key")
	SystemDataStatusLastModifiedByTypeManagedIdentity = SystemDataStatusLastModifiedByType("ManagedIdentity")
	SystemDataStatusLastModifiedByTypeUser            = SystemDataStatusLastModifiedByType("User")
)
