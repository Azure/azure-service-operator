// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_FirewallRule_Spec_ARM struct {
	// Name: Resource name.
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *ServerFirewallRuleProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_FirewallRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule Servers_FirewallRule_Spec_ARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (rule *Servers_FirewallRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/firewallRules"
func (rule *Servers_FirewallRule_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/firewallRules"
}

// The properties of a server firewall rule.
type ServerFirewallRuleProperties_ARM struct {
	// EndIpAddress: The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to
	// startIpAddress. Use value '0.0.0.0' for all Azure-internal IP addresses.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// StartIpAddress: The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' for all
	// Azure-internal IP addresses.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}
