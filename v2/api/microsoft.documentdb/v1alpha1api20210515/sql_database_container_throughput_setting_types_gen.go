// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515storage"
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

// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources=sqldatabasecontainerthroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources={sqldatabasecontainerthroughputsettings/status,sqldatabasecontainerthroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_throughputSettings
type SqlDatabaseContainerThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                           `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerThroughputSetting{}

// GetConditions returns the conditions of the resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) GetConditions() conditions.Conditions {
	return sqlDatabaseContainerThroughputSetting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) SetConditions(conditions conditions.Conditions) {
	sqlDatabaseContainerThroughputSetting.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-microsoft-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerthroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=sqldatabasecontainerthroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabasecontainerthroughputsettings.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabaseContainerThroughputSetting{}

// Default applies defaults to the SqlDatabaseContainerThroughputSetting resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) Default() {
	sqlDatabaseContainerThroughputSetting.defaultImpl()
	var temp interface{} = sqlDatabaseContainerThroughputSetting
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseContainerThroughputSetting resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) defaultImpl() {}

var _ genruntime.KubernetesResource = &SqlDatabaseContainerThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (sqlDatabaseContainerThroughputSetting SqlDatabaseContainerThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &sqlDatabaseContainerThroughputSetting.Spec
}

// GetStatus returns the status of this resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &sqlDatabaseContainerThroughputSetting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(sqlDatabaseContainerThroughputSetting.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: sqlDatabaseContainerThroughputSetting.Namespace,
		Name:      sqlDatabaseContainerThroughputSetting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		sqlDatabaseContainerThroughputSetting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	sqlDatabaseContainerThroughputSetting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-microsoft-documentdb-azure-com-v1alpha1api20210515-sqldatabasecontainerthroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=sqldatabasecontainerthroughputsettings,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabasecontainerthroughputsettings.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabaseContainerThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) ValidateCreate() error {
	validations := sqlDatabaseContainerThroughputSetting.createValidations()
	var temp interface{} = sqlDatabaseContainerThroughputSetting
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
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) ValidateDelete() error {
	validations := sqlDatabaseContainerThroughputSetting.deleteValidations()
	var temp interface{} = sqlDatabaseContainerThroughputSetting
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
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) ValidateUpdate(old runtime.Object) error {
	validations := sqlDatabaseContainerThroughputSetting.updateValidations()
	var temp interface{} = sqlDatabaseContainerThroughputSetting
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
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) createValidations() []func() error {
	return []func() error{sqlDatabaseContainerThroughputSetting.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return sqlDatabaseContainerThroughputSetting.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&sqlDatabaseContainerThroughputSetting.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabaseContainerThroughputSetting populates our SqlDatabaseContainerThroughputSetting from the provided source SqlDatabaseContainerThroughputSetting
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) AssignPropertiesFromSqlDatabaseContainerThroughputSetting(source *v1alpha1api20210515storage.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	sqlDatabaseContainerThroughputSetting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec()")
	}
	sqlDatabaseContainerThroughputSetting.Spec = spec

	// Status
	var status ThroughputSettingsGetResults_Status
	err = status.AssignPropertiesFromThroughputSettingsGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromThroughputSettingsGetResultsStatus()")
	}
	sqlDatabaseContainerThroughputSetting.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerThroughputSetting populates the provided destination SqlDatabaseContainerThroughputSetting from our SqlDatabaseContainerThroughputSetting
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) AssignPropertiesToSqlDatabaseContainerThroughputSetting(destination *v1alpha1api20210515storage.SqlDatabaseContainerThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *sqlDatabaseContainerThroughputSetting.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
	err := sqlDatabaseContainerThroughputSetting.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210515storage.ThroughputSettingsGetResults_Status
	err = sqlDatabaseContainerThroughputSetting.Status.AssignPropertiesToThroughputSettingsGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToThroughputSettingsGetResultsStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (sqlDatabaseContainerThroughputSetting *SqlDatabaseContainerThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: sqlDatabaseContainerThroughputSetting.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerThroughputSetting",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_throughputSettings
type SqlDatabaseContainerThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerThroughputSetting `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecAPIVersion string

const DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecAPIVersion20210515 = DatabaseAccountsSqlDatabasesContainersThroughputSettingsSpecAPIVersion("2021-05-15")

type DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec struct {
	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"SqlDatabaseContainer"`

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

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if databaseAccountsSqlDatabasesContainersThroughputSettingsSpec == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM

	// Set property ‘Location’:
	if databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Location != nil {
		location := *databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	resourceARM, err := databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Resource.ConvertToARM(resolved)
	if err != nil {
		return nil, err
	}
	result.Properties.Resource = resourceARM.(ThroughputSettingsResourceARM)

	// Set property ‘Tags’:
	if databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesContainersThroughputSettings_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Location = &location
	}

	// Set property ‘Owner’:
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	var resource ThroughputSettingsResource
	err := resource.PopulateFromARM(owner, typedInput.Properties.Resource)
	if err != nil {
		return err
	}
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Resource = resource

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec from the provided source
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec)
	if ok {
		// Populate our instance from source
		return databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec)
	if ok {
		// Populate destination from our instance
		return databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{}
	err := databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec populates our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec from the provided source DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(source *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) error {

	// Location
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Owner = source.Owner.Copy()

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignPropertiesFromThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesFromThroughputSettingsResource()")
		}
		databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Resource = resource
	} else {
		databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Resource = ThroughputSettingsResource{}
	}

	// Tags
	databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec populates the provided destination DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec from our DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersThroughputSettingsSpec(destination *v1alpha1api20210515storage.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.OriginalVersion()

	// Owner
	destination.Owner = databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Owner.Copy()

	// Resource
	var resource v1alpha1api20210515storage.ThroughputSettingsResource
	err := databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Resource.AssignPropertiesToThroughputSettingsResource(&resource)
	if err != nil {
		return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesToThroughputSettingsResource()")
	}
	destination.Resource = &resource

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(databaseAccountsSqlDatabasesContainersThroughputSettingsSpec.Tags)

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
func (databaseAccountsSqlDatabasesContainersThroughputSettingsSpec *DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseContainerThroughputSetting{}, &SqlDatabaseContainerThroughputSettingList{})
}
