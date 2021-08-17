// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package storageaccount

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"k8s.io/apimachinery/pkg/runtime"
)

// NewManager returns an instance of the Storage Account Client
func NewManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *azureStorageManager {
	return &azureStorageManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

type StorageManager interface {
	CreateStorage(ctx context.Context,
		groupName string,
		storageAccountName string,
		location string,
		sku azurev1alpha1.StorageAccountSku,
		kind azurev1alpha1.StorageAccountKind,
		tags map[string]*string,
		accessTier azurev1alpha1.StorageAccountAccessTier,
		enableHTTPsTrafficOnly *bool, dataLakeEnabled *bool, networkRule *storage.NetworkRuleSet) (pollingURL string, result storage.Account, err error)

	// Get gets the description of the specified storage account.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// storageAccountName - the name of the storage account
	GetStorage(ctx context.Context, resourceGroupName string, storageAccountName string) (result storage.Account, err error)

	// DeleteStorage removes the storage account
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// storageAccountName - the name of the storage account
	DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error)

	ListKeys(ctx context.Context, groupName string, storageAccountName string) (result storage.AccountListKeysResult, err error)

	StoreSecrets(ctx context.Context,
		resourceGroupName string,
		accountName string,
		instance *v1alpha1.StorageAccount) error

	resourcemanager.ARMClient
}
