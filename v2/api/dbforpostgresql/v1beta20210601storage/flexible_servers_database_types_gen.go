// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601storage

import (
	v1beta20220120ps "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20220120previewstorage"
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
// Storage version of v1beta20210601.FlexibleServersDatabase
// Deprecated version of FlexibleServersDatabase. Use v1api20210601.FlexibleServersDatabase instead
type FlexibleServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_Database_Spec   `json:"spec,omitempty"`
	Status            FlexibleServers_Database_STATUS `json:"status,omitempty"`
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
	// intermediate variable for conversion
	var source v1beta20220120ps.FlexibleServersDatabase

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = database.AssignProperties_From_FlexibleServersDatabase(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to database")
	}

	return nil
}

// ConvertTo populates the provided hub FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v1beta20220120ps.FlexibleServersDatabase
	err := database.AssignProperties_To_FlexibleServersDatabase(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from database")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ genruntime.KubernetesResource = &FlexibleServersDatabase{}

// AzureName returns the Azure name of the resource
func (database *FlexibleServersDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/databases"
func (database *FlexibleServersDatabase) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *FlexibleServersDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServers_Database_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (database *FlexibleServersDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return database.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (database *FlexibleServersDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServers_Database_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServers_Database_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignProperties_From_FlexibleServersDatabase populates our FlexibleServersDatabase from the provided source FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignProperties_From_FlexibleServersDatabase(source *v1beta20220120ps.FlexibleServersDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServers_Database_Spec
	err := spec.AssignProperties_From_FlexibleServers_Database_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Database_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status FlexibleServers_Database_STATUS
	err = status.AssignProperties_From_FlexibleServers_Database_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_FlexibleServers_Database_STATUS() to populate field Status")
	}
	database.Status = status

	// Invoke the augmentConversionForFlexibleServersDatabase interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersDatabase populates the provided destination FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignProperties_To_FlexibleServersDatabase(destination *v1beta20220120ps.FlexibleServersDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec v1beta20220120ps.FlexibleServers_Database_Spec
	err := database.Spec.AssignProperties_To_FlexibleServers_Database_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Database_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1beta20220120ps.FlexibleServers_Database_STATUS
	err = database.Status.AssignProperties_To_FlexibleServers_Database_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_FlexibleServers_Database_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFlexibleServersDatabase interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

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
// Storage version of v1beta20210601.FlexibleServersDatabase
// Deprecated version of FlexibleServersDatabase. Use v1api20210601.FlexibleServersDatabase instead
type FlexibleServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersDatabase `json:"items"`
}

type augmentConversionForFlexibleServersDatabase interface {
	AssignPropertiesFrom(src *v1beta20220120ps.FlexibleServersDatabase) error
	AssignPropertiesTo(dst *v1beta20220120ps.FlexibleServersDatabase) error
}

// Storage version of v1beta20210601.FlexibleServers_Database_Spec
type FlexibleServers_Database_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Charset         *string `json:"charset,omitempty"`
	Collation       *string `json:"collation,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner       *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_Database_Spec{}

// ConvertSpecFrom populates our FlexibleServers_Database_Spec from the provided source
func (database *FlexibleServers_Database_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1beta20220120ps.FlexibleServers_Database_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_FlexibleServers_Database_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1beta20220120ps.FlexibleServers_Database_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_FlexibleServers_Database_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_Database_Spec
func (database *FlexibleServers_Database_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1beta20220120ps.FlexibleServers_Database_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_FlexibleServers_Database_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20220120ps.FlexibleServers_Database_Spec{}
	err := database.AssignProperties_To_FlexibleServers_Database_Spec(dst)
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

// AssignProperties_From_FlexibleServers_Database_Spec populates our FlexibleServers_Database_Spec from the provided source FlexibleServers_Database_Spec
func (database *FlexibleServers_Database_Spec) AssignProperties_From_FlexibleServers_Database_Spec(source *v1beta20220120ps.FlexibleServers_Database_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	database.AzureName = source.AzureName

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// OriginalVersion
	database.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServers_Database_Spec interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServers_Database_Spec); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Database_Spec populates the provided destination FlexibleServers_Database_Spec from our FlexibleServers_Database_Spec
func (database *FlexibleServers_Database_Spec) AssignProperties_To_FlexibleServers_Database_Spec(destination *v1beta20220120ps.FlexibleServers_Database_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// AzureName
	destination.AzureName = database.AzureName

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServers_Database_Spec interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServers_Database_Spec); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20210601.FlexibleServers_Database_STATUS
// Deprecated version of FlexibleServers_Database_STATUS. Use v1api20210601.FlexibleServers_Database_STATUS instead
type FlexibleServers_Database_STATUS struct {
	Charset     *string                `json:"charset,omitempty"`
	Collation   *string                `json:"collation,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServers_Database_STATUS{}

// ConvertStatusFrom populates our FlexibleServers_Database_STATUS from the provided source
func (database *FlexibleServers_Database_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1beta20220120ps.FlexibleServers_Database_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_FlexibleServers_Database_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1beta20220120ps.FlexibleServers_Database_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_FlexibleServers_Database_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServers_Database_STATUS
func (database *FlexibleServers_Database_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1beta20220120ps.FlexibleServers_Database_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_FlexibleServers_Database_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20220120ps.FlexibleServers_Database_STATUS{}
	err := database.AssignProperties_To_FlexibleServers_Database_STATUS(dst)
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

// AssignProperties_From_FlexibleServers_Database_STATUS populates our FlexibleServers_Database_STATUS from the provided source FlexibleServers_Database_STATUS
func (database *FlexibleServers_Database_STATUS) AssignProperties_From_FlexibleServers_Database_STATUS(source *v1beta20220120ps.FlexibleServers_Database_STATUS) error {
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
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
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

	// Invoke the augmentConversionForFlexibleServers_Database_STATUS interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServers_Database_STATUS); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServers_Database_STATUS populates the provided destination FlexibleServers_Database_STATUS from our FlexibleServers_Database_STATUS
func (database *FlexibleServers_Database_STATUS) AssignProperties_To_FlexibleServers_Database_STATUS(destination *v1beta20220120ps.FlexibleServers_Database_STATUS) error {
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
		var systemDatum v1beta20220120ps.SystemData_STATUS
		err := database.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
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

	// Invoke the augmentConversionForFlexibleServers_Database_STATUS interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServers_Database_STATUS); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServers_Database_Spec interface {
	AssignPropertiesFrom(src *v1beta20220120ps.FlexibleServers_Database_Spec) error
	AssignPropertiesTo(dst *v1beta20220120ps.FlexibleServers_Database_Spec) error
}

type augmentConversionForFlexibleServers_Database_STATUS interface {
	AssignPropertiesFrom(src *v1beta20220120ps.FlexibleServers_Database_STATUS) error
	AssignPropertiesTo(dst *v1beta20220120ps.FlexibleServers_Database_STATUS) error
}

func init() {
	SchemeBuilder.Register(&FlexibleServersDatabase{}, &FlexibleServersDatabaseList{})
}
