// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220101

type FlexibleServers_Configuration_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a configuration.
type ConfigurationProperties_STATUS_ARM struct {
	// AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	// CurrentValue: Current value of the configuration.
	CurrentValue *string `json:"currentValue,omitempty"`

	// DataType: Data type of the configuration.
	DataType *string `json:"dataType,omitempty"`

	// DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	// DocumentationLink: The link used to get the document from community or Azure site.
	DocumentationLink *string `json:"documentationLink,omitempty"`

	// IsConfigPendingRestart: If is the configuration pending restart or not.
	IsConfigPendingRestart *ConfigurationProperties_IsConfigPendingRestart_STATUS `json:"isConfigPendingRestart,omitempty"`

	// IsDynamicConfig: If is the configuration dynamic.
	IsDynamicConfig *ConfigurationProperties_IsDynamicConfig_STATUS `json:"isDynamicConfig,omitempty"`

	// IsReadOnly: If is the configuration read only.
	IsReadOnly *ConfigurationProperties_IsReadOnly_STATUS `json:"isReadOnly,omitempty"`

	// Source: Source of the configuration.
	Source *ConfigurationProperties_Source_STATUS `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}