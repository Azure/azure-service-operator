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

	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20200501"
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
	publicIPAddress := &network.PublicIPAddresses{
		TypeMeta: metav1.TypeMeta{
			Kind: reflect.TypeOf(network.PublicIPAddresses{}).Name(),
		},
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddresses_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			Properties: network.PublicIPAddressPropertiesFormat{
				PublicIPAllocationMethod: network.PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic,
			},
		},
	}

	tc.CreateResourceAndWait(publicIPAddress)

	// LoadBalancer
	loadBalancerSku := network.LoadBalancerSkuNameStandard
	lbName := tc.Namer.GenerateName("loadbalancer")
	lbFrontendName := "LoadBalancerFrontend"
	loadBalancer := &network.LoadBalancer{
		ObjectMeta: tc.MakeObjectMetaWithName(lbName),
		Spec: network.LoadBalancers_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &network.LoadBalancerSku{
				Name: &loadBalancerSku,
			},
			Properties: network.LoadBalancerPropertiesFormat{
				FrontendIPConfigurations: []network.FrontendIPConfiguration{
					{
						Name: lbFrontendName,
						Properties: &network.FrontendIPConfigurationPropertiesFormat{
							PublicIPAddress: &network.SubResource{
								Reference: tc.MakeReferenceFromResource(publicIPAddress),
							},
						},
					},
				},
				// TODO: The below stuff isn't really necessary for LB CRUD but is required for VMSS...
				InboundNatPools: []network.InboundNatPool{
					{
						Name: "MyFancyNatPool",
						Properties: &network.InboundNatPoolPropertiesFormat{
							FrontendIPConfiguration: network.SubResource{
								Reference: genruntime.ResourceReference{
									// TODO: This is still really awkward
									ARMID: tc.MakeARMId(rg.Name, "Microsoft.Network", "loadBalancers", lbName, "frontendIPConfigurations", lbFrontendName),
								},
							},
							Protocol:               network.InboundNatPoolPropertiesFormatProtocolTcp,
							FrontendPortRangeStart: 50000,
							FrontendPortRangeEnd:   51000,
							BackendPort:            22,
						},
					},
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
	exists, retryAfter, err := tc.AzureClient.HeadResource(ctx, armId, "2020-05-01")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
