/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	network2024 "github.com/Azure/azure-service-operator/v2/api/network/v1api20240101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_ApplicationGateway_20220701_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	appGatewayName := "appgatewaytest"
	vnet := newVNet20201101(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/8"})
	subnet := newSubnet20201101(tc, vnet, "10.0.0.0/24")

	publicIPAddress := newPublicIP20201101(tc, testcommon.AsOwner(rg))
	tc.CreateResourcesAndWait(publicIPAddress, vnet, subnet)

	tc.Expect(vnet.Status.Id).ToNot(BeNil())
	tc.Expect(subnet.Status.Id).ToNot(BeNil())

	appGtsListenerName := "app-gw-http-listener-1"
	appGtwFHttpListenerID, err := getFrontendPortsARMID(tc, rg, appGatewayName, "httpListeners", appGtsListenerName)
	tc.Expect(err).To(BeNil())

	appGtwFIPConfig := defineApplicationGatewayIPConfiguration20220701(tc.MakeReferenceFromResource(subnet))
	appGtwFeIpConfig, appGtwFeIpConfigID := defineApplicationGatewayFrontendIPConfiguration20220701(tc, rg, appGatewayName, tc.MakeReferenceFromResource(subnet), tc.MakeReferenceFromResource(publicIPAddress))
	appGtwFEPorts, appGtwFePortsID := defineApplicationGatewayFrontendPort20220701(tc, rg, appGatewayName)
	appGtwBackendPools, appGtwBackendPoolsID := defineApplicationGatewayBackendAddressPool20220701(tc, rg, appGatewayName)
	appGtwBackendHttpSettings, appGtwBackendHttpSettingsID := defineApplicationGatewayBackendHttpSettings20220701(tc, rg, appGatewayName)
	appGtwWaf := defineApplicationGatewayWebApplicationFirewallPolicy(tc, rg)

	applicationGateway := &network.ApplicationGateway{
		ObjectMeta: tc.MakeObjectMetaWithName(appGatewayName),
		Spec: network.ApplicationGateway_Spec{
			AutoscaleConfiguration: &network.ApplicationGatewayAutoscaleConfiguration{
				MaxCapacity: to.Ptr(3),
				MinCapacity: to.Ptr(1),
			},
			AzureName: appGatewayName,
			Location:  to.Ptr(rg.Location()),
			Owner:     testcommon.AsOwner(rg),
			Sku: &network.ApplicationGatewaySku{
				Name: to.Ptr(network.ApplicationGatewaySku_Name(network.ApplicationGatewaySku_Name_STATUS_WAF_V2)),
				Tier: to.Ptr(network.ApplicationGatewaySku_Tier(network.ApplicationGatewaySku_Tier_STATUS_WAF_V2)),
			},
			GatewayIPConfigurations:  appGtwFIPConfig,
			FrontendIPConfigurations: appGtwFeIpConfig,
			FrontendPorts:            appGtwFEPorts,
			FirewallPolicy: &network.SubResource{
				Reference: tc.MakeReferenceFromResource(appGtwWaf),
			},
			HttpListeners: []network.ApplicationGatewayHttpListener{
				{
					Name: to.Ptr(appGtsListenerName),
					FrontendIPConfiguration: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwFeIpConfigID,
						},
					},
					FrontendPort: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwFePortsID,
						},
					},
					Protocol:  to.Ptr(network.ApplicationGatewayProtocol_Http),
					HostNames: []string{"test.contoso.com"},
				},
			},
			BackendAddressPools:           appGtwBackendPools,
			BackendHttpSettingsCollection: appGtwBackendHttpSettings,
			RequestRoutingRules: []network.ApplicationGatewayRequestRoutingRule{
				{
					Name: to.Ptr("app-gtw-routing-rule-1"),
					BackendAddressPool: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwBackendPoolsID,
						},
					},
					BackendHttpSettings: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwBackendHttpSettingsID,
						},
					},
					HttpListener: &network.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwFHttpListenerID,
						},
					},
					RuleType: to.Ptr(network.ApplicationGatewayRequestRoutingRulePropertiesFormat_RuleType_Basic),
					Priority: to.Ptr(100),
				},
			},
		},
	}
	tc.CreateResourcesAndWait(applicationGateway, appGtwWaf)
	tc.Expect(applicationGateway.Status.Id).ToNot(BeNil())
	tc.Expect(applicationGateway.Status.Sku).ToNot(BeNil())
	tc.Expect(applicationGateway.Status.Sku.Tier).ToNot(BeNil())
	tc.Expect(*applicationGateway.Status.Sku.Tier).To(Equal(network.ApplicationGatewaySku_Tier_STATUS_WAF_V2))

	armId := *applicationGateway.Status.Id
	tc.DeleteResourcesAndWait(applicationGateway, publicIPAddress, subnet, vnet)
	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func defineApplicationGatewayIPConfiguration20220701(subnet *genruntime.ResourceReference) []network.ApplicationGatewayIPConfiguration_ApplicationGateway_SubResourceEmbedded {
	subResName := "app-gw-ip-config-1"
	appGtwFeIpConfig := []network.ApplicationGatewayIPConfiguration_ApplicationGateway_SubResourceEmbedded{
		{
			Name: to.Ptr(subResName),
			Subnet: &network.SubResource{
				Reference: subnet,
			},
		},
	}

	return appGtwFeIpConfig
}

// See https://learn.microsoft.com/en-us/azure/web-application-firewall/ag/quick-create-template
func defineApplicationGatewayWebApplicationFirewallPolicy(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *network2024.WebApplicationFirewallPolicy {
	firewallName := "app-gw-waf"
	firewall := &network2024.WebApplicationFirewallPolicy{
		ObjectMeta: tc.MakeObjectMeta(firewallName),
		Spec: network2024.WebApplicationFirewallPolicy_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			CustomRules: []network2024.WebApplicationFirewallCustomRule{
				{
					Name:     to.Ptr("rule1"),
					Priority: to.Ptr(100),
					RuleType: to.Ptr(network2024.WebApplicationFirewallCustomRule_RuleType_MatchRule),
					Action:   to.Ptr(network2024.WebApplicationFirewallCustomRule_Action_Block),
					MatchConditions: []network2024.MatchCondition{
						{
							MatchVariables: []network2024.MatchVariable{
								{
									VariableName: to.Ptr(network2024.MatchVariable_VariableName_RemoteAddr),
								},
							},
							Operator:         to.Ptr(network2024.MatchCondition_Operator_IPMatch),
							NegationConditon: to.Ptr(true),
							MatchValues: []string{
								"10.10.10.0/24",
							},
						},
					},
				},
			},
			PolicySettings: &network2024.PolicySettings{
				RequestBodyCheck:       to.Ptr(true),
				MaxRequestBodySizeInKb: to.Ptr(128),
				FileUploadLimitInMb:    to.Ptr(100),
				State:                  to.Ptr(network2024.PolicySettings_State_Enabled),
				Mode:                   to.Ptr(network2024.PolicySettings_Mode_Prevention),
			},
			ManagedRules: &network2024.ManagedRulesDefinition{
				ManagedRuleSets: []network2024.ManagedRuleSet{
					{
						RuleSetType:    to.Ptr("OWASP"),
						RuleSetVersion: to.Ptr("3.2"),
					},
				},
			},
		},
	}

	return firewall
}

func defineApplicationGatewayFrontendIPConfiguration20220701(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	appGatewayName string,
	subnet *genruntime.ResourceReference,
	publicIP *genruntime.ResourceReference,
) ([]network.ApplicationGatewayFrontendIPConfiguration, string) {
	appGtwFeIpName := "app-gw-fip-config-1"
	appGtwFeIpName2 := "app-gw-fip-config-2"
	appGtwFeIpConfig := []network.ApplicationGatewayFrontendIPConfiguration{
		{
			Name:                      to.Ptr(appGtwFeIpName),
			PrivateIPAddress:          to.Ptr("10.0.0.10"),
			PrivateIPAllocationMethod: to.Ptr(network.IPAllocationMethod_Static),
			Subnet: &network.SubResource{
				Reference: subnet,
			},
		},
		{
			Name: to.Ptr(appGtwFeIpName2),
			PublicIPAddress: &network.SubResource{
				Reference: publicIP,
			},
		},
	}
	appGtwARMID, err := getFrontendPortsARMID(tc, rg, appGatewayName, "frontendIPConfigurations", appGtwFeIpName)
	tc.Expect(err).To(BeNil())
	return appGtwFeIpConfig, appGtwARMID
}

func defineApplicationGatewayFrontendPort20220701(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayFrontendPort, string) {
	appGtwFePortName := "app-gw-fe-port-1"
	AppGtwFEPorts := []network.ApplicationGatewayFrontendPort{
		{
			Name: to.Ptr(appGtwFePortName),
			Port: to.Ptr(80),
		},
	}
	appGtwFePortsID, err := getFrontendPortsARMID(tc, rg, appGatewayName, "frontendPorts", appGtwFePortName)
	tc.Expect(err).To(BeNil())
	return AppGtwFEPorts, appGtwFePortsID
}

func defineApplicationGatewayBackendAddressPool20220701(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayBackendAddressPool, string) {
	subResname := "app-gw-be-pool-1"
	subRes := []network.ApplicationGatewayBackendAddressPool{
		{
			Name: to.Ptr(subResname),
			BackendAddresses: []network.ApplicationGatewayBackendAddress{
				{
					IpAddress: to.Ptr("10.0.1.1"),
				},
			},
		},
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, appGatewayName, "backendAddressPools", subResname)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewayBackendHttpSettings20220701(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayBackendHttpSettings, string) {
	subResname := "app-gw-be-http-setting-1"
	subRes := []network.ApplicationGatewayBackendHttpSettings{
		{
			Name:                           to.Ptr(subResname),
			Port:                           to.Ptr(8443),
			Protocol:                       to.Ptr(network.ApplicationGatewayProtocol_Http),
			CookieBasedAffinity:            to.Ptr(network.ApplicationGatewayBackendHttpSettingsPropertiesFormat_CookieBasedAffinity_Disabled),
			PickHostNameFromBackendAddress: to.Ptr(false),
		},
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, appGatewayName, "backendHttpSettingsCollection", subResname)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func getFrontendPortsARMID(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string, appGtwSubResourceTypeName string, subResName string) (string, error) {
	return genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		"applicationGateways",
		appGatewayName,
		appGtwSubResourceTypeName,
		subResName,
	)
}
