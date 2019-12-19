/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/k8s-infra/api/azmetav1"
	"github.com/Azure/k8s-infra/pkg/zips"
)

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	azmetav1.TrackedResourceSpec `json:",inline"`
}

// ResourceGroupStatus defines the observed state of ResourceGroup
type ResourceGroupStatus struct {
	azmetav1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// ResourceGroup is the Schema for the resourcegroups API
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceGroupSpec   `json:"spec,omitempty"`
	Status ResourceGroupStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroupList contains a list of ResourceGroup
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

func (rg *ResourceGroup) ToResource() (zips.Resource, error) {
	return zips.Resource{
		ID:         rg.Status.ID,
		Name:       rg.Name,
		Type:       "Microsoft.Resources/resourceGroup",
		Location:   rg.Spec.Location,
		APIVersion: "2018-05-01",
	}, nil
}

func (rg *ResourceGroup) FromResource(zips.Resource) (runtime.Object, error) {
	// TODO: fill in with appropriate logic
	return rg, nil
}

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
	SchemeBuilder.Register(&unstructured.Unstructured{}, &unstructured.UnstructuredList{})
}
