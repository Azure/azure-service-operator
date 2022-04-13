/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
)

func Test_ReconcilePolicy_SkipReconcileAddedAlongWithTagsChange_ReconcileIsSkipped(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.ProvisioningState).To(Equal(to.StringPtr("Succeeded")))
	tc.Expect(rg.Status.ID).ToNot(BeNil())

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
	tc.Expect(rg.Status.ProvisioningState).To(Equal(to.StringPtr("Succeeded")))
	tc.Expect(rg.Status.ID).ToNot(BeNil())

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

func testDeleteSkipped(t *testing.T, policy string) {
	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWait(rg)

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.ProvisioningState).To(Equal(to.StringPtr("Succeeded")))
	tc.Expect(rg.Status.ID).ToNot(BeNil())
	armId := *rg.Status.ID

	// Update to skip reconcile
	old := rg.DeepCopy()
	rg.Annotations["serviceoperator.azure.com/reconcile-policy"] = policy
	tc.Patch(old, rg)

	tc.DeleteResourceAndWait(rg)

	// Ensure that the resource group was NOT deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
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
	exists, _, err = tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		"2020-06-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
