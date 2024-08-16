// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=authorization.azure.com,resources=roleassignments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=authorization.azure.com,resources={roleassignments/status,roleassignments/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220401.RoleAssignment
// Generator information:
// - Generated from: /authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/authorization-RoleAssignmentsCalls.json
// - ARM URI: /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}
type RoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RoleAssignment_Spec   `json:"spec,omitempty"`
	Status            RoleAssignment_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RoleAssignment{}

// GetConditions returns the conditions of the resource
func (assignment *RoleAssignment) GetConditions() conditions.Conditions {
	return assignment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (assignment *RoleAssignment) SetConditions(conditions conditions.Conditions) {
	assignment.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *RoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-04-01"
func (assignment RoleAssignment) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (assignment *RoleAssignment) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeExtension
}

// GetSpec returns the specification of this resource
func (assignment *RoleAssignment) GetSpec() genruntime.ConvertibleSpec {
	return &assignment.Spec
}

// GetStatus returns the status of this resource
func (assignment *RoleAssignment) GetStatus() genruntime.ConvertibleStatus {
	return &assignment.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (assignment *RoleAssignment) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Authorization/roleAssignments"
func (assignment *RoleAssignment) GetType() string {
	return "Microsoft.Authorization/roleAssignments"
}

// NewEmptyStatus returns a new empty (blank) status
func (assignment *RoleAssignment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RoleAssignment_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (assignment *RoleAssignment) Owner() *genruntime.ResourceReference {
	return assignment.Spec.Owner.AsResourceReference()
}

// SetStatus sets the status of this resource
func (assignment *RoleAssignment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RoleAssignment_STATUS); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st RoleAssignment_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// Hub marks that this RoleAssignment is the hub type for conversion
func (assignment *RoleAssignment) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *RoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion,
		Kind:    "RoleAssignment",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220401.RoleAssignment
// Generator information:
// - Generated from: /authorization/resource-manager/Microsoft.Authorization/stable/2022-04-01/authorization-RoleAssignmentsCalls.json
// - ARM URI: /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Storage version of v1api20220401.APIVersion
// +kubebuilder:validation:Enum={"2022-04-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-04-01")

// Storage version of v1api20220401.RoleAssignment_Spec
type RoleAssignment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName        string  `json:"azureName,omitempty"`
	Condition        *string `json:"condition,omitempty"`
	ConditionVersion *string `json:"conditionVersion,omitempty"`

	// DelegatedManagedIdentityResourceReference: Id of the delegated managed identity resource
	DelegatedManagedIdentityResourceReference *genruntime.ResourceReference `armReference:"DelegatedManagedIdentityResourceId" json:"delegatedManagedIdentityResourceReference,omitempty"`
	Description                               *string                       `json:"description,omitempty"`
	OriginalVersion                           string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. This resource is an
	// extension resource, which means that any other Azure resource can be its owner.
	Owner                 *genruntime.ArbitraryOwnerReference `json:"owner,omitempty"`
	PrincipalId           *string                             `json:"principalId,omitempty" optionalConfigMapPair:"PrincipalId"`
	PrincipalIdFromConfig *genruntime.ConfigMapReference      `json:"principalIdFromConfig,omitempty" optionalConfigMapPair:"PrincipalId"`
	PrincipalType         *string                             `json:"principalType,omitempty"`
	PropertyBag           genruntime.PropertyBag              `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	// RoleDefinitionReference: The role definition ID.
	RoleDefinitionReference *genruntime.ResourceReference `armReference:"RoleDefinitionId" json:"roleDefinitionReference,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RoleAssignment_Spec{}

// ConvertSpecFrom populates our RoleAssignment_Spec from the provided source
func (assignment *RoleAssignment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == assignment {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(assignment)
}

// ConvertSpecTo populates the provided destination from our RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == assignment {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(assignment)
}

// Storage version of v1api20220401.RoleAssignment_STATUS
// Role Assignments
type RoleAssignment_STATUS struct {
	Condition                          *string                `json:"condition,omitempty"`
	ConditionVersion                   *string                `json:"conditionVersion,omitempty"`
	Conditions                         []conditions.Condition `json:"conditions,omitempty"`
	CreatedBy                          *string                `json:"createdBy,omitempty"`
	CreatedOn                          *string                `json:"createdOn,omitempty"`
	DelegatedManagedIdentityResourceId *string                `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string                `json:"description,omitempty"`
	Id                                 *string                `json:"id,omitempty"`
	Name                               *string                `json:"name,omitempty"`
	PrincipalId                        *string                `json:"principalId,omitempty"`
	PrincipalType                      *string                `json:"principalType,omitempty"`
	PropertyBag                        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RoleDefinitionId                   *string                `json:"roleDefinitionId,omitempty"`
	Scope                              *string                `json:"scope,omitempty"`
	Type                               *string                `json:"type,omitempty"`
	UpdatedBy                          *string                `json:"updatedBy,omitempty"`
	UpdatedOn                          *string                `json:"updatedOn,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RoleAssignment_STATUS{}

// ConvertStatusFrom populates our RoleAssignment_STATUS from the provided source
func (assignment *RoleAssignment_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == assignment {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(assignment)
}

// ConvertStatusTo populates the provided destination from our RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == assignment {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(assignment)
}

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
