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
	vaultskey "github.com/Azure/azure-service-operator/v2/api/keyvault/v20230701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Test_KeyVault_VaultsKey_20230701_CRUD exercises the full VaultsKey lifecycle end-to-end via envtest:
// creating an RSA key owned by a Vault, confirming Ready, exercising the delete-blocking behavior
// under the default "manage" reconcile-policy (see VaultsKeyExtension.Delete and the reconciler guard
// clause in generic_reconciler.go), and then exercising the detach path
// (reconcile-policy=detach-on-delete + the vaultskey.aso.io/acknowledge-key-retained-enabled
// annotation), which does allow the Kubernetes resource to be removed while leaving the key enabled
// in Azure.
//
// NOTE: this test requires a recorded HTTP cassette
// (v2/internal/controllers/recordings/Test_KeyVault_VaultsKey_20230701_CRUD.yaml) to run in
// record/replay mode against a real sequence of ARM interactions. That cassette does not exist yet and
// this sandbox has no live Azure credentials to record one (recording requires an actual subscription
// and issuing real requests once, which are then replayed on subsequent runs). Fabricating a synthetic
// cassette by hand would risk misrepresenting the true shape of ARM's responses (including the
// provisioningState polling sequence for CreateIfNotExist), so rather than do that, this test is left
// in place - fully written to the same structure/conventions as other tests in this file - and skipped
// with a clear TODO. The lower-level behaviors this test exercises (VaultsKeyExtension.Delete,
// RequireDetachAcknowledgement, and the reconciler guard clause) are already covered by fast,
// deterministic unit tests using a fake ARM HTTP client, see:
//   - v2/api/keyvault/customizations/vaults_key_extension_test.go
//   - v2/internal/reconcilers/generic/generic_reconciler_test.go
//
// TODO: record a real cassette for this test (requires a live Azure subscription) and remove the
// t.Skip below.
func Test_KeyVault_VaultsKey_20230701_CRUD(t *testing.T) {
	t.Parallel()
	t.Skip("no recorded HTTP cassette available for this test in this environment (no live Azure " +
		"credentials to record one) - see comment on this test for details")

	tc := globalTestContext.ForTest(t)

	// Use a resource group scoped to this test only, so that the Vault (and any Keys within it) it
	// contains can be torn down via RG cascade-delete rather than relying on VaultsKey's own delete
	// path (which, under the default policy, never actually deletes the key - see below). This avoids
	// orphaned keys accumulating in a shared vault/RG across repeated test runs.
	rg := tc.CreateTestResourceGroupAndWait()

	vault := &keyvault.Vault{
		ObjectMeta: tc.MakeObjectMeta("vaultskeytest"),
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

	key := &vaultskey.VaultsKey{
		ObjectMeta: tc.MakeObjectMeta("rsakey"),
		Spec: vaultskey.VaultsKey_Spec{
			Owner: testcommon.AsOwner(vault),
			Properties: &vaultskey.KeyProperties{
				Kty:     to.Ptr(vaultskey.KeyProperties_Kty_RSA),
				KeySize: to.Ptr(2048),
				Attributes: &vaultskey.KeyAttributes{
					Enabled:    to.Ptr(true),
					Exportable: to.Ptr(false),
				},
			},
		},
	}

	// Create the VaultsKey and confirm it reaches Ready.
	tc.CreateResourceAndWaitWithoutCleanup(key)
	tc.Expect(key.Status.Id).ToNot(BeNil())

	// --- Delete-blocking behavior under the default "manage" reconcile-policy ---
	//
	// Deleting the CR should NOT actually remove it: VaultsKeyExtension.Delete disables the key in
	// Azure but always blocks finalizer removal (there is no ARM DELETE for keys), surfacing a
	// KeyDeletionBlocked condition/event in the meantime.
	tc.DeleteResource(key)

	objectKey := client.ObjectKeyFromObject(key)
	tc.Eventually(func() string {
		fresh := &vaultskey.VaultsKey{}
		tc.GetResource(objectKey, fresh)
		for _, cond := range fresh.Status.Conditions {
			if cond.Type == "Ready" {
				return cond.Reason
			}
		}
		return ""
	}).Should(Equal("KeyDeletionBlocked"), "deletion should be blocked with reason KeyDeletionBlocked under the default manage policy")

	// The resource should still exist in Kubernetes (finalizer not removed).
	fresh := &vaultskey.VaultsKey{}
	tc.GetResource(objectKey, fresh)
	tc.Expect(fresh.DeletionTimestamp).ToNot(BeNil())
	tc.Expect(controllerutil.ContainsFinalizer(fresh, genruntime.ReconcilerFinalizer)).To(BeTrue())

	// --- Detach path ---
	//
	// Setting both the namespace/object reconcile-policy=detach-on-delete AND the VaultsKey-specific
	// acknowledgment annotation should allow the finalizer to be removed / the CR to be deleted, while
	// the underlying key remains enabled and live in Azure (this is the entire point of the
	// acknowledgment requirement - see RequireDetachAcknowledgement).
	old := fresh.DeepCopy()
	if fresh.Annotations == nil {
		fresh.Annotations = map[string]string{}
	}
	fresh.Annotations[annotations.ReconcilePolicy] = string(annotations.ReconcilePolicyDetachOnDelete)
	fresh.Annotations["vaultskey.aso.io/acknowledge-key-retained-enabled"] = "true"
	tc.Patch(old, fresh)

	tc.DeleteResourceAndWait(fresh)

	// tc.DeleteResourceAndWait already asserts (via tc.Match.BeDeleted()) that the CR is gone from
	// Kubernetes. We cannot assert against real Azure state here without a live call against the
	// recorded cassette (the key is understood to remain enabled in Azure per the design - it is
	// never deleted by this code path).

	// Teardown: delete the Resource Group (cascade-deletes the Vault and any Keys within it) rather
	// than relying on VaultsKey's own delete capability, which - under the default policy - never
	// actually deletes the key in Azure. This is the critical step to avoid orphaned-key accumulation
	// across repeated live/recorded test runs.
	tc.DeleteResourceAndWait(rg)
}
