// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200801previewstorage

import (
	"fmt"
	v20220401s "github.com/Azure/azure-service-operator/v2/api/authorization/v1api20220401storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20200801preview.RoleAssignment
// Generator information:
// - Generated from: /authorization/resource-manager/Microsoft.Authorization/preview/2020-08-01-preview/authorization-RoleAssignmentsCalls.json
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

var _ conversion.Convertible = &RoleAssignment{}

// ConvertFrom populates our RoleAssignment from the provided hub RoleAssignment
func (assignment *RoleAssignment) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20220401s.RoleAssignment)
	if !ok {
		return fmt.Errorf("expected authorization/v1api20220401storage/RoleAssignment but received %T instead", hub)
	}

	return assignment.AssignProperties_From_RoleAssignment(source)
}

// ConvertTo populates the provided hub RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20220401s.RoleAssignment)
	if !ok {
		return fmt.Errorf("expected authorization/v1api20220401storage/RoleAssignment but received %T instead", hub)
	}

	return assignment.AssignProperties_To_RoleAssignment(destination)
}

var _ genruntime.KubernetesResource = &RoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *RoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-08-01-preview"
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

// AssignProperties_From_RoleAssignment populates our RoleAssignment from the provided source RoleAssignment
func (assignment *RoleAssignment) AssignProperties_From_RoleAssignment(source *v20220401s.RoleAssignment) error {

	// ObjectMeta
	assignment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RoleAssignment_Spec
	err := spec.AssignProperties_From_RoleAssignment_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RoleAssignment_Spec() to populate field Spec")
	}
	assignment.Spec = spec

	// Status
	var status RoleAssignment_STATUS
	err = status.AssignProperties_From_RoleAssignment_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_RoleAssignment_STATUS() to populate field Status")
	}
	assignment.Status = status

	// Invoke the augmentConversionForRoleAssignment interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForRoleAssignment); ok {
		err := augmentedAssignment.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RoleAssignment populates the provided destination RoleAssignment from our RoleAssignment
func (assignment *RoleAssignment) AssignProperties_To_RoleAssignment(destination *v20220401s.RoleAssignment) error {

	// ObjectMeta
	destination.ObjectMeta = *assignment.ObjectMeta.DeepCopy()

	// Spec
	var spec v20220401s.RoleAssignment_Spec
	err := assignment.Spec.AssignProperties_To_RoleAssignment_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RoleAssignment_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20220401s.RoleAssignment_STATUS
	err = assignment.Status.AssignProperties_To_RoleAssignment_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_RoleAssignment_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForRoleAssignment interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForRoleAssignment); ok {
		err := augmentedAssignment.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *RoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion,
		Kind:    "RoleAssignment",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20200801preview.RoleAssignment
// Generator information:
// - Generated from: /authorization/resource-manager/Microsoft.Authorization/preview/2020-08-01-preview/authorization-RoleAssignmentsCalls.json
// - ARM URI: /{scope}/providers/Microsoft.Authorization/roleAssignments/{roleAssignmentName}
type RoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoleAssignment `json:"items"`
}

// Storage version of v1api20200801preview.APIVersion
// +kubebuilder:validation:Enum={"2020-08-01-preview"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-08-01-preview")

type augmentConversionForRoleAssignment interface {
	AssignPropertiesFrom(src *v20220401s.RoleAssignment) error
	AssignPropertiesTo(dst *v20220401s.RoleAssignment) error
}

// Storage version of v1api20200801preview.RoleAssignment_Spec
type RoleAssignment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                          string  `json:"azureName,omitempty"`
	Condition                          *string `json:"condition,omitempty"`
	ConditionVersion                   *string `json:"conditionVersion,omitempty"`
	DelegatedManagedIdentityResourceId *string `json:"delegatedManagedIdentityResourceId,omitempty"`
	Description                        *string `json:"description,omitempty"`
	OriginalVersion                    string  `json:"originalVersion,omitempty"`

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
	src, ok := source.(*v20220401s.RoleAssignment_Spec)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_RoleAssignment_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20220401s.RoleAssignment_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_RoleAssignment_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20220401s.RoleAssignment_Spec)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_RoleAssignment_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20220401s.RoleAssignment_Spec{}
	err := assignment.AssignProperties_To_RoleAssignment_Spec(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_RoleAssignment_Spec populates our RoleAssignment_Spec from the provided source RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) AssignProperties_From_RoleAssignment_Spec(source *v20220401s.RoleAssignment_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	assignment.AzureName = source.AzureName

	// Condition
	assignment.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignment.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	assignment.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignment.Description = genruntime.ClonePointerToString(source.Description)

	// OriginalVersion
	assignment.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		assignment.Owner = &owner
	} else {
		assignment.Owner = nil
	}

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalIdFromConfig
	if source.PrincipalIdFromConfig != nil {
		principalIdFromConfig := source.PrincipalIdFromConfig.Copy()
		assignment.PrincipalIdFromConfig = &principalIdFromConfig
	} else {
		assignment.PrincipalIdFromConfig = nil
	}

	// PrincipalType
	assignment.PrincipalType = genruntime.ClonePointerToString(source.PrincipalType)

	// RoleDefinitionReference
	if source.RoleDefinitionReference != nil {
		roleDefinitionReference := source.RoleDefinitionReference.Copy()
		assignment.RoleDefinitionReference = &roleDefinitionReference
	} else {
		assignment.RoleDefinitionReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		assignment.PropertyBag = propertyBag
	} else {
		assignment.PropertyBag = nil
	}

	// Invoke the augmentConversionForRoleAssignment_Spec interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForRoleAssignment_Spec); ok {
		err := augmentedAssignment.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RoleAssignment_Spec populates the provided destination RoleAssignment_Spec from our RoleAssignment_Spec
func (assignment *RoleAssignment_Spec) AssignProperties_To_RoleAssignment_Spec(destination *v20220401s.RoleAssignment_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(assignment.PropertyBag)

	// AzureName
	destination.AzureName = assignment.AzureName

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignment.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignment.ConditionVersion)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignment.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignment.Description)

	// OriginalVersion
	destination.OriginalVersion = assignment.OriginalVersion

	// Owner
	if assignment.Owner != nil {
		owner := assignment.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalIdFromConfig
	if assignment.PrincipalIdFromConfig != nil {
		principalIdFromConfig := assignment.PrincipalIdFromConfig.Copy()
		destination.PrincipalIdFromConfig = &principalIdFromConfig
	} else {
		destination.PrincipalIdFromConfig = nil
	}

	// PrincipalType
	destination.PrincipalType = genruntime.ClonePointerToString(assignment.PrincipalType)

	// RoleDefinitionReference
	if assignment.RoleDefinitionReference != nil {
		roleDefinitionReference := assignment.RoleDefinitionReference.Copy()
		destination.RoleDefinitionReference = &roleDefinitionReference
	} else {
		destination.RoleDefinitionReference = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRoleAssignment_Spec interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForRoleAssignment_Spec); ok {
		err := augmentedAssignment.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20200801preview.RoleAssignment_STATUS
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
	src, ok := source.(*v20220401s.RoleAssignment_STATUS)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_RoleAssignment_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20220401s.RoleAssignment_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_RoleAssignment_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20220401s.RoleAssignment_STATUS)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_RoleAssignment_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20220401s.RoleAssignment_STATUS{}
	err := assignment.AssignProperties_To_RoleAssignment_STATUS(dst)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_RoleAssignment_STATUS populates our RoleAssignment_STATUS from the provided source RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) AssignProperties_From_RoleAssignment_STATUS(source *v20220401s.RoleAssignment_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Condition
	assignment.Condition = genruntime.ClonePointerToString(source.Condition)

	// ConditionVersion
	assignment.ConditionVersion = genruntime.ClonePointerToString(source.ConditionVersion)

	// Conditions
	assignment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedBy
	assignment.CreatedBy = genruntime.ClonePointerToString(source.CreatedBy)

	// CreatedOn
	assignment.CreatedOn = genruntime.ClonePointerToString(source.CreatedOn)

	// DelegatedManagedIdentityResourceId
	assignment.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(source.DelegatedManagedIdentityResourceId)

	// Description
	assignment.Description = genruntime.ClonePointerToString(source.Description)

	// Id
	assignment.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	assignment.Name = genruntime.ClonePointerToString(source.Name)

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// PrincipalType
	assignment.PrincipalType = genruntime.ClonePointerToString(source.PrincipalType)

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// Type
	assignment.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedBy
	assignment.UpdatedBy = genruntime.ClonePointerToString(source.UpdatedBy)

	// UpdatedOn
	assignment.UpdatedOn = genruntime.ClonePointerToString(source.UpdatedOn)

	// Update the property bag
	if len(propertyBag) > 0 {
		assignment.PropertyBag = propertyBag
	} else {
		assignment.PropertyBag = nil
	}

	// Invoke the augmentConversionForRoleAssignment_STATUS interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForRoleAssignment_STATUS); ok {
		err := augmentedAssignment.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_RoleAssignment_STATUS populates the provided destination RoleAssignment_STATUS from our RoleAssignment_STATUS
func (assignment *RoleAssignment_STATUS) AssignProperties_To_RoleAssignment_STATUS(destination *v20220401s.RoleAssignment_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(assignment.PropertyBag)

	// Condition
	destination.Condition = genruntime.ClonePointerToString(assignment.Condition)

	// ConditionVersion
	destination.ConditionVersion = genruntime.ClonePointerToString(assignment.ConditionVersion)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(assignment.Conditions)

	// CreatedBy
	destination.CreatedBy = genruntime.ClonePointerToString(assignment.CreatedBy)

	// CreatedOn
	destination.CreatedOn = genruntime.ClonePointerToString(assignment.CreatedOn)

	// DelegatedManagedIdentityResourceId
	destination.DelegatedManagedIdentityResourceId = genruntime.ClonePointerToString(assignment.DelegatedManagedIdentityResourceId)

	// Description
	destination.Description = genruntime.ClonePointerToString(assignment.Description)

	// Id
	destination.Id = genruntime.ClonePointerToString(assignment.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(assignment.Name)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// PrincipalType
	destination.PrincipalType = genruntime.ClonePointerToString(assignment.PrincipalType)

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// Type
	destination.Type = genruntime.ClonePointerToString(assignment.Type)

	// UpdatedBy
	destination.UpdatedBy = genruntime.ClonePointerToString(assignment.UpdatedBy)

	// UpdatedOn
	destination.UpdatedOn = genruntime.ClonePointerToString(assignment.UpdatedOn)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForRoleAssignment_STATUS interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForRoleAssignment_STATUS); ok {
		err := augmentedAssignment.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForRoleAssignment_Spec interface {
	AssignPropertiesFrom(src *v20220401s.RoleAssignment_Spec) error
	AssignPropertiesTo(dst *v20220401s.RoleAssignment_Spec) error
}

type augmentConversionForRoleAssignment_STATUS interface {
	AssignPropertiesFrom(src *v20220401s.RoleAssignment_STATUS) error
	AssignPropertiesTo(dst *v20220401s.RoleAssignment_STATUS) error
}

func init() {
	SchemeBuilder.Register(&RoleAssignment{}, &RoleAssignmentList{})
}
