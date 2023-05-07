// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

type DnsZones_A_STATUS_ARM struct {
	// Etag: The etag of the record set.
	Etag *string `json:"etag,omitempty"`

	// Id: The ID of the record set.
	Id *string `json:"id,omitempty"`

	// Name: The name of the record set.
	Name *string `json:"name,omitempty"`

	// Properties: The properties of the record set.
	Properties *RecordSetProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the record set.
	Type *string `json:"type,omitempty"`
}
