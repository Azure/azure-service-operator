/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/go-autorest/autorest/to"
)

func Test_Networking_VirtualNetworkGateway_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	defer tc.DeleteResourceAndWait(rg)

	vnet := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/16"})
	publicIPAddress := newPublicIp(tc, testcommon.AsOwner(rg))

	subnet := &network.VirtualNetworksSubnet{
		// Name should always be 'gatewaysubnet'
		ObjectMeta: tc.MakeObjectMetaWithName("gatewaysubnet"),
		Spec: network.VirtualNetworks_Subnets_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.StringPtr("10.0.0.0/24"),
		},
	}

	tc.CreateResourcesAndWait(vnet, subnet, publicIPAddress)

	gateway := newVnetGateway(tc, publicIPAddress, subnet, rg)

	tc.CreateResourceAndWait(gateway)
	tc.DeleteResourceAndWait(gateway)

}

func newVnetGateway(tc *testcommon.KubePerTestContext, publicIPAddress *network.PublicIPAddress, subnet *network.VirtualNetworksSubnet, rg *resources.ResourceGroup) *network.VirtualNetworkGateway {
	gatewayType := network.VirtualNetworkGateways_Spec_Properties_GatewayType_Vpn
	skuName := network.VirtualNetworkGatewaySku_Name_VpnGw2
	skuTier := network.VirtualNetworkGatewaySku_Tier_VpnGw2
	vpnGatewayGen := network.VirtualNetworkGateways_Spec_Properties_VpnGatewayGeneration_Generation1
	vpnType := network.VirtualNetworkGateways_Spec_Properties_VpnType_RouteBased

	return &network.VirtualNetworkGateway{
		ObjectMeta: tc.MakeObjectMeta("gateway"),
		Spec: network.VirtualNetworkGateway_Spec{
			GatewayType: &gatewayType,
			IpConfigurations: []network.VirtualNetworkGatewayIPConfiguration{{
				Name: to.StringPtr("config1"),
				PublicIPAddress: &network.SubResource{
					Reference: tc.MakeReferenceFromResource(publicIPAddress),
				},
				Subnet: &network.SubResource{
					Reference: tc.MakeReferenceFromResource(subnet),
				},
			}},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.VirtualNetworkGatewaySku{
				Name: &skuName,
				Tier: &skuTier,
			},
			VpnGatewayGeneration: &vpnGatewayGen,
			VpnType:              &vpnType,
		},
	}
}
