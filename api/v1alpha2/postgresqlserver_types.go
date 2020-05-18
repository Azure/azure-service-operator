// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PostgreSQLServerSpec defines the desired state of PostgreSQLServer
type PostgreSQLServerSpec struct {
	Location               string              `json:"location"`
	ResourceGroup          string              `json:"resourceGroup,omitempty"`
	Sku                    AzureDBsSQLSku      `json:"sku,omitempty"`
	ServerVersion          ServerVersion       `json:"serverVersion,omitempty"`
	SSLEnforcement         SslEnforcementEnum  `json:"sslEnforcement,omitempty"`
	KeyVaultToStoreSecrets string              `json:"keyVaultToStoreSecrets,omitempty"`
	CreateMode             string              `json:"createMode,omitempty"`
	StorageProfile         *PSQLStorageProfile `json:"storageProfile,omitempty"`
	ReplicaProperties      ReplicaProperties   `json:"replicaProperties,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// PostgreSQLServer is the Schema for the postgresqlservers API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type PostgreSQLServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PostgreSQLServerSpec `json:"spec,omitempty"`
	Status ASOStatus            `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PostgreSQLServerList contains a list of PostgreSQLServer
type PostgreSQLServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PostgreSQLServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PostgreSQLServer{}, &PostgreSQLServerList{})
}
