// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

// Deprecated version of RedisPatchSchedules_Spec. Use v1beta20201201.RedisPatchSchedules_Spec instead
type RedisPatchSchedules_SpecARM struct {
	Location   *string             `json:"location,omitempty"`
	Name       string              `json:"name,omitempty"`
	Properties *ScheduleEntriesARM `json:"properties,omitempty"`
	Tags       map[string]string   `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisPatchSchedules_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (schedules RedisPatchSchedules_SpecARM) GetAPIVersion() string {
	return string(APIVersionValue)
}

// GetName returns the Name of the resource
func (schedules *RedisPatchSchedules_SpecARM) GetName() string {
	return schedules.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (schedules *RedisPatchSchedules_SpecARM) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// Deprecated version of ScheduleEntries. Use v1beta20201201.ScheduleEntries instead
type ScheduleEntriesARM struct {
	ScheduleEntries []ScheduleEntryARM `json:"scheduleEntries,omitempty"`
}

// Deprecated version of ScheduleEntry. Use v1beta20201201.ScheduleEntry instead
type ScheduleEntryARM struct {
	DayOfWeek         *ScheduleEntryDayOfWeek `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                 `json:"maintenanceWindow,omitempty"`
	StartHourUtc      *int                    `json:"startHourUtc,omitempty"`
}
