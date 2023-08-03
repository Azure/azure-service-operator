/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"log"
	"testing"

	. "github.com/onsi/gomega"

	network2020 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_ApplicationGateway_HTTPS_Listener_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	appGatewayName := "appgatewaytest"
	vnet := newVNet(tc, testcommon.AsOwner(rg), []string{"10.0.0.0/8"})
	subnet := &network2020.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network2020.VirtualNetworks_Subnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.0.0/24"),
			AzureName:     "AzureAppGtwSubnet", // The subnet must have this name in Azure for Bastion to work
		},
	}
	publicIPAddress := newPublicIp(tc, testcommon.AsOwner(rg))
	tc.CreateResourceAndWait(publicIPAddress)
	tc.CreateResourceAndWait(vnet)
	tc.Expect(vnet.Status.Id).ToNot(BeNil())
	tc.CreateResourceAndWait(subnet)
	tc.Expect(subnet.Status.Id).ToNot(BeNil())
	appGtsListnerName := "app-gw-http-listner-1"
	httpListnerParams := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "httpListeners",
		"param-val2": appGtsListnerName,
	}
	appGtwFHttpListnerID, err := getFrontendPortsARMID(tc, rg, httpListnerParams)
	tc.Expect(err).To(BeNil())
	appGtwFIPConfig := defineApplicationGatewayIPConfiguration(tc.MakeReferenceFromResource(subnet))
	appGtwFeIpConfig, appGtwFeIpConfigID := defineApplicationGatewayFrontendIPConfiguration(tc, rg, appGatewayName, tc.MakeReferenceFromResource(subnet), tc.MakeReferenceFromResource(publicIPAddress))
	appGtwFEPorts, appGtwFePortsID := defineApplicationGatewayFrontendPort(tc, rg, appGatewayName)
	appGtwSslCerts, appGtwSslCertsID := defineApplicationGatewaySslCertificate(tc, rg, appGatewayName)
	appGtwSslProfiles, appGtwSslProfilesID := defineApplicationGatewaySslProfile(tc, rg, appGatewayName)
	appGtwBackendPools, appGtwBackendPoolsID := defineApplicationGatewayBackendAddressPool(tc, rg, appGatewayName)
	appGtwBackendHttpSettings, appGtwBackendHttpSettingsID := defineApplicationGatewayBackendHttpSettings(tc, rg, appGatewayName)
	appGtwWafConfig, appGtwWafConfigID := defineApplicationGatewayWebApplicationFirewallConfiguration(tc, rg, appGatewayName)
	log.Println(appGtwSslProfilesID)
	log.Println(appGtwWafConfigID)
	applicationGateway := &network.ApplicationGateway{
		ObjectMeta: tc.MakeObjectMetaWithName(appGatewayName),
		Spec: network.ApplicationGateway_Spec{
			AutoscaleConfiguration: &network.ApplicationGatewayAutoscaleConfiguration{
				MaxCapacity: to.Ptr[int](3),
				MinCapacity: to.Ptr[int](1),
			},
			AzureName: appGatewayName,
			Location:  to.Ptr[string](rg.Location()),
			Owner:     testcommon.AsOwner(rg),
			Sku: &network.ApplicationGatewaySku{
				Name: to.Ptr[network.ApplicationGatewaySku_Name](network.ApplicationGatewaySku_Name(network.ApplicationGatewaySku_Name_STATUS_WAF_V2)),
				Tier: to.Ptr[network.ApplicationGatewaySku_Tier](network.ApplicationGatewaySku_Tier(network.ApplicationGatewaySku_Tier_STATUS_WAF_V2)),
			},
			GatewayIPConfigurations:  appGtwFIPConfig,
			FrontendIPConfigurations: appGtwFeIpConfig,

			FrontendPorts:   appGtwFEPorts,
			SslCertificates: appGtwSslCerts,
			SslProfiles:     appGtwSslProfiles,
			SslPolicy: to.Ptr[network.ApplicationGatewaySslPolicy](network.ApplicationGatewaySslPolicy{
				PolicyName: to.Ptr[network.PolicyNameEnum](network.PolicyNameEnum(network.PolicyNameEnum_AppGwSslPolicy20220101S)),
				PolicyType: to.Ptr[network.ApplicationGatewaySslPolicy_PolicyType](network.ApplicationGatewaySslPolicy_PolicyType(network.ApplicationGatewaySslPolicy_PolicyType_Predefined)),
			}),
			WebApplicationFirewallConfiguration: appGtwWafConfig,
			HttpListeners: []network.ApplicationGatewayHttpListener{
				{
					Name: to.Ptr[string](appGtsListnerName),
					FrontendIPConfiguration: &network.ApplicationGatewaySubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwFeIpConfigID,
						},
					},
					FrontendPort: &network.ApplicationGatewaySubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwFePortsID,
						},
					},
					Protocol:  to.Ptr[network.ApplicationGatewayProtocol](network.ApplicationGatewayProtocol(network.ApplicationGatewayProtocol_Https)),
					HostNames: []string{"test.contoso.com"},
					SslCertificate: &network.ApplicationGatewaySubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwSslCertsID,
						},
					},
				},
			},
			BackendAddressPools:           appGtwBackendPools,
			BackendHttpSettingsCollection: appGtwBackendHttpSettings,
			RequestRoutingRules: []network.ApplicationGatewayRequestRoutingRule{
				{
					Name: to.Ptr[string]("app-gtw-routing-rule-1"),
					BackendAddressPool: &network.ApplicationGatewaySubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwBackendPoolsID,
						},
					},
					BackendHttpSettings: &network.ApplicationGatewaySubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwBackendHttpSettingsID,
						},
					},
					HttpListener: &network.ApplicationGatewaySubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: appGtwFHttpListnerID,
						},
					},
					RuleType: to.Ptr[network.ApplicationGatewayRequestRoutingRulePropertiesFormat_RuleType](network.ApplicationGatewayRequestRoutingRulePropertiesFormat_RuleType(network.ApplicationGatewayRequestRoutingRulePropertiesFormat_RuleType_Basic)),
					Priority: to.Ptr[int](100),
				},
			},
		},
	}
	tc.CreateResourceAndWait(applicationGateway)
	tc.Expect(applicationGateway.Status.Id).ToNot(BeNil())
	tc.Expect(applicationGateway.Status.Sku.Tier).To(Equal(network.ApplicationGatewaySku_Tier_STATUS_WAF_V2))
	armId := *applicationGateway.Status.Id
	tc.DeleteResourceAndWait(applicationGateway)
	tc.DeleteResourceAndWait(publicIPAddress)
	tc.DeleteResourceAndWait(subnet)
	tc.DeleteResourceAndWait(vnet)
	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func defineApplicationGatewayIPConfiguration(subnet *genruntime.ResourceReference) []network.ApplicationGatewayIPConfiguration_ApplicationGateway_SubResourceEmbedded {
	subResName := "app-gw-ip-config-1"
	appGtwFeIpConfig := []network.ApplicationGatewayIPConfiguration_ApplicationGateway_SubResourceEmbedded{
		{
			Name: to.Ptr[string](subResName),
			Subnet: &network.ApplicationGatewaySubResource{
				Reference: subnet,
			},
		},
	}

	return appGtwFeIpConfig
}

func defineApplicationGatewayWebApplicationFirewallConfiguration(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) (*network.ApplicationGatewayWebApplicationFirewallConfiguration, string) {
	subResname := "app-gw-waf-config-1"
	subRes := to.Ptr[network.ApplicationGatewayWebApplicationFirewallConfiguration](network.ApplicationGatewayWebApplicationFirewallConfiguration{
		Enabled:        to.Ptr[bool](true),
		FirewallMode:   to.Ptr[network.ApplicationGatewayWebApplicationFirewallConfiguration_FirewallMode](network.ApplicationGatewayWebApplicationFirewallConfiguration_FirewallMode(network.ApplicationGatewayWebApplicationFirewallConfiguration_FirewallMode_Detection)),
		RuleSetType:    to.Ptr[string]("OWASP"),
		RuleSetVersion: to.Ptr[string]("3.2"),
	})
	armid_params := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "applicationGatewayWebApplicationFirewallConfiguration",
		"param-val2": subResname,
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, armid_params)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewayFrontendIPConfiguration(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string, subnet *genruntime.ResourceReference, publicIP *genruntime.ResourceReference) ([]network.ApplicationGatewayFrontendIPConfiguration, string) {
	appGtwFeIpName := "app-gw-fip-config-1"
	appGtwFeIpName2 := "app-gw-fip-config-2"
	appGtwFeIpConfig := []network.ApplicationGatewayFrontendIPConfiguration{
		{
			Name:                      to.Ptr[string](appGtwFeIpName),
			PrivateIPAddress:          to.Ptr[string]("10.0.0.10"),
			PrivateIPAllocationMethod: to.Ptr[network.IPAllocationMethod](network.IPAllocationMethod("Static")),
			Subnet: &network.ApplicationGatewaySubResource{
				Reference: subnet,
			},
		},
		{
			Name: to.Ptr[string](appGtwFeIpName2),
			PublicIPAddress: &network.ApplicationGatewaySubResource{
				Reference: publicIP,
			},
		},
	}
	armid_params := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "frontendIPConfigurations",
		"param-val2": appGtwFeIpName,
	}
	appGtwARMID, err := getFrontendPortsARMID(tc, rg, armid_params)
	tc.Expect(err).To(BeNil())
	return appGtwFeIpConfig, appGtwARMID
}

func defineApplicationGatewayFrontendPort(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayFrontendPort, string) {
	appGtwFePortName := "app-gw-fe-port-1"
	AppGtwFEPorts := []network.ApplicationGatewayFrontendPort{
		{
			Name: to.Ptr[string](appGtwFePortName),
			Port: to.Ptr[int](443),
		},
	}
	armid_params := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "frontendPorts",
		"param-val2": appGtwFePortName,
	}
	appGtwFePortsID, err := getFrontendPortsARMID(tc, rg, armid_params)
	tc.Expect(err).To(BeNil())
	return AppGtwFEPorts, appGtwFePortsID
}

func defineApplicationGatewaySslCertificate(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewaySslCertificate, string) {
	subResname := "app-gw-sslcert-1"
	subRes := []network.ApplicationGatewaySslCertificate{
		{
			//KeyVaultSecretId: to.Ptr[string]("https://keyvaultname.vault.azure.net/secrets/secretname"),
			Name: to.Ptr[string](subResname),
			Data: to.Ptr[string](`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`),
		},
	}
	armid_params := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "sslCertificates",
		"param-val2": subResname,
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, armid_params)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewaySslProfile(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewaySslProfile, string) {
	subResname := "app-gw-ssl-profile-1"
	subRes := []network.ApplicationGatewaySslProfile{
		{
			Name: to.Ptr[string](subResname),
			SslPolicy: &network.ApplicationGatewaySslPolicy{
				//MinProtocolVersion: to.Ptr[network.ProtocolsEnum](network.ProtocolsEnum(network.ProtocolsEnum_STATUS_TLSv1_2)),
				/*CipherSuites: []network.CipherSuitesEnum{
					network.CipherSuitesEnum(network.CipherSuitesEnum_TLS_RSA_WITH_AES_256_CBC_SHA256),
					network.CipherSuitesEnum(network.CipherSuitesEnum_TLS_RSA_WITH_AES_256_CBC_SHA),
				},*/
				PolicyName: to.Ptr[network.PolicyNameEnum](network.PolicyNameEnum(network.PolicyNameEnum_AppGwSslPolicy20220101S)),
				PolicyType: to.Ptr[network.ApplicationGatewaySslPolicy_PolicyType](network.ApplicationGatewaySslPolicy_PolicyType(network.ApplicationGatewaySslPolicy_PolicyType_Predefined)),
			},
		},
	}
	armid_params := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "sslProfiles",
		"param-val2": subResname,
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, armid_params)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewayBackendAddressPool(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayBackendAddressPool, string) {
	subResname := "app-gw-be-pool-1"
	subRes := []network.ApplicationGatewayBackendAddressPool{
		{
			Name: to.Ptr[string](subResname),
			BackendAddresses: []network.ApplicationGatewayBackendAddress{
				{
					IpAddress: to.Ptr[string]("10.0.1.1"),
				},
			},
		},
	}
	armid_params := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "backendAddressPools",
		"param-val2": subResname,
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, armid_params)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewayBackendHttpSettings(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayBackendHttpSettings, string) {
	subResname := "app-gw-be-pool-1"
	subRes := []network.ApplicationGatewayBackendHttpSettings{
		{
			Name:                           to.Ptr[string](subResname),
			Port:                           to.Ptr[int](8443),
			Protocol:                       to.Ptr[network.ApplicationGatewayProtocol](network.ApplicationGatewayProtocol(network.ApplicationGatewayProtocol_Https)),
			CookieBasedAffinity:            to.Ptr[network.ApplicationGatewayBackendHttpSettingsPropertiesFormat_CookieBasedAffinity](network.ApplicationGatewayBackendHttpSettingsPropertiesFormat_CookieBasedAffinity_Disabled),
			PickHostNameFromBackendAddress: to.Ptr[bool](false),
		},
	}
	armid_params := map[string]string{
		"param-key1": "applicationGateways",
		"param-val1": appGatewayName,
		"param-key2": "backendHttpSettingsCollection",
		"param-val2": subResname,
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, armid_params)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func getFrontendPortsARMID(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, params map[string]string) (string, error) {
	return genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		params["param-key1"],
		params["param-val1"],
		params["param-key2"],
		params["param-val2"],
	)
}
