/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network2020 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Refer to the follow ARM template examples:
// https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.network/azure-bastion/azuredeploy.json
// https://github.com/Azure/azure-quickstart-templates/blob/master/demos/private-aks-cluster/bastion.bicep

func Test_Networking_BastionHost_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/8"})
	subnet := &network2020.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network2020.VirtualNetworks_Subnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.0.0/24"),
			AzureName:     "AzureBastionSubnet", // The subnet must have this name in Azure for Bastion to work
		},
	}
	publicIPAddress := newPublicIp(tc, testcommon.AsOwner(rg))

	host := &network.BastionHost{
		ObjectMeta: tc.MakeObjectMeta("bastion"),
		Spec: network.BastionHost_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			IpConfigurations: []network.BastionHostIPConfiguration{
				{
					Name: to.Ptr("IpConf"),
					PublicIPAddress: &network.BastionHostSubResource{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
					Subnet: &network.BastionHostSubResource{
						Reference: tc.MakeReferenceFromResource(subnet),
					},
				},
			},
		},
	}

	tc.CreateResourcesAndWait(vnet, subnet, publicIPAddress, host)

	tc.Expect(host.Status.Id).ToNot(BeNil())
	armID := *publicIPAddress.Status.Id

	tc.DeleteResourceAndWait(host)

	tc.DeleteResourceAndWait(publicIPAddress)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armID, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
