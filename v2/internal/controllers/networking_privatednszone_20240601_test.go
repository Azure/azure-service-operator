/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network20240301 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20240601"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_PrivateDnsZone_20240601_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	name := tc.Namer.GenerateName("pdz") + ".com"
	zone := newPrivateDNSZone20240601(tc, name, rg)

	tc.CreateResourceAndWait(zone)

	tc.Expect(zone.Status.MaxNumberOfRecordSets).ToNot(BeNil())
	tc.Expect(zone.Status.MaxNumberOfVirtualNetworkLinks).ToNot(BeNil())

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_VirtualNetworkLinks_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				PrivateDNSZone_VirtualNetworkLinks_20240601_CRUD(tc, zone, rg)
			},
		},
		testcommon.Subtest{
			Name: "Test_Recordset_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				PrivateDNSZones_CNAME_Record_20240601_Test(tc, zone)
			},
		},
	)

	tc.DeleteResourceAndWait(zone)

	// Ensure that the resource was really deleted in Azure
	armId, hasID := genruntime.GetResourceID(zone)
	tc.Expect(hasID).To(BeTrue())

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func PrivateDNSZones_CNAME_Record_20240601_Test(tc *testcommon.KubePerTestContext, zone *network.PrivateDnsZone) {
	record := &network.PrivateDnsZonesCNAMERecord{
		ObjectMeta: tc.MakeObjectMetaWithName("record"),
		Spec: network.PrivateDnsZonesCNAMERecord_Spec{
			CnameRecord: &network.CnameRecord{Cname: to.Ptr("asotest.com")},
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

func PrivateDNSZone_VirtualNetworkLinks_20240601_CRUD(tc *testcommon.KubePerTestContext, zone *network.PrivateDnsZone, rg *resources.ResourceGroup) {
	vnet := newVMVirtualNetwork20240301(tc, testcommon.AsOwner(rg))
	link := newVirtualNetworkLink20240601(tc, zone, vnet)

	tc.CreateResourcesAndWait(vnet, link)

	tc.Expect(link.Status.Id).ToNot(BeNil())

	old := link.DeepCopy()
	key := "foo"
	link.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, link)
	tc.Expect(link.Status.Tags).To(HaveKey(key))

	tc.DeleteResource(link)
}

func newPrivateDNSZone20240601(tc *testcommon.KubePerTestContext, name string, rg *resources.ResourceGroup) *network.PrivateDnsZone {
	zone := &network.PrivateDnsZone{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: network.PrivateDnsZone_Spec{
			Location: to.Ptr("global"),
			Owner:    testcommon.AsOwner(rg),
		},
	}
	return zone
}

func newVirtualNetworkLink20240601(tc *testcommon.KubePerTestContext, dnsZone *network.PrivateDnsZone, vnet *network20240301.VirtualNetwork) *network.PrivateDnsZonesVirtualNetworkLink {
	links := &network.PrivateDnsZonesVirtualNetworkLink{
		ObjectMeta: tc.MakeObjectMetaWithName(dnsZone.Name + "-link"),
		Spec: network.PrivateDnsZonesVirtualNetworkLink_Spec{
			Location:            to.Ptr("global"),
			Owner:               testcommon.AsOwner(dnsZone),
			VirtualNetwork:      &network.SubResource{Reference: tc.MakeReferenceFromResource(vnet)},
			RegistrationEnabled: to.Ptr(false),
		},
	}
	return links
}
