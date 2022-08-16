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

	skuFamily := keyvault.Sku_Family_A
	skuName := keyvault.Sku_Name_Standard

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
						Certificates: []keyvault.Permissions_Certificates{keyvault.Permissions_Certificates_Get},
						Keys:         []keyvault.Permissions_Keys{keyvault.Permissions_Keys_Get},
						Secrets:      []keyvault.Permissions_Secrets{keyvault.Permissions_Secrets_Get},
						Storage:      []keyvault.Permissions_Storage{keyvault.Permissions_Storage_Get},
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
