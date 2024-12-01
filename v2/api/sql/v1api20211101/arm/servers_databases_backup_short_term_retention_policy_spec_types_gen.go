// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type ServersDatabasesBackupShortTermRetentionPolicy_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *BackupShortTermRetentionPolicyProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &ServersDatabasesBackupShortTermRetentionPolicy_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (policy ServersDatabasesBackupShortTermRetentionPolicy_Spec) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (policy *ServersDatabasesBackupShortTermRetentionPolicy_Spec) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"
func (policy *ServersDatabasesBackupShortTermRetentionPolicy_Spec) GetType() string {
	return "Microsoft.Sql/servers/databases/backupShortTermRetentionPolicies"
}

// Properties of a short term retention policy
type BackupShortTermRetentionPolicyProperties struct {
	// DiffBackupIntervalInHours: The differential backup interval in hours. This is how many interval hours between each
	// differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours `json:"diffBackupIntervalInHours,omitempty"`

	// RetentionDays: The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `json:"retentionDays,omitempty"`
}

// +kubebuilder:validation:Enum={12,24}
type BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours int

const (
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_12 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours(12)
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_24 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours(24)
)
