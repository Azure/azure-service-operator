/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Networking_NatGateway_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	publicIPAddress := newPublicIp(tc, testcommon.AsOwner(rg))

	natgw := &network.NatGateway{
		ObjectMeta: tc.MakeObjectMetaWithName("natgw"),
		Spec: network.NatGateway_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			PublicIpAddresses: []network.ApplicationGatewaySubResource{
				{
					Reference: tc.MakeReferenceFromResource(publicIPAddress),
				},
			},
			Sku: &network.NatGatewaySku{
				Name: to.Ptr(network.NatGatewaySku_Name_Standard),
			},
		},
	}

	tc.CreateResourcesAndWait(publicIPAddress, natgw)

	tc.Expect(publicIPAddress.Status.Id).ToNot(BeNil())
	tc.Expect(natgw.Status.Id).ToNot(BeNil())
	armId := *natgw.Status.Id

	old := natgw.DeepCopy()
	key := "foo"
	natgw.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, natgw)
	tc.Expect(natgw.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(natgw)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

}
