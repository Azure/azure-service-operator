// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115/storage"
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
// Storage version of v1api20210515.SqlRoleAssignment
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/rbac.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlRoleAssignments/{roleAssignmentId}
type SqlRoleAssignment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlRoleAssignment_Spec   `json:"spec,omitempty"`
	Status            SqlRoleAssignment_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlRoleAssignment{}

// GetConditions returns the conditions of the resource
func (assignment *SqlRoleAssignment) GetConditions() conditions.Conditions {
	return assignment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (assignment *SqlRoleAssignment) SetConditions(conditions conditions.Conditions) {
	assignment.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlRoleAssignment{}

// ConvertFrom populates our SqlRoleAssignment from the provided hub SqlRoleAssignment
func (assignment *SqlRoleAssignment) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.SqlRoleAssignment)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/SqlRoleAssignment but received %T instead", hub)
	}

	return assignment.AssignProperties_From_SqlRoleAssignment(source)
}

// ConvertTo populates the provided hub SqlRoleAssignment from our SqlRoleAssignment
func (assignment *SqlRoleAssignment) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.SqlRoleAssignment)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/SqlRoleAssignment but received %T instead", hub)
	}

	return assignment.AssignProperties_To_SqlRoleAssignment(destination)
}

var _ genruntime.KubernetesResource = &SqlRoleAssignment{}

// AzureName returns the Azure name of the resource
func (assignment *SqlRoleAssignment) AzureName() string {
	return assignment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (assignment SqlRoleAssignment) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceScope returns the scope of the resource
func (assignment *SqlRoleAssignment) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (assignment *SqlRoleAssignment) GetSpec() genruntime.ConvertibleSpec {
	return &assignment.Spec
}

// GetStatus returns the status of this resource
func (assignment *SqlRoleAssignment) GetStatus() genruntime.ConvertibleStatus {
	return &assignment.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (assignment *SqlRoleAssignment) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
func (assignment *SqlRoleAssignment) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlRoleAssignments"
}

// NewEmptyStatus returns a new empty (blank) status
func (assignment *SqlRoleAssignment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlRoleAssignment_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (assignment *SqlRoleAssignment) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(assignment.Spec)
	return assignment.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (assignment *SqlRoleAssignment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlRoleAssignment_STATUS); ok {
		assignment.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlRoleAssignment_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	assignment.Status = st
	return nil
}

// AssignProperties_From_SqlRoleAssignment populates our SqlRoleAssignment from the provided source SqlRoleAssignment
func (assignment *SqlRoleAssignment) AssignProperties_From_SqlRoleAssignment(source *storage.SqlRoleAssignment) error {

	// ObjectMeta
	assignment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec SqlRoleAssignment_Spec
	err := spec.AssignProperties_From_SqlRoleAssignment_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_SqlRoleAssignment_Spec() to populate field Spec")
	}
	assignment.Spec = spec

	// Status
	var status SqlRoleAssignment_STATUS
	err = status.AssignProperties_From_SqlRoleAssignment_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_SqlRoleAssignment_STATUS() to populate field Status")
	}
	assignment.Status = status

	// Invoke the augmentConversionForSqlRoleAssignment interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForSqlRoleAssignment); ok {
		err := augmentedAssignment.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlRoleAssignment populates the provided destination SqlRoleAssignment from our SqlRoleAssignment
func (assignment *SqlRoleAssignment) AssignProperties_To_SqlRoleAssignment(destination *storage.SqlRoleAssignment) error {

	// ObjectMeta
	destination.ObjectMeta = *assignment.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.SqlRoleAssignment_Spec
	err := assignment.Spec.AssignProperties_To_SqlRoleAssignment_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_SqlRoleAssignment_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.SqlRoleAssignment_STATUS
	err = assignment.Status.AssignProperties_To_SqlRoleAssignment_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_SqlRoleAssignment_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForSqlRoleAssignment interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForSqlRoleAssignment); ok {
		err := augmentedAssignment.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (assignment *SqlRoleAssignment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: assignment.Spec.OriginalVersion,
		Kind:    "SqlRoleAssignment",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210515.SqlRoleAssignment
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/rbac.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlRoleAssignments/{roleAssignmentId}
type SqlRoleAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlRoleAssignment `json:"items"`
}

type augmentConversionForSqlRoleAssignment interface {
	AssignPropertiesFrom(src *storage.SqlRoleAssignment) error
	AssignPropertiesTo(dst *storage.SqlRoleAssignment) error
}

// Storage version of v1api20210515.SqlRoleAssignment_Spec
type SqlRoleAssignment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/DatabaseAccount resource
	Owner                 *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"DatabaseAccount"`
	PrincipalId           *string                            `json:"principalId,omitempty" optionalConfigMapPair:"PrincipalId"`
	PrincipalIdFromConfig *genruntime.ConfigMapReference     `json:"principalIdFromConfig,omitempty" optionalConfigMapPair:"PrincipalId"`
	PropertyBag           genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RoleDefinitionId      *string                            `json:"roleDefinitionId,omitempty"`
	Scope                 *string                            `json:"scope,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SqlRoleAssignment_Spec{}

// ConvertSpecFrom populates our SqlRoleAssignment_Spec from the provided source
func (assignment *SqlRoleAssignment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.SqlRoleAssignment_Spec)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_SqlRoleAssignment_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlRoleAssignment_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_SqlRoleAssignment_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our SqlRoleAssignment_Spec
func (assignment *SqlRoleAssignment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.SqlRoleAssignment_Spec)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_SqlRoleAssignment_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlRoleAssignment_Spec{}
	err := assignment.AssignProperties_To_SqlRoleAssignment_Spec(dst)
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

// AssignProperties_From_SqlRoleAssignment_Spec populates our SqlRoleAssignment_Spec from the provided source SqlRoleAssignment_Spec
func (assignment *SqlRoleAssignment_Spec) AssignProperties_From_SqlRoleAssignment_Spec(source *storage.SqlRoleAssignment_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	assignment.AzureName = source.AzureName

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

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// Update the property bag
	if len(propertyBag) > 0 {
		assignment.PropertyBag = propertyBag
	} else {
		assignment.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlRoleAssignment_Spec interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForSqlRoleAssignment_Spec); ok {
		err := augmentedAssignment.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlRoleAssignment_Spec populates the provided destination SqlRoleAssignment_Spec from our SqlRoleAssignment_Spec
func (assignment *SqlRoleAssignment_Spec) AssignProperties_To_SqlRoleAssignment_Spec(destination *storage.SqlRoleAssignment_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(assignment.PropertyBag)

	// AzureName
	destination.AzureName = assignment.AzureName

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

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlRoleAssignment_Spec interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForSqlRoleAssignment_Spec); ok {
		err := augmentedAssignment.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210515.SqlRoleAssignment_STATUS
type SqlRoleAssignment_STATUS struct {
	Conditions       []conditions.Condition `json:"conditions,omitempty"`
	Id               *string                `json:"id,omitempty"`
	Name             *string                `json:"name,omitempty"`
	PrincipalId      *string                `json:"principalId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RoleDefinitionId *string                `json:"roleDefinitionId,omitempty"`
	Scope            *string                `json:"scope,omitempty"`
	Type             *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlRoleAssignment_STATUS{}

// ConvertStatusFrom populates our SqlRoleAssignment_STATUS from the provided source
func (assignment *SqlRoleAssignment_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.SqlRoleAssignment_STATUS)
	if ok {
		// Populate our instance from source
		return assignment.AssignProperties_From_SqlRoleAssignment_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlRoleAssignment_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = assignment.AssignProperties_From_SqlRoleAssignment_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlRoleAssignment_STATUS
func (assignment *SqlRoleAssignment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.SqlRoleAssignment_STATUS)
	if ok {
		// Populate destination from our instance
		return assignment.AssignProperties_To_SqlRoleAssignment_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlRoleAssignment_STATUS{}
	err := assignment.AssignProperties_To_SqlRoleAssignment_STATUS(dst)
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

// AssignProperties_From_SqlRoleAssignment_STATUS populates our SqlRoleAssignment_STATUS from the provided source SqlRoleAssignment_STATUS
func (assignment *SqlRoleAssignment_STATUS) AssignProperties_From_SqlRoleAssignment_STATUS(source *storage.SqlRoleAssignment_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	assignment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	assignment.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	assignment.Name = genruntime.ClonePointerToString(source.Name)

	// PrincipalId
	assignment.PrincipalId = genruntime.ClonePointerToString(source.PrincipalId)

	// RoleDefinitionId
	assignment.RoleDefinitionId = genruntime.ClonePointerToString(source.RoleDefinitionId)

	// Scope
	assignment.Scope = genruntime.ClonePointerToString(source.Scope)

	// Type
	assignment.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		assignment.PropertyBag = propertyBag
	} else {
		assignment.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlRoleAssignment_STATUS interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForSqlRoleAssignment_STATUS); ok {
		err := augmentedAssignment.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlRoleAssignment_STATUS populates the provided destination SqlRoleAssignment_STATUS from our SqlRoleAssignment_STATUS
func (assignment *SqlRoleAssignment_STATUS) AssignProperties_To_SqlRoleAssignment_STATUS(destination *storage.SqlRoleAssignment_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(assignment.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(assignment.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(assignment.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(assignment.Name)

	// PrincipalId
	destination.PrincipalId = genruntime.ClonePointerToString(assignment.PrincipalId)

	// RoleDefinitionId
	destination.RoleDefinitionId = genruntime.ClonePointerToString(assignment.RoleDefinitionId)

	// Scope
	destination.Scope = genruntime.ClonePointerToString(assignment.Scope)

	// Type
	destination.Type = genruntime.ClonePointerToString(assignment.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlRoleAssignment_STATUS interface (if implemented) to customize the conversion
	var assignmentAsAny any = assignment
	if augmentedAssignment, ok := assignmentAsAny.(augmentConversionForSqlRoleAssignment_STATUS); ok {
		err := augmentedAssignment.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlRoleAssignment_Spec interface {
	AssignPropertiesFrom(src *storage.SqlRoleAssignment_Spec) error
	AssignPropertiesTo(dst *storage.SqlRoleAssignment_Spec) error
}

type augmentConversionForSqlRoleAssignment_STATUS interface {
	AssignPropertiesFrom(src *storage.SqlRoleAssignment_STATUS) error
	AssignPropertiesTo(dst *storage.SqlRoleAssignment_STATUS) error
}

func init() {
	SchemeBuilder.Register(&SqlRoleAssignment{}, &SqlRoleAssignmentList{})
}
