/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/k8s-infra/pkg/zips"
)

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	APIVersion string            `json:"apiVersion,omitempty"`
	Location   string            `json:"location,omitempty"`
	ManagedBy  string            `json:"managedBy,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
}

// ResourceGroupStatus defines the observed state of ResourceGroup
type ResourceGroupStatus struct {
	ID                string `json:"id,omitempty"`
	ProvisioningState string `json:"provisioningState,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

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

func (*ResourceGroup) Hub() {
	fmt.Println("Hub was called!")
}

func (rg *ResourceGroup) ToResource() (zips.Resource, error) {
	res := zips.Resource{
		ID:                rg.Status.ID,
		Type:              "Microsoft.Resources/resourceGroups",
		Name:              rg.Name,
		APIVersion:        rg.Spec.APIVersion,
		Location:          rg.Spec.Location,
		Tags:              rg.Spec.Tags,
		ManagedBy:         rg.Spec.ManagedBy,
		ProvisioningState: rg.Status.ProvisioningState,
	}

	return *res.SetAnnotations(rg.Annotations), nil
}

func (rg *ResourceGroup) FromResource(res zips.Resource) error {
	rg.Status.ID = res.ID
	rg.Status.ProvisioningState = res.ProvisioningState
	rg.Spec.Tags = res.Tags
	rg.Spec.ManagedBy = res.ManagedBy
	return nil
}

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
