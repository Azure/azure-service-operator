// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
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
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_firewallRules
type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFirewallRules_Spec  `json:"spec,omitempty"`
	Status            RedisFirewallRule_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisFirewallRule{}

// GetConditions returns the conditions of the resource
func (redisFirewallRule *RedisFirewallRule) GetConditions() conditions.Conditions {
	return redisFirewallRule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (redisFirewallRule *RedisFirewallRule) SetConditions(conditions conditions.Conditions) {
	redisFirewallRule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisFirewallRule{}

// ConvertFrom populates our RedisFirewallRule from the provided hub RedisFirewallRule
func (redisFirewallRule *RedisFirewallRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20201201storage.RedisFirewallRule)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisFirewallRule but received %T instead", hub)
	}

	return redisFirewallRule.AssignPropertiesFromRedisFirewallRule(source)
}

// ConvertTo populates the provided hub RedisFirewallRule from our RedisFirewallRule
func (redisFirewallRule *RedisFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20201201storage.RedisFirewallRule)
	if !ok {
		return fmt.Errorf("expected storage:cache/v1alpha1api20201201storage/RedisFirewallRule but received %T instead", hub)
	}

	return redisFirewallRule.AssignPropertiesToRedisFirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redisfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &RedisFirewallRule{}

// Default applies defaults to the RedisFirewallRule resource
func (redisFirewallRule *RedisFirewallRule) Default() {
	redisFirewallRule.defaultImpl()
	var temp interface{} = redisFirewallRule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (redisFirewallRule *RedisFirewallRule) defaultAzureName() {
	if redisFirewallRule.Spec.AzureName == "" {
		redisFirewallRule.Spec.AzureName = redisFirewallRule.Name
	}
}

// defaultImpl applies the code generated defaults to the RedisFirewallRule resource
func (redisFirewallRule *RedisFirewallRule) defaultImpl() { redisFirewallRule.defaultAzureName() }

var _ genruntime.KubernetesResource = &RedisFirewallRule{}

// AzureName returns the Azure name of the resource
func (redisFirewallRule *RedisFirewallRule) AzureName() string {
	return redisFirewallRule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (redisFirewallRule RedisFirewallRule) GetAPIVersion() string {
	return "2020-12-01"
}

// GetResourceKind returns the kind of the resource
func (redisFirewallRule *RedisFirewallRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (redisFirewallRule *RedisFirewallRule) GetSpec() genruntime.ConvertibleSpec {
	return &redisFirewallRule.Spec
}

// GetStatus returns the status of this resource
func (redisFirewallRule *RedisFirewallRule) GetStatus() genruntime.ConvertibleStatus {
	return &redisFirewallRule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (redisFirewallRule *RedisFirewallRule) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (redisFirewallRule *RedisFirewallRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisFirewallRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (redisFirewallRule *RedisFirewallRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(redisFirewallRule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  redisFirewallRule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (redisFirewallRule *RedisFirewallRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisFirewallRule_Status); ok {
		redisFirewallRule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisFirewallRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	redisFirewallRule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redisfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &RedisFirewallRule{}

// ValidateCreate validates the creation of the resource
func (redisFirewallRule *RedisFirewallRule) ValidateCreate() error {
	validations := redisFirewallRule.createValidations()
	var temp interface{} = redisFirewallRule
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
func (redisFirewallRule *RedisFirewallRule) ValidateDelete() error {
	validations := redisFirewallRule.deleteValidations()
	var temp interface{} = redisFirewallRule
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
func (redisFirewallRule *RedisFirewallRule) ValidateUpdate(old runtime.Object) error {
	validations := redisFirewallRule.updateValidations()
	var temp interface{} = redisFirewallRule
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
func (redisFirewallRule *RedisFirewallRule) createValidations() []func() error {
	return []func() error{redisFirewallRule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (redisFirewallRule *RedisFirewallRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (redisFirewallRule *RedisFirewallRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return redisFirewallRule.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (redisFirewallRule *RedisFirewallRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&redisFirewallRule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromRedisFirewallRule populates our RedisFirewallRule from the provided source RedisFirewallRule
func (redisFirewallRule *RedisFirewallRule) AssignPropertiesFromRedisFirewallRule(source *v1alpha1api20201201storage.RedisFirewallRule) error {

	// ObjectMeta
	redisFirewallRule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisFirewallRules_Spec
	err := spec.AssignPropertiesFromRedisFirewallRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisFirewallRulesSpec() to populate field Spec")
	}
	redisFirewallRule.Spec = spec

	// Status
	var status RedisFirewallRule_Status
	err = status.AssignPropertiesFromRedisFirewallRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisFirewallRuleStatus() to populate field Status")
	}
	redisFirewallRule.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisFirewallRule populates the provided destination RedisFirewallRule from our RedisFirewallRule
func (redisFirewallRule *RedisFirewallRule) AssignPropertiesToRedisFirewallRule(destination *v1alpha1api20201201storage.RedisFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *redisFirewallRule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20201201storage.RedisFirewallRules_Spec
	err := redisFirewallRule.Spec.AssignPropertiesToRedisFirewallRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisFirewallRulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20201201storage.RedisFirewallRule_Status
	err = redisFirewallRule.Status.AssignPropertiesToRedisFirewallRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisFirewallRuleStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (redisFirewallRule *RedisFirewallRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: redisFirewallRule.Spec.OriginalVersion(),
		Kind:    "RedisFirewallRule",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_firewallRules
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFirewallRule `json:"items"`
}

type RedisFirewallRule_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	//Id: Fully qualified resource ID for the resource. Ex -
	///subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	//Name: The name of the resource
	Name *string `json:"name,omitempty"`

	//StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`

	//Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or
	//"Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisFirewallRule_Status{}

// ConvertStatusFrom populates our RedisFirewallRule_Status from the provided source
func (redisFirewallRuleStatus *RedisFirewallRule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisFirewallRule_Status)
	if ok {
		// Populate our instance from source
		return redisFirewallRuleStatus.AssignPropertiesFromRedisFirewallRuleStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisFirewallRule_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = redisFirewallRuleStatus.AssignPropertiesFromRedisFirewallRuleStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisFirewallRule_Status
func (redisFirewallRuleStatus *RedisFirewallRule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisFirewallRule_Status)
	if ok {
		// Populate destination from our instance
		return redisFirewallRuleStatus.AssignPropertiesToRedisFirewallRuleStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisFirewallRule_Status{}
	err := redisFirewallRuleStatus.AssignPropertiesToRedisFirewallRuleStatus(dst)
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
func (redisFirewallRuleStatus *RedisFirewallRule_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisFirewallRule_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (redisFirewallRuleStatus *RedisFirewallRule_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisFirewallRule_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisFirewallRule_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘EndIP’:
	// copying flattened property:
	if typedInput.Properties != nil {
		redisFirewallRuleStatus.EndIP = &typedInput.Properties.EndIP
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		redisFirewallRuleStatus.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		redisFirewallRuleStatus.Name = &name
	}

	// Set property ‘StartIP’:
	// copying flattened property:
	if typedInput.Properties != nil {
		redisFirewallRuleStatus.StartIP = &typedInput.Properties.StartIP
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		redisFirewallRuleStatus.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisFirewallRuleStatus populates our RedisFirewallRule_Status from the provided source RedisFirewallRule_Status
func (redisFirewallRuleStatus *RedisFirewallRule_Status) AssignPropertiesFromRedisFirewallRuleStatus(source *v1alpha1api20201201storage.RedisFirewallRule_Status) error {

	// Conditions
	redisFirewallRuleStatus.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// EndIP
	redisFirewallRuleStatus.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// Id
	redisFirewallRuleStatus.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	redisFirewallRuleStatus.Name = genruntime.ClonePointerToString(source.Name)

	// StartIP
	redisFirewallRuleStatus.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// Type
	redisFirewallRuleStatus.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRedisFirewallRuleStatus populates the provided destination RedisFirewallRule_Status from our RedisFirewallRule_Status
func (redisFirewallRuleStatus *RedisFirewallRule_Status) AssignPropertiesToRedisFirewallRuleStatus(destination *v1alpha1api20201201storage.RedisFirewallRule_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(redisFirewallRuleStatus.Conditions)

	// EndIP
	destination.EndIP = genruntime.ClonePointerToString(redisFirewallRuleStatus.EndIP)

	// Id
	destination.Id = genruntime.ClonePointerToString(redisFirewallRuleStatus.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(redisFirewallRuleStatus.Name)

	// StartIP
	destination.StartIP = genruntime.ClonePointerToString(redisFirewallRuleStatus.StartIP)

	// Type
	destination.Type = genruntime.ClonePointerToString(redisFirewallRuleStatus.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2020-12-01"}
type RedisFirewallRulesSpecAPIVersion string

const RedisFirewallRulesSpecAPIVersion20201201 = RedisFirewallRulesSpecAPIVersion("2020-12-01")

type RedisFirewallRules_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	// +kubebuilder:validation:Required
	//EndIP: highest IP address included in the range
	EndIP string `json:"endIP"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner" kind:"Redis"`

	// +kubebuilder:validation:Required
	//StartIP: lowest IP address included in the range
	StartIP string `json:"startIP"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RedisFirewallRules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if redisFirewallRulesSpec == nil {
		return nil, nil
	}
	var result RedisFirewallRules_SpecARM

	// Set property ‘Location’:
	if redisFirewallRulesSpec.Location != nil {
		location := *redisFirewallRulesSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	result.Properties.EndIP = redisFirewallRulesSpec.EndIP
	result.Properties.StartIP = redisFirewallRulesSpec.StartIP

	// Set property ‘Tags’:
	if redisFirewallRulesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range redisFirewallRulesSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisFirewallRules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisFirewallRules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisFirewallRules_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	redisFirewallRulesSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘EndIP’:
	// copying flattened property:
	redisFirewallRulesSpec.EndIP = typedInput.Properties.EndIP

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		redisFirewallRulesSpec.Location = &location
	}

	// Set property ‘Owner’:
	redisFirewallRulesSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘StartIP’:
	// copying flattened property:
	redisFirewallRulesSpec.StartIP = typedInput.Properties.StartIP

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		redisFirewallRulesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			redisFirewallRulesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisFirewallRules_Spec{}

// ConvertSpecFrom populates our RedisFirewallRules_Spec from the provided source
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20201201storage.RedisFirewallRules_Spec)
	if ok {
		// Populate our instance from source
		return redisFirewallRulesSpec.AssignPropertiesFromRedisFirewallRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20201201storage.RedisFirewallRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = redisFirewallRulesSpec.AssignPropertiesFromRedisFirewallRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisFirewallRules_Spec
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20201201storage.RedisFirewallRules_Spec)
	if ok {
		// Populate destination from our instance
		return redisFirewallRulesSpec.AssignPropertiesToRedisFirewallRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20201201storage.RedisFirewallRules_Spec{}
	err := redisFirewallRulesSpec.AssignPropertiesToRedisFirewallRulesSpec(dst)
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
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) AssignPropertiesFromRedisFirewallRulesSpec(source *v1alpha1api20201201storage.RedisFirewallRules_Spec) error {

	// AzureName
	redisFirewallRulesSpec.AzureName = source.AzureName

	// EndIP
	redisFirewallRulesSpec.EndIP = genruntime.GetOptionalStringValue(source.EndIP)

	// Location
	redisFirewallRulesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	redisFirewallRulesSpec.Owner = source.Owner.Copy()

	// StartIP
	redisFirewallRulesSpec.StartIP = genruntime.GetOptionalStringValue(source.StartIP)

	// Tags
	redisFirewallRulesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToRedisFirewallRulesSpec populates the provided destination RedisFirewallRules_Spec from our RedisFirewallRules_Spec
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) AssignPropertiesToRedisFirewallRulesSpec(destination *v1alpha1api20201201storage.RedisFirewallRules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = redisFirewallRulesSpec.AzureName

	// EndIP
	endIP := redisFirewallRulesSpec.EndIP
	destination.EndIP = &endIP

	// Location
	destination.Location = genruntime.ClonePointerToString(redisFirewallRulesSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = redisFirewallRulesSpec.OriginalVersion()

	// Owner
	destination.Owner = redisFirewallRulesSpec.Owner.Copy()

	// StartIP
	startIP := redisFirewallRulesSpec.StartIP
	destination.StartIP = &startIP

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(redisFirewallRulesSpec.Tags)

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
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (redisFirewallRulesSpec *RedisFirewallRules_Spec) SetAzureName(azureName string) {
	redisFirewallRulesSpec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&RedisFirewallRule{}, &RedisFirewallRuleList{})
}
