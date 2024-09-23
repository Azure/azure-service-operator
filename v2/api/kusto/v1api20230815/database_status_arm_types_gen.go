// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230815

import "encoding/json"

type Database_STATUS_ARM struct {
	// ReadOnlyFollowing: Mutually exclusive with all other properties
	ReadOnlyFollowing *ReadOnlyFollowingDatabase_STATUS_ARM `json:"readOnlyFollowing,omitempty"`

	// ReadWrite: Mutually exclusive with all other properties
	ReadWrite *ReadWriteDatabase_STATUS_ARM `json:"readWrite,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because Database_STATUS_ARM represents a discriminated union (JSON OneOf)
func (database Database_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if database.ReadOnlyFollowing != nil {
		return json.Marshal(database.ReadOnlyFollowing)
	}
	if database.ReadWrite != nil {
		return json.Marshal(database.ReadWrite)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the Database_STATUS_ARM
func (database *Database_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["kind"]
	if discriminator == "ReadOnlyFollowing" {
		database.ReadOnlyFollowing = &ReadOnlyFollowingDatabase_STATUS_ARM{}
		return json.Unmarshal(data, database.ReadOnlyFollowing)
	}
	if discriminator == "ReadWrite" {
		database.ReadWrite = &ReadWriteDatabase_STATUS_ARM{}
		return json.Unmarshal(data, database.ReadWrite)
	}

	// No error
	return nil
}

type ReadOnlyFollowingDatabase_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the database
	Kind ReadOnlyFollowingDatabase_Kind_STATUS_ARM `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The database properties.
	Properties *ReadOnlyFollowingDatabaseProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ReadWriteDatabase_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Kind: Kind of the database
	Kind ReadWriteDatabase_Kind_STATUS_ARM `json:"kind,omitempty"`

	// Location: Resource location.
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The database properties.
	Properties *ReadWriteDatabaseProperties_STATUS_ARM `json:"properties,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

type ReadOnlyFollowingDatabase_Kind_STATUS_ARM string

const ReadOnlyFollowingDatabase_Kind_STATUS_ARM_ReadOnlyFollowing = ReadOnlyFollowingDatabase_Kind_STATUS_ARM("ReadOnlyFollowing")

// Mapping from string to ReadOnlyFollowingDatabase_Kind_STATUS_ARM
var readOnlyFollowingDatabase_Kind_STATUS_ARM_Values = map[string]ReadOnlyFollowingDatabase_Kind_STATUS_ARM{
	"readonlyfollowing": ReadOnlyFollowingDatabase_Kind_STATUS_ARM_ReadOnlyFollowing,
}

// Class representing the Kusto database properties.
type ReadOnlyFollowingDatabaseProperties_STATUS_ARM struct {
	// AttachedDatabaseConfigurationName: The name of the attached database configuration cluster
	AttachedDatabaseConfigurationName *string `json:"attachedDatabaseConfigurationName,omitempty"`

	// DatabaseShareOrigin: The origin of the following setup.
	DatabaseShareOrigin *DatabaseShareOrigin_STATUS_ARM `json:"databaseShareOrigin,omitempty"`

	// HotCachePeriod: The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod *string `json:"hotCachePeriod,omitempty"`

	// LeaderClusterResourceId: The name of the leader cluster
	LeaderClusterResourceId *string `json:"leaderClusterResourceId,omitempty"`

	// OriginalDatabaseName: The original database name, before databaseNameOverride or databaseNamePrefix where applied.
	OriginalDatabaseName *string `json:"originalDatabaseName,omitempty"`

	// PrincipalsModificationKind: The principals modification kind of the database
	PrincipalsModificationKind *ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM `json:"principalsModificationKind,omitempty"`

	// ProvisioningState: The provisioned state of the resource.
	ProvisioningState *ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// SoftDeletePeriod: The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod *string `json:"softDeletePeriod,omitempty"`

	// Statistics: The statistics of the database.
	Statistics *DatabaseStatistics_STATUS_ARM `json:"statistics,omitempty"`

	// SuspensionDetails: The database suspension details. If the database is suspended, this object contains information
	// related to the database's suspension state.
	SuspensionDetails *SuspensionDetails_STATUS_ARM `json:"suspensionDetails,omitempty"`

	// TableLevelSharingProperties: Table level sharing specifications
	TableLevelSharingProperties *TableLevelSharingProperties_STATUS_ARM `json:"tableLevelSharingProperties,omitempty"`
}

type ReadWriteDatabase_Kind_STATUS_ARM string

const ReadWriteDatabase_Kind_STATUS_ARM_ReadWrite = ReadWriteDatabase_Kind_STATUS_ARM("ReadWrite")

// Mapping from string to ReadWriteDatabase_Kind_STATUS_ARM
var readWriteDatabase_Kind_STATUS_ARM_Values = map[string]ReadWriteDatabase_Kind_STATUS_ARM{
	"readwrite": ReadWriteDatabase_Kind_STATUS_ARM_ReadWrite,
}

// Class representing the Kusto database properties.
type ReadWriteDatabaseProperties_STATUS_ARM struct {
	// HotCachePeriod: The time the data should be kept in cache for fast queries in TimeSpan.
	HotCachePeriod *string `json:"hotCachePeriod,omitempty"`

	// IsFollowed: Indicates whether the database is followed.
	IsFollowed *bool `json:"isFollowed,omitempty"`

	// KeyVaultProperties: KeyVault properties for the database encryption.
	KeyVaultProperties *KeyVaultProperties_STATUS_ARM `json:"keyVaultProperties,omitempty"`

	// ProvisioningState: The provisioned state of the resource.
	ProvisioningState *ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// SoftDeletePeriod: The time the data should be kept before it stops being accessible to queries in TimeSpan.
	SoftDeletePeriod *string `json:"softDeletePeriod,omitempty"`

	// Statistics: The statistics of the database.
	Statistics *DatabaseStatistics_STATUS_ARM `json:"statistics,omitempty"`

	// SuspensionDetails: The database suspension details. If the database is suspended, this object contains information
	// related to the database's suspension state.
	SuspensionDetails *SuspensionDetails_STATUS_ARM `json:"suspensionDetails,omitempty"`
}

// The origin of the following setup.
type DatabaseShareOrigin_STATUS_ARM string

const (
	DatabaseShareOrigin_STATUS_ARM_DataShare = DatabaseShareOrigin_STATUS_ARM("DataShare")
	DatabaseShareOrigin_STATUS_ARM_Direct    = DatabaseShareOrigin_STATUS_ARM("Direct")
	DatabaseShareOrigin_STATUS_ARM_Other     = DatabaseShareOrigin_STATUS_ARM("Other")
)

// Mapping from string to DatabaseShareOrigin_STATUS_ARM
var databaseShareOrigin_STATUS_ARM_Values = map[string]DatabaseShareOrigin_STATUS_ARM{
	"datashare": DatabaseShareOrigin_STATUS_ARM_DataShare,
	"direct":    DatabaseShareOrigin_STATUS_ARM_Direct,
	"other":     DatabaseShareOrigin_STATUS_ARM_Other,
}

// A class that contains database statistics information.
type DatabaseStatistics_STATUS_ARM struct {
	// Size: The database size - the total size of compressed data and index in bytes.
	Size *float64 `json:"size,omitempty"`
}

type ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM string

const (
	ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_None    = ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM("None")
	ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_Replace = ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM("Replace")
	ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_Union   = ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM("Union")
)

// Mapping from string to ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM
var readOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_Values = map[string]ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM{
	"none":    ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_None,
	"replace": ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_Replace,
	"union":   ReadOnlyFollowingDatabaseProperties_PrincipalsModificationKind_STATUS_ARM_Union,
}

// The database suspension details. If the database is suspended, this object contains information related to the
// database's suspension state.
type SuspensionDetails_STATUS_ARM struct {
	// SuspensionStartDate: The starting date and time of the suspension state.
	SuspensionStartDate *string `json:"suspensionStartDate,omitempty"`
}

// Tables that will be included and excluded in the follower database
type TableLevelSharingProperties_STATUS_ARM struct {
	// ExternalTablesToExclude: List of external tables to exclude from the follower database
	ExternalTablesToExclude []string `json:"externalTablesToExclude,omitempty"`

	// ExternalTablesToInclude: List of external tables to include in the follower database
	ExternalTablesToInclude []string `json:"externalTablesToInclude,omitempty"`

	// FunctionsToExclude: List of functions to exclude from the follower database
	FunctionsToExclude []string `json:"functionsToExclude,omitempty"`

	// FunctionsToInclude: List of functions to include in the follower database
	FunctionsToInclude []string `json:"functionsToInclude,omitempty"`

	// MaterializedViewsToExclude: List of materialized views to exclude from the follower database
	MaterializedViewsToExclude []string `json:"materializedViewsToExclude,omitempty"`

	// MaterializedViewsToInclude: List of materialized views to include in the follower database
	MaterializedViewsToInclude []string `json:"materializedViewsToInclude,omitempty"`

	// TablesToExclude: List of tables to exclude from the follower database
	TablesToExclude []string `json:"tablesToExclude,omitempty"`

	// TablesToInclude: List of tables to include in the follower database
	TablesToInclude []string `json:"tablesToInclude,omitempty"`
}
