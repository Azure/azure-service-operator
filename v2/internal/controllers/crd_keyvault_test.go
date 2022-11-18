/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1beta20210401preview"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
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

func Test_KeyVault_Vault_FromConfig_CRUD(t *testing.T) {

	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	defer tc.DeleteResourcesAndWait(rg)

	configMapName := "my-configmap"
	principalIdKey := "principalId"
	tenantIdKey := "tenantId"
	clientIdKey := "clientId"

	// Create a dummy managed identity which we will assign to a role
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					ClientId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  clientIdKey,
					},
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  principalIdKey,
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  tenantIdKey,
					},
				},
			},
		},
	}

	vault := newVault(tc, rg)

	accessPolicyFromConfig := keyvault.AccessPolicyEntry{
		ApplicationIdFromConfig: &genruntime.ConfigMapReference{
			Name: configMapName,
			Key:  clientIdKey,
		},
		ObjectIdFromConfig: &genruntime.ConfigMapReference{
			Name: configMapName,
			Key:  principalIdKey,
		},
		TenantIdFromConfig: &genruntime.ConfigMapReference{
			Name: configMapName,
			Key:  tenantIdKey,
		},
		Permissions: &keyvault.Permissions{
			Certificates: []keyvault.Permissions_Certificates{keyvault.Permissions_Certificates_Get},
			Keys:         []keyvault.Permissions_Keys{keyvault.Permissions_Keys_Get},
			Secrets:      []keyvault.Permissions_Secrets{keyvault.Permissions_Secrets_Get},
			Storage:      []keyvault.Permissions_Storage{keyvault.Permissions_Storage_Get},
		},
	}

	vault.Spec.Properties.AccessPolicies = append(vault.Spec.Properties.AccessPolicies, accessPolicyFromConfig)

	tc.CreateResourcesAndWait(mi, vault)
	tc.DeleteResourcesAndWait(vault, mi)

}

func newVault(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *keyvault.Vault {

	skuFamily := keyvault.Sku_Family_A
	skuName := keyvault.Sku_Name_Standard

	return &keyvault.Vault{
		ObjectMeta: tc.MakeObjectMeta("keyvault"),
		Spec: keyvault.Vault_Spec{
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
