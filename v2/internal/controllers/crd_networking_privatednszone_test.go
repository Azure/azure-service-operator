/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20180901"
	"github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_PrivateDnsZone_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	name := tc.Namer.GenerateName("pdz") + ".com"
	zone := newPrivateDNSZone(tc, name, rg)

	tc.CreateResourceAndWait(zone)

	tc.Expect(zone.Status.MaxNumberOfRecordSets).ToNot(BeNil())
	tc.Expect(zone.Status.MaxNumberOfVirtualNetworkLinks).ToNot(BeNil())

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_WorkspaceCompute_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				VirtualNetworkLinks_CRUD(tc, zone, rg)
			},
		},
	)

	tc.DeleteResourceAndWait(zone)

	// Ensure that the resource was really deleted in Azure
	armId, hasID := genruntime.GetResourceID(zone)
	tc.Expect(hasID).To(BeTrue())

	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func VirtualNetworkLinks_CRUD(tc *testcommon.KubePerTestContext, zone *network.PrivateDnsZone, rg *resources.ResourceGroup) {
	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	tc.CreateResourceAndWait(vnet)

	links := newVirtualNetworkLink(tc, zone, vnet)

	tc.CreateResourcesAndWait(links)

	tc.Expect(links.Status.Id).ToNot(BeNil())

	old := links.DeepCopy()
	key := "foo"
	links.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, links)
	tc.Expect(links.Status.Tags).To(HaveKey(key))

	tc.DeleteResource(links)

}

func newPrivateDNSZone(tc *testcommon.KubePerTestContext, name string, rg *resources.ResourceGroup) *network.PrivateDnsZone {
	zone := &network.PrivateDnsZone{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: network.PrivateDnsZone_Spec{
			Location: to.StringPtr("global"),
			Owner:    testcommon.AsOwner(rg),
		},
	}
	return zone
}

func newVirtualNetworkLink(tc *testcommon.KubePerTestContext, dnsZone *network.PrivateDnsZone, vnet *v1beta20201101.VirtualNetwork) *network.PrivateDnsZonesVirtualNetworkLink {
	links := &network.PrivateDnsZonesVirtualNetworkLink{
		ObjectMeta: tc.MakeObjectMetaWithName(dnsZone.Name + "-link"),
		Spec: network.PrivateDnsZones_VirtualNetworkLink_Spec{
			Location:            to.StringPtr("global"),
			Owner:               testcommon.AsOwner(dnsZone),
			VirtualNetwork:      &network.SubResource{Reference: tc.MakeReferenceFromResource(vnet)},
			RegistrationEnabled: to.BoolPtr(false),
		},
	}
	return links
}
