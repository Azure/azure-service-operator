// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.NetworkSecurityGroupsSecurityRule
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups_securityRules
type NetworkSecurityGroupsSecurityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkSecurityGroupsSecurityRules_Spec                                   `json:"spec,omitempty"`
	Status            SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NetworkSecurityGroupsSecurityRule{}

// GetConditions returns the conditions of the resource
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) GetConditions() conditions.Conditions {
	return networkSecurityGroupsSecurityRule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) SetConditions(conditions conditions.Conditions) {
	networkSecurityGroupsSecurityRule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NetworkSecurityGroupsSecurityRule{}

// AzureName returns the Azure name of the resource
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) AzureName() string {
	return networkSecurityGroupsSecurityRule.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) GetSpec() genruntime.ConvertibleSpec {
	return &networkSecurityGroupsSecurityRule.Spec
}

// GetStatus returns the status of this resource
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) GetStatus() genruntime.ConvertibleStatus {
	return &networkSecurityGroupsSecurityRule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups/securityRules"
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) GetType() string {
	return "Microsoft.Network/networkSecurityGroups/securityRules"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(networkSecurityGroupsSecurityRule.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: networkSecurityGroupsSecurityRule.Namespace, Name: networkSecurityGroupsSecurityRule.Spec.Owner.Name}
}

// SetStatus sets the status of this resource
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded); ok {
		networkSecurityGroupsSecurityRule.Status = *st
		return nil
	}

	// Convert status to required version
	var st SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	networkSecurityGroupsSecurityRule.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (networkSecurityGroupsSecurityRule *NetworkSecurityGroupsSecurityRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: networkSecurityGroupsSecurityRule.Spec.OriginalVersion,
		Kind:    "NetworkSecurityGroupsSecurityRule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.NetworkSecurityGroupsSecurityRule
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkSecurityGroups_securityRules
type NetworkSecurityGroupsSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroupsSecurityRule `json:"items"`
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroupsSecurityRules_Spec
type NetworkSecurityGroupsSecurityRules_Spec struct {
	Access *string `json:"access,omitempty"`

	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                            string        `json:"azureName"`
	Description                          *string       `json:"description,omitempty"`
	DestinationAddressPrefix             *string       `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string      `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []SubResource `json:"destinationApplicationSecurityGroups,omitempty"`
	DestinationPortRange                 *string       `json:"destinationPortRange,omitempty"`
	DestinationPortRanges                []string      `json:"destinationPortRanges,omitempty"`
	Direction                            *string       `json:"direction,omitempty"`
	Location                             *string       `json:"location,omitempty"`
	OriginalVersion                      string        `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner                           genruntime.KnownResourceReference `group:"microsoft.network.azure.com" json:"owner" kind:"NetworkSecurityGroup"`
	Priority                        *int                              `json:"priority,omitempty"`
	PropertyBag                     genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Protocol                        *string                           `json:"protocol,omitempty"`
	SourceAddressPrefix             *string                           `json:"sourceAddressPrefix,omitempty"`
	SourceAddressPrefixes           []string                          `json:"sourceAddressPrefixes,omitempty"`
	SourceApplicationSecurityGroups []SubResource                     `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                 *string                           `json:"sourcePortRange,omitempty"`
	SourcePortRanges                []string                          `json:"sourcePortRanges,omitempty"`
	Tags                            map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NetworkSecurityGroupsSecurityRules_Spec{}

// ConvertSpecFrom populates our NetworkSecurityGroupsSecurityRules_Spec from the provided source
func (networkSecurityGroupsSecurityRulesSpec *NetworkSecurityGroupsSecurityRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == networkSecurityGroupsSecurityRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(networkSecurityGroupsSecurityRulesSpec)
}

// ConvertSpecTo populates the provided destination from our NetworkSecurityGroupsSecurityRules_Spec
func (networkSecurityGroupsSecurityRulesSpec *NetworkSecurityGroupsSecurityRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == networkSecurityGroupsSecurityRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(networkSecurityGroupsSecurityRulesSpec)
}

//Storage version of v1alpha1api20201101.SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
//Generated from:
type SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded struct {
	Access                               *string                                                                                 `json:"access,omitempty"`
	Conditions                           []conditions.Condition                                                                  `json:"conditions,omitempty"`
	Description                          *string                                                                                 `json:"description,omitempty"`
	DestinationAddressPrefix             *string                                                                                 `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string                                                                                `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"destinationApplicationSecurityGroups,omitempty"`
	DestinationPortRange                 *string                                                                                 `json:"destinationPortRange,omitempty"`
	DestinationPortRanges                []string                                                                                `json:"destinationPortRanges,omitempty"`
	Direction                            *string                                                                                 `json:"direction,omitempty"`
	Etag                                 *string                                                                                 `json:"etag,omitempty"`
	Id                                   *string                                                                                 `json:"id,omitempty"`
	Name                                 *string                                                                                 `json:"name,omitempty"`
	Priority                             *int                                                                                    `json:"priority,omitempty"`
	PropertyBag                          genruntime.PropertyBag                                                                  `json:"$propertyBag,omitempty"`
	Protocol                             *string                                                                                 `json:"protocol,omitempty"`
	ProvisioningState                    *string                                                                                 `json:"provisioningState,omitempty"`
	SourceAddressPrefix                  *string                                                                                 `json:"sourceAddressPrefix,omitempty"`
	SourceAddressPrefixes                []string                                                                                `json:"sourceAddressPrefixes,omitempty"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                      *string                                                                                 `json:"sourcePortRange,omitempty"`
	SourcePortRanges                     []string                                                                                `json:"sourcePortRanges,omitempty"`
	Type                                 *string                                                                                 `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded{}

// ConvertStatusFrom populates our SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded from the provided source
func (securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded *SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded)
}

// ConvertStatusTo populates the provided destination from our SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
func (securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded *SecurityRule_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(securityRuleStatusNetworkSecurityGroupsSecurityRuleSubResourceEmbedded)
}

//Storage version of v1alpha1api20201101.ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded
//Generated from:
type ApplicationSecurityGroup_Status_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.SubResource
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	//Reference: Resource ID.
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroupsSecurityRule{}, &NetworkSecurityGroupsSecurityRuleList{})
}
