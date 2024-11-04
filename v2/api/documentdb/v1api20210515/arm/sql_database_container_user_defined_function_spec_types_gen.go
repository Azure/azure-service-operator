// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type SqlDatabaseContainerUserDefinedFunction_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to create and update Azure Cosmos DB userDefinedFunction.
	Properties *SqlUserDefinedFunctionCreateUpdateProperties `json:"properties,omitempty"`
	Tags       map[string]string                             `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &SqlDatabaseContainerUserDefinedFunction_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (function SqlDatabaseContainerUserDefinedFunction_Spec) GetAPIVersion() string {
	return "2021-05-15"
}

// GetName returns the Name of the resource
func (function *SqlDatabaseContainerUserDefinedFunction_Spec) GetName() string {
	return function.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
func (function *SqlDatabaseContainerUserDefinedFunction_Spec) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
}

// Properties to create and update Azure Cosmos DB userDefinedFunction.
type SqlUserDefinedFunctionCreateUpdateProperties struct {
	// Options: A key-value pair of options to be applied for the request. This corresponds to the headers sent with the
	// request.
	Options *CreateUpdateOptions `json:"options,omitempty"`

	// Resource: The standard JSON format of a userDefinedFunction
	Resource *SqlUserDefinedFunctionResource `json:"resource,omitempty"`
}

// Cosmos DB SQL userDefinedFunction resource object
type SqlUserDefinedFunctionResource struct {
	// Body: Body of the User Defined Function
	Body *string `json:"body,omitempty"`

	// Id: Name of the Cosmos DB SQL userDefinedFunction
	Id *string `json:"id,omitempty"`
}