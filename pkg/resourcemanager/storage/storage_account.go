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

package storage

import (
	"context"
	"fmt"
	s "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-04-01/storage"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"log"
)

func getStorageClient() s.AccountsClient {
	storageClient := s.NewAccountsClient(config.SubscriptionID())
	auth, _ := iam.GetResourceManagementAuthorizer()
	storageClient.Authorizer = auth
	storageClient.AddToUserAgent(config.UserAgent())
	return storageClient
}

// CreateStorageAccountAndWait creates a storage account
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
// accountKind - the Kind of the account
func CreateStorageAccountAndWait(ctx context.Context, resourceGroupName string, accountName string, accountKind s.Kind, location string) (*s.Account, error) {
	storageClient := getStorageClient()

	log.Println(fmt.Sprintf("Creating storage account '%s' for resource group: %s", accountName, resourceGroupName))

	accCreateFuture, err := storageClient.Create(
		ctx,
		resourceGroupName,
		accountName,
		s.AccountCreateParameters{
			Sku: &s.Sku{
				Name: s.StandardLRS, // TODO: de-hardcode this
			},
			Kind:     accountKind,
			Location: &location,
		})

	if err != nil {
		return nil, err
	}

	err = accCreateFuture.WaitForCompletionRef(ctx, storageClient.Client)
	if err != nil {
		return nil, err
	}

	result, err := accCreateFuture.Result(storageClient)

	return &result, err

}

// Get gets the description of the specified storage account.
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
func GetStorageAccount(ctx context.Context, resourceGroupName string, accountName string) (result s.Account, err error) {
	storageClient := getStorageClient()
	return storageClient.GetProperties(ctx, resourceGroupName, accountName, "")
}

// DeleteStorageAccount creates a storage account
// Parameters:
// resourceGroupName - name of the resource group within the azure subscription.
// accountName - the name of the storage account
func DeleteStorageAccount(ctx context.Context, resourceGroupName string, accountName string) (result autorest.Response, err error) {
	storageClient := getStorageClient()
	log.Println(fmt.Sprintf("Deleting storage account '%s' for resource group: %s", accountName, resourceGroupName))

	return storageClient.Delete(ctx,
		resourceGroupName,
		accountName)
}
