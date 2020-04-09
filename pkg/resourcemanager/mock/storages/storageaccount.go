// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storages

import (
	"context"
	"errors"
	"net/http"

	storage "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type storageResource struct {
	resourceGroupName  string
	storageAccountName string
	StorageAccount     storage.Account
}

type mockStorageManager struct {
	storageResource []storageResource
}

func findStorage(res []storageResource, predicate func(storageResource) bool) (int, storageResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, storageResource{}
}

type Finder interface {
}

type StorageResources []storageResource

func (srs *StorageResources) Find(predicate func(storageResource) bool) {

}

func (manager *mockStorageManager) CreateStorage(ctx context.Context, groupName string,
	storageAccountName string,
	location string,
	sku azurev1alpha1.StorageAccountSku,
	kind azurev1alpha1.StorageAccountKind,
	tags map[string]*string,
	accessTier azurev1alpha1.StorageAccountAccessTier,
	enableHTTPsTrafficOnly *bool, dataLakeEnabled *bool) (result storage.Account, err error) {
	s := storageResource{
		resourceGroupName:  groupName,
		storageAccountName: storageAccountName,
		StorageAccount: storage.Account{
			Response: helpers.GetRestResponse(201),
			Tags:     tags,
			Location: to.StringPtr(location),
			Name:     to.StringPtr(storageAccountName),
			AccountProperties: &storage.AccountProperties{
				ProvisioningState: storage.Succeeded,
			},
		},
	}

	manager.storageResource = append(manager.storageResource, s)
	return s.StorageAccount, nil
}

// Get gets the description of the specified storage account.
func (manager *mockStorageManager) GetStorage(ctx context.Context, resourceGroupName string, accountName string) (storage.Account, error) {
	groups := manager.storageResource

	index, group := findStorage(groups, func(g storageResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.storageAccountName == accountName
	})

	if index == -1 {
		return storage.Account{}, errors.New("storage account not found")
	}

	return group.StorageAccount, nil
}

// removes the storage account
func (manager *mockStorageManager) DeleteStorage(ctx context.Context, resourceGroupName string, accountName string) (autorest.Response, error) {
	groups := manager.storageResource

	index, _ := findStorage(groups, func(g storageResource) bool {
		return g.resourceGroupName == resourceGroupName &&
			g.storageAccountName == accountName
	})

	if index == -1 {
		return helpers.GetRestResponse(http.StatusNotFound), errors.New("storage account not found")
	}

	manager.storageResource = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(http.StatusOK), nil
}
