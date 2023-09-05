// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

import (
	"fmt"
	v1beta20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1beta20211101storage"
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
// Deprecated version of NamespacesEventhubsConsumerGroup. Use v1api20211101.NamespacesEventhubsConsumerGroup instead
type NamespacesEventhubsConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Eventhubs_Consumergroup_Spec   `json:"spec,omitempty"`
	Status            Namespaces_Eventhubs_Consumergroup_STATUS `json:"status,omitempty"`
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
	// intermediate variable for conversion
	var source v1beta20211101s.NamespacesEventhubsConsumerGroup

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = group.AssignProperties_From_NamespacesEventhubsConsumerGroup(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to group")
	}

	return nil
}

// ConvertTo populates the provided hub NamespacesEventhubsConsumerGroup from our NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination v1beta20211101s.NamespacesEventhubsConsumerGroup
	err := group.AssignProperties_To_NamespacesEventhubsConsumerGroup(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from group")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1beta20211101-namespaceseventhubsconsumergroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsconsumergroups,verbs=create;update,versions=v1beta20211101,name=default.v1beta20211101.namespaceseventhubsconsumergroups.eventhub.azure.com,admissionReviewVersions=v1

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

var _ genruntime.KubernetesResource = &NamespacesEventhubsConsumerGroup{}

// AzureName returns the Azure name of the resource
func (group *NamespacesEventhubsConsumerGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (group NamespacesEventhubsConsumerGroup) GetAPIVersion() string {
	return string(APIVersion_Value)
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

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
func (group *NamespacesEventhubsConsumerGroup) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs/consumergroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *NamespacesEventhubsConsumerGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Namespaces_Eventhubs_Consumergroup_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (group *NamespacesEventhubsConsumerGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return group.Spec.Owner.AsResourceReference(ownerGroup, ownerKind)
}

// SetStatus sets the status of this resource
func (group *NamespacesEventhubsConsumerGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Namespaces_Eventhubs_Consumergroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st Namespaces_Eventhubs_Consumergroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1beta20211101-namespaceseventhubsconsumergroup,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsconsumergroups,verbs=create;update,versions=v1beta20211101,name=validate.v1beta20211101.namespaceseventhubsconsumergroups.eventhub.azure.com,admissionReviewVersions=v1

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
	return []func() (admission.Warnings, error){group.validateResourceReferences, group.validateOwnerReference}
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
	}
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

// validateWriteOnceProperties validates all WriteOnce properties
func (group *NamespacesEventhubsConsumerGroup) validateWriteOnceProperties(old runtime.Object) (admission.Warnings, error) {
	oldObj, ok := old.(*NamespacesEventhubsConsumerGroup)
	if !ok {
		return nil, nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, group)
}

// AssignProperties_From_NamespacesEventhubsConsumerGroup populates our NamespacesEventhubsConsumerGroup from the provided source NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_From_NamespacesEventhubsConsumerGroup(source *v1beta20211101s.NamespacesEventhubsConsumerGroup) error {

	// ObjectMeta
	group.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec Namespaces_Eventhubs_Consumergroup_Spec
	err := spec.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec() to populate field Spec")
	}
	group.Spec = spec

	// Status
	var status Namespaces_Eventhubs_Consumergroup_STATUS
	err = status.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS() to populate field Status")
	}
	group.Status = status

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsConsumerGroup populates the provided destination NamespacesEventhubsConsumerGroup from our NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_To_NamespacesEventhubsConsumerGroup(destination *v1beta20211101s.NamespacesEventhubsConsumerGroup) error {

	// ObjectMeta
	destination.ObjectMeta = *group.ObjectMeta.DeepCopy()

	// Spec
	var spec v1beta20211101s.Namespaces_Eventhubs_Consumergroup_Spec
	err := group.Spec.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1beta20211101s.Namespaces_Eventhubs_Consumergroup_STATUS
	err = group.Status.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS() to populate field Status")
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
// Deprecated version of NamespacesEventhubsConsumerGroup. Use v1api20211101.NamespacesEventhubsConsumerGroup instead
type NamespacesEventhubsConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsConsumerGroup `json:"items"`
}

type Namespaces_Eventhubs_Consumergroup_Spec struct {
	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string `json:"azureName,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner        *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`
	UserMetadata *string                            `json:"userMetadata,omitempty"`
}

var _ genruntime.ARMTransformer = &Namespaces_Eventhubs_Consumergroup_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if consumergroup == nil {
		return nil, nil
	}
	result := &Namespaces_Eventhubs_Consumergroup_Spec_ARM{}

	// Set property "Name":
	result.Name = resolved.Name

	// Set property "Properties":
	if consumergroup.UserMetadata != nil {
		result.Properties = &Namespaces_Eventhubs_Consumergroup_Properties_Spec_ARM{}
	}
	if consumergroup.UserMetadata != nil {
		userMetadata := *consumergroup.UserMetadata
		result.Properties.UserMetadata = &userMetadata
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Namespaces_Eventhubs_Consumergroup_Spec_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Namespaces_Eventhubs_Consumergroup_Spec_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Namespaces_Eventhubs_Consumergroup_Spec_ARM, got %T", armInput)
	}

	// Set property "AzureName":
	consumergroup.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property "Owner":
	consumergroup.Owner = &genruntime.KnownResourceReference{
		Name:  owner.Name,
		ARMID: owner.ARMID,
	}

	// Set property "UserMetadata":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UserMetadata != nil {
			userMetadata := *typedInput.Properties.UserMetadata
			consumergroup.UserMetadata = &userMetadata
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &Namespaces_Eventhubs_Consumergroup_Spec{}

// ConvertSpecFrom populates our Namespaces_Eventhubs_Consumergroup_Spec from the provided source
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1beta20211101s.Namespaces_Eventhubs_Consumergroup_Spec)
	if ok {
		// Populate our instance from source
		return consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(src)
	}

	// Convert to an intermediate form
	src = &v1beta20211101s.Namespaces_Eventhubs_Consumergroup_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our Namespaces_Eventhubs_Consumergroup_Spec
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1beta20211101s.Namespaces_Eventhubs_Consumergroup_Spec)
	if ok {
		// Populate destination from our instance
		return consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20211101s.Namespaces_Eventhubs_Consumergroup_Spec{}
	err := consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(dst)
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

// AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec populates our Namespaces_Eventhubs_Consumergroup_Spec from the provided source Namespaces_Eventhubs_Consumergroup_Spec
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(source *v1beta20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error {

	// AzureName
	consumergroup.AzureName = source.AzureName

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		consumergroup.Owner = &owner
	} else {
		consumergroup.Owner = nil
	}

	// UserMetadata
	consumergroup.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec populates the provided destination Namespaces_Eventhubs_Consumergroup_Spec from our Namespaces_Eventhubs_Consumergroup_Spec
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(destination *v1beta20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = consumergroup.AzureName

	// OriginalVersion
	destination.OriginalVersion = consumergroup.OriginalVersion()

	// Owner
	if consumergroup.Owner != nil {
		owner := consumergroup.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// UserMetadata
	destination.UserMetadata = genruntime.ClonePointerToString(consumergroup.UserMetadata)

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
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) SetAzureName(azureName string) {
	consumergroup.AzureName = azureName
}

// Deprecated version of Namespaces_Eventhubs_Consumergroup_STATUS. Use v1api20211101.Namespaces_Eventhubs_Consumergroup_STATUS instead
type Namespaces_Eventhubs_Consumergroup_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions   []conditions.Condition `json:"conditions,omitempty"`
	CreatedAt    *string                `json:"createdAt,omitempty"`
	Id           *string                `json:"id,omitempty"`
	Location     *string                `json:"location,omitempty"`
	Name         *string                `json:"name,omitempty"`
	SystemData   *SystemData_STATUS     `json:"systemData,omitempty"`
	Type         *string                `json:"type,omitempty"`
	UpdatedAt    *string                `json:"updatedAt,omitempty"`
	UserMetadata *string                `json:"userMetadata,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Namespaces_Eventhubs_Consumergroup_STATUS{}

// ConvertStatusFrom populates our Namespaces_Eventhubs_Consumergroup_STATUS from the provided source
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1beta20211101s.Namespaces_Eventhubs_Consumergroup_STATUS)
	if ok {
		// Populate our instance from source
		return consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(src)
	}

	// Convert to an intermediate form
	src = &v1beta20211101s.Namespaces_Eventhubs_Consumergroup_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our Namespaces_Eventhubs_Consumergroup_STATUS
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1beta20211101s.Namespaces_Eventhubs_Consumergroup_STATUS)
	if ok {
		// Populate destination from our instance
		return consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &v1beta20211101s.Namespaces_Eventhubs_Consumergroup_STATUS{}
	err := consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(dst)
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

var _ genruntime.FromARMConverter = &Namespaces_Eventhubs_Consumergroup_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Namespaces_Eventhubs_Consumergroup_STATUS_ARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Namespaces_Eventhubs_Consumergroup_STATUS_ARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Namespaces_Eventhubs_Consumergroup_STATUS_ARM, got %T", armInput)
	}

	// no assignment for property "Conditions"

	// Set property "CreatedAt":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedAt != nil {
			createdAt := *typedInput.Properties.CreatedAt
			consumergroup.CreatedAt = &createdAt
		}
	}

	// Set property "Id":
	if typedInput.Id != nil {
		id := *typedInput.Id
		consumergroup.Id = &id
	}

	// Set property "Location":
	if typedInput.Location != nil {
		location := *typedInput.Location
		consumergroup.Location = &location
	}

	// Set property "Name":
	if typedInput.Name != nil {
		name := *typedInput.Name
		consumergroup.Name = &name
	}

	// Set property "SystemData":
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		consumergroup.SystemData = &systemData
	}

	// Set property "Type":
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		consumergroup.Type = &typeVar
	}

	// Set property "UpdatedAt":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedAt != nil {
			updatedAt := *typedInput.Properties.UpdatedAt
			consumergroup.UpdatedAt = &updatedAt
		}
	}

	// Set property "UserMetadata":
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UserMetadata != nil {
			userMetadata := *typedInput.Properties.UserMetadata
			consumergroup.UserMetadata = &userMetadata
		}
	}

	// No error
	return nil
}

// AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS populates our Namespaces_Eventhubs_Consumergroup_STATUS from the provided source Namespaces_Eventhubs_Consumergroup_STATUS
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) AssignProperties_From_Namespaces_Eventhubs_Consumergroup_STATUS(source *v1beta20211101s.Namespaces_Eventhubs_Consumergroup_STATUS) error {

	// Conditions
	consumergroup.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// CreatedAt
	consumergroup.CreatedAt = genruntime.ClonePointerToString(source.CreatedAt)

	// Id
	consumergroup.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	consumergroup.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	consumergroup.Name = genruntime.ClonePointerToString(source.Name)

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_STATUS
		err := systemDatum.AssignProperties_From_SystemData_STATUS(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
		}
		consumergroup.SystemData = &systemDatum
	} else {
		consumergroup.SystemData = nil
	}

	// Type
	consumergroup.Type = genruntime.ClonePointerToString(source.Type)

	// UpdatedAt
	consumergroup.UpdatedAt = genruntime.ClonePointerToString(source.UpdatedAt)

	// UserMetadata
	consumergroup.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS populates the provided destination Namespaces_Eventhubs_Consumergroup_STATUS from our Namespaces_Eventhubs_Consumergroup_STATUS
func (consumergroup *Namespaces_Eventhubs_Consumergroup_STATUS) AssignProperties_To_Namespaces_Eventhubs_Consumergroup_STATUS(destination *v1beta20211101s.Namespaces_Eventhubs_Consumergroup_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(consumergroup.Conditions)

	// CreatedAt
	destination.CreatedAt = genruntime.ClonePointerToString(consumergroup.CreatedAt)

	// Id
	destination.Id = genruntime.ClonePointerToString(consumergroup.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(consumergroup.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(consumergroup.Name)

	// SystemData
	if consumergroup.SystemData != nil {
		var systemDatum v1beta20211101s.SystemData_STATUS
		err := consumergroup.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(consumergroup.Type)

	// UpdatedAt
	destination.UpdatedAt = genruntime.ClonePointerToString(consumergroup.UpdatedAt)

	// UserMetadata
	destination.UserMetadata = genruntime.ClonePointerToString(consumergroup.UserMetadata)

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
