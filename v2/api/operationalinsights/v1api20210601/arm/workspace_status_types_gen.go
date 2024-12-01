// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// The top level Workspace resource container.
type Workspace_STATUS struct {
	// Etag: The etag of the workspace.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Workspace properties.
	Properties *WorkspaceProperties_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// Workspace properties.
type WorkspaceProperties_STATUS struct {
	// CreatedDate: Workspace creation date.
	CreatedDate *string `json:"createdDate,omitempty"`

	// CustomerId: This is a read-only property. Represents the ID associated with the workspace.
	CustomerId *string `json:"customerId,omitempty"`

	// Features: Workspace features.
	Features *WorkspaceFeatures_STATUS `json:"features,omitempty"`

	// ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.
	ForceCmkForQuery *bool `json:"forceCmkForQuery,omitempty"`

	// ModifiedDate: Workspace modification date.
	ModifiedDate *string `json:"modifiedDate,omitempty"`

	// PrivateLinkScopedResources: List of linked private link scope resources.
	PrivateLinkScopedResources []PrivateLinkScopedResource_STATUS `json:"privateLinkScopedResources,omitempty"`

	// ProvisioningState: The provisioning state of the workspace.
	ProvisioningState *WorkspaceProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *PublicNetworkAccessType_STATUS `json:"publicNetworkAccessForIngestion,omitempty"`

	// PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *PublicNetworkAccessType_STATUS `json:"publicNetworkAccessForQuery,omitempty"`

	// RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
	// documentation for details.
	RetentionInDays *int `json:"retentionInDays,omitempty"`

	// Sku: The SKU of the workspace.
	Sku *WorkspaceSku_STATUS `json:"sku,omitempty"`

	// WorkspaceCapping: The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCapping_STATUS `json:"workspaceCapping,omitempty"`
}

// The private link scope resource reference.
type PrivateLinkScopedResource_STATUS struct {
	// ResourceId: The full resource Id of the private link scope resource.
	ResourceId *string `json:"resourceId,omitempty"`

	// ScopeId: The private link scope unique Identifier.
	ScopeId *string `json:"scopeId,omitempty"`
}

// The network access type for operating on the Log Analytics Workspace. By default it is Enabled
type PublicNetworkAccessType_STATUS string

const (
	PublicNetworkAccessType_STATUS_Disabled = PublicNetworkAccessType_STATUS("Disabled")
	PublicNetworkAccessType_STATUS_Enabled  = PublicNetworkAccessType_STATUS("Enabled")
)

// Mapping from string to PublicNetworkAccessType_STATUS
var publicNetworkAccessType_STATUS_Values = map[string]PublicNetworkAccessType_STATUS{
	"disabled": PublicNetworkAccessType_STATUS_Disabled,
	"enabled":  PublicNetworkAccessType_STATUS_Enabled,
}

// The daily volume cap for ingestion.
type WorkspaceCapping_STATUS struct {
	// DailyQuotaGb: The workspace daily quota for ingestion.
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty"`

	// DataIngestionStatus: The status of data ingestion for this workspace.
	DataIngestionStatus *WorkspaceCapping_DataIngestionStatus_STATUS `json:"dataIngestionStatus,omitempty"`

	// QuotaNextResetTime: The time when the quota will be rest.
	QuotaNextResetTime *string `json:"quotaNextResetTime,omitempty"`
}

// Workspace features.
type WorkspaceFeatures_STATUS struct {
	// ClusterResourceId: Dedicated LA cluster resourceId that is linked to the workspaces.
	ClusterResourceId *string `json:"clusterResourceId,omitempty"`

	// DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	// EnableDataExport: Flag that indicate if data should be exported.
	EnableDataExport *bool `json:"enableDataExport,omitempty"`

	// EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission to use - resource or workspace or both.
	EnableLogAccessUsingOnlyResourcePermissions *bool `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`

	// ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"immediatePurgeDataOn30Days,omitempty"`
}

type WorkspaceProperties_ProvisioningState_STATUS string

const (
	WorkspaceProperties_ProvisioningState_STATUS_Canceled            = WorkspaceProperties_ProvisioningState_STATUS("Canceled")
	WorkspaceProperties_ProvisioningState_STATUS_Creating            = WorkspaceProperties_ProvisioningState_STATUS("Creating")
	WorkspaceProperties_ProvisioningState_STATUS_Deleting            = WorkspaceProperties_ProvisioningState_STATUS("Deleting")
	WorkspaceProperties_ProvisioningState_STATUS_Failed              = WorkspaceProperties_ProvisioningState_STATUS("Failed")
	WorkspaceProperties_ProvisioningState_STATUS_ProvisioningAccount = WorkspaceProperties_ProvisioningState_STATUS("ProvisioningAccount")
	WorkspaceProperties_ProvisioningState_STATUS_Succeeded           = WorkspaceProperties_ProvisioningState_STATUS("Succeeded")
	WorkspaceProperties_ProvisioningState_STATUS_Updating            = WorkspaceProperties_ProvisioningState_STATUS("Updating")
)

// Mapping from string to WorkspaceProperties_ProvisioningState_STATUS
var workspaceProperties_ProvisioningState_STATUS_Values = map[string]WorkspaceProperties_ProvisioningState_STATUS{
	"canceled":            WorkspaceProperties_ProvisioningState_STATUS_Canceled,
	"creating":            WorkspaceProperties_ProvisioningState_STATUS_Creating,
	"deleting":            WorkspaceProperties_ProvisioningState_STATUS_Deleting,
	"failed":              WorkspaceProperties_ProvisioningState_STATUS_Failed,
	"provisioningaccount": WorkspaceProperties_ProvisioningState_STATUS_ProvisioningAccount,
	"succeeded":           WorkspaceProperties_ProvisioningState_STATUS_Succeeded,
	"updating":            WorkspaceProperties_ProvisioningState_STATUS_Updating,
}

// The SKU (tier) of a workspace.
type WorkspaceSku_STATUS struct {
	// CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
	// selected.
	CapacityReservationLevel *WorkspaceSku_CapacityReservationLevel_STATUS `json:"capacityReservationLevel,omitempty"`

	// LastSkuUpdate: The last time when the sku was updated.
	LastSkuUpdate *string `json:"lastSkuUpdate,omitempty"`

	// Name: The name of the SKU.
	Name *WorkspaceSku_Name_STATUS `json:"name,omitempty"`
}

type WorkspaceCapping_DataIngestionStatus_STATUS string

const (
	WorkspaceCapping_DataIngestionStatus_STATUS_ApproachingQuota      = WorkspaceCapping_DataIngestionStatus_STATUS("ApproachingQuota")
	WorkspaceCapping_DataIngestionStatus_STATUS_ForceOff              = WorkspaceCapping_DataIngestionStatus_STATUS("ForceOff")
	WorkspaceCapping_DataIngestionStatus_STATUS_ForceOn               = WorkspaceCapping_DataIngestionStatus_STATUS("ForceOn")
	WorkspaceCapping_DataIngestionStatus_STATUS_OverQuota             = WorkspaceCapping_DataIngestionStatus_STATUS("OverQuota")
	WorkspaceCapping_DataIngestionStatus_STATUS_RespectQuota          = WorkspaceCapping_DataIngestionStatus_STATUS("RespectQuota")
	WorkspaceCapping_DataIngestionStatus_STATUS_SubscriptionSuspended = WorkspaceCapping_DataIngestionStatus_STATUS("SubscriptionSuspended")
)

// Mapping from string to WorkspaceCapping_DataIngestionStatus_STATUS
var workspaceCapping_DataIngestionStatus_STATUS_Values = map[string]WorkspaceCapping_DataIngestionStatus_STATUS{
	"approachingquota":      WorkspaceCapping_DataIngestionStatus_STATUS_ApproachingQuota,
	"forceoff":              WorkspaceCapping_DataIngestionStatus_STATUS_ForceOff,
	"forceon":               WorkspaceCapping_DataIngestionStatus_STATUS_ForceOn,
	"overquota":             WorkspaceCapping_DataIngestionStatus_STATUS_OverQuota,
	"respectquota":          WorkspaceCapping_DataIngestionStatus_STATUS_RespectQuota,
	"subscriptionsuspended": WorkspaceCapping_DataIngestionStatus_STATUS_SubscriptionSuspended,
}

type WorkspaceSku_CapacityReservationLevel_STATUS int

const (
	WorkspaceSku_CapacityReservationLevel_STATUS_100  = WorkspaceSku_CapacityReservationLevel_STATUS(100)
	WorkspaceSku_CapacityReservationLevel_STATUS_200  = WorkspaceSku_CapacityReservationLevel_STATUS(200)
	WorkspaceSku_CapacityReservationLevel_STATUS_300  = WorkspaceSku_CapacityReservationLevel_STATUS(300)
	WorkspaceSku_CapacityReservationLevel_STATUS_400  = WorkspaceSku_CapacityReservationLevel_STATUS(400)
	WorkspaceSku_CapacityReservationLevel_STATUS_500  = WorkspaceSku_CapacityReservationLevel_STATUS(500)
	WorkspaceSku_CapacityReservationLevel_STATUS_1000 = WorkspaceSku_CapacityReservationLevel_STATUS(1000)
	WorkspaceSku_CapacityReservationLevel_STATUS_2000 = WorkspaceSku_CapacityReservationLevel_STATUS(2000)
	WorkspaceSku_CapacityReservationLevel_STATUS_5000 = WorkspaceSku_CapacityReservationLevel_STATUS(5000)
)

type WorkspaceSku_Name_STATUS string

const (
	WorkspaceSku_Name_STATUS_CapacityReservation = WorkspaceSku_Name_STATUS("CapacityReservation")
	WorkspaceSku_Name_STATUS_Free                = WorkspaceSku_Name_STATUS("Free")
	WorkspaceSku_Name_STATUS_LACluster           = WorkspaceSku_Name_STATUS("LACluster")
	WorkspaceSku_Name_STATUS_PerGB2018           = WorkspaceSku_Name_STATUS("PerGB2018")
	WorkspaceSku_Name_STATUS_PerNode             = WorkspaceSku_Name_STATUS("PerNode")
	WorkspaceSku_Name_STATUS_Premium             = WorkspaceSku_Name_STATUS("Premium")
	WorkspaceSku_Name_STATUS_Standalone          = WorkspaceSku_Name_STATUS("Standalone")
	WorkspaceSku_Name_STATUS_Standard            = WorkspaceSku_Name_STATUS("Standard")
)

// Mapping from string to WorkspaceSku_Name_STATUS
var workspaceSku_Name_STATUS_Values = map[string]WorkspaceSku_Name_STATUS{
	"capacityreservation": WorkspaceSku_Name_STATUS_CapacityReservation,
	"free":                WorkspaceSku_Name_STATUS_Free,
	"lacluster":           WorkspaceSku_Name_STATUS_LACluster,
	"pergb2018":           WorkspaceSku_Name_STATUS_PerGB2018,
	"pernode":             WorkspaceSku_Name_STATUS_PerNode,
	"premium":             WorkspaceSku_Name_STATUS_Premium,
	"standalone":          WorkspaceSku_Name_STATUS_Standalone,
	"standard":            WorkspaceSku_Name_STATUS_Standard,
}
