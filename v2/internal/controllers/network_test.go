/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func newLoadBalancerForVMSS(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	publicIPAddress *network.PublicIPAddress,
) *network.LoadBalancer {
	loadBalancerSku := network.LoadBalancerSku_Name_Standard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.TransportProtocol_Tcp

	// TODO: Getting this is SUPER awkward
	frontIPConfigurationARMID, err := genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		"loadBalancers",
		lbName,
		"frontendIPConfigurations",
		lbFrontendName)
	if err != nil {
		panic(err)
	}

	return &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded{
				{
					Name: &lbFrontendName,
					PublicIPAddress: &network.PublicIPAddressSpec_LoadBalancer_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			InboundNatPools: []network.InboundNatPool{
				{
					Name: to.Ptr("MyFancyNatPool"),
					FrontendIPConfiguration: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: frontIPConfigurationARMID,
						},
					},
					Protocol:               &protocol,
					FrontendPortRangeStart: to.Ptr(50_000),
					FrontendPortRangeEnd:   to.Ptr(51_000),
					BackendPort:            to.Ptr(22),
				},
			},
		},
	}
}

func newPublicIPAddressForVMSS(
	tc *testcommon.KubePerTestContext,
	owner *genruntime.KnownResourceReference,
) *network.PublicIPAddress {
	publicIPAddressSku := network.PublicIPAddressSku_Name_Standard
	allocationMethod := network.IPAllocationMethod_Static
	return &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &network.PublicIPAddressSku{
				Name: &publicIPAddressSku,
			},
			PublicIPAllocationMethod: &allocationMethod,
		},
	}
}

func newVMSubnet(
	tc *testcommon.KubePerTestContext,
	owner *genruntime.KnownResourceReference,
) *network.VirtualNetworksSubnet {
	return &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworks_Subnet_Spec{
			Owner:         owner,
			AddressPrefix: to.Ptr("10.0.0.0/24"),
		},
	}
}

func newVMVirtualNetwork(
	tc *testcommon.KubePerTestContext,
	owner *genruntime.KnownResourceReference,
) *network.VirtualNetwork {
	return &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("vn")),
		Spec: network.VirtualNetwork_Spec{
			Owner:    owner,
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
}
