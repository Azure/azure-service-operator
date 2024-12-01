// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type EventSubscription_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of the event subscription.
	Properties *EventSubscriptionProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &EventSubscription_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (subscription EventSubscription_Spec) GetAPIVersion() string {
	return "2020-06-01"
}

// GetName returns the Name of the resource
func (subscription *EventSubscription_Spec) GetName() string {
	return subscription.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/eventSubscriptions"
func (subscription *EventSubscription_Spec) GetType() string {
	return "Microsoft.EventGrid/eventSubscriptions"
}

// Properties of the Event Subscription.
type EventSubscriptionProperties struct {
	// DeadLetterDestination: The DeadLetter destination of the event subscription.
	DeadLetterDestination *DeadLetterDestination `json:"deadLetterDestination,omitempty"`

	// Destination: Information about the destination where events have to be delivered for the event subscription.
	Destination *EventSubscriptionDestination `json:"destination,omitempty"`

	// EventDeliverySchema: The event delivery schema for the event subscription.
	EventDeliverySchema *EventSubscriptionProperties_EventDeliverySchema `json:"eventDeliverySchema,omitempty"`

	// ExpirationTimeUtc: Expiration time of the event subscription.
	ExpirationTimeUtc *string `json:"expirationTimeUtc,omitempty"`

	// Filter: Information about the filter for the event subscription.
	Filter *EventSubscriptionFilter `json:"filter,omitempty"`

	// Labels: List of user defined labels.
	Labels []string `json:"labels,omitempty"`

	// RetryPolicy: The retry policy for events. This can be used to configure maximum number of delivery attempts and time to
	// live for events.
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`
}

type DeadLetterDestination struct {
	// StorageBlob: Mutually exclusive with all other properties
	StorageBlob *StorageBlobDeadLetterDestination `json:"storageBlob,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DeadLetterDestination represents a discriminated union (JSON OneOf)
func (destination DeadLetterDestination) MarshalJSON() ([]byte, error) {
	if destination.StorageBlob != nil {
		return json.Marshal(destination.StorageBlob)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the DeadLetterDestination
func (destination *DeadLetterDestination) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "StorageBlob" {
		destination.StorageBlob = &StorageBlobDeadLetterDestination{}
		return json.Unmarshal(data, destination.StorageBlob)
	}

	// No error
	return nil
}

type EventSubscriptionDestination struct {
	// AzureFunction: Mutually exclusive with all other properties
	AzureFunction *AzureFunctionEventSubscriptionDestination `json:"azureFunction,omitempty"`

	// EventHub: Mutually exclusive with all other properties
	EventHub *EventHubEventSubscriptionDestination `json:"eventHub,omitempty"`

	// HybridConnection: Mutually exclusive with all other properties
	HybridConnection *HybridConnectionEventSubscriptionDestination `json:"hybridConnection,omitempty"`

	// ServiceBusQueue: Mutually exclusive with all other properties
	ServiceBusQueue *ServiceBusQueueEventSubscriptionDestination `json:"serviceBusQueue,omitempty"`

	// ServiceBusTopic: Mutually exclusive with all other properties
	ServiceBusTopic *ServiceBusTopicEventSubscriptionDestination `json:"serviceBusTopic,omitempty"`

	// StorageQueue: Mutually exclusive with all other properties
	StorageQueue *StorageQueueEventSubscriptionDestination `json:"storageQueue,omitempty"`

	// WebHook: Mutually exclusive with all other properties
	WebHook *WebHookEventSubscriptionDestination `json:"webHook,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because EventSubscriptionDestination represents a discriminated union (JSON OneOf)
func (destination EventSubscriptionDestination) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON unmarshals the EventSubscriptionDestination
func (destination *EventSubscriptionDestination) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "AzureFunction" {
		destination.AzureFunction = &AzureFunctionEventSubscriptionDestination{}
		return json.Unmarshal(data, destination.AzureFunction)
	}
	if discriminator == "EventHub" {
		destination.EventHub = &EventHubEventSubscriptionDestination{}
		return json.Unmarshal(data, destination.EventHub)
	}
	if discriminator == "HybridConnection" {
		destination.HybridConnection = &HybridConnectionEventSubscriptionDestination{}
		return json.Unmarshal(data, destination.HybridConnection)
	}
	if discriminator == "ServiceBusQueue" {
		destination.ServiceBusQueue = &ServiceBusQueueEventSubscriptionDestination{}
		return json.Unmarshal(data, destination.ServiceBusQueue)
	}
	if discriminator == "ServiceBusTopic" {
		destination.ServiceBusTopic = &ServiceBusTopicEventSubscriptionDestination{}
		return json.Unmarshal(data, destination.ServiceBusTopic)
	}
	if discriminator == "StorageQueue" {
		destination.StorageQueue = &StorageQueueEventSubscriptionDestination{}
		return json.Unmarshal(data, destination.StorageQueue)
	}
	if discriminator == "WebHook" {
		destination.WebHook = &WebHookEventSubscriptionDestination{}
		return json.Unmarshal(data, destination.WebHook)
	}

	// No error
	return nil
}

// Filter for the Event Subscription.
type EventSubscriptionFilter struct {
	// AdvancedFilters: An array of advanced filters that are used for filtering event subscriptions.
	AdvancedFilters []AdvancedFilter `json:"advancedFilters,omitempty"`

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

// +kubebuilder:validation:Enum={"CloudEventSchemaV1_0","CustomInputSchema","EventGridSchema"}
type EventSubscriptionProperties_EventDeliverySchema string

const (
	EventSubscriptionProperties_EventDeliverySchema_CloudEventSchemaV1_0 = EventSubscriptionProperties_EventDeliverySchema("CloudEventSchemaV1_0")
	EventSubscriptionProperties_EventDeliverySchema_CustomInputSchema    = EventSubscriptionProperties_EventDeliverySchema("CustomInputSchema")
	EventSubscriptionProperties_EventDeliverySchema_EventGridSchema      = EventSubscriptionProperties_EventDeliverySchema("EventGridSchema")
)

// Mapping from string to EventSubscriptionProperties_EventDeliverySchema
var eventSubscriptionProperties_EventDeliverySchema_Values = map[string]EventSubscriptionProperties_EventDeliverySchema{
	"cloudeventschemav1_0": EventSubscriptionProperties_EventDeliverySchema_CloudEventSchemaV1_0,
	"custominputschema":    EventSubscriptionProperties_EventDeliverySchema_CustomInputSchema,
	"eventgridschema":      EventSubscriptionProperties_EventDeliverySchema_EventGridSchema,
}

// Information about the retry policy for an event subscription.
type RetryPolicy struct {
	// EventTimeToLiveInMinutes: Time To Live (in minutes) for events.
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`

	// MaxDeliveryAttempts: Maximum number of delivery retry attempts for events.
	MaxDeliveryAttempts *int `json:"maxDeliveryAttempts,omitempty"`
}

type AdvancedFilter struct {
	// BoolEquals: Mutually exclusive with all other properties
	BoolEquals *BoolEqualsAdvancedFilter `json:"boolEquals,omitempty"`

	// NumberGreaterThan: Mutually exclusive with all other properties
	NumberGreaterThan *NumberGreaterThanAdvancedFilter `json:"numberGreaterThan,omitempty"`

	// NumberGreaterThanOrEquals: Mutually exclusive with all other properties
	NumberGreaterThanOrEquals *NumberGreaterThanOrEqualsAdvancedFilter `json:"numberGreaterThanOrEquals,omitempty"`

	// NumberIn: Mutually exclusive with all other properties
	NumberIn *NumberInAdvancedFilter `json:"numberIn,omitempty"`

	// NumberLessThan: Mutually exclusive with all other properties
	NumberLessThan *NumberLessThanAdvancedFilter `json:"numberLessThan,omitempty"`

	// NumberLessThanOrEquals: Mutually exclusive with all other properties
	NumberLessThanOrEquals *NumberLessThanOrEqualsAdvancedFilter `json:"numberLessThanOrEquals,omitempty"`

	// NumberNotIn: Mutually exclusive with all other properties
	NumberNotIn *NumberNotInAdvancedFilter `json:"numberNotIn,omitempty"`

	// StringBeginsWith: Mutually exclusive with all other properties
	StringBeginsWith *StringBeginsWithAdvancedFilter `json:"stringBeginsWith,omitempty"`

	// StringContains: Mutually exclusive with all other properties
	StringContains *StringContainsAdvancedFilter `json:"stringContains,omitempty"`

	// StringEndsWith: Mutually exclusive with all other properties
	StringEndsWith *StringEndsWithAdvancedFilter `json:"stringEndsWith,omitempty"`

	// StringIn: Mutually exclusive with all other properties
	StringIn *StringInAdvancedFilter `json:"stringIn,omitempty"`

	// StringNotIn: Mutually exclusive with all other properties
	StringNotIn *StringNotInAdvancedFilter `json:"stringNotIn,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because AdvancedFilter represents a discriminated union (JSON OneOf)
func (filter AdvancedFilter) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON unmarshals the AdvancedFilter
func (filter *AdvancedFilter) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["operatorType"]
	if discriminator == "BoolEquals" {
		filter.BoolEquals = &BoolEqualsAdvancedFilter{}
		return json.Unmarshal(data, filter.BoolEquals)
	}
	if discriminator == "NumberGreaterThan" {
		filter.NumberGreaterThan = &NumberGreaterThanAdvancedFilter{}
		return json.Unmarshal(data, filter.NumberGreaterThan)
	}
	if discriminator == "NumberGreaterThanOrEquals" {
		filter.NumberGreaterThanOrEquals = &NumberGreaterThanOrEqualsAdvancedFilter{}
		return json.Unmarshal(data, filter.NumberGreaterThanOrEquals)
	}
	if discriminator == "NumberIn" {
		filter.NumberIn = &NumberInAdvancedFilter{}
		return json.Unmarshal(data, filter.NumberIn)
	}
	if discriminator == "NumberLessThan" {
		filter.NumberLessThan = &NumberLessThanAdvancedFilter{}
		return json.Unmarshal(data, filter.NumberLessThan)
	}
	if discriminator == "NumberLessThanOrEquals" {
		filter.NumberLessThanOrEquals = &NumberLessThanOrEqualsAdvancedFilter{}
		return json.Unmarshal(data, filter.NumberLessThanOrEquals)
	}
	if discriminator == "NumberNotIn" {
		filter.NumberNotIn = &NumberNotInAdvancedFilter{}
		return json.Unmarshal(data, filter.NumberNotIn)
	}
	if discriminator == "StringBeginsWith" {
		filter.StringBeginsWith = &StringBeginsWithAdvancedFilter{}
		return json.Unmarshal(data, filter.StringBeginsWith)
	}
	if discriminator == "StringContains" {
		filter.StringContains = &StringContainsAdvancedFilter{}
		return json.Unmarshal(data, filter.StringContains)
	}
	if discriminator == "StringEndsWith" {
		filter.StringEndsWith = &StringEndsWithAdvancedFilter{}
		return json.Unmarshal(data, filter.StringEndsWith)
	}
	if discriminator == "StringIn" {
		filter.StringIn = &StringInAdvancedFilter{}
		return json.Unmarshal(data, filter.StringIn)
	}
	if discriminator == "StringNotIn" {
		filter.StringNotIn = &StringNotInAdvancedFilter{}
		return json.Unmarshal(data, filter.StringNotIn)
	}

	// No error
	return nil
}

type AzureFunctionEventSubscriptionDestination struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType AzureFunctionEventSubscriptionDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: Azure Function Properties of the event subscription destination.
	Properties *AzureFunctionEventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

type EventHubEventSubscriptionDestination struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType EventHubEventSubscriptionDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: Event Hub Properties of the event subscription destination.
	Properties *EventHubEventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

type HybridConnectionEventSubscriptionDestination struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType HybridConnectionEventSubscriptionDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: Hybrid connection Properties of the event subscription destination.
	Properties *HybridConnectionEventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

type ServiceBusQueueEventSubscriptionDestination struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType ServiceBusQueueEventSubscriptionDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: Service Bus Properties of the event subscription destination.
	Properties *ServiceBusQueueEventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

type ServiceBusTopicEventSubscriptionDestination struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType ServiceBusTopicEventSubscriptionDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: Service Bus Topic Properties of the event subscription destination.
	Properties *ServiceBusTopicEventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

type StorageBlobDeadLetterDestination struct {
	// EndpointType: Type of the endpoint for the dead letter destination
	EndpointType StorageBlobDeadLetterDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: The properties of the Storage Blob based deadletter destination
	Properties *StorageBlobDeadLetterDestinationProperties `json:"properties,omitempty"`
}

type StorageQueueEventSubscriptionDestination struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType StorageQueueEventSubscriptionDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: Storage Queue Properties of the event subscription destination.
	Properties *StorageQueueEventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

type WebHookEventSubscriptionDestination struct {
	// EndpointType: Type of the endpoint for the event subscription destination.
	EndpointType WebHookEventSubscriptionDestination_EndpointType `json:"endpointType,omitempty"`

	// Properties: WebHook Properties of the event subscription destination.
	Properties *WebHookEventSubscriptionDestinationProperties `json:"properties,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureFunction"}
type AzureFunctionEventSubscriptionDestination_EndpointType string

const AzureFunctionEventSubscriptionDestination_EndpointType_AzureFunction = AzureFunctionEventSubscriptionDestination_EndpointType("AzureFunction")

// Mapping from string to AzureFunctionEventSubscriptionDestination_EndpointType
var azureFunctionEventSubscriptionDestination_EndpointType_Values = map[string]AzureFunctionEventSubscriptionDestination_EndpointType{
	"azurefunction": AzureFunctionEventSubscriptionDestination_EndpointType_AzureFunction,
}

// The properties that represent the Azure Function destination of an event subscription.
type AzureFunctionEventSubscriptionDestinationProperties struct {
	// MaxEventsPerBatch: Maximum number of events per batch.
	MaxEventsPerBatch *int `json:"maxEventsPerBatch,omitempty"`

	// PreferredBatchSizeInKilobytes: Preferred batch size in Kilobytes.
	PreferredBatchSizeInKilobytes *int    `json:"preferredBatchSizeInKilobytes,omitempty"`
	ResourceId                    *string `json:"resourceId,omitempty"`
}

type BoolEqualsAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType BoolEqualsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Value: The boolean filter value.
	Value *bool `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"EventHub"}
type EventHubEventSubscriptionDestination_EndpointType string

const EventHubEventSubscriptionDestination_EndpointType_EventHub = EventHubEventSubscriptionDestination_EndpointType("EventHub")

// Mapping from string to EventHubEventSubscriptionDestination_EndpointType
var eventHubEventSubscriptionDestination_EndpointType_Values = map[string]EventHubEventSubscriptionDestination_EndpointType{
	"eventhub": EventHubEventSubscriptionDestination_EndpointType_EventHub,
}

// The properties for a event hub destination.
type EventHubEventSubscriptionDestinationProperties struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"HybridConnection"}
type HybridConnectionEventSubscriptionDestination_EndpointType string

const HybridConnectionEventSubscriptionDestination_EndpointType_HybridConnection = HybridConnectionEventSubscriptionDestination_EndpointType("HybridConnection")

// Mapping from string to HybridConnectionEventSubscriptionDestination_EndpointType
var hybridConnectionEventSubscriptionDestination_EndpointType_Values = map[string]HybridConnectionEventSubscriptionDestination_EndpointType{
	"hybridconnection": HybridConnectionEventSubscriptionDestination_EndpointType_HybridConnection,
}

// The properties for a hybrid connection destination.
type HybridConnectionEventSubscriptionDestinationProperties struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

type NumberGreaterThanAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberGreaterThanAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberGreaterThanOrEqualsAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberGreaterThanOrEqualsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberInAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []float64 `json:"values,omitempty"`
}

type NumberLessThanAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberLessThanAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberLessThanOrEqualsAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberLessThanOrEqualsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Value: The filter value.
	Value *float64 `json:"value,omitempty"`
}

type NumberNotInAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType NumberNotInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []float64 `json:"values,omitempty"`
}

// +kubebuilder:validation:Enum={"ServiceBusQueue"}
type ServiceBusQueueEventSubscriptionDestination_EndpointType string

const ServiceBusQueueEventSubscriptionDestination_EndpointType_ServiceBusQueue = ServiceBusQueueEventSubscriptionDestination_EndpointType("ServiceBusQueue")

// Mapping from string to ServiceBusQueueEventSubscriptionDestination_EndpointType
var serviceBusQueueEventSubscriptionDestination_EndpointType_Values = map[string]ServiceBusQueueEventSubscriptionDestination_EndpointType{
	"servicebusqueue": ServiceBusQueueEventSubscriptionDestination_EndpointType_ServiceBusQueue,
}

// The properties that represent the Service Bus destination of an event subscription.
type ServiceBusQueueEventSubscriptionDestinationProperties struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"ServiceBusTopic"}
type ServiceBusTopicEventSubscriptionDestination_EndpointType string

const ServiceBusTopicEventSubscriptionDestination_EndpointType_ServiceBusTopic = ServiceBusTopicEventSubscriptionDestination_EndpointType("ServiceBusTopic")

// Mapping from string to ServiceBusTopicEventSubscriptionDestination_EndpointType
var serviceBusTopicEventSubscriptionDestination_EndpointType_Values = map[string]ServiceBusTopicEventSubscriptionDestination_EndpointType{
	"servicebustopic": ServiceBusTopicEventSubscriptionDestination_EndpointType_ServiceBusTopic,
}

// The properties that represent the Service Bus Topic destination of an event subscription.
type ServiceBusTopicEventSubscriptionDestinationProperties struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"StorageBlob"}
type StorageBlobDeadLetterDestination_EndpointType string

const StorageBlobDeadLetterDestination_EndpointType_StorageBlob = StorageBlobDeadLetterDestination_EndpointType("StorageBlob")

// Mapping from string to StorageBlobDeadLetterDestination_EndpointType
var storageBlobDeadLetterDestination_EndpointType_Values = map[string]StorageBlobDeadLetterDestination_EndpointType{
	"storageblob": StorageBlobDeadLetterDestination_EndpointType_StorageBlob,
}

// Properties of the storage blob based dead letter destination.
type StorageBlobDeadLetterDestinationProperties struct {
	// BlobContainerName: The name of the Storage blob container that is the destination of the deadletter events
	BlobContainerName *string `json:"blobContainerName,omitempty"`
	ResourceId        *string `json:"resourceId,omitempty"`
}

// +kubebuilder:validation:Enum={"StorageQueue"}
type StorageQueueEventSubscriptionDestination_EndpointType string

const StorageQueueEventSubscriptionDestination_EndpointType_StorageQueue = StorageQueueEventSubscriptionDestination_EndpointType("StorageQueue")

// Mapping from string to StorageQueueEventSubscriptionDestination_EndpointType
var storageQueueEventSubscriptionDestination_EndpointType_Values = map[string]StorageQueueEventSubscriptionDestination_EndpointType{
	"storagequeue": StorageQueueEventSubscriptionDestination_EndpointType_StorageQueue,
}

// The properties for a storage queue destination.
type StorageQueueEventSubscriptionDestinationProperties struct {
	// QueueName: The name of the Storage queue under a storage account that is the destination of an event subscription.
	QueueName  *string `json:"queueName,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}

type StringBeginsWithAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringBeginsWithAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringContainsAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringContainsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringEndsWithAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringEndsWithAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringInAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

type StringNotInAdvancedFilter struct {
	// Key: The field/property in the event based on which you want to filter.
	Key *string `json:"key,omitempty"`

	// OperatorType: The operator type used for filtering, e.g., NumberIn, StringContains, BoolEquals and others.
	OperatorType StringNotInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`

	// Values: The set of filter values.
	Values []string `json:"values,omitempty"`
}

// +kubebuilder:validation:Enum={"WebHook"}
type WebHookEventSubscriptionDestination_EndpointType string

const WebHookEventSubscriptionDestination_EndpointType_WebHook = WebHookEventSubscriptionDestination_EndpointType("WebHook")

// Mapping from string to WebHookEventSubscriptionDestination_EndpointType
var webHookEventSubscriptionDestination_EndpointType_Values = map[string]WebHookEventSubscriptionDestination_EndpointType{
	"webhook": WebHookEventSubscriptionDestination_EndpointType_WebHook,
}

// Information about the webhook destination properties for an event subscription.
type WebHookEventSubscriptionDestinationProperties struct {
	// AzureActiveDirectoryApplicationIdOrUri: The Azure Active Directory Application ID or URI to get the access token that
	// will be included as the bearer token in delivery requests.
	AzureActiveDirectoryApplicationIdOrUri *string `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`

	// AzureActiveDirectoryTenantId: The Azure Active Directory Tenant ID to get the access token that will be included as the
	// bearer token in delivery requests.
	AzureActiveDirectoryTenantId *string `json:"azureActiveDirectoryTenantId,omitempty"`

	// EndpointUrl: The URL that represents the endpoint of the destination of an event subscription.
	EndpointUrl *string `json:"endpointUrl,omitempty"`

	// MaxEventsPerBatch: Maximum number of events per batch.
	MaxEventsPerBatch *int `json:"maxEventsPerBatch,omitempty"`

	// PreferredBatchSizeInKilobytes: Preferred batch size in Kilobytes.
	PreferredBatchSizeInKilobytes *int `json:"preferredBatchSizeInKilobytes,omitempty"`
}

// +kubebuilder:validation:Enum={"BoolEquals"}
type BoolEqualsAdvancedFilter_OperatorType string

const BoolEqualsAdvancedFilter_OperatorType_BoolEquals = BoolEqualsAdvancedFilter_OperatorType("BoolEquals")

// Mapping from string to BoolEqualsAdvancedFilter_OperatorType
var boolEqualsAdvancedFilter_OperatorType_Values = map[string]BoolEqualsAdvancedFilter_OperatorType{
	"boolequals": BoolEqualsAdvancedFilter_OperatorType_BoolEquals,
}

// +kubebuilder:validation:Enum={"NumberGreaterThan"}
type NumberGreaterThanAdvancedFilter_OperatorType string

const NumberGreaterThanAdvancedFilter_OperatorType_NumberGreaterThan = NumberGreaterThanAdvancedFilter_OperatorType("NumberGreaterThan")

// Mapping from string to NumberGreaterThanAdvancedFilter_OperatorType
var numberGreaterThanAdvancedFilter_OperatorType_Values = map[string]NumberGreaterThanAdvancedFilter_OperatorType{
	"numbergreaterthan": NumberGreaterThanAdvancedFilter_OperatorType_NumberGreaterThan,
}

// +kubebuilder:validation:Enum={"NumberGreaterThanOrEquals"}
type NumberGreaterThanOrEqualsAdvancedFilter_OperatorType string

const NumberGreaterThanOrEqualsAdvancedFilter_OperatorType_NumberGreaterThanOrEquals = NumberGreaterThanOrEqualsAdvancedFilter_OperatorType("NumberGreaterThanOrEquals")

// Mapping from string to NumberGreaterThanOrEqualsAdvancedFilter_OperatorType
var numberGreaterThanOrEqualsAdvancedFilter_OperatorType_Values = map[string]NumberGreaterThanOrEqualsAdvancedFilter_OperatorType{
	"numbergreaterthanorequals": NumberGreaterThanOrEqualsAdvancedFilter_OperatorType_NumberGreaterThanOrEquals,
}

// +kubebuilder:validation:Enum={"NumberIn"}
type NumberInAdvancedFilter_OperatorType string

const NumberInAdvancedFilter_OperatorType_NumberIn = NumberInAdvancedFilter_OperatorType("NumberIn")

// Mapping from string to NumberInAdvancedFilter_OperatorType
var numberInAdvancedFilter_OperatorType_Values = map[string]NumberInAdvancedFilter_OperatorType{
	"numberin": NumberInAdvancedFilter_OperatorType_NumberIn,
}

// +kubebuilder:validation:Enum={"NumberLessThan"}
type NumberLessThanAdvancedFilter_OperatorType string

const NumberLessThanAdvancedFilter_OperatorType_NumberLessThan = NumberLessThanAdvancedFilter_OperatorType("NumberLessThan")

// Mapping from string to NumberLessThanAdvancedFilter_OperatorType
var numberLessThanAdvancedFilter_OperatorType_Values = map[string]NumberLessThanAdvancedFilter_OperatorType{
	"numberlessthan": NumberLessThanAdvancedFilter_OperatorType_NumberLessThan,
}

// +kubebuilder:validation:Enum={"NumberLessThanOrEquals"}
type NumberLessThanOrEqualsAdvancedFilter_OperatorType string

const NumberLessThanOrEqualsAdvancedFilter_OperatorType_NumberLessThanOrEquals = NumberLessThanOrEqualsAdvancedFilter_OperatorType("NumberLessThanOrEquals")

// Mapping from string to NumberLessThanOrEqualsAdvancedFilter_OperatorType
var numberLessThanOrEqualsAdvancedFilter_OperatorType_Values = map[string]NumberLessThanOrEqualsAdvancedFilter_OperatorType{
	"numberlessthanorequals": NumberLessThanOrEqualsAdvancedFilter_OperatorType_NumberLessThanOrEquals,
}

// +kubebuilder:validation:Enum={"NumberNotIn"}
type NumberNotInAdvancedFilter_OperatorType string

const NumberNotInAdvancedFilter_OperatorType_NumberNotIn = NumberNotInAdvancedFilter_OperatorType("NumberNotIn")

// Mapping from string to NumberNotInAdvancedFilter_OperatorType
var numberNotInAdvancedFilter_OperatorType_Values = map[string]NumberNotInAdvancedFilter_OperatorType{
	"numbernotin": NumberNotInAdvancedFilter_OperatorType_NumberNotIn,
}

// +kubebuilder:validation:Enum={"StringBeginsWith"}
type StringBeginsWithAdvancedFilter_OperatorType string

const StringBeginsWithAdvancedFilter_OperatorType_StringBeginsWith = StringBeginsWithAdvancedFilter_OperatorType("StringBeginsWith")

// Mapping from string to StringBeginsWithAdvancedFilter_OperatorType
var stringBeginsWithAdvancedFilter_OperatorType_Values = map[string]StringBeginsWithAdvancedFilter_OperatorType{
	"stringbeginswith": StringBeginsWithAdvancedFilter_OperatorType_StringBeginsWith,
}

// +kubebuilder:validation:Enum={"StringContains"}
type StringContainsAdvancedFilter_OperatorType string

const StringContainsAdvancedFilter_OperatorType_StringContains = StringContainsAdvancedFilter_OperatorType("StringContains")

// Mapping from string to StringContainsAdvancedFilter_OperatorType
var stringContainsAdvancedFilter_OperatorType_Values = map[string]StringContainsAdvancedFilter_OperatorType{
	"stringcontains": StringContainsAdvancedFilter_OperatorType_StringContains,
}

// +kubebuilder:validation:Enum={"StringEndsWith"}
type StringEndsWithAdvancedFilter_OperatorType string

const StringEndsWithAdvancedFilter_OperatorType_StringEndsWith = StringEndsWithAdvancedFilter_OperatorType("StringEndsWith")

// Mapping from string to StringEndsWithAdvancedFilter_OperatorType
var stringEndsWithAdvancedFilter_OperatorType_Values = map[string]StringEndsWithAdvancedFilter_OperatorType{
	"stringendswith": StringEndsWithAdvancedFilter_OperatorType_StringEndsWith,
}

// +kubebuilder:validation:Enum={"StringIn"}
type StringInAdvancedFilter_OperatorType string

const StringInAdvancedFilter_OperatorType_StringIn = StringInAdvancedFilter_OperatorType("StringIn")

// Mapping from string to StringInAdvancedFilter_OperatorType
var stringInAdvancedFilter_OperatorType_Values = map[string]StringInAdvancedFilter_OperatorType{
	"stringin": StringInAdvancedFilter_OperatorType_StringIn,
}

// +kubebuilder:validation:Enum={"StringNotIn"}
type StringNotInAdvancedFilter_OperatorType string

const StringNotInAdvancedFilter_OperatorType_StringNotIn = StringNotInAdvancedFilter_OperatorType("StringNotIn")

// Mapping from string to StringNotInAdvancedFilter_OperatorType
var stringNotInAdvancedFilter_OperatorType_Values = map[string]StringNotInAdvancedFilter_OperatorType{
	"stringnotin": StringNotInAdvancedFilter_OperatorType_StringNotIn,
}
