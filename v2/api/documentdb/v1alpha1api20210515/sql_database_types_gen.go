// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515

import (
	"fmt"
	alpha20210515s "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515storage"
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
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabase struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabase_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccountsSqlDatabase_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabase{}

// GetConditions returns the conditions of the resource
func (database *SqlDatabase) GetConditions() conditions.Conditions {
	return database.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (database *SqlDatabase) SetConditions(conditions conditions.Conditions) {
	database.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabase{}

// ConvertFrom populates our SqlDatabase from the provided hub SqlDatabase
func (database *SqlDatabase) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20210515s.SqlDatabase

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = database.AssignPropertiesFromSqlDatabase(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to database")
	}

	return nil
}

// ConvertTo populates the provided hub SqlDatabase from our SqlDatabase
func (database *SqlDatabase) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210515s.SqlDatabase
	err := database.AssignPropertiesToSqlDatabase(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from database")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1alpha1api20210515-sqldatabase,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabases,verbs=create;update,versions=v1alpha1api20210515,name=default.v1alpha1api20210515.sqldatabases.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabase{}

// Default applies defaults to the SqlDatabase resource
func (database *SqlDatabase) Default() {
	database.defaultImpl()
	var temp interface{} = database
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (database *SqlDatabase) defaultAzureName() {
	if database.Spec.AzureName == "" {
		database.Spec.AzureName = database.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabase resource
func (database *SqlDatabase) defaultImpl() { database.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabase{}

// AzureName returns the Azure name of the resource
func (database *SqlDatabase) AzureName() string {
	return database.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (database SqlDatabase) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceKind returns the kind of the resource
func (database *SqlDatabase) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (database *SqlDatabase) GetSpec() genruntime.ConvertibleSpec {
	return &database.Spec
}

// GetStatus returns the status of this resource
func (database *SqlDatabase) GetStatus() genruntime.ConvertibleStatus {
	return &database.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
func (database *SqlDatabase) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases"
}

// NewEmptyStatus returns a new empty (blank) status
func (database *SqlDatabase) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DatabaseAccountsSqlDatabase_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (database *SqlDatabase) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(database.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  database.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (database *SqlDatabase) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DatabaseAccountsSqlDatabase_STATUS); ok {
		database.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccountsSqlDatabase_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	database.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1alpha1api20210515-sqldatabase,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabases,verbs=create;update,versions=v1alpha1api20210515,name=validate.v1alpha1api20210515.sqldatabases.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabase{}

// ValidateCreate validates the creation of the resource
func (database *SqlDatabase) ValidateCreate() error {
	validations := database.createValidations()
	var temp interface{} = database
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
func (database *SqlDatabase) ValidateDelete() error {
	validations := database.deleteValidations()
	var temp interface{} = database
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
func (database *SqlDatabase) ValidateUpdate(old runtime.Object) error {
	validations := database.updateValidations()
	var temp interface{} = database
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
func (database *SqlDatabase) createValidations() []func() error {
	return []func() error{database.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (database *SqlDatabase) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (database *SqlDatabase) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return database.validateResourceReferences()
		},
		database.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (database *SqlDatabase) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&database.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (database *SqlDatabase) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*SqlDatabase)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, database)
}

// AssignPropertiesFromSqlDatabase populates our SqlDatabase from the provided source SqlDatabase
func (database *SqlDatabase) AssignPropertiesFromSqlDatabase(source *alpha20210515s.SqlDatabase) error {

	// ObjectMeta
	database.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabase_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec() to populate field Spec")
	}
	database.Spec = spec

	// Status
	var status DatabaseAccountsSqlDatabase_STATUS
	err = status.AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS() to populate field Status")
	}
	database.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabase populates the provided destination SqlDatabase from our SqlDatabase
func (database *SqlDatabase) AssignPropertiesToSqlDatabase(destination *alpha20210515s.SqlDatabase) error {

	// ObjectMeta
	destination.ObjectMeta = *database.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210515s.DatabaseAccountsSqlDatabase_Spec
	err := database.Spec.AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabase_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210515s.DatabaseAccountsSqlDatabase_STATUS
	err = database.Status.AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (database *SqlDatabase) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: database.Spec.OriginalVersion(),
		Kind:    "SqlDatabase",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of SqlDatabase. Use v1beta20210515.SqlDatabase instead
type SqlDatabaseList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabase `json:"items"`
}

// Deprecated version of DatabaseAccountsSqlDatabase_STATUS. Use v1beta20210515.DatabaseAccountsSqlDatabase_STATUS instead
type DatabaseAccountsSqlDatabase_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition                    `json:"conditions,omitempty"`
	Id         *string                                   `json:"id,omitempty"`
	Location   *string                                   `json:"location,omitempty"`
	Name       *string                                   `json:"name,omitempty"`
	Options    *OptionsResource_STATUS                   `json:"options,omitempty"`
	Resource   *SqlDatabaseGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags       map[string]string                         `json:"tags,omitempty"`
	Type       *string                                   `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccountsSqlDatabase_STATUS{}

// ConvertStatusFrom populates our DatabaseAccountsSqlDatabase_STATUS from the provided source
func (database *DatabaseAccountsSqlDatabase_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210515s.DatabaseAccountsSqlDatabase_STATUS)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.DatabaseAccountsSqlDatabase_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccountsSqlDatabase_STATUS
func (database *DatabaseAccountsSqlDatabase_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210515s.DatabaseAccountsSqlDatabase_STATUS)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.DatabaseAccountsSqlDatabase_STATUS{}
	err := database.AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(dst)
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

var _ genruntime.FromARMConverter = &DatabaseAccountsSqlDatabase_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *DatabaseAccountsSqlDatabase_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabase_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *DatabaseAccountsSqlDatabase_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabase_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabase_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		database.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		database.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		database.Name = &name
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 OptionsResource_STATUS
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			database.Options = &options
		}
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlDatabaseGetProperties_Resource_STATUS
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			database.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		database.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			database.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		database.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS populates our DatabaseAccountsSqlDatabase_STATUS from the provided source DatabaseAccountsSqlDatabase_STATUS
func (database *DatabaseAccountsSqlDatabase_STATUS) AssignPropertiesFromDatabaseAccountsSqlDatabase_STATUS(source *alpha20210515s.DatabaseAccountsSqlDatabase_STATUS) error {

	// Conditions
	database.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	database.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	database.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	database.Name = genruntime.ClonePointerToString(source.Name)

	// Options
	if source.Options != nil {
		var option OptionsResource_STATUS
		err := option.AssignPropertiesFromOptionsResource_STATUS(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromOptionsResource_STATUS() to populate field Options")
		}
		database.Options = &option
	} else {
		database.Options = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseGetProperties_Resource_STATUS
		err := resource.AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS() to populate field Resource")
		}
		database.Resource = &resource
	} else {
		database.Resource = nil
	}

	// Tags
	database.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	database.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS populates the provided destination DatabaseAccountsSqlDatabase_STATUS from our DatabaseAccountsSqlDatabase_STATUS
func (database *DatabaseAccountsSqlDatabase_STATUS) AssignPropertiesToDatabaseAccountsSqlDatabase_STATUS(destination *alpha20210515s.DatabaseAccountsSqlDatabase_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(database.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(database.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(database.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(database.Name)

	// Options
	if database.Options != nil {
		var option alpha20210515s.OptionsResource_STATUS
		err := database.Options.AssignPropertiesToOptionsResource_STATUS(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToOptionsResource_STATUS() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// Resource
	if database.Resource != nil {
		var resource alpha20210515s.SqlDatabaseGetProperties_Resource_STATUS
		err := database.Resource.AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(database.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(database.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type DatabaseAccountsSqlDatabase_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string               `json:"azureName,omitempty"`
	Location  *string              `json:"location,omitempty"`
	Options   *CreateUpdateOptions `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/DatabaseAccount resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"DatabaseAccount"`

	// +kubebuilder:validation:Required
	Resource *SqlDatabaseResource `json:"resource,omitempty"`
	Tags     map[string]string    `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabase_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (database *DatabaseAccountsSqlDatabase_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if database == nil {
		return nil, nil
	}
	result := &DatabaseAccountsSqlDatabase_SpecARM{}

	// Set property ‘AzureName’:
	result.AzureName = database.AzureName

	// Set property ‘Location’:
	if database.Location != nil {
		location := *database.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if database.Options != nil || database.Resource != nil {
		result.Properties = &SqlDatabaseCreateUpdatePropertiesARM{}
	}
	if database.Options != nil {
		optionsARM, err := (*database.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := *optionsARM.(*CreateUpdateOptionsARM)
		result.Properties.Options = &options
	}
	if database.Resource != nil {
		resourceARM, err := (*database.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resourceARM.(*SqlDatabaseResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if database.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range database.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (database *DatabaseAccountsSqlDatabase_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabase_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (database *DatabaseAccountsSqlDatabase_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabase_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabase_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	database.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		database.Location = &location
	}

	// Set property ‘Options’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 CreateUpdateOptions
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			database.Options = &options
		}
	}

	// Set property ‘Owner’:
	database.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlDatabaseResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			database.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		database.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			database.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabase_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabase_Spec from the provided source
func (database *DatabaseAccountsSqlDatabase_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210515s.DatabaseAccountsSqlDatabase_Spec)
	if ok {
		// Populate our instance from source
		return database.AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210515s.DatabaseAccountsSqlDatabase_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = database.AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabase_Spec
func (database *DatabaseAccountsSqlDatabase_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210515s.DatabaseAccountsSqlDatabase_Spec)
	if ok {
		// Populate destination from our instance
		return database.AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210515s.DatabaseAccountsSqlDatabase_Spec{}
	err := database.AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec populates our DatabaseAccountsSqlDatabase_Spec from the provided source DatabaseAccountsSqlDatabase_Spec
func (database *DatabaseAccountsSqlDatabase_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabase_Spec(source *alpha20210515s.DatabaseAccountsSqlDatabase_Spec) error {

	// AzureName
	database.AzureName = source.AzureName

	// Location
	database.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		database.Options = &option
	} else {
		database.Options = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		database.Owner = &owner
	} else {
		database.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlDatabaseResource
		err := resource.AssignPropertiesFromSqlDatabaseResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlDatabaseResource() to populate field Resource")
		}
		database.Resource = &resource
	} else {
		database.Resource = nil
	}

	// Tags
	database.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabase_Spec populates the provided destination DatabaseAccountsSqlDatabase_Spec from our DatabaseAccountsSqlDatabase_Spec
func (database *DatabaseAccountsSqlDatabase_Spec) AssignPropertiesToDatabaseAccountsSqlDatabase_Spec(destination *alpha20210515s.DatabaseAccountsSqlDatabase_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = database.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(database.Location)

	// Options
	if database.Options != nil {
		var option alpha20210515s.CreateUpdateOptions
		err := database.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = database.OriginalVersion()

	// Owner
	if database.Owner != nil {
		owner := database.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if database.Resource != nil {
		var resource alpha20210515s.SqlDatabaseResource
		err := database.Resource.AssignPropertiesToSqlDatabaseResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlDatabaseResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(database.Tags)

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
func (database *DatabaseAccountsSqlDatabase_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (database *DatabaseAccountsSqlDatabase_Spec) SetAzureName(azureName string) {
	database.AzureName = azureName
}

// Deprecated version of SqlDatabaseGetProperties_Resource_STATUS. Use v1beta20210515.SqlDatabaseGetProperties_Resource_STATUS instead
type SqlDatabaseGetProperties_Resource_STATUS struct {
	Colls *string  `json:"_colls,omitempty"`
	Etag  *string  `json:"_etag,omitempty"`
	Id    *string  `json:"id,omitempty"`
	Rid   *string  `json:"_rid,omitempty"`
	Ts    *float64 `json:"_ts,omitempty"`
	Users *string  `json:"_users,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlDatabaseGetProperties_Resource_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlDatabaseGetProperties_Resource_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseGetProperties_Resource_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlDatabaseGetProperties_Resource_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseGetProperties_Resource_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseGetProperties_Resource_STATUSARM, got %T", armInput)
	}

	// Set property ‘Colls’:
	if typedInput.Colls != nil {
		colls := *typedInput.Colls
		resource.Colls = &colls
	}

	// Set property ‘Etag’:
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		resource.Etag = &etag
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// Set property ‘Rid’:
	if typedInput.Rid != nil {
		rid := *typedInput.Rid
		resource.Rid = &rid
	}

	// Set property ‘Ts’:
	if typedInput.Ts != nil {
		ts := *typedInput.Ts
		resource.Ts = &ts
	}

	// Set property ‘Users’:
	if typedInput.Users != nil {
		users := *typedInput.Users
		resource.Users = &users
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS populates our SqlDatabaseGetProperties_Resource_STATUS from the provided source SqlDatabaseGetProperties_Resource_STATUS
func (resource *SqlDatabaseGetProperties_Resource_STATUS) AssignPropertiesFromSqlDatabaseGetProperties_Resource_STATUS(source *alpha20210515s.SqlDatabaseGetProperties_Resource_STATUS) error {

	// Colls
	resource.Colls = genruntime.ClonePointerToString(source.Colls)

	// Etag
	resource.Etag = genruntime.ClonePointerToString(source.Etag)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// Rid
	resource.Rid = genruntime.ClonePointerToString(source.Rid)

	// Ts
	if source.Ts != nil {
		t := *source.Ts
		resource.Ts = &t
	} else {
		resource.Ts = nil
	}

	// Users
	resource.Users = genruntime.ClonePointerToString(source.Users)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS populates the provided destination SqlDatabaseGetProperties_Resource_STATUS from our SqlDatabaseGetProperties_Resource_STATUS
func (resource *SqlDatabaseGetProperties_Resource_STATUS) AssignPropertiesToSqlDatabaseGetProperties_Resource_STATUS(destination *alpha20210515s.SqlDatabaseGetProperties_Resource_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Colls
	destination.Colls = genruntime.ClonePointerToString(resource.Colls)

	// Etag
	destination.Etag = genruntime.ClonePointerToString(resource.Etag)

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

	// Rid
	destination.Rid = genruntime.ClonePointerToString(resource.Rid)

	// Ts
	if resource.Ts != nil {
		t := *resource.Ts
		destination.Ts = &t
	} else {
		destination.Ts = nil
	}

	// Users
	destination.Users = genruntime.ClonePointerToString(resource.Users)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of SqlDatabaseResource. Use v1beta20210515.SqlDatabaseResource instead
type SqlDatabaseResource struct {
	// +kubebuilder:validation:Required
	Id *string `json:"id,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlDatabaseResource{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlDatabaseResource) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	result := &SqlDatabaseResourceARM{}

	// Set property ‘Id’:
	if resource.Id != nil {
		id := *resource.Id
		result.Id = &id
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlDatabaseResource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlDatabaseResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlDatabaseResource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlDatabaseResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlDatabaseResourceARM, got %T", armInput)
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlDatabaseResource populates our SqlDatabaseResource from the provided source SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesFromSqlDatabaseResource(source *alpha20210515s.SqlDatabaseResource) error {

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseResource populates the provided destination SqlDatabaseResource from our SqlDatabaseResource
func (resource *SqlDatabaseResource) AssignPropertiesToSqlDatabaseResource(destination *alpha20210515s.SqlDatabaseResource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Id
	destination.Id = genruntime.ClonePointerToString(resource.Id)

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
