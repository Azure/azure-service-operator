/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PostgreSQLUserSpec defines the desired state of PostgreSqlUser
type PostgreSQLUserSpec struct {
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

	// The databases of which this user should inherit the ownership role.
	OwnedDatabases []string `json:"ownedDatabases,omitempty"`

	// This database is meant as a default database for users and applications to connect to.
	// The ASO needs an existing database to connect to but does not run database specific queries.
	MaintenanceDatabase    string `json:"maintenanceDatabase,omitempty"`
	AdminSecret            string `json:"adminSecret,omitempty"`
	AdminSecretKeyVault    string `json:"adminSecretKeyVault,omitempty"`
	Username               string `json:"username,omitempty"`
	KeyVaultToStoreSecrets string `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PostgreSQLUser is the Schema for the postgresqlusers API
// +kubebuilder:resource:shortName=psqlu
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type PostgreSQLUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgreSQLUserSpec `json:"spec,omitempty"`
	Status ASOStatus          `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgreSQLUserList contains a list of PostgreSQLUser
type PostgreSQLUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgreSQLUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PostgreSQLUser{}, &PostgreSQLUserList{})
}

// IsSubmitted checks if psqluser is provisioning
func (s *PostgreSQLUser) IsSubmitted() bool {
	return s.Status.Provisioning || s.Status.Provisioned
}
