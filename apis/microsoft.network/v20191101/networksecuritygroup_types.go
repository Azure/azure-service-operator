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
	NetworkSecurityGroupSpecProperties struct {
		SecurityRuleRefs []azcorev1.KnownTypeReference `json:"securityRules,omitempty"`
	}

	// NetworkSecurityGroupSpec defines the desired state of NetworkSecurityGroup
	NetworkSecurityGroupSpec struct {
		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef"`

		// Location of the VNET in Azure
		// +kubebuilder:validation:Required
		Location string `json:"location"`

		// Tags are user defined key value pairs
		// +optional
		Tags map[string]string `json:"tags,omitempty"`

		// Properties of the Virtual Network
		Properties *NetworkSecurityGroupSpecProperties `json:"properties,omitempty"`
	}

	// NetworkSecurityGroupStatus defines the observed state of NetworkSecurityGroup
	NetworkSecurityGroupStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// NetworkSecurityGroup is the Schema for the networksecuritygroups API
	NetworkSecurityGroup struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   NetworkSecurityGroupSpec   `json:"spec,omitempty"`
		Status NetworkSecurityGroupStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// NetworkSecurityGroupList contains a list of NetworkSecurityGroup
	NetworkSecurityGroupList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []NetworkSecurityGroup `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&NetworkSecurityGroup{}, &NetworkSecurityGroupList{})
}
