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
	LoadBalancerSpecProperties struct {
		BackendAddressPoolRefs      []azcorev1.KnownTypeReference `json:"backendAddressPools,omitempty" group:"microsoft.network.infra.azure.com" kind:"BackendAddressPool" owned:"true"`
		FrontendIPConfigurationRefs []azcorev1.KnownTypeReference `json:"frontendIPConfigurationRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"FrontendIPConfiguration" owned:"true"`
		InboundNatRuleRefs          []azcorev1.KnownTypeReference `json:"inboundNatPoolRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"InboundNatRule" owned:"true"`
		LoadBalancingRuleRefs       []azcorev1.KnownTypeReference `json:"loadBalancingRuleRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"LoadBalancingRule" owned:"true"`
	}

	// LoadBalancerSpec defines the desired state of LoadBalancer
	LoadBalancerSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion,omitempty"`
		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef" group:"microsoft.resources.infra.azure.com" kind:"ResourceGroup"`

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
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

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
