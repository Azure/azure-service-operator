// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type NamespacesTopicsSubscription_STATUS struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of subscriptions resource.
	Properties *SBSubscriptionProperties_STATUS `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

// Description of Subscription Resource.
type SBSubscriptionProperties_STATUS struct {
	// AccessedAt: Last time there was a receive request to this subscription.
	AccessedAt *string `json:"accessedAt,omitempty"`

	// AutoDeleteOnIdle: ISO 8061 timeSpan idle interval after which the topic is automatically deleted. The minimum duration
	// is 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// ClientAffineProperties: Properties specific to client affine subscriptions.
	ClientAffineProperties *SBClientAffineProperties_STATUS `json:"clientAffineProperties,omitempty"`

	// CountDetails: Message count details
	CountDetails *MessageCountDetails_STATUS `json:"countDetails,omitempty"`

	// CreatedAt: Exact time the message was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	// DeadLetteringOnFilterEvaluationExceptions: Value that indicates whether a subscription has dead letter support on filter
	// evaluation exceptions.
	DeadLetteringOnFilterEvaluationExceptions *bool `json:"deadLetteringOnFilterEvaluationExceptions,omitempty"`

	// DeadLetteringOnMessageExpiration: Value that indicates whether a subscription has dead letter support when a message
	// expires.
	DeadLetteringOnMessageExpiration *bool `json:"deadLetteringOnMessageExpiration,omitempty"`

	// DefaultMessageTimeToLive: ISO 8061 Default message timespan to live value. This is the duration after which the message
	// expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
	// set on a message itself.
	DefaultMessageTimeToLive *string `json:"defaultMessageTimeToLive,omitempty"`

	// DuplicateDetectionHistoryTimeWindow: ISO 8601 timeSpan structure that defines the duration of the duplicate detection
	// history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`

	// EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty"`

	// ForwardDeadLetteredMessagesTo: Queue/Topic name to forward the Dead Letter message
	ForwardDeadLetteredMessagesTo *string `json:"forwardDeadLetteredMessagesTo,omitempty"`

	// ForwardTo: Queue/Topic name to forward the messages
	ForwardTo *string `json:"forwardTo,omitempty"`

	// IsClientAffine: Value that indicates whether the subscription has an affinity to the client id.
	IsClientAffine *bool `json:"isClientAffine,omitempty"`

	// LockDuration: ISO 8061 lock duration timespan for the subscription. The default value is 1 minute.
	LockDuration *string `json:"lockDuration,omitempty"`

	// MaxDeliveryCount: Number of maximum deliveries.
	MaxDeliveryCount *int `json:"maxDeliveryCount,omitempty"`

	// MessageCount: Number of messages.
	MessageCount *int `json:"messageCount,omitempty"`

	// RequiresSession: Value indicating if a subscription supports the concept of sessions.
	RequiresSession *bool `json:"requiresSession,omitempty"`

	// Status: Enumerates the possible values for the status of a messaging entity.
	Status *EntityStatus_STATUS `json:"status,omitempty"`

	// UpdatedAt: The exact time the message was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

// Properties specific to client affine subscriptions.
type SBClientAffineProperties_STATUS struct {
	// ClientId: Indicates the Client ID of the application that created the client-affine subscription.
	ClientId *string `json:"clientId,omitempty"`

	// IsDurable: For client-affine subscriptions, this value indicates whether the subscription is durable or not.
	IsDurable *bool `json:"isDurable,omitempty"`

	// IsShared: For client-affine subscriptions, this value indicates whether the subscription is shared or not.
	IsShared *bool `json:"isShared,omitempty"`
}
