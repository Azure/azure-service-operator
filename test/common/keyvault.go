/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package common

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/keyvault/mgmt/2018-02-14/keyvault"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/gofrs/uuid"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	kvhelper "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
)

func CreateKeyVaultSoftDeleteEnabled(ctx context.Context, creds config.Credentials, resourceGroupName string, vaultName string, location string, objectID *string) error {
	vaultsClient, err := kvhelper.GetKeyVaultClient(creds)
	if err != nil {
		return errors.Wrapf(err, "couldn't get vaults client")
	}

	id, err := uuid.FromString(creds.TenantID())
	if err != nil {
		return errors.Wrapf(err, "couldn't convert tenantID to UUID")
	}

	accessPolicies, err := CreateKeyVaultTestAccessPolicies(creds, objectID)
	if err != nil {
		return nil
	}

	params := keyvault.VaultCreateOrUpdateParameters{
		Properties: &keyvault.VaultProperties{
			TenantID:         &id,
			AccessPolicies:   &accessPolicies,
			EnableSoftDelete: to.BoolPtr(true),
			Sku: &keyvault.Sku{
				Family: to.StringPtr("A"),
				Name:   keyvault.Standard,
			},
		},
		Location: to.StringPtr(location),
	}

	future, err := vaultsClient.CreateOrUpdate(ctx, resourceGroupName, vaultName, params)
	if err != nil {
		return err
	}

	return future.WaitForCompletionRef(ctx, vaultsClient.Client)
}

// CreateVaultWithAccessPolicies creates a new key vault and provides access policies to the specified user - used in test
func CreateVaultWithAccessPolicies(ctx context.Context, creds config.Credentials, groupName string, vaultName string, location string, objectID *string) error {
	vaultsClient, err := kvhelper.GetKeyVaultClient(creds)
	if err != nil {
		return errors.Wrapf(err, "couldn't get vaults client")
	}
	id, err := uuid.FromString(creds.TenantID())
	if err != nil {
		return errors.Wrapf(err, "couldn't convert tenantID to UUID")
	}

	apList, err := CreateKeyVaultTestAccessPolicies(creds, objectID)
	if err != nil {
		return nil
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
		return err
	}

	return future.WaitForCompletionRef(ctx, vaultsClient.Client)
}

func CreateKeyVaultTestAccessPolicies(creds config.Credentials, objectID *string) ([]keyvault.AccessPolicyEntry, error) {
	id, err := uuid.FromString(creds.TenantID())
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't convert tenantID to UUID")
	}

	apList := []keyvault.AccessPolicyEntry{
		{
			TenantID: &id,
			ObjectID: objectID,
			Permissions: &keyvault.Permissions{
				Keys: &[]keyvault.KeyPermissions{
					keyvault.KeyPermissionsCreate,
				},
				Secrets: &[]keyvault.SecretPermissions{
					keyvault.SecretPermissionsSet,
					keyvault.SecretPermissionsGet,
					keyvault.SecretPermissionsDelete,
					keyvault.SecretPermissionsList,
					keyvault.SecretPermissionsRecover,
				},
			},
		},
	}

	return apList, nil
}
