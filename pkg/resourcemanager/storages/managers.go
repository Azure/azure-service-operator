// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
)

type StorageManagers struct {
	StorageAccount storageaccount.StorageManager
	BlobContainer  blobcontainer.BlobContainerManager
}

var AzureStorageManagers = StorageManagers{
	StorageAccount: storageaccount.New(),
	BlobContainer:  blobcontainer.New(),
}
