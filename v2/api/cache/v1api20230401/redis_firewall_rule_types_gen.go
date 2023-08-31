// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230401

import (
	"fmt"
	v20230401s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230401storage"
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
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-04-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules/{ruleName}
type RedisFirewallRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Redis_FirewallRule_Spec   `json:"spec,omitempty"`
	Status            Redis_FirewallRule_STATUS `json:"status,omitempty"`
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
	source, ok := hub.(*v20230401s.RedisFirewallRule)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230401storage/RedisFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_RedisFirewallRule(source)
}

// ConvertTo populates the provided hub RedisFirewallRule from our RedisFirewallRule
func (rule *RedisFirewallRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20230401s.RedisFirewallRule)
	if !ok {
		return fmt.Errorf("expected cache/v1api20230401storage/RedisFirewallRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_RedisFirewallRule(destination)
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1api20230401-redisfirewallrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1api20230401,name=default.v1api20230401.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RedisFirewallRule{}

// Default applies defaults to the RedisFirewallRule resource
func (rule *RedisFirewallRule) Default() {
	rule.defaultImpl()
	var temp any = rule
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

var _ genruntime.ImportableResource = &RedisFirewallRule{}

// InitializeSpec initializes the spec for this resource from the given status
func (rule *RedisFirewallRule) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Redis_FirewallRule_STATUS); ok {
		return rule.Spec.Initialize_From_Redis_FirewallRule_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Redis_FirewallRule_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &RedisFirewallRule{}

// AzureName returns the Azure name of the resource
func (rule *RedisFirewallRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-04-01"
func (rule RedisFirewallRule) GetAPIVersion() string {
	return string(APIVersion_Value)
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
	return &Redis_FirewallRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
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
	if st, ok := status.(*Redis_FirewallRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Redis_FirewallRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1api20230401-redisfirewallrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redisfirewallrules,verbs=create;update,versions=v1api20230401,name=validate.v1api20230401.redisfirewallrules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RedisFirewallRule{}

// ValidateCreate validates the creation of the resource
func (rule *RedisFirewallRule) ValidateCreate() (admission.Warnings, error) {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *RedisFirewallRule) ValidateDelete() (admission.Warnings, error) {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (rule *RedisFirewallRule) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (rule *RedisFirewallRule) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *RedisFirewallRule) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *RedisFirewallRule) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *RedisFirewallRule) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *RedisFirewallRule) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*RedisFirewallRule)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_RedisFirewallRule populates our RedisFirewallRule from the provided source RedisFirewallRule
func (rule *RedisFirewallRule) AssignProperties_From_RedisFirewallRule(source *v20230401s.RedisFirewallRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Redis_FirewallRule_Spec
	err := spec.AssignProperties_From_Redis_FirewallRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_FirewallRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status Redis_FirewallRule_STATUS
	err = status.AssignProperties_From_Redis_FirewallRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Redis_FirewallRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_RedisFirewallRule populates the provided destination RedisFirewallRule from our RedisFirewallRule
func (rule *RedisFirewallRule) AssignProperties_To_RedisFirewallRule(destination *v20230401s.RedisFirewallRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20230401s.Redis_FirewallRule_Spec
	err := rule.Spec.AssignProperties_To_Redis_FirewallRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_FirewallRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20230401s.Redis_FirewallRule_STATUS
	err = rule.Status.AssignProperties_To_Redis_FirewallRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Redis_FirewallRule_STATUS() to populate field Status")
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
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-04-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{cacheName}/firewallRules/{ruleName}
type RedisFirewallRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisFirewallRule `json:"items"`
}

type Redis_FirewallRule_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`

	// +kubebuilder:validation:Required
	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`
}

var _ genruntime.ARMTransformer = &Redis_FirewallRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *Redis_FirewallRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &Redis_FirewallRule_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if rule.EndIP != nil || rule.StartIP != nil {
		result.Properties = &RedisFirewallRuleProperties_ARM{}
	}
	if rule.EndIP != nil {
		endIP := *rule.EndIP
		result.Properties.EndIP = &endIP
	}
	if rule.StartIP != nil {
		startIP := *rule.StartIP
		result.Properties.StartIP = &startIP
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Redis_FirewallRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Redis_FirewallRule_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Redis_FirewallRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Redis_FirewallRule_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Redis_FirewallRule_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "EndIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIP != nil {
			endIP := *typedInput.Properties.EndIP
			rule.EndIP = &endIP
		}
	}

	// Set property "Owner":
	rule.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property "StartIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIP != nil {
			startIP := *typedInput.Properties.StartIP
			rule.StartIP = &startIP
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Redis_FirewallRule_Spec{}

// ConvertSpecFrom populates our Redis_FirewallRule_Spec from the provided source
func (rule *Redis_FirewallRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20230401s.Redis_FirewallRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Redis_FirewallRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20230401s.Redis_FirewallRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Redis_FirewallRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Redis_FirewallRule_Spec
func (rule *Redis_FirewallRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20230401s.Redis_FirewallRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Redis_FirewallRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20230401s.Redis_FirewallRule_Spec{}
	err := rule.AssignProperties_To_Redis_FirewallRule_Spec(dst)
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

// AssignProperties_From_Redis_FirewallRule_Spec populates our Redis_FirewallRule_Spec from the provided source Redis_FirewallRule_Spec
func (rule *Redis_FirewallRule_Spec) AssignProperties_From_Redis_FirewallRule_Spec(source *v20230401s.Redis_FirewallRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// EndIP
	rule.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// StartIP
	rule.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// No error
	return nil
}

// AssignProperties_To_Redis_FirewallRule_Spec populates the provided destination Redis_FirewallRule_Spec from our Redis_FirewallRule_Spec
func (rule *Redis_FirewallRule_Spec) AssignProperties_To_Redis_FirewallRule_Spec(destination *v20230401s.Redis_FirewallRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// EndIP
	destination.EndIP = genruntime.ClonePointerToString(rule.EndIP)

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// StartIP
	destination.StartIP = genruntime.ClonePointerToString(rule.StartIP)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_Redis_FirewallRule_STATUS populates our Redis_FirewallRule_Spec from the provided source Redis_FirewallRule_STATUS
func (rule *Redis_FirewallRule_Spec) Initialize_From_Redis_FirewallRule_STATUS(source *Redis_FirewallRule_STATUS) error {

	// EndIP
	rule.EndIP = genruntime.ClonePointerToString(source.EndIP)

	// StartIP
	rule.StartIP = genruntime.ClonePointerToString(source.StartIP)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (rule *Redis_FirewallRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *Redis_FirewallRule_Spec) SetAzureName(azureName string) { rule.AzureName = azureName }

type Redis_FirewallRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Redis_FirewallRule_STATUS{}

// ConvertStatusFrom populates our Redis_FirewallRule_STATUS from the provided source
func (rule *Redis_FirewallRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20230401s.Redis_FirewallRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Redis_FirewallRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20230401s.Redis_FirewallRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Redis_FirewallRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Redis_FirewallRule_STATUS
func (rule *Redis_FirewallRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20230401s.Redis_FirewallRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Redis_FirewallRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20230401s.Redis_FirewallRule_STATUS{}
	err := rule.AssignProperties_To_Redis_FirewallRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Redis_FirewallRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Redis_FirewallRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Redis_FirewallRule_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Redis_FirewallRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Redis_FirewallRule_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Redis_FirewallRule_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "EndIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.EndIP != nil {
			endIP := *typedInput.Properties.EndIP
			rule.EndIP = &endIP
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

	// Set property "StartIP":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.StartIP != nil {
			startIP := *typedInput.Properties.StartIP
			rule.StartIP = &startIP
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

// AssignProperties_From_Redis_FirewallRule_STATUS populates our Redis_FirewallRule_STATUS from the provided source Redis_FirewallRule_STATUS
func (rule *Redis_FirewallRule_STATUS) AssignProperties_From_Redis_FirewallRule_STATUS(source *v20230401s.Redis_FirewallRule_STATUS) error {

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

// AssignProperties_To_Redis_FirewallRule_STATUS populates the provided destination Redis_FirewallRule_STATUS from our Redis_FirewallRule_STATUS
func (rule *Redis_FirewallRule_STATUS) AssignProperties_To_Redis_FirewallRule_STATUS(destination *v20230401s.Redis_FirewallRule_STATUS) error {
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

func init() {
	SchemeBuilder.Register(&RedisFirewallRule{}, &RedisFirewallRuleList{})
}
