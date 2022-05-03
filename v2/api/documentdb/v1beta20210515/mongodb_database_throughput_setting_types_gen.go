// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210515

import (
	"fmt"
	v20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515storage"
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
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsMongodbDatabasesThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                     `json:"status,omitempty"`
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
	source, ok := hub.(*v20210515s.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignPropertiesFromMongodbDatabaseThroughputSetting(source)
}

// ConvertTo populates the provided hub MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.MongodbDatabaseThroughputSetting)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/MongodbDatabaseThroughputSetting but received %T instead", hub)
	}

	return setting.AssignPropertiesToMongodbDatabaseThroughputSetting(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1beta20210515-mongodbdatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1beta20210515,name=default.v1beta20210515.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &MongodbDatabaseThroughputSetting{}

// Default applies defaults to the MongodbDatabaseThroughputSetting resource
func (setting *MongodbDatabaseThroughputSetting) Default() {
	setting.defaultImpl()
	var temp interface{} = setting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the MongodbDatabaseThroughputSetting resource
func (setting *MongodbDatabaseThroughputSetting) defaultImpl() {}

var _ genruntime.KubernetesResource = &MongodbDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (setting *MongodbDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (setting MongodbDatabaseThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (setting *MongodbDatabaseThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (setting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
func (setting *MongodbDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (setting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  setting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (setting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1beta20210515-mongodbdatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1beta20210515,name=validate.v1beta20210515.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &MongodbDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateCreate() error {
	validations := setting.createValidations()
	var temp interface{} = setting
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
func (setting *MongodbDatabaseThroughputSetting) ValidateDelete() error {
	validations := setting.deleteValidations()
	var temp interface{} = setting
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
func (setting *MongodbDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := setting.updateValidations()
	var temp interface{} = setting
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
func (setting *MongodbDatabaseThroughputSetting) createValidations() []func() error {
	return []func() error{setting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (setting *MongodbDatabaseThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (setting *MongodbDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return setting.validateResourceReferences()
		},
		setting.validateImmutableProperties}
}

// validateImmutableProperties validates all immutable properties
func (setting *MongodbDatabaseThroughputSetting) validateImmutableProperties(old runtime.Object) error {
	oldObj, ok := old.(*MongodbDatabaseThroughputSetting)
	if !ok {
		return nil
	}
	return genruntime.ValidateImmutableProperties(oldObj, setting)
}

// validateResourceReferences validates all resource references
func (setting *MongodbDatabaseThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromMongodbDatabaseThroughputSetting populates our MongodbDatabaseThroughputSetting from the provided source MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignPropertiesFromMongodbDatabaseThroughputSetting(source *v20210515s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status ThroughputSettingsGetResults_Status
	err = status.AssignPropertiesFromThroughputSettingsGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsGetResultsStatus() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignPropertiesToMongodbDatabaseThroughputSetting populates the provided destination MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignPropertiesToMongodbDatabaseThroughputSetting(destination *v20210515s.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
	err := setting.Spec.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.ThroughputSettingsGetResults_Status
	err = setting.Status.AssignPropertiesToThroughputSettingsGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsGetResultsStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (setting *MongodbDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: setting.Spec.OriginalVersion(),
		Kind:    "MongodbDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_mongodbDatabases_throughputSettings
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion string

const DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion20210515 = DatabaseAccountsMongodbDatabasesThroughputSettingsSpecAPIVersion("2021-05-15")

type DatabaseAccountsMongodbDatabasesThroughputSettings_Spec struct {
	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a documentdb.azure.com/MongodbDatabase resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"MongodbDatabase"`

	// +kubebuilder:validation:Required
	//Resource: Cosmos DB resource throughput object. Either throughput is required or autoscaleSettings is required, but not
	//both.
	Resource *ThroughputSettingsResource `json:"resource,omitempty"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
	//resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
	//type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph",
	//"DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if settings == nil {
		return nil, nil
	}
	var result DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM

	// Set property ‘Location’:
	if settings.Location != nil {
		location := *settings.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if settings.Resource != nil {
		result.Properties = &ThroughputSettingsUpdatePropertiesARM{}
	}
	if settings.Resource != nil {
		resourceARM, err := (*settings.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := resourceARM.(ThroughputSettingsResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if settings.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range settings.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsMongodbDatabasesThroughputSettings_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		settings.Location = &location
	}

	// Set property ‘Owner’:
	settings.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			settings.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		settings.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			settings.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec from the provided source
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec)
	if ok {
		// Populate our instance from source
		return settings.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = settings.AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec)
	if ok {
		// Populate destination from our instance
		return settings.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{}
	err := settings.AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec populates our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec from the provided source DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) AssignPropertiesFromDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(source *v20210515s.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) error {

	// Location
	settings.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		settings.Owner = &owner
	} else {
		settings.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignPropertiesFromThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromThroughputSettingsResource() to populate field Resource")
		}
		settings.Resource = &resource
	} else {
		settings.Resource = nil
	}

	// Tags
	settings.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec populates the provided destination DatabaseAccountsMongodbDatabasesThroughputSettings_Spec from our DatabaseAccountsMongodbDatabasesThroughputSettings_Spec
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) AssignPropertiesToDatabaseAccountsMongodbDatabasesThroughputSettingsSpec(destination *v20210515s.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(settings.Location)

	// OriginalVersion
	destination.OriginalVersion = settings.OriginalVersion()

	// Owner
	if settings.Owner != nil {
		owner := settings.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if settings.Resource != nil {
		var resource v20210515s.ThroughputSettingsResource
		err := settings.Resource.AssignPropertiesToThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(settings.Tags)

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
func (settings *DatabaseAccountsMongodbDatabasesThroughputSettings_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}
