// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSqlActionSpec defines the desired state of AzureSqlAction
type AzureSqlActionSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	ResourceGroup string `json:"resourcegroup"`
	ActionName    string `json:"actionname"`
	ServerName    string `json:"servername"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureSqlAction is the Schema for the azuresqlactions API
type AzureSqlAction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSqlActionSpec `json:"spec,omitempty"`
	Status ASOStatus          `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSqlActionList contains a list of AzureSqlAction
type AzureSqlActionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSqlAction `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSqlAction{}, &AzureSqlActionList{})
}

func (s *AzureSqlAction) IsSubmitted() bool {
	return s.Status.Provisioned || s.Status.Provisioning
}
