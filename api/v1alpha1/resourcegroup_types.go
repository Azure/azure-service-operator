// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ResourceGroupSpec defines the desired state of ResourceGroup
type ResourceGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location string `json:"location"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// ResourceGroup is the Schema for the resourcegroups API
// +kubebuilder:resource:shortName=rg,path=resourcegroups
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Provisioning",type="string",JSONPath=".status.provisioning"
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceGroupSpec `json:"spec,omitempty"`
	Status ASOStatus         `json:"status,omitempty"`
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

func (resourceGroup *ResourceGroup) IsBeingDeleted() bool {
	return !resourceGroup.ObjectMeta.DeletionTimestamp.IsZero()
}

func (resourceGroup *ResourceGroup) IsSubmitted() bool {
	return resourceGroup.Status.Provisioning || resourceGroup.Status.Provisioned

}

func (resourceGroup *ResourceGroup) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(resourceGroup.ObjectMeta.Finalizers, finalizerName)
}

func (resourceGroup *ResourceGroup) AddFinalizer(finalizerName string) {
	resourceGroup.ObjectMeta.Finalizers = append(resourceGroup.ObjectMeta.Finalizers, finalizerName)
}

func (resourceGroup *ResourceGroup) RemoveFinalizer(finalizerName string) {
	resourceGroup.ObjectMeta.Finalizers = helpers.RemoveString(resourceGroup.ObjectMeta.Finalizers, finalizerName)
}
