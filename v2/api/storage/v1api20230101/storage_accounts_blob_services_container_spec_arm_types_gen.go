// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccounts_BlobServices_Container_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of the blob container.
	Properties *ContainerProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccounts_BlobServices_Container_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (container StorageAccounts_BlobServices_Container_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (container *StorageAccounts_BlobServices_Container_Spec_ARM) GetName() string {
	return container.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices/containers"
func (container *StorageAccounts_BlobServices_Container_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices/containers"
}

// The properties of a container.
type ContainerProperties_ARM struct {
	// DefaultEncryptionScope: Default the container to use specified encryption scope for all writes.
	DefaultEncryptionScope *string `json:"defaultEncryptionScope,omitempty"`

	// DenyEncryptionScopeOverride: Block override of encryption scope from the container default.
	DenyEncryptionScopeOverride *bool `json:"denyEncryptionScopeOverride,omitempty"`

	// EnableNfsV3AllSquash: Enable NFSv3 all squash on blob container.
	EnableNfsV3AllSquash *bool `json:"enableNfsV3AllSquash,omitempty"`

	// EnableNfsV3RootSquash: Enable NFSv3 root squash on blob container.
	EnableNfsV3RootSquash *bool `json:"enableNfsV3RootSquash,omitempty"`

	// ImmutableStorageWithVersioning: The object level immutability property of the container. The property is immutable and
	// can only be set to true at the container creation time. Existing containers must undergo a migration process.
	ImmutableStorageWithVersioning *ImmutableStorageWithVersioning_ARM `json:"immutableStorageWithVersioning,omitempty"`

	// Metadata: A name-value pair to associate with the container as metadata.
	Metadata map[string]string `json:"metadata"`

	// PublicAccess: Specifies whether data in the container may be accessed publicly and the level of access.
	PublicAccess *ContainerProperties_PublicAccess_ARM `json:"publicAccess,omitempty"`
}

// +kubebuilder:validation:Enum={"Blob","Container","None"}
type ContainerProperties_PublicAccess_ARM string

const (
	ContainerProperties_PublicAccess_ARM_Blob      = ContainerProperties_PublicAccess_ARM("Blob")
	ContainerProperties_PublicAccess_ARM_Container = ContainerProperties_PublicAccess_ARM("Container")
	ContainerProperties_PublicAccess_ARM_None      = ContainerProperties_PublicAccess_ARM("None")
)

// Mapping from string to ContainerProperties_PublicAccess_ARM
var containerProperties_PublicAccess_ARM_Values = map[string]ContainerProperties_PublicAccess_ARM{
	"blob":      ContainerProperties_PublicAccess_ARM_Blob,
	"container": ContainerProperties_PublicAccess_ARM_Container,
	"none":      ContainerProperties_PublicAccess_ARM_None,
}

// Object level immutability properties of the container.
type ImmutableStorageWithVersioning_ARM struct {
	// Enabled: This is an immutable property, when set to true it enables object level immutability at the container level.
	Enabled *bool `json:"enabled,omitempty"`
}
