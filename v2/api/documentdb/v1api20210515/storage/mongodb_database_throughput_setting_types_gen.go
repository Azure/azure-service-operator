// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115/storage"
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
// Storage version of v1api20210515.MongodbDatabaseThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongodbDatabaseThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            MongodbDatabaseThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &MongodbDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *MongodbDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *MongodbDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &MongodbDatabaseThroughputSetting{}

// ConvertFrom populates our MongodbDatabaseThroughputSetting from the provided hub MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_MongodbDatabaseThroughputSetting(source)
}

// ConvertTo populates the provided hub MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_MongodbDatabaseThroughputSetting(destination)
}

var _ configmaps.Exporter = &MongodbDatabaseThroughputSetting{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (setting *MongodbDatabaseThroughputSetting) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &MongodbDatabaseThroughputSetting{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (setting *MongodbDatabaseThroughputSetting) SecretDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &MongodbDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *MongodbDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting MongodbDatabaseThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceScope returns the scope of the resource
func (setting *MongodbDatabaseThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *MongodbDatabaseThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
func (setting *MongodbDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &MongodbDatabaseThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*MongodbDatabaseThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongodbDatabaseThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// AssignProperties_From_MongodbDatabaseThroughputSetting populates our MongodbDatabaseThroughputSetting from the provided source MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_From_MongodbDatabaseThroughputSetting(source *storage.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec MongodbDatabaseThroughputSetting_Spec
	err := spec.AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_MongodbDatabaseThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status MongodbDatabaseThroughputSetting_STATUS
	err = status.AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// Invoke the augmentConversionForMongodbDatabaseThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForMongodbDatabaseThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting populates the provided destination MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_To_MongodbDatabaseThroughputSetting(destination *storage.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.MongodbDatabaseThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_MongodbDatabaseThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.MongodbDatabaseThroughputSetting_STATUS
	err = setting.Status.AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForMongodbDatabaseThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForMongodbDatabaseThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *MongodbDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "MongodbDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210515.MongodbDatabaseThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

type augmentConversionForMongodbDatabaseThroughputSetting interface {
	AssignPropertiesFrom(src *storage.MongodbDatabaseThroughputSetting) error
	AssignPropertiesTo(dst *storage.MongodbDatabaseThroughputSetting) error
}

// Storage version of v1api20210515.MongodbDatabaseThroughputSetting_Spec
type MongodbDatabaseThroughputSetting_Spec struct {
	Location        *string                                       `json:"location,omitempty"`
	OperatorSpec    *MongodbDatabaseThroughputSettingOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/MongodbDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"MongodbDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &MongodbDatabaseThroughputSetting_Spec{}

// ConvertSpecFrom populates our MongodbDatabaseThroughputSetting_Spec from the provided source
func (setting *MongodbDatabaseThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.MongodbDatabaseThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.MongodbDatabaseThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our MongodbDatabaseThroughputSetting_Spec
func (setting *MongodbDatabaseThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.MongodbDatabaseThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.MongodbDatabaseThroughputSetting_Spec{}
	err := setting.AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(dst)
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

// AssignProperties_From_MongodbDatabaseThroughputSetting_Spec populates our MongodbDatabaseThroughputSetting_Spec from the provided source MongodbDatabaseThroughputSetting_Spec
func (setting *MongodbDatabaseThroughputSetting_Spec) AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(source *storage.MongodbDatabaseThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec MongodbDatabaseThroughputSettingOperatorSpec
		err := operatorSpec.AssignProperties_From_MongodbDatabaseThroughputSettingOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_MongodbDatabaseThroughputSettingOperatorSpec() to populate field OperatorSpec")
		}
		setting.OperatorSpec = &operatorSpec
	} else {
		setting.OperatorSpec = nil
	}

	// OriginalVersion
	setting.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignProperties_From_ThroughputSettingsResource(source.Resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ThroughputSettingsResource() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		setting.PropertyBag = propertyBag
	} else {
		setting.PropertyBag = nil
	}

	// Invoke the augmentConversionForMongodbDatabaseThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForMongodbDatabaseThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting_Spec populates the provided destination MongodbDatabaseThroughputSetting_Spec from our MongodbDatabaseThroughputSetting_Spec
func (setting *MongodbDatabaseThroughputSetting_Spec) AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(destination *storage.MongodbDatabaseThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(setting.PropertyBag)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OperatorSpec
	if setting.OperatorSpec != nil {
		var operatorSpec storage.MongodbDatabaseThroughputSettingOperatorSpec
		err := setting.OperatorSpec.AssignProperties_To_MongodbDatabaseThroughputSettingOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_MongodbDatabaseThroughputSettingOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if setting.Resource != nil {
		var resource storage.ThroughputSettingsResource
		err := setting.Resource.AssignProperties_To_ThroughputSettingsResource(&resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForMongodbDatabaseThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForMongodbDatabaseThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210515.MongodbDatabaseThroughputSetting_STATUS
type MongodbDatabaseThroughputSetting_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongodbDatabaseThroughputSetting_STATUS{}

// ConvertStatusFrom populates our MongodbDatabaseThroughputSetting_STATUS from the provided source
func (setting *MongodbDatabaseThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.MongodbDatabaseThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.MongodbDatabaseThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our MongodbDatabaseThroughputSetting_STATUS
func (setting *MongodbDatabaseThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.MongodbDatabaseThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.MongodbDatabaseThroughputSetting_STATUS{}
	err := setting.AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(dst)
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

// AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS populates our MongodbDatabaseThroughputSetting_STATUS from the provided source MongodbDatabaseThroughputSetting_STATUS
func (setting *MongodbDatabaseThroughputSetting_STATUS) AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(source *storage.MongodbDatabaseThroughputSetting_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		setting.PropertyBag = propertyBag
	} else {
		setting.PropertyBag = nil
	}

	// Invoke the augmentConversionForMongodbDatabaseThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForMongodbDatabaseThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS populates the provided destination MongodbDatabaseThroughputSetting_STATUS from our MongodbDatabaseThroughputSetting_STATUS
func (setting *MongodbDatabaseThroughputSetting_STATUS) AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(destination *storage.MongodbDatabaseThroughputSetting_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(setting.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// Resource
	if setting.Resource != nil {
		var resource storage.ThroughputSettingsGetProperties_Resource_STATUS
		err := setting.Resource.AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForMongodbDatabaseThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForMongodbDatabaseThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForMongodbDatabaseThroughputSetting_Spec interface {
	AssignPropertiesFrom(src *storage.MongodbDatabaseThroughputSetting_Spec) error
	AssignPropertiesTo(dst *storage.MongodbDatabaseThroughputSetting_Spec) error
}

type augmentConversionForMongodbDatabaseThroughputSetting_STATUS interface {
	AssignPropertiesFrom(src *storage.MongodbDatabaseThroughputSetting_STATUS) error
	AssignPropertiesTo(dst *storage.MongodbDatabaseThroughputSetting_STATUS) error
}

// Storage version of v1api20210515.MongodbDatabaseThroughputSettingOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type MongodbDatabaseThroughputSettingOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_MongodbDatabaseThroughputSettingOperatorSpec populates our MongodbDatabaseThroughputSettingOperatorSpec from the provided source MongodbDatabaseThroughputSettingOperatorSpec
func (operator *MongodbDatabaseThroughputSettingOperatorSpec) AssignProperties_From_MongodbDatabaseThroughputSettingOperatorSpec(source *storage.MongodbDatabaseThroughputSettingOperatorSpec) error {
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

	// Invoke the augmentConversionForMongodbDatabaseThroughputSettingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForMongodbDatabaseThroughputSettingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSettingOperatorSpec populates the provided destination MongodbDatabaseThroughputSettingOperatorSpec from our MongodbDatabaseThroughputSettingOperatorSpec
func (operator *MongodbDatabaseThroughputSettingOperatorSpec) AssignProperties_To_MongodbDatabaseThroughputSettingOperatorSpec(destination *storage.MongodbDatabaseThroughputSettingOperatorSpec) error {
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

	// Invoke the augmentConversionForMongodbDatabaseThroughputSettingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForMongodbDatabaseThroughputSettingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForMongodbDatabaseThroughputSettingOperatorSpec interface {
	AssignPropertiesFrom(src *storage.MongodbDatabaseThroughputSettingOperatorSpec) error
	AssignPropertiesTo(dst *storage.MongodbDatabaseThroughputSettingOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}
