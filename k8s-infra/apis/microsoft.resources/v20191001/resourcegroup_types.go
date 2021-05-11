/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191001

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type (
	// ResourceGroupSpec defines the desired state of ResourceGroup
	ResourceGroupSpec struct {
		// Location is the Azure location for the group (eg westus2, southcentralus, etc...)
		Location string `json:"location,omitempty"`

		// ManagedBy is the management group responsible for managing this group
		// +optional
		ManagedBy string `json:"managedBy,omitempty"`

		// Tags are user defined key value pairs
		// +optional
		Tags map[string]string `json:"tags,omitempty"`
	}

	// ResourceGroupStatus defines the observed state of ResourceGroup
	ResourceGroupStatus struct {
		ID                string `json:"id,omitempty"`
		ProvisioningState string `json:"provisioningState,omitempty"`
	}

	// +kubebuilder:object:root=true

	// ResourceGroup is the Schema for the resourcegroups API
	ResourceGroup struct {
		metav1.TypeMeta   `json:",inline"`
		metav1.ObjectMeta `json:"metadata,omitempty"`

		Spec   ResourceGroupSpec   `json:"spec,omitempty"`
		Status ResourceGroupStatus `json:"status,omitempty"`
	}

	// +kubebuilder:object:root=true

	// ResourceGroupList contains a list of ResourceGroup
	ResourceGroupList struct {
		metav1.TypeMeta `json:",inline"`
		metav1.ListMeta `json:"metadata,omitempty"`
		Items           []ResourceGroup `json:"items"`
	}
)

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
