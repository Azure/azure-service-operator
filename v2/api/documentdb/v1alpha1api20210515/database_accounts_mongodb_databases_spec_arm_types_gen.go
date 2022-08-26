// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_mongodb_database__spec_arm_types_gen.go
// Deprecated version of DatabaseAccountsMongodbDatabase_Spec. Use v1beta20210515.DatabaseAccountsMongodbDatabase_Spec instead
type DatabaseAccountsMongodbDatabase_SpecARM struct {
	AzureName  string                                    `json:"azureName,omitempty"`
========
// Deprecated version of DatabaseAccounts_MongodbDatabases_Spec. Use v1beta20210515.DatabaseAccounts_MongodbDatabases_Spec instead
type DatabaseAccounts_MongodbDatabases_SpecARM struct {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_mongodb_databases_spec_arm_types_gen.go
	Location   *string                                   `json:"location,omitempty"`
	Name       string                                    `json:"name,omitempty"`
	Properties *MongoDBDatabaseCreateUpdatePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
}

<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_mongodb_database__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &DatabaseAccountsMongodbDatabase_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database DatabaseAccountsMongodbDatabase_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &DatabaseAccounts_MongodbDatabases_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (databases DatabaseAccounts_MongodbDatabases_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_mongodb_databases_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/documentdb/v1alpha1api20210515/database_accounts_mongodb_database__spec_arm_types_gen.go
func (database *DatabaseAccountsMongodbDatabase_SpecARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
func (database *DatabaseAccountsMongodbDatabase_SpecARM) GetType() string {
========
func (databases *DatabaseAccounts_MongodbDatabases_SpecARM) GetName() string {
	return databases.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
func (databases *DatabaseAccounts_MongodbDatabases_SpecARM) GetType() string {
>>>>>>>> main:v2/api/documentdb/v1alpha1api20210515/database_accounts_mongodb_databases_spec_arm_types_gen.go
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
}

// Deprecated version of MongoDBDatabaseCreateUpdateProperties. Use v1beta20210515.MongoDBDatabaseCreateUpdateProperties instead
type MongoDBDatabaseCreateUpdatePropertiesARM struct {
	Options  *CreateUpdateOptionsARM     `json:"options,omitempty"`
	Resource *MongoDBDatabaseResourceARM `json:"resource,omitempty"`
}

// Deprecated version of CreateUpdateOptions. Use v1beta20210515.CreateUpdateOptions instead
type CreateUpdateOptionsARM struct {
	AutoscaleSettings *AutoscaleSettingsARM `json:"autoscaleSettings,omitempty"`
	Throughput        *int                  `json:"throughput,omitempty"`
}

// Deprecated version of MongoDBDatabaseResource. Use v1beta20210515.MongoDBDatabaseResource instead
type MongoDBDatabaseResourceARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of AutoscaleSettings. Use v1beta20210515.AutoscaleSettings instead
type AutoscaleSettingsARM struct {
	MaxThroughput *int `json:"maxThroughput,omitempty"`
}
