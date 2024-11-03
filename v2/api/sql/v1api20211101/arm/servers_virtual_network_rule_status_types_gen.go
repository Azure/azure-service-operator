// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type ServersVirtualNetworkRule_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *VirtualNetworkRuleProperties_STATUS `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of a virtual network rule.
type VirtualNetworkRuleProperties_STATUS struct {
	// IgnoreMissingVnetServiceEndpoint: Create firewall rule before the virtual network has vnet service endpoint enabled.
	IgnoreMissingVnetServiceEndpoint *bool `json:"ignoreMissingVnetServiceEndpoint,omitempty"`

	// State: Virtual Network Rule State
	State *VirtualNetworkRuleProperties_State_STATUS `json:"state,omitempty"`

	// VirtualNetworkSubnetId: The ARM resource id of the virtual network subnet.
	VirtualNetworkSubnetId *string `json:"virtualNetworkSubnetId,omitempty"`
}

type VirtualNetworkRuleProperties_State_STATUS string

const (
	VirtualNetworkRuleProperties_State_STATUS_Deleting     = VirtualNetworkRuleProperties_State_STATUS("Deleting")
	VirtualNetworkRuleProperties_State_STATUS_Failed       = VirtualNetworkRuleProperties_State_STATUS("Failed")
	VirtualNetworkRuleProperties_State_STATUS_InProgress   = VirtualNetworkRuleProperties_State_STATUS("InProgress")
	VirtualNetworkRuleProperties_State_STATUS_Initializing = VirtualNetworkRuleProperties_State_STATUS("Initializing")
	VirtualNetworkRuleProperties_State_STATUS_Ready        = VirtualNetworkRuleProperties_State_STATUS("Ready")
	VirtualNetworkRuleProperties_State_STATUS_Unknown      = VirtualNetworkRuleProperties_State_STATUS("Unknown")
)

// Mapping from string to VirtualNetworkRuleProperties_State_STATUS
var virtualNetworkRuleProperties_State_STATUS_Values = map[string]VirtualNetworkRuleProperties_State_STATUS{
	"deleting":     VirtualNetworkRuleProperties_State_STATUS_Deleting,
	"failed":       VirtualNetworkRuleProperties_State_STATUS_Failed,
	"inprogress":   VirtualNetworkRuleProperties_State_STATUS_InProgress,
	"initializing": VirtualNetworkRuleProperties_State_STATUS_Initializing,
	"ready":        VirtualNetworkRuleProperties_State_STATUS_Ready,
	"unknown":      VirtualNetworkRuleProperties_State_STATUS_Unknown,
}
