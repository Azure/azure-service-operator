// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkSecurityGroupsSecurityRules_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the security rule.
	Properties *SecurityRulePropertiesFormatARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkSecurityGroupsSecurityRules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (rules NetworkSecurityGroupsSecurityRules_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (rules NetworkSecurityGroupsSecurityRules_SpecARM) GetName() string {
	return rules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups/securityRules"
func (rules NetworkSecurityGroupsSecurityRules_SpecARM) GetType() string {
	return "Microsoft.Network/networkSecurityGroups/securityRules"
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SecurityRulePropertiesFormat
type SecurityRulePropertiesFormatARM struct {
	// Access: The network traffic is allowed or denied.
	Access *SecurityRulePropertiesFormatAccess `json:"access,omitempty"`

	// Description: A description for this rule. Restricted to 140 chars.
	Description *string `json:"description,omitempty"`

	// DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to
	// match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty"`

	// DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty"`

	// DestinationApplicationSecurityGroups: The application security group specified as destination.
	DestinationApplicationSecurityGroups []SubResourceARM `json:"destinationApplicationSecurityGroups,omitempty"`

	// DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used
	// to match all ports.
	DestinationPortRange *string `json:"destinationPortRange,omitempty"`

	// DestinationPortRanges: The destination port ranges.
	DestinationPortRanges []string `json:"destinationPortRanges,omitempty"`

	// Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction *SecurityRulePropertiesFormatDirection `json:"direction,omitempty"`

	// Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
	// rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `json:"priority,omitempty"`

	// Protocol: Network protocol this rule applies to.
	Protocol *SecurityRulePropertiesFormatProtocol `json:"protocol,omitempty"`

	// SourceAddressPrefix: The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags
	// such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies
	// where network traffic originates from.
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty"`

	// SourceAddressPrefixes: The CIDR or source IP ranges.
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty"`

	// SourceApplicationSecurityGroups: The application security group specified as source.
	SourceApplicationSecurityGroups []SubResourceARM `json:"sourceApplicationSecurityGroups,omitempty"`

	// SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match
	// all ports.
	SourcePortRange *string `json:"sourcePortRange,omitempty"`

	// SourcePortRanges: The source port ranges.
	SourcePortRanges []string `json:"sourcePortRanges,omitempty"`
}
