/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701"
	vaultkey "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Test_KeyVault_VaultKey_20230701_CRUD exercises the full VaultKey lifecycle end-to-end via envtest:
// creating an RSA key owned by a Vault, confirming Ready, exercising the delete-blocking behavior
// under the default "manage" reconcile-policy (Microsoft.KeyVault/vaults/keys has no ARM DELETE
// operation, so the standard DeleteNotPossibleInAzure handling blocks finalizer removal while the
// key still exists), and then exercising the standard detach path
// (reconcile-policy=detach-on-delete), which allows the Kubernetes resource to be removed while
// leaving the key in Azure.
//
// NOTE: this test requires a recorded HTTP cassette
// (v2/internal/controllers/recordings/Test_KeyVault_VaultKey_20230701_CRUD.yaml) to run in
// record/replay mode against a real sequence of ARM interactions. That cassette does not exist yet and
// this sandbox has no live Azure credentials to record one (recording requires an actual subscription
// and issuing real requests once, which are then replayed on subsequent runs). Fabricating a synthetic
// cassette by hand would risk misrepresenting the true shape of ARM's responses (including the
// provisioningState polling sequence for CreateIfNotExist), so rather than do that, this test is left
// in place - fully written to the same structure/conventions as other tests in this file - and skipped
// with a clear TODO.
//
// TODO: record a real cassette for this test (requires a live Azure subscription) and remove the
// t.Skip below.
func Test_KeyVault_VaultKey_20230701_CRUD(t *testing.T) {
	t.Parallel()
	t.Skip("no recorded HTTP cassette available for this test in this environment (no live Azure " +
		"credentials to record one) - see comment on this test for details")

	tc := globalTestContext.ForTest(t)

	// Use a resource group scoped to this test only, so that the Vault (and any Keys within it) it
	// contains can be torn down via RG cascade-delete rather than relying on VaultKey's own delete
	// path (keys cannot be deleted through ARM - see below). This avoids orphaned keys accumulating
	// in a shared vault/RG across repeated test runs.
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

	key := &vaultkey.VaultKey{
		ObjectMeta: tc.MakeObjectMeta("rsakey"),
		Spec: vaultkey.VaultKey_Spec{
			Owner: testcommon.AsOwner(vault),
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

	// Create the VaultKey and confirm it reaches Ready.
	tc.CreateResourceAndWaitWithoutCleanup(key)
	tc.Expect(key.Status.Id).ToNot(BeNil())

	// --- Delete-blocking behavior under the default "manage" reconcile-policy ---
	//
	// Deleting the CR should NOT actually remove it: Microsoft.KeyVault/vaults/keys has no ARM DELETE
	// operation (VaultKey.GetSupportedOperations returns only Get and Put), so the reconciler's
	// standard DeleteNotPossibleInAzure handling blocks finalizer removal while the key still exists,
	// surfacing a DeletionNotSupportedInAzure condition in the meantime. No delete is ever attempted
	// against the key.
	tc.DeleteResource(key)

	objectKey := client.ObjectKeyFromObject(key)
	tc.Eventually(func() string {
		fresh := &vaultkey.VaultKey{}
		tc.GetResource(objectKey, fresh)
		for _, cond := range fresh.Status.Conditions {
			if cond.Type == "Ready" {
				return cond.Reason
			}
		}
		return ""
	}).Should(Equal("DeletionNotSupportedInAzure"), "deletion should be blocked with reason DeletionNotSupportedInAzure under the default manage policy")

	// The resource should still exist in Kubernetes (finalizer not removed).
	fresh := &vaultkey.VaultKey{}
	tc.GetResource(objectKey, fresh)
	tc.Expect(fresh.DeletionTimestamp).ToNot(BeNil())
	tc.Expect(controllerutil.ContainsFinalizer(fresh, genruntime.ReconcilerFinalizer)).To(BeTrue())

	// --- Detach path ---
	//
	// Setting the standard reconcile-policy=detach-on-delete annotation should allow the finalizer
	// to be removed / the CR to be deleted, while the underlying key remains in Azure.
	old := fresh.DeepCopy()
	if fresh.Annotations == nil {
		fresh.Annotations = map[string]string{}
	}
	fresh.Annotations[annotations.ReconcilePolicy] = string(annotations.ReconcilePolicyDetachOnDelete)
	tc.Patch(old, fresh)

	tc.DeleteResourceAndWait(fresh)

	// tc.DeleteResourceAndWait already asserts (via tc.Match.BeDeleted()) that the CR is gone from
	// Kubernetes. We cannot assert against real Azure state here without a live call against the
	// recorded cassette (the key is understood to remain in Azure per the design - it is never
	// deleted by this code path).

	// Teardown: delete the Resource Group (cascade-deletes the Vault and any Keys within it) rather
	// than relying on VaultKey's own delete capability, which never actually deletes the key in
	// Azure. This is the critical step to avoid orphaned-key accumulation across repeated
	// live/recorded test runs.
	tc.DeleteResourceAndWait(rg)
}
