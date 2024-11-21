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
// Storage version of v1api20210515.SqlDatabaseContainerThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseContainerThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            SqlDatabaseContainerThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseContainerThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerThroughputSetting{}

// ConvertFrom populates our SqlDatabaseContainerThroughputSetting from the provided hub SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.SqlDatabaseContainerThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/SqlDatabaseContainerThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerThroughputSetting from our SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.SqlDatabaseContainerThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1api20231115/storage/SqlDatabaseContainerThroughputSetting but received %T instead", hub)
	}

	return setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting(destination)
}

var _ configmaps.Exporter = &SqlDatabaseContainerThroughputSetting{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (setting *SqlDatabaseContainerThroughputSetting) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SqlDatabaseContainerThroughputSetting{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (setting *SqlDatabaseContainerThroughputSetting) SecretDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseContainerThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseContainerThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *SqlDatabaseContainerThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (setting *SqlDatabaseContainerThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseContainerThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseContainerThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *SqlDatabaseContainerThroughputSetting) Owner() *genruntime.ResourceReference {
	if setting.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseContainerThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseContainerThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseContainerThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// AssignProperties_From_SqlDatabaseContainerThroughputSetting populates our SqlDatabaseContainerThroughputSetting from the provided source SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) AssignProperties_From_SqlDatabaseContainerThroughputSetting(source *storage.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec SqlDatabaseContainerThroughputSetting_Spec
	err := spec.AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status SqlDatabaseContainerThroughputSetting_STATUS
	err = status.AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseContainerThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerThroughputSetting populates the provided destination SqlDatabaseContainerThroughputSetting from our SqlDatabaseContainerThroughputSetting
func (setting *SqlDatabaseContainerThroughputSetting) AssignProperties_To_SqlDatabaseContainerThroughputSetting(destination *storage.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.SqlDatabaseContainerThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.SqlDatabaseContainerThroughputSetting_STATUS
	err = setting.Status.AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseContainerThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseContainerThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "SqlDatabaseContainerThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210515.SqlDatabaseContainerThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/containers/{containerName}/throughputSettings/default
type SqlDatabaseContainerThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerThroughputSetting `json:"items"`
}

type augmentConversionForSqlDatabaseContainerThroughputSetting interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseContainerThroughputSetting) error
	AssignPropertiesTo(dst *storage.SqlDatabaseContainerThroughputSetting) error
}

// Storage version of v1api20210515.SqlDatabaseContainerThroughputSetting_Spec
type SqlDatabaseContainerThroughputSetting_Spec struct {
	Location        *string                                            `json:"location,omitempty"`
	OperatorSpec    *SqlDatabaseContainerThroughputSettingOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                             `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SqlDatabaseContainerThroughputSetting_Spec{}

// ConvertSpecFrom populates our SqlDatabaseContainerThroughputSetting_Spec from the provided source
func (setting *SqlDatabaseContainerThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.SqlDatabaseContainerThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabaseContainerThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our SqlDatabaseContainerThroughputSetting_Spec
func (setting *SqlDatabaseContainerThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.SqlDatabaseContainerThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabaseContainerThroughputSetting_Spec{}
	err := setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(dst)
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

// AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec populates our SqlDatabaseContainerThroughputSetting_Spec from the provided source SqlDatabaseContainerThroughputSetting_Spec
func (setting *SqlDatabaseContainerThroughputSetting_Spec) AssignProperties_From_SqlDatabaseContainerThroughputSetting_Spec(source *storage.SqlDatabaseContainerThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec SqlDatabaseContainerThroughputSettingOperatorSpec
		err := operatorSpec.AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec() to populate field OperatorSpec")
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

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseContainerThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec populates the provided destination SqlDatabaseContainerThroughputSetting_Spec from our SqlDatabaseContainerThroughputSetting_Spec
func (setting *SqlDatabaseContainerThroughputSetting_Spec) AssignProperties_To_SqlDatabaseContainerThroughputSetting_Spec(destination *storage.SqlDatabaseContainerThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(setting.PropertyBag)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OperatorSpec
	if setting.OperatorSpec != nil {
		var operatorSpec storage.SqlDatabaseContainerThroughputSettingOperatorSpec
		err := setting.OperatorSpec.AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec() to populate field OperatorSpec")
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

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseContainerThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210515.SqlDatabaseContainerThroughputSetting_STATUS
type SqlDatabaseContainerThroughputSetting_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseContainerThroughputSetting_STATUS{}

// ConvertStatusFrom populates our SqlDatabaseContainerThroughputSetting_STATUS from the provided source
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.SqlDatabaseContainerThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabaseContainerThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseContainerThroughputSetting_STATUS
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.SqlDatabaseContainerThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabaseContainerThroughputSetting_STATUS{}
	err := setting.AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(dst)
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

// AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS populates our SqlDatabaseContainerThroughputSetting_STATUS from the provided source SqlDatabaseContainerThroughputSetting_STATUS
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) AssignProperties_From_SqlDatabaseContainerThroughputSetting_STATUS(source *storage.SqlDatabaseContainerThroughputSetting_STATUS) error {
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

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseContainerThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS populates the provided destination SqlDatabaseContainerThroughputSetting_STATUS from our SqlDatabaseContainerThroughputSetting_STATUS
func (setting *SqlDatabaseContainerThroughputSetting_STATUS) AssignProperties_To_SqlDatabaseContainerThroughputSetting_STATUS(destination *storage.SqlDatabaseContainerThroughputSetting_STATUS) error {
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

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseContainerThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlDatabaseContainerThroughputSetting_Spec interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseContainerThroughputSetting_Spec) error
	AssignPropertiesTo(dst *storage.SqlDatabaseContainerThroughputSetting_Spec) error
}

type augmentConversionForSqlDatabaseContainerThroughputSetting_STATUS interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseContainerThroughputSetting_STATUS) error
	AssignPropertiesTo(dst *storage.SqlDatabaseContainerThroughputSetting_STATUS) error
}

// Storage version of v1api20210515.SqlDatabaseContainerThroughputSettingOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SqlDatabaseContainerThroughputSettingOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec populates our SqlDatabaseContainerThroughputSettingOperatorSpec from the provided source SqlDatabaseContainerThroughputSettingOperatorSpec
func (operator *SqlDatabaseContainerThroughputSettingOperatorSpec) AssignProperties_From_SqlDatabaseContainerThroughputSettingOperatorSpec(source *storage.SqlDatabaseContainerThroughputSettingOperatorSpec) error {
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

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSettingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForSqlDatabaseContainerThroughputSettingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec populates the provided destination SqlDatabaseContainerThroughputSettingOperatorSpec from our SqlDatabaseContainerThroughputSettingOperatorSpec
func (operator *SqlDatabaseContainerThroughputSettingOperatorSpec) AssignProperties_To_SqlDatabaseContainerThroughputSettingOperatorSpec(destination *storage.SqlDatabaseContainerThroughputSettingOperatorSpec) error {
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

	// Invoke the augmentConversionForSqlDatabaseContainerThroughputSettingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForSqlDatabaseContainerThroughputSettingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlDatabaseContainerThroughputSettingOperatorSpec interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseContainerThroughputSettingOperatorSpec) error
	AssignPropertiesTo(dst *storage.SqlDatabaseContainerThroughputSettingOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerThroughputSetting{}, &SqlDatabaseContainerThroughputSettingList{})
}
