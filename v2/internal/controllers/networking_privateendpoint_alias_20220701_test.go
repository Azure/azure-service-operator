/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v20250301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// Test_Networking_PrivateEndpoint_Alias_20250301_CRUD verifies that a PrivateEndpoint can be created
// using a PrivateLinkService alias (wellKnownName) instead of an ARM ID or Kubernetes resource reference.
// See https://github.com/Azure/azure-service-operator/issues/4531
func Test_Networking_PrivateEndpoint_Alias_20250301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create networking prerequisites for the PrivateLinkService
	vnet := newVirtualNetwork20250301(tc, testcommon.AsOwner(rg))

	privateLinkServiceNetworkPolicies := network.SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Disabled
	plsSubnet := newSubnet20250301(tc, vnet, "10.0.0.0/24")
	plsSubnet.Spec.PrivateLinkServiceNetworkPolicies = &privateLinkServiceNetworkPolicies

	lb, frontendIPConfigurationARMID := newLoadBalancerForPLS20250301(tc, rg, plsSubnet)

	tc.CreateResourcesAndWait(vnet, plsSubnet, lb)

	// Create a PrivateLinkService and export its alias to a ConfigMap
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
						Reference: tc.MakeReferenceFromResource(plsSubnet),
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
	tc.Expect(pls.Status.Alias).ToNot(BeNil())

	plsAlias := *pls.Status.Alias
	tc.T.Logf("PrivateLinkService alias: %s", plsAlias)

	// Create a second subnet for the PrivateEndpoint (can't reuse the PLS subnet)
	peSubnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("pesubnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.1.0/24"),
		},
	}
	tc.CreateResourceAndWait(peSubnet)

	// Create a PrivateEndpoint using the PLS alias as wellKnownName.
	// Aliases must use ManualPrivateLinkServiceConnections (not PrivateLinkServiceConnections) —
	// Azure requires the manual path when connecting by alias.
	// Because we set AutoApproval on the PLS, this will be auto-approved despite using the manual path.
	endpoint := &network.PrivateEndpoint{
		ObjectMeta: tc.MakeObjectMeta("endpoint"),
		Spec: network.PrivateEndpoint_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			ManualPrivateLinkServiceConnections: []network.PrivateLinkServiceConnection{
				{
					Name: to.Ptr("aliasConnection"),
					PrivateLinkServiceReference: &genruntime.WellKnownResourceReference{
						WellKnownName: plsAlias,
					},
				},
			},
			Subnet: &network.Subnet_PrivateEndpoint_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(peSubnet),
			},
		},
	}

	tc.CreateResourceAndWait(endpoint)
	tc.Expect(endpoint.Status.Id).ToNot(BeNil())
	endpointArmId := *endpoint.Status.Id

	tc.DeleteResourceAndWait(endpoint)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, endpointArmId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newVirtualNetwork20250301(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.VirtualNetwork {
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
