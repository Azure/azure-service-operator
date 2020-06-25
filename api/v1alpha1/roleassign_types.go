// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RoleAssignSpec defines the desired state of RoleAssign
type RoleAssignSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	PrincipalID string `json:"principalId"`
	Role        string `json:"role"`
	Scope       string `json:"scope"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
// RoleAssign is the Schema for the roleassigns API
type RoleAssign struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RoleAssignSpec `json:"spec,omitempty"`
	Status ASOStatus      `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleAssignList contains a list of RoleAssign
type RoleAssignList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssign `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RoleAssign{}, &RoleAssignList{})
}
