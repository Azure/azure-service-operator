/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"
	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220702"
	"github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401preview"
	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20210401preview"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Compute_DiskEncryptionSet_CRUD(t *testing.T) {

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	configMapName := "my-configmap"
	principalIdKey := "principalId"
	tenantIdKey := "tenantId"
	clientIdKey := "clientId"
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

	kv := newVaultForDiskEncryptionSet("kv", tc, rg)
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
			Keys: []keyvault.Permissions_Keys{
				keyvault.Permissions_Keys_Get,
				keyvault.Permissions_Keys_UnwrapKey,
				keyvault.Permissions_Keys_WrapKey,
			},
			Secrets: []keyvault.Permissions_Secrets{keyvault.Permissions_Secrets_Get},
			Storage: []keyvault.Permissions_Storage{keyvault.Permissions_Storage_Get},
		},
	}
	kv.Spec.Properties.AccessPolicies = []keyvault.AccessPolicyEntry{accessPolicyFromConfig}
	tc.CreateResourcesAndWait(mi, kv)
	key := createKeyVaultKey(tc, kv, rg)

	diskEncryptionSet := &compute.DiskEncryptionSet{
		ObjectMeta: tc.MakeObjectMeta("set"),
		Spec: compute.DiskEncryptionSet_Spec{
			ActiveKey: &compute.KeyForDiskEncryptionSet{
				KeyUrl: key.Properties.KeyURIWithVersion,
				SourceVault: &compute.SourceVault{
					Reference: tc.MakeReferenceFromResource(kv),
				},
			},
			EncryptionType: to.Ptr(compute.DiskEncryptionSetType_EncryptionAtRestWithCustomerKey),
			Identity: &compute.EncryptionSetIdentity{
				Type: to.Ptr(compute.EncryptionSetIdentity_Type_SystemAssigned),
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(diskEncryptionSet)
	defer tc.DeleteResourceAndWait(diskEncryptionSet)
}

func newVaultForDiskEncryptionSet(name string, tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *keyvault.Vault {

	return &keyvault.Vault{
		ObjectMeta: tc.MakeObjectMeta(name),
		Spec: keyvault.Vault_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &keyvault.VaultProperties{
				Sku: &keyvault.Sku{
					Family: to.Ptr(keyvault.Sku_Family_A),
					Name:   to.Ptr(keyvault.Sku_Name_Standard),
				},
				TenantId:                  to.Ptr(tc.AzureTenant),
				EnabledForDiskEncryption:  to.Ptr(true),
				SoftDeleteRetentionInDays: to.Ptr(7),
			},
		},
	}
}

func createKeyVaultKey(tc *testcommon.KubePerTestContext, kv *v1api20210401preview.Vault, rg *resources.ResourceGroup) armkeyvault.Key {
	client, err := armkeyvault.NewKeysClient(tc.AzureSubscription, tc.AzureClient.Creds(), nil)
	tc.Expect(err).To(BeNil())

	keyProperties := armkeyvault.KeyCreateParameters{
		Properties: &armkeyvault.KeyProperties{
			Attributes: &armkeyvault.KeyAttributes{
				Enabled: to.Ptr(true),
			},
			KeySize: to.Ptr(int32(2048)),
			Kty:     to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
		},
	}

	keyResp, err := client.CreateIfNotExist(context.TODO(), rg.Name, kv.Name, "testkey", keyProperties, nil)
	tc.Expect(err).To(BeNil())

	return keyResp.Key
}
