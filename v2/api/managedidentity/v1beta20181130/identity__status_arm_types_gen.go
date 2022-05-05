// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20181130

type Identity_StatusARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties associated with the identity.
	Properties *UserAssignedIdentityProperties_StatusARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type UserAssignedIdentityProperties_StatusARM struct {
	// ClientId: The id of the app associated with the identity. This is a random generated UUID by MSI.
	ClientId *string `json:"clientId,omitempty"`

	// PrincipalId: The id of the service principal object associated with the created identity.
	PrincipalId *string `json:"principalId,omitempty"`

	// TenantId: The id of the tenant which the identity belongs to.
	TenantId *string `json:"tenantId,omitempty"`
}
