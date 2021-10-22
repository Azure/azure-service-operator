// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM struct {
	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: Cosmos DB userDefinedFunction name.
	Name string `json:"name"`

	//Properties: Properties to create and update Azure Cosmos DB userDefinedFunction.
	Properties SqlUserDefinedFunctionCreateUpdatePropertiesARM `json:"properties"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags
	//can be used in viewing and grouping this resource (across resource groups). A
	//maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For
	//example, the default experience for a template type is set with
	//"defaultExperience": "Cassandra". Current "defaultExperience" values also
	//include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM) GetAPIVersion() string {
	return "2021-05-15"
}

// GetName returns the Name of the resource
func (databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM) GetName() string {
	return databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
func (databaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecARM DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionCreateUpdateProperties
type SqlUserDefinedFunctionCreateUpdatePropertiesARM struct {
	//Options: CreateUpdateOptions are a list of key-value pairs that describe the
	//resource. Supported keys are "If-Match", "If-None-Match", "Session-Token" and
	//"Throughput"
	Options *CreateUpdateOptionsARM `json:"options,omitempty"`

	//Resource: Cosmos DB SQL userDefinedFunction resource object
	Resource SqlUserDefinedFunctionResourceARM `json:"resource"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionResource
type SqlUserDefinedFunctionResourceARM struct {
	//Body: Body of the User Defined Function
	Body *string `json:"body,omitempty"`

	//Id: Name of the Cosmos DB SQL userDefinedFunction
	Id string `json:"id"`
}
