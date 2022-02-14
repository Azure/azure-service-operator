// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200202

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Components_SpecARM struct {
	//Etag: Resource etag
	Etag *string `json:"etag,omitempty"`

	//Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
	//values should typically be one of the following: web, ios, other, store, java, phone.
	Kind string `json:"kind"`

	//Location: Location to deploy resource to
	Location string `json:"location,omitempty"`

	//Name: The name of the Application Insights component resource.
	Name string `json:"name"`

	//Properties: Properties that define an Application Insights component resource.
	Properties ApplicationInsightsComponentPropertiesARM `json:"properties"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Components_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-02-02"
func (components Components_SpecARM) GetAPIVersion() string {
	return "2020-02-02"
}

// GetName returns the Name of the resource
func (components Components_SpecARM) GetName() string {
	return components.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/components"
func (components Components_SpecARM) GetType() string {
	return "Microsoft.Insights/components"
}

//Generated from: https://schema.management.azure.com/schemas/2020-02-02/Microsoft.Insights.Application.json#/definitions/ApplicationInsightsComponentProperties
type ApplicationInsightsComponentPropertiesARM struct {
	//ApplicationType: Type of application being monitored.
	ApplicationType ApplicationInsightsComponentPropertiesApplicationType `json:"Application_Type"`

	//DisableIpMasking: Disable IP masking.
	DisableIpMasking *bool `json:"DisableIpMasking,omitempty"`

	//DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"DisableLocalAuth,omitempty"`

	//FlowType: Used by the Application Insights system to determine what kind of flow this component was created by. This is
	//to be set to 'Bluefield' when creating/updating a component via the REST API.
	FlowType *ApplicationInsightsComponentPropertiesFlowType `json:"Flow_Type,omitempty"`

	//ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.
	ForceCustomerStorageForProfiler *bool `json:"ForceCustomerStorageForProfiler,omitempty"`

	//HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
	//with HockeyApp.
	HockeyAppId *string `json:"HockeyAppId,omitempty"`

	//ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"ImmediatePurgeDataOn30Days,omitempty"`

	//IngestionMode: Indicates the flow of the ingestion.
	IngestionMode *ApplicationInsightsComponentPropertiesIngestionMode `json:"IngestionMode,omitempty"`

	//PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.
	PublicNetworkAccessForIngestion *ApplicationInsightsComponentPropertiesPublicNetworkAccessForIngestion `json:"publicNetworkAccessForIngestion,omitempty"`

	//PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.
	PublicNetworkAccessForQuery *ApplicationInsightsComponentPropertiesPublicNetworkAccessForQuery `json:"publicNetworkAccessForQuery,omitempty"`

	//RequestSource: Describes what tool created this Application Insights component. Customers using this API should set this
	//to the default 'rest'.
	RequestSource *ApplicationInsightsComponentPropertiesRequestSource `json:"Request_Source,omitempty"`

	//RetentionInDays: Retention period in days.
	RetentionInDays *int `json:"RetentionInDays,omitempty"`

	//SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
	//Application Insights telemetry.
	SamplingPercentage  *float64 `json:"SamplingPercentage,omitempty"`
	WorkspaceResourceId *string  `json:"workspaceResourceId,omitempty"`
}
