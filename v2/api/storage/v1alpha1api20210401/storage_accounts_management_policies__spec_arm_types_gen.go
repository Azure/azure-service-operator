// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

//Deprecated version of StorageAccountsManagementPolicies_Spec. Use v1beta20210401.StorageAccountsManagementPolicies_Spec instead
type StorageAccountsManagementPolicies_SpecARM struct {
	Name       string                         `json:"name,omitempty"`
	Properties *ManagementPolicyPropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string              `json:"tags,omitempty"`
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

//Deprecated version of ManagementPolicyProperties. Use v1beta20210401.ManagementPolicyProperties instead
type ManagementPolicyPropertiesARM struct {
	Policy *ManagementPolicySchemaARM `json:"policy,omitempty"`
}

//Deprecated version of ManagementPolicySchema. Use v1beta20210401.ManagementPolicySchema instead
type ManagementPolicySchemaARM struct {
	Rules []ManagementPolicyRuleARM `json:"rules,omitempty"`
}

//Deprecated version of ManagementPolicyRule. Use v1beta20210401.ManagementPolicyRule instead
type ManagementPolicyRuleARM struct {
	Definition *ManagementPolicyDefinitionARM `json:"definition,omitempty"`
	Enabled    *bool                          `json:"enabled,omitempty"`
	Name       *string                        `json:"name,omitempty"`
	Type       *ManagementPolicyRuleType      `json:"type,omitempty"`
}

//Deprecated version of ManagementPolicyDefinition. Use v1beta20210401.ManagementPolicyDefinition instead
type ManagementPolicyDefinitionARM struct {
	Actions *ManagementPolicyActionARM `json:"actions,omitempty"`
	Filters *ManagementPolicyFilterARM `json:"filters,omitempty"`
}

//Deprecated version of ManagementPolicyRuleType. Use v1beta20210401.ManagementPolicyRuleType instead
// +kubebuilder:validation:Enum={"Lifecycle"}
type ManagementPolicyRuleType string

const ManagementPolicyRuleTypeLifecycle = ManagementPolicyRuleType("Lifecycle")

//Deprecated version of ManagementPolicyAction. Use v1beta20210401.ManagementPolicyAction instead
type ManagementPolicyActionARM struct {
	BaseBlob *ManagementPolicyBaseBlobARM `json:"baseBlob,omitempty"`
	Snapshot *ManagementPolicySnapShotARM `json:"snapshot,omitempty"`
	Version  *ManagementPolicyVersionARM  `json:"version,omitempty"`
}

//Deprecated version of ManagementPolicyFilter. Use v1beta20210401.ManagementPolicyFilter instead
type ManagementPolicyFilterARM struct {
	BlobIndexMatch []TagFilterARM `json:"blobIndexMatch,omitempty"`
	BlobTypes      []string       `json:"blobTypes,omitempty"`
	PrefixMatch    []string       `json:"prefixMatch,omitempty"`
}

//Deprecated version of ManagementPolicyBaseBlob. Use v1beta20210401.ManagementPolicyBaseBlob instead
type ManagementPolicyBaseBlobARM struct {
	Delete                      *DateAfterModificationARM `json:"delete,omitempty"`
	EnableAutoTierToHotFromCool *bool                     `json:"enableAutoTierToHotFromCool,omitempty"`
	TierToArchive               *DateAfterModificationARM `json:"tierToArchive,omitempty"`
	TierToCool                  *DateAfterModificationARM `json:"tierToCool,omitempty"`
}

//Deprecated version of ManagementPolicySnapShot. Use v1beta20210401.ManagementPolicySnapShot instead
type ManagementPolicySnapShotARM struct {
	Delete        *DateAfterCreationARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreationARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreationARM `json:"tierToCool,omitempty"`
}

//Deprecated version of ManagementPolicyVersion. Use v1beta20210401.ManagementPolicyVersion instead
type ManagementPolicyVersionARM struct {
	Delete        *DateAfterCreationARM `json:"delete,omitempty"`
	TierToArchive *DateAfterCreationARM `json:"tierToArchive,omitempty"`
	TierToCool    *DateAfterCreationARM `json:"tierToCool,omitempty"`
}

//Deprecated version of TagFilter. Use v1beta20210401.TagFilter instead
type TagFilterARM struct {
	Name  *string `json:"name,omitempty"`
	Op    *string `json:"op,omitempty"`
	Value *string `json:"value,omitempty"`
}

//Deprecated version of DateAfterCreation. Use v1beta20210401.DateAfterCreation instead
type DateAfterCreationARM struct {
	DaysAfterCreationGreaterThan *int `json:"daysAfterCreationGreaterThan,omitempty"`
}

//Deprecated version of DateAfterModification. Use v1beta20210401.DateAfterModification instead
type DateAfterModificationARM struct {
	DaysAfterLastAccessTimeGreaterThan *int `json:"daysAfterLastAccessTimeGreaterThan,omitempty"`
	DaysAfterModificationGreaterThan   *int `json:"daysAfterModificationGreaterThan,omitempty"`
}
