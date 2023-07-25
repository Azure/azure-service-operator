// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101storage"
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
// Storage version of v1beta20211101.NamespacesAuthorizationRule
// Deprecated version of NamespacesAuthorizationRule. Use v1api20211101.NamespacesAuthorizationRule instead
type NamespacesAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_AuthorizationRule_Spec   `json:"spec,omitempty"`
	Status            Namespaces_AuthorizationRule_STATUS `json:"status,omitempty"`
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
	source, ok := hub.(*v20211101s.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1api20211101storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_NamespacesAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.NamespacesAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1api20211101storage/NamespacesAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_NamespacesAuthorizationRule(destination)
}

var _ genruntime.KubernetesResource = &NamespacesAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesAuthorizationRule) GetAPIVersion() string {
	return string(APIVersion_Value)
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/authorizationRules"
func (rule *NamespacesAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Namespaces_AuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *NamespacesAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Namespaces_AuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Namespaces_AuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// AssignProperties_From_NamespacesAuthorizationRule populates our NamespacesAuthorizationRule from the provided source NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignProperties_From_NamespacesAuthorizationRule(source *v20211101s.NamespacesAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Namespaces_AuthorizationRule_Spec
	err := spec.AssignProperties_From_Namespaces_AuthorizationRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_AuthorizationRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status Namespaces_AuthorizationRule_STATUS
	err = status.AssignProperties_From_Namespaces_AuthorizationRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_AuthorizationRule_STATUS() to populate field Status")
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
func (rule *NamespacesAuthorizationRule) AssignProperties_To_NamespacesAuthorizationRule(destination *v20211101s.NamespacesAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Namespaces_AuthorizationRule_Spec
	err := rule.Spec.AssignProperties_To_Namespaces_AuthorizationRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_AuthorizationRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Namespaces_AuthorizationRule_STATUS
	err = rule.Status.AssignProperties_To_Namespaces_AuthorizationRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_AuthorizationRule_STATUS() to populate field Status")
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
// Storage version of v1beta20211101.NamespacesAuthorizationRule
// Deprecated version of NamespacesAuthorizationRule. Use v1api20211101.NamespacesAuthorizationRule instead
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

type augmentConversionForNamespacesAuthorizationRule interface {
	AssignPropertiesFrom(src *v20211101s.NamespacesAuthorizationRule) error
	AssignPropertiesTo(dst *v20211101s.NamespacesAuthorizationRule) error
}

// Storage version of v1beta20211101.Namespaces_AuthorizationRule_Spec
type Namespaces_AuthorizationRule_Spec struct {
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/Namespace resource
	Owner       *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Rights      []string                           `json:"rights,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_AuthorizationRule_Spec{}

// ConvertSpecFrom populates our Namespaces_AuthorizationRule_Spec from the provided source
func (rule *Namespaces_AuthorizationRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Namespaces_AuthorizationRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Namespaces_AuthorizationRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Namespaces_AuthorizationRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Namespaces_AuthorizationRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Namespaces_AuthorizationRule_Spec
func (rule *Namespaces_AuthorizationRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Namespaces_AuthorizationRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Namespaces_AuthorizationRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Namespaces_AuthorizationRule_Spec{}
	err := rule.AssignProperties_To_Namespaces_AuthorizationRule_Spec(dst)
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

// AssignProperties_From_Namespaces_AuthorizationRule_Spec populates our Namespaces_AuthorizationRule_Spec from the provided source Namespaces_AuthorizationRule_Spec
func (rule *Namespaces_AuthorizationRule_Spec) AssignProperties_From_Namespaces_AuthorizationRule_Spec(source *v20211101s.Namespaces_AuthorizationRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	rule.AzureName = source.AzureName

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

	// Invoke the augmentConversionForNamespaces_AuthorizationRule_Spec interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespaces_AuthorizationRule_Spec); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_AuthorizationRule_Spec populates the provided destination Namespaces_AuthorizationRule_Spec from our Namespaces_AuthorizationRule_Spec
func (rule *Namespaces_AuthorizationRule_Spec) AssignProperties_To_Namespaces_AuthorizationRule_Spec(destination *v20211101s.Namespaces_AuthorizationRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// AzureName
	destination.AzureName = rule.AzureName

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

	// Invoke the augmentConversionForNamespaces_AuthorizationRule_Spec interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespaces_AuthorizationRule_Spec); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1beta20211101.Namespaces_AuthorizationRule_STATUS
// Deprecated version of Namespaces_AuthorizationRule_STATUS. Use v1api20211101.Namespaces_AuthorizationRule_STATUS instead
type Namespaces_AuthorizationRule_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rights      []string               `json:"rights,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Namespaces_AuthorizationRule_STATUS{}

// ConvertStatusFrom populates our Namespaces_AuthorizationRule_STATUS from the provided source
func (rule *Namespaces_AuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Namespaces_AuthorizationRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Namespaces_AuthorizationRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Namespaces_AuthorizationRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Namespaces_AuthorizationRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Namespaces_AuthorizationRule_STATUS
func (rule *Namespaces_AuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Namespaces_AuthorizationRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Namespaces_AuthorizationRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Namespaces_AuthorizationRule_STATUS{}
	err := rule.AssignProperties_To_Namespaces_AuthorizationRule_STATUS(dst)
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

// AssignProperties_From_Namespaces_AuthorizationRule_STATUS populates our Namespaces_AuthorizationRule_STATUS from the provided source Namespaces_AuthorizationRule_STATUS
func (rule *Namespaces_AuthorizationRule_STATUS) AssignProperties_From_Namespaces_AuthorizationRule_STATUS(source *v20211101s.Namespaces_AuthorizationRule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	rule.Location = genruntime.ClonePointerToString(source.Location)

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

	// Invoke the augmentConversionForNamespaces_AuthorizationRule_STATUS interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespaces_AuthorizationRule_STATUS); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_AuthorizationRule_STATUS populates the provided destination Namespaces_AuthorizationRule_STATUS from our Namespaces_AuthorizationRule_STATUS
func (rule *Namespaces_AuthorizationRule_STATUS) AssignProperties_To_Namespaces_AuthorizationRule_STATUS(destination *v20211101s.Namespaces_AuthorizationRule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(rule.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// Rights
	destination.Rights = genruntime.CloneSliceOfString(rule.Rights)

	// SystemData
	if rule.SystemData != nil {
		var systemDatum v20211101s.SystemData_STATUS
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

	// Invoke the augmentConversionForNamespaces_AuthorizationRule_STATUS interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForNamespaces_AuthorizationRule_STATUS); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return errors.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForNamespaces_AuthorizationRule_Spec interface {
	AssignPropertiesFrom(src *v20211101s.Namespaces_AuthorizationRule_Spec) error
	AssignPropertiesTo(dst *v20211101s.Namespaces_AuthorizationRule_Spec) error
}

type augmentConversionForNamespaces_AuthorizationRule_STATUS interface {
	AssignPropertiesFrom(src *v20211101s.Namespaces_AuthorizationRule_STATUS) error
	AssignPropertiesTo(dst *v20211101s.Namespaces_AuthorizationRule_STATUS) error
}

func init() {
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
