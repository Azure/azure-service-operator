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

// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources=sqldatabases,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=microsoft.documentdb.azure.com,resources={sqldatabases/status,sqldatabases/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases
type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabases_Spec `json:"spec,omitempty"`
	Status            SqlDatabaseGetResults_Status      `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabase{}

// GetConditions returns the conditions of the resource
func (sqlDatabase *SqlDatabase) GetConditions() conditions.Conditions {
	return sqlDatabase.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (sqlDatabase *SqlDatabase) SetConditions(conditions conditions.Conditions) {
	sqlDatabase.Status.Conditions = conditions
}

// +kubebuilder:webhook:path=/mutate-microsoft-documentdb-azure-com-v1alpha1api20210515-sqldatabase,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=sqldatabases,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabases.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabase{}

// Default applies defaults to the SqlDatabase resource
func (sqlDatabase *SqlDatabase) Default() {
	sqlDatabase.defaultImpl()
	var temp interface{} = sqlDatabase
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (sqlDatabase *SqlDatabase) defaultAzureName() {
	if sqlDatabase.Spec.AzureName == "" {
		sqlDatabase.Spec.AzureName = sqlDatabase.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabase resource
func (sqlDatabase *SqlDatabase) defaultImpl() { sqlDatabase.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabase{}

// AzureName returns the Azure name of the resource
func (sqlDatabase *SqlDatabase) AzureName() string {
	return sqlDatabase.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (sqlDatabase *SqlDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (sqlDatabase *SqlDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &sqlDatabase.Spec
}

// GetStatus returns the status of this resource
func (sqlDatabase *SqlDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &sqlDatabase.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
func (sqlDatabase *SqlDatabase) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
}

// NewEmptyStatus returns a new empty (blank) status
func (sqlDatabase *SqlDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlDatabaseGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (sqlDatabase *SqlDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(sqlDatabase.Spec)
	return &genruntime.ResourceReference{
		Group:     group,
		Kind:      kind,
		Namespace: sqlDatabase.Namespace,
		Name:      sqlDatabase.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (sqlDatabase *SqlDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlDatabaseGetResults_Status); ok {
		sqlDatabase.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlDatabaseGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	sqlDatabase.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-microsoft-documentdb-azure-com-v1alpha1api20210515-sqldatabase,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=microsoft.documentdb.azure.com,resources=sqldatabases,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabases.microsoft.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabase{}

// ValidateCreate validates the creation of the resource
func (sqlDatabase *SqlDatabase) ValidateCreate() error {
	validations := sqlDatabase.createValidations()
	var temp interface{} = sqlDatabase
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
func (sqlDatabase *SqlDatabase) ValidateDelete() error {
	validations := sqlDatabase.deleteValidations()
	var temp interface{} = sqlDatabase
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
func (sqlDatabase *SqlDatabase) ValidateUpdate(old runtime.Object) error {
	validations := sqlDatabase.updateValidations()
	var temp interface{} = sqlDatabase
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
func (sqlDatabase *SqlDatabase) createValidations() []func() error {
	return []func() error{sqlDatabase.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (sqlDatabase *SqlDatabase) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (sqlDatabase *SqlDatabase) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return sqlDatabase.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (sqlDatabase *SqlDatabase) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&sqlDatabase.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabase populates our SqlDatabase from the provided source SqlDatabase
func (sqlDatabase *SqlDatabase) AssignPropertiesFromSqlDatabase(source *v1alpha1api20210515storage.SqlDatabase) error {

	// Spec
	var spec DatabaseAccountsSqlDatabases_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec()")
	}
	sqlDatabase.Spec = spec

	// Status
	var status SqlDatabaseGetResults_Status
	err = status.AssignPropertiesFromSqlDatabaseGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromSqlDatabaseGetResultsStatus()")
	}
	sqlDatabase.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabase populates the provided destination SqlDatabase from our SqlDatabase
func (sqlDatabase *SqlDatabase) AssignPropertiesToSqlDatabase(destination *v1alpha1api20210515storage.SqlDatabase) error {

	// Spec
	var spec v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec
	err := sqlDatabase.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToDatabaseAccountsSqlDatabasesSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20210515storage.SqlDatabaseGetResults_Status
	err = sqlDatabase.Status.AssignPropertiesToSqlDatabaseGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToSqlDatabaseGetResultsStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (sqlDatabase *SqlDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: sqlDatabase.Spec.OriginalVersion(),
		Kind:    "SqlDatabase",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsSqlDatabasesSpecAPIVersion string

const DatabaseAccountsSqlDatabasesSpecAPIVersion20210515 = DatabaseAccountsSqlDatabasesSpecAPIVersion("2021-05-15")

type DatabaseAccountsSqlDatabases_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Options: CreateUpdateOptions are a list of key-value pairs that describe the
	//resource. Supported keys are "If-Match", "If-None-Match", "Session-Token" and
	//"Throughput"
	Options *CreateUpdateOptions `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"microsoft.documentdb.azure.com" json:"owner" kind:"DatabaseAccount"`

	// +kubebuilder:validation:Required
	//Resource: Cosmos DB SQL database resource object
	Resource SqlDatabaseResource `json:"resource"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags
	//can be used in viewing and grouping this resource (across resource groups). A
	//maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For
	//example, the default experience for a template type is set with
	//"defaultExperience": "Cassandra". Current "defaultExperience" values also
	//include "Table", "Graph", "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabases_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if databaseAccountsSqlDatabasesSpec == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabases_SpecARM

	// Set property ‘Location’:
	if databaseAccountsSqlDatabasesSpec.Location != nil {
		location := *databaseAccountsSqlDatabasesSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if databaseAccountsSqlDatabasesSpec.Options != nil {
		optionsARM, err := (*databaseAccountsSqlDatabasesSpec.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := optionsARM.(CreateUpdateOptionsARM)
		result.Properties.Options = &options
	}
	resourceARM, err := databaseAccountsSqlDatabasesSpec.Resource.ConvertToARM(resolved)
	if err != nil {
		return nil, err
	}
	result.Properties.Resource = resourceARM.(SqlDatabaseResourceARM)

	// Set property ‘Tags’:
	if databaseAccountsSqlDatabasesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range databaseAccountsSqlDatabasesSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) CreateEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabases_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabases_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabases_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	databaseAccountsSqlDatabasesSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		databaseAccountsSqlDatabasesSpec.Location = &location
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties.Options != nil {
		var options1 CreateUpdateOptions
		err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
		if err != nil {
			return err
		}
		options := options1
		databaseAccountsSqlDatabasesSpec.Options = &options
	}

	// Set property ‘Owner’:
	databaseAccountsSqlDatabasesSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	var resource SqlDatabaseResource
	err := resource.PopulateFromARM(owner, typedInput.Properties.Resource)
	if err != nil {
		return err
	}
	databaseAccountsSqlDatabasesSpec.Resource = resource

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		databaseAccountsSqlDatabasesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			databaseAccountsSqlDatabasesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabases_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabases_Spec from the provided source
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec)
	if ok {
		// Populate our instance from source
		return databaseAccountsSqlDatabasesSpec.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = databaseAccountsSqlDatabasesSpec.AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabases_Spec
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec)
	if ok {
		// Populate destination from our instance
		return databaseAccountsSqlDatabasesSpec.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec{}
	err := databaseAccountsSqlDatabasesSpec.AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec populates our DatabaseAccountsSqlDatabases_Spec from the provided source DatabaseAccountsSqlDatabases_Spec
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesSpec(source *v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec) error {

	// AzureName
	databaseAccountsSqlDatabasesSpec.AzureName = source.AzureName

	// Location
	databaseAccountsSqlDatabasesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "populating Options from Options, calling AssignPropertiesFromCreateUpdateOptions()")
		}
		databaseAccountsSqlDatabasesSpec.Options = &option
	} else {
		databaseAccountsSqlDatabasesSpec.Options = nil
	}

	// Owner
	databaseAccountsSqlDatabasesSpec.Owner = source.Owner.Copy()

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseResource
		err := resource.AssignPropertiesFromSqlDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesFromSqlDatabaseResource()")
		}
		databaseAccountsSqlDatabasesSpec.Resource = resource
	} else {
		databaseAccountsSqlDatabasesSpec.Resource = SqlDatabaseResource{}
	}

	// Tags
	databaseAccountsSqlDatabasesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesSpec populates the provided destination DatabaseAccountsSqlDatabases_Spec from our DatabaseAccountsSqlDatabases_Spec
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesSpec(destination *v1alpha1api20210515storage.DatabaseAccountsSqlDatabases_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = databaseAccountsSqlDatabasesSpec.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(databaseAccountsSqlDatabasesSpec.Location)

	// Options
	if databaseAccountsSqlDatabasesSpec.Options != nil {
		var option v1alpha1api20210515storage.CreateUpdateOptions
		err := (*databaseAccountsSqlDatabasesSpec.Options).AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "populating Options from Options, calling AssignPropertiesToCreateUpdateOptions()")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = databaseAccountsSqlDatabasesSpec.OriginalVersion()

	// Owner
	destination.Owner = databaseAccountsSqlDatabasesSpec.Owner.Copy()

	// Resource
	var resource v1alpha1api20210515storage.SqlDatabaseResource
	err := databaseAccountsSqlDatabasesSpec.Resource.AssignPropertiesToSqlDatabaseResource(&resource)
	if err != nil {
		return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesToSqlDatabaseResource()")
	}
	destination.Resource = &resource

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(databaseAccountsSqlDatabasesSpec.Tags)

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
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (databaseAccountsSqlDatabasesSpec *DatabaseAccountsSqlDatabases_Spec) SetAzureName(azureName string) {
	databaseAccountsSqlDatabasesSpec.AzureName = azureName
}

//Generated from:
type SqlDatabaseGetResults_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name     *string                                   `json:"name,omitempty"`
	Options  *OptionsResource_Status                   `json:"options,omitempty"`
	Resource *SqlDatabaseGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags     map[string]string                         `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlDatabaseGetResults_Status{}

// ConvertStatusFrom populates our SqlDatabaseGetResults_Status from the provided source
func (sqlDatabaseGetResultsStatus *SqlDatabaseGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20210515storage.SqlDatabaseGetResults_Status)
	if ok {
		// Populate our instance from source
		return sqlDatabaseGetResultsStatus.AssignPropertiesFromSqlDatabaseGetResultsStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20210515storage.SqlDatabaseGetResults_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = sqlDatabaseGetResultsStatus.AssignPropertiesFromSqlDatabaseGetResultsStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlDatabaseGetResults_Status
func (sqlDatabaseGetResultsStatus *SqlDatabaseGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20210515storage.SqlDatabaseGetResults_Status)
	if ok {
		// Populate destination from our instance
		return sqlDatabaseGetResultsStatus.AssignPropertiesToSqlDatabaseGetResultsStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20210515storage.SqlDatabaseGetResults_Status{}
	err := sqlDatabaseGetResultsStatus.AssignPropertiesToSqlDatabaseGetResultsStatus(dst)
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

var _ genruntime.FromARMConverter = &SqlDatabaseGetResults_Status{}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (sqlDatabaseGetResultsStatus *SqlDatabaseGetResults_Status) CreateEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseGetResults_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (sqlDatabaseGetResultsStatus *SqlDatabaseGetResults_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseGetResults_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseGetResults_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		sqlDatabaseGetResultsStatus.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		sqlDatabaseGetResultsStatus.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		sqlDatabaseGetResultsStatus.Name = &name
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 OptionsResource_Status
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			sqlDatabaseGetResultsStatus.Options = &options
		}
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlDatabaseGetProperties_Status_Resource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			sqlDatabaseGetResultsStatus.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		sqlDatabaseGetResultsStatus.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			sqlDatabaseGetResultsStatus.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		sqlDatabaseGetResultsStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseGetResultsStatus populates our SqlDatabaseGetResults_Status from the provided source SqlDatabaseGetResults_Status
func (sqlDatabaseGetResultsStatus *SqlDatabaseGetResults_Status) AssignPropertiesFromSqlDatabaseGetResultsStatus(source *v1alpha1api20210515storage.SqlDatabaseGetResults_Status) error {

	// Conditions
	sqlDatabaseGetResultsStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	sqlDatabaseGetResultsStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	sqlDatabaseGetResultsStatus.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	sqlDatabaseGetResultsStatus.Name = genruntime.ClonePointerToString(source.Name)

	// Options
	if source.Options != nil {
		var option OptionsResource_Status
		err := option.AssignPropertiesFromOptionsResourceStatus(source.Options)
		if err != nil {
			return errors.Wrap(err, "populating Options from Options, calling AssignPropertiesFromOptionsResourceStatus()")
		}
		sqlDatabaseGetResultsStatus.Options = &option
	} else {
		sqlDatabaseGetResultsStatus.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseGetProperties_Status_Resource
		err := resource.AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource()")
		}
		sqlDatabaseGetResultsStatus.Resource = &resource
	} else {
		sqlDatabaseGetResultsStatus.Resource = nil
	}

	// Tags
	sqlDatabaseGetResultsStatus.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	sqlDatabaseGetResultsStatus.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseGetResultsStatus populates the provided destination SqlDatabaseGetResults_Status from our SqlDatabaseGetResults_Status
func (sqlDatabaseGetResultsStatus *SqlDatabaseGetResults_Status) AssignPropertiesToSqlDatabaseGetResultsStatus(destination *v1alpha1api20210515storage.SqlDatabaseGetResults_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(sqlDatabaseGetResultsStatus.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(sqlDatabaseGetResultsStatus.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(sqlDatabaseGetResultsStatus.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(sqlDatabaseGetResultsStatus.Name)

	// Options
	if sqlDatabaseGetResultsStatus.Options != nil {
		var option v1alpha1api20210515storage.OptionsResource_Status
		err := (*sqlDatabaseGetResultsStatus.Options).AssignPropertiesToOptionsResourceStatus(&option)
		if err != nil {
			return errors.Wrap(err, "populating Options from Options, calling AssignPropertiesToOptionsResourceStatus()")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if sqlDatabaseGetResultsStatus.Resource != nil {
		var resource v1alpha1api20210515storage.SqlDatabaseGetProperties_Status_Resource
		err := (*sqlDatabaseGetResultsStatus.Resource).AssignPropertiesToSqlDatabaseGetPropertiesStatusResource(&resource)
		if err != nil {
			return errors.Wrap(err, "populating Resource from Resource, calling AssignPropertiesToSqlDatabaseGetPropertiesStatusResource()")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(sqlDatabaseGetResultsStatus.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(sqlDatabaseGetResultsStatus.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type SqlDatabaseGetProperties_Status_Resource struct {
	//Colls: A system generated property that specified the addressable path of the
	//collections resource.
	Colls *string `json:"_colls,omitempty"`

	//Etag: A system generated property representing the resource etag required for
	//optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	// +kubebuilder:validation:Required
	//Id: Name of the Cosmos DB SQL database
	Id string `json:"id"`

	//Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	//Ts: A system generated property that denotes the last updated timestamp of the
	//resource.
	Ts *float64 `json:"_ts,omitempty"`

	//Users: A system generated property that specifies the addressable path of the
	//users resource.
	Users *string `json:"_users,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlDatabaseGetProperties_Status_Resource{}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (sqlDatabaseGetPropertiesStatusResource *SqlDatabaseGetProperties_Status_Resource) CreateEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseGetProperties_Status_ResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (sqlDatabaseGetPropertiesStatusResource *SqlDatabaseGetProperties_Status_Resource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseGetProperties_Status_ResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseGetProperties_Status_ResourceARM, got %T", armInput)
	}

	// Set property ‘Colls’:
	if typedInput.Colls != nil {
		colls := *typedInput.Colls
		sqlDatabaseGetPropertiesStatusResource.Colls = &colls
	}

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		sqlDatabaseGetPropertiesStatusResource.Etag = &etag
	}

	// Set property ‘Id’:
	sqlDatabaseGetPropertiesStatusResource.Id = typedInput.Id

	// Set property ‘Rid’:
	if typedInput.Rid != nil {
		rid := *typedInput.Rid
		sqlDatabaseGetPropertiesStatusResource.Rid = &rid
	}

	// Set property ‘Ts’:
	if typedInput.Ts != nil {
		ts := *typedInput.Ts
		sqlDatabaseGetPropertiesStatusResource.Ts = &ts
	}

	// Set property ‘Users’:
	if typedInput.Users != nil {
		users := *typedInput.Users
		sqlDatabaseGetPropertiesStatusResource.Users = &users
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource populates our SqlDatabaseGetProperties_Status_Resource from the provided source SqlDatabaseGetProperties_Status_Resource
func (sqlDatabaseGetPropertiesStatusResource *SqlDatabaseGetProperties_Status_Resource) AssignPropertiesFromSqlDatabaseGetPropertiesStatusResource(source *v1alpha1api20210515storage.SqlDatabaseGetProperties_Status_Resource) error {

	// Colls
	sqlDatabaseGetPropertiesStatusResource.Colls = genruntime.ClonePointerToString(source.Colls)

	// Etag
	sqlDatabaseGetPropertiesStatusResource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	sqlDatabaseGetPropertiesStatusResource.Id = genruntime.GetOptionalStringValue(source.Id)

	// Rid
	sqlDatabaseGetPropertiesStatusResource.Rid = genruntime.ClonePointerToString(source.Rid)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		sqlDatabaseGetPropertiesStatusResource.Ts = &t
	} else {
		sqlDatabaseGetPropertiesStatusResource.Ts = nil
	}

	// Users
	sqlDatabaseGetPropertiesStatusResource.Users = genruntime.ClonePointerToString(source.Users)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseGetPropertiesStatusResource populates the provided destination SqlDatabaseGetProperties_Status_Resource from our SqlDatabaseGetProperties_Status_Resource
func (sqlDatabaseGetPropertiesStatusResource *SqlDatabaseGetProperties_Status_Resource) AssignPropertiesToSqlDatabaseGetPropertiesStatusResource(destination *v1alpha1api20210515storage.SqlDatabaseGetProperties_Status_Resource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Colls
	destination.Colls = genruntime.ClonePointerToString(sqlDatabaseGetPropertiesStatusResource.Colls)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(sqlDatabaseGetPropertiesStatusResource.Etag)

	// Id
	id := sqlDatabaseGetPropertiesStatusResource.Id
	destination.Id = &id

	// Rid
	destination.Rid = genruntime.ClonePointerToString(sqlDatabaseGetPropertiesStatusResource.Rid)

	// Ts
	if sqlDatabaseGetPropertiesStatusResource.Ts != nil {
		t := *sqlDatabaseGetPropertiesStatusResource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Users
	destination.Users = genruntime.ClonePointerToString(sqlDatabaseGetPropertiesStatusResource.Users)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlDatabaseResource
type SqlDatabaseResource struct {
	// +kubebuilder:validation:Required
	//Id: Name of the Cosmos DB SQL database
	Id string `json:"id"`
}

var _ genruntime.ARMTransformer = &SqlDatabaseResource{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (sqlDatabaseResource *SqlDatabaseResource) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if sqlDatabaseResource == nil {
		return nil, nil
	}
	var result SqlDatabaseResourceARM

	// Set property ‘Id’:
	result.Id = sqlDatabaseResource.Id
	return result, nil
}

// CreateEmptyARMValue returns an empty ARM value suitable for deserializing into
func (sqlDatabaseResource *SqlDatabaseResource) CreateEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (sqlDatabaseResource *SqlDatabaseResource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseResourceARM, got %T", armInput)
	}

	// Set property ‘Id’:
	sqlDatabaseResource.Id = typedInput.Id

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseResource populates our SqlDatabaseResource from the provided source SqlDatabaseResource
func (sqlDatabaseResource *SqlDatabaseResource) AssignPropertiesFromSqlDatabaseResource(source *v1alpha1api20210515storage.SqlDatabaseResource) error {

	// Id
	sqlDatabaseResource.Id = genruntime.GetOptionalStringValue(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseResource populates the provided destination SqlDatabaseResource from our SqlDatabaseResource
func (sqlDatabaseResource *SqlDatabaseResource) AssignPropertiesToSqlDatabaseResource(destination *v1alpha1api20210515storage.SqlDatabaseResource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Id
	id := sqlDatabaseResource.Id
	destination.Id = &id

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
	SchemeBuilder.Register(&SqlDatabase{}, &SqlDatabaseList{})
}
