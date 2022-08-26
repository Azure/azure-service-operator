// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_database__spec_arm_types_gen.go
type FlexibleServersDatabase_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`
	Name      string `json:"name,omitempty"`
========
type FlexibleServers_Databases_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the database.
	Name string `json:"name,omitempty"`
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_databases_spec_arm_types_gen.go

	// Properties: The properties of a database.
	Properties *DatabasePropertiesARM `json:"properties,omitempty"`
}

<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_database__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &FlexibleServersDatabase_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (database FlexibleServersDatabase_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &FlexibleServers_Databases_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (databases FlexibleServers_Databases_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_databases_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/dbformysql/v1beta20210501/flexible_servers_database__spec_arm_types_gen.go
func (database *FlexibleServersDatabase_SpecARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/databases"
func (database *FlexibleServersDatabase_SpecARM) GetType() string {
========
func (databases *FlexibleServers_Databases_SpecARM) GetName() string {
	return databases.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/databases"
func (databases *FlexibleServers_Databases_SpecARM) GetType() string {
>>>>>>>> main:v2/api/dbformysql/v1beta20210501/flexible_servers_databases_spec_arm_types_gen.go
	return "Microsoft.DBforMySQL/flexibleServers/databases"
}

type DatabasePropertiesARM struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}
