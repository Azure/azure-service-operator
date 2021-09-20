/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/hack/generated/apis/microsoft.network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_Networking_VirtualNetwork_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vn := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Location: testcommon.DefaultTestRegion,
			AddressSpace: network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/8"},
			},
		},
	}

	tc.CreateResourceAndWait(vn)

	tc.Expect(vn.Status.Id).ToNot(BeNil())
	armId := *vn.Status.Id

	tc.DeleteResourceAndWait(vn)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(network.VirtualNetworksSpecAPIVersion20201101))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
