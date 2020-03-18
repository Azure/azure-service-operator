// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"context"
	"fmt"
	"log"

	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
)

type azureBlobContainerManager struct{}

func getContainerClient() s.BlobContainersClient {
	containersClient := s.NewBlobContainersClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	containersClient.Authorizer = auth
	containersClient.AddToUserAgent(config.UserAgent())
	return containersClient
}

// Creates a blob container in a storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
// accessLevel - 'PublicAccessContainer', 'PublicAccessBlob', or 'PublicAccessNone'
func (_ *azureBlobContainerManager) CreateBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string, accessLevel s.PublicAccess) (*s.BlobContainer, error) {
	containerClient := getContainerClient()

	blobContainerProperties := s.ContainerProperties{
		PublicAccess: accessLevel,
	}

	log.Println(fmt.Sprintf("Creating blob container '%s' in storage account: %s", containerName, accountName))

	container, err := containerClient.Create(
		ctx,
		resourceGroupName,
		accountName,
		containerName,
		s.BlobContainer{ContainerProperties: &blobContainerProperties})

	if err != nil {
		return nil, err
	}

	return &container, err
}

// Get gets the description of the specified blob container.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (_ *azureBlobContainerManager) GetBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result s.BlobContainer, err error) {
	containerClient := getContainerClient()
	return containerClient.Get(ctx, resourceGroupName, accountName, containerName)
}

// Deletes a blob container in a storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// containerName - the name of the container
func (_ *azureBlobContainerManager) DeleteBlobContainer(ctx context.Context, resourceGroupName string, accountName string, containerName string) (result autorest.Response, err error) {
	containerClient := getContainerClient()
	log.Println(fmt.Sprintf("Deleting blob container '%s' for resource group: %s", containerName, accountName))

	return containerClient.Delete(ctx,
		resourceGroupName,
		accountName,
		containerName)
}
