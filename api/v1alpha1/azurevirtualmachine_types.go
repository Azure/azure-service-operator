// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureVirtualMachineSpec defines the desired state of AzureVirtualMachine
type AzureVirtualMachineSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location             string `json:"location"`
	ResourceGroup        string `json:"resourceGroup"`
	VMSize               string `json:"vmSize"`
	OSType               OSType `json:"osType"`
	AdminUserName        string `json:"adminUserName"`
	SSHPublicKeyData     string `json:"sshPublicKeyData,omitempty"`
	NetworkInterfaceName string `json:"networkInterfaceName"`
	PlatformImageURN     string `json:"platformImageURN"`
}

type OSType string

const (
	// Windows ...
	Windows OSType = "Windows"
	// Linux ...
	Linux OSType = "Linux"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureVirtualMachine is the Schema for the azurevirtualmachines API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureVirtualMachine struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureVirtualMachineSpec `json:"spec,omitempty"`
	Status ASOStatus               `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureVirtualMachineList contains a list of AzureVirtualMachine
type AzureVirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureVirtualMachine `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureVirtualMachine{}, &AzureVirtualMachineList{})
}
