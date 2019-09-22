package keyvaults

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	pkghelpers "github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type keyVaultResource struct {
	resourceGroupName string
	vaultName         string
	KeyVault          keyvault.Vault
}

type MockKeyVaultManager struct {
	keyVaultResources []keyVaultResource
}

func findKeyVault(res []keyVaultResource, predicate func(keyVaultResource) bool) (int, keyVaultResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, keyVaultResource{}
}

// CreateVault creates a new key vault
func (manager *MockKeyVaultManager) CreateVault(ctx context.Context, groupName string, vaultName string, location string) (keyvault.Vault, error) {
	v := keyvault.Vault{
		Response:   helpers.GetRestResponse(200),
		Properties: &keyvault.VaultProperties{},
		ID:         to.StringPtr(pkghelpers.RandomString(10)),
		Name:       to.StringPtr(vaultName),
		Location:   to.StringPtr(location),
	}

	_, err := manager.GetVault(ctx, groupName, vaultName)
	if err != nil {
		manager.keyVaultResources = append(manager.keyVaultResources, keyVaultResource{
			resourceGroupName: groupName,
			vaultName:         vaultName,
			KeyVault:          v,
		})
	}

	return v, nil
}

// DeleteVault removes the resource group named by env var
func (manager *MockKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaults := manager.keyVaultResources

	index, _ := findKeyVault(vaults, func(g keyVaultResource) bool {
		return g.resourceGroupName == groupName &&
			g.vaultName == vaultName
	})

	if index == -1 {
		return helpers.GetRestResponse(404), errors.New("key vault not found")
	}

	manager.keyVaultResources = append(vaults[:index], vaults[index+1:]...)

	return helpers.GetRestResponse(200), nil
}

// Returns an existing keyvault instance
func (manager *MockKeyVaultManager) GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error) {
	vaults := manager.keyVaultResources

	index, v := findKeyVault(vaults, func(g keyVaultResource) bool {
		return g.resourceGroupName == groupName &&
			g.vaultName == vaultName
	})

	if index == -1 {
		return keyvault.Vault{
			Response: helpers.GetRestResponse(404),
		}, errors.New("key vault not found")
	}

	return v.KeyVault, nil
}
