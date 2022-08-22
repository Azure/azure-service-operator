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
// Storage version of v1alpha1api20210515.SqlDatabase
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabases_Spec `json:"spec,omitempty"`
	Status            SqlDatabaseGetResults_STATUS      `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabase{}

// GetConditions returns the conditions of the resource
func (database *SqlDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *SqlDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabase{}

// ConvertFrom populates our SqlDatabase from the provided hub SqlDatabase
func (database *SqlDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesFromSqlDatabase(source)
}

// ConvertTo populates the provided hub SqlDatabase from our SqlDatabase
func (database *SqlDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesToSqlDatabase(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabase{}

// AzureName returns the Azure name of the resource
func (database *SqlDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database SqlDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (database *SqlDatabase) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *SqlDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *SqlDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
func (database *SqlDatabase) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *SqlDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseGetResults_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *SqlDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *SqlDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseGetResults_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignPropertiesFromSqlDatabase populates our SqlDatabase from the provided source SqlDatabase
func (database *SqlDatabase) AssignPropertiesFromSqlDatabase(source *v20210515s.SqlDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabases_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status SqlDatabaseGetResults_STATUS
	err = status.AssignPropertiesFromSqlDatabaseGetResultsSTATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseGetResultsSTATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabase populates the provided destination SqlDatabase from our SqlDatabase
func (database *SqlDatabase) AssignPropertiesToSqlDatabase(destination *v20210515s.SqlDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsSqlDatabases_Spec
	err := database.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.SqlDatabaseGetResults_STATUS
	err = database.Status.AssignPropertiesToSqlDatabaseGetResultsSTATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseGetResultsSTATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *SqlDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "SqlDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.SqlDatabase
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabases_Spec
type DatabaseAccountsSqlDatabases_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string               `json:"azureName,omitempty"`
	Location        *string              `json:"location,omitempty"`
	Options         *CreateUpdateOptions `json:"options,omitempty"`
	OriginalVersion string               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/DatabaseAccount resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"DatabaseAccount"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *SqlDatabaseResource               `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabases_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabases_Spec from the provided source
func (databases *DatabaseAccountsSqlDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabases_Spec)
	if ok {
		// Populate our instance from source
		return databases.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabases_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databases.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabases_Spec
func (databases *DatabaseAccountsSqlDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabases_Spec)
	if ok {
		// Populate destination from our instance
		return databases.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabases_Spec{}
	err := databases.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec populates our DatabaseAccountsSqlDatabases_Spec from the provided source DatabaseAccountsSqlDatabases_Spec
func (databases *DatabaseAccountsSqlDatabases_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(source *v20210515s.DatabaseAccountsSqlDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	databases.AzureName = source.AzureName

	// Location
	databases.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		databases.Options = &option
	} else {
		databases.Options = nil
	}

	// OriginalVersion
	databases.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		databases.Owner = &owner
	} else {
		databases.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseResource
		err := resource.AssignPropertiesFromSqlDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseResource() to populate field Resource")
		}
		databases.Resource = &resource
	} else {
		databases.Resource = nil
	}

	// Tags
	databases.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		databases.PropertyBag = propertyBag
	} else {
		databases.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesSpec populates the provided destination DatabaseAccountsSqlDatabases_Spec from our DatabaseAccountsSqlDatabases_Spec
func (databases *DatabaseAccountsSqlDatabases_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(destination *v20210515s.DatabaseAccountsSqlDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(databases.PropertyBag)

	// AzureName
	destination.AzureName = databases.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(databases.Location)

	// Options
	if databases.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := databases.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = databases.OriginalVersion

	// Owner
	if databases.Owner != nil {
		owner := databases.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if databases.Resource != nil {
		var resource v20210515s.SqlDatabaseResource
		err := databases.Resource.AssignPropertiesToSqlDatabaseResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(databases.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlDatabaseGetResults_STATUS
// Deprecated version of SqlDatabaseGetResults_STATUS. Use v1beta20210515.SqlDatabaseGetResults_STATUS instead
type SqlDatabaseGetResults_STATUS struct {
	Conditions  []conditions.Condition                    `json:"conditions,omitempty"`
	Id          *string                                   `json:"id,omitempty"`
	Location    *string                                   `json:"location,omitempty"`
	Name        *string                                   `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                   `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	Resource    *SqlDatabaseGetProperties_STATUS_Resource `json:"resource,omitempty"`
	Tags        map[string]string                         `json:"tags,omitempty"`
	Type        *string                                   `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseGetResults_STATUS{}

// ConvertStatusFrom populates our SqlDatabaseGetResults_STATUS from the provided source
func (results *SqlDatabaseGetResults_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.SqlDatabaseGetResults_STATUS)
	if ok {
		// Populate our instance from source
		return results.AssignPropertiesFromSqlDatabaseGetResultsSTATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.SqlDatabaseGetResults_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignPropertiesFromSqlDatabaseGetResultsSTATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseGetResults_STATUS
func (results *SqlDatabaseGetResults_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.SqlDatabaseGetResults_STATUS)
	if ok {
		// Populate destination from our instance
		return results.AssignPropertiesToSqlDatabaseGetResultsSTATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.SqlDatabaseGetResults_STATUS{}
	err := results.AssignPropertiesToSqlDatabaseGetResultsSTATUS(dst)
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

// AssignPropertiesFromSqlDatabaseGetResultsSTATUS populates our SqlDatabaseGetResults_STATUS from the provided source SqlDatabaseGetResults_STATUS
func (results *SqlDatabaseGetResults_STATUS) AssignPropertiesFromSqlDatabaseGetResultsSTATUS(source *v20210515s.SqlDatabaseGetResults_STATUS) error {
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

	// Options
	if source.Options != nil {
		var option OptionsResource_STATUS
		err := option.AssignPropertiesFromOptionsResourceSTATUS(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromOptionsResourceSTATUS() to populate field Options")
		}
		results.Options = &option
	} else {
		results.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseGetProperties_STATUS_Resource
		err := resource.AssignPropertiesFromSqlDatabaseGetPropertiesSTATUSResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseGetPropertiesSTATUSResource() to populate field Resource")
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

// AssignPropertiesToSqlDatabaseGetResultsSTATUS populates the provided destination SqlDatabaseGetResults_STATUS from our SqlDatabaseGetResults_STATUS
func (results *SqlDatabaseGetResults_STATUS) AssignPropertiesToSqlDatabaseGetResultsSTATUS(destination *v20210515s.SqlDatabaseGetResults_STATUS) error {
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

	// Options
	if results.Options != nil {
		var option v20210515s.OptionsResource_STATUS
		err := results.Options.AssignPropertiesToOptionsResourceSTATUS(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToOptionsResourceSTATUS() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if results.Resource != nil {
		var resource v20210515s.SqlDatabaseGetProperties_STATUS_Resource
		err := results.Resource.AssignPropertiesToSqlDatabaseGetPropertiesSTATUSResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseGetPropertiesSTATUSResource() to populate field Resource")
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

// Storage version of v1alpha1api20210515.SqlDatabaseGetProperties_STATUS_Resource
// Deprecated version of SqlDatabaseGetProperties_STATUS_Resource. Use v1beta20210515.SqlDatabaseGetProperties_STATUS_Resource instead
type SqlDatabaseGetProperties_STATUS_Resource struct {
	Colls       *string                `json:"_colls,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
	Users       *string                `json:"_users,omitempty"`
}

// AssignPropertiesFromSqlDatabaseGetPropertiesSTATUSResource populates our SqlDatabaseGetProperties_STATUS_Resource from the provided source SqlDatabaseGetProperties_STATUS_Resource
func (resource *SqlDatabaseGetProperties_STATUS_Resource) AssignPropertiesFromSqlDatabaseGetPropertiesSTATUSResource(source *v20210515s.SqlDatabaseGetProperties_STATUS_Resource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Colls
	resource.Colls = genruntime.ClonePointerToString(source.Colls)

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

	// Users
	resource.Users = genruntime.ClonePointerToString(source.Users)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseGetPropertiesSTATUSResource populates the provided destination SqlDatabaseGetProperties_STATUS_Resource from our SqlDatabaseGetProperties_STATUS_Resource
func (resource *SqlDatabaseGetProperties_STATUS_Resource) AssignPropertiesToSqlDatabaseGetPropertiesSTATUSResource(destination *v20210515s.SqlDatabaseGetProperties_STATUS_Resource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Colls
	destination.Colls = genruntime.ClonePointerToString(resource.Colls)

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

	// Users
	destination.Users = genruntime.ClonePointerToString(resource.Users)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.SqlDatabaseResource
// Deprecated version of SqlDatabaseResource. Use v1beta20210515.SqlDatabaseResource instead
type SqlDatabaseResource struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignPropertiesFromSqlDatabaseResource populates our SqlDatabaseResource from the provided source SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesFromSqlDatabaseResource(source *v20210515s.SqlDatabaseResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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

// AssignPropertiesToSqlDatabaseResource populates the provided destination SqlDatabaseResource from our SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesToSqlDatabaseResource(destination *v20210515s.SqlDatabaseResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

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
	SchemeBuilder.Register(&SqlDatabase{}, &SqlDatabaseList{})
}
