// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"fmt"
	v20220801s "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
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
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicyfragments.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policyFragments/{id}
type PolicyFragment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Service_PolicyFragment_Spec   `json:"spec,omitempty"`
	Status            Service_PolicyFragment_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PolicyFragment{}

// GetConditions returns the conditions of the resource
func (fragment *PolicyFragment) GetConditions() conditions.Conditions {
	return fragment.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (fragment *PolicyFragment) SetConditions(conditions conditions.Conditions) {
	fragment.Status.Conditions = conditions
}

var _ conversion.Convertible = &PolicyFragment{}

// ConvertFrom populates our PolicyFragment from the provided hub PolicyFragment
func (fragment *PolicyFragment) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20220801s.PolicyFragment)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/PolicyFragment but received %T instead", hub)
	}

	return fragment.AssignProperties_From_PolicyFragment(source)
}

// ConvertTo populates the provided hub PolicyFragment from our PolicyFragment
func (fragment *PolicyFragment) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20220801s.PolicyFragment)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/PolicyFragment but received %T instead", hub)
	}

	return fragment.AssignProperties_To_PolicyFragment(destination)
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20220801-policyfragment,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=policyfragments,verbs=create;update,versions=v1api20220801,name=default.v1api20220801.policyfragments.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &PolicyFragment{}

// Default applies defaults to the PolicyFragment resource
func (fragment *PolicyFragment) Default() {
	fragment.defaultImpl()
	var temp any = fragment
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (fragment *PolicyFragment) defaultAzureName() {
	if fragment.Spec.AzureName == "" {
		fragment.Spec.AzureName = fragment.Name
	}
}

// defaultImpl applies the code generated defaults to the PolicyFragment resource
func (fragment *PolicyFragment) defaultImpl() { fragment.defaultAzureName() }

var _ genruntime.ImportableResource = &PolicyFragment{}

// InitializeSpec initializes the spec for this resource from the given status
func (fragment *PolicyFragment) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Service_PolicyFragment_STATUS); ok {
		return fragment.Spec.Initialize_From_Service_PolicyFragment_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Service_PolicyFragment_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &PolicyFragment{}

// AzureName returns the Azure name of the resource
func (fragment *PolicyFragment) AzureName() string {
	return fragment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (fragment PolicyFragment) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (fragment *PolicyFragment) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (fragment *PolicyFragment) GetSpec() genruntime.ConvertibleSpec {
	return &fragment.Spec
}

// GetStatus returns the status of this resource
func (fragment *PolicyFragment) GetStatus() genruntime.ConvertibleStatus {
	return &fragment.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (fragment *PolicyFragment) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/policyFragments"
func (fragment *PolicyFragment) GetType() string {
	return "Microsoft.ApiManagement/service/policyFragments"
}

// NewEmptyStatus returns a new empty (blank) status
func (fragment *PolicyFragment) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Service_PolicyFragment_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (fragment *PolicyFragment) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(fragment.Spec)
	return fragment.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (fragment *PolicyFragment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Service_PolicyFragment_STATUS); ok {
		fragment.Status = *st
		return nil
	}

	// Convert status to required version
	var st Service_PolicyFragment_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	fragment.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20220801-policyfragment,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=policyfragments,verbs=create;update,versions=v1api20220801,name=validate.v1api20220801.policyfragments.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &PolicyFragment{}

// ValidateCreate validates the creation of the resource
func (fragment *PolicyFragment) ValidateCreate() (admission.Warnings, error) {
	validations := fragment.createValidations()
	var temp any = fragment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (fragment *PolicyFragment) ValidateDelete() (admission.Warnings, error) {
	validations := fragment.deleteValidations()
	var temp any = fragment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (fragment *PolicyFragment) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := fragment.updateValidations()
	var temp any = fragment
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (fragment *PolicyFragment) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){fragment.validateResourceReferences, fragment.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (fragment *PolicyFragment) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (fragment *PolicyFragment) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return fragment.validateResourceReferences()
		},
		fragment.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return fragment.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (fragment *PolicyFragment) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(fragment)
}

// validateResourceReferences validates all resource references
func (fragment *PolicyFragment) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&fragment.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (fragment *PolicyFragment) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*PolicyFragment)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, fragment)
}

// AssignProperties_From_PolicyFragment populates our PolicyFragment from the provided source PolicyFragment
func (fragment *PolicyFragment) AssignProperties_From_PolicyFragment(source *v20220801s.PolicyFragment) error {

	// ObjectMeta
	fragment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Service_PolicyFragment_Spec
	err := spec.AssignProperties_From_Service_PolicyFragment_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_PolicyFragment_Spec() to populate field Spec")
	}
	fragment.Spec = spec

	// Status
	var status Service_PolicyFragment_STATUS
	err = status.AssignProperties_From_Service_PolicyFragment_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_PolicyFragment_STATUS() to populate field Status")
	}
	fragment.Status = status

	// No error
	return nil
}

// AssignProperties_To_PolicyFragment populates the provided destination PolicyFragment from our PolicyFragment
func (fragment *PolicyFragment) AssignProperties_To_PolicyFragment(destination *v20220801s.PolicyFragment) error {

	// ObjectMeta
	destination.ObjectMeta = *fragment.ObjectMeta.DeepCopy()

	// Spec
	var spec v20220801s.Service_PolicyFragment_Spec
	err := fragment.Spec.AssignProperties_To_Service_PolicyFragment_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_PolicyFragment_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20220801s.Service_PolicyFragment_STATUS
	err = fragment.Status.AssignProperties_To_Service_PolicyFragment_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_PolicyFragment_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (fragment *PolicyFragment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: fragment.Spec.OriginalVersion(),
		Kind:    "PolicyFragment",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicyfragments.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policyFragments/{id}
type PolicyFragmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyFragment `json:"items"`
}

type Service_PolicyFragment_Spec struct {
	// +kubebuilder:validation:MaxLength=80
	// +kubebuilder:validation:MinLength=1
	// +kubebuilder:validation:Pattern="(^[\\w]+$)|(^[\\w][\\w\\-]+[\\w]$)"
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:MaxLength=1000
	// +kubebuilder:validation:MinLength=0
	// Description: Policy fragment description.
	Description *string `json:"description,omitempty"`

	// Format: Format of the policy fragment content.
	Format *PolicyFragmentContractProperties_Format `json:"format,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`

	// +kubebuilder:validation:Required
	// Value: Contents of the policy fragment.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ARMTransformer = &Service_PolicyFragment_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (fragment *Service_PolicyFragment_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if fragment == nil {
		return nil, nil
	}
	result := &Service_PolicyFragment_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if fragment.Description != nil ||
		fragment.Format != nil ||
		fragment.Value != nil {
		result.Properties = &PolicyFragmentContractProperties_ARM{}
	}
	if fragment.Description != nil {
		description := *fragment.Description
		result.Properties.Description = &description
	}
	if fragment.Format != nil {
		format := *fragment.Format
		result.Properties.Format = &format
	}
	if fragment.Value != nil {
		value := *fragment.Value
		result.Properties.Value = &value
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (fragment *Service_PolicyFragment_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Service_PolicyFragment_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (fragment *Service_PolicyFragment_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Service_PolicyFragment_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Service_PolicyFragment_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	fragment.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			fragment.Description = &description
		}
	}

	// Set property "Format":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Format != nil {
			format := *typedInput.Properties.Format
			fragment.Format = &format
		}
	}

	// Set property "Owner":
	fragment.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			fragment.Value = &value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Service_PolicyFragment_Spec{}

// ConvertSpecFrom populates our Service_PolicyFragment_Spec from the provided source
func (fragment *Service_PolicyFragment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20220801s.Service_PolicyFragment_Spec)
	if ok {
		// Populate our instance from source
		return fragment.AssignProperties_From_Service_PolicyFragment_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20220801s.Service_PolicyFragment_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = fragment.AssignProperties_From_Service_PolicyFragment_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Service_PolicyFragment_Spec
func (fragment *Service_PolicyFragment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20220801s.Service_PolicyFragment_Spec)
	if ok {
		// Populate destination from our instance
		return fragment.AssignProperties_To_Service_PolicyFragment_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20220801s.Service_PolicyFragment_Spec{}
	err := fragment.AssignProperties_To_Service_PolicyFragment_Spec(dst)
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

// AssignProperties_From_Service_PolicyFragment_Spec populates our Service_PolicyFragment_Spec from the provided source Service_PolicyFragment_Spec
func (fragment *Service_PolicyFragment_Spec) AssignProperties_From_Service_PolicyFragment_Spec(source *v20220801s.Service_PolicyFragment_Spec) error {

	// AzureName
	fragment.AzureName = source.AzureName

	// Description
	if source.Description != nil {
		description := *source.Description
		fragment.Description = &description
	} else {
		fragment.Description = nil
	}

	// Format
	if source.Format != nil {
		format := PolicyFragmentContractProperties_Format(*source.Format)
		fragment.Format = &format
	} else {
		fragment.Format = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		fragment.Owner = &owner
	} else {
		fragment.Owner = nil
	}

	// Value
	fragment.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_Service_PolicyFragment_Spec populates the provided destination Service_PolicyFragment_Spec from our Service_PolicyFragment_Spec
func (fragment *Service_PolicyFragment_Spec) AssignProperties_To_Service_PolicyFragment_Spec(destination *v20220801s.Service_PolicyFragment_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = fragment.AzureName

	// Description
	if fragment.Description != nil {
		description := *fragment.Description
		destination.Description = &description
	} else {
		destination.Description = nil
	}

	// Format
	if fragment.Format != nil {
		format := string(*fragment.Format)
		destination.Format = &format
	} else {
		destination.Format = nil
	}

	// OriginalVersion
	destination.OriginalVersion = fragment.OriginalVersion()

	// Owner
	if fragment.Owner != nil {
		owner := fragment.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Value
	destination.Value = genruntime.ClonePointerToString(fragment.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Service_PolicyFragment_STATUS populates our Service_PolicyFragment_Spec from the provided source Service_PolicyFragment_STATUS
func (fragment *Service_PolicyFragment_Spec) Initialize_From_Service_PolicyFragment_STATUS(source *Service_PolicyFragment_STATUS) error {

	// Description
	if source.Description != nil {
		description := *source.Description
		fragment.Description = &description
	} else {
		fragment.Description = nil
	}

	// Format
	if source.Format != nil {
		format := PolicyFragmentContractProperties_Format(*source.Format)
		fragment.Format = &format
	} else {
		fragment.Format = nil
	}

	// Value
	fragment.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (fragment *Service_PolicyFragment_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (fragment *Service_PolicyFragment_Spec) SetAzureName(azureName string) {
	fragment.AzureName = azureName
}

type Service_PolicyFragment_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Description: Policy fragment description.
	Description *string `json:"description,omitempty"`

	// Format: Format of the policy fragment content.
	Format *PolicyFragmentContractProperties_Format_STATUS `json:"format,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Value: Contents of the policy fragment.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Service_PolicyFragment_STATUS{}

// ConvertStatusFrom populates our Service_PolicyFragment_STATUS from the provided source
func (fragment *Service_PolicyFragment_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20220801s.Service_PolicyFragment_STATUS)
	if ok {
		// Populate our instance from source
		return fragment.AssignProperties_From_Service_PolicyFragment_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20220801s.Service_PolicyFragment_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = fragment.AssignProperties_From_Service_PolicyFragment_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Service_PolicyFragment_STATUS
func (fragment *Service_PolicyFragment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20220801s.Service_PolicyFragment_STATUS)
	if ok {
		// Populate destination from our instance
		return fragment.AssignProperties_To_Service_PolicyFragment_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20220801s.Service_PolicyFragment_STATUS{}
	err := fragment.AssignProperties_To_Service_PolicyFragment_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Service_PolicyFragment_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (fragment *Service_PolicyFragment_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Service_PolicyFragment_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (fragment *Service_PolicyFragment_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Service_PolicyFragment_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Service_PolicyFragment_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Description":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Description != nil {
			description := *typedInput.Properties.Description
			fragment.Description = &description
		}
	}

	// Set property "Format":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Format != nil {
			format := *typedInput.Properties.Format
			fragment.Format = &format
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		fragment.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		fragment.Name = &name
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		fragment.Type = &typeVar
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			fragment.Value = &value
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Service_PolicyFragment_STATUS populates our Service_PolicyFragment_STATUS from the provided source Service_PolicyFragment_STATUS
func (fragment *Service_PolicyFragment_STATUS) AssignProperties_From_Service_PolicyFragment_STATUS(source *v20220801s.Service_PolicyFragment_STATUS) error {

	// Conditions
	fragment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Description
	fragment.Description = genruntime.ClonePointerToString(source.Description)

	// Format
	if source.Format != nil {
		format := PolicyFragmentContractProperties_Format_STATUS(*source.Format)
		fragment.Format = &format
	} else {
		fragment.Format = nil
	}

	// Id
	fragment.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	fragment.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	fragment.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	fragment.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_Service_PolicyFragment_STATUS populates the provided destination Service_PolicyFragment_STATUS from our Service_PolicyFragment_STATUS
func (fragment *Service_PolicyFragment_STATUS) AssignProperties_To_Service_PolicyFragment_STATUS(destination *v20220801s.Service_PolicyFragment_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(fragment.Conditions)

	// Description
	destination.Description = genruntime.ClonePointerToString(fragment.Description)

	// Format
	if fragment.Format != nil {
		format := string(*fragment.Format)
		destination.Format = &format
	} else {
		destination.Format = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(fragment.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(fragment.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(fragment.Type)

	// Value
	destination.Value = genruntime.ClonePointerToString(fragment.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"rawxml","xml"}
type PolicyFragmentContractProperties_Format string

const (
	PolicyFragmentContractProperties_Format_Rawxml = PolicyFragmentContractProperties_Format("rawxml")
	PolicyFragmentContractProperties_Format_Xml    = PolicyFragmentContractProperties_Format("xml")
)

type PolicyFragmentContractProperties_Format_STATUS string

const (
	PolicyFragmentContractProperties_Format_STATUS_Rawxml = PolicyFragmentContractProperties_Format_STATUS("rawxml")
	PolicyFragmentContractProperties_Format_STATUS_Xml    = PolicyFragmentContractProperties_Format_STATUS("xml")
)

func init() {
	SchemeBuilder.Register(&PolicyFragment{}, &PolicyFragmentList{})
}
