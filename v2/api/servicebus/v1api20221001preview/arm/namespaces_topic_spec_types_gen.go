// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesTopic_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of topic resource.
	Properties *SBTopicProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesTopic_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-10-01-preview"
func (topic NamespacesTopic_Spec) GetAPIVersion() string {
	return "2022-10-01-preview"
}

// GetName returns the Name of the resource
func (topic *NamespacesTopic_Spec) GetName() string {
	return topic.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics"
func (topic *NamespacesTopic_Spec) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics"
}

// The Topic Properties definition.
type SBTopicProperties struct {
	// AutoDeleteOnIdle: ISO 8601 timespan idle interval after which the topic is automatically deleted. The minimum duration
	// is 5 minutes.
	AutoDeleteOnIdle *string `json:"autoDeleteOnIdle,omitempty"`

	// DefaultMessageTimeToLive: ISO 8601 Default message timespan to live value. This is the duration after which the message
	// expires, starting from when the message is sent to Service Bus. This is the default value used when TimeToLive is not
	// set on a message itself.
	DefaultMessageTimeToLive *string `json:"defaultMessageTimeToLive,omitempty"`

	// DuplicateDetectionHistoryTimeWindow: ISO8601 timespan structure that defines the duration of the duplicate detection
	// history. The default value is 10 minutes.
	DuplicateDetectionHistoryTimeWindow *string `json:"duplicateDetectionHistoryTimeWindow,omitempty"`

	// EnableBatchedOperations: Value that indicates whether server-side batched operations are enabled.
	EnableBatchedOperations *bool `json:"enableBatchedOperations,omitempty"`

	// EnableExpress: Value that indicates whether Express Entities are enabled. An express topic holds a message in memory
	// temporarily before writing it to persistent storage.
	EnableExpress *bool `json:"enableExpress,omitempty"`

	// EnablePartitioning: Value that indicates whether the topic to be partitioned across multiple message brokers is enabled.
	EnablePartitioning *bool `json:"enablePartitioning,omitempty"`

	// MaxMessageSizeInKilobytes: Maximum size (in KB) of the message payload that can be accepted by the topic. This property
	// is only used in Premium today and default is 1024.
	MaxMessageSizeInKilobytes *int `json:"maxMessageSizeInKilobytes,omitempty"`

	// MaxSizeInMegabytes: Maximum size of the topic in megabytes, which is the size of the memory allocated for the topic.
	// Default is 1024.
	MaxSizeInMegabytes *int `json:"maxSizeInMegabytes,omitempty"`

	// RequiresDuplicateDetection: Value indicating if this topic requires duplicate detection.
	RequiresDuplicateDetection *bool `json:"requiresDuplicateDetection,omitempty"`

	// SupportOrdering: Value that indicates whether the topic supports ordering.
	SupportOrdering *bool `json:"supportOrdering,omitempty"`
}
