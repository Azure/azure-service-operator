// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

import (
	"fmt"
	v1api20200601s "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601storage"
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
// - Generated from: /resources/resource-manager/Microsoft.Resources/stable/2020-06-01/resources.json
// - ARM URI: /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}
type ResourceGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ResourceGroup_Spec   `json:"spec,omitempty"`
	Status            ResourceGroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ResourceGroup{}

// GetConditions returns the conditions of the resource
func (group *ResourceGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *ResourceGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ conversion.Convertible = &ResourceGroup{}

// ConvertFrom populates our ResourceGroup from the provided hub ResourceGroup
func (group *ResourceGroup) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1api20200601s.ResourceGroup)
	if !ok {
		return fmt.Errorf("expected resources/v1api20200601storage/ResourceGroup but received %T instead", hub)
	}

	return group.AssignProperties_From_ResourceGroup(source)
}

// ConvertTo populates the provided hub ResourceGroup from our ResourceGroup
func (group *ResourceGroup) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1api20200601s.ResourceGroup)
	if !ok {
		return fmt.Errorf("expected resources/v1api20200601storage/ResourceGroup but received %T instead", hub)
	}

	return group.AssignProperties_To_ResourceGroup(destination)
}

// +kubebuilder:webhook:path=/mutate-resources-azure-com-v1api20200601-resourcegroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=resources.azure.com,resources=resourcegroups,verbs=create;update,versions=v1api20200601,name=default.v1api20200601.resourcegroups.resources.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ResourceGroup{}

// Default applies defaults to the ResourceGroup resource
func (group *ResourceGroup) Default() {
	group.defaultImpl()
	var temp any = group
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (group *ResourceGroup) defaultAzureName() {
	if group.Spec.AzureName == "" {
		group.Spec.AzureName = group.Name
	}
}

// defaultImpl applies the code generated defaults to the ResourceGroup resource
func (group *ResourceGroup) defaultImpl() { group.defaultAzureName() }

var _ genruntime.ImportableResource = &ResourceGroup{}

// InitializeSpec initializes the spec for this resource from the given status
func (group *ResourceGroup) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*ResourceGroup_STATUS); ok {
		return group.Spec.Initialize_From_ResourceGroup_STATUS(s)
	}

	return fmt.Errorf("expected Status of type ResourceGroup_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ResourceGroup{}

// AzureName returns the Azure name of the resource
func (group *ResourceGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (group ResourceGroup) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (group *ResourceGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeLocation
}

// GetSpec returns the specification of this resource
func (group *ResourceGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *ResourceGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Resources/resourceGroups"
func (group *ResourceGroup) GetType() string {
	return "Microsoft.Resources/resourceGroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *ResourceGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ResourceGroup_STATUS{}
}

// Owner returns nil as Location scoped resources never have an owner
func (group *ResourceGroup) Owner() *genruntime.ResourceReference {
	return nil
}

// SetStatus sets the status of this resource
func (group *ResourceGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ResourceGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st ResourceGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

var _ genruntime.LocatableResource = &ResourceGroup{}

// Location returns the location of the resource
func (group *ResourceGroup) Location() string {
	if group.Spec.Location == nil {
		return ""
	}
	return *group.Spec.Location
}

// +kubebuilder:webhook:path=/validate-resources-azure-com-v1api20200601-resourcegroup,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=resources.azure.com,resources=resourcegroups,verbs=create;update,versions=v1api20200601,name=validate.v1api20200601.resourcegroups.resources.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ResourceGroup{}

// ValidateCreate validates the creation of the resource
func (group *ResourceGroup) ValidateCreate() (admission.Warnings, error) {
	validations := group.createValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (group *ResourceGroup) ValidateDelete() (admission.Warnings, error) {
	validations := group.deleteValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (group *ResourceGroup) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := group.updateValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (group *ResourceGroup) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){group.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (group *ResourceGroup) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (group *ResourceGroup) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return group.validateResourceReferences()
		},
		group.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (group *ResourceGroup) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&group.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (group *ResourceGroup) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ResourceGroup)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, group)
}

// AssignProperties_From_ResourceGroup populates our ResourceGroup from the provided source ResourceGroup
func (group *ResourceGroup) AssignProperties_From_ResourceGroup(source *v1api20200601s.ResourceGroup) error {

	// ObjectMeta
	group.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec ResourceGroup_Spec
	err := spec.AssignProperties_From_ResourceGroup_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ResourceGroup_Spec() to populate field Spec")
	}
	group.Spec = spec

	// Status
	var status ResourceGroup_STATUS
	err = status.AssignProperties_From_ResourceGroup_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ResourceGroup_STATUS() to populate field Status")
	}
	group.Status = status

	// No error
	return nil
}

// AssignProperties_To_ResourceGroup populates the provided destination ResourceGroup from our ResourceGroup
func (group *ResourceGroup) AssignProperties_To_ResourceGroup(destination *v1api20200601s.ResourceGroup) error {

	// ObjectMeta
	destination.ObjectMeta = *group.ObjectMeta.DeepCopy()

	// Spec
	var spec v1api20200601s.ResourceGroup_Spec
	err := group.Spec.AssignProperties_To_ResourceGroup_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ResourceGroup_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1api20200601s.ResourceGroup_STATUS
	err = group.Status.AssignProperties_To_ResourceGroup_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ResourceGroup_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *ResourceGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion(),
		Kind:    "ResourceGroup",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /resources/resource-manager/Microsoft.Resources/stable/2020-06-01/resources.json
// - ARM URI: /subscriptions/{subscriptionId}/resourcegroups/{resourceGroupName}
type ResourceGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ResourceGroup `json:"items"`
}

// +kubebuilder:validation:Enum={"2020-06-01"}
type APIVersion string

const APIVersion_Value = APIVersion("2020-06-01")

type ResourceGroup_Spec struct {
	// +kubebuilder:validation:MaxLength=90
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Location: The location of the resource group. It cannot be changed after the resource group has been created. It must be
	// one of the supported Azure locations.
	Location *string `json:"location,omitempty"`

	// ManagedBy: The ID of the resource that manages this resource group.
	ManagedBy *string `json:"managedBy,omitempty"`

	// Tags: The tags attached to the resource group.
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &ResourceGroup_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (group *ResourceGroup_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if group == nil {
		return nil, nil
	}
	result := &ResourceGroup_Spec_ARM{}

	// Set property ‘Location’:
	if group.Location != nil {
		location := *group.Location
		result.Location = &location
	}

	// Set property ‘ManagedBy’:
	if group.ManagedBy != nil {
		managedBy := *group.ManagedBy
		result.ManagedBy = &managedBy
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Tags’:
	if group.Tags != nil {
		result.Tags = make(map[string]string, len(group.Tags))
		for key, value := range group.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (group *ResourceGroup_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ResourceGroup_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (group *ResourceGroup_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroup_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ResourceGroup_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	group.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		group.Location = &location
	}

	// Set property ‘ManagedBy’:
	if typedInput.ManagedBy != nil {
		managedBy := *typedInput.ManagedBy
		group.ManagedBy = &managedBy
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		group.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			group.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &ResourceGroup_Spec{}

// ConvertSpecFrom populates our ResourceGroup_Spec from the provided source
func (group *ResourceGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1api20200601s.ResourceGroup_Spec)
	if ok {
		// Populate our instance from source
		return group.AssignProperties_From_ResourceGroup_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1api20200601s.ResourceGroup_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = group.AssignProperties_From_ResourceGroup_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our ResourceGroup_Spec
func (group *ResourceGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1api20200601s.ResourceGroup_Spec)
	if ok {
		// Populate destination from our instance
		return group.AssignProperties_To_ResourceGroup_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20200601s.ResourceGroup_Spec{}
	err := group.AssignProperties_To_ResourceGroup_Spec(dst)
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

// AssignProperties_From_ResourceGroup_Spec populates our ResourceGroup_Spec from the provided source ResourceGroup_Spec
func (group *ResourceGroup_Spec) AssignProperties_From_ResourceGroup_Spec(source *v1api20200601s.ResourceGroup_Spec) error {

	// AzureName
	group.AzureName = source.AzureName

	// Location
	group.Location = genruntime.ClonePointerToString(source.Location)

	// ManagedBy
	group.ManagedBy = genruntime.ClonePointerToString(source.ManagedBy)

	// Tags
	group.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_ResourceGroup_Spec populates the provided destination ResourceGroup_Spec from our ResourceGroup_Spec
func (group *ResourceGroup_Spec) AssignProperties_To_ResourceGroup_Spec(destination *v1api20200601s.ResourceGroup_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = group.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(group.Location)

	// ManagedBy
	destination.ManagedBy = genruntime.ClonePointerToString(group.ManagedBy)

	// OriginalVersion
	destination.OriginalVersion = group.OriginalVersion()

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(group.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_ResourceGroup_STATUS populates our ResourceGroup_Spec from the provided source ResourceGroup_STATUS
func (group *ResourceGroup_Spec) Initialize_From_ResourceGroup_STATUS(source *ResourceGroup_STATUS) error {

	// Location
	group.Location = genruntime.ClonePointerToString(source.Location)

	// ManagedBy
	group.ManagedBy = genruntime.ClonePointerToString(source.ManagedBy)

	// Tags
	group.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (group *ResourceGroup_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (group *ResourceGroup_Spec) SetAzureName(azureName string) { group.AzureName = azureName }

// Resource group information.
type ResourceGroup_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: The ID of the resource group.
	Id *string `json:"id,omitempty"`

	// Location: The location of the resource group. It cannot be changed after the resource group has been created. It must be
	// one of the supported Azure locations.
	Location *string `json:"location,omitempty"`

	// ManagedBy: The ID of the resource that manages this resource group.
	ManagedBy *string `json:"managedBy,omitempty"`

	// Name: The name of the resource group.
	Name *string `json:"name,omitempty"`

	// Properties: The resource group properties.
	Properties *ResourceGroupProperties_STATUS `json:"properties,omitempty"`

	// Tags: The tags attached to the resource group.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource group.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &ResourceGroup_STATUS{}

// ConvertStatusFrom populates our ResourceGroup_STATUS from the provided source
func (group *ResourceGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1api20200601s.ResourceGroup_STATUS)
	if ok {
		// Populate our instance from source
		return group.AssignProperties_From_ResourceGroup_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1api20200601s.ResourceGroup_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = group.AssignProperties_From_ResourceGroup_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our ResourceGroup_STATUS
func (group *ResourceGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1api20200601s.ResourceGroup_STATUS)
	if ok {
		// Populate destination from our instance
		return group.AssignProperties_To_ResourceGroup_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1api20200601s.ResourceGroup_STATUS{}
	err := group.AssignProperties_To_ResourceGroup_STATUS(dst)
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

var _ genruntime.FromARMConverter = &ResourceGroup_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (group *ResourceGroup_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ResourceGroup_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (group *ResourceGroup_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroup_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ResourceGroup_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		group.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		group.Location = &location
	}

	// Set property ‘ManagedBy’:
	if typedInput.ManagedBy != nil {
		managedBy := *typedInput.ManagedBy
		group.ManagedBy = &managedBy
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		group.Name = &name
	}

	// Set property ‘Properties’:
	if typedInput.Properties != nil {
		var properties1 ResourceGroupProperties_STATUS
		err := properties1.PopulateFromARM(owner, *typedInput.Properties)
		if err != nil {
			return err
		}
		properties := properties1
		group.Properties = &properties
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		group.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			group.Tags[key] = value
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		group.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_ResourceGroup_STATUS populates our ResourceGroup_STATUS from the provided source ResourceGroup_STATUS
func (group *ResourceGroup_STATUS) AssignProperties_From_ResourceGroup_STATUS(source *v1api20200601s.ResourceGroup_STATUS) error {

	// Conditions
	group.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	group.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	group.Location = genruntime.ClonePointerToString(source.Location)

	// ManagedBy
	group.ManagedBy = genruntime.ClonePointerToString(source.ManagedBy)

	// Name
	group.Name = genruntime.ClonePointerToString(source.Name)

	// Properties
	if source.Properties != nil {
		var property ResourceGroupProperties_STATUS
		err := property.AssignProperties_From_ResourceGroupProperties_STATUS(source.Properties)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_ResourceGroupProperties_STATUS() to populate field Properties")
		}
		group.Properties = &property
	} else {
		group.Properties = nil
	}

	// Tags
	group.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Type
	group.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_ResourceGroup_STATUS populates the provided destination ResourceGroup_STATUS from our ResourceGroup_STATUS
func (group *ResourceGroup_STATUS) AssignProperties_To_ResourceGroup_STATUS(destination *v1api20200601s.ResourceGroup_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(group.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(group.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(group.Location)

	// ManagedBy
	destination.ManagedBy = genruntime.ClonePointerToString(group.ManagedBy)

	// Name
	destination.Name = genruntime.ClonePointerToString(group.Name)

	// Properties
	if group.Properties != nil {
		var property v1api20200601s.ResourceGroupProperties_STATUS
		err := group.Properties.AssignProperties_To_ResourceGroupProperties_STATUS(&property)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_ResourceGroupProperties_STATUS() to populate field Properties")
		}
		destination.Properties = &property
	} else {
		destination.Properties = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(group.Tags)

	// Type
	destination.Type = genruntime.ClonePointerToString(group.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// The resource group properties.
type ResourceGroupProperties_STATUS struct {
	// ProvisioningState: The provisioning state.
	ProvisioningState *string `json:"provisioningState,omitempty"`
}

var _ genruntime.FromARMConverter = &ResourceGroupProperties_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (properties *ResourceGroupProperties_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ResourceGroupProperties_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (properties *ResourceGroupProperties_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ResourceGroupProperties_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ResourceGroupProperties_STATUS_ARM, got %T", armInput)
	}

	// Set property ‘ProvisioningState’:
	if typedInput.ProvisioningState != nil {
		provisioningState := *typedInput.ProvisioningState
		properties.ProvisioningState = &provisioningState
	}

	// No error
	return nil
}

// AssignProperties_From_ResourceGroupProperties_STATUS populates our ResourceGroupProperties_STATUS from the provided source ResourceGroupProperties_STATUS
func (properties *ResourceGroupProperties_STATUS) AssignProperties_From_ResourceGroupProperties_STATUS(source *v1api20200601s.ResourceGroupProperties_STATUS) error {

	// ProvisioningState
	properties.ProvisioningState = genruntime.ClonePointerToString(source.ProvisioningState)

	// No error
	return nil
}

// AssignProperties_To_ResourceGroupProperties_STATUS populates the provided destination ResourceGroupProperties_STATUS from our ResourceGroupProperties_STATUS
func (properties *ResourceGroupProperties_STATUS) AssignProperties_To_ResourceGroupProperties_STATUS(destination *v1api20200601s.ResourceGroupProperties_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ProvisioningState
	destination.ProvisioningState = genruntime.ClonePointerToString(properties.ProvisioningState)

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
	SchemeBuilder.Register(&ResourceGroup{}, &ResourceGroupList{})
}
