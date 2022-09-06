// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

// Deprecated version of Configuration_STATUS. Use v1beta20210601.Configuration_STATUS instead
type Configuration_STATUSARM struct {
	Id         *string                            `json:"id,omitempty"`
	Name       *string                            `json:"name,omitempty"`
	Properties *ConfigurationProperties_STATUSARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUSARM              `json:"systemData,omitempty"`
	Type       *string                            `json:"type,omitempty"`
}

// Deprecated version of ConfigurationProperties_STATUS. Use v1beta20210601.ConfigurationProperties_STATUS instead
type ConfigurationProperties_STATUSARM struct {
	AllowedValues          *string                                  `json:"allowedValues,omitempty"`
	DataType               *ConfigurationProperties_STATUS_DataType `json:"dataType,omitempty"`
	DefaultValue           *string                                  `json:"defaultValue,omitempty"`
	Description            *string                                  `json:"description,omitempty"`
	DocumentationLink      *string                                  `json:"documentationLink,omitempty"`
	IsConfigPendingRestart *bool                                    `json:"isConfigPendingRestart,omitempty"`
	IsDynamicConfig        *bool                                    `json:"isDynamicConfig,omitempty"`
	IsReadOnly             *bool                                    `json:"isReadOnly,omitempty"`
	Source                 *string                                  `json:"source,omitempty"`
	Unit                   *string                                  `json:"unit,omitempty"`
	Value                  *string                                  `json:"value,omitempty"`
}

// Deprecated version of SystemData_STATUS. Use v1beta20210601.SystemData_STATUS instead
type SystemData_STATUSARM struct {
	CreatedAt          *string                               `json:"createdAt,omitempty"`
	CreatedBy          *string                               `json:"createdBy,omitempty"`
	CreatedByType      *SystemData_STATUS_CreatedByType      `json:"createdByType,omitempty"`
	LastModifiedAt     *string                               `json:"lastModifiedAt,omitempty"`
	LastModifiedBy     *string                               `json:"lastModifiedBy,omitempty"`
	LastModifiedByType *SystemData_STATUS_LastModifiedByType `json:"lastModifiedByType,omitempty"`
}

// Deprecated version of ConfigurationProperties_STATUS_DataType. Use
// v1beta20210601.ConfigurationProperties_STATUS_DataType instead
type ConfigurationProperties_STATUS_DataType string

const (
	ConfigurationProperties_STATUS_DataType_Boolean     = ConfigurationProperties_STATUS_DataType("Boolean")
	ConfigurationProperties_STATUS_DataType_Enumeration = ConfigurationProperties_STATUS_DataType("Enumeration")
	ConfigurationProperties_STATUS_DataType_Integer     = ConfigurationProperties_STATUS_DataType("Integer")
	ConfigurationProperties_STATUS_DataType_Numeric     = ConfigurationProperties_STATUS_DataType("Numeric")
)

// Deprecated version of SystemData_STATUS_CreatedByType. Use v1beta20210601.SystemData_STATUS_CreatedByType instead
type SystemData_STATUS_CreatedByType string

const (
	SystemData_STATUS_CreatedByType_Application     = SystemData_STATUS_CreatedByType("Application")
	SystemData_STATUS_CreatedByType_Key             = SystemData_STATUS_CreatedByType("Key")
	SystemData_STATUS_CreatedByType_ManagedIdentity = SystemData_STATUS_CreatedByType("ManagedIdentity")
	SystemData_STATUS_CreatedByType_User            = SystemData_STATUS_CreatedByType("User")
)

// Deprecated version of SystemData_STATUS_LastModifiedByType. Use v1beta20210601.SystemData_STATUS_LastModifiedByType
// instead
type SystemData_STATUS_LastModifiedByType string

const (
	SystemData_STATUS_LastModifiedByType_Application     = SystemData_STATUS_LastModifiedByType("Application")
	SystemData_STATUS_LastModifiedByType_Key             = SystemData_STATUS_LastModifiedByType("Key")
	SystemData_STATUS_LastModifiedByType_ManagedIdentity = SystemData_STATUS_LastModifiedByType("ManagedIdentity")
	SystemData_STATUS_LastModifiedByType_User            = SystemData_STATUS_LastModifiedByType("User")
)
