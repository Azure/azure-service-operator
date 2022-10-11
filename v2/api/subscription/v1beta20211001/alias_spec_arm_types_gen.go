// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211001

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Alias_Spec_ARM struct {
	// Name: AliasName is the name for the subscription creation request. Note that this is not the same as subscription name
	// and this doesn’t have any other lifecycle need beyond the request for subscription creation.
	Name string `json:"name,omitempty"`

	// Properties: Put subscription properties.
	Properties *PutAliasRequestProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Alias_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-10-01"
func (alias Alias_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (alias *Alias_Spec_ARM) GetName() string {
	return alias.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Subscription/aliases"
func (alias *Alias_Spec_ARM) GetType() string {
	return "Microsoft.Subscription/aliases"
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.Subscription.json#/definitions/PutAliasRequestProperties
type PutAliasRequestProperties_ARM struct {
	// AdditionalProperties: Put subscription additional properties.
	AdditionalProperties *PutAliasRequestAdditionalProperties_ARM `json:"additionalProperties,omitempty"`

	// BillingScope: Billing scope of the subscription.
	// For CustomerLed and FieldLed -
	// /billingAccounts/{billingAccountName}/billingProfiles/{billingProfileName}/invoiceSections/{invoiceSectionName}
	// For PartnerLed - /billingAccounts/{billingAccountName}/customers/{customerName}
	// For Legacy EA - /billingAccounts/{billingAccountName}/enrollmentAccounts/{enrollmentAccountName}
	BillingScope *string `json:"billingScope,omitempty"`

	// DisplayName: The friendly name of the subscription.
	DisplayName *string `json:"displayName,omitempty"`

	// ResellerId: Reseller Id
	ResellerId *string `json:"resellerId,omitempty"`

	// SubscriptionId: This parameter can be used to create alias for existing subscription Id
	SubscriptionId *string                             `json:"subscriptionId,omitempty"`
	Workload       *PutAliasRequestProperties_Workload `json:"workload,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2021-10-01/Microsoft.Subscription.json#/definitions/PutAliasRequestAdditionalProperties
type PutAliasRequestAdditionalProperties_ARM struct {
	// ManagementGroupId: Management group Id for the subscription.
	ManagementGroupId *string `json:"managementGroupId,omitempty"`

	// SubscriptionOwnerId: Owner Id of the subscription
	SubscriptionOwnerId *string `json:"subscriptionOwnerId,omitempty"`

	// SubscriptionTenantId: Tenant Id of the subscription
	SubscriptionTenantId *string `json:"subscriptionTenantId,omitempty"`

	// Tags: Tags for the subscription
	Tags map[string]string `json:"tags,omitempty"`
}

// +kubebuilder:validation:Enum={"DevTest","Production"}
type PutAliasRequestProperties_Workload string

const (
	PutAliasRequestProperties_Workload_DevTest    = PutAliasRequestProperties_Workload("DevTest")
	PutAliasRequestProperties_Workload_Production = PutAliasRequestProperties_Workload("Production")
)
