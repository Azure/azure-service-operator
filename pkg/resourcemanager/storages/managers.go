// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import "github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"

type StorageManagers struct {
	Storage       storageaccount.StorageManager
	BlobContainer BlobContainerManager
	FileSystem    FileSystemManager
}

var AzureStorageManagers = StorageManagers{
	Storage:       storageaccount.New(),
	BlobContainer: &azureBlobContainerManager{},
	FileSystem:    &azureFileSystemManager{},
}
