// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
)

type StorageManagers struct {
	Storage       storageaccount.StorageManager
	BlobContainer blobcontainer.BlobContainerManager
	FileSystem    FileSystemManager
}

var AzureStorageManagers = StorageManagers{
	Storage:       storageaccount.New(),
	BlobContainer: blobcontainer.New(),
	FileSystem:    &azureFileSystemManager{},
}
