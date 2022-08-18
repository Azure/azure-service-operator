// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisPatchSchedules_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Default string modeled as parameter for auto generation to work correctly.
	Name string `json:"name,omitempty"`

	// Properties: List of patch schedules for a Redis cache.
	Properties *ScheduleEntriesARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisPatchSchedules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedules RedisPatchSchedules_SpecARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (schedules *RedisPatchSchedules_SpecARM) GetName() string {
	return schedules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedules *RedisPatchSchedules_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntries
type ScheduleEntriesARM struct {
	// ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntryARM `json:"scheduleEntries,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry
type ScheduleEntryARM struct {
	// DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek *ScheduleEntryDayOfWeek `json:"dayOfWeek,omitempty"`

	// MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc *int `json:"startHourUtc,omitempty"`
}
