// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AzureSQLUserSpec defines the desired state of SqlUser
type AzureSQLUserSpec struct {
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

	// optional
	SubscriptionID         string   `json:"subscriptionId,omitempty"`
	AdminSecret            string   `json:"adminSecret,omitempty"`
	AdminSecretKeyVault    string   `json:"adminSecretKeyVault,omitempty"`
	Username               string   `json:"username,omitempty"`
	KeyVaultToStoreSecrets string   `json:"keyVaultToStoreSecrets,omitempty"`
	KeyVaultSecretPrefix   string   `json:"keyVaultSecretPrefix,omitempty"` // TODO: Remove this in a future version?
	KeyVaultSecretFormats  []string `json:"keyVaultSecretFormats,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// AzureSQLUser is the Schema for the sqlusers API
// +kubebuilder:resource:shortName=asqlu
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type AzureSQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AzureSQLUserSpec `json:"spec,omitempty"`
	Status ASOStatus        `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AzureSQLUserList contains a list of SqlUser
type AzureSQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AzureSQLUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&AzureSQLUser{}, &AzureSQLUserList{})
}

// IsSubmitted checks if sqluser is provisioning
func (s *AzureSQLUser) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}
