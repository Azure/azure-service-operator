// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

type StorageManagers struct {
	Storage       StorageManager
	BlobContainer BlobContainerManager
	FileSystem    FileSystemManager
}

var AzureStorageManagers = StorageManagers{
	Storage:       &azureStorageManager{},
	BlobContainer: &azureBlobContainerManager{},
	FileSystem:    &azureFileSystemManager{},
}
