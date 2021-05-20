// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLUserSpec defines the desired state of MySqlUser
type MySQLUserSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Server string `json:"server"`

	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`

	// optional

	// The server-level roles assigned ot the user.
	Roles []string `json:"roles,omitempty"`

	// The database-level roles assigned to the user (keyed by
	// database name).
	DatabaseRoles map[string][]string `json:"databaseRoles,omitempty"`

	AdminSecret            string `json:"adminSecret,omitempty"`
	AdminSecretKeyVault    string `json:"adminSecretKeyVault,omitempty"`
	Username               string `json:"username,omitempty"`
	KeyVaultToStoreSecrets string `json:"keyVaultToStoreSecrets,omitempty"`
}

func (s MySQLUserSpec) GetAdminSecretName() string {
	adminSecretName := s.AdminSecret
	if len(adminSecretName) == 0 {
		adminSecretName = s.Server
	}

	return adminSecretName
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MySQLUser is the Schema for the mysqlusers API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLUserSpec `json:"spec,omitempty"`
	Status ASOStatus     `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLUserList contains a list of MySQLUser
type MySQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLUser{}, &MySQLUserList{})
}

// IsSubmitted checks if sqluser is provisioning
func (s *MySQLUser) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}
