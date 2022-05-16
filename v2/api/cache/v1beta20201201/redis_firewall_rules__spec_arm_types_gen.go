// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisFirewallRules_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the firewall rule.
	Name string `json:"name,omitempty"`

	// Properties: Specifies a range of IP addresses permitted to connect to the cache
	Properties *RedisFirewallRulePropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisFirewallRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (rules RedisFirewallRules_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (rules RedisFirewallRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (rules RedisFirewallRules_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/RedisFirewallRuleProperties
type RedisFirewallRulePropertiesARM struct {
	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`
}
