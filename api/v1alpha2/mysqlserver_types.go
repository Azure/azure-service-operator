// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// MySQLServerSpec defines the desired state of MySQLServer
type MySQLServerSpec struct {
	Location               string               `json:"location"`
	ResourceGroup          string               `json:"resourceGroup,omitempty"`
	Sku                    AzureDBsSQLSku       `json:"sku,omitempty"`
	ServerVersion          ServerVersion        `json:"serverVersion,omitempty"`
	SSLEnforcement         SslEnforcementEnum   `json:"sslEnforcement,omitempty"`
	CreateMode             string               `json:"createMode,omitempty"`
	ReplicaProperties      ReplicaProperties    `json:"replicaProperties,omitempty"`
	StorageProfile         *MySQLStorageProfile `json:"storageProfile,omitempty"`
	KeyVaultToStoreSecrets string               `json:"keyVaultToStoreSecrets,omitempty"`
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
			StorageProfile: &MySQLStorageProfile{
				BackupRetentionDays: to.Int32Ptr(10),
				GeoRedundantBackup:  "Disabled",
				StorageMB:           to.Int32Ptr(5120),
				StorageAutogrow:     "Disabled",
			},
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
