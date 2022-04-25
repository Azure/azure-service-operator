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
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_storedProcedures
type SqlDatabaseContainerStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec `json:"spec,omitempty"`
	Status            SqlStoredProcedureGetResults_Status                         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerStoredProcedure{}

// GetConditions returns the conditions of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetConditions() conditions.Conditions {
	return procedure.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (procedure *SqlDatabaseContainerStoredProcedure) SetConditions(conditions conditions.Conditions) {
	procedure.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerStoredProcedure{}

// ConvertFrom populates our SqlDatabaseContainerStoredProcedure from the provided hub SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabaseContainerStoredProcedure)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerStoredProcedure but received %T instead", hub)
	}

	return procedure.AssignPropertiesFromSqlDatabaseContainerStoredProcedure(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseContainerStoredProcedure)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerStoredProcedure but received %T instead", hub)
	}

	return procedure.AssignPropertiesToSqlDatabaseContainerStoredProcedure(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1beta20210515-sqldatabasecontainerstoredprocedure,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1beta20210515,name=default.v1beta20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabaseContainerStoredProcedure{}

// Default applies defaults to the SqlDatabaseContainerStoredProcedure resource
func (procedure *SqlDatabaseContainerStoredProcedure) Default() {
	procedure.defaultImpl()
	var temp interface{} = procedure
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (procedure *SqlDatabaseContainerStoredProcedure) defaultAzureName() {
	if procedure.Spec.AzureName == "" {
		procedure.Spec.AzureName = procedure.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseContainerStoredProcedure resource
func (procedure *SqlDatabaseContainerStoredProcedure) defaultImpl() { procedure.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabaseContainerStoredProcedure{}

// AzureName returns the Azure name of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) AzureName() string {
	return procedure.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (procedure SqlDatabaseContainerStoredProcedure) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetSpec() genruntime.ConvertibleSpec {
	return &procedure.Spec
}

// GetStatus returns the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetStatus() genruntime.ConvertibleStatus {
	return &procedure.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
func (procedure *SqlDatabaseContainerStoredProcedure) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/storedProcedures"
}

// NewEmptyStatus returns a new empty (blank) status
func (procedure *SqlDatabaseContainerStoredProcedure) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlStoredProcedureGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (procedure *SqlDatabaseContainerStoredProcedure) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(procedure.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  procedure.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (procedure *SqlDatabaseContainerStoredProcedure) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlStoredProcedureGetResults_Status); ok {
		procedure.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlStoredProcedureGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	procedure.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1beta20210515-sqldatabasecontainerstoredprocedure,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1beta20210515,name=validate.v1beta20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabaseContainerStoredProcedure{}

// ValidateCreate validates the creation of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateCreate() error {
	validations := procedure.createValidations()
	var temp interface{} = procedure
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
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateDelete() error {
	validations := procedure.deleteValidations()
	var temp interface{} = procedure
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
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateUpdate(old runtime.Object) error {
	validations := procedure.updateValidations()
	var temp interface{} = procedure
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
func (procedure *SqlDatabaseContainerStoredProcedure) createValidations() []func() error {
	return []func() error{procedure.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return procedure.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (procedure *SqlDatabaseContainerStoredProcedure) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&procedure.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabaseContainerStoredProcedure populates our SqlDatabaseContainerStoredProcedure from the provided source SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesFromSqlDatabaseContainerStoredProcedure(source *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	procedure.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec() to populate field Spec")
	}
	procedure.Spec = spec

	// Status
	var status SqlStoredProcedureGetResults_Status
	err = status.AssignPropertiesFromSqlStoredProcedureGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureGetResultsStatus() to populate field Status")
	}
	procedure.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerStoredProcedure populates the provided destination SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignPropertiesToSqlDatabaseContainerStoredProcedure(destination *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	destination.ObjectMeta = *procedure.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
	err := procedure.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.SqlStoredProcedureGetResults_Status
	err = procedure.Status.AssignPropertiesToSqlStoredProcedureGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureGetResultsStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (procedure *SqlDatabaseContainerStoredProcedure) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: procedure.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerStoredProcedure",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_storedProcedures
type SqlDatabaseContainerStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerStoredProcedure `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecAPIVersion string

const DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecAPIVersion20210515 = DatabaseAccountsSqlDatabasesContainersStoredProceduresSpecAPIVersion("2021-05-15")

type DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are "If-Match",
	//"If-None-Match", "Session-Token" and "Throughput"
	Options *CreateUpdateOptions `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	//Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	//controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	//reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`

	// +kubebuilder:validation:Required
	//Resource: Cosmos DB SQL storedProcedure resource object
	Resource *SqlStoredProcedureResource `json:"resource,omitempty"`

	//Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
	//resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	//greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
	//type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph",
	//"DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if procedures == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM

	// Set property ‘Location’:
	if procedures.Location != nil {
		location := *procedures.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if procedures.Options != nil || procedures.Resource != nil {
		result.Properties = &SqlStoredProcedureCreateUpdatePropertiesARM{}
	}
	if procedures.Options != nil {
		optionsARM, err := (*procedures.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := optionsARM.(CreateUpdateOptionsARM)
		result.Properties.Options = &options
	}
	if procedures.Resource != nil {
		resourceARM, err := (*procedures.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := resourceARM.(SqlStoredProcedureResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if procedures.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range procedures.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesContainersStoredProcedures_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	procedures.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		procedures.Location = &location
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
			procedures.Options = &options
		}
	}

	// Set property ‘Owner’:
	procedures.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlStoredProcedureResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			procedures.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		procedures.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			procedures.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec from the provided source
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec)
	if ok {
		// Populate our instance from source
		return procedures.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = procedures.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec)
	if ok {
		// Populate destination from our instance
		return procedures.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{}
	err := procedures.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec populates our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec from the provided source DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(source *v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) error {

	// AzureName
	procedures.AzureName = source.AzureName

	// Location
	procedures.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		procedures.Options = &option
	} else {
		procedures.Options = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		procedures.Owner = &owner
	} else {
		procedures.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureResource
		err := resource.AssignPropertiesFromSqlStoredProcedureResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureResource() to populate field Resource")
		}
		procedures.Resource = &resource
	} else {
		procedures.Resource = nil
	}

	// Tags
	procedures.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec populates the provided destination DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec from our DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersStoredProceduresSpec(destination *v20210515s.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = procedures.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(procedures.Location)

	// Options
	if procedures.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := procedures.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = procedures.OriginalVersion()

	// Owner
	if procedures.Owner != nil {
		owner := procedures.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if procedures.Resource != nil {
		var resource v20210515s.SqlStoredProcedureResource
		err := procedures.Resource.AssignPropertiesToSqlStoredProcedureResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(procedures.Tags)

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
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (procedures *DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec) SetAzureName(azureName string) {
	procedures.AzureName = azureName
}

type SqlStoredProcedureGetResults_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	//Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	//Name: The name of the ARM resource.
	Name     *string                                          `json:"name,omitempty"`
	Resource *SqlStoredProcedureGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags     map[string]string                                `json:"tags,omitempty"`

	//Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlStoredProcedureGetResults_Status{}

// ConvertStatusFrom populates our SqlStoredProcedureGetResults_Status from the provided source
func (results *SqlStoredProcedureGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.SqlStoredProcedureGetResults_Status)
	if ok {
		// Populate our instance from source
		return results.AssignPropertiesFromSqlStoredProcedureGetResultsStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.SqlStoredProcedureGetResults_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignPropertiesFromSqlStoredProcedureGetResultsStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlStoredProcedureGetResults_Status
func (results *SqlStoredProcedureGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.SqlStoredProcedureGetResults_Status)
	if ok {
		// Populate destination from our instance
		return results.AssignPropertiesToSqlStoredProcedureGetResultsStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.SqlStoredProcedureGetResults_Status{}
	err := results.AssignPropertiesToSqlStoredProcedureGetResultsStatus(dst)
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

var _ genruntime.FromARMConverter = &SqlStoredProcedureGetResults_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (results *SqlStoredProcedureGetResults_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureGetResults_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (results *SqlStoredProcedureGetResults_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureGetResults_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureGetResults_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		results.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		results.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		results.Name = &name
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlStoredProcedureGetProperties_Status_Resource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			results.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		results.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			results.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		results.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureGetResultsStatus populates our SqlStoredProcedureGetResults_Status from the provided source SqlStoredProcedureGetResults_Status
func (results *SqlStoredProcedureGetResults_Status) AssignPropertiesFromSqlStoredProcedureGetResultsStatus(source *v20210515s.SqlStoredProcedureGetResults_Status) error {

	// Conditions
	results.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	results.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	results.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	results.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureGetProperties_Status_Resource
		err := resource.AssignPropertiesFromSqlStoredProcedureGetPropertiesStatusResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlStoredProcedureGetPropertiesStatusResource() to populate field Resource")
		}
		results.Resource = &resource
	} else {
		results.Resource = nil
	}

	// Tags
	results.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	results.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureGetResultsStatus populates the provided destination SqlStoredProcedureGetResults_Status from our SqlStoredProcedureGetResults_Status
func (results *SqlStoredProcedureGetResults_Status) AssignPropertiesToSqlStoredProcedureGetResultsStatus(destination *v20210515s.SqlStoredProcedureGetResults_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(results.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(results.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(results.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(results.Name)

	// Resource
	if results.Resource != nil {
		var resource v20210515s.SqlStoredProcedureGetProperties_Status_Resource
		err := results.Resource.AssignPropertiesToSqlStoredProcedureGetPropertiesStatusResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlStoredProcedureGetPropertiesStatusResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(results.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(results.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type SqlStoredProcedureGetProperties_Status_Resource struct {
	//Body: Body of the Stored Procedure
	Body *string `json:"body,omitempty"`

	//Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	//Id: Name of the Cosmos DB SQL storedProcedure
	Id *string `json:"id,omitempty"`

	//Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	//Ts: A system generated property that denotes the last updated timestamp of the resource.
	Ts *float64 `json:"_ts,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlStoredProcedureGetProperties_Status_Resource{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureGetProperties_Status_Resource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureGetProperties_Status_ResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureGetProperties_Status_Resource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureGetProperties_Status_ResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureGetProperties_Status_ResourceARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
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

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureGetPropertiesStatusResource populates our SqlStoredProcedureGetProperties_Status_Resource from the provided source SqlStoredProcedureGetProperties_Status_Resource
func (resource *SqlStoredProcedureGetProperties_Status_Resource) AssignPropertiesFromSqlStoredProcedureGetPropertiesStatusResource(source *v20210515s.SqlStoredProcedureGetProperties_Status_Resource) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

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

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureGetPropertiesStatusResource populates the provided destination SqlStoredProcedureGetProperties_Status_Resource from our SqlStoredProcedureGetProperties_Status_Resource
func (resource *SqlStoredProcedureGetProperties_Status_Resource) AssignPropertiesToSqlStoredProcedureGetPropertiesStatusResource(destination *v20210515s.SqlStoredProcedureGetProperties_Status_Resource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

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

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlStoredProcedureResource
type SqlStoredProcedureResource struct {
	//Body: Body of the Stored Procedure
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	//Id: Name of the Cosmos DB SQL storedProcedure
	Id *string `json:"id,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlStoredProcedureResource{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlStoredProcedureResource) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	var result SqlStoredProcedureResourceARM

	// Set property ‘Body’:
	if resource.Body != nil {
		body := *resource.Body
		result.Body = &body
	}

	// Set property ‘Id’:
	if resource.Id != nil {
		id := *resource.Id
		result.Id = &id
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureResource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureResource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureResourceARM, got %T", armInput)
	}

	// Set property ‘Body’:
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// No error
	return nil
}

// AssignPropertiesFromSqlStoredProcedureResource populates our SqlStoredProcedureResource from the provided source SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignPropertiesFromSqlStoredProcedureResource(source *v20210515s.SqlStoredProcedureResource) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlStoredProcedureResource populates the provided destination SqlStoredProcedureResource from our SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignPropertiesToSqlStoredProcedureResource(destination *v20210515s.SqlStoredProcedureResource) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Body
	destination.Body = genruntime.ClonePointerToString(resource.Body)

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
	SchemeBuilder.Register(&SqlDatabaseContainerStoredProcedure{}, &SqlDatabaseContainerStoredProcedureList{})
}
