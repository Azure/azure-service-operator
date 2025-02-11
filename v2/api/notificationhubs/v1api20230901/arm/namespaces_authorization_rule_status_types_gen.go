// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type NamespacesAuthorizationRule_STATUS struct {
	// Id: Fully qualified resource ID for the resource. E.g.
	// "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"
	Id *string `json:"id,omitempty"`

	// Location: Deprecated - only for compatibility.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: SharedAccessAuthorizationRule properties.
	Properties *SharedAccessAuthorizationRuleProperties_STATUS `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Tags: Deprecated - only for compatibility.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// SharedAccessAuthorizationRule properties.
type SharedAccessAuthorizationRuleProperties_STATUS struct {
	// ClaimType: Gets a string that describes the claim type
	ClaimType *string `json:"claimType,omitempty"`

	// ClaimValue: Gets a string that describes the claim value
	ClaimValue *string `json:"claimValue,omitempty"`

	// CreatedTime: Gets the created time for this rule
	CreatedTime *string `json:"createdTime,omitempty"`

	// KeyName: Gets a string that describes the authorization rule.
	KeyName *string `json:"keyName,omitempty"`

	// ModifiedTime: Gets the last modified time for this rule
	ModifiedTime *string `json:"modifiedTime,omitempty"`

	// Revision: Gets the revision number for the rule
	Revision *int `json:"revision,omitempty"`

	// Rights: Gets or sets the rights associated with the rule.
	Rights []AccessRights_STATUS `json:"rights,omitempty"`
}

// Defines values for AccessRights.
type AccessRights_STATUS string

const (
	AccessRights_STATUS_Listen = AccessRights_STATUS("Listen")
	AccessRights_STATUS_Manage = AccessRights_STATUS("Manage")
	AccessRights_STATUS_Send   = AccessRights_STATUS("Send")
)

// Mapping from string to AccessRights_STATUS
var accessRights_STATUS_Values = map[string]AccessRights_STATUS{
	"listen": AccessRights_STATUS_Listen,
	"manage": AccessRights_STATUS_Manage,
	"send":   AccessRights_STATUS_Send,
}
