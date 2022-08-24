/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1beta20210401preview"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

// If recording this test, might need to manually purge the old KeyVault: az keyvault purge --name asotest-keyvault-ngmgjs

func Test_KeyVault_Vault_CRUD(t *testing.T) {

	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	defer tc.DeleteResourcesAndWait(rg)

	vault := newVault(tc, rg)

	tc.CreateResourceAndWait(vault)
	tc.DeleteResourceAndWait(vault)

}

func newVault(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *keyvault.Vault {

	skuFamily := keyvault.SkuFamily_A
	skuName := keyvault.SkuName_Standard

	return &keyvault.Vault{
		ObjectMeta: tc.MakeObjectMeta("keyvault"),
		Spec: keyvault.Vaults_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &keyvault.VaultProperties{
				AccessPolicies: []keyvault.AccessPolicyEntry{{
					ApplicationId: to.StringPtr(tc.AzureTenant),
					ObjectId:      to.StringPtr(tc.AzureTenant),
					Permissions: &keyvault.Permissions{
						Certificates: []keyvault.PermissionsCertificates{keyvault.PermissionsCertificates_Get},
						Keys:         []keyvault.PermissionsKeys{keyvault.PermissionsKeys_Get},
						Secrets:      []keyvault.PermissionsSecrets{keyvault.PermissionsSecrets_Get},
						Storage:      []keyvault.PermissionsStorage{keyvault.PermissionsStorage_Get},
					},
					TenantId: to.StringPtr(tc.AzureTenant),
				}},
				Sku: &keyvault.Sku{
					Family: &skuFamily,
					Name:   &skuName,
				},
				TenantId: to.StringPtr(tc.AzureTenant),
			},
		},
	}
}
