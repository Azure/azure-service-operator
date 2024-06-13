// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20240402preview

type ManagedClusters_TrustedAccessRoleBinding_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties for trusted access role binding
	Properties *TrustedAccessRoleBindingProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Properties for trusted access role binding
type TrustedAccessRoleBindingProperties_STATUS_ARM struct {
	// ProvisioningState: The current provisioning state of trusted access role binding.
	ProvisioningState *TrustedAccessRoleBindingProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// Roles: A list of roles to bind, each item is a resource type qualified role name. For example:
	// 'Microsoft.MachineLearningServices/workspaces/reader'.
	Roles []string `json:"roles"`

	// SourceResourceId: The ARM resource ID of source resource that trusted access is configured for.
	SourceResourceId *string `json:"sourceResourceId,omitempty"`
}

type TrustedAccessRoleBindingProperties_ProvisioningState_STATUS string

const (
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Canceled  = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Canceled")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Deleting  = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Deleting")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Failed    = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Failed")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Succeeded = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Succeeded")
	TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Updating  = TrustedAccessRoleBindingProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to TrustedAccessRoleBindingProperties_ProvisioningState_STATUS
var trustedAccessRoleBindingProperties_ProvisioningState_STATUS_Values = map[string]TrustedAccessRoleBindingProperties_ProvisioningState_STATUS{
	"canceled":  TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Canceled,
	"deleting":  TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Deleting,
	"failed":    TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Failed,
	"succeeded": TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Succeeded,
	"updating":  TrustedAccessRoleBindingProperties_ProvisioningState_STATUS_Updating,
}
