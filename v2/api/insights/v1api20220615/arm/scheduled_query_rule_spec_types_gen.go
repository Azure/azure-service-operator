// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ScheduledQueryRule_Spec struct {
	// Kind: Indicates the type of scheduled query rule. The default is LogAlert.
	Kind *ScheduledQueryRule_Kind_Spec `json:"kind,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: The rule properties of the resource.
	Properties *ScheduledQueryRuleProperties `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ScheduledQueryRule_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-06-15"
func (rule ScheduledQueryRule_Spec) GetAPIVersion() string {
	return "2022-06-15"
}

// GetName returns the Name of the resource
func (rule *ScheduledQueryRule_Spec) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Insights/scheduledQueryRules"
func (rule *ScheduledQueryRule_Spec) GetType() string {
	return "Microsoft.Insights/scheduledQueryRules"
}

// +kubebuilder:validation:Enum={"LogAlert","LogToMetric"}
type ScheduledQueryRule_Kind_Spec string

const (
	ScheduledQueryRule_Kind_Spec_LogAlert    = ScheduledQueryRule_Kind_Spec("LogAlert")
	ScheduledQueryRule_Kind_Spec_LogToMetric = ScheduledQueryRule_Kind_Spec("LogToMetric")
)

// Mapping from string to ScheduledQueryRule_Kind_Spec
var scheduledQueryRule_Kind_Spec_Values = map[string]ScheduledQueryRule_Kind_Spec{
	"logalert":    ScheduledQueryRule_Kind_Spec_LogAlert,
	"logtometric": ScheduledQueryRule_Kind_Spec_LogToMetric,
}

// scheduled query rule Definition
type ScheduledQueryRuleProperties struct {
	// Actions: Actions to invoke when the alert fires.
	Actions *Actions `json:"actions,omitempty"`

	// AutoMitigate: The flag that indicates whether the alert should be automatically resolved or not. The default is true.
	// Relevant only for rules of the kind LogAlert.
	AutoMitigate *bool `json:"autoMitigate,omitempty"`

	// CheckWorkspaceAlertsStorageConfigured: The flag which indicates whether this scheduled query rule should be stored in
	// the customer's storage. The default is false. Relevant only for rules of the kind LogAlert.
	CheckWorkspaceAlertsStorageConfigured *bool `json:"checkWorkspaceAlertsStorageConfigured,omitempty"`

	// Criteria: The rule criteria that defines the conditions of the scheduled query rule.
	Criteria *ScheduledQueryRuleCriteria `json:"criteria,omitempty"`

	// Description: The description of the scheduled query rule.
	Description *string `json:"description,omitempty"`

	// DisplayName: The display name of the alert rule
	DisplayName *string `json:"displayName,omitempty"`

	// Enabled: The flag which indicates whether this scheduled query rule is enabled. Value should be true or false
	Enabled *bool `json:"enabled,omitempty"`

	// EvaluationFrequency: How often the scheduled query rule is evaluated represented in ISO 8601 duration format. Relevant
	// and required only for rules of the kind LogAlert.
	EvaluationFrequency *string `json:"evaluationFrequency,omitempty"`

	// MuteActionsDuration: Mute actions for the chosen period of time (in ISO 8601 duration format) after the alert is fired.
	// Relevant only for rules of the kind LogAlert.
	MuteActionsDuration *string `json:"muteActionsDuration,omitempty"`

	// OverrideQueryTimeRange: If specified then overrides the query time range (default is
	// WindowSize*NumberOfEvaluationPeriods). Relevant only for rules of the kind LogAlert.
	OverrideQueryTimeRange *string  `json:"overrideQueryTimeRange,omitempty"`
	Scopes                 []string `json:"scopes,omitempty"`

	// Severity: Severity of the alert. Should be an integer between [0-4]. Value of 0 is severest. Relevant and required only
	// for rules of the kind LogAlert.
	Severity *ScheduledQueryRuleProperties_Severity `json:"severity,omitempty"`

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

// Actions to invoke when the alert fires.
type Actions struct {
	ActionGroups []string `json:"actionGroups,omitempty"`

	// CustomProperties: The properties of an alert payload.
	CustomProperties map[string]string `json:"customProperties,omitempty"`
}

// The rule criteria that defines the conditions of the scheduled query rule.
type ScheduledQueryRuleCriteria struct {
	// AllOf: A list of conditions to evaluate against the specified scopes
	AllOf []Condition `json:"allOf,omitempty"`
}

// +kubebuilder:validation:Enum={0,1,2,3,4}
type ScheduledQueryRuleProperties_Severity int

const (
	ScheduledQueryRuleProperties_Severity_0 = ScheduledQueryRuleProperties_Severity(0)
	ScheduledQueryRuleProperties_Severity_1 = ScheduledQueryRuleProperties_Severity(1)
	ScheduledQueryRuleProperties_Severity_2 = ScheduledQueryRuleProperties_Severity(2)
	ScheduledQueryRuleProperties_Severity_3 = ScheduledQueryRuleProperties_Severity(3)
	ScheduledQueryRuleProperties_Severity_4 = ScheduledQueryRuleProperties_Severity(4)
)

// A condition of the scheduled query rule.
type Condition struct {
	// Dimensions: List of Dimensions conditions
	Dimensions []Dimension `json:"dimensions,omitempty"`

	// FailingPeriods: The minimum number of violations required within the selected lookback time window required to raise an
	// alert. Relevant only for rules of the kind LogAlert.
	FailingPeriods *Condition_FailingPeriods `json:"failingPeriods,omitempty"`

	// MetricMeasureColumn: The column containing the metric measure number. Relevant only for rules of the kind LogAlert.
	MetricMeasureColumn *string `json:"metricMeasureColumn,omitempty"`

	// MetricName: The name of the metric to be sent. Relevant and required only for rules of the kind LogToMetric.
	MetricName *string `json:"metricName,omitempty"`

	// Operator: The criteria operator. Relevant and required only for rules of the kind LogAlert.
	Operator *Condition_Operator `json:"operator,omitempty"`

	// Query: Log query alert
	Query            *string `json:"query,omitempty"`
	ResourceIdColumn *string `json:"resourceIdColumn,omitempty"`

	// Threshold: the criteria threshold value that activates the alert. Relevant and required only for rules of the kind
	// LogAlert.
	Threshold *float64 `json:"threshold,omitempty"`

	// TimeAggregation: Aggregation type. Relevant and required only for rules of the kind LogAlert.
	TimeAggregation *Condition_TimeAggregation `json:"timeAggregation,omitempty"`
}

type Condition_FailingPeriods struct {
	// MinFailingPeriodsToAlert: The number of violations to trigger an alert. Should be smaller or equal to
	// numberOfEvaluationPeriods. Default value is 1
	MinFailingPeriodsToAlert *int `json:"minFailingPeriodsToAlert,omitempty"`

	// NumberOfEvaluationPeriods: The number of aggregated lookback points. The lookback time window is calculated based on the
	// aggregation granularity (windowSize) and the selected number of aggregated points. Default value is 1
	NumberOfEvaluationPeriods *int `json:"numberOfEvaluationPeriods,omitempty"`
}

// +kubebuilder:validation:Enum={"Equals","GreaterThan","GreaterThanOrEqual","LessThan","LessThanOrEqual"}
type Condition_Operator string

const (
	Condition_Operator_Equals             = Condition_Operator("Equals")
	Condition_Operator_GreaterThan        = Condition_Operator("GreaterThan")
	Condition_Operator_GreaterThanOrEqual = Condition_Operator("GreaterThanOrEqual")
	Condition_Operator_LessThan           = Condition_Operator("LessThan")
	Condition_Operator_LessThanOrEqual    = Condition_Operator("LessThanOrEqual")
)

// Mapping from string to Condition_Operator
var condition_Operator_Values = map[string]Condition_Operator{
	"equals":             Condition_Operator_Equals,
	"greaterthan":        Condition_Operator_GreaterThan,
	"greaterthanorequal": Condition_Operator_GreaterThanOrEqual,
	"lessthan":           Condition_Operator_LessThan,
	"lessthanorequal":    Condition_Operator_LessThanOrEqual,
}

// +kubebuilder:validation:Enum={"Average","Count","Maximum","Minimum","Total"}
type Condition_TimeAggregation string

const (
	Condition_TimeAggregation_Average = Condition_TimeAggregation("Average")
	Condition_TimeAggregation_Count   = Condition_TimeAggregation("Count")
	Condition_TimeAggregation_Maximum = Condition_TimeAggregation("Maximum")
	Condition_TimeAggregation_Minimum = Condition_TimeAggregation("Minimum")
	Condition_TimeAggregation_Total   = Condition_TimeAggregation("Total")
)

// Mapping from string to Condition_TimeAggregation
var condition_TimeAggregation_Values = map[string]Condition_TimeAggregation{
	"average": Condition_TimeAggregation_Average,
	"count":   Condition_TimeAggregation_Count,
	"maximum": Condition_TimeAggregation_Maximum,
	"minimum": Condition_TimeAggregation_Minimum,
	"total":   Condition_TimeAggregation_Total,
}

// Dimension splitting and filtering definition
type Dimension struct {
	// Name: Name of the dimension
	Name *string `json:"name,omitempty"`

	// Operator: Operator for dimension values
	Operator *Dimension_Operator `json:"operator,omitempty"`

	// Values: List of dimension values
	Values []string `json:"values,omitempty"`
}

// +kubebuilder:validation:Enum={"Exclude","Include"}
type Dimension_Operator string

const (
	Dimension_Operator_Exclude = Dimension_Operator("Exclude")
	Dimension_Operator_Include = Dimension_Operator("Include")
)

// Mapping from string to Dimension_Operator
var dimension_Operator_Values = map[string]Dimension_Operator{
	"exclude": Dimension_Operator_Exclude,
	"include": Dimension_Operator_Include,
}
