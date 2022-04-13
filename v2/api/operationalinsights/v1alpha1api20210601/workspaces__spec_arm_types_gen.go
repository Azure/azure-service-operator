// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

//Deprecated version of Workspaces_Spec. Use v1beta20210601.Workspaces_Spec instead
type Workspaces_SpecARM struct {
	ETag       *string                 `json:"eTag,omitempty"`
	Location   *string                 `json:"location,omitempty"`
	Name       string                  `json:"name,omitempty"`
	Properties *WorkspacePropertiesARM `json:"properties,omitempty"`
	Tags       map[string]string       `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Workspaces_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (workspaces Workspaces_SpecARM) GetAPIVersion() string {
	return "2021-06-01"
}

// GetName returns the Name of the resource
func (workspaces Workspaces_SpecARM) GetName() string {
	return workspaces.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.OperationalInsights/workspaces"
func (workspaces Workspaces_SpecARM) GetType() string {
	return "Microsoft.OperationalInsights/workspaces"
}

//Deprecated version of WorkspaceProperties. Use v1beta20210601.WorkspaceProperties instead
type WorkspacePropertiesARM struct {
	Features                        *WorkspaceFeaturesARM                               `json:"features,omitempty"`
	ForceCmkForQuery                *bool                                               `json:"forceCmkForQuery,omitempty"`
	ProvisioningState               *WorkspacePropertiesProvisioningState               `json:"provisioningState,omitempty"`
	PublicNetworkAccessForIngestion *WorkspacePropertiesPublicNetworkAccessForIngestion `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *WorkspacePropertiesPublicNetworkAccessForQuery     `json:"publicNetworkAccessForQuery,omitempty"`
	RetentionInDays                 *int                                                `json:"retentionInDays,omitempty"`
	Sku                             *WorkspaceSkuARM                                    `json:"sku,omitempty"`
	WorkspaceCapping                *WorkspaceCappingARM                                `json:"workspaceCapping,omitempty"`
}

//Deprecated version of WorkspaceCapping. Use v1beta20210601.WorkspaceCapping instead
type WorkspaceCappingARM struct {
	DailyQuotaGb *float64 `json:"dailyQuotaGb,omitempty"`
}

//Deprecated version of WorkspaceFeatures. Use v1beta20210601.WorkspaceFeatures instead
type WorkspaceFeaturesARM struct {
	AdditionalProperties                        map[string]v1.JSON `json:"additionalProperties,omitempty"`
	ClusterResourceId                           *string            `json:"clusterResourceId,omitempty"`
	DisableLocalAuth                            *bool              `json:"disableLocalAuth,omitempty"`
	EnableDataExport                            *bool              `json:"enableDataExport,omitempty"`
	EnableLogAccessUsingOnlyResourcePermissions *bool              `json:"enableLogAccessUsingOnlyResourcePermissions,omitempty"`
	ImmediatePurgeDataOn30Days                  *bool              `json:"immediatePurgeDataOn30Days,omitempty"`
}

//Deprecated version of WorkspaceSku. Use v1beta20210601.WorkspaceSku instead
type WorkspaceSkuARM struct {
	CapacityReservationLevel *int              `json:"capacityReservationLevel,omitempty"`
	Name                     *WorkspaceSkuName `json:"name,omitempty"`
}
