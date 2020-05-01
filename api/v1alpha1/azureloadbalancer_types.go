// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureLoadBalancerSpec defines the desired state of AzureLoadBalancer
type AzureLoadBalancerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location               string `json:"location"`
	ResourceGroup          string `json:"resourceGroup"`
	PublicIPAddressName    string `json:"publicIPAddressName"`
	BackendAddressPoolName string `json:"backendAddressPoolName"`
	InboundNatPoolName     string `json:"inboundNatPoolName"`
	FrontendPortRangeStart int    `json:"frontendPortRangeStart"`
	FrontendPortRangeEnd   int    `json:"frontendPortRangeEnd"`
	BackendPort            int    `json:"backendPort"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureLoadBalancer is the Schema for the azureloadbalancers API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureLoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureLoadBalancerSpec `json:"spec,omitempty"`
	Status ASOStatus             `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureLoadBalancerList contains a list of AzureLoadBalancer
type AzureLoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureLoadBalancer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureLoadBalancer{}, &AzureLoadBalancerList{})
}
