// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import "encoding/json"

type EventSubscription_STATUS_ARM struct {
	// Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	// Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the event subscription.
	Properties *EventSubscriptionProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to Event Subscription resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

type EventSubscriptionProperties_STATUS_ARM struct {
	// DeadLetterDestination: The DeadLetter destination of the event subscription.
	DeadLetterDestination *DeadLetterDestination_STATUS_ARM `json:"deadLetterDestination,omitempty"`

	// Destination: Information about the destination where events have to be delivered for the event subscription.
	Destination *EventSubscriptionDestination_STATUS_ARM `json:"destination,omitempty"`

	// EventDeliverySchema: The event delivery schema for the event subscription.
	EventDeliverySchema *EventSubscriptionProperties_EventDeliverySchema_STATUS `json:"eventDeliverySchema,omitempty"`

	// ExpirationTimeUtc: Expiration time of the event subscription.
	ExpirationTimeUtc *string `json:"expirationTimeUtc,omitempty"`

	// Filter: Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter_STATUS_ARM `json:"filter,omitempty"`

	// Labels: List of user defined labels.
	Labels []string `json:"labels,omitempty"`

	// ProvisioningState: Provisioning state of the event subscription.
	ProvisioningState *EventSubscriptionProperties_ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// RetryPolicy: The retry policy for events. This can be used to configure maximum number of delivery attempts and time to
	// live for events.
	RetryPolicy *RetryPolicy_STATUS_ARM `json:"retryPolicy,omitempty"`

	// Topic: Name of the topic of the event subscription.
	Topic *string `json:"topic,omitempty"`
}

type DeadLetterDestination_STATUS_ARM struct {
	// StorageBlob: Mutually exclusive with all other properties
	StorageBlob *StorageBlobDeadLetterDestination_STATUS_ARM `json:"storageBlob,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DeadLetterDestination_STATUS_ARM represents a discriminated union (JSON OneOf)
func (destination DeadLetterDestination_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if destination.StorageBlob != nil {
		return json.Marshal(destination.StorageBlob)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the DeadLetterDestination_STATUS_ARM
func (destination *DeadLetterDestination_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "StorageBlob" {
		destination.StorageBlob = &StorageBlobDeadLetterDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.StorageBlob)
	}

	// No error
	return nil
}

type EventSubscriptionDestination_STATUS_ARM struct {
	// AzureFunction: Mutually exclusive with all other properties
	AzureFunction *AzureFunctionEventSubscriptionDestination_STATUS_ARM `json:"azureFunction,omitempty"`

	// EventHub: Mutually exclusive with all other properties
	EventHub *EventHubEventSubscriptionDestination_STATUS_ARM `json:"eventHub,omitempty"`

	// HybridConnection: Mutually exclusive with all other properties
	HybridConnection *HybridConnectionEventSubscriptionDestination_STATUS_ARM `json:"hybridConnection,omitempty"`

	// ServiceBusQueue: Mutually exclusive with all other properties
	ServiceBusQueue *ServiceBusQueueEventSubscriptionDestination_STATUS_ARM `json:"serviceBusQueue,omitempty"`

	// ServiceBusTopic: Mutually exclusive with all other properties
	ServiceBusTopic *ServiceBusTopicEventSubscriptionDestination_STATUS_ARM `json:"serviceBusTopic,omitempty"`

	// StorageQueue: Mutually exclusive with all other properties
	StorageQueue *StorageQueueEventSubscriptionDestination_STATUS_ARM `json:"storageQueue,omitempty"`

	// WebHook: Mutually exclusive with all other properties
	WebHook *WebHookEventSubscriptionDestination_STATUS_ARM `json:"webHook,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because EventSubscriptionDestination_STATUS_ARM represents a discriminated union (JSON OneOf)
func (destination EventSubscriptionDestination_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if destination.AzureFunction != nil {
		return json.Marshal(destination.AzureFunction)
	}
	if destination.EventHub != nil {
		return json.Marshal(destination.EventHub)
	}
	if destination.HybridConnection != nil {
		return json.Marshal(destination.HybridConnection)
	}
	if destination.ServiceBusQueue != nil {
		return json.Marshal(destination.ServiceBusQueue)
	}
	if destination.ServiceBusTopic != nil {
		return json.Marshal(destination.ServiceBusTopic)
	}
	if destination.StorageQueue != nil {
		return json.Marshal(destination.StorageQueue)
	}
	if destination.WebHook != nil {
		return json.Marshal(destination.WebHook)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the EventSubscriptionDestination_STATUS_ARM
func (destination *EventSubscriptionDestination_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "AzureFunction" {
		destination.AzureFunction = &AzureFunctionEventSubscriptionDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.AzureFunction)
	}
	if discriminator == "EventHub" {
		destination.EventHub = &EventHubEventSubscriptionDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.EventHub)
	}
	if discriminator == "HybridConnection" {
		destination.HybridConnection = &HybridConnectionEventSubscriptionDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.HybridConnection)
	}
	if discriminator == "ServiceBusQueue" {
		destination.ServiceBusQueue = &ServiceBusQueueEventSubscriptionDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.ServiceBusQueue)
	}
	if discriminator == "ServiceBusTopic" {
		destination.ServiceBusTopic = &ServiceBusTopicEventSubscriptionDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.ServiceBusTopic)
	}
	if discriminator == "StorageQueue" {
		destination.StorageQueue = &StorageQueueEventSubscriptionDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.StorageQueue)
	}
	if discriminator == "WebHook" {
		destination.WebHook = &WebHookEventSubscriptionDestination_STATUS_ARM{}
		return json.Unmarshal(data, destination.WebHook)
	}

	// No error
	return nil
}

type EventSubscriptionFilter_STATUS_ARM struct {
	// AdvancedFilters: An array of advanced filters that are used for filtering event subscriptions.
	AdvancedFilters []AdvancedFilter_STATUS_ARM `json:"advancedFilters,omitempty"`

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

type RetryPolicy_STATUS_ARM struct {
	// EventTimeToLiveInMinutes: Time To Live (in minutes) for events.
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`

	// MaxDeliveryAttempts: Maximum number of delivery retry attempts for events.
	MaxDeliveryAttempts *int `json:"maxDeliveryAttempts,omitempty"`
}

type AdvancedFilter_STATUS_ARM struct {
	// BoolEquals: Mutually exclusive with all other properties
	BoolEquals *BoolEqualsAdvancedFilter_STATUS_ARM `json:"boolEquals,omitempty"`

	// NumberGreaterThan: Mutually exclusive with all other properties
	NumberGreaterThan *NumberGreaterThanAdvancedFilter_STATUS_ARM `json:"numberGreaterThan,omitempty"`

	// NumberGreaterThanOrEquals: Mutually exclusive with all other properties
	NumberGreaterThanOrEquals *NumberGreaterThanOrEqualsAdvancedFilter_STATUS_ARM `json:"numberGreaterThanOrEquals,omitempty"`

	// NumberIn: Mutually exclusive with all other properties
	NumberIn *NumberInAdvancedFilter_STATUS_ARM `json:"numberIn,omitempty"`

	// NumberLessThan: Mutually exclusive with all other properties
	NumberLessThan *NumberLessThanAdvancedFilter_STATUS_ARM `json:"numberLessThan,omitempty"`

	// NumberLessThanOrEquals: Mutually exclusive with all other properties
	NumberLessThanOrEquals *NumberLessThanOrEqualsAdvancedFilter_STATUS_ARM `json:"numberLessThanOrEquals,omitempty"`

	// NumberNotIn: Mutually exclusive with all other properties
	NumberNotIn *NumberNotInAdvancedFilter_STATUS_ARM `json:"numberNotIn,omitempty"`

	// StringBeginsWith: Mutually exclusive with all other properties
	StringBeginsWith *StringBeginsWithAdvancedFilter_STATUS_ARM `json:"stringBeginsWith,omitempty"`

	// StringContains: Mutually exclusive with all other properties
	StringContains *StringContainsAdvancedFilter_STATUS_ARM `json:"stringContains,omitempty"`

	// StringEndsWith: Mutually exclusive with all other properties
	StringEndsWith *StringEndsWithAdvancedFilter_STATUS_ARM `json:"stringEndsWith,omitempty"`

	// StringIn: Mutually exclusive with all other properties
	StringIn *StringInAdvancedFilter_STATUS_ARM `json:"stringIn,omitempty"`

	// StringNotIn: Mutually exclusive with all other properties
	StringNotIn *StringNotInAdvancedFilter_STATUS_ARM `json:"stringNotIn,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because AdvancedFilter_STATUS_ARM represents a discriminated union (JSON OneOf)
func (filter AdvancedFilter_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if filter.BoolEquals != nil {
		return json.Marshal(filter.BoolEquals)
	}
	if filter.NumberGreaterThan != nil {
		return json.Marshal(filter.NumberGreaterThan)
	}
	if filter.NumberGreaterThanOrEquals != nil {
		return json.Marshal(filter.NumberGreaterThanOrEquals)
	}
	if filter.NumberIn != nil {
		return json.Marshal(filter.NumberIn)
	}
	if filter.NumberLessThan != nil {
		return json.Marshal(filter.NumberLessThan)
	}
	if filter.NumberLessThanOrEquals != nil {
		return json.Marshal(filter.NumberLessThanOrEquals)
	}
	if filter.NumberNotIn != nil {
		return json.Marshal(filter.NumberNotIn)
	}
	if filter.StringBeginsWith != nil {
		return json.Marshal(filter.StringBeginsWith)
	}
	if filter.StringContains != nil {
		return json.Marshal(filter.StringContains)
	}
	if filter.StringEndsWith != nil {
		return json.Marshal(filter.StringEndsWith)
	}
	if filter.StringIn != nil {
		return json.Marshal(filter.StringIn)
	}
	if filter.StringNotIn != nil {
		return json.Marshal(filter.StringNotIn)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the AdvancedFilter_STATUS_ARM
func (filter *AdvancedFilter_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["operatorType"]
	if discriminator == "BoolEquals" {
		filter.BoolEquals = &BoolEqualsAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.BoolEquals)
	}
	if discriminator == "NumberGreaterThan" {
		filter.NumberGreaterThan = &NumberGreaterThanAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.NumberGreaterThan)
	}
	if discriminator == "NumberGreaterThanOrEquals" {
		filter.NumberGreaterThanOrEquals = &NumberGreaterThanOrEqualsAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.NumberGreaterThanOrEquals)
	}
	if discriminator == "NumberIn" {
		filter.NumberIn = &NumberInAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.NumberIn)
	}
	if discriminator == "NumberLessThan" {
		filter.NumberLessThan = &NumberLessThanAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.NumberLessThan)
	}
	if discriminator == "NumberLessThanOrEquals" {
		filter.NumberLessThanOrEquals = &NumberLessThanOrEqualsAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.NumberLessThanOrEquals)
	}
	if discriminator == "NumberNotIn" {
		filter.NumberNotIn = &NumberNotInAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.NumberNotIn)
	}
	if discriminator == "StringBeginsWith" {
		filter.StringBeginsWith = &StringBeginsWithAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.StringBeginsWith)
	}
	if discriminator == "StringContains" {
		filter.StringContains = &StringContainsAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.StringContains)
	}
	if discriminator == "StringEndsWith" {
		filter.StringEndsWith = &StringEndsWithAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.StringEndsWith)
	}
	if discriminator == "StringIn" {
		filter.StringIn = &StringInAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.StringIn)
	}
	if discriminator == "StringNotIn" {
		filter.StringNotIn = &StringNotInAdvancedFilter_STATUS_ARM{}
		return json.Unmarshal(data, filter.StringNotIn)
	}

	// No error
	return nil
}

type AzureFunctionEventSubscriptionDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType AzureFunctionEventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: Azure Function Properties of the event subscription destination.
	Properties *AzureFunctionEventSubscriptionDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type EventHubEventSubscriptionDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType EventHubEventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: Event Hub Properties of the event subscription destination.
	Properties *EventHubEventSubscriptionDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type HybridConnectionEventSubscriptionDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType HybridConnectionEventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: Hybrid connection Properties of the event subscription destination.
	Properties *HybridConnectionEventSubscriptionDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type ServiceBusQueueEventSubscriptionDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType ServiceBusQueueEventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: Service Bus Properties of the event subscription destination.
	Properties *ServiceBusQueueEventSubscriptionDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type ServiceBusTopicEventSubscriptionDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType ServiceBusTopicEventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: Service Bus Topic Properties of the event subscription destination.
	Properties *ServiceBusTopicEventSubscriptionDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type StorageBlobDeadLetterDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the dead letter destination
	EndpointType StorageBlobDeadLetterDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: The properties of the Storage Blob based deadletter destination
	Properties *StorageBlobDeadLetterDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type StorageQueueEventSubscriptionDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType StorageQueueEventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: Storage Queue Properties of the event subscription destination.
	Properties *StorageQueueEventSubscriptionDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type WebHookEventSubscriptionDestination_STATUS_ARM struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType WebHookEventSubscriptionDestination_EndpointType_STATUS `json:"endpointType,omitempty"`

	// Properties: WebHook Properties of the event subscription destination.
	Properties *WebHookEventSubscriptionDestinationProperties_STATUS_ARM `json:"properties,omitempty"`
}

type AzureFunctionEventSubscriptionDestinationProperties_STATUS_ARM struct {
	// MaxEventsPerBatch: Maximum number of events per batch.
	MaxEventsPerBatch *int `json:"maxEventsPerBatch,omitempty"`

	// PreferredBatchSizeInKilobytes: Preferred batch size in Kilobytes.
	PreferredBatchSizeInKilobytes *int `json:"preferredBatchSizeInKilobytes,omitempty"`

	// ResourceId: The Azure Resource Id that represents the endpoint of the Azure Function destination of an event
	// subscription.
	ResourceId *string `json:"resourceId,omitempty"`
}

type BoolEqualsAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType BoolEqualsAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Value: The boolean filter value.
	Value *bool `json:"value,omitempty"`
}

type EventHubEventSubscriptionDestinationProperties_STATUS_ARM struct {
	// ResourceId: The Azure Resource Id that represents the endpoint of an Event Hub destination of an event subscription.
	ResourceId *string `json:"resourceId,omitempty"`
}

type HybridConnectionEventSubscriptionDestinationProperties_STATUS_ARM struct {
	// ResourceId: The Azure Resource ID of an hybrid connection that is the destination of an event subscription.
	ResourceId *string `json:"resourceId,omitempty"`
}

type NumberGreaterThanAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberGreaterThanAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberGreaterThanOrEqualsAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberGreaterThanOrEqualsAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberInAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberInAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []float64 `json:"values,omitempty"`
}

type NumberLessThanAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberLessThanAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberLessThanOrEqualsAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberLessThanOrEqualsAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberNotInAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberNotInAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []float64 `json:"values,omitempty"`
}

type ServiceBusQueueEventSubscriptionDestinationProperties_STATUS_ARM struct {
	// ResourceId: The Azure Resource Id that represents the endpoint of the Service Bus destination of an event subscription.
	ResourceId *string `json:"resourceId,omitempty"`
}

type ServiceBusTopicEventSubscriptionDestinationProperties_STATUS_ARM struct {
	// ResourceId: The Azure Resource Id that represents the endpoint of the Service Bus Topic destination of an event
	// subscription.
	ResourceId *string `json:"resourceId,omitempty"`
}

type StorageBlobDeadLetterDestinationProperties_STATUS_ARM struct {
	// BlobContainerName: The name of the Storage blob container that is the destination of the deadletter events
	BlobContainerName *string `json:"blobContainerName,omitempty"`

	// ResourceId: The Azure Resource ID of the storage account that is the destination of the deadletter events
	ResourceId *string `json:"resourceId,omitempty"`
}

type StorageQueueEventSubscriptionDestinationProperties_STATUS_ARM struct {
	// QueueName: The name of the Storage queue under a storage account that is the destination of an event subscription.
	QueueName *string `json:"queueName,omitempty"`

	// ResourceId: The Azure Resource ID of the storage account that contains the queue that is the destination of an event
	// subscription.
	ResourceId *string `json:"resourceId,omitempty"`
}

type StringBeginsWithAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringBeginsWithAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringContainsAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringContainsAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringEndsWithAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringEndsWithAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringInAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringInAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringNotInAdvancedFilter_STATUS_ARM struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringNotInAdvancedFilter_OperatorType_STATUS `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type WebHookEventSubscriptionDestinationProperties_STATUS_ARM struct {
	// AzureActiveDirectoryApplicationIdOrUri: The Azure Active Directory Application ID or URI to get the access token that
	// will be included as the bearer token in delivery requests.
	AzureActiveDirectoryApplicationIdOrUri *string `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`

	// AzureActiveDirectoryTenantId: The Azure Active Directory Tenant ID to get the access token that will be included as the
	// bearer token in delivery requests.
	AzureActiveDirectoryTenantId *string `json:"azureActiveDirectoryTenantId,omitempty"`

	// EndpointBaseUrl: The base URL that represents the endpoint of the destination of an event subscription.
	EndpointBaseUrl *string `json:"endpointBaseUrl,omitempty"`

	// MaxEventsPerBatch: Maximum number of events per batch.
	MaxEventsPerBatch *int `json:"maxEventsPerBatch,omitempty"`

	// PreferredBatchSizeInKilobytes: Preferred batch size in Kilobytes.
	PreferredBatchSizeInKilobytes *int `json:"preferredBatchSizeInKilobytes,omitempty"`
}
