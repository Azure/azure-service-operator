// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210501

import (
	"fmt"
	alpha20210501s "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1alpha1api20210501storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	kerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Deprecated version of FlexibleServersFirewallRule. Use v1beta20210501.FlexibleServersFirewallRule instead
type FlexibleServersFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersFirewallRules_Spec `json:"spec,omitempty"`
	Status            FirewallRule_Status               `json:"status,omitempty"`
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
	// intermediate variable for conversion
	var source alpha20210501s.FlexibleServersFirewallRule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = rule.AssignPropertiesFromFlexibleServersFirewallRule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to rule")
	}

	return nil
}

// ConvertTo populates the provided hub FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20210501s.FlexibleServersFirewallRule
	err := rule.AssignPropertiesToFlexibleServersFirewallRule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from rule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-dbformysql-azure-com-v1alpha1api20210501-flexibleserversfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210501,name=default.v1alpha1api20210501.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &FlexibleServersFirewallRule{}

// Default applies defaults to the FlexibleServersFirewallRule resource
func (rule *FlexibleServersFirewallRule) Default() {
	rule.defaultImpl()
	var temp interface{} = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *FlexibleServersFirewallRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the FlexibleServersFirewallRule resource
func (rule *FlexibleServersFirewallRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &FlexibleServersFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *FlexibleServersFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-01"
func (rule FlexibleServersFirewallRule) GetAPIVersion() string {
	return string(APIVersionValue)
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/firewallRules"
func (rule *FlexibleServersFirewallRule) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *FlexibleServersFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &FirewallRule_Status{}
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
	if st, ok := status.(*FirewallRule_Status); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st FirewallRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-dbformysql-azure-com-v1alpha1api20210501-flexibleserversfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=dbformysql.azure.com,resources=flexibleserversfirewallrules,verbs=create;update,versions=v1alpha1api20210501,name=validate.v1alpha1api20210501.flexibleserversfirewallrules.dbformysql.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &FlexibleServersFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *FlexibleServersFirewallRule) ValidateCreate() error {
	validations := rule.createValidations()
	var temp interface{} = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateDelete validates the deletion of the resource
func (rule *FlexibleServersFirewallRule) ValidateDelete() error {
	validations := rule.deleteValidations()
	var temp interface{} = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation()
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// ValidateUpdate validates an update of the resource
func (rule *FlexibleServersFirewallRule) ValidateUpdate(old runtime.Object) error {
	validations := rule.updateValidations()
	var temp interface{} = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	var errs []error
	for _, validation := range validations {
		err := validation(old)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return kerrors.NewAggregate(errs)
}

// createValidations validates the creation of the resource
func (rule *FlexibleServersFirewallRule) createValidations() []func() error {
	return []func() error{rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *FlexibleServersFirewallRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (rule *FlexibleServersFirewallRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *FlexibleServersFirewallRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *FlexibleServersFirewallRule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*FlexibleServersFirewallRule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignPropertiesFromFlexibleServersFirewallRule populates our FlexibleServersFirewallRule from the provided source FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesFromFlexibleServersFirewallRule(source *alpha20210501s.FlexibleServersFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec FlexibleServersFirewallRules_Spec
	err := spec.AssignPropertiesFromFlexibleServersFirewallRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFlexibleServersFirewallRulesSpec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status FirewallRule_Status
	err = status.AssignPropertiesFromFirewallRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromFirewallRuleStatus() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRule populates the provided destination FlexibleServersFirewallRule from our FlexibleServersFirewallRule
func (rule *FlexibleServersFirewallRule) AssignPropertiesToFlexibleServersFirewallRule(destination *alpha20210501s.FlexibleServersFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20210501s.FlexibleServersFirewallRules_Spec
	err := rule.Spec.AssignPropertiesToFlexibleServersFirewallRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFlexibleServersFirewallRulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20210501s.FirewallRule_Status
	err = rule.Status.AssignPropertiesToFirewallRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToFirewallRuleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *FlexibleServersFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "FlexibleServersFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of FlexibleServersFirewallRule. Use v1beta20210501.FlexibleServersFirewallRule instead
type FlexibleServersFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersFirewallRule `json:"items"`
}

// Deprecated version of FirewallRule_Status. Use v1beta20210501.FirewallRule_Status instead
type FirewallRule_Status struct {
	// Conditions: The observed state of the resource
	Conditions     []conditions.Condition `json:"conditions,omitempty"`
	EndIpAddress   *string                `json:"endIpAddress,omitempty"`
	Id             *string                `json:"id,omitempty"`
	Name           *string                `json:"name,omitempty"`
	StartIpAddress *string                `json:"startIpAddress,omitempty"`
	SystemData     *SystemData_Status     `json:"systemData,omitempty"`
	Type           *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &FirewallRule_Status{}

// ConvertStatusFrom populates our FirewallRule_Status from the provided source
func (rule *FirewallRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20210501s.FirewallRule_Status)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromFirewallRuleStatus(src)
	}

	// Convert to an intermediate form
	src = &alpha20210501s.FirewallRule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromFirewallRuleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our FirewallRule_Status
func (rule *FirewallRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20210501s.FirewallRule_Status)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToFirewallRuleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210501s.FirewallRule_Status{}
	err := rule.AssignPropertiesToFirewallRuleStatus(dst)
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

var _ genruntime.FromARMConverter = &FirewallRule_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *FirewallRule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FirewallRule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *FirewallRule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FirewallRule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FirewallRule_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rule.EndIpAddress = &endIpAddress
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		rule.Name = &name
	}

	// Set property ‘StartIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rule.StartIpAddress = &startIpAddress
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_Status
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		rule.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromFirewallRuleStatus populates our FirewallRule_Status from the provided source FirewallRule_Status
func (rule *FirewallRule_Status) AssignPropertiesFromFirewallRuleStatus(source *alpha20210501s.FirewallRule_Status) error {

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
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		rule.SystemData = &systemDatum
	} else {
		rule.SystemData = nil
	}

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToFirewallRuleStatus populates the provided destination FirewallRule_Status from our FirewallRule_Status
func (rule *FirewallRule_Status) AssignPropertiesToFirewallRuleStatus(destination *alpha20210501s.FirewallRule_Status) error {
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

	// SystemData
	if rule.SystemData != nil {
		var systemDatum alpha20210501s.SystemData_Status
		err := rule.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
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

type FlexibleServersFirewallRules_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	EndIpAddress *string `json:"endIpAddress,omitempty"`
	Location     *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a dbformysql.azure.com/FlexibleServer resource
	Owner *genruntime.KnownResourceReference `group:"dbformysql.azure.com" json:"owner,omitempty" kind:"FlexibleServer"`

	// +kubebuilder:validation:Required
	// +kubebuilder:validation:Pattern="^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	StartIpAddress *string           `json:"startIpAddress,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &FlexibleServersFirewallRules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rules *FlexibleServersFirewallRules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rules == nil {
		return nil, nil
	}
	result := &FlexibleServersFirewallRules_SpecARM{}

	// Set property ‘Location’:
	if rules.Location != nil {
		location := *rules.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if rules.EndIpAddress != nil || rules.StartIpAddress != nil {
		result.Properties = &FirewallRulePropertiesARM{}
	}
	if rules.EndIpAddress != nil {
		endIpAddress := *rules.EndIpAddress
		result.Properties.EndIpAddress = &endIpAddress
	}
	if rules.StartIpAddress != nil {
		startIpAddress := *rules.StartIpAddress
		result.Properties.StartIpAddress = &startIpAddress
	}

	// Set property ‘Tags’:
	if rules.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range rules.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rules *FlexibleServersFirewallRules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &FlexibleServersFirewallRules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rules *FlexibleServersFirewallRules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(FlexibleServersFirewallRules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected FlexibleServersFirewallRules_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	rules.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘EndIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIpAddress != nil {
			endIpAddress := *typedInput.Properties.EndIpAddress
			rules.EndIpAddress = &endIpAddress
		}
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		rules.Location = &location
	}

	// Set property ‘Owner’:
	rules.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘StartIpAddress’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIpAddress != nil {
			startIpAddress := *typedInput.Properties.StartIpAddress
			rules.StartIpAddress = &startIpAddress
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		rules.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			rules.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &FlexibleServersFirewallRules_Spec{}

// ConvertSpecFrom populates our FlexibleServersFirewallRules_Spec from the provided source
func (rules *FlexibleServersFirewallRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20210501s.FlexibleServersFirewallRules_Spec)
	if ok {
		// Populate our instance from source
		return rules.AssignPropertiesFromFlexibleServersFirewallRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20210501s.FlexibleServersFirewallRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rules.AssignPropertiesFromFlexibleServersFirewallRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our FlexibleServersFirewallRules_Spec
func (rules *FlexibleServersFirewallRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20210501s.FlexibleServersFirewallRules_Spec)
	if ok {
		// Populate destination from our instance
		return rules.AssignPropertiesToFlexibleServersFirewallRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20210501s.FlexibleServersFirewallRules_Spec{}
	err := rules.AssignPropertiesToFlexibleServersFirewallRulesSpec(dst)
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

// AssignPropertiesFromFlexibleServersFirewallRulesSpec populates our FlexibleServersFirewallRules_Spec from the provided source FlexibleServersFirewallRules_Spec
func (rules *FlexibleServersFirewallRules_Spec) AssignPropertiesFromFlexibleServersFirewallRulesSpec(source *alpha20210501s.FlexibleServersFirewallRules_Spec) error {

	// AzureName
	rules.AzureName = source.AzureName

	// EndIpAddress
	if source.EndIpAddress != nil {
		endIpAddress := *source.EndIpAddress
		rules.EndIpAddress = &endIpAddress
	} else {
		rules.EndIpAddress = nil
	}

	// Location
	rules.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rules.Owner = &owner
	} else {
		rules.Owner = nil
	}

	// StartIpAddress
	if source.StartIpAddress != nil {
		startIpAddress := *source.StartIpAddress
		rules.StartIpAddress = &startIpAddress
	} else {
		rules.StartIpAddress = nil
	}

	// Tags
	rules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToFlexibleServersFirewallRulesSpec populates the provided destination FlexibleServersFirewallRules_Spec from our FlexibleServersFirewallRules_Spec
func (rules *FlexibleServersFirewallRules_Spec) AssignPropertiesToFlexibleServersFirewallRulesSpec(destination *alpha20210501s.FlexibleServersFirewallRules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rules.AzureName

	// EndIpAddress
	if rules.EndIpAddress != nil {
		endIpAddress := *rules.EndIpAddress
		destination.EndIpAddress = &endIpAddress
	} else {
		destination.EndIpAddress = nil
	}

	// Location
	destination.Location = genruntime.ClonePointerToString(rules.Location)

	// OriginalVersion
	destination.OriginalVersion = rules.OriginalVersion()

	// Owner
	if rules.Owner != nil {
		owner := rules.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// StartIpAddress
	if rules.StartIpAddress != nil {
		startIpAddress := *rules.StartIpAddress
		destination.StartIpAddress = &startIpAddress
	} else {
		destination.StartIpAddress = nil
	}

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

// OriginalVersion returns the original API version used to create the resource.
func (rules *FlexibleServersFirewallRules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rules *FlexibleServersFirewallRules_Spec) SetAzureName(azureName string) {
	rules.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&FlexibleServersFirewallRule{}, &FlexibleServersFirewallRuleList{})
}
