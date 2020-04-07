// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storageaccount

import (
	"context"
	"errors"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type azureStorageManager struct{}

// ParseNetworkPolicy - helper function to parse network policies from Kubernetes spec
func ParseNetworkPolicy(ruleSet *v1alpha1.StorageNetworkRuleSet) storage.NetworkRuleSet {
	var bypass storage.Bypass

	switch ruleSet.Bypass {
	case "AzureServices":
		bypass = storage.AzureServices
	case "None":
		bypass = storage.None
	case "Logging":
		bypass = storage.Logging
	case "Metrics":
		bypass = storage.Metrics
	default:
		bypass = storage.AzureServices
	}

	var defaultAction storage.DefaultAction
	switch ruleSet.DefaultAction {
	case "Allow":
		defaultAction = storage.DefaultActionAllow
	case "Deny":
		defaultAction = storage.DefaultActionDeny
	default:
		defaultAction = storage.DefaultActionDeny
	}

	var ipInstances []storage.IPRule
	if ruleSet.IPRule != nil {
		for _, i := range *ruleSet.IPRule {
			subnetID := i.IPAddressOrRange
			ipInstances = append(ipInstances, storage.IPRule{
				IPAddressOrRange: subnetID,
				Action:           storage.Allow,
			})
		}
	}

	var vnetInstances []storage.VirtualNetworkRule
	if ruleSet.VirtualNetworkRules != nil {
		for _, i := range *ruleSet.VirtualNetworkRules {
			ventID := i.VirtualNetworkResourceID
			vnetInstances = append(vnetInstances, storage.VirtualNetworkRule{
				VirtualNetworkResourceID: ventID,
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
func (_ *azureStorageManager) CreateStorage(ctx context.Context, instance *v1alpha1.Storage,
	groupName string,
	storageAccountName string,
	location string,
	sku azurev1alpha1.StorageSku,
	kind azurev1alpha1.StorageKind,
	tags map[string]*string,
	accessTier azurev1alpha1.StorageAccessTier,
	enableHTTPsTrafficOnly *bool, dataLakeEnabled *bool) (result storage.Account, err error) {

	storagesClient := getStoragesClient()

	//Check if name is available
	storageType := "Microsoft.Storage/storageAccounts"
	checkAccountParams := storage.AccountCheckNameAvailabilityParameters{Name: &storageAccountName, Type: &storageType}
	checkNameResult, err := storagesClient.CheckNameAvailability(ctx, checkAccountParams)
	if err != nil {
		return result, err
	}
	if dataLakeEnabled == to.BoolPtr(true) && kind != "StorageV2" {
		return result, errors.New("unable to create datalake enabled storage account")
	}
	if *checkNameResult.NameAvailable == false {
		if checkNameResult.Reason == storage.AccountNameInvalid {
			return result, errors.New("AccountNameInvalid")
		} else if checkNameResult.Reason == storage.AlreadyExists {
			return result, errors.New("AlreadyExists")
		}
	}

	sSku := storage.Sku{Name: storage.SkuName(sku.Name)}
	sKind := storage.Kind(kind)
	sAccessTier := storage.AccessTier(accessTier)

	var networkAcls storage.NetworkRuleSet
	if instance.Spec.NetworkRule != nil {
		networkAcls = ParseNetworkPolicy(instance.Spec.NetworkRule)
	} else {
		networkAcls = storage.NetworkRuleSet{}
	}

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
			NetworkRuleSet:         &networkAcls,
		},
	}

	//log.Println(fmt.Sprintf("creating storage '%s' in resource group '%s' and location: %v", storageAccountName, groupName, location))
	future, err := storagesClient.Create(ctx, groupName, storageAccountName, params)
	if err != nil {
		return result, err
	}

	return future.Result(storagesClient)

}

// Get gets the description of the specified storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
func (_ *azureStorageManager) GetStorage(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.GetProperties(ctx, resourceGroupName, accountName, "")
}

// DeleteStorage removes the resource group named by env var
func (_ *azureStorageManager) DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.Delete(ctx, groupName, storageAccountName)
}

func (_ *azureStorageManager) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountListKeysResult, err error) {
	storagesClient := getStoragesClient()
	return storagesClient.ListKeys(ctx, resourceGroupName, accountName, storage.Kerb)
}
