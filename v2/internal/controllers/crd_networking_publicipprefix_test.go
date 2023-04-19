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

func Test_Networking_PublicIPPrefix_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	publicIpPrefix := &network.PublicIPPrefix{
		ObjectMeta: tc.MakeObjectMeta("ipprefix"),
		Spec: network.PublicIPPrefix_Spec{
			Location:               tc.AzureRegion,
			Owner:                  testcommon.AsOwner(rg),
			PrefixLength:           to.Ptr(28),
			PublicIPAddressVersion: to.Ptr(network.IPVersion_IPv4),
			Sku: &network.PublicIPPrefixSku{
				Name: to.Ptr(network.PublicIPPrefixSku_Name_Standard),
				Tier: to.Ptr(network.PublicIPPrefixSku_Tier_Regional),
			},
		},
	}

	tc.CreateResourceAndWait(publicIpPrefix)

	tc.Expect(publicIpPrefix.Status.Id).ToNot(BeNil())
	armId := *publicIpPrefix.Status.Id

	old := publicIpPrefix.DeepCopy()
	key := "foo"
	publicIpPrefix.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, publicIpPrefix)
	tc.Expect(publicIpPrefix.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(publicIpPrefix)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

}
