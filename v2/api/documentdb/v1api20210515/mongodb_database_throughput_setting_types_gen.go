// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20210515/storage"
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
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default
type MongodbDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MongodbDatabaseThroughputSetting_Spec   `json:"spec,omitempty"`
	Status            MongodbDatabaseThroughputSetting_STATUS `json:"status,omitempty"`
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
	// intermediate variable for conversion
	var source storage.MongodbDatabaseThroughputSetting

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = setting.AssignProperties_From_MongodbDatabaseThroughputSetting(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to setting")
	}

	return nil
}

// ConvertTo populates the provided hub MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination storage.MongodbDatabaseThroughputSetting
	err := setting.AssignProperties_To_MongodbDatabaseThroughputSetting(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from setting")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-documentdb-azure-com-v1api20210515-mongodbdatabasethroughputsetting,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1api20210515,name=default.v1api20210515.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &MongodbDatabaseThroughputSetting{}

// Default applies defaults to the MongodbDatabaseThroughputSetting resource
func (setting *MongodbDatabaseThroughputSetting) Default() {
	setting.defaultImpl()
	var temp any = setting
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

// GetResourceScope returns the scope of the resource
func (setting *MongodbDatabaseThroughputSetting) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (setting *MongodbDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &setting.Spec
}

// GetStatus returns the status of this resource
func (setting *MongodbDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &setting.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (setting *MongodbDatabaseThroughputSetting) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
func (setting *MongodbDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/mongodbDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (setting *MongodbDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &MongodbDatabaseThroughputSetting_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (setting *MongodbDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(setting.Spec)
	return setting.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (setting *MongodbDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*MongodbDatabaseThroughputSetting_STATUS); ok {
		setting.Status = *st
		return nil
	}

	// Convert status to required version
	var st MongodbDatabaseThroughputSetting_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	setting.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-documentdb-azure-com-v1api20210515-mongodbdatabasethroughputsetting,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=documentdb.azure.com,resources=mongodbdatabasethroughputsettings,verbs=create;update,versions=v1api20210515,name=validate.v1api20210515.mongodbdatabasethroughputsettings.documentdb.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &MongodbDatabaseThroughputSetting{}

// ValidateCreate validates the creation of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateCreate() (admission.Warnings, error) {
	validations := setting.createValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateDelete() (admission.Warnings, error) {
	validations := setting.deleteValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (setting *MongodbDatabaseThroughputSetting) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := setting.updateValidations()
	var temp any = setting
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (setting *MongodbDatabaseThroughputSetting) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){setting.validateResourceReferences, setting.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (setting *MongodbDatabaseThroughputSetting) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (setting *MongodbDatabaseThroughputSetting) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return setting.validateResourceReferences()
		},
		setting.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return setting.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (setting *MongodbDatabaseThroughputSetting) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(setting)
}

// validateResourceReferences validates all resource references
func (setting *MongodbDatabaseThroughputSetting) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&setting.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (setting *MongodbDatabaseThroughputSetting) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*MongodbDatabaseThroughputSetting)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, setting)
}

// AssignProperties_From_MongodbDatabaseThroughputSetting populates our MongodbDatabaseThroughputSetting from the provided source MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_From_MongodbDatabaseThroughputSetting(source *storage.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	setting.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec MongodbDatabaseThroughputSetting_Spec
	err := spec.AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_MongodbDatabaseThroughputSetting_Spec() to populate field Spec")
	}
	setting.Spec = spec

	// Status
	var status MongodbDatabaseThroughputSetting_STATUS
	err = status.AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS() to populate field Status")
	}
	setting.Status = status

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting populates the provided destination MongodbDatabaseThroughputSetting from our MongodbDatabaseThroughputSetting
func (setting *MongodbDatabaseThroughputSetting) AssignProperties_To_MongodbDatabaseThroughputSetting(destination *storage.MongodbDatabaseThroughputSetting) error {

	// ObjectMeta
	destination.ObjectMeta = *setting.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.MongodbDatabaseThroughputSetting_Spec
	err := setting.Spec.AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_MongodbDatabaseThroughputSetting_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.MongodbDatabaseThroughputSetting_STATUS
	err = setting.Status.AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS() to populate field Status")
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
// Generator information:
// - Generated from: /cosmos-db/resource-manager/Microsoft.DocumentDB/stable/2021-05-15/cosmos-db.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DocumentDB/databaseAccounts/{accountName}/mongodbDatabases/{databaseName}/throughputSettings/default
type MongodbDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MongodbDatabaseThroughputSetting `json:"items"`
}

type MongodbDatabaseThroughputSetting_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a documentdb.azure.com/MongodbDatabase resource
	Owner *genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner,omitempty" kind:"MongodbDatabase"`

	// +kubebuilder:validation:Required
	// Resource: The standard JSON format of a resource throughput
	Resource *ThroughputSettingsResource `json:"resource,omitempty"`
	Tags     map[string]string           `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &MongodbDatabaseThroughputSetting_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (setting *MongodbDatabaseThroughputSetting_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if setting == nil {
		return nil, nil
	}
	result := &MongodbDatabaseThroughputSetting_Spec_ARM{}

	// Set property "Location":
	if setting.Location != nil {
		location := *setting.Location
		result.Location = &location
	}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if setting.Resource != nil {
		result.Properties = &ThroughputSettingsUpdateProperties_ARM{}
	}
	if setting.Resource != nil {
		resource_ARM, err := (*setting.Resource).ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		resource := *resource_ARM.(*ThroughputSettingsResource_ARM)
		result.Properties.Resource = &resource
	}

	// Set property "Tags":
	if setting.Tags != nil {
		result.Tags = make(map[string]string, len(setting.Tags))
		for key, value := range setting.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *MongodbDatabaseThroughputSetting_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &MongodbDatabaseThroughputSetting_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *MongodbDatabaseThroughputSetting_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(MongodbDatabaseThroughputSetting_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected MongodbDatabaseThroughputSetting_Spec_ARM, got %T", armInput)
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property "Owner":
	setting.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsResource
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &MongodbDatabaseThroughputSetting_Spec{}

// ConvertSpecFrom populates our MongodbDatabaseThroughputSetting_Spec from the provided source
func (setting *MongodbDatabaseThroughputSetting_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.MongodbDatabaseThroughputSetting_Spec)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.MongodbDatabaseThroughputSetting_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our MongodbDatabaseThroughputSetting_Spec
func (setting *MongodbDatabaseThroughputSetting_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.MongodbDatabaseThroughputSetting_Spec)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.MongodbDatabaseThroughputSetting_Spec{}
	err := setting.AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(dst)
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

// AssignProperties_From_MongodbDatabaseThroughputSetting_Spec populates our MongodbDatabaseThroughputSetting_Spec from the provided source MongodbDatabaseThroughputSetting_Spec
func (setting *MongodbDatabaseThroughputSetting_Spec) AssignProperties_From_MongodbDatabaseThroughputSetting_Spec(source *storage.MongodbDatabaseThroughputSetting_Spec) error {

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		setting.Owner = &owner
	} else {
		setting.Owner = nil
	}

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsResource
		err := resource.AssignProperties_From_ThroughputSettingsResource(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsResource() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting_Spec populates the provided destination MongodbDatabaseThroughputSetting_Spec from our MongodbDatabaseThroughputSetting_Spec
func (setting *MongodbDatabaseThroughputSetting_Spec) AssignProperties_To_MongodbDatabaseThroughputSetting_Spec(destination *storage.MongodbDatabaseThroughputSetting_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// OriginalVersion
	destination.OriginalVersion = setting.OriginalVersion()

	// Owner
	if setting.Owner != nil {
		owner := setting.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Resource
	if setting.Resource != nil {
		var resource storage.ThroughputSettingsResource
		err := setting.Resource.AssignProperties_To_ThroughputSettingsResource(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsResource() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

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
func (setting *MongodbDatabaseThroughputSetting_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type MongodbDatabaseThroughputSetting_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: The unique resource identifier of the ARM resource.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`

	// Name: The name of the ARM resource.
	Name     *string                                          `json:"name,omitempty"`
	Resource *ThroughputSettingsGetProperties_Resource_STATUS `json:"resource,omitempty"`
	Tags     map[string]string                                `json:"tags,omitempty"`

	// Type: The type of Azure resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &MongodbDatabaseThroughputSetting_STATUS{}

// ConvertStatusFrom populates our MongodbDatabaseThroughputSetting_STATUS from the provided source
func (setting *MongodbDatabaseThroughputSetting_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.MongodbDatabaseThroughputSetting_STATUS)
	if ok {
		// Populate our instance from source
		return setting.AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.MongodbDatabaseThroughputSetting_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = setting.AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our MongodbDatabaseThroughputSetting_STATUS
func (setting *MongodbDatabaseThroughputSetting_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.MongodbDatabaseThroughputSetting_STATUS)
	if ok {
		// Populate destination from our instance
		return setting.AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.MongodbDatabaseThroughputSetting_STATUS{}
	err := setting.AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(dst)
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

var _ genruntime.FromARMConverter = &MongodbDatabaseThroughputSetting_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (setting *MongodbDatabaseThroughputSetting_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &MongodbDatabaseThroughputSetting_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (setting *MongodbDatabaseThroughputSetting_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(MongodbDatabaseThroughputSetting_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected MongodbDatabaseThroughputSetting_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		setting.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		setting.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		setting.Name = &name
	}

	// Set property "Resource":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Resource != nil {
			var resource1 ThroughputSettingsGetProperties_Resource_STATUS
			err := resource1.PopulateFromARM(owner, *typedInput.Properties.Resource)
			if err != nil {
				return err
			}
			resource := resource1
			setting.Resource = &resource
		}
	}

	// Set property "Tags":
	if typedInput.Tags != nil {
		setting.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			setting.Tags[key] = value
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		setting.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS populates our MongodbDatabaseThroughputSetting_STATUS from the provided source MongodbDatabaseThroughputSetting_STATUS
func (setting *MongodbDatabaseThroughputSetting_STATUS) AssignProperties_From_MongodbDatabaseThroughputSetting_STATUS(source *storage.MongodbDatabaseThroughputSetting_STATUS) error {

	// Conditions
	setting.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	setting.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	setting.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	setting.Name = genruntime.ClonePointerToString(source.Name)

	// Resource
	if source.Resource != nil {
		var resource ThroughputSettingsGetProperties_Resource_STATUS
		err := resource.AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS(source.Resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		setting.Resource = &resource
	} else {
		setting.Resource = nil
	}

	// Tags
	setting.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	setting.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS populates the provided destination MongodbDatabaseThroughputSetting_STATUS from our MongodbDatabaseThroughputSetting_STATUS
func (setting *MongodbDatabaseThroughputSetting_STATUS) AssignProperties_To_MongodbDatabaseThroughputSetting_STATUS(destination *storage.MongodbDatabaseThroughputSetting_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(setting.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(setting.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(setting.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(setting.Name)

	// Resource
	if setting.Resource != nil {
		var resource storage.ThroughputSettingsGetProperties_Resource_STATUS
		err := setting.Resource.AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS(&resource)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ThroughputSettingsGetProperties_Resource_STATUS() to populate field Resource")
		}
		destination.Resource = &resource
	} else {
		destination.Resource = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(setting.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(setting.Type)

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
	SchemeBuilder.Register(&MongodbDatabaseThroughputSetting{}, &MongodbDatabaseThroughputSettingList{})
}
