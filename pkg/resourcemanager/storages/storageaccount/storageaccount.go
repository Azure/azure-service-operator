// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storageaccount

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemgrconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

const templateForConnectionString = "DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=%s"

type azureStorageManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

// ParseNetworkPolicy - helper function to parse network policies from Kubernetes spec
func ParseNetworkPolicy(ruleSet *v1alpha1.StorageNetworkRuleSet) storage.NetworkRuleSet {

	bypass := storage.AzureServices
	switch ruleSet.Bypass {
	case "AzureServices":
		bypass = storage.AzureServices
	case "None":
		bypass = storage.None
	case "Logging":
		bypass = storage.Logging
	case "Metrics":
		bypass = storage.Metrics
	}

	defaultAction := storage.DefaultActionDeny
	if strings.EqualFold(ruleSet.DefaultAction, "allow") {
		defaultAction = storage.DefaultActionAllow
	}

	var ipInstances []storage.IPRule
	if ruleSet.IPRules != nil {
		for _, i := range *ruleSet.IPRules {
			ipmask := i.IPAddressOrRange
			ipInstances = append(ipInstances, storage.IPRule{
				IPAddressOrRange: ipmask,
				Action:           storage.Allow,
			})
		}
	}

	var vnetInstances []storage.VirtualNetworkRule
	if ruleSet.VirtualNetworkRules != nil {
		for _, i := range *ruleSet.VirtualNetworkRules {
			vnetID := i.SubnetId
			vnetInstances = append(vnetInstances, storage.VirtualNetworkRule{
				VirtualNetworkResourceID: vnetID,
				Action:                   storage.Allow,
			})
		}
	}

	return storage.NetworkRuleSet{
		Bypass:              bypass,
		DefaultAction:       defaultAction,
		IPRules:             &ipInstances,
		VirtualNetworkRules: &vnetInstances,
	}
}

func getStoragesClient() storage.AccountsClient {
	storagesClient := storage.NewAccountsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Printf("failed to initialize authorizer: %v\n", err)
	}
	storagesClient.Authorizer = a
	storagesClient.AddToUserAgent(config.UserAgent())
	return storagesClient
}

// CreateStorage creates a new storage account
func (_ *azureStorageManager) CreateStorage(ctx context.Context,
	groupName string,
	storageAccountName string,
	location string,
	sku azurev1alpha1.StorageAccountSku,
	kind azurev1alpha1.StorageAccountKind,
	tags map[string]*string,
	accessTier azurev1alpha1.StorageAccountAccessTier,
	enableHTTPsTrafficOnly *bool, dataLakeEnabled *bool, networkRule *storage.NetworkRuleSet) (pollingURL string, result storage.Account, err error) {

	storagesClient := getStoragesClient()

	//Check if name is available
	storageType := "Microsoft.Storage/storageAccounts"
	checkAccountParams := storage.AccountCheckNameAvailabilityParameters{Name: &storageAccountName, Type: &storageType}
	checkNameResult, err := storagesClient.CheckNameAvailability(ctx, checkAccountParams)
	if err != nil {
		return "", result, err
	}
	if dataLakeEnabled == to.BoolPtr(true) && kind != "StorageV2" {
		err = errors.New("unable to create datalake enabled storage account")
		return
	}
	if *checkNameResult.NameAvailable == false {
		if checkNameResult.Reason == storage.AccountNameInvalid {
			err = errors.New("AccountNameInvalid")
			return
		} else if checkNameResult.Reason == storage.AlreadyExists {
			err = errors.New("AlreadyExists")
			return
		}
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
			IsHnsEnabled:           dataLakeEnabled,
			NetworkRuleSet:         networkRule,
		},
	}

	//log.Println(fmt.Sprintf("creating storage '%s' in resource group '%s' and location: %v", storageAccountName, groupName, location))
	future, err := storagesClient.Create(ctx, groupName, storageAccountName, params)
	if err != nil {
		return "", result, err
	}

	result, err = future.Result(storagesClient)

	return future.PollingURL(), result, err

}

// Get gets the description of the specified storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// storageAccountName - the name of the storage account
func (_ *azureStorageManager) GetStorage(ctx context.Context, resourceGroupName string, storageAccountName string) (result storage.Account, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.GetProperties(ctx, resourceGroupName, storageAccountName, "")
}

// DeleteStorage removes the resource group named by env var
func (_ *azureStorageManager) DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.Delete(ctx, groupName, storageAccountName)
}

// ListKeys lists the keys
func (_ *azureStorageManager) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountListKeysResult, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.ListKeys(ctx, resourceGroupName, accountName, storage.Kerb)
}

// StoreSecrets upserts the secret information for this storage account
func (s *azureStorageManager) StoreSecrets(ctx context.Context, resourceGroupName string, accountName string, instance *v1alpha1.StorageAccount) error {

	// get the keys
	keyResult, err := s.ListKeys(ctx, resourceGroupName, accountName)
	if err != nil {
		return err
	}
	if keyResult.Keys == nil {
		return fmt.Errorf("No keys were returned from ListKeys")
	}
	keys := *keyResult.Keys
	storageEndpointSuffix := resourcemgrconfig.Environment().StorageEndpointSuffix

	// build the connection string
	data := map[string][]byte{
		"StorageAccountName": []byte(accountName),
	}
	for i, key := range keys {
		data[fmt.Sprintf("ConnectionString%v", i)] = []byte(fmt.Sprintf(templateForConnectionString, accountName, *key.Value, storageEndpointSuffix))
		data[fmt.Sprintf("Key%v", i)] = []byte(*key.Value)
	}

	// upsert
	key := types.NamespacedName{
		Name:      fmt.Sprintf("StorageAccount-%s-%s", resourceGroupName, accountName),
		Namespace: instance.Namespace,
	}
	return s.SecretClient.Upsert(ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
	)
}
