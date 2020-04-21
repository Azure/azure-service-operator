// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type StorageManagers struct {
	StorageAccount storageaccount.StorageManager
	BlobContainer  blobcontainer.BlobContainerManager
	FileSystem     FileSystemManager
}

func AzureStorageManagers(secretClient secrets.SecretClient) StorageManagers {
	return StorageManagers{
		StorageAccount: storageaccount.New(secretClient),
		BlobContainer:  blobcontainer.New(),
		FileSystem: &azureFileSystemManager{
			SecretClient: secretClient,
		},
	}
}

var EmptyAzureStorageManagers = StorageManagers{
	StorageAccount: storageaccount.EmptyStorageManager(),
	BlobContainer:  blobcontainer.EmptyBlobContainerManager(),
	FileSystem:     &azureFileSystemManager{},
}
