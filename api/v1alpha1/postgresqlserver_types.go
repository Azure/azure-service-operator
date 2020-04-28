// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PostgreSQLServerSpec defines the desired state of PostgreSQLServer
type PostgreSQLServerSpec struct {
	Location               string             `json:"location"`
	ResourceGroup          string             `json:"resourceGroup,omitempty"`
	Sku                    AzureDBsSQLSku     `json:"sku,omitempty"`
	ServerVersion          ServerVersion      `json:"serverVersion,omitempty"`
	SSLEnforcement         SslEnforcementEnum `json:"sslEnforcement,omitempty"`
	KeyVaultToStoreSecrets string             `json:"keyVaultToStoreSecrets,omitempty"`
	CreateMode             string             `json:"createMode,omitempty"`
	ReplicaProperties      ReplicaProperties  `json:"replicaProperties,omitempty"`
}

type AzureDBsSQLSku struct {
	// Name - The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
	Name string `json:"name,omitempty"`
	// Tier - The tier of the particular SKU, e.g. Basic. Possible values include: 'Basic', 'GeneralPurpose', 'MemoryOptimized'
	Tier SkuTier `json:"tier,omitempty"`
	// Capacity - The scale up/out capacity, representing server's compute units.
	Capacity int32 `json:"capacity,omitempty"`
	// Size - The size code, to be interpreted by resource as appropriate.
	Size string `json:"size,omitempty"`
	// Family - The family of hardware.
	Family string `json:"family,omitempty"`
}

// ServerVersion enumerates the values for server version.
type ServerVersion string

const (
	// NineFullStopFive ...
	NineFullStopFive ServerVersion = "9.5"
	// NineFullStopSix ...
	NineFullStopSix ServerVersion = "9.6"
	// OneOne ...
	OneOne ServerVersion = "11"
	// OneZero ...
	OneZero ServerVersion = "10"
	// OneZeroFullStopTwo ...
	OneZeroFullStopTwo ServerVersion = "10.2"
	// OneZeroFullStopZero ...
	OneZeroFullStopZero ServerVersion = "10.0"
)

type SkuTier string

const (
	// Basic ...
	PSQLBasic SkuTier = "Basic"
	// GeneralPurpose ...
	PSQLGeneralPurpose SkuTier = "GeneralPurpose"
	// MemoryOptimized ...
	PSQLMemoryOptimized SkuTier = "MemoryOptimized"
)

type SslEnforcementEnum string

const (
	// SslEnforcementEnumDisabled ...
	SslEnforcementEnumDisabled SslEnforcementEnum = "Disabled"
	// SslEnforcementEnumEnabled ...
	SslEnforcementEnumEnabled SslEnforcementEnum = "Enabled"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

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
