// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

// Deprecated version of StorageAccounts_BlobService_STATUS. Use v1api20210401.StorageAccounts_BlobService_STATUS instead
type StorageAccounts_BlobService_STATUS_ARM struct {
	Id         *string                                            `json:"id,omitempty"`
	Name       *string                                            `json:"name,omitempty"`
	Properties *StorageAccounts_BlobService_Properties_STATUS_ARM `json:"properties,omitempty"`
	Sku        *Sku_STATUS_ARM                                    `json:"sku,omitempty"`
	Type       *string                                            `json:"type,omitempty"`
}

// Deprecated version of StorageAccounts_BlobService_Properties_STATUS. Use v1api20210401.StorageAccounts_BlobService_Properties_STATUS instead
type StorageAccounts_BlobService_Properties_STATUS_ARM struct {
	AutomaticSnapshotPolicyEnabled *bool                                    `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed_STATUS_ARM                   `json:"changeFeed,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_STATUS_ARM        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules_STATUS_ARM                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                                  `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy_STATUS_ARM        `json:"deleteRetentionPolicy,omitempty"`
	IsVersioningEnabled            *bool                                    `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy_STATUS_ARM `json:"lastAccessTimeTrackingPolicy,omitempty"`
	RestorePolicy                  *RestorePolicyProperties_STATUS_ARM      `json:"restorePolicy,omitempty"`
}

// Deprecated version of ChangeFeed_STATUS. Use v1api20210401.ChangeFeed_STATUS instead
type ChangeFeed_STATUS_ARM struct {
	Enabled         *bool `json:"enabled,omitempty"`
	RetentionInDays *int  `json:"retentionInDays,omitempty"`
}

// Deprecated version of CorsRules_STATUS. Use v1api20210401.CorsRules_STATUS instead
type CorsRules_STATUS_ARM struct {
	CorsRules []CorsRule_STATUS_ARM `json:"corsRules"`
}

// Deprecated version of DeleteRetentionPolicy_STATUS. Use v1api20210401.DeleteRetentionPolicy_STATUS instead
type DeleteRetentionPolicy_STATUS_ARM struct {
	Days    *int  `json:"days,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
}

// Deprecated version of LastAccessTimeTrackingPolicy_STATUS. Use v1api20210401.LastAccessTimeTrackingPolicy_STATUS instead
type LastAccessTimeTrackingPolicy_STATUS_ARM struct {
	BlobType                  []string                                  `json:"blobType"`
	Enable                    *bool                                     `json:"enable,omitempty"`
	Name                      *LastAccessTimeTrackingPolicy_Name_STATUS `json:"name,omitempty"`
	TrackingGranularityInDays *int                                      `json:"trackingGranularityInDays,omitempty"`
}

// Deprecated version of RestorePolicyProperties_STATUS. Use v1api20210401.RestorePolicyProperties_STATUS instead
type RestorePolicyProperties_STATUS_ARM struct {
	Days            *int    `json:"days,omitempty"`
	Enabled         *bool   `json:"enabled,omitempty"`
	LastEnabledTime *string `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string `json:"minRestoreTime,omitempty"`
}

// Deprecated version of CorsRule_STATUS. Use v1api20210401.CorsRule_STATUS instead
type CorsRule_STATUS_ARM struct {
	AllowedHeaders  []string                         `json:"allowedHeaders"`
	AllowedMethods  []CorsRule_AllowedMethods_STATUS `json:"allowedMethods"`
	AllowedOrigins  []string                         `json:"allowedOrigins"`
	ExposedHeaders  []string                         `json:"exposedHeaders"`
	MaxAgeInSeconds *int                             `json:"maxAgeInSeconds,omitempty"`
}
