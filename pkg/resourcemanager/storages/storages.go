package storages

import (
	"context"
	//"encoding/json"
	"errors"
	"fmt"
	"log"

	//uuid "github.com/satori/go.uuid"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	//azureV1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	//"github.com/Azure/azure-service-operator/pkg/client/deployment"
	//"github.com/Azure/azure-service-operator/pkg/template"
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// New generates a new object
// func New(storage *azureV1alpha1.Storage) *Template {
// 	return &Template{
// 		Storage: storage,
// 	}
// }

// // Template defines the dynamodb cfts
// type Template struct {
// 	Storage *azureV1alpha1.Storage
// }

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
	sku azurev1.StorageSku,
	kind azurev1.StorageKind,
	tags map[string]*string,
	accessTier azurev1.StorageAccessTier,
	enableHTTPsTrafficOnly *bool) (storage.Account, error) {
	storagesClient := getStoragesClient()
	// id, err := uuid.FromString(config.TenantID())
	// if err != nil {
	// 	return storage.Account{}, err
	// }
	log.Println("Storage:AccountName" + storageAccountName)
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
	//mskuname := string(sku.Name)
	//test := storage.Sku{Name{string(sku.Name)}
	//msku := storage.Sku{Name : mskuname}//"test"}
	//t2 := storage.Sku(sku)
	sSku := storage.Sku{Name: storage.SkuName(sku.Name)}
	//sKind := string(kind)
	//sKind2 := storage.Kind{storage.Kind(kind)}
	sKind := storage.Kind(kind)
	//sKind := storage.Kind{Kind : kind}
	sAccessTier := storage.AccessTier(accessTier)

	params := storage.AccountCreateParameters{
		Location: to.StringPtr(location),
		Sku:      &sSku,
		//storage.Sku{Name{sku.Name}},
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
	return future.Result(storagesClient)
}

// DeleteStorage removes the resource group named by env var
func DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.Delete(ctx, groupName, storageAccountName)
}

// Pre-Refactor
// func (t *Template) CreateDeployment(ctx context.Context, resourceGroupName string) (string, error) {
// 	deploymentName := uuid.NewV4().String()
// 	asset, err := template.Asset("storage.json")
// 	templateContents := make(map[string]interface{})
// 	json.Unmarshal(asset, &templateContents)
// 	params := map[string]interface{}{
// 		"location": map[string]interface{}{
// 			"value": t.Storage.Spec.Location,
// 		},
// 		"accountType": map[string]interface{}{
// 			"value": t.Storage.Spec.Sku.Name,
// 		},
// 		"kind": map[string]interface{}{
// 			"value": t.Storage.Spec.Kind,
// 		},
// 		"accessTier": map[string]interface{}{
// 			"value": t.Storage.Spec.AccessTier,
// 		},
// 		"supportsHttpsTrafficOnly": map[string]interface{}{
// 			"value": *t.Storage.Spec.EnableHTTPSTrafficOnly,
// 		},
// 	}

// 	err = deployment.CreateDeployment(ctx, resourceGroupName, deploymentName, &templateContents, &params)
// 	return deploymentName, err
// }
