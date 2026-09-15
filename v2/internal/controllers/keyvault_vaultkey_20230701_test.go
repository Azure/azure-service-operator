/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701"
	vaultkey "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Test_KeyVault_VaultKey_20230701_CRUD exercises the full VaultKey lifecycle end-to-end via
// envtest: creating an RSA key owned by a Vault, confirming Ready, updating a mutable property
// (applied through the Key Vault data plane by the VaultKeyExtension), deleting the resource with
// deleteMode=delete (soft-deleting the key via the data plane), and finally the default detach
// behaviour, where deleting the CR leaves the key untouched in Azure.
//
// NOTE: this test requires a recorded HTTP cassette
// (v2/internal/controllers/recordings/Test_KeyVault_VaultKey_20230701_CRUD.yaml) to run in
// record/replay mode. The cassette must capture BOTH planes: the ARM interactions (including the
// CreateIfNotExist polling sequence) and the Key Vault data-plane traffic the VaultKeyExtension
// issues (GetKey/UpdateKey/DeleteKey, including the 401 challenge handshake the azkeys client
// performs). That cassette does not exist yet and this sandbox has no live Azure credentials to
// record one. Fabricating a synthetic cassette by hand would risk misrepresenting the true shape
// of the service's responses, so rather than do that, this test is left in place - fully written
// to the same structure/conventions as other tests in this file - and skipped with a clear TODO.
//
// TODO: record a real cassette for this test (requires a live Azure subscription) and remove the
// t.Skip below.
func Test_KeyVault_VaultKey_20230701_CRUD(t *testing.T) {
	t.Parallel()
	t.Skip("no recorded HTTP cassette available for this test in this environment (no live Azure " +
		"credentials to record one) - see comment on this test for details")

	tc := globalTestContext.ForTest(t)

	// Use a resource group scoped to this test only, so everything it contains can be torn down
	// via RG cascade-delete regardless of how the individual key deletions behave.
	rg := tc.CreateTestResourceGroupAndWait()

	vault := &keyvault.Vault{
		ObjectMeta: tc.MakeObjectMeta("vaultkeytest"),
		Spec: keyvault.Vault_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &keyvault.VaultProperties{
				CreateMode: to.Ptr(keyvault.VaultProperties_CreateMode_CreateOrRecover),
				Sku: &keyvault.Sku{
					Family: to.Ptr(keyvault.Sku_Family_A),
					Name:   to.Ptr(keyvault.Sku_Name_Standard),
				},
				TenantId:                  to.Ptr(tc.AzureTenant),
				EnableRbacAuthorization:   to.Ptr(true),
				SoftDeleteRetentionInDays: to.Ptr(7),
			},
		},
	}
	tc.CreateResourceAndWait(vault)

	// --- Create: ARM generates the key material ---
	key := &vaultkey.VaultKey{
		ObjectMeta: tc.MakeObjectMeta("rsakey"),
		Spec: vaultkey.VaultKey_Spec{
			Owner: testcommon.AsOwner(vault),
			OperatorSpec: &vaultkey.VaultKeyOperatorSpec{
				// Soft-delete the key in Azure when this resource is deleted, so this test
				// also covers the data-plane Deleter path.
				DeleteMode: to.Ptr("delete"),
			},
			Properties: &vaultkey.KeyProperties{
				Kty:     to.Ptr(vaultkey.KeyProperties_Kty_RSA),
				KeySize: to.Ptr(2048),
				Attributes: &vaultkey.KeyAttributes{
					Enabled:    to.Ptr(true),
					Exportable: to.Ptr(false),
				},
			},
		},
	}

	tc.CreateResourceAndWaitWithoutCleanup(key)
	tc.Expect(key.Status.Id).ToNot(BeNil())

	// --- Update: mutable properties are applied through the data plane ---
	//
	// Disabling the key is a data-plane UpdateKey performed by the VaultKeyExtension's
	// PostReconcileCheck; the resource reaches Ready at the new generation once the update has
	// been applied. (The refreshed attribute value becomes visible in status one reconcile
	// later, so this test does not assert on status here.)
	old := key.DeepCopy()
	key.Spec.Properties.Attributes.Enabled = to.Ptr(false)
	tc.PatchResourceAndWait(old, key)

	// --- Delete with deleteMode=delete: the key is soft-deleted via the data plane ---
	tc.DeleteResourceAndWait(key)

	// --- Default detach behaviour: deleting the CR leaves the key untouched in Azure ---
	ecKey := &vaultkey.VaultKey{
		ObjectMeta: tc.MakeObjectMeta("eckey"),
		Spec: vaultkey.VaultKey_Spec{
			Owner: testcommon.AsOwner(vault),
			Properties: &vaultkey.KeyProperties{
				Kty:       to.Ptr(vaultkey.KeyProperties_Kty_EC),
				CurveName: to.Ptr(vaultkey.KeyProperties_CurveName_P256),
			},
		},
	}

	tc.CreateResourceAndWaitWithoutCleanup(ecKey)
	tc.DeleteResourceAndWait(ecKey)
	// The CR is gone; with no deleteMode set the key itself remains live in the vault (detach is
	// the default and touches nothing in Azure). Azure-side state can't be asserted here without
	// an extra data-plane call; the RG teardown below cleans the key up.

	// Teardown: delete the Resource Group, cascade-deleting the Vault and any keys still in it
	// (including the detached EC key above).
	tc.DeleteResourceAndWait(rg)
}
