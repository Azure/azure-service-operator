// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec. Use v1beta20210515.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec instead
type DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM struct {
	Location   *string                                `json:"location,omitempty"`
	Name       string                                 `json:"name,omitempty"`
	Properties *ThroughputSettingsUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                      `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (settings DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM) GetName() string {
	return settings.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (settings *DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}
