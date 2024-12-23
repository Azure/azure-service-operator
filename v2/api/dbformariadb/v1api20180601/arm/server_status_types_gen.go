// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Represents a server.
type Server_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the server.
	Properties *ServerProperties_STATUS `json:"properties,omitempty"`

	// Sku: The SKU (pricing tier) of the server.
	Sku *Sku_STATUS `json:"sku,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// The properties of a server.
type ServerProperties_STATUS struct {
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
	PrivateEndpointConnections []ServerPrivateEndpointConnection_STATUS `json:"privateEndpointConnections,omitempty"`

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
	StorageProfile *StorageProfile_STATUS `json:"storageProfile,omitempty"`

	// UserVisibleState: A state of a server that is visible to user.
	UserVisibleState *ServerProperties_UserVisibleState_STATUS `json:"userVisibleState,omitempty"`

	// Version: Server version.
	Version *ServerVersion_STATUS `json:"version,omitempty"`
}

// Billing information related properties of a server.
type Sku_STATUS struct {
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

// Enforce a minimal Tls version for the server.
type MinimalTlsVersion_STATUS string

const (
	MinimalTlsVersion_STATUS_TLS1_0                 = MinimalTlsVersion_STATUS("TLS1_0")
	MinimalTlsVersion_STATUS_TLS1_1                 = MinimalTlsVersion_STATUS("TLS1_1")
	MinimalTlsVersion_STATUS_TLS1_2                 = MinimalTlsVersion_STATUS("TLS1_2")
	MinimalTlsVersion_STATUS_TLSEnforcementDisabled = MinimalTlsVersion_STATUS("TLSEnforcementDisabled")
)

// Mapping from string to MinimalTlsVersion_STATUS
var minimalTlsVersion_STATUS_Values = map[string]MinimalTlsVersion_STATUS{
	"tls1_0":                 MinimalTlsVersion_STATUS_TLS1_0,
	"tls1_1":                 MinimalTlsVersion_STATUS_TLS1_1,
	"tls1_2":                 MinimalTlsVersion_STATUS_TLS1_2,
	"tlsenforcementdisabled": MinimalTlsVersion_STATUS_TLSEnforcementDisabled,
}

// Whether or not public network access is allowed for this server. Value is optional but if passed in, must be 'Enabled'
// or 'Disabled'
type PublicNetworkAccess_STATUS string

const (
	PublicNetworkAccess_STATUS_Disabled = PublicNetworkAccess_STATUS("Disabled")
	PublicNetworkAccess_STATUS_Enabled  = PublicNetworkAccess_STATUS("Enabled")
)

// Mapping from string to PublicNetworkAccess_STATUS
var publicNetworkAccess_STATUS_Values = map[string]PublicNetworkAccess_STATUS{
	"disabled": PublicNetworkAccess_STATUS_Disabled,
	"enabled":  PublicNetworkAccess_STATUS_Enabled,
}

// A private endpoint connection under a server
type ServerPrivateEndpointConnection_STATUS struct {
	// Id: Resource Id of the private endpoint connection.
	Id *string `json:"id,omitempty"`

	// Properties: Private endpoint connection properties
	Properties *ServerPrivateEndpointConnectionProperties_STATUS `json:"properties,omitempty"`
}

type ServerProperties_UserVisibleState_STATUS string

const (
	ServerProperties_UserVisibleState_STATUS_Disabled = ServerProperties_UserVisibleState_STATUS("Disabled")
	ServerProperties_UserVisibleState_STATUS_Dropping = ServerProperties_UserVisibleState_STATUS("Dropping")
	ServerProperties_UserVisibleState_STATUS_Ready    = ServerProperties_UserVisibleState_STATUS("Ready")
)

// Mapping from string to ServerProperties_UserVisibleState_STATUS
var serverProperties_UserVisibleState_STATUS_Values = map[string]ServerProperties_UserVisibleState_STATUS{
	"disabled": ServerProperties_UserVisibleState_STATUS_Disabled,
	"dropping": ServerProperties_UserVisibleState_STATUS_Dropping,
	"ready":    ServerProperties_UserVisibleState_STATUS_Ready,
}

// The version of a server.
type ServerVersion_STATUS string

const (
	ServerVersion_STATUS_102 = ServerVersion_STATUS("10.2")
	ServerVersion_STATUS_103 = ServerVersion_STATUS("10.3")
)

// Mapping from string to ServerVersion_STATUS
var serverVersion_STATUS_Values = map[string]ServerVersion_STATUS{
	"10.2": ServerVersion_STATUS_102,
	"10.3": ServerVersion_STATUS_103,
}

type Sku_Tier_STATUS string

const (
	Sku_Tier_STATUS_Basic           = Sku_Tier_STATUS("Basic")
	Sku_Tier_STATUS_GeneralPurpose  = Sku_Tier_STATUS("GeneralPurpose")
	Sku_Tier_STATUS_MemoryOptimized = Sku_Tier_STATUS("MemoryOptimized")
)

// Mapping from string to Sku_Tier_STATUS
var sku_Tier_STATUS_Values = map[string]Sku_Tier_STATUS{
	"basic":           Sku_Tier_STATUS_Basic,
	"generalpurpose":  Sku_Tier_STATUS_GeneralPurpose,
	"memoryoptimized": Sku_Tier_STATUS_MemoryOptimized,
}

// Enable ssl enforcement or not when connect to server.
type SslEnforcement_STATUS string

const (
	SslEnforcement_STATUS_Disabled = SslEnforcement_STATUS("Disabled")
	SslEnforcement_STATUS_Enabled  = SslEnforcement_STATUS("Enabled")
)

// Mapping from string to SslEnforcement_STATUS
var sslEnforcement_STATUS_Values = map[string]SslEnforcement_STATUS{
	"disabled": SslEnforcement_STATUS_Disabled,
	"enabled":  SslEnforcement_STATUS_Enabled,
}

// Storage Profile properties of a server
type StorageProfile_STATUS struct {
	// BackupRetentionDays: Backup retention days for the server.
	BackupRetentionDays *int `json:"backupRetentionDays,omitempty"`

	// GeoRedundantBackup: Enable Geo-redundant or not for server backup.
	GeoRedundantBackup *StorageProfile_GeoRedundantBackup_STATUS `json:"geoRedundantBackup,omitempty"`

	// StorageAutogrow: Enable Storage Auto Grow.
	StorageAutogrow *StorageProfile_StorageAutogrow_STATUS `json:"storageAutogrow,omitempty"`

	// StorageMB: Max storage allowed for a server.
	StorageMB *int `json:"storageMB,omitempty"`
}

// Properties of a private endpoint connection.
type ServerPrivateEndpointConnectionProperties_STATUS struct {
	// PrivateEndpoint: Private endpoint which the connection belongs to.
	PrivateEndpoint *PrivateEndpointProperty_STATUS `json:"privateEndpoint,omitempty"`

	// PrivateLinkServiceConnectionState: Connection state of the private endpoint connection.
	PrivateLinkServiceConnectionState *ServerPrivateLinkServiceConnectionStateProperty_STATUS `json:"privateLinkServiceConnectionState,omitempty"`

	// ProvisioningState: State of the private endpoint connection.
	ProvisioningState *ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
}

type StorageProfile_GeoRedundantBackup_STATUS string

const (
	StorageProfile_GeoRedundantBackup_STATUS_Disabled = StorageProfile_GeoRedundantBackup_STATUS("Disabled")
	StorageProfile_GeoRedundantBackup_STATUS_Enabled  = StorageProfile_GeoRedundantBackup_STATUS("Enabled")
)

// Mapping from string to StorageProfile_GeoRedundantBackup_STATUS
var storageProfile_GeoRedundantBackup_STATUS_Values = map[string]StorageProfile_GeoRedundantBackup_STATUS{
	"disabled": StorageProfile_GeoRedundantBackup_STATUS_Disabled,
	"enabled":  StorageProfile_GeoRedundantBackup_STATUS_Enabled,
}

type StorageProfile_StorageAutogrow_STATUS string

const (
	StorageProfile_StorageAutogrow_STATUS_Disabled = StorageProfile_StorageAutogrow_STATUS("Disabled")
	StorageProfile_StorageAutogrow_STATUS_Enabled  = StorageProfile_StorageAutogrow_STATUS("Enabled")
)

// Mapping from string to StorageProfile_StorageAutogrow_STATUS
var storageProfile_StorageAutogrow_STATUS_Values = map[string]StorageProfile_StorageAutogrow_STATUS{
	"disabled": StorageProfile_StorageAutogrow_STATUS_Disabled,
	"enabled":  StorageProfile_StorageAutogrow_STATUS_Enabled,
}

type PrivateEndpointProperty_STATUS struct {
	// Id: Resource id of the private endpoint.
	Id *string `json:"id,omitempty"`
}

type ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS string

const (
	ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Approving = ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS("Approving")
	ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Dropping  = ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS("Dropping")
	ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Failed    = ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS("Failed")
	ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Ready     = ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS("Ready")
	ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Rejecting = ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS("Rejecting")
)

// Mapping from string to ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS
var serverPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Values = map[string]ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS{
	"approving": ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Approving,
	"dropping":  ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Dropping,
	"failed":    ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Failed,
	"ready":     ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Ready,
	"rejecting": ServerPrivateEndpointConnectionProperties_ProvisioningState_STATUS_Rejecting,
}

type ServerPrivateLinkServiceConnectionStateProperty_STATUS struct {
	// ActionsRequired: The actions required for private link service connection.
	ActionsRequired *ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS `json:"actionsRequired,omitempty"`

	// Description: The private link service connection description.
	Description *string `json:"description,omitempty"`

	// Status: The private link service connection status.
	Status *ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS `json:"status,omitempty"`
}

type ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS string

const ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS_None = ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS("None")

// Mapping from string to ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS
var serverPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS_Values = map[string]ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS{
	"none": ServerPrivateLinkServiceConnectionStateProperty_ActionsRequired_STATUS_None,
}

type ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS string

const (
	ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Approved     = ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS("Approved")
	ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Disconnected = ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS("Disconnected")
	ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Pending      = ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS("Pending")
	ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Rejected     = ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS("Rejected")
)

// Mapping from string to ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS
var serverPrivateLinkServiceConnectionStateProperty_Status_STATUS_Values = map[string]ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS{
	"approved":     ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Approved,
	"disconnected": ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Disconnected,
	"pending":      ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Pending,
	"rejected":     ServerPrivateLinkServiceConnectionStateProperty_Status_STATUS_Rejected,
}
