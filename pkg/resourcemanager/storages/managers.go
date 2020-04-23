// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
)

type StorageManagers struct {
	StorageAccount storageaccount.StorageManager
	BlobContainer  blobcontainer.BlobContainerManager
}

<<<<<<< HEAD
func AzureStorageManagers(secretClient secrets.SecretClient, scheme *runtime.Scheme) StorageManagers {
	return StorageManagers{
		StorageAccount: storageaccount.New(secretClient, scheme),
		BlobContainer:  blobcontainer.New(),
		FileSystem: &azureFileSystemManager{
			SecretClient: secretClient,
			Scheme:       scheme,
		},
	}
}

var EmptyAzureStorageManagers = StorageManagers{
	StorageAccount: storageaccount.EmptyStorageManager(),
	BlobContainer:  blobcontainer.EmptyBlobContainerManager(),
	FileSystem:     &azureFileSystemManager{},
=======
var AzureStorageManagers = StorageManagers{
	StorageAccount: storageaccount.New(),
	BlobContainer:  blobcontainer.New(),
>>>>>>> 81d67f98519f5d0e8dba4aa13b157cd7b1e44a44
}
