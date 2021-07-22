/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"sigs.k8s.io/controller-runtime/pkg/client"

	. "github.com/onsi/gomega"

	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
)

func Test_ResourceGroup_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWait(rg)

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.ProvisioningState).To(Equal(string(armclient.SucceededProvisioningState)))
	tc.Expect(rg.Status.ID).ToNot(BeNil())
	armId := rg.Status.ID

	// Update the tags
	patcher := tc.NewResourcePatcher(rg)

	rg.Spec.Tags["tag1"] = "value1"
	patcher.Patch(rg)

	objectKey := client.ObjectKeyFromObject(rg)

	// ensure they get updated
	tc.Eventually(func() map[string]string {
		newRG := &resources.ResourceGroup{}
		tc.GetResource(objectKey, newRG)
		return newRG.Status.Tags
	}).Should(HaveKeyWithValue("tag1", "value1"))

	tc.DeleteResourceAndWait(rg)

	// Ensure that the resource group was really deleted in Azure
	// TODO: Do we want to just use an SDK here? This process is quite icky as is...
	exists, _, err := tc.AzureClient.HeadResource(
		tc.Ctx,
		armId,
		"2020-06-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
