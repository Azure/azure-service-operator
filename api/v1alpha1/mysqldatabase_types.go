// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLDatabaseSpec defines the desired state of MySQLDatabase
type MySQLDatabaseSpec struct {
	ResourceGroup string `json:"resourceGroup,omitempty"`
	Server        string `json:"server,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MySQLDatabase is the Schema for the mysqldatabases API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLDatabaseSpec `json:"spec,omitempty"`
	Status ASOStatus         `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLDatabaseList contains a list of MySQLDatabase
type MySQLDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLDatabase `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLDatabase{}, &MySQLDatabaseList{})
}
