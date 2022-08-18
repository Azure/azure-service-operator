// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import (
	"fmt"
	alpha20201201s "github.com/Azure/azure-service-operator/v2/api/cache/v1alpha1api20201201storage"
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
// Deprecated version of RedisPatchSchedule. Use v1beta20201201.RedisPatchSchedule instead
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedules_Spec  `json:"spec,omitempty"`
	Status            RedisPatchSchedule_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisPatchSchedule{}

// GetConditions returns the conditions of the resource
func (schedule *RedisPatchSchedule) GetConditions() conditions.Conditions {
	return schedule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (schedule *RedisPatchSchedule) SetConditions(conditions conditions.Conditions) {
	schedule.Status.Conditions = conditions
}

var _ conversion.Convertible = &RedisPatchSchedule{}

// ConvertFrom populates our RedisPatchSchedule from the provided hub RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertFrom(hub conversion.Hub) error {
	// intermediate variable for conversion
	var source alpha20201201s.RedisPatchSchedule

	err := source.ConvertFrom(hub)
	if err != nil {
		return errors.Wrap(err, "converting from hub to source")
	}

	err = schedule.AssignPropertiesFromRedisPatchSchedule(&source)
	if err != nil {
		return errors.Wrap(err, "converting from source to schedule")
	}

	return nil
}

// ConvertTo populates the provided hub RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) ConvertTo(hub conversion.Hub) error {
	// intermediate variable for conversion
	var destination alpha20201201s.RedisPatchSchedule
	err := schedule.AssignPropertiesToRedisPatchSchedule(&destination)
	if err != nil {
		return errors.Wrap(err, "converting to destination from schedule")
	}
	err = destination.ConvertTo(hub)
	if err != nil {
		return errors.Wrap(err, "converting from destination to hub")
	}

	return nil
}

// +kubebuilder:webhook:path=/mutate-cache-azure-com-v1alpha1api20201201-redispatchschedule,mutating=true,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1alpha1api20201201,name=default.v1alpha1api20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Defaulter = &RedisPatchSchedule{}

// Default applies defaults to the RedisPatchSchedule resource
func (schedule *RedisPatchSchedule) Default() {
	schedule.defaultImpl()
	var temp interface{} = schedule
	if runtimeDefaulter, ok := temp.(genruntime.Defaulter); ok {
		runtimeDefaulter.CustomDefault()
	}
}

// defaultImpl applies the code generated defaults to the RedisPatchSchedule resource
func (schedule *RedisPatchSchedule) defaultImpl() {}

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource (always "default")
func (schedule *RedisPatchSchedule) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (schedule *RedisPatchSchedule) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (schedule *RedisPatchSchedule) GetSpec() genruntime.ConvertibleSpec {
	return &schedule.Spec
}

// GetStatus returns the status of this resource
func (schedule *RedisPatchSchedule) GetStatus() genruntime.ConvertibleStatus {
	return &schedule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedule *RedisPatchSchedule) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// NewEmptyStatus returns a new empty (blank) status
func (schedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisPatchSchedule_STATUS{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (schedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(schedule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  schedule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (schedule *RedisPatchSchedule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisPatchSchedule_STATUS); ok {
		schedule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisPatchSchedule_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	schedule.Status = st
	return nil
}

// +kubebuilder:webhook:path=/validate-cache-azure-com-v1alpha1api20201201-redispatchschedule,mutating=false,sideEffects=None,matchPolicy=Exact,failurePolicy=fail,groups=cache.azure.com,resources=redispatchschedules,verbs=create;update,versions=v1alpha1api20201201,name=validate.v1alpha1api20201201.redispatchschedules.cache.azure.com,admissionReviewVersions=v1

var _ admission.Validator = &RedisPatchSchedule{}

// ValidateCreate validates the creation of the resource
func (schedule *RedisPatchSchedule) ValidateCreate() error {
	validations := schedule.createValidations()
	var temp interface{} = schedule
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
func (schedule *RedisPatchSchedule) ValidateDelete() error {
	validations := schedule.deleteValidations()
	var temp interface{} = schedule
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
func (schedule *RedisPatchSchedule) ValidateUpdate(old runtime.Object) error {
	validations := schedule.updateValidations()
	var temp interface{} = schedule
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
func (schedule *RedisPatchSchedule) createValidations() []func() error {
	return []func() error{schedule.validateResourceReferences}
}

// deleteValidations validates the deletion of the resource
func (schedule *RedisPatchSchedule) deleteValidations() []func() error {
	return nil
}

// updateValidations validates the update of the resource
func (schedule *RedisPatchSchedule) updateValidations() []func(old runtime.Object) error {
	return []func(old runtime.Object) error{
		func(old runtime.Object) error {
			return schedule.validateResourceReferences()
		},
		schedule.validateWriteOnceProperties}
}

// validateResourceReferences validates all resource references
func (schedule *RedisPatchSchedule) validateResourceReferences() error {
	refs, err := reflecthelpers.FindResourceReferences(&schedule.Spec)
	if err != nil {
		return err
	}
	return genruntime.ValidateResourceReferences(refs)
}

// validateWriteOnceProperties validates all WriteOnce properties
func (schedule *RedisPatchSchedule) validateWriteOnceProperties(old runtime.Object) error {
	oldObj, ok := old.(*RedisPatchSchedule)
	if !ok {
		return nil
	}

	return genruntime.ValidateWriteOnceProperties(oldObj, schedule)
}

// AssignPropertiesFromRedisPatchSchedule populates our RedisPatchSchedule from the provided source RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesFromRedisPatchSchedule(source *alpha20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	schedule.ObjectMeta = *source.ObjectMeta.DeepCopy()

	// Spec
	var spec RedisPatchSchedules_Spec
	err := spec.AssignPropertiesFromRedisPatchSchedulesSpec(&source.Spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchSchedulesSpec() to populate field Spec")
	}
	schedule.Spec = spec

	// Status
	var status RedisPatchSchedule_STATUS
	err = status.AssignPropertiesFromRedisPatchScheduleSTATUS(&source.Status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesFromRedisPatchScheduleSTATUS() to populate field Status")
	}
	schedule.Status = status

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedule populates the provided destination RedisPatchSchedule from our RedisPatchSchedule
func (schedule *RedisPatchSchedule) AssignPropertiesToRedisPatchSchedule(destination *alpha20201201s.RedisPatchSchedule) error {

	// ObjectMeta
	destination.ObjectMeta = *schedule.ObjectMeta.DeepCopy()

	// Spec
	var spec alpha20201201s.RedisPatchSchedules_Spec
	err := schedule.Spec.AssignPropertiesToRedisPatchSchedulesSpec(&spec)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchSchedulesSpec() to populate field Spec")
	}
	destination.Spec = spec

	// Status
	var status alpha20201201s.RedisPatchSchedule_STATUS
	err = schedule.Status.AssignPropertiesToRedisPatchScheduleSTATUS(&status)
	if err != nil {
		return errors.Wrap(err, "calling AssignPropertiesToRedisPatchScheduleSTATUS() to populate field Status")
	}
	destination.Status = status

	// No error
	return nil
}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (schedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: schedule.Spec.OriginalVersion(),
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
// Deprecated version of RedisPatchSchedule. Use v1beta20201201.RedisPatchSchedule instead
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

// Deprecated version of RedisPatchSchedule_STATUS. Use v1beta20201201.RedisPatchSchedule_STATUS instead
type RedisPatchSchedule_STATUS struct {
	// Conditions: The observed state of the resource
	Conditions      []conditions.Condition `json:"conditions,omitempty"`
	Id              *string                `json:"id,omitempty"`
	Location        *string                `json:"location,omitempty"`
	Name            *string                `json:"name,omitempty"`
	ScheduleEntries []ScheduleEntry_STATUS `json:"scheduleEntries,omitempty"`
	Type            *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_STATUS{}

// ConvertStatusFrom populates our RedisPatchSchedule_STATUS from the provided source
func (schedule *RedisPatchSchedule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	src, ok := source.(*alpha20201201s.RedisPatchSchedule_STATUS)
	if ok {
		// Populate our instance from source
		return schedule.AssignPropertiesFromRedisPatchScheduleSTATUS(src)
	}

	// Convert to an intermediate form
	src = &alpha20201201s.RedisPatchSchedule_STATUS{}
	err := src.ConvertStatusFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertStatusFrom()")
	}

	// Update our instance from src
	err = schedule.AssignPropertiesFromRedisPatchScheduleSTATUS(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertStatusFrom()")
	}

	return nil
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_STATUS
func (schedule *RedisPatchSchedule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	dst, ok := destination.(*alpha20201201s.RedisPatchSchedule_STATUS)
	if ok {
		// Populate destination from our instance
		return schedule.AssignPropertiesToRedisPatchScheduleSTATUS(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20201201s.RedisPatchSchedule_STATUS{}
	err := schedule.AssignPropertiesToRedisPatchScheduleSTATUS(dst)
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

var _ genruntime.FromARMConverter = &RedisPatchSchedule_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (schedule *RedisPatchSchedule_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedule_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (schedule *RedisPatchSchedule_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedule_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedule_STATUSARM, got %T", armInput)
	}

	// no assignment for property ‘Conditions’

	// Set property ‘Id’:
	if typedInput.Id != nil {
		id := *typedInput.Id
		schedule.Id = &id
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		schedule.Location = &location
	}

	// Set property ‘Name’:
	if typedInput.Name != nil {
		name := *typedInput.Name
		schedule.Name = &name
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.ScheduleEntries {
			var item1 ScheduleEntry_STATUS
			err := item1.PopulateFromARM(owner, item)
			if err != nil {
				return err
			}
			schedule.ScheduleEntries = append(schedule.ScheduleEntries, item1)
		}
	}

	// Set property ‘Type’:
	if typedInput.Type != nil {
		typeVar := *typedInput.Type
		schedule.Type = &typeVar
	}

	// No error
	return nil
}

// AssignPropertiesFromRedisPatchScheduleSTATUS populates our RedisPatchSchedule_STATUS from the provided source RedisPatchSchedule_STATUS
func (schedule *RedisPatchSchedule_STATUS) AssignPropertiesFromRedisPatchScheduleSTATUS(source *alpha20201201s.RedisPatchSchedule_STATUS) error {

	// Conditions
	schedule.Conditions = genruntime.CloneSliceOfCondition(source.Conditions)

	// Id
	schedule.Id = genruntime.ClonePointerToString(source.Id)

	// Location
	schedule.Location = genruntime.ClonePointerToString(source.Location)

	// Name
	schedule.Name = genruntime.ClonePointerToString(source.Name)

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry_STATUS, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry_STATUS
			err := scheduleEntry.AssignPropertiesFromScheduleEntrySTATUS(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntrySTATUS() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedule.ScheduleEntries = scheduleEntryList
	} else {
		schedule.ScheduleEntries = nil
	}

	// Type
	schedule.Type = genruntime.ClonePointerToString(source.Type)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchScheduleSTATUS populates the provided destination RedisPatchSchedule_STATUS from our RedisPatchSchedule_STATUS
func (schedule *RedisPatchSchedule_STATUS) AssignPropertiesToRedisPatchScheduleSTATUS(destination *alpha20201201s.RedisPatchSchedule_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Conditions
	destination.Conditions = genruntime.CloneSliceOfCondition(schedule.Conditions)

	// Id
	destination.Id = genruntime.ClonePointerToString(schedule.Id)

	// Location
	destination.Location = genruntime.ClonePointerToString(schedule.Location)

	// Name
	destination.Name = genruntime.ClonePointerToString(schedule.Name)

	// ScheduleEntries
	if schedule.ScheduleEntries != nil {
		scheduleEntryList := make([]alpha20201201s.ScheduleEntry_STATUS, len(schedule.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedule.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry alpha20201201s.ScheduleEntry_STATUS
			err := scheduleEntryItem.AssignPropertiesToScheduleEntrySTATUS(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntrySTATUS() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Type
	destination.Type = genruntime.ClonePointerToString(schedule.Type)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

type RedisPatchSchedules_Spec struct {
	Location *string `json:"location,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`

	// +kubebuilder:validation:Required
	ScheduleEntries []ScheduleEntry   `json:"scheduleEntries,omitempty"`
	Tags            map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMTransformer = &RedisPatchSchedules_Spec{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (schedules *RedisPatchSchedules_Spec) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if schedules == nil {
		return nil, nil
	}
	result := &RedisPatchSchedules_SpecARM{}

	// Set property ‘Location’:
	if schedules.Location != nil {
		location := *schedules.Location
		result.Location = &location
	}

	// Set property ‘Name’:
	result.Name = resolved.Name

	// Set property ‘Properties’:
	if schedules.ScheduleEntries != nil {
		result.Properties = &ScheduleEntriesARM{}
	}
	for _, item := range schedules.ScheduleEntries {
		itemARM, err := item.ConvertToARM(resolved)
		if err != nil {
			return nil, err
		}
		result.Properties.ScheduleEntries = append(result.Properties.ScheduleEntries, *itemARM.(*ScheduleEntryARM))
	}

	// Set property ‘Tags’:
	if schedules.Tags != nil {
		result.Tags = make(map[string]string, len(schedules.Tags))
		for key, value := range schedules.Tags {
			result.Tags[key] = value
		}
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (schedules *RedisPatchSchedules_Spec) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &RedisPatchSchedules_SpecARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (schedules *RedisPatchSchedules_Spec) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(RedisPatchSchedules_SpecARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected RedisPatchSchedules_SpecARM, got %T", armInput)
	}

	// Set property ‘Location’:
	if typedInput.Location != nil {
		location := *typedInput.Location
		schedules.Location = &location
	}

	// Set property ‘Owner’:
	schedules.Owner = &genruntime.KnownResourceReference{
		Name: owner.Name,
	}

	// Set property ‘ScheduleEntries’:
	// copying flattened property:
	if typedInput.Properties != nil {
		for _, item := range typedInput.Properties.ScheduleEntries {
			var item1 ScheduleEntry
			err := item1.PopulateFromARM(owner, item)
			if err != nil {
				return err
			}
			schedules.ScheduleEntries = append(schedules.ScheduleEntries, item1)
		}
	}

	// Set property ‘Tags’:
	if typedInput.Tags != nil {
		schedules.Tags = make(map[string]string, len(typedInput.Tags))
		for key, value := range typedInput.Tags {
			schedules.Tags[key] = value
		}
	}

	// No error
	return nil
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedules_Spec{}

// ConvertSpecFrom populates our RedisPatchSchedules_Spec from the provided source
func (schedules *RedisPatchSchedules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	src, ok := source.(*alpha20201201s.RedisPatchSchedules_Spec)
	if ok {
		// Populate our instance from source
		return schedules.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	}

	// Convert to an intermediate form
	src = &alpha20201201s.RedisPatchSchedules_Spec{}
	err := src.ConvertSpecFrom(source)
	if err != nil {
		return errors.Wrap(err, "initial step of conversion in ConvertSpecFrom()")
	}

	// Update our instance from src
	err = schedules.AssignPropertiesFromRedisPatchSchedulesSpec(src)
	if err != nil {
		return errors.Wrap(err, "final step of conversion in ConvertSpecFrom()")
	}

	return nil
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	dst, ok := destination.(*alpha20201201s.RedisPatchSchedules_Spec)
	if ok {
		// Populate destination from our instance
		return schedules.AssignPropertiesToRedisPatchSchedulesSpec(dst)
	}

	// Convert to an intermediate form
	dst = &alpha20201201s.RedisPatchSchedules_Spec{}
	err := schedules.AssignPropertiesToRedisPatchSchedulesSpec(dst)
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

// AssignPropertiesFromRedisPatchSchedulesSpec populates our RedisPatchSchedules_Spec from the provided source RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) AssignPropertiesFromRedisPatchSchedulesSpec(source *alpha20201201s.RedisPatchSchedules_Spec) error {

	// Location
	schedules.Location = genruntime.ClonePointerToString(source.Location)

	// Owner
	if source.Owner != nil {
		owner := source.Owner.Copy()
		schedules.Owner = &owner
	} else {
		schedules.Owner = nil
	}

	// ScheduleEntries
	if source.ScheduleEntries != nil {
		scheduleEntryList := make([]ScheduleEntry, len(source.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range source.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry ScheduleEntry
			err := scheduleEntry.AssignPropertiesFromScheduleEntry(&scheduleEntryItem)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesFromScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		schedules.ScheduleEntries = scheduleEntryList
	} else {
		schedules.ScheduleEntries = nil
	}

	// Tags
	schedules.Tags = genruntime.CloneMapOfStringToString(source.Tags)

	// No error
	return nil
}

// AssignPropertiesToRedisPatchSchedulesSpec populates the provided destination RedisPatchSchedules_Spec from our RedisPatchSchedules_Spec
func (schedules *RedisPatchSchedules_Spec) AssignPropertiesToRedisPatchSchedulesSpec(destination *alpha20201201s.RedisPatchSchedules_Spec) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// Location
	destination.Location = genruntime.ClonePointerToString(schedules.Location)

	// OriginalVersion
	destination.OriginalVersion = schedules.OriginalVersion()

	// Owner
	if schedules.Owner != nil {
		owner := schedules.Owner.Copy()
		destination.Owner = &owner
	} else {
		destination.Owner = nil
	}

	// ScheduleEntries
	if schedules.ScheduleEntries != nil {
		scheduleEntryList := make([]alpha20201201s.ScheduleEntry, len(schedules.ScheduleEntries))
		for scheduleEntryIndex, scheduleEntryItem := range schedules.ScheduleEntries {
			// Shadow the loop variable to avoid aliasing
			scheduleEntryItem := scheduleEntryItem
			var scheduleEntry alpha20201201s.ScheduleEntry
			err := scheduleEntryItem.AssignPropertiesToScheduleEntry(&scheduleEntry)
			if err != nil {
				return errors.Wrap(err, "calling AssignPropertiesToScheduleEntry() to populate field ScheduleEntries")
			}
			scheduleEntryList[scheduleEntryIndex] = scheduleEntry
		}
		destination.ScheduleEntries = scheduleEntryList
	} else {
		destination.ScheduleEntries = nil
	}

	// Tags
	destination.Tags = genruntime.CloneMapOfStringToString(schedules.Tags)

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
func (schedules *RedisPatchSchedules_Spec) OriginalVersion() string {
	return GroupVersion.Version
}

// Deprecated version of ScheduleEntry. Use v1beta20201201.ScheduleEntry instead
type ScheduleEntry struct {
	// +kubebuilder:validation:Required
	DayOfWeek         *ScheduleEntryDayOfWeek `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                 `json:"maintenanceWindow,omitempty"`

	// +kubebuilder:validation:Required
	StartHourUtc *int `json:"startHourUtc,omitempty"`
}

var _ genruntime.ARMTransformer = &ScheduleEntry{}

// ConvertToARM converts from a Kubernetes CRD object to an ARM object
func (entry *ScheduleEntry) ConvertToARM(resolved genruntime.ConvertToARMResolvedDetails) (interface{}, error) {
	if entry == nil {
		return nil, nil
	}
	result := &ScheduleEntryARM{}

	// Set property ‘DayOfWeek’:
	if entry.DayOfWeek != nil {
		dayOfWeek := *entry.DayOfWeek
		result.DayOfWeek = &dayOfWeek
	}

	// Set property ‘MaintenanceWindow’:
	if entry.MaintenanceWindow != nil {
		maintenanceWindow := *entry.MaintenanceWindow
		result.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	if entry.StartHourUtc != nil {
		startHourUtc := *entry.StartHourUtc
		result.StartHourUtc = &startHourUtc
	}
	return result, nil
}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (entry *ScheduleEntry) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntryARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (entry *ScheduleEntry) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntryARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntryARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	if typedInput.DayOfWeek != nil {
		dayOfWeek := *typedInput.DayOfWeek
		entry.DayOfWeek = &dayOfWeek
	}

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	if typedInput.StartHourUtc != nil {
		startHourUtc := *typedInput.StartHourUtc
		entry.StartHourUtc = &startHourUtc
	}

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntry populates our ScheduleEntry from the provided source ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesFromScheduleEntry(source *alpha20201201s.ScheduleEntry) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		dayOfWeek := ScheduleEntryDayOfWeek(*source.DayOfWeek)
		entry.DayOfWeek = &dayOfWeek
	} else {
		entry.DayOfWeek = nil
	}

	// MaintenanceWindow
	if source.MaintenanceWindow != nil {
		maintenanceWindow := *source.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	} else {
		entry.MaintenanceWindow = nil
	}

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntry populates the provided destination ScheduleEntry from our ScheduleEntry
func (entry *ScheduleEntry) AssignPropertiesToScheduleEntry(destination *alpha20201201s.ScheduleEntry) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	if entry.DayOfWeek != nil {
		dayOfWeek := string(*entry.DayOfWeek)
		destination.DayOfWeek = &dayOfWeek
	} else {
		destination.DayOfWeek = nil
	}

	// MaintenanceWindow
	if entry.MaintenanceWindow != nil {
		maintenanceWindow := *entry.MaintenanceWindow
		destination.MaintenanceWindow = &maintenanceWindow
	} else {
		destination.MaintenanceWindow = nil
	}

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of ScheduleEntry_STATUS. Use v1beta20201201.ScheduleEntry_STATUS instead
type ScheduleEntry_STATUS struct {
	DayOfWeek         *ScheduleEntrySTATUSDayOfWeek `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                       `json:"maintenanceWindow,omitempty"`
	StartHourUtc      *int                          `json:"startHourUtc,omitempty"`
}

var _ genruntime.FromARMConverter = &ScheduleEntry_STATUS{}

// NewEmptyARMValue returns an empty ARM value suitable for deserializing into
func (entry *ScheduleEntry_STATUS) NewEmptyARMValue() genruntime.ARMResourceStatus {
	return &ScheduleEntry_STATUSARM{}
}

// PopulateFromARM populates a Kubernetes CRD object from an Azure ARM object
func (entry *ScheduleEntry_STATUS) PopulateFromARM(owner genruntime.ArbitraryOwnerReference, armInput interface{}) error {
	typedInput, ok := armInput.(ScheduleEntry_STATUSARM)
	if !ok {
		return fmt.Errorf("unexpected type supplied for PopulateFromARM() function. Expected ScheduleEntry_STATUSARM, got %T", armInput)
	}

	// Set property ‘DayOfWeek’:
	if typedInput.DayOfWeek != nil {
		dayOfWeek := *typedInput.DayOfWeek
		entry.DayOfWeek = &dayOfWeek
	}

	// Set property ‘MaintenanceWindow’:
	if typedInput.MaintenanceWindow != nil {
		maintenanceWindow := *typedInput.MaintenanceWindow
		entry.MaintenanceWindow = &maintenanceWindow
	}

	// Set property ‘StartHourUtc’:
	if typedInput.StartHourUtc != nil {
		startHourUtc := *typedInput.StartHourUtc
		entry.StartHourUtc = &startHourUtc
	}

	// No error
	return nil
}

// AssignPropertiesFromScheduleEntrySTATUS populates our ScheduleEntry_STATUS from the provided source ScheduleEntry_STATUS
func (entry *ScheduleEntry_STATUS) AssignPropertiesFromScheduleEntrySTATUS(source *alpha20201201s.ScheduleEntry_STATUS) error {

	// DayOfWeek
	if source.DayOfWeek != nil {
		dayOfWeek := ScheduleEntrySTATUSDayOfWeek(*source.DayOfWeek)
		entry.DayOfWeek = &dayOfWeek
	} else {
		entry.DayOfWeek = nil
	}

	// MaintenanceWindow
	entry.MaintenanceWindow = genruntime.ClonePointerToString(source.MaintenanceWindow)

	// StartHourUtc
	entry.StartHourUtc = genruntime.ClonePointerToInt(source.StartHourUtc)

	// No error
	return nil
}

// AssignPropertiesToScheduleEntrySTATUS populates the provided destination ScheduleEntry_STATUS from our ScheduleEntry_STATUS
func (entry *ScheduleEntry_STATUS) AssignPropertiesToScheduleEntrySTATUS(destination *alpha20201201s.ScheduleEntry_STATUS) error {
	// Create a new property bag
	propertyBag := genruntime.NewPropertyBag()

	// DayOfWeek
	if entry.DayOfWeek != nil {
		dayOfWeek := string(*entry.DayOfWeek)
		destination.DayOfWeek = &dayOfWeek
	} else {
		destination.DayOfWeek = nil
	}

	// MaintenanceWindow
	destination.MaintenanceWindow = genruntime.ClonePointerToString(entry.MaintenanceWindow)

	// StartHourUtc
	destination.StartHourUtc = genruntime.ClonePointerToInt(entry.StartHourUtc)

	// Update the property bag
	if len(propertyBag) > 0 {
		destination.PropertyBag = propertyBag
	} else {
		destination.PropertyBag = nil
	}

	// No error
	return nil
}

// Deprecated version of ScheduleEntryDayOfWeek. Use v1beta20201201.ScheduleEntryDayOfWeek instead
// +kubebuilder:validation:Enum={"Everyday","Friday","Monday","Saturday","Sunday","Thursday","Tuesday","Wednesday","Weekend"}
type ScheduleEntryDayOfWeek string

const (
	ScheduleEntryDayOfWeek_Everyday  = ScheduleEntryDayOfWeek("Everyday")
	ScheduleEntryDayOfWeek_Friday    = ScheduleEntryDayOfWeek("Friday")
	ScheduleEntryDayOfWeek_Monday    = ScheduleEntryDayOfWeek("Monday")
	ScheduleEntryDayOfWeek_Saturday  = ScheduleEntryDayOfWeek("Saturday")
	ScheduleEntryDayOfWeek_Sunday    = ScheduleEntryDayOfWeek("Sunday")
	ScheduleEntryDayOfWeek_Thursday  = ScheduleEntryDayOfWeek("Thursday")
	ScheduleEntryDayOfWeek_Tuesday   = ScheduleEntryDayOfWeek("Tuesday")
	ScheduleEntryDayOfWeek_Wednesday = ScheduleEntryDayOfWeek("Wednesday")
	ScheduleEntryDayOfWeek_Weekend   = ScheduleEntryDayOfWeek("Weekend")
)

// Deprecated version of ScheduleEntrySTATUSDayOfWeek. Use v1beta20201201.ScheduleEntrySTATUSDayOfWeek instead
type ScheduleEntrySTATUSDayOfWeek string

const (
	ScheduleEntrySTATUSDayOfWeek_Everyday  = ScheduleEntrySTATUSDayOfWeek("Everyday")
	ScheduleEntrySTATUSDayOfWeek_Friday    = ScheduleEntrySTATUSDayOfWeek("Friday")
	ScheduleEntrySTATUSDayOfWeek_Monday    = ScheduleEntrySTATUSDayOfWeek("Monday")
	ScheduleEntrySTATUSDayOfWeek_Saturday  = ScheduleEntrySTATUSDayOfWeek("Saturday")
	ScheduleEntrySTATUSDayOfWeek_Sunday    = ScheduleEntrySTATUSDayOfWeek("Sunday")
	ScheduleEntrySTATUSDayOfWeek_Thursday  = ScheduleEntrySTATUSDayOfWeek("Thursday")
	ScheduleEntrySTATUSDayOfWeek_Tuesday   = ScheduleEntrySTATUSDayOfWeek("Tuesday")
	ScheduleEntrySTATUSDayOfWeek_Wednesday = ScheduleEntrySTATUSDayOfWeek("Wednesday")
	ScheduleEntrySTATUSDayOfWeek_Weekend   = ScheduleEntrySTATUSDayOfWeek("Weekend")
)

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}
