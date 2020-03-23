/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azcorev1 "github.com/Azure/k8s-infra/apis/core/v1"
)

// NetworkSecurityGroupSpec defines the desired state of NetworkSecurityGroup
type (
	NetworkSecurityGroupSpecProperties struct {
		SecurityRuleRefs []azcorev1.KnownTypeReference `json:"securityRules,omitempty" group:"microsoft.network.infra.azure.com" kind:"SecurityRule" owned:"true"`
	}

	NetworkSecurityGroupSpec struct {
		// +k8s:conversion-gen=false
		APIVersion string `json:"apiVersion,omitempty"`

		// ResourceGroupRef is the Azure Resource Group the VirtualNetwork resides within
		// +kubebuilder:validation:Required
		ResourceGroupRef *azcorev1.KnownTypeReference `json:"resourceGroupRef" group:"microsoft.resources.infra.azure.com" kind:"ResourceGroup"`

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
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

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
