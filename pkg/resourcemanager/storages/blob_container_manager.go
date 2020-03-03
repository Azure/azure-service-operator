// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

type BlobContainerManager interface {
	CreateBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string, accessLevel s.PublicAccess) (*storage.BlobContainer, error)

	// Get gets the description of the specified blob container.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// accountName - the name of the storage account
	// containerName - the name of the container
	GetBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result storage.BlobContainer, err error)

	// Deletes a blob container in a storage account.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// accountName - the name of the storage account
	// containerName - the name of the container
	DeleteBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result autorest.Response, err error)
}
