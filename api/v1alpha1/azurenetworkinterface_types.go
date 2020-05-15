// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureNetworkInterfaceSpec defines the desired state of AzureNetworkInterface
type AzureNetworkInterfaceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location            string `json:"location"`
	ResourceGroup       string `json:"resourceGroup"`
	VNetName            string `json:"vnetName"`
	SubnetName          string `json:"subnetName"`
	PublicIPAddressName string `json:"publicIPAddressName"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureNetworkInterface is the Schema for the azurenetworkinterfaces API
// +kubebuilder:resource:shortName=ni,path=azurenetworkinterfaces
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureNetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureNetworkInterfaceSpec `json:"spec,omitempty"`
	Status ASOStatus                 `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureNetworkInterfaceList contains a list of AzureNetworkInterface
type AzureNetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureNetworkInterface `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureNetworkInterface{}, &AzureNetworkInterfaceList{})
}
