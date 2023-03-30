// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

type DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_STATUS_ARM struct {
	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of an Azure Cosmos DB resource throughput
	Properties *ThroughputSettingsGetProperties_STATUS_ARM `json:"properties,omitempty"`
	Tags       map[string]string                           `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}
