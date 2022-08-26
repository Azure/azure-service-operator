// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=networksecuritygroupssecurityrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={networksecuritygroupssecurityrules/status,networksecuritygroupssecurityrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20201101.NetworkSecurityGroupsSecurityRule
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/networkSecurityGroup.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}
type NetworkSecurityGroupsSecurityRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
<<<<<<< HEAD
	Spec              NetworkSecurityGroupsSecurityRule_Spec   `json:"spec,omitempty"`
	Status            NetworkSecurityGroupsSecurityRule_STATUS `json:"status,omitempty"`
=======
	Spec              NetworkSecurityGroups_SecurityRules_Spec                                   `json:"spec,omitempty"`
	Status            SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded `json:"status,omitempty"`
>>>>>>> main
}

var _ conditions.Conditioner = &NetworkSecurityGroupsSecurityRule{}

// GetConditions returns the conditions of the resource
func (rule *NetworkSecurityGroupsSecurityRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NetworkSecurityGroupsSecurityRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NetworkSecurityGroupsSecurityRule{}

// AzureName returns the Azure name of the resource
func (rule *NetworkSecurityGroupsSecurityRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (rule NetworkSecurityGroupsSecurityRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (rule *NetworkSecurityGroupsSecurityRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *NetworkSecurityGroupsSecurityRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NetworkSecurityGroupsSecurityRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkSecurityGroups/securityRules"
func (rule *NetworkSecurityGroupsSecurityRule) GetType() string {
	return "Microsoft.Network/networkSecurityGroups/securityRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NetworkSecurityGroupsSecurityRule) NewEmptyStatus() genruntime.ConvertibleStatus {
<<<<<<< HEAD
	return &NetworkSecurityGroupsSecurityRule_STATUS{}
=======
	return &SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded{}
>>>>>>> main
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NetworkSecurityGroupsSecurityRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NetworkSecurityGroupsSecurityRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
<<<<<<< HEAD
	if st, ok := status.(*NetworkSecurityGroupsSecurityRule_STATUS); ok {
=======
	if st, ok := status.(*SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded); ok {
>>>>>>> main
		rule.Status = *st
		return nil
	}

	// Convert status to required version
<<<<<<< HEAD
	var st NetworkSecurityGroupsSecurityRule_STATUS
=======
	var st SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
>>>>>>> main
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this NetworkSecurityGroupsSecurityRule is the hub type for conversion
func (rule *NetworkSecurityGroupsSecurityRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NetworkSecurityGroupsSecurityRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "NetworkSecurityGroupsSecurityRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20201101.NetworkSecurityGroupsSecurityRule
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2020-11-01/networkSecurityGroup.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkSecurityGroups/{networkSecurityGroupName}/securityRules/{securityRuleName}
type NetworkSecurityGroupsSecurityRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkSecurityGroupsSecurityRule `json:"items"`
}

<<<<<<< HEAD
// Storage version of v1beta20201101.NetworkSecurityGroupsSecurityRule_Spec
type NetworkSecurityGroupsSecurityRule_Spec struct {
=======
// Storage version of v1beta20201101.NetworkSecurityGroups_SecurityRules_Spec
type NetworkSecurityGroups_SecurityRules_Spec struct {
>>>>>>> main
	Access *string `json:"access,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                            string                         `json:"azureName,omitempty"`
	Description                          *string                        `json:"description,omitempty"`
	DestinationAddressPrefix             *string                        `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string                       `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroupSpec `json:"destinationApplicationSecurityGroups,omitempty"`
	DestinationPortRange                 *string                        `json:"destinationPortRange,omitempty"`
	DestinationPortRanges                []string                       `json:"destinationPortRanges,omitempty"`
	Direction                            *string                        `json:"direction,omitempty"`
	OriginalVersion                      string                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/NetworkSecurityGroup resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"NetworkSecurityGroup"`
	Priority    *int                               `json:"priority,omitempty"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Protocol    *string                            `json:"protocol,omitempty"`

	// Reference: Resource ID.
	Reference                       *genruntime.ResourceReference  `armReference:"Id" json:"reference,omitempty"`
	SourceAddressPrefix             *string                        `json:"sourceAddressPrefix,omitempty"`
	SourceAddressPrefixes           []string                       `json:"sourceAddressPrefixes,omitempty"`
	SourceApplicationSecurityGroups []ApplicationSecurityGroupSpec `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                 *string                        `json:"sourcePortRange,omitempty"`
	SourcePortRanges                []string                       `json:"sourcePortRanges,omitempty"`
	Type                            *string                        `json:"type,omitempty"`
}

<<<<<<< HEAD
var _ genruntime.ConvertibleSpec = &NetworkSecurityGroupsSecurityRule_Spec{}

// ConvertSpecFrom populates our NetworkSecurityGroupsSecurityRule_Spec from the provided source
func (rule *NetworkSecurityGroupsSecurityRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rule {
=======
var _ genruntime.ConvertibleSpec = &NetworkSecurityGroups_SecurityRules_Spec{}

// ConvertSpecFrom populates our NetworkSecurityGroups_SecurityRules_Spec from the provided source
func (rules *NetworkSecurityGroups_SecurityRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rules {
>>>>>>> main
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rule)
}

<<<<<<< HEAD
// ConvertSpecTo populates the provided destination from our NetworkSecurityGroupsSecurityRule_Spec
func (rule *NetworkSecurityGroupsSecurityRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rule {
=======
// ConvertSpecTo populates the provided destination from our NetworkSecurityGroups_SecurityRules_Spec
func (rules *NetworkSecurityGroups_SecurityRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rules {
>>>>>>> main
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rule)
}

<<<<<<< HEAD
// Storage version of v1beta20201101.NetworkSecurityGroupsSecurityRule_STATUS
type NetworkSecurityGroupsSecurityRule_STATUS struct {
	Access                               *string                                                                                 `json:"access,omitempty"`
	Conditions                           []conditions.Condition                                                                  `json:"conditions,omitempty"`
	Description                          *string                                                                                 `json:"description,omitempty"`
	DestinationAddressPrefix             *string                                                                                 `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string                                                                                `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"destinationApplicationSecurityGroups,omitempty"`
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
	SourceApplicationSecurityGroups      []ApplicationSecurityGroup_STATUS_NetworkSecurityGroupsSecurityRule_SubResourceEmbedded `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                      *string                                                                                 `json:"sourcePortRange,omitempty"`
	SourcePortRanges                     []string                                                                                `json:"sourcePortRanges,omitempty"`
	Type                                 *string                                                                                 `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NetworkSecurityGroupsSecurityRule_STATUS{}

// ConvertStatusFrom populates our NetworkSecurityGroupsSecurityRule_STATUS from the provided source
func (rule *NetworkSecurityGroupsSecurityRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == rule {
=======
// Storage version of v1beta20201101.SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
type SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded struct {
	Access                               *string                                                                                  `json:"access,omitempty"`
	Conditions                           []conditions.Condition                                                                   `json:"conditions,omitempty"`
	Description                          *string                                                                                  `json:"description,omitempty"`
	DestinationAddressPrefix             *string                                                                                  `json:"destinationAddressPrefix,omitempty"`
	DestinationAddressPrefixes           []string                                                                                 `json:"destinationAddressPrefixes,omitempty"`
	DestinationApplicationSecurityGroups []ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded `json:"destinationApplicationSecurityGroups,omitempty"`
	DestinationPortRange                 *string                                                                                  `json:"destinationPortRange,omitempty"`
	DestinationPortRanges                []string                                                                                 `json:"destinationPortRanges,omitempty"`
	Direction                            *string                                                                                  `json:"direction,omitempty"`
	Etag                                 *string                                                                                  `json:"etag,omitempty"`
	Id                                   *string                                                                                  `json:"id,omitempty"`
	Name                                 *string                                                                                  `json:"name,omitempty"`
	Priority                             *int                                                                                     `json:"priority,omitempty"`
	PropertyBag                          genruntime.PropertyBag                                                                   `json:"$propertyBag,omitempty"`
	Protocol                             *string                                                                                  `json:"protocol,omitempty"`
	ProvisioningState                    *string                                                                                  `json:"provisioningState,omitempty"`
	SourceAddressPrefix                  *string                                                                                  `json:"sourceAddressPrefix,omitempty"`
	SourceAddressPrefixes                []string                                                                                 `json:"sourceAddressPrefixes,omitempty"`
	SourceApplicationSecurityGroups      []ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded `json:"sourceApplicationSecurityGroups,omitempty"`
	SourcePortRange                      *string                                                                                  `json:"sourcePortRange,omitempty"`
	SourcePortRanges                     []string                                                                                 `json:"sourcePortRanges,omitempty"`
	Type                                 *string                                                                                  `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded{}

// ConvertStatusFrom populates our SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded from the provided source
func (embedded *SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == embedded {
>>>>>>> main
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(rule)
}

<<<<<<< HEAD
// ConvertStatusTo populates the provided destination from our NetworkSecurityGroupsSecurityRule_STATUS
func (rule *NetworkSecurityGroupsSecurityRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == rule {
=======
// ConvertStatusTo populates the provided destination from our SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
func (embedded *SecurityRule_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == embedded {
>>>>>>> main
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(rule)
}

// Storage version of v1beta20201101.ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded
type ApplicationSecurityGroup_STATUS_NetworkSecurityGroups_SecurityRule_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1beta20201101.ApplicationSecurityGroupSpec
type ApplicationSecurityGroupSpec struct {
	Location    *string                `json:"location,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource ID.
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
	Tags      map[string]string             `json:"tags,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroupsSecurityRule{}, &NetworkSecurityGroupsSecurityRuleList{})
}
