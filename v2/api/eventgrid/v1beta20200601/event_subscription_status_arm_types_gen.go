// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

type EventSubscription_STATUSARM struct {
	// Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	// Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the event subscription.
	Properties *EventSubscriptionProperties_STATUSARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to Event Subscription resource.
	SystemData *SystemData_STATUSARM `json:"systemData,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type EventSubscriptionProperties_STATUSARM struct {
	// DeadLetterDestination: The DeadLetter destination of the event subscription.
	DeadLetterDestination *DeadLetterDestination_STATUSARM `json:"deadLetterDestination,omitempty"`

	// Destination: Information about the destination where events have to be delivered for the event subscription.
	Destination *EventSubscriptionDestination_STATUSARM `json:"destination,omitempty"`

	// EventDeliverySchema: The event delivery schema for the event subscription.
<<<<<<< HEAD
	EventDeliverySchema *EventSubscriptionProperties_EventDeliverySchema_STATUS `json:"eventDeliverySchema,omitempty"`
=======
	EventDeliverySchema *EventSubscriptionProperties_STATUS_EventDeliverySchema `json:"eventDeliverySchema,omitempty"`
>>>>>>> main

	// ExpirationTimeUtc: Expiration time of the event subscription.
	ExpirationTimeUtc *string `json:"expirationTimeUtc,omitempty"`

	// Filter: Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter_STATUSARM `json:"filter,omitempty"`

	// Labels: List of user defined labels.
	Labels []string `json:"labels,omitempty"`

	// ProvisioningState: Provisioning state of the event subscription.
<<<<<<< HEAD
	ProvisioningState *EventSubscriptionProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`
=======
	ProvisioningState *EventSubscriptionProperties_STATUS_ProvisioningState `json:"provisioningState,omitempty"`
>>>>>>> main

	// RetryPolicy: The retry policy for events. This can be used to configure maximum number of delivery attempts and time to
	// live for events.
	RetryPolicy *RetryPolicy_STATUSARM `json:"retryPolicy,omitempty"`

	// Topic: Name of the topic of the event subscription.
	Topic *string `json:"topic,omitempty"`
}

type DeadLetterDestination_STATUSARM struct {
	// EndpointType: Type of the endpoint for the dead letter destination
<<<<<<< HEAD
	EndpointType *DeadLetterDestination_EndpointType_STATUS `json:"endpointType,omitempty"`
=======
	EndpointType *DeadLetterDestination_STATUS_EndpointType `json:"endpointType,omitempty"`
>>>>>>> main
}

type EventSubscriptionDestination_STATUSARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
<<<<<<< HEAD
	EndpointType *EventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`
=======
	EndpointType *EventSubscriptionDestination_STATUS_EndpointType `json:"endpointType,omitempty"`
>>>>>>> main
}

type EventSubscriptionFilter_STATUSARM struct {
	// AdvancedFilters: An array of advanced filters that are used for filtering event subscriptions.
	AdvancedFilters []AdvancedFilter_STATUSARM `json:"advancedFilters,omitempty"`

	// IncludedEventTypes: A list of applicable event types that need to be part of the event subscription. If it is desired to
	// subscribe to all default event types, set the IncludedEventTypes to null.
	IncludedEventTypes []string `json:"includedEventTypes,omitempty"`

	// IsSubjectCaseSensitive: Specifies if the SubjectBeginsWith and SubjectEndsWith properties of the filter
	// should be compared in a case sensitive manner.
	IsSubjectCaseSensitive *bool `json:"isSubjectCaseSensitive,omitempty"`

	// SubjectBeginsWith: An optional string to filter events for an event subscription based on a resource path prefix.
	// The format of this depends on the publisher of the events.
	// Wildcard characters are not supported in this path.
	SubjectBeginsWith *string `json:"subjectBeginsWith,omitempty"`

	// SubjectEndsWith: An optional string to filter events for an event subscription based on a resource path suffix.
	// Wildcard characters are not supported in this path.
	SubjectEndsWith *string `json:"subjectEndsWith,omitempty"`
}

type RetryPolicy_STATUSARM struct {
	// EventTimeToLiveInMinutes: Time To Live (in minutes) for events.
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`

	// MaxDeliveryAttempts: Maximum number of delivery retry attempts for events.
	MaxDeliveryAttempts *int `json:"maxDeliveryAttempts,omitempty"`
}

type AdvancedFilter_STATUSARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
<<<<<<< HEAD
	OperatorType *AdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`
=======
	OperatorType *AdvancedFilter_STATUS_OperatorType `json:"operatorType,omitempty"`
>>>>>>> main
}
