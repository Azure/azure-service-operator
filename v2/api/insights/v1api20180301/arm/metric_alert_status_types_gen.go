// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type MetricAlert_STATUS struct {
	// Id: Azure resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Azure resource name
	Name *string `json:"name,omitempty"`

	// Properties: The alert rule properties of the resource.
	Properties *MetricAlertProperties_STATUS `json:"properties,omitempty"`

	// Tags: Resource tags
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Azure resource type
	Type *string `json:"type,omitempty"`
}

// An alert rule.
type MetricAlertProperties_STATUS struct {
	// Actions: the array of actions that are performed when the alert rule becomes active, and when an alert condition is
	// resolved.
	Actions []MetricAlertAction_STATUS `json:"actions,omitempty"`

	// AutoMitigate: the flag that indicates whether the alert should be auto resolved or not. The default is true.
	AutoMitigate *bool `json:"autoMitigate,omitempty"`

	// Criteria: defines the specific alert criteria information.
	Criteria *MetricAlertCriteria_STATUS `json:"criteria,omitempty"`

	// Description: the description of the metric alert that will be included in the alert email.
	Description *string `json:"description,omitempty"`

	// Enabled: the flag that indicates whether the metric alert is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// EvaluationFrequency: how often the metric alert is evaluated represented in ISO 8601 duration format.
	EvaluationFrequency *string `json:"evaluationFrequency,omitempty"`

	// IsMigrated: the value indicating whether this alert rule is migrated.
	IsMigrated *bool `json:"isMigrated,omitempty"`

	// LastUpdatedTime: Last time the rule was updated in ISO8601 format.
	LastUpdatedTime *string `json:"lastUpdatedTime,omitempty"`

	// Scopes: the list of resource id's that this metric alert is scoped to.
	Scopes []string `json:"scopes,omitempty"`

	// Severity: Alert severity {0, 1, 2, 3, 4}
	Severity *int `json:"severity,omitempty"`

	// TargetResourceRegion: the region of the target resource(s) on which the alert is created/updated. Mandatory if the scope
	// contains a subscription, resource group, or more than one resource.
	TargetResourceRegion *string `json:"targetResourceRegion,omitempty"`

	// TargetResourceType: the resource type of the target resource(s) on which the alert is created/updated. Mandatory if the
	// scope contains a subscription, resource group, or more than one resource.
	TargetResourceType *string `json:"targetResourceType,omitempty"`

	// WindowSize: the period of time (in ISO 8601 duration format) that is used to monitor alert activity based on the
	// threshold.
	WindowSize *string `json:"windowSize,omitempty"`
}

// An alert action.
type MetricAlertAction_STATUS struct {
	// ActionGroupId: the id of the action group to use.
	ActionGroupId *string `json:"actionGroupId,omitempty"`

	// WebHookProperties: This field allows specifying custom properties, which would be appended to the alert payload sent as
	// input to the webhook.
	WebHookProperties map[string]string `json:"webHookProperties,omitempty"`
}

type MetricAlertCriteria_STATUS struct {
	// MicrosoftAzureMonitorMultipleResourceMultipleMetric: Mutually exclusive with all other properties
	MicrosoftAzureMonitorMultipleResourceMultipleMetric *MetricAlertMultipleResourceMultipleMetricCriteria_STATUS `json:"microsoftAzureMonitorMultipleResourceMultipleMetricCriteria,omitempty"`

	// MicrosoftAzureMonitorSingleResourceMultipleMetric: Mutually exclusive with all other properties
	MicrosoftAzureMonitorSingleResourceMultipleMetric *MetricAlertSingleResourceMultipleMetricCriteria_STATUS `json:"microsoftAzureMonitorSingleResourceMultipleMetricCriteria,omitempty"`

	// MicrosoftAzureMonitorWebtestLocationAvailability: Mutually exclusive with all other properties
	MicrosoftAzureMonitorWebtestLocationAvailability *WebtestLocationAvailabilityCriteria_STATUS `json:"microsoftAzureMonitorWebtestLocationAvailabilityCriteria,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because MetricAlertCriteria_STATUS represents a discriminated union (JSON OneOf)
func (criteria MetricAlertCriteria_STATUS) MarshalJSON() ([]byte, error) {
	if criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric != nil {
		return json.Marshal(criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric)
	}
	if criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric != nil {
		return json.Marshal(criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric)
	}
	if criteria.MicrosoftAzureMonitorWebtestLocationAvailability != nil {
		return json.Marshal(criteria.MicrosoftAzureMonitorWebtestLocationAvailability)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the MetricAlertCriteria_STATUS
func (criteria *MetricAlertCriteria_STATUS) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["odata.type"]
	if discriminator == "Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria" {
		criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric = &MetricAlertMultipleResourceMultipleMetricCriteria_STATUS{}
		return json.Unmarshal(data, criteria.MicrosoftAzureMonitorMultipleResourceMultipleMetric)
	}
	if discriminator == "Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria" {
		criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric = &MetricAlertSingleResourceMultipleMetricCriteria_STATUS{}
		return json.Unmarshal(data, criteria.MicrosoftAzureMonitorSingleResourceMultipleMetric)
	}
	if discriminator == "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria" {
		criteria.MicrosoftAzureMonitorWebtestLocationAvailability = &WebtestLocationAvailabilityCriteria_STATUS{}
		return json.Unmarshal(data, criteria.MicrosoftAzureMonitorWebtestLocationAvailability)
	}

	// No error
	return nil
}

type MetricAlertMultipleResourceMultipleMetricCriteria_STATUS struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// AllOf: the list of multiple metric criteria for this 'all of' operation.
	AllOf []MultiMetricCriteria_STATUS `json:"allOf,omitempty"`

	// OdataType: specifies the type of the alert criteria.
	OdataType MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS `json:"odata.type,omitempty"`
}

type MetricAlertSingleResourceMultipleMetricCriteria_STATUS struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// AllOf: The list of metric criteria for this 'all of' operation.
	AllOf []MetricCriteria_STATUS `json:"allOf,omitempty"`

	// OdataType: specifies the type of the alert criteria.
	OdataType MetricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS `json:"odata.type,omitempty"`
}

type WebtestLocationAvailabilityCriteria_STATUS struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// ComponentId: The Application Insights resource Id.
	ComponentId *string `json:"componentId,omitempty"`

	// FailedLocationCount: The number of failed locations.
	FailedLocationCount *float64 `json:"failedLocationCount,omitempty"`

	// OdataType: specifies the type of the alert criteria.
	OdataType WebtestLocationAvailabilityCriteria_OdataType_STATUS `json:"odata.type,omitempty"`

	// WebTestId: The Application Insights web test Id.
	WebTestId *string `json:"webTestId,omitempty"`
}

type MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS string

const MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS_MicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria = MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS("Microsoft.Azure.Monitor.MultipleResourceMultipleMetricCriteria")

// Mapping from string to MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS
var metricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS_Values = map[string]MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS{
	"microsoft.azure.monitor.multipleresourcemultiplemetriccriteria": MetricAlertMultipleResourceMultipleMetricCriteria_OdataType_STATUS_MicrosoftAzureMonitorMultipleResourceMultipleMetricCriteria,
}

type MetricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS string

const MetricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria = MetricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS("Microsoft.Azure.Monitor.SingleResourceMultipleMetricCriteria")

// Mapping from string to MetricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS
var metricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS_Values = map[string]MetricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS{
	"microsoft.azure.monitor.singleresourcemultiplemetriccriteria": MetricAlertSingleResourceMultipleMetricCriteria_OdataType_STATUS_MicrosoftAzureMonitorSingleResourceMultipleMetricCriteria,
}

type MetricCriteria_STATUS struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// CriterionType: Specifies the type of threshold criteria
	CriterionType MetricCriteria_CriterionType_STATUS `json:"criterionType,omitempty"`

	// Dimensions: List of dimension conditions.
	Dimensions []MetricDimension_STATUS `json:"dimensions,omitempty"`

	// MetricName: Name of the metric.
	MetricName *string `json:"metricName,omitempty"`

	// MetricNamespace: Namespace of the metric.
	MetricNamespace *string `json:"metricNamespace,omitempty"`

	// Name: Name of the criteria.
	Name *string `json:"name,omitempty"`

	// Operator: the criteria operator.
	Operator *MetricCriteria_Operator_STATUS `json:"operator,omitempty"`

	// SkipMetricValidation: Allows creating an alert rule on a custom metric that isn't yet emitted, by causing the metric
	// validation to be skipped.
	SkipMetricValidation *bool `json:"skipMetricValidation,omitempty"`

	// Threshold: the criteria threshold value that activates the alert.
	Threshold *float64 `json:"threshold,omitempty"`

	// TimeAggregation: the criteria time aggregation types.
	TimeAggregation *MetricCriteria_TimeAggregation_STATUS `json:"timeAggregation,omitempty"`
}

type MultiMetricCriteria_STATUS struct {
	// Dynamic: Mutually exclusive with all other properties
	Dynamic *DynamicMetricCriteria_STATUS `json:"dynamicThresholdCriterion,omitempty"`

	// Static: Mutually exclusive with all other properties
	Static *MetricCriteria_STATUS `json:"staticThresholdCriterion,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because MultiMetricCriteria_STATUS represents a discriminated union (JSON OneOf)
func (criteria MultiMetricCriteria_STATUS) MarshalJSON() ([]byte, error) {
	if criteria.Dynamic != nil {
		return json.Marshal(criteria.Dynamic)
	}
	if criteria.Static != nil {
		return json.Marshal(criteria.Static)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the MultiMetricCriteria_STATUS
func (criteria *MultiMetricCriteria_STATUS) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["criterionType"]
	if discriminator == "DynamicThresholdCriterion" {
		criteria.Dynamic = &DynamicMetricCriteria_STATUS{}
		return json.Unmarshal(data, criteria.Dynamic)
	}
	if discriminator == "StaticThresholdCriterion" {
		criteria.Static = &MetricCriteria_STATUS{}
		return json.Unmarshal(data, criteria.Static)
	}

	// No error
	return nil
}

type WebtestLocationAvailabilityCriteria_OdataType_STATUS string

const WebtestLocationAvailabilityCriteria_OdataType_STATUS_MicrosoftAzureMonitorWebtestLocationAvailabilityCriteria = WebtestLocationAvailabilityCriteria_OdataType_STATUS("Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria")

// Mapping from string to WebtestLocationAvailabilityCriteria_OdataType_STATUS
var webtestLocationAvailabilityCriteria_OdataType_STATUS_Values = map[string]WebtestLocationAvailabilityCriteria_OdataType_STATUS{
	"microsoft.azure.monitor.webtestlocationavailabilitycriteria": WebtestLocationAvailabilityCriteria_OdataType_STATUS_MicrosoftAzureMonitorWebtestLocationAvailabilityCriteria,
}

type DynamicMetricCriteria_STATUS struct {
	AdditionalProperties map[string]v1.JSON `json:"additionalProperties,omitempty"`

	// AlertSensitivity: The extent of deviation required to trigger an alert. This will affect how tight the threshold is to
	// the metric series pattern.
	AlertSensitivity *DynamicMetricCriteria_AlertSensitivity_STATUS `json:"alertSensitivity,omitempty"`

	// CriterionType: Specifies the type of threshold criteria
	CriterionType DynamicMetricCriteria_CriterionType_STATUS `json:"criterionType,omitempty"`

	// Dimensions: List of dimension conditions.
	Dimensions []MetricDimension_STATUS `json:"dimensions,omitempty"`

	// FailingPeriods: The minimum number of violations required within the selected lookback time window required to raise an
	// alert.
	FailingPeriods *DynamicThresholdFailingPeriods_STATUS `json:"failingPeriods,omitempty"`

	// IgnoreDataBefore: Use this option to set the date from which to start learning the metric historical data and calculate
	// the dynamic thresholds (in ISO8601 format)
	IgnoreDataBefore *string `json:"ignoreDataBefore,omitempty"`

	// MetricName: Name of the metric.
	MetricName *string `json:"metricName,omitempty"`

	// MetricNamespace: Namespace of the metric.
	MetricNamespace *string `json:"metricNamespace,omitempty"`

	// Name: Name of the criteria.
	Name *string `json:"name,omitempty"`

	// Operator: The operator used to compare the metric value against the threshold.
	Operator *DynamicMetricCriteria_Operator_STATUS `json:"operator,omitempty"`

	// SkipMetricValidation: Allows creating an alert rule on a custom metric that isn't yet emitted, by causing the metric
	// validation to be skipped.
	SkipMetricValidation *bool `json:"skipMetricValidation,omitempty"`

	// TimeAggregation: the criteria time aggregation types.
	TimeAggregation *DynamicMetricCriteria_TimeAggregation_STATUS `json:"timeAggregation,omitempty"`
}

type MetricCriteria_CriterionType_STATUS string

const MetricCriteria_CriterionType_STATUS_StaticThresholdCriterion = MetricCriteria_CriterionType_STATUS("StaticThresholdCriterion")

// Mapping from string to MetricCriteria_CriterionType_STATUS
var metricCriteria_CriterionType_STATUS_Values = map[string]MetricCriteria_CriterionType_STATUS{
	"staticthresholdcriterion": MetricCriteria_CriterionType_STATUS_StaticThresholdCriterion,
}

type MetricCriteria_Operator_STATUS string

const (
	MetricCriteria_Operator_STATUS_Equals             = MetricCriteria_Operator_STATUS("Equals")
	MetricCriteria_Operator_STATUS_GreaterThan        = MetricCriteria_Operator_STATUS("GreaterThan")
	MetricCriteria_Operator_STATUS_GreaterThanOrEqual = MetricCriteria_Operator_STATUS("GreaterThanOrEqual")
	MetricCriteria_Operator_STATUS_LessThan           = MetricCriteria_Operator_STATUS("LessThan")
	MetricCriteria_Operator_STATUS_LessThanOrEqual    = MetricCriteria_Operator_STATUS("LessThanOrEqual")
)

// Mapping from string to MetricCriteria_Operator_STATUS
var metricCriteria_Operator_STATUS_Values = map[string]MetricCriteria_Operator_STATUS{
	"equals":             MetricCriteria_Operator_STATUS_Equals,
	"greaterthan":        MetricCriteria_Operator_STATUS_GreaterThan,
	"greaterthanorequal": MetricCriteria_Operator_STATUS_GreaterThanOrEqual,
	"lessthan":           MetricCriteria_Operator_STATUS_LessThan,
	"lessthanorequal":    MetricCriteria_Operator_STATUS_LessThanOrEqual,
}

type MetricCriteria_TimeAggregation_STATUS string

const (
	MetricCriteria_TimeAggregation_STATUS_Average = MetricCriteria_TimeAggregation_STATUS("Average")
	MetricCriteria_TimeAggregation_STATUS_Count   = MetricCriteria_TimeAggregation_STATUS("Count")
	MetricCriteria_TimeAggregation_STATUS_Maximum = MetricCriteria_TimeAggregation_STATUS("Maximum")
	MetricCriteria_TimeAggregation_STATUS_Minimum = MetricCriteria_TimeAggregation_STATUS("Minimum")
	MetricCriteria_TimeAggregation_STATUS_Total   = MetricCriteria_TimeAggregation_STATUS("Total")
)

// Mapping from string to MetricCriteria_TimeAggregation_STATUS
var metricCriteria_TimeAggregation_STATUS_Values = map[string]MetricCriteria_TimeAggregation_STATUS{
	"average": MetricCriteria_TimeAggregation_STATUS_Average,
	"count":   MetricCriteria_TimeAggregation_STATUS_Count,
	"maximum": MetricCriteria_TimeAggregation_STATUS_Maximum,
	"minimum": MetricCriteria_TimeAggregation_STATUS_Minimum,
	"total":   MetricCriteria_TimeAggregation_STATUS_Total,
}

// Specifies a metric dimension.
type MetricDimension_STATUS struct {
	// Name: Name of the dimension.
	Name *string `json:"name,omitempty"`

	// Operator: the dimension operator. Only 'Include' and 'Exclude' are supported
	Operator *string `json:"operator,omitempty"`

	// Values: list of dimension values.
	Values []string `json:"values,omitempty"`
}

type DynamicMetricCriteria_AlertSensitivity_STATUS string

const (
	DynamicMetricCriteria_AlertSensitivity_STATUS_High   = DynamicMetricCriteria_AlertSensitivity_STATUS("High")
	DynamicMetricCriteria_AlertSensitivity_STATUS_Low    = DynamicMetricCriteria_AlertSensitivity_STATUS("Low")
	DynamicMetricCriteria_AlertSensitivity_STATUS_Medium = DynamicMetricCriteria_AlertSensitivity_STATUS("Medium")
)

// Mapping from string to DynamicMetricCriteria_AlertSensitivity_STATUS
var dynamicMetricCriteria_AlertSensitivity_STATUS_Values = map[string]DynamicMetricCriteria_AlertSensitivity_STATUS{
	"high":   DynamicMetricCriteria_AlertSensitivity_STATUS_High,
	"low":    DynamicMetricCriteria_AlertSensitivity_STATUS_Low,
	"medium": DynamicMetricCriteria_AlertSensitivity_STATUS_Medium,
}

type DynamicMetricCriteria_CriterionType_STATUS string

const DynamicMetricCriteria_CriterionType_STATUS_DynamicThresholdCriterion = DynamicMetricCriteria_CriterionType_STATUS("DynamicThresholdCriterion")

// Mapping from string to DynamicMetricCriteria_CriterionType_STATUS
var dynamicMetricCriteria_CriterionType_STATUS_Values = map[string]DynamicMetricCriteria_CriterionType_STATUS{
	"dynamicthresholdcriterion": DynamicMetricCriteria_CriterionType_STATUS_DynamicThresholdCriterion,
}

type DynamicMetricCriteria_Operator_STATUS string

const (
	DynamicMetricCriteria_Operator_STATUS_GreaterOrLessThan = DynamicMetricCriteria_Operator_STATUS("GreaterOrLessThan")
	DynamicMetricCriteria_Operator_STATUS_GreaterThan       = DynamicMetricCriteria_Operator_STATUS("GreaterThan")
	DynamicMetricCriteria_Operator_STATUS_LessThan          = DynamicMetricCriteria_Operator_STATUS("LessThan")
)

// Mapping from string to DynamicMetricCriteria_Operator_STATUS
var dynamicMetricCriteria_Operator_STATUS_Values = map[string]DynamicMetricCriteria_Operator_STATUS{
	"greaterorlessthan": DynamicMetricCriteria_Operator_STATUS_GreaterOrLessThan,
	"greaterthan":       DynamicMetricCriteria_Operator_STATUS_GreaterThan,
	"lessthan":          DynamicMetricCriteria_Operator_STATUS_LessThan,
}

type DynamicMetricCriteria_TimeAggregation_STATUS string

const (
	DynamicMetricCriteria_TimeAggregation_STATUS_Average = DynamicMetricCriteria_TimeAggregation_STATUS("Average")
	DynamicMetricCriteria_TimeAggregation_STATUS_Count   = DynamicMetricCriteria_TimeAggregation_STATUS("Count")
	DynamicMetricCriteria_TimeAggregation_STATUS_Maximum = DynamicMetricCriteria_TimeAggregation_STATUS("Maximum")
	DynamicMetricCriteria_TimeAggregation_STATUS_Minimum = DynamicMetricCriteria_TimeAggregation_STATUS("Minimum")
	DynamicMetricCriteria_TimeAggregation_STATUS_Total   = DynamicMetricCriteria_TimeAggregation_STATUS("Total")
)

// Mapping from string to DynamicMetricCriteria_TimeAggregation_STATUS
var dynamicMetricCriteria_TimeAggregation_STATUS_Values = map[string]DynamicMetricCriteria_TimeAggregation_STATUS{
	"average": DynamicMetricCriteria_TimeAggregation_STATUS_Average,
	"count":   DynamicMetricCriteria_TimeAggregation_STATUS_Count,
	"maximum": DynamicMetricCriteria_TimeAggregation_STATUS_Maximum,
	"minimum": DynamicMetricCriteria_TimeAggregation_STATUS_Minimum,
	"total":   DynamicMetricCriteria_TimeAggregation_STATUS_Total,
}

// The minimum number of violations required within the selected lookback time window required to raise an alert.
type DynamicThresholdFailingPeriods_STATUS struct {
	// MinFailingPeriodsToAlert: The number of violations to trigger an alert. Should be smaller or equal to
	// numberOfEvaluationPeriods.
	MinFailingPeriodsToAlert *float64 `json:"minFailingPeriodsToAlert,omitempty"`

	// NumberOfEvaluationPeriods: The number of aggregated lookback points. The lookback time window is calculated based on the
	// aggregation granularity (windowSize) and the selected number of aggregated points.
	NumberOfEvaluationPeriods *float64 `json:"numberOfEvaluationPeriods,omitempty"`
}
