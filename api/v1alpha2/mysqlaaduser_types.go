// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLAADUserSpec defines the desired state of MySQLAADUser
type MySQLAADUserSpec struct {
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	Server string `json:"server"`

	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`

	// AAD ID is the ID of the user in Azure Active Directory.
	// When creating a user for a managed identity this must be the client id (sometimes called app id) of the managed identity.
	// When creating a user for a "normal" (non-managed identity) user or group, this is the OID of the user or group.
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Required
	AADID string `json:"aadId,omitempty"`

	// optional

	// The server-level roles assigned to the user.
	Roles []string `json:"roles"`

	// The database-level roles assigned to the user (keyed by
	// database name).
	DatabaseRoles map[string][]string `json:"databaseRoles,omitempty"`

	Username string `json:"username,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MySQLAADUser is the Schema for an AAD user for MySQL
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLAADUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLAADUserSpec `json:"spec,omitempty"`
	Status ASOStatus        `json:"status,omitempty"`
}

func (u MySQLAADUser) Username() string {
	username := u.Name
	if u.Spec.Username != "" {
		username = u.Spec.Username
	}

	return username
}

// +kubebuilder:object:root=true

// MySQLAADUserList contains a list of MySQLAADUser
type MySQLAADUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLAADUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLAADUser{}, &MySQLAADUserList{})
}
