// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220801

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
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
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}
type Policy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Policy_Spec   `json:"spec,omitempty"`
	Status            Policy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &Policy{}

// GetConditions returns the conditions of the resource
func (policy *Policy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *Policy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &Policy{}

// ConvertFrom populates our Policy from the provided hub Policy
func (policy *Policy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.Policy)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/Policy but received %T instead", hub)
	}

	return policy.AssignProperties_From_Policy(source)
}

// ConvertTo populates the provided hub Policy from our Policy
func (policy *Policy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.Policy)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/Policy but received %T instead", hub)
	}

	return policy.AssignProperties_To_Policy(destination)
}

// +kubebuilder:webhook:path=/mutate-apimanagement-azure-com-v1api20220801-policy,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=policies,verbs=create;update,versions=v1api20220801,name=default.v1api20220801.policies.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &Policy{}

// Default applies defaults to the Policy resource
func (policy *Policy) Default() {
	policy.defaultImpl()
	var temp any = policy
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the Policy resource
func (policy *Policy) defaultImpl() {}

var _ genruntime.ImportableResource = &Policy{}

// InitializeSpec initializes the spec for this resource from the given status
func (policy *Policy) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Policy_STATUS); ok {
		return policy.Spec.Initialize_From_Policy_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Policy_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &Policy{}

// AzureName returns the Azure name of the resource (always "policy")
func (policy *Policy) AzureName() string {
	return "policy"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (policy Policy) GetAPIVersion() string {
	return "2022-08-01"
}

// GetResourceScope returns the scope of the resource
func (policy *Policy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *Policy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *Policy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *Policy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/policies"
func (policy *Policy) GetType() string {
	return "Microsoft.ApiManagement/service/policies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *Policy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Policy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *Policy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *Policy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Policy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Policy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-apimanagement-azure-com-v1api20220801-policy,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=apimanagement.azure.com,resources=policies,verbs=create;update,versions=v1api20220801,name=validate.v1api20220801.policies.apimanagement.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &Policy{}

// ValidateCreate validates the creation of the resource
func (policy *Policy) ValidateCreate() (admission.Warnings, error) {
	validations := policy.createValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (policy *Policy) ValidateDelete() (admission.Warnings, error) {
	validations := policy.deleteValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (policy *Policy) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := policy.updateValidations()
	var temp any = policy
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (policy *Policy) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){policy.validateResourceReferences, policy.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (policy *Policy) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (policy *Policy) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateResourceReferences()
		},
		policy.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return policy.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (policy *Policy) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(policy)
}

// validateResourceReferences validates all resource references
func (policy *Policy) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&policy.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (policy *Policy) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*Policy)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, policy)
}

// AssignProperties_From_Policy populates our Policy from the provided source Policy
func (policy *Policy) AssignProperties_From_Policy(source *storage.Policy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Policy_Spec
	err := spec.AssignProperties_From_Policy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Policy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Policy_STATUS
	err = status.AssignProperties_From_Policy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Policy_STATUS() to populate field Status")
	}
	policy.Status = status

	// No error
	return nil
}

// AssignProperties_To_Policy populates the provided destination Policy from our Policy
func (policy *Policy) AssignProperties_To_Policy(destination *storage.Policy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Policy_Spec
	err := policy.Spec.AssignProperties_To_Policy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Policy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Policy_STATUS
	err = policy.Status.AssignProperties_To_Policy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Policy_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *Policy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion(),
		Kind:    "Policy",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimpolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

type Policy_Spec struct {
	// Format: Format of the policyContent.
	Format *PolicyContractProperties_Format `json:"format,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`

	// +kubebuilder:validation:Required
	// Value: Contents of the Policy as defined by the format.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ARMTransformer = &Policy_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (policy *Policy_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if policy == nil {
		return nil, nil
	}
	result := &Policy_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if policy.Format != nil || policy.Value != nil {
		result.Properties = &PolicyContractProperties_ARM{}
	}
	if policy.Format != nil {
		var temp string
		temp = string(*policy.Format)
		format := PolicyContractProperties_Format_ARM(temp)
		result.Properties.Format = &format
	}
	if policy.Value != nil {
		value := *policy.Value
		result.Properties.Value = &value
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Policy_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Policy_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Policy_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Policy_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Policy_Spec_ARM, got %T", armInput)
	}

	// Set property "Format":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Format != nil {
			var temp string
			temp = string(*typedInput.Properties.Format)
			format := PolicyContractProperties_Format(temp)
			policy.Format = &format
		}
	}

	// Set property "Owner":
	policy.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			policy.Value = &value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Policy_Spec{}

// ConvertSpecFrom populates our Policy_Spec from the provided source
func (policy *Policy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Policy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Policy_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Policy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Policy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Policy_Spec
func (policy *Policy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Policy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Policy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Policy_Spec{}
	err := policy.AssignProperties_To_Policy_Spec(dst)
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

// AssignProperties_From_Policy_Spec populates our Policy_Spec from the provided source Policy_Spec
func (policy *Policy_Spec) AssignProperties_From_Policy_Spec(source *storage.Policy_Spec) error {

	// Format
	if source.Format != nil {
		format := *source.Format
		formatTemp := genruntime.ToEnum(format, policyContractProperties_Format_Values)
		policy.Format = &formatTemp
	} else {
		policy.Format = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_Policy_Spec populates the provided destination Policy_Spec from our Policy_Spec
func (policy *Policy_Spec) AssignProperties_To_Policy_Spec(destination *storage.Policy_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Format
	if policy.Format != nil {
		format := string(*policy.Format)
		destination.Format = &format
	} else {
		destination.Format = nil
	}

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion()

	// Owner
	if policy.Owner != nil {
		owner := policy.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Value
	destination.Value = genruntime.ClonePointerToString(policy.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Policy_STATUS populates our Policy_Spec from the provided source Policy_STATUS
func (policy *Policy_Spec) Initialize_From_Policy_STATUS(source *Policy_STATUS) error {

	// Format
	if source.Format != nil {
		format := genruntime.ToEnum(string(*source.Format), policyContractProperties_Format_Values)
		policy.Format = &format
	} else {
		policy.Format = nil
	}

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (policy *Policy_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

type Policy_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Format: Format of the policyContent.
	Format *PolicyContractProperties_Format_STATUS `json:"format,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`

	// Value: Contents of the Policy as defined by the format.
	Value *string `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Policy_STATUS{}

// ConvertStatusFrom populates our Policy_STATUS from the provided source
func (policy *Policy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Policy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Policy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Policy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Policy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Policy_STATUS
func (policy *Policy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Policy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Policy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Policy_STATUS{}
	err := policy.AssignProperties_To_Policy_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Policy_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (policy *Policy_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Policy_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (policy *Policy_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Policy_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Policy_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Format":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Format != nil {
			var temp string
			temp = string(*typedInput.Properties.Format)
			format := PolicyContractProperties_Format_STATUS(temp)
			policy.Format = &format
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		policy.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		policy.Name = &name
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		policy.Type = &typeVar
	}

	// Set property "Value":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.Value != nil {
			value := *typedInput.Properties.Value
			policy.Value = &value
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Policy_STATUS populates our Policy_STATUS from the provided source Policy_STATUS
func (policy *Policy_STATUS) AssignProperties_From_Policy_STATUS(source *storage.Policy_STATUS) error {

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Format
	if source.Format != nil {
		format := *source.Format
		formatTemp := genruntime.ToEnum(format, policyContractProperties_Format_STATUS_Values)
		policy.Format = &formatTemp
	} else {
		policy.Format = nil
	}

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// No error
	return nil
}

// AssignProperties_To_Policy_STATUS populates the provided destination Policy_STATUS from our Policy_STATUS
func (policy *Policy_STATUS) AssignProperties_To_Policy_STATUS(destination *storage.Policy_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// Format
	if policy.Format != nil {
		format := string(*policy.Format)
		destination.Format = &format
	} else {
		destination.Format = nil
	}

	// Id
	destination.Id = genruntime.ClonePointerToString(policy.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(policy.Name)

	// Type
	destination.Type = genruntime.ClonePointerToString(policy.Type)

	// Value
	destination.Value = genruntime.ClonePointerToString(policy.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"rawxml","rawxml-link","xml","xml-link"}
type PolicyContractProperties_Format string

const (
	PolicyContractProperties_Format_Rawxml     = PolicyContractProperties_Format("rawxml")
	PolicyContractProperties_Format_RawxmlLink = PolicyContractProperties_Format("rawxml-link")
	PolicyContractProperties_Format_Xml        = PolicyContractProperties_Format("xml")
	PolicyContractProperties_Format_XmlLink    = PolicyContractProperties_Format("xml-link")
)

// Mapping from string to PolicyContractProperties_Format
var policyContractProperties_Format_Values = map[string]PolicyContractProperties_Format{
	"rawxml":      PolicyContractProperties_Format_Rawxml,
	"rawxml-link": PolicyContractProperties_Format_RawxmlLink,
	"xml":         PolicyContractProperties_Format_Xml,
	"xml-link":    PolicyContractProperties_Format_XmlLink,
}

type PolicyContractProperties_Format_STATUS string

const (
	PolicyContractProperties_Format_STATUS_Rawxml     = PolicyContractProperties_Format_STATUS("rawxml")
	PolicyContractProperties_Format_STATUS_RawxmlLink = PolicyContractProperties_Format_STATUS("rawxml-link")
	PolicyContractProperties_Format_STATUS_Xml        = PolicyContractProperties_Format_STATUS("xml")
	PolicyContractProperties_Format_STATUS_XmlLink    = PolicyContractProperties_Format_STATUS("xml-link")
)

// Mapping from string to PolicyContractProperties_Format_STATUS
var policyContractProperties_Format_STATUS_Values = map[string]PolicyContractProperties_Format_STATUS{
	"rawxml":      PolicyContractProperties_Format_STATUS_Rawxml,
	"rawxml-link": PolicyContractProperties_Format_STATUS_RawxmlLink,
	"xml":         PolicyContractProperties_Format_STATUS_Xml,
	"xml-link":    PolicyContractProperties_Format_STATUS_XmlLink,
}

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
