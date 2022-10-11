// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec. Use v1beta20210515.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec instead
type DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_ARM struct {
	Location   *string                                 `json:"location,omitempty"`
	Name       string                                  `json:"name,omitempty"`
	Properties *ThroughputSettingsUpdateProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string                       `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_ARM) GetName() string {
	return setting.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (setting *DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec_ARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
}
