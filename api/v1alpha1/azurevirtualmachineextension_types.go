// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureVirtualMachineExtensionSpec defines the desired state of AzureVirtualMachineExtension
type AzureVirtualMachineExtensionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location                string `json:"location"`
	ResourceGroup           string `json:"resourceGroup"`
	VMName                  string `json:"vmName"`
	AutoUpgradeMinorVersion bool   `json:"autoUpgradeMinorVersion"`
	ForceUpdateTag          string `json:"forceUpdateTag"`
	Publisher               string `json:"publisher"`
	TypeName                string `json:"typeName"`
	TypeHandlerVersion      string `json:"typeHandlerVersion"`
	Settings                string `json:"settings,omitempty"`
	ProtectedSettings       string `json:"protectedSettings,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureVirtualMachineExtension is the Schema for the azurevirtualmachineextensions API
// +kubebuilder:resource:shortName=vmext,path=azurevirtualmachineextension
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureVirtualMachineExtension struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureVirtualMachineExtensionSpec `json:"spec,omitempty"`
	Status ASOStatus                        `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureVirtualMachineExtensionList contains a list of AzureVirtualMachineExtension
type AzureVirtualMachineExtensionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureVirtualMachineExtension `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureVirtualMachineExtension{}, &AzureVirtualMachineExtensionList{})
}
