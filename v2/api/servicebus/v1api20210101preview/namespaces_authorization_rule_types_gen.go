// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210101preview

import (
	"fmt"
	v20210101ps "github.com/Azure/azure-service-operator/v2/api/servicebus/v1api20210101preview/storage"
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
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/AuthorizationRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}
type NamespacesAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_AuthorizationRule_Spec   `json:"spec,omitempty"`
	Status            Namespaces_AuthorizationRule_STATUS `json:"status,omitempty"`
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

var _ conversion.Convertible = &NamespacesAuthorizationRule{}

// ConvertFrom populates our NamespacesAuthorizationRule from the provided hub NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source v20210101ps.NamespacesAuthorizationRule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = rule.AssignProperties_From_NamespacesAuthorizationRule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to rule")
	}

	return nil
}

// ConvertTo populates the provided hub NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v20210101ps.NamespacesAuthorizationRule
	err := rule.AssignProperties_To_NamespacesAuthorizationRule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from rule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-servicebus-azure-com-v1api20210101preview-namespacesauthorizationrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=servicebus.azure.com,resources=namespacesauthorizationrules,verbs=create;update,versions=v1api20210101preview,name=default.v1api20210101preview.namespacesauthorizationrules.servicebus.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &NamespacesAuthorizationRule{}

// Default applies defaults to the NamespacesAuthorizationRule resource
func (rule *NamespacesAuthorizationRule) Default() {
	rule.defaultImpl()
	var temp any = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *NamespacesAuthorizationRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the NamespacesAuthorizationRule resource
func (rule *NamespacesAuthorizationRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &NamespacesAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-01-01-preview"
func (rule NamespacesAuthorizationRule) GetAPIVersion() string {
	return string(APIVersion_Value)
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
	return &Namespaces_AuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *NamespacesAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *NamespacesAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Namespaces_AuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Namespaces_AuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-servicebus-azure-com-v1api20210101preview-namespacesauthorizationrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=servicebus.azure.com,resources=namespacesauthorizationrules,verbs=create;update,versions=v1api20210101preview,name=validate.v1api20210101preview.namespacesauthorizationrules.servicebus.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &NamespacesAuthorizationRule{}

// ValidateCreate validates the creation of the resource
func (rule *NamespacesAuthorizationRule) ValidateCreate() (admission.Warnings, error) {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *NamespacesAuthorizationRule) ValidateDelete() (admission.Warnings, error) {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (rule *NamespacesAuthorizationRule) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (rule *NamespacesAuthorizationRule) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){rule.validateResourceReferences, rule.validateOwnerReference, rule.validateSecretDestinations}
}

// deleteValidations validates the deletion of the resource
func (rule *NamespacesAuthorizationRule) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *NamespacesAuthorizationRule) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return rule.validateSecretDestinations()
		},
	}
}

// validateOwnerReference validates the owner field
func (rule *NamespacesAuthorizationRule) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(rule)
}

// validateResourceReferences validates all resource references
func (rule *NamespacesAuthorizationRule) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (rule *NamespacesAuthorizationRule) validateSecretDestinations() (admission.Warnings, error) {
	if rule.Spec.OperatorSpec == nil {
		return nil, nil
	}
	if rule.Spec.OperatorSpec.Secrets == nil {
		return nil, nil
	}
	toValidate := []*genruntime.SecretDestination{
		rule.Spec.OperatorSpec.Secrets.PrimaryConnectionString,
		rule.Spec.OperatorSpec.Secrets.PrimaryKey,
		rule.Spec.OperatorSpec.Secrets.SecondaryConnectionString,
		rule.Spec.OperatorSpec.Secrets.SecondaryKey,
	}
	return genruntime.ValidateSecretDestinations(toValidate)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *NamespacesAuthorizationRule) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*NamespacesAuthorizationRule)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_NamespacesAuthorizationRule populates our NamespacesAuthorizationRule from the provided source NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignProperties_From_NamespacesAuthorizationRule(source *v20210101ps.NamespacesAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Namespaces_AuthorizationRule_Spec
	err := spec.AssignProperties_From_Namespaces_AuthorizationRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_AuthorizationRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status Namespaces_AuthorizationRule_STATUS
	err = status.AssignProperties_From_Namespaces_AuthorizationRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_AuthorizationRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRule populates the provided destination NamespacesAuthorizationRule from our NamespacesAuthorizationRule
func (rule *NamespacesAuthorizationRule) AssignProperties_To_NamespacesAuthorizationRule(destination *v20210101ps.NamespacesAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20210101ps.Namespaces_AuthorizationRule_Spec
	err := rule.Spec.AssignProperties_To_Namespaces_AuthorizationRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_AuthorizationRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20210101ps.Namespaces_AuthorizationRule_STATUS
	err = rule.Status.AssignProperties_To_Namespaces_AuthorizationRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_AuthorizationRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "NamespacesAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /servicebus/resource-manager/Microsoft.ServiceBus/preview/2021-01-01-preview/AuthorizationRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ServiceBus/namespaces/{namespaceName}/AuthorizationRules/{authorizationRuleName}
type NamespacesAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesAuthorizationRule `json:"items"`
}

type Namespaces_AuthorizationRule_Spec struct {
	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *NamespacesAuthorizationRuleOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a servicebus.azure.com/Namespace resource
	Owner *genruntime.KnownResourceReference `group:"servicebus.azure.com" json:"owner,omitempty" kind:"Namespace"`

	// +kubebuilder:validation:Required
	// Rights: The rights associated with the rule.
	Rights []Namespaces_AuthorizationRule_Properties_Rights_Spec `json:"rights,omitempty"`
}

var _ genruntime.ARMTransformer = &Namespaces_AuthorizationRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *Namespaces_AuthorizationRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &Namespaces_AuthorizationRule_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if rule.Rights != nil {
		result.Properties = &Namespaces_AuthorizationRule_Properties_Spec_ARM{}
	}
	for _, item := range rule.Rights {
		result.Properties.Rights = append(result.Properties.Rights, item)
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Namespaces_AuthorizationRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Namespaces_AuthorizationRule_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Namespaces_AuthorizationRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Namespaces_AuthorizationRule_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Namespaces_AuthorizationRule_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	rule.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "Rights":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Rights {
			rule.Rights = append(rule.Rights, item)
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Namespaces_AuthorizationRule_Spec{}

// ConvertSpecFrom populates our Namespaces_AuthorizationRule_Spec from the provided source
func (rule *Namespaces_AuthorizationRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20210101ps.Namespaces_AuthorizationRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Namespaces_AuthorizationRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.Namespaces_AuthorizationRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Namespaces_AuthorizationRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Namespaces_AuthorizationRule_Spec
func (rule *Namespaces_AuthorizationRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20210101ps.Namespaces_AuthorizationRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Namespaces_AuthorizationRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.Namespaces_AuthorizationRule_Spec{}
	err := rule.AssignProperties_To_Namespaces_AuthorizationRule_Spec(dst)
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

// AssignProperties_From_Namespaces_AuthorizationRule_Spec populates our Namespaces_AuthorizationRule_Spec from the provided source Namespaces_AuthorizationRule_Spec
func (rule *Namespaces_AuthorizationRule_Spec) AssignProperties_From_Namespaces_AuthorizationRule_Spec(source *v20210101ps.Namespaces_AuthorizationRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec NamespacesAuthorizationRuleOperatorSpec
		err := operatorSpec.AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec(source.OperatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec() to populate field OperatorSpec")
		}
		rule.OperatorSpec = &operatorSpec
	} else {
		rule.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// Rights
	if source.Rights != nil {
		rightList := make([]Namespaces_AuthorizationRule_Properties_Rights_Spec, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = Namespaces_AuthorizationRule_Properties_Rights_Spec(rightItem)
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_AuthorizationRule_Spec populates the provided destination Namespaces_AuthorizationRule_Spec from our Namespaces_AuthorizationRule_Spec
func (rule *Namespaces_AuthorizationRule_Spec) AssignProperties_To_Namespaces_AuthorizationRule_Spec(destination *v20210101ps.Namespaces_AuthorizationRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// OperatorSpec
	if rule.OperatorSpec != nil {
		var operatorSpec v20210101ps.NamespacesAuthorizationRuleOperatorSpec
		err := rule.OperatorSpec.AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec(&operatorSpec)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = rule.OriginalVersion()

	// Owner
	if rule.Owner != nil {
		owner := rule.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Rights
	if rule.Rights != nil {
		rightList := make([]string, len(rule.Rights))
		for rightIndex, rightItem := range rule.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = string(rightItem)
		}
		destination.Rights = rightList
	} else {
		destination.Rights = nil
	}

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
func (rule *Namespaces_AuthorizationRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *Namespaces_AuthorizationRule_Spec) SetAzureName(azureName string) {
	rule.AzureName = azureName
}

type Namespaces_AuthorizationRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Resource Id
	Id *string `json:"id,omitempty"`

	// Name: Resource name
	Name *string `json:"name,omitempty"`

	// Rights: The rights associated with the rule.
	Rights []Namespaces_AuthorizationRule_Properties_Rights_STATUS `json:"rights,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: Resource type
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Namespaces_AuthorizationRule_STATUS{}

// ConvertStatusFrom populates our Namespaces_AuthorizationRule_STATUS from the provided source
func (rule *Namespaces_AuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20210101ps.Namespaces_AuthorizationRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Namespaces_AuthorizationRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20210101ps.Namespaces_AuthorizationRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Namespaces_AuthorizationRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Namespaces_AuthorizationRule_STATUS
func (rule *Namespaces_AuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20210101ps.Namespaces_AuthorizationRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Namespaces_AuthorizationRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20210101ps.Namespaces_AuthorizationRule_STATUS{}
	err := rule.AssignProperties_To_Namespaces_AuthorizationRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Namespaces_AuthorizationRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Namespaces_AuthorizationRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Namespaces_AuthorizationRule_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Namespaces_AuthorizationRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Namespaces_AuthorizationRule_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Namespaces_AuthorizationRule_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

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

	// Set property "Rights":
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Rights {
			rule.Rights = append(rule.Rights, item)
		}
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		rule.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		rule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignProperties_From_Namespaces_AuthorizationRule_STATUS populates our Namespaces_AuthorizationRule_STATUS from the provided source Namespaces_AuthorizationRule_STATUS
func (rule *Namespaces_AuthorizationRule_STATUS) AssignProperties_From_Namespaces_AuthorizationRule_STATUS(source *v20210101ps.Namespaces_AuthorizationRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// Rights
	if source.Rights != nil {
		rightList := make([]Namespaces_AuthorizationRule_Properties_Rights_STATUS, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = Namespaces_AuthorizationRule_Properties_Rights_STATUS(rightItem)
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
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

// AssignProperties_To_Namespaces_AuthorizationRule_STATUS populates the provided destination Namespaces_AuthorizationRule_STATUS from our Namespaces_AuthorizationRule_STATUS
func (rule *Namespaces_AuthorizationRule_STATUS) AssignProperties_To_Namespaces_AuthorizationRule_STATUS(destination *v20210101ps.Namespaces_AuthorizationRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(rule.Name)

	// Rights
	if rule.Rights != nil {
		rightList := make([]string, len(rule.Rights))
		for rightIndex, rightItem := range rule.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = string(rightItem)
		}
		destination.Rights = rightList
	} else {
		destination.Rights = nil
	}

	// SystemData
	if rule.SystemData != nil {
		var systemDatum v20210101ps.SystemData_STATUS
		err := rule.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
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

// +kubebuilder:validation:Enum={"Listen","Manage","Send"}
type Namespaces_AuthorizationRule_Properties_Rights_Spec string

const (
	Namespaces_AuthorizationRule_Properties_Rights_Spec_Listen = Namespaces_AuthorizationRule_Properties_Rights_Spec("Listen")
	Namespaces_AuthorizationRule_Properties_Rights_Spec_Manage = Namespaces_AuthorizationRule_Properties_Rights_Spec("Manage")
	Namespaces_AuthorizationRule_Properties_Rights_Spec_Send   = Namespaces_AuthorizationRule_Properties_Rights_Spec("Send")
)

type Namespaces_AuthorizationRule_Properties_Rights_STATUS string

const (
	Namespaces_AuthorizationRule_Properties_Rights_STATUS_Listen = Namespaces_AuthorizationRule_Properties_Rights_STATUS("Listen")
	Namespaces_AuthorizationRule_Properties_Rights_STATUS_Manage = Namespaces_AuthorizationRule_Properties_Rights_STATUS("Manage")
	Namespaces_AuthorizationRule_Properties_Rights_STATUS_Send   = Namespaces_AuthorizationRule_Properties_Rights_STATUS("Send")
)

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesAuthorizationRuleOperatorSpec struct {
	// Secrets: configures where to place Azure generated secrets.
	Secrets *NamespacesAuthorizationRuleOperatorSecrets `json:"secrets,omitempty"`
}

// AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec populates our NamespacesAuthorizationRuleOperatorSpec from the provided source NamespacesAuthorizationRuleOperatorSpec
func (operator *NamespacesAuthorizationRuleOperatorSpec) AssignProperties_From_NamespacesAuthorizationRuleOperatorSpec(source *v20210101ps.NamespacesAuthorizationRuleOperatorSpec) error {

	// Secrets
	if source.Secrets != nil {
		var secret NamespacesAuthorizationRuleOperatorSecrets
		err := secret.AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets(source.Secrets)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets() to populate field Secrets")
		}
		operator.Secrets = &secret
	} else {
		operator.Secrets = nil
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec populates the provided destination NamespacesAuthorizationRuleOperatorSpec from our NamespacesAuthorizationRuleOperatorSpec
func (operator *NamespacesAuthorizationRuleOperatorSpec) AssignProperties_To_NamespacesAuthorizationRuleOperatorSpec(destination *v20210101ps.NamespacesAuthorizationRuleOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Secrets
	if operator.Secrets != nil {
		var secret v20210101ps.NamespacesAuthorizationRuleOperatorSecrets
		err := operator.Secrets.AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets(&secret)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets() to populate field Secrets")
		}
		destination.Secrets = &secret
	} else {
		destination.Secrets = nil
	}

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type NamespacesAuthorizationRuleOperatorSecrets struct {
	// PrimaryConnectionString: indicates where the PrimaryConnectionString secret should be placed. If omitted, the secret
	// will not be retrieved from Azure.
	PrimaryConnectionString *genruntime.SecretDestination `json:"primaryConnectionString,omitempty"`

	// PrimaryKey: indicates where the PrimaryKey secret should be placed. If omitted, the secret will not be retrieved from
	// Azure.
	PrimaryKey *genruntime.SecretDestination `json:"primaryKey,omitempty"`

	// SecondaryConnectionString: indicates where the SecondaryConnectionString secret should be placed. If omitted, the secret
	// will not be retrieved from Azure.
	SecondaryConnectionString *genruntime.SecretDestination `json:"secondaryConnectionString,omitempty"`

	// SecondaryKey: indicates where the SecondaryKey secret should be placed. If omitted, the secret will not be retrieved
	// from Azure.
	SecondaryKey *genruntime.SecretDestination `json:"secondaryKey,omitempty"`
}

// AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets populates our NamespacesAuthorizationRuleOperatorSecrets from the provided source NamespacesAuthorizationRuleOperatorSecrets
func (secrets *NamespacesAuthorizationRuleOperatorSecrets) AssignProperties_From_NamespacesAuthorizationRuleOperatorSecrets(source *v20210101ps.NamespacesAuthorizationRuleOperatorSecrets) error {

	// PrimaryConnectionString
	if source.PrimaryConnectionString != nil {
		primaryConnectionString := source.PrimaryConnectionString.Copy()
		secrets.PrimaryConnectionString = &primaryConnectionString
	} else {
		secrets.PrimaryConnectionString = nil
	}

	// PrimaryKey
	if source.PrimaryKey != nil {
		primaryKey := source.PrimaryKey.Copy()
		secrets.PrimaryKey = &primaryKey
	} else {
		secrets.PrimaryKey = nil
	}

	// SecondaryConnectionString
	if source.SecondaryConnectionString != nil {
		secondaryConnectionString := source.SecondaryConnectionString.Copy()
		secrets.SecondaryConnectionString = &secondaryConnectionString
	} else {
		secrets.SecondaryConnectionString = nil
	}

	// SecondaryKey
	if source.SecondaryKey != nil {
		secondaryKey := source.SecondaryKey.Copy()
		secrets.SecondaryKey = &secondaryKey
	} else {
		secrets.SecondaryKey = nil
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets populates the provided destination NamespacesAuthorizationRuleOperatorSecrets from our NamespacesAuthorizationRuleOperatorSecrets
func (secrets *NamespacesAuthorizationRuleOperatorSecrets) AssignProperties_To_NamespacesAuthorizationRuleOperatorSecrets(destination *v20210101ps.NamespacesAuthorizationRuleOperatorSecrets) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// PrimaryConnectionString
	if secrets.PrimaryConnectionString != nil {
		primaryConnectionString := secrets.PrimaryConnectionString.Copy()
		destination.PrimaryConnectionString = &primaryConnectionString
	} else {
		destination.PrimaryConnectionString = nil
	}

	// PrimaryKey
	if secrets.PrimaryKey != nil {
		primaryKey := secrets.PrimaryKey.Copy()
		destination.PrimaryKey = &primaryKey
	} else {
		destination.PrimaryKey = nil
	}

	// SecondaryConnectionString
	if secrets.SecondaryConnectionString != nil {
		secondaryConnectionString := secrets.SecondaryConnectionString.Copy()
		destination.SecondaryConnectionString = &secondaryConnectionString
	} else {
		destination.SecondaryConnectionString = nil
	}

	// SecondaryKey
	if secrets.SecondaryKey != nil {
		secondaryKey := secrets.SecondaryKey.Copy()
		destination.SecondaryKey = &secondaryKey
	} else {
		destination.SecondaryKey = nil
	}

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
	SchemeBuilder.Register(&NamespacesAuthorizationRule{}, &NamespacesAuthorizationRuleList{})
}
