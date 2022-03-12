/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
)

func Test_Resources_ResourceGroup_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Create a resource group
	rg := tc.NewTestResourceGroup()
	tc.CreateResourceAndWait(rg)

	// check properties
	tc.Expect(rg.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(rg.Status.ProvisioningState).To(Equal(to.StringPtr("Succeeded")))
	tc.Expect(rg.Status.ID).ToNot(BeNil())
	armId := *rg.Status.ID

	// Update the tags
	old := rg.DeepCopy()
	rg.Spec.Tags["tag1"] = "value1"
	tc.PatchResourceAndWait(old, rg)
	tc.Expect(rg.Status.Tags).To(HaveKeyWithValue("tag1", "value1"))

	tc.DeleteResourceAndWait(rg)

	// Ensure that the resource group was really deleted in Azure
	// TODO: Do we want to just use an SDK here? This process is quite icky as is...
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		"2020-06-01")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
