/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20171001"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_VirtualNetwork_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	rg, err := testContext.CreateNewTestResourceGroup(testcommon.WaitForCreation)
	g.Expect(err).ToNot(HaveOccurred())

	vn := &network.VirtualNetwork{
		ObjectMeta: testContext.MakeObjectMetaWithName(testContext.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Location: &testcommon.DefaultTestRegion,
			Properties: network.VirtualNetworkPropertiesFormat{
				AddressSpace: &network.AddressSpace{
					AddressPrefixes: []string{"10.0.0.0/8"},
				},
			},
		},
	}

	// Create
	g.Expect(testContext.KubeClient.Create(ctx, vn)).To(Succeed())
	g.Eventually(vn, remainingTime(t)).Should(testContext.Match.BeProvisioned(ctx))

	g.Expect(vn.Status.Id).ToNot(BeNil())
	armId := *vn.Status.Id

	// Delete
	g.Expect(testContext.KubeClient.Delete(ctx, vn)).To(Succeed())
	g.Eventually(vn, remainingTime(t)).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := testContext.AzureClient.HeadResource(ctx, armId, "2017-10-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
