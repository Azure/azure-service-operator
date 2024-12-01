// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type Autoscalesetting_STATUS struct {
	// Id: Azure resource Id
	Id *string `json:"id,omitempty"`

	// Location: Resource location
	Location *string `json:"location,omitempty"`

	// Name: Azure resource name
	Name *string `json:"name,omitempty"`

	// Properties: The autoscale setting of the resource.
	Properties *AutoscaleSetting_STATUS `json:"properties,omitempty"`

	// SystemData: The system metadata related to the response.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Gets or sets a list of key value pairs that describe the resource. These tags can be used in viewing and grouping
	// this resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	// greater in length than 128 characters and a value no greater in length than 256 characters.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: Azure resource type
	Type *string `json:"type,omitempty"`
}

// A setting that contains all of the configuration for the automatic scaling of a resource.
type AutoscaleSetting_STATUS struct {
	// Enabled: the enabled flag. Specifies whether automatic scaling is enabled for the resource. The default value is 'false'.
	Enabled *bool `json:"enabled,omitempty"`

	// Name: the name of the autoscale setting.
	Name *string `json:"name,omitempty"`

	// Notifications: the collection of notifications.
	Notifications []AutoscaleNotification_STATUS `json:"notifications,omitempty"`

	// PredictiveAutoscalePolicy: the predictive autoscale policy mode.
	PredictiveAutoscalePolicy *PredictiveAutoscalePolicy_STATUS `json:"predictiveAutoscalePolicy,omitempty"`

	// Profiles: the collection of automatic scaling profiles that specify different scaling parameters for different time
	// periods. A maximum of 20 profiles can be specified.
	Profiles []AutoscaleProfile_STATUS `json:"profiles,omitempty"`

	// TargetResourceLocation: the location of the resource that the autoscale setting should be added to.
	TargetResourceLocation *string `json:"targetResourceLocation,omitempty"`

	// TargetResourceUri: the resource identifier of the resource that the autoscale setting should be added to.
	TargetResourceUri *string `json:"targetResourceUri,omitempty"`
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

// Autoscale notification.
type AutoscaleNotification_STATUS struct {
	// Email: the email notification.
	Email *EmailNotification_STATUS `json:"email,omitempty"`

	// Operation: the operation associated with the notification and its value must be "scale"
	Operation *AutoscaleNotification_Operation_STATUS `json:"operation,omitempty"`

	// Webhooks: the collection of webhook notifications.
	Webhooks []WebhookNotification_STATUS `json:"webhooks,omitempty"`
}

// Autoscale profile.
type AutoscaleProfile_STATUS struct {
	// Capacity: the number of instances that can be used during this profile.
	Capacity *ScaleCapacity_STATUS `json:"capacity,omitempty"`

	// FixedDate: the specific date-time for the profile. This element is not used if the Recurrence element is used.
	FixedDate *TimeWindow_STATUS `json:"fixedDate,omitempty"`

	// Name: the name of the profile.
	Name *string `json:"name,omitempty"`

	// Recurrence: the repeating times at which this profile begins. This element is not used if the FixedDate element is used.
	Recurrence *Recurrence_STATUS `json:"recurrence,omitempty"`

	// Rules: the collection of rules that provide the triggers and parameters for the scaling action. A maximum of 10 rules
	// can be specified.
	Rules []ScaleRule_STATUS `json:"rules,omitempty"`
}

// The parameters for enabling predictive autoscale.
type PredictiveAutoscalePolicy_STATUS struct {
	// ScaleLookAheadTime: the amount of time to specify by which instances are launched in advance. It must be between 1
	// minute and 60 minutes in ISO 8601 format.
	ScaleLookAheadTime *string `json:"scaleLookAheadTime,omitempty"`

	// ScaleMode: the predictive autoscale mode
	ScaleMode *PredictiveAutoscalePolicy_ScaleMode_STATUS `json:"scaleMode,omitempty"`
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

type AutoscaleNotification_Operation_STATUS string

const AutoscaleNotification_Operation_STATUS_Scale = AutoscaleNotification_Operation_STATUS("Scale")

// Mapping from string to AutoscaleNotification_Operation_STATUS
var autoscaleNotification_Operation_STATUS_Values = map[string]AutoscaleNotification_Operation_STATUS{
	"scale": AutoscaleNotification_Operation_STATUS_Scale,
}

// Email notification of an autoscale event.
type EmailNotification_STATUS struct {
	// CustomEmails: the custom e-mails list. This value can be null or empty, in which case this attribute will be ignored.
	CustomEmails []string `json:"customEmails,omitempty"`

	// SendToSubscriptionAdministrator: a value indicating whether to send email to subscription administrator.
	SendToSubscriptionAdministrator *bool `json:"sendToSubscriptionAdministrator,omitempty"`

	// SendToSubscriptionCoAdministrators: a value indicating whether to send email to subscription co-administrators.
	SendToSubscriptionCoAdministrators *bool `json:"sendToSubscriptionCoAdministrators,omitempty"`
}

type PredictiveAutoscalePolicy_ScaleMode_STATUS string

const (
	PredictiveAutoscalePolicy_ScaleMode_STATUS_Disabled     = PredictiveAutoscalePolicy_ScaleMode_STATUS("Disabled")
	PredictiveAutoscalePolicy_ScaleMode_STATUS_Enabled      = PredictiveAutoscalePolicy_ScaleMode_STATUS("Enabled")
	PredictiveAutoscalePolicy_ScaleMode_STATUS_ForecastOnly = PredictiveAutoscalePolicy_ScaleMode_STATUS("ForecastOnly")
)

// Mapping from string to PredictiveAutoscalePolicy_ScaleMode_STATUS
var predictiveAutoscalePolicy_ScaleMode_STATUS_Values = map[string]PredictiveAutoscalePolicy_ScaleMode_STATUS{
	"disabled":     PredictiveAutoscalePolicy_ScaleMode_STATUS_Disabled,
	"enabled":      PredictiveAutoscalePolicy_ScaleMode_STATUS_Enabled,
	"forecastonly": PredictiveAutoscalePolicy_ScaleMode_STATUS_ForecastOnly,
}

// The repeating times at which this profile begins. This element is not used if the FixedDate element is used.
type Recurrence_STATUS struct {
	// Frequency: the recurrence frequency. How often the schedule profile should take effect. This value must be Week, meaning
	// each week will have the same set of profiles. For example, to set a daily schedule, set schedule to every day of the
	// week. The frequency property specifies that the schedule is repeated weekly.
	Frequency *Recurrence_Frequency_STATUS `json:"frequency,omitempty"`

	// Schedule: the scheduling constraints for when the profile begins.
	Schedule *RecurrentSchedule_STATUS `json:"schedule,omitempty"`
}

// The number of instances that can be used during this profile.
type ScaleCapacity_STATUS struct {
	// Default: the number of instances that will be set if metrics are not available for evaluation. The default is only used
	// if the current instance count is lower than the default.
	Default *string `json:"default,omitempty"`

	// Maximum: the maximum number of instances for the resource. The actual maximum number of instances is limited by the
	// cores that are available in the subscription.
	Maximum *string `json:"maximum,omitempty"`

	// Minimum: the minimum number of instances for the resource.
	Minimum *string `json:"minimum,omitempty"`
}

// A rule that provide the triggers and parameters for the scaling action.
type ScaleRule_STATUS struct {
	// MetricTrigger: the trigger that results in a scaling action.
	MetricTrigger *MetricTrigger_STATUS `json:"metricTrigger,omitempty"`

	// ScaleAction: the parameters for the scaling action.
	ScaleAction *ScaleAction_STATUS `json:"scaleAction,omitempty"`
}

// A specific date-time for the profile.
type TimeWindow_STATUS struct {
	// End: the end time for the profile in ISO 8601 format.
	End *string `json:"end,omitempty"`

	// Start: the start time for the profile in ISO 8601 format.
	Start *string `json:"start,omitempty"`

	// TimeZone: the timezone of the start and end times for the profile. Some examples of valid time zones are: Dateline
	// Standard Time, UTC-11, Hawaiian Standard Time, Alaskan Standard Time, Pacific Standard Time (Mexico), Pacific Standard
	// Time, US Mountain Standard Time, Mountain Standard Time (Mexico), Mountain Standard Time, Central America Standard Time,
	// Central Standard Time, Central Standard Time (Mexico), Canada Central Standard Time, SA Pacific Standard Time, Eastern
	// Standard Time, US Eastern Standard Time, Venezuela Standard Time, Paraguay Standard Time, Atlantic Standard Time,
	// Central Brazilian Standard Time, SA Western Standard Time, Pacific SA Standard Time, Newfoundland Standard Time, E.
	// South America Standard Time, Argentina Standard Time, SA Eastern Standard Time, Greenland Standard Time, Montevideo
	// Standard Time, Bahia Standard Time, UTC-02, Mid-Atlantic Standard Time, Azores Standard Time, Cape Verde Standard Time,
	// Morocco Standard Time, UTC, GMT Standard Time, Greenwich Standard Time, W. Europe Standard Time, Central Europe Standard
	// Time, Romance Standard Time, Central European Standard Time, W. Central Africa Standard Time, Namibia Standard Time,
	// Jordan Standard Time, GTB Standard Time, Middle East Standard Time, Egypt Standard Time, Syria Standard Time, E. Europe
	// Standard Time, South Africa Standard Time, FLE Standard Time, Turkey Standard Time, Israel Standard Time, Kaliningrad
	// Standard Time, Libya Standard Time, Arabic Standard Time, Arab Standard Time, Belarus Standard Time, Russian Standard
	// Time, E. Africa Standard Time, Iran Standard Time, Arabian Standard Time, Azerbaijan Standard Time, Russia Time Zone 3,
	// Mauritius Standard Time, Georgian Standard Time, Caucasus Standard Time, Afghanistan Standard Time, West Asia Standard
	// Time, Ekaterinburg Standard Time, Pakistan Standard Time, India Standard Time, Sri Lanka Standard Time, Nepal Standard
	// Time, Central Asia Standard Time, Bangladesh Standard Time, N. Central Asia Standard Time, Myanmar Standard Time, SE
	// Asia Standard Time, North Asia Standard Time, China Standard Time, North Asia East Standard Time, Singapore Standard
	// Time, W. Australia Standard Time, Taipei Standard Time, Ulaanbaatar Standard Time, Tokyo Standard Time, Korea Standard
	// Time, Yakutsk Standard Time, Cen. Australia Standard Time, AUS Central Standard Time, E. Australia Standard Time, AUS
	// Eastern Standard Time, West Pacific Standard Time, Tasmania Standard Time, Magadan Standard Time, Vladivostok Standard
	// Time, Russia Time Zone 10, Central Pacific Standard Time, Russia Time Zone 11, New Zealand Standard Time, UTC+12, Fiji
	// Standard Time, Kamchatka Standard Time, Tonga Standard Time, Samoa Standard Time, Line Islands Standard Time
	TimeZone *string `json:"timeZone,omitempty"`
}

// Webhook notification of an autoscale event.
type WebhookNotification_STATUS struct {
	// Properties: a property bag of settings. This value can be empty.
	Properties map[string]string `json:"properties,omitempty"`

	// ServiceUri: the service address to receive the notification.
	ServiceUri *string `json:"serviceUri,omitempty"`
}

// The trigger that results in a scaling action.
type MetricTrigger_STATUS struct {
	// Dimensions: List of dimension conditions. For example:
	// [{"DimensionName":"AppName","Operator":"Equals","Values":["App1"]},{"DimensionName":"Deployment","Operator":"Equals","Values":["default"]}].
	Dimensions []ScaleRuleMetricDimension_STATUS `json:"dimensions,omitempty"`

	// DividePerInstance: a value indicating whether metric should divide per instance.
	DividePerInstance *bool `json:"dividePerInstance,omitempty"`

	// MetricName: the name of the metric that defines what the rule monitors.
	MetricName *string `json:"metricName,omitempty"`

	// MetricNamespace: the namespace of the metric that defines what the rule monitors.
	MetricNamespace *string `json:"metricNamespace,omitempty"`

	// MetricResourceLocation: the location of the resource the rule monitors.
	MetricResourceLocation *string `json:"metricResourceLocation,omitempty"`

	// MetricResourceUri: the resource identifier of the resource the rule monitors.
	MetricResourceUri *string `json:"metricResourceUri,omitempty"`

	// Operator: the operator that is used to compare the metric data and the threshold.
	Operator *MetricTrigger_Operator_STATUS `json:"operator,omitempty"`

	// Statistic: the metric statistic type. How the metrics from multiple instances are combined.
	Statistic *MetricTrigger_Statistic_STATUS `json:"statistic,omitempty"`

	// Threshold: the threshold of the metric that triggers the scale action.
	Threshold *float64 `json:"threshold,omitempty"`

	// TimeAggregation: time aggregation type. How the data that is collected should be combined over time. The default value
	// is Average.
	TimeAggregation *MetricTrigger_TimeAggregation_STATUS `json:"timeAggregation,omitempty"`

	// TimeGrain: the granularity of metrics the rule monitors. Must be one of the predefined values returned from metric
	// definitions for the metric. Must be between 12 hours and 1 minute.
	TimeGrain *string `json:"timeGrain,omitempty"`

	// TimeWindow: the range of time in which instance data is collected. This value must be greater than the delay in metric
	// collection, which can vary from resource-to-resource. Must be between 12 hours and 5 minutes.
	TimeWindow *string `json:"timeWindow,omitempty"`
}

// +kubebuilder:validation:Enum={"Day","Hour","Minute","Month","None","Second","Week","Year"}
type Recurrence_Frequency_STATUS string

const (
	Recurrence_Frequency_STATUS_Day    = Recurrence_Frequency_STATUS("Day")
	Recurrence_Frequency_STATUS_Hour   = Recurrence_Frequency_STATUS("Hour")
	Recurrence_Frequency_STATUS_Minute = Recurrence_Frequency_STATUS("Minute")
	Recurrence_Frequency_STATUS_Month  = Recurrence_Frequency_STATUS("Month")
	Recurrence_Frequency_STATUS_None   = Recurrence_Frequency_STATUS("None")
	Recurrence_Frequency_STATUS_Second = Recurrence_Frequency_STATUS("Second")
	Recurrence_Frequency_STATUS_Week   = Recurrence_Frequency_STATUS("Week")
	Recurrence_Frequency_STATUS_Year   = Recurrence_Frequency_STATUS("Year")
)

// Mapping from string to Recurrence_Frequency_STATUS
var recurrence_Frequency_STATUS_Values = map[string]Recurrence_Frequency_STATUS{
	"day":    Recurrence_Frequency_STATUS_Day,
	"hour":   Recurrence_Frequency_STATUS_Hour,
	"minute": Recurrence_Frequency_STATUS_Minute,
	"month":  Recurrence_Frequency_STATUS_Month,
	"none":   Recurrence_Frequency_STATUS_None,
	"second": Recurrence_Frequency_STATUS_Second,
	"week":   Recurrence_Frequency_STATUS_Week,
	"year":   Recurrence_Frequency_STATUS_Year,
}

// The scheduling constraints for when the profile begins.
type RecurrentSchedule_STATUS struct {
	// Days: the collection of days that the profile takes effect on. Possible values are Sunday through Saturday.
	Days []string `json:"days,omitempty"`

	// Hours: A collection of hours that the profile takes effect on. Values supported are 0 to 23 on the 24-hour clock (AM/PM
	// times are not supported).
	Hours []int `json:"hours,omitempty"`

	// Minutes: A collection of minutes at which the profile takes effect at.
	Minutes []int `json:"minutes,omitempty"`

	// TimeZone: the timezone for the hours of the profile. Some examples of valid time zones are: Dateline Standard Time,
	// UTC-11, Hawaiian Standard Time, Alaskan Standard Time, Pacific Standard Time (Mexico), Pacific Standard Time, US
	// Mountain Standard Time, Mountain Standard Time (Mexico), Mountain Standard Time, Central America Standard Time, Central
	// Standard Time, Central Standard Time (Mexico), Canada Central Standard Time, SA Pacific Standard Time, Eastern Standard
	// Time, US Eastern Standard Time, Venezuela Standard Time, Paraguay Standard Time, Atlantic Standard Time, Central
	// Brazilian Standard Time, SA Western Standard Time, Pacific SA Standard Time, Newfoundland Standard Time, E. South
	// America Standard Time, Argentina Standard Time, SA Eastern Standard Time, Greenland Standard Time, Montevideo Standard
	// Time, Bahia Standard Time, UTC-02, Mid-Atlantic Standard Time, Azores Standard Time, Cape Verde Standard Time, Morocco
	// Standard Time, UTC, GMT Standard Time, Greenwich Standard Time, W. Europe Standard Time, Central Europe Standard Time,
	// Romance Standard Time, Central European Standard Time, W. Central Africa Standard Time, Namibia Standard Time, Jordan
	// Standard Time, GTB Standard Time, Middle East Standard Time, Egypt Standard Time, Syria Standard Time, E. Europe
	// Standard Time, South Africa Standard Time, FLE Standard Time, Turkey Standard Time, Israel Standard Time, Kaliningrad
	// Standard Time, Libya Standard Time, Arabic Standard Time, Arab Standard Time, Belarus Standard Time, Russian Standard
	// Time, E. Africa Standard Time, Iran Standard Time, Arabian Standard Time, Azerbaijan Standard Time, Russia Time Zone 3,
	// Mauritius Standard Time, Georgian Standard Time, Caucasus Standard Time, Afghanistan Standard Time, West Asia Standard
	// Time, Ekaterinburg Standard Time, Pakistan Standard Time, India Standard Time, Sri Lanka Standard Time, Nepal Standard
	// Time, Central Asia Standard Time, Bangladesh Standard Time, N. Central Asia Standard Time, Myanmar Standard Time, SE
	// Asia Standard Time, North Asia Standard Time, China Standard Time, North Asia East Standard Time, Singapore Standard
	// Time, W. Australia Standard Time, Taipei Standard Time, Ulaanbaatar Standard Time, Tokyo Standard Time, Korea Standard
	// Time, Yakutsk Standard Time, Cen. Australia Standard Time, AUS Central Standard Time, E. Australia Standard Time, AUS
	// Eastern Standard Time, West Pacific Standard Time, Tasmania Standard Time, Magadan Standard Time, Vladivostok Standard
	// Time, Russia Time Zone 10, Central Pacific Standard Time, Russia Time Zone 11, New Zealand Standard Time, UTC+12, Fiji
	// Standard Time, Kamchatka Standard Time, Tonga Standard Time, Samoa Standard Time, Line Islands Standard Time
	TimeZone *string `json:"timeZone,omitempty"`
}

// The parameters for the scaling action.
type ScaleAction_STATUS struct {
	// Cooldown: the amount of time to wait since the last scaling action before this action occurs. It must be between 1 week
	// and 1 minute in ISO 8601 format.
	Cooldown *string `json:"cooldown,omitempty"`

	// Direction: the scale direction. Whether the scaling action increases or decreases the number of instances.
	Direction *ScaleAction_Direction_STATUS `json:"direction,omitempty"`

	// Type: the type of action that should occur when the scale rule fires.
	Type *ScaleAction_Type_STATUS `json:"type,omitempty"`

	// Value: the number of instances that are involved in the scaling action. This value must be 1 or greater. The default
	// value is 1.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"Equals","GreaterThan","GreaterThanOrEqual","LessThan","LessThanOrEqual","NotEquals"}
type MetricTrigger_Operator_STATUS string

const (
	MetricTrigger_Operator_STATUS_Equals             = MetricTrigger_Operator_STATUS("Equals")
	MetricTrigger_Operator_STATUS_GreaterThan        = MetricTrigger_Operator_STATUS("GreaterThan")
	MetricTrigger_Operator_STATUS_GreaterThanOrEqual = MetricTrigger_Operator_STATUS("GreaterThanOrEqual")
	MetricTrigger_Operator_STATUS_LessThan           = MetricTrigger_Operator_STATUS("LessThan")
	MetricTrigger_Operator_STATUS_LessThanOrEqual    = MetricTrigger_Operator_STATUS("LessThanOrEqual")
	MetricTrigger_Operator_STATUS_NotEquals          = MetricTrigger_Operator_STATUS("NotEquals")
)

// Mapping from string to MetricTrigger_Operator_STATUS
var metricTrigger_Operator_STATUS_Values = map[string]MetricTrigger_Operator_STATUS{
	"equals":             MetricTrigger_Operator_STATUS_Equals,
	"greaterthan":        MetricTrigger_Operator_STATUS_GreaterThan,
	"greaterthanorequal": MetricTrigger_Operator_STATUS_GreaterThanOrEqual,
	"lessthan":           MetricTrigger_Operator_STATUS_LessThan,
	"lessthanorequal":    MetricTrigger_Operator_STATUS_LessThanOrEqual,
	"notequals":          MetricTrigger_Operator_STATUS_NotEquals,
}

// +kubebuilder:validation:Enum={"Average","Count","Max","Min","Sum"}
type MetricTrigger_Statistic_STATUS string

const (
	MetricTrigger_Statistic_STATUS_Average = MetricTrigger_Statistic_STATUS("Average")
	MetricTrigger_Statistic_STATUS_Count   = MetricTrigger_Statistic_STATUS("Count")
	MetricTrigger_Statistic_STATUS_Max     = MetricTrigger_Statistic_STATUS("Max")
	MetricTrigger_Statistic_STATUS_Min     = MetricTrigger_Statistic_STATUS("Min")
	MetricTrigger_Statistic_STATUS_Sum     = MetricTrigger_Statistic_STATUS("Sum")
)

// Mapping from string to MetricTrigger_Statistic_STATUS
var metricTrigger_Statistic_STATUS_Values = map[string]MetricTrigger_Statistic_STATUS{
	"average": MetricTrigger_Statistic_STATUS_Average,
	"count":   MetricTrigger_Statistic_STATUS_Count,
	"max":     MetricTrigger_Statistic_STATUS_Max,
	"min":     MetricTrigger_Statistic_STATUS_Min,
	"sum":     MetricTrigger_Statistic_STATUS_Sum,
}

// +kubebuilder:validation:Enum={"Average","Count","Last","Maximum","Minimum","Total"}
type MetricTrigger_TimeAggregation_STATUS string

const (
	MetricTrigger_TimeAggregation_STATUS_Average = MetricTrigger_TimeAggregation_STATUS("Average")
	MetricTrigger_TimeAggregation_STATUS_Count   = MetricTrigger_TimeAggregation_STATUS("Count")
	MetricTrigger_TimeAggregation_STATUS_Last    = MetricTrigger_TimeAggregation_STATUS("Last")
	MetricTrigger_TimeAggregation_STATUS_Maximum = MetricTrigger_TimeAggregation_STATUS("Maximum")
	MetricTrigger_TimeAggregation_STATUS_Minimum = MetricTrigger_TimeAggregation_STATUS("Minimum")
	MetricTrigger_TimeAggregation_STATUS_Total   = MetricTrigger_TimeAggregation_STATUS("Total")
)

// Mapping from string to MetricTrigger_TimeAggregation_STATUS
var metricTrigger_TimeAggregation_STATUS_Values = map[string]MetricTrigger_TimeAggregation_STATUS{
	"average": MetricTrigger_TimeAggregation_STATUS_Average,
	"count":   MetricTrigger_TimeAggregation_STATUS_Count,
	"last":    MetricTrigger_TimeAggregation_STATUS_Last,
	"maximum": MetricTrigger_TimeAggregation_STATUS_Maximum,
	"minimum": MetricTrigger_TimeAggregation_STATUS_Minimum,
	"total":   MetricTrigger_TimeAggregation_STATUS_Total,
}

// +kubebuilder:validation:Enum={"Decrease","Increase","None"}
type ScaleAction_Direction_STATUS string

const (
	ScaleAction_Direction_STATUS_Decrease = ScaleAction_Direction_STATUS("Decrease")
	ScaleAction_Direction_STATUS_Increase = ScaleAction_Direction_STATUS("Increase")
	ScaleAction_Direction_STATUS_None     = ScaleAction_Direction_STATUS("None")
)

// Mapping from string to ScaleAction_Direction_STATUS
var scaleAction_Direction_STATUS_Values = map[string]ScaleAction_Direction_STATUS{
	"decrease": ScaleAction_Direction_STATUS_Decrease,
	"increase": ScaleAction_Direction_STATUS_Increase,
	"none":     ScaleAction_Direction_STATUS_None,
}

// +kubebuilder:validation:Enum={"ChangeCount","ExactCount","PercentChangeCount","ServiceAllowedNextValue"}
type ScaleAction_Type_STATUS string

const (
	ScaleAction_Type_STATUS_ChangeCount             = ScaleAction_Type_STATUS("ChangeCount")
	ScaleAction_Type_STATUS_ExactCount              = ScaleAction_Type_STATUS("ExactCount")
	ScaleAction_Type_STATUS_PercentChangeCount      = ScaleAction_Type_STATUS("PercentChangeCount")
	ScaleAction_Type_STATUS_ServiceAllowedNextValue = ScaleAction_Type_STATUS("ServiceAllowedNextValue")
)

// Mapping from string to ScaleAction_Type_STATUS
var scaleAction_Type_STATUS_Values = map[string]ScaleAction_Type_STATUS{
	"changecount":             ScaleAction_Type_STATUS_ChangeCount,
	"exactcount":              ScaleAction_Type_STATUS_ExactCount,
	"percentchangecount":      ScaleAction_Type_STATUS_PercentChangeCount,
	"serviceallowednextvalue": ScaleAction_Type_STATUS_ServiceAllowedNextValue,
}

// Specifies an auto scale rule metric dimension.
type ScaleRuleMetricDimension_STATUS struct {
	// DimensionName: Name of the dimension.
	DimensionName *string `json:"DimensionName,omitempty"`

	// Operator: the dimension operator. Only 'Equals' and 'NotEquals' are supported. 'Equals' being equal to any of the
	// values. 'NotEquals' being not equal to all of the values
	Operator *ScaleRuleMetricDimension_Operator_STATUS `json:"Operator,omitempty"`

	// Values: list of dimension values. For example: ["App1","App2"].
	Values []string `json:"Values,omitempty"`
}

// +kubebuilder:validation:Enum={"Equals","NotEquals"}
type ScaleRuleMetricDimension_Operator_STATUS string

const (
	ScaleRuleMetricDimension_Operator_STATUS_Equals    = ScaleRuleMetricDimension_Operator_STATUS("Equals")
	ScaleRuleMetricDimension_Operator_STATUS_NotEquals = ScaleRuleMetricDimension_Operator_STATUS("NotEquals")
)

// Mapping from string to ScaleRuleMetricDimension_Operator_STATUS
var scaleRuleMetricDimension_Operator_STATUS_Values = map[string]ScaleRuleMetricDimension_Operator_STATUS{
	"equals":    ScaleRuleMetricDimension_Operator_STATUS_Equals,
	"notequals": ScaleRuleMetricDimension_Operator_STATUS_NotEquals,
}
