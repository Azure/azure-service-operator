/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	network20201101 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	network20220701 "github.com/Azure/azure-service-operator/v2/api/network/v1beta20220701"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_PrivateLinkService_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))

	privateLinkServiceNetworkPolicies := network20201101.SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Disabled
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))
	subnet.Spec.PrivateLinkServiceNetworkPolicies = &privateLinkServiceNetworkPolicies

	lb, frontendIPConfigurationARMID := newLbForPLS(tc, rg, subnet)

	tc.CreateResourcesAndWait(vnet, subnet, lb)

	ipAllocationMethod := network20220701.IPAllocationMethod_Dynamic
	ipVersion := network20220701.IPVersion_IPv4
	plsConfigMap := "plsconfig"
	pls := &network20220701.PrivateLinkService{
		ObjectMeta: tc.MakeObjectMeta("pls"),
		Spec: network20220701.PrivateLinkService_Spec{
			AutoApproval: &network20220701.ResourceSet{
				Subscriptions: []string{tc.AzureSubscription},
			},
			IpConfigurations: []network20220701.PrivateLinkServiceIpConfiguration{
				{
					Name:                      to.StringPtr("config"),
					PrivateIPAddressVersion:   &ipVersion,
					PrivateIPAllocationMethod: &ipAllocationMethod,
					Subnet: &network20220701.Subnet_PrivateLinkService_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(subnet),
					},
				},
			},
			LoadBalancerFrontendIpConfigurations: []network20220701.FrontendIPConfiguration_PrivateLinkService_SubResourceEmbedded{
				{
					Reference: &genruntime.ResourceReference{
						ARMID: frontendIPConfigurationARMID,
					},
				},
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Visibility: &network20220701.ResourceSet{
				Subscriptions: []string{tc.AzureSubscription},
			},
			OperatorSpec: &network20220701.PrivateLinkServiceOperatorSpec{
				ConfigMaps: &network20220701.PrivateLinkServiceOperatorConfigMaps{
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
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network20220701.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

}

func newLbForPLS(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, subnet *network20201101.VirtualNetworksSubnet) (*network20201101.LoadBalancer, string) {
	ipAllocationMethod := network20201101.IPAllocationMethod_Dynamic
	// LoadBalancer
	loadBalancerSku := network20201101.LoadBalancerSku_Name_Standard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network20201101.TransportProtocol_Tcp

	frontendIPConfigurationARMID, err := getFrontendIPConfigurationARMID(tc, rg, lbName, lbFrontendName)
	tc.Expect(err).To(BeNil())

	loadBalancer := &network20201101.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network20201101.LoadBalancer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network20201101.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network20201101.FrontendIPConfiguration_LoadBalancer_SubResourceEmbedded{
				{
					Name:                      &lbFrontendName,
					PrivateIPAllocationMethod: &ipAllocationMethod,
					Subnet: &network20201101.Subnet_LoadBalancer_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(subnet),
					},
				},
			},
			InboundNatPools: []network20201101.InboundNatPool{
				{
					Name: to.StringPtr("MyFancyNatPool"),
					FrontendIPConfiguration: &network20201101.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: frontendIPConfigurationARMID,
						},
					},
					Protocol:               &protocol,
					FrontendPortRangeStart: to.IntPtr(50_000),
					FrontendPortRangeEnd:   to.IntPtr(51_000),
					BackendPort:            to.IntPtr(22),
				},
			},
		},
	}

	return loadBalancer, frontendIPConfigurationARMID
}
