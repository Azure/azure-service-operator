// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20211101

type Servers_Databases_BackupShortTermRetentionPolicy_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`

	// Name: Resource name.
	Name *string `json:"name,omitempty"`

	// Properties: Resource properties.
	Properties *BackupShortTermRetentionPolicyProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: Resource type.
	Type *string `json:"type,omitempty"`
}

// Properties of a short term retention policy
type BackupShortTermRetentionPolicyProperties_STATUS_ARM struct {
	// DiffBackupIntervalInHours: The differential backup interval in hours. This is how many interval hours between each
	// differential backup will be supported. This is only applicable to live databases but not dropped databases.
	DiffBackupIntervalInHours *BackupShortTermRetentionPolicyProperties_DiffBackupIntervalInHours_STATUS `json:"diffBackupIntervalInHours,omitempty"`

	// RetentionDays: The backup retention period in days. This is how many days Point-in-Time Restore will be supported.
	RetentionDays *int `json:"retentionDays,omitempty"`
}
