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

	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20200501"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/util/patch"
)

func Test_PublicIP_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	rg, err := testContext.CreateNewTestResourceGroup(testcommon.WaitForCreation)
	g.Expect(err).ToNot(HaveOccurred())

	// Public IP Address
	// TODO: Note the microsoft.networking package also defines a PublicIPAddress type, so
	// TODO: depluralization of this resource doesn't work because of the collision.
	sku := network.PublicIPAddressSkuNameStandard
	publicIPAddress := &network.PublicIPAddresses{
		ObjectMeta: testContext.MakeObjectMetaWithName(testContext.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddresses_Spec{
			Location: testContext.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			Properties: network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic,
			},
		},
	}

	err = testContext.KubeClient.Create(ctx, publicIPAddress)
	g.Expect(err).ToNot(HaveOccurred())

	// It should be created in Kubernetes
	g.Eventually(publicIPAddress).Should(testContext.Match.BeProvisioned(ctx))
	g.Expect(publicIPAddress.Status.Id).ToNot(BeNil())
	armId := *publicIPAddress.Status.Id

	// Perform a simple patch
	patchHelper, err := patch.NewHelper(publicIPAddress, testContext.KubeClient)
	g.Expect(err).ToNot(HaveOccurred())

	idleTimeoutInMinutes := 7
	publicIPAddress.Spec.Properties.IdleTimeoutInMinutes = &idleTimeoutInMinutes
	err = patchHelper.Patch(ctx, publicIPAddress)
	g.Expect(err).ToNot(HaveOccurred())

	objectKey, err := client.ObjectKeyFromObject(publicIPAddress)
	g.Expect(err).ToNot(HaveOccurred())

	// ensure state got updated in Azure
	g.Eventually(func() *int {
		updatedIP := &network.PublicIPAddresses{}
		g.Expect(testContext.KubeClient.Get(ctx, objectKey, updatedIP)).To(Succeed())
		return updatedIP.Status.Properties.IdleTimeoutInMinutes
	}, remainingTime(t)).Should(Equal(&idleTimeoutInMinutes))

	// Delete
	err = testContext.KubeClient.Delete(ctx, publicIPAddress)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(publicIPAddress).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := testContext.AzureClient.HeadResource(ctx, armId, "2020-05-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
