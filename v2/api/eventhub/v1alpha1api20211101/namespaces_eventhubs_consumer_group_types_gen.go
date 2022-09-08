// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20211101

import (
	"fmt"
	alpha20211101s "github.com/Azure/azure-service-operator/v2/api/eventhub/v1alpha1api20211101storage"
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
// Deprecated version of NamespacesEventhubsConsumerGroup. Use v1beta20211101.NamespacesEventhubsConsumerGroup instead
type NamespacesEventhubsConsumerGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Namespaces_Eventhubs_Consumergroup_Spec `json:"spec,omitempty"`
	Status            ConsumerGroup_STATUS                    `json:"status,omitempty"`
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
	var source alpha20211101s.NamespacesEventhubsConsumerGroup

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
	var destination alpha20211101s.NamespacesEventhubsConsumerGroup
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

// +kubebuilder:webhook:path=/mutate-eventhub-azure-com-v1alpha1api20211101-namespaceseventhubsconsumergroup,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsconsumergroups,verbs=create;update,versions=v1alpha1api20211101,name=default.v1alpha1api20211101.namespaceseventhubsconsumergroups.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &NamespacesEventhubsConsumerGroup{}

// Default applies defaults to the NamespacesEventhubsConsumerGroup resource
func (group *NamespacesEventhubsConsumerGroup) Default() {
	group.defaultImpl()
	var temp interface{} = group
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
	return &ConsumerGroup_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (group *NamespacesEventhubsConsumerGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return &genruntime.ResourceReference{
		Group: ownerGroup,
		Kind:  ownerKind,
		Name:  group.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (group *NamespacesEventhubsConsumerGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ConsumerGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st ConsumerGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventhub-azure-com-v1alpha1api20211101-namespaceseventhubsconsumergroup,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventhub.azure.com,resources=namespaceseventhubsconsumergroups,verbs=create;update,versions=v1alpha1api20211101,name=validate.v1alpha1api20211101.namespaceseventhubsconsumergroups.eventhub.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &NamespacesEventhubsConsumerGroup{}

// ValidateCreate validates the creation of the resource
func (group *NamespacesEventhubsConsumerGroup) ValidateCreate() error {
	validations := group.createValidations()
	var temp interface{} = group
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
func (group *NamespacesEventhubsConsumerGroup) ValidateDelete() error {
	validations := group.deleteValidations()
	var temp interface{} = group
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
func (group *NamespacesEventhubsConsumerGroup) ValidateUpdate(old runtime.Object) error {
	validations := group.updateValidations()
	var temp interface{} = group
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
func (group *NamespacesEventhubsConsumerGroup) createValidations() []func() error {
	return []func() error{group.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (group *NamespacesEventhubsConsumerGroup) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (group *NamespacesEventhubsConsumerGroup) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return group.validateResourceReferences()
		},
		group.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (group *NamespacesEventhubsConsumerGroup) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&group.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (group *NamespacesEventhubsConsumerGroup) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*NamespacesEventhubsConsumerGroup)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, group)
}

// AssignProperties_From_NamespacesEventhubsConsumerGroup populates our NamespacesEventhubsConsumerGroup from the provided source NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_From_NamespacesEventhubsConsumerGroup(source *alpha20211101s.NamespacesEventhubsConsumerGroup) error {

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
	var status ConsumerGroup_STATUS
	err = status.AssignProperties_From_ConsumerGroup_STATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_From_ConsumerGroup_STATUS() to populate field Status")
	}
	group.Status = status

	// No error
	return nil
}

// AssignProperties_To_NamespacesEventhubsConsumerGroup populates the provided destination NamespacesEventhubsConsumerGroup from our NamespacesEventhubsConsumerGroup
func (group *NamespacesEventhubsConsumerGroup) AssignProperties_To_NamespacesEventhubsConsumerGroup(destination *alpha20211101s.NamespacesEventhubsConsumerGroup) error {

	// ObjectMeta
	destination.ObjectMeta = *group.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20211101s.Namespaces_Eventhubs_Consumergroup_Spec
	err := group.Spec.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20211101s.ConsumerGroup_STATUS
	err = group.Status.AssignProperties_To_ConsumerGroup_STATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignProperties_To_ConsumerGroup_STATUS() to populate field Status")
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
// Deprecated version of NamespacesEventhubsConsumerGroup. Use v1beta20211101.NamespacesEventhubsConsumerGroup instead
type NamespacesEventhubsConsumerGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NamespacesEventhubsConsumerGroup `json:"items"`
}

// Deprecated version of ConsumerGroup_STATUS. Use v1beta20211101.ConsumerGroup_STATUS instead
type ConsumerGroup_STATUS struct {
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

var _ genruntime.ConvertibleStatus = &ConsumerGroup_STATUS{}

// ConvertStatusFrom populates our ConsumerGroup_STATUS from the provided source
func (group *ConsumerGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20211101s.ConsumerGroup_STATUS)
	if ok {
		// Populate our instance from source
		return group.AssignProperties_From_ConsumerGroup_STATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20211101s.ConsumerGroup_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = group.AssignProperties_From_ConsumerGroup_STATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our ConsumerGroup_STATUS
func (group *ConsumerGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20211101s.ConsumerGroup_STATUS)
	if ok {
		// Populate destination from our instance
		return group.AssignProperties_To_ConsumerGroup_STATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20211101s.ConsumerGroup_STATUS{}
	err := group.AssignProperties_To_ConsumerGroup_STATUS(dst)
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

var _ genruntime.FromARMConverter = &ConsumerGroup_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (group *ConsumerGroup_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ConsumerGroup_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (group *ConsumerGroup_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ConsumerGroup_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ConsumerGroup_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘CreatedAt’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.CreatedAt != nil {
			createdAt := *typedInput.Properties.CreatedAt
			group.CreatedAt = &createdAt
		}
	}

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		group.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		group.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		group.Name = &name
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_STATUS
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		group.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		group.Type = &typeVar
	}

	// Set property ‘UpdatedAt’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.UpdatedAt != nil {
			updatedAt := *typedInput.Properties.UpdatedAt
			group.UpdatedAt = &updatedAt
		}
	}

	// Set property ‘UserMetadata’:
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

// AssignProperties_From_ConsumerGroup_STATUS populates our ConsumerGroup_STATUS from the provided source ConsumerGroup_STATUS
func (group *ConsumerGroup_STATUS) AssignProperties_From_ConsumerGroup_STATUS(source *alpha20211101s.ConsumerGroup_STATUS) error {

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
			return errors.Wrap(err, "calling AssignProperties_From_SystemData_STATUS() to populate field SystemData")
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

// AssignProperties_To_ConsumerGroup_STATUS populates the provided destination ConsumerGroup_STATUS from our ConsumerGroup_STATUS
func (group *ConsumerGroup_STATUS) AssignProperties_To_ConsumerGroup_STATUS(destination *alpha20211101s.ConsumerGroup_STATUS) error {
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
		var systemDatum alpha20211101s.SystemData_STATUS
		err := group.SystemData.AssignProperties_To_SystemData_STATUS(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignProperties_To_SystemData_STATUS() to populate field SystemData")
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

type Namespaces_Eventhubs_Consumergroup_Spec struct {
	// +kubebuilder:validation:MaxLength=50
	// +kubebuilder:validation:MinLength=1
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName string  `json:"azureName,omitempty"`
	Location  *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a eventhub.azure.com/NamespacesEventhub resource
	Owner        *genruntime.KnownResourceReference `group:"eventhub.azure.com" json:"owner,omitempty" kind:"NamespacesEventhub"`
	Tags         map[string]string                  `json:"tags,omitempty"`
	UserMetadata *string                            `json:"userMetadata,omitempty"`
}

var _ genruntime.ARMTransformer = &Namespaces_Eventhubs_Consumergroup_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if consumergroup == nil {
		return nil, nil
	}
	result := &Namespaces_Eventhubs_Consumergroup_SpecARM{}

	// Set property ‘Location’:
	if consumergroup.Location != nil {
		location := *consumergroup.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if consumergroup.UserMetadata != nil {
		result.Properties = &ConsumerGroupPropertiesARM{}
	}
	if consumergroup.UserMetadata != nil {
		userMetadata := *consumergroup.UserMetadata
		result.Properties.UserMetadata = &userMetadata
	}

	// Set property ‘Tags’:
	if consumergroup.Tags != nil {
		result.Tags = make(map[string]string, len(consumergroup.Tags))
		for key, value := range consumergroup.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &Namespaces_Eventhubs_Consumergroup_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(Namespaces_Eventhubs_Consumergroup_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected Namespaces_Eventhubs_Consumergroup_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	consumergroup.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		consumergroup.Location = &location
	}

	// Set property ‘Owner’:
	consumergroup.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		consumergroup.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			consumergroup.Tags[key] = value
		}
	}

	// Set property ‘UserMetadata’:
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
	src, ok := source.(*alpha20211101s.Namespaces_Eventhubs_Consumergroup_Spec)
	if ok {
		// Populate our instance from source
		return consumergroup.AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(src)
	}

	// Convert to an intermediate form
	src = &alpha20211101s.Namespaces_Eventhubs_Consumergroup_Spec{}
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
	dst, ok := destination.(*alpha20211101s.Namespaces_Eventhubs_Consumergroup_Spec)
	if ok {
		// Populate destination from our instance
		return consumergroup.AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20211101s.Namespaces_Eventhubs_Consumergroup_Spec{}
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
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) AssignProperties_From_Namespaces_Eventhubs_Consumergroup_Spec(source *alpha20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error {

	// AzureName
	consumergroup.AzureName = source.AzureName

	// Location
	consumergroup.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		consumergroup.Owner = &owner
	} else {
		consumergroup.Owner = nil
	}

	// Tags
	consumergroup.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// UserMetadata
	consumergroup.UserMetadata = genruntime.ClonePointerToString(source.UserMetadata)

	// No error
	return nil
}

// AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec populates the provided destination Namespaces_Eventhubs_Consumergroup_Spec from our Namespaces_Eventhubs_Consumergroup_Spec
func (consumergroup *Namespaces_Eventhubs_Consumergroup_Spec) AssignProperties_To_Namespaces_Eventhubs_Consumergroup_Spec(destination *alpha20211101s.Namespaces_Eventhubs_Consumergroup_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = consumergroup.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(consumergroup.Location)

	// OriginalVersion
	destination.OriginalVersion = consumergroup.OriginalVersion()

	// Owner
	if consumergroup.Owner != nil {
		owner := consumergroup.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(consumergroup.Tags)

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

func init() {
	SchemeBuilder.Register(&NamespacesEventhubsConsumerGroup{}, &NamespacesEventhubsConsumerGroupList{})
}
