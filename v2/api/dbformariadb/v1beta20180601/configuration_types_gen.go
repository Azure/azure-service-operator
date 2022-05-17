// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20180601

import (
	"fmt"
	v20180601s "github.com/Azure/azure-service-operator/v2/api/dbformariadb/v1beta20180601storage"
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
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/resourceDefinitions/servers_configurations
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServersConfigurations_Spec `json:"spec,omitempty"`
	Status            Configuration_Status       `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Configuration{}

// GetConditions returns the conditions of the resource
func (configuration *Configuration) GetConditions() conditions.Conditions {
	return configuration.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (configuration *Configuration) SetConditions(conditions conditions.Conditions) {
	configuration.Status.Conditions = conditions
}

var _ conversion.Convertible = &Configuration{}

// ConvertFrom populates our Configuration from the provided hub Configuration
func (configuration *Configuration) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20180601s.Configuration)
	if !ok {
		return fmt.Errorf("expected dbformariadb/v1beta20180601storage/Configuration but received %T instead", hub)
	}

	return configuration.AssignPropertiesFromConfiguration(source)
}

// ConvertTo populates the provided hub Configuration from our Configuration
func (configuration *Configuration) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20180601s.Configuration)
	if !ok {
		return fmt.Errorf("expected dbformariadb/v1beta20180601storage/Configuration but received %T instead", hub)
	}

	return configuration.AssignPropertiesToConfiguration(destination)
}

// +kubebuilder:webhook:path=/mutate-dbformariadb-azure-com-v1beta20180601-configuration,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=configurations,verbs=create;update,versions=v1beta20180601,name=default.v1beta20180601.configurations.dbformariadb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &Configuration{}

// Default applies defaults to the Configuration resource
func (configuration *Configuration) Default() {
	configuration.defaultImpl()
	var temp interface{} = configuration
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (configuration *Configuration) defaultAzureName() {
	if configuration.Spec.AzureName == "" {
		configuration.Spec.AzureName = configuration.Name
	}
}

// defaultImpl applies the code generated defaults to the Configuration resource
func (configuration *Configuration) defaultImpl() { configuration.defaultAzureName() }

var _ genruntime.KubernetesResource = &Configuration{}

// AzureName returns the Azure name of the resource
func (configuration *Configuration) AzureName() string {
	return configuration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2018-06-01"
func (configuration Configuration) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (configuration *Configuration) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (configuration *Configuration) GetSpec() genruntime.ConvertibleSpec {
	return &configuration.Spec
}

// GetStatus returns the status of this resource
func (configuration *Configuration) GetStatus() genruntime.ConvertibleStatus {
	return &configuration.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMariaDB/servers/configurations"
func (configuration *Configuration) GetType() string {
	return "Microsoft.DBforMariaDB/servers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (configuration *Configuration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Configuration_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (configuration *Configuration) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(configuration.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  configuration.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (configuration *Configuration) SetStatus(status genruntime.ConvertibleStatus) error {
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

// +kubebuilder:webhook:path=/validate-dbformariadb-azure-com-v1beta20180601-configuration,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=configurations,verbs=create;update,versions=v1beta20180601,name=validate.v1beta20180601.configurations.dbformariadb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &Configuration{}

// ValidateCreate validates the creation of the resource
func (configuration *Configuration) ValidateCreate() error {
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
func (configuration *Configuration) ValidateDelete() error {
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
func (configuration *Configuration) ValidateUpdate(old runtime.Object) error {
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
func (configuration *Configuration) createValidations() []func() error {
	return []func() error{configuration.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (configuration *Configuration) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (configuration *Configuration) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return configuration.validateResourceReferences()
		},
		configuration.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (configuration *Configuration) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&configuration.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (configuration *Configuration) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*Configuration)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, configuration)
}

// AssignPropertiesFromConfiguration populates our Configuration from the provided source Configuration
func (configuration *Configuration) AssignPropertiesFromConfiguration(source *v20180601s.Configuration) error {

	// ObjectMeta
	configuration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec ServersConfigurations_Spec
	err := spec.AssignPropertiesFromServersConfigurationsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromServersConfigurationsSpec() to populate field Spec")
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

// AssignPropertiesToConfiguration populates the provided destination Configuration from our Configuration
func (configuration *Configuration) AssignPropertiesToConfiguration(destination *v20180601s.Configuration) error {

	// ObjectMeta
	destination.ObjectMeta = *configuration.ObjectMeta.DeepCopy()

	// Spec
	var spec v20180601s.ServersConfigurations_Spec
	err := configuration.Spec.AssignPropertiesToServersConfigurationsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToServersConfigurationsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20180601s.Configuration_Status
	err = configuration.Status.AssignPropertiesToConfigurationStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToConfigurationStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (configuration *Configuration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: configuration.Spec.OriginalVersion(),
		Kind:    "Configuration",
	}
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2018-06-01/Microsoft.DBforMariaDB.json#/resourceDefinitions/servers_configurations
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// +kubebuilder:validation:Enum={"2018-06-01"}
type APIVersion string

const APIVersionValue = APIVersion("2018-06-01")

type Configuration_Status struct {
	// AllowedValues: Allowed values of the configuration.
	AllowedValues *string `json:"allowedValues,omitempty"`

	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// DataType: Data type of the configuration.
	DataType *string `json:"dataType,omitempty"`

	// DefaultValue: Default value of the configuration.
	DefaultValue *string `json:"defaultValue,omitempty"`

	// Description: Description of the configuration.
	Description *string `json:"description,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Configuration_Status{}

// ConvertStatusFrom populates our Configuration_Status from the provided source
func (configuration *Configuration_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20180601s.Configuration_Status)
	if ok {
		// Populate our instance from source
		return configuration.AssignPropertiesFromConfigurationStatus(src)
	}

	// Convert to an intermediate form
	src = &v20180601s.Configuration_Status{}
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
	dst, ok := destination.(*v20180601s.Configuration_Status)
	if ok {
		// Populate destination from our instance
		return configuration.AssignPropertiesToConfigurationStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20180601s.Configuration_Status{}
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
func (configuration *Configuration_Status) AssignPropertiesFromConfigurationStatus(source *v20180601s.Configuration_Status) error {

	// AllowedValues
	configuration.AllowedValues = genruntime.ClonePointerToString(source.AllowedValues)

	// Conditions
	configuration.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// DataType
	configuration.DataType = genruntime.ClonePointerToString(source.DataType)

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

	// Type
	configuration.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	configuration.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignPropertiesToConfigurationStatus populates the provided destination Configuration_Status from our Configuration_Status
func (configuration *Configuration_Status) AssignPropertiesToConfigurationStatus(destination *v20180601s.Configuration_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AllowedValues
	destination.AllowedValues = genruntime.ClonePointerToString(configuration.AllowedValues)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(configuration.Conditions)

	// DataType
	destination.DataType = genruntime.ClonePointerToString(configuration.DataType)

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

type ServersConfigurations_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformariadb.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"dbformariadb.azure.com" json:"owner,omitempty" kind:"Server"`

	// Source: Source of the configuration.
	Source *string `json:"source,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ARMTransformer = &ServersConfigurations_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (configurations *ServersConfigurations_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if configurations == nil {
		return nil, nil
	}
	var result ServersConfigurations_SpecARM

	// Set property ‘Location’:
	if configurations.Location != nil {
		location := *configurations.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if configurations.Source != nil || configurations.Value != nil {
		result.Properties = &ConfigurationPropertiesARM{}
	}
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
func (configurations *ServersConfigurations_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ServersConfigurations_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configurations *ServersConfigurations_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ServersConfigurations_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ServersConfigurations_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	configurations.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		configurations.Location = &location
	}

	// Set property ‘Owner’:
	configurations.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Source’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Source != nil {
			source := *typedInput.Properties.Source
			configurations.Source = &source
		}
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
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			configurations.Value = &value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &ServersConfigurations_Spec{}

// ConvertSpecFrom populates our ServersConfigurations_Spec from the provided source
func (configurations *ServersConfigurations_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20180601s.ServersConfigurations_Spec)
	if ok {
		// Populate our instance from source
		return configurations.AssignPropertiesFromServersConfigurationsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20180601s.ServersConfigurations_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = configurations.AssignPropertiesFromServersConfigurationsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our ServersConfigurations_Spec
func (configurations *ServersConfigurations_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20180601s.ServersConfigurations_Spec)
	if ok {
		// Populate destination from our instance
		return configurations.AssignPropertiesToServersConfigurationsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20180601s.ServersConfigurations_Spec{}
	err := configurations.AssignPropertiesToServersConfigurationsSpec(dst)
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

// AssignPropertiesFromServersConfigurationsSpec populates our ServersConfigurations_Spec from the provided source ServersConfigurations_Spec
func (configurations *ServersConfigurations_Spec) AssignPropertiesFromServersConfigurationsSpec(source *v20180601s.ServersConfigurations_Spec) error {

	// AzureName
	configurations.AzureName = source.AzureName

	// Location
	configurations.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		configurations.Owner = &owner
	} else {
		configurations.Owner = nil
	}

	// Source
	configurations.Source = genruntime.ClonePointerToString(source.Source)

	// Tags
	configurations.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Value
	configurations.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignPropertiesToServersConfigurationsSpec populates the provided destination ServersConfigurations_Spec from our ServersConfigurations_Spec
func (configurations *ServersConfigurations_Spec) AssignPropertiesToServersConfigurationsSpec(destination *v20180601s.ServersConfigurations_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = configurations.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(configurations.Location)

	// OriginalVersion
	destination.OriginalVersion = configurations.OriginalVersion()

	// Owner
	if configurations.Owner != nil {
		owner := configurations.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

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
func (configurations *ServersConfigurations_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (configurations *ServersConfigurations_Spec) SetAzureName(azureName string) {
	configurations.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
