// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

import (
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210601.FlexibleServersFirewallRule
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_firewallRules
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersFirewallRules_Spec `json:"spec,omitempty"`
	Status            FirewallRule_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersFirewallRule{}

// GetConditions returns the conditions of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetConditions() conditions.Conditions {
	return flexibleServersFirewallRule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) SetConditions(conditions conditions.Conditions) {
	flexibleServersFirewallRule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FlexibleServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) AzureName() string {
	return flexibleServersFirewallRule.Spec.AzureName
}

// GetResourceKind returns the kind of the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &flexibleServersFirewallRule.Spec
}

// GetStatus returns the status of this resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &flexibleServersFirewallRule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(flexibleServersFirewallRule.Spec)
	return &genruntime.ResourceReference{Group: group, Kind: kind, Namespace: flexibleServersFirewallRule.Namespace, Name: flexibleServersFirewallRule.Spec.Owner.Name}
}

// SetStatus sets the status of this resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FirewallRule_Status); ok {
		flexibleServersFirewallRule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FirewallRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	flexibleServersFirewallRule.Status = st
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (flexibleServersFirewallRule *FlexibleServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: flexibleServersFirewallRule.Spec.OriginalVersion,
		Kind:    "FlexibleServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210601.FlexibleServersFirewallRule
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_firewallRules
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

//Storage version of v1alpha1api20210601.FirewallRule_Status
//Generated from:
type FirewallRule_Status struct {
	Conditions     []conditions.Condition `json:"conditions,omitempty"`
	EndIpAddress   *string                `json:"endIpAddress,omitempty"`
	Id             *string                `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartIpAddress *string                `json:"startIpAddress,omitempty"`
	SystemData     *SystemData_Status     `json:"systemData,omitempty"`
	Type           *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FirewallRule_Status{}

// ConvertStatusFrom populates our FirewallRule_Status from the provided source
func (firewallRuleStatus *FirewallRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == firewallRuleStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(firewallRuleStatus)
}

// ConvertStatusTo populates the provided destination from our FirewallRule_Status
func (firewallRuleStatus *FirewallRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == firewallRuleStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(firewallRuleStatus)
}

//Storage version of v1alpha1api20210601.FlexibleServersFirewallRules_Spec
type FlexibleServersFirewallRules_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	EndIpAddress    *string `json:"endIpAddress,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner          genruntime.KnownResourceReference `group:"microsoft.dbforpostgresql.azure.com" json:"owner" kind:"FlexibleServer"`
	PropertyBag    genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	StartIpAddress *string                           `json:"startIpAddress,omitempty"`
	Tags           map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRules_Spec{}

// ConvertSpecFrom populates our FlexibleServersFirewallRules_Spec from the provided source
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == flexibleServersFirewallRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(flexibleServersFirewallRulesSpec)
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRules_Spec
func (flexibleServersFirewallRulesSpec *FlexibleServersFirewallRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == flexibleServersFirewallRulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(flexibleServersFirewallRulesSpec)
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
