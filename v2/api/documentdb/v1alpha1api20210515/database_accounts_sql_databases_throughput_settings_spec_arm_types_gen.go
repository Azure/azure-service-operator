// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_throughput_setting__spec_arm_types_gen.go
// Deprecated version of DatabaseAccountsSqlDatabasesThroughputSetting_Spec. Use v1beta20210515.DatabaseAccountsSqlDatabasesThroughputSetting_Spec instead
type DatabaseAccountsSqlDatabasesThroughputSetting_SpecARM struct {
	AzureName  string                                 `json:"azureName,omitempty"`
========
// Deprecated version of DatabaseAccounts_SqlDatabases_ThroughputSettings_Spec. Use v1beta20210515.DatabaseAccounts_SqlDatabases_ThroughputSettings_Spec instead
type DatabaseAccounts_SqlDatabases_ThroughputSettings_SpecARM struct {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_throughput_settings_spec_arm_types_gen.go
	Location   *string                                `json:"location,omitempty"`
	Name       string                                 `json:"name,omitempty"`
	Properties *ThroughputSettingsUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                      `json:"tags,omitempty"`
}

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_throughput_setting__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesThroughputSetting_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting DatabaseAccountsSqlDatabasesThroughputSetting_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &DatabaseAccounts_SqlDatabases_ThroughputSettings_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (settings DatabaseAccounts_SqlDatabases_ThroughputSettings_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_throughput_settings_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_throughput_setting__spec_arm_types_gen.go
func (setting *DatabaseAccountsSqlDatabasesThroughputSetting_SpecARM) GetName() string {
	return setting.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (setting *DatabaseAccountsSqlDatabasesThroughputSetting_SpecARM) GetType() string {
========
func (settings *DatabaseAccounts_SqlDatabases_ThroughputSettings_SpecARM) GetName() string {
	return settings.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (settings *DatabaseAccounts_SqlDatabases_ThroughputSettings_SpecARM) GetType() string {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_sql_databases_throughput_settings_spec_arm_types_gen.go
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
}
