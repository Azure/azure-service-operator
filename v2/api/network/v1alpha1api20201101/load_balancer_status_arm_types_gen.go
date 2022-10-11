// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

// Deprecated version of LoadBalancer_STATUS. Use v1beta20201101.LoadBalancer_STATUS instead
type LoadBalancer_STATUS_ARM struct {
	Etag             *string                                  `json:"etag,omitempty"`
	ExtendedLocation *ExtendedLocation_STATUS_ARM             `json:"extendedLocation,omitempty"`
	Id               *string                                  `json:"id,omitempty"`
	Location         *string                                  `json:"location,omitempty"`
	Name             *string                                  `json:"name,omitempty"`
	Properties       *LoadBalancerPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Sku              *LoadBalancerSku_STATUS_ARM              `json:"sku,omitempty"`
	Tags             map[string]string                        `json:"tags,omitempty"`
	Type             *string                                  `json:"type,omitempty"`
}

// Deprecated version of ExtendedLocation_STATUS. Use v1beta20201101.ExtendedLocation_STATUS instead
type ExtendedLocation_STATUS_ARM struct {
	Name *string                      `json:"name,omitempty"`
	Type *ExtendedLocationType_STATUS `json:"type,omitempty"`
}

// Deprecated version of LoadBalancerPropertiesFormat_STATUS. Use v1beta20201101.LoadBalancerPropertiesFormat_STATUS instead
type LoadBalancerPropertiesFormat_STATUS_ARM struct {
	BackendAddressPools      []BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded_ARM      `json:"backendAddressPools,omitempty"`
	FrontendIPConfigurations []FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded_ARM `json:"frontendIPConfigurations,omitempty"`
	InboundNatPools          []InboundNatPool_STATUS_ARM                                           `json:"inboundNatPools,omitempty"`
	InboundNatRules          []InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded_ARM          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       []LoadBalancingRule_STATUS_ARM                                        `json:"loadBalancingRules,omitempty"`
	OutboundRules            []OutboundRule_STATUS_ARM                                             `json:"outboundRules,omitempty"`
	Probes                   []Probe_STATUS_ARM                                                    `json:"probes,omitempty"`
	ProvisioningState        *ProvisioningState_STATUS                                             `json:"provisioningState,omitempty"`
	ResourceGuid             *string                                                               `json:"resourceGuid,omitempty"`
}

// Deprecated version of LoadBalancerSku_STATUS. Use v1beta20201101.LoadBalancerSku_STATUS instead
type LoadBalancerSku_STATUS_ARM struct {
	Name *LoadBalancerSku_Name_STATUS `json:"name,omitempty"`
	Tier *LoadBalancerSku_Tier_STATUS `json:"tier,omitempty"`
}

// Deprecated version of BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded. Use v1beta20201101.BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded instead
type BackendAddressPool_STATUS_LoadBalancer_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of ExtendedLocationType_STATUS. Use v1beta20201101.ExtendedLocationType_STATUS instead
type ExtendedLocationType_STATUS string

const ExtendedLocationType_STATUS_EdgeZone = ExtendedLocationType_STATUS("EdgeZone")

// Deprecated version of FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded. Use v1beta20201101.FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded instead
type FrontendIPConfiguration_STATUS_LoadBalancer_SubResourceEmbedded_ARM struct {
	Etag       *string                                                                              `json:"etag,omitempty"`
	Id         *string                                                                              `json:"id,omitempty"`
	Name       *string                                                                              `json:"name,omitempty"`
	Properties *FrontendIPConfigurationPropertiesFormat_STATUS_LoadBalancer_SubResourceEmbedded_ARM `json:"properties,omitempty"`
	Type       *string                                                                              `json:"type,omitempty"`
	Zones      []string                                                                             `json:"zones,omitempty"`
}

// Deprecated version of InboundNatPool_STATUS. Use v1beta20201101.InboundNatPool_STATUS instead
type InboundNatPool_STATUS_ARM struct {
	Etag       *string                                    `json:"etag,omitempty"`
	Id         *string                                    `json:"id,omitempty"`
	Name       *string                                    `json:"name,omitempty"`
	Properties *InboundNatPoolPropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                    `json:"type,omitempty"`
}

// Deprecated version of InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded. Use v1beta20201101.InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded instead
type InboundNatRule_STATUS_LoadBalancer_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}

// Deprecated version of LoadBalancerSku_Name_STATUS. Use v1beta20201101.LoadBalancerSku_Name_STATUS instead
type LoadBalancerSku_Name_STATUS string

const (
	LoadBalancerSku_Name_STATUS_Basic    = LoadBalancerSku_Name_STATUS("Basic")
	LoadBalancerSku_Name_STATUS_Standard = LoadBalancerSku_Name_STATUS("Standard")
)

// Deprecated version of LoadBalancerSku_Tier_STATUS. Use v1beta20201101.LoadBalancerSku_Tier_STATUS instead
type LoadBalancerSku_Tier_STATUS string

const (
	LoadBalancerSku_Tier_STATUS_Global   = LoadBalancerSku_Tier_STATUS("Global")
	LoadBalancerSku_Tier_STATUS_Regional = LoadBalancerSku_Tier_STATUS("Regional")
)

// Deprecated version of LoadBalancingRule_STATUS. Use v1beta20201101.LoadBalancingRule_STATUS instead
type LoadBalancingRule_STATUS_ARM struct {
	Etag       *string                                       `json:"etag,omitempty"`
	Id         *string                                       `json:"id,omitempty"`
	Name       *string                                       `json:"name,omitempty"`
	Properties *LoadBalancingRulePropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                       `json:"type,omitempty"`
}

// Deprecated version of OutboundRule_STATUS. Use v1beta20201101.OutboundRule_STATUS instead
type OutboundRule_STATUS_ARM struct {
	Etag       *string                                  `json:"etag,omitempty"`
	Id         *string                                  `json:"id,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	Properties *OutboundRulePropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                  `json:"type,omitempty"`
}

// Deprecated version of Probe_STATUS. Use v1beta20201101.Probe_STATUS instead
type Probe_STATUS_ARM struct {
	Etag       *string                           `json:"etag,omitempty"`
	Id         *string                           `json:"id,omitempty"`
	Name       *string                           `json:"name,omitempty"`
	Properties *ProbePropertiesFormat_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                           `json:"type,omitempty"`
}

// Deprecated version of FrontendIPConfigurationPropertiesFormat_STATUS_LoadBalancer_SubResourceEmbedded. Use v1beta20201101.FrontendIPConfigurationPropertiesFormat_STATUS_LoadBalancer_SubResourceEmbedded instead
type FrontendIPConfigurationPropertiesFormat_STATUS_LoadBalancer_SubResourceEmbedded_ARM struct {
	InboundNatPools           []SubResource_STATUS_ARM                                     `json:"inboundNatPools,omitempty"`
	InboundNatRules           []SubResource_STATUS_ARM                                     `json:"inboundNatRules,omitempty"`
	LoadBalancingRules        []SubResource_STATUS_ARM                                     `json:"loadBalancingRules,omitempty"`
	OutboundRules             []SubResource_STATUS_ARM                                     `json:"outboundRules,omitempty"`
	PrivateIPAddress          *string                                                      `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *IPVersion_STATUS                                            `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *IPAllocationMethod_STATUS                                   `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *ProvisioningState_STATUS                                    `json:"provisioningState,omitempty"`
	PublicIPAddress           *PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded_ARM `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *SubResource_STATUS_ARM                                      `json:"publicIPPrefix,omitempty"`
	Subnet                    *Subnet_STATUS_LoadBalancer_SubResourceEmbedded_ARM          `json:"subnet,omitempty"`
}

// Deprecated version of InboundNatPoolPropertiesFormat_STATUS. Use v1beta20201101.InboundNatPoolPropertiesFormat_STATUS instead
type InboundNatPoolPropertiesFormat_STATUS_ARM struct {
	BackendPort             *int                      `json:"backendPort,omitempty"`
	EnableFloatingIP        *bool                     `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                     `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS_ARM   `json:"frontendIPConfiguration,omitempty"`
	FrontendPortRangeEnd    *int                      `json:"frontendPortRangeEnd,omitempty"`
	FrontendPortRangeStart  *int                      `json:"frontendPortRangeStart,omitempty"`
	IdleTimeoutInMinutes    *int                      `json:"idleTimeoutInMinutes,omitempty"`
	Protocol                *TransportProtocol_STATUS `json:"protocol,omitempty"`
	ProvisioningState       *ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

// Deprecated version of LoadBalancingRulePropertiesFormat_STATUS. Use v1beta20201101.LoadBalancingRulePropertiesFormat_STATUS instead
type LoadBalancingRulePropertiesFormat_STATUS_ARM struct {
	BackendAddressPool      *SubResource_STATUS_ARM                                    `json:"backendAddressPool,omitempty"`
	BackendPort             *int                                                       `json:"backendPort,omitempty"`
	DisableOutboundSnat     *bool                                                      `json:"disableOutboundSnat,omitempty"`
	EnableFloatingIP        *bool                                                      `json:"enableFloatingIP,omitempty"`
	EnableTcpReset          *bool                                                      `json:"enableTcpReset,omitempty"`
	FrontendIPConfiguration *SubResource_STATUS_ARM                                    `json:"frontendIPConfiguration,omitempty"`
	FrontendPort            *int                                                       `json:"frontendPort,omitempty"`
	IdleTimeoutInMinutes    *int                                                       `json:"idleTimeoutInMinutes,omitempty"`
	LoadDistribution        *LoadBalancingRulePropertiesFormat_LoadDistribution_STATUS `json:"loadDistribution,omitempty"`
	Probe                   *SubResource_STATUS_ARM                                    `json:"probe,omitempty"`
	Protocol                *TransportProtocol_STATUS                                  `json:"protocol,omitempty"`
	ProvisioningState       *ProvisioningState_STATUS                                  `json:"provisioningState,omitempty"`
}

// Deprecated version of OutboundRulePropertiesFormat_STATUS. Use v1beta20201101.OutboundRulePropertiesFormat_STATUS instead
type OutboundRulePropertiesFormat_STATUS_ARM struct {
	AllocatedOutboundPorts   *int                                          `json:"allocatedOutboundPorts,omitempty"`
	BackendAddressPool       *SubResource_STATUS_ARM                       `json:"backendAddressPool,omitempty"`
	EnableTcpReset           *bool                                         `json:"enableTcpReset,omitempty"`
	FrontendIPConfigurations []SubResource_STATUS_ARM                      `json:"frontendIPConfigurations,omitempty"`
	IdleTimeoutInMinutes     *int                                          `json:"idleTimeoutInMinutes,omitempty"`
	Protocol                 *OutboundRulePropertiesFormat_Protocol_STATUS `json:"protocol,omitempty"`
	ProvisioningState        *ProvisioningState_STATUS                     `json:"provisioningState,omitempty"`
}

// Deprecated version of ProbePropertiesFormat_STATUS. Use v1beta20201101.ProbePropertiesFormat_STATUS instead
type ProbePropertiesFormat_STATUS_ARM struct {
	IntervalInSeconds  *int                                   `json:"intervalInSeconds,omitempty"`
	LoadBalancingRules []SubResource_STATUS_ARM               `json:"loadBalancingRules,omitempty"`
	NumberOfProbes     *int                                   `json:"numberOfProbes,omitempty"`
	Port               *int                                   `json:"port,omitempty"`
	Protocol           *ProbePropertiesFormat_Protocol_STATUS `json:"protocol,omitempty"`
	ProvisioningState  *ProvisioningState_STATUS              `json:"provisioningState,omitempty"`
	RequestPath        *string                                `json:"requestPath,omitempty"`
}

// Deprecated version of PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded. Use v1beta20201101.PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded instead
type PublicIPAddress_STATUS_LoadBalancer_SubResourceEmbedded_ARM struct {
	ExtendedLocation *ExtendedLocation_STATUS_ARM   `json:"extendedLocation,omitempty"`
	Id               *string                        `json:"id,omitempty"`
	Sku              *PublicIPAddressSku_STATUS_ARM `json:"sku,omitempty"`
	Zones            []string                       `json:"zones,omitempty"`
}

// Deprecated version of Subnet_STATUS_LoadBalancer_SubResourceEmbedded. Use v1beta20201101.Subnet_STATUS_LoadBalancer_SubResourceEmbedded instead
type Subnet_STATUS_LoadBalancer_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}
