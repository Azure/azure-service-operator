// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ServersDatabasesBackupLongTermRetentionPolicy_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *BaseLongTermRetentionPolicyProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ServersDatabasesBackupLongTermRetentionPolicy_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesBackupLongTermRetentionPolicy_Spec) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (policy *ServersDatabasesBackupLongTermRetentionPolicy_Spec) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
func (policy *ServersDatabasesBackupLongTermRetentionPolicy_Spec) GetType() string {
	return "Microsoft.Sql/servers/databases/backupLongTermRetentionPolicies"
}

// Properties of a long term retention policy
type BaseLongTermRetentionPolicyProperties struct {
	// MonthlyRetention: The monthly retention policy for an LTR backup in an ISO 8601 format.
	MonthlyRetention *string `json:"monthlyRetention,omitempty"`

	// WeekOfYear: The week of year to take the yearly backup in an ISO 8601 format.
	WeekOfYear *int `json:"weekOfYear,omitempty"`

	// WeeklyRetention: The weekly retention policy for an LTR backup in an ISO 8601 format.
	WeeklyRetention *string `json:"weeklyRetention,omitempty"`

	// YearlyRetention: The yearly retention policy for an LTR backup in an ISO 8601 format.
	YearlyRetention *string `json:"yearlyRetention,omitempty"`
}