/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	OutboundRuleSpecProperties struct {
		AllocatedOutboundPorts int  `json:"allocatedOutboundPorts,omitempty"`
		EnableTCPReset         bool `json:"enableTcpReset,omitempty"`
		IdleTimeoutInMinutes   int  `json:"idleTimeoutInMinutes,omitempty"`
		// +kubebuilder:validation:Enum=All;Tcp;Udp
		Protocol string `json:"protocol,omitempty"`
	}

	// OutboundRuleSpec defines the desired state of OutboundRule
	OutboundRuleSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string                      `json:"apiVersion,omitempty"`
		Properties *OutboundRuleSpecProperties `json:"properties,omitempty"`
	}

	// OutboundRuleStatus defines the observed state of OutboundRule
	OutboundRuleStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// OutboundRule is the Schema for the outboundrules API
	OutboundRule struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   OutboundRuleSpec   `json:"spec,omitempty"`
		Status OutboundRuleStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// OutboundRuleList contains a list of OutboundRule
	OutboundRuleList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []OutboundRule `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&OutboundRule{}, &OutboundRuleList{})
}
