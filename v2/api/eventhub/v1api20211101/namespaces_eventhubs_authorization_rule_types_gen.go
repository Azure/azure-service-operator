// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	v20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101/storage"
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
// - Generated from: /eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/AuthorizationRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}
type NamespacesEventhubsAuthorizationRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Eventhubs_AuthorizationRule_Spec   `json:"spec,omitempty"`
	Status            Namespaces_Eventhubs_AuthorizationRule_STATUS `json:"status,omitempty"`
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
		return fmt.Errorf("expected eventhub/v1api20211101/storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_From_NamespacesEventhubsAuthorizationRule(source)
}

// ConvertTo populates the provided hub NamespacesEventhubsAuthorizationRule from our NamespacesEventhubsAuthorizationRule
func (rule *NamespacesEventhubsAuthorizationRule) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v20211101s.NamespacesEventhubsAuthorizationRule)
	if !ok {
		return fmt.Errorf("expected eventhub/v1api20211101/storage/NamespacesEventhubsAuthorizationRule but received %T instead", hub)
	}

	return rule.AssignProperties_To_NamespacesEventhubsAuthorizationRule(destination)
}

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1api20211101-namespaceseventhubsauthorizationrule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &NamespacesEventhubsAuthorizationRule{}

// Default applies defaults to the NamespacesEventhubsAuthorizationRule resource
func (rule *NamespacesEventhubsAuthorizationRule) Default() {
	rule.defaultImpl()
	var temp any = rule
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

var _ genruntime.ImportableResource = &NamespacesEventhubsAuthorizationRule{}

// InitializeSpec initializes the spec for this resource from the given status
func (rule *NamespacesEventhubsAuthorizationRule) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*Namespaces_Eventhubs_AuthorizationRule_STATUS); ok {
		return rule.Spec.Initialize_From_Namespaces_Eventhubs_AuthorizationRule_STATUS(s)
	}

	return fmt.Errorf("expected Status of type Namespaces_Eventhubs_AuthorizationRule_STATUS but received %T instead", status)
}

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

// GetSupportedOperations returns the operations supported by the resource
func (rule *NamespacesEventhubsAuthorizationRule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
func (rule *NamespacesEventhubsAuthorizationRule) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/authorizationRules"
}

// NewEmptyStatus returns a new empty (blank) status
func (rule *NamespacesEventhubsAuthorizationRule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Namespaces_Eventhubs_AuthorizationRule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (rule *NamespacesEventhubsAuthorizationRule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(rule.Spec)
	return rule.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (rule *NamespacesEventhubsAuthorizationRule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Namespaces_Eventhubs_AuthorizationRule_STATUS); ok {
		rule.Status = *st
		return nil
	}

	// Convert status to required version
	var st Namespaces_Eventhubs_AuthorizationRule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	rule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1api20211101-namespaceseventhubsauthorizationrule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsauthorizationrules,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.namespaceseventhubsauthorizationrules.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &NamespacesEventhubsAuthorizationRule{}

// ValidateCreate validates the creation of the resource
func (rule *NamespacesEventhubsAuthorizationRule) ValidateCreate() (admission.Warnings, error) {
	validations := rule.createValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (rule *NamespacesEventhubsAuthorizationRule) ValidateDelete() (admission.Warnings, error) {
	validations := rule.deleteValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (rule *NamespacesEventhubsAuthorizationRule) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := rule.updateValidations()
	var temp any = rule
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (rule *NamespacesEventhubsAuthorizationRule) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){rule.validateResourceReferences, rule.validateOwnerReference}
}

// deleteValidations validates the deletion of the resource
func (rule *NamespacesEventhubsAuthorizationRule) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (rule *NamespacesEventhubsAuthorizationRule) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
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
func (rule *NamespacesEventhubsAuthorizationRule) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(rule)
}

// validateResourceReferences validates all resource references
func (rule *NamespacesEventhubsAuthorizationRule) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&rule.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (rule *NamespacesEventhubsAuthorizationRule) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*NamespacesEventhubsAuthorizationRule)
	if !ok {
		return nil, nil
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
	var status Namespaces_Eventhubs_AuthorizationRule_STATUS
	err = status.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS() to populate field Status")
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
	var status v20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS
	err = rule.Status.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS() to populate field Status")
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
// Generator information:
// - Generated from: /eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/AuthorizationRules.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/authorizationRules/{authorizationRuleName}
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

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`

	// +kubebuilder:validation:Required
	// Rights: The rights associated with the rule.
	Rights []Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec `json:"rights,omitempty"`
}

var _ genruntime.ARMTransformer = &Namespaces_Eventhubs_AuthorizationRule_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if rule == nil {
		return nil, nil
	}
	result := &Namespaces_Eventhubs_AuthorizationRule_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if rule.Rights != nil {
		result.Properties = &Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM{}
	}
	for _, item := range rule.Rights {
		result.Properties.Rights = append(result.Properties.Rights, item)
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

	// Set property "AzureName":
	rule.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

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

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		rule.Owner = &owner
	} else {
		rule.Owner = nil
	}

	// Rights
	if source.Rights != nil {
		rightList := make([]Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec(rightItem)
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
	}

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec populates the provided destination Namespaces_Eventhubs_AuthorizationRule_Spec from our Namespaces_Eventhubs_AuthorizationRule_Spec
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_Spec(destination *v20211101s.Namespaces_Eventhubs_AuthorizationRule_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = rule.AzureName

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

// Initialize_From_Namespaces_Eventhubs_AuthorizationRule_STATUS populates our Namespaces_Eventhubs_AuthorizationRule_Spec from the provided source Namespaces_Eventhubs_AuthorizationRule_STATUS
func (rule *Namespaces_Eventhubs_AuthorizationRule_Spec) Initialize_From_Namespaces_Eventhubs_AuthorizationRule_STATUS(source *Namespaces_Eventhubs_AuthorizationRule_STATUS) error {

	// Rights
	if source.Rights != nil {
		rightList := make([]Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			right := Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec(rightItem)
			rightList[rightIndex] = right
		}
		rule.Rights = rightList
	} else {
		rule.Rights = nil
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

type Namespaces_Eventhubs_AuthorizationRule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Rights: The rights associated with the rule.
	Rights []Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS `json:"rights,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Namespaces_Eventhubs_AuthorizationRule_STATUS{}

// ConvertStatusFrom populates our Namespaces_Eventhubs_AuthorizationRule_STATUS from the provided source
func (rule *Namespaces_Eventhubs_AuthorizationRule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS)
	if ok {
		// Populate our instance from source
		return rule.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = rule.AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Namespaces_Eventhubs_AuthorizationRule_STATUS
func (rule *Namespaces_Eventhubs_AuthorizationRule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS)
	if ok {
		// Populate destination from our instance
		return rule.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS{}
	err := rule.AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Namespaces_Eventhubs_AuthorizationRule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (rule *Namespaces_Eventhubs_AuthorizationRule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Namespaces_Eventhubs_AuthorizationRule_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (rule *Namespaces_Eventhubs_AuthorizationRule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Namespaces_Eventhubs_AuthorizationRule_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Namespaces_Eventhubs_AuthorizationRule_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		rule.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		rule.Location = &location
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

// AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS populates our Namespaces_Eventhubs_AuthorizationRule_STATUS from the provided source Namespaces_Eventhubs_AuthorizationRule_STATUS
func (rule *Namespaces_Eventhubs_AuthorizationRule_STATUS) AssignProperties_From_Namespaces_Eventhubs_AuthorizationRule_STATUS(source *v20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS) error {

	// Conditions
	rule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	rule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	rule.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	rule.Name = genruntime.ClonePointerToString(source.Name)

	// Rights
	if source.Rights != nil {
		rightList := make([]Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS, len(source.Rights))
		for rightIndex, rightItem := range source.Rights {
			// Shadow the loop variable to avoid aliasing
			rightItem := rightItem
			rightList[rightIndex] = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS(rightItem)
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

// AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS populates the provided destination Namespaces_Eventhubs_AuthorizationRule_STATUS from our Namespaces_Eventhubs_AuthorizationRule_STATUS
func (rule *Namespaces_Eventhubs_AuthorizationRule_STATUS) AssignProperties_To_Namespaces_Eventhubs_AuthorizationRule_STATUS(destination *v20211101s.Namespaces_Eventhubs_AuthorizationRule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(rule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(rule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(rule.Location)

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
		var systemDatum v20211101s.SystemData_STATUS
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
type Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec string

const (
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_Listen = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec("Listen")
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_Manage = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec("Manage")
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_Send   = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec("Send")
)

type Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS string

const (
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Listen = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS("Listen")
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Manage = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS("Manage")
	Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Send   = Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS("Send")
)

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsAuthorizationRule{}, &NamespacesEventhubsAuthorizationRuleList{})
}
