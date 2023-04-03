// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220131preview

type UserAssignedIdentities_FederatedIdentityCredential_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The properties associated with the federated identity credential.
	Properties *FederatedIdentityCredentialProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties associated with a federated identity credential.
type FederatedIdentityCredentialProperties_STATUS_ARM struct {
	// Audiences: The list of audiences that can appear in the issued token.
	Audiences []string `json:"audiences,omitempty"`

	// Issuer: The URL of the issuer to be trusted.
	Issuer *string `json:"issuer,omitempty"`

	// Subject: The identifier of the external identity.
	Subject *string `json:"subject,omitempty"`
}
