/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20240301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Refer to https://learn.microsoft.com/en-us/azure/firewall-manager/quick-firewall-policy-bicep for example firewall configuration
func Test_Networking_Firewall_20240301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	vnet := newVNet20240301(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/16"})
	publicIP := newPublicIP20240301(tc, testcommon.AsOwner(rg))
	subnet := newSubnet20240301(tc, vnet, "10.0.0.0/24")
	subnet.Name = "azurefirewallsubnet" // Name must be azurefirewallsubnet for Azure Firewall

	firewallPolicy := &network.FirewallPolicy{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("firewallpolicy")),
		Spec: network.FirewallPolicy_Spec{
			Owner:           testcommon.AsOwner(rg),
			Location:        tc.AzureRegion,
			ThreatIntelMode: to.Ptr(network.AzureFirewallThreatIntelMode_Alert),
		},
	}
	firewall := &network.AzureFirewall{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("firewall")),
		Spec: network.AzureFirewall_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			IpConfigurations: []network.AzureFirewallIPConfiguration{
				{
					Name: to.Ptr("firewall-ipconfig"),
					Subnet: &network.SubResource{
						Reference: tc.MakeReferenceFromResource(subnet),
					},
					PublicIPAddress: &network.SubResource{
						Reference: tc.MakeReferenceFromResource(publicIP),
					},
				},
			},
			FirewallPolicy: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(firewallPolicy),
			},
		},
	}

	tc.CreateResourcesAndWait(vnet, publicIP, subnet, firewallPolicy, firewall)
	tc.Expect(firewall.Status.Id).ToNot(BeNil())
	armId := *firewall.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_FirewallPolicy_RuleCollectionGroup_CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FirewallPolicy_RuleCollectionGroup_20240301_CRUD(tc, firewallPolicy)
			},
		},
	)

	// Update firewall to enable autoscaling
	old := firewall.DeepCopy()
	firewall.Spec.AutoscaleConfiguration = &network.AzureFirewallAutoscaleConfiguration{
		MaxCapacity: to.Ptr(4),
	}
	tc.PatchResourceAndWait(old, firewall)
	tc.Expect(firewall.Status.AutoscaleConfiguration.MaxCapacity).To(Equal(to.Ptr(4)))

	tc.DeleteResourceAndWait(firewall)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func FirewallPolicy_RuleCollectionGroup_20240301_CRUD(tc *testcommon.KubePerTestContext, firewallPolicy *network.FirewallPolicy) {
	ruleCollectionGroup := &network.FirewallPoliciesRuleCollectionGroup{
		ObjectMeta: tc.MakeObjectMeta("rulecollectiongroup"),
		Spec: network.FirewallPoliciesRuleCollectionGroup_Spec{
			Owner:    testcommon.AsOwner(firewallPolicy),
			Priority: to.Ptr(200),
			RuleCollections: []network.FirewallPolicyRuleCollection{
				{
					FirewallPolicyFilter: &network.FirewallPolicyFilterRuleCollection{
						Name:               to.Ptr("TestRuleCollection"),
						Priority:           to.Ptr(1000),
						RuleCollectionType: to.Ptr(network.FirewallPolicyFilterRuleCollection_RuleCollectionType_FirewallPolicyFilterRuleCollection),
						Action: &network.FirewallPolicyFilterRuleCollectionAction{
							Type: to.Ptr(network.FirewallPolicyFilterRuleCollectionActionType_Allow),
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(ruleCollectionGroup)

	tc.Expect(ruleCollectionGroup.Status.Id).ToNot(BeNil())

	old := ruleCollectionGroup.DeepCopy()
	newPriority := 300
	ruleCollectionGroup.Spec.Priority = &newPriority

	tc.PatchResourceAndWait(old, ruleCollectionGroup)
	tc.Expect(ruleCollectionGroup.Status.Priority).To(BeEquivalentTo(&newPriority))

	tc.DeleteResourceAndWait(ruleCollectionGroup)
}
