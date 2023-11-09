// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}
type ServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Servers_FirewallRule_Spec   `json:"spec,omitempty"`
	Status            Servers_FirewallRule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &ServersFirewallRule{}

// GetConditions returns the conditions of the resource
func (rule *ServersFirewallRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *ServersFirewallRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &ServersFirewallRule{}

// ConvertFrom populates our ServersFirewallRule from the provided hub ServersFirewallRule
func (rule *ServersFirewallRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.ServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_ServersFirewallRule(source)
}

// ConvertTo populates the provided hub ServersFirewallRule from our ServersFirewallRule
func (rule *ServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.ServersFirewallRule)
	if !ok {
		return fmt.Errorf("expected sql/v1api20211101storage/ServersFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_ServersFirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-sql-azure-com-v1api20211101-serversfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversfirewallrules,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.serversfirewallrules.sql.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &ServersFirewallRule{}

// Default applies defaults to the ServersFirewallRule resource
func (rule *ServersFirewallRule) Default() {
	rule.defaultImpl()
	var temp any = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *ServersFirewallRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the ServersFirewallRule resource
func (rule *ServersFirewallRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.ImportableResource = &ServersFirewallRule{}

// InitializeSpec initializes the spec for this resource from the given status
func (rule *ServersFirewallRule) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Servers_FirewallRule_STATUS); ok {
		return rule.Spec.Initialize_From_Servers_FirewallRule_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Servers_FirewallRule_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &ServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *ServersFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule ServersFirewallRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (rule *ServersFirewallRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *ServersFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *ServersFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (rule *ServersFirewallRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/firewallRules"
func (rule *ServersFirewallRule) GetType() string {
	return "Microsoft.Sql/servers/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *ServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Servers_FirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *ServersFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *ServersFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Servers_FirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Servers_FirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-sql-azure-com-v1api20211101-serversfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=sql.azure.com,resources=serversfirewallrules,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.serversfirewallrules.sql.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &ServersFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *ServersFirewallRule) ValidateCreate() (admission.Warnings, error) {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *ServersFirewallRule) ValidateDelete() (admission.Warnings, error) {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (rule *ServersFirewallRule) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (rule *ServersFirewallRule) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){rule.validateResourceReferences, rule.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (rule *ServersFirewallRule) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *ServersFirewallRule) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateOwnerReference()
		},
	}
}

// validateOwnerReference validates the owner field
func (rule *ServersFirewallRule) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(rule)
}

// validateResourceReferences validates all resource references
func (rule *ServersFirewallRule) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *ServersFirewallRule) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*ServersFirewallRule)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_ServersFirewallRule populates our ServersFirewallRule from the provided source ServersFirewallRule
func (rule *ServersFirewallRule) AssignProperties_From_ServersFirewallRule(source *v20211101s.ServersFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Servers_FirewallRule_Spec
	err := spec.AssignProperties_From_Servers_FirewallRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_FirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status Servers_FirewallRule_STATUS
	err = status.AssignProperties_From_Servers_FirewallRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Servers_FirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_ServersFirewallRule populates the provided destination ServersFirewallRule from our ServersFirewallRule
func (rule *ServersFirewallRule) AssignProperties_To_ServersFirewallRule(destination *v20211101s.ServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Servers_FirewallRule_Spec
	err := rule.Spec.AssignProperties_To_Servers_FirewallRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_FirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.Servers_FirewallRule_STATUS
	err = rule.Status.AssignProperties_To_Servers_FirewallRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Servers_FirewallRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *ServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "ServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /sql/resource-manager/Microsoft.Sql/stable/2021-11-01/FirewallRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Sql/servers/{serverName}/firewallRules/{firewallRuleName}
type ServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServersFirewallRule `json:"items"`
}

type Servers_FirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// EndIpAddress: The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to
	// startIpAddress. Use value '0.0.0.0' for all Azure-internal IP addresses.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a sql.azure.com/Server resource
	Owner *genruntime.KnownResourceReference `group:"sql.azure.com" json:"owner,omitempty" kind:"Server"`

	// StartIpAddress: The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' for all
	// Azure-internal IP addresses.
	StartIpAddress *string `json:"startIpAddress,omitempty"`
}

var _ genruntime.ARMTransformer = &Servers_FirewallRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *Servers_FirewallRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &Servers_FirewallRule_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if rule.EndIpAddress != nil || rule.StartIpAddress != nil {
		result.Properties = &ServerFirewallRuleProperties_ARM{}
	}
	if rule.EndIpAddress != nil {
		endIpAddress := *rule.EndIpAddress
		result.Properties.EndIpAddress = &endIpAddress
	}
	if rule.StartIpAddress != nil {
		startIpAddress := *rule.StartIpAddress
		result.Properties.StartIpAddress = &startIpAddress
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Servers_FirewallRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_FirewallRule_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Servers_FirewallRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_FirewallRule_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_FirewallRule_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "EndIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rule.EndIpAddress = &endIpAddress
		}
	}

	// Set property "Owner":
	rule.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "StartIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rule.StartIpAddress = &startIpAddress
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Servers_FirewallRule_Spec{}

// ConvertSpecFrom populates our Servers_FirewallRule_Spec from the provided source
func (rule *Servers_FirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Servers_FirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Servers_FirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_FirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Servers_FirewallRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Servers_FirewallRule_Spec
func (rule *Servers_FirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Servers_FirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Servers_FirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_FirewallRule_Spec{}
	err := rule.AssignProperties_To_Servers_FirewallRule_Spec(dst)
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

// AssignProperties_From_Servers_FirewallRule_Spec populates our Servers_FirewallRule_Spec from the provided source Servers_FirewallRule_Spec
func (rule *Servers_FirewallRule_Spec) AssignProperties_From_Servers_FirewallRule_Spec(source *v20211101s.Servers_FirewallRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// EndIpAddress
	rule.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// StartIpAddress
	rule.StartIpAddress = genruntime.ClonePointerToString(source.StartIpAddress)

	// No error
	return nil
}

// AssignProperties_To_Servers_FirewallRule_Spec populates the provided destination Servers_FirewallRule_Spec from our Servers_FirewallRule_Spec
func (rule *Servers_FirewallRule_Spec) AssignProperties_To_Servers_FirewallRule_Spec(destination *v20211101s.Servers_FirewallRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIpAddress
	destination.EndIpAddress = genruntime.ClonePointerToString(rule.EndIpAddress)

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

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

// Initialize_From_Servers_FirewallRule_STATUS populates our Servers_FirewallRule_Spec from the provided source Servers_FirewallRule_STATUS
func (rule *Servers_FirewallRule_Spec) Initialize_From_Servers_FirewallRule_STATUS(source *Servers_FirewallRule_STATUS) error {

	// EndIpAddress
	rule.EndIpAddress = genruntime.ClonePointerToString(source.EndIpAddress)

	// StartIpAddress
	rule.StartIpAddress = genruntime.ClonePointerToString(source.StartIpAddress)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (rule *Servers_FirewallRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *Servers_FirewallRule_Spec) SetAzureName(azureName string) { rule.AzureName = azureName }

type Servers_FirewallRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// EndIpAddress: The end IP address of the firewall rule. Must be IPv4 format. Must be greater than or equal to
	// startIpAddress. Use value '0.0.0.0' for all Azure-internal IP addresses.
	EndIpAddress *string `json:"endIpAddress,omitempty"`

	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// StartIpAddress: The start IP address of the firewall rule. Must be IPv4 format. Use value '0.0.0.0' for all
	// Azure-internal IP addresses.
	StartIpAddress *string `json:"startIpAddress,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Servers_FirewallRule_STATUS{}

// ConvertStatusFrom populates our Servers_FirewallRule_STATUS from the provided source
func (rule *Servers_FirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Servers_FirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Servers_FirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Servers_FirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Servers_FirewallRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Servers_FirewallRule_STATUS
func (rule *Servers_FirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Servers_FirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Servers_FirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Servers_FirewallRule_STATUS{}
	err := rule.AssignProperties_To_Servers_FirewallRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Servers_FirewallRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Servers_FirewallRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Servers_FirewallRule_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Servers_FirewallRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Servers_FirewallRule_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Servers_FirewallRule_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "EndIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rule.EndIpAddress = &endIpAddress
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property "StartIpAddress":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rule.StartIpAddress = &startIpAddress
		}
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Servers_FirewallRule_STATUS populates our Servers_FirewallRule_STATUS from the provided source Servers_FirewallRule_STATUS
func (rule *Servers_FirewallRule_STATUS) AssignProperties_From_Servers_FirewallRule_STATUS(source *v20211101s.Servers_FirewallRule_STATUS) error {

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

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignProperties_To_Servers_FirewallRule_STATUS populates the provided destination Servers_FirewallRule_STATUS from our Servers_FirewallRule_STATUS
func (rule *Servers_FirewallRule_STATUS) AssignProperties_To_Servers_FirewallRule_STATUS(destination *v20211101s.Servers_FirewallRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

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

func init() {
	SchemeBuilder.Register(&ServersFirewallRule{}, &ServersFirewallRuleList{})
}
