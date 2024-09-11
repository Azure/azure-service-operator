// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230501preview.ProductPolicy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}
type ProductPolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Service_Products_Policy_Spec   `json:"spec,omitempty"`
	Status            Service_Products_Policy_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ProductPolicy{}

// GetConditions returns the conditions of the resource
func (policy *ProductPolicy) GetConditions() conditions.Conditions {
	return policy.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (policy *ProductPolicy) SetConditions(conditions conditions.Conditions) {
	policy.Status.Conditions = conditions
}

var _ conversion.Convertible = &ProductPolicy{}

// ConvertFrom populates our ProductPolicy from the provided hub ProductPolicy
func (policy *ProductPolicy) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.ProductPolicy)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ProductPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_From_ProductPolicy(source)
}

// ConvertTo populates the provided hub ProductPolicy from our ProductPolicy
func (policy *ProductPolicy) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.ProductPolicy)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/ProductPolicy but received %T instead", hub)
	}

	return policy.AssignProperties_To_ProductPolicy(destination)
}

var _ genruntime.KubernetesResource = &ProductPolicy{}

// AzureName returns the Azure name of the resource (always "policy")
func (policy *ProductPolicy) AzureName() string {
	return "policy"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (policy ProductPolicy) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetResourceScope returns the scope of the resource
func (policy *ProductPolicy) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (policy *ProductPolicy) GetSpec() genruntime.ConvertibleSpec {
	return &policy.Spec
}

// GetStatus returns the status of this resource
func (policy *ProductPolicy) GetStatus() genruntime.ConvertibleStatus {
	return &policy.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (policy *ProductPolicy) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationHead,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/products/policies"
func (policy *ProductPolicy) GetType() string {
	return "Microsoft.ApiManagement/service/products/policies"
}

// NewEmptyStatus returns a new empty (blank) status
func (policy *ProductPolicy) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Service_Products_Policy_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (policy *ProductPolicy) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(policy.Spec)
	return policy.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (policy *ProductPolicy) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Service_Products_Policy_STATUS); ok {
		policy.Status = *st
		return nil
	}

	// Convert status to required version
	var st Service_Products_Policy_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// AssignProperties_From_ProductPolicy populates our ProductPolicy from the provided source ProductPolicy
func (policy *ProductPolicy) AssignProperties_From_ProductPolicy(source *storage.ProductPolicy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Service_Products_Policy_Spec
	err := spec.AssignProperties_From_Service_Products_Policy_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_Products_Policy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Service_Products_Policy_STATUS
	err = status.AssignProperties_From_Service_Products_Policy_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Service_Products_Policy_STATUS() to populate field Status")
	}
	policy.Status = status

	// Invoke the augmentConversionForProductPolicy interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForProductPolicy); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_ProductPolicy populates the provided destination ProductPolicy from our ProductPolicy
func (policy *ProductPolicy) AssignProperties_To_ProductPolicy(destination *storage.ProductPolicy) error {

	// ObjectMeta
	destination.ObjectMeta = *policy.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.Service_Products_Policy_Spec
	err := policy.Spec.AssignProperties_To_Service_Products_Policy_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_Products_Policy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Service_Products_Policy_STATUS
	err = policy.Status.AssignProperties_To_Service_Products_Policy_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Service_Products_Policy_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForProductPolicy interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForProductPolicy); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ProductPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "ProductPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501preview.ProductPolicy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}
type ProductPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductPolicy `json:"items"`
}

type augmentConversionForProductPolicy interface {
	AssignPropertiesFrom(src *storage.ProductPolicy) error
	AssignPropertiesTo(dst *storage.ProductPolicy) error
}

// Storage version of v1api20230501preview.Service_Products_Policy_Spec
type Service_Products_Policy_Spec struct {
	Format          *string `json:"format,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Product resource
	Owner       *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Product"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Value       *string                            `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Service_Products_Policy_Spec{}

// ConvertSpecFrom populates our Service_Products_Policy_Spec from the provided source
func (policy *Service_Products_Policy_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.Service_Products_Policy_Spec)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Service_Products_Policy_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_Products_Policy_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Service_Products_Policy_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Service_Products_Policy_Spec
func (policy *Service_Products_Policy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.Service_Products_Policy_Spec)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Service_Products_Policy_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_Products_Policy_Spec{}
	err := policy.AssignProperties_To_Service_Products_Policy_Spec(dst)
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

// AssignProperties_From_Service_Products_Policy_Spec populates our Service_Products_Policy_Spec from the provided source Service_Products_Policy_Spec
func (policy *Service_Products_Policy_Spec) AssignProperties_From_Service_Products_Policy_Spec(source *storage.Service_Products_Policy_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Format
	policy.Format = genruntime.ClonePointerToString(source.Format)

	// OriginalVersion
	policy.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		policy.Owner = &owner
	} else {
		policy.Owner = nil
	}

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		policy.PropertyBag = propertyBag
	} else {
		policy.PropertyBag = nil
	}

	// Invoke the augmentConversionForService_Products_Policy_Spec interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForService_Products_Policy_Spec); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Service_Products_Policy_Spec populates the provided destination Service_Products_Policy_Spec from our Service_Products_Policy_Spec
func (policy *Service_Products_Policy_Spec) AssignProperties_To_Service_Products_Policy_Spec(destination *storage.Service_Products_Policy_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(policy.PropertyBag)

	// Format
	destination.Format = genruntime.ClonePointerToString(policy.Format)

	// OriginalVersion
	destination.OriginalVersion = policy.OriginalVersion

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

	// Invoke the augmentConversionForService_Products_Policy_Spec interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForService_Products_Policy_Spec); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.Service_Products_Policy_STATUS
type Service_Products_Policy_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Format      *string                `json:"format,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Value       *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Service_Products_Policy_STATUS{}

// ConvertStatusFrom populates our Service_Products_Policy_STATUS from the provided source
func (policy *Service_Products_Policy_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.Service_Products_Policy_STATUS)
	if ok {
		// Populate our instance from source
		return policy.AssignProperties_From_Service_Products_Policy_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.Service_Products_Policy_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Service_Products_Policy_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Service_Products_Policy_STATUS
func (policy *Service_Products_Policy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.Service_Products_Policy_STATUS)
	if ok {
		// Populate destination from our instance
		return policy.AssignProperties_To_Service_Products_Policy_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.Service_Products_Policy_STATUS{}
	err := policy.AssignProperties_To_Service_Products_Policy_STATUS(dst)
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

// AssignProperties_From_Service_Products_Policy_STATUS populates our Service_Products_Policy_STATUS from the provided source Service_Products_Policy_STATUS
func (policy *Service_Products_Policy_STATUS) AssignProperties_From_Service_Products_Policy_STATUS(source *storage.Service_Products_Policy_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	policy.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Format
	policy.Format = genruntime.ClonePointerToString(source.Format)

	// Id
	policy.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	policy.Name = genruntime.ClonePointerToString(source.Name)

	// Type
	policy.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	policy.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		policy.PropertyBag = propertyBag
	} else {
		policy.PropertyBag = nil
	}

	// Invoke the augmentConversionForService_Products_Policy_STATUS interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForService_Products_Policy_STATUS); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Service_Products_Policy_STATUS populates the provided destination Service_Products_Policy_STATUS from our Service_Products_Policy_STATUS
func (policy *Service_Products_Policy_STATUS) AssignProperties_To_Service_Products_Policy_STATUS(destination *storage.Service_Products_Policy_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(policy.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(policy.Conditions)

	// Format
	destination.Format = genruntime.ClonePointerToString(policy.Format)

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

	// Invoke the augmentConversionForService_Products_Policy_STATUS interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForService_Products_Policy_STATUS); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForService_Products_Policy_Spec interface {
	AssignPropertiesFrom(src *storage.Service_Products_Policy_Spec) error
	AssignPropertiesTo(dst *storage.Service_Products_Policy_Spec) error
}

type augmentConversionForService_Products_Policy_STATUS interface {
	AssignPropertiesFrom(src *storage.Service_Products_Policy_STATUS) error
	AssignPropertiesTo(dst *storage.Service_Products_Policy_STATUS) error
}

func init() {
	SchemeBuilder.Register(&ProductPolicy{}, &ProductPolicyList{})
}
