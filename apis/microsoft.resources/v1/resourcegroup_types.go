/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/k8s-infra/pkg/zips"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type AzRGSpec struct {
	Name      string            `json:"name,omitempty"`
	Location  string            `json:"location,omitempty"`
	ManagedBy string            `json:"managedBy,omitempty"`
	Tags      map[string]string `json:"tags,omitempty"`
}

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Foo is an example field of ResourceGroup. Edit ResourceGroup_types.go to remove/update
	APIVersion string            `json:"apiVersion,omitempty"`
	Location   string            `json:"location,omitempty"`
	ManagedBy  string            `json:"managedBy,omitempty"`
	Tags       map[string]string `json:"tags,omitempty"`
}

// ResourceGroupStatus defines the observed state of ResourceGroup
type ResourceGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
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

func (rg *ResourceGroup) ToResource() zips.Resource {
	return zips.Resource{
		ID:         rg.Status.ID,
		Type:       "Microsoft.Resources/resourceGroups",
		Name:       rg.Name,
		APIVersion: rg.Spec.APIVersion,
		Location:   rg.Spec.Location,
		Tags:       rg.Spec.Tags,
		ManagedBy:  rg.Spec.ManagedBy,
	}
}

func (rg *ResourceGroup) FromResource(res zips.Resource) runtime.Object {
	rg.Status.ID = res.ID
	rg.Status.ProvisioningState = res.ProvisioningState
	rg.Spec.Tags = res.Tags
	rg.Spec.ManagedBy = res.ManagedBy
	return rg
}

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
