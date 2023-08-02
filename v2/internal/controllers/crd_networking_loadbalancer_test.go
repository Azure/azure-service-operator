/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"reflect"
	"testing"

	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_LoadBalancer_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Public IP Address
	sku := network.PublicIPAddressSku_Name_Standard
	allocationMethod := network.IPAllocationMethod_Static
	publicIPAddress := &network.PublicIPAddress{
		TypeMeta: metav1.TypeMeta{
			Kind: reflect.TypeOf(network.PublicIPAddress{}).Name(),
		},
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			PublicIPAllocationMethod: &allocationMethod,
		},
	}

	tc.CreateResourceAndWait(publicIPAddress)

	// LoadBalancer
	loadBalancerSku := network.LoadBalancerSku_Name_Standard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.TransportProtocol_Tcp

	frontendIPConfigurationARMID, err := getfrontendIPConfigurationARMID(tc, rg, lbName, lbFrontendName)
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
					Name: &lbFrontendName,
					PublicIPAddress: &network.PublicIPAddressSpec_LoadBalancer_SubResourceEmbedded{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			// TODO: The below stuff isn't really necessary for LB CRUD but is required for VMSS...
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

	tc.CreateResourceAndWait(loadBalancer)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_LoadBalancer_InboundNatRule_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				LoadBalancer_InboundNatRule_CRUD(tc, loadBalancer, frontendIPConfigurationARMID)
			},
		},
	)

	// It should be created in Kubernetes
	g.Expect(loadBalancer.Status.Id).ToNot(BeNil())
	armId := *loadBalancer.Status.Id

	tc.DeleteResourceAndWait(loadBalancer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(ctx, armId, string(network.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}

func LoadBalancer_InboundNatRule_CRUD(tc *testcommon.KubePerTestContext, lb *network.LoadBalancer, frontendIPConfigurationARMID string) {

	natRule := &network.LoadBalancersInboundNatRule{
		ObjectMeta: tc.MakeObjectMeta("rule"),
		Spec: network.LoadBalancers_InboundNatRule_Spec{
			BackendPort: to.Ptr(22),
			FrontendIPConfiguration: &network.SubResource{
				Reference: &genruntime.ResourceReference{
					ARMID: frontendIPConfigurationARMID,
				},
			},
			FrontendPort: to.Ptr(28),
			Owner:        testcommon.AsOwner(lb),
		},
	}

	tc.CreateResourceAndWait(natRule)

	tc.Expect(natRule.Status.Id).ToNot(BeNil())

	old := natRule.DeepCopy()
	port := 20
	natRule.Spec.FrontendPort = &port

	tc.PatchResourceAndWait(old, natRule)
	tc.Expect(natRule.Status.FrontendPort).To(BeEquivalentTo(&port))

	tc.DeleteResourceAndWait(natRule)

}

// TODO: This is still really awkward
func getFrontendIPConfigurationARMID(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, lbName string, lbFrontendName string) (string, error) {
	return genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		"loadBalancers",
		lbName,
		"frontendIPConfigurations",
		lbFrontendName)
}
