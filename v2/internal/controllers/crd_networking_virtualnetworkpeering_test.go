/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_VirtualNetworkPeering_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet1 := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/16"})
	vnet2 := newVNet(tc, testcommon.AsOwner(rg), []string{"10.1.0.0/16"})

	tc.CreateResourcesAndWait(vnet1, vnet2)

	peering := &network.VirtualNetworksVirtualNetworkPeering{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vgateway")),
		Spec: network.VirtualNetworks_VirtualNetworkPeering_Spec{
			Owner: testcommon.AsOwner(vnet1),
			RemoteVirtualNetwork: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(vnet2),
			},
		},
	}

	tc.CreateResourceAndWait(peering)

	tc.Expect(peering.Status.Id).ToNot(BeNil())
	armId := *peering.Status.Id
	tc.Expect(peering.Status.RemoteVirtualNetwork.Id).ToNot(BeNil())
	tc.Expect(*peering.Status.RemoteVirtualNetwork.Id).To(ContainSubstring(vnet2.AzureName()))

	// Update peering to enable traffic forwarding
	old := peering.DeepCopy()
	peering.Spec.AllowForwardedTraffic = to.BoolPtr(true)
	tc.PatchResourceAndWait(old, peering)
	tc.Expect(peering.Status.AllowForwardedTraffic).To(Equal(to.BoolPtr(true)))

	tc.DeleteResourceAndWait(peering)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Test_Networking_VirtualNetworkPeering_CreatedThenVNETUpdated_PeeringStillExists(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet1 := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/16"})
	vnet2 := newVNet(tc, testcommon.AsOwner(rg), []string{"10.1.0.0/16"})

	tc.CreateResourcesAndWait(vnet1, vnet2)

	peering := &network.VirtualNetworksVirtualNetworkPeering{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vgateway")),
		Spec: network.VirtualNetworks_VirtualNetworkPeering_Spec{
			Owner: testcommon.AsOwner(vnet1),
			RemoteVirtualNetwork: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(vnet2),
			},
		},
	}

	tc.CreateResourceAndWait(peering)
	tc.Expect(peering.Status.Id).ToNot(BeNil())
	armId := *peering.Status.Id

	// Now update the VNET
	old := vnet1.DeepCopy()
	vnet1.Spec.Tags = map[string]string{
		"taters": "boil 'em, mash 'em, stick 'em in a stew",
	}
	tc.PatchResourceAndWait(old, vnet1)

	// Now ensure that the VirtualNetworkPeering still exists
	exists, _, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())
}
