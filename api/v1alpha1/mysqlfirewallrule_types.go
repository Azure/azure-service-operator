// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLFirewallRuleSpec defines the desired state of MySQLFirewallRule
type MySQLFirewallRuleSpec struct {
	ResourceGroup  string `json:"resourceGroup"`
	Server         string `json:"server"`
	StartIPAddress string `json:"startIpAddress"`
	EndIPAddress   string `json:"endIpAddress"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MySQLFirewallRule is the Schema for the mysqlfirewallrules API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLFirewallRuleSpec `json:"spec,omitempty"`
	Status ASOStatus             `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLFirewallRuleList contains a list of MySQLFirewallRule
type MySQLFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLFirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLFirewallRule{}, &MySQLFirewallRuleList{})
}
