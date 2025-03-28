/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20240302"
	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Compute_DiskEncryptionSet_20240302_CRUD(t *testing.T) {
	t.Parallel()

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

	kv := newVaultForDiskEncryptionSet("kv2", tc, rg)
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

	tc.Expect(diskEncryptionSet.Status.Id).ToNot(BeNil())
	armId := *diskEncryptionSet.Status.Id

	tc.DeleteResourceAndWait(diskEncryptionSet)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(compute.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
