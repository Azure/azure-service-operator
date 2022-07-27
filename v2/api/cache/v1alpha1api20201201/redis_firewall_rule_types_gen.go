// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	alpha20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
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
// Deprecated version of RedisFirewallRule. Use v1beta20201201.RedisFirewallRule instead
type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFirewallRules_Spec  `json:"spec,omitempty"`
	Status            RedisFirewallRule_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisFirewallRule{}

// GetConditions returns the conditions of the resource
func (rule *RedisFirewallRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *RedisFirewallRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisFirewallRule{}

// ConvertFrom populates our RedisFirewallRule from the provided hub RedisFirewallRule
func (rule *RedisFirewallRule) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20201201s.RedisFirewallRule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = rule.AssignPropertiesFromRedisFirewallRule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to rule")
	}

	return nil
}

// ConvertTo populates the provided hub RedisFirewallRule from our RedisFirewallRule
func (rule *RedisFirewallRule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20201201s.RedisFirewallRule
	err := rule.AssignPropertiesToRedisFirewallRule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from rule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redisfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RedisFirewallRule{}

// Default applies defaults to the RedisFirewallRule resource
func (rule *RedisFirewallRule) Default() {
	rule.defaultImpl()
	var temp interface{} = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *RedisFirewallRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisFirewallRule resource
func (rule *RedisFirewallRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &RedisFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *RedisFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (rule RedisFirewallRule) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetResourceScope returns the scope of the resource
func (rule *RedisFirewallRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *RedisFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *RedisFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (rule *RedisFirewallRule) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *RedisFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisFirewallRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *RedisFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *RedisFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisFirewallRule_Status); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisFirewallRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redisfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RedisFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *RedisFirewallRule) ValidateCreate() error {
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
func (rule *RedisFirewallRule) ValidateDelete() error {
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
func (rule *RedisFirewallRule) ValidateUpdate(old runtime.Object) error {
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
func (rule *RedisFirewallRule) createValidations() []func() error {
	return []func() error{rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *RedisFirewallRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (rule *RedisFirewallRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *RedisFirewallRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *RedisFirewallRule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RedisFirewallRule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignPropertiesFromRedisFirewallRule populates our RedisFirewallRule from the provided source RedisFirewallRule
func (rule *RedisFirewallRule) AssignPropertiesFromRedisFirewallRule(source *alpha20201201s.RedisFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisFirewallRules_Spec
	err := spec.AssignPropertiesFromRedisFirewallRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisFirewallRulesSpec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status RedisFirewallRule_Status
	err = status.AssignPropertiesFromRedisFirewallRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisFirewallRuleStatus() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisFirewallRule populates the provided destination RedisFirewallRule from our RedisFirewallRule
func (rule *RedisFirewallRule) AssignPropertiesToRedisFirewallRule(destination *alpha20201201s.RedisFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20201201s.RedisFirewallRules_Spec
	err := rule.Spec.AssignPropertiesToRedisFirewallRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisFirewallRulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20201201s.RedisFirewallRule_Status
	err = rule.Status.AssignPropertiesToRedisFirewallRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisFirewallRuleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *RedisFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "RedisFirewallRule",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of RedisFirewallRule. Use v1beta20201201.RedisFirewallRule instead
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFirewallRule `json:"items"`
}

// Deprecated version of RedisFirewallRule_Status. Use v1beta20201201.RedisFirewallRule_Status instead
type RedisFirewallRule_Status struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`
	EndIP      *string                `json:"endIP,omitempty"`
	Id         *string                `json:"id,omitempty"`
	Name       *string                `json:"name,omitempty"`
	StartIP    *string                `json:"startIP,omitempty"`
	Type       *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisFirewallRule_Status{}

// ConvertStatusFrom populates our RedisFirewallRule_Status from the provided source
func (rule *RedisFirewallRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20201201s.RedisFirewallRule_Status)
	if ok {
		// Populate our instance from source
		return rule.AssignPropertiesFromRedisFirewallRuleStatus(src)
	}

	// Convert to an intermediate form
	src = &alpha20201201s.RedisFirewallRule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignPropertiesFromRedisFirewallRuleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisFirewallRule_Status
func (rule *RedisFirewallRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20201201s.RedisFirewallRule_Status)
	if ok {
		// Populate destination from our instance
		return rule.AssignPropertiesToRedisFirewallRuleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20201201s.RedisFirewallRule_Status{}
	err := rule.AssignPropertiesToRedisFirewallRuleStatus(dst)
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

var _ genruntime.FromARMConverter = &RedisFirewallRule_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *RedisFirewallRule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisFirewallRule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *RedisFirewallRule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisFirewallRule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisFirewallRule_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘EndIP’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIP != nil {
			endIP := *typedInput.Properties.EndIP
			rule.EndIP = &endIP
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

	// Set property ‘StartIP’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIP != nil {
			startIP := *typedInput.Properties.StartIP
			rule.StartIP = &startIP
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisFirewallRuleStatus populates our RedisFirewallRule_Status from the provided source RedisFirewallRule_Status
func (rule *RedisFirewallRule_Status) AssignPropertiesFromRedisFirewallRuleStatus(source *alpha20201201s.RedisFirewallRule_Status) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIP
	rule.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// StartIP
	rule.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// Type
	rule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRedisFirewallRuleStatus populates the provided destination RedisFirewallRule_Status from our RedisFirewallRule_Status
func (rule *RedisFirewallRule_Status) AssignPropertiesToRedisFirewallRuleStatus(destination *alpha20201201s.RedisFirewallRule_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// EndIP
	destination.EndIP = genruntime.ClonePointerToString(rule.EndIP)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// StartIP
	destination.StartIP = genruntime.ClonePointerToString(rule.StartIP)

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

type RedisFirewallRules_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	EndIP    *string `json:"endIP,omitempty"`
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`

	// +kubebuilder:validation:Required
	StartIP *string           `json:"startIP,omitempty"`
	Tags    map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RedisFirewallRules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rules *RedisFirewallRules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rules == nil {
		return nil, nil
	}
	result := &RedisFirewallRules_SpecARM{}

	// Set property ‘Location’:
	if rules.Location != nil {
		location := *rules.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if rules.EndIP != nil || rules.StartIP != nil {
		result.Properties = &RedisFirewallRulePropertiesARM{}
	}
	if rules.EndIP != nil {
		endIP := *rules.EndIP
		result.Properties.EndIP = &endIP
	}
	if rules.StartIP != nil {
		startIP := *rules.StartIP
		result.Properties.StartIP = &startIP
	}

	// Set property ‘Tags’:
	if rules.Tags != nil {
		result.Tags = make(map[string]string, len(rules.Tags))
		for key, value := range rules.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rules *RedisFirewallRules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisFirewallRules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rules *RedisFirewallRules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisFirewallRules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisFirewallRules_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	rules.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘EndIP’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIP != nil {
			endIP := *typedInput.Properties.EndIP
			rules.EndIP = &endIP
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

	// Set property ‘StartIP’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIP != nil {
			startIP := *typedInput.Properties.StartIP
			rules.StartIP = &startIP
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		rules.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			rules.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisFirewallRules_Spec{}

// ConvertSpecFrom populates our RedisFirewallRules_Spec from the provided source
func (rules *RedisFirewallRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20201201s.RedisFirewallRules_Spec)
	if ok {
		// Populate our instance from source
		return rules.AssignPropertiesFromRedisFirewallRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20201201s.RedisFirewallRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rules.AssignPropertiesFromRedisFirewallRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisFirewallRules_Spec
func (rules *RedisFirewallRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20201201s.RedisFirewallRules_Spec)
	if ok {
		// Populate destination from our instance
		return rules.AssignPropertiesToRedisFirewallRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20201201s.RedisFirewallRules_Spec{}
	err := rules.AssignPropertiesToRedisFirewallRulesSpec(dst)
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

// AssignPropertiesFromRedisFirewallRulesSpec populates our RedisFirewallRules_Spec from the provided source RedisFirewallRules_Spec
func (rules *RedisFirewallRules_Spec) AssignPropertiesFromRedisFirewallRulesSpec(source *alpha20201201s.RedisFirewallRules_Spec) error {

	// AzureName
	rules.AzureName = source.AzureName

	// EndIP
	rules.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// Location
	rules.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rules.Owner = &owner
	} else {
		rules.Owner = nil
	}

	// StartIP
	rules.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// Tags
	rules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToRedisFirewallRulesSpec populates the provided destination RedisFirewallRules_Spec from our RedisFirewallRules_Spec
func (rules *RedisFirewallRules_Spec) AssignPropertiesToRedisFirewallRulesSpec(destination *alpha20201201s.RedisFirewallRules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rules.AzureName

	// EndIP
	destination.EndIP = genruntime.ClonePointerToString(rules.EndIP)

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

	// StartIP
	destination.StartIP = genruntime.ClonePointerToString(rules.StartIP)

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
func (rules *RedisFirewallRules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rules *RedisFirewallRules_Spec) SetAzureName(azureName string) { rules.AzureName = azureName }

func init() {
	SchemeBuilder.Register(&RedisFirewallRule{}, &RedisFirewallRuleList{})
}
