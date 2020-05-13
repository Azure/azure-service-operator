// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzurePublicIPAddressSpec defines the desired state of AzurePublicIPAddress
type AzurePublicIPAddressSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location                 string `json:"location"`
	ResourceGroup            string `json:"resourceGroup"`
	PublicIPAllocationMethod string `json:"publicIPAllocationMethod"`
	IdleTimeoutInMinutes     int    `json:"idleTimeoutInMinutes"`
	PublicIPAddressVersion   string `json:"publicIPAddressVersion"`
	SkuName                  string `json:"skuName"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzurePublicIPAddress is the Schema for the azurepublicipaddresses API
// +kubebuilder:resource:shortName=pipa,path=azurepublicipaddress
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzurePublicIPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzurePublicIPAddressSpec `json:"spec,omitempty"`
	Status ASOStatus                `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzurePublicIPAddressList contains a list of AzurePublicIPAddress
type AzurePublicIPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzurePublicIPAddress `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzurePublicIPAddress{}, &AzurePublicIPAddressList{})
}
