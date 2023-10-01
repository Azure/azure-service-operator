/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401preview"
	"github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_KeyVault_WhenRecoverSpecified_RecoversSuccessfully(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create our Resource Group for testing
	rg := tc.CreateTestResourceGroupAndWait()
	defer tc.DeleteResourcesAndWait(rg)

	// Create our original KeyVault; this might just be a create, or it might be a recover, if this test
	// has been run live recently.
	tc.LogSectionf("Create original KeyVault with key")
	vault := newSoftDeletingKeyVault(tc, rg, "aso-kv-gs", to.Ptr(keyvault.VaultProperties_CreateMode_CreateOrRecover))
	tc.CreateResourceAndWait(vault)

	createKeyVaultKey(tc, vault, rg)

	tc.LogSectionf("Delete original KeyVault")
	tc.DeleteResourceAndWait(vault)

	// Create our replacement KeyVault by recovering the one we just deleted
	tc.LogSectionf("Recover original KeyVault")
	replacementVault := newSoftDeletingKeyVault(tc, rg, "aso-kv-gs", to.Ptr(keyvault.VaultProperties_CreateMode_CreateOrRecover))

	tc.CreateResourceAndWait(replacementVault)
}

func newSoftDeletingKeyVault(
	tc *testcommon.KubePerTestContext,
	rg *v1api20200601.ResourceGroup,
	name string,
	mode *keyvault.VaultProperties_CreateMode,
) *keyvault.Vault {
	return &keyvault.Vault{
		// We don't use a random name because we need to ensure the two keyvaults collide
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: keyvault.Vault_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &keyvault.VaultProperties{
				CreateMode: mode,
				AccessPolicies: []keyvault.AccessPolicyEntry{{
					ApplicationId: to.Ptr(tc.AzureTenant),
					ObjectId:      to.Ptr(tc.AzureTenant),
					Permissions: &keyvault.Permissions{
						Certificates: []keyvault.Permissions_Certificates{
							keyvault.Permissions_Certificates_Get,
						},
						Keys: []keyvault.Permissions_Keys{
							keyvault.Permissions_Keys_Get,
							keyvault.Permissions_Keys_List,
						},
						Secrets: []keyvault.Permissions_Secrets{
							keyvault.Permissions_Secrets_Get,
						},
						Storage: []keyvault.Permissions_Storage{
							keyvault.Permissions_Storage_Get,
						},
					},
					TenantId: to.Ptr(tc.AzureTenant),
				}},
				Sku: &keyvault.Sku{
					Family: to.Ptr(keyvault.Sku_Family_A),
					Name:   to.Ptr(keyvault.Sku_Name_Standard),
				},
				TenantId:         to.Ptr(tc.AzureTenant),
				EnableSoftDelete: to.Ptr(true),
			},
		},
	}
}
