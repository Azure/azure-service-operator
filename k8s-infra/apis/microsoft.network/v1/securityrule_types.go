/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	SecurityRuleSpecProperties struct {
		// +kubebuilder:validation:Enum=Allow;Deny
		Access string `json:"access,omitempty"`

		// +kubebuiler:validation:MaxLength=140
		Description string `json:"description,omitempty"`

		DestinationAddressPrefix string   `json:"destinationAddressPrefix,omitempty"`
		DestinationPortRange     string   `json:"destinationPortRange,omitempty"`
		DestinationPortRanges    []string `json:"destinationPortRanges,omitempty"`

		// +kubebuilder:validation:Enum=Inbound;Outbound
		Direction string `json:"direction,omitempty"`
		Priority  int    `json:"priority,omitempty"`

		// +kubebuilder:validation:Enum=*;Ah;Esp;Icmp;Tcp;Udp
		Protocol string `json:"protocol,omitempty"`

		ProvisioningState     string   `json:"provisioningState,omitempty"`
		SourceAddressPrefix   string   `json:"sourceAddressPrefix,omitempty"`
		SourceAddressPrefixes string   `json:"sourceAddressPrefixes,omitempty"`
		SourcePortRange       string   `json:"sourcePortRange,omitempty"`
		SourcePortRanges      []string `json:"sourcePortRanges,omitempty"`
	}

	// SecurityRuleSpec defines the desired state of SecurityRule
	SecurityRuleSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string                      `json:"apiVersion,omitempty"`
		Properties *SecurityRuleSpecProperties `json:"properties,omitempty"`
	}

	// SecurityRuleStatus defines the observed state of SecurityRule
	SecurityRuleStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// SecurityRule is the Schema for the securityrules API
	SecurityRule struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   SecurityRuleSpec   `json:"spec,omitempty"`
		Status SecurityRuleStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// SecurityRuleList contains a list of SecurityRule
	SecurityRuleList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []SecurityRule `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&SecurityRule{}, &SecurityRuleList{})
}
