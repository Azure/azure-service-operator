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
	network20200601 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20200601"
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
			Name: "Test_VirtualNetworkLinks_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				PrivateDNSZone_VirtualNetworkLinks_CRUD(tc, zone, rg)
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

func PrivateDNSZone_VirtualNetworkLinks_CRUD(tc *testcommon.KubePerTestContext, zone *network.PrivateDnsZone, rg *resources.ResourceGroup) {
	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	link := newVirtualNetworkLink(tc, zone, vnet)

	tc.CreateResourcesAndWait(vnet, link)

	tc.Expect(link.Status.Id).ToNot(BeNil())

	old := link.DeepCopy()
	key := "foo"
	link.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, link)
	tc.Expect(link.Status.Tags).To(HaveKey(key))

	tc.DeleteResource(link)
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

func newVirtualNetworkLink(tc *testcommon.KubePerTestContext, dnsZone *network.PrivateDnsZone, vnet *v1beta20201101.VirtualNetwork) *network20200601.PrivateDnsZonesVirtualNetworkLink {
	links := &network20200601.PrivateDnsZonesVirtualNetworkLink{
		ObjectMeta: tc.MakeObjectMetaWithName(dnsZone.Name + "-link"),
		Spec: network20200601.PrivateDnsZones_VirtualNetworkLink_Spec{
			Location:            to.StringPtr("global"),
			Owner:               testcommon.AsOwner(dnsZone),
			VirtualNetwork:      &network20200601.SubResource{Reference: tc.MakeReferenceFromResource(vnet)},
			RegistrationEnabled: to.BoolPtr(false),
		},
	}
	return links
}
