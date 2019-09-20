package storages

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
)

type BlobContainerResource struct {
	ResourceGroupName  string
	StorageAccountName string
	BlobContainerName  string
	BlobContainer      storage.BlobContainer
}

type mockBlobContainerManager struct {
	blobContainerResource []BlobContainerResource
}

func findBlobContainer(res []BlobContainerResource, predicate func(BlobContainerResource) bool) (int, BlobContainerResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, BlobContainerResource{}
}

// Creates a blob container in a storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (manager *mockBlobContainerManager) CreateBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (*storage.BlobContainer, error) {
	bc := BlobContainerResource{
		ResourceGroupName:  resourceGroupName,
		StorageAccountName: accountName,
		BlobContainerName:  containerName,
		BlobContainer:      storage.BlobContainer{},
	}
	manager.blobContainerResource = append(manager.blobContainerResource, bc)
	return &bc.BlobContainer, nil
}

// Get gets the description of the specified blob container.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (manager *mockBlobContainerManager) GetBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (storage.BlobContainer, error) {
	containers := manager.blobContainerResource

	index, c := findBlobContainer(containers, func(g BlobContainerResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.StorageAccountName == accountName &&
			g.BlobContainerName == containerName
	})

	if index == -1 {
		return storage.BlobContainer{}, errors.New("blob container not found")
	}

	return c.BlobContainer, nil
}

// Deletes a blob container in a storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (manager *mockBlobContainerManager) DeleteBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (autorest.Response, error) {
	containers := manager.blobContainerResource

	index, _ := findBlobContainer(containers, func(g BlobContainerResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.StorageAccountName == accountName &&
			g.BlobContainerName == containerName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("blob container not found")
	}

	manager.blobContainerResource = append(containers[:index], containers[index+1:]...)

	return helpers.GetRestResponse(200), nil
}
