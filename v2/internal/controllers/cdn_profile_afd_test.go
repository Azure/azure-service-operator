/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	cdn "github.com/Azure/azure-service-operator/v2/api/cdn/v1api20230501"
	frontdoor "github.com/Azure/azure-service-operator/v2/api/network.frontdoor/v1api20220501"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// https://learn.microsoft.com/en-us/azure/frontdoor/create-front-door-template
// https://learn.microsoft.com/en-us/azure/frontdoor/front-door-quickstart-template-samples?pivots=front-door-standard-premium
// https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.cdn/front-door-standard-premium-custom-domain-customer-certificate/azuredeploy.json

func Test_CDN_Profile_AFD_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	sku := cdn.Sku_Name_Standard_AzureFrontDoor
	profile := &cdn.Profile{
		ObjectMeta: tc.MakeObjectMeta("cdnprofile"),
		Spec: cdn.Profile_Spec{
			Location: to.Ptr("Global"),
			Owner:    testcommon.AsOwner(rg),
			Sku:      &cdn.Sku{Name: &sku},
		},
	}

	// TODO: See what comes back on the profile based on the child resources being created...

	// Difficult to create these in subtests given the diamond dependency for routes and security policies,
	// so instead we test them all here.
	endpoint := newAFDEndpoint(tc, profile)
	originGroup := newAFDOriginGroup(tc, profile)
	origin := newAFDOrigin(tc, originGroup)
	ruleSet := newRuleSet(tc, profile)
	rule := newRule(tc, ruleSet)
	route := newRoute(tc, endpoint, originGroup, ruleSet)
	customDomain := newAFDCustomDomain(tc, profile)
	firewall := newAFDFirewall(tc, rg)
	securityPolicy := newSecurityPolicy(tc, profile, endpoint, firewall)
	// Can't easily test secret as it requires creation of a keyvault secret, but the resource itself is quite simple
	// so it seems safe to exclude from testing here

	tc.CreateResourcesAndWait(profile, endpoint, originGroup, origin, route, customDomain, firewall, securityPolicy, ruleSet, rule)

	tc.Expect(*profile.Status.Location).To(Equal("Global"))
	tc.Expect(string(*profile.Status.Sku.Name)).To(Equal("Standard_AzureFrontDoor"))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CDN Profile Update",
			Test: func(tc *testcommon.KubePerTestContext) {
				CDN_Profile_Update(tc, profile)
			},
		},
		testcommon.Subtest{
			Name: "CDN AfdEndpoint Update",
			Test: func(tc *testcommon.KubePerTestContext) {
				CDN_AfdEndpoint_Update(tc, endpoint)
			},
		},
		testcommon.Subtest{
			Name: "CDN AfdOriginGroup Update",
			Test: func(tc *testcommon.KubePerTestContext) {
				CDN_AfdOriginGroup_Update(tc, originGroup)
			},
		},
		testcommon.Subtest{
			Name: "CDN AfdOrigin Update",
			Test: func(tc *testcommon.KubePerTestContext) {
				CDN_AfdOrigin_Update(tc, origin)
			},
		},
		testcommon.Subtest{
			Name: "CDN Rule Update",
			Test: func(tc *testcommon.KubePerTestContext) {
				CDN_Rule_Update(tc, rule)
			},
		},
		testcommon.Subtest{
			Name: "CDN Route Update",
			Test: func(tc *testcommon.KubePerTestContext) {
				CDN_Route_Update(tc, route)
			},
		},
	)

	// Unroll the resources in reverse order to ensure that delete works
	tc.DeleteResourceAndWait(route)
	tc.Expect(route.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(origin)
	tc.Expect(origin.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(originGroup)
	tc.Expect(originGroup.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(rule)
	tc.Expect(rule.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(ruleSet)
	tc.Expect(ruleSet.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(securityPolicy)
	tc.Expect(securityPolicy.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(endpoint)
	tc.Expect(endpoint.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(customDomain)
	tc.Expect(customDomain.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))

	tc.DeleteResourceAndWait(profile)
	tc.Expect(profile.Status.Id).To(tc.MatchAzure.BeDeleted(string(cdn.APIVersion_Value)))
}

func newAFDEndpoint(tc *testcommon.KubePerTestContext, profile *cdn.Profile) *cdn.AfdEndpoint {
	return &cdn.AfdEndpoint{
		ObjectMeta: tc.MakeObjectMeta("cdn-endpoint"),
		Spec: cdn.AfdEndpoint_Spec{
			Owner:        testcommon.AsOwner(profile),
			Location:     to.Ptr("Global"),
			EnabledState: to.Ptr(cdn.AFDEndpointProperties_EnabledState_Enabled),
		},
	}
}

func newAFDOriginGroup(tc *testcommon.KubePerTestContext, profile *cdn.Profile) *cdn.AfdOriginGroup {
	return &cdn.AfdOriginGroup{
		ObjectMeta: tc.MakeObjectMeta("origingroup"),
		Spec: cdn.AfdOriginGroup_Spec{
			Owner: testcommon.AsOwner(profile),
			LoadBalancingSettings: &cdn.LoadBalancingSettingsParameters{
				SampleSize:                to.Ptr(4),
				SuccessfulSamplesRequired: to.Ptr(3),
			},
			HealthProbeSettings: &cdn.HealthProbeParameters{
				ProbePath:              to.Ptr("/"),
				ProbeRequestType:       to.Ptr(cdn.HealthProbeParameters_ProbeRequestType_HEAD),
				ProbeProtocol:          to.Ptr(cdn.HealthProbeParameters_ProbeProtocol_Http),
				ProbeIntervalInSeconds: to.Ptr(100),
			},
		},
	}
}

func newAFDOrigin(tc *testcommon.KubePerTestContext, originGroup *cdn.AfdOriginGroup) *cdn.AfdOrigin {
	return &cdn.AfdOrigin{
		ObjectMeta: tc.MakeObjectMeta("origin"),
		Spec: cdn.AfdOrigin_Spec{
			Owner:    testcommon.AsOwner(originGroup),
			HostName: to.Ptr("example.com"),
			HttpPort: to.Ptr(80),
		},
	}
}

func newRoute(tc *testcommon.KubePerTestContext, endpoint *cdn.AfdEndpoint, originGroup *cdn.AfdOriginGroup, ruleSet *cdn.RuleSet) *cdn.Route {
	return &cdn.Route{
		ObjectMeta: tc.MakeObjectMeta("route"),
		Spec: cdn.Route_Spec{
			Owner: testcommon.AsOwner(endpoint),
			OriginGroup: &cdn.ResourceReference{
				Reference: tc.MakeReferenceFromResource(originGroup),
			},
			RuleSets: []cdn.ResourceReference{
				{
					Reference: tc.MakeReferenceFromResource(ruleSet),
				},
			},
			SupportedProtocols: []cdn.AFDEndpointProtocols{
				cdn.AFDEndpointProtocols_Http,
				cdn.AFDEndpointProtocols_Https,
			},
			PatternsToMatch: []string{
				"/*",
			},
			ForwardingProtocol:  to.Ptr(cdn.RouteProperties_ForwardingProtocol_HttpOnly),
			LinkToDefaultDomain: to.Ptr(cdn.RouteProperties_LinkToDefaultDomain_Enabled),
			HttpsRedirect:       to.Ptr(cdn.RouteProperties_HttpsRedirect_Enabled),
		},
	}
}

func newAFDCustomDomain(tc *testcommon.KubePerTestContext, profile *cdn.Profile) *cdn.AfdCustomDomain {
	return &cdn.AfdCustomDomain{
		ObjectMeta: tc.MakeObjectMeta("custdomain"),
		Spec: cdn.AfdCustomDomain_Spec{
			Owner:    testcommon.AsOwner(profile),
			HostName: to.Ptr("example.com"),
		},
	}
}

func newSecurityPolicy(tc *testcommon.KubePerTestContext, profile *cdn.Profile, endpoint *cdn.AfdEndpoint, firewall *frontdoor.WebApplicationFirewallPolicy) *cdn.SecurityPolicy {
	return &cdn.SecurityPolicy{
		ObjectMeta: tc.MakeObjectMeta("securitypolicy"),
		Spec: cdn.SecurityPolicy_Spec{
			Owner: testcommon.AsOwner(profile),
			Parameters: &cdn.SecurityPolicyPropertiesParameters{
				WebApplicationFirewall: &cdn.SecurityPolicyWebApplicationFirewallParameters{
					Type: to.Ptr(cdn.SecurityPolicyWebApplicationFirewallParameters_Type_WebApplicationFirewall),
					WafPolicy: &cdn.ResourceReference{
						Reference: tc.MakeReferenceFromResource(firewall),
					},
					Associations: []cdn.SecurityPolicyWebApplicationFirewallAssociation{
						{
							Domains: []cdn.ActivatedResourceReference{
								{
									Reference: tc.MakeReferenceFromResource(endpoint),
								},
							},
							PatternsToMatch: []string{
								"/*",
							},
						},
					},
				},
			},
		},
	}
}

func newRuleSet(tc *testcommon.KubePerTestContext, profile *cdn.Profile) *cdn.RuleSet {
	return &cdn.RuleSet{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("ruleset")),
		Spec: cdn.RuleSet_Spec{
			Owner: testcommon.AsOwner(profile),
		},
	}
}

func newRule(tc *testcommon.KubePerTestContext, rule *cdn.RuleSet) *cdn.Rule {
	return &cdn.Rule{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("rule")),
		Spec: cdn.Rule_Spec{
			Owner: testcommon.AsOwner(rule),
			RuleConditions: []cdn.DeliveryRuleCondition{
				{
					UrlPath: &cdn.DeliveryRuleUrlPathCondition{
						Name: to.Ptr(cdn.DeliveryRuleUrlPathCondition_Name_UrlPath),
						Parameters: &cdn.UrlPathMatchConditionParameters{
							Operator:        to.Ptr(cdn.UrlPathMatchConditionParameters_Operator_BeginsWith),
							NegateCondition: to.Ptr(false),
							MatchValues: []string{
								"secure/",
							},
							Transforms: []cdn.Transform{
								cdn.Transform_Lowercase,
							},
							TypeName: to.Ptr(cdn.UrlPathMatchConditionParameters_TypeName_DeliveryRuleUrlPathMatchConditionParameters),
						},
					},
				},
			},
			Actions: []cdn.DeliveryRuleAction{
				{
					UrlRedirect: &cdn.UrlRedirectAction{
						Name: to.Ptr(cdn.UrlRedirectAction_Name_UrlRedirect),
						Parameters: &cdn.UrlRedirectActionParameters{
							RedirectType:        to.Ptr(cdn.UrlRedirectActionParameters_RedirectType_TemporaryRedirect),
							DestinationProtocol: to.Ptr(cdn.UrlRedirectActionParameters_DestinationProtocol_Https),
							CustomHostname:      to.Ptr("microsoft.com"),
							CustomPath:          to.Ptr("/"),
							TypeName:            to.Ptr(cdn.UrlRedirectActionParameters_TypeName_DeliveryRuleUrlRedirectActionParameters),
						},
					},
				},
			},
		},
	}
}

func CDN_Profile_Update(tc *testcommon.KubePerTestContext, profile *cdn.Profile) {
	// Should start out nil
	tc.Expect(profile.Spec.OriginResponseTimeoutSeconds).To(BeNil())

	old := profile.DeepCopy()
	profile.Spec.OriginResponseTimeoutSeconds = to.Ptr(37)

	tc.PatchResourceAndWait(old, profile)

	tc.Expect(profile.Status.OriginResponseTimeoutSeconds).ToNot(BeNil())
	tc.Expect(*profile.Status.OriginResponseTimeoutSeconds).To(Equal(37))
}

func CDN_AfdEndpoint_Update(tc *testcommon.KubePerTestContext, endpoint *cdn.AfdEndpoint) {
	old := endpoint.DeepCopy()
	endpoint.Spec.Tags = map[string]string{
		"cheese": "string",
	}

	tc.PatchResourceAndWait(old, endpoint)

	tc.Expect(endpoint.Status.Tags).To(HaveKey("cheese"))
}

func CDN_AfdOriginGroup_Update(tc *testcommon.KubePerTestContext, originGroup *cdn.AfdOriginGroup) {
	old := originGroup.DeepCopy()
	originGroup.Spec.HealthProbeSettings.ProbeRequestType = to.Ptr(cdn.HealthProbeParameters_ProbeRequestType_GET)
	originGroup.Spec.HealthProbeSettings.ProbeIntervalInSeconds = to.Ptr(50)

	tc.PatchResourceAndWait(old, originGroup)

	tc.Expect(originGroup.Status.HealthProbeSettings).ToNot(BeNil())
	tc.Expect(originGroup.Status.HealthProbeSettings.ProbeRequestType).ToNot(BeNil())
	tc.Expect(*originGroup.Status.HealthProbeSettings.ProbeRequestType).To(Equal(cdn.HealthProbeParameters_ProbeRequestType_STATUS_GET))
	tc.Expect(originGroup.Status.HealthProbeSettings.ProbeIntervalInSeconds).ToNot(BeNil())
	tc.Expect(*originGroup.Status.HealthProbeSettings.ProbeIntervalInSeconds).To(Equal(50))
}

func CDN_AfdOrigin_Update(tc *testcommon.KubePerTestContext, origin *cdn.AfdOrigin) {
	old := origin.DeepCopy()
	origin.Spec.HttpPort = to.Ptr(8080)

	tc.PatchResourceAndWait(old, origin)

	tc.Expect(origin.Status.HttpPort).ToNot(BeNil())
	tc.Expect(*origin.Status.HttpPort).To(Equal(8080))
}

func CDN_Rule_Update(tc *testcommon.KubePerTestContext, rule *cdn.Rule) {
	old := rule.DeepCopy()
	rule.Spec.Order = to.Ptr(7)

	tc.PatchResourceAndWait(old, rule)

	tc.Expect(rule.Status.Order).ToNot(BeNil())
	tc.Expect(*rule.Status.Order).To(Equal(7))
}

func CDN_Route_Update(tc *testcommon.KubePerTestContext, route *cdn.Route) {
	old := route.DeepCopy()
	route.Spec.PatternsToMatch = []string{
		"/foo/*",
		"/bar/*",
	}

	tc.PatchResourceAndWait(old, route)

	tc.Expect(route.Status.PatternsToMatch).To(HaveLen(2))
}
