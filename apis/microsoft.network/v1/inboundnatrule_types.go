/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	InboundNatRuleSpecProperties struct {
		BackendPort          int  `json:"backendPort,omitempty"`
		EnableFloatingIP     bool `json:"enableFloatingIP,omitempty"`
		EnableTCPReset       bool `json:"enableTcpReset,omitempty"`
		FrontendPort         int  `json:"frontendPort,omitempty"`
		IdleTimeoutInMinutes int  `json:"idleTimeoutInMinutes,omitempty"`
		// +kubebuilder:validation:Enum=All;Tcp;Udp
		Protocol string `json:"protocol,omitempty"`
	}

	// InboundNatRuleSpec defines the desired state of InboundNatRule
	InboundNatRuleSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion,omitempty"`

		// Properties of the Virtual Network
		Properties *InboundNatRuleSpecProperties `json:"properties,omitempty"`
	}

	// InboundNatRuleStatus defines the observed state of InboundNatRule
	InboundNatRuleStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// InboundNatRule is the Schema for the inboundnatrules API
	InboundNatRule struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   InboundNatRuleSpec   `json:"spec,omitempty"`
		Status InboundNatRuleStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// InboundNatRuleList contains a list of InboundNatRule
	InboundNatRuleList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []InboundNatRule `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&InboundNatRule{}, &InboundNatRuleList{})
}
