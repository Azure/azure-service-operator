// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201201

type RedisPatchSchedule_STATUSARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: List of patch schedules for a Redis cache.
	Properties *ScheduleEntries_STATUSARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ScheduleEntries_STATUSARM struct {
	// ScheduleEntries: List of patch schedules for a Redis cache.
	ScheduleEntries []ScheduleEntry_STATUSARM `json:"scheduleEntries,omitempty"`
}

type ScheduleEntry_STATUSARM struct {
	// DayOfWeek: Day of the week when a cache can be patched.
	DayOfWeek *ScheduleEntry_STATUS_DayOfWeek `json:"dayOfWeek,omitempty"`

	// MaintenanceWindow: ISO8601 timespan specifying how much time cache patching can take.
	MaintenanceWindow *string `json:"maintenanceWindow,omitempty"`

	// StartHourUtc: Start hour after which cache patching can start.
	StartHourUtc *int `json:"startHourUtc,omitempty"`
}
