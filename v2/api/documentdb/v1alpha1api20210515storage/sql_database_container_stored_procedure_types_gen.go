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
	Spec              DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec `json:"spec,omitempty"`
	Status            SqlStoredProcedureGetResults_STATUS                           `json:"status,omitempty"`
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

	return procedure.AssignProperties_From_SqlDatabaseContainerStoredProcedure(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseContainerStoredProcedure)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerStoredProcedure but received %T instead", hub)
	}

	return procedure.AssignProperties_To_SqlDatabaseContainerStoredProcedure(destination)
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
	return &SqlStoredProcedureGetResults_STATUS{}
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
	if st, ok := status.(*SqlStoredProcedureGetResults_STATUS); ok {
		procedure.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlStoredProcedureGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	procedure.Status = st
	return nil
}

// AssignProperties_From_SqlDatabaseContainerStoredProcedure populates our SqlDatabaseContainerStoredProcedure from the provided source SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignProperties_From_SqlDatabaseContainerStoredProcedure(source *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	procedure.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec() to populate field Spec")
	}
	procedure.Spec = spec

	// Status
	var status SqlStoredProcedureGetResults_STATUS
	err = status.AssignProperties_From_SqlStoredProcedureGetResults_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_SqlStoredProcedureGetResults_STATUS() to populate field Status")
	}
	procedure.Status = status

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerStoredProcedure populates the provided destination SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignProperties_To_SqlDatabaseContainerStoredProcedure(destination *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	destination.ObjectMeta = *procedure.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
	err := procedure.Spec.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.SqlStoredProcedureGetResults_STATUS
	err = procedure.Status.AssignProperties_To_SqlStoredProcedureGetResults_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_SqlStoredProcedureGetResults_STATUS() to populate field Status")
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

// Storage version of v1alpha1api20210515.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
type DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlStoredProcedureResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec from the provided source
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec)
	if ok {
		// Populate our instance from source
		return procedure.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = procedure.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec)
	if ok {
		// Populate destination from our instance
		return procedure.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}
	err := procedure.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(dst)
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

// AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec populates our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec from the provided source DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(source *v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	procedure.AzureName = source.AzureName

	// Location
	procedure.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignProperties_From_CreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CreateUpdateOptions() to populate field Options")
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
		err := resource.AssignProperties_From_SqlStoredProcedureResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlStoredProcedureResource() to populate field Resource")
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

// AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec populates the provided destination DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec from our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(destination *v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(procedure.PropertyBag)

	// AzureName
	destination.AzureName = procedure.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(procedure.Location)

	// Options
	if procedure.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := procedure.Options.AssignProperties_To_CreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CreateUpdateOptions() to populate field Options")
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
		err := procedure.Resource.AssignProperties_To_SqlStoredProcedureResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlStoredProcedureResource() to populate field Resource")
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

// Storage version of v1alpha1api20210515.SqlStoredProcedureGetResults_STATUS
// Deprecated version of SqlStoredProcedureGetResults_STATUS. Use v1beta20210515.SqlStoredProcedureGetResults_STATUS instead
type SqlStoredProcedureGetResults_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *SqlStoredProcedureGetProperties_STATUS_Resource `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlStoredProcedureGetResults_STATUS{}

// ConvertStatusFrom populates our SqlStoredProcedureGetResults_STATUS from the provided source
func (results *SqlStoredProcedureGetResults_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.SqlStoredProcedureGetResults_STATUS)
	if ok {
		// Populate our instance from source
		return results.AssignProperties_From_SqlStoredProcedureGetResults_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.SqlStoredProcedureGetResults_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignProperties_From_SqlStoredProcedureGetResults_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlStoredProcedureGetResults_STATUS
func (results *SqlStoredProcedureGetResults_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.SqlStoredProcedureGetResults_STATUS)
	if ok {
		// Populate destination from our instance
		return results.AssignProperties_To_SqlStoredProcedureGetResults_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.SqlStoredProcedureGetResults_STATUS{}
	err := results.AssignProperties_To_SqlStoredProcedureGetResults_STATUS(dst)
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

// AssignProperties_From_SqlStoredProcedureGetResults_STATUS populates our SqlStoredProcedureGetResults_STATUS from the provided source SqlStoredProcedureGetResults_STATUS
func (results *SqlStoredProcedureGetResults_STATUS) AssignProperties_From_SqlStoredProcedureGetResults_STATUS(source *v20210515s.SqlStoredProcedureGetResults_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	results.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	results.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	results.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	results.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureGetProperties_STATUS_Resource
		err := resource.AssignProperties_From_SqlStoredProcedureGetProperties_STATUS_Resource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlStoredProcedureGetProperties_STATUS_Resource() to populate field Resource")
		}
		results.Resource = &resource
	} else {
		results.Resource = nil
	}

	// Tags
	results.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	results.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		results.PropertyBag = propertyBag
	} else {
		results.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_SqlStoredProcedureGetResults_STATUS populates the provided destination SqlStoredProcedureGetResults_STATUS from our SqlStoredProcedureGetResults_STATUS
func (results *SqlStoredProcedureGetResults_STATUS) AssignProperties_To_SqlStoredProcedureGetResults_STATUS(destination *v20210515s.SqlStoredProcedureGetResults_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(results.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(results.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(results.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(results.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(results.Name)

	// Resource
	if results.Resource != nil {
		var resource v20210515s.SqlStoredProcedureGetProperties_STATUS_Resource
		err := results.Resource.AssignProperties_To_SqlStoredProcedureGetProperties_STATUS_Resource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlStoredProcedureGetProperties_STATUS_Resource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(results.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(results.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlStoredProcedureGetProperties_STATUS_Resource
// Deprecated version of SqlStoredProcedureGetProperties_STATUS_Resource. Use v1beta20210515.SqlStoredProcedureGetProperties_STATUS_Resource instead
type SqlStoredProcedureGetProperties_STATUS_Resource struct {
	Body        *string                `json:"body,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

// AssignProperties_From_SqlStoredProcedureGetProperties_STATUS_Resource populates our SqlStoredProcedureGetProperties_STATUS_Resource from the provided source SqlStoredProcedureGetProperties_STATUS_Resource
func (resource *SqlStoredProcedureGetProperties_STATUS_Resource) AssignProperties_From_SqlStoredProcedureGetProperties_STATUS_Resource(source *v20210515s.SqlStoredProcedureGetProperties_STATUS_Resource) error {
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

// AssignProperties_To_SqlStoredProcedureGetProperties_STATUS_Resource populates the provided destination SqlStoredProcedureGetProperties_STATUS_Resource from our SqlStoredProcedureGetProperties_STATUS_Resource
func (resource *SqlStoredProcedureGetProperties_STATUS_Resource) AssignProperties_To_SqlStoredProcedureGetProperties_STATUS_Resource(destination *v20210515s.SqlStoredProcedureGetProperties_STATUS_Resource) error {
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

// AssignProperties_From_SqlStoredProcedureResource populates our SqlStoredProcedureResource from the provided source SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignProperties_From_SqlStoredProcedureResource(source *v20210515s.SqlStoredProcedureResource) error {
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

// AssignProperties_To_SqlStoredProcedureResource populates the provided destination SqlStoredProcedureResource from our SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignProperties_To_SqlStoredProcedureResource(destination *v20210515s.SqlStoredProcedureResource) error {
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
