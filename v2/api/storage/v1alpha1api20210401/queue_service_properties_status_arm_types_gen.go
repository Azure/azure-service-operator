// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

// Deprecated version of QueueServiceProperties_STATUS. Use v1beta20210401.QueueServiceProperties_STATUS instead
type QueueServiceProperties_STATUS_ARM struct {
	Id         *string                                       `json:"id,omitempty"`
	Name       *string                                       `json:"name,omitempty"`
	Properties *QueueServiceProperties_Properties_STATUS_ARM `json:"properties,omitempty"`
	Type       *string                                       `json:"type,omitempty"`
}

// Deprecated version of QueueServiceProperties_Properties_STATUS. Use v1beta20210401.QueueServiceProperties_Properties_STATUS instead
type QueueServiceProperties_Properties_STATUS_ARM struct {
	Cors *CorsRules_STATUS_ARM `json:"cors,omitempty"`
}
