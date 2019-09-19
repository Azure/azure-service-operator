package keyvaults

import (
	"context"
	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
)

var AzureKeyVaultManager KeyVaultManager = &azureKeyVaultManager{}

type KeyVaultManager interface {
	CreateVault(ctx context.Context, groupName string, vaultName string, location string) (keyvault.Vault, error)

	// DeleteVault removes the resource group named by env var
	DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error)

	// CheckExistence checks for the presence of a keyvault instance on Azure
	GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error)
}
