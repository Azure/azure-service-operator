// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

// Deprecated version of EventSubscription_STATUS. Use v1beta20200601.EventSubscription_STATUS instead
type EventSubscription_STATUSARM struct {
	Id         *string                                `json:"id,omitempty"`
	Name       *string                                `json:"name,omitempty"`
	Properties *EventSubscriptionProperties_STATUSARM `json:"properties,omitempty"`
	SystemData *SystemData_STATUSARM                  `json:"systemData,omitempty"`
	Type       *string                                `json:"type,omitempty"`
}

// Deprecated version of EventSubscriptionProperties_STATUS. Use v1beta20200601.EventSubscriptionProperties_STATUS instead
type EventSubscriptionProperties_STATUSARM struct {
	DeadLetterDestination *DeadLetterDestination_STATUSARM                        `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination_STATUSARM                 `json:"destination,omitempty"`
	EventDeliverySchema   *EventSubscriptionProperties_STATUS_EventDeliverySchema `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                                                 `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter_STATUSARM                      `json:"filter,omitempty"`
	Labels                []string                                                `json:"labels,omitempty"`
	ProvisioningState     *EventSubscriptionProperties_STATUS_ProvisioningState   `json:"provisioningState,omitempty"`
	RetryPolicy           *RetryPolicy_STATUSARM                                  `json:"retryPolicy,omitempty"`
	Topic                 *string                                                 `json:"topic,omitempty"`
}

// Deprecated version of DeadLetterDestination_STATUS. Use v1beta20200601.DeadLetterDestination_STATUS instead
type DeadLetterDestination_STATUSARM struct {
	EndpointType *DeadLetterDestination_STATUS_EndpointType `json:"endpointType,omitempty"`
}

// Deprecated version of EventSubscriptionDestination_STATUS. Use v1beta20200601.EventSubscriptionDestination_STATUS instead
type EventSubscriptionDestination_STATUSARM struct {
	EndpointType *EventSubscriptionDestination_STATUS_EndpointType `json:"endpointType,omitempty"`
}

// Deprecated version of EventSubscriptionFilter_STATUS. Use v1beta20200601.EventSubscriptionFilter_STATUS instead
type EventSubscriptionFilter_STATUSARM struct {
	AdvancedFilters        []AdvancedFilter_STATUSARM `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string                   `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                      `json:"isSubjectCaseSensitive,omitempty"`
	SubjectBeginsWith      *string                    `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string                    `json:"subjectEndsWith,omitempty"`
}

// Deprecated version of RetryPolicy_STATUS. Use v1beta20200601.RetryPolicy_STATUS instead
type RetryPolicy_STATUSARM struct {
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int `json:"maxDeliveryAttempts,omitempty"`
}

// Deprecated version of AdvancedFilter_STATUS. Use v1beta20200601.AdvancedFilter_STATUS instead
type AdvancedFilter_STATUSARM struct {
	Key          *string                             `json:"key,omitempty"`
	OperatorType *AdvancedFilter_STATUS_OperatorType `json:"operatorType,omitempty"`
}
