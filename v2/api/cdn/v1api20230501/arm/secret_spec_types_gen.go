// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type Secret_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the Secret to create.
	Properties *SecretProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Secret_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (secret Secret_Spec) GetAPIVersion() string {
	return "2023-05-01"
}

// GetName returns the Name of the resource
func (secret *Secret_Spec) GetName() string {
	return secret.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/secrets"
func (secret *Secret_Spec) GetType() string {
	return "Microsoft.Cdn/profiles/secrets"
}

// The JSON object that contains the properties of the Secret to create.
type SecretProperties struct {
	// Parameters: object which contains secret parameters
	Parameters *SecretParameters `json:"parameters,omitempty"`
}

type SecretParameters struct {
	// AzureFirstPartyManagedCertificate: Mutually exclusive with all other properties
	AzureFirstPartyManagedCertificate *AzureFirstPartyManagedCertificateParameters `json:"azureFirstPartyManagedCertificate,omitempty"`

	// CustomerCertificate: Mutually exclusive with all other properties
	CustomerCertificate *CustomerCertificateParameters `json:"customerCertificate,omitempty"`

	// ManagedCertificate: Mutually exclusive with all other properties
	ManagedCertificate *ManagedCertificateParameters `json:"managedCertificate,omitempty"`

	// UrlSigningKey: Mutually exclusive with all other properties
	UrlSigningKey *UrlSigningKeyParameters `json:"urlSigningKey,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because SecretParameters represents a discriminated union (JSON OneOf)
func (parameters SecretParameters) MarshalJSON() ([]byte, error) {
	if parameters.AzureFirstPartyManagedCertificate != nil {
		return json.Marshal(parameters.AzureFirstPartyManagedCertificate)
	}
	if parameters.CustomerCertificate != nil {
		return json.Marshal(parameters.CustomerCertificate)
	}
	if parameters.ManagedCertificate != nil {
		return json.Marshal(parameters.ManagedCertificate)
	}
	if parameters.UrlSigningKey != nil {
		return json.Marshal(parameters.UrlSigningKey)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the SecretParameters
func (parameters *SecretParameters) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["type"]
	if discriminator == "AzureFirstPartyManagedCertificate" {
		parameters.AzureFirstPartyManagedCertificate = &AzureFirstPartyManagedCertificateParameters{}
		return json.Unmarshal(data, parameters.AzureFirstPartyManagedCertificate)
	}
	if discriminator == "CustomerCertificate" {
		parameters.CustomerCertificate = &CustomerCertificateParameters{}
		return json.Unmarshal(data, parameters.CustomerCertificate)
	}
	if discriminator == "ManagedCertificate" {
		parameters.ManagedCertificate = &ManagedCertificateParameters{}
		return json.Unmarshal(data, parameters.ManagedCertificate)
	}
	if discriminator == "UrlSigningKey" {
		parameters.UrlSigningKey = &UrlSigningKeyParameters{}
		return json.Unmarshal(data, parameters.UrlSigningKey)
	}

	// No error
	return nil
}

type AzureFirstPartyManagedCertificateParameters struct {
	// SubjectAlternativeNames: The list of SANs.
	SubjectAlternativeNames []string                                         `json:"subjectAlternativeNames,omitempty"`
	Type                    AzureFirstPartyManagedCertificateParameters_Type `json:"type,omitempty"`
}

type CustomerCertificateParameters struct {
	// SecretSource: Resource reference to the Azure Key Vault certificate. Expected to be in format of
	// /subscriptions/{​​​​​​​​​subscriptionId}​​​​​​​​​/resourceGroups/{​​​​​​​​​resourceGroupName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/providers/Microsoft.KeyVault/vaults/{vaultName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/secrets/{certificateName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​
	SecretSource *ResourceReference `json:"secretSource,omitempty"`

	// SecretVersion: Version of the secret to be used
	SecretVersion *string `json:"secretVersion,omitempty"`

	// SubjectAlternativeNames: The list of SANs.
	SubjectAlternativeNames []string                           `json:"subjectAlternativeNames,omitempty"`
	Type                    CustomerCertificateParameters_Type `json:"type,omitempty"`

	// UseLatestVersion: Whether to use the latest version for the certificate
	UseLatestVersion *bool `json:"useLatestVersion,omitempty"`
}

type ManagedCertificateParameters struct {
	Type ManagedCertificateParameters_Type `json:"type,omitempty"`
}

type UrlSigningKeyParameters struct {
	// KeyId: Defines the customer defined key Id. This id will exist in the incoming request to indicate the key used to form
	// the hash.
	KeyId *string `json:"keyId,omitempty"`

	// SecretSource: Resource reference to the Azure Key Vault secret. Expected to be in format of
	// /subscriptions/{​​​​​​​​​subscriptionId}​​​​​​​​​/resourceGroups/{​​​​​​​​​resourceGroupName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/providers/Microsoft.KeyVault/vaults/{vaultName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​/secrets/{secretName}​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​​
	SecretSource *ResourceReference `json:"secretSource,omitempty"`

	// SecretVersion: Version of the secret to be used
	SecretVersion *string                      `json:"secretVersion,omitempty"`
	Type          UrlSigningKeyParameters_Type `json:"type,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureFirstPartyManagedCertificate"}
type AzureFirstPartyManagedCertificateParameters_Type string

const AzureFirstPartyManagedCertificateParameters_Type_AzureFirstPartyManagedCertificate = AzureFirstPartyManagedCertificateParameters_Type("AzureFirstPartyManagedCertificate")

// Mapping from string to AzureFirstPartyManagedCertificateParameters_Type
var azureFirstPartyManagedCertificateParameters_Type_Values = map[string]AzureFirstPartyManagedCertificateParameters_Type{
	"azurefirstpartymanagedcertificate": AzureFirstPartyManagedCertificateParameters_Type_AzureFirstPartyManagedCertificate,
}

// +kubebuilder:validation:Enum={"CustomerCertificate"}
type CustomerCertificateParameters_Type string

const CustomerCertificateParameters_Type_CustomerCertificate = CustomerCertificateParameters_Type("CustomerCertificate")

// Mapping from string to CustomerCertificateParameters_Type
var customerCertificateParameters_Type_Values = map[string]CustomerCertificateParameters_Type{
	"customercertificate": CustomerCertificateParameters_Type_CustomerCertificate,
}

// +kubebuilder:validation:Enum={"ManagedCertificate"}
type ManagedCertificateParameters_Type string

const ManagedCertificateParameters_Type_ManagedCertificate = ManagedCertificateParameters_Type("ManagedCertificate")

// Mapping from string to ManagedCertificateParameters_Type
var managedCertificateParameters_Type_Values = map[string]ManagedCertificateParameters_Type{
	"managedcertificate": ManagedCertificateParameters_Type_ManagedCertificate,
}

// +kubebuilder:validation:Enum={"UrlSigningKey"}
type UrlSigningKeyParameters_Type string

const UrlSigningKeyParameters_Type_UrlSigningKey = UrlSigningKeyParameters_Type("UrlSigningKey")

// Mapping from string to UrlSigningKeyParameters_Type
var urlSigningKeyParameters_Type_Values = map[string]UrlSigningKeyParameters_Type{
	"urlsigningkey": UrlSigningKeyParameters_Type_UrlSigningKey,
}
