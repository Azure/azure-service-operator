package storages

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	apiv1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

func getStoragesClient() storage.AccountsClient {
	storagesClient := storage.NewAccountsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	storagesClient.Authorizer = a
	storagesClient.AddToUserAgent(config.UserAgent())
	return storagesClient
}

// CreateStorage creates a new storage account
func CreateStorage(ctx context.Context, groupName string,
	storageAccountName string,
	location string,
	sku apiv1.StorageSku,
	kind apiv1.StorageKind,
	tags map[string]*string,
	accessTier apiv1.StorageAccessTier,
	enableHTTPsTrafficOnly *bool) (storage.Account, error) {
	storagesClient := getStoragesClient()

	storageType := "Microsoft.Storage/storageAccounts"
	checkAccountParams := storage.AccountCheckNameAvailabilityParameters{Name: &storageAccountName, Type: &storageType}
	result, err := storagesClient.CheckNameAvailability(ctx, checkAccountParams)
	if err != nil {
		return storage.Account{}, err
	}

	if *result.NameAvailable == false {
		log.Fatalf("storage account not available: %v\n", result.Reason)
		return storage.Account{}, errors.New("storage account not available")
	}

	sSku := storage.Sku{Name: storage.SkuName(sku.Name)}
	sKind := storage.Kind(kind)
	sAccessTier := storage.AccessTier(accessTier)

	params := storage.AccountCreateParameters{
		Location: to.StringPtr(location),
		Sku:      &sSku,
		Kind:     sKind,
		Tags:     tags,
		Identity: nil,
		AccountPropertiesCreateParameters: &storage.AccountPropertiesCreateParameters{
			AccessTier:             sAccessTier,
			EnableHTTPSTrafficOnly: enableHTTPsTrafficOnly,
		},
	}

	log.Println(fmt.Sprintf("creating storage '%s' in resource group '%s' and location: %v", storageAccountName, groupName, location))
	future, err := storagesClient.Create(ctx, groupName, storageAccountName, params)

	err = future.WaitForCompletionRef(ctx, storagesClient.Client)
	if err != nil {
		return storage.Account{}, err
	}

	return future.Result(storagesClient)
}

// Get gets the description of the specified storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
func GetStorage(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.GetProperties(ctx, resourceGroupName, accountName, "")
}

// DeleteStorage removes the resource group named by env var
func DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.Delete(ctx, groupName, storageAccountName)
}
