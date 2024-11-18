// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import (
	"fmt"
	arm "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101/arm"
	storage "github.com/Azure/azure-service-operator/v2/api/eventhub/v1api20211101/storage"
	"github.com/Azure/azure-service-operator/v2/internal/reflecthelpers"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/configmaps"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/secrets"
	"github.com/rotisserie/eris"
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
// - Generated from: /eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/consumergroups.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}
type NamespacesEventhubsConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NamespacesEventhubsConsumerGroup_Spec   `json:"spec,omitempty"`
	Status            NamespacesEventhubsConsumerGroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NamespacesEventhubsConsumerGroup{}

// GetConditions returns the conditions of the resource
func (group *NamespacesEventhubsConsumerGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *NamespacesEventhubsConsumerGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ conversion.Convertible = &NamespacesEventhubsConsumerGroup{}

// ConvertFrom populates our NamespacesEventhubsConsumerGroup from the provided hub NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*storage.NamespacesEventhubsConsumerGroup)
	if !ok {
		return fmt.Errorf("expected eventhub/v1api20211101/storage/NamespacesEventhubsConsumerGroup but received %T instead", hub)
	}

	return group.AssignProperties_From_NamespacesEventhubsConsumerGroup(source)
}

// ConvertTo populates the provided hub NamespacesEventhubsConsumerGroup from our NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*storage.NamespacesEventhubsConsumerGroup)
	if !ok {
		return fmt.Errorf("expected eventhub/v1api20211101/storage/NamespacesEventhubsConsumerGroup but received %T instead", hub)
	}

	return group.AssignProperties_To_NamespacesEventhubsConsumerGroup(destination)
}

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1api20211101-namespaceseventhubsconsumergroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsconsumergroups,verbs=create;update,versions=v1api20211101,name=default.v1api20211101.namespaceseventhubsconsumergroups.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &NamespacesEventhubsConsumerGroup{}

// Default applies defaults to the NamespacesEventhubsConsumerGroup resource
func (group *NamespacesEventhubsConsumerGroup) Default() {
	group.defaultImpl()
	var temp any = group
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (group *NamespacesEventhubsConsumerGroup) defaultAzureName() {
	if group.Spec.AzureName == "" {
		group.Spec.AzureName = group.Name
	}
}

// defaultImpl applies the code generated defaults to the NamespacesEventhubsConsumerGroup resource
func (group *NamespacesEventhubsConsumerGroup) defaultImpl() { group.defaultAzureName() }

var _ configmaps.Exporter = &NamespacesEventhubsConsumerGroup{}

// ConfigMapDestinationExpressions returns the Spec.OperatorSpec.ConfigMapExpressions property
func (group *NamespacesEventhubsConsumerGroup) ConfigMapDestinationExpressions() []*core.DestinationExpression {
	if group.Spec.OperatorSpec == nil {
		return nil
	}
	return group.Spec.OperatorSpec.ConfigMapExpressions
}

var _ secrets.Exporter = &NamespacesEventhubsConsumerGroup{}

// SecretDestinationExpressions returns the Spec.OperatorSpec.SecretExpressions property
func (group *NamespacesEventhubsConsumerGroup) SecretDestinationExpressions() []*core.DestinationExpression {
	if group.Spec.OperatorSpec == nil {
		return nil
	}
	return group.Spec.OperatorSpec.SecretExpressions
}

var _ genruntime.ImportableResource = &NamespacesEventhubsConsumerGroup{}

// InitializeSpec initializes the spec for this resource from the given status
func (group *NamespacesEventhubsConsumerGroup) InitializeSpec(status genruntime.ConvertibleStatus) error {
	if s, ok := status.(*NamespacesEventhubsConsumerGroup_STATUS); ok {
		return group.Spec.Initialize_From_NamespacesEventhubsConsumerGroup_STATUS(s)
	}

	return fmt.Errorf("expected Status of type NamespacesEventhubsConsumerGroup_STATUS but received %T instead", status)
}

var _ genruntime.KubernetesResource = &NamespacesEventhubsConsumerGroup{}

// AzureName returns the Azure name of the resource
func (group *NamespacesEventhubsConsumerGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (group NamespacesEventhubsConsumerGroup) GetAPIVersion() string {
	return "2021-11-01"
}

// GetResourceScope returns the scope of the resource
func (group *NamespacesEventhubsConsumerGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (group *NamespacesEventhubsConsumerGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *NamespacesEventhubsConsumerGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (group *NamespacesEventhubsConsumerGroup) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
func (group *NamespacesEventhubsConsumerGroup) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *NamespacesEventhubsConsumerGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NamespacesEventhubsConsumerGroup_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (group *NamespacesEventhubsConsumerGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return group.Spec.Owner.AsResourceReference(ownerGroup, ownerKind)
}

// SetStatus sets the status of this resource
func (group *NamespacesEventhubsConsumerGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NamespacesEventhubsConsumerGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st NamespacesEventhubsConsumerGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return eris.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1api20211101-namespaceseventhubsconsumergroup,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsconsumergroups,verbs=create;update,versions=v1api20211101,name=validate.v1api20211101.namespaceseventhubsconsumergroups.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &NamespacesEventhubsConsumerGroup{}

// ValidateCreate validates the creation of the resource
func (group *NamespacesEventhubsConsumerGroup) ValidateCreate() (admission.Warnings, error) {
	validations := group.createValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.CreateValidations()...)
	}
	return genruntime.ValidateCreate(validations)
}

// ValidateDelete validates the deletion of the resource
func (group *NamespacesEventhubsConsumerGroup) ValidateDelete() (admission.Warnings, error) {
	validations := group.deleteValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.DeleteValidations()...)
	}
	return genruntime.ValidateDelete(validations)
}

// ValidateUpdate validates an update of the resource
func (group *NamespacesEventhubsConsumerGroup) ValidateUpdate(old runtime.Object) (admission.Warnings, error) {
	validations := group.updateValidations()
	var temp any = group
	if runtimeValidator, ok := temp.(genruntime.Validator); ok {
		validations = append(validations, runtimeValidator.UpdateValidations()...)
	}
	return genruntime.ValidateUpdate(old, validations)
}

// createValidations validates the creation of the resource
func (group *NamespacesEventhubsConsumerGroup) createValidations() []func() (admission.Warnings, error) {
	return []func() (admission.Warnings, error){group.validateResourceReferences, group.validateOwnerReference, group.validateSecretDestinations, group.validateConfigMapDestinations}
}

// deleteValidations validates the deletion of the resource
func (group *NamespacesEventhubsConsumerGroup) deleteValidations() []func() (admission.Warnings, error) {
	return nil
}

// updateValidations validates the update of the resource
func (group *NamespacesEventhubsConsumerGroup) updateValidations() []func(old runtime.Object) (admission.Warnings, error) {
	return []func(old runtime.Object) (admission.Warnings, error){
		func(old runtime.Object) (admission.Warnings, error) {
			return group.validateResourceReferences()
		},
		group.validateWriteOnceProperties,
		func(old runtime.Object) (admission.Warnings, error) {
			return group.validateOwnerReference()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return group.validateSecretDestinations()
		},
		func(old runtime.Object) (admission.Warnings, error) {
			return group.validateConfigMapDestinations()
		},
	}
}

// validateConfigMapDestinations validates there are no colliding genruntime.ConfigMapDestinations
func (group *NamespacesEventhubsConsumerGroup) validateConfigMapDestinations() (admission.Warnings, error) {
	if group.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return configmaps.ValidateDestinations(group, nil, group.Spec.OperatorSpec.ConfigMapExpressions)
}

// validateOwnerReference validates the owner field
func (group *NamespacesEventhubsConsumerGroup) validateOwnerReference() (admission.Warnings, error) {
	return genruntime.ValidateOwner(group)
}

// validateResourceReferences validates all resource references
func (group *NamespacesEventhubsConsumerGroup) validateResourceReferences() (admission.Warnings, error) {
	refs, err := reflecthelpers.FindResourceReferences(&group.Spec)
	if err != nil {
		return nil, err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateSecretDestinations validates there are no colliding genruntime.SecretDestination's
func (group *NamespacesEventhubsConsumerGroup) validateSecretDestinations() (admission.Warnings, error) {
	if group.Spec.OperatorSpec == nil {
		return nil, nil
	}
	return secrets.ValidateDestinations(group, nil, group.Spec.OperatorSpec.SecretExpressions)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (group *NamespacesEventhubsConsumerGroup) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*NamespacesEventhubsConsumerGroup)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, group)
}

// AssignProperties_From_NamespacesEventhubsConsumerGroup populates our NamespacesEventhubsConsumerGroup from the provided source NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_From_NamespacesEventhubsConsumerGroup(source *storage.NamespacesEventhubsConsumerGroup) error {

	// ObjectMeta
	group.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec NamespacesEventhubsConsumerGroup_Spec
	err := spec.AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec(&source.Spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec() to populate field Spec")
	}
	group.Spec = spec

	// Status
	var status NamespacesEventhubsConsumerGroup_STATUS
	err = status.AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS(&source.Status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS() to populate field Status")
	}
	group.Status = status

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsConsumerGroup populates the provided destination NamespacesEventhubsConsumerGroup from our NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_To_NamespacesEventhubsConsumerGroup(destination *storage.NamespacesEventhubsConsumerGroup) error {

	// ObjectMeta
	destination.ObjectMeta = *group.ObjectMeta.DeepCopy()

	// Spec
	var spec storage.NamespacesEventhubsConsumerGroup_Spec
	err := group.Spec.AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec(&spec)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status storage.NamespacesEventhubsConsumerGroup_STATUS
	err = group.Status.AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS(&status)
	if err != nil {
		return eris.Wrap(err, "calling AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *NamespacesEventhubsConsumerGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion(),
		Kind:    "NamespacesEventhubsConsumerGroup",
	}
}

// +kubebuilder:object:root=true
// Generator information:
// - Generated from: /eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/consumergroups.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.EventHub/namespaces/{namespaceName}/eventhubs/{eventHubName}/consumergroups/{consumerGroupName}
type NamespacesEventhubsConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsConsumerGroup `json:"items"`
}

type NamespacesEventhubsConsumerGroup_Spec struct {
	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// OperatorSpec: The specification for configuring operator behavior. This field is interpreted by the operator and not
	// passed directly to Azure
	OperatorSpec *NamespacesEventhubsConsumerGroupOperatorSpec `json:"operatorSpec,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`

	// UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
	// used to store descriptive data, such as list of teams and their contact information also user-defined configuration
	// settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}

var _ genruntime.ARMTransformer = &NamespacesEventhubsConsumerGroup_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (group *NamespacesEventhubsConsumerGroup_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if group == nil {
		return nil, nil
	}
	result := &arm.NamespacesEventhubsConsumerGroup_Spec{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if group.UserMetadata != nil {
		result.Properties = &arm.Namespaces_Eventhubs_Consumergroup_Properties_Spec{}
	}
	if group.UserMetadata != nil {
		userMetadata := *group.UserMetadata
		result.Properties.UserMetadata = &userMetadata
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (group *NamespacesEventhubsConsumerGroup_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.NamespacesEventhubsConsumerGroup_Spec{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (group *NamespacesEventhubsConsumerGroup_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.NamespacesEventhubsConsumerGroup_Spec)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.NamespacesEventhubsConsumerGroup_Spec, got %T", armInput)
	}

	// Set property "AzureName":
	group.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// no assignment for property "OperatorSpec"

	// Set property "Owner":
	group.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "UserMetadata":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UserMetadata != nil {
			userMetadata := *typedInput.Properties.UserMetadata
			group.UserMetadata = &userMetadata
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &NamespacesEventhubsConsumerGroup_Spec{}

// ConvertSpecFrom populates our NamespacesEventhubsConsumerGroup_Spec from the provided source
func (group *NamespacesEventhubsConsumerGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*storage.NamespacesEventhubsConsumerGroup_Spec)
	if ok {
		// Populate our instance from source
		return group.AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec(src)
	}

	// Convert to an intermediate form
	src = &storage.NamespacesEventhubsConsumerGroup_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = group.AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our NamespacesEventhubsConsumerGroup_Spec
func (group *NamespacesEventhubsConsumerGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*storage.NamespacesEventhubsConsumerGroup_Spec)
	if ok {
		// Populate destination from our instance
		return group.AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &storage.NamespacesEventhubsConsumerGroup_Spec{}
	err := group.AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertSpecTo()")
	}

	// Update dst from our instance
	err = dst.ConvertSpecTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertSpecTo()")
	}

	return nil
}

// AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec populates our NamespacesEventhubsConsumerGroup_Spec from the provided source NamespacesEventhubsConsumerGroup_Spec
func (group *NamespacesEventhubsConsumerGroup_Spec) AssignProperties_From_NamespacesEventhubsConsumerGroup_Spec(source *storage.NamespacesEventhubsConsumerGroup_Spec) error {

	// AzureName
	group.AzureName = source.AzureName

	// OperatorSpec
	if source.OperatorSpec != nil {
		var operatorSpec NamespacesEventhubsConsumerGroupOperatorSpec
		err := operatorSpec.AssignProperties_From_NamespacesEventhubsConsumerGroupOperatorSpec(source.OperatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_NamespacesEventhubsConsumerGroupOperatorSpec() to populate field OperatorSpec")
		}
		group.OperatorSpec = &operatorSpec
	} else {
		group.OperatorSpec = nil
	}

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		group.Owner = &owner
	} else {
		group.Owner = nil
	}

	// UserMetadata
	group.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec populates the provided destination NamespacesEventhubsConsumerGroup_Spec from our NamespacesEventhubsConsumerGroup_Spec
func (group *NamespacesEventhubsConsumerGroup_Spec) AssignProperties_To_NamespacesEventhubsConsumerGroup_Spec(destination *storage.NamespacesEventhubsConsumerGroup_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = group.AzureName

	// OperatorSpec
	if group.OperatorSpec != nil {
		var operatorSpec storage.NamespacesEventhubsConsumerGroupOperatorSpec
		err := group.OperatorSpec.AssignProperties_To_NamespacesEventhubsConsumerGroupOperatorSpec(&operatorSpec)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_NamespacesEventhubsConsumerGroupOperatorSpec() to populate field OperatorSpec")
		}
		destination.OperatorSpec = &operatorSpec
	} else {
		destination.OperatorSpec = nil
	}

	// OriginalVersion
	destination.OriginalVersion = group.OriginalVersion()

	// Owner
	if group.Owner != nil {
		owner := group.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// UserMetadata
	destination.UserMetadata = genruntime.ClonePointerToString(group.UserMetadata)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Initialize_From_NamespacesEventhubsConsumerGroup_STATUS populates our NamespacesEventhubsConsumerGroup_Spec from the provided source NamespacesEventhubsConsumerGroup_STATUS
func (group *NamespacesEventhubsConsumerGroup_Spec) Initialize_From_NamespacesEventhubsConsumerGroup_STATUS(source *NamespacesEventhubsConsumerGroup_STATUS) error {

	// UserMetadata
	group.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// No error
	return nil
}

// OriginalVersion returns the original API version used to create the resource.
func (group *NamespacesEventhubsConsumerGroup_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (group *NamespacesEventhubsConsumerGroup_Spec) SetAzureName(azureName string) {
	group.AzureName = azureName
}

type NamespacesEventhubsConsumerGroup_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	// CreatedAt: Exact time the message was created.
	CreatedAt *string `json:"createdAt,omitempty"`

	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// SystemData: The system meta data relating to this resource.
	SystemData *SystemData_STATUS `json:"systemData,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.EventHub/Namespaces" or "Microsoft.EventHub/Namespaces/EventHubs"
	Type *string `json:"type,omitempty"`

	// UpdatedAt: The exact time the message was updated.
	UpdatedAt *string `json:"updatedAt,omitempty"`

	// UserMetadata: User Metadata is a placeholder to store user-defined string data with maximum length 1024. e.g. it can be
	// used to store descriptive data, such as list of teams and their contact information also user-defined configuration
	// settings can be stored.
	UserMetadata *string `json:"userMetadata,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NamespacesEventhubsConsumerGroup_STATUS{}

// ConvertStatusFrom populates our NamespacesEventhubsConsumerGroup_STATUS from the provided source
func (group *NamespacesEventhubsConsumerGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*storage.NamespacesEventhubsConsumerGroup_STATUS)
	if ok {
		// Populate our instance from source
		return group.AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS(src)
	}

	// Convert to an intermediate form
	src = &storage.NamespacesEventhubsConsumerGroup_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = group.AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS(src)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our NamespacesEventhubsConsumerGroup_STATUS
func (group *NamespacesEventhubsConsumerGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*storage.NamespacesEventhubsConsumerGroup_STATUS)
	if ok {
		// Populate destination from our instance
		return group.AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &storage.NamespacesEventhubsConsumerGroup_STATUS{}
	err := group.AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS(dst)
	if err != nil {
		return eris.Wrap(err, "initial step of conversion in ConvertStatusTo()")
	}

	// Update dst from our instance
	err = dst.ConvertStatusTo(destination)
	if err != nil {
		return eris.Wrap(err, "final step of conversion in ConvertStatusTo()")
	}

	return nil
}

var _ genruntime.FromARMConverter = &NamespacesEventhubsConsumerGroup_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (group *NamespacesEventhubsConsumerGroup_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &arm.NamespacesEventhubsConsumerGroup_STATUS{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (group *NamespacesEventhubsConsumerGroup_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(arm.NamespacesEventhubsConsumerGroup_STATUS)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected arm.NamespacesEventhubsConsumerGroup_STATUS, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "CreatedAt":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedAt != nil {
			createdAt := *typedInput.Properties.CreatedAt
			group.CreatedAt = &createdAt
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		group.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		group.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		group.Name = &name
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		group.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		group.Type = &typeVar
	}

	// Set property "UpdatedAt":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedAt != nil {
			updatedAt := *typedInput.Properties.UpdatedAt
			group.UpdatedAt = &updatedAt
		}
	}

	// Set property "UserMetadata":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UserMetadata != nil {
			userMetadata := *typedInput.Properties.UserMetadata
			group.UserMetadata = &userMetadata
		}
	}

	// No error
	return nil
}

// AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS populates our NamespacesEventhubsConsumerGroup_STATUS from the provided source NamespacesEventhubsConsumerGroup_STATUS
func (group *NamespacesEventhubsConsumerGroup_STATUS) AssignProperties_From_NamespacesEventhubsConsumerGroup_STATUS(source *storage.NamespacesEventhubsConsumerGroup_STATUS) error {

	// Conditions
	group.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedAt
	group.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// Id
	group.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	group.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	group.Name = genruntime.ClonePointerToString(source.Name)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		group.SystemData = &systemDatum
	} else {
		group.SystemData = nil
	}

	// Type
	group.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	group.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// UserMetadata
	group.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS populates the provided destination NamespacesEventhubsConsumerGroup_STATUS from our NamespacesEventhubsConsumerGroup_STATUS
func (group *NamespacesEventhubsConsumerGroup_STATUS) AssignProperties_To_NamespacesEventhubsConsumerGroup_STATUS(destination *storage.NamespacesEventhubsConsumerGroup_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(group.Conditions)

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(group.CreatedAt)

	// Id
	destination.Id = genruntime.ClonePointerToString(group.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(group.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(group.Name)

	// SystemData
	if group.SystemData != nil {
		var systemDatum storage.SystemData_STATUS
		err := group.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return eris.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(group.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(group.UpdatedAt)

	// UserMetadata
	destination.UserMetadata = genruntime.ClonePointerToString(group.UserMetadata)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Details for configuring operator behavior. Fields in this struct are interpreted by the operator directly rather than being passed to Azure
type NamespacesEventhubsConsumerGroupOperatorSpec struct {
	// ConfigMapExpressions: configures where to place operator written dynamic ConfigMaps (created with CEL expressions).
	ConfigMapExpressions []*core.DestinationExpression `json:"configMapExpressions,omitempty"`

	// SecretExpressions: configures where to place operator written dynamic secrets (created with CEL expressions).
	SecretExpressions []*core.DestinationExpression `json:"secretExpressions,omitempty"`
}

// AssignProperties_From_NamespacesEventhubsConsumerGroupOperatorSpec populates our NamespacesEventhubsConsumerGroupOperatorSpec from the provided source NamespacesEventhubsConsumerGroupOperatorSpec
func (operator *NamespacesEventhubsConsumerGroupOperatorSpec) AssignProperties_From_NamespacesEventhubsConsumerGroupOperatorSpec(source *storage.NamespacesEventhubsConsumerGroupOperatorSpec) error {

	// ConfigMapExpressions
	if source.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(source.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range source.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		operator.ConfigMapExpressions = configMapExpressionList
	} else {
		operator.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if source.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(source.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range source.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		operator.SecretExpressions = secretExpressionList
	} else {
		operator.SecretExpressions = nil
	}

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsConsumerGroupOperatorSpec populates the provided destination NamespacesEventhubsConsumerGroupOperatorSpec from our NamespacesEventhubsConsumerGroupOperatorSpec
func (operator *NamespacesEventhubsConsumerGroupOperatorSpec) AssignProperties_To_NamespacesEventhubsConsumerGroupOperatorSpec(destination *storage.NamespacesEventhubsConsumerGroupOperatorSpec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// ConfigMapExpressions
	if operator.ConfigMapExpressions != nil {
		configMapExpressionList := make([]*core.DestinationExpression, len(operator.ConfigMapExpressions))
		for configMapExpressionIndex, configMapExpressionItem := range operator.ConfigMapExpressions {
			// Shadow the loop variable to avoid aliasing
			configMapExpressionItem := configMapExpressionItem
			if configMapExpressionItem != nil {
				configMapExpression := *configMapExpressionItem.DeepCopy()
				configMapExpressionList[configMapExpressionIndex] = &configMapExpression
			} else {
				configMapExpressionList[configMapExpressionIndex] = nil
			}
		}
		destination.ConfigMapExpressions = configMapExpressionList
	} else {
		destination.ConfigMapExpressions = nil
	}

	// SecretExpressions
	if operator.SecretExpressions != nil {
		secretExpressionList := make([]*core.DestinationExpression, len(operator.SecretExpressions))
		for secretExpressionIndex, secretExpressionItem := range operator.SecretExpressions {
			// Shadow the loop variable to avoid aliasing
			secretExpressionItem := secretExpressionItem
			if secretExpressionItem != nil {
				secretExpression := *secretExpressionItem.DeepCopy()
				secretExpressionList[secretExpressionIndex] = &secretExpression
			} else {
				secretExpressionList[secretExpressionIndex] = nil
			}
		}
		destination.SecretExpressions = secretExpressionList
	} else {
		destination.SecretExpressions = nil
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
	SchemeBuilder.Register(&NamespacesEventhubsConsumerGroup{}, &NamespacesEventhubsConsumerGroupList{})
}
