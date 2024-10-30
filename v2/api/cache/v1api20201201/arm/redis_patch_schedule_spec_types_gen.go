// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisPatchSchedule_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: List of patch schedules for a Redis cache.
	Properties *ScheduleEntries `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisPatchSchedule_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedule RedisPatchSchedule_Spec) GetAPIVersion() string {
	return "2020-12-01"
}

// GetName returns the Name of the resource
func (schedule *RedisPatchSchedule_Spec) GetName() string {
	return schedule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedule *RedisPatchSchedule_Spec) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// List of patch schedules for a Redis cache.
type ScheduleEntries struct {
	// ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry `json:"scheduleEntries,omitempty"`
}

// Patch schedule entry for a Premium Redis Cache.
type ScheduleEntry struct {
	// DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek *ScheduleEntry_DayOfWeek `json:"dayOfWeek,omitempty"`

	// MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc *int `json:"startHourUtc,omitempty"`
}

// +kubebuilder:validation:Enum={"Everyday","Friday","Monday","Saturday","Sunday","Thursday","Tuesday","Wednesday","Weekend"}
type ScheduleEntry_DayOfWeek string

const (
	ScheduleEntry_DayOfWeek_Everyday  = ScheduleEntry_DayOfWeek("Everyday")
	ScheduleEntry_DayOfWeek_Friday    = ScheduleEntry_DayOfWeek("Friday")
	ScheduleEntry_DayOfWeek_Monday    = ScheduleEntry_DayOfWeek("Monday")
	ScheduleEntry_DayOfWeek_Saturday  = ScheduleEntry_DayOfWeek("Saturday")
	ScheduleEntry_DayOfWeek_Sunday    = ScheduleEntry_DayOfWeek("Sunday")
	ScheduleEntry_DayOfWeek_Thursday  = ScheduleEntry_DayOfWeek("Thursday")
	ScheduleEntry_DayOfWeek_Tuesday   = ScheduleEntry_DayOfWeek("Tuesday")
	ScheduleEntry_DayOfWeek_Wednesday = ScheduleEntry_DayOfWeek("Wednesday")
	ScheduleEntry_DayOfWeek_Weekend   = ScheduleEntry_DayOfWeek("Weekend")
)

// Mapping from string to ScheduleEntry_DayOfWeek
var scheduleEntry_DayOfWeek_Values = map[string]ScheduleEntry_DayOfWeek{
	"everyday":  ScheduleEntry_DayOfWeek_Everyday,
	"friday":    ScheduleEntry_DayOfWeek_Friday,
	"monday":    ScheduleEntry_DayOfWeek_Monday,
	"saturday":  ScheduleEntry_DayOfWeek_Saturday,
	"sunday":    ScheduleEntry_DayOfWeek_Sunday,
	"thursday":  ScheduleEntry_DayOfWeek_Thursday,
	"tuesday":   ScheduleEntry_DayOfWeek_Tuesday,
	"wednesday": ScheduleEntry_DayOfWeek_Wednesday,
	"weekend":   ScheduleEntry_DayOfWeek_Weekend,
}
