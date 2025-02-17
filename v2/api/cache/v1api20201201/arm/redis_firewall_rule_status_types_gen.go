// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type RedisFirewallRule_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: redis cache firewall rule properties
	Properties *RedisFirewallRuleProperties_STATUS `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Specifies a range of IP addresses permitted to connect to the cache
type RedisFirewallRuleProperties_STATUS struct {
	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`
}
