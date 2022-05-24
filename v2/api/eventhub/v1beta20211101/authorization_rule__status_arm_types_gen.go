// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type AuthorizationRule_StatusARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties supplied to create or update AuthorizationRule
	Properties *AuthorizationRule_Status_PropertiesARM `json:"properties,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_StatusARM `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

type AuthorizationRule_Status_PropertiesARM struct {
	// Rights: The rights associated with the rule.
	Rights []string `json:"rights,omitempty"`
}

type SystemData_StatusARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *string `json:"createdByType,omitempty"`

	// LastModifiedAt: The type of identity that last modified the resource.
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *string `json:"lastModifiedByType,omitempty"`
}
