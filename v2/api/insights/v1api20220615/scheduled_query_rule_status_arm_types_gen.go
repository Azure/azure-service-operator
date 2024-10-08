// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220615

type ScheduledQueryRule_STATUS_ARM struct {
	// Etag: The etag field is *not* required. If it is provided in the response body, it must also be provided as a header per
	// the normal etag convention.  Entity tags are used for comparing two or more entities from the same requested resource.
	// HTTP/1.1 uses entity tags in the etag (section 14.19), If-Match (section 14.24), If-None-Match (section 14.26), and
	// If-Range (section 14.27) header fields.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Kind: Indicates the type of scheduled query rule. The default is LogAlert.
	Kind *ScheduledQueryRule_Kind_STATUS_ARM `json:"kind,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The rule properties of the resource.
	Properties *ScheduledQueryRuleProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: SystemData of ScheduledQueryRule.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ScheduledQueryRule_Kind_STATUS_ARM string

const (
	ScheduledQueryRule_Kind_STATUS_ARM_LogAlert    = ScheduledQueryRule_Kind_STATUS_ARM("LogAlert")
	ScheduledQueryRule_Kind_STATUS_ARM_LogToMetric = ScheduledQueryRule_Kind_STATUS_ARM("LogToMetric")
)

// Mapping from string to ScheduledQueryRule_Kind_STATUS_ARM
var scheduledQueryRule_Kind_STATUS_ARM_Values = map[string]ScheduledQueryRule_Kind_STATUS_ARM{
	"logalert":    ScheduledQueryRule_Kind_STATUS_ARM_LogAlert,
	"logtometric": ScheduledQueryRule_Kind_STATUS_ARM_LogToMetric,
}

// scheduled query rule Definition
type ScheduledQueryRuleProperties_STATUS_ARM struct {
	// Actions: Actions to invoke when the alert fires.
	Actions *Actions_STATUS_ARM `json:"actions,omitempty"`

	// AutoMitigate: The flag that indicates whether the alert should be automatically resolved or not. The default is true.
	// Relevant only for rules of the kind LogAlert.
	AutoMitigate *bool `json:"autoMitigate,omitempty"`

	// CheckWorkspaceAlertsStorageConfigured: The flag which indicates whether this scheduled query rule should be stored in
	// the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured *bool `json:"checkWorkspaceAlertsStorageConfigured,omitempty"`

	// CreatedWithApiVersion: The api-version used when creating this alert rule
	CreatedWithApiVersion *string `json:"createdWithApiVersion,omitempty"`

	// Criteria: The rule criteria that defines the conditions of the scheduled query rule.
	Criteria *ScheduledQueryRuleCriteria_STATUS_ARM `json:"criteria,omitempty"`

	// Description: The description of the scheduled query rule.
	Description *string `json:"description,omitempty"`

	// DisplayName: The display name of the alert rule
	DisplayName *string `json:"displayName,omitempty"`

	// Enabled: The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled *bool `json:"enabled,omitempty"`

	// EvaluationFrequency: How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant
	// and required only for rules of the kind LogAlert.
	EvaluationFrequency *string `json:"evaluationFrequency,omitempty"`

	// IsLegacyLogAnalyticsRule: True if alert rule is legacy Log Analytic rule
	IsLegacyLogAnalyticsRule *bool `json:"isLegacyLogAnalyticsRule,omitempty"`

	// IsWorkspaceAlertsStorageConfigured: The flag which indicates whether this scheduled query rule has been configured to be
	// stored in the customer's storage. The default is false.
	IsWorkspaceAlertsStorageConfigured *bool `json:"isWorkspaceAlertsStorageConfigured,omitempty"`

	// MuteActionsDuration: Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
	// Relevant only for rules of the kind LogAlert.
	MuteActionsDuration *string `json:"muteActionsDuration,omitempty"`

	// OverrideQueryTimeRange: If specified then overrides the query time range (default is
	// WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange *string `json:"overrideQueryTimeRange,omitempty"`

	// Scopes: The list of resource id's that this scheduled query rule is scoped to.
	Scopes []string `json:"scopes,omitempty"`

	// Severity: Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only
	// for rules of the kind LogAlert.
	Severity *ScheduledQueryRuleProperties_Severity_STATUS_ARM `json:"severity,omitempty"`

	// SkipQueryValidation: The flag which indicates whether the provided query should be validated or not. The default is
	// false. Relevant only for rules of the kind LogAlert.
	SkipQueryValidation *bool `json:"skipQueryValidation,omitempty"`

	// TargetResourceTypes: List of resource type of the target resource(s) on which the alert is created/updated. For example
	// if the scope is a resource group and targetResourceTypes is Microsoft.Compute/virtualMachines, then a different alert
	// will be fired for each virtual machine in the resource group which meet the alert criteria. Relevant only for rules of
	// the kind LogAlert
	TargetResourceTypes []string `json:"targetResourceTypes,omitempty"`

	// WindowSize: The period of time (in ISO 8601 duration format) on which the Alert query will be executed (bin size).
	// Relevant and required only for rules of the kind LogAlert.
	WindowSize *string `json:"windowSize,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS_ARM `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS_ARM `json:"lastModifiedByType,omitempty"`
}

// Actions to invoke when the alert fires.
type Actions_STATUS_ARM struct {
	// ActionGroups: Action Group resource Ids to invoke when the alert fires.
	ActionGroups []string `json:"actionGroups,omitempty"`

	// CustomProperties: The properties of an alert payload.
	CustomProperties map[string]string `json:"customProperties,omitempty"`
}

// The rule criteria that defines the conditions of the scheduled query rule.
type ScheduledQueryRuleCriteria_STATUS_ARM struct {
	// AllOf: A list of conditions to evaluate against the specified scopes
	AllOf []Condition_STATUS_ARM `json:"allOf,omitempty"`
}

type ScheduledQueryRuleProperties_Severity_STATUS_ARM int

const (
	ScheduledQueryRuleProperties_Severity_STATUS_ARM_0 = ScheduledQueryRuleProperties_Severity_STATUS_ARM(0)
	ScheduledQueryRuleProperties_Severity_STATUS_ARM_1 = ScheduledQueryRuleProperties_Severity_STATUS_ARM(1)
	ScheduledQueryRuleProperties_Severity_STATUS_ARM_2 = ScheduledQueryRuleProperties_Severity_STATUS_ARM(2)
	ScheduledQueryRuleProperties_Severity_STATUS_ARM_3 = ScheduledQueryRuleProperties_Severity_STATUS_ARM(3)
	ScheduledQueryRuleProperties_Severity_STATUS_ARM_4 = ScheduledQueryRuleProperties_Severity_STATUS_ARM(4)
)

type SystemData_CreatedByType_STATUS_ARM string

const (
	SystemData_CreatedByType_STATUS_ARM_Application     = SystemData_CreatedByType_STATUS_ARM("Application")
	SystemData_CreatedByType_STATUS_ARM_Key             = SystemData_CreatedByType_STATUS_ARM("Key")
	SystemData_CreatedByType_STATUS_ARM_ManagedIdentity = SystemData_CreatedByType_STATUS_ARM("ManagedIdentity")
	SystemData_CreatedByType_STATUS_ARM_User            = SystemData_CreatedByType_STATUS_ARM("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS_ARM
var systemData_CreatedByType_STATUS_ARM_Values = map[string]SystemData_CreatedByType_STATUS_ARM{
	"application":     SystemData_CreatedByType_STATUS_ARM_Application,
	"key":             SystemData_CreatedByType_STATUS_ARM_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ARM_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_ARM_User,
}

type SystemData_LastModifiedByType_STATUS_ARM string

const (
	SystemData_LastModifiedByType_STATUS_ARM_Application     = SystemData_LastModifiedByType_STATUS_ARM("Application")
	SystemData_LastModifiedByType_STATUS_ARM_Key             = SystemData_LastModifiedByType_STATUS_ARM("Key")
	SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity = SystemData_LastModifiedByType_STATUS_ARM("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_ARM_User            = SystemData_LastModifiedByType_STATUS_ARM("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS_ARM
var systemData_LastModifiedByType_STATUS_ARM_Values = map[string]SystemData_LastModifiedByType_STATUS_ARM{
	"application":     SystemData_LastModifiedByType_STATUS_ARM_Application,
	"key":             SystemData_LastModifiedByType_STATUS_ARM_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ARM_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_ARM_User,
}

// A condition of the scheduled query rule.
type Condition_STATUS_ARM struct {
	// Dimensions: List of Dimensions conditions
	Dimensions []Dimension_STATUS_ARM `json:"dimensions,omitempty"`

	// FailingPeriods: The minimum number of violations required within the selected lookback time window required to raise an
	// alert. Relevant only for rules of the kind LogAlert.
	FailingPeriods *Condition_FailingPeriods_STATUS_ARM `json:"failingPeriods,omitempty"`

	// MetricMeasureColumn: The column containing the metric measure number. Relevant only for rules of the kind LogAlert.
	MetricMeasureColumn *string `json:"metricMeasureColumn,omitempty"`

	// MetricName: The name of the metric to be sent. Relevant and required only for rules of the kind LogToMetric.
	MetricName *string `json:"metricName,omitempty"`

	// Operator: The criteria operator. Relevant and required only for rules of the kind LogAlert.
	Operator *Condition_Operator_STATUS_ARM `json:"operator,omitempty"`

	// Query: Log query alert
	Query *string `json:"query,omitempty"`

	// ResourceIdColumn: The column containing the resource id. The content of the column must be a uri formatted as resource
	// id. Relevant only for rules of the kind LogAlert.
	ResourceIdColumn *string `json:"resourceIdColumn,omitempty"`

	// Threshold: the criteria threshold value that activates the alert. Relevant and required only for rules of the kind
	// LogAlert.
	Threshold *float64 `json:"threshold,omitempty"`

	// TimeAggregation: Aggregation type. Relevant and required only for rules of the kind LogAlert.
	TimeAggregation *Condition_TimeAggregation_STATUS_ARM `json:"timeAggregation,omitempty"`
}

type Condition_FailingPeriods_STATUS_ARM struct {
	// MinFailingPeriodsToAlert: The number of violations to trigger an alert. Should be smaller or equal to
	// numberOfEvaluationPeriods. Default value is 1
	MinFailingPeriodsToAlert *int `json:"minFailingPeriodsToAlert,omitempty"`

	// NumberOfEvaluationPeriods: The number of aggregated lookback points. The lookback time window is calculated based on the
	// aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
	NumberOfEvaluationPeriods *int `json:"numberOfEvaluationPeriods,omitempty"`
}

type Condition_Operator_STATUS_ARM string

const (
	Condition_Operator_STATUS_ARM_Equals             = Condition_Operator_STATUS_ARM("Equals")
	Condition_Operator_STATUS_ARM_GreaterThan        = Condition_Operator_STATUS_ARM("GreaterThan")
	Condition_Operator_STATUS_ARM_GreaterThanOrEqual = Condition_Operator_STATUS_ARM("GreaterThanOrEqual")
	Condition_Operator_STATUS_ARM_LessThan           = Condition_Operator_STATUS_ARM("LessThan")
	Condition_Operator_STATUS_ARM_LessThanOrEqual    = Condition_Operator_STATUS_ARM("LessThanOrEqual")
)

// Mapping from string to Condition_Operator_STATUS_ARM
var condition_Operator_STATUS_ARM_Values = map[string]Condition_Operator_STATUS_ARM{
	"equals":             Condition_Operator_STATUS_ARM_Equals,
	"greaterthan":        Condition_Operator_STATUS_ARM_GreaterThan,
	"greaterthanorequal": Condition_Operator_STATUS_ARM_GreaterThanOrEqual,
	"lessthan":           Condition_Operator_STATUS_ARM_LessThan,
	"lessthanorequal":    Condition_Operator_STATUS_ARM_LessThanOrEqual,
}

type Condition_TimeAggregation_STATUS_ARM string

const (
	Condition_TimeAggregation_STATUS_ARM_Average = Condition_TimeAggregation_STATUS_ARM("Average")
	Condition_TimeAggregation_STATUS_ARM_Count   = Condition_TimeAggregation_STATUS_ARM("Count")
	Condition_TimeAggregation_STATUS_ARM_Maximum = Condition_TimeAggregation_STATUS_ARM("Maximum")
	Condition_TimeAggregation_STATUS_ARM_Minimum = Condition_TimeAggregation_STATUS_ARM("Minimum")
	Condition_TimeAggregation_STATUS_ARM_Total   = Condition_TimeAggregation_STATUS_ARM("Total")
)

// Mapping from string to Condition_TimeAggregation_STATUS_ARM
var condition_TimeAggregation_STATUS_ARM_Values = map[string]Condition_TimeAggregation_STATUS_ARM{
	"average": Condition_TimeAggregation_STATUS_ARM_Average,
	"count":   Condition_TimeAggregation_STATUS_ARM_Count,
	"maximum": Condition_TimeAggregation_STATUS_ARM_Maximum,
	"minimum": Condition_TimeAggregation_STATUS_ARM_Minimum,
	"total":   Condition_TimeAggregation_STATUS_ARM_Total,
}

// Dimension splitting and filtering definition
type Dimension_STATUS_ARM struct {
	// Name: Name of the dimension
	Name *string `json:"name,omitempty"`

	// Operator: Operator for dimension values
	Operator *Dimension_Operator_STATUS_ARM `json:"operator,omitempty"`

	// Values: List of dimension values
	Values []string `json:"values,omitempty"`
}

type Dimension_Operator_STATUS_ARM string

const (
	Dimension_Operator_STATUS_ARM_Exclude = Dimension_Operator_STATUS_ARM("Exclude")
	Dimension_Operator_STATUS_ARM_Include = Dimension_Operator_STATUS_ARM("Include")
)

// Mapping from string to Dimension_Operator_STATUS_ARM
var dimension_Operator_STATUS_ARM_Values = map[string]Dimension_Operator_STATUS_ARM{
	"exclude": Dimension_Operator_STATUS_ARM_Exclude,
	"include": Dimension_Operator_STATUS_ARM_Include,
}
