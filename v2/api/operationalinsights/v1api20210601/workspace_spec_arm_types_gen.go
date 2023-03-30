// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Workspace_Spec_ARM struct {
	// Etag: The etag of the workspace.
	Etag *string `json:"etag,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Workspace properties.
	Properties *WorkspaceProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspace_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspace Workspace_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (workspace *Workspace_Spec_ARM) GetName() string {
	return workspace.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.OperationalInsights/workspaces"
func (workspace *Workspace_Spec_ARM) GetType() string {
	return "Microsoft.OperationalInsights/workspaces"
}

// Workspace properties.
type WorkspaceProperties_ARM struct {
	// Features: Workspace features.
	Features *WorkspaceFeatures_ARM `json:"features,omitempty"`

	// ForceCmkForQuery: Indicates whether customer managed storage is mandatory for query management.
	ForceCmkForQuery *bool `json:"forceCmkForQuery,omitempty"`

	// ProvisioningState: The provisioning state of the workspace.
	ProvisioningState *WorkspaceProperties_ProvisioningState `json:"provisioningState,omitempty"`

	// PublicNetworkAccessForIngestion: The network access type for accessing Log Analytics ingestion.
	PublicNetworkAccessForIngestion *PublicNetworkAccessType `json:"publicNetworkAccessForIngestion,omitempty"`

	// PublicNetworkAccessForQuery: The network access type for accessing Log Analytics query.
	PublicNetworkAccessForQuery *PublicNetworkAccessType `json:"publicNetworkAccessForQuery,omitempty"`

	// RetentionInDays: The workspace data retention in days. Allowed values are per pricing plan. See pricing tiers
	// documentation for details.
	RetentionInDays *int `json:"retentionInDays,omitempty"`

	// Sku: The SKU of the workspace.
	Sku *WorkspaceSku_ARM `json:"sku,omitempty"`

	// WorkspaceCapping: The daily volume cap for ingestion.
	WorkspaceCapping *WorkspaceCapping_ARM `json:"workspaceCapping,omitempty"`
}

// The daily volume cap for ingestion.
type WorkspaceCapping_ARM struct {
	// DailyQuotaGb: The workspace daily quota for ingestion.
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty"`
}

// Workspace features.
type WorkspaceFeatures_ARM struct {
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

// The SKU (tier) of a workspace.
type WorkspaceSku_ARM struct {
	// CapacityReservationLevel: The capacity reservation level in GB for this workspace, when CapacityReservation sku is
	// selected.
	CapacityReservationLevel *WorkspaceSku_CapacityReservationLevel `json:"capacityReservationLevel,omitempty"`

	// Name: The name of the SKU.
	Name *WorkspaceSku_Name `json:"name,omitempty"`
}
