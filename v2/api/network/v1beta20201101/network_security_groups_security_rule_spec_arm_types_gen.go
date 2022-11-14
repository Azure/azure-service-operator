// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NetworkSecurityGroups_SecurityRule_Spec_ARM struct {
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name string `json:"name,omitempty"`

	// Properties: Properties of the security rule.
	Properties *SecurityRulePropertiesFormat_ARM `json:"properties,omitempty"`

	// Type: The type of the resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NetworkSecurityGroups_SecurityRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (rule NetworkSecurityGroups_SecurityRule_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (rule *NetworkSecurityGroups_SecurityRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups/securityRules"
func (rule *NetworkSecurityGroups_SecurityRule_Spec_ARM) GetType() string {
	return "Microsoft.Network/networkSecurityGroups/securityRules"
}

type SecurityRulePropertiesFormat_ARM struct {
	// Access: The network traffic is allowed or denied.
	Access *SecurityRuleAccess `json:"access,omitempty"`

	// Description: A description for this rule. Restricted to 140 chars.
	Description *string `json:"description,omitempty"`

	// DestinationAddressPrefix: The destination address prefix. CIDR or destination IP range. Asterisk '*' can also be used to
	// match all source IPs. Default tags such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used.
	DestinationAddressPrefix *string `json:"destinationAddressPrefix,omitempty"`

	// DestinationAddressPrefixes: The destination address prefixes. CIDR or destination IP ranges.
	DestinationAddressPrefixes []string `json:"destinationAddressPrefixes,omitempty"`

	// DestinationApplicationSecurityGroups: The application security group specified as destination.
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupSpec_ARM `json:"destinationApplicationSecurityGroups,omitempty"`

	// DestinationPortRange: The destination port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used
	// to match all ports.
	DestinationPortRange *string `json:"destinationPortRange,omitempty"`

	// DestinationPortRanges: The destination port ranges.
	DestinationPortRanges []string `json:"destinationPortRanges,omitempty"`

	// Direction: The direction of the rule. The direction specifies if rule will be evaluated on incoming or outgoing traffic.
	Direction *SecurityRuleDirection `json:"direction,omitempty"`

	// Priority: The priority of the rule. The value can be between 100 and 4096. The priority number must be unique for each
	// rule in the collection. The lower the priority number, the higher the priority of the rule.
	Priority *int `json:"priority,omitempty"`

	// Protocol: Network protocol this rule applies to.
	Protocol *SecurityRulePropertiesFormat_Protocol `json:"protocol,omitempty"`

	// SourceAddressPrefix: The CIDR or source IP range. Asterisk '*' can also be used to match all source IPs. Default tags
	// such as 'VirtualNetwork', 'AzureLoadBalancer' and 'Internet' can also be used. If this is an ingress rule, specifies
	// where network traffic originates from.
	SourceAddressPrefix *string `json:"sourceAddressPrefix,omitempty"`

	// SourceAddressPrefixes: The CIDR or source IP ranges.
	SourceAddressPrefixes []string `json:"sourceAddressPrefixes,omitempty"`

	// SourceApplicationSecurityGroups: The application security group specified as source.
	SourceApplicationSecurityGroups []ApplicationSecurityGroupSpec_ARM `json:"sourceApplicationSecurityGroups,omitempty"`

	// SourcePortRange: The source port or range. Integer or range between 0 and 65535. Asterisk '*' can also be used to match
	// all ports.
	SourcePortRange *string `json:"sourcePortRange,omitempty"`

	// SourcePortRanges: The source port ranges.
	SourcePortRanges []string `json:"sourcePortRanges,omitempty"`
}

type ApplicationSecurityGroupSpec_ARM struct {
	Id *string `json:"id,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}
