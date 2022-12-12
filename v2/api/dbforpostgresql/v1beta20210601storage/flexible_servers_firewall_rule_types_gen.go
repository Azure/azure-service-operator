// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources=flexibleserversfirewallrules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources={flexibleserversfirewallrules/status,flexibleserversfirewallrules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1beta20210601.FlexibleServersFirewallRule
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/postgresql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServers_FirewallRule_Spec   `json:"spec,omitempty"`
	Status            FlexibleServers_FirewallRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersFirewallRule{}

// GetConditions returns the conditions of the resource
func (rule *FlexibleServersFirewallRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *FlexibleServersFirewallRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FlexibleServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *FlexibleServersFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (rule FlexibleServersFirewallRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (rule *FlexibleServersFirewallRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *FlexibleServersFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *FlexibleServersFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
func (rule *FlexibleServersFirewallRule) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *FlexibleServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServers_FirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *FlexibleServersFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *FlexibleServersFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServers_FirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServers_FirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// Hub marks that this FlexibleServersFirewallRule is the hub type for conversion
func (rule *FlexibleServersFirewallRule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *FlexibleServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "FlexibleServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1beta20210601.FlexibleServersFirewallRule
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2021-06-01/postgresql.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

// Storage version of v1beta20210601.FlexibleServers_FirewallRule_Spec
type FlexibleServers_FirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	EndIpAddress    *string `json:"endIpAddress,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner          *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	StartIpAddress *string                            `json:"startIpAddress,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServers_FirewallRule_Spec{}

// ConvertSpecFrom populates our FlexibleServers_FirewallRule_Spec from the provided source
func (rule *FlexibleServers_FirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(rule)
}

// ConvertSpecTo populates the provided destination from our FlexibleServers_FirewallRule_Spec
func (rule *FlexibleServers_FirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(rule)
}

// Storage version of v1beta20210601.FlexibleServers_FirewallRule_STATUS
type FlexibleServers_FirewallRule_STATUS struct {
	Conditions     []conditions.Condition `json:"conditions,omitempty"`
	EndIpAddress   *string                `json:"endIpAddress,omitempty"`
	Id             *string                `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartIpAddress *string                `json:"startIpAddress,omitempty"`
	SystemData     *SystemData_STATUS     `json:"systemData,omitempty"`
	Type           *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServers_FirewallRule_STATUS{}

// ConvertStatusFrom populates our FlexibleServers_FirewallRule_STATUS from the provided source
func (rule *FlexibleServers_FirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(rule)
}

// ConvertStatusTo populates the provided destination from our FlexibleServers_FirewallRule_STATUS
func (rule *FlexibleServers_FirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == rule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(rule)
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
