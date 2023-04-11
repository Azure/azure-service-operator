// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *BaseLongTermRetentionPolicyProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
func (policy *Servers_Databases_BackupLongTermRetentionPolicy_Spec_ARM) GetType() string {
	return "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
}

// Properties of a long term retention policy
type BaseLongTermRetentionPolicyProperties_ARM struct {
	// MonthlyRetention: The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `json:"monthlyRetention,omitempty"`

	// WeekOfYear: The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `json:"weekOfYear,omitempty"`

	// WeeklyRetention: The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `json:"weeklyRetention,omitempty"`

	// YearlyRetention: The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `json:"yearlyRetention,omitempty"`
}
