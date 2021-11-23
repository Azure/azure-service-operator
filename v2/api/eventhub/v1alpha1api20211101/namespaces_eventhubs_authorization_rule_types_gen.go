// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
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
//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsAuthorizationRules_Spec `json:"spec,omitempty"`
	Status            AuthorizationRule_Status                   `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetConditions() conditions.Conditions {
	return namespacesEventhubsAuthorizationRule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	namespacesEventhubsAuthorizationRule.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesEventhubsAuthorizationRule{}

// ConvertFrom populates our NamespacesEventhubsAuthorizationRule from the provided hub NamespacesEventhubsAuthorizationRule
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected storage:eventhub/v1alpha1api20211101storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return namespacesEventhubsAuthorizationRule.AssignPropertiesFromNamespacesEventhubsAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected storage:eventhub/v1alpha1api20211101storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return namespacesEventhubsAuthorizationRule.AssignPropertiesToNamespacesEventhubsAuthorizationRule(destination)
}

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1alpha1api20211101-namespaceseventhubsauthorizationrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1alpha1api20211101,name=default.v1alpha1api20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &NamespacesEventhubsAuthorizationRule{}

// Default applies defaults to the NamespacesEventhubsAuthorizationRule resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) Default() {
	namespacesEventhubsAuthorizationRule.defaultImpl()
	var temp interface{} = namespacesEventhubsAuthorizationRule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) defaultAzureName() {
	if namespacesEventhubsAuthorizationRule.Spec.AzureName == "" {
		namespacesEventhubsAuthorizationRule.Spec.AzureName = namespacesEventhubsAuthorizationRule.Name
	}
}

// defaultImpl applies the code generated defaults to the NamespacesEventhubsAuthorizationRule resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) defaultImpl() {
	namespacesEventhubsAuthorizationRule.defaultAzureName()
}

var _ genruntime.KubernetesResource = &NamespacesEventhubsAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) AzureName() string {
	return namespacesEventhubsAuthorizationRule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (namespacesEventhubsAuthorizationRule NamespacesEventhubsAuthorizationRule) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceKind returns the kind of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &namespacesEventhubsAuthorizationRule.Spec
}

// GetStatus returns the status of this resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &namespacesEventhubsAuthorizationRule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationRule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(namespacesEventhubsAuthorizationRule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  namespacesEventhubsAuthorizationRule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AuthorizationRule_Status); ok {
		namespacesEventhubsAuthorizationRule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	namespacesEventhubsAuthorizationRule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1alpha1api20211101-namespaceseventhubsauthorizationrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1alpha1api20211101,name=validate.v1alpha1api20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &NamespacesEventhubsAuthorizationRule{}

// ValidateCreate validates the creation of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) ValidateCreate() error {
	validations := namespacesEventhubsAuthorizationRule.createValidations()
	var temp interface{} = namespacesEventhubsAuthorizationRule
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
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) ValidateDelete() error {
	validations := namespacesEventhubsAuthorizationRule.deleteValidations()
	var temp interface{} = namespacesEventhubsAuthorizationRule
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
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) ValidateUpdate(old runtime.Object) error {
	validations := namespacesEventhubsAuthorizationRule.updateValidations()
	var temp interface{} = namespacesEventhubsAuthorizationRule
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
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) createValidations() []func() error {
	return []func() error{namespacesEventhubsAuthorizationRule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return namespacesEventhubsAuthorizationRule.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&namespacesEventhubsAuthorizationRule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromNamespacesEventhubsAuthorizationRule populates our NamespacesEventhubsAuthorizationRule from the provided source NamespacesEventhubsAuthorizationRule
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) AssignPropertiesFromNamespacesEventhubsAuthorizationRule(source *v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	namespacesEventhubsAuthorizationRule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesEventhubsAuthorizationRules_Spec
	err := spec.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec()")
	}
	namespacesEventhubsAuthorizationRule.Spec = spec

	// Status
	var status AuthorizationRule_Status
	err = status.AssignPropertiesFromAuthorizationRuleStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesFromAuthorizationRuleStatus()")
	}
	namespacesEventhubsAuthorizationRule.Status = status

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRule populates the provided destination NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) AssignPropertiesToNamespacesEventhubsAuthorizationRule(destination *v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *namespacesEventhubsAuthorizationRule.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec
	err := namespacesEventhubsAuthorizationRule.Spec.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "populating Spec from Spec, calling AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec()")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20211101storage.AuthorizationRule_Status
	err = namespacesEventhubsAuthorizationRule.Status.AssignPropertiesToAuthorizationRuleStatus(&status)
	if err != nil {
		return errors.Wrap(err, "populating Status from Status, calling AssignPropertiesToAuthorizationRuleStatus()")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (namespacesEventhubsAuthorizationRule *NamespacesEventhubsAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: namespacesEventhubsAuthorizationRule.Spec.OriginalVersion(),
		Kind:    "NamespacesEventhubsAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules
type NamespacesEventhubsAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsAuthorizationRule `json:"items"`
}

// +kubebuilder:validation:Enum={"2021-11-01"}
type NamespacesEventhubsAuthorizationRulesSpecAPIVersion string

const NamespacesEventhubsAuthorizationRulesSpecAPIVersion20211101 = NamespacesEventhubsAuthorizationRulesSpecAPIVersion("2021-11-01")

type NamespacesEventhubsAuthorizationRules_Spec struct {
	// +kubebuilder:validation:MinLength=1
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner" kind:"NamespacesEventhub"`

	// +kubebuilder:validation:Required
	//Rights: The rights associated with the rule.
	Rights []AuthorizationRulePropertiesRights `json:"rights"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &NamespacesEventhubsAuthorizationRules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if namespacesEventhubsAuthorizationRulesSpec == nil {
		return nil, nil
	}
	var result NamespacesEventhubsAuthorizationRules_SpecARM

	// Set property ‘Location’:
	if namespacesEventhubsAuthorizationRulesSpec.Location != nil {
		location := *namespacesEventhubsAuthorizationRulesSpec.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	for _, item := range namespacesEventhubsAuthorizationRulesSpec.Rights {
		result.Properties.Rights = append(result.Properties.Rights, item)
	}

	// Set property ‘Tags’:
	if namespacesEventhubsAuthorizationRulesSpec.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range namespacesEventhubsAuthorizationRulesSpec.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &NamespacesEventhubsAuthorizationRules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(NamespacesEventhubsAuthorizationRules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected NamespacesEventhubsAuthorizationRules_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	namespacesEventhubsAuthorizationRulesSpec.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		namespacesEventhubsAuthorizationRulesSpec.Location = &location
	}

	// Set property ‘Owner’:
	namespacesEventhubsAuthorizationRulesSpec.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Rights’:
	// copying flattened property:
	for _, item := range typedInput.Properties.Rights {
		namespacesEventhubsAuthorizationRulesSpec.Rights = append(namespacesEventhubsAuthorizationRulesSpec.Rights, item)
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		namespacesEventhubsAuthorizationRulesSpec.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			namespacesEventhubsAuthorizationRulesSpec.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsAuthorizationRules_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubsAuthorizationRules_Spec from the provided source
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec)
	if ok {
		// Populate our instance from source
		return namespacesEventhubsAuthorizationRulesSpec.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = namespacesEventhubsAuthorizationRulesSpec.AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsAuthorizationRules_Spec
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec)
	if ok {
		// Populate destination from our instance
		return namespacesEventhubsAuthorizationRulesSpec.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec{}
	err := namespacesEventhubsAuthorizationRulesSpec.AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(dst)
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

// AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec populates our NamespacesEventhubsAuthorizationRules_Spec from the provided source NamespacesEventhubsAuthorizationRules_Spec
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) AssignPropertiesFromNamespacesEventhubsAuthorizationRulesSpec(source *v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec) error {

	// AzureName
	namespacesEventhubsAuthorizationRulesSpec.AzureName = source.AzureName

	// Location
	namespacesEventhubsAuthorizationRulesSpec.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	namespacesEventhubsAuthorizationRulesSpec.Owner = source.Owner.Copy()

	// Rights
	if source.Rights != nil {
		rightList := make([]AuthorizationRulePropertiesRights, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = AuthorizationRulePropertiesRights(rightItem)
		}
		namespacesEventhubsAuthorizationRulesSpec.Rights = rightList
	} else {
		namespacesEventhubsAuthorizationRulesSpec.Rights = nil
	}

	// Tags
	namespacesEventhubsAuthorizationRulesSpec.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec populates the provided destination NamespacesEventhubsAuthorizationRules_Spec from our NamespacesEventhubsAuthorizationRules_Spec
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) AssignPropertiesToNamespacesEventhubsAuthorizationRulesSpec(destination *v1alpha1api20211101storage.NamespacesEventhubsAuthorizationRules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = namespacesEventhubsAuthorizationRulesSpec.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(namespacesEventhubsAuthorizationRulesSpec.Location)

	// OriginalVersion
	destination.OriginalVersion = namespacesEventhubsAuthorizationRulesSpec.OriginalVersion()

	// Owner
	destination.Owner = namespacesEventhubsAuthorizationRulesSpec.Owner.Copy()

	// Rights
	if namespacesEventhubsAuthorizationRulesSpec.Rights != nil {
		rightList := make([]string, len(namespacesEventhubsAuthorizationRulesSpec.Rights))
		for rightIndex, rightItem := range namespacesEventhubsAuthorizationRulesSpec.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = string(rightItem)
		}
		destination.Rights = rightList
	} else {
		destination.Rights = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(namespacesEventhubsAuthorizationRulesSpec.Tags)

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
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (namespacesEventhubsAuthorizationRulesSpec *NamespacesEventhubsAuthorizationRules_Spec) SetAzureName(azureName string) {
	namespacesEventhubsAuthorizationRulesSpec.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
