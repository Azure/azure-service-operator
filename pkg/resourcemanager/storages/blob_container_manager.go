/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storages

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
)

type BlobContainerManager interface {
	CreateBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (*storage.BlobContainer, error)

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
