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
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
)

type blobContainerResource struct {
	resourceGroupName  string
	storageAccountName string
	blobContainerName  string
	blobContainer      storage.BlobContainer
}

type mockBlobContainerManager struct {
	blobContainerResource []blobContainerResource
}

func findBlobContainer(res []blobContainerResource, predicate func(blobContainerResource) bool) (int, blobContainerResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, blobContainerResource{}
}

// Creates a blob container in a storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (manager *mockBlobContainerManager) CreateBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (*storage.BlobContainer, error) {
	bc := blobContainerResource{
		resourceGroupName:  resourceGroupName,
		storageAccountName: accountName,
		blobContainerName:  containerName,
		blobContainer:      storage.BlobContainer{},
	}
	manager.blobContainerResource = append(manager.blobContainerResource, bc)
	return &bc.blobContainer, nil
}

// Get gets the description of the specified blob container.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (manager *mockBlobContainerManager) GetBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (storage.BlobContainer, error) {
	containers := manager.blobContainerResource

	index, c := findBlobContainer(containers, func(g blobContainerResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.storageAccountName == accountName &&
			g.blobContainerName == containerName
	})

	if index == -1 {
		return storage.BlobContainer{}, errors.New("blob container not found")
	}

	return c.blobContainer, nil
}

// Deletes a blob container in a storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (manager *mockBlobContainerManager) DeleteBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (autorest.Response, error) {
	containers := manager.blobContainerResource

	index, _ := findBlobContainer(containers, func(g blobContainerResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.storageAccountName == accountName &&
			g.blobContainerName == containerName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("blob container not found")
	}

	manager.blobContainerResource = append(containers[:index], containers[index+1:]...)

	return helpers.GetRestResponse(200), nil
}
