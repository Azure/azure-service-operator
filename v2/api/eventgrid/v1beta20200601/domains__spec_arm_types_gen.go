// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Domains_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the Domain.
	Properties *DomainPropertiesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Domains_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (domains Domains_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (domains *Domains_SpecARM) GetName() string {
	return domains.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains"
func (domains *Domains_SpecARM) GetType() string {
	return "Microsoft.EventGrid/domains"
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/DomainProperties
type DomainPropertiesARM struct {
	// InboundIpRules: This can be used to restrict traffic from specific IPs instead of all IPs. Note: These are considered
	// only if PublicNetworkAccess is enabled.
	InboundIpRules []InboundIpRuleARM `json:"inboundIpRules,omitempty"`

	// InputSchema: This determines the format that Event Grid should expect for incoming events published to the domain.
	InputSchema *DomainPropertiesInputSchema `json:"inputSchema,omitempty"`

	// InputSchemaMapping: By default, Event Grid expects events to be in the Event Grid event schema. Specifying an input
	// schema mapping enables publishing to Event Grid using a custom input schema. Currently, the only supported type of
	// InputSchemaMapping is 'JsonInputSchemaMapping'.
	InputSchemaMapping *JsonInputSchemaMappingARM `json:"inputSchemaMapping,omitempty"`

	// PublicNetworkAccess: This determines if traffic is allowed over public network. By default it is enabled.
	// You can further restrict to specific IPs by configuring <seealso
	// cref="P:Microsoft.Azure.Events.ResourceProvider.Common.Contracts.DomainProperties.InboundIpRules" />.
	PublicNetworkAccess *DomainPropertiesPublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/InboundIpRule
type InboundIpRuleARM struct {
	// Action: Action to perform based on the match or no match of the IpMask.
	Action *InboundIpRuleAction `json:"action,omitempty"`

	// IpMask: IP Address in CIDR notation e.g., 10.0.0.0/8.
	IpMask *string `json:"ipMask,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonInputSchemaMapping
type JsonInputSchemaMappingARM struct {
	InputSchemaMappingType *JsonInputSchemaMappingInputSchemaMappingType `json:"inputSchemaMappingType,omitempty"`

	// Properties: This can be used to map properties of a source schema (or default values, for certain supported properties)
	// to properties of the EventGridEvent schema.
	Properties *JsonInputSchemaMappingPropertiesARM `json:"properties,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonInputSchemaMappingProperties
type JsonInputSchemaMappingPropertiesARM struct {
	// DataVersion: This is used to express the source of an input schema mapping for a single target field
	// in the Event Grid Event schema. This is currently used in the mappings for the 'subject',
	// 'eventtype' and 'dataversion' properties. This represents a field in the input event schema
	// along with a default value to be used, and at least one of these two properties should be provided.
	DataVersion *JsonFieldWithDefaultARM `json:"dataVersion,omitempty"`

	// EventTime: This is used to express the source of an input schema mapping for a single target field in the Event Grid
	// Event schema. This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a
	// field in the input event schema.
	EventTime *JsonFieldARM `json:"eventTime,omitempty"`

	// EventType: This is used to express the source of an input schema mapping for a single target field
	// in the Event Grid Event schema. This is currently used in the mappings for the 'subject',
	// 'eventtype' and 'dataversion' properties. This represents a field in the input event schema
	// along with a default value to be used, and at least one of these two properties should be provided.
	EventType *JsonFieldWithDefaultARM `json:"eventType,omitempty"`

	// Id: This is used to express the source of an input schema mapping for a single target field in the Event Grid Event
	// schema. This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a field
	// in the input event schema.
	Id *JsonFieldARM `json:"id,omitempty"`

	// Subject: This is used to express the source of an input schema mapping for a single target field
	// in the Event Grid Event schema. This is currently used in the mappings for the 'subject',
	// 'eventtype' and 'dataversion' properties. This represents a field in the input event schema
	// along with a default value to be used, and at least one of these two properties should be provided.
	Subject *JsonFieldWithDefaultARM `json:"subject,omitempty"`

	// Topic: This is used to express the source of an input schema mapping for a single target field in the Event Grid Event
	// schema. This is currently used in the mappings for the 'id', 'topic' and 'eventtime' properties. This represents a field
	// in the input event schema.
	Topic *JsonFieldARM `json:"topic,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonField
type JsonFieldARM struct {
	// SourceField: Name of a field in the input event schema that's to be used as the source of a mapping.
	SourceField *string `json:"sourceField,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/definitions/JsonFieldWithDefault
type JsonFieldWithDefaultARM struct {
	// DefaultValue: The default value to be used for mapping when a SourceField is not provided or if there's no property with
	// the specified name in the published JSON event payload.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// SourceField: Name of a field in the input event schema that's to be used as the source of a mapping.
	SourceField *string `json:"sourceField,omitempty"`
}
