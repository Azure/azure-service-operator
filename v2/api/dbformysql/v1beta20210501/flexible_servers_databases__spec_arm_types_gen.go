// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersDatabases_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the database.
	Name string `json:"name,omitempty"`

	// Properties: The properties of a database.
	Properties *DatabasePropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersDatabases_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (databases FlexibleServersDatabases_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (databases FlexibleServersDatabases_SpecARM) GetName() string {
	return databases.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/databases"
func (databases FlexibleServersDatabases_SpecARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/databases"
}

// Generated from: https://schema.management.azure.com/schemas/2021-05-01/Microsoft.DBforMySQL.json#/definitions/DatabaseProperties
type DatabasePropertiesARM struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}
