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

	auth "github.com/Azure/azure-sdk-for-go/services/graphrbac/1.6/graphrbac"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	uuid "github.com/satori/go.uuid"
)

type azureKeyVaultManager struct{}

func getVaultsClient() (keyvault.VaultsClient, error) {
	client := keyvault.NewVaultsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		client = keyvault.VaultsClient{}
	} else {
		client.Authorizer = a
		client.AddToUserAgent(config.UserAgent())
	}
	return client, err
}

func getObjectID(ctx context.Context, tenantID string, clientID string) (*string, error) {
	appclient := auth.NewApplicationsClient(tenantID)
	a, err := iam.GetGraphAuthorizer()
	if err != nil {
		return nil, err
	}
	appclient.Authorizer = a
	appclient.AddToUserAgent(config.UserAgent())

	result, err := appclient.GetServicePrincipalsIDByAppID(ctx, clientID)
	var functionRet *string = nil
	if err == nil {
		functionRet = result.Value
	}
	return functionRet, err
}

// CreateVault creates a new key vault
func (_ *azureKeyVaultManager) CreateVault(ctx context.Context, groupName string, vaultName string, location string) (keyvault.Vault, error) {
	vaultsClient, err := getVaultsClient()
	if err != nil {
		return keyvault.Vault{}, err
	}
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

	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)

	return future.Result(vaultsClient)
}

// CreateVaultWithAccessPolicies creates a new key vault and provides access policies to the specified user
func (_ *azureKeyVaultManager) CreateVaultWithAccessPolicies(ctx context.Context, groupName string, vaultName string, location string, clientID string) (keyvault.Vault, error) {
	vaultsClient, err := getVaultsClient()
	if err != nil {
		return keyvault.Vault{}, err
	}
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
		objID, err := getObjectID(ctx, config.TenantID(), clientID)
		if err != nil {
			return keyvault.Vault{}, err
		}
		if objID != nil {
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

	future, err := vaultsClient.CreateOrUpdate(ctx, groupName, vaultName, params)
	if err != nil {
		return keyvault.Vault{}, err
	}

	return future.Result(vaultsClient)
}

// DeleteVault removes the resource group named by env var
func (_ *azureKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaultsClient, err := getVaultsClient()
	if err != nil {
		return autorest.Response{}, err
	}
	return vaultsClient.Delete(ctx, groupName, vaultName)
}

// CheckExistence checks for the presence of a keyvault instance on Azure
func (_ *azureKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaultsClient, err := getVaultsClient()
	if err != nil {
		return keyvault.Vault{}, err
	}
	return vaultsClient.Get(ctx, groupName, vaultName)

}
