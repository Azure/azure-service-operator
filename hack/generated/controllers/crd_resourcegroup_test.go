/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	resources "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/armclient"
)

func Test_ResourceGroup_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	kubeClient := testContext.KubeClient

	// Create a resource group
	rg := testContext.NewTestResourceGroup()
	g.Expect(kubeClient.Create(ctx, rg)).To(Succeed())

	// It should be created in Kubernetes
	g.Eventually(rg, remainingTime(t)).Should(testContext.Match.BeProvisioned(ctx))

	// check properties
	g.Expect(rg.Status.Location).To(Equal(testContext.AzureRegion))
	g.Expect(rg.Status.Properties.ProvisioningState).To(Equal(string(armclient.SucceededProvisioningState)))
	g.Expect(rg.Status.ID).ToNot(BeNil())
	armId := rg.Status.ID

	// Update the tags
	rg.Spec.Tags["tag1"] = "value1"
	g.Expect(kubeClient.Update(ctx, rg)).To(Succeed())

	objectKey, err := client.ObjectKeyFromObject(rg)
	g.Expect(err).ToNot(HaveOccurred())

	// ensure they get updated
	g.Eventually(func() map[string]string {
		newRG := &resources.ResourceGroup{}
		g.Expect(kubeClient.Get(ctx, objectKey, newRG)).To(Succeed())
		return newRG.Status.Tags
	}, remainingTime(t)).Should(HaveKeyWithValue("tag1", "value1"))

	// Delete the resource group
	g.Expect(kubeClient.Delete(ctx, rg)).To(Succeed())
	g.Eventually(rg, remainingTime(t)).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource group was really deleted in Azure
	// TODO: Do we want to just use an SDK here? This process is quite icky as is...
	exists, _, err := testContext.AzureClient.HeadResource(
		ctx,
		armId,
		"2020-06-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(exists).To(BeFalse())
}
