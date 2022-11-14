// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// Deprecated version of Component_Spec. Use v1beta20200202.Component_Spec instead
type Component_Spec_ARM struct {
	Etag       *string                                     `json:"etag,omitempty"`
	Kind       *string                                     `json:"kind,omitempty"`
	Location   *string                                     `json:"location,omitempty"`
	Name       string                                      `json:"name,omitempty"`
	Properties *ApplicationInsightsComponentProperties_ARM `json:"properties,omitempty"`
	Tags       *v1.JSON                                    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Component_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-02-02"
func (component Component_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (component *Component_Spec_ARM) GetName() string {
	return component.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/components"
func (component *Component_Spec_ARM) GetType() string {
	return "Microsoft.Insights/components"
}

// Deprecated version of ApplicationInsightsComponentProperties. Use v1beta20200202.ApplicationInsightsComponentProperties instead
type ApplicationInsightsComponentProperties_ARM struct {
	Application_Type                *ApplicationInsightsComponentProperties_Application_Type `json:"Application_Type,omitempty"`
	DisableIpMasking                *bool                                                    `json:"DisableIpMasking,omitempty"`
	DisableLocalAuth                *bool                                                    `json:"DisableLocalAuth,omitempty"`
	Flow_Type                       *ApplicationInsightsComponentProperties_Flow_Type        `json:"Flow_Type,omitempty"`
	ForceCustomerStorageForProfiler *bool                                                    `json:"ForceCustomerStorageForProfiler,omitempty"`
	HockeyAppId                     *string                                                  `json:"HockeyAppId,omitempty"`
	ImmediatePurgeDataOn30Days      *bool                                                    `json:"ImmediatePurgeDataOn30Days,omitempty"`
	IngestionMode                   *ApplicationInsightsComponentProperties_IngestionMode    `json:"IngestionMode,omitempty"`
	PublicNetworkAccessForIngestion *PublicNetworkAccessType                                 `json:"publicNetworkAccessForIngestion,omitempty"`
	PublicNetworkAccessForQuery     *PublicNetworkAccessType                                 `json:"publicNetworkAccessForQuery,omitempty"`
	Request_Source                  *ApplicationInsightsComponentProperties_Request_Source   `json:"Request_Source,omitempty"`
	RetentionInDays                 *int                                                     `json:"RetentionInDays,omitempty"`
	SamplingPercentage              *float64                                                 `json:"SamplingPercentage,omitempty"`
	WorkspaceResourceId             *string                                                  `json:"workspaceResourceId,omitempty"`
}
