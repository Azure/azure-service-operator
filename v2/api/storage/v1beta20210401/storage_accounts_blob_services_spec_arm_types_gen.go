// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20210401

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

<<<<<<<< HEAD:v2/api/storage/v1beta20210401/storage_accounts_blob_service__spec_arm_types_gen.go
type StorageAccountsBlobService_SpecARM struct {
	AzureName string `json:"azureName,omitempty"`
	Name      string `json:"name,omitempty"`
========
type StorageAccounts_BlobServices_SpecARM struct {
	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: The name of the blob Service within the specified storage account. Blob Service Name must be 'default'
	Name string `json:"name,omitempty"`
>>>>>>>> main:v2/api/storage/v1beta20210401/storage_accounts_blob_services_spec_arm_types_gen.go

	// Properties: The properties of a storage account’s Blob service.
	Properties *StorageAccountsBlobService_Spec_PropertiesARM `json:"properties,omitempty"`
}

<<<<<<<< HEAD:v2/api/storage/v1beta20210401/storage_accounts_blob_service__spec_arm_types_gen.go
var _ genruntime.ARMResourceSpec = &StorageAccountsBlobService_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccountsBlobService_SpecARM) GetAPIVersion() string {
========
var _ genruntime.ARMResourceSpec = &StorageAccounts_BlobServices_SpecARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (services StorageAccounts_BlobServices_SpecARM) GetAPIVersion() string {
>>>>>>>> main:v2/api/storage/v1beta20210401/storage_accounts_blob_services_spec_arm_types_gen.go
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
<<<<<<<< HEAD:v2/api/storage/v1beta20210401/storage_accounts_blob_service__spec_arm_types_gen.go
func (service *StorageAccountsBlobService_SpecARM) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices"
func (service *StorageAccountsBlobService_SpecARM) GetType() string {
========
func (services *StorageAccounts_BlobServices_SpecARM) GetName() string {
	return services.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices"
func (services *StorageAccounts_BlobServices_SpecARM) GetType() string {
>>>>>>>> main:v2/api/storage/v1beta20210401/storage_accounts_blob_services_spec_arm_types_gen.go
	return "Microsoft.Storage/storageAccounts/blobServices"
}

type StorageAccountsBlobService_Spec_PropertiesARM struct {
	// AutomaticSnapshotPolicyEnabled: Deprecated in favor of isVersioningEnabled property.
	AutomaticSnapshotPolicyEnabled *bool `json:"automaticSnapshotPolicyEnabled,omitempty"`

	// ChangeFeed: The blob service properties for change feed events.
	ChangeFeed *ChangeFeedARM `json:"changeFeed,omitempty"`

	// ContainerDeleteRetentionPolicy: The blob service properties for container soft delete.
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicyARM `json:"containerDeleteRetentionPolicy,omitempty"`

	// Cors: Specifies CORS rules for the Blob service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Blob service.
	Cors *CorsRulesARM `json:"cors,omitempty"`

	// DefaultServiceVersion: DefaultServiceVersion indicates the default version to use for requests to the Blob service if an
	// incoming request’s version is not specified. Possible values include version 2008-10-27 and all more recent versions.
	DefaultServiceVersion *string `json:"defaultServiceVersion,omitempty"`

	// DeleteRetentionPolicy: The blob service properties for blob soft delete.
	DeleteRetentionPolicy *DeleteRetentionPolicyARM `json:"deleteRetentionPolicy,omitempty"`

	// IsVersioningEnabled: Versioning is enabled if set to true.
	IsVersioningEnabled *bool `json:"isVersioningEnabled,omitempty"`

	// LastAccessTimeTrackingPolicy: The blob service property to configure last access time based tracking policy.
	LastAccessTimeTrackingPolicy *LastAccessTimeTrackingPolicyARM `json:"lastAccessTimeTrackingPolicy,omitempty"`

	// RestorePolicy: The blob service properties for blob restore policy.
	RestorePolicy *RestorePolicyPropertiesARM `json:"restorePolicy,omitempty"`
}

type ChangeFeedARM struct {
	// Enabled: Indicates whether change feed event logging is enabled for the Blob service.
	Enabled *bool `json:"enabled,omitempty"`

	// RetentionInDays: Indicates the duration of changeFeed retention in days. Minimum value is 1 day and maximum value is
	// 146000 days (400 years). A null value indicates an infinite retention of the change feed.
	RetentionInDays *int `json:"retentionInDays,omitempty"`
}

type CorsRulesARM struct {
	// CorsRules: The List of CORS rules. You can include up to five CorsRule elements in the request.
	CorsRules []CorsRuleARM `json:"corsRules,omitempty"`
}

type DeleteRetentionPolicyARM struct {
	// Days: Indicates the number of days that the deleted item should be retained. The minimum specified value can be 1 and
	// the maximum value can be 365.
	Days *int `json:"days,omitempty"`

	// Enabled: Indicates whether DeleteRetentionPolicy is enabled.
	Enabled *bool `json:"enabled,omitempty"`
}

type LastAccessTimeTrackingPolicyARM struct {
	// BlobType: An array of predefined supported blob types. Only blockBlob is the supported value. This field is currently
	// read only
	BlobType []string `json:"blobType,omitempty"`

	// Enable: When set to true last access time based tracking is enabled.
	Enable *bool `json:"enable,omitempty"`

<<<<<<<< HEAD:v2/api/storage/v1beta20210401/storage_accounts_blob_service__spec_arm_types_gen.go
	// Name: Name of the policy. The valid value is AccessTimeTracking. This field is currently read only
========
	// Name: Name of the policy. The valid value is AccessTimeTracking. This field is currently read only.
>>>>>>>> main:v2/api/storage/v1beta20210401/storage_accounts_blob_services_spec_arm_types_gen.go
	Name *LastAccessTimeTrackingPolicy_Name `json:"name,omitempty"`

	// TrackingGranularityInDays: The field specifies blob object tracking granularity in days, typically how often the blob
	// object should be tracked.This field is currently read only with value as 1
	TrackingGranularityInDays *int `json:"trackingGranularityInDays,omitempty"`
}

type RestorePolicyPropertiesARM struct {
	// Days: how long this blob can be restored. It should be great than zero and less than DeleteRetentionPolicy.days.
	Days *int `json:"days,omitempty"`

	// Enabled: Blob restore is enabled if set to true.
	Enabled *bool `json:"enabled,omitempty"`
}

type CorsRuleARM struct {
	// AllowedHeaders: Required if CorsRule element is present. A list of headers allowed to be part of the cross-origin
	// request.
	AllowedHeaders []string `json:"allowedHeaders,omitempty"`

	// AllowedMethods: Required if CorsRule element is present. A list of HTTP methods that are allowed to be executed by the
	// origin.
	AllowedMethods []CorsRule_AllowedMethods `json:"allowedMethods,omitempty"`

	// AllowedOrigins: Required if CorsRule element is present. A list of origin domains that will be allowed via CORS, or "*"
	// to allow all domains
	AllowedOrigins []string `json:"allowedOrigins,omitempty"`

	// ExposedHeaders: Required if CorsRule element is present. A list of response headers to expose to CORS clients.
	ExposedHeaders []string `json:"exposedHeaders,omitempty"`

	// MaxAgeInSeconds: Required if CorsRule element is present. The number of seconds that the client/browser should cache a
	// preflight response.
	MaxAgeInSeconds *int `json:"maxAgeInSeconds,omitempty"`
}
