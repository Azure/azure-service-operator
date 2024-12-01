// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersDatabase_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a database.
	Properties *DatabaseProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersDatabase_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-30"
func (database FlexibleServersDatabase_Spec) GetAPIVersion() string {
	return "2023-06-30"
}

// GetName returns the Name of the resource
func (database *FlexibleServersDatabase_Spec) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/databases"
func (database *FlexibleServersDatabase_Spec) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/databases"
}

// The properties of a database.
type DatabaseProperties struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}
