/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_VirtualNetworkPeering_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet1 := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}

	vnet2 := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.1.0.0/16"},
			},
		},
	}

	tc.CreateResourcesAndWait(vnet1, vnet2)

	peering := &network.VirtualNetworksVirtualNetworkPeering{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vgateway")),
		Spec: network.VirtualNetworksVirtualNetworkPeerings_Spec{
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
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.VirtualNetworksVirtualNetworkPeeringsSpecAPIVersion20201101))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
