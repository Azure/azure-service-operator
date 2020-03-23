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
	FrontendIPConfigurationSpecProperties struct {
		PrivateIPAddress string `json:"privateIPAddress,omitempty"`
		// +kubebuilder:validation:Enum=IPv4;IPv6
		PrivateIPAddressVersion string `json:"privateIPAddressVersion,omitempty"`
		// +kubebuilder:validation:Enum=Dynamic;Static
		PrivateIPAllocationMethod string                       `json:"privateIPAllocationMethod,omitempty"`
		PublicIPAddressRef        *azcorev1.KnownTypeReference `json:"publicIPAddressRef,omitempty" group:"microsoft.network.infra.azure.com" kind:"PublicIPAddress"`
		SubnetRef                 *azcorev1.KnownTypeReference `json:"subnetRef,omitempty" group:"microsoft.network.infra.azure.com" kind:"Subnet"`
		Zones                     []string                     `json:"zones,omitempty"`
	}

	// FrontendIPConfigurationSpec defines the desired state of FrontendIPConfiguration
	FrontendIPConfigurationSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion,omitempty"`

		// Properties of the Virtual Network
		Properties *FrontendIPConfigurationSpecProperties `json:"properties,omitempty"`
	}

	// FrontendIPConfigurationStatus defines the observed state of FrontendIPConfiguration
	FrontendIPConfigurationStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// FrontendIPConfiguration is the Schema for the frontendipconfigurations API
	FrontendIPConfiguration struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   FrontendIPConfigurationSpec   `json:"spec,omitempty"`
		Status FrontendIPConfigurationStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// FrontendIPConfigurationList contains a list of FrontendIPConfiguration
	FrontendIPConfigurationList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []FrontendIPConfiguration `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&FrontendIPConfiguration{}, &FrontendIPConfigurationList{})
}
