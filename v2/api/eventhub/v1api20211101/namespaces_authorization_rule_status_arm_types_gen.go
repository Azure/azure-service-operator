// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

type NamespacesAuthorizationRule_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties supplied to create or update AuthorizationRule
	Properties *Namespaces_AuthorizationRule_Properties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type Namespaces_AuthorizationRule_Properties_STATUS_ARM struct {
	// Rights: The rights associated with the rule.
	Rights []Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM `json:"rights,omitempty"`
}

type Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM string

const (
	Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM_Listen = Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM("Listen")
	Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM_Manage = Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM("Manage")
	Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM_Send   = Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM("Send")
)

// Mapping from string to Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM
var namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM_Values = map[string]Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM{
	"listen": Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM_Listen,
	"manage": Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM_Manage,
	"send":   Namespaces_AuthorizationRule_Properties_Rights_STATUS_ARM_Send,
}
