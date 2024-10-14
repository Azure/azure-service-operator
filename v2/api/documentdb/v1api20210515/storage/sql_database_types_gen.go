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
// Storage version of v1api20210515.SqlDatabase
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}
type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabase_Spec   `json:"spec,omitempty"`
	Status            SqlDatabase_STATUS `json:"status,omitempty"`
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
	source, ok := hub.(*storage.SqlDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/SqlDatabase but received %T instead", hub)
	}

	return database.AssignProperties_From_SqlDatabase(source)
}

// ConvertTo populates the provided hub SqlDatabase from our SqlDatabase
func (database *SqlDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.SqlDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/SqlDatabase but received %T instead", hub)
	}

	return database.AssignProperties_To_SqlDatabase(destination)
}

var _ genruntime.KubernetesResource = &SqlDatabase{}

// AzureName returns the Azure name of the resource
func (database *SqlDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database SqlDatabase) GetAPIVersion() string {
	return "2021-05-15"
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

// GetSupportedOperations returns the operations supported by the resource
func (database *SqlDatabase) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
func (database *SqlDatabase) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *SqlDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabase_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (database *SqlDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return database.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (database *SqlDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabase_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabase_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignProperties_From_SqlDatabase populates our SqlDatabase from the provided source SqlDatabase
func (database *SqlDatabase) AssignProperties_From_SqlDatabase(source *storage.SqlDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec SqlDatabase_Spec
	err := spec.AssignProperties_From_SqlDatabase_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_SqlDatabase_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status SqlDatabase_STATUS
	err = status.AssignProperties_From_SqlDatabase_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_SqlDatabase_STATUS() to populate field Status")
	}
	database.Status = status

	// Invoke the augmentConversionForSqlDatabase interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForSqlDatabase); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabase populates the provided destination SqlDatabase from our SqlDatabase
func (database *SqlDatabase) AssignProperties_To_SqlDatabase(destination *storage.SqlDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.SqlDatabase_Spec
	err := database.Spec.AssignProperties_To_SqlDatabase_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_SqlDatabase_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.SqlDatabase_STATUS
	err = database.Status.AssignProperties_To_SqlDatabase_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_SqlDatabase_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForSqlDatabase interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForSqlDatabase); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

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
// Storage version of v1api20210515.SqlDatabase
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

type augmentConversionForSqlDatabase interface {
	AssignPropertiesFrom(src *storage.SqlDatabase) error
	AssignPropertiesTo(dst *storage.SqlDatabase) error
}

// Storage version of v1api20210515.SqlDatabase_Spec
type SqlDatabase_Spec struct {
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

var _ genruntime.ConvertibleSpec = &SqlDatabase_Spec{}

// ConvertSpecFrom populates our SqlDatabase_Spec from the provided source
func (database *SqlDatabase_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.SqlDatabase_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_SqlDatabase_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabase_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_SqlDatabase_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our SqlDatabase_Spec
func (database *SqlDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.SqlDatabase_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_SqlDatabase_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabase_Spec{}
	err := database.AssignProperties_To_SqlDatabase_Spec(dst)
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

// AssignProperties_From_SqlDatabase_Spec populates our SqlDatabase_Spec from the provided source SqlDatabase_Spec
func (database *SqlDatabase_Spec) AssignProperties_From_SqlDatabase_Spec(source *storage.SqlDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	database.AzureName = source.AzureName

	// Location
	database.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignProperties_From_CreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CreateUpdateOptions() to populate field Options")
		}
		database.Options = &option
	} else {
		database.Options = nil
	}

	// OriginalVersion
	database.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseResource
		err := resource.AssignProperties_From_SqlDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlDatabaseResource() to populate field Resource")
		}
		database.Resource = &resource
	} else {
		database.Resource = nil
	}

	// Tags
	database.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlDatabase_Spec interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForSqlDatabase_Spec); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabase_Spec populates the provided destination SqlDatabase_Spec from our SqlDatabase_Spec
func (database *SqlDatabase_Spec) AssignProperties_To_SqlDatabase_Spec(destination *storage.SqlDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// AzureName
	destination.AzureName = database.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(database.Location)

	// Options
	if database.Options != nil {
		var option storage.CreateUpdateOptions
		err := database.Options.AssignProperties_To_CreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if database.Resource != nil {
		var resource storage.SqlDatabaseResource
		err := database.Resource.AssignProperties_To_SqlDatabaseResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlDatabaseResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(database.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlDatabase_Spec interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForSqlDatabase_Spec); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210515.SqlDatabase_STATUS
type SqlDatabase_STATUS struct {
	Conditions  []conditions.Condition                    `json:"conditions,omitempty"`
	Id          *string                                   `json:"id,omitempty"`
	Location    *string                                   `json:"location,omitempty"`
	Name        *string                                   `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                   `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	Resource    *SqlDatabaseGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                         `json:"tags,omitempty"`
	Type        *string                                   `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabase_STATUS{}

// ConvertStatusFrom populates our SqlDatabase_STATUS from the provided source
func (database *SqlDatabase_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.SqlDatabase_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_SqlDatabase_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabase_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_SqlDatabase_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlDatabase_STATUS
func (database *SqlDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.SqlDatabase_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_SqlDatabase_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabase_STATUS{}
	err := database.AssignProperties_To_SqlDatabase_STATUS(dst)
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

// AssignProperties_From_SqlDatabase_STATUS populates our SqlDatabase_STATUS from the provided source SqlDatabase_STATUS
func (database *SqlDatabase_STATUS) AssignProperties_From_SqlDatabase_STATUS(source *storage.SqlDatabase_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	database.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// Options
	if source.Options != nil {
		var option OptionsResource_STATUS
		err := option.AssignProperties_From_OptionsResource_STATUS(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_OptionsResource_STATUS() to populate field Options")
		}
		database.Options = &option
	} else {
		database.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_SqlDatabaseGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlDatabaseGetProperties_Resource_STATUS() to populate field Resource")
		}
		database.Resource = &resource
	} else {
		database.Resource = nil
	}

	// Tags
	database.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlDatabase_STATUS interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForSqlDatabase_STATUS); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabase_STATUS populates the provided destination SqlDatabase_STATUS from our SqlDatabase_STATUS
func (database *SqlDatabase_STATUS) AssignProperties_To_SqlDatabase_STATUS(destination *storage.SqlDatabase_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(database.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// Options
	if database.Options != nil {
		var option storage.OptionsResource_STATUS
		err := database.Options.AssignProperties_To_OptionsResource_STATUS(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_OptionsResource_STATUS() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if database.Resource != nil {
		var resource storage.SqlDatabaseGetProperties_Resource_STATUS
		err := database.Resource.AssignProperties_To_SqlDatabaseGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlDatabaseGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(database.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlDatabase_STATUS interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForSqlDatabase_STATUS); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlDatabase_Spec interface {
	AssignPropertiesFrom(src *storage.SqlDatabase_Spec) error
	AssignPropertiesTo(dst *storage.SqlDatabase_Spec) error
}

type augmentConversionForSqlDatabase_STATUS interface {
	AssignPropertiesFrom(src *storage.SqlDatabase_STATUS) error
	AssignPropertiesTo(dst *storage.SqlDatabase_STATUS) error
}

// Storage version of v1api20210515.SqlDatabaseGetProperties_Resource_STATUS
type SqlDatabaseGetProperties_Resource_STATUS struct {
	Colls       *string                `json:"_colls,omitempty"`
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
	Users       *string                `json:"_users,omitempty"`
}

// AssignProperties_From_SqlDatabaseGetProperties_Resource_STATUS populates our SqlDatabaseGetProperties_Resource_STATUS from the provided source SqlDatabaseGetProperties_Resource_STATUS
func (resource *SqlDatabaseGetProperties_Resource_STATUS) AssignProperties_From_SqlDatabaseGetProperties_Resource_STATUS(source *storage.SqlDatabaseGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Colls
	resource.Colls = genruntime.ClonePointerToString(source.Colls)

	// CreateMode
	if source.CreateMode != nil {
		propertyBag.Add("CreateMode", *source.CreateMode)
	} else {
		propertyBag.Remove("CreateMode")
	}

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// RestoreParameters
	if source.RestoreParameters != nil {
		propertyBag.Add("RestoreParameters", *source.RestoreParameters)
	} else {
		propertyBag.Remove("RestoreParameters")
	}

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

	// Invoke the augmentConversionForSqlDatabaseGetProperties_Resource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlDatabaseGetProperties_Resource_STATUS); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseGetProperties_Resource_STATUS populates the provided destination SqlDatabaseGetProperties_Resource_STATUS from our SqlDatabaseGetProperties_Resource_STATUS
func (resource *SqlDatabaseGetProperties_Resource_STATUS) AssignProperties_To_SqlDatabaseGetProperties_Resource_STATUS(destination *storage.SqlDatabaseGetProperties_Resource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// Colls
	destination.Colls = genruntime.ClonePointerToString(resource.Colls)

	// CreateMode
	if propertyBag.Contains("CreateMode") {
		var createMode string
		err := propertyBag.Pull("CreateMode", &createMode)
		if err != nil {
			return errors.Wrap(err, "pulling 'CreateMode' from propertyBag")
		}

		destination.CreateMode = &createMode
	} else {
		destination.CreateMode = nil
	}

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// RestoreParameters
	if propertyBag.Contains("RestoreParameters") {
		var restoreParameter storage.RestoreParametersBase_STATUS
		err := propertyBag.Pull("RestoreParameters", &restoreParameter)
		if err != nil {
			return errors.Wrap(err, "pulling 'RestoreParameters' from propertyBag")
		}

		destination.RestoreParameters = &restoreParameter
	} else {
		destination.RestoreParameters = nil
	}

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

	// Invoke the augmentConversionForSqlDatabaseGetProperties_Resource_STATUS interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlDatabaseGetProperties_Resource_STATUS); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210515.SqlDatabaseResource
// Cosmos DB SQL database resource object
type SqlDatabaseResource struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_SqlDatabaseResource populates our SqlDatabaseResource from the provided source SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignProperties_From_SqlDatabaseResource(source *storage.SqlDatabaseResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// CreateMode
	if source.CreateMode != nil {
		propertyBag.Add("CreateMode", *source.CreateMode)
	} else {
		propertyBag.Remove("CreateMode")
	}

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// RestoreParameters
	if source.RestoreParameters != nil {
		propertyBag.Add("RestoreParameters", *source.RestoreParameters)
	} else {
		propertyBag.Remove("RestoreParameters")
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlDatabaseResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlDatabaseResource); ok {
		err := augmentedResource.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseResource populates the provided destination SqlDatabaseResource from our SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignProperties_To_SqlDatabaseResource(destination *storage.SqlDatabaseResource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// CreateMode
	if propertyBag.Contains("CreateMode") {
		var createMode string
		err := propertyBag.Pull("CreateMode", &createMode)
		if err != nil {
			return errors.Wrap(err, "pulling 'CreateMode' from propertyBag")
		}

		destination.CreateMode = &createMode
	} else {
		destination.CreateMode = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// RestoreParameters
	if propertyBag.Contains("RestoreParameters") {
		var restoreParameter storage.RestoreParametersBase
		err := propertyBag.Pull("RestoreParameters", &restoreParameter)
		if err != nil {
			return errors.Wrap(err, "pulling 'RestoreParameters' from propertyBag")
		}

		destination.RestoreParameters = &restoreParameter
	} else {
		destination.RestoreParameters = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForSqlDatabaseResource interface (if implemented) to customize the conversion
	var resourceAsAny any = resource
	if augmentedResource, ok := resourceAsAny.(augmentConversionForSqlDatabaseResource); ok {
		err := augmentedResource.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlDatabaseGetProperties_Resource_STATUS interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseGetProperties_Resource_STATUS) error
	AssignPropertiesTo(dst *storage.SqlDatabaseGetProperties_Resource_STATUS) error
}

type augmentConversionForSqlDatabaseResource interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseResource) error
	AssignPropertiesTo(dst *storage.SqlDatabaseResource) error
}

func init() {
	SchemeBuilder.Register(&SqlDatabase{}, &SqlDatabaseList{})
}
