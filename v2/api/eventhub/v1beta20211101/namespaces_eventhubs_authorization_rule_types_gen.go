// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Eventhubs_AuthorizationRule_Spec `json:"spec,omitempty"`
	Status            AuthorizationRule_STATUS                    `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsAuthorizationRule{}

// GetConditions returns the conditions of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetConditions() conditions.Conditions {
	return rule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (rule *NamespacesEventhubsAuthorizationRule) SetConditions(conditions conditions.Conditions) {
	rule.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesEventhubsAuthorizationRule{}

// ConvertFrom populates our NamespacesEventhubsAuthorizationRule from the provided hub NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v20211101s.NamespacesEventhubsAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_NamespacesEventhubsAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.NamespacesEventhubsAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1beta20211101storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_NamespacesEventhubsAuthorizationRule(destination)
}

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1beta20211101-namespaceseventhubsauthorizationrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1beta20211101,name=default.v1beta20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &NamespacesEventhubsAuthorizationRule{}

// Default applies defaults to the NamespacesEventhubsAuthorizationRule resource
func (rule *NamespacesEventhubsAuthorizationRule) Default() {
	rule.defaultImpl()
	var temp interface{} = rule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (rule *NamespacesEventhubsAuthorizationRule) defaultAzureName() {
	if rule.Spec.AzureName == "" {
		rule.Spec.AzureName = rule.Name
	}
}

// defaultImpl applies the code generated defaults to the NamespacesEventhubsAuthorizationRule resource
func (rule *NamespacesEventhubsAuthorizationRule) defaultImpl() { rule.defaultAzureName() }

var _ genruntime.KubernetesResource = &NamespacesEventhubsAuthorizationRule{}

// AzureName returns the Azure name of the resource
func (rule *NamespacesEventhubsAuthorizationRule) AzureName() string {
	return rule.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (rule NamespacesEventhubsAuthorizationRule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetSpec() genruntime.ConvertibleSpec {
	return &rule.Spec
}

// GetStatus returns the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) GetStatus() genruntime.ConvertibleStatus {
	return &rule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (rule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &AuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (rule *NamespacesEventhubsAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  rule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*AuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st AuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1beta20211101-namespaceseventhubsauthorizationrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1beta20211101,name=validate.v1beta20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &NamespacesEventhubsAuthorizationRule{}

// ValidateCreate validates the creation of the resource
func (rule *NamespacesEventhubsAuthorizationRule) ValidateCreate() error {
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
func (rule *NamespacesEventhubsAuthorizationRule) ValidateDelete() error {
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
func (rule *NamespacesEventhubsAuthorizationRule) ValidateUpdate(old runtime.Object) error {
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
func (rule *NamespacesEventhubsAuthorizationRule) createValidations() []func() error {
	return []func() error{rule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (rule *NamespacesEventhubsAuthorizationRule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (rule *NamespacesEventhubsAuthorizationRule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return rule.validateResourceReferences()
		},
		rule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (rule *NamespacesEventhubsAuthorizationRule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *NamespacesEventhubsAuthorizationRule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*NamespacesEventhubsAuthorizationRule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, rule)
}

// AssignProperties_From_NamespacesEventhubsAuthorizationRule populates our NamespacesEventhubsAuthorizationRule from the provided source NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) AssignProperties_From_NamespacesEventhubsAuthorizationRule(source *v20211101s.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	rule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Namespaces_Eventhubs_AuthorizationRule_Spec
	err := spec.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec() to populate field Spec")
	}
	rule.Spec = spec

	// Status
	var status AuthorizationRule_STATUS
	err = status.AssignProperties_From_AuthorizationRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_AuthorizationRule_STATUS() to populate field Status")
	}
	rule.Status = status

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsAuthorizationRule populates the provided destination NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) AssignProperties_To_NamespacesEventhubsAuthorizationRule(destination *v20211101s.NamespacesEventhubsAuthorizationRule) error {

	// ObjectMeta
	destination.ObjectMeta = *rule.ObjectMeta.DeepCopy()

	// Spec
	var spec v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec
	err := rule.Spec.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v20211101s.AuthorizationRule_STATUS
	err = rule.Status.AssignProperties_To_AuthorizationRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_AuthorizationRule_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (rule *NamespacesEventhubsAuthorizationRule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: rule.Spec.OriginalVersion(),
		Kind:    "NamespacesEventhubsAuthorizationRule",
	}
}

// +kubebuilder:object:root=true
// Generated from: https://schema.management.azure.com/schemas/2021-11-01/Microsoft.EventHub.json#/resourceDefinitions/namespaces_eventhubs_authorizationRules
type NamespacesEventhubsAuthorizationRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsAuthorizationRule `json:"items"`
}

type Namespaces_Eventhubs_AuthorizationRule_Spec struct {
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`

	// +kubebuilder:validation:Required
	// Rights: The rights associated with the rule.
	Rights []AuthorizationRuleProperties_Rights `json:"rights,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &Namespaces_Eventhubs_AuthorizationRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &Namespaces_Eventhubs_AuthorizationRule_Spec_ARM{}

	// Set property ‘Location’:
	if rule.Location != nil {
		location := *rule.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if rule.Rights != nil {
		result.Properties = &AuthorizationRuleProperties_ARM{}
	}
	for _, item := range rule.Rights {
		result.Properties.Rights = append(result.Properties.Rights, item)
	}

	// Set property ‘Tags’:
	if rule.Tags != nil {
		result.Tags = make(map[string]string, len(rule.Tags))
		for key, value := range rule.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Namespaces_Eventhubs_AuthorizationRule_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Namespaces_Eventhubs_AuthorizationRule_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Namespaces_Eventhubs_AuthorizationRule_Spec_ARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		rule.Location = &location
	}

	// Set property ‘Owner’:
	rule.Owner = &genruntime.KnownResourceReference{Name: owner.Name}

	// Set property ‘Rights’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.Rights {
			rule.Rights = append(rule.Rights, item)
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		rule.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			rule.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Namespaces_Eventhubs_AuthorizationRule_Spec{}

// ConvertSpecFrom populates our Namespaces_Eventhubs_AuthorizationRule_Spec from the provided source
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Namespaces_Eventhubs_AuthorizationRule_Spec
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec{}
	err := rule.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec(dst)
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

// AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec populates our Namespaces_Eventhubs_AuthorizationRule_Spec from the provided source Namespaces_Eventhubs_AuthorizationRule_Spec
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_Spec(source *v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec) error {

	// AzureName
	rule.AzureName = source.AzureName

	// Location
	rule.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// Rights
	if source.Rights != nil {
		rightList := make([]AuthorizationRuleProperties_Rights, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = AuthorizationRuleProperties_Rights(rightItem)
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
	}

	// Tags
	rule.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec populates the provided destination Namespaces_Eventhubs_AuthorizationRule_Spec from our Namespaces_Eventhubs_AuthorizationRule_Spec
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec(destination *v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(rule.Location)

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

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(rule.Tags)

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
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) SetAzureName(azureName string) {
	rule.AzureName = azureName
}

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
