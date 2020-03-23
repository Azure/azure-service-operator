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
	NetworkInterfaceIPConfigurationSpecProperties struct {
		Primary          bool   `json:"primary,omitempty"`
		PrivateIPAddress string `json:"privateIPAddress,omitempty"`
		// +kubebuilder:validation:Enum=IPv4;IPv6
		PrivateIPAddressVersion string `json:"privateIPAddressVersion,omitempty"`
		// +kubebuilder:validation:Enum=Dynamic;Static
		PrivateIPAllocationMethod string                       `json:"privateIPAllocationMethod,omitempty"`
		PublicIPAddressRef        *azcorev1.KnownTypeReference `json:"publicIPAddressRef,omitempty"`
		SubnetRef                 *azcorev1.KnownTypeReference `json:"subnetRef,omitempty" `
	}

	// NetworkInterfaceIPConfigurationSpec defines the desired state of NetworkInterfaceIPConfiguration
	NetworkInterfaceIPConfigurationSpec struct {
		Properties *NetworkInterfaceIPConfigurationSpecProperties `json:"properties,omitempty"`
	}

	// NetworkInterfaceIPConfigurationStatus defines the observed state of NetworkInterfaceIPConfiguration
	NetworkInterfaceIPConfigurationStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// NetworkInterfaceIPConfiguration is the Schema for the networkinterfaceipconfigurations API
	NetworkInterfaceIPConfiguration struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   NetworkInterfaceIPConfigurationSpec   `json:"spec,omitempty"`
		Status NetworkInterfaceIPConfigurationStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// NetworkInterfaceIPConfigurationList contains a list of NetworkInterfaceIPConfiguration
	NetworkInterfaceIPConfigurationList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []NetworkInterfaceIPConfiguration `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&NetworkInterfaceIPConfiguration{}, &NetworkInterfaceIPConfigurationList{})
}
