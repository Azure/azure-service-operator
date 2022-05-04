// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

// Deprecated version of Topic_Status. Use v1beta20200601.Topic_Status instead
type Topic_StatusARM struct {
	Id         *string                    `json:"id,omitempty"`
	Location   *string                    `json:"location,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	Properties *TopicProperties_StatusARM `json:"properties,omitempty"`
	SystemData *SystemData_StatusARM      `json:"systemData,omitempty"`
	Tags       map[string]string          `json:"tags,omitempty"`
	Type       *string                    `json:"type,omitempty"`
}

// Deprecated version of TopicProperties_Status. Use v1beta20200601.TopicProperties_Status instead
type TopicProperties_StatusARM struct {
	Endpoint                   *string                                                         `json:"endpoint,omitempty"`
	InboundIpRules             []InboundIpRule_StatusARM                                       `json:"inboundIpRules,omitempty"`
	InputSchema                *TopicPropertiesStatusInputSchema                               `json:"inputSchema,omitempty"`
	InputSchemaMapping         *InputSchemaMapping_StatusARM                                   `json:"inputSchemaMapping,omitempty"`
	MetricResourceId           *string                                                         `json:"metricResourceId,omitempty"`
	PrivateEndpointConnections []PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *TopicPropertiesStatusProvisioningState                         `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *TopicPropertiesStatusPublicNetworkAccess                       `json:"publicNetworkAccess,omitempty"`
}

// Deprecated version of PrivateEndpointConnection_Status_Topic_SubResourceEmbedded. Use v1beta20200601.PrivateEndpointConnection_Status_Topic_SubResourceEmbedded instead
type PrivateEndpointConnection_Status_Topic_SubResourceEmbeddedARM struct {
	Id *string `json:"id,omitempty"`
}
