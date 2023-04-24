/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20180501"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_DnsZone_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	name := tc.Namer.GenerateName("dz") + ".com"
	zone := newDNSZone(tc, name, rg)

	tc.CreateResourceAndWait(zone)

	tc.Expect(zone.Status.MaxNumberOfRecordSets).ToNot(BeNil())

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_DnsZoneRecordset_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DNSZones_CNAME_Record_Test(tc, zone)
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

func DNSZones_CNAME_Record_Test(tc *testcommon.KubePerTestContext, zone *network.DnsZone) {
	record := &network.DnsZonesCNAMERecord{
		ObjectMeta: tc.MakeObjectMetaWithName("record"),
		Spec: network.DnsZones_CNAME_Spec{
			CNAMERecord: &network.CnameRecord{Cname: to.Ptr("asotest.com")},
			Owner:       testcommon.AsOwner(zone),
			TTL:         to.Ptr(3600),
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

func newDNSZone(tc *testcommon.KubePerTestContext, name string, rg *resources.ResourceGroup) *network.DnsZone {
	zone := &network.DnsZone{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: network.DnsZone_Spec{
			Location: to.Ptr("global"),
			Owner:    testcommon.AsOwner(rg),
		},
	}
	return zone
}
