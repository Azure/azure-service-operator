/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_ReconcilePolicy_SkipReconcileAddedAlongWithTagsChange_ReconcileIsSkipped(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.Properties.ProvisioningState).To(Equal(to.Ptr("Succeeded")))
	tc.Expect(rg.Status.Id).ToNot(BeNil())

	// Update the tags but also skip reconcile
	old := rg.DeepCopy()
	rg.Spec.Tags["tag1"] = "value1"
	rg.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	tc.PatchResourceAndWait(old, rg)
	tc.Expect(rg.Status.Tags).ToNot(HaveKey("tag1"))

	// Stop skipping reconcile
	old = rg.DeepCopy()
	delete(rg.Annotations, "serviceoperator.azure.com/reconcile-policy")
	tc.Patch(old, rg)

	// ensure they get updated
	objectKey := client.ObjectKeyFromObject(rg)
	tc.Eventually(func() map[string]string {
		newRG := &resources.ResourceGroup{}
		tc.GetResource(objectKey, newRG)
		return newRG.Status.Tags
	}).Should(HaveKeyWithValue("tag1", "value1"))
}

func Test_ReconcilePolicy_UnknownPolicyIsIgnored(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.Properties.ProvisioningState).To(Equal(to.Ptr("Succeeded")))
	tc.Expect(rg.Status.Id).ToNot(BeNil())

	// Update the tags but also reconcile policy
	old := rg.DeepCopy()
	rg.Spec.Tags["tag1"] = "value1"
	rg.Annotations["serviceoperator.azure.com/reconcile-policy"] = "UNKNOWN"
	tc.PatchResourceAndWait(old, rg)
	tc.Expect(rg.Status.Tags).To(HaveKeyWithValue("tag1", "value1"))
}

func Test_ReconcilePolicy_DetachOnDelete_SkipsDelete(t *testing.T) {
	t.Parallel()
	testDeleteSkipped(t, "detach-on-delete")
}

func Test_ReconcilePolicy_Skip_SkipsDelete(t *testing.T) {
	t.Parallel()
	testDeleteSkipped(t, "skip")
}

func Test_ReconcilePolicy_SkippedParentDeleted_ChildIssuesDeleteToAzure(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	tc.Expect(rg.Status.Id).ToNot(BeNil())
	rgResourceId := *rg.Status.Id

	// Update the tags to skip reconcile
	old := rg.DeepCopy()
	rg.Spec.Tags["tag1"] = "value1"
	rg.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	tc.PatchResourceAndWait(old, rg)

	// Create a child resource in this RG
	acct := newStorageAccount(tc, rg)
	tc.CreateResourceAndWait(acct)

	tc.Expect(acct.Status.Id).ToNot(BeNil())
	resourceId := *acct.Status.Id

	// Delete the resource group that has reconcile-policy skip set
	defer func() {
		resp, err := tc.AzureClient.BeginDeleteByID(tc.Ctx, rgResourceId, rg.GetAPIVersion())
		tc.Expect(err).ToNot(HaveOccurred())
		_, err = resp.Poller.PollUntilDone(tc.Ctx, nil)
		tc.Expect(err).ToNot(HaveOccurred())
	}()
	tc.DeleteResourceAndWait(rg) // This doesn't do anything in Azure because of the skip policy

	tc.DeleteResourceAndWait(acct)
	// Ensure that the account was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		resourceId,
		string(storage.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Test_ReconcilePolicySkip_CreatesSecretAndConfigMap(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a storage account
	accessTier := storage.StorageAccountPropertiesCreateParameters_AccessTier_Hot
	kind := storage.StorageAccount_Kind_Spec_StorageV2
	sku := storage.SkuName_Standard_LRS
	acct := &storage.StorageAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("stor")),
		Spec: storage.StorageAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Sku: &storage.Sku{
				Name: &sku,
			},
			AccessTier: &accessTier,
		},
	}

	origAcct := acct.DeepCopy()
	tc.CreateResourceAndWait(acct)

	// Now annotate the resource with reconcile-policy: skip and delete it (which will leave it in Azure)
	old := acct.DeepCopy()
	acct.Annotations = make(map[string]string)
	acct.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	tc.Patch(old, acct)
	// Delete it
	tc.DeleteResourceAndWait(acct)

	// Create it again but this time adopt it and also set operatorSpec to export a configMap
	secretName := "storagekeys"
	secretKey := "key1"
	configMapName := "storageconfig"
	blobEndpointKey := "blobEndpoint"
	origAcct.Annotations = make(map[string]string)
	origAcct.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	origAcct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		ConfigMaps: &storage.StorageAccountOperatorConfigMaps{
			BlobEndpoint: &genruntime.ConfigMapDestination{
				Name: configMapName,
				Key:  blobEndpointKey,
			},
		},
		Secrets: &storage.StorageAccountOperatorSecrets{
			Key1: &genruntime.SecretDestination{
				Name: secretName,
				Key:  secretKey,
			},
		},
	}

	acct = origAcct.DeepCopy()

	tc.CreateResourcesAndWait(acct)

	// The configmap and secret should also be exported
	tc.ExpectConfigMapHasKeys(configMapName, blobEndpointKey)
	tc.ExpectSecretHasKeys(secretName, secretKey)

	// Now un-skip things so that deletion happens correctly
	old = acct.DeepCopy()
	delete(acct.Annotations, "serviceoperator.azure.com/reconcile-policy")
	tc.Patch(old, acct)
}

func testDeleteSkipped(t *testing.T, policy string) {
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWait(rg)

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.Properties.ProvisioningState).To(Equal(to.Ptr("Succeeded")))
	tc.Expect(rg.Status.Id).ToNot(BeNil())
	armId := *rg.Status.Id

	// Update to skip reconcile
	old := rg.DeepCopy()
	rg.Status.Conditions[0].ObservedGeneration = -1 // This is a hack so that we can tell when reconcile has happened to avoid a race
	tc.PatchStatus(old, rg)

	rg.Annotations["serviceoperator.azure.com/reconcile-policy"] = policy
	tc.Patch(old, rg)
	rv := rg.GetResourceVersion()
	print(rv)
	tc.Eventually(rg).Should(tc.Match.BeProvisioned(0))

	tc.DeleteResourceAndWait(rg)

	// Ensure that the resource group was NOT deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		"2020-06-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	// Create a fresh copy of the same RG - this adopts the existing RG
	newRG := &resources.ResourceGroup{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: rg.Namespace,
			Name:      rg.Name,
		},
		Spec: rg.Spec,
	}
	tc.CreateResourceAndWait(newRG)
	// Delete it
	tc.DeleteResourceAndWait(newRG)

	// Ensure that now the RG was deleted
	exists, _, err = tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		"2020-06-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
