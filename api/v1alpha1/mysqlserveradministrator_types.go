// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLServerAdministrator is the Schema for the mysqlserveradministrator API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:shortName=mysqladmin
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLServerAdministrator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MySQLServerAdministratorSpec `json:"spec,omitempty"`
	Status            ASOStatus                    `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
type MySQLServerAdministratorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLServerAdministrator `json:"items"`
}

type MySQLServerAdministratorSpec struct {

	// +kubebuilder:validation:Pattern=^[-\w\._\(\)]+$
	// +kubebuilder:validation:MinLength:1
	// +kubebuilder:validation:Required
	ResourceGroup string `json:"resourceGroup"`

	// +kubebuilder:validation:Required
	Server string `json:"server"`

	// +kubebuilder:validation:Required
	//AdministratorType: The type of administrator.
	AdministratorType MySQLServerAdministratorType `json:"administratorType"`

	//Login: The server administrator login account name.
	//For example: "myuser@microsoft.com" might be the login if specifying an AAD user. "my-mi" might be the name of a managed identity
	// +kubebuilder:validation:Required
	Login string `json:"login"`

	//Sid: The server administrator Sid (Secure ID). If creating an AAD user, this is the OID of the entity in AAD.
	// +kubebuilder:validation:Required
	Sid string `json:"sid"`

	//TenantId: The server Active Directory Administrator tenant id.
	// +kubebuilder:validation:Required
	TenantId string `json:"tenantId"`
}

// +kubebuilder:validation:Enum={"ActiveDirectory"}
type MySQLServerAdministratorType string

const MySQLServerAdministratorTypeActiveDirectory = MySQLServerAdministratorType("ActiveDirectory")

func init() {
	SchemeBuilder.Register(&MySQLServerAdministrator{}, &MySQLServerAdministratorList{})
}
