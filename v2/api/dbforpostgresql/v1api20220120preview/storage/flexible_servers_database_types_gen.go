// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	storage "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
// Storage version of v1api20220120preview.FlexibleServersDatabase
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/Databases.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/databases/{databaseName}
type FlexibleServersDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersDatabase_Spec   `json:"spec,omitempty"`
	Status            FlexibleServersDatabase_STATUS `json:"status,omitempty"`
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
	var source storage.FlexibleServersDatabase

	err := source.ConvertFrom(hub)
	if err != nil {
		return eris.Wrap(err, "converting from hub to source")
	}

	err = database.AssignProperties_From_FlexibleServersDatabase(&source)
	if err != nil {
		return eris.Wrap(err, "converting from source to database")
	}

	return nil
}

// ConvertTo populates the provided hub FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.FlexibleServersDatabase
	err := database.AssignProperties_To_FlexibleServersDatabase(&destination)
	if err != nil {
		return eris.Wrap(err, "converting to destination from database")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return eris.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ configmaps.Exporter = &FlexibleServersDatabase{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (database *FlexibleServersDatabase) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if database.Spec.OperatorSpec == nil {
		return nil
	}
	return database.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &FlexibleServersDatabase{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (database *FlexibleServersDatabase) SecretDestinationExpressions() []*core.DestinationExpression {
	if database.Spec.OperatorSpec == nil {
		return nil
	}
	return database.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &FlexibleServersDatabase{}

// AzureName returns the Azure name of the resource
func (database *FlexibleServersDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-20-preview"
func (database FlexibleServersDatabase) GetAPIVersion() string {
	return "2022-01-20-preview"
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

// GetSupportedOperations returns the operations supported by the resource
func (database *FlexibleServersDatabase) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/databases"
func (database *FlexibleServersDatabase) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/databases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *FlexibleServersDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServersDatabase_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (database *FlexibleServersDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return database.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (database *FlexibleServersDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServersDatabase_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersDatabase_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// AssignProperties_From_FlexibleServersDatabase populates our FlexibleServersDatabase from the provided source FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignProperties_From_FlexibleServersDatabase(source *storage.FlexibleServersDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersDatabase_Spec
	err := spec.AssignProperties_From_FlexibleServersDatabase_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersDatabase_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status FlexibleServersDatabase_STATUS
	err = status.AssignProperties_From_FlexibleServersDatabase_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersDatabase_STATUS() to populate field Status")
	}
	database.Status = status

	// Invoke the augmentConversionForFlexibleServersDatabase interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersDatabase populates the provided destination FlexibleServersDatabase from our FlexibleServersDatabase
func (database *FlexibleServersDatabase) AssignProperties_To_FlexibleServersDatabase(destination *storage.FlexibleServersDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.FlexibleServersDatabase_Spec
	err := database.Spec.AssignProperties_To_FlexibleServersDatabase_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersDatabase_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.FlexibleServersDatabase_STATUS
	err = database.Status.AssignProperties_To_FlexibleServersDatabase_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersDatabase_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFlexibleServersDatabase interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
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
// Storage version of v1api20220120preview.FlexibleServersDatabase
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/preview/2022-01-20-preview/Databases.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/databases/{databaseName}
type FlexibleServersDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersDatabase `json:"items"`
}

type augmentConversionForFlexibleServersDatabase interface {
	AssignPropertiesFrom(src *storage.FlexibleServersDatabase) error
	AssignPropertiesTo(dst *storage.FlexibleServersDatabase) error
}

// Storage version of v1api20220120preview.FlexibleServersDatabase_Spec
type FlexibleServersDatabase_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                               `json:"azureName,omitempty"`
	Charset         *string                              `json:"charset,omitempty"`
	Collation       *string                              `json:"collation,omitempty"`
	OperatorSpec    *FlexibleServersDatabaseOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                               `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner       *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersDatabase_Spec{}

// ConvertSpecFrom populates our FlexibleServersDatabase_Spec from the provided source
func (database *FlexibleServersDatabase_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.FlexibleServersDatabase_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_FlexibleServersDatabase_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersDatabase_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_FlexibleServersDatabase_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersDatabase_Spec
func (database *FlexibleServersDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.FlexibleServersDatabase_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_FlexibleServersDatabase_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersDatabase_Spec{}
	err := database.AssignProperties_To_FlexibleServersDatabase_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServersDatabase_Spec populates our FlexibleServersDatabase_Spec from the provided source FlexibleServersDatabase_Spec
func (database *FlexibleServersDatabase_Spec) AssignProperties_From_FlexibleServersDatabase_Spec(source *storage.FlexibleServersDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	database.AzureName = source.AzureName

	// Charset
	database.Charset = genruntime.ClonePointerToString(source.Charset)

	// Collation
	database.Collation = genruntime.ClonePointerToString(source.Collation)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec FlexibleServersDatabaseOperatorSpec
		err := operatorSpec.AssignProperties_From_FlexibleServersDatabaseOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersDatabaseOperatorSpec() to populate field OperatorSpec")
		}
		database.OperatorSpec = &operatorSpec
	} else {
		database.OperatorSpec = nil
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

	// Update the property bag
	if len(propertyBag) > 0 {
		database.PropertyBag = propertyBag
	} else {
		database.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersDatabase_Spec interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase_Spec); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersDatabase_Spec populates the provided destination FlexibleServersDatabase_Spec from our FlexibleServersDatabase_Spec
func (database *FlexibleServersDatabase_Spec) AssignProperties_To_FlexibleServersDatabase_Spec(destination *storage.FlexibleServersDatabase_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(database.PropertyBag)

	// AzureName
	destination.AzureName = database.AzureName

	// Charset
	destination.Charset = genruntime.ClonePointerToString(database.Charset)

	// Collation
	destination.Collation = genruntime.ClonePointerToString(database.Collation)

	// OperatorSpec
	if database.OperatorSpec != nil {
		var operatorSpec storage.FlexibleServersDatabaseOperatorSpec
		err := database.OperatorSpec.AssignProperties_To_FlexibleServersDatabaseOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersDatabaseOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
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

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersDatabase_Spec interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase_Spec); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220120preview.FlexibleServersDatabase_STATUS
type FlexibleServersDatabase_STATUS struct {
	Charset     *string                `json:"charset,omitempty"`
	Collation   *string                `json:"collation,omitempty"`
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServersDatabase_STATUS{}

// ConvertStatusFrom populates our FlexibleServersDatabase_STATUS from the provided source
func (database *FlexibleServersDatabase_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.FlexibleServersDatabase_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignProperties_From_FlexibleServersDatabase_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersDatabase_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignProperties_From_FlexibleServersDatabase_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServersDatabase_STATUS
func (database *FlexibleServersDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.FlexibleServersDatabase_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignProperties_To_FlexibleServersDatabase_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersDatabase_STATUS{}
	err := database.AssignProperties_To_FlexibleServersDatabase_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_FlexibleServersDatabase_STATUS populates our FlexibleServersDatabase_STATUS from the provided source FlexibleServersDatabase_STATUS
func (database *FlexibleServersDatabase_STATUS) AssignProperties_From_FlexibleServersDatabase_STATUS(source *storage.FlexibleServersDatabase_STATUS) error {
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
			return eris.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
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

	// Invoke the augmentConversionForFlexibleServersDatabase_STATUS interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase_STATUS); ok {
		err := augmentedDatabase.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersDatabase_STATUS populates the provided destination FlexibleServersDatabase_STATUS from our FlexibleServersDatabase_STATUS
func (database *FlexibleServersDatabase_STATUS) AssignProperties_To_FlexibleServersDatabase_STATUS(destination *storage.FlexibleServersDatabase_STATUS) error {
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
		var systemDatum storage.SystemData_STATUS
		err := database.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
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

	// Invoke the augmentConversionForFlexibleServersDatabase_STATUS interface (if implemented) to customize the conversion
	var databaseAsAny any = database
	if augmentedDatabase, ok := databaseAsAny.(augmentConversionForFlexibleServersDatabase_STATUS); ok {
		err := augmentedDatabase.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersDatabase_Spec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersDatabase_Spec) error
	AssignPropertiesTo(dst *storage.FlexibleServersDatabase_Spec) error
}

type augmentConversionForFlexibleServersDatabase_STATUS interface {
	AssignPropertiesFrom(src *storage.FlexibleServersDatabase_STATUS) error
	AssignPropertiesTo(dst *storage.FlexibleServersDatabase_STATUS) error
}

// Storage version of v1api20220120preview.FlexibleServersDatabaseOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FlexibleServersDatabaseOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_FlexibleServersDatabaseOperatorSpec populates our FlexibleServersDatabaseOperatorSpec from the provided source FlexibleServersDatabaseOperatorSpec
func (operator *FlexibleServersDatabaseOperatorSpec) AssignProperties_From_FlexibleServersDatabaseOperatorSpec(source *storage.FlexibleServersDatabaseOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersDatabaseOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersDatabaseOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersDatabaseOperatorSpec populates the provided destination FlexibleServersDatabaseOperatorSpec from our FlexibleServersDatabaseOperatorSpec
func (operator *FlexibleServersDatabaseOperatorSpec) AssignProperties_To_FlexibleServersDatabaseOperatorSpec(destination *storage.FlexibleServersDatabaseOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersDatabaseOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersDatabaseOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersDatabaseOperatorSpec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersDatabaseOperatorSpec) error
	AssignPropertiesTo(dst *storage.FlexibleServersDatabaseOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&FlexibleServersDatabase{}, &FlexibleServersDatabaseList{})
}
