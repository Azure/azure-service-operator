// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type PrometheusRuleGroup_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The Prometheus rule group properties of the resource.
	Properties *PrometheusRuleGroupProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// An Azure Prometheus rule group.
type PrometheusRuleGroupProperties_STATUS struct {
	// ClusterName: Apply rule to data from a specific cluster.
	ClusterName *string `json:"clusterName,omitempty"`

	// Description: Rule group description.
	Description *string `json:"description,omitempty"`

	// Enabled: Enable/disable rule group.
	Enabled *bool `json:"enabled,omitempty"`

	// Interval: The interval in which to run the Prometheus rule group represented in ISO 8601 duration format. Should be
	// between 1 and 15 minutes
	Interval *string `json:"interval,omitempty"`

	// Rules: Defines the rules in the Prometheus rule group.
	Rules []PrometheusRule_STATUS `json:"rules,omitempty"`

	// Scopes: Target Azure Monitor workspaces resource ids. This api-version is currently limited to creating with one scope.
	// This may change in future.
	Scopes []string `json:"scopes,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// An Azure Prometheus alerting or recording rule.
type PrometheusRule_STATUS struct {
	// Actions: Actions that are performed when the alert rule becomes active, and when an alert condition is resolved.
	Actions []PrometheusRuleGroupAction_STATUS `json:"actions,omitempty"`

	// Alert: Alert rule name.
	Alert *string `json:"alert,omitempty"`

	// Annotations: The annotations clause specifies a set of informational labels that can be used to store longer additional
	// information such as alert descriptions or runbook links. The annotation values can be templated.
	Annotations map[string]string `json:"annotations,omitempty"`

	// Enabled: Enable/disable rule.
	Enabled *bool `json:"enabled,omitempty"`

	// Expression: The PromQL expression to evaluate. https://prometheus.io/docs/prometheus/latest/querying/basics/. Evaluated
	// periodically as given by 'interval', and the result recorded as a new set of time series with the metric name as given
	// by 'record'.
	Expression *string `json:"expression,omitempty"`

	// For: The amount of time alert must be active before firing.
	For *string `json:"for,omitempty"`

	// Labels: Labels to add or overwrite before storing the result.
	Labels map[string]string `json:"labels,omitempty"`

	// Record: Recorded metrics name.
	Record *string `json:"record,omitempty"`

	// ResolveConfiguration: Defines the configuration for resolving fired alerts. Only relevant for alerts.
	ResolveConfiguration *PrometheusRuleResolveConfiguration_STATUS `json:"resolveConfiguration,omitempty"`

	// Severity: The severity of the alerts fired by the rule. Must be between 0 and 4.
	Severity *int `json:"severity,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

// An alert action. Only relevant for alerts.
type PrometheusRuleGroupAction_STATUS struct {
	// ActionGroupId: The resource id of the action group to use.
	ActionGroupId *string `json:"actionGroupId,omitempty"`

	// ActionProperties: The properties of an action group object.
	ActionProperties map[string]string `json:"actionProperties,omitempty"`
}

// Specifies the Prometheus alert rule configuration.
type PrometheusRuleResolveConfiguration_STATUS struct {
	// AutoResolved: Enable alert auto-resolution.
	AutoResolved *bool `json:"autoResolved,omitempty"`

	// TimeToResolve: Alert auto-resolution timeout.
	TimeToResolve *string `json:"timeToResolve,omitempty"`
}
