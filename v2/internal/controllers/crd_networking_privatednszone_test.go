/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20180901"
	network20200601 "github.com/Azure/azure-service-operator/v2/api/network/v1api20200601"
	network20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
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
		testcommon.Subtest{
			Name: "Test_Recordset_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				PrivateDNSZones_CNAME_Record_Test(tc, zone)
			},
		},
	)

	tc.DeleteResourceAndWait(zone)

	//Ensure that the resource was really deleted in Azure
	armId, hasID := genruntime.GetResourceID(zone)
	tc.Expect(hasID).To(BeTrue())

	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func PrivateDNSZones_CNAME_Record_Test(tc *testcommon.KubePerTestContext, zone *network.PrivateDnsZone) {
	record := &network20200601.PrivateDnsZonesCNAMERecord{
		ObjectMeta: tc.MakeObjectMetaWithName("record"),
		Spec: network20200601.PrivateDnsZones_CNAME_Spec{
			CnameRecord: &network20200601.CnameRecord{Cname: to.Ptr("asotest.com")},
			Owner:       testcommon.AsOwner(zone),
			Ttl:         to.Ptr(3600),
		},
	}

	tc.CreateResourceAndWait(record)

	tc.Expect(record.Status.Id).ToNot(BeNil())

	old := record.DeepCopy()
	key := "foo"
	record.Spec.Metadata = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, record)
	tc.Expect(record.Status.Metadata).To(HaveKey(key))

	tc.DeleteResourceAndWait(record)
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
			Location: to.Ptr("global"),
			Owner:    testcommon.AsOwner(rg),
		},
	}
	return zone
}

func newVirtualNetworkLink(tc *testcommon.KubePerTestContext, dnsZone *network.PrivateDnsZone, vnet *network20201101.VirtualNetwork) *network20200601.PrivateDnsZonesVirtualNetworkLink {
	links := &network20200601.PrivateDnsZonesVirtualNetworkLink{
		ObjectMeta: tc.MakeObjectMetaWithName(dnsZone.Name + "-link"),
		Spec: network20200601.PrivateDnsZones_VirtualNetworkLink_Spec{
			Location:            to.Ptr("global"),
			Owner:               testcommon.AsOwner(dnsZone),
			VirtualNetwork:      &network20200601.SubResource{Reference: tc.MakeReferenceFromResource(vnet)},
			RegistrationEnabled: to.Ptr(false),
		},
	}
	return links
}
