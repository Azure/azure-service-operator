// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"fmt"
	storage "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20240801/storage"
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
// Storage version of v1api20221201.FlexibleServersFirewallRule
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersFirewallRule_Spec   `json:"spec,omitempty"`
	Status            FlexibleServersFirewallRule_STATUS `json:"status,omitempty"`
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

var _ conversion.Convertible = &FlexibleServersFirewallRule{}

// ConvertFrom populates our FlexibleServersFirewallRule from the provided hub FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1api20240801/storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_FlexibleServersFirewallRule(source)
}

// ConvertTo populates the provided hub FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1api20240801/storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_FlexibleServersFirewallRule(destination)
}

var _ configmaps.Exporter = &FlexibleServersFirewallRule{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (rule *FlexibleServersFirewallRule) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &FlexibleServersFirewallRule{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (rule *FlexibleServersFirewallRule) SecretDestinationExpressions() []*core.DestinationExpression {
	if rule.Spec.OperatorSpec == nil {
		return nil
	}
	return rule.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.KubernetesResource = &FlexibleServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *FlexibleServersFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-12-01"
func (rule FlexibleServersFirewallRule) GetAPIVersion() string {
	return "2022-12-01"
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

// GetSupportedOperations returns the operations supported by the resource
func (rule *FlexibleServersFirewallRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
func (rule *FlexibleServersFirewallRule) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *FlexibleServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FlexibleServersFirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *FlexibleServersFirewallRule) Owner() *genruntime.ResourceReference {
	if rule.Spec.Owner == nil {
		return nil
	}

	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *FlexibleServersFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*FlexibleServersFirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersFirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// AssignProperties_From_FlexibleServersFirewallRule populates our FlexibleServersFirewallRule from the provided source FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignProperties_From_FlexibleServersFirewallRule(source *storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersFirewallRule_Spec
	err := spec.AssignProperties_From_FlexibleServersFirewallRule_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status FlexibleServersFirewallRule_STATUS
	err = status.AssignProperties_From_FlexibleServersFirewallRule_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// Invoke the augmentConversionForFlexibleServersFirewallRule interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForFlexibleServersFirewallRule); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRule populates the provided destination FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignProperties_To_FlexibleServersFirewallRule(destination *storage.FlexibleServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.FlexibleServersFirewallRule_Spec
	err := rule.Spec.AssignProperties_To_FlexibleServersFirewallRule_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.FlexibleServersFirewallRule_STATUS
	err = rule.Status.AssignProperties_To_FlexibleServersFirewallRule_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// Invoke the augmentConversionForFlexibleServersFirewallRule interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForFlexibleServersFirewallRule); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *FlexibleServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion,
		Kind:    "FlexibleServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20221201.FlexibleServersFirewallRule
// Generator information:
// - Generated from: /postgresql/resource-manager/Microsoft.DBforPostgreSQL/stable/2022-12-01/FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/firewallRules/{firewallRuleName}
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

type augmentConversionForFlexibleServersFirewallRule interface {
	AssignPropertiesFrom(src *storage.FlexibleServersFirewallRule) error
	AssignPropertiesTo(dst *storage.FlexibleServersFirewallRule) error
}

// Storage version of v1api20221201.FlexibleServersFirewallRule_Spec
type FlexibleServersFirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string                                   `json:"azureName,omitempty"`
	EndIpAddress    *string                                  `json:"endIpAddress,omitempty"`
	OperatorSpec    *FlexibleServersFirewallRuleOperatorSpec `json:"operatorSpec,omitempty"`
	OriginalVersion string                                   `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbforpostgresql.azure.com/FlexibleServer resource
	Owner          *genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`
	PropertyBag    genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	StartIpAddress *string                            `json:"startIpAddress,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRule_Spec{}

// ConvertSpecFrom populates our FlexibleServersFirewallRule_Spec from the provided source
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_FlexibleServersFirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersFirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_FlexibleServersFirewallRule_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_FlexibleServersFirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersFirewallRule_Spec{}
	err := rule.AssignProperties_To_FlexibleServersFirewallRule_Spec(dst)
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

// AssignProperties_From_FlexibleServersFirewallRule_Spec populates our FlexibleServersFirewallRule_Spec from the provided source FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignProperties_From_FlexibleServersFirewallRule_Spec(source *storage.FlexibleServersFirewallRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	rule.AzureName = source.AzureName

	// EndIpAddress
	rule.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec FlexibleServersFirewallRuleOperatorSpec
		err := operatorSpec.AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec() to populate field OperatorSpec")
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

	// StartIpAddress
	rule.StartIpAddress = genruntime.ClonePointerToString(source.StartIpAddress)

	// Update the property bag
	if len(propertyBag) > 0 {
		rule.PropertyBag = propertyBag
	} else {
		rule.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersFirewallRule_Spec interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForFlexibleServersFirewallRule_Spec); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRule_Spec populates the provided destination FlexibleServersFirewallRule_Spec from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignProperties_To_FlexibleServersFirewallRule_Spec(destination *storage.FlexibleServersFirewallRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIpAddress
	destination.EndIpAddress = genruntime.ClonePointerToString(rule.EndIpAddress)

	// OperatorSpec
	if rule.OperatorSpec != nil {
		var operatorSpec storage.FlexibleServersFirewallRuleOperatorSpec
		err := rule.OperatorSpec.AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec() to populate field OperatorSpec")
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

	// StartIpAddress
	destination.StartIpAddress = genruntime.ClonePointerToString(rule.StartIpAddress)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// Invoke the augmentConversionForFlexibleServersFirewallRule_Spec interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForFlexibleServersFirewallRule_Spec); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

// Storage version of v1api20221201.FlexibleServersFirewallRule_STATUS
type FlexibleServersFirewallRule_STATUS struct {
	Conditions     []conditions.Condition `json:"conditions,omitempty"`
	EndIpAddress   *string                `json:"endIpAddress,omitempty"`
	Id             *string                `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	PropertyBag    genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartIpAddress *string                `json:"startIpAddress,omitempty"`
	SystemData     *SystemData_STATUS     `json:"systemData,omitempty"`
	Type           *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FlexibleServersFirewallRule_STATUS{}

// ConvertStatusFrom populates our FlexibleServersFirewallRule_STATUS from the provided source
func (rule *FlexibleServersFirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_FlexibleServersFirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.FlexibleServersFirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_FlexibleServersFirewallRule_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_FlexibleServersFirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.FlexibleServersFirewallRule_STATUS{}
	err := rule.AssignProperties_To_FlexibleServersFirewallRule_STATUS(dst)
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

// AssignProperties_From_FlexibleServersFirewallRule_STATUS populates our FlexibleServersFirewallRule_STATUS from the provided source FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignProperties_From_FlexibleServersFirewallRule_STATUS(source *storage.FlexibleServersFirewallRule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIpAddress
	rule.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// StartIpAddress
	rule.StartIpAddress = genruntime.ClonePointerToString(source.StartIpAddress)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
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

	// Invoke the augmentConversionForFlexibleServersFirewallRule_STATUS interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForFlexibleServersFirewallRule_STATUS); ok {
		err := augmentedRule.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRule_STATUS populates the provided destination FlexibleServersFirewallRule_STATUS from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignProperties_To_FlexibleServersFirewallRule_STATUS(destination *storage.FlexibleServersFirewallRule_STATUS) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// EndIpAddress
	destination.EndIpAddress = genruntime.ClonePointerToString(rule.EndIpAddress)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// StartIpAddress
	destination.StartIpAddress = genruntime.ClonePointerToString(rule.StartIpAddress)

	// SystemData
	if rule.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := rule.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
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

	// Invoke the augmentConversionForFlexibleServersFirewallRule_STATUS interface (if implemented) to customize the conversion
	var ruleAsAny any = rule
	if augmentedRule, ok := ruleAsAny.(augmentConversionForFlexibleServersFirewallRule_STATUS); ok {
		err := augmentedRule.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersFirewallRule_Spec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersFirewallRule_Spec) error
	AssignPropertiesTo(dst *storage.FlexibleServersFirewallRule_Spec) error
}

type augmentConversionForFlexibleServersFirewallRule_STATUS interface {
	AssignPropertiesFrom(src *storage.FlexibleServersFirewallRule_STATUS) error
	AssignPropertiesTo(dst *storage.FlexibleServersFirewallRule_STATUS) error
}

// Storage version of v1api20221201.FlexibleServersFirewallRuleOperatorSpec
// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type FlexibleServersFirewallRuleOperatorSpec struct {
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`
	PropertyBag          genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	SecretExpressions    []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec populates our FlexibleServersFirewallRuleOperatorSpec from the provided source FlexibleServersFirewallRuleOperatorSpec
func (operator *FlexibleServersFirewallRuleOperatorSpec) AssignProperties_From_FlexibleServersFirewallRuleOperatorSpec(source *storage.FlexibleServersFirewallRuleOperatorSpec) error {
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

	// Invoke the augmentConversionForFlexibleServersFirewallRuleOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersFirewallRuleOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesFrom(source)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesFrom() for conversion")
		}
	}

	// No error
	return nil
}

// AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec populates the provided destination FlexibleServersFirewallRuleOperatorSpec from our FlexibleServersFirewallRuleOperatorSpec
func (operator *FlexibleServersFirewallRuleOperatorSpec) AssignProperties_To_FlexibleServersFirewallRuleOperatorSpec(destination *storage.FlexibleServersFirewallRuleOperatorSpec) error {
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

	// Invoke the augmentConversionForFlexibleServersFirewallRuleOperatorSpec interface (if implemented) to customize the conversion
	var operatorAsAny any = operator
	if augmentedOperator, ok := operatorAsAny.(augmentConversionForFlexibleServersFirewallRuleOperatorSpec); ok {
		err := augmentedOperator.AssignPropertiesTo(destination)
		if err != nil {
			return eris.Wrap(err, "calling augmented AssignPropertiesTo() for conversion")
		}
	}

	// No error
	return nil
}

type augmentConversionForFlexibleServersFirewallRuleOperatorSpec interface {
	AssignPropertiesFrom(src *storage.FlexibleServersFirewallRuleOperatorSpec) error
	AssignPropertiesTo(dst *storage.FlexibleServersFirewallRuleOperatorSpec) error
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
