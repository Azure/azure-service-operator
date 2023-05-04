// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EventhubNamespaceSpec defines the desired state of EventhubNamespace
type EventhubNamespaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location   string                      `json:"location"`
	Sku        EventhubNamespaceSku        `json:"sku,omitempty"`
	Properties EventhubNamespaceProperties `json:"properties,omitempty"`
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string                        `json:"resourceGroup"`
	NetworkRule   *EventhubNamespaceNetworkRule `json:"networkRule,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// EventhubNamespace is the Schema for the eventhubnamespaces API
// +kubebuilder:resource:shortName=ehns
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type EventhubNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventhubNamespaceSpec `json:"spec,omitempty"`
	Status ASOStatus             `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceList contains a list of EventhubNamespace
type EventhubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubNamespace `json:"items"`
}

// EventhubNamespaceSku defines the sku
type EventhubNamespaceSku struct {
	Name     string `json:"name,omitempty"`     //allowedValues "Basic", "Standard"
	Tier     string `json:"tier,omitempty"`     //allowedValues "Basic", "Standard"
	Capacity int32  `json:"capacity,omitempty"` //allowedValues 1, 2, 4
}

// EventhubNamespaceProperties defines the namespace properties
type EventhubNamespaceProperties struct {
	IsAutoInflateEnabled   bool  `json:"isAutoInflateEnabled,omitempty"`
	MaximumThroughputUnits int32 `json:"maximumThroughputUnits,omitempty"`
	KafkaEnabled           bool  `json:"kafkaEnabled,omitempty"`
}

// EventhubNamespaceNetworkRule defines the namespace network rule
type EventhubNamespaceNetworkRule struct {
	DefaultAction DefaultAction `json:"defaultAction,omitempty"`
	// VirtualNetworkRules - List VirtualNetwork Rules
	VirtualNetworkRules *[]VirtualNetworkRules `json:"virtualNetworkRules,omitempty"`
	// IPRules - List of IpRules
	IPRules *[]IPRules `json:"ipRules,omitempty"`
}

// DefaultAction defined as a string
type DefaultAction string

const (
	// Allow ...
	Allow DefaultAction = "Allow"
	// Deny ...
	Deny DefaultAction = "Deny"
)

type VirtualNetworkRules struct {
	// Subnet - Full Resource ID of Virtual Network Subnet
	SubnetID string `json:"subnetId,omitempty"`
	// IgnoreMissingVnetServiceEndpoint - Value that indicates whether to ignore missing VNet Service Endpoint
	IgnoreMissingServiceEndpoint bool `json:"ignoreMissingServiceEndpoint,omitempty"`
}

type IPRules struct {
	// IPMask - IPv4 address 1.1.1.1 or CIDR notation 1.1.0.0/24
	IPMask *string `json:"ipMask,omitempty"`
}

func init() {
	SchemeBuilder.Register(&EventhubNamespace{}, &EventhubNamespaceList{})
}

func (eventhubNamespace *EventhubNamespace) IsSubmitted() bool {
	return eventhubNamespace.Status.Provisioning || eventhubNamespace.Status.Provisioned

}

func (eventhubNamespace *EventhubNamespace) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(eventhubNamespace.ObjectMeta.Finalizers, finalizerName)
}

func (eventhubNamespace *EventhubNamespace) AddFinalizer(finalizerName string) {
	eventhubNamespace.ObjectMeta.Finalizers = append(eventhubNamespace.ObjectMeta.Finalizers, finalizerName)
}

func (eventhubNamespace *EventhubNamespace) RemoveFinalizer(finalizerName string) {
	eventhubNamespace.ObjectMeta.Finalizers = helpers.RemoveString(eventhubNamespace.ObjectMeta.Finalizers, finalizerName)
}
