// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200202

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Component_Spec_ARM struct {
	// Etag: Resource etag
	Etag *string `json:"etag,omitempty"`

	// Kind: The kind of application that this component refers to, used to customize UI. This value is a freeform string,
	// values should typically be one of the following: web, ios, other, store, java, phone.
	Kind *string `json:"kind,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties that define an Application Insights component resource.
	Properties *ApplicationInsightsComponentProperties_ARM `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`
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

// Properties that define an Application Insights component resource.
type ApplicationInsightsComponentProperties_ARM struct {
	// Application_Type: Type of application being monitored.
	Application_Type *ApplicationInsightsComponentProperties_Application_Type_ARM `json:"Application_Type,omitempty"`

	// DisableIpMasking: Disable IP masking.
	DisableIpMasking *bool `json:"DisableIpMasking,omitempty"`

	// DisableLocalAuth: Disable Non-AAD based Auth.
	DisableLocalAuth *bool `json:"DisableLocalAuth,omitempty"`

	// Flow_Type: Used by the Application Insights system to determine what kind of flow this component was created by. This is
	// to be set to 'Bluefield' when creating/updating a component via the REST API.
	Flow_Type *ApplicationInsightsComponentProperties_Flow_Type_ARM `json:"Flow_Type,omitempty"`

	// ForceCustomerStorageForProfiler: Force users to create their own storage account for profiler and debugger.
	ForceCustomerStorageForProfiler *bool `json:"ForceCustomerStorageForProfiler,omitempty"`

	// HockeyAppId: The unique application ID created when a new application is added to HockeyApp, used for communications
	// with HockeyApp.
	HockeyAppId *string `json:"HockeyAppId,omitempty"`

	// ImmediatePurgeDataOn30Days: Purge data immediately after 30 days.
	ImmediatePurgeDataOn30Days *bool `json:"ImmediatePurgeDataOn30Days,omitempty"`

	// IngestionMode: Indicates the flow of the ingestion.
	IngestionMode *ApplicationInsightsComponentProperties_IngestionMode_ARM `json:"IngestionMode,omitempty"`

	// PublicNetworkAccessForIngestion: The network access type for accessing Application Insights ingestion.
	PublicNetworkAccessForIngestion *PublicNetworkAccessType_ARM `json:"publicNetworkAccessForIngestion,omitempty"`

	// PublicNetworkAccessForQuery: The network access type for accessing Application Insights query.
	PublicNetworkAccessForQuery *PublicNetworkAccessType_ARM `json:"publicNetworkAccessForQuery,omitempty"`

	// Request_Source: Describes what tool created this Application Insights component. Customers using this API should set
	// this to the default 'rest'.
	Request_Source *ApplicationInsightsComponentProperties_Request_Source_ARM `json:"Request_Source,omitempty"`

	// RetentionInDays: Retention period in days.
	RetentionInDays *int `json:"RetentionInDays,omitempty"`

	// SamplingPercentage: Percentage of the data produced by the application being monitored that is being sampled for
	// Application Insights telemetry.
	SamplingPercentage  *float64 `json:"SamplingPercentage,omitempty"`
	WorkspaceResourceId *string  `json:"workspaceResourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"other","web"}
type ApplicationInsightsComponentProperties_Application_Type_ARM string

const (
	ApplicationInsightsComponentProperties_Application_Type_ARM_Other = ApplicationInsightsComponentProperties_Application_Type_ARM("other")
	ApplicationInsightsComponentProperties_Application_Type_ARM_Web   = ApplicationInsightsComponentProperties_Application_Type_ARM("web")
)

// Mapping from string to ApplicationInsightsComponentProperties_Application_Type_ARM
var applicationInsightsComponentProperties_Application_Type_ARM_Values = map[string]ApplicationInsightsComponentProperties_Application_Type_ARM{
	"other": ApplicationInsightsComponentProperties_Application_Type_ARM_Other,
	"web":   ApplicationInsightsComponentProperties_Application_Type_ARM_Web,
}

// +kubebuilder:validation:Enum={"Bluefield"}
type ApplicationInsightsComponentProperties_Flow_Type_ARM string

const ApplicationInsightsComponentProperties_Flow_Type_ARM_Bluefield = ApplicationInsightsComponentProperties_Flow_Type_ARM("Bluefield")

// Mapping from string to ApplicationInsightsComponentProperties_Flow_Type_ARM
var applicationInsightsComponentProperties_Flow_Type_ARM_Values = map[string]ApplicationInsightsComponentProperties_Flow_Type_ARM{
	"bluefield": ApplicationInsightsComponentProperties_Flow_Type_ARM_Bluefield,
}

// +kubebuilder:validation:Enum={"ApplicationInsights","ApplicationInsightsWithDiagnosticSettings","LogAnalytics"}
type ApplicationInsightsComponentProperties_IngestionMode_ARM string

const (
	ApplicationInsightsComponentProperties_IngestionMode_ARM_ApplicationInsights                       = ApplicationInsightsComponentProperties_IngestionMode_ARM("ApplicationInsights")
	ApplicationInsightsComponentProperties_IngestionMode_ARM_ApplicationInsightsWithDiagnosticSettings = ApplicationInsightsComponentProperties_IngestionMode_ARM("ApplicationInsightsWithDiagnosticSettings")
	ApplicationInsightsComponentProperties_IngestionMode_ARM_LogAnalytics                              = ApplicationInsightsComponentProperties_IngestionMode_ARM("LogAnalytics")
)

// Mapping from string to ApplicationInsightsComponentProperties_IngestionMode_ARM
var applicationInsightsComponentProperties_IngestionMode_ARM_Values = map[string]ApplicationInsightsComponentProperties_IngestionMode_ARM{
	"applicationinsights":                       ApplicationInsightsComponentProperties_IngestionMode_ARM_ApplicationInsights,
	"applicationinsightswithdiagnosticsettings": ApplicationInsightsComponentProperties_IngestionMode_ARM_ApplicationInsightsWithDiagnosticSettings,
	"loganalytics":                              ApplicationInsightsComponentProperties_IngestionMode_ARM_LogAnalytics,
}

// +kubebuilder:validation:Enum={"rest"}
type ApplicationInsightsComponentProperties_Request_Source_ARM string

const ApplicationInsightsComponentProperties_Request_Source_ARM_Rest = ApplicationInsightsComponentProperties_Request_Source_ARM("rest")

// Mapping from string to ApplicationInsightsComponentProperties_Request_Source_ARM
var applicationInsightsComponentProperties_Request_Source_ARM_Values = map[string]ApplicationInsightsComponentProperties_Request_Source_ARM{
	"rest": ApplicationInsightsComponentProperties_Request_Source_ARM_Rest,
}

// The network access type for operating on the Application Insights Component. By default it is Enabled
// +kubebuilder:validation:Enum={"Disabled","Enabled"}
type PublicNetworkAccessType_ARM string

const (
	PublicNetworkAccessType_ARM_Disabled = PublicNetworkAccessType_ARM("Disabled")
	PublicNetworkAccessType_ARM_Enabled  = PublicNetworkAccessType_ARM("Enabled")
)

// Mapping from string to PublicNetworkAccessType_ARM
var publicNetworkAccessType_ARM_Values = map[string]PublicNetworkAccessType_ARM{
	"disabled": PublicNetworkAccessType_ARM_Disabled,
	"enabled":  PublicNetworkAccessType_ARM_Enabled,
}
