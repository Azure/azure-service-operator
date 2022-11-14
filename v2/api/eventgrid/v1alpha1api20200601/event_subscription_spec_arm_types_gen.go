// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Deprecated version of EventSubscription_Spec. Use v1beta20200601.EventSubscription_Spec instead
type EventSubscription_Spec_ARM struct {
	Name       string                           `json:"name,omitempty"`
	Properties *EventSubscriptionProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &EventSubscription_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (subscription EventSubscription_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (subscription *EventSubscription_Spec_ARM) GetName() string {
	return subscription.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/eventSubscriptions"
func (subscription *EventSubscription_Spec_ARM) GetType() string {
	return "Microsoft.EventGrid/eventSubscriptions"
}

// Deprecated version of EventSubscriptionProperties. Use v1beta20200601.EventSubscriptionProperties instead
type EventSubscriptionProperties_ARM struct {
	DeadLetterDestination *DeadLetterDestination_ARM                       `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination_ARM                `json:"destination,omitempty"`
	EventDeliverySchema   *EventSubscriptionProperties_EventDeliverySchema `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                                          `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter_ARM                     `json:"filter,omitempty"`
	Labels                []string                                         `json:"labels,omitempty"`
	RetryPolicy           *RetryPolicy_ARM                                 `json:"retryPolicy,omitempty"`
}

// Deprecated version of DeadLetterDestination. Use v1beta20200601.DeadLetterDestination instead
type DeadLetterDestination_ARM struct {
	StorageBlob *StorageBlobDeadLetterDestination_ARM `json:"storageBlob,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DeadLetterDestination_ARM represents a discriminated union (JSON OneOf)
func (destination DeadLetterDestination_ARM) MarshalJSON() ([]byte, error) {
	if destination.StorageBlob != nil {
		return json.Marshal(destination.StorageBlob)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the DeadLetterDestination_ARM
func (destination *DeadLetterDestination_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "StorageBlob" {
		destination.StorageBlob = &StorageBlobDeadLetterDestination_ARM{}
		return json.Unmarshal(data, destination.StorageBlob)
	}

	// No error
	return nil
}

// Deprecated version of EventSubscriptionDestination. Use v1beta20200601.EventSubscriptionDestination instead
type EventSubscriptionDestination_ARM struct {
	AzureFunction    *AzureFunctionEventSubscriptionDestination_ARM    `json:"azureFunction,omitempty"`
	EventHub         *EventHubEventSubscriptionDestination_ARM         `json:"eventHub,omitempty"`
	HybridConnection *HybridConnectionEventSubscriptionDestination_ARM `json:"hybridConnection,omitempty"`
	ServiceBusQueue  *ServiceBusQueueEventSubscriptionDestination_ARM  `json:"serviceBusQueue,omitempty"`
	ServiceBusTopic  *ServiceBusTopicEventSubscriptionDestination_ARM  `json:"serviceBusTopic,omitempty"`
	StorageQueue     *StorageQueueEventSubscriptionDestination_ARM     `json:"storageQueue,omitempty"`
	WebHook          *WebHookEventSubscriptionDestination_ARM          `json:"webHook,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because EventSubscriptionDestination_ARM represents a discriminated union (JSON OneOf)
func (destination EventSubscriptionDestination_ARM) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON unmarshals the EventSubscriptionDestination_ARM
func (destination *EventSubscriptionDestination_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["endpointType"]
	if discriminator == "AzureFunction" {
		destination.AzureFunction = &AzureFunctionEventSubscriptionDestination_ARM{}
		return json.Unmarshal(data, destination.AzureFunction)
	}
	if discriminator == "EventHub" {
		destination.EventHub = &EventHubEventSubscriptionDestination_ARM{}
		return json.Unmarshal(data, destination.EventHub)
	}
	if discriminator == "HybridConnection" {
		destination.HybridConnection = &HybridConnectionEventSubscriptionDestination_ARM{}
		return json.Unmarshal(data, destination.HybridConnection)
	}
	if discriminator == "ServiceBusQueue" {
		destination.ServiceBusQueue = &ServiceBusQueueEventSubscriptionDestination_ARM{}
		return json.Unmarshal(data, destination.ServiceBusQueue)
	}
	if discriminator == "ServiceBusTopic" {
		destination.ServiceBusTopic = &ServiceBusTopicEventSubscriptionDestination_ARM{}
		return json.Unmarshal(data, destination.ServiceBusTopic)
	}
	if discriminator == "StorageQueue" {
		destination.StorageQueue = &StorageQueueEventSubscriptionDestination_ARM{}
		return json.Unmarshal(data, destination.StorageQueue)
	}
	if discriminator == "WebHook" {
		destination.WebHook = &WebHookEventSubscriptionDestination_ARM{}
		return json.Unmarshal(data, destination.WebHook)
	}

	// No error
	return nil
}

// Deprecated version of EventSubscriptionFilter. Use v1beta20200601.EventSubscriptionFilter instead
type EventSubscriptionFilter_ARM struct {
	AdvancedFilters        []AdvancedFilter_ARM `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string             `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                `json:"isSubjectCaseSensitive,omitempty"`
	SubjectBeginsWith      *string              `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string              `json:"subjectEndsWith,omitempty"`
}

// Deprecated version of RetryPolicy. Use v1beta20200601.RetryPolicy instead
type RetryPolicy_ARM struct {
	EventTimeToLiveInMinutes *int `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int `json:"maxDeliveryAttempts,omitempty"`
}

// Deprecated version of AdvancedFilter. Use v1beta20200601.AdvancedFilter instead
type AdvancedFilter_ARM struct {
	BoolEquals                *BoolEqualsAdvancedFilter_ARM                `json:"boolEquals,omitempty"`
	NumberGreaterThan         *NumberGreaterThanAdvancedFilter_ARM         `json:"numberGreaterThan,omitempty"`
	NumberGreaterThanOrEquals *NumberGreaterThanOrEqualsAdvancedFilter_ARM `json:"numberGreaterThanOrEquals,omitempty"`
	NumberIn                  *NumberInAdvancedFilter_ARM                  `json:"numberIn,omitempty"`
	NumberLessThan            *NumberLessThanAdvancedFilter_ARM            `json:"numberLessThan,omitempty"`
	NumberLessThanOrEquals    *NumberLessThanOrEqualsAdvancedFilter_ARM    `json:"numberLessThanOrEquals,omitempty"`
	NumberNotIn               *NumberNotInAdvancedFilter_ARM               `json:"numberNotIn,omitempty"`
	StringBeginsWith          *StringBeginsWithAdvancedFilter_ARM          `json:"stringBeginsWith,omitempty"`
	StringContains            *StringContainsAdvancedFilter_ARM            `json:"stringContains,omitempty"`
	StringEndsWith            *StringEndsWithAdvancedFilter_ARM            `json:"stringEndsWith,omitempty"`
	StringIn                  *StringInAdvancedFilter_ARM                  `json:"stringIn,omitempty"`
	StringNotIn               *StringNotInAdvancedFilter_ARM               `json:"stringNotIn,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because AdvancedFilter_ARM represents a discriminated union (JSON OneOf)
func (filter AdvancedFilter_ARM) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON unmarshals the AdvancedFilter_ARM
func (filter *AdvancedFilter_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["operatorType"]
	if discriminator == "BoolEquals" {
		filter.BoolEquals = &BoolEqualsAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.BoolEquals)
	}
	if discriminator == "NumberGreaterThan" {
		filter.NumberGreaterThan = &NumberGreaterThanAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.NumberGreaterThan)
	}
	if discriminator == "NumberGreaterThanOrEquals" {
		filter.NumberGreaterThanOrEquals = &NumberGreaterThanOrEqualsAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.NumberGreaterThanOrEquals)
	}
	if discriminator == "NumberIn" {
		filter.NumberIn = &NumberInAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.NumberIn)
	}
	if discriminator == "NumberLessThan" {
		filter.NumberLessThan = &NumberLessThanAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.NumberLessThan)
	}
	if discriminator == "NumberLessThanOrEquals" {
		filter.NumberLessThanOrEquals = &NumberLessThanOrEqualsAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.NumberLessThanOrEquals)
	}
	if discriminator == "NumberNotIn" {
		filter.NumberNotIn = &NumberNotInAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.NumberNotIn)
	}
	if discriminator == "StringBeginsWith" {
		filter.StringBeginsWith = &StringBeginsWithAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.StringBeginsWith)
	}
	if discriminator == "StringContains" {
		filter.StringContains = &StringContainsAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.StringContains)
	}
	if discriminator == "StringEndsWith" {
		filter.StringEndsWith = &StringEndsWithAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.StringEndsWith)
	}
	if discriminator == "StringIn" {
		filter.StringIn = &StringInAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.StringIn)
	}
	if discriminator == "StringNotIn" {
		filter.StringNotIn = &StringNotInAdvancedFilter_ARM{}
		return json.Unmarshal(data, filter.StringNotIn)
	}

	// No error
	return nil
}

// Deprecated version of AzureFunctionEventSubscriptionDestination. Use v1beta20200601.AzureFunctionEventSubscriptionDestination instead
type AzureFunctionEventSubscriptionDestination_ARM struct {
	EndpointType AzureFunctionEventSubscriptionDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *AzureFunctionEventSubscriptionDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of EventHubEventSubscriptionDestination. Use v1beta20200601.EventHubEventSubscriptionDestination instead
type EventHubEventSubscriptionDestination_ARM struct {
	EndpointType EventHubEventSubscriptionDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *EventHubEventSubscriptionDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of HybridConnectionEventSubscriptionDestination. Use v1beta20200601.HybridConnectionEventSubscriptionDestination instead
type HybridConnectionEventSubscriptionDestination_ARM struct {
	EndpointType HybridConnectionEventSubscriptionDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *HybridConnectionEventSubscriptionDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of ServiceBusQueueEventSubscriptionDestination. Use v1beta20200601.ServiceBusQueueEventSubscriptionDestination instead
type ServiceBusQueueEventSubscriptionDestination_ARM struct {
	EndpointType ServiceBusQueueEventSubscriptionDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *ServiceBusQueueEventSubscriptionDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of ServiceBusTopicEventSubscriptionDestination. Use v1beta20200601.ServiceBusTopicEventSubscriptionDestination instead
type ServiceBusTopicEventSubscriptionDestination_ARM struct {
	EndpointType ServiceBusTopicEventSubscriptionDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *ServiceBusTopicEventSubscriptionDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of StorageBlobDeadLetterDestination. Use v1beta20200601.StorageBlobDeadLetterDestination instead
type StorageBlobDeadLetterDestination_ARM struct {
	EndpointType StorageBlobDeadLetterDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *StorageBlobDeadLetterDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of StorageQueueEventSubscriptionDestination. Use v1beta20200601.StorageQueueEventSubscriptionDestination instead
type StorageQueueEventSubscriptionDestination_ARM struct {
	EndpointType StorageQueueEventSubscriptionDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *StorageQueueEventSubscriptionDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of WebHookEventSubscriptionDestination. Use v1beta20200601.WebHookEventSubscriptionDestination instead
type WebHookEventSubscriptionDestination_ARM struct {
	EndpointType WebHookEventSubscriptionDestination_EndpointType   `json:"endpointType,omitempty"`
	Properties   *WebHookEventSubscriptionDestinationProperties_ARM `json:"properties,omitempty"`
}

// Deprecated version of AzureFunctionEventSubscriptionDestinationProperties. Use v1beta20200601.AzureFunctionEventSubscriptionDestinationProperties instead
type AzureFunctionEventSubscriptionDestinationProperties_ARM struct {
	MaxEventsPerBatch             *int    `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes *int    `json:"preferredBatchSizeInKilobytes,omitempty"`
	ResourceId                    *string `json:"resourceId,omitempty"`
}

// Deprecated version of BoolEqualsAdvancedFilter. Use v1beta20200601.BoolEqualsAdvancedFilter instead
type BoolEqualsAdvancedFilter_ARM struct {
	Key          *string                               `json:"key,omitempty"`
	OperatorType BoolEqualsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Value        *bool                                 `json:"value,omitempty"`
}

// Deprecated version of EventHubEventSubscriptionDestinationProperties. Use v1beta20200601.EventHubEventSubscriptionDestinationProperties instead
type EventHubEventSubscriptionDestinationProperties_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

// Deprecated version of HybridConnectionEventSubscriptionDestinationProperties. Use v1beta20200601.HybridConnectionEventSubscriptionDestinationProperties instead
type HybridConnectionEventSubscriptionDestinationProperties_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

// Deprecated version of NumberGreaterThanAdvancedFilter. Use v1beta20200601.NumberGreaterThanAdvancedFilter instead
type NumberGreaterThanAdvancedFilter_ARM struct {
	Key          *string                                      `json:"key,omitempty"`
	OperatorType NumberGreaterThanAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Value        *float64                                     `json:"value,omitempty"`
}

// Deprecated version of NumberGreaterThanOrEqualsAdvancedFilter. Use v1beta20200601.NumberGreaterThanOrEqualsAdvancedFilter instead
type NumberGreaterThanOrEqualsAdvancedFilter_ARM struct {
	Key          *string                                              `json:"key,omitempty"`
	OperatorType NumberGreaterThanOrEqualsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Value        *float64                                             `json:"value,omitempty"`
}

// Deprecated version of NumberInAdvancedFilter. Use v1beta20200601.NumberInAdvancedFilter instead
type NumberInAdvancedFilter_ARM struct {
	Key          *string                             `json:"key,omitempty"`
	OperatorType NumberInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Values       []float64                           `json:"values,omitempty"`
}

// Deprecated version of NumberLessThanAdvancedFilter. Use v1beta20200601.NumberLessThanAdvancedFilter instead
type NumberLessThanAdvancedFilter_ARM struct {
	Key          *string                                   `json:"key,omitempty"`
	OperatorType NumberLessThanAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Value        *float64                                  `json:"value,omitempty"`
}

// Deprecated version of NumberLessThanOrEqualsAdvancedFilter. Use v1beta20200601.NumberLessThanOrEqualsAdvancedFilter instead
type NumberLessThanOrEqualsAdvancedFilter_ARM struct {
	Key          *string                                           `json:"key,omitempty"`
	OperatorType NumberLessThanOrEqualsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Value        *float64                                          `json:"value,omitempty"`
}

// Deprecated version of NumberNotInAdvancedFilter. Use v1beta20200601.NumberNotInAdvancedFilter instead
type NumberNotInAdvancedFilter_ARM struct {
	Key          *string                                `json:"key,omitempty"`
	OperatorType NumberNotInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Values       []float64                              `json:"values,omitempty"`
}

// Deprecated version of ServiceBusQueueEventSubscriptionDestinationProperties. Use v1beta20200601.ServiceBusQueueEventSubscriptionDestinationProperties instead
type ServiceBusQueueEventSubscriptionDestinationProperties_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

// Deprecated version of ServiceBusTopicEventSubscriptionDestinationProperties. Use v1beta20200601.ServiceBusTopicEventSubscriptionDestinationProperties instead
type ServiceBusTopicEventSubscriptionDestinationProperties_ARM struct {
	ResourceId *string `json:"resourceId,omitempty"`
}

// Deprecated version of StorageBlobDeadLetterDestinationProperties. Use v1beta20200601.StorageBlobDeadLetterDestinationProperties instead
type StorageBlobDeadLetterDestinationProperties_ARM struct {
	BlobContainerName *string `json:"blobContainerName,omitempty"`
	ResourceId        *string `json:"resourceId,omitempty"`
}

// Deprecated version of StorageQueueEventSubscriptionDestinationProperties. Use v1beta20200601.StorageQueueEventSubscriptionDestinationProperties instead
type StorageQueueEventSubscriptionDestinationProperties_ARM struct {
	QueueName  *string `json:"queueName,omitempty"`
	ResourceId *string `json:"resourceId,omitempty"`
}

// Deprecated version of StringBeginsWithAdvancedFilter. Use v1beta20200601.StringBeginsWithAdvancedFilter instead
type StringBeginsWithAdvancedFilter_ARM struct {
	Key          *string                                     `json:"key,omitempty"`
	OperatorType StringBeginsWithAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Values       []string                                    `json:"values,omitempty"`
}

// Deprecated version of StringContainsAdvancedFilter. Use v1beta20200601.StringContainsAdvancedFilter instead
type StringContainsAdvancedFilter_ARM struct {
	Key          *string                                   `json:"key,omitempty"`
	OperatorType StringContainsAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Values       []string                                  `json:"values,omitempty"`
}

// Deprecated version of StringEndsWithAdvancedFilter. Use v1beta20200601.StringEndsWithAdvancedFilter instead
type StringEndsWithAdvancedFilter_ARM struct {
	Key          *string                                   `json:"key,omitempty"`
	OperatorType StringEndsWithAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Values       []string                                  `json:"values,omitempty"`
}

// Deprecated version of StringInAdvancedFilter. Use v1beta20200601.StringInAdvancedFilter instead
type StringInAdvancedFilter_ARM struct {
	Key          *string                             `json:"key,omitempty"`
	OperatorType StringInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Values       []string                            `json:"values,omitempty"`
}

// Deprecated version of StringNotInAdvancedFilter. Use v1beta20200601.StringNotInAdvancedFilter instead
type StringNotInAdvancedFilter_ARM struct {
	Key          *string                                `json:"key,omitempty"`
	OperatorType StringNotInAdvancedFilter_OperatorType `json:"operatorType,omitempty"`
	Values       []string                               `json:"values,omitempty"`
}

// Deprecated version of WebHookEventSubscriptionDestinationProperties. Use v1beta20200601.WebHookEventSubscriptionDestinationProperties instead
type WebHookEventSubscriptionDestinationProperties_ARM struct {
	AzureActiveDirectoryApplicationIdOrUri *string `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`
	AzureActiveDirectoryTenantId           *string `json:"azureActiveDirectoryTenantId,omitempty"`
	EndpointUrl                            *string `json:"endpointUrl,omitempty"`
	MaxEventsPerBatch                      *int    `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes          *int    `json:"preferredBatchSizeInKilobytes,omitempty"`
}
