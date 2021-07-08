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

	"github.com/Azure/go-autorest/autorest/to"

	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/genruntime"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_LoadBalancer_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateNewTestResourceGroupAndWait()

	// Public IP Address
	sku := network.PublicIPAddressSkuNameStandard
	publicIPAddress := &network.PublicIPAddress{
		TypeMeta: metav1.TypeMeta{
			Kind: reflect.TypeOf(network.PublicIPAddress{}).Name(),
		},
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddresses_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			PublicIPAllocationMethod: network.PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic,
		},
	}

	tc.CreateResourceAndWait(publicIPAddress)

	// LoadBalancer
	loadBalancerSku := network.LoadBalancerSkuNameStandard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	protocol := network.InboundNatPoolPropertiesFormatProtocolTcp
	loadBalancer := &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancers_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			FrontendIPConfigurations: []network.LoadBalancers_Spec_Properties_FrontendIPConfigurations{
				{
					Name: lbFrontendName,
					PublicIPAddress: &network.SubResource{
						Reference: tc.MakeReferenceFromResource(publicIPAddress),
					},
				},
			},
			// TODO: The below stuff isn't really necessary for LB CRUD but is required for VMSS...
			InboundNatPools: []network.LoadBalancers_Spec_Properties_InboundNatPools{
				{
					Name: "MyFancyNatPool",
					FrontendIPConfiguration: &network.SubResource{
						Reference: genruntime.ResourceReference{
							// TODO: This is still really awkward
							ARMID: tc.MakeARMId(rg.Name, "Microsoft.Network", "loadBalancers", lbName, "frontendIPConfigurations", lbFrontendName),
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

	tc.CreateResourceAndWait(loadBalancer)

	// It should be created in Kubernetes
	g.Expect(loadBalancer.Status.Id).ToNot(BeNil())
	armId := *loadBalancer.Status.Id

	tc.DeleteResourceAndWait(loadBalancer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(ctx, armId, string(network.LoadBalancersSpecAPIVersion20201101))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
