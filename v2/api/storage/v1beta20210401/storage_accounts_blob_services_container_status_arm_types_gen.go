// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

// Deprecated version of StorageAccounts_BlobServices_Container_STATUS. Use v1api20210401.StorageAccounts_BlobServices_Container_STATUS instead
type StorageAccounts_BlobServices_Container_STATUS_ARM struct {
	Etag       *string                         `json:"etag"`
	Id         *string                         `json:"id"`
	Name       *string                         `json:"name"`
	Properties *ContainerProperties_STATUS_ARM `json:"properties"`
	Type       *string                         `json:"type"`
}

// Deprecated version of ContainerProperties_STATUS. Use v1api20210401.ContainerProperties_STATUS instead
type ContainerProperties_STATUS_ARM struct {
	DefaultEncryptionScope         *string                                    `json:"defaultEncryptionScope"`
	Deleted                        *bool                                      `json:"deleted"`
	DeletedTime                    *string                                    `json:"deletedTime"`
	DenyEncryptionScopeOverride    *bool                                      `json:"denyEncryptionScopeOverride"`
	HasImmutabilityPolicy          *bool                                      `json:"hasImmutabilityPolicy"`
	HasLegalHold                   *bool                                      `json:"hasLegalHold"`
	ImmutabilityPolicy             *ImmutabilityPolicyProperties_STATUS_ARM   `json:"immutabilityPolicy"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_STATUS_ARM `json:"immutableStorageWithVersioning"`
	LastModifiedTime               *string                                    `json:"lastModifiedTime"`
	LeaseDuration                  *ContainerProperties_LeaseDuration_STATUS  `json:"leaseDuration"`
	LeaseState                     *ContainerProperties_LeaseState_STATUS     `json:"leaseState"`
	LeaseStatus                    *ContainerProperties_LeaseStatus_STATUS    `json:"leaseStatus"`
	LegalHold                      *LegalHoldProperties_STATUS_ARM            `json:"legalHold"`
	Metadata                       map[string]string                          `json:"metadata"`
	PublicAccess                   *ContainerProperties_PublicAccess_STATUS   `json:"publicAccess"`
	RemainingRetentionDays         *int                                       `json:"remainingRetentionDays"`
	Version                        *string                                    `json:"version"`
}

// Deprecated version of ImmutabilityPolicyProperties_STATUS. Use v1api20210401.ImmutabilityPolicyProperties_STATUS instead
type ImmutabilityPolicyProperties_STATUS_ARM struct {
	Etag          *string                                `json:"etag"`
	Properties    *ImmutabilityPolicyProperty_STATUS_ARM `json:"properties"`
	UpdateHistory []UpdateHistoryProperty_STATUS_ARM     `json:"updateHistory"`
}

// Deprecated version of ImmutableStorageWithVersioning_STATUS. Use v1api20210401.ImmutableStorageWithVersioning_STATUS instead
type ImmutableStorageWithVersioning_STATUS_ARM struct {
	Enabled        *bool                                                 `json:"enabled"`
	MigrationState *ImmutableStorageWithVersioning_MigrationState_STATUS `json:"migrationState"`
	TimeStamp      *string                                               `json:"timeStamp"`
}

// Deprecated version of LegalHoldProperties_STATUS. Use v1api20210401.LegalHoldProperties_STATUS instead
type LegalHoldProperties_STATUS_ARM struct {
	HasLegalHold *bool                    `json:"hasLegalHold"`
	Tags         []TagProperty_STATUS_ARM `json:"tags"`
}

// Deprecated version of ImmutabilityPolicyProperty_STATUS. Use v1api20210401.ImmutabilityPolicyProperty_STATUS instead
type ImmutabilityPolicyProperty_STATUS_ARM struct {
	AllowProtectedAppendWrites            *bool                                    `json:"allowProtectedAppendWrites"`
	ImmutabilityPeriodSinceCreationInDays *int                                     `json:"immutabilityPeriodSinceCreationInDays"`
	State                                 *ImmutabilityPolicyProperty_State_STATUS `json:"state"`
}

// Deprecated version of TagProperty_STATUS. Use v1api20210401.TagProperty_STATUS instead
type TagProperty_STATUS_ARM struct {
	ObjectIdentifier *string `json:"objectIdentifier"`
	Tag              *string `json:"tag"`
	TenantId         *string `json:"tenantId"`
	Timestamp        *string `json:"timestamp"`
	Upn              *string `json:"upn"`
}

// Deprecated version of UpdateHistoryProperty_STATUS. Use v1api20210401.UpdateHistoryProperty_STATUS instead
type UpdateHistoryProperty_STATUS_ARM struct {
	ImmutabilityPeriodSinceCreationInDays *int                                 `json:"immutabilityPeriodSinceCreationInDays"`
	ObjectIdentifier                      *string                              `json:"objectIdentifier"`
	TenantId                              *string                              `json:"tenantId"`
	Timestamp                             *string                              `json:"timestamp"`
	Update                                *UpdateHistoryProperty_Update_STATUS `json:"update"`
	Upn                                   *string                              `json:"upn"`
}
