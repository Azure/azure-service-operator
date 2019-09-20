package storages

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	apiv1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
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
	sku apiv1.StorageSku,
	kind apiv1.StorageKind,
	tags map[string]*string,
	accessTier apiv1.StorageAccessTier,
	enableHTTPsTrafficOnly *bool) (*storage.Account, error) {
	s := storageResource{
		resourceGroupName:  groupName,
		storageAccountName: storageAccountName,
		StorageAccount: storage.Account{
			Response: helpers.GetRestResponse(201),
			Tags:     tags,
			Location: to.StringPtr(location),
			Name:     to.StringPtr(storageAccountName),
		},
	}

	manager.storageResource = append(manager.storageResource, s)
	return &s.StorageAccount, nil
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
		return helpers.GetRestResponse(404), errors.New("storage account not found")
	}

	manager.storageResource = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(200), nil
}
