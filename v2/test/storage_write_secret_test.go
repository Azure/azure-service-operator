/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"

	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"

	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1beta20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// This test exists to ensure that we have the right RBAC permissions on secrets, which isn't covered by
// envtest.
func Test_StorageAccount_Secret_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	storageKeysSecret := "storagekeys"
	accessTier := storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot
<<<<<<< HEAD
	kind := storage.StorageAccount_Spec_Kind_StorageV2
	sku := storage.SkuName_Standard_LRS
=======
	kind := storage.StorageAccount_Kind_Spec_StorageV2
	sku := storage.Sku_Name_Standard_LRS
>>>>>>> main
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Sku: &storage.Sku{
				Name: &sku,
			},
			// TODO: They mark this property as optional but actually it is required
			AccessTier: &accessTier,
			OperatorSpec: &storage.StorageAccountOperatorSpec{
				Secrets: &storage.StorageAccountOperatorSecrets{
					Key1: &genruntime.SecretDestination{Name: storageKeysSecret, Key: "key1"},
				},
			},
		},
	}

	tc.CreateResourceAndWait(acct)

	tc.ExpectSecretHasKeys(storageKeysSecret, "key1")

	// Patch to add a new key to the secret
	old := acct.DeepCopy()
	acct.Spec.OperatorSpec.Secrets.Key2 = &genruntime.SecretDestination{Name: storageKeysSecret, Key: "key2"}
	tc.PatchResourceAndWait(old, acct)

	tc.ExpectSecretHasKeys(storageKeysSecret, "key1", "key2")

	// Delete the account and ensure the secret is now gone
	tc.DeleteResourceAndWait(acct)

	tc.ExpectResourceDoesNotExist(types.NamespacedName{Namespace: tc.Namespace, Name: storageKeysSecret}, &v1.Secret{})
}
