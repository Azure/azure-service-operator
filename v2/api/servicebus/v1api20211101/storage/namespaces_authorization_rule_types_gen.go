// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=servicebus.azure.com,resources=namespacesauthorizationrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=servicebus.azure.com,resources={namespacesauthorizationrules/status,namespacesauthorizationrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20211101.NamespacesAuthorizationRule
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/AuthorizationRules.json
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

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesAuthorizationRule) GetAPIVersion() string {
	return "2021-11-01"
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
	if rule.Spec.Owner == nil {
		return nil
	}

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
		return eris.Wrap(err, "failed to convert status")
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
// Storage version of v1api20211101.NamespacesAuthorizationRule
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/AuthorizationRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

// Storage version of v1api20211101.NamespacesAuthorizationRule_Spec
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
	if source == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rule)
}

// ConvertSpecTo populates the provided destination from our NamespacesAuthorizationRule_Spec
func (rule *NamespacesAuthorizationRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rule)
}

// Storage version of v1api20211101.NamespacesAuthorizationRule_STATUS
type NamespacesAuthorizationRule_STATUS struct {
	Conditions  []conditions.Condition `json:"conditions,omitempty"`
	Id          *string                `json:"id,omitempty"`
	Location    *string                `json:"location,omitempty"`
	Name        *string                `json:"name,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Rights      []string               `json:"rights,omitempty"`
	SystemData  *SystemData_STATUS     `json:"systemData,omitempty"`
	Type        *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesAuthorizationRule_STATUS{}

// ConvertStatusFrom populates our NamespacesAuthorizationRule_STATUS from the provided source
func (rule *NamespacesAuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(rule)
}

// ConvertStatusTo populates the provided destination from our NamespacesAuthorizationRule_STATUS
func (rule *NamespacesAuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == rule {
		return eris.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(rule)
}

// Storage version of v1api20211101.NamespacesAuthorizationRuleOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesAuthorizationRuleOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression               `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag                      `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression               `json:"secretExpressions,omitempty"`
	Secrets              *NamespacesAuthorizationRuleOperatorSecrets `json:"secrets,omitempty"`
}

// Storage version of v1api20211101.NamespacesAuthorizationRuleOperatorSecrets
type NamespacesAuthorizationRuleOperatorSecrets struct {
	PrimaryConnectionString   *genruntime.SecretDestination `json:"primaryConnectionString,omitempty"`
	PrimaryKey                *genruntime.SecretDestination `json:"primaryKey,omitempty"`
	PropertyBag               genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecondaryConnectionString *genruntime.SecretDestination `json:"secondaryConnectionString,omitempty"`
	SecondaryKey              *genruntime.SecretDestination `json:"secondaryKey,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
