// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storageaccount

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemgrconfig "github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

const templateForConnectionString = "DefaultEndpointsProtocol=https;AccountName=%s;AccountKey=%s;EndpointSuffix=%s"

type azureStorageManager struct {
	Creds        config.Credentials
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

func getStorageClient(creds config.Credentials) (storage.AccountsClient, error) {
	storageClient := storage.NewAccountsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return storage.AccountsClient{}, err
	}
	storageClient.Authorizer = a
	storageClient.AddToUserAgent(config.UserAgent())
	return storageClient, nil
}

// CreateStorage creates a new storage account
func (m *azureStorageManager) CreateStorage(ctx context.Context,
	groupName string,
	storageAccountName string,
	location string,
	sku azurev1alpha1.StorageAccountSku,
	kind azurev1alpha1.StorageAccountKind,
	tags map[string]*string,
	accessTier azurev1alpha1.StorageAccountAccessTier,
	enableHTTPsTrafficOnly *bool, dataLakeEnabled *bool, networkRule *storage.NetworkRuleSet) (pollingURL string, result storage.Account, err error) {

	storageClient, err := getStorageClient(m.Creds)
	if err != nil {
		return "", storage.Account{}, err
	}

	//Check if name is available
	storageType := "Microsoft.Storage/storageAccounts"
	checkAccountParams := storage.AccountCheckNameAvailabilityParameters{Name: &storageAccountName, Type: &storageType}
	checkNameResult, err := storageClient.CheckNameAvailability(ctx, checkAccountParams)
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

	future, err := storageClient.Create(ctx, groupName, storageAccountName, params)
	if err != nil {
		return "", result, err
	}

	result, err = future.Result(storageClient)

	return future.PollingURL(), result, err

}

// Get gets the description of the specified storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// storageAccountName - the name of the storage account
func (m *azureStorageManager) GetStorage(ctx context.Context, resourceGroupName string, storageAccountName string) (result storage.Account, err error) {
	storageClient, err := getStorageClient(m.Creds)
	if err != nil {
		return storage.Account{}, err
	}

	return storageClient.GetProperties(ctx, resourceGroupName, storageAccountName, "")
}

// DeleteStorage removes the resource group named by env var
func (m *azureStorageManager) DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error) {
	storageClient, err := getStorageClient(m.Creds)
	if err != nil {
		return autorest.Response{
			Response: &http.Response{
				StatusCode: 500,
			},
		}, err
	}

	return storageClient.Delete(ctx, groupName, storageAccountName)
}

func (m *azureStorageManager) ListKeys(ctx context.Context, resourceGroupName string, accountName string) (result storage.AccountListKeysResult, err error) {
	storageClient, err := getStorageClient(m.Creds)
	if err != nil {
		return storage.AccountListKeysResult{}, err
	}

	return storageClient.ListKeys(ctx, resourceGroupName, accountName, storage.Kerb)
}

// StoreSecrets upserts the secret information for this storage account
func (m *azureStorageManager) StoreSecrets(ctx context.Context, resourceGroupName string, accountName string, instance *v1alpha1.StorageAccount) error {

	// get the keys
	keyResult, err := m.ListKeys(ctx, resourceGroupName, accountName)
	if err != nil {
		return err
	}
	if keyResult.Keys == nil {
		return errors.New("No keys were returned from ListKeys")
	}
	keys := *keyResult.Keys
	storageEndpointSuffix := resourcemgrconfig.Environment().StorageEndpointSuffix

	// build the connection string
	data := map[string][]byte{
		"StorageAccountName": []byte(accountName),
	}
	for i, key := range keys {
		data[fmt.Sprintf("connectionString%d", i)] = []byte(fmt.Sprintf(templateForConnectionString, accountName, *key.Value, storageEndpointSuffix))
		data[fmt.Sprintf("key%d", i)] = []byte(*key.Value)
	}

	// upsert
	secretKey := m.makeSecretKey(instance)
	return m.SecretClient.Upsert(ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(m.Scheme),
	)
}

func (m *azureStorageManager) makeSecretKey(instance *v1alpha1.StorageAccount) secrets.SecretKey {
	if m.SecretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		return secrets.SecretKey{
			Name:      fmt.Sprintf("storageaccount-%s-%s", instance.Spec.ResourceGroup, instance.Name),
			Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind,
		}
	}
	return secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
}
