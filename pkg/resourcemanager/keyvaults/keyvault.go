/*

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

package keyvaults

import (
	"context"
	"fmt"
	"log"

	auth "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/satori/go.uuid"
)

type azureKeyVaultManager struct{}

func getVaultsClient() keyvault.VaultsClient {
	vaultsClient := keyvault.NewVaultsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	vaultsClient.Authorizer = a
	vaultsClient.AddToUserAgent(config.UserAgent())
	return vaultsClient
}

func getObjectID(ctx context.Context, tenantID string, clientID string) *string {
	appclient := auth.NewApplicationsClient(tenantID)
	a, err := iam.GetGraphAuthorizer()
	if err != nil {
		return nil
	}
	appclient.Authorizer = a
	appclient.AddToUserAgent(config.UserAgent())

	result, err := appclient.GetServicePrincipalsIDByAppID(ctx, clientID)
	if err != nil {
		return nil
	}
	return result.Value
}

// CreateVault creates a new key vault
func (_ *azureKeyVaultManager) CreateVault(ctx context.Context, groupName string, vaultName string, location string) (keyvault.Vault, error) {
	vaultsClient := getVaultsClient()
	id, err := uuid.FromString(config.TenantID())
	if err != nil {
		return keyvault.Vault{}, err
	}

	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:       &id,
			AccessPolicies: &[]keyvault.AccessPolicyEntry{},
			Sku: &keyvault.Sku{
				Family: to.StringPtr("A"),
				Name:   keyvault.Standard,
			},
		},
		Location: to.StringPtr(location),
	}

	log.Println(fmt.Sprintf("creating keyvault '%s' in resource group '%s' and location: %v", vaultName, groupName, location))
	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)

	return future.Result(vaultsClient)
}

// CreateVaultWithAccessPolicies creates a new key vault and provides access policies to the specified user
func (_ *azureKeyVaultManager) CreateVaultWithAccessPolicies(ctx context.Context, groupName string, vaultName string, location string, clientID string) (keyvault.Vault, error) {
	vaultsClient := getVaultsClient()
	id, err := uuid.FromString(config.TenantID())
	if err != nil {
		return keyvault.Vault{}, err
	}

	apList := []keyvault.AccessPolicyEntry{}
	ap := keyvault.AccessPolicyEntry{
		TenantID: &id,
		Permissions: &keyvault.Permissions{
			Keys: &[]keyvault.KeyPermissions{
				keyvault.KeyPermissionsCreate,
			},
			Secrets: &[]keyvault.SecretPermissions{
				keyvault.SecretPermissionsSet,
				keyvault.SecretPermissionsGet,
				keyvault.SecretPermissionsDelete,
				keyvault.SecretPermissionsList,
			},
		},
	}
	if clientID != "" {
		if objID := getObjectID(ctx, config.TenantID(), clientID); objID != nil {
			ap.ObjectID = objID
			apList = append(apList, ap)
		}

	}

	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:       &id,
			AccessPolicies: &apList,
			Sku: &keyvault.Sku{
				Family: to.StringPtr("A"),
				Name:   keyvault.Standard,
			},
		},
		Location: to.StringPtr(location),
	}

	log.Println(fmt.Sprintf("creating keyvault '%s' in resource group '%s' and location: %v with access policies granted to %v", vaultName, groupName, location, clientID))
	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)
	if err != nil {
		return keyvault.Vault{}, err
	}

	return future.Result(vaultsClient)
}

// DeleteVault removes the resource group named by env var
func (_ *azureKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaultsClient := getVaultsClient()
	return vaultsClient.Delete(ctx, groupName, vaultName)
}

// CheckExistence checks for the presence of a keyvault instance on Azure
func (_ *azureKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaultsClient := getVaultsClient()
	return vaultsClient.Get(ctx, groupName, vaultName)

}
