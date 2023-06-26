/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_VirtualNetwork_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/8"})

	tc.CreateResourceAndWait(vnet)

	tc.Expect(vnet.Status.Id).ToNot(BeNil())
	armId := *vnet.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Subnet CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Subnet_CRUD(tc, vnet)
			},
		},
	)

	tc.DeleteResourceAndWait(vnet)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Subnet_CRUD(tc *testcommon.KubePerTestContext, vnet *network.VirtualNetwork) {
	subnet := newSubnet(tc, vnet, "10.0.0.0/24")

	tc.CreateResourceAndWait(subnet)
	defer tc.DeleteResourceAndWait(subnet)

	tc.Expect(subnet.Status.Id).ToNot(BeNil())
	tc.Expect(subnet.Status.AddressPrefix).To(Equal(to.Ptr("10.0.0.0/24")))

	// Update the subnet
	old := subnet.DeepCopy()
	subnet.Spec.Delegations = []network.Delegation{
		{
			Name:        to.Ptr("mydelegation"),
			ServiceName: to.Ptr("Microsoft.DBforMySQL/serversv2"),
		},
	}
	tc.PatchResourceAndWait(old, subnet)
	tc.Expect(subnet.Status.Delegations).To(HaveLen(1))
}

func newSubnet(tc *testcommon.KubePerTestContext, vnet *network.VirtualNetwork, addressPrefix string) *network.VirtualNetworksSubnet {
	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworks_Subnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr(addressPrefix),
		},
	}
	return subnet
}

func Test_Networking_Subnet_CreatedThenVNETUpdated_SubnetStillExists(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/16"})

	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworks_Subnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}

	tc.CreateResourceAndWait(vnet)
	tc.CreateResourceAndWait(subnet)
	tc.Expect(subnet.Status.Id).ToNot(BeNil())
	armId := *subnet.Status.Id

	// Now update the VNET
	old := vnet.DeepCopy()
	vnet.Spec.Tags = map[string]string{
		"taters": "boil 'em, mash 'em, stick 'em in a stew",
	}
	tc.PatchResourceAndWait(old, vnet)

	// Now ensure that the Subnet still exists
	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())
}

func newVNet(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference, addressPrefixes []string) *network.VirtualNetwork {
	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetwork_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: addressPrefixes,
			},
		},
	}
	return vnet
}

func Test_Networking_VirtualNetworkAndSubnetAdopted_SubnetsStillExist(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	originalVnet := vnet.DeepCopy()
	originalSubnet := subnet.DeepCopy()

	networkInterface := newVMNetworkInterface(tc, testcommon.AsOwner(rg), subnet)
	secret := createVMPasswordSecretAndRef(tc)
	vm := newVirtualMachine20201201(tc, rg, networkInterface, secret)
	tc.CreateResourcesAndWait(vm, vnet, subnet, networkInterface)

	// Annotate with skip
	oldSubnet := subnet.DeepCopy()
	subnet.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	// Can't use "PatchAndWait" here because it waits for generation to change but with just
	// this annotation change the generation will not change
	tc.Patch(oldSubnet, subnet)

	oldVnet := vnet.DeepCopy()
	vnet.Annotations["serviceoperator.azure.com/reconcile-policy"] = "skip"
	// Can't use "PatchAndWait" here because it waits for generation to change but with just
	// this annotation change the generation will not change
	tc.Patch(oldVnet, vnet)

	// Delete the VNET/Subnets
	tc.DeleteResourcesAndWait(vnet, subnet)

	// Now create the subnet/vnet again, this should be a no-op as it just adopts the existing resources
	subnet = originalSubnet
	vnet = originalVnet
	// If we mistakenly try to delete the subnets, this Create will fail with an error saying that the subnet is in use
	tc.CreateResourcesAndWait(vnet, subnet)
}
