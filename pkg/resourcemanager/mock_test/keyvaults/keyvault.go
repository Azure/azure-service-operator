package keyvaults

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mock_test/helpers"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type KeyVaultResource struct {
	ResourceGroupName string
	VaultName     string
	KeyVault       keyvault.Vault
}

type MockKeyVaultManager struct {
	keyVaultResources []KeyVaultResource
}

func findKeyVault(res []KeyVaultResource, predicate func(KeyVaultResource) bool) (int, KeyVaultResource) {
	for index, r := range res {
		if predicate(r) {
			return index, r
		}
	}
	return -1, KeyVaultResource{}
}


// CreateVault creates a new key vault
func (manager *MockKeyVaultManager) CreateVault(ctx context.Context, groupName string, vaultName string, location string) (keyvault.Vault, error) {
	v := keyvault.Vault{
		Response:   helpers.GetRestResponse(200),
		Properties: &keyvault.VaultProperties{},
		Name:       to.StringPtr(vaultName),
		Location:   to.StringPtr(location),
	}
	manager.keyVaultResources = append(manager.keyVaultResources, KeyVaultResource{
		ResourceGroupName: groupName,
		VaultName:         vaultName,
		KeyVault:          v,
	})

	return v, nil
}

// DeleteVault removes the resource group named by env var
func (manager *MockKeyVaultManager) DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error) {
	vaults := manager.keyVaultResources

	index, _ := findKeyVault(vaults, func(g KeyVaultResource) bool {
		return g.ResourceGroupName == groupName &&
			g.VaultName == vaultName
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

	index, v := findKeyVault(vaults, func(g KeyVaultResource) bool {
		return g.ResourceGroupName == groupName &&
			g.VaultName == vaultName
	})

	if index == -1 {
		return keyvault.Vault{
			Response: helpers.GetRestResponse(404),
		}, errors.New("key vault not found")
	}

	return v.KeyVault, nil
}
