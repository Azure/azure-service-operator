// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PostgreSQLFirewallRuleSpec defines the desired state of PostgreSQLFirewallRule
type PostgreSQLFirewallRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroup  string `json:"resourceGroup"`
	Server         string `json:"server"`
	StartIPAddress string `json:"startIpAddress"`
	EndIPAddress   string `json:"endIpAddress"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// PostgreSQLFirewallRule is the Schema for the postgresqlfirewallrules API
// +kubebuilder:resource:shortName=psqlfwr,path=postgresqlfirewallrule
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type PostgreSQLFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgreSQLFirewallRuleSpec `json:"spec,omitempty"`
	Status ASOStatus                  `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgreSQLFirewallRuleList contains a list of PostgreSQLFirewallRule
type PostgreSQLFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgreSQLFirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PostgreSQLFirewallRule{}, &PostgreSQLFirewallRuleList{})
}
