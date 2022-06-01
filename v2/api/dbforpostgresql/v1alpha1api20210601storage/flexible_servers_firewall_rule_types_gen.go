// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

import (
	"fmt"
	v20210601s "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1beta20210601storage"
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
// Storage version of v1alpha1api20210601.FlexibleServersFirewallRule
// Deprecated version of FlexibleServersFirewallRule. Use v1beta20210601.FlexibleServersFirewallRule instead
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
	source, ok := hub.(*v20210601s.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1beta20210601storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignPropertiesFromFlexibleServersFirewallRule(source)
}

// ConvertTo populates the provided hub FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20210601s.FlexibleServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected dbforpostgresql/v1beta20210601storage/FlexibleServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignPropertiesToFlexibleServersFirewallRule(destination)
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

// GetResourceKind returns the kind of the resource
func (rule *FlexibleServersFirewallRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
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
	return &FlexibleServersFirewallRule_STATUS{}
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
	if st, ok := status.(*FlexibleServersFirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FlexibleServersFirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// AssignPropertiesFromFlexibleServersFirewallRule populates our FlexibleServersFirewallRule from the provided source FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesFromFlexibleServersFirewallRule(source *v20210601s.FlexibleServersFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersFirewallRule_Spec
	err := spec.AssignPropertiesFromFlexibleServersFirewallRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status FlexibleServersFirewallRule_STATUS
	err = status.AssignPropertiesFromFlexibleServersFirewallRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule populates the provided destination FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesToFlexibleServersFirewallRule(destination *v20210601s.FlexibleServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210601s.FlexibleServersFirewallRule_Spec
	err := rule.Spec.AssignPropertiesToFlexibleServersFirewallRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersFirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210601s.FlexibleServersFirewallRule_STATUS
	err = rule.Status.AssignPropertiesToFlexibleServersFirewallRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersFirewallRule_STATUS() to populate field Status")
	}
	destination.Status = status

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
// Storage version of v1alpha1api20210601.FlexibleServersFirewallRule
// Deprecated version of FlexibleServersFirewallRule. Use v1beta20210601.FlexibleServersFirewallRule instead
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

// Storage version of v1alpha1api20210601.FlexibleServersFirewallRule_STATUS
// Deprecated version of FlexibleServersFirewallRule_STATUS. Use v1beta20210601.FlexibleServersFirewallRule_STATUS instead
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
	src, ok := source.(*v20210601s.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromFlexibleServersFirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.FlexibleServersFirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromFlexibleServersFirewallRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210601s.FlexibleServersFirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToFlexibleServersFirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.FlexibleServersFirewallRule_STATUS{}
	err := rule.AssignPropertiesToFlexibleServersFirewallRule_STATUS(dst)
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

// AssignPropertiesFromFlexibleServersFirewallRule_STATUS populates our FlexibleServersFirewallRule_STATUS from the provided source FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignPropertiesFromFlexibleServersFirewallRule_STATUS(source *v20210601s.FlexibleServersFirewallRule_STATUS) error {
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
		err := systemDatum.AssignPropertiesFromSystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemData_STATUS() to populate field SystemData")
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

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule_STATUS populates the provided destination FlexibleServersFirewallRule_STATUS from our FlexibleServersFirewallRule_STATUS
func (rule *FlexibleServersFirewallRule_STATUS) AssignPropertiesToFlexibleServersFirewallRule_STATUS(destination *v20210601s.FlexibleServersFirewallRule_STATUS) error {
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
		var systemDatum v20210601s.SystemData_STATUS
		err := rule.SystemData.AssignPropertiesToSystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemData_STATUS() to populate field SystemData")
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

	// No error
	return nil
}

// Storage version of v1alpha1api20210601.FlexibleServersFirewallRule_Spec
type FlexibleServersFirewallRule_Spec struct {
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

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRule_Spec{}

// ConvertSpecFrom populates our FlexibleServersFirewallRule_Spec from the provided source
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210601s.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromFlexibleServersFirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210601s.FlexibleServersFirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromFlexibleServersFirewallRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210601s.FlexibleServersFirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToFlexibleServersFirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210601s.FlexibleServersFirewallRule_Spec{}
	err := rule.AssignPropertiesToFlexibleServersFirewallRule_Spec(dst)
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

// AssignPropertiesFromFlexibleServersFirewallRule_Spec populates our FlexibleServersFirewallRule_Spec from the provided source FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignPropertiesFromFlexibleServersFirewallRule_Spec(source *v20210601s.FlexibleServersFirewallRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(source.PropertyBag)

	// AzureName
	rule.AzureName = source.AzureName

	// EndIpAddress
	rule.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

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

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule_Spec populates the provided destination FlexibleServersFirewallRule_Spec from our FlexibleServersFirewallRule_Spec
func (rule *FlexibleServersFirewallRule_Spec) AssignPropertiesToFlexibleServersFirewallRule_Spec(destination *v20210601s.FlexibleServersFirewallRule_Spec) error {
	// Clone the existing property bag
	propertyBag := genruntime.NewPropertyBag(rule.PropertyBag)

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIpAddress
	destination.EndIpAddress = genruntime.ClonePointerToString(rule.EndIpAddress)

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

	// No error
	return nil
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
