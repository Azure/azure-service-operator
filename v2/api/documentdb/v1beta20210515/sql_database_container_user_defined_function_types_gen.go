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
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_userDefinedFunctions
type SqlDatabaseContainerUserDefinedFunction struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec `json:"spec,omitempty"`
	Status            SqlUserDefinedFunctionGetResults_Status                         `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseContainerUserDefinedFunction{}

// GetConditions returns the conditions of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetConditions() conditions.Conditions {
	return function.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (function *SqlDatabaseContainerUserDefinedFunction) SetConditions(conditions conditions.Conditions) {
	function.Status.Conditions = conditions
}

var _ conversion.Convertible = &SqlDatabaseContainerUserDefinedFunction{}

// ConvertFrom populates our SqlDatabaseContainerUserDefinedFunction from the provided hub SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20210515s.SqlDatabaseContainerUserDefinedFunction)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerUserDefinedFunction but received %T instead", hub)
	}

	return function.AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction(source)
}

// ConvertTo populates the provided hub SqlDatabaseContainerUserDefinedFunction from our SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210515s.SqlDatabaseContainerUserDefinedFunction)
	if !ok {
		return fmt.Errorf("expected documentdb/v1beta20210515storage/SqlDatabaseContainerUserDefinedFunction but received %T instead", hub)
	}

	return function.AssignPropertiesToSqlDatabaseContainerUserDefinedFunction(destination)
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1beta20210515-sqldatabasecontaineruserdefinedfunction,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontaineruserdefinedfunctions,verbs=create;update,versions=v1beta20210515,name=default.v1beta20210515.sqldatabasecontaineruserdefinedfunctions.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &SqlDatabaseContainerUserDefinedFunction{}

// Default applies defaults to the SqlDatabaseContainerUserDefinedFunction resource
func (function *SqlDatabaseContainerUserDefinedFunction) Default() {
	function.defaultImpl()
	var temp interface{} = function
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (function *SqlDatabaseContainerUserDefinedFunction) defaultAzureName() {
	if function.Spec.AzureName == "" {
		function.Spec.AzureName = function.Name
	}
}

// defaultImpl applies the code generated defaults to the SqlDatabaseContainerUserDefinedFunction resource
func (function *SqlDatabaseContainerUserDefinedFunction) defaultImpl() { function.defaultAzureName() }

var _ genruntime.KubernetesResource = &SqlDatabaseContainerUserDefinedFunction{}

// AzureName returns the Azure name of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) AzureName() string {
	return function.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (function SqlDatabaseContainerUserDefinedFunction) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetSpec() genruntime.ConvertibleSpec {
	return &function.Spec
}

// GetStatus returns the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) GetStatus() genruntime.ConvertibleStatus {
	return &function.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
func (function *SqlDatabaseContainerUserDefinedFunction) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/userDefinedFunctions"
}

// NewEmptyStatus returns a new empty (blank) status
func (function *SqlDatabaseContainerUserDefinedFunction) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &SqlUserDefinedFunctionGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (function *SqlDatabaseContainerUserDefinedFunction) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(function.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  function.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (function *SqlDatabaseContainerUserDefinedFunction) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*SqlUserDefinedFunctionGetResults_Status); ok {
		function.Status = *st
		return nil
	}

	// Convert status to required version
	var st SqlUserDefinedFunctionGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	function.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1beta20210515-sqldatabasecontaineruserdefinedfunction,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontaineruserdefinedfunctions,verbs=create;update,versions=v1beta20210515,name=validate.v1beta20210515.sqldatabasecontaineruserdefinedfunctions.documentdb.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &SqlDatabaseContainerUserDefinedFunction{}

// ValidateCreate validates the creation of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) ValidateCreate() error {
	validations := function.createValidations()
	var temp interface{} = function
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
func (function *SqlDatabaseContainerUserDefinedFunction) ValidateDelete() error {
	validations := function.deleteValidations()
	var temp interface{} = function
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
func (function *SqlDatabaseContainerUserDefinedFunction) ValidateUpdate(old runtime.Object) error {
	validations := function.updateValidations()
	var temp interface{} = function
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
func (function *SqlDatabaseContainerUserDefinedFunction) createValidations() []func() error {
	return []func() error{function.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (function *SqlDatabaseContainerUserDefinedFunction) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return function.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (function *SqlDatabaseContainerUserDefinedFunction) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&function.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction populates our SqlDatabaseContainerUserDefinedFunction from the provided source SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) AssignPropertiesFromSqlDatabaseContainerUserDefinedFunction(source *v20210515s.SqlDatabaseContainerUserDefinedFunction) error {

	// ObjectMeta
	function.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
	err := spec.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec() to populate field Spec")
	}
	function.Spec = spec

	// Status
	var status SqlUserDefinedFunctionGetResults_Status
	err = status.AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus() to populate field Status")
	}
	function.Status = status

	// No error
	return nil
}

// AssignPropertiesToSqlDatabaseContainerUserDefinedFunction populates the provided destination SqlDatabaseContainerUserDefinedFunction from our SqlDatabaseContainerUserDefinedFunction
func (function *SqlDatabaseContainerUserDefinedFunction) AssignPropertiesToSqlDatabaseContainerUserDefinedFunction(destination *v20210515s.SqlDatabaseContainerUserDefinedFunction) error {

	// ObjectMeta
	destination.ObjectMeta = *function.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
	err := function.Spec.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.SqlUserDefinedFunctionGetResults_Status
	err = function.Status.AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (function *SqlDatabaseContainerUserDefinedFunction) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: function.Spec.OriginalVersion(),
		Kind:    "SqlDatabaseContainerUserDefinedFunction",
	}
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_containers_userDefinedFunctions
type SqlDatabaseContainerUserDefinedFunctionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerUserDefinedFunction `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-05-15"}
type DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecAPIVersion string

const DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecAPIVersion20210515 = DatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpecAPIVersion("2021-05-15")

type DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Options: CreateUpdateOptions are a list of key-value pairs that describe the resource. Supported keys are "If-Match",
	// "If-None-Match", "Session-Token" and "Throughput"
	Options *CreateUpdateOptions `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`

	// +kubebuilder:validation:Required
	// Resource: Cosmos DB SQL userDefinedFunction resource object
	Resource *SqlUserDefinedFunctionResource `json:"resource,omitempty"`

	// Tags: Tags are a list of key-value pairs that describe the resource. These tags can be used in viewing and grouping this
	// resource (across resource groups). A maximum of 15 tags can be provided for a resource. Each tag must have a key no
	// greater than 128 characters and value no greater than 256 characters. For example, the default experience for a template
	// type is set with "defaultExperience": "Cassandra". Current "defaultExperience" values also include "Table", "Graph",
	// "DocumentDB", and "MongoDB".
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if functions == nil {
		return nil, nil
	}
	var result DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM

	// Set property ‘Location’:
	if functions.Location != nil {
		location := *functions.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if functions.Options != nil || functions.Resource != nil {
		result.Properties = &SqlUserDefinedFunctionCreateUpdatePropertiesARM{}
	}
	if functions.Options != nil {
		optionsARM, err := (*functions.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := optionsARM.(CreateUpdateOptionsARM)
		result.Properties.Options = &options
	}
	if functions.Resource != nil {
		resourceARM, err := (*functions.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := resourceARM.(SqlUserDefinedFunctionResourceARM)
		result.Properties.Resource = &resource
	}

	// Set property ‘Tags’:
	if functions.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range functions.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	functions.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		functions.Location = &location
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
			functions.Options = &options
		}
	}

	// Set property ‘Owner’:
	functions.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Resource’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlUserDefinedFunctionResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			functions.Resource = &resource
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		functions.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			functions.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec from the provided source
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec)
	if ok {
		// Populate our instance from source
		return functions.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = functions.AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec)
	if ok {
		// Populate destination from our instance
		return functions.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{}
	err := functions.AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(dst)
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

// AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec populates our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec from the provided source DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) AssignPropertiesFromDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(source *v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) error {

	// AzureName
	functions.AzureName = source.AzureName

	// Location
	functions.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignPropertiesFromCreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromCreateUpdateOptions() to populate field Options")
		}
		functions.Options = &option
	} else {
		functions.Options = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		functions.Owner = &owner
	} else {
		functions.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlUserDefinedFunctionResource
		err := resource.AssignPropertiesFromSqlUserDefinedFunctionResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlUserDefinedFunctionResource() to populate field Resource")
		}
		functions.Resource = &resource
	} else {
		functions.Resource = nil
	}

	// Tags
	functions.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec populates the provided destination DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec from our DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) AssignPropertiesToDatabaseAccountsSqlDatabasesContainersUserDefinedFunctionsSpec(destination *v20210515s.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = functions.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(functions.Location)

	// Options
	if functions.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := functions.Options.AssignPropertiesToCreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToCreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = functions.OriginalVersion()

	// Owner
	if functions.Owner != nil {
		owner := functions.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if functions.Resource != nil {
		var resource v20210515s.SqlUserDefinedFunctionResource
		err := functions.Resource.AssignPropertiesToSqlUserDefinedFunctionResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlUserDefinedFunctionResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(functions.Tags)

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
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (functions *DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec) SetAzureName(azureName string) {
	functions.AzureName = azureName
}

type SqlUserDefinedFunctionGetResults_Status struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name     *string                                              `json:"name,omitempty"`
	Resource *SqlUserDefinedFunctionGetProperties_Status_Resource `json:"resource,omitempty"`
	Tags     map[string]string                                    `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &SqlUserDefinedFunctionGetResults_Status{}

// ConvertStatusFrom populates our SqlUserDefinedFunctionGetResults_Status from the provided source
func (results *SqlUserDefinedFunctionGetResults_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.SqlUserDefinedFunctionGetResults_Status)
	if ok {
		// Populate our instance from source
		return results.AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.SqlUserDefinedFunctionGetResults_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = results.AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our SqlUserDefinedFunctionGetResults_Status
func (results *SqlUserDefinedFunctionGetResults_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.SqlUserDefinedFunctionGetResults_Status)
	if ok {
		// Populate destination from our instance
		return results.AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.SqlUserDefinedFunctionGetResults_Status{}
	err := results.AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus(dst)
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

var _ genruntime.FromARMConverter = &SqlUserDefinedFunctionGetResults_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (results *SqlUserDefinedFunctionGetResults_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlUserDefinedFunctionGetResults_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (results *SqlUserDefinedFunctionGetResults_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlUserDefinedFunctionGetResults_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlUserDefinedFunctionGetResults_StatusARM, got %T", armInput)
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
			var resource1 SqlUserDefinedFunctionGetProperties_Status_Resource
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

// AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus populates our SqlUserDefinedFunctionGetResults_Status from the provided source SqlUserDefinedFunctionGetResults_Status
func (results *SqlUserDefinedFunctionGetResults_Status) AssignPropertiesFromSqlUserDefinedFunctionGetResultsStatus(source *v20210515s.SqlUserDefinedFunctionGetResults_Status) error {

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
		var resource SqlUserDefinedFunctionGetProperties_Status_Resource
		err := resource.AssignPropertiesFromSqlUserDefinedFunctionGetPropertiesStatusResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSqlUserDefinedFunctionGetPropertiesStatusResource() to populate field Resource")
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

// AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus populates the provided destination SqlUserDefinedFunctionGetResults_Status from our SqlUserDefinedFunctionGetResults_Status
func (results *SqlUserDefinedFunctionGetResults_Status) AssignPropertiesToSqlUserDefinedFunctionGetResultsStatus(destination *v20210515s.SqlUserDefinedFunctionGetResults_Status) error {
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
		var resource v20210515s.SqlUserDefinedFunctionGetProperties_Status_Resource
		err := results.Resource.AssignPropertiesToSqlUserDefinedFunctionGetPropertiesStatusResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSqlUserDefinedFunctionGetPropertiesStatusResource() to populate field Resource")
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

type SqlUserDefinedFunctionGetProperties_Status_Resource struct {
	// Body: Body of the User Defined Function
	Body *string `json:"body,omitempty"`

	// Etag: A system generated property representing the resource etag required for optimistic concurrency control.
	Etag *string `json:"_etag,omitempty"`

	// Id: Name of the Cosmos DB SQL userDefinedFunction
	Id *string `json:"id,omitempty"`

	// Rid: A system generated property. A unique identifier.
	Rid *string `json:"_rid,omitempty"`

	// Ts: A system generated property that denotes the last updated timestamp of the resource.
	Ts *float64 `json:"_ts,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlUserDefinedFunctionGetProperties_Status_Resource{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlUserDefinedFunctionGetProperties_Status_Resource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlUserDefinedFunctionGetProperties_Status_ResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlUserDefinedFunctionGetProperties_Status_Resource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlUserDefinedFunctionGetProperties_Status_ResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlUserDefinedFunctionGetProperties_Status_ResourceARM, got %T", armInput)
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

// AssignPropertiesFromSqlUserDefinedFunctionGetPropertiesStatusResource populates our SqlUserDefinedFunctionGetProperties_Status_Resource from the provided source SqlUserDefinedFunctionGetProperties_Status_Resource
func (resource *SqlUserDefinedFunctionGetProperties_Status_Resource) AssignPropertiesFromSqlUserDefinedFunctionGetPropertiesStatusResource(source *v20210515s.SqlUserDefinedFunctionGetProperties_Status_Resource) error {

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

// AssignPropertiesToSqlUserDefinedFunctionGetPropertiesStatusResource populates the provided destination SqlUserDefinedFunctionGetProperties_Status_Resource from our SqlUserDefinedFunctionGetProperties_Status_Resource
func (resource *SqlUserDefinedFunctionGetProperties_Status_Resource) AssignPropertiesToSqlUserDefinedFunctionGetPropertiesStatusResource(destination *v20210515s.SqlUserDefinedFunctionGetProperties_Status_Resource) error {
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

// Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/definitions/SqlUserDefinedFunctionResource
type SqlUserDefinedFunctionResource struct {
	// Body: Body of the User Defined Function
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	// Id: Name of the Cosmos DB SQL userDefinedFunction
	Id *string `json:"id,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlUserDefinedFunctionResource{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlUserDefinedFunctionResource) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	var result SqlUserDefinedFunctionResourceARM

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
func (resource *SqlUserDefinedFunctionResource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlUserDefinedFunctionResourceARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlUserDefinedFunctionResource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlUserDefinedFunctionResourceARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlUserDefinedFunctionResourceARM, got %T", armInput)
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

// AssignPropertiesFromSqlUserDefinedFunctionResource populates our SqlUserDefinedFunctionResource from the provided source SqlUserDefinedFunctionResource
func (resource *SqlUserDefinedFunctionResource) AssignPropertiesFromSqlUserDefinedFunctionResource(source *v20210515s.SqlUserDefinedFunctionResource) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// No error
	return nil
}

// AssignPropertiesToSqlUserDefinedFunctionResource populates the provided destination SqlUserDefinedFunctionResource from our SqlUserDefinedFunctionResource
func (resource *SqlUserDefinedFunctionResource) AssignPropertiesToSqlUserDefinedFunctionResource(destination *v20210515s.SqlUserDefinedFunctionResource) error {
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
	SchemeBuilder.Register(&SqlDatabaseContainerUserDefinedFunction{}, &SqlDatabaseContainerUserDefinedFunctionList{})
}
