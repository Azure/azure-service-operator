/*
Copyright 2019 microsoft.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package storages

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/go-autorest/autorest"
)

type StorageManager interface {
	CreateStorage(ctx context.Context, groupName string,
		storageAccountName string,
		location string,
		sku azurev1alpha1.StorageSku,
		kind azurev1alpha1.StorageKind,
		tags map[string]*string,
		accessTier azurev1alpha1.StorageAccessTier,
		enableHTTPsTrafficOnly *bool) (result storage.Account, err error)

	// Get gets the description of the specified storage account.
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// accountName - the name of the storage account
	GetStorage(ctx context.Context, resourceGroupName string, accountName string) (result storage.Account, err error)

	// DeleteStorage removes the storage account
	// Parameters:
	// resourceGroupName - name of the resource group within the azure subscription.
	// accountName - the name of the storage account
	DeleteStorage(ctx context.Context, groupName string, storageAccountName string) (result autorest.Response, err error)
}
