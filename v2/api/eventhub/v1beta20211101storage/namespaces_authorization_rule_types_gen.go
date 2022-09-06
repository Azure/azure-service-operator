// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=eventhub.azure.com,resources=namespacesauthorizationrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=eventhub.azure.com,resources={namespacesauthorizationrules/status,namespacesauthorizationrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20211101.NamespacesAuthorizationRule
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_authorizationRules
type NamespacesAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_AuthorizationRules_Spec `json:"spec,omitempty"`
	Status            AuthorizationRule_STATUS           `json:"status,omitempty"`
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
	return &AuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
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
	if st, ok := status.(*AuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this NamespacesAuthorizationRule is the hub type for conversion
func (rule *NamespacesAuthorizationRule) Hub() {}

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
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_authorizationRules
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

// Storage version of v1beta20211101.AuthorizationRule_STATUS
type AuthorizationRule_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rights      []string               `json:"rights,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &AuthorizationRule_STATUS{}

// ConvertStatusFrom populates our AuthorizationRule_STATUS from the provided source
func (rule *AuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(rule)
}

// ConvertStatusTo populates the provided destination from our AuthorizationRule_STATUS
func (rule *AuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(rule)
}

// Storage version of v1beta20211101.Namespaces_AuthorizationRules_Spec
type Namespaces_AuthorizationRules_Spec struct {
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/Namespace resource
	Owner       *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"Namespace"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Rights      []string                           `json:"rights,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &Namespaces_AuthorizationRules_Spec{}

// ConvertSpecFrom populates our Namespaces_AuthorizationRules_Spec from the provided source
func (rules *Namespaces_AuthorizationRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rules {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rules)
}

// ConvertSpecTo populates the provided destination from our Namespaces_AuthorizationRules_Spec
func (rules *Namespaces_AuthorizationRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rules {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rules)
}

func init() {
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
