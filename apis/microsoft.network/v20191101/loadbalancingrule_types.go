/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191101

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

type (
	LoadBalancerRuleSpecProperties struct {
		BackendPort                int                          `json:"backendPort,omitempty"`
		DisableOutboundSnat        bool                         `json:"disableOutboundSnat,omitempty"`
		EnableFloatingIP           bool                         `json:"enableFloatingIP,omitempty"`
		EnableTCPReset             bool                         `json:"enableTCPReset,omitempty"`
		FrontendPort               int                          `json:"frontendPort,omitempty"`
		FrontendIPConfigurationRef *azcorev1.KnownTypeReference `json:"frontendIPConfigurationRef,omitempty"`
		IdleTimeoutInMinutes       int                          `json:"idleTimeoutInMinutes,omitempty"`
		// +kubebuilder:validation:Enum=Default;SourceIP;SourceIPProtocol
		LoadDistribution string `json:"loadDistribution,omitempty"`
		// +kubebuilder:validation:Enum=All;Tcp;Udp
		Protocol string `json:"protocol,omitempty"`
	}

	// LoadBalancingRuleSpec defines the desired state of LoadBalancingRule
	LoadBalancingRuleSpec struct {
		Properties *LoadBalancerRuleSpecProperties `json:"properties,omitempty"`
	}

	// LoadBalancingRuleStatus defines the observed state of LoadBalancingRule
	LoadBalancingRuleStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// LoadBalancingRule is the Schema for the loadbalancingrules API
	LoadBalancingRule struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   LoadBalancingRuleSpec   `json:"spec,omitempty"`
		Status LoadBalancingRuleStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// LoadBalancingRuleList contains a list of LoadBalancingRule
	LoadBalancingRuleList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []LoadBalancingRule `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&LoadBalancingRule{}, &LoadBalancingRuleList{})
}
