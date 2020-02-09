/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v20191001

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	v1 "github.com/Azure/k8s-infra/apis/microsoft.resources/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

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
type ResourceGroupStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ID                string `json:"id,omitempty"`
	ProvisioningState string `json:"provisioningState,omitempty"`
	Type              string `json:"type,omitempty"`
}

// +kubebuilder:object:root=true

// ResourceGroup is the Schema for the resourcegroups API
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceGroupSpec   `json:"spec,omitempty"`
	Status ResourceGroupStatus `json:"status,omitempty"`
}

func (rg *ResourceGroup) ConvertTo(dstRaw conversion.Hub) error {
	to := dstRaw.(*v1.ResourceGroup)
	to.ObjectMeta = rg.ObjectMeta
	to.Spec.APIVersion = "2019-10-01"
	to.Spec.Location = rg.Spec.Location
	to.Spec.ManagedBy = rg.Spec.ManagedBy
	to.Spec.Tags = rg.Spec.Tags
	to.Status.ID = rg.Status.ID
	to.Status.ProvisioningState = rg.Status.ProvisioningState
	return nil
}

func (rg *ResourceGroup) ConvertFrom(src conversion.Hub) error {
	from := src.(*v1.ResourceGroup)
	rg.ObjectMeta = from.ObjectMeta
	rg.Spec.Location = from.Spec.Location
	rg.Spec.ManagedBy = from.Spec.ManagedBy
	rg.Spec.Tags = from.Spec.Tags
	rg.Status.ID = from.Status.ID
	rg.Status.ProvisioningState = from.Status.ProvisioningState
	return nil
}

// +kubebuilder:object:root=true

// ResourceGroupList contains a list of ResourceGroup
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
