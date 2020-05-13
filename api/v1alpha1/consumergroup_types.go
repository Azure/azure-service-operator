// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ConsumerGroupSpec defines the desired state of ConsumerGroup
type ConsumerGroupSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroup     string `json:"resourceGroup,omitempty"`
	Namespace         string `json:"namespace,omitempty"`
	Eventhub          string `json:"eventHub,omitempty"`
	ConsumerGroupName string `json:"consumerGroupName,omitempty"` // optional, falls back to ObjectMeta.Name
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ConsumerGroup is the Schema for the consumergroups API
// +kubebuilder:resource:shortName=cg,path=consumergroup
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type ConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConsumerGroupSpec `json:"spec,omitempty"`
	Status ASOStatus         `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConsumerGroupList contains a list of ConsumerGroup
type ConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConsumerGroup `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConsumerGroup{}, &ConsumerGroupList{})
}

func init() {
	SchemeBuilder.Register(&ConsumerGroup{}, &ConsumerGroupList{})
}

func (consumerGroup *ConsumerGroup) IsSubmitted() bool {
	return consumerGroup.Status.Provisioning || consumerGroup.Status.Provisioned

}

func (consumerGroup *ConsumerGroup) HasFinalizer(finalizerName string) bool {
	return helpers.ContainsString(consumerGroup.ObjectMeta.Finalizers, finalizerName)
}

func (consumerGroup *ConsumerGroup) AddFinalizer(finalizerName string) {
	consumerGroup.ObjectMeta.Finalizers = append(consumerGroup.ObjectMeta.Finalizers, finalizerName)
}

func (consumerGroup *ConsumerGroup) RemoveFinalizer(finalizerName string) {
	consumerGroup.ObjectMeta.Finalizers = helpers.RemoveString(consumerGroup.ObjectMeta.Finalizers, finalizerName)
}
