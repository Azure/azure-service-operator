// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of Topic_Spec. Use v1beta20200601.Topic_Spec instead
type Topic_Spec_ARM struct {
	Location   *string              `json:"location,omitempty"`
	Name       string               `json:"name,omitempty"`
	Properties *TopicProperties_ARM `json:"properties,omitempty"`
	Tags       map[string]string    `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Topic_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic Topic_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (topic *Topic_Spec_ARM) GetName() string {
	return topic.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/topics"
func (topic *Topic_Spec_ARM) GetType() string {
	return "Microsoft.EventGrid/topics"
}

// Deprecated version of TopicProperties. Use v1beta20200601.TopicProperties instead
type TopicProperties_ARM struct {
	InboundIpRules      []InboundIpRule_ARM                  `json:"inboundIpRules,omitempty"`
	InputSchema         *TopicProperties_InputSchema         `json:"inputSchema,omitempty"`
	InputSchemaMapping  *InputSchemaMapping_ARM              `json:"inputSchemaMapping,omitempty"`
	PublicNetworkAccess *TopicProperties_PublicNetworkAccess `json:"publicNetworkAccess,omitempty"`
}
