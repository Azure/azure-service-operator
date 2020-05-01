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

func NewAzureStorageManagers(secretClient secrets.SecretClient, scheme *runtime.Scheme) StorageManagers {
	return StorageManagers{
		StorageAccount: storageaccount.New(secretClient, scheme),
		BlobContainer:  blobcontainer.New(),
	}
}
