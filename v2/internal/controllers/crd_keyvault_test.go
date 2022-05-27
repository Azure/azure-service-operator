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

	skuFamily := keyvault.Sku_FamilyA
	skuName := keyvault.Sku_NameStandard

	//TODO: This value here is a random generated string to comply with `^[0-9a-fA-F]{8}(-[0-9a-fA-F]{4}){3}-[0-9a-fA-F]{12}$` regex.
	// Tried random generator here which gets stored in the recordings and going to differ on each run resulting in test failure.
	// We kinda need some static value here
	str := "1C793267-c310-d4ae-7BD5-5Af5BEF875D3"

	return &keyvault.Vault{
		ObjectMeta: tc.MakeObjectMeta("keyvau"),
		Spec: keyvault.Vault_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &keyvault.VaultProperties{
				AccessPolicies: []keyvault.AccessPolicyEntry{{
					ApplicationId: to.StringPtr(str),
					ObjectId:      to.StringPtr(str),
					Permissions: &keyvault.Permissions{
						Certificates: []keyvault.Permissions_Certificates{keyvault.Permissions_CertificatesGet},
						Keys:         []keyvault.Permissions_Keys{keyvault.Permissions_KeysGet},
						Secrets:      []keyvault.Permissions_Secrets{keyvault.Permissions_SecretsGet},
						Storage:      []keyvault.Permissions_Storage{keyvault.Permissions_StorageGet},
					},
					TenantId: to.StringPtr(str),
				}},
				Sku: &keyvault.Sku{
					Family: &skuFamily,
					Name:   &skuName,
				},
				TenantId: to.StringPtr(str),
			},
		},
	}
}
