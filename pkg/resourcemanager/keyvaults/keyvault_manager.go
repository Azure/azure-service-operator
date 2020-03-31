// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package keyvaults

import (
	"context"

	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

var AzureKeyVaultManager KeyVaultManager = &azureKeyVaultManager{}

type KeyVaultManager interface {
	CreateVault(ctx context.Context, instance *azurev1alpha1.KeyVault, sku azurev1alpha1.KeyVaultSku, tags map[string]*string) (keyvault.Vault, error)

	// CreateVault and grant access to the specific user ID
	CreateVaultWithAccessPolicies(ctx context.Context, groupName string, vaultName string, location string, userID string, sku azurev1alpha1.KeyVaultSku) (keyvault.Vault, error)

	// DeleteVault removes the resource group named by env var
	DeleteVault(ctx context.Context, groupName string, vaultName string) (result autorest.Response, err error)

	// CheckExistence checks for the presence of a keyvault instance on Azure
	GetVault(ctx context.Context, groupName string, vaultName string) (result keyvault.Vault, err error)

	// also embed async client methods
	resourcemanager.ARMClient
}
