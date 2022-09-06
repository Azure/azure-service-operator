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
// Storage version of v1alpha1api20210515.MongodbDatabase
// Deprecated version of MongodbDatabase. Use v1beta20210515.MongodbDatabase instead
type MongodbDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_MongodbDatabases_Spec `json:"spec,omitempty"`
	Status            MongoDBDatabaseGetResults_STATUS       `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabase{}

// GetConditions returns the conditions of the resource
func (database *MongodbDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *MongodbDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &MongodbDatabase{}

// ConvertFrom populates our MongodbDatabase from the provided hub MongodbDatabase
func (database *MongodbDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.MongodbDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabase but received %T instead", hub)
	}

	return database.AssignProperties_From_MongodbDatabase(source)
}

// ConvertTo populates the provided hub MongodbDatabase from our MongodbDatabase
func (database *MongodbDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.MongodbDatabase)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabase but received %T instead", hub)
	}

	return database.AssignProperties_To_MongodbDatabase(destination)
}

var _ genruntime.KubernetesResource = &MongodbDatabase{}

// AzureName returns the Azure name of the resource
func (database *MongodbDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database MongodbDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (database *MongodbDatabase) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *MongodbDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *MongodbDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
func (database *MongodbDatabase) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *MongodbDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &MongoDBDatabaseGetResults_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *MongodbDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *MongodbDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*MongoDBDatabaseGetResults_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongoDBDatabaseGetResults_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignProperties_From_MongodbDatabase populates our MongodbDatabase from the provided source MongodbDatabase
func (database *MongodbDatabase) AssignProperties_From_MongodbDatabase(source *v20210515s.MongodbDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_MongodbDatabases_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_MongodbDatabases_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_MongodbDatabases_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status MongoDBDatabaseGetResults_STATUS
	err = status.AssignProperties_From_MongoDBDatabaseGetResults_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_MongoDBDatabaseGetResults_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabase populates the provided destination MongodbDatabase from our MongodbDatabase
func (database *MongodbDatabase) AssignProperties_To_MongodbDatabase(destination *v20210515s.MongodbDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccounts_MongodbDatabases_Spec
	err := database.Spec.AssignProperties_To_DatabaseAccounts_MongodbDatabases_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_MongodbDatabases_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.MongoDBDatabaseGetResults_STATUS
	err = database.Status.AssignProperties_To_MongoDBDatabaseGetResults_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_MongoDBDatabaseGetResults_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *MongodbDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "MongodbDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210515.MongodbDatabase
// Deprecated version of MongodbDatabase. Use v1beta20210515.MongodbDatabase instead
type MongodbDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabase `json:"items"`
}

// Storage version of v1alpha1api20210515.DatabaseAccounts_MongodbDatabases_Spec
type DatabaseAccounts_MongodbDatabases_Spec struct {
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
	Resource    *MongoDBDatabaseResource           `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_MongodbDatabases_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_MongodbDatabases_Spec from the provided source
func (databases *DatabaseAccounts_MongodbDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_MongodbDatabases_Spec)
	if ok {
		// Populate our instance from source
		return databases.AssignProperties_From_DatabaseAccounts_MongodbDatabases_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_MongodbDatabases_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databases.AssignProperties_From_DatabaseAccounts_MongodbDatabases_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_MongodbDatabases_Spec
func (databases *DatabaseAccounts_MongodbDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_MongodbDatabases_Spec)
	if ok {
		// Populate destination from our instance
		return databases.AssignProperties_To_DatabaseAccounts_MongodbDatabases_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_MongodbDatabases_Spec{}
	err := databases.AssignProperties_To_DatabaseAccounts_MongodbDatabases_Spec(dst)
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

// AssignProperties_From_DatabaseAccounts_MongodbDatabases_Spec populates our DatabaseAccounts_MongodbDatabases_Spec from the provided source DatabaseAccounts_MongodbDatabases_Spec
func (databases *DatabaseAccounts_MongodbDatabases_Spec) AssignProperties_From_DatabaseAccounts_MongodbDatabases_Spec(source *v20210515s.DatabaseAccounts_MongodbDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	databases.AzureName = source.AzureName

	// Location
	databases.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignProperties_From_CreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CreateUpdateOptions() to populate field Options")
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
		var resource MongoDBDatabaseResource
		err := resource.AssignProperties_From_MongoDBDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_MongoDBDatabaseResource() to populate field Resource")
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

// AssignProperties_To_DatabaseAccounts_MongodbDatabases_Spec populates the provided destination DatabaseAccounts_MongodbDatabases_Spec from our DatabaseAccounts_MongodbDatabases_Spec
func (databases *DatabaseAccounts_MongodbDatabases_Spec) AssignProperties_To_DatabaseAccounts_MongodbDatabases_Spec(destination *v20210515s.DatabaseAccounts_MongodbDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(databases.PropertyBag)

	// AzureName
	destination.AzureName = databases.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(databases.Location)

	// Options
	if databases.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := databases.Options.AssignProperties_To_CreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CreateUpdateOptions() to populate field Options")
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
		var resource v20210515s.MongoDBDatabaseResource
		err := databases.Resource.AssignProperties_To_MongoDBDatabaseResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_MongoDBDatabaseResource() to populate field Resource")
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

// Storage version of v1alpha1api20210515.MongoDBDatabaseGetResults_STATUS
// Deprecated version of MongoDBDatabaseGetResults_STATUS. Use v1beta20210515.MongoDBDatabaseGetResults_STATUS instead
type MongoDBDatabaseGetResults_STATUS struct {
	Conditions  []conditions.Condition                        `json:"conditions,omitempty"`
	Id          *string                                       `json:"id,omitempty"`
	Location    *string                                       `json:"location,omitempty"`
	Name        *string                                       `json:"name,omitempty"`
	Options     *OptionsResource_STATUS                       `json:"options,omitempty"`
	PropertyBag genruntime.PropertyBag                        `json:"$propertyBag,omitempty"`
	Resource    *MongoDBDatabaseGetProperties_STATUS_Resource `json:"resource,omitempty"`
	Tags        map[string]string                             `json:"tags,omitempty"`
	Type        *string                                       `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongoDBDatabaseGetResults_STATUS{}

// ConvertStatusFrom populates our MongoDBDatabaseGetResults_STATUS from the provided source
func (results *MongoDBDatabaseGetResults_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.MongoDBDatabaseGetResults_STATUS)
	if ok {
		// Populate our instance from source
		return results.AssignProperties_From_MongoDBDatabaseGetResults_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.MongoDBDatabaseGetResults_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignProperties_From_MongoDBDatabaseGetResults_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our MongoDBDatabaseGetResults_STATUS
func (results *MongoDBDatabaseGetResults_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.MongoDBDatabaseGetResults_STATUS)
	if ok {
		// Populate destination from our instance
		return results.AssignProperties_To_MongoDBDatabaseGetResults_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.MongoDBDatabaseGetResults_STATUS{}
	err := results.AssignProperties_To_MongoDBDatabaseGetResults_STATUS(dst)
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

// AssignProperties_From_MongoDBDatabaseGetResults_STATUS populates our MongoDBDatabaseGetResults_STATUS from the provided source MongoDBDatabaseGetResults_STATUS
func (results *MongoDBDatabaseGetResults_STATUS) AssignProperties_From_MongoDBDatabaseGetResults_STATUS(source *v20210515s.MongoDBDatabaseGetResults_STATUS) error {
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
		err := option.AssignProperties_From_OptionsResource_STATUS(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_OptionsResource_STATUS() to populate field Options")
		}
		results.Options = &option
	} else {
		results.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource MongoDBDatabaseGetProperties_STATUS_Resource
		err := resource.AssignProperties_From_MongoDBDatabaseGetProperties_STATUS_Resource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_MongoDBDatabaseGetProperties_STATUS_Resource() to populate field Resource")
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

// AssignProperties_To_MongoDBDatabaseGetResults_STATUS populates the provided destination MongoDBDatabaseGetResults_STATUS from our MongoDBDatabaseGetResults_STATUS
func (results *MongoDBDatabaseGetResults_STATUS) AssignProperties_To_MongoDBDatabaseGetResults_STATUS(destination *v20210515s.MongoDBDatabaseGetResults_STATUS) error {
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
		err := results.Options.AssignProperties_To_OptionsResource_STATUS(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_OptionsResource_STATUS() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if results.Resource != nil {
		var resource v20210515s.MongoDBDatabaseGetProperties_STATUS_Resource
		err := results.Resource.AssignProperties_To_MongoDBDatabaseGetProperties_STATUS_Resource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_MongoDBDatabaseGetProperties_STATUS_Resource() to populate field Resource")
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

// Storage version of v1alpha1api20210515.CreateUpdateOptions
// Deprecated version of CreateUpdateOptions. Use v1beta20210515.CreateUpdateOptions instead
type CreateUpdateOptions struct {
	AutoscaleSettings *AutoscaleSettings     `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Throughput        *int                   `json:"throughput,omitempty"`
}

// AssignProperties_From_CreateUpdateOptions populates our CreateUpdateOptions from the provided source CreateUpdateOptions
func (options *CreateUpdateOptions) AssignProperties_From_CreateUpdateOptions(source *v20210515s.CreateUpdateOptions) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoscaleSettings
	if source.AutoscaleSettings != nil {
		var autoscaleSetting AutoscaleSettings
		err := autoscaleSetting.AssignProperties_From_AutoscaleSettings(source.AutoscaleSettings)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_AutoscaleSettings() to populate field AutoscaleSettings")
		}
		options.AutoscaleSettings = &autoscaleSetting
	} else {
		options.AutoscaleSettings = nil
	}

	// Throughput
	options.Throughput = genruntime.ClonePointerToInt(source.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		options.PropertyBag = propertyBag
	} else {
		options.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_CreateUpdateOptions populates the provided destination CreateUpdateOptions from our CreateUpdateOptions
func (options *CreateUpdateOptions) AssignProperties_To_CreateUpdateOptions(destination *v20210515s.CreateUpdateOptions) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(options.PropertyBag)

	// AutoscaleSettings
	if options.AutoscaleSettings != nil {
		var autoscaleSetting v20210515s.AutoscaleSettings
		err := options.AutoscaleSettings.AssignProperties_To_AutoscaleSettings(&autoscaleSetting)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_AutoscaleSettings() to populate field AutoscaleSettings")
		}
		destination.AutoscaleSettings = &autoscaleSetting
	} else {
		destination.AutoscaleSettings = nil
	}

	// Throughput
	destination.Throughput = genruntime.ClonePointerToInt(options.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.MongoDBDatabaseGetProperties_STATUS_Resource
// Deprecated version of MongoDBDatabaseGetProperties_STATUS_Resource. Use v1beta20210515.MongoDBDatabaseGetProperties_STATUS_Resource instead
type MongoDBDatabaseGetProperties_STATUS_Resource struct {
	Etag        *string                `json:"_etag,omitempty"`
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rid         *string                `json:"_rid,omitempty"`
	Ts          *float64               `json:"_ts,omitempty"`
}

// AssignProperties_From_MongoDBDatabaseGetProperties_STATUS_Resource populates our MongoDBDatabaseGetProperties_STATUS_Resource from the provided source MongoDBDatabaseGetProperties_STATUS_Resource
func (resource *MongoDBDatabaseGetProperties_STATUS_Resource) AssignProperties_From_MongoDBDatabaseGetProperties_STATUS_Resource(source *v20210515s.MongoDBDatabaseGetProperties_STATUS_Resource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

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

// AssignProperties_To_MongoDBDatabaseGetProperties_STATUS_Resource populates the provided destination MongoDBDatabaseGetProperties_STATUS_Resource from our MongoDBDatabaseGetProperties_STATUS_Resource
func (resource *MongoDBDatabaseGetProperties_STATUS_Resource) AssignProperties_To_MongoDBDatabaseGetProperties_STATUS_Resource(destination *v20210515s.MongoDBDatabaseGetProperties_STATUS_Resource) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

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

// Storage version of v1alpha1api20210515.MongoDBDatabaseResource
// Deprecated version of MongoDBDatabaseResource. Use v1beta20210515.MongoDBDatabaseResource instead
type MongoDBDatabaseResource struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_MongoDBDatabaseResource populates our MongoDBDatabaseResource from the provided source MongoDBDatabaseResource
func (resource *MongoDBDatabaseResource) AssignProperties_From_MongoDBDatabaseResource(source *v20210515s.MongoDBDatabaseResource) error {
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

// AssignProperties_To_MongoDBDatabaseResource populates the provided destination MongoDBDatabaseResource from our MongoDBDatabaseResource
func (resource *MongoDBDatabaseResource) AssignProperties_To_MongoDBDatabaseResource(destination *v20210515s.MongoDBDatabaseResource) error {
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

// Storage version of v1alpha1api20210515.OptionsResource_STATUS
// Deprecated version of OptionsResource_STATUS. Use v1beta20210515.OptionsResource_STATUS instead
type OptionsResource_STATUS struct {
	AutoscaleSettings *AutoscaleSettings_STATUS `json:"autoscaleSettings,omitempty"`
	PropertyBag       genruntime.PropertyBag    `json:"$propertyBag,omitempty"`
	Throughput        *int                      `json:"throughput,omitempty"`
}

// AssignProperties_From_OptionsResource_STATUS populates our OptionsResource_STATUS from the provided source OptionsResource_STATUS
func (resource *OptionsResource_STATUS) AssignProperties_From_OptionsResource_STATUS(source *v20210515s.OptionsResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AutoscaleSettings
	if source.AutoscaleSettings != nil {
		var autoscaleSetting AutoscaleSettings_STATUS
		err := autoscaleSetting.AssignProperties_From_AutoscaleSettings_STATUS(source.AutoscaleSettings)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_AutoscaleSettings_STATUS() to populate field AutoscaleSettings")
		}
		resource.AutoscaleSettings = &autoscaleSetting
	} else {
		resource.AutoscaleSettings = nil
	}

	// Throughput
	resource.Throughput = genruntime.ClonePointerToInt(source.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		resource.PropertyBag = propertyBag
	} else {
		resource.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_OptionsResource_STATUS populates the provided destination OptionsResource_STATUS from our OptionsResource_STATUS
func (resource *OptionsResource_STATUS) AssignProperties_To_OptionsResource_STATUS(destination *v20210515s.OptionsResource_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(resource.PropertyBag)

	// AutoscaleSettings
	if resource.AutoscaleSettings != nil {
		var autoscaleSetting v20210515s.AutoscaleSettings_STATUS
		err := resource.AutoscaleSettings.AssignProperties_To_AutoscaleSettings_STATUS(&autoscaleSetting)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_AutoscaleSettings_STATUS() to populate field AutoscaleSettings")
		}
		destination.AutoscaleSettings = &autoscaleSetting
	} else {
		destination.AutoscaleSettings = nil
	}

	// Throughput
	destination.Throughput = genruntime.ClonePointerToInt(resource.Throughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.AutoscaleSettings
// Deprecated version of AutoscaleSettings. Use v1beta20210515.AutoscaleSettings instead
type AutoscaleSettings struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_AutoscaleSettings populates our AutoscaleSettings from the provided source AutoscaleSettings
func (settings *AutoscaleSettings) AssignProperties_From_AutoscaleSettings(source *v20210515s.AutoscaleSettings) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// MaxThroughput
	settings.MaxThroughput = genruntime.ClonePointerToInt(source.MaxThroughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_AutoscaleSettings populates the provided destination AutoscaleSettings from our AutoscaleSettings
func (settings *AutoscaleSettings) AssignProperties_To_AutoscaleSettings(destination *v20210515s.AutoscaleSettings) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// MaxThroughput
	destination.MaxThroughput = genruntime.ClonePointerToInt(settings.MaxThroughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210515.AutoscaleSettings_STATUS
// Deprecated version of AutoscaleSettings_STATUS. Use v1beta20210515.AutoscaleSettings_STATUS instead
type AutoscaleSettings_STATUS struct {
	MaxThroughput *int                   `json:"maxThroughput,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// AssignProperties_From_AutoscaleSettings_STATUS populates our AutoscaleSettings_STATUS from the provided source AutoscaleSettings_STATUS
func (settings *AutoscaleSettings_STATUS) AssignProperties_From_AutoscaleSettings_STATUS(source *v20210515s.AutoscaleSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// MaxThroughput
	settings.MaxThroughput = genruntime.ClonePointerToInt(source.MaxThroughput)

	// Update the property bag
	if len(propertyBag) > 0 {
		settings.PropertyBag = propertyBag
	} else {
		settings.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignProperties_To_AutoscaleSettings_STATUS populates the provided destination AutoscaleSettings_STATUS from our AutoscaleSettings_STATUS
func (settings *AutoscaleSettings_STATUS) AssignProperties_To_AutoscaleSettings_STATUS(destination *v20210515s.AutoscaleSettings_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(settings.PropertyBag)

	// MaxThroughput
	destination.MaxThroughput = genruntime.ClonePointerToInt(settings.MaxThroughput)

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
	SchemeBuilder.Register(&MongodbDatabase{}, &MongodbDatabaseList{})
}
