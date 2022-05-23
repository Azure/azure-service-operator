// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RedisFirewallRules_Spec. Use v1beta20201201.RedisFirewallRules_Spec instead
type RedisFirewallRules_SpecARM struct {
	Location   *string                         `json:"location,omitempty"`
	Name       string                          `json:"name,omitempty"`
	Properties *RedisFirewallRulePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string               `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisFirewallRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (rules RedisFirewallRules_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (rules *RedisFirewallRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (rules *RedisFirewallRules_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// Deprecated version of RedisFirewallRuleProperties. Use v1beta20201201.RedisFirewallRuleProperties instead
type RedisFirewallRulePropertiesARM struct {
	EndIP   *string `json:"endIP,omitempty"`
	StartIP *string `json:"startIP,omitempty"`
}
