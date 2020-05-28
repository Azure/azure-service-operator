// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureDiskSpec defines the desired state of AzureDisk
type AzureDiskSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string `json:"location"`
	ResourceGroup string `json:"resourceGroup"`
	CreateOption  string `json:"createOption"`
	DiskSizeGB    int32  `json:"diskSizeGB"`
}

// +kubebuilder:object:root=true

// AzureDisk is the Schema for the azuredisks API
type AzureDisk struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureDiskSpec `json:"spec,omitempty"`
	Status ASOStatus     `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureDiskList contains a list of AzureDisk
type AzureDiskList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureDisk `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureDisk{}, &AzureDiskList{})
}
