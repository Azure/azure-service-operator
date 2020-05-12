// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package v1alpha2

import (
	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

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

// MySQLServerSpec defines the desired state of MySQLServer
type MySQLServerSpec struct {
	Location               string                      `json:"location"`
	ResourceGroup          string                      `json:"resourceGroup,omitempty"`
	Sku                    v1alpha1.AzureDBsSQLSku     `json:"sku,omitempty"`
	ServerVersion          v1alpha1.ServerVersion      `json:"serverVersion,omitempty"`
	SSLEnforcement         v1alpha1.SslEnforcementEnum `json:"sslEnforcement,omitempty"`
	CreateMode             string                      `json:"createMode,omitempty"`
	ReplicaProperties      v1alpha1.ReplicaProperties  `json:"replicaProperties,omitempty"`
	StorageProfile         *StorageProfile             `json:"storageProfile,omitempty"`
	KeyVaultToStoreSecrets string                      `json:"keyVaultToStoreSecrets,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// MySQLServer is the Schema for the mysqlservers API
// +kubebuilder:printcolumn:name="Provisioned",type="string",JSONPath=".status.provisioned"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.message"
type MySQLServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MySQLServerSpec    `json:"spec,omitempty"`
	Status v1alpha1.ASOStatus `json:"status,omitempty"`
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
			Sku: v1alpha1.AzureDBsSQLSku{
				Name:     "GP_Gen5_4",
				Tier:     v1alpha1.SkuTier("GeneralPurpose"),
				Family:   "Gen5",
				Size:     "51200",
				Capacity: 4,
			},
			ServerVersion:  v1alpha1.ServerVersion("8.0"),
			SSLEnforcement: v1alpha1.SslEnforcementEnumEnabled,
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
			ReplicaProperties: v1alpha1.ReplicaProperties{
				SourceServerId: sourceserverid,
			},
		},
	}
}
