// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of NetworkSecurityGroups_SecurityRule_Spec. Use v1beta20201101.NetworkSecurityGroups_SecurityRule_Spec instead
type NetworkSecurityGroups_SecurityRule_Spec_ARM struct {
	Name       string                            `json:"name,omitempty"`
	Properties *SecurityRulePropertiesFormat_ARM `json:"properties,omitempty"`
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

// Deprecated version of SecurityRulePropertiesFormat. Use v1beta20201101.SecurityRulePropertiesFormat instead
type SecurityRulePropertiesFormat_ARM struct {
	Access                               *SecurityRuleAccess                                                                       `json:"access,omitempty"`
	Description                          *string                                                                                   `json:"description,omitempty"`
	DestinationAddressPrefix             *string                                                                                   `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string                                                                                  `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM `json:"destinationApplicationSecurityGroups,omitempty"`
	DestinationPortRange                 *string                                                                                   `json:"destinationPortRange,omitempty"`
	DestinationPortRanges                []string                                                                                  `json:"destinationPortRanges,omitempty"`
	Direction                            *SecurityRuleDirection                                                                    `json:"direction,omitempty"`
	Priority                             *int                                                                                      `json:"priority,omitempty"`
	Protocol                             *SecurityRulePropertiesFormat_Protocol                                                    `json:"protocol,omitempty"`
	SourceAddressPrefix                  *string                                                                                   `json:"sourceAddressPrefix,omitempty"`
	SourceAddressPrefixes                []string                                                                                  `json:"sourceAddressPrefixes,omitempty"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                      *string                                                                                   `json:"sourcePortRange,omitempty"`
	SourcePortRanges                     []string                                                                                  `json:"sourcePortRanges,omitempty"`
}

// Deprecated version of ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded. Use v1beta20201101.ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded instead
type ApplicationSecurityGroupSpec_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded_ARM struct {
	Id *string `json:"id,omitempty"`
}
