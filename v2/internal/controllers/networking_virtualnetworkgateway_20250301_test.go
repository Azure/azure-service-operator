/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"
	"time"

	network "github.com/Azure/azure-service-operator/v2/api/network/v20250301"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Networking_VirtualNetworkGateway_20250301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	defer tc.DeleteResourceAndWait(rg)

	vnet := newVNet20250301(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/16"})

	// AZ VPN gateway SKUs require a zoned public IP
	publicIPSku := network.PublicIPAddressSku_Name_Standard
	publicIPAllocation := network.IPAllocationMethod_Static
	publicIPAddress := &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.PublicIPAddressSku{
				Name: &publicIPSku,
			},
			PublicIPAllocationMethod: &publicIPAllocation,
			Zones:                    []string{"1", "2", "3"},
		},
	}

	subnet := &network.VirtualNetworksSubnet{
		// Name should always be 'gatewaysubnet'
		ObjectMeta: tc.MakeObjectMetaWithName("gatewaysubnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}

	tc.CreateResourcesAndWait(vnet, subnet, publicIPAddress)

	gateway := newVnetGateway20250301(tc, publicIPAddress, subnet, rg)

	// AZ VPN gateways take longer to provision than non-AZ (~30-60 min)
	gen := gateway.GetGeneration()
	tc.CreateResource(gateway)
	tc.Eventually(gateway, tc.CustomOperationTimeout(90*time.Minute), tc.PollingInterval()).Should(tc.Match.BeProvisioned(gen))
	tc.DeleteResourceAndWait(gateway)
}

func newVnetGateway20250301(tc *testcommon.KubePerTestContext, publicIPAddress *network.PublicIPAddress, subnet *network.VirtualNetworksSubnet, rg *resources.ResourceGroup) *network.VirtualNetworkGateway {
	gatewayType := network.VirtualNetworkGatewayPropertiesFormat_GatewayType_Vpn
	skuName := network.VirtualNetworkGatewaySku_Name_VpnGw1AZ
	skuTier := network.VirtualNetworkGatewaySku_Tier_VpnGw1AZ
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
