// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of FlexibleServersFirewallRules_Spec. Use v1beta20210501.FlexibleServersFirewallRules_Spec instead
type FlexibleServersFirewallRules_SpecARM struct {
	Location   *string                    `json:"location,omitempty"`
	Name       string                     `json:"name,omitempty"`
	Properties *FirewallRulePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string          `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersFirewallRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (rules FlexibleServersFirewallRules_SpecARM) GetAPIVersion() string {
	return "2021-05-01"
}

// GetName returns the Name of the resource
func (rules FlexibleServersFirewallRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/firewallRules"
func (rules FlexibleServersFirewallRules_SpecARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/firewallRules"
}

// Deprecated version of FirewallRuleProperties. Use v1beta20210501.FirewallRuleProperties instead
type FirewallRulePropertiesARM struct {
	EndIpAddress   *string `json:"endIpAddress,omitempty"`
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}
