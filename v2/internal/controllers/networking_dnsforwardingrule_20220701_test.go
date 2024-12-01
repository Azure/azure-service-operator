/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Networking_DNSForwardingRuleSet_20220701_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVNet20201101(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/8"})
	resolver := newDnsResolver(tc, rg, vnet)
	subnet := newSubnet20201101(tc, vnet, "10.225.0.0/28")
	outboundEP := newDnsResolversOutboundEndpoint(tc, resolver, subnet)

	tc.CreateResourcesAndWait(vnet, subnet, resolver, outboundEP)

	ruleSet := &network.DnsForwardingRuleset{
		ObjectMeta: tc.MakeObjectMeta("ruleset"),
		Spec: network.DnsForwardingRuleset_Spec{
			DnsResolverOutboundEndpoints: []network.SubResource{
				{
					Reference: tc.MakeReferenceFromResource(outboundEP),
				},
			},
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourcesAndWait(ruleSet)
	tc.Expect(ruleSet.Status.Id).ToNot(BeNil())
	armId := *ruleSet.Status.Id

	old := ruleSet.DeepCopy()
	key := "foo"
	ruleSet.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, ruleSet)
	tc.Expect(ruleSet.Status.Tags).To(HaveKey(key))

	// Run sub-tests on storage account
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "DnsForwardingRuleset ForwardingRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DnsForwardingRuleset_ForwardingRules_20220701_CRUD(tc, ruleSet)
			},
		},
		testcommon.Subtest{
			Name: "DnsForwardingRuleset VirtualNetwork Link CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DnsForwardingRuleset_VirtualNetworkLink_20220701_CRUD(tc, rg, ruleSet)
			},
		},
	)

	tc.DeleteResourceAndWait(ruleSet)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func DnsForwardingRuleset_ForwardingRules_20220701_CRUD(tc *testcommon.KubePerTestContext, set *network.DnsForwardingRuleset) {
	rule := &network.DnsForwardingRuleSetsForwardingRule{
		ObjectMeta: tc.MakeObjectMeta("rule"),
		Spec: network.DnsForwardingRuleSetsForwardingRule_Spec{
			DomainName:          to.Ptr("test."),
			ForwardingRuleState: to.Ptr(network.ForwardingRuleProperties_ForwardingRuleState_Disabled),
			Owner:               testcommon.AsOwner(set),
			TargetDnsServers: []network.TargetDnsServer{
				{
					IpAddress: to.Ptr("192.168.1.1"),
					Port:      to.Ptr(53),
				},
			},
		},
	}

	tc.CreateResourceAndWait(rule)
	tc.Expect(rule.Status.Id).ToNot(BeNil())
	armId := *rule.Status.Id

	tc.DeleteResourceAndWait(rule)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func DnsForwardingRuleset_VirtualNetworkLink_20220701_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, set *network.DnsForwardingRuleset) {
	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	link := &network.DnsForwardingRuleSetsVirtualNetworkLink{
		ObjectMeta: tc.MakeObjectMeta("rule"),
		Spec: network.DnsForwardingRuleSetsVirtualNetworkLink_Spec{
			Owner: testcommon.AsOwner(set),
			VirtualNetwork: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(vnet),
			},
		},
	}

	tc.CreateResourcesAndWait(link, vnet)
	tc.Expect(link.Status.Id).ToNot(BeNil())
	armId := *link.Status.Id

	tc.DeleteResourcesAndWait(link, vnet)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
