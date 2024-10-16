// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesTopicsSubscriptionsRule_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of Rule resource
	Properties *Ruleproperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesTopicsSubscriptionsRule_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesTopicsSubscriptionsRule_Spec_ARM) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (rule *NamespacesTopicsSubscriptionsRule_Spec_ARM) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/topics/subscriptions/rules"
func (rule *NamespacesTopicsSubscriptionsRule_Spec_ARM) GetType() string {
	return "Microsoft.ServiceBus/namespaces/topics/subscriptions/rules"
}

// Description of Rule Resource.
type Ruleproperties_ARM struct {
	// Action: Represents the filter actions which are allowed for the transformation of a message that have been matched by a
	// filter expression.
	Action *Action_ARM `json:"action,omitempty"`

	// CorrelationFilter: Properties of correlationFilter
	CorrelationFilter *CorrelationFilter_ARM `json:"correlationFilter,omitempty"`

	// FilterType: Filter type that is evaluated against a BrokeredMessage.
	FilterType *FilterType_ARM `json:"filterType,omitempty"`

	// SqlFilter: Properties of sqlFilter
	SqlFilter *SqlFilter_ARM `json:"sqlFilter,omitempty"`
}

// Represents the filter actions which are allowed for the transformation of a message that have been matched by a filter
// expression.
type Action_ARM struct {
	// CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
	// currently hard-coded to 20.
	CompatibilityLevel *int `json:"compatibilityLevel,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SqlExpression: SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `json:"sqlExpression,omitempty"`
}

// Represents the correlation filter expression.
type CorrelationFilter_ARM struct {
	// ContentType: Content type of the message.
	ContentType *string `json:"contentType,omitempty"`

	// CorrelationId: Identifier of the correlation.
	CorrelationId *string `json:"correlationId,omitempty"`

	// Label: Application specific label.
	Label *string `json:"label,omitempty"`

	// MessageId: Identifier of the message.
	MessageId *string `json:"messageId,omitempty"`

	// Properties: dictionary object for custom filters
	Properties map[string]string `json:"properties,omitempty"`

	// ReplyTo: Address of the queue to reply to.
	ReplyTo *string `json:"replyTo,omitempty"`

	// ReplyToSessionId: Session identifier to reply to.
	ReplyToSessionId *string `json:"replyToSessionId,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SessionId: Session identifier.
	SessionId *string `json:"sessionId,omitempty"`

	// To: Address to send to.
	To *string `json:"to,omitempty"`
}

// Rule filter types
// +kubebuilder:validation:Enum={"CorrelationFilter","SqlFilter"}
type FilterType_ARM string

const (
	FilterType_ARM_CorrelationFilter = FilterType_ARM("CorrelationFilter")
	FilterType_ARM_SqlFilter         = FilterType_ARM("SqlFilter")
)

// Mapping from string to FilterType_ARM
var filterType_ARM_Values = map[string]FilterType_ARM{
	"correlationfilter": FilterType_ARM_CorrelationFilter,
	"sqlfilter":         FilterType_ARM_SqlFilter,
}

// Represents a filter which is a composition of an expression and an action that is executed in the pub/sub pipeline.
type SqlFilter_ARM struct {
	// CompatibilityLevel: This property is reserved for future use. An integer value showing the compatibility level,
	// currently hard-coded to 20.
	CompatibilityLevel *int `json:"compatibilityLevel,omitempty"`

	// RequiresPreprocessing: Value that indicates whether the rule action requires preprocessing.
	RequiresPreprocessing *bool `json:"requiresPreprocessing,omitempty"`

	// SqlExpression: The SQL expression. e.g. MyProperty='ABC'
	SqlExpression *string `json:"sqlExpression,omitempty"`
}
