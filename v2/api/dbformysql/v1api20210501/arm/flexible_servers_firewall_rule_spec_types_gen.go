// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersFirewallRule_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a firewall rule.
	Properties *FirewallRuleProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersFirewallRule_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (rule FlexibleServersFirewallRule_Spec) GetAPIVersion() string {
	return "2021-05-01"
}

// GetName returns the Name of the resource
func (rule *FlexibleServersFirewallRule_Spec) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/firewallRules"
func (rule *FlexibleServersFirewallRule_Spec) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/firewallRules"
}

// The properties of a server firewall rule.
type FirewallRuleProperties struct {
	// EndIpAddress: The end IP address of the server firewall rule. Must be IPv4 format.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// StartIpAddress: The start IP address of the server firewall rule. Must be IPv4 format.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}
