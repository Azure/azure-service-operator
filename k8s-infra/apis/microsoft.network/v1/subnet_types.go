/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	// SubnetProperties are the properties of the subnet
	SubnetProperties struct {
		// AddressPrefix for the subnet, eg. 10.0.0.0/24
		AddressPrefix string `json:"addressPrefix,omitempty"`

		// AddressPrefixes are a list of address prefixes for a subnet
		AddressPrefixes []string `json:"addressPrefixes,omitempty"`
	}

	// SubnetSpec is a subnet in a Virtual Network
	SubnetSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion"`
		// Properties of the subnet
		Properties SubnetProperties `json:"properties,omitempty"`
	}

	// SubnetStatus defines the observed state of Subnet
	SubnetStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// Subnet is the Schema for the subnets API
	Subnet struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   SubnetSpec   `json:"spec,omitempty"`
		Status SubnetStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// SubnetList contains a list of Subnet
	SubnetList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []Subnet `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&Subnet{}, &SubnetList{})
}
