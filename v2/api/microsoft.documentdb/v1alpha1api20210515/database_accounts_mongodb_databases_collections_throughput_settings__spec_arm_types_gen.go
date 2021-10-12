// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM struct {
	//APIVersion: API Version of the resource type, optional when apiProfile is used
	//on the template
	APIVersion DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecAPIVersion `json:"apiVersion"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: Name of the resource
	Name string `json:"name"`

	//Properties: Properties to update Azure Cosmos DB resource throughput.
	Properties ThroughputSettingsUpdatePropertiesARM `json:"properties"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags
	//can be used in viewing and grouping this resource (across resource groups). A
	//maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For
	//example, the default experience for a template type is set with
	//"defaultExperience": "Cassandra". Current "defaultExperience" values also
	//include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`

	//Type: Resource type
	Type DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecType `json:"type"`
}

var _ genruntime.ARMResourceSpec = &DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM{}

// GetAPIVersion returns the APIVersion of the resource
func (databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM) GetAPIVersion() string {
	return string(databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM.APIVersion)
}

// GetName returns the Name of the resource
func (databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM) GetName() string {
	return databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM.Name
}

// GetType returns the Type of the resource
func (databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_SpecARM) GetType() string {
	return string(databaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecARM.Type)
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecAPIVersion string

const DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecAPIVersion20210515 = DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecAPIVersion("2021-05-15")

// +kubebuilder:validation:Enum={"Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections/throughputSettings"}
type DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecType string

const DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecTypeMicrosoftDocumentDBDatabaseAccountsMongodbDatabasesCollectionsThroughputSettings = DatabaseAccountsMongodbDatabasesCollectionsThroughputSettingsSpecType("Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/collections/throughputSettings")

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsUpdateProperties
type ThroughputSettingsUpdatePropertiesARM struct {
	//Resource: Cosmos DB resource throughput object. Either throughput is required or
	//autoscaleSettings is required, but not both.
	Resource ThroughputSettingsResourceARM `json:"resource"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputSettingsResource
type ThroughputSettingsResourceARM struct {
	//AutoscaleSettings: Cosmos DB provisioned throughput settings object
	AutoscaleSettings *AutoscaleSettingsResourceARM `json:"autoscaleSettings,omitempty"`

	//Throughput: Value of the Cosmos DB resource throughput. Either throughput is
	//required or autoscaleSettings is required, but not both.
	Throughput *int `json:"throughput,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoscaleSettingsResource
type AutoscaleSettingsResourceARM struct {
	//AutoUpgradePolicy: Cosmos DB resource auto-upgrade policy
	AutoUpgradePolicy *AutoUpgradePolicyResourceARM `json:"autoUpgradePolicy,omitempty"`

	//MaxThroughput: Represents maximum throughput container can scale up to.
	MaxThroughput int `json:"maxThroughput"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/AutoUpgradePolicyResource
type AutoUpgradePolicyResourceARM struct {
	//ThroughputPolicy: Cosmos DB resource throughput policy
	ThroughputPolicy *ThroughputPolicyResourceARM `json:"throughputPolicy,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/ThroughputPolicyResource
type ThroughputPolicyResourceARM struct {
	//IncrementPercent: Represents the percentage by which throughput can increase
	//every time throughput policy kicks in.
	IncrementPercent *int `json:"incrementPercent,omitempty"`

	//IsEnabled: Determines whether the ThroughputPolicy is active or not
	IsEnabled *bool `json:"isEnabled,omitempty"`
}
