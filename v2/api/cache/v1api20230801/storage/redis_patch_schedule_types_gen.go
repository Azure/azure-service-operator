// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redispatchschedules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redispatchschedules/status,redispatchschedules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230801.RedisPatchSchedule
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-08-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/default
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedule_Spec   `json:"spec,omitempty"`
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

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource (always "default")
func (schedule *RedisPatchSchedule) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-08-01"
func (schedule RedisPatchSchedule) GetAPIVersion() string {
	return "2023-08-01"
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

// GetSupportedOperations returns the operations supported by the resource
func (schedule *RedisPatchSchedule) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedule *RedisPatchSchedule) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// NewEmptyStatus returns a new empty (blank) status
func (schedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisPatchSchedule_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (schedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(schedule.Spec)
	return schedule.Spec.Owner.AsResourceReference(group, kind)
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

// Hub marks that this RedisPatchSchedule is the hub type for conversion
func (schedule *RedisPatchSchedule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (schedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: schedule.Spec.OriginalVersion,
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230801.RedisPatchSchedule
// Generator information:
// - Generated from: /redis/resource-manager/Microsoft.Cache/stable/2023-08-01/redis.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Cache/redis/{name}/patchSchedules/default
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

// Storage version of v1api20230801.RedisPatchSchedule_Spec
type RedisPatchSchedule_Spec struct {
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a cache.azure.com/Redis resource
	Owner           *genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner,omitempty" kind:"Redis"`
	PropertyBag     genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry                    `json:"scheduleEntries,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedule_Spec{}

// ConvertSpecFrom populates our RedisPatchSchedule_Spec from the provided source
func (schedule *RedisPatchSchedule_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == schedule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(schedule)
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedule_Spec
func (schedule *RedisPatchSchedule_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == schedule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(schedule)
}

// Storage version of v1api20230801.RedisPatchSchedule_STATUS
type RedisPatchSchedule_STATUS struct {
	Conditions      []conditions.Condition `json:"conditions,omitempty"`
	Id              *string                `json:"id,omitempty"`
	Location        *string                `json:"location,omitempty"`
	Name            *string                `json:"name,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry_STATUS `json:"scheduleEntries,omitempty"`
	Type            *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_STATUS{}

// ConvertStatusFrom populates our RedisPatchSchedule_STATUS from the provided source
func (schedule *RedisPatchSchedule_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == schedule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(schedule)
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_STATUS
func (schedule *RedisPatchSchedule_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == schedule {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(schedule)
}

// Storage version of v1api20230801.ScheduleEntry
// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntry struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

// Storage version of v1api20230801.ScheduleEntry_STATUS
// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntry_STATUS struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}
