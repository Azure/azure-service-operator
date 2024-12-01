// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamedValue_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: NamedValue entity contract properties for PUT operation.
	Properties *NamedValueCreateContractProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamedValue_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (value NamedValue_Spec) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetName returns the Name of the resource
func (value *NamedValue_Spec) GetName() string {
	return value.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/namedValues"
func (value *NamedValue_Spec) GetType() string {
	return "Microsoft.ApiManagement/service/namedValues"
}

// NamedValue Contract properties.
type NamedValueCreateContractProperties struct {
	// DisplayName: Unique name of NamedValue. It may contain only letters, digits, period, dash, and underscore characters.
	DisplayName *string `json:"displayName,omitempty"`

	// KeyVault: KeyVault location details of the namedValue.
	KeyVault *KeyVaultContractCreateProperties `json:"keyVault,omitempty"`

	// Secret: Determines whether the value is a secret and should be encrypted or not. Default value is false.
	Secret *bool `json:"secret,omitempty"`

	// Tags: Optional tags that when provided can be used to filter the NamedValue list.
	Tags []string `json:"tags,omitempty"`

	// Value: Value of the NamedValue. Can contain policy expressions. It may not be empty or consist only of whitespace. This
	// property will not be filled on 'GET' operations! Use '/listSecrets' POST request to get the value.
	Value *string `json:"value,omitempty"`
}

// Create keyVault contract details.
type KeyVaultContractCreateProperties struct {
	// IdentityClientId: Null for SystemAssignedIdentity or Client Id for UserAssignedIdentity , which will be used to access
	// key vault secret.
	IdentityClientId *string `json:"identityClientId,omitempty" optionalConfigMapPair:"IdentityClientId"`

	// SecretIdentifier: Key vault secret identifier for fetching secret. Providing a versioned secret will prevent
	// auto-refresh. This requires API Management service to be configured with aka.ms/apimmsi
	SecretIdentifier *string `json:"secretIdentifier,omitempty"`
}
