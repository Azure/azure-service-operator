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
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources={flexibleserversconfigurations/status,flexibleserversconfigurations/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
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
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetConditions() conditions.Conditions {
	return flexibleServersConfiguration.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (flexibleServersConfiguration *FlexibleServersConfiguration) SetConditions(conditions conditions.Conditions) {
	flexibleServersConfiguration.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-dbforpostgresql-azure-com-v1alpha1api20210601-flexibleserversconfiguration,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1alpha1api20210601,name=default.v1alpha1api20210601.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &FlexibleServersConfiguration{}

// Default applies defaults to the FlexibleServersConfiguration resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) Default() {
	flexibleServersConfiguration.defaultImpl()
	var temp interface{} = flexibleServersConfiguration
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (flexibleServersConfiguration *FlexibleServersConfiguration) defaultAzureName() {
	if flexibleServersConfiguration.Spec.AzureName == "" {
		flexibleServersConfiguration.Spec.AzureName = flexibleServersConfiguration.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersConfiguration resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) defaultImpl() {
	flexibleServersConfiguration.defaultAzureName()
}

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) AzureName() string {
	return flexibleServersConfiguration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (flexibleServersConfiguration FlexibleServersConfiguration) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceKind returns the kind of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetSpec() genruntime.ConvertibleSpec {
	return &flexibleServersConfiguration.Spec
}

// GetStatus returns the status of this resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetStatus() genruntime.ConvertibleStatus {
	return &flexibleServersConfiguration.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (flexibleServersConfiguration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Configuration_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (flexibleServersConfiguration *FlexibleServersConfiguration) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(flexibleServersConfiguration.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: flexibleServersConfiguration.Namespace,
		Name:      flexibleServersConfiguration.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Configuration_Status); ok {
		flexibleServersConfiguration.Status = *st
		return nil
	}

	// Convert status to required version
	var st Configuration_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	flexibleServersConfiguration.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbforpostgresql-azure-com-v1alpha1api20210601-flexibleserversconfiguration,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=create;update,versions=v1alpha1api20210601,name=validate.v1alpha1api20210601.flexibleserversconfigurations.dbforpostgresql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &FlexibleServersConfiguration{}

// ValidateCreate validates the creation of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) ValidateCreate() error {
	validations := flexibleServersConfiguration.createValidations()
	var temp interface{} = flexibleServersConfiguration
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
func (flexibleServersConfiguration *FlexibleServersConfiguration) ValidateDelete() error {
	validations := flexibleServersConfiguration.deleteValidations()
	var temp interface{} = flexibleServersConfiguration
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
func (flexibleServersConfiguration *FlexibleServersConfiguration) ValidateUpdate(old runtime.Object) error {
	validations := flexibleServersConfiguration.updateValidations()
	var temp interface{} = flexibleServersConfiguration
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
func (flexibleServersConfiguration *FlexibleServersConfiguration) createValidations() []func() error {
	return []func() error{flexibleServersConfiguration.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return flexibleServersConfiguration.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (flexibleServersConfiguration *FlexibleServersConfiguration) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&flexibleServersConfiguration.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromFlexibleServersConfiguration populates our FlexibleServersConfiguration from the provided source FlexibleServersConfiguration
func (flexibleServersConfiguration *FlexibleServersConfiguration) AssignPropertiesFromFlexibleServersConfiguration(source *v1alpha1api20210601storage.FlexibleServersConfiguration) error {

	// ObjectMeta
	flexibleServersConfiguration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersConfigurations_Spec
	err := spec.AssignPropertiesFromFlexibleServersConfigurationsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromFlexibleServersConfigurationsSpec()")
	}
	flexibleServersConfiguration.Spec = spec

	// Status
	var status Configuration_Status
	err = status.AssignPropertiesFromConfigurationStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromConfigurationStatus()")
	}
	flexibleServersConfiguration.Status = status

	// TypeMeta
	flexibleServersConfiguration.TypeMeta = source.TypeMeta

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersConfiguration populates the provided destination FlexibleServersConfiguration from our FlexibleServersConfiguration
func (flexibleServersConfiguration *FlexibleServersConfiguration) AssignPropertiesToFlexibleServersConfiguration(destination *v1alpha1api20210601storage.FlexibleServersConfiguration) error {

	// ObjectMeta
	destination.ObjectMeta = *flexibleServersConfiguration.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210601storage.FlexibleServersConfigurations_Spec
	err := flexibleServersConfiguration.Spec.AssignPropertiesToFlexibleServersConfigurationsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToFlexibleServersConfigurationsSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210601storage.Configuration_Status
	err = flexibleServersConfiguration.Status.AssignPropertiesToConfigurationStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToConfigurationStatus()")
	}
	destination.Status = status

	// TypeMeta
	destination.TypeMeta = flexibleServersConfiguration.TypeMeta

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: flexibleServersConfiguration.Spec.OriginalVersion(),
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
func (configurationStatus *Configuration_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210601storage.Configuration_Status)
	if ok {
		// Populate our instance from source
		return configurationStatus.AssignPropertiesFromConfigurationStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210601storage.Configuration_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = configurationStatus.AssignPropertiesFromConfigurationStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Configuration_Status
func (configurationStatus *Configuration_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210601storage.Configuration_Status)
	if ok {
		// Populate destination from our instance
		return configurationStatus.AssignPropertiesToConfigurationStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210601storage.Configuration_Status{}
	err := configurationStatus.AssignPropertiesToConfigurationStatus(dst)
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
func (configurationStatus *Configuration_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Configuration_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configurationStatus *Configuration_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Configuration_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Configuration_StatusARM, got %T", armInput)
	}

	// Set property ‘AllowedValues’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AllowedValues != nil {
			allowedValues := *typedInput.Properties.AllowedValues
			configurationStatus.AllowedValues = &allowedValues
		}
	}

	// no assignment for property ‘Conditions’

	// Set property ‘DataType’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DataType != nil {
			dataType := *typedInput.Properties.DataType
			configurationStatus.DataType = &dataType
		}
	}

	// Set property ‘DefaultValue’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DefaultValue != nil {
			defaultValue := *typedInput.Properties.DefaultValue
			configurationStatus.DefaultValue = &defaultValue
		}
	}

	// Set property ‘Description’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			configurationStatus.Description = &description
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		configurationStatus.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		configurationStatus.Name = &name
	}

	// Set property ‘Source’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Source != nil {
			source := *typedInput.Properties.Source
			configurationStatus.Source = &source
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
		configurationStatus.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		configurationStatus.Type = &typeVar
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			configurationStatus.Value = &value
		}
	}

	// No error
	return nil
}

// AssignPropertiesFromConfigurationStatus populates our Configuration_Status from the provided source Configuration_Status
func (configurationStatus *Configuration_Status) AssignPropertiesFromConfigurationStatus(source *v1alpha1api20210601storage.Configuration_Status) error {

	// AllowedValues
	configurationStatus.AllowedValues = genruntime.ClonePointerToString(source.AllowedValues)

	// Conditions
	configurationStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DataType
	if source.DataType != nil {
		dataType := ConfigurationPropertiesStatusDataType(*source.DataType)
		configurationStatus.DataType = &dataType
	} else {
		configurationStatus.DataType = nil
	}

	// DefaultValue
	configurationStatus.DefaultValue = genruntime.ClonePointerToString(source.DefaultValue)

	// Description
	configurationStatus.Description = genruntime.ClonePointerToString(source.Description)

	// Id
	configurationStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	configurationStatus.Name = genruntime.ClonePointerToString(source.Name)

	// Source
	configurationStatus.Source = genruntime.ClonePointerToString(source.Source)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesFromSystemDataStatus()")
		}
		configurationStatus.SystemData = &systemDatum
	} else {
		configurationStatus.SystemData = nil
	}

	// Type
	configurationStatus.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	configurationStatus.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignPropertiesToConfigurationStatus populates the provided destination Configuration_Status from our Configuration_Status
func (configurationStatus *Configuration_Status) AssignPropertiesToConfigurationStatus(destination *v1alpha1api20210601storage.Configuration_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AllowedValues
	destination.AllowedValues = genruntime.ClonePointerToString(configurationStatus.AllowedValues)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(configurationStatus.Conditions)

	// DataType
	if configurationStatus.DataType != nil {
		dataType := string(*configurationStatus.DataType)
		destination.DataType = &dataType
	} else {
		destination.DataType = nil
	}

	// DefaultValue
	destination.DefaultValue = genruntime.ClonePointerToString(configurationStatus.DefaultValue)

	// Description
	destination.Description = genruntime.ClonePointerToString(configurationStatus.Description)

	// Id
	destination.Id = genruntime.ClonePointerToString(configurationStatus.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(configurationStatus.Name)

	// Source
	destination.Source = genruntime.ClonePointerToString(configurationStatus.Source)

	// SystemData
	if configurationStatus.SystemData != nil {
		var systemDatum v1alpha1api20210601storage.SystemData_Status
		err := (*configurationStatus.SystemData).AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "populating SystemData from SystemData, calling AssignPropertiesToSystemDataStatus()")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(configurationStatus.Type)

	// Value
	destination.Value = genruntime.ClonePointerToString(configurationStatus.Value)

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
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if flexibleServersConfigurationsSpec == nil {
		return nil, nil
	}
	var result FlexibleServersConfigurations_SpecARM

	// Set property ‘Location’:
	if flexibleServersConfigurationsSpec.Location != nil {
		location := *flexibleServersConfigurationsSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if flexibleServersConfigurationsSpec.Source != nil {
		source := *flexibleServersConfigurationsSpec.Source
		result.Properties.Source = &source
	}
	if flexibleServersConfigurationsSpec.Value != nil {
		value := *flexibleServersConfigurationsSpec.Value
		result.Properties.Value = &value
	}

	// Set property ‘Tags’:
	if flexibleServersConfigurationsSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, tagsValue := range flexibleServersConfigurationsSpec.Tags {
			result.Tags[key] = tagsValue
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersConfigurations_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersConfigurations_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersConfigurations_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	flexibleServersConfigurationsSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		flexibleServersConfigurationsSpec.Location = &location
	}

	// Set property ‘Owner’:
	flexibleServersConfigurationsSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Source’:
	// copying flattened property:
	if typedInput.Properties.Source != nil {
		source := *typedInput.Properties.Source
		flexibleServersConfigurationsSpec.Source = &source
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		flexibleServersConfigurationsSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			flexibleServersConfigurationsSpec.Tags[key] = value
		}
	}

	// Set property ‘Value’:
	// copying flattened property:
	if typedInput.Properties.Value != nil {
		value := *typedInput.Properties.Value
		flexibleServersConfigurationsSpec.Value = &value
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersConfigurations_Spec{}

// ConvertSpecFrom populates our FlexibleServersConfigurations_Spec from the provided source
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210601storage.FlexibleServersConfigurations_Spec)
	if ok {
		// Populate our instance from source
		return flexibleServersConfigurationsSpec.AssignPropertiesFromFlexibleServersConfigurationsSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210601storage.FlexibleServersConfigurations_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = flexibleServersConfigurationsSpec.AssignPropertiesFromFlexibleServersConfigurationsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersConfigurations_Spec
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210601storage.FlexibleServersConfigurations_Spec)
	if ok {
		// Populate destination from our instance
		return flexibleServersConfigurationsSpec.AssignPropertiesToFlexibleServersConfigurationsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210601storage.FlexibleServersConfigurations_Spec{}
	err := flexibleServersConfigurationsSpec.AssignPropertiesToFlexibleServersConfigurationsSpec(dst)
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
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) AssignPropertiesFromFlexibleServersConfigurationsSpec(source *v1alpha1api20210601storage.FlexibleServersConfigurations_Spec) error {

	// AzureName
	flexibleServersConfigurationsSpec.AzureName = source.AzureName

	// Location
	flexibleServersConfigurationsSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	flexibleServersConfigurationsSpec.Owner = source.Owner.Copy()

	// Source
	flexibleServersConfigurationsSpec.Source = genruntime.ClonePointerToString(source.Source)

	// Tags
	flexibleServersConfigurationsSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Value
	flexibleServersConfigurationsSpec.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersConfigurationsSpec populates the provided destination FlexibleServersConfigurations_Spec from our FlexibleServersConfigurations_Spec
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) AssignPropertiesToFlexibleServersConfigurationsSpec(destination *v1alpha1api20210601storage.FlexibleServersConfigurations_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = flexibleServersConfigurationsSpec.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(flexibleServersConfigurationsSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = flexibleServersConfigurationsSpec.OriginalVersion()

	// Owner
	destination.Owner = flexibleServersConfigurationsSpec.Owner.Copy()

	// Source
	destination.Source = genruntime.ClonePointerToString(flexibleServersConfigurationsSpec.Source)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(flexibleServersConfigurationsSpec.Tags)

	// Value
	destination.Value = genruntime.ClonePointerToString(flexibleServersConfigurationsSpec.Value)

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
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) SetAzureName(azureName string) {
	flexibleServersConfigurationsSpec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
