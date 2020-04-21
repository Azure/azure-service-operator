// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSQLManagedUserSpec defines the desired state of AzureSQLManagedUser
type AzureSQLManagedUserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Server                  string   `json:"server"`
	DbName                  string   `json:"dbName"`
	ResourceGroup           string   `json:"resourceGroup,omitempty"`
	Roles                   []string `json:"roles"`
	ManagedIdentityName     string   `json:"managedIdentityName,omitempty"`
	ManagedIdentityObjectId string   `json:"managedIdentityObjectId"`
	KeyVaultToStoreSecrets  string   `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AzureSQLManagedUser is the Schema for the azuresqlmanagedusers API
type AzureSQLManagedUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSQLManagedUserSpec `json:"spec,omitempty"`
	Status ASOStatus               `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSQLManagedUserList contains a list of AzureSQLManagedUser
type AzureSQLManagedUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSQLManagedUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSQLManagedUser{}, &AzureSQLManagedUserList{})
}
