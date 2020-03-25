// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EventhubNamespaceNetworkRuleSpec defines the desired state of EventhubNamespaceNetworkRule
type EventhubNamespaceNetworkRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Namespace     string        `json:"namespace"`
	ResourceGroup string        `json:"resourceGroup"`
	DefaultAction DefaultAction `json:"defaultAction,omitempty"`
	// VirtualNetworkRules - List VirtualNetwork Rules
	VirtualNetworkRules *[]VirtualNetworkRules `json:"virtualNetworkRules,omitempty"`
	// IPRules - List of IpRules
	IPRules *[]IPRules `json:"ipRules,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// EventhubNamespaceNetworkRule is the Schema for the eventhubnamespacenetworkrules API
type EventhubNamespaceNetworkRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EventhubNamespaceNetworkRuleSpec `json:"spec,omitempty"`
	Status ASOStatus                        `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EventhubNamespaceNetworkRuleList contains a list of EventhubNamespaceNetworkRule
type EventhubNamespaceNetworkRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventhubNamespaceNetworkRule `json:"items"`
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
	SchemeBuilder.Register(&EventhubNamespaceNetworkRule{}, &EventhubNamespaceNetworkRuleList{})
}
