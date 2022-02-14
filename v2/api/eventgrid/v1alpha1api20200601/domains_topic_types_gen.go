// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20200601

import (
	"fmt"
	"github.com/Azure/azure-service-operator/v2/api/eventgrid/v1alpha1api20200601storage"
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
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopic struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainsTopics_Spec `json:"spec,omitempty"`
	Status            DomainTopic_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DomainsTopic{}

// GetConditions returns the conditions of the resource
func (topic *DomainsTopic) GetConditions() conditions.Conditions {
	return topic.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (topic *DomainsTopic) SetConditions(conditions conditions.Conditions) {
	topic.Status.Conditions = conditions
}

var _ conversion.Convertible = &DomainsTopic{}

// ConvertFrom populates our DomainsTopic from the provided hub DomainsTopic
func (topic *DomainsTopic) ConvertFrom(hub conversion.Hub) error {
	source, ok := hub.(*v1alpha1api20200601storage.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected storage:eventgrid/v1alpha1api20200601storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignPropertiesFromDomainsTopic(source)
}

// ConvertTo populates the provided hub DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) ConvertTo(hub conversion.Hub) error {
	destination, ok := hub.(*v1alpha1api20200601storage.DomainsTopic)
	if !ok {
		return fmt.Errorf("expected storage:eventgrid/v1alpha1api20200601storage/DomainsTopic but received %T instead", hub)
	}

	return topic.AssignPropertiesToDomainsTopic(destination)
}

// +kubebuilder:webhook:path=/mutate-eventgrid-azure-com-v1alpha1api20200601-domainstopic,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=domainstopics,verbs=create;update,versions=v1alpha1api20200601,name=default.v1alpha1api20200601.domainstopics.eventgrid.azure.com,admissionReviewVersions=v1beta1

var _ admission.Defaulter = &DomainsTopic{}

// Default applies defaults to the DomainsTopic resource
func (topic *DomainsTopic) Default() {
	topic.defaultImpl()
	var temp interface{} = topic
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultAzureName defaults the Azure name of the resource to the Kubernetes name
func (topic *DomainsTopic) defaultAzureName() {
	if topic.Spec.AzureName == "" {
		topic.Spec.AzureName = topic.Name
	}
}

// defaultImpl applies the code generated defaults to the DomainsTopic resource
func (topic *DomainsTopic) defaultImpl() { topic.defaultAzureName() }

var _ genruntime.KubernetesResource = &DomainsTopic{}

// AzureName returns the Azure name of the resource
func (topic *DomainsTopic) AzureName() string {
	return topic.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-06-01"
func (topic DomainsTopic) GetAPIVersion() string {
	return "2020-06-01"
}

// GetResourceKind returns the kind of the resource
func (topic *DomainsTopic) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (topic *DomainsTopic) GetSpec() genruntime.ConvertibleSpec {
	return &topic.Spec
}

// GetStatus returns the status of this resource
func (topic *DomainsTopic) GetStatus() genruntime.ConvertibleStatus {
	return &topic.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventGrid/domains/topics"
func (topic *DomainsTopic) GetType() string {
	return "Microsoft.EventGrid/domains/topics"
}

// NewEmptyStatus returns a new empty (blank) status
func (topic *DomainsTopic) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DomainTopic_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (topic *DomainsTopic) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(topic.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  topic.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (topic *DomainsTopic) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DomainTopic_Status); ok {
		topic.Status = *st
		return nil
	}

	// Convert status to required version
	var st DomainTopic_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	topic.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-eventgrid-azure-com-v1alpha1api20200601-domainstopic,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=eventgrid.azure.com,resources=domainstopics,verbs=create;update,versions=v1alpha1api20200601,name=validate.v1alpha1api20200601.domainstopics.eventgrid.azure.com,admissionReviewVersions=v1beta1

var _ admission.Validator = &DomainsTopic{}

// ValidateCreate validates the creation of the resource
func (topic *DomainsTopic) ValidateCreate() error {
	validations := topic.createValidations()
	var temp interface{} = topic
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
func (topic *DomainsTopic) ValidateDelete() error {
	validations := topic.deleteValidations()
	var temp interface{} = topic
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
func (topic *DomainsTopic) ValidateUpdate(old runtime.Object) error {
	validations := topic.updateValidations()
	var temp interface{} = topic
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
func (topic *DomainsTopic) createValidations() []func() error {
	return []func() error{topic.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (topic *DomainsTopic) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (topic *DomainsTopic) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return topic.validateResourceReferences()
		},
	}
}

// validateResourceReferences validates all resource references
func (topic *DomainsTopic) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&topic.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// AssignPropertiesFromDomainsTopic populates our DomainsTopic from the provided source DomainsTopic
func (topic *DomainsTopic) AssignPropertiesFromDomainsTopic(source *v1alpha1api20200601storage.DomainsTopic) error {

	// ObjectMeta
	topic.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec DomainsTopics_Spec
	err := spec.AssignPropertiesFromDomainsTopicsSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDomainsTopicsSpec() to populate field Spec")
	}
	topic.Spec = spec

	// Status
	var status DomainTopic_Status
	err = status.AssignPropertiesFromDomainTopicStatus(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromDomainTopicStatus() to populate field Status")
	}
	topic.Status = status

	// No error
	return nil
}

// AssignPropertiesToDomainsTopic populates the provided destination DomainsTopic from our DomainsTopic
func (topic *DomainsTopic) AssignPropertiesToDomainsTopic(destination *v1alpha1api20200601storage.DomainsTopic) error {

	// ObjectMeta
	destination.ObjectMeta = *topic.ObjectMeta.DeepCopy()

	// Spec
	var spec v1alpha1api20200601storage.DomainsTopics_Spec
	err := topic.Spec.AssignPropertiesToDomainsTopicsSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDomainsTopicsSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status v1alpha1api20200601storage.DomainTopic_Status
	err = topic.Status.AssignPropertiesToDomainTopicStatus(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToDomainTopicStatus() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (topic *DomainsTopic) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: topic.Spec.OriginalVersion(),
		Kind:    "DomainsTopic",
	}
}

// +kubebuilder:object:root=true
//Generated from: https://schema.management.azure.com/schemas/2020-06-01/Microsoft.EventGrid.json#/resourceDefinitions/domains_topics
type DomainsTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DomainsTopic `json:"items"`
}

type DomainTopic_Status struct {
	//Conditions: The observed state of the resource
	Conditions []conditions.Condition `json:"conditions,omitempty"`

	//Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	//Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	//ProvisioningState: Provisioning state of the domain topic.
	ProvisioningState *DomainTopicPropertiesStatusProvisioningState `json:"provisioningState,omitempty"`

	//SystemData: The system metadata relating to Domain Topic resource.
	SystemData *SystemData_Status `json:"systemData,omitempty"`

	//Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DomainTopic_Status{}

// ConvertStatusFrom populates our DomainTopic_Status from the provided source
func (topic *DomainTopic_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*v1alpha1api20200601storage.DomainTopic_Status)
	if ok {
		// Populate our instance from source
		return topic.AssignPropertiesFromDomainTopicStatus(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20200601storage.DomainTopic_Status{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = topic.AssignPropertiesFromDomainTopicStatus(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our DomainTopic_Status
func (topic *DomainTopic_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*v1alpha1api20200601storage.DomainTopic_Status)
	if ok {
		// Populate destination from our instance
		return topic.AssignPropertiesToDomainTopicStatus(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20200601storage.DomainTopic_Status{}
	err := topic.AssignPropertiesToDomainTopicStatus(dst)
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

var _ genruntime.FromARMConverter = &DomainTopic_Status{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (topic *DomainTopic_Status) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DomainTopic_StatusARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (topic *DomainTopic_Status) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DomainTopic_StatusARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DomainTopic_StatusARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		topic.Id = &id
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		topic.Name = &name
	}

	// Set property ‘ProvisioningState’:
	// copying flattened property:
	if typedInput.Properties != nil {
		if typedInput.Properties.ProvisioningState != nil {
			provisioningState := *typedInput.Properties.ProvisioningState
			topic.ProvisioningState = &provisioningState
		}
	}

	// Set property ‘SystemData’:
	if typedInput.SystemData != nil {
		var systemData1 SystemData_Status
		err := systemData1.PopulateFromARM(owner, *typedInput.SystemData)
		if err != nil {
			return err
		}
		systemData := systemData1
		topic.SystemData = &systemData
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		topic.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromDomainTopicStatus populates our DomainTopic_Status from the provided source DomainTopic_Status
func (topic *DomainTopic_Status) AssignPropertiesFromDomainTopicStatus(source *v1alpha1api20200601storage.DomainTopic_Status) error {

	// Conditions
	topic.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	topic.Id = genruntime.ClonePointerToString(source.Id)

	// Name
	topic.Name = genruntime.ClonePointerToString(source.Name)

	// ProvisioningState
	if source.ProvisioningState != nil {
		provisioningState := DomainTopicPropertiesStatusProvisioningState(*source.ProvisioningState)
		topic.ProvisioningState = &provisioningState
	} else {
		topic.ProvisioningState = nil
	}

	// SystemData
	if source.SystemData != nil {
		var systemDatum SystemData_Status
		err := systemDatum.AssignPropertiesFromSystemDataStatus(source.SystemData)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesFromSystemDataStatus() to populate field SystemData")
		}
		topic.SystemData = &systemDatum
	} else {
		topic.SystemData = nil
	}

	// Type
	topic.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToDomainTopicStatus populates the provided destination DomainTopic_Status from our DomainTopic_Status
func (topic *DomainTopic_Status) AssignPropertiesToDomainTopicStatus(destination *v1alpha1api20200601storage.DomainTopic_Status) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(topic.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(topic.Id)

	// Name
	destination.Name = genruntime.ClonePointerToString(topic.Name)

	// ProvisioningState
	if topic.ProvisioningState != nil {
		provisioningState := string(*topic.ProvisioningState)
		destination.ProvisioningState = &provisioningState
	} else {
		destination.ProvisioningState = nil
	}

	// SystemData
	if topic.SystemData != nil {
		var systemDatum v1alpha1api20200601storage.SystemData_Status
		err := topic.SystemData.AssignPropertiesToSystemDataStatus(&systemDatum)
		if err != nil {
			return errors.Wrap(err, "calling AssignPropertiesToSystemDataStatus() to populate field SystemData")
		}
		destination.SystemData = &systemDatum
	} else {
		destination.SystemData = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(topic.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"2020-06-01"}
type DomainsTopicsSpecAPIVersion string

const DomainsTopicsSpecAPIVersion20200601 = DomainsTopicsSpecAPIVersion("2020-06-01")

type DomainsTopics_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	//doesn't have to be.
	AzureName string `json:"azureName"`

	//Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	Owner genruntime.KnownResourceReference `group:"eventgrid.azure.com" json:"owner" kind:"Domain"`

	//Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &DomainsTopics_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (topics *DomainsTopics_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if topics == nil {
		return nil, nil
	}
	var result DomainsTopics_SpecARM

	// Set property ‘Location’:
	if topics.Location != nil {
		location := *topics.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Tags’:
	if topics.Tags != nil {
		result.Tags = make(map[string]string)
		for key, value := range topics.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (topics *DomainsTopics_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &DomainsTopics_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (topics *DomainsTopics_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(DomainsTopics_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected DomainsTopics_SpecARM, got %T", armInput)
	}

	// Set property ‘AzureName’:
	topics.SetAzureName(genruntime.ExtractKubernetesResourceNameFromARMName(typedInput.Name))

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		topics.Location = &location
	}

	// Set property ‘Owner’:
	topics.Owner = genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		topics.Tags = make(map[string]string)
		for key, value := range typedInput.Tags {
			topics.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &DomainsTopics_Spec{}

// ConvertSpecFrom populates our DomainsTopics_Spec from the provided source
func (topics *DomainsTopics_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*v1alpha1api20200601storage.DomainsTopics_Spec)
	if ok {
		// Populate our instance from source
		return topics.AssignPropertiesFromDomainsTopicsSpec(src)
	}

	// Convert to an intermediate form
	src = &v1alpha1api20200601storage.DomainsTopics_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = topics.AssignPropertiesFromDomainsTopicsSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our DomainsTopics_Spec
func (topics *DomainsTopics_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*v1alpha1api20200601storage.DomainsTopics_Spec)
	if ok {
		// Populate destination from our instance
		return topics.AssignPropertiesToDomainsTopicsSpec(dst)
	}

	// Convert to an intermediate form
	dst = &v1alpha1api20200601storage.DomainsTopics_Spec{}
	err := topics.AssignPropertiesToDomainsTopicsSpec(dst)
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

// AssignPropertiesFromDomainsTopicsSpec populates our DomainsTopics_Spec from the provided source DomainsTopics_Spec
func (topics *DomainsTopics_Spec) AssignPropertiesFromDomainsTopicsSpec(source *v1alpha1api20200601storage.DomainsTopics_Spec) error {

	// AzureName
	topics.AzureName = source.AzureName

	// Location
	topics.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	topics.Owner = source.Owner.Copy()

	// Tags
	topics.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToDomainsTopicsSpec populates the provided destination DomainsTopics_Spec from our DomainsTopics_Spec
func (topics *DomainsTopics_Spec) AssignPropertiesToDomainsTopicsSpec(destination *v1alpha1api20200601storage.DomainsTopics_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// AzureName
	destination.AzureName = topics.AzureName

	// Location
	destination.Location = genruntime.ClonePointerToString(topics.Location)

	// OriginalVersion
	destination.OriginalVersion = topics.OriginalVersion()

	// Owner
	destination.Owner = topics.Owner.Copy()

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(topics.Tags)

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
func (topics *DomainsTopics_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// SetAzureName sets the Azure name of the resource
func (topics *DomainsTopics_Spec) SetAzureName(azureName string) { topics.AzureName = azureName }

func init() {
	SchemeBuilder.Register(&DomainsTopic{}, &DomainsTopicList{})
}
