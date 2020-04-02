// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureManagedDiskSpec defines the desired state of AzureManagedDisk
type AzureManagedDiskSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string `json:"location"`
	ResourceGroup string `json:"resourceGroup"`
	DiskSizeGB    int    `json:"diskSizeGB"`
	CreateOption  string `json:"createOption"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureManagedDisk is the Schema for the azuremanageddisks API
type AzureManagedDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureManagedDiskSpec `json:"spec,omitempty"`
	Status ASOStatus            `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureManagedDiskList contains a list of AzureManagedDisk
type AzureManagedDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureManagedDisk `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureManagedDisk{}, &AzureManagedDiskList{})
}
