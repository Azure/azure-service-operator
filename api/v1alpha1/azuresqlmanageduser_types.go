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

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Server string `json:"server"`

	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	DbName string `json:"dbName"`

	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`

	// +kubebuilder:validation:Required
	Roles []string `json:"roles"`

	SubscriptionID string `json:"subscriptionId,omitempty"`

	ManagedIdentityName     string `json:"managedIdentityName,omitempty"`
	ManagedIdentityClientId string `json:"managedIdentityClientId"`
	KeyVaultSecretPrefix    string `json:"keyVaultSecretPrefix,omitempty"`
	KeyVaultToStoreSecrets  string `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// AzureSQLManagedUser is the Schema for the azuresqlmanagedusers API
// +kubebuilder:resource:shortName=asqlmu
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
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
