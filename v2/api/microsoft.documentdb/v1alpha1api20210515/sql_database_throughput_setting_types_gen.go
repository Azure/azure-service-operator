// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515storage"
	"github.com/Azure/azure-service-operator/v2/internal/controller/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources={sqldatabasethroughputsettings/status,sqldatabasethroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings
type SqlDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                 `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return sqlDatabaseThroughputSetting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	sqlDatabaseThroughputSetting.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-microsoft-documentdb-azure-com-v1alpha1api20210515-sqldatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabasethroughputsettings.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabaseThroughputSetting{}

// Default applies defaults to the SqlDatabaseThroughputSetting resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) Default() {
	sqlDatabaseThroughputSetting.defaultImpl()
	var temp interface{} = sqlDatabaseThroughputSetting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseThroughputSetting resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) defaultImpl() {}

var _ genruntime.KubernetesResource = &SqlDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetResourceKind returns the kind of the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &sqlDatabaseThroughputSetting.Spec
}

// GetStatus returns the status of this resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &sqlDatabaseThroughputSetting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(sqlDatabaseThroughputSetting.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: sqlDatabaseThroughputSetting.Namespace,
		Name:      sqlDatabaseThroughputSetting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		sqlDatabaseThroughputSetting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	sqlDatabaseThroughputSetting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-microsoft-documentdb-azure-com-v1alpha1api20210515-sqldatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabasethroughputsettings.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) ValidateCreate() error {
	validations := sqlDatabaseThroughputSetting.createValidations()
	var temp interface{} = sqlDatabaseThroughputSetting
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
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) ValidateDelete() error {
	validations := sqlDatabaseThroughputSetting.deleteValidations()
	var temp interface{} = sqlDatabaseThroughputSetting
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
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := sqlDatabaseThroughputSetting.updateValidations()
	var temp interface{} = sqlDatabaseThroughputSetting
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
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) createValidations() []func() error {
	return []func() error{sqlDatabaseThroughputSetting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return sqlDatabaseThroughputSetting.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&sqlDatabaseThroughputSetting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabaseThroughputSetting populates our SqlDatabaseThroughputSetting from the provided source SqlDatabaseThroughputSetting
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) AssignPropertiesFromSqlDatabaseThroughputSetting(source *v1alpha1api20210515storage.SqlDatabaseThroughputSetting) error {

	// Spec
	var spec DatabaseAccountsSqlDatabasesThroughputSettings_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec()")
	}
	sqlDatabaseThroughputSetting.Spec = spec

	// Status
	var status ThroughputSettingsGetResults_Status
	err = status.AssignPropertiesFromThroughputSettingsGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromThroughputSettingsGetResultsStatus()")
	}
	sqlDatabaseThroughputSetting.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseThroughputSetting populates the provided destination SqlDatabaseThroughputSetting from our SqlDatabaseThroughputSetting
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) AssignPropertiesToSqlDatabaseThroughputSetting(destination *v1alpha1api20210515storage.SqlDatabaseThroughputSetting) error {

	// Spec
	var spec v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec
	err := sqlDatabaseThroughputSetting.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210515storage.ThroughputSettingsGetResults_Status
	err = sqlDatabaseThroughputSetting.Status.AssignPropertiesToThroughputSettingsGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToThroughputSettingsGetResultsStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: sqlDatabaseThroughputSetting.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings
type SqlDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseThroughputSetting `json:"items"`
}

type DatabaseAccountsSqlDatabasesThroughputSettings_Spec struct {
	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"SqlDatabase"`

	// +kubebuilder:validation:Required
	//Resource: Cosmos DB resource throughput object. Either throughput is required or
	//autoscaleSettings is required, but not both.
	Resource ThroughputSettingsResource `json:"resource"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags
	//can be used in viewing and grouping this resource (across resource groups). A
	//maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For
	//example, the default experience for a template type is set with
	//"defaultExperience": "Cassandra". Current "defaultExperience" values also
	//include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesThroughputSettings_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if databaseAccountsSqlDatabasesThroughputSettingsSpec == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM

	// Set property ‘APIVersion’:
	result.APIVersion = DatabaseAccountsSqlDatabasesThroughputSettingsSpecAPIVersion20210515

	// Set property ‘Location’:
	if databaseAccountsSqlDatabasesThroughputSettingsSpec.Location != nil {
		location := *databaseAccountsSqlDatabasesThroughputSettingsSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	resourceARM, err := databaseAccountsSqlDatabasesThroughputSettingsSpec.Resource.ConvertToARM(resolved)
	if err != nil {
		return nil, err
	}
	result.Properties.Resource = resourceARM.(ThroughputSettingsResourceARM)

	// Set property ‘Tags’:
	if databaseAccountsSqlDatabasesThroughputSettingsSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range databaseAccountsSqlDatabasesThroughputSettingsSpec.Tags {
			result.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	result.Type = DatabaseAccountsSqlDatabasesThroughputSettingsSpecTypeMicrosoftDocumentDBDatabaseAccountsSqlDatabasesThroughputSettings
	return result, nil
}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) CreateEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesThroughputSettings_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		databaseAccountsSqlDatabasesThroughputSettingsSpec.Location = &location
	}

	// Set property ‘Owner’:
	databaseAccountsSqlDatabasesThroughputSettingsSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	var resource ThroughputSettingsResource
	err := resource.PopulateFromARM(owner, typedInput.Properties.Resource)
	if err != nil {
		return err
	}
	databaseAccountsSqlDatabasesThroughputSettingsSpec.Resource = resource

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		databaseAccountsSqlDatabasesThroughputSettingsSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			databaseAccountsSqlDatabasesThroughputSettingsSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesThroughputSettings_Spec from the provided source
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec)
	if ok {
		// Populate our instance from source
		return databaseAccountsSqlDatabasesThroughputSettingsSpec.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databaseAccountsSqlDatabasesThroughputSettingsSpec.AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesThroughputSettings_Spec
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec)
	if ok {
		// Populate destination from our instance
		return databaseAccountsSqlDatabasesThroughputSettingsSpec.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec{}
	err := databaseAccountsSqlDatabasesThroughputSettingsSpec.AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec populates our DatabaseAccountsSqlDatabasesThroughputSettings_Spec from the provided source DatabaseAccountsSqlDatabasesThroughputSettings_Spec
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesThroughputSettingsSpec(source *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec) error {

	// Location
	if source.Location != nil {
		location := *source.Location
		databaseAccountsSqlDatabasesThroughputSettingsSpec.Location = &location
	} else {
		databaseAccountsSqlDatabasesThroughputSettingsSpec.Location = nil
	}

	// Owner
	databaseAccountsSqlDatabasesThroughputSettingsSpec.Owner = source.Owner.Copy()

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignPropertiesFromThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesFromThroughputSettingsResource()")
		}
		databaseAccountsSqlDatabasesThroughputSettingsSpec.Resource = resource
	} else {
		databaseAccountsSqlDatabasesThroughputSettingsSpec.Resource = ThroughputSettingsResource{}
	}

	// Tags
	tagMap := make(map[string]string)
	for tagKey, tagValue := range source.Tags {
		// Shadow the loop variable to avoid aliasing
		tagValue := tagValue
		tagMap[tagKey] = tagValue
	}
	databaseAccountsSqlDatabasesThroughputSettingsSpec.Tags = tagMap

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec populates the provided destination DatabaseAccountsSqlDatabasesThroughputSettings_Spec from our DatabaseAccountsSqlDatabasesThroughputSettings_Spec
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesThroughputSettingsSpec(destination *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesThroughputSettings_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	if databaseAccountsSqlDatabasesThroughputSettingsSpec.Location != nil {
		location := *databaseAccountsSqlDatabasesThroughputSettingsSpec.Location
		destination.Location = &location
	} else {
		destination.Location = nil
	}

	// OriginalVersion
	destination.OriginalVersion = databaseAccountsSqlDatabasesThroughputSettingsSpec.OriginalVersion()

	// Owner
	destination.Owner = databaseAccountsSqlDatabasesThroughputSettingsSpec.Owner.Copy()

	// Resource
	var resource v1alpha1api20210515storage.ThroughputSettingsResource
	err := databaseAccountsSqlDatabasesThroughputSettingsSpec.Resource.AssignPropertiesToThroughputSettingsResource(&resource)
	if err != nil {
		return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesToThroughputSettingsResource()")
	}
	destination.Resource = &resource

	// Tags
	tagMap := make(map[string]string)
	for tagKey, tagValue := range databaseAccountsSqlDatabasesThroughputSettingsSpec.Tags {
		// Shadow the loop variable to avoid aliasing
		tagValue := tagValue
		tagMap[tagKey] = tagValue
	}
	destination.Tags = tagMap

	// Update the property bag
	destination.PropertyBag = propertyBag

	// No error
	return nil
}

func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseThroughputSetting{}, &SqlDatabaseThroughputSettingList{})
}
