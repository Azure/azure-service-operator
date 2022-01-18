// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1alpha1api20210601storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_configurations
type FlexibleServersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersConfigurations_Spec `json:"spec,omitempty"`
	Status            Configuration_Status               `json:"status,omitempty"`
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
	source, ok := hub.(*v1alpha1api20210601storage.FlexibleServersConfiguration)
	if !ok {
		return fmt.Errorf("expected storage:dbforpostgresql/v1alpha1api20210601storage/FlexibleServersConfiguration but received %T instead", hub)
	}

	return configuration.AssignPropertiesFromFlexibleServersConfiguration(source)
}

// ConvertTo populates the provided hub FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20210601storage.FlexibleServersConfiguration)
	if !ok {
		return fmt.Errorf("expected storage:dbforpostgresql/v1alpha1api20210601storage/FlexibleServersConfiguration but received %T instead", hub)
	}

	return configuration.AssignPropertiesToFlexibleServersConfiguration(destination)
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1alpha1api20210601-flexibleserversconfiguration,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1alpha1api20210601,name=default.v1alpha1api20210601.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &FlexibleServersConfiguration{}

// Default applies defaults to the FlexibleServersConfiguration resource
func (configuration *FlexibleServersConfiguration) Default() {
	configuration.defaultImpl()
	var temp interface{} = configuration
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (configuration *FlexibleServersConfiguration) defaultAzureName() {
	if configuration.Spec.AzureName == "" {
		configuration.Spec.AzureName = configuration.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersConfiguration resource
func (configuration *FlexibleServersConfiguration) defaultImpl() { configuration.defaultAzureName() }

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (configuration *FlexibleServersConfiguration) AzureName() string {
	return configuration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (configuration FlexibleServersConfiguration) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceKind returns the kind of the resource
func (configuration *FlexibleServersConfiguration) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (configuration *FlexibleServersConfiguration) GetSpec() genruntime.ConvertibleSpec {
	return &configuration.Spec
}

// GetStatus returns the status of this resource
func (configuration *FlexibleServersConfiguration) GetStatus() genruntime.ConvertibleStatus {
	return &configuration.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (configuration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Configuration_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (configuration *FlexibleServersConfiguration) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(configuration.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  configuration.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (configuration *FlexibleServersConfiguration) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Configuration_Status); ok {
		configuration.Status = *st
		return nil
	}

	// Convert status to required version
	var st Configuration_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	configuration.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1alpha1api20210601-flexibleserversconfiguration,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1alpha1api20210601,name=validate.v1alpha1api20210601.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &FlexibleServersConfiguration{}

// ValidateCreate validates the creation of the resource
func (configuration *FlexibleServersConfiguration) ValidateCreate() error {
	validations := configuration.createValidations()
	var temp interface{} = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (configuration *FlexibleServersConfiguration) ValidateDelete() error {
	validations := configuration.deleteValidations()
	var temp interface{} = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (configuration *FlexibleServersConfiguration) ValidateUpdate(old runtime.Object) error {
	validations := configuration.updateValidations()
	var temp interface{} = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (configuration *FlexibleServersConfiguration) createValidations() []func() error {
	return []func() error{configuration.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (configuration *FlexibleServersConfiguration) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (configuration *FlexibleServersConfiguration) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return configuration.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (configuration *FlexibleServersConfiguration) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&configuration.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromFlexibleServersConfiguration populates our FlexibleServersConfiguration from the provided source FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignPropertiesFromFlexibleServersConfiguration(source *v1alpha1api20210601storage.FlexibleServersConfiguration) error {

	// ObjectMeta
	configuration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersConfigurations_Spec
	err := spec.AssignPropertiesFromFlexibleServersConfigurationsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersConfigurationsSpec() to populate field Spec")
	}
	configuration.Spec = spec

	// Status
	var status Configuration_Status
	err = status.AssignPropertiesFromConfigurationStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromConfigurationStatus() to populate field Status")
	}
	configuration.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersConfiguration populates the provided destination FlexibleServersConfiguration from our FlexibleServersConfiguration
func (configuration *FlexibleServersConfiguration) AssignPropertiesToFlexibleServersConfiguration(destination *v1alpha1api20210601storage.FlexibleServersConfiguration) error {

	// ObjectMeta
	destination.ObjectMeta = *configuration.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210601storage.FlexibleServersConfigurations_Spec
	err := configuration.Spec.AssignPropertiesToFlexibleServersConfigurationsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersConfigurationsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210601storage.Configuration_Status
	err = configuration.Status.AssignPropertiesToConfigurationStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToConfigurationStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (configuration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: configuration.Spec.OriginalVersion(),
		Kind:    "FlexibleServersConfiguration",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_configurations
type FlexibleServersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersConfiguration `json:"items"`
}

type Configuration_Status struct {
	//AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//DataType: Data type of the configuration.
	DataType *ConfigurationPropertiesStatusDataType `json:"dataType,omitempty"`

	//DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	//Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	//SystemData: The system metadata relating to this resource.
	SystemData *SystemData_Status `json:"systemData,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	//Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Configuration_Status{}

// ConvertStatusFrom populates our Configuration_Status from the provided source
func (configuration *Configuration_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210601storage.Configuration_Status)
	if ok {
		// Populate our instance from source
		return configuration.AssignPropertiesFromConfigurationStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210601storage.Configuration_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = configuration.AssignPropertiesFromConfigurationStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Configuration_Status
func (configuration *Configuration_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210601storage.Configuration_Status)
	if ok {
		// Populate destination from our instance
		return configuration.AssignPropertiesToConfigurationStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210601storage.Configuration_Status{}
	err := configuration.AssignPropertiesToConfigurationStatus(dst)
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

var _ genruntime.FromARMConverter = &Configuration_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (configuration *Configuration_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Configuration_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configuration *Configuration_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Configuration_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Configuration_StatusARM, got %T", armInput)
	}

	// Set property ‘AllowedValues’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AllowedValues != nil {
			allowedValues := *typedInput.Properties.AllowedValues
			configuration.AllowedValues = &allowedValues
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘DataType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DataType != nil {
			dataType := *typedInput.Properties.DataType
			configuration.DataType = &dataType
		}
	}

	// Set property ‘DefaultValue’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DefaultValue != nil {
			defaultValue := *typedInput.Properties.DefaultValue
			configuration.DefaultValue = &defaultValue
		}
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			configuration.Description = &description
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		configuration.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		configuration.Name = &name
	}

	// Set property ‘Source’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Source != nil {
			source := *typedInput.Properties.Source
			configuration.Source = &source
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_Status
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		configuration.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		configuration.Type = &typeVar
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			configuration.Value = &value
		}
	}

	// No error
	return nil
}

// AssignPropertiesFromConfigurationStatus populates our Configuration_Status from the provided source Configuration_Status
func (configuration *Configuration_Status) AssignPropertiesFromConfigurationStatus(source *v1alpha1api20210601storage.Configuration_Status) error {

	// AllowedValues
	configuration.AllowedValues = genruntime.ClonePointerToString(source.AllowedValues)

	// Conditions
	configuration.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DataType
	if source.DataType != nil {
		dataType := ConfigurationPropertiesStatusDataType(*source.DataType)
		configuration.DataType = &dataType
	} else {
		configuration.DataType = nil
	}

	// DefaultValue
	configuration.DefaultValue = genruntime.ClonePointerToString(source.DefaultValue)

	// Description
	configuration.Description = genruntime.ClonePointerToString(source.Description)

	// Id
	configuration.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	configuration.Name = genruntime.ClonePointerToString(source.Name)

	// Source
	configuration.Source = genruntime.ClonePointerToString(source.Source)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		configuration.SystemData = &systemDatum
	} else {
		configuration.SystemData = nil
	}

	// Type
	configuration.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignPropertiesToConfigurationStatus populates the provided destination Configuration_Status from our Configuration_Status
func (configuration *Configuration_Status) AssignPropertiesToConfigurationStatus(destination *v1alpha1api20210601storage.Configuration_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AllowedValues
	destination.AllowedValues = genruntime.ClonePointerToString(configuration.AllowedValues)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(configuration.Conditions)

	// DataType
	if configuration.DataType != nil {
		dataType := string(*configuration.DataType)
		destination.DataType = &dataType
	} else {
		destination.DataType = nil
	}

	// DefaultValue
	destination.DefaultValue = genruntime.ClonePointerToString(configuration.DefaultValue)

	// Description
	destination.Description = genruntime.ClonePointerToString(configuration.Description)

	// Id
	destination.Id = genruntime.ClonePointerToString(configuration.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(configuration.Name)

	// Source
	destination.Source = genruntime.ClonePointerToString(configuration.Source)

	// SystemData
	if configuration.SystemData != nil {
		var systemDatum v1alpha1api20210601storage.SystemData_Status
		err := configuration.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
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

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2021-06-01"}
type FlexibleServersConfigurationsSpecAPIVersion string

const FlexibleServersConfigurationsSpecAPIVersion20210601 = FlexibleServersConfigurationsSpecAPIVersion("2021-06-01")

type FlexibleServersConfigurations_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner" kind:"FlexibleServer"`

	//Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	//Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ARMTransformer = &FlexibleServersConfigurations_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (configurations *FlexibleServersConfigurations_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if configurations == nil {
		return nil, nil
	}
	var result FlexibleServersConfigurations_SpecARM

	// Set property ‘Location’:
	if configurations.Location != nil {
		location := *configurations.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if configurations.Source != nil {
		source := *configurations.Source
		result.Properties.Source = &source
	}
	if configurations.Value != nil {
		value := *configurations.Value
		result.Properties.Value = &value
	}

	// Set property ‘Tags’:
	if configurations.Tags != nil {
		result.Tags = make(map[string]string)
		for key, tagsValue := range configurations.Tags {
			result.Tags[key] = tagsValue
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (configurations *FlexibleServersConfigurations_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersConfigurations_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configurations *FlexibleServersConfigurations_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersConfigurations_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersConfigurations_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	configurations.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		configurations.Location = &location
	}

	// Set property ‘Owner’:
	configurations.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Source’:
	// copying flattened property:
	if typedInput.Properties.Source != nil {
		source := *typedInput.Properties.Source
		configurations.Source = &source
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		configurations.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			configurations.Tags[key] = value
		}
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties.Value != nil {
		value := *typedInput.Properties.Value
		configurations.Value = &value
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersConfigurations_Spec{}

// ConvertSpecFrom populates our FlexibleServersConfigurations_Spec from the provided source
func (configurations *FlexibleServersConfigurations_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210601storage.FlexibleServersConfigurations_Spec)
	if ok {
		// Populate our instance from source
		return configurations.AssignPropertiesFromFlexibleServersConfigurationsSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210601storage.FlexibleServersConfigurations_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = configurations.AssignPropertiesFromFlexibleServersConfigurationsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersConfigurations_Spec
func (configurations *FlexibleServersConfigurations_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210601storage.FlexibleServersConfigurations_Spec)
	if ok {
		// Populate destination from our instance
		return configurations.AssignPropertiesToFlexibleServersConfigurationsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210601storage.FlexibleServersConfigurations_Spec{}
	err := configurations.AssignPropertiesToFlexibleServersConfigurationsSpec(dst)
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

// AssignPropertiesFromFlexibleServersConfigurationsSpec populates our FlexibleServersConfigurations_Spec from the provided source FlexibleServersConfigurations_Spec
func (configurations *FlexibleServersConfigurations_Spec) AssignPropertiesFromFlexibleServersConfigurationsSpec(source *v1alpha1api20210601storage.FlexibleServersConfigurations_Spec) error {

	// AzureName
	configurations.AzureName = source.AzureName

	// Location
	configurations.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	configurations.Owner = source.Owner.Copy()

	// Source
	configurations.Source = genruntime.ClonePointerToString(source.Source)

	// Tags
	configurations.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Value
	configurations.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersConfigurationsSpec populates the provided destination FlexibleServersConfigurations_Spec from our FlexibleServersConfigurations_Spec
func (configurations *FlexibleServersConfigurations_Spec) AssignPropertiesToFlexibleServersConfigurationsSpec(destination *v1alpha1api20210601storage.FlexibleServersConfigurations_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = configurations.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(configurations.Location)

	// OriginalVersion
	destination.OriginalVersion = configurations.OriginalVersion()

	// Owner
	destination.Owner = configurations.Owner.Copy()

	// Source
	destination.Source = genruntime.ClonePointerToString(configurations.Source)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(configurations.Tags)

	// Value
	destination.Value = genruntime.ClonePointerToString(configurations.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (configurations *FlexibleServersConfigurations_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (configurations *FlexibleServersConfigurations_Spec) SetAzureName(azureName string) {
	configurations.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
