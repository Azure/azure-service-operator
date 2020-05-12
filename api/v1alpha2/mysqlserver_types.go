// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

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

type StorageProfile struct {
	// BackupRetentionDays - Backup retention days for the server.
	BackupRetentionDays *int32 `json:"backupRetentionDays,omitempty"`
	// GeoRedundantBackup - Enable Geo-redundant or not for server backup. Possible values include: 'Enabled', 'Disabled'
	GeoRedundantBackup mysql.GeoRedundantBackup `json:"geoRedundantBackup,omitempty"`
	// StorageMB - Max storage allowed for a server.
	StorageMB *int32 `json:"storageMB,omitempty"`
	// StorageAutogrow - Enable Storage Auto Grow. Possible values include: 'StorageAutogrowEnabled', 'StorageAutogrowDisabled'
	StorageAutogrow mysql.StorageAutogrow `json:"storageAutogrow,omitempty"`
}

type ReplicaProperties struct {
	SourceServerId string `json:"sourceServerId,omitempty"`
}

// MySQLServerSpec defines the desired state of MySQLServer
type MySQLServerSpec struct {
	Location               string             `json:"location"`
	ResourceGroup          string             `json:"resourceGroup,omitempty"`
	Sku                    AzureDBsSQLSku     `json:"sku,omitempty"`
	ServerVersion          ServerVersion      `json:"serverVersion,omitempty"`
	SSLEnforcement         SslEnforcementEnum `json:"sslEnforcement,omitempty"`
	CreateMode             string             `json:"createMode,omitempty"`
	ReplicaProperties      ReplicaProperties  `json:"replicaProperties,omitempty"`
	StorageProfile         *StorageProfile    `json:"storageProfile,omitempty"`
	KeyVaultToStoreSecrets string             `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// MySQLServer is the Schema for the mysqlservers API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLServerSpec `json:"spec,omitempty"`
	Status ASOStatus       `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MySQLServerList contains a list of MySQLServer
type MySQLServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MySQLServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MySQLServer{}, &MySQLServerList{})
}

func NewDefaultMySQLServer(name, resourceGroup, location string) *MySQLServer {
	return &MySQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		Spec: MySQLServerSpec{
			Location:      location,
			ResourceGroup: resourceGroup,
			Sku: AzureDBsSQLSku{
				Name:     "GP_Gen5_4",
				Tier:     SkuTier("GeneralPurpose"),
				Family:   "Gen5",
				Size:     "51200",
				Capacity: 4,
			},
			ServerVersion:  ServerVersion("8.0"),
			SSLEnforcement: SslEnforcementEnumEnabled,
			CreateMode:     "Default",
		},
	}
}

func NewReplicaMySQLServer(name, resourceGroup, location string, sourceserverid string) *MySQLServer {
	return &MySQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: "default",
		},
		Spec: MySQLServerSpec{
			Location:      location,
			ResourceGroup: resourceGroup,
			CreateMode:    "Replica",
			ReplicaProperties: ReplicaProperties{
				SourceServerId: sourceserverid,
			},
		},
	}
}
