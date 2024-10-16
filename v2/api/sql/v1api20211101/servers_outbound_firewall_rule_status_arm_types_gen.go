// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

type ServersOutboundFirewallRule_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *OutboundFirewallRuleProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// The properties of an outbound firewall rule.
type OutboundFirewallRuleProperties_STATUS_ARM struct {
	// ProvisioningState: The state of the outbound rule.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}
