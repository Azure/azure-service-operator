// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventgrid.azure.com,resources=eventsubscriptions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventgrid.azure.com,resources={eventsubscriptions/status,eventsubscriptions/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1beta20200601.EventSubscription
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/unknown_resourceDefinitions/eventSubscriptions
type EventSubscription struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventSubscriptions_Spec  `json:"spec,omitempty"`
	Status            EventSubscription_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &EventSubscription{}

// GetConditions returns the conditions of the resource
func (subscription *EventSubscription) GetConditions() conditions.Conditions {
	return subscription.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (subscription *EventSubscription) SetConditions(conditions conditions.Conditions) {
	subscription.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &EventSubscription{}

// AzureName returns the Azure name of the resource
func (subscription *EventSubscription) AzureName() string {
	return subscription.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (subscription EventSubscription) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceKind returns the kind of the resource
func (subscription *EventSubscription) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindExtension
}

// GetSpec returns the specification of this resource
func (subscription *EventSubscription) GetSpec() genruntime.ConvertibleSpec {
	return &subscription.Spec
}

// GetStatus returns the status of this resource
func (subscription *EventSubscription) GetStatus() genruntime.ConvertibleStatus {
	return &subscription.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/eventSubscriptions"
func (subscription *EventSubscription) GetType() string {
	return "Microsoft.EventGrid/eventSubscriptions"
}

// NewEmptyStatus returns a new empty (blank) status
func (subscription *EventSubscription) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &EventSubscription_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (subscription *EventSubscription) Owner() *genruntime.ResourceReference {
	return &genruntime.ResourceReference{
		Group: subscription.Spec.Owner.Group,
		Kind:  subscription.Spec.Owner.Kind,
		Name:  subscription.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (subscription *EventSubscription) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*EventSubscription_Status); ok {
		subscription.Status = *st
		return nil
	}

	// Convert status to required version
	var st EventSubscription_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	subscription.Status = st
	return nil
}

// Hub marks that this EventSubscription is the hub type for conversion
func (subscription *EventSubscription) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (subscription *EventSubscription) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: subscription.Spec.OriginalVersion,
		Kind:    "EventSubscription",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1beta20200601.EventSubscription
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/unknown_resourceDefinitions/eventSubscriptions
type EventSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EventSubscription `json:"items"`
}

//Storage version of v1beta20200601.EventSubscription_Status
type EventSubscription_Status struct {
	Conditions            []conditions.Condition               `json:"conditions,omitempty"`
	DeadLetterDestination *DeadLetterDestination_Status        `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination_Status `json:"destination,omitempty"`
	EventDeliverySchema   *string                              `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                              `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter_Status      `json:"filter,omitempty"`
	Id                    *string                              `json:"id,omitempty"`
	Labels                []string                             `json:"labels,omitempty"`
	Name                  *string                              `json:"name,omitempty"`
	PropertyBag           genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	ProvisioningState     *string                              `json:"provisioningState,omitempty"`
	RetryPolicy           *RetryPolicy_Status                  `json:"retryPolicy,omitempty"`
	SystemData            *SystemData_Status                   `json:"systemData,omitempty"`
	Topic                 *string                              `json:"topic,omitempty"`
	Type                  *string                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &EventSubscription_Status{}

// ConvertStatusFrom populates our EventSubscription_Status from the provided source
func (subscription *EventSubscription_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(subscription)
}

// ConvertStatusTo populates the provided destination from our EventSubscription_Status
func (subscription *EventSubscription_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == subscription {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(subscription)
}

//Storage version of v1beta20200601.EventSubscriptions_Spec
type EventSubscriptions_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName             string                            `json:"azureName,omitempty"`
	DeadLetterDestination *StorageBlobDeadLetterDestination `json:"deadLetterDestination,omitempty"`
	Destination           *EventSubscriptionDestination     `json:"destination,omitempty"`
	EventDeliverySchema   *string                           `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc     *string                           `json:"expirationTimeUtc,omitempty"`
	Filter                *EventSubscriptionFilter          `json:"filter,omitempty"`
	Labels                []string                          `json:"labels,omitempty"`
	Location              *string                           `json:"location,omitempty"`
	OriginalVersion       string                            `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	//extension resource, which means that any other Azure resource can be its owner.
	Owner       *genruntime.ArbitraryOwnerReference `json:"owner,omitempty"`
	PropertyBag genruntime.PropertyBag              `json:"$propertyBag,omitempty"`
	RetryPolicy *RetryPolicy                        `json:"retryPolicy,omitempty"`
	Tags        map[string]string                   `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &EventSubscriptions_Spec{}

// ConvertSpecFrom populates our EventSubscriptions_Spec from the provided source
func (subscriptions *EventSubscriptions_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == subscriptions {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(subscriptions)
}

// ConvertSpecTo populates the provided destination from our EventSubscriptions_Spec
func (subscriptions *EventSubscriptions_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == subscriptions {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(subscriptions)
}

//Storage version of v1beta20200601.DeadLetterDestination_Status
type DeadLetterDestination_Status struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.EventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventSubscriptionDestination
type EventSubscriptionDestination struct {
	AzureFunction    *AzureFunctionEventSubscriptionDestination    `json:"azureFunctionEventSubscriptionDestination,omitempty"`
	EventHub         *EventHubEventSubscriptionDestination         `json:"eventHubEventSubscriptionDestination,omitempty"`
	HybridConnection *HybridConnectionEventSubscriptionDestination `json:"hybridConnectionEventSubscriptionDestination,omitempty"`
	PropertyBag      genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	ServiceBusQueue  *ServiceBusQueueEventSubscriptionDestination  `json:"serviceBusQueueEventSubscriptionDestination,omitempty"`
	ServiceBusTopic  *ServiceBusTopicEventSubscriptionDestination  `json:"serviceBusTopicEventSubscriptionDestination,omitempty"`
	StorageQueue     *StorageQueueEventSubscriptionDestination     `json:"storageQueueEventSubscriptionDestination,omitempty"`
	WebHook          *WebHookEventSubscriptionDestination          `json:"webHookEventSubscriptionDestination,omitempty"`
}

//Storage version of v1beta20200601.EventSubscriptionDestination_Status
type EventSubscriptionDestination_Status struct {
	EndpointType *string                `json:"endpointType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.EventSubscriptionFilter
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventSubscriptionFilter
type EventSubscriptionFilter struct {
	AdvancedFilters        []AdvancedFilter       `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string               `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                  `json:"isSubjectCaseSensitive,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SubjectBeginsWith      *string                `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string                `json:"subjectEndsWith,omitempty"`
}

//Storage version of v1beta20200601.EventSubscriptionFilter_Status
type EventSubscriptionFilter_Status struct {
	AdvancedFilters        []AdvancedFilter_Status `json:"advancedFilters,omitempty"`
	IncludedEventTypes     []string                `json:"includedEventTypes,omitempty"`
	IsSubjectCaseSensitive *bool                   `json:"isSubjectCaseSensitive,omitempty"`
	PropertyBag            genruntime.PropertyBag  `json:"$propertyBag,omitempty"`
	SubjectBeginsWith      *string                 `json:"subjectBeginsWith,omitempty"`
	SubjectEndsWith        *string                 `json:"subjectEndsWith,omitempty"`
}

//Storage version of v1beta20200601.RetryPolicy
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/RetryPolicy
type RetryPolicy struct {
	EventTimeToLiveInMinutes *int                   `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int                   `json:"maxDeliveryAttempts,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.RetryPolicy_Status
type RetryPolicy_Status struct {
	EventTimeToLiveInMinutes *int                   `json:"eventTimeToLiveInMinutes,omitempty"`
	MaxDeliveryAttempts      *int                   `json:"maxDeliveryAttempts,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.StorageBlobDeadLetterDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageBlobDeadLetterDestination
type StorageBlobDeadLetterDestination struct {
	EndpointType *string                                     `json:"endpointType,omitempty"`
	Properties   *StorageBlobDeadLetterDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                      `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/AdvancedFilter
type AdvancedFilter struct {
	BoolEquals                *AdvancedFilter_BoolEquals                `json:"boolEqualsAdvancedFilter,omitempty"`
	NumberGreaterThan         *AdvancedFilter_NumberGreaterThan         `json:"numberGreaterThanAdvancedFilter,omitempty"`
	NumberGreaterThanOrEquals *AdvancedFilter_NumberGreaterThanOrEquals `json:"numberGreaterThanOrEqualsAdvancedFilter,omitempty"`
	NumberIn                  *AdvancedFilter_NumberIn                  `json:"numberInAdvancedFilter,omitempty"`
	NumberLessThan            *AdvancedFilter_NumberLessThan            `json:"numberLessThanAdvancedFilter,omitempty"`
	NumberLessThanOrEquals    *AdvancedFilter_NumberLessThanOrEquals    `json:"numberLessThanOrEqualsAdvancedFilter,omitempty"`
	NumberNotIn               *AdvancedFilter_NumberNotIn               `json:"numberNotInAdvancedFilter,omitempty"`
	PropertyBag               genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	StringBeginsWith          *AdvancedFilter_StringBeginsWith          `json:"stringBeginsWithAdvancedFilter,omitempty"`
	StringContains            *AdvancedFilter_StringContains            `json:"stringContainsAdvancedFilter,omitempty"`
	StringEndsWith            *AdvancedFilter_StringEndsWith            `json:"stringEndsWithAdvancedFilter,omitempty"`
	StringIn                  *AdvancedFilter_StringIn                  `json:"stringInAdvancedFilter,omitempty"`
	StringNotIn               *AdvancedFilter_StringNotIn               `json:"stringNotInAdvancedFilter,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_Status
type AdvancedFilter_Status struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.AzureFunctionEventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/AzureFunctionEventSubscriptionDestination
type AzureFunctionEventSubscriptionDestination struct {
	EndpointType *string                                              `json:"endpointType,omitempty"`
	Properties   *AzureFunctionEventSubscriptionDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                               `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.EventHubEventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventHubEventSubscriptionDestination
type EventHubEventSubscriptionDestination struct {
	EndpointType *string                                         `json:"endpointType,omitempty"`
	Properties   *EventHubEventSubscriptionDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                          `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.HybridConnectionEventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/HybridConnectionEventSubscriptionDestination
type HybridConnectionEventSubscriptionDestination struct {
	EndpointType *string                                                 `json:"endpointType,omitempty"`
	Properties   *HybridConnectionEventSubscriptionDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                                  `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.ServiceBusQueueEventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusQueueEventSubscriptionDestination
type ServiceBusQueueEventSubscriptionDestination struct {
	EndpointType *string                                                `json:"endpointType,omitempty"`
	Properties   *ServiceBusQueueEventSubscriptionDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.ServiceBusTopicEventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusTopicEventSubscriptionDestination
type ServiceBusTopicEventSubscriptionDestination struct {
	EndpointType *string                                                `json:"endpointType,omitempty"`
	Properties   *ServiceBusTopicEventSubscriptionDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                                 `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.StorageBlobDeadLetterDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageBlobDeadLetterDestinationProperties
type StorageBlobDeadLetterDestinationProperties struct {
	BlobContainerName *string                `json:"blobContainerName,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//ResourceReference: The Azure Resource ID of the storage account that is the destination of the deadletter events
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

//Storage version of v1beta20200601.StorageQueueEventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageQueueEventSubscriptionDestination
type StorageQueueEventSubscriptionDestination struct {
	EndpointType *string                                             `json:"endpointType,omitempty"`
	Properties   *StorageQueueEventSubscriptionDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                              `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.WebHookEventSubscriptionDestination
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/WebHookEventSubscriptionDestination
type WebHookEventSubscriptionDestination struct {
	EndpointType *string                                        `json:"endpointType,omitempty"`
	Properties   *WebHookEventSubscriptionDestinationProperties `json:"properties,omitempty"`
	PropertyBag  genruntime.PropertyBag                         `json:"$propertyBag,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_BoolEquals
type AdvancedFilter_BoolEquals struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *bool                  `json:"value,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_NumberGreaterThan
type AdvancedFilter_NumberGreaterThan struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_NumberGreaterThanOrEquals
type AdvancedFilter_NumberGreaterThanOrEquals struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_NumberIn
type AdvancedFilter_NumberIn struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []float64              `json:"values,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_NumberLessThan
type AdvancedFilter_NumberLessThan struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_NumberLessThanOrEquals
type AdvancedFilter_NumberLessThanOrEquals struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Value        *float64               `json:"value,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_NumberNotIn
type AdvancedFilter_NumberNotIn struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []float64              `json:"values,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_StringBeginsWith
type AdvancedFilter_StringBeginsWith struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_StringContains
type AdvancedFilter_StringContains struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_StringEndsWith
type AdvancedFilter_StringEndsWith struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_StringIn
type AdvancedFilter_StringIn struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

//Storage version of v1beta20200601.AdvancedFilter_StringNotIn
type AdvancedFilter_StringNotIn struct {
	Key          *string                `json:"key,omitempty"`
	OperatorType *string                `json:"operatorType,omitempty"`
	PropertyBag  genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Values       []string               `json:"values,omitempty"`
}

//Storage version of v1beta20200601.AzureFunctionEventSubscriptionDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/AzureFunctionEventSubscriptionDestinationProperties
type AzureFunctionEventSubscriptionDestinationProperties struct {
	MaxEventsPerBatch             *int                   `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes *int                   `json:"preferredBatchSizeInKilobytes,omitempty"`
	PropertyBag                   genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//ResourceReference: The Azure Resource Id that represents the endpoint of the Azure Function destination of an event
	//subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

//Storage version of v1beta20200601.EventHubEventSubscriptionDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/EventHubEventSubscriptionDestinationProperties
type EventHubEventSubscriptionDestinationProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//ResourceReference: The Azure Resource Id that represents the endpoint of an Event Hub destination of an event
	//subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

//Storage version of v1beta20200601.HybridConnectionEventSubscriptionDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/HybridConnectionEventSubscriptionDestinationProperties
type HybridConnectionEventSubscriptionDestinationProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//ResourceReference: The Azure Resource ID of an hybrid connection that is the destination of an event subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

//Storage version of v1beta20200601.ServiceBusQueueEventSubscriptionDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusQueueEventSubscriptionDestinationProperties
type ServiceBusQueueEventSubscriptionDestinationProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//ResourceReference: The Azure Resource Id that represents the endpoint of the Service Bus destination of an event
	//subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

//Storage version of v1beta20200601.ServiceBusTopicEventSubscriptionDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/ServiceBusTopicEventSubscriptionDestinationProperties
type ServiceBusTopicEventSubscriptionDestinationProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	//ResourceReference: The Azure Resource Id that represents the endpoint of the Service Bus Topic destination of an event
	//subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

//Storage version of v1beta20200601.StorageQueueEventSubscriptionDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/StorageQueueEventSubscriptionDestinationProperties
type StorageQueueEventSubscriptionDestinationProperties struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	QueueName   *string                `json:"queueName,omitempty"`

	//ResourceReference: The Azure Resource ID of the storage account that contains the queue that is the destination of an
	//event subscription.
	ResourceReference *genruntime.ResourceReference `armReference:"ResourceId" json:"resourceReference,omitempty"`
}

//Storage version of v1beta20200601.WebHookEventSubscriptionDestinationProperties
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/WebHookEventSubscriptionDestinationProperties
type WebHookEventSubscriptionDestinationProperties struct {
	AzureActiveDirectoryApplicationIdOrUri *string                `json:"azureActiveDirectoryApplicationIdOrUri,omitempty"`
	AzureActiveDirectoryTenantId           *string                `json:"azureActiveDirectoryTenantId,omitempty"`
	EndpointUrl                            *string                `json:"endpointUrl,omitempty"`
	MaxEventsPerBatch                      *int                   `json:"maxEventsPerBatch,omitempty"`
	PreferredBatchSizeInKilobytes          *int                   `json:"preferredBatchSizeInKilobytes,omitempty"`
	PropertyBag                            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&EventSubscription{}, &EventSubscriptionList{})
}
