// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureVMScaleSetSpec defines the desired state of AzureVMScaleSet
type AzureVMScaleSetSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location               string `json:"location"`
	ResourceGroup          string `json:"resourceGroup"`
	VMSize                 string `json:"vmSize"`
	Capacity               int    `json:"capacity"`
	OSType                 OSType `json:"osType"`
	AdminUserName          string `json:"adminUserName"`
	SSHPublicKeyData       string `json:"sshPublicKeyData,omitempty"`
	PlatformImageURN       string `json:"platformImageURN"`
	VirtualNetworkName     string `json:"virtualNetworkName"`
	SubnetName             string `json:"subnetName"`
	LoadBalancerName       string `json:"loadBalancerName"`
	BackendAddressPoolName string `json:"backendAddressPoolName"`
	InboundNatPoolName     string `json:"inboundNatPoolName"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureVMScaleSet is the Schema for the azurevmscalesets API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureVMScaleSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureVMScaleSetSpec `json:"spec,omitempty"`
	Status ASOStatus           `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureVMScaleSetList contains a list of AzureVMScaleSet
type AzureVMScaleSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureVMScaleSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureVMScaleSet{}, &AzureVMScaleSetList{})
}
