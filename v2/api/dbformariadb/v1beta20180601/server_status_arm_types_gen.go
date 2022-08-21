// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

type Server_STATUSARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_STATUSARM `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_STATUSARM `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ServerProperties_STATUSARM struct {
	// AdministratorLogin: The administrator's login name of a server. Can only be specified when the server is being created
	// (and is required for creation).
	AdministratorLogin *string `json:"administratorLogin,omitempty"`

	// EarliestRestoreDate: Earliest restore point creation time (ISO8601 format)
	EarliestRestoreDate *string `json:"earliestRestoreDate,omitempty"`

	// FullyQualifiedDomainName: The fully qualified domain name of a server.
	FullyQualifiedDomainName *string `json:"fullyQualifiedDomainName,omitempty"`

	// MasterServerId: The master server id of a replica server.
	MasterServerId *string `json:"masterServerId,omitempty"`

	// MinimalTlsVersion: Enforce a minimal Tls version for the server.
	MinimalTlsVersion *MinimalTlsVersion_STATUS `json:"minimalTlsVersion,omitempty"`

	// PrivateEndpointConnections: List of private endpoint connections on a server
	PrivateEndpointConnections []ServerPrivateEndpointConnection_STATUSARM `json:"privateEndpointConnections,omitempty"`

	// PublicNetworkAccess: Whether or not public network access is allowed for this server. Value is optional but if passed
	// in, must be 'Enabled' or 'Disabled'
	PublicNetworkAccess *PublicNetworkAccess_STATUS `json:"publicNetworkAccess,omitempty"`

	// ReplicaCapacity: The maximum number of replicas that a master server can have.
	ReplicaCapacity *int `json:"replicaCapacity,omitempty"`

	// ReplicationRole: The replication role of the server.
	ReplicationRole *string `json:"replicationRole,omitempty"`

	// SslEnforcement: Enable ssl enforcement or not when connect to server.
	SslEnforcement *SslEnforcement_STATUS `json:"sslEnforcement,omitempty"`

	// StorageProfile: Storage profile of a server.
	StorageProfile *StorageProfile_STATUSARM `json:"storageProfile,omitempty"`

	// UserVisibleState: A state of a server that is visible to user.
	UserVisibleState *ServerProperties_UserVisibleState_STATUS `json:"userVisibleState,omitempty"`

	// Version: Server version.
	Version *ServerVersion_STATUS `json:"version,omitempty"`
}

type Sku_STATUSARM struct {
	// Capacity: The scale up/out capacity, representing server's compute units.
	Capacity *int `json:"capacity,omitempty"`

	// Family: The family of hardware.
	Family *string `json:"family,omitempty"`

	// Name: The name of the sku, typically, tier + family + cores, e.g. B_Gen4_1, GP_Gen5_8.
	Name *string `json:"name,omitempty"`

	// Size: The size code, to be interpreted by resource as appropriate.
	Size *string `json:"size,omitempty"`

	// Tier: The tier of the particular SKU, e.g. Basic.
	Tier *Sku_Tier_STATUS `json:"tier,omitempty"`
}

type ServerPrivateEndpointConnection_STATUSARM struct {
	// Id: Resource Id of the private endpoint connection.
	Id *string `json:"id,omitempty"`

	// Properties: Private endpoint connection properties
	Properties *ServerPrivateEndpointConnectionProperties_STATUSARM `json:"properties,omitempty"`
}

type Sku_Tier_STATUS string

const (
	Sku_Tier_Basic_STATUS           = Sku_Tier_STATUS("Basic")
	Sku_Tier_GeneralPurpose_STATUS  = Sku_Tier_STATUS("GeneralPurpose")
	Sku_Tier_MemoryOptimized_STATUS = Sku_Tier_STATUS("MemoryOptimized")
)

type StorageProfile_STATUSARM struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: Enable Geo-redundant or not for server backup.
	GeoRedundantBackup *StorageProfile_GeoRedundantBackup_STATUS `json:"geoRedundantBackup,omitempty"`

	// StorageAutogrow: Enable Storage Auto Grow.
	StorageAutogrow *StorageProfile_StorageAutogrow_STATUS `json:"storageAutogrow,omitempty"`

	// StorageMB: Max storage allowed for a server.
	StorageMB *int `json:"storageMB,omitempty"`
}

type ServerPrivateEndpointConnectionProperties_STATUSARM struct {
	// PrivateEndpoint: Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty_STATUSARM `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *ServerPrivateLinkServiceConnectionStateProperty_STATUSARM `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: State of the private endpoint connection.
	ProvisioningState *ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type PrivateEndpointProperty_STATUSARM struct {
	// Id: Resource id of the private endpoint.
	Id *string `json:"id,omitempty"`
}

type ServerPrivateLinkServiceConnectionStateProperty_STATUSARM struct {
	// ActionsRequired: The actions required for private link service connection.
	ActionsRequired *ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS `json:"actionsRequired,omitempty"`

	// Description: The private link service connection description.
	Description *string `json:"description,omitempty"`

	// Status: The private link service connection status.
	Status *ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS `json:"status,omitempty"`
}
