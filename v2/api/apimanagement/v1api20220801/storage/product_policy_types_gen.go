// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=apimanagement.azure.com,resources=productpolicies,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=apimanagement.azure.com,resources={productpolicies/status,productpolicies/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220801.ProductPolicy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimproducts.json
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

var _ genruntime.KubernetesResource = &ProductPolicy{}

// AzureName returns the Azure name of the resource (always "policy")
func (policy *ProductPolicy) AzureName() string {
	return "policy"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-08-01"
func (policy ProductPolicy) GetAPIVersion() string {
	return string(APIVersion_Value)
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

// Hub marks that this ProductPolicy is the hub type for conversion
func (policy *ProductPolicy) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *ProductPolicy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "ProductPolicy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220801.ProductPolicy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/apimproducts.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/products/{productId}/policies/{policyId}
type ProductPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProductPolicy `json:"items"`
}

// Storage version of v1api20220801.Service_Products_Policy_Spec
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
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(policy)
}

// ConvertSpecTo populates the provided destination from our Service_Products_Policy_Spec
func (policy *Service_Products_Policy_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(policy)
}

// Storage version of v1api20220801.Service_Products_Policy_STATUS
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
	if source == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(policy)
}

// ConvertStatusTo populates the provided destination from our Service_Products_Policy_STATUS
func (policy *Service_Products_Policy_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == policy {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(policy)
}

func init() {
	SchemeBuilder.Register(&ProductPolicy{}, &ProductPolicyList{})
}
