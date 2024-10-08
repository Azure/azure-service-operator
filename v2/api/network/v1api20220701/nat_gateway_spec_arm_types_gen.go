// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NatGateway_Spec_ARM struct {
	// Location: Resource location.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Nat Gateway properties.
	Properties *NatGatewayPropertiesFormat_ARM `json:"properties,omitempty"`

	// Sku: The nat gateway SKU.
	Sku *NatGatewaySku_ARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Zones: A list of availability zones denoting the zone in which Nat Gateway should be deployed.
	Zones []string `json:"zones,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NatGateway_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (gateway NatGateway_Spec_ARM) GetAPIVersion() string {
	return "2022-07-01"
}

// GetName returns the Name of the resource
func (gateway *NatGateway_Spec_ARM) GetName() string {
	return gateway.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/natGateways"
func (gateway *NatGateway_Spec_ARM) GetType() string {
	return "Microsoft.Network/natGateways"
}

// Nat Gateway properties.
type NatGatewayPropertiesFormat_ARM struct {
	// IdleTimeoutInMinutes: The idle timeout of the nat gateway.
	IdleTimeoutInMinutes *int `json:"idleTimeoutInMinutes,omitempty"`

	// PublicIpAddresses: An array of public ip addresses associated with the nat gateway resource.
	PublicIpAddresses []ApplicationGatewaySubResource_ARM `json:"publicIpAddresses,omitempty"`

	// PublicIpPrefixes: An array of public ip prefixes associated with the nat gateway resource.
	PublicIpPrefixes []ApplicationGatewaySubResource_ARM `json:"publicIpPrefixes,omitempty"`
}

// SKU of nat gateway.
type NatGatewaySku_ARM struct {
	// Name: Name of Nat Gateway SKU.
	Name *NatGatewaySku_Name_ARM `json:"name,omitempty"`
}

// +kubebuilder:validation:Enum={"Standard"}
type NatGatewaySku_Name_ARM string

const NatGatewaySku_Name_ARM_Standard = NatGatewaySku_Name_ARM("Standard")

// Mapping from string to NatGatewaySku_Name_ARM
var natGatewaySku_Name_ARM_Values = map[string]NatGatewaySku_Name_ARM{
	"standard": NatGatewaySku_Name_ARM_Standard,
}
