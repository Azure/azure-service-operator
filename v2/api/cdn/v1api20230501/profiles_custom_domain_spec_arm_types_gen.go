// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Profiles_CustomDomain_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The JSON object that contains the properties of the domain to create.
	Properties *AFDDomainProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Profiles_CustomDomain_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01"
func (domain Profiles_CustomDomain_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (domain *Profiles_CustomDomain_Spec_ARM) GetName() string {
	return domain.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cdn/profiles/customDomains"
func (domain *Profiles_CustomDomain_Spec_ARM) GetType() string {
	return "Microsoft.Cdn/profiles/customDomains"
}

// The JSON object that contains the properties of the domain to create.
type AFDDomainProperties_ARM struct {
	// AzureDnsZone: Resource reference to the Azure DNS zone
	AzureDnsZone *ResourceReference_ARM `json:"azureDnsZone,omitempty"`

	// ExtendedProperties: Key-Value pair representing migration properties for domains.
	ExtendedProperties map[string]string `json:"extendedProperties,omitempty"`

	// HostName: The host name of the domain. Must be a domain name.
	HostName *string `json:"hostName,omitempty"`

	// PreValidatedCustomDomainResourceId: Resource reference to the Azure resource where custom domain ownership was
	// prevalidated
	PreValidatedCustomDomainResourceId *ResourceReference_ARM `json:"preValidatedCustomDomainResourceId,omitempty"`

	// TlsSettings: The configuration specifying how to enable HTTPS for the domain - using AzureFrontDoor managed certificate
	// or user's own certificate. If not specified, enabling ssl uses AzureFrontDoor managed certificate by default.
	TlsSettings *AFDDomainHttpsParameters_ARM `json:"tlsSettings,omitempty"`
}

// The JSON object that contains the properties to secure a domain.
type AFDDomainHttpsParameters_ARM struct {
	// CertificateType: Defines the source of the SSL certificate.
	CertificateType *AFDDomainHttpsParameters_CertificateType `json:"certificateType,omitempty"`

	// MinimumTlsVersion: TLS protocol version that will be used for Https
	MinimumTlsVersion *AFDDomainHttpsParameters_MinimumTlsVersion `json:"minimumTlsVersion,omitempty"`

	// Secret: Resource reference to the secret. ie. subs/rg/profile/secret
	Secret *ResourceReference_ARM `json:"secret,omitempty"`
}
