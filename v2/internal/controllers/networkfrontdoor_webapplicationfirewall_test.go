/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	frontdoor "github.com/Azure/azure-service-operator/v2/api/network.frontdoor/v1api20220501"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// Test built from the following sample:
// https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.cdn/front-door-standard-premium-rate-limit/azuredeploy.json

func Test_NetworkFrontDoorFirewallPolicy_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	afdFirewall := newAFDFirewall(tc, rg)

	tc.CreateResourceAndWait(afdFirewall)

	tc.Expect(afdFirewall.Status.CustomRules.Rules).To(HaveLen(1))
	tc.Expect(afdFirewall.Status.Id).ToNot(BeNil())

	armID := *afdFirewall.Status.Id

	// Perform a simple patch
	old := afdFirewall.DeepCopy()
	afdFirewall.Spec.CustomRules.Rules = append(
		afdFirewall.Spec.CustomRules.Rules,
		frontdoor.CustomRule{
			Name:         to.Ptr("BlockTrafficFromIPRanges"),
			Priority:     to.Ptr(101),
			EnabledState: to.Ptr(frontdoor.CustomRule_EnabledState_Enabled),
			RuleType:     to.Ptr(frontdoor.CustomRule_RuleType_MatchRule),
			Action:       to.Ptr(frontdoor.ActionType_Block),
			MatchConditions: []frontdoor.MatchCondition{
				{
					MatchVariable: to.Ptr(frontdoor.MatchCondition_MatchVariable_RemoteAddr),
					Operator:      to.Ptr(frontdoor.MatchCondition_Operator_IPMatch),
					MatchValue: []string{
						"198.51.100.100",
					},
				},
			},
		})
	tc.PatchResourceAndWait(old, afdFirewall)
	tc.Expect(afdFirewall.Status.CustomRules.Rules).To(HaveLen(2))

	// Ensure delete works
	tc.DeleteResourceAndWait(afdFirewall)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armID,
		string(frontdoor.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func newAFDFirewall(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *frontdoor.WebApplicationFirewallPolicy {
	return &frontdoor.WebApplicationFirewallPolicy{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("firewall")),
		Spec: frontdoor.WebApplicationFirewallPolicy_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: to.Ptr("global"),
			Sku: &frontdoor.Sku{
				Name: to.Ptr(frontdoor.Sku_Name_Standard_AzureFrontDoor),
			},
			PolicySettings: &frontdoor.PolicySettings{
				EnabledState: to.Ptr(frontdoor.PolicySettings_EnabledState_Enabled),
				Mode:         to.Ptr(frontdoor.PolicySettings_Mode_Detection),
			},
			CustomRules: &frontdoor.CustomRuleList{
				Rules: []frontdoor.CustomRule{
					{
						Name:                       to.Ptr("RateLimit"),
						Priority:                   to.Ptr(100),
						EnabledState:               to.Ptr(frontdoor.CustomRule_EnabledState_Enabled),
						RuleType:                   to.Ptr(frontdoor.CustomRule_RuleType_RateLimitRule),
						RateLimitThreshold:         to.Ptr(10),
						RateLimitDurationInMinutes: to.Ptr(5),
						Action:                     to.Ptr(frontdoor.ActionType_Block),
						MatchConditions: []frontdoor.MatchCondition{
							{
								MatchVariable:   to.Ptr(frontdoor.MatchCondition_MatchVariable_RemoteAddr),
								Operator:        to.Ptr(frontdoor.MatchCondition_Operator_IPMatch),
								NegateCondition: to.Ptr(true),
								MatchValue: []string{
									"192.0.2.0/24",
								},
							},
						},
					},
				},
			},
		},
	}
}
