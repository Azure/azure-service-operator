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
	"github.com/Azure/azure-service-operator/v2/internal/testcommon/vcr"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Test_KeyVault_VaultKey_20230701_CRUD exercises the VaultKey lifecycle: creating an RSA key owned
// by a Vault, applying a mutable-property change through the Key Vault data plane, soft-deleting the
// key on resource deletion with deleteMode=delete, and the default detach behaviour, where deleting
// the resource leaves the key in Azure.
//
// The recording must capture both planes: the ARM interactions and the Key Vault data-plane traffic
// the VaultKeyExtension issues (including the azkeys 401 challenge handshake).
func Test_KeyVault_VaultKey_20230701_CRUD(t *testing.T) {
	t.Parallel()

	// TODO: remove once the recording has been added. Live runs (-live) are unaffected; to record, run
	// `go test ./internal/controllers -run Test_KeyVault_VaultKey_20230701_CRUD -args -live` with Azure
	// credentials and RECORD_REPLAY unset.
	if !*isLive {
		exists, err := vcr.CassetteFileExists("recordings/" + t.Name())
		if err != nil {
			t.Fatal(err)
		}
		if !exists {
			t.Skip("VaultKey recording pending: see the TODO above for how to record it")
		}
	}

	tc := globalTestContext.ForTest(t)

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

	// Create: ARM generates the key material
	key := &vaultkey.VaultKey{
		ObjectMeta: tc.MakeObjectMeta("rsakey"),
		Spec: vaultkey.VaultKey_Spec{
			Owner: testcommon.AsOwner(vault),
			OperatorSpec: &vaultkey.VaultKeyOperatorSpec{
				// Soft-delete the key in Azure when this resource is deleted, covering the data-plane Deleter path
				DeleteMode: to.Ptr("delete"),
			},
			Properties: &vaultkey.KeyProperties{
				Kty:     to.Ptr(vaultkey.KeyProperties_Kty_RSA),
				KeySize: to.Ptr(2048),
				Attributes: &vaultkey.KeyAttributes{
					Enabled:    to.Ptr(true),
					Exportable: to.Ptr(false),
				},
				RotationPolicy: &vaultkey.RotationPolicy{
					Attributes: &vaultkey.KeyRotationPolicyAttributes{ExpiryTime: to.Ptr("P2Y")},
					LifetimeActions: []vaultkey.LifetimeAction{
						{
							Action:  &vaultkey.Action{Type: to.Ptr(vaultkey.Action_Type_Rotate)},
							Trigger: &vaultkey.Trigger{TimeAfterCreate: to.Ptr("P90D")},
						},
						{
							// A non-default notify trigger, so the recording shows the service keeps it
							Action:  &vaultkey.Action{Type: to.Ptr(vaultkey.Action_Type_Notify)},
							Trigger: &vaultkey.Trigger{TimeBeforeExpiry: to.Ptr("P60D")},
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWaitWithoutCleanup(key)
	tc.Expect(key.Status.Id).ToNot(BeNil())
	armID := *key.Status.Id

	// Update: a mutable property is applied through the data plane by the VaultKeyExtension's
	// PostReconcileCheck before the resource reaches Ready at the new generation
	old := key.DeepCopy()
	key.Spec.Properties.Attributes.Enabled = to.Ptr(false)
	tc.PatchResourceAndWait(old, key)

	var live struct {
		Properties struct {
			Attributes struct {
				Enabled *bool `json:"enabled"`
			} `json:"attributes"`
		} `json:"properties"`
	}
	_, err := tc.AzureClient.GetByID(tc.Ctx, armID, string(vaultkey.APIVersion_Value), &live)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(live.Properties.Attributes.Enabled).To(HaveValue(BeFalse()))

	// Delete with deleteMode=delete: the key is soft-deleted via the data plane, so ARM no longer
	// finds it (the recording will confirm ARM answers 404 for a soft-deleted key)
	tc.DeleteResourceAndWait(key)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armID, string(vaultkey.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())

	// Default detach behaviour: deleting the resource leaves the key in Azure
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
	tc.Expect(ecKey.Status.Id).ToNot(BeNil())
	ecArmID := *ecKey.Status.Id

	tc.DeleteResourceAndWait(ecKey)

	exists, _, err = tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, ecArmID, string(vaultkey.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	// Teardown: delete the Resource Group, cascade-deleting the Vault and the detached key
	tc.DeleteResourceAndWait(rg)
}
