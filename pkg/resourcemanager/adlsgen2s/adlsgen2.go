package adlsgen2s

import (
	"context"
	// "github.com/Azure/azure-sdk-for-go/services/storage/datalake/2019-10-31/storagedatalake"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"

	"errors"
	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"log"
)

type azureAdlsGen2Manager struct{}

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

func (_ *azureAdlsGen2Manager) CreateAdlsGen2(ctx context.Context, groupName string, datalakeName string, location string, sku azurev1alpha1.StorageSku, kind azurev1alpha1.StorageKind, tags map[string]*string, accessTier azurev1alpha1.StorageAccessTier, enableHTTPSTrafficOnly *bool) (*storage.Account, error) {
	// TODO: this is copy and pasted from storage.go. Figure out a better way to refactor so there isn't duplicate code
	storagesClient := getStoragesClient()

	//Check if name is available
	storageType := "Microsoft.Storage/storageAccounts"
	checkAccountParams := storage.AccountCheckNameAvailabilityParameters{Name: &datalakeName, Type: &storageType}
	checkNameResult, err := storagesClient.CheckNameAvailability(ctx, checkAccountParams)
	if err != nil {
		return nil, err
	}

	if *checkNameResult.NameAvailable == false {
		log.Fatalf("storage account not available: %v\n", checkNameResult.Reason)
		return nil, errors.New("storage account name not available")
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
			EnableHTTPSTrafficOnly: enableHTTPSTrafficOnly,
			IsHnsEnabled:           to.BoolPtr(true),
		},
	}

	// TODO: check to make sure that kind = StorageV2
	//log.Println(fmt.Sprintf("creating storage '%s' in resource group '%s' and location: %v", storageAccountName, groupName, location))
	future, err := storagesClient.Create(ctx, groupName, datalakeName, params)
	if err != nil {
		return nil, err
	}

	err = future.WaitForCompletionRef(ctx, storagesClient.Client)
	if err != nil {
		return nil, err
	}
	result, err := future.Result(storagesClient)
	return &result, err
}

func (_ *azureAdlsGen2Manager) GetAdlsGen2(ctx context.Context, groupName string, datalakeName string) (result storage.Account, err error) {
	adlsClient := getStoragesClient()
	return adlsClient.GetProperties(ctx, groupName, datalakeName, "")
}

func (_ *azureAdlsGen2Manager) DeleteAdlsGen2(ctx context.Context, groupName string, datalakeName string) (result autorest.Response, err error) {

	return autorest.Response{}, err
}

// func getFsClient(accountName string) storagedatalake.FilesystemClient {
// 	xmsversion := "2019-10-31"
// 	fsClient := storagedatalake.NewFilesystemClient(xmsversion, accountName)

// 	a, err := iam.GetResourceManagementAuthorizer()
// 	if err != nil {
// 		log.Fatalf("failed to initialize authorizer: %v\n", err)
// 	}
// 	fsClient.Authorizer = a
// 	fsClient.AddToUserAgent(config.UserAgent())
// 	return fsClient
// }
