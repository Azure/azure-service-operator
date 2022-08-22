// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401

// Deprecated version of BlobContainer_STATUS. Use v1beta20210401.BlobContainer_STATUS instead
type BlobContainer_STATUSARM struct {
	Etag       *string                        `json:"etag,omitempty"`
	Id         *string                        `json:"id,omitempty"`
	Name       *string                        `json:"name,omitempty"`
	Properties *ContainerProperties_STATUSARM `json:"properties,omitempty"`
	Type       *string                        `json:"type,omitempty"`
}

// Deprecated version of ContainerProperties_STATUS. Use v1beta20210401.ContainerProperties_STATUS instead
type ContainerProperties_STATUSARM struct {
	DefaultEncryptionScope         *string                                   `json:"defaultEncryptionScope,omitempty"`
	Deleted                        *bool                                     `json:"deleted,omitempty"`
	DeletedTime                    *string                                   `json:"deletedTime,omitempty"`
	DenyEncryptionScopeOverride    *bool                                     `json:"denyEncryptionScopeOverride,omitempty"`
	HasImmutabilityPolicy          *bool                                     `json:"hasImmutabilityPolicy,omitempty"`
	HasLegalHold                   *bool                                     `json:"hasLegalHold,omitempty"`
	ImmutabilityPolicy             *ImmutabilityPolicyProperties_STATUSARM   `json:"immutabilityPolicy,omitempty"`
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_STATUSARM `json:"immutableStorageWithVersioning,omitempty"`
	LastModifiedTime               *string                                   `json:"lastModifiedTime,omitempty"`
	LeaseDuration                  *ContainerPropertiesSTATUSLeaseDuration   `json:"leaseDuration,omitempty"`
	LeaseState                     *ContainerPropertiesSTATUSLeaseState      `json:"leaseState,omitempty"`
	LeaseStatus                    *ContainerPropertiesSTATUSLeaseStatus     `json:"leaseStatus,omitempty"`
	LegalHold                      *LegalHoldProperties_STATUSARM            `json:"legalHold,omitempty"`
	Metadata                       map[string]string                         `json:"metadata,omitempty"`
	PublicAccess                   *ContainerPropertiesSTATUSPublicAccess    `json:"publicAccess,omitempty"`
	RemainingRetentionDays         *int                                      `json:"remainingRetentionDays,omitempty"`
	Version                        *string                                   `json:"version,omitempty"`
}

// Deprecated version of ContainerPropertiesSTATUSLeaseDuration. Use v1beta20210401.ContainerPropertiesSTATUSLeaseDuration
// instead
type ContainerPropertiesSTATUSLeaseDuration string

const (
	ContainerPropertiesSTATUSLeaseDuration_Fixed    = ContainerPropertiesSTATUSLeaseDuration("Fixed")
	ContainerPropertiesSTATUSLeaseDuration_Infinite = ContainerPropertiesSTATUSLeaseDuration("Infinite")
)

// Deprecated version of ContainerPropertiesSTATUSLeaseState. Use v1beta20210401.ContainerPropertiesSTATUSLeaseState instead
type ContainerPropertiesSTATUSLeaseState string

const (
	ContainerPropertiesSTATUSLeaseState_Available = ContainerPropertiesSTATUSLeaseState("Available")
	ContainerPropertiesSTATUSLeaseState_Breaking  = ContainerPropertiesSTATUSLeaseState("Breaking")
	ContainerPropertiesSTATUSLeaseState_Broken    = ContainerPropertiesSTATUSLeaseState("Broken")
	ContainerPropertiesSTATUSLeaseState_Expired   = ContainerPropertiesSTATUSLeaseState("Expired")
	ContainerPropertiesSTATUSLeaseState_Leased    = ContainerPropertiesSTATUSLeaseState("Leased")
)

// Deprecated version of ContainerPropertiesSTATUSLeaseStatus. Use v1beta20210401.ContainerPropertiesSTATUSLeaseStatus
// instead
type ContainerPropertiesSTATUSLeaseStatus string

const (
	ContainerPropertiesSTATUSLeaseStatus_Locked   = ContainerPropertiesSTATUSLeaseStatus("Locked")
	ContainerPropertiesSTATUSLeaseStatus_Unlocked = ContainerPropertiesSTATUSLeaseStatus("Unlocked")
)

// Deprecated version of ContainerPropertiesSTATUSPublicAccess. Use v1beta20210401.ContainerPropertiesSTATUSPublicAccess
// instead
type ContainerPropertiesSTATUSPublicAccess string

const (
	ContainerPropertiesSTATUSPublicAccess_Blob      = ContainerPropertiesSTATUSPublicAccess("Blob")
	ContainerPropertiesSTATUSPublicAccess_Container = ContainerPropertiesSTATUSPublicAccess("Container")
	ContainerPropertiesSTATUSPublicAccess_None      = ContainerPropertiesSTATUSPublicAccess("None")
)

// Deprecated version of ImmutabilityPolicyProperties_STATUS. Use v1beta20210401.ImmutabilityPolicyProperties_STATUS instead
type ImmutabilityPolicyProperties_STATUSARM struct {
	Etag          *string                               `json:"etag,omitempty"`
	Properties    *ImmutabilityPolicyProperty_STATUSARM `json:"properties,omitempty"`
	UpdateHistory []UpdateHistoryProperty_STATUSARM     `json:"updateHistory,omitempty"`
}

// Deprecated version of ImmutableStorageWithVersioning_STATUS. Use v1beta20210401.ImmutableStorageWithVersioning_STATUS instead
type ImmutableStorageWithVersioning_STATUSARM struct {
	Enabled        *bool                                               `json:"enabled,omitempty"`
	MigrationState *ImmutableStorageWithVersioningSTATUSMigrationState `json:"migrationState,omitempty"`
	TimeStamp      *string                                             `json:"timeStamp,omitempty"`
}

// Deprecated version of LegalHoldProperties_STATUS. Use v1beta20210401.LegalHoldProperties_STATUS instead
type LegalHoldProperties_STATUSARM struct {
	HasLegalHold *bool                   `json:"hasLegalHold,omitempty"`
	Tags         []TagProperty_STATUSARM `json:"tags,omitempty"`
}

// Deprecated version of ImmutabilityPolicyProperty_STATUS. Use v1beta20210401.ImmutabilityPolicyProperty_STATUS instead
type ImmutabilityPolicyProperty_STATUSARM struct {
	AllowProtectedAppendWrites            *bool                                  `json:"allowProtectedAppendWrites,omitempty"`
	ImmutabilityPeriodSinceCreationInDays *int                                   `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	State                                 *ImmutabilityPolicyPropertySTATUSState `json:"state,omitempty"`
}

// Deprecated version of ImmutableStorageWithVersioningSTATUSMigrationState. Use
// v1beta20210401.ImmutableStorageWithVersioningSTATUSMigrationState instead
type ImmutableStorageWithVersioningSTATUSMigrationState string

const (
	ImmutableStorageWithVersioningSTATUSMigrationState_Completed  = ImmutableStorageWithVersioningSTATUSMigrationState("Completed")
	ImmutableStorageWithVersioningSTATUSMigrationState_InProgress = ImmutableStorageWithVersioningSTATUSMigrationState("InProgress")
)

// Deprecated version of TagProperty_STATUS. Use v1beta20210401.TagProperty_STATUS instead
type TagProperty_STATUSARM struct {
	ObjectIdentifier *string `json:"objectIdentifier,omitempty"`
	Tag              *string `json:"tag,omitempty"`
	TenantId         *string `json:"tenantId,omitempty"`
	Timestamp        *string `json:"timestamp,omitempty"`
	Upn              *string `json:"upn,omitempty"`
}

// Deprecated version of UpdateHistoryProperty_STATUS. Use v1beta20210401.UpdateHistoryProperty_STATUS instead
type UpdateHistoryProperty_STATUSARM struct {
	ImmutabilityPeriodSinceCreationInDays *int                               `json:"immutabilityPeriodSinceCreationInDays,omitempty"`
	ObjectIdentifier                      *string                            `json:"objectIdentifier,omitempty"`
	TenantId                              *string                            `json:"tenantId,omitempty"`
	Timestamp                             *string                            `json:"timestamp,omitempty"`
	Update                                *UpdateHistoryPropertySTATUSUpdate `json:"update,omitempty"`
	Upn                                   *string                            `json:"upn,omitempty"`
}

// Deprecated version of UpdateHistoryPropertySTATUSUpdate. Use v1beta20210401.UpdateHistoryPropertySTATUSUpdate instead
type UpdateHistoryPropertySTATUSUpdate string

const (
	UpdateHistoryPropertySTATUSUpdate_Extend = UpdateHistoryPropertySTATUSUpdate("extend")
	UpdateHistoryPropertySTATUSUpdate_Lock   = UpdateHistoryPropertySTATUSUpdate("lock")
	UpdateHistoryPropertySTATUSUpdate_Put    = UpdateHistoryPropertySTATUSUpdate("put")
)
