/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ManagedIdentity_ResourceCanBeCreated(t *testing.T) {
	t.Parallel()

	// The identity used by this test is always based on AAD Pod
	// Identity, so all we need to do is create and delete a resource
	// group and ensure that it works.

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Test creating another resource to ensure conversion webhooks
	// are working - resource group is the only one without multiple
	// versions (and so no conversion webhooks).
	vnet := network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetworks_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/8"},
			},
		},
	}
	tc.CreateResourceAndWait(&vnet)

	tc.DeleteResourcesAndWait(&vnet, rg)
}
