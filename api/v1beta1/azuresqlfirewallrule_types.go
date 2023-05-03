// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSqlFirewallRuleSpec defines the desired state of AzureSqlFirewallRule
type AzureSqlFirewallRuleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Server         string `json:"server"`
	SubscriptionID string `json:"subscriptionID,omitempty"`
	StartIPAddress string `json:"startIpAddress,omitempty"`
	EndIPAddress   string `json:"endIpAddress,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// AzureSqlFirewallRule is the Schema for the azuresqlfirewallrules API
// +kubebuilder:resource:shortName=asqlfwr
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureSqlFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlFirewallRuleSpec `json:"spec,omitempty"`
	Status ASOStatus                `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlFirewallRuleList contains a list of AzureSqlFirewallRule
type AzureSqlFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlFirewallRule `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSqlFirewallRule{}, &AzureSqlFirewallRuleList{})
}

// IsSubmitted returns whether a particular firewallrule has been processed or is being processed
func (s *AzureSqlFirewallRule) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}

// NewAzureSQLFirewallRule returns a filled struct prt
func NewAzureSQLFirewallRule(names types.NamespacedName, resourceGroup, server, from, to string) *AzureSqlFirewallRule {
	return &AzureSqlFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      names.Name,
			Namespace: names.Namespace,
		},
		Spec: AzureSqlFirewallRuleSpec{
			ResourceGroup:  resourceGroup,
			Server:         server,
			StartIPAddress: from,
			EndIPAddress:   to,
		},
	}
}
