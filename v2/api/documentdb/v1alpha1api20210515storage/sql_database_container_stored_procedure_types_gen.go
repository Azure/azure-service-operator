// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"fmt"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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
// Storage version of v1alpha1api20210515.SqlDatabaseContainerStoredProcedure
// Deprecated version of SqlDatabaseContainerStoredProcedure. Use v1beta20210515.SqlDatabaseContainerStoredProcedure instead
type SqlDatabaseContainerStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerStoredProcedure{}

// GetConditions returns the conditions of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetConditions() conditions.Conditions {
	return procedure.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (procedure *SqlDatabaseContainerStoredProcedure) SetConditions(conditions conditions.Conditions) {
	procedure.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerStoredProcedure{}

// ConvertFrom populates our SqlDatabaseContainerStoredProcedure from the provided hub SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabaseContainerStoredProcedure)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerStoredProcedure but received %T instead", hub)
	}

	return procedure.AssignPropertiesFromSqlDatabaseContainerStoredProcedure(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseContainerStoredProcedure)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerStoredProcedure but received %T instead", hub)
	}

	return procedure.AssignPropertiesToSqlDatabaseContainerStoredProcedure(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerStoredProcedure{}

// AzureName returns the Azure name of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) AzureName() string {
	return procedure.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedure SqlDatabaseContainerStoredProcedure) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetSpec() genruntime.ConvertibleSpec {
	return &procedure.Spec
}

// GetStatus returns the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetStatus() genruntime.ConvertibleStatus {
	return &procedure.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedure *SqlDatabaseContainerStoredProcedure) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

// NewEmptyStatus returns a new empty (blank) status
func (procedure *SqlDatabaseContainerStoredProcedure) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (procedure *SqlDatabaseContainerStoredProcedure) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(procedure.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  procedure.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS); ok {
		procedure.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	procedure.Status = st
	return nil
}

// AssignPropertiesFromSqlDatabaseContainerStoredProcedure populates our SqlDatabaseContainerStoredProcedure from the provided source SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesFromSqlDatabaseContainerStoredProcedure(source *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	procedure.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec() to populate field Spec")
	}
	procedure.Spec = spec

	// Status
	var status DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
	err = status.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS() to populate field Status")
	}
	procedure.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerStoredProcedure populates the provided destination SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesToSqlDatabaseContainerStoredProcedure(destination *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	destination.ObjectMeta = *procedure.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
	err := procedure.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
	err = procedure.Status.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (procedure *SqlDatabaseContainerStoredProcedure) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: procedure.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerStoredProcedure",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.SqlDatabaseContainerStoredProcedure
// Deprecated version of SqlDatabaseContainerStoredProcedure. Use v1beta20210515.SqlDatabaseContainerStoredProcedure instead
type SqlDatabaseContainerStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerStoredProcedure `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
type DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/DatabaseAccountsSqlDatabasesContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"DatabaseAccountsSqlDatabasesContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlStoredProcedureResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec from the provided source
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec)
	if ok {
		// Populate our instance from source
		return procedure.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = procedure.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec)
	if ok {
		// Populate destination from our instance
		return procedure.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec{}
	err := procedure.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec populates our DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec from the provided source DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(source *v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	procedure.AzureName = source.AzureName

	// Location
	procedure.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		procedure.Options = &option
	} else {
		procedure.Options = nil
	}

	// OriginalVersion
	procedure.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		procedure.Owner = &owner
	} else {
		procedure.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureResource
		err := resource.AssignPropertiesFromSqlStoredProcedureResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureResource() to populate field Resource")
		}
		procedure.Resource = &resource
	} else {
		procedure.Resource = nil
	}

	// Tags
	procedure.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		procedure.PropertyBag = propertyBag
	} else {
		procedure.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec populates the provided destination DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec from our DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec(destination *v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(procedure.PropertyBag)

	// AzureName
	destination.AzureName = procedure.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(procedure.Location)

	// Options
	if procedure.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := procedure.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = procedure.OriginalVersion

	// Owner
	if procedure.Owner != nil {
		owner := procedure.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if procedure.Resource != nil {
		var resource v20210515s.SqlStoredProcedureResource
		err := procedure.Resource.AssignPropertiesToSqlStoredProcedureResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(procedure.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
// Deprecated version of DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS. Use v1beta20210515.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS instead
type DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *SqlStoredProcedureGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS{}

// ConvertStatusFrom populates our DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS from the provided source
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS)
	if ok {
		// Populate our instance from source
		return procedure.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = procedure.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS)
	if ok {
		// Populate destination from our instance
		return procedure.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS{}
	err := procedure.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS populates our DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS from the provided source DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(source *v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	procedure.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	procedure.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	procedure.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	procedure.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureGetProperties_Resource_STATUS
		err := resource.AssignPropertiesFromSqlStoredProcedureGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureGetProperties_Resource_STATUS() to populate field Resource")
		}
		procedure.Resource = &resource
	} else {
		procedure.Resource = nil
	}

	// Tags
	procedure.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	procedure.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		procedure.PropertyBag = propertyBag
	} else {
		procedure.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS populates the provided destination DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS from our DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS
func (procedure *DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS(destination *v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedure_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(procedure.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(procedure.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(procedure.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(procedure.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(procedure.Name)

	// Resource
	if procedure.Resource != nil {
		var resource v20210515s.SqlStoredProcedureGetProperties_Resource_STATUS
		err := procedure.Resource.AssignPropertiesToSqlStoredProcedureGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(procedure.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(procedure.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlStoredProcedureGetProperties_Resource_STATUS
// Deprecated version of SqlStoredProcedureGetProperties_Resource_STATUS. Use v1beta20210515.SqlStoredProcedureGetProperties_Resource_STATUS instead
type SqlStoredProcedureGetProperties_Resource_STATUS struct {
	Body        *string                `json:"body,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

// AssignPropertiesFromSqlStoredProcedureGetProperties_Resource_STATUS populates our SqlStoredProcedureGetProperties_Resource_STATUS from the provided source SqlStoredProcedureGetProperties_Resource_STATUS
func (resource *SqlStoredProcedureGetProperties_Resource_STATUS) AssignPropertiesFromSqlStoredProcedureGetProperties_Resource_STATUS(source *v20210515s.SqlStoredProcedureGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Rid
	resource.Rid = genruntime.ClonePointerToString(source.Rid)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		resource.Ts = &t
	} else {
		resource.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureGetProperties_Resource_STATUS populates the provided destination SqlStoredProcedureGetProperties_Resource_STATUS from our SqlStoredProcedureGetProperties_Resource_STATUS
func (resource *SqlStoredProcedureGetProperties_Resource_STATUS) AssignPropertiesToSqlStoredProcedureGetProperties_Resource_STATUS(destination *v20210515s.SqlStoredProcedureGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Rid
	destination.Rid = genruntime.ClonePointerToString(resource.Rid)

	// Ts
	if resource.Ts != nil {
		t := *resource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlStoredProcedureResource
// Deprecated version of SqlStoredProcedureResource. Use v1beta20210515.SqlStoredProcedureResource instead
type SqlStoredProcedureResource struct {
	Body        *string                `json:"body,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSqlStoredProcedureResource populates our SqlStoredProcedureResource from the provided source SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignPropertiesFromSqlStoredProcedureResource(source *v20210515s.SqlStoredProcedureResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureResource populates the provided destination SqlStoredProcedureResource from our SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignPropertiesToSqlStoredProcedureResource(destination *v20210515s.SqlStoredProcedureResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerStoredProcedure{}, &SqlDatabaseContainerStoredProcedureList{})
}
