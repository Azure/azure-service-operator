// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
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
// Storage version of v1api20210515.SqlDatabaseThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/throughputSettings/default
type SqlDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlDatabaseThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            SqlDatabaseThroughputSetting_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (setting *SqlDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return setting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (setting *SqlDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	setting.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseThroughputSetting{}

// ConvertFrom populates our SqlDatabaseThroughputSetting from the provided hub SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.SqlDatabaseThroughputSetting

	err := source.ConvertFrom(hub)
	if err != nil {
		return eris.Wrap(err, "converting from hub to source")
	}

	err = setting.AssignProperties_From_SqlDatabaseThroughputSetting(&source)
	if err != nil {
		return eris.Wrap(err, "converting from source to setting")
	}

	return nil
}

// ConvertTo populates the provided hub SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.SqlDatabaseThroughputSetting
	err := setting.AssignProperties_To_SqlDatabaseThroughputSetting(&destination)
	if err != nil {
		return eris.Wrap(err, "converting to destination from setting")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return eris.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ configmaps.Exporter = &SqlDatabaseThroughputSetting{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (setting *SqlDatabaseThroughputSetting) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &SqlDatabaseThroughputSetting{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (setting *SqlDatabaseThroughputSetting) SecretDestinationExpressions() []*core.DestinationExpression {
	if setting.Spec.OperatorSpec == nil {
		return nil
	}
	return setting.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &SqlDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *SqlDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting SqlDatabaseThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceScope returns the scope of the resource
func (setting *SqlDatabaseThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *SqlDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *SqlDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *SqlDatabaseThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (setting *SqlDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *SqlDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *SqlDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	if setting.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *SqlDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// AssignProperties_From_SqlDatabaseThroughputSetting populates our SqlDatabaseThroughputSetting from the provided source SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignProperties_From_SqlDatabaseThroughputSetting(source *storage.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec SqlDatabaseThroughputSetting_Spec
	err := spec.AssignProperties_From_SqlDatabaseThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status SqlDatabaseThroughputSetting_STATUS
	err = status.AssignProperties_From_SqlDatabaseThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseThroughputSetting populates the provided destination SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (setting *SqlDatabaseThroughputSetting) AssignProperties_To_SqlDatabaseThroughputSetting(destination *storage.SqlDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.SqlDatabaseThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_SqlDatabaseThroughputSetting_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.SqlDatabaseThroughputSetting_STATUS
	err = setting.Status.AssignProperties_To_SqlDatabaseThroughputSetting_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseThroughputSetting_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *SqlDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion,
		Kind:    "SqlDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210515.SqlDatabaseThroughputSetting
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/sqlDatabases/{databaseName}/throughputSettings/default
type SqlDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseThroughputSetting `json:"items"`
}

type augmentConversionForSqlDatabaseThroughputSetting interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseThroughputSetting) error
	AssignPropertiesTo(dst *storage.SqlDatabaseThroughputSetting) error
}

// Storage version of v1api20210515.SqlDatabaseThroughputSetting_Spec
type SqlDatabaseThroughputSetting_Spec struct {
	Location        *string                                   `json:"location,omitempty"`
	OperatorSpec    *SqlDatabaseThroughputSettingOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                    `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabase resource
	Owner       *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabase"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource        `json:"resource,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &SqlDatabaseThroughputSetting_Spec{}

// ConvertSpecFrom populates our SqlDatabaseThroughputSetting_Spec from the provided source
func (setting *SqlDatabaseThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.SqlDatabaseThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_SqlDatabaseThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabaseThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_SqlDatabaseThroughputSetting_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our SqlDatabaseThroughputSetting_Spec
func (setting *SqlDatabaseThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.SqlDatabaseThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_SqlDatabaseThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabaseThroughputSetting_Spec{}
	err := setting.AssignProperties_To_SqlDatabaseThroughputSetting_Spec(dst)
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

// AssignProperties_From_SqlDatabaseThroughputSetting_Spec populates our SqlDatabaseThroughputSetting_Spec from the provided source SqlDatabaseThroughputSetting_Spec
func (setting *SqlDatabaseThroughputSetting_Spec) AssignProperties_From_SqlDatabaseThroughputSetting_Spec(source *storage.SqlDatabaseThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec SqlDatabaseThroughputSettingOperatorSpec
		err := operatorSpec.AssignProperties_From_SqlDatabaseThroughputSettingOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SqlDatabaseThroughputSettingOperatorSpec() to populate field OperatorSpec")
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

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseThroughputSetting_Spec populates the provided destination SqlDatabaseThroughputSetting_Spec from our SqlDatabaseThroughputSetting_Spec
func (setting *SqlDatabaseThroughputSetting_Spec) AssignProperties_To_SqlDatabaseThroughputSetting_Spec(destination *storage.SqlDatabaseThroughputSetting_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(setting.PropertyBag)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OperatorSpec
	if setting.OperatorSpec != nil {
		var operatorSpec storage.SqlDatabaseThroughputSettingOperatorSpec
		err := setting.OperatorSpec.AssignProperties_To_SqlDatabaseThroughputSettingOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SqlDatabaseThroughputSettingOperatorSpec() to populate field OperatorSpec")
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

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting_Spec interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting_Spec); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210515.SqlDatabaseThroughputSetting_STATUS
type SqlDatabaseThroughputSetting_STATUS struct {
	Conditions  []conditions.Condition                           `json:"conditions,omitempty"`
	Id          *string                                          `json:"id,omitempty"`
	Location    *string                                          `json:"location,omitempty"`
	Name        *string                                          `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag                           `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags        map[string]string                                `json:"tags,omitempty"`
	Type        *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseThroughputSetting_STATUS{}

// ConvertStatusFrom populates our SqlDatabaseThroughputSetting_STATUS from the provided source
func (setting *SqlDatabaseThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.SqlDatabaseThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_SqlDatabaseThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.SqlDatabaseThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_SqlDatabaseThroughputSetting_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseThroughputSetting_STATUS
func (setting *SqlDatabaseThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.SqlDatabaseThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_SqlDatabaseThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.SqlDatabaseThroughputSetting_STATUS{}
	err := setting.AssignProperties_To_SqlDatabaseThroughputSetting_STATUS(dst)
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

// AssignProperties_From_SqlDatabaseThroughputSetting_STATUS populates our SqlDatabaseThroughputSetting_STATUS from the provided source SqlDatabaseThroughputSetting_STATUS
func (setting *SqlDatabaseThroughputSetting_STATUS) AssignProperties_From_SqlDatabaseThroughputSetting_STATUS(source *storage.SqlDatabaseThroughputSetting_STATUS) error {
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

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseThroughputSetting_STATUS populates the provided destination SqlDatabaseThroughputSetting_STATUS from our SqlDatabaseThroughputSetting_STATUS
func (setting *SqlDatabaseThroughputSetting_STATUS) AssignProperties_To_SqlDatabaseThroughputSetting_STATUS(destination *storage.SqlDatabaseThroughputSetting_STATUS) error {
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

	// Invoke the augmentConversionForSqlDatabaseThroughputSetting_STATUS interface (if implemented) to customize the conversion
	var settingAsAny any = setting
	if augmentedSetting, ok := settingAsAny.(augmentConversionForSqlDatabaseThroughputSetting_STATUS); ok {
		err := augmentedSetting.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlDatabaseThroughputSetting_Spec interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseThroughputSetting_Spec) error
	AssignPropertiesTo(dst *storage.SqlDatabaseThroughputSetting_Spec) error
}

type augmentConversionForSqlDatabaseThroughputSetting_STATUS interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseThroughputSetting_STATUS) error
	AssignPropertiesTo(dst *storage.SqlDatabaseThroughputSetting_STATUS) error
}

// Storage version of v1api20210515.SqlDatabaseThroughputSettingOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type SqlDatabaseThroughputSettingOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_SqlDatabaseThroughputSettingOperatorSpec populates our SqlDatabaseThroughputSettingOperatorSpec from the provided source SqlDatabaseThroughputSettingOperatorSpec
func (operator *SqlDatabaseThroughputSettingOperatorSpec) AssignProperties_From_SqlDatabaseThroughputSettingOperatorSpec(source *storage.SqlDatabaseThroughputSettingOperatorSpec) error {
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

	// Invoke the augmentConversionForSqlDatabaseThroughputSettingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForSqlDatabaseThroughputSettingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseThroughputSettingOperatorSpec populates the provided destination SqlDatabaseThroughputSettingOperatorSpec from our SqlDatabaseThroughputSettingOperatorSpec
func (operator *SqlDatabaseThroughputSettingOperatorSpec) AssignProperties_To_SqlDatabaseThroughputSettingOperatorSpec(destination *storage.SqlDatabaseThroughputSettingOperatorSpec) error {
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

	// Invoke the augmentConversionForSqlDatabaseThroughputSettingOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForSqlDatabaseThroughputSettingOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForSqlDatabaseThroughputSettingOperatorSpec interface {
	AssignPropertiesFrom(src *storage.SqlDatabaseThroughputSettingOperatorSpec) error
	AssignPropertiesTo(dst *storage.SqlDatabaseThroughputSettingOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseThroughputSetting{}, &SqlDatabaseThroughputSettingList{})
}
