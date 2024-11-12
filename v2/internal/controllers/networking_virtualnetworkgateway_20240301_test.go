/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Networking_VirtualNetworkGateway_20240301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	defer tc.DeleteResourceAndWait(rg)

	vnet := newVNet20240301(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/16"})
	publicIPAddress := newPublicIP20240301(tc, testcommon.AsOwner(rg))

	subnet := &network.VirtualNetworksSubnet{
		// Name should always be 'gatewaysubnet'
		ObjectMeta: tc.MakeObjectMetaWithName("gatewaysubnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}

	tc.CreateResourcesAndWait(vnet, subnet, publicIPAddress)

	gateway := newVnetGateway20240301(tc, publicIPAddress, subnet, rg)

	tc.CreateResourceAndWait(gateway)
	tc.DeleteResourceAndWait(gateway)
}

func newVnetGateway20240301(tc *testcommon.KubePerTestContext, publicIPAddress *network.PublicIPAddress, subnet *network.VirtualNetworksSubnet, rg *resources.ResourceGroup) *network.VirtualNetworkGateway {
	gatewayType := network.VirtualNetworkGatewayPropertiesFormat_GatewayType_Vpn
	skuName := network.VirtualNetworkGatewaySku_Name_VpnGw2
	skuTier := network.VirtualNetworkGatewaySku_Tier_VpnGw2
	vpnGatewayGen := network.VirtualNetworkGatewayPropertiesFormat_VpnGatewayGeneration_Generation1
	vpnType := network.VirtualNetworkGatewayPropertiesFormat_VpnType_RouteBased

	return &network.VirtualNetworkGateway{
		ObjectMeta: tc.MakeObjectMeta("gateway"),
		Spec: network.VirtualNetworkGateway_Spec{
			GatewayType: &gatewayType,
			IpConfigurations: []network.VirtualNetworkGatewayIPConfiguration{{
				Name: to.Ptr("config1"),
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
