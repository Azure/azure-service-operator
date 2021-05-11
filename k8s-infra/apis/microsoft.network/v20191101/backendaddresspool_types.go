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
	BackendAddressPoolSpecProperties struct {
		BackendIPConfigurationRefs []azcorev1.KnownTypeReference `json:"backendIPConfigurations,omitempty"`
		LoadBalancingRuleRefs      []azcorev1.KnownTypeReference `json:"loadBalancingRuleRefs,omitempty"`
		OutboundRuleRefs           []azcorev1.KnownTypeReference `json:"outboundRuleRefs,omitempty"`
	}

	BackendAddressPoolSpec struct {
		Properties *BackendAddressPoolSpecProperties `json:"properties,omitempty"`
	}

	// BackendAddressPoolStatus defines the observed state of BackendAddressPool
	BackendAddressPoolStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// BackendAddressPool is the Schema for the backendaddresspools API
	BackendAddressPool struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   BackendAddressPoolSpec   `json:"spec,omitempty"`
		Status BackendAddressPoolStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// BackendAddressPoolList contains a list of BackendAddressPool
	BackendAddressPoolList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []BackendAddressPool `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&BackendAddressPool{}, &BackendAddressPoolList{})
}
