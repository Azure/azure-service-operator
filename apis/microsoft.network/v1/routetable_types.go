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
	// RouteTableSpecProperties are the resource specific properties
	RouteTableSpecProperties struct {
		DisableBGPRoutePropagation bool                          `json:"disableBgpRoutePropagation,omitempty"`
		RouteRefs                  []azcorev1.KnownTypeReference `json:"routeRefs,omitempty" group:"microsoft.network.infra.azure.com" kind:"Route" owned:"true"`
	}

	// RouteTableSpec defines the desired state of RouteTable
	RouteTableSpec struct {
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
		Properties *RouteTableSpecProperties `json:"properties,omitempty"`
	}

	// RouteTableStatus defines the observed state of RouteTable
	RouteTableStatus struct {
		ID string `json:"id,omitempty"`
		// +k8s:conversion-gen=false
		DeploymentID      string `json:"deploymentId,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true
	// +kubebuilder:subresource:status
	// +kubebuilder:storageversion

	// RouteTable is the Schema for the routetables API
	RouteTable struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   RouteTableSpec   `json:"spec,omitempty"`
		Status RouteTableStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// RouteTableList contains a list of RouteTable
	RouteTableList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []RouteTable `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
}
