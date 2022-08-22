/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20180901"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
)

func Test_Networking_PrivateDnsZone_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	name := tc.Namer.GenerateName("pdz") + ".com"
	zone := &network.PrivateDnsZone{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: network.PrivateDnsZones_Spec{
			Location: to.StringPtr("global"),
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(zone)

	tc.Expect(zone.Status.MaxNumberOfRecordSets).ToNot(BeNil())
	tc.Expect(zone.Status.MaxNumberOfVirtualNetworkLinks).ToNot(BeNil())

	tc.DeleteResourceAndWait(zone)

	// Ensure that the resource was really deleted in Azure
	armId, hasID := genruntime.GetResourceID(zone)
	tc.Expect(hasID).To(BeTrue())

	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
