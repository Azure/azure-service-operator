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
	LoadBalancerSpecProperties struct {
		BackendAddressPoolRefs      []azcorev1.KnownTypeReference `json:"backendAddressPools,omitempty"`
		FrontendIPConfigurationRefs []azcorev1.KnownTypeReference `json:"frontendIPConfigurationRefs,omitempty"`
		InboundNatRuleRefs          []azcorev1.KnownTypeReference `json:"inboundNatPoolRefs,omitempty"`
		LoadBalancingRuleRefs       []azcorev1.KnownTypeReference `json:"loadBalancingRuleRefs,omitempty"`
	}

	// LoadBalancerSpec defines the desired state of LoadBalancer
	LoadBalancerSpec struct {
		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef"`

		// Location of the VNET in Azure
		// +kubebuilder:validation:Required
		Location string `json:"location"`

		// +kubebuilder:validation:Enum=Basic;Standard
		SKU string `json:"sku,omitempty"`

		// Tags are user defined key value pairs
		// +optional
		Tags map[string]string `json:"tags,omitempty"`

		// Properties of the Virtual Network
		Properties *LoadBalancerSpecProperties `json:"properties,omitempty"`
	}

	// LoadBalancerStatus defines the observed state of LoadBalancer
	LoadBalancerStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// LoadBalancer is the Schema for the loadbalancers API
	LoadBalancer struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   LoadBalancerSpec   `json:"spec,omitempty"`
		Status LoadBalancerStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// LoadBalancerList contains a list of LoadBalancer
	LoadBalancerList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []LoadBalancer `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
