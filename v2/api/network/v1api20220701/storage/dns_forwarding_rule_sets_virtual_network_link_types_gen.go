// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=dnsforwardingrulesetsvirtualnetworklinks,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnsforwardingrulesetsvirtualnetworklinks/status,dnsforwardingrulesetsvirtualnetworklinks/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.DnsForwardingRuleSetsVirtualNetworkLink
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks/{virtualNetworkLinkName}
type DnsForwardingRuleSetsVirtualNetworkLink struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsForwardingRuleSetsVirtualNetworkLink_Spec   `json:"spec,omitempty"`
	Status            DnsForwardingRuleSetsVirtualNetworkLink_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsForwardingRuleSetsVirtualNetworkLink{}

// GetConditions returns the conditions of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) GetConditions() conditions.Conditions {
	return link.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (link *DnsForwardingRuleSetsVirtualNetworkLink) SetConditions(conditions conditions.Conditions) {
	link.Status.Conditions = conditions
}

var _ configmaps.Exporter = &DnsForwardingRuleSetsVirtualNetworkLink{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (link *DnsForwardingRuleSetsVirtualNetworkLink) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if link.Spec.OperatorSpec == nil {
		return nil
	}
	return link.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &DnsForwardingRuleSetsVirtualNetworkLink{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (link *DnsForwardingRuleSetsVirtualNetworkLink) SecretDestinationExpressions() []*core.DestinationExpression {
	if link.Spec.OperatorSpec == nil {
		return nil
	}
	return link.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &DnsForwardingRuleSetsVirtualNetworkLink{}

// AzureName returns the Azure name of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) AzureName() string {
	return link.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (link DnsForwardingRuleSetsVirtualNetworkLink) GetAPIVersion() string {
	return "2022-07-01"
}

// GetResourceScope returns the scope of the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) GetSpec() genruntime.ConvertibleSpec {
	return &link.Spec
}

// GetStatus returns the status of this resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) GetStatus() genruntime.ConvertibleStatus {
	return &link.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsForwardingRulesets/virtualNetworkLinks"
func (link *DnsForwardingRuleSetsVirtualNetworkLink) GetType() string {
	return "Microsoft.Network/dnsForwardingRulesets/virtualNetworkLinks"
}

// NewEmptyStatus returns a new empty (blank) status
func (link *DnsForwardingRuleSetsVirtualNetworkLink) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsForwardingRuleSetsVirtualNetworkLink_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (link *DnsForwardingRuleSetsVirtualNetworkLink) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(link.Spec)
	return link.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsForwardingRuleSetsVirtualNetworkLink_STATUS); ok {
		link.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsForwardingRuleSetsVirtualNetworkLink_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	link.Status = st
	return nil
}

// Hub marks that this DnsForwardingRuleSetsVirtualNetworkLink is the hub type for conversion
func (link *DnsForwardingRuleSetsVirtualNetworkLink) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (link *DnsForwardingRuleSetsVirtualNetworkLink) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: link.Spec.OriginalVersion,
		Kind:    "DnsForwardingRuleSetsVirtualNetworkLink",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.DnsForwardingRuleSetsVirtualNetworkLink
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsForwardingRulesets/{dnsForwardingRulesetName}/virtualNetworkLinks/{virtualNetworkLinkName}
type DnsForwardingRuleSetsVirtualNetworkLinkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsForwardingRuleSetsVirtualNetworkLink `json:"items"`
}

// Storage version of v1api20220701.DnsForwardingRuleSetsVirtualNetworkLink_Spec
type DnsForwardingRuleSetsVirtualNetworkLink_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                               `json:"azureName,omitempty"`
	Metadata        map[string]string                                    `json:"metadata,omitempty"`
	OperatorSpec    *DnsForwardingRuleSetsVirtualNetworkLinkOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/DnsForwardingRuleset resource
	Owner          *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"DnsForwardingRuleset"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	VirtualNetwork *SubResource                       `json:"virtualNetwork,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsForwardingRuleSetsVirtualNetworkLink_Spec{}

// ConvertSpecFrom populates our DnsForwardingRuleSetsVirtualNetworkLink_Spec from the provided source
func (link *DnsForwardingRuleSetsVirtualNetworkLink_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == link {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(link)
}

// ConvertSpecTo populates the provided destination from our DnsForwardingRuleSetsVirtualNetworkLink_Spec
func (link *DnsForwardingRuleSetsVirtualNetworkLink_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == link {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(link)
}

// Storage version of v1api20220701.DnsForwardingRuleSetsVirtualNetworkLink_STATUS
type DnsForwardingRuleSetsVirtualNetworkLink_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Etag              *string                `json:"etag,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Metadata          map[string]string      `json:"metadata,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	SystemData        *SystemData_STATUS     `json:"systemData,omitempty"`
	Type              *string                `json:"type,omitempty"`
	VirtualNetwork    *SubResource_STATUS    `json:"virtualNetwork,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsForwardingRuleSetsVirtualNetworkLink_STATUS{}

// ConvertStatusFrom populates our DnsForwardingRuleSetsVirtualNetworkLink_STATUS from the provided source
func (link *DnsForwardingRuleSetsVirtualNetworkLink_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == link {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(link)
}

// ConvertStatusTo populates the provided destination from our DnsForwardingRuleSetsVirtualNetworkLink_STATUS
func (link *DnsForwardingRuleSetsVirtualNetworkLink_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == link {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(link)
}

// Storage version of v1api20220701.DnsForwardingRuleSetsVirtualNetworkLinkOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type DnsForwardingRuleSetsVirtualNetworkLinkOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DnsForwardingRuleSetsVirtualNetworkLink{}, &DnsForwardingRuleSetsVirtualNetworkLinkList{})
}
