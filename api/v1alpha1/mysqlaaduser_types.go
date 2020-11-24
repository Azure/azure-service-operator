// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLAADUserSpec defines the desired state of MySQLAADUser
type MySQLAADUserSpec struct {
	// +kubebuilder:validation:MinLength:1
	// +kubebuilder:validation:Required
	Server string `json:"server"`

	// +kubebuilder:validation:MinLength:1
	// +kubebuilder:validation:Required
	DBName string `json:"dbName"`

	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength:1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`

	// The roles assigned to the user. A user must have at least one role.
	// +kubebuilder:validation:MinLength:1
	// +kubebuilder:validation:Required
	Roles []string `json:"roles"`

	// Note: We current do not support arbitrary AAD users (although the MySQL API does).

	// ClientID is the client ID of the identity backing the user.
	// +kubebuilder:validation:Required
	ClientID string `json:"clientId,omitempty"`

	// optional
	Username string `json:"username,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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
