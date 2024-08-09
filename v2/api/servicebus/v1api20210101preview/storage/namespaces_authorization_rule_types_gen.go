// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
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
// Storage version of v1api20210101preview.NamespacesAuthorizationRule
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/AuthorizationRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}
type NamespacesAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesAuthorizationRule_Spec   `json:"spec,omitempty"`
	Status            NamespacesAuthorizationRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesAuthorizationRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesAuthorizationRule{}

// ConvertFrom populates our NamespacesAuthorizationRule from the provided hub NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected servicebus/v1api20211101/storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_NamespacesAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected servicebus/v1api20211101/storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_NamespacesAuthorizationRule(destination)
}

var _ configmaps.Exporter = &NamespacesAuthorizationRule{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (rule *NamespacesAuthorizationRule) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &NamespacesAuthorizationRule{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (rule *NamespacesAuthorizationRule) SecretDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &NamespacesAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (rule NamespacesAuthorizationRule) GetAPIVersion() string {
	return "2021-01-01-preview"
}

// GetResourceScope returns the scope of the resource
func (rule *NamespacesAuthorizationRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *NamespacesAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (rule *NamespacesAuthorizationRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ServiceBus/namespaces/AuthorizationRules"
func (rule *NamespacesAuthorizationRule) GetType() string {
	return "Microsoft.ServiceBus/namespaces/AuthorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesAuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *NamespacesAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *NamespacesAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesAuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesAuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// AssignProperties_From_NamespacesAuthorizationRule populates our NamespacesAuthorizationRule from the provided source NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignProperties_From_NamespacesAuthorizationRule(source *storage.NamespacesAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesAuthorizationRule_Spec
	err := spec.AssignProperties_From_NamespacesAuthorizationRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_NamespacesAuthorizationRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status NamespacesAuthorizationRule_STATUS
	err = status.AssignProperties_From_NamespacesAuthorizationRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_NamespacesAuthorizationRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// Invoke the augmentConversionForNamespacesAuthorizationRule interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespacesAuthorizationRule); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRule populates the provided destination NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignProperties_To_NamespacesAuthorizationRule(destination *storage.NamespacesAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.NamespacesAuthorizationRule_Spec
	err := rule.Spec.AssignProperties_To_NamespacesAuthorizationRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_NamespacesAuthorizationRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.NamespacesAuthorizationRule_STATUS
	err = rule.Status.AssignProperties_To_NamespacesAuthorizationRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_NamespacesAuthorizationRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForNamespacesAuthorizationRule interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespacesAuthorizationRule); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "NamespacesAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20210101preview.NamespacesAuthorizationRule
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/AuthorizationRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

type augmentConversionForNamespacesAuthorizationRule interface {
	AssignPropertiesFrom(src *storage.NamespacesAuthorizationRule) error
	AssignPropertiesTo(dst *storage.NamespacesAuthorizationRule) error
}

// Storage version of v1api20210101preview.NamespacesAuthorizationRule_Spec
type NamespacesAuthorizationRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                   `json:"azureName,omitempty"`
	OperatorSpec    *NamespacesAuthorizationRuleOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/Namespace resource
	Owner       *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Rights      []string                           `json:"rights,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesAuthorizationRule_Spec{}

// ConvertSpecFrom populates our NamespacesAuthorizationRule_Spec from the provided source
func (rule *NamespacesAuthorizationRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.NamespacesAuthorizationRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_NamespacesAuthorizationRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.NamespacesAuthorizationRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_NamespacesAuthorizationRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesAuthorizationRule_Spec
func (rule *NamespacesAuthorizationRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.NamespacesAuthorizationRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_NamespacesAuthorizationRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.NamespacesAuthorizationRule_Spec{}
	err := rule.AssignProperties_To_NamespacesAuthorizationRule_Spec(dst)
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

// AssignProperties_From_NamespacesAuthorizationRule_Spec populates our NamespacesAuthorizationRule_Spec from the provided source NamespacesAuthorizationRule_Spec
func (rule *NamespacesAuthorizationRule_Spec) AssignProperties_From_NamespacesAuthorizationRule_Spec(source *storage.NamespacesAuthorizationRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	rule.AzureName = source.AzureName

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec NamespacesAuthorizationRuleOperatorSpec
		err := operatorSpec.AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec() to populate field OperatorSpec")
		}
		rule.OperatorSpec = &operatorSpec
	} else {
		rule.OperatorSpec = nil
	}

	// OriginalVersion
	rule.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// Rights
	rule.Rights = genruntime.CloneSliceOfString(source.Rights)

	// Update the property bag
	if len(propertyBag) > 0 {
		rule.PropertyBag = propertyBag
	} else {
		rule.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRule_Spec interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespacesAuthorizationRule_Spec); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRule_Spec populates the provided destination NamespacesAuthorizationRule_Spec from our NamespacesAuthorizationRule_Spec
func (rule *NamespacesAuthorizationRule_Spec) AssignProperties_To_NamespacesAuthorizationRule_Spec(destination *storage.NamespacesAuthorizationRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// AzureName
	destination.AzureName = rule.AzureName

	// OperatorSpec
	if rule.OperatorSpec != nil {
		var operatorSpec storage.NamespacesAuthorizationRuleOperatorSpec
		err := rule.OperatorSpec.AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Rights
	destination.Rights = genruntime.CloneSliceOfString(rule.Rights)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRule_Spec interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespacesAuthorizationRule_Spec); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20210101preview.NamespacesAuthorizationRule_STATUS
type NamespacesAuthorizationRule_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rights      []string               `json:"rights,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesAuthorizationRule_STATUS{}

// ConvertStatusFrom populates our NamespacesAuthorizationRule_STATUS from the provided source
func (rule *NamespacesAuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.NamespacesAuthorizationRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_NamespacesAuthorizationRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.NamespacesAuthorizationRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_NamespacesAuthorizationRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our NamespacesAuthorizationRule_STATUS
func (rule *NamespacesAuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.NamespacesAuthorizationRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_NamespacesAuthorizationRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.NamespacesAuthorizationRule_STATUS{}
	err := rule.AssignProperties_To_NamespacesAuthorizationRule_STATUS(dst)
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

// AssignProperties_From_NamespacesAuthorizationRule_STATUS populates our NamespacesAuthorizationRule_STATUS from the provided source NamespacesAuthorizationRule_STATUS
func (rule *NamespacesAuthorizationRule_STATUS) AssignProperties_From_NamespacesAuthorizationRule_STATUS(source *storage.NamespacesAuthorizationRule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	if source.Location != nil {
		propertyBag.Add("Location", *source.Location)
	} else {
		propertyBag.Remove("Location")
	}

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// Rights
	rule.Rights = genruntime.CloneSliceOfString(source.Rights)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		rule.SystemData = &systemDatum
	} else {
		rule.SystemData = nil
	}

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		rule.PropertyBag = propertyBag
	} else {
		rule.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRule_STATUS interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespacesAuthorizationRule_STATUS); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRule_STATUS populates the provided destination NamespacesAuthorizationRule_STATUS from our NamespacesAuthorizationRule_STATUS
func (rule *NamespacesAuthorizationRule_STATUS) AssignProperties_To_NamespacesAuthorizationRule_STATUS(destination *storage.NamespacesAuthorizationRule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Location
	if propertyBag.Contains("Location") {
		var location string
		err := propertyBag.Pull("Location", &location)
		if err != nil {
			return errors.Wrap(err, "pulling 'Location' from propertyBag")
		}

		destination.Location = &location
	} else {
		destination.Location = nil
	}

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// Rights
	destination.Rights = genruntime.CloneSliceOfString(rule.Rights)

	// SystemData
	if rule.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := rule.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(rule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRule_STATUS interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespacesAuthorizationRule_STATUS); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespacesAuthorizationRule_Spec interface {
	AssignPropertiesFrom(src *storage.NamespacesAuthorizationRule_Spec) error
	AssignPropertiesTo(dst *storage.NamespacesAuthorizationRule_Spec) error
}

type augmentConversionForNamespacesAuthorizationRule_STATUS interface {
	AssignPropertiesFrom(src *storage.NamespacesAuthorizationRule_STATUS) error
	AssignPropertiesTo(dst *storage.NamespacesAuthorizationRule_STATUS) error
}

// Storage version of v1api20210101preview.NamespacesAuthorizationRuleOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesAuthorizationRuleOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression               `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag                      `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression               `json:"secretExpressions,omitempty"`
	Secrets              *NamespacesAuthorizationRuleOperatorSecrets `json:"secrets,omitempty"`
}

// AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec populates our NamespacesAuthorizationRuleOperatorSpec from the provided source NamespacesAuthorizationRuleOperatorSpec
func (operator *NamespacesAuthorizationRuleOperatorSpec) AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec(source *storage.NamespacesAuthorizationRuleOperatorSpec) error {
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

	// Secrets
	if source.Secrets != nil {
		var secret NamespacesAuthorizationRuleOperatorSecrets
		err := secret.AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets(source.Secrets)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets() to populate field Secrets")
		}
		operator.Secrets = &secret
	} else {
		operator.Secrets = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		operator.PropertyBag = propertyBag
	} else {
		operator.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRuleOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForNamespacesAuthorizationRuleOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec populates the provided destination NamespacesAuthorizationRuleOperatorSpec from our NamespacesAuthorizationRuleOperatorSpec
func (operator *NamespacesAuthorizationRuleOperatorSpec) AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec(destination *storage.NamespacesAuthorizationRuleOperatorSpec) error {
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

	// Secrets
	if operator.Secrets != nil {
		var secret storage.NamespacesAuthorizationRuleOperatorSecrets
		err := operator.Secrets.AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets(&secret)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets() to populate field Secrets")
		}
		destination.Secrets = &secret
	} else {
		destination.Secrets = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRuleOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForNamespacesAuthorizationRuleOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespacesAuthorizationRuleOperatorSpec interface {
	AssignPropertiesFrom(src *storage.NamespacesAuthorizationRuleOperatorSpec) error
	AssignPropertiesTo(dst *storage.NamespacesAuthorizationRuleOperatorSpec) error
}

// Storage version of v1api20210101preview.NamespacesAuthorizationRuleOperatorSecrets
type NamespacesAuthorizationRuleOperatorSecrets struct {
	PrimaryConnectionString   *genruntime.SecretDestination `json:"primaryConnectionString,omitempty"`
	PrimaryKey                *genruntime.SecretDestination `json:"primaryKey,omitempty"`
	PropertyBag               genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryConnectionString *genruntime.SecretDestination `json:"secondaryConnectionString,omitempty"`
	SecondaryKey              *genruntime.SecretDestination `json:"secondaryKey,omitempty"`
}

// AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets populates our NamespacesAuthorizationRuleOperatorSecrets from the provided source NamespacesAuthorizationRuleOperatorSecrets
func (secrets *NamespacesAuthorizationRuleOperatorSecrets) AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets(source *storage.NamespacesAuthorizationRuleOperatorSecrets) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// PrimaryConnectionString
	if source.PrimaryConnectionString != nil {
		primaryConnectionString := source.PrimaryConnectionString.Copy()
		secrets.PrimaryConnectionString = &primaryConnectionString
	} else {
		secrets.PrimaryConnectionString = nil
	}

	// PrimaryKey
	if source.PrimaryKey != nil {
		primaryKey := source.PrimaryKey.Copy()
		secrets.PrimaryKey = &primaryKey
	} else {
		secrets.PrimaryKey = nil
	}

	// SecondaryConnectionString
	if source.SecondaryConnectionString != nil {
		secondaryConnectionString := source.SecondaryConnectionString.Copy()
		secrets.SecondaryConnectionString = &secondaryConnectionString
	} else {
		secrets.SecondaryConnectionString = nil
	}

	// SecondaryKey
	if source.SecondaryKey != nil {
		secondaryKey := source.SecondaryKey.Copy()
		secrets.SecondaryKey = &secondaryKey
	} else {
		secrets.SecondaryKey = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		secrets.PropertyBag = propertyBag
	} else {
		secrets.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRuleOperatorSecrets interface (if implemented) to customize the conversion
	var secretsAsAny any = secrets
	if augmentedSecrets, ok := secretsAsAny.(augmentConversionForNamespacesAuthorizationRuleOperatorSecrets); ok {
		err := augmentedSecrets.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets populates the provided destination NamespacesAuthorizationRuleOperatorSecrets from our NamespacesAuthorizationRuleOperatorSecrets
func (secrets *NamespacesAuthorizationRuleOperatorSecrets) AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets(destination *storage.NamespacesAuthorizationRuleOperatorSecrets) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(secrets.PropertyBag)

	// PrimaryConnectionString
	if secrets.PrimaryConnectionString != nil {
		primaryConnectionString := secrets.PrimaryConnectionString.Copy()
		destination.PrimaryConnectionString = &primaryConnectionString
	} else {
		destination.PrimaryConnectionString = nil
	}

	// PrimaryKey
	if secrets.PrimaryKey != nil {
		primaryKey := secrets.PrimaryKey.Copy()
		destination.PrimaryKey = &primaryKey
	} else {
		destination.PrimaryKey = nil
	}

	// SecondaryConnectionString
	if secrets.SecondaryConnectionString != nil {
		secondaryConnectionString := secrets.SecondaryConnectionString.Copy()
		destination.SecondaryConnectionString = &secondaryConnectionString
	} else {
		destination.SecondaryConnectionString = nil
	}

	// SecondaryKey
	if secrets.SecondaryKey != nil {
		secondaryKey := secrets.SecondaryKey.Copy()
		destination.SecondaryKey = &secondaryKey
	} else {
		destination.SecondaryKey = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForNamespacesAuthorizationRuleOperatorSecrets interface (if implemented) to customize the conversion
	var secretsAsAny any = secrets
	if augmentedSecrets, ok := secretsAsAny.(augmentConversionForNamespacesAuthorizationRuleOperatorSecrets); ok {
		err := augmentedSecrets.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespacesAuthorizationRuleOperatorSecrets interface {
	AssignPropertiesFrom(src *storage.NamespacesAuthorizationRuleOperatorSecrets) error
	AssignPropertiesTo(dst *storage.NamespacesAuthorizationRuleOperatorSecrets) error
}

func init() {
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
