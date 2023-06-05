/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Networking_DnsResolver_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/8"})

	resolver := newDnsResolver(tc, rg, vnet)

	tc.CreateResourcesAndWait(vnet, resolver)
	tc.Expect(vnet.Status.Id).ToNot(BeNil())
	tc.Expect(resolver.Status.Id).ToNot(BeNil())
	armId := *resolver.Status.Id

	old := resolver.DeepCopy()
	key := "foo"
	resolver.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, resolver)
	tc.Expect(resolver.Status.Tags).To(HaveKey(key))

	// Run sub-tests on storage account
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "DnsResolver InboundEndpoint CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DnsResolver_InboundEndpoint_CRUD(tc, resolver, vnet)
			},
		},
		testcommon.Subtest{
			Name: "DnsResolver OutboundEndpoint CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DnsResolver_OutboundEndpoint_CRUD(tc, resolver, vnet)
			},
		},
	)

	tc.DeleteResourceAndWait(resolver)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(v1api20220701.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func DnsResolver_InboundEndpoint_CRUD(tc *testcommon.KubePerTestContext, resolver *v1api20220701.DnsResolver, vnet *v1api20201101.VirtualNetwork) {
	subnet := newSubnet(tc, vnet, "10.0.0.0/24")
	tc.CreateResourceAndWait(subnet)
	defer tc.DeleteResourceAndWait(subnet)

	inbound := &v1api20220701.DnsResolversInboundEndpoint{
		ObjectMeta: tc.MakeObjectMeta("inbound"),
		Spec: v1api20220701.DnsResolvers_InboundEndpoint_Spec{
			IpConfigurations: []v1api20220701.IpConfiguration{
				{
					PrivateIpAllocationMethod: to.Ptr(v1api20220701.IpConfiguration_PrivateIpAllocationMethod_Dynamic),
					Subnet:                    &v1api20220701.DnsresolverSubResource{Reference: tc.MakeReferenceFromResource(subnet)},
				},
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(resolver),
		},
	}

	tc.CreateResourceAndWait(inbound)
	tc.Expect(inbound.Status.Id).ToNot(BeNil())
	armId := *inbound.Status.Id

	old := inbound.DeepCopy()
	key := "foo"
	inbound.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, inbound)
	tc.Expect(inbound.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(inbound)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(v1api20220701.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func DnsResolver_OutboundEndpoint_CRUD(tc *testcommon.KubePerTestContext, resolver *v1api20220701.DnsResolver, vnet *v1api20201101.VirtualNetwork) {
	subnet := newSubnet(tc, vnet, "10.225.0.0/28")
	tc.CreateResourceAndWait(subnet)
	defer tc.DeleteResourceAndWait(subnet)

	outbound := newDnsResolversOutboundEndpoint(tc, resolver, subnet)

	tc.CreateResourceAndWait(outbound)
	tc.Expect(outbound.Status.Id).ToNot(BeNil())
	armId := *outbound.Status.Id

	old := outbound.DeepCopy()
	key := "foo"
	outbound.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, outbound)
	tc.Expect(outbound.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(outbound)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(v1api20220701.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newDnsResolversOutboundEndpoint(tc *testcommon.KubePerTestContext, resolver *v1api20220701.DnsResolver, subnet *v1api20201101.VirtualNetworksSubnet) *v1api20220701.DnsResolversOutboundEndpoint {
	outbound := &v1api20220701.DnsResolversOutboundEndpoint{
		ObjectMeta: tc.MakeObjectMeta("outbound"),
		Spec: v1api20220701.DnsResolvers_OutboundEndpoint_Spec{
			Subnet:   &v1api20220701.DnsresolverSubResource{Reference: tc.MakeReferenceFromResource(subnet)},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(resolver),
		},
	}
	return outbound
}

func newDnsResolver(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, vnet *v1api20201101.VirtualNetwork) *v1api20220701.DnsResolver {
	resolver := &v1api20220701.DnsResolver{
		ObjectMeta: tc.MakeObjectMeta("resolver"),
		Spec: v1api20220701.DnsResolver_Spec{
			Location:       tc.AzureRegion,
			Owner:          testcommon.AsOwner(rg),
			VirtualNetwork: &v1api20220701.DnsresolverSubResource{Reference: tc.MakeReferenceFromResource(vnet)},
		},
	}
	return resolver
}
