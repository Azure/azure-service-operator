// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501storage

import (
	"fmt"
	v20210501s "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1beta20210501storage"
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
// Storage version of v1alpha1api20210501.FlexibleServersDatabase
// Deprecated version of FlexibleServersDatabase. Use v1beta20210501.FlexibleServersDatabase instead
type FlexibleServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersDatabases_Spec `json:"spec,omitempty"`
	Status            Database_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersDatabase{}

// GetConditions returns the conditions of the resource
func (database *FlexibleServersDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *FlexibleServersDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersDatabase{}

// ConvertFrom populates our FlexibleServersDatabase from the provided hub FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210501s.FlexibleServersDatabase)
	if !ok {
		return fmt.Errorf("expected dbformysql/v1beta20210501storage/FlexibleServersDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesFromFlexibleServersDatabase(source)
}

// ConvertTo populates the provided hub FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210501s.FlexibleServersDatabase)
	if !ok {
		return fmt.Errorf("expected dbformysql/v1beta20210501storage/FlexibleServersDatabase but received %T instead", hub)
	}

	return database.AssignPropertiesToFlexibleServersDatabase(destination)
}

var _ genruntime.KubernetesResource = &FlexibleServersDatabase{}

// AzureName returns the Azure name of the resource
func (database *FlexibleServersDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (database FlexibleServersDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (database *FlexibleServersDatabase) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (database *FlexibleServersDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *FlexibleServersDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/databases"
func (database *FlexibleServersDatabase) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *FlexibleServersDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Database_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *FlexibleServersDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *FlexibleServersDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Database_Status); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st Database_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignPropertiesFromFlexibleServersDatabase populates our FlexibleServersDatabase from the provided source FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignPropertiesFromFlexibleServersDatabase(source *v20210501s.FlexibleServersDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersDatabases_Spec
	err := spec.AssignPropertiesFromFlexibleServersDatabasesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersDatabasesSpec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status Database_Status
	err = status.AssignPropertiesFromDatabaseStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseStatus() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersDatabase populates the provided destination FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignPropertiesToFlexibleServersDatabase(destination *v20210501s.FlexibleServersDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210501s.FlexibleServersDatabases_Spec
	err := database.Spec.AssignPropertiesToFlexibleServersDatabasesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersDatabasesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210501s.Database_Status
	err = database.Status.AssignPropertiesToDatabaseStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *FlexibleServersDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion,
		Kind:    "FlexibleServersDatabase",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20210501.FlexibleServersDatabase
// Deprecated version of FlexibleServersDatabase. Use v1beta20210501.FlexibleServersDatabase instead
type FlexibleServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersDatabase `json:"items"`
}

// Storage version of v1alpha1api20210501.Database_Status
// Deprecated version of Database_Status. Use v1beta20210501.Database_Status instead
type Database_Status struct {
	Charset     *string                `json:"charset,omitempty"`
	Collation   *string                `json:"collation,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_Status     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Database_Status{}

// ConvertStatusFrom populates our Database_Status from the provided source
func (database *Database_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210501s.Database_Status)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210501s.Database_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Database_Status
func (database *Database_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210501s.Database_Status)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210501s.Database_Status{}
	err := database.AssignPropertiesToDatabaseStatus(dst)
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

// AssignPropertiesFromDatabaseStatus populates our Database_Status from the provided source Database_Status
func (database *Database_Status) AssignPropertiesFromDatabaseStatus(source *v20210501s.Database_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		database.SystemData = &systemDatum
	} else {
		database.SystemData = nil
	}

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToDatabaseStatus populates the provided destination Database_Status from our Database_Status
func (database *Database_Status) AssignPropertiesToDatabaseStatus(destination *v20210501s.Database_Status) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// SystemData
	if database.SystemData != nil {
		var systemDatum v20210501s.SystemData_Status
		err := database.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Storage version of v1alpha1api20210501.FlexibleServersDatabases_Spec
type FlexibleServersDatabases_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Charset         *string `json:"charset,omitempty"`
	Collation       *string `json:"collation,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformysql.azure.com/FlexibleServer resource
	Owner       *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersDatabases_Spec{}

// ConvertSpecFrom populates our FlexibleServersDatabases_Spec from the provided source
func (databases *FlexibleServersDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210501s.FlexibleServersDatabases_Spec)
	if ok {
		// Populate our instance from source
		return databases.AssignPropertiesFromFlexibleServersDatabasesSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210501s.FlexibleServersDatabases_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databases.AssignPropertiesFromFlexibleServersDatabasesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersDatabases_Spec
func (databases *FlexibleServersDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210501s.FlexibleServersDatabases_Spec)
	if ok {
		// Populate destination from our instance
		return databases.AssignPropertiesToFlexibleServersDatabasesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210501s.FlexibleServersDatabases_Spec{}
	err := databases.AssignPropertiesToFlexibleServersDatabasesSpec(dst)
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

// AssignPropertiesFromFlexibleServersDatabasesSpec populates our FlexibleServersDatabases_Spec from the provided source FlexibleServersDatabases_Spec
func (databases *FlexibleServersDatabases_Spec) AssignPropertiesFromFlexibleServersDatabasesSpec(source *v20210501s.FlexibleServersDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	databases.AzureName = source.AzureName

	// Charset
	databases.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	databases.Collation = genruntime.ClonePointerToString(source.Collation)

	// Location
	databases.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	databases.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		databases.Owner = &owner
	} else {
		databases.Owner = nil
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

// AssignPropertiesToFlexibleServersDatabasesSpec populates the provided destination FlexibleServersDatabases_Spec from our FlexibleServersDatabases_Spec
func (databases *FlexibleServersDatabases_Spec) AssignPropertiesToFlexibleServersDatabasesSpec(destination *v20210501s.FlexibleServersDatabases_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(databases.PropertyBag)

	// AzureName
	destination.AzureName = databases.AzureName

	// Charset
	destination.Charset = genruntime.ClonePointerToString(databases.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(databases.Collation)

	// Location
	destination.Location = genruntime.ClonePointerToString(databases.Location)

	// OriginalVersion
	destination.OriginalVersion = databases.OriginalVersion

	// Owner
	if databases.Owner != nil {
		owner := databases.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
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

func init() {
	SchemeBuilder.Register(&FlexibleServersDatabase{}, &FlexibleServersDatabaseList{})
}
