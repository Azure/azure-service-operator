// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210401

type StorageAccountsManagementPolicy_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Returns the Storage Account Data Policies Rules.
	Properties *ManagementPolicyProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The Storage Account ManagementPolicy properties.
type ManagementPolicyProperties_STATUS_ARM struct {
	// LastModifiedTime: Returns the date and time the ManagementPolicies was last modified.
	LastModifiedTime *string `json:"lastModifiedTime,omitempty"`

	// Policy: The Storage Account ManagementPolicy, in JSON format. See more details in:
	// https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Policy *ManagementPolicySchema_STATUS_ARM `json:"policy,omitempty"`
}

// The Storage Account ManagementPolicies Rules. See more details in:
// https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
type ManagementPolicySchema_STATUS_ARM struct {
	// Rules: The Storage Account ManagementPolicies Rules. See more details in:
	// https://docs.microsoft.com/en-us/azure/storage/common/storage-lifecycle-managment-concepts.
	Rules []ManagementPolicyRule_STATUS_ARM `json:"rules"`
}

// An object that wraps the Lifecycle rule. Each rule is uniquely defined by name.
type ManagementPolicyRule_STATUS_ARM struct {
	// Definition: An object that defines the Lifecycle rule.
	Definition *ManagementPolicyDefinition_STATUS_ARM `json:"definition,omitempty"`

	// Enabled: Rule is enabled if set to true.
	Enabled *bool `json:"enabled,omitempty"`

	// Name: A rule name can contain any combination of alpha numeric characters. Rule name is case-sensitive. It must be
	// unique within a policy.
	Name *string `json:"name,omitempty"`

	// Type: The valid value is Lifecycle
	Type *ManagementPolicyRule_Type_STATUS_ARM `json:"type,omitempty"`
}

// An object that defines the Lifecycle rule. Each definition is made up with a filters set and an actions set.
type ManagementPolicyDefinition_STATUS_ARM struct {
	// Actions: An object that defines the action set.
	Actions *ManagementPolicyAction_STATUS_ARM `json:"actions,omitempty"`

	// Filters: An object that defines the filter set.
	Filters *ManagementPolicyFilter_STATUS_ARM `json:"filters,omitempty"`
}

type ManagementPolicyRule_Type_STATUS_ARM string

const ManagementPolicyRule_Type_STATUS_ARM_Lifecycle = ManagementPolicyRule_Type_STATUS_ARM("Lifecycle")

// Mapping from string to ManagementPolicyRule_Type_STATUS_ARM
var managementPolicyRule_Type_STATUS_ARM_Values = map[string]ManagementPolicyRule_Type_STATUS_ARM{
	"lifecycle": ManagementPolicyRule_Type_STATUS_ARM_Lifecycle,
}

// Actions are applied to the filtered blobs when the execution condition is met.
type ManagementPolicyAction_STATUS_ARM struct {
	// BaseBlob: The management policy action for base blob
	BaseBlob *ManagementPolicyBaseBlob_STATUS_ARM `json:"baseBlob,omitempty"`

	// Snapshot: The management policy action for snapshot
	Snapshot *ManagementPolicySnapShot_STATUS_ARM `json:"snapshot,omitempty"`

	// Version: The management policy action for version
	Version *ManagementPolicyVersion_STATUS_ARM `json:"version,omitempty"`
}

// Filters limit rule actions to a subset of blobs within the storage account. If multiple filters are defined, a logical
// AND is performed on all filters.
type ManagementPolicyFilter_STATUS_ARM struct {
	// BlobIndexMatch: An array of blob index tag based filters, there can be at most 10 tag filters
	BlobIndexMatch []TagFilter_STATUS_ARM `json:"blobIndexMatch"`

	// BlobTypes: An array of predefined enum values. Currently blockBlob supports all tiering and delete actions. Only delete
	// actions are supported for appendBlob.
	BlobTypes []string `json:"blobTypes"`

	// PrefixMatch: An array of strings for prefixes to be match.
	PrefixMatch []string `json:"prefixMatch"`
}

// Management policy action for base blob.
type ManagementPolicyBaseBlob_STATUS_ARM struct {
	// Delete: The function to delete the blob
	Delete *DateAfterModification_STATUS_ARM `json:"delete,omitempty"`

	// EnableAutoTierToHotFromCool: This property enables auto tiering of a blob from cool to hot on a blob access. This
	// property requires tierToCool.daysAfterLastAccessTimeGreaterThan.
	EnableAutoTierToHotFromCool *bool `json:"enableAutoTierToHotFromCool,omitempty"`

	// TierToArchive: The function to tier blobs to archive storage. Support blobs currently at Hot or Cool tier
	TierToArchive *DateAfterModification_STATUS_ARM `json:"tierToArchive,omitempty"`

	// TierToCool: The function to tier blobs to cool storage. Support blobs currently at Hot tier
	TierToCool *DateAfterModification_STATUS_ARM `json:"tierToCool,omitempty"`
}

// Management policy action for snapshot.
type ManagementPolicySnapShot_STATUS_ARM struct {
	// Delete: The function to delete the blob snapshot
	Delete *DateAfterCreation_STATUS_ARM `json:"delete,omitempty"`

	// TierToArchive: The function to tier blob snapshot to archive storage. Support blob snapshot currently at Hot or Cool tier
	TierToArchive *DateAfterCreation_STATUS_ARM `json:"tierToArchive,omitempty"`

	// TierToCool: The function to tier blob snapshot to cool storage. Support blob snapshot currently at Hot tier
	TierToCool *DateAfterCreation_STATUS_ARM `json:"tierToCool,omitempty"`
}

// Management policy action for blob version.
type ManagementPolicyVersion_STATUS_ARM struct {
	// Delete: The function to delete the blob version
	Delete *DateAfterCreation_STATUS_ARM `json:"delete,omitempty"`

	// TierToArchive: The function to tier blob version to archive storage. Support blob version currently at Hot or Cool tier
	TierToArchive *DateAfterCreation_STATUS_ARM `json:"tierToArchive,omitempty"`

	// TierToCool: The function to tier blob version to cool storage. Support blob version currently at Hot tier
	TierToCool *DateAfterCreation_STATUS_ARM `json:"tierToCool,omitempty"`
}

// Blob index tag based filtering for blob objects
type TagFilter_STATUS_ARM struct {
	// Name: This is the filter tag name, it can have 1 - 128 characters
	Name *string `json:"name,omitempty"`

	// Op: This is the comparison operator which is used for object comparison and filtering. Only == (equality operator) is
	// currently supported
	Op *string `json:"op,omitempty"`

	// Value: This is the filter tag value field used for tag based filtering, it can have 0 - 256 characters
	Value *string `json:"value,omitempty"`
}

// Object to define the number of days after creation.
type DateAfterCreation_STATUS_ARM struct {
	// DaysAfterCreationGreaterThan: Value indicating the age in days after creation
	DaysAfterCreationGreaterThan *float64 `json:"daysAfterCreationGreaterThan,omitempty"`
}

// Object to define the number of days after object last modification Or last access. Properties
// daysAfterModificationGreaterThan and daysAfterLastAccessTimeGreaterThan are mutually exclusive.
type DateAfterModification_STATUS_ARM struct {
	// DaysAfterLastAccessTimeGreaterThan: Value indicating the age in days after last blob access. This property can only be
	// used in conjunction with last access time tracking policy
	DaysAfterLastAccessTimeGreaterThan *float64 `json:"daysAfterLastAccessTimeGreaterThan,omitempty"`

	// DaysAfterModificationGreaterThan: Value indicating the age in days after last modification
	DaysAfterModificationGreaterThan *float64 `json:"daysAfterModificationGreaterThan,omitempty"`
}
