// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisCacheFirewallRuleSpec defines the desired state of RedisCacheFirewallRule
type RedisCacheFirewallRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ResourceGroup string                           `json:"resourceGroup"`
	CacheName     string                           `json:"redisCache"`
	Properties    RedisCacheFirewallRuleProperties `json:"properties"`
}

// RedisCacheFirewallRuleProperties the parameters of the RedisCacheFirewallRule
type RedisCacheFirewallRuleProperties struct {
	StartIP string `json:"startIP"`
	EndIP   string `json:"endIP"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedisCacheFirewallRule is the Schema for the rediscachefirewallrules API
// +kubebuilder:resource:shortName=rcfwr
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type RedisCacheFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisCacheFirewallRuleSpec `json:"spec,omitempty"`
	Status ASOStatus                  `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisCacheFirewallRuleList contains a list of RedisCacheFirewallRule
type RedisCacheFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisCacheFirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisCacheFirewallRule{}, &RedisCacheFirewallRuleList{})
}
