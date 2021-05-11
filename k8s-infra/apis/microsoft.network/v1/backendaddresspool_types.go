/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

type (
	BackendAddressPoolSpecProperties struct {
		BackendIPConfigurationRefs []azcorev1.KnownTypeReference `json:"backendIPConfigurations,omitempty" group:"microsoft.network.infra.azure.com" kind:"NetworkInterfaceIPConfiguration"`
		LoadBalancingRuleRefs      []azcorev1.KnownTypeReference `json:"loadBalancingRuleRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"LoadBalancingRule"`
		OutboundRuleRefs           []azcorev1.KnownTypeReference `json:"outboundRuleRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"OutboundRule"`
	}

	BackendAddressPoolSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion"`

		Properties *BackendAddressPoolSpecProperties `json:"properties,omitempty"`
	}

	// BackendAddressPoolStatus defines the observed state of BackendAddressPool
	BackendAddressPoolStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

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
