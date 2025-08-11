/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	corev1 "k8s.io/api/core/v1"
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

func Test_ReconcilePolicy_SkipsDelete(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		policy      string
		onNamespace bool
	}{
		{
			name:        "DetachOnDelete_OnResource",
			policy:      "detach-on-delete",
			onNamespace: false,
		},
		{
			name:        "Skip_OnResource",
			policy:      "skip",
			onNamespace: false,
		},
		{
			name:        "DetachOnDelete_OnNamespace",
			policy:      "detach-on-delete",
			onNamespace: true,
		},
		{
			name:        "Skip_OnNamespace",
			policy:      "skip",
			onNamespace: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			testDeleteSkipped(
				t,
				deleteSkippedOptions{
					policy:      tc.policy,
					onNamespace: tc.onNamespace,
				})
		})
	}
}

func Test_ReconcilePolicy_SkipReconcileAddedToNamespace_ReconcileIsSkipped(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.Properties.ProvisioningState).To(Equal(to.Ptr("Succeeded")))
	tc.Expect(rg.Status.Id).ToNot(BeNil())

	// Update the namespace to skip reconcile
	ns := &corev1.Namespace{}
	tc.GetResource(client.ObjectKey{Name: tc.Namespace}, ns)
	oldNS := ns.DeepCopy()
	ns.Annotations = make(map[string]string)
	ns.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	tc.Patch(oldNS, ns)

	// Update the tags
	old := rg.DeepCopy()
	rg.Spec.Tags["tag1"] = "value1"
	tc.PatchResourceAndWait(old, rg)
	tc.Expect(rg.Status.Tags).ToNot(HaveKey("tag1"))

	// Stop skipping reconcile
	oldNS = ns.DeepCopy()
	delete(ns.Annotations, "serviceoperator.azure.com/reconcile-policy")
	tc.Patch(oldNS, ns)

	// ensure the tags get updated
	objectKey := client.ObjectKeyFromObject(rg)
	tc.Eventually(func() map[string]string {
		newRG := &resources.ResourceGroup{}
		tc.GetResource(objectKey, newRG)
		return newRG.Status.Tags
	}).Should(HaveKeyWithValue("tag1", "value1"))
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

type deleteSkippedOptions struct {
	policy      string
	onNamespace bool
}

func testDeleteSkipped(t *testing.T, opts deleteSkippedOptions) {
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWait(rg)

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.Properties.ProvisioningState).To(Equal(to.Ptr("Succeeded")))
	tc.Expect(rg.Status.Id).ToNot(BeNil())
	armId := *rg.Status.Id

	if !opts.onNamespace {
		if opts.policy == "skip" {
			// This is a hack so that we can tell when reconcile has happened to avoid a race in the recording.
			// It's only required in the skip case because for detach-on-delete we don't reconcile at all
			// See HasReconcilePolicyAnnotationChanged
			old := rg.DeepCopy()
			rg.Status.Conditions[0].ObservedGeneration = -1
			tc.PatchStatus(old, rg)
		}

		// Update to skip reconcile
		old := rg.DeepCopy()
		rg.Annotations["serviceoperator.azure.com/reconcile-policy"] = opts.policy
		tc.Patch(old, rg)
		// TODO: There may still be a race here where delete in the operator code gets the old cached value without the policy in the detach-on-delete case,
		// TODO: since there is no reconcile after the annotation patch (see HasReconcilePolicyAnnotationChanged), so this eventually finishes quickly.
		tc.Eventually(rg).Should(tc.Match.BeProvisioned(-1))
	} else {
		// Update the namespace to skip reconcile
		ns := &corev1.Namespace{}
		tc.GetResource(client.ObjectKey{Name: tc.Namespace}, ns)
		oldNS := ns.DeepCopy()
		ns.Annotations = make(map[string]string)
		ns.Annotations["serviceoperator.azure.com/reconcile-policy"] = opts.policy
		tc.Patch(oldNS, ns)
	}

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

	if opts.onNamespace {
		// Remove skip reconcile annotation from the namespace
		ns := &corev1.Namespace{}
		tc.GetResource(client.ObjectKey{Name: tc.Namespace}, ns)
		oldNS := ns.DeepCopy()
		delete(ns.Annotations, "serviceoperator.azure.com/reconcile-policy")
		tc.Patch(oldNS, ns)
	}

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
