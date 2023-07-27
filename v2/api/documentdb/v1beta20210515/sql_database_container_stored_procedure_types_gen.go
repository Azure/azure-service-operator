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
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of SqlDatabaseContainerStoredProcedure. Use v1api20210515.SqlDatabaseContainerStoredProcedure instead
type SqlDatabaseContainerStoredProcedure struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec   `json:"spec,omitempty"`
	Status            DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS `json:"status,omitempty"`
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
	// intermediate variable for conversion
	var source v20210515s.SqlDatabaseContainerStoredProcedure

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = procedure.AssignProperties_From_SqlDatabaseContainerStoredProcedure(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to procedure")
	}

	return nil
}

// ConvertTo populates the provided hub SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20210515s.SqlDatabaseContainerStoredProcedure
	err := procedure.AssignProperties_To_SqlDatabaseContainerStoredProcedure(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from procedure")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1beta20210515-sqldatabasecontainerstoredprocedure,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1beta20210515,name=default.v1beta20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &SqlDatabaseContainerStoredProcedure{}

// Default applies defaults to the SqlDatabaseContainerStoredProcedure resource
func (procedure *SqlDatabaseContainerStoredProcedure) Default() {
	procedure.defaultImpl()
	var temp any = procedure
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
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
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
	return &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS{}
}

// Owner returns the ResourceReference of the owner
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
	if st, ok := status.(*DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS); ok {
		procedure.Status = *st
		return nil
	}

	// Convert status to required version
	var st DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	procedure.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1beta20210515-sqldatabasecontainerstoredprocedure,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=sqldatabasecontainerstoredprocedures,verbs=create;update,versions=v1beta20210515,name=validate.v1beta20210515.sqldatabasecontainerstoredprocedures.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &SqlDatabaseContainerStoredProcedure{}

// ValidateCreate validates the creation of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateCreate() (admission.Warnings, error) {
	validations := procedure.createValidations()
	var temp any = procedure
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateDelete() (admission.Warnings, error) {
	validations := procedure.deleteValidations()
	var temp any = procedure
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := procedure.updateValidations()
	var temp any = procedure
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){procedure.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (procedure *SqlDatabaseContainerStoredProcedure) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return procedure.validateResourceReferences()
		},
		procedure.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (procedure *SqlDatabaseContainerStoredProcedure) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&procedure.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (procedure *SqlDatabaseContainerStoredProcedure) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*SqlDatabaseContainerStoredProcedure)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, procedure)
}

// AssignProperties_From_SqlDatabaseContainerStoredProcedure populates our SqlDatabaseContainerStoredProcedure from the provided source SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignProperties_From_SqlDatabaseContainerStoredProcedure(source *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	procedure.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
	err := spec.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec() to populate field Spec")
	}
	procedure.Spec = spec

	// Status
	var status DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS
	err = status.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS() to populate field Status")
	}
	procedure.Status = status

	// No error
	return nil
}

// AssignProperties_To_SqlDatabaseContainerStoredProcedure populates the provided destination SqlDatabaseContainerStoredProcedure from our SqlDatabaseContainerStoredProcedure
func (procedure *SqlDatabaseContainerStoredProcedure) AssignProperties_To_SqlDatabaseContainerStoredProcedure(destination *v20210515s.SqlDatabaseContainerStoredProcedure) error {

	// ObjectMeta
	destination.ObjectMeta = *procedure.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
	err := procedure.Spec.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS
	err = procedure.Status.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS() to populate field Status")
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
// Deprecated version of SqlDatabaseContainerStoredProcedure. Use v1api20210515.SqlDatabaseContainerStoredProcedure instead
type SqlDatabaseContainerStoredProcedureList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseContainerStoredProcedure `json:"items"`
}

type DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string               `json:"azureName,omitempty"`
	Location  *string              `json:"location,omitempty"`
	Options   *CreateUpdateOptions `json:"options,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/SqlDatabaseContainer resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"SqlDatabaseContainer"`

	// +kubebuilder:validation:Required
	Resource *SqlStoredProcedureResource `json:"resource,omitempty"`
	Tags     map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if procedure == nil {
		return nil, nil
	}
	result := &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM{}

	// Set property "Location":
	if procedure.Location != nil {
		location := *procedure.Location
		result.Location = &location
	}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if procedure.Options != nil || procedure.Resource != nil {
		result.Properties = &SqlStoredProcedureCreateUpdateProperties_ARM{}
	}
	if procedure.Options != nil {
		options_ARM, err := (*procedure.Options).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		options := *options_ARM.(*CreateUpdateOptions_ARM)
		result.Properties.Options = &options
	}
	if procedure.Resource != nil {
		resource_ARM, err := (*procedure.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resource_ARM.(*SqlStoredProcedureResource_ARM)
		result.Properties.Resource = &resource
	}

	// Set property "Tags":
	if procedure.Tags != nil {
		result.Tags = make(map[string]string, len(procedure.Tags))
		for key, value := range procedure.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	procedure.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		procedure.Location = &location
	}

	// Set property "Options":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Options != nil {
			var options1 CreateUpdateOptions
			err := options1.PopulateFromARM(owner, *typedInput.Properties.Options)
			if err != nil {
				return err
			}
			options := options1
			procedure.Options = &options
		}
	}

	// Set property "Owner":
	procedure.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlStoredProcedureResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			procedure.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		procedure.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			procedure.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}

// ConvertSpecFrom populates our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec from the provided source
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec)
	if ok {
		// Populate our instance from source
		return procedure.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = procedure.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec)
	if ok {
		// Populate destination from our instance
		return procedure.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{}
	err := procedure.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(dst)
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

// AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec populates our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec from the provided source DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(source *v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) error {

	// AzureName
	procedure.AzureName = source.AzureName

	// Location
	procedure.Location = genruntime.ClonePointerToString(source.Location)

	// Options
	if source.Options != nil {
		var option CreateUpdateOptions
		err := option.AssignProperties_From_CreateUpdateOptions(source.Options)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_CreateUpdateOptions() to populate field Options")
		}
		procedure.Options = &option
	} else {
		procedure.Options = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		procedure.Owner = &owner
	} else {
		procedure.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureResource
		err := resource.AssignProperties_From_SqlStoredProcedureResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlStoredProcedureResource() to populate field Resource")
		}
		procedure.Resource = &resource
	} else {
		procedure.Resource = nil
	}

	// Tags
	procedure.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec populates the provided destination DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec from our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec(destination *v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = procedure.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(procedure.Location)

	// Options
	if procedure.Options != nil {
		var option v20210515s.CreateUpdateOptions
		err := procedure.Options.AssignProperties_To_CreateUpdateOptions(&option)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_CreateUpdateOptions() to populate field Options")
		}
		destination.Options = &option
	} else {
		destination.Options = nil
	}

	// OriginalVersion
	destination.OriginalVersion = procedure.OriginalVersion()

	// Owner
	if procedure.Owner != nil {
		owner := procedure.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if procedure.Resource != nil {
		var resource v20210515s.SqlStoredProcedureResource
		err := procedure.Resource.AssignProperties_To_SqlStoredProcedureResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlStoredProcedureResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(procedure.Tags)

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
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec) SetAzureName(azureName string) {
	procedure.AzureName = azureName
}

// Deprecated version of DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS. Use v1api20210515.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS instead
type DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition                           `json:"conditions,omitempty"`
	Id         *string                                          `json:"id,omitempty"`
	Location   *string                                          `json:"location,omitempty"`
	Name       *string                                          `json:"name,omitempty"`
	Resource   *SqlStoredProcedureGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags       map[string]string                                `json:"tags,omitempty"`
	Type       *string                                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS{}

// ConvertStatusFrom populates our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS from the provided source
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS)
	if ok {
		// Populate our instance from source
		return procedure.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = procedure.AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS)
	if ok {
		// Populate destination from our instance
		return procedure.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS{}
	err := procedure.AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(dst)
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

var _ genruntime.FromARMConverter = &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		procedure.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		procedure.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		procedure.Name = &name
	}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 SqlStoredProcedureGetProperties_Resource_STATUS
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			procedure.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		procedure.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			procedure.Tags[key] = value
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		procedure.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS populates our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS from the provided source DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) AssignProperties_From_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(source *v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) error {

	// Conditions
	procedure.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	procedure.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	procedure.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	procedure.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource SqlStoredProcedureGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_SqlStoredProcedureGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SqlStoredProcedureGetProperties_Resource_STATUS() to populate field Resource")
		}
		procedure.Resource = &resource
	} else {
		procedure.Resource = nil
	}

	// Tags
	procedure.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	procedure.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS populates the provided destination DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS from our DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS
func (procedure *DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) AssignProperties_To_DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS(destination *v20210515s.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(procedure.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(procedure.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(procedure.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(procedure.Name)

	// Resource
	if procedure.Resource != nil {
		var resource v20210515s.SqlStoredProcedureGetProperties_Resource_STATUS
		err := procedure.Resource.AssignProperties_To_SqlStoredProcedureGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SqlStoredProcedureGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(procedure.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(procedure.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of SqlStoredProcedureGetProperties_Resource_STATUS. Use v1api20210515.SqlStoredProcedureGetProperties_Resource_STATUS instead
type SqlStoredProcedureGetProperties_Resource_STATUS struct {
	Body *string  `json:"body,omitempty"`
	Etag *string  `json:"_etag,omitempty"`
	Id   *string  `json:"id,omitempty"`
	Rid  *string  `json:"_rid,omitempty"`
	Ts   *float64 `json:"_ts,omitempty"`
}

var _ genruntime.FromARMConverter = &SqlStoredProcedureGetProperties_Resource_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureGetProperties_Resource_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureGetProperties_Resource_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureGetProperties_Resource_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureGetProperties_Resource_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureGetProperties_Resource_STATUS_ARM, got %T", armInput)
	}

	// Set property "Body":
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property "Etag":
	if typedInput.Etag != nil {
		etag := *typedInput.Etag
		resource.Etag = &etag
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// Set property "Rid":
	if typedInput.Rid != nil {
		rid := *typedInput.Rid
		resource.Rid = &rid
	}

	// Set property "Ts":
	if typedInput.Ts != nil {
		ts := *typedInput.Ts
		resource.Ts = &ts
	}

	// No error
	return nil
}

// AssignProperties_From_SqlStoredProcedureGetProperties_Resource_STATUS populates our SqlStoredProcedureGetProperties_Resource_STATUS from the provided source SqlStoredProcedureGetProperties_Resource_STATUS
func (resource *SqlStoredProcedureGetProperties_Resource_STATUS) AssignProperties_From_SqlStoredProcedureGetProperties_Resource_STATUS(source *v20210515s.SqlStoredProcedureGetProperties_Resource_STATUS) error {

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

// AssignProperties_To_SqlStoredProcedureGetProperties_Resource_STATUS populates the provided destination SqlStoredProcedureGetProperties_Resource_STATUS from our SqlStoredProcedureGetProperties_Resource_STATUS
func (resource *SqlStoredProcedureGetProperties_Resource_STATUS) AssignProperties_To_SqlStoredProcedureGetProperties_Resource_STATUS(destination *v20210515s.SqlStoredProcedureGetProperties_Resource_STATUS) error {
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

// Deprecated version of SqlStoredProcedureResource. Use v1api20210515.SqlStoredProcedureResource instead
type SqlStoredProcedureResource struct {
	Body *string `json:"body,omitempty"`

	// +kubebuilder:validation:Required
	Id *string `json:"id,omitempty"`
}

var _ genruntime.ARMTransformer = &SqlStoredProcedureResource{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (resource *SqlStoredProcedureResource) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if resource == nil {
		return nil, nil
	}
	result := &SqlStoredProcedureResource_ARM{}

	// Set property "Body":
	if resource.Body != nil {
		body := *resource.Body
		result.Body = &body
	}

	// Set property "Id":
	if resource.Id != nil {
		id := *resource.Id
		result.Id = &id
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (resource *SqlStoredProcedureResource) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &SqlStoredProcedureResource_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (resource *SqlStoredProcedureResource) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(SqlStoredProcedureResource_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected SqlStoredProcedureResource_ARM, got %T", armInput)
	}

	// Set property "Body":
	if typedInput.Body != nil {
		body := *typedInput.Body
		resource.Body = &body
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		resource.Id = &id
	}

	// No error
	return nil
}

// AssignProperties_From_SqlStoredProcedureResource populates our SqlStoredProcedureResource from the provided source SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignProperties_From_SqlStoredProcedureResource(source *v20210515s.SqlStoredProcedureResource) error {

	// Body
	resource.Body = genruntime.ClonePointerToString(source.Body)

	// Id
	resource.Id = genruntime.ClonePointerToString(source.Id)

	// No error
	return nil
}

// AssignProperties_To_SqlStoredProcedureResource populates the provided destination SqlStoredProcedureResource from our SqlStoredProcedureResource
func (resource *SqlStoredProcedureResource) AssignProperties_To_SqlStoredProcedureResource(destination *v20210515s.SqlStoredProcedureResource) error {
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
