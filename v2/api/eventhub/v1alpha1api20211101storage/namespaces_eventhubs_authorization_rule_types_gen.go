// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101storage

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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
// Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRule
// Deprecated version of NamespacesEventhubsAuthorizationRule. Use v1beta20211101.NamespacesEventhubsAuthorizationRule instead
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsAuthorizationRules_Spec `json:"spec,omitempty"`
	Status            AuthorizationRule_Status                   `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesEventhubsAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesEventhubsAuthorizationRule{}

// ConvertFrom populates our NamespacesEventhubsAuthorizationRule from the provided hub NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.NamespacesEventhubsAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignPropertiesFromNamespacesEventhubsAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.NamespacesEventhubsAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignPropertiesToNamespacesEventhubsAuthorizationRule(destination)
}

var _ genruntime.KubernetesResource = &NamespacesEventhubsAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesEventhubsAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesEventhubsAuthorizationRule) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceKind returns the kind of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (rule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NamespacesEventhubsAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AuthorizationRule_Status); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// AssignPropertiesFromNamespacesEventhubsAuthorizationRule populates our NamespacesEventhubsAuthorizationRule from the provided source NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) AssignPropertiesFromNamespacesEventhubsAuthorizationRule(source *v20211101s.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesEventhubsAuthorizationRules_Spec
	err := spec.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status AuthorizationRule_Status
	err = status.AssignPropertiesFromAuthorizationRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromAuthorizationRuleStatus() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRule populates the provided destination NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) AssignPropertiesToNamespacesEventhubsAuthorizationRule(destination *v20211101s.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.NamespacesEventhubsAuthorizationRules_Spec
	err := rule.Spec.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.AuthorizationRule_Status
	err = rule.Status.AssignPropertiesToAuthorizationRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToAuthorizationRuleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesEventhubsAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "NamespacesEventhubsAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRule
// Deprecated version of NamespacesEventhubsAuthorizationRule. Use v1beta20211101.NamespacesEventhubsAuthorizationRule instead
type NamespacesEventhubsAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsAuthorizationRule `json:"items"`
}

// Storage version of v1alpha1api20211101.NamespacesEventhubsAuthorizationRules_Spec
type NamespacesEventhubsAuthorizationRules_Spec struct {
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner       *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Rights      []string                           `json:"rights,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsAuthorizationRules_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubsAuthorizationRules_Spec from the provided source
func (rules *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.NamespacesEventhubsAuthorizationRules_Spec)
	if ok {
		// Populate our instance from source
		return rules.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.NamespacesEventhubsAuthorizationRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rules.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsAuthorizationRules_Spec
func (rules *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.NamespacesEventhubsAuthorizationRules_Spec)
	if ok {
		// Populate destination from our instance
		return rules.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.NamespacesEventhubsAuthorizationRules_Spec{}
	err := rules.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(dst)
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

// AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec populates our NamespacesEventhubsAuthorizationRules_Spec from the provided source NamespacesEventhubsAuthorizationRules_Spec
func (rules *NamespacesEventhubsAuthorizationRules_Spec) AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(source *v20211101s.NamespacesEventhubsAuthorizationRules_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	rules.AzureName = source.AzureName

	// Location
	rules.Location = genruntime.ClonePointerToString(source.Location)

	// OriginalVersion
	rules.OriginalVersion = source.OriginalVersion

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rules.Owner = &owner
	} else {
		rules.Owner = nil
	}

	// Rights
	rules.Rights = genruntime.CloneSliceOfString(source.Rights)

	// Tags
	rules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// Update the property bag
	if len(propertyBag) > 0 {
		rules.PropertyBag = propertyBag
	} else {
		rules.PropertyBag = nil
	}

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec populates the provided destination NamespacesEventhubsAuthorizationRules_Spec from our NamespacesEventhubsAuthorizationRules_Spec
func (rules *NamespacesEventhubsAuthorizationRules_Spec) AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(destination *v20211101s.NamespacesEventhubsAuthorizationRules_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rules.PropertyBag)

	// AzureName
	destination.AzureName = rules.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(rules.Location)

	// OriginalVersion
	destination.OriginalVersion = rules.OriginalVersion

	// Owner
	if rules.Owner != nil {
		owner := rules.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Rights
	destination.Rights = genruntime.CloneSliceOfString(rules.Rights)

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(rules.Tags)

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
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
