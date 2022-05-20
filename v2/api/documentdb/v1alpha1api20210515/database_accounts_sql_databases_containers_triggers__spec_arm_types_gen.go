// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccountsSqlDatabasesContainersTriggers_Spec. Use v1beta20210515.DatabaseAccountsSqlDatabasesContainersTriggers_Spec instead
type DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM struct {
	Location   *string                              `json:"location,omitempty"`
	Name       string                               `json:"name,omitempty"`
	Properties *SqlTriggerCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (triggers DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (triggers *DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM) GetName() string {
	return triggers.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
func (triggers *DatabaseAccountsSqlDatabasesContainersTriggers_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/triggers"
}

// Deprecated version of SqlTriggerCreateUpdateProperties. Use v1beta20210515.SqlTriggerCreateUpdateProperties instead
type SqlTriggerCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM `json:"options,omitempty"`
	Resource *SqlTriggerResourceARM  `json:"resource,omitempty"`
}

// Deprecated version of SqlTriggerResource. Use v1beta20210515.SqlTriggerResource instead
type SqlTriggerResourceARM struct {
	Body             *string                             `json:"body,omitempty"`
	Id               *string                             `json:"id,omitempty"`
	TriggerOperation *SqlTriggerResourceTriggerOperation `json:"triggerOperation,omitempty"`
	TriggerType      *SqlTriggerResourceTriggerType      `json:"triggerType,omitempty"`
}

// Deprecated version of SqlTriggerResourceTriggerOperation. Use v1beta20210515.SqlTriggerResourceTriggerOperation instead
// +kubebuilder:validation:Enum={"All","Create","Delete","Replace","Update"}
type SqlTriggerResourceTriggerOperation string

const (
	SqlTriggerResourceTriggerOperationAll     = SqlTriggerResourceTriggerOperation("All")
	SqlTriggerResourceTriggerOperationCreate  = SqlTriggerResourceTriggerOperation("Create")
	SqlTriggerResourceTriggerOperationDelete  = SqlTriggerResourceTriggerOperation("Delete")
	SqlTriggerResourceTriggerOperationReplace = SqlTriggerResourceTriggerOperation("Replace")
	SqlTriggerResourceTriggerOperationUpdate  = SqlTriggerResourceTriggerOperation("Update")
)

// Deprecated version of SqlTriggerResourceTriggerType. Use v1beta20210515.SqlTriggerResourceTriggerType instead
// +kubebuilder:validation:Enum={"Post","Pre"}
type SqlTriggerResourceTriggerType string

const (
	SqlTriggerResourceTriggerTypePost = SqlTriggerResourceTriggerType("Post")
	SqlTriggerResourceTriggerTypePre  = SqlTriggerResourceTriggerType("Pre")
)
