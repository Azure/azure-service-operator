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

type StorageResource struct {
	ResourceGroupName  string
	StorageAccountName string
	StorageAccount     storage.Account
}

type mockStorageManager struct {
	storageResource []StorageResource
}

func findStorage(res []StorageResource, predicate func(StorageResource) bool) (int, StorageResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, StorageResource{}
}

type Finder interface {
}

type StorageResources []StorageResource

func (srs *StorageResources) Find(predicate func(StorageResource) bool) {

}

func (manager *mockStorageManager) CreateStorage(ctx context.Context, groupName string,
	storageAccountName string,
	location string,
	sku apiv1.StorageSku,
	kind apiv1.StorageKind,
	tags map[string]*string,
	accessTier apiv1.StorageAccessTier,
	enableHTTPsTrafficOnly *bool) (*storage.Account, error) {
	s := StorageResource{
		ResourceGroupName:  groupName,
		StorageAccountName: storageAccountName,
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

	index, group := findStorage(groups, func(g StorageResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.StorageAccountName == accountName
	})

	if index == -1 {
		return storage.Account{}, errors.New("storage account not found")
	}

	return group.StorageAccount, nil
}

// removes the storage account
func (manager *mockStorageManager) DeleteStorage(ctx context.Context, resourceGroupName string, accountName string) (autorest.Response, error) {
	groups := manager.storageResource

	index, _ := findStorage(groups, func(g StorageResource) bool {
		return g.ResourceGroupName == resourceGroupName &&
			g.StorageAccountName == accountName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("storage account not found")
	}

	manager.storageResource = append(groups[:index], groups[index+1:]...)

	return helpers.GetRestResponse(200), nil
}
