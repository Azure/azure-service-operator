// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

type ServersDatabasesBackupShortTermRetentionPolicy_STATUS struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *BackupShortTermRetentionPolicyProperties_STATUS `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of a short term retention policy
type BackupShortTermRetentionPolicyProperties_STATUS struct {
	// DiffBackupIntervalInHours: The differential backup interval in hours. This is how many interval hours between each
	// differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS `json:"diffBackupIntervalInHours,omitempty"`

	// RetentionDays: The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `json:"retentionDays,omitempty"`
}

type BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS int

const (
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS_12 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS(12)
	BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS_24 = BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS(24)
)
