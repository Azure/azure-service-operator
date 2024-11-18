// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/apimanagement/v1api20220801/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
// Storage version of v1api20230501preview.Policy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimpolicies.json
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

var _ configmaps.Exporter = &Policy{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (policy *Policy) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &Policy{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (policy *Policy) SecretDestinationExpressions() []*core.DestinationExpression {
	if policy.Spec.OperatorSpec == nil {
		return nil
	}
	return policy.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &Policy{}

// AzureName returns the Azure name of the resource (always "policy")
func (policy *Policy) AzureName() string {
	return "policy"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (policy Policy) GetAPIVersion() string {
	return "2023-05-01-preview"
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
		return eris.Wrap(err, "failed to convert status")
	}

	policy.Status = st
	return nil
}

// AssignProperties_From_Policy populates our Policy from the provided source Policy
func (policy *Policy) AssignProperties_From_Policy(source *storage.Policy) error {

	// ObjectMeta
	policy.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Policy_Spec
	err := spec.AssignProperties_From_Policy_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_Policy_Spec() to populate field Spec")
	}
	policy.Spec = spec

	// Status
	var status Policy_STATUS
	err = status.AssignProperties_From_Policy_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_Policy_STATUS() to populate field Status")
	}
	policy.Status = status

	// Invoke the augmentConversionForPolicy interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForPolicy); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

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
		return eris.Wrap(err, "calling AssignProperties_To_Policy_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.Policy_STATUS
	err = policy.Status.AssignProperties_To_Policy_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_Policy_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForPolicy interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForPolicy); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (policy *Policy) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: policy.Spec.OriginalVersion,
		Kind:    "Policy",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501preview.Policy
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimpolicies.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policies/{policyId}
type PolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Policy `json:"items"`
}

type augmentConversionForPolicy interface {
	AssignPropertiesFrom(src *storage.Policy) error
	AssignPropertiesTo(dst *storage.Policy) error
}

// Storage version of v1api20230501preview.Policy_Spec
type Policy_Spec struct {
	Format          *string             `json:"format,omitempty"`
	OperatorSpec    *PolicyOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string              `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner       *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Value       *string                            `json:"value,omitempty"`
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
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Policy_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
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
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_Policy_Spec populates our Policy_Spec from the provided source Policy_Spec
func (policy *Policy_Spec) AssignProperties_From_Policy_Spec(source *storage.Policy_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Format
	policy.Format = genruntime.ClonePointerToString(source.Format)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec PolicyOperatorSpec
		err := operatorSpec.AssignProperties_From_PolicyOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_PolicyOperatorSpec() to populate field OperatorSpec")
		}
		policy.OperatorSpec = &operatorSpec
	} else {
		policy.OperatorSpec = nil
	}

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

	// Invoke the augmentConversionForPolicy_Spec interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForPolicy_Spec); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Policy_Spec populates the provided destination Policy_Spec from our Policy_Spec
func (policy *Policy_Spec) AssignProperties_To_Policy_Spec(destination *storage.Policy_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(policy.PropertyBag)

	// Format
	destination.Format = genruntime.ClonePointerToString(policy.Format)

	// OperatorSpec
	if policy.OperatorSpec != nil {
		var operatorSpec storage.PolicyOperatorSpec
		err := policy.OperatorSpec.AssignProperties_To_PolicyOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_PolicyOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

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

	// Invoke the augmentConversionForPolicy_Spec interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForPolicy_Spec); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.Policy_STATUS
type Policy_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Format      *string                `json:"format,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Type        *string                `json:"type,omitempty"`
	Value       *string                `json:"value,omitempty"`
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
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = policy.AssignProperties_From_Policy_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
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
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

// AssignProperties_From_Policy_STATUS populates our Policy_STATUS from the provided source Policy_STATUS
func (policy *Policy_STATUS) AssignProperties_From_Policy_STATUS(source *storage.Policy_STATUS) error {
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

	// Invoke the augmentConversionForPolicy_STATUS interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForPolicy_STATUS); ok {
		err := augmentedPolicy.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Policy_STATUS populates the provided destination Policy_STATUS from our Policy_STATUS
func (policy *Policy_STATUS) AssignProperties_To_Policy_STATUS(destination *storage.Policy_STATUS) error {
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

	// Invoke the augmentConversionForPolicy_STATUS interface (if implemented) to customize the conversion
	var policyAsAny any = policy
	if augmentedPolicy, ok := policyAsAny.(augmentConversionForPolicy_STATUS); ok {
		err := augmentedPolicy.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForPolicy_Spec interface {
	AssignPropertiesFrom(src *storage.Policy_Spec) error
	AssignPropertiesTo(dst *storage.Policy_Spec) error
}

type augmentConversionForPolicy_STATUS interface {
	AssignPropertiesFrom(src *storage.Policy_STATUS) error
	AssignPropertiesTo(dst *storage.Policy_STATUS) error
}

// Storage version of v1api20230501preview.PolicyOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type PolicyOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_PolicyOperatorSpec populates our PolicyOperatorSpec from the provided source PolicyOperatorSpec
func (operator *PolicyOperatorSpec) AssignProperties_From_PolicyOperatorSpec(source *storage.PolicyOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForPolicyOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForPolicyOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PolicyOperatorSpec populates the provided destination PolicyOperatorSpec from our PolicyOperatorSpec
func (operator *PolicyOperatorSpec) AssignProperties_To_PolicyOperatorSpec(destination *storage.PolicyOperatorSpec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(operator.PropertyBag)

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForPolicyOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForPolicyOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForPolicyOperatorSpec interface {
	AssignPropertiesFrom(src *storage.PolicyOperatorSpec) error
	AssignPropertiesTo(dst *storage.PolicyOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&Policy{}, &PolicyList{})
}
