// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsManagementPolicies_SpecARM struct {
	//Name: The name of the Storage Account Management Policy. It should always be 'default'
	Name string `json:"name,omitempty"`

	//Properties: The Storage Account ManagementPolicy properties.
	Properties *ManagementPolicyPropertiesARM `json:"properties,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsManagementPolicies_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (policies StorageAccountsManagementPolicies_SpecARM) GetAPIVersion() string {
	return "2021-04-01"
}

// GetName returns the Name of the resource
func (policies StorageAccountsManagementPolicies_SpecARM) GetName() string {
	return policies.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/managementPolicies"
func (policies StorageAccountsManagementPolicies_SpecARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/managementPolicies"
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicyProperties
type ManagementPolicyPropertiesARM struct {
	//Policy: The Storage Account ManagementPolicies Rules. See more details in:
	//https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Policy *ManagementPolicySchemaARM `json:"policy,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicySchema
type ManagementPolicySchemaARM struct {
	//Rules: The Storage Account ManagementPolicies Rules. See more details in:
	//https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Rules []ManagementPolicyRuleARM `json:"rules,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicyRule
type ManagementPolicyRuleARM struct {
	//Definition: An object that defines the Lifecycle rule. Each definition is made up with a filters set and an actions set.
	Definition *ManagementPolicyDefinitionARM `json:"definition,omitempty"`

	//Enabled: Rule is enabled if set to true.
	Enabled *bool `json:"enabled,omitempty"`

	//Name: A rule name can contain any combination of alpha numeric characters. Rule name is case-sensitive. It must be
	//unique within a policy.
	Name *string `json:"name,omitempty"`

	//Type: The valid value is Lifecycle
	Type *ManagementPolicyRuleType `json:"type,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicyDefinition
type ManagementPolicyDefinitionARM struct {
	//Actions: Actions are applied to the filtered blobs when the execution condition is met.
	Actions *ManagementPolicyActionARM `json:"actions,omitempty"`

	//Filters: Filters limit rule actions to a subset of blobs within the storage account. If multiple filters are defined, a
	//logical AND is performed on all filters.
	Filters *ManagementPolicyFilterARM `json:"filters,omitempty"`
}

// +kubebuilder:validation:Enum={"Lifecycle"}
type ManagementPolicyRuleType string

const ManagementPolicyRuleTypeLifecycle = ManagementPolicyRuleType("Lifecycle")

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicyAction
type ManagementPolicyActionARM struct {
	//BaseBlob: Management policy action for base blob.
	BaseBlob *ManagementPolicyBaseBlobARM `json:"baseBlob,omitempty"`

	//Snapshot: Management policy action for snapshot.
	Snapshot *ManagementPolicySnapShotARM `json:"snapshot,omitempty"`

	//Version: Management policy action for blob version.
	Version *ManagementPolicyVersionARM `json:"version,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicyFilter
type ManagementPolicyFilterARM struct {
	//BlobIndexMatch: An array of blob index tag based filters, there can be at most 10 tag filters
	BlobIndexMatch []TagFilterARM `json:"blobIndexMatch,omitempty"`

	//BlobTypes: An array of predefined enum values. Currently blockBlob supports all tiering and delete actions. Only delete
	//actions are supported for appendBlob.
	BlobTypes []string `json:"blobTypes,omitempty"`

	//PrefixMatch: An array of strings for prefixes to be match.
	PrefixMatch []string `json:"prefixMatch,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicyBaseBlob
type ManagementPolicyBaseBlobARM struct {
	//Delete: Object to define the number of days after object last modification Or last access. Properties
	//daysAfterModificationGreaterThan and daysAfterLastAccessTimeGreaterThan are mutually exclusive.
	Delete *DateAfterModificationARM `json:"delete,omitempty"`

	//EnableAutoTierToHotFromCool: This property enables auto tiering of a blob from cool to hot on a blob access. This
	//property requires tierToCool.daysAfterLastAccessTimeGreaterThan.
	EnableAutoTierToHotFromCool *bool `json:"enableAutoTierToHotFromCool,omitempty"`

	//TierToArchive: Object to define the number of days after object last modification Or last access. Properties
	//daysAfterModificationGreaterThan and daysAfterLastAccessTimeGreaterThan are mutually exclusive.
	TierToArchive *DateAfterModificationARM `json:"tierToArchive,omitempty"`

	//TierToCool: Object to define the number of days after object last modification Or last access. Properties
	//daysAfterModificationGreaterThan and daysAfterLastAccessTimeGreaterThan are mutually exclusive.
	TierToCool *DateAfterModificationARM `json:"tierToCool,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicySnapShot
type ManagementPolicySnapShotARM struct {
	//Delete: Object to define the number of days after creation.
	Delete *DateAfterCreationARM `json:"delete,omitempty"`

	//TierToArchive: Object to define the number of days after creation.
	TierToArchive *DateAfterCreationARM `json:"tierToArchive,omitempty"`

	//TierToCool: Object to define the number of days after creation.
	TierToCool *DateAfterCreationARM `json:"tierToCool,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ManagementPolicyVersion
type ManagementPolicyVersionARM struct {
	//Delete: Object to define the number of days after creation.
	Delete *DateAfterCreationARM `json:"delete,omitempty"`

	//TierToArchive: Object to define the number of days after creation.
	TierToArchive *DateAfterCreationARM `json:"tierToArchive,omitempty"`

	//TierToCool: Object to define the number of days after creation.
	TierToCool *DateAfterCreationARM `json:"tierToCool,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/TagFilter
type TagFilterARM struct {
	//Name: This is the filter tag name, it can have 1 - 128 characters
	Name *string `json:"name,omitempty"`

	//Op: This is the comparison operator which is used for object comparison and filtering. Only == (equality operator) is
	//currently supported
	Op *string `json:"op,omitempty"`

	//Value: This is the filter tag value field used for tag based filtering, it can have 0 - 256 characters
	Value *string `json:"value,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/DateAfterCreation
type DateAfterCreationARM struct {
	//DaysAfterCreationGreaterThan: Value indicating the age in days after creation
	DaysAfterCreationGreaterThan *int `json:"daysAfterCreationGreaterThan,omitempty"`
}

//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/DateAfterModification
type DateAfterModificationARM struct {
	//DaysAfterLastAccessTimeGreaterThan: Value indicating the age in days after last blob access. This property can only be
	//used in conjunction with last access time tracking policy
	DaysAfterLastAccessTimeGreaterThan *int `json:"daysAfterLastAccessTimeGreaterThan,omitempty"`

	//DaysAfterModificationGreaterThan: Value indicating the age in days after last modification
	DaysAfterModificationGreaterThan *int `json:"daysAfterModificationGreaterThan,omitempty"`
}
