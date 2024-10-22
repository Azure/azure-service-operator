// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210501preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type DiagnosticSetting_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of a Diagnostic Settings Resource.
	Properties *DiagnosticSettings_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &DiagnosticSetting_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01-preview"
func (setting DiagnosticSetting_Spec_ARM) GetAPIVersion() string {
	return "2021-05-01-preview"
}

// GetName returns the Name of the resource
func (setting *DiagnosticSetting_Spec_ARM) GetName() string {
	return setting.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/diagnosticSettings"
func (setting *DiagnosticSetting_Spec_ARM) GetType() string {
	return "Microsoft.Insights/diagnosticSettings"
}

// The diagnostic settings.
type DiagnosticSettings_ARM struct {
	EventHubAuthorizationRuleId *string `json:"eventHubAuthorizationRuleId,omitempty"`

	// EventHubName: The name of the event hub. If none is specified, the default event hub will be selected.
	EventHubName *string `json:"eventHubName,omitempty"`

	// LogAnalyticsDestinationType: A string indicating whether the export to Log Analytics should use the default destination
	// type, i.e. AzureDiagnostics, or use a destination type constructed as follows: <normalized service identity>_<normalized
	// category name>. Possible values are: Dedicated and null (null is default.)
	LogAnalyticsDestinationType *string `json:"logAnalyticsDestinationType,omitempty"`

	// Logs: The list of logs settings.
	Logs                 []LogSettings_ARM `json:"logs,omitempty"`
	MarketplacePartnerId *string           `json:"marketplacePartnerId,omitempty"`

	// Metrics: The list of metric settings.
	Metrics []MetricSettings_ARM `json:"metrics,omitempty"`

	// ServiceBusRuleId: The service bus rule Id of the diagnostic setting. This is here to maintain backwards compatibility.
	ServiceBusRuleId *string `json:"serviceBusRuleId,omitempty"`
	StorageAccountId *string `json:"storageAccountId,omitempty"`
	WorkspaceId      *string `json:"workspaceId,omitempty"`
}

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular log.
type LogSettings_ARM struct {
	// Category: Name of a Diagnostic Log category for a resource type this setting is applied to. To obtain the list of
	// Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
	Category *string `json:"category,omitempty"`

	// CategoryGroup: Name of a Diagnostic Log category group for a resource type this setting is applied to. To obtain the
	// list of Diagnostic Log categories for a resource, first perform a GET diagnostic settings operation.
	CategoryGroup *string `json:"categoryGroup,omitempty"`

	// Enabled: a value indicating whether this log is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// RetentionPolicy: the retention policy for this log.
	RetentionPolicy *RetentionPolicy_ARM `json:"retentionPolicy,omitempty"`
}

// Part of MultiTenantDiagnosticSettings. Specifies the settings for a particular metric.
type MetricSettings_ARM struct {
	// Category: Name of a Diagnostic Metric category for a resource type this setting is applied to. To obtain the list of
	// Diagnostic metric categories for a resource, first perform a GET diagnostic settings operation.
	Category *string `json:"category,omitempty"`

	// Enabled: a value indicating whether this category is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// RetentionPolicy: the retention policy for this category.
	RetentionPolicy *RetentionPolicy_ARM `json:"retentionPolicy,omitempty"`

	// TimeGrain: the timegrain of the metric in ISO8601 format.
	TimeGrain *string `json:"timeGrain,omitempty"`
}

// Specifies the retention policy for the log.
type RetentionPolicy_ARM struct {
	// Days: the number of days for the retention in days. A value of 0 will retain the events indefinitely.
	Days *int `json:"days,omitempty"`

	// Enabled: a value indicating whether the retention policy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}
