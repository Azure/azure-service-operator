// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	storage "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20230630/storage"
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
// Storage version of v1api20220101.FlexibleServersConfiguration
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/Configurations/stable/2022-01-01/Configurations.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/configurations/{configurationName}
type FlexibleServersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersConfiguration_Spec   `json:"spec,omitempty"`
	Status            FlexibleServersConfiguration_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersConfiguration{}

// GetConditions returns the conditions of the resource
func (configuration *FlexibleServersConfiguration) GetConditions() conditions.Conditions {
	return configuration.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (configuration *FlexibleServersConfiguration) SetConditions(conditions conditions.Conditions) {
	configuration.Status.Conditions = conditions
}

var _ conversion.Convertible = &FlexibleServersConfiguration{}

// ConvertFrom populates our FlexibleServersConfiguration from the provided hub FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source storage.FlexibleServersConfiguration

	err := source.ConvertFrom(hub)
	if err != nil {
		return eris.Wrap(err, "converting from hub to source")
	}

	err = configuration.AssignProperties_From_FlexibleServersConfiguration(&source)
	if err != nil {
		return eris.Wrap(err, "converting from source to configuration")
	}

	return nil
}

// ConvertTo populates the provided hub FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.FlexibleServersConfiguration
	err := configuration.AssignProperties_To_FlexibleServersConfiguration(&destination)
	if err != nil {
		return eris.Wrap(err, "converting to destination from configuration")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return eris.Wrap(err, "converting from destination to hub")
	}

	return nil
}

var _ configmaps.Exporter = &FlexibleServersConfiguration{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (configuration *FlexibleServersConfiguration) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if configuration.Spec.OperatorSpec == nil {
		return nil
	}
	return configuration.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &FlexibleServersConfiguration{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (configuration *FlexibleServersConfiguration) SecretDestinationExpressions() []*core.DestinationExpression {
	if configuration.Spec.OperatorSpec == nil {
		return nil
	}
	return configuration.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (configuration *FlexibleServersConfiguration) AzureName() string {
	return configuration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-01-01"
func (configuration FlexibleServersConfiguration) GetAPIVersion() string {
	return "2022-01-01"
}

// GetResourceScope returns the scope of the resource
func (configuration *FlexibleServersConfiguration) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (configuration *FlexibleServersConfiguration) GetSpec() genruntime.ConvertibleSpec {
	return &configuration.Spec
}

// GetStatus returns the status of this resource
func (configuration *FlexibleServersConfiguration) GetStatus() genruntime.ConvertibleStatus {
	return &configuration.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (configuration *FlexibleServersConfiguration) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (configuration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServersConfiguration_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (configuration *FlexibleServersConfiguration) Owner() *genruntime.ResourceReference {
	if configuration.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(configuration.Spec)
	return configuration.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (configuration *FlexibleServersConfiguration) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServersConfiguration_STATUS); ok {
		configuration.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersConfiguration_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	configuration.Status = st
	return nil
}

// AssignProperties_From_FlexibleServersConfiguration populates our FlexibleServersConfiguration from the provided source FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignProperties_From_FlexibleServersConfiguration(source *storage.FlexibleServersConfiguration) error {

	// ObjectMeta
	configuration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersConfiguration_Spec
	err := spec.AssignProperties_From_FlexibleServersConfiguration_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersConfiguration_Spec() to populate field Spec")
	}
	configuration.Spec = spec

	// Status
	var status FlexibleServersConfiguration_STATUS
	err = status.AssignProperties_From_FlexibleServersConfiguration_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersConfiguration_STATUS() to populate field Status")
	}
	configuration.Status = status

	// Invoke the augmentConversionForFlexibleServersConfiguration interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration); ok {
		err := augmentedConfiguration.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersConfiguration populates the provided destination FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignProperties_To_FlexibleServersConfiguration(destination *storage.FlexibleServersConfiguration) error {

	// ObjectMeta
	destination.ObjectMeta = *configuration.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.FlexibleServersConfiguration_Spec
	err := configuration.Spec.AssignProperties_To_FlexibleServersConfiguration_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersConfiguration_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.FlexibleServersConfiguration_STATUS
	err = configuration.Status.AssignProperties_To_FlexibleServersConfiguration_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersConfiguration_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFlexibleServersConfiguration interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration); ok {
		err := augmentedConfiguration.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (configuration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: configuration.Spec.OriginalVersion,
		Kind:    "FlexibleServersConfiguration",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220101.FlexibleServersConfiguration
// Generator information:
// - Generated from: /mysql/resource-manager/Microsoft.DBforMySQL/Configurations/stable/2022-01-01/Configurations.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforMySQL/flexibleServers/{serverName}/configurations/{configurationName}
type FlexibleServersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersConfiguration `json:"items"`
}

type augmentConversionForFlexibleServersConfiguration interface {
	AssignPropertiesFrom(src *storage.FlexibleServersConfiguration) error
	AssignPropertiesTo(dst *storage.FlexibleServersConfiguration) error
}

// Storage version of v1api20220101.FlexibleServersConfiguration_Spec
type FlexibleServersConfiguration_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                    `json:"azureName,omitempty"`
	CurrentValue    *string                                   `json:"currentValue,omitempty"`
	OperatorSpec    *FlexibleServersConfigurationOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                    `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformysql.azure.com/FlexibleServer resource
	Owner       *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Source      *string                            `json:"source,omitempty"`
	Value       *string                            `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersConfiguration_Spec{}

// ConvertSpecFrom populates our FlexibleServersConfiguration_Spec from the provided source
func (configuration *FlexibleServersConfiguration_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.FlexibleServersConfiguration_Spec)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_FlexibleServersConfiguration_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersConfiguration_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_FlexibleServersConfiguration_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersConfiguration_Spec
func (configuration *FlexibleServersConfiguration_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.FlexibleServersConfiguration_Spec)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_FlexibleServersConfiguration_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersConfiguration_Spec{}
	err := configuration.AssignProperties_To_FlexibleServersConfiguration_Spec(dst)
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

// AssignProperties_From_FlexibleServersConfiguration_Spec populates our FlexibleServersConfiguration_Spec from the provided source FlexibleServersConfiguration_Spec
func (configuration *FlexibleServersConfiguration_Spec) AssignProperties_From_FlexibleServersConfiguration_Spec(source *storage.FlexibleServersConfiguration_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	configuration.AzureName = source.AzureName

	// CurrentValue
	configuration.CurrentValue = genruntime.ClonePointerToString(source.CurrentValue)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec FlexibleServersConfigurationOperatorSpec
		err := operatorSpec.AssignProperties_From_FlexibleServersConfigurationOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersConfigurationOperatorSpec() to populate field OperatorSpec")
		}
		configuration.OperatorSpec = &operatorSpec
	} else {
		configuration.OperatorSpec = nil
	}

	// OriginalVersion
	configuration.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		configuration.Owner = &owner
	} else {
		configuration.Owner = nil
	}

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		configuration.PropertyBag = propertyBag
	} else {
		configuration.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersConfiguration_Spec interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration_Spec); ok {
		err := augmentedConfiguration.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersConfiguration_Spec populates the provided destination FlexibleServersConfiguration_Spec from our FlexibleServersConfiguration_Spec
func (configuration *FlexibleServersConfiguration_Spec) AssignProperties_To_FlexibleServersConfiguration_Spec(destination *storage.FlexibleServersConfiguration_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(configuration.PropertyBag)

	// AzureName
	destination.AzureName = configuration.AzureName

	// CurrentValue
	destination.CurrentValue = genruntime.ClonePointerToString(configuration.CurrentValue)

	// OperatorSpec
	if configuration.OperatorSpec != nil {
		var operatorSpec storage.FlexibleServersConfigurationOperatorSpec
		err := configuration.OperatorSpec.AssignProperties_To_FlexibleServersConfigurationOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersConfigurationOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = configuration.OriginalVersion

	// Owner
	if configuration.Owner != nil {
		owner := configuration.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// Value
	destination.Value = genruntime.ClonePointerToString(configuration.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersConfiguration_Spec interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration_Spec); ok {
		err := augmentedConfiguration.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20220101.FlexibleServersConfiguration_STATUS
type FlexibleServersConfiguration_STATUS struct {
	AllowedValues          *string                `json:"allowedValues,omitempty"`
	Conditions             []conditions.Condition `json:"conditions,omitempty"`
	CurrentValue           *string                `json:"currentValue,omitempty"`
	DataType               *string                `json:"dataType,omitempty"`
	DefaultValue           *string                `json:"defaultValue,omitempty"`
	Description            *string                `json:"description,omitempty"`
	DocumentationLink      *string                `json:"documentationLink,omitempty"`
	Id                     *string                `json:"id,omitempty"`
	IsConfigPendingRestart *string                `json:"isConfigPendingRestart,omitempty"`
	IsDynamicConfig        *string                `json:"isDynamicConfig,omitempty"`
	IsReadOnly             *string                `json:"isReadOnly,omitempty"`
	Name                   *string                `json:"name,omitempty"`
	PropertyBag            genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Source                 *string                `json:"source,omitempty"`
	SystemData             *SystemData_STATUS     `json:"systemData,omitempty"`
	Type                   *string                `json:"type,omitempty"`
	Value                  *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServersConfiguration_STATUS{}

// ConvertStatusFrom populates our FlexibleServersConfiguration_STATUS from the provided source
func (configuration *FlexibleServersConfiguration_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.FlexibleServersConfiguration_STATUS)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_FlexibleServersConfiguration_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersConfiguration_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_FlexibleServersConfiguration_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServersConfiguration_STATUS
func (configuration *FlexibleServersConfiguration_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.FlexibleServersConfiguration_STATUS)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_FlexibleServersConfiguration_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersConfiguration_STATUS{}
	err := configuration.AssignProperties_To_FlexibleServersConfiguration_STATUS(dst)
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

// AssignProperties_From_FlexibleServersConfiguration_STATUS populates our FlexibleServersConfiguration_STATUS from the provided source FlexibleServersConfiguration_STATUS
func (configuration *FlexibleServersConfiguration_STATUS) AssignProperties_From_FlexibleServersConfiguration_STATUS(source *storage.FlexibleServersConfiguration_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AllowedValues
	configuration.AllowedValues = genruntime.ClonePointerToString(source.AllowedValues)

	// Conditions
	configuration.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CurrentValue
	configuration.CurrentValue = genruntime.ClonePointerToString(source.CurrentValue)

	// DataType
	configuration.DataType = genruntime.ClonePointerToString(source.DataType)

	// DefaultValue
	configuration.DefaultValue = genruntime.ClonePointerToString(source.DefaultValue)

	// Description
	configuration.Description = genruntime.ClonePointerToString(source.Description)

	// DocumentationLink
	configuration.DocumentationLink = genruntime.ClonePointerToString(source.DocumentationLink)

	// Id
	configuration.Id = genruntime.ClonePointerToString(source.Id)

	// IsConfigPendingRestart
	configuration.IsConfigPendingRestart = genruntime.ClonePointerToString(source.IsConfigPendingRestart)

	// IsDynamicConfig
	configuration.IsDynamicConfig = genruntime.ClonePointerToString(source.IsDynamicConfig)

	// IsReadOnly
	configuration.IsReadOnly = genruntime.ClonePointerToString(source.IsReadOnly)

	// Name
	configuration.Name = genruntime.ClonePointerToString(source.Name)

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		configuration.SystemData = &systemDatum
	} else {
		configuration.SystemData = nil
	}

	// Type
	configuration.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		configuration.PropertyBag = propertyBag
	} else {
		configuration.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersConfiguration_STATUS interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration_STATUS); ok {
		err := augmentedConfiguration.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersConfiguration_STATUS populates the provided destination FlexibleServersConfiguration_STATUS from our FlexibleServersConfiguration_STATUS
func (configuration *FlexibleServersConfiguration_STATUS) AssignProperties_To_FlexibleServersConfiguration_STATUS(destination *storage.FlexibleServersConfiguration_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(configuration.PropertyBag)

	// AllowedValues
	destination.AllowedValues = genruntime.ClonePointerToString(configuration.AllowedValues)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(configuration.Conditions)

	// CurrentValue
	destination.CurrentValue = genruntime.ClonePointerToString(configuration.CurrentValue)

	// DataType
	destination.DataType = genruntime.ClonePointerToString(configuration.DataType)

	// DefaultValue
	destination.DefaultValue = genruntime.ClonePointerToString(configuration.DefaultValue)

	// Description
	destination.Description = genruntime.ClonePointerToString(configuration.Description)

	// DocumentationLink
	destination.DocumentationLink = genruntime.ClonePointerToString(configuration.DocumentationLink)

	// Id
	destination.Id = genruntime.ClonePointerToString(configuration.Id)

	// IsConfigPendingRestart
	destination.IsConfigPendingRestart = genruntime.ClonePointerToString(configuration.IsConfigPendingRestart)

	// IsDynamicConfig
	destination.IsDynamicConfig = genruntime.ClonePointerToString(configuration.IsDynamicConfig)

	// IsReadOnly
	destination.IsReadOnly = genruntime.ClonePointerToString(configuration.IsReadOnly)

	// Name
	destination.Name = genruntime.ClonePointerToString(configuration.Name)

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// SystemData
	if configuration.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := configuration.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(configuration.Type)

	// Value
	destination.Value = genruntime.ClonePointerToString(configuration.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersConfiguration_STATUS interface (if implemented) to customize the conversion
	var configurationAsAny any = configuration
	if augmentedConfiguration, ok := configurationAsAny.(augmentConversionForFlexibleServersConfiguration_STATUS); ok {
		err := augmentedConfiguration.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersConfiguration_Spec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersConfiguration_Spec) error
	AssignPropertiesTo(dst *storage.FlexibleServersConfiguration_Spec) error
}

type augmentConversionForFlexibleServersConfiguration_STATUS interface {
	AssignPropertiesFrom(src *storage.FlexibleServersConfiguration_STATUS) error
	AssignPropertiesTo(dst *storage.FlexibleServersConfiguration_STATUS) error
}

// Storage version of v1api20220101.FlexibleServersConfigurationOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FlexibleServersConfigurationOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_FlexibleServersConfigurationOperatorSpec populates our FlexibleServersConfigurationOperatorSpec from the provided source FlexibleServersConfigurationOperatorSpec
func (operator *FlexibleServersConfigurationOperatorSpec) AssignProperties_From_FlexibleServersConfigurationOperatorSpec(source *storage.FlexibleServersConfigurationOperatorSpec) error {
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

	// Invoke the augmentConversionForFlexibleServersConfigurationOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersConfigurationOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersConfigurationOperatorSpec populates the provided destination FlexibleServersConfigurationOperatorSpec from our FlexibleServersConfigurationOperatorSpec
func (operator *FlexibleServersConfigurationOperatorSpec) AssignProperties_To_FlexibleServersConfigurationOperatorSpec(destination *storage.FlexibleServersConfigurationOperatorSpec) error {
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

	// Invoke the augmentConversionForFlexibleServersConfigurationOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersConfigurationOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersConfigurationOperatorSpec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersConfigurationOperatorSpec) error
	AssignPropertiesTo(dst *storage.FlexibleServersConfigurationOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
