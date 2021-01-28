/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/k8s-infra/hack/generated/pkg/armclient"
)

func Test_ResourceGroup_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	// Create a resource group
	rg := testContext.NewTestResourceGroup()
	err = testContext.KubeClient.Create(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())

	// It should be created in Kubernetes
	g.Eventually(rg).Should(testContext.Match.BeProvisioned(ctx))

	g.Expect(rg.Status.Location).To(Equal(testContext.AzureRegion))
	g.Expect(rg.Status.Properties.ProvisioningState).To(Equal(string(armclient.SucceededProvisioningState)))
	g.Expect(rg.Status.ID).ToNot(BeNil())
	armId := rg.Status.ID

	// Delete the resource group
	err = testContext.KubeClient.Delete(ctx, rg)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(rg).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource group was really deleted in Azure
	// TODO: Do we want to just use an SDK here? This process is quite icky as is...
	exists, _, err := testContext.AzureClient.HeadResource(
		ctx,
		armId,
		"2020-06-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(exists).To(BeFalse())
}
