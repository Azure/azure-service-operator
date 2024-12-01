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
// Storage version of v1api20230501preview.PolicyFragment
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimpolicyfragments.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policyFragments/{id}
type PolicyFragment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PolicyFragment_Spec   `json:"spec,omitempty"`
	Status            PolicyFragment_STATUS `json:"status,omitempty"`
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
	source, ok := hub.(*storage.PolicyFragment)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/PolicyFragment but received %T instead", hub)
	}

	return fragment.AssignProperties_From_PolicyFragment(source)
}

// ConvertTo populates the provided hub PolicyFragment from our PolicyFragment
func (fragment *PolicyFragment) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.PolicyFragment)
	if !ok {
		return fmt.Errorf("expected apimanagement/v1api20220801/storage/PolicyFragment but received %T instead", hub)
	}

	return fragment.AssignProperties_To_PolicyFragment(destination)
}

var _ configmaps.Exporter = &PolicyFragment{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (fragment *PolicyFragment) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if fragment.Spec.OperatorSpec == nil {
		return nil
	}
	return fragment.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &PolicyFragment{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (fragment *PolicyFragment) SecretDestinationExpressions() []*core.DestinationExpression {
	if fragment.Spec.OperatorSpec == nil {
		return nil
	}
	return fragment.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &PolicyFragment{}

// AzureName returns the Azure name of the resource
func (fragment *PolicyFragment) AzureName() string {
	return fragment.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (fragment PolicyFragment) GetAPIVersion() string {
	return "2023-05-01-preview"
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
	return &PolicyFragment_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (fragment *PolicyFragment) Owner() *genruntime.ResourceReference {
	if fragment.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(fragment.Spec)
	return fragment.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (fragment *PolicyFragment) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PolicyFragment_STATUS); ok {
		fragment.Status = *st
		return nil
	}

	// Convert status to required version
	var st PolicyFragment_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	fragment.Status = st
	return nil
}

// AssignProperties_From_PolicyFragment populates our PolicyFragment from the provided source PolicyFragment
func (fragment *PolicyFragment) AssignProperties_From_PolicyFragment(source *storage.PolicyFragment) error {

	// ObjectMeta
	fragment.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec PolicyFragment_Spec
	err := spec.AssignProperties_From_PolicyFragment_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_PolicyFragment_Spec() to populate field Spec")
	}
	fragment.Spec = spec

	// Status
	var status PolicyFragment_STATUS
	err = status.AssignProperties_From_PolicyFragment_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_PolicyFragment_STATUS() to populate field Status")
	}
	fragment.Status = status

	// Invoke the augmentConversionForPolicyFragment interface (if implemented) to customize the conversion
	var fragmentAsAny any = fragment
	if augmentedFragment, ok := fragmentAsAny.(augmentConversionForPolicyFragment); ok {
		err := augmentedFragment.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PolicyFragment populates the provided destination PolicyFragment from our PolicyFragment
func (fragment *PolicyFragment) AssignProperties_To_PolicyFragment(destination *storage.PolicyFragment) error {

	// ObjectMeta
	destination.ObjectMeta = *fragment.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.PolicyFragment_Spec
	err := fragment.Spec.AssignProperties_To_PolicyFragment_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_PolicyFragment_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.PolicyFragment_STATUS
	err = fragment.Status.AssignProperties_To_PolicyFragment_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_PolicyFragment_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForPolicyFragment interface (if implemented) to customize the conversion
	var fragmentAsAny any = fragment
	if augmentedFragment, ok := fragmentAsAny.(augmentConversionForPolicyFragment); ok {
		err := augmentedFragment.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (fragment *PolicyFragment) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: fragment.Spec.OriginalVersion,
		Kind:    "PolicyFragment",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230501preview.PolicyFragment
// Generator information:
// - Generated from: /apimanagement/resource-manager/Microsoft.ApiManagement/preview/2023-05-01-preview/apimpolicyfragments.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ApiManagement/service/{serviceName}/policyFragments/{id}
type PolicyFragmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PolicyFragment `json:"items"`
}

type augmentConversionForPolicyFragment interface {
	AssignPropertiesFrom(src *storage.PolicyFragment) error
	AssignPropertiesTo(dst *storage.PolicyFragment) error
}

// Storage version of v1api20230501preview.PolicyFragment_Spec
type PolicyFragment_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                      `json:"azureName,omitempty"`
	Description     *string                     `json:"description,omitempty"`
	Format          *string                     `json:"format,omitempty"`
	OperatorSpec    *PolicyFragmentOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                      `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a apimanagement.azure.com/Service resource
	Owner       *genruntime.KnownResourceReference `group:"apimanagement.azure.com" json:"owner,omitempty" kind:"Service"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Value       *string                            `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PolicyFragment_Spec{}

// ConvertSpecFrom populates our PolicyFragment_Spec from the provided source
func (fragment *PolicyFragment_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.PolicyFragment_Spec)
	if ok {
		// Populate our instance from source
		return fragment.AssignProperties_From_PolicyFragment_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.PolicyFragment_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = fragment.AssignProperties_From_PolicyFragment_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our PolicyFragment_Spec
func (fragment *PolicyFragment_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.PolicyFragment_Spec)
	if ok {
		// Populate destination from our instance
		return fragment.AssignProperties_To_PolicyFragment_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.PolicyFragment_Spec{}
	err := fragment.AssignProperties_To_PolicyFragment_Spec(dst)
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

// AssignProperties_From_PolicyFragment_Spec populates our PolicyFragment_Spec from the provided source PolicyFragment_Spec
func (fragment *PolicyFragment_Spec) AssignProperties_From_PolicyFragment_Spec(source *storage.PolicyFragment_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	fragment.AzureName = source.AzureName

	// Description
	fragment.Description = genruntime.ClonePointerToString(source.Description)

	// Format
	fragment.Format = genruntime.ClonePointerToString(source.Format)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec PolicyFragmentOperatorSpec
		err := operatorSpec.AssignProperties_From_PolicyFragmentOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_PolicyFragmentOperatorSpec() to populate field OperatorSpec")
		}
		fragment.OperatorSpec = &operatorSpec
	} else {
		fragment.OperatorSpec = nil
	}

	// OriginalVersion
	fragment.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		fragment.Owner = &owner
	} else {
		fragment.Owner = nil
	}

	// Value
	fragment.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		fragment.PropertyBag = propertyBag
	} else {
		fragment.PropertyBag = nil
	}

	// Invoke the augmentConversionForPolicyFragment_Spec interface (if implemented) to customize the conversion
	var fragmentAsAny any = fragment
	if augmentedFragment, ok := fragmentAsAny.(augmentConversionForPolicyFragment_Spec); ok {
		err := augmentedFragment.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PolicyFragment_Spec populates the provided destination PolicyFragment_Spec from our PolicyFragment_Spec
func (fragment *PolicyFragment_Spec) AssignProperties_To_PolicyFragment_Spec(destination *storage.PolicyFragment_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(fragment.PropertyBag)

	// AzureName
	destination.AzureName = fragment.AzureName

	// Description
	destination.Description = genruntime.ClonePointerToString(fragment.Description)

	// Format
	destination.Format = genruntime.ClonePointerToString(fragment.Format)

	// OperatorSpec
	if fragment.OperatorSpec != nil {
		var operatorSpec storage.PolicyFragmentOperatorSpec
		err := fragment.OperatorSpec.AssignProperties_To_PolicyFragmentOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_PolicyFragmentOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = fragment.OriginalVersion

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

	// Invoke the augmentConversionForPolicyFragment_Spec interface (if implemented) to customize the conversion
	var fragmentAsAny any = fragment
	if augmentedFragment, ok := fragmentAsAny.(augmentConversionForPolicyFragment_Spec); ok {
		err := augmentedFragment.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20230501preview.PolicyFragment_STATUS
type PolicyFragment_STATUS struct {
	Conditions        []conditions.Condition `json:"conditions,omitempty"`
	Description       *string                `json:"description,omitempty"`
	Format            *string                `json:"format,omitempty"`
	Id                *string                `json:"id,omitempty"`
	Name              *string                `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	Type              *string                `json:"type,omitempty"`
	Value             *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PolicyFragment_STATUS{}

// ConvertStatusFrom populates our PolicyFragment_STATUS from the provided source
func (fragment *PolicyFragment_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.PolicyFragment_STATUS)
	if ok {
		// Populate our instance from source
		return fragment.AssignProperties_From_PolicyFragment_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.PolicyFragment_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = fragment.AssignProperties_From_PolicyFragment_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our PolicyFragment_STATUS
func (fragment *PolicyFragment_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.PolicyFragment_STATUS)
	if ok {
		// Populate destination from our instance
		return fragment.AssignProperties_To_PolicyFragment_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.PolicyFragment_STATUS{}
	err := fragment.AssignProperties_To_PolicyFragment_STATUS(dst)
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

// AssignProperties_From_PolicyFragment_STATUS populates our PolicyFragment_STATUS from the provided source PolicyFragment_STATUS
func (fragment *PolicyFragment_STATUS) AssignProperties_From_PolicyFragment_STATUS(source *storage.PolicyFragment_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	fragment.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Description
	fragment.Description = genruntime.ClonePointerToString(source.Description)

	// Format
	fragment.Format = genruntime.ClonePointerToString(source.Format)

	// Id
	fragment.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	fragment.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if propertyBag.Contains("ProvisioningState") {
		var provisioningState string
		err := propertyBag.Pull("ProvisioningState", &provisioningState)
		if err != nil {
			return eris.Wrap(err, "pulling 'ProvisioningState' from propertyBag")
		}

		fragment.ProvisioningState = &provisioningState
	} else {
		fragment.ProvisioningState = nil
	}

	// Type
	fragment.Type = genruntime.ClonePointerToString(source.Type)

	// Value
	fragment.Value = genruntime.ClonePointerToString(source.Value)

	// Update the property bag
	if len(propertyBag) > 0 {
		fragment.PropertyBag = propertyBag
	} else {
		fragment.PropertyBag = nil
	}

	// Invoke the augmentConversionForPolicyFragment_STATUS interface (if implemented) to customize the conversion
	var fragmentAsAny any = fragment
	if augmentedFragment, ok := fragmentAsAny.(augmentConversionForPolicyFragment_STATUS); ok {
		err := augmentedFragment.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PolicyFragment_STATUS populates the provided destination PolicyFragment_STATUS from our PolicyFragment_STATUS
func (fragment *PolicyFragment_STATUS) AssignProperties_To_PolicyFragment_STATUS(destination *storage.PolicyFragment_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(fragment.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(fragment.Conditions)

	// Description
	destination.Description = genruntime.ClonePointerToString(fragment.Description)

	// Format
	destination.Format = genruntime.ClonePointerToString(fragment.Format)

	// Id
	destination.Id = genruntime.ClonePointerToString(fragment.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(fragment.Name)

	// ProvisioningState
	if fragment.ProvisioningState != nil {
		propertyBag.Add("ProvisioningState", *fragment.ProvisioningState)
	} else {
		propertyBag.Remove("ProvisioningState")
	}

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

	// Invoke the augmentConversionForPolicyFragment_STATUS interface (if implemented) to customize the conversion
	var fragmentAsAny any = fragment
	if augmentedFragment, ok := fragmentAsAny.(augmentConversionForPolicyFragment_STATUS); ok {
		err := augmentedFragment.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForPolicyFragment_Spec interface {
	AssignPropertiesFrom(src *storage.PolicyFragment_Spec) error
	AssignPropertiesTo(dst *storage.PolicyFragment_Spec) error
}

type augmentConversionForPolicyFragment_STATUS interface {
	AssignPropertiesFrom(src *storage.PolicyFragment_STATUS) error
	AssignPropertiesTo(dst *storage.PolicyFragment_STATUS) error
}

// Storage version of v1api20230501preview.PolicyFragmentOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type PolicyFragmentOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_PolicyFragmentOperatorSpec populates our PolicyFragmentOperatorSpec from the provided source PolicyFragmentOperatorSpec
func (operator *PolicyFragmentOperatorSpec) AssignProperties_From_PolicyFragmentOperatorSpec(source *storage.PolicyFragmentOperatorSpec) error {
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

	// Invoke the augmentConversionForPolicyFragmentOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForPolicyFragmentOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_PolicyFragmentOperatorSpec populates the provided destination PolicyFragmentOperatorSpec from our PolicyFragmentOperatorSpec
func (operator *PolicyFragmentOperatorSpec) AssignProperties_To_PolicyFragmentOperatorSpec(destination *storage.PolicyFragmentOperatorSpec) error {
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

	// Invoke the augmentConversionForPolicyFragmentOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForPolicyFragmentOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForPolicyFragmentOperatorSpec interface {
	AssignPropertiesFrom(src *storage.PolicyFragmentOperatorSpec) error
	AssignPropertiesTo(dst *storage.PolicyFragmentOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&PolicyFragment{}, &PolicyFragmentList{})
}
