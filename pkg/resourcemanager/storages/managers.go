// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/blobcontainer"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/storages/storageaccount"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"k8s.io/apimachinery/pkg/runtime"
)

type StorageManagers struct {
	StorageAccount storageaccount.StorageManager
	BlobContainer  blobcontainer.BlobContainerManager
}

func NewAzureStorageManagers(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) StorageManagers {
	return StorageManagers{
		StorageAccount: storageaccount.NewManager(creds, secretClient, scheme),
		BlobContainer:  blobcontainer.NewManager(creds),
	}
}
