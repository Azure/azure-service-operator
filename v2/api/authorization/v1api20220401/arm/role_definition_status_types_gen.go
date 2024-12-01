// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Role definition.
type RoleDefinition_STATUS struct {
	// Id: The role definition ID.
	Id *string `json:"id,omitempty"`

	// Name: The role definition name.
	Name *string `json:"name,omitempty"`

	// Properties: Role definition properties.
	Properties *RoleDefinitionProperties_STATUS `json:"properties,omitempty"`

	// Type: The role definition type.
	Type *string `json:"type,omitempty"`
}

// Role definition properties.
type RoleDefinitionProperties_STATUS struct {
	// AssignableScopes: Role definition assignable scopes.
	AssignableScopes []string `json:"assignableScopes,omitempty"`

	// CreatedBy: Id of the user who created the assignment
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedOn: Time it was created
	CreatedOn *string `json:"createdOn,omitempty"`

	// Description: The role definition description.
	Description *string `json:"description,omitempty"`

	// Permissions: Role definition permissions.
	Permissions []Permission_STATUS `json:"permissions,omitempty"`

	// RoleName: The role name.
	RoleName *string `json:"roleName,omitempty"`

	// Type: The role type.
	Type *string `json:"type,omitempty"`

	// UpdatedBy: Id of the user who updated the assignment
	UpdatedBy *string `json:"updatedBy,omitempty"`

	// UpdatedOn: Time it was updated
	UpdatedOn *string `json:"updatedOn,omitempty"`
}

// Role definition permissions.
type Permission_STATUS struct {
	// Actions: Allowed actions.
	Actions []string `json:"actions,omitempty"`

	// DataActions: Allowed Data actions.
	DataActions []string `json:"dataActions,omitempty"`

	// NotActions: Denied actions.
	NotActions []string `json:"notActions,omitempty"`

	// NotDataActions: Denied Data actions.
	NotDataActions []string `json:"notDataActions,omitempty"`
}
