/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_PrivateLinkService_20240301_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))

	privateLinkServiceNetworkPolicies := network.SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Disabled
	subnet := newVMSubnet20240301(tc, testcommon.AsOwner(vnet))
	subnet.Spec.PrivateLinkServiceNetworkPolicies = &privateLinkServiceNetworkPolicies

	lb, frontendIPConfigurationARMID := newLoadBalancerForPLS20240301(tc, rg, subnet)

	tc.CreateResourcesAndWait(vnet, subnet, lb)

	ipAllocationMethod := network.IPAllocationMethod_Dynamic
	ipVersion := network.IPVersion_IPv4
	plsConfigMap := "plsconfig"
	pls := &network.PrivateLinkService{
		ObjectMeta: tc.MakeObjectMeta("pls"),
		Spec: network.PrivateLinkService_Spec{
			AutoApproval: &network.ResourceSet{
				Subscriptions: []string{tc.AzureSubscription},
			},
			IpConfigurations: []network.PrivateLinkServiceIpConfiguration{
				{
					Name:                      to.Ptr("config"),
					PrivateIPAddressVersion:   &ipVersion,
					PrivateIPAllocationMethod: &ipAllocationMethod,
					Subnet: &network.Subnet_PrivateLinkService_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(subnet),
					},
				},
			},
			LoadBalancerFrontendIpConfigurations: []network.FrontendIPConfiguration_PrivateLinkService_SubResourceEmbedded{
				{
					Reference: &genruntime.ResourceReference{
						ARMID: frontendIPConfigurationARMID,
					},
				},
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Visibility: &network.ResourceSet{
				Subscriptions: []string{tc.AzureSubscription},
			},
			OperatorSpec: &network.PrivateLinkServiceOperatorSpec{
				ConfigMaps: &network.PrivateLinkServiceOperatorConfigMaps{
					Alias: &genruntime.ConfigMapDestination{
						Name: plsConfigMap,
						Key:  "alias",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(pls)

	tc.Expect(pls.Status.Id).ToNot(BeNil())
	armId := *pls.Status.Id

	tc.ExpectConfigMapHasKeysAndValues(plsConfigMap, "alias", *pls.Status.Alias)

	old := pls.DeepCopy()
	key := "foo"
	pls.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, pls)
	tc.Expect(pls.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(pls)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newLoadBalancerForPLS20240301(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, subnet *network.VirtualNetworksSubnet) (*network.LoadBalancer, string) {
	ipAllocationMethod := network.IPAllocationMethod_Dynamic
	// LoadBalancer
	loadBalancerSku := network.LoadBalancerSku_Name_Standard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.TransportProtocol_Tcp

	frontendIPConfigurationARMID, err := getFrontendIPConfigurationARMID(tc, rg, lbName, lbFrontendName)
	tc.Expect(err).To(BeNil())

	loadBalancer := &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded{
				{
					Name:                      &lbFrontendName,
					PrivateIPAllocationMethod: &ipAllocationMethod,
					Subnet: &network.Subnet_LoadBalancer_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(subnet),
					},
				},
			},
			InboundNatPools: []network.InboundNatPool{
				{
					Name: to.Ptr("MyFancyNatPool"),
					FrontendIPConfiguration: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: frontendIPConfigurationARMID,
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

	return loadBalancer, frontendIPConfigurationARMID
}
