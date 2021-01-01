// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ContainerRegistrySku the SKU of the Container Registry
type ContainerRegistrySku struct {
	// Name - The SKU name. Required for account creation; optional for update.
	// Possible values include: 'Basic', 'Classic', 'Premium' or 'Standard'
	Name string `json:"name,omitempty"`
}

// AzureContainerRegistrySpec defines the desired state of AzureContainerRegistry
type AzureContainerRegistrySpec struct {
	AdminUserEnabled bool                 `json:"adminUserEnabled"`
	Location         string               `json:"location"`
	ResourceGroup    string               `json:"resourceGroup"`
	Sku              ContainerRegistrySku `json:"sku,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureContainerRegistry is the Schema for the azurecontainerregistries API
// +kubebuilder:resource:shortName=acr
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureContainerRegistry struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureContainerRegistrySpec `json:"spec,omitempty"`
	Status ASOStatus                  `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureContainerRegistryList contains a list of AzureContainerRegistry
type AzureContainerRegistryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureContainerRegistry `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureContainerRegistry{}, &AzureContainerRegistryList{})
}
