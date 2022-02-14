// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

type Workspace_StatusARM struct {
	//ETag: The ETag of the workspace.
	ETag *string `json:"eTag,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Properties: Workspace properties.
	Properties *WorkspaceProperties_StatusARM `json:"properties,omitempty"`

	//Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type WorkspaceProperties_StatusARM struct {
	//CreatedDate: Workspace creation date.
	CreatedDate *string `json:"createdDate,omitempty"`

	//CustomerId: This is a read-only property. Represents the ID associated with the workspace.
	CustomerId *string `json:"customerId,omitempty"`

	//Features: Workspace features.
	Features *WorkspaceFeatures_StatusARM `json:"features,omitempty"`

	//ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.
	ForceCmkForQuery *bool `json:"forceCmkForQuery,omitempty"`

	//ModifiedDate: Workspace modification date.
	ModifiedDate *string `json:"modifiedDate,omitempty"`

	//PrivateLinkScopedResources: List of linked private link scope resources.
	PrivateLinkScopedResources []PrivateLinkScopedResource_StatusARM `json:"privateLinkScopedResources,omitempty"`

	//ProvisioningState: The provisioning state of the workspace.
	ProvisioningState *WorkspacePropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	//PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *PublicNetworkAccessType_Status `json:"publicNetworkAccessForIngestion,omitempty"`

	//PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *PublicNetworkAccessType_Status `json:"publicNetworkAccessForQuery,omitempty"`

	//RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
	//documentation for details.
	RetentionInDays *int `json:"retentionInDays,omitempty"`

	//Sku: The SKU of the workspace.
	Sku *WorkspaceSku_StatusARM `json:"sku,omitempty"`

	//WorkspaceCapping: The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCapping_StatusARM `json:"workspaceCapping,omitempty"`
}

type PrivateLinkScopedResource_StatusARM struct {
	//ResourceId: The full resource Id of the private link scope resource.
	ResourceId *string `json:"resourceId,omitempty"`

	//ScopeId: The private link scope unique Identifier.
	ScopeId *string `json:"scopeId,omitempty"`
}

type WorkspaceCapping_StatusARM struct {
	//DailyQuotaGb: The workspace daily quota for ingestion.
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty"`

	//DataIngestionStatus: The status of data ingestion for this workspace.
	DataIngestionStatus *WorkspaceCappingStatusDataIngestionStatus `json:"dataIngestionStatus,omitempty"`

	//QuotaNextResetTime: The time when the quota will be rest.
	QuotaNextResetTime *string `json:"quotaNextResetTime,omitempty"`
}

type WorkspaceFeatures_StatusARM struct {
	//ClusterResourceId: Dedicated LA cluster resourceId that is linked to the workspaces.
	ClusterResourceId *string `json:"clusterResourceId,omitempty"`

	//DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"disableLocalAuth,omitempty"`

	//EnableDataExport: Flag that indicate if data should be exported.
	EnableDataExport *bool `json:"enableDataExport,omitempty"`

	//EnableLogAccessUsingOnlyResourcePermissions: Flag that indicate which permission to use - resource or workspace or both.
	EnableLogAccessUsingOnlyResourcePermissions *bool `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`

	//ImmediatePurgeDataOn30Days: Flag that describes if we want to remove the data after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"immediatePurgeDataOn30Days,omitempty"`
}

type WorkspaceSku_StatusARM struct {
	//CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
	//selected.
	CapacityReservationLevel *WorkspaceSkuStatusCapacityReservationLevel `json:"capacityReservationLevel,omitempty"`

	//LastSkuUpdate: The last time when the sku was updated.
	LastSkuUpdate *string `json:"lastSkuUpdate,omitempty"`

	//Name: The name of the SKU.
	Name WorkspaceSkuStatusName `json:"name"`
}
