/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	network "github.com/Azure/azure-service-operator/v2/api/microsoft.network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"
)

func Test_Networking_VirtualNetwork_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: testcommon.DefaultTestRegion,
			AddressSpace: network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/8"},
			},
		},
	}

	tc.CreateResourceAndWait(vnet)

	tc.Expect(vnet.Status.Id).ToNot(BeNil())
	armId := *vnet.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Subnet CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				Subnet_CRUD(testContext, vnet)
			},
		},
	)

	tc.DeleteResourceAndWait(vnet)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(network.VirtualNetworksSpecAPIVersion20201101))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Subnet_CRUD(tc testcommon.KubePerTestContext, vnet *network.VirtualNetwork) {
	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnets_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: "10.0.0.0/24",
		},
	}

	tc.CreateResourceAndWait(subnet)
	defer tc.DeleteResourceAndWait(subnet)

	tc.Expect(subnet.Status.Id).ToNot(BeNil())
	tc.Expect(subnet.Status.AddressPrefix).To(Equal(to.StringPtr("10.0.0.0/24")))

	// Update the subnet
	old := subnet.DeepCopy()
	subnet.Spec.Delegations = []network.VirtualNetworksSubnets_Spec_Properties_Delegations{
		{
			Name:        "mydelegation",
			ServiceName: to.StringPtr("Microsoft.DBforMySQL/serversv2"),
		},
	}
	tc.Patch(old, subnet)

	objectKey := client.ObjectKeyFromObject(subnet)

	// ensure state got updated in Azure
	tc.Eventually(func() []network.Delegation_Status {
		updated := &network.VirtualNetworksSubnet{}
		tc.GetResource(objectKey, updated)
		return updated.Status.Delegations
	}).Should(HaveLen(1))
}
