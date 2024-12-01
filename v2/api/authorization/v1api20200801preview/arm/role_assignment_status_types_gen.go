// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

// Role Assignments
type RoleAssignment_STATUS struct {
	// Id: The role assignment ID.
	Id *string `json:"id,omitempty"`

	// Name: The role assignment name.
	Name *string `json:"name,omitempty"`

	// Properties: Role assignment properties.
	Properties *RoleAssignmentProperties_STATUS `json:"properties,omitempty"`

	// Type: The role assignment type.
	Type *string `json:"type,omitempty"`
}

// Role assignment properties.
type RoleAssignmentProperties_STATUS struct {
	// Condition: The conditions on the role assignment. This limits the resources it can be assigned to. e.g.:
	// @Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName] StringEqualsIgnoreCase
	// 'foo_storage_container'
	Condition *string `json:"condition,omitempty"`

	// ConditionVersion: Version of the condition. Currently accepted value is '2.0'
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	// CreatedBy: Id of the user who created the assignment
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedOn: Time it was created
	CreatedOn *string `json:"createdOn,omitempty"`

	// DelegatedManagedIdentityResourceId: Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`

	// Description: Description of role assignment
	Description *string `json:"description,omitempty"`

	// PrincipalId: The principal ID.
	PrincipalId *string `json:"principalId,omitempty"`

	// PrincipalType: The principal type of the assigned principal ID.
	PrincipalType *RoleAssignmentProperties_PrincipalType_STATUS `json:"principalType,omitempty"`

	// RoleDefinitionId: The role definition ID.
	RoleDefinitionId *string `json:"roleDefinitionId,omitempty"`

	// Scope: The role assignment scope.
	Scope *string `json:"scope,omitempty"`

	// UpdatedBy: Id of the user who updated the assignment
	UpdatedBy *string `json:"updatedBy,omitempty"`

	// UpdatedOn: Time it was updated
	UpdatedOn *string `json:"updatedOn,omitempty"`
}

type RoleAssignmentProperties_PrincipalType_STATUS string

const (
	RoleAssignmentProperties_PrincipalType_STATUS_ForeignGroup     = RoleAssignmentProperties_PrincipalType_STATUS("ForeignGroup")
	RoleAssignmentProperties_PrincipalType_STATUS_Group            = RoleAssignmentProperties_PrincipalType_STATUS("Group")
	RoleAssignmentProperties_PrincipalType_STATUS_ServicePrincipal = RoleAssignmentProperties_PrincipalType_STATUS("ServicePrincipal")
	RoleAssignmentProperties_PrincipalType_STATUS_User             = RoleAssignmentProperties_PrincipalType_STATUS("User")
)

// Mapping from string to RoleAssignmentProperties_PrincipalType_STATUS
var roleAssignmentProperties_PrincipalType_STATUS_Values = map[string]RoleAssignmentProperties_PrincipalType_STATUS{
	"foreigngroup":     RoleAssignmentProperties_PrincipalType_STATUS_ForeignGroup,
	"group":            RoleAssignmentProperties_PrincipalType_STATUS_Group,
	"serviceprincipal": RoleAssignmentProperties_PrincipalType_STATUS_ServicePrincipal,
	"user":             RoleAssignmentProperties_PrincipalType_STATUS_User,
}
