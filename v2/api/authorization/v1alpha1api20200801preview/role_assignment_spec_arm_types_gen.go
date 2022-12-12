// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200801preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RoleAssignment_Spec. Use v1beta20200801preview.RoleAssignment_Spec instead
type RoleAssignment_Spec_ARM struct {
	Name       string                        `json:"name,omitempty"`
	Properties *RoleAssignmentProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RoleAssignment_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-08-01-preview"
func (assignment RoleAssignment_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (assignment *RoleAssignment_Spec_ARM) GetName() string {
	return assignment.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Authorization/roleAssignments"
func (assignment *RoleAssignment_Spec_ARM) GetType() string {
	return "Microsoft.Authorization/roleAssignments"
}

// Deprecated version of RoleAssignmentProperties. Use v1beta20200801preview.RoleAssignmentProperties instead
type RoleAssignmentProperties_ARM struct {
	Condition                          *string                                 `json:"condition,omitempty"`
	ConditionVersion                   *string                                 `json:"conditionVersion,omitempty"`
	DelegatedManagedIdentityResourceId *string                                 `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string                                 `json:"description,omitempty"`
	PrincipalId                        *string                                 `json:"principalId,omitempty" optionalConfigMapPair:"PrincipalId"`
	PrincipalType                      *RoleAssignmentProperties_PrincipalType `json:"principalType,omitempty"`
	RoleDefinitionId                   *string                                 `json:"roleDefinitionId,omitempty"`
}
