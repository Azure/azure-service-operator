// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	kvops "github.com/Azure/azure-sdk-for-go/services/keyvault/v7.0/keyvault"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// KeyVaultKeySpec defines the desired state of KeyVaultKey
type KeyVaultKeySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Location      string                      `json:"location"`
	ResourceGroup string                      `json:"resourceGroup"`
	KeyVault      string                      `json:"keyVault,omitempty"`
	KeySize       int32                       `json:"keySize,omitempty"`
	Type          kvops.JSONWebKeyType        `json:"type,omitempty"`
	Operations    []kvops.JSONWebKeyOperation `json:"operations,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// KeyVaultKey is the Schema for the keyvaultkeys API
// +kubebuilder:resource:shortName=kvk,path=keyvaultkey
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type KeyVaultKey struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KeyVaultKeySpec `json:"spec,omitempty"`
	Status ASOStatus       `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// KeyVaultKeyList contains a list of KeyVaultKey
type KeyVaultKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KeyVaultKey `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KeyVaultKey{}, &KeyVaultKeyList{})
}
