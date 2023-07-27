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
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of Configuration. Use v1api20180601.Configuration instead
type Configuration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_Configuration_Spec   `json:"spec,omitempty"`
	Status            Servers_Configuration_STATUS `json:"status,omitempty"`
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
	// intermediate variable for conversion
	var source v20180601s.Configuration

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = configuration.AssignProperties_From_Configuration(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to configuration")
	}

	return nil
}

// ConvertTo populates the provided hub Configuration from our Configuration
func (configuration *Configuration) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20180601s.Configuration
	err := configuration.AssignProperties_To_Configuration(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from configuration")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-dbformariadb-azure-com-v1beta20180601-configuration,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=configurations,verbs=create;update,versions=v1beta20180601,name=default.v1beta20180601.configurations.dbformariadb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &Configuration{}

// Default applies defaults to the Configuration resource
func (configuration *Configuration) Default() {
	configuration.defaultImpl()
	var temp any = configuration
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
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (configuration *Configuration) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
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
	return &Servers_Configuration_STATUS{}
}

// Owner returns the ResourceReference of the owner
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
	if st, ok := status.(*Servers_Configuration_STATUS); ok {
		configuration.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_Configuration_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	configuration.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformariadb-azure-com-v1beta20180601-configuration,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformariadb.azure.com,resources=configurations,verbs=create;update,versions=v1beta20180601,name=validate.v1beta20180601.configurations.dbformariadb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &Configuration{}

// ValidateCreate validates the creation of the resource
func (configuration *Configuration) ValidateCreate() (admission.Warnings, error) {
	validations := configuration.createValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (configuration *Configuration) ValidateDelete() (admission.Warnings, error) {
	validations := configuration.deleteValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (configuration *Configuration) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := configuration.updateValidations()
	var temp any = configuration
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (configuration *Configuration) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){configuration.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (configuration *Configuration) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (configuration *Configuration) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return configuration.validateResourceReferences()
		},
		configuration.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (configuration *Configuration) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&configuration.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (configuration *Configuration) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*Configuration)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, configuration)
}

// AssignProperties_From_Configuration populates our Configuration from the provided source Configuration
func (configuration *Configuration) AssignProperties_From_Configuration(source *v20180601s.Configuration) error {

	// ObjectMeta
	configuration.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_Configuration_Spec
	err := spec.AssignProperties_From_Servers_Configuration_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Configuration_Spec() to populate field Spec")
	}
	configuration.Spec = spec

	// Status
	var status Servers_Configuration_STATUS
	err = status.AssignProperties_From_Servers_Configuration_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_Configuration_STATUS() to populate field Status")
	}
	configuration.Status = status

	// No error
	return nil
}

// AssignProperties_To_Configuration populates the provided destination Configuration from our Configuration
func (configuration *Configuration) AssignProperties_To_Configuration(destination *v20180601s.Configuration) error {

	// ObjectMeta
	destination.ObjectMeta = *configuration.ObjectMeta.DeepCopy()

	// Spec
	var spec v20180601s.Servers_Configuration_Spec
	err := configuration.Spec.AssignProperties_To_Servers_Configuration_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Configuration_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20180601s.Servers_Configuration_STATUS
	err = configuration.Status.AssignProperties_To_Servers_Configuration_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_Configuration_STATUS() to populate field Status")
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
// Deprecated version of Configuration. Use v1api20180601.Configuration instead
type ConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Configuration `json:"items"`
}

// Deprecated version of APIVersion. Use v1api20180601.APIVersion instead
// +kubebuilder:validation:Enum={"2018-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2018-06-01")

type Servers_Configuration_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformariadb.azure.com/Server resource
	Owner  *genruntime.KnownResourceReference `group:"dbformariadb.azure.com" json:"owner,omitempty" kind:"Server"`
	Source *string                            `json:"source,omitempty"`
	Value  *string                            `json:"value,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_Configuration_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (configuration *Servers_Configuration_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if configuration == nil {
		return nil, nil
	}
	result := &Servers_Configuration_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if configuration.Source != nil || configuration.Value != nil {
		result.Properties = &ConfigurationProperties_ARM{}
	}
	if configuration.Source != nil {
		source := *configuration.Source
		result.Properties.Source = &source
	}
	if configuration.Value != nil {
		value := *configuration.Value
		result.Properties.Value = &value
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (configuration *Servers_Configuration_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Configuration_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configuration *Servers_Configuration_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Configuration_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Configuration_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	configuration.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Owner":
	configuration.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property "Source":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Source != nil {
			source := *typedInput.Properties.Source
			configuration.Source = &source
		}
	}

	// Set property "Value":
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

var _ genruntime.ConvertibleSpec = &Servers_Configuration_Spec{}

// ConvertSpecFrom populates our Servers_Configuration_Spec from the provided source
func (configuration *Servers_Configuration_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20180601s.Servers_Configuration_Spec)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_Servers_Configuration_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20180601s.Servers_Configuration_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_Servers_Configuration_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_Configuration_Spec
func (configuration *Servers_Configuration_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20180601s.Servers_Configuration_Spec)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_Servers_Configuration_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20180601s.Servers_Configuration_Spec{}
	err := configuration.AssignProperties_To_Servers_Configuration_Spec(dst)
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

// AssignProperties_From_Servers_Configuration_Spec populates our Servers_Configuration_Spec from the provided source Servers_Configuration_Spec
func (configuration *Servers_Configuration_Spec) AssignProperties_From_Servers_Configuration_Spec(source *v20180601s.Servers_Configuration_Spec) error {

	// AzureName
	configuration.AzureName = source.AzureName

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

	// No error
	return nil
}

// AssignProperties_To_Servers_Configuration_Spec populates the provided destination Servers_Configuration_Spec from our Servers_Configuration_Spec
func (configuration *Servers_Configuration_Spec) AssignProperties_To_Servers_Configuration_Spec(destination *v20180601s.Servers_Configuration_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = configuration.AzureName

	// OriginalVersion
	destination.OriginalVersion = configuration.OriginalVersion()

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

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (configuration *Servers_Configuration_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (configuration *Servers_Configuration_Spec) SetAzureName(azureName string) {
	configuration.AzureName = azureName
}

// Deprecated version of Servers_Configuration_STATUS. Use v1api20180601.Servers_Configuration_STATUS instead
type Servers_Configuration_STATUS struct {
	AllowedValues *string `json:"allowedValues,omitempty"`

	// Conditions: The observed state of the resource
	Conditions   []conditions.Condition `json:"conditions,omitempty"`
	DataType     *string                `json:"dataType,omitempty"`
	DefaultValue *string                `json:"defaultValue,omitempty"`
	Description  *string                `json:"description,omitempty"`
	Id           *string                `json:"id,omitempty"`
	Name         *string                `json:"name,omitempty"`
	Source       *string                `json:"source,omitempty"`
	Type         *string                `json:"type,omitempty"`
	Value        *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_Configuration_STATUS{}

// ConvertStatusFrom populates our Servers_Configuration_STATUS from the provided source
func (configuration *Servers_Configuration_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20180601s.Servers_Configuration_STATUS)
	if ok {
		// Populate our instance from source
		return configuration.AssignProperties_From_Servers_Configuration_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20180601s.Servers_Configuration_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = configuration.AssignProperties_From_Servers_Configuration_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_Configuration_STATUS
func (configuration *Servers_Configuration_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20180601s.Servers_Configuration_STATUS)
	if ok {
		// Populate destination from our instance
		return configuration.AssignProperties_To_Servers_Configuration_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20180601s.Servers_Configuration_STATUS{}
	err := configuration.AssignProperties_To_Servers_Configuration_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_Configuration_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (configuration *Servers_Configuration_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_Configuration_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (configuration *Servers_Configuration_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_Configuration_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_Configuration_STATUS_ARM, got %T", armInput)
	}

	// Set property "AllowedValues":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.AllowedValues != nil {
			allowedValues := *typedInput.Properties.AllowedValues
			configuration.AllowedValues = &allowedValues
		}
	}

	// no assignment for property "Conditions"

	// Set property "DataType":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DataType != nil {
			dataType := *typedInput.Properties.DataType
			configuration.DataType = &dataType
		}
	}

	// Set property "DefaultValue":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.DefaultValue != nil {
			defaultValue := *typedInput.Properties.DefaultValue
			configuration.DefaultValue = &defaultValue
		}
	}

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			configuration.Description = &description
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		configuration.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		configuration.Name = &name
	}

	// Set property "Source":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Source != nil {
			source := *typedInput.Properties.Source
			configuration.Source = &source
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		configuration.Type = &typeVar
	}

	// Set property "Value":
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

// AssignProperties_From_Servers_Configuration_STATUS populates our Servers_Configuration_STATUS from the provided source Servers_Configuration_STATUS
func (configuration *Servers_Configuration_STATUS) AssignProperties_From_Servers_Configuration_STATUS(source *v20180601s.Servers_Configuration_STATUS) error {

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

// AssignProperties_To_Servers_Configuration_STATUS populates the provided destination Servers_Configuration_STATUS from our Servers_Configuration_STATUS
func (configuration *Servers_Configuration_STATUS) AssignProperties_To_Servers_Configuration_STATUS(destination *v20180601s.Servers_Configuration_STATUS) error {
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

func init() {
	SchemeBuilder.Register(&Configuration{}, &ConfigurationList{})
}
