/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
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
	tc.CreateResourcesAndWait(publicIPAddress, vnet, subnet)

	tc.Expect(vnet.Status.Id).ToNot(BeNil())
	tc.Expect(subnet.Status.Id).ToNot(BeNil())

	appGtsListenerName := "app-gw-http-listener-1"
	appGtwFHttpListenerID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "httpListeners", appGtsListenerName)
	tc.Expect(err).To(BeNil())

	appGtwFIPConfig := defineApplicationGatewayIPConfiguration(tc.MakeReferenceFromResource(subnet))
	appGtwFeIpConfig, appGtwFeIpConfigID := defineApplicationGatewayFrontendIPConfiguration(tc, rg, appGatewayName, tc.MakeReferenceFromResource(subnet), tc.MakeReferenceFromResource(publicIPAddress))
	appGtwFEPorts, appGtwFePortsID := defineApplicationGatewayFrontendPort(tc, rg, appGatewayName)
	appGtwSslCerts, appGtwSslCertsID := defineApplicationGatewaySslCertificate(tc, rg, appGatewayName)
	appGtwSslProfiles, appGtwSslProfilesID := defineApplicationGatewaySslProfile(tc, rg, appGatewayName)
	appGtwBackendPools, appGtwBackendPoolsID := defineApplicationGatewayBackendAddressPool(tc, rg, appGatewayName)
	appGtwBackendHttpSettings, appGtwBackendHttpSettingsID := defineApplicationGatewayBackendHttpSettings(tc, rg, appGatewayName)
	appGtwWafConfig, appGtwWafConfigID := defineApplicationGatewayWebApplicationFirewallConfiguration(tc, rg, appGatewayName)

	tc.T.Log("SSL Profile ID: ", appGtwSslProfilesID)
	tc.T.Log("Firewall ARM ID", appGtwWafConfigID)

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

			FrontendPorts:   appGtwFEPorts,
			SslCertificates: appGtwSslCerts,
			SslProfiles:     appGtwSslProfiles,
			SslPolicy: to.Ptr(network.ApplicationGatewaySslPolicy{
				PolicyName: to.Ptr(network.PolicyNameEnum(network.PolicyNameEnum_AppGwSslPolicy20220101S)),
				PolicyType: to.Ptr(network.ApplicationGatewaySslPolicy_PolicyType(network.ApplicationGatewaySslPolicy_PolicyType_Predefined)),
			}),
			WebApplicationFirewallConfiguration: appGtwWafConfig,
			HttpListeners: []network.ApplicationGatewayHttpListener{
				{
					Name: to.Ptr(appGtsListenerName),
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
					Protocol:  to.Ptr(network.ApplicationGatewayProtocol(network.ApplicationGatewayProtocol_Https)),
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
					Name: to.Ptr("app-gtw-routing-rule-1"),
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
							ARMID: appGtwFHttpListenerID,
						},
					},
					RuleType: to.Ptr(network.ApplicationGatewayRequestRoutingRulePropertiesFormat_RuleType(network.ApplicationGatewayRequestRoutingRulePropertiesFormat_RuleType_Basic)),
					Priority: to.Ptr(100),
				},
			},
		},
	}
	tc.CreateResourceAndWait(applicationGateway)
	tc.Expect(applicationGateway.Status.Id).ToNot(BeNil())
	tc.Expect(applicationGateway.Status.Sku).ToNot(BeNil())
	tc.Expect(applicationGateway.Status.Sku.Tier).ToNot(BeNil())
	tc.Expect(*applicationGateway.Status.Sku.Tier).To(Equal(network.ApplicationGatewaySku_Tier_STATUS_WAF_V2))

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
			Name: to.Ptr(subResName),
			Subnet: &network.ApplicationGatewaySubResource{
				Reference: subnet,
			},
		},
	}

	return appGtwFeIpConfig
}

func defineApplicationGatewayWebApplicationFirewallConfiguration(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) (*network.ApplicationGatewayWebApplicationFirewallConfiguration, string) {
	subResname := "app-gw-waf-config-1"
	subRes := to.Ptr(network.ApplicationGatewayWebApplicationFirewallConfiguration{
		Enabled:        to.Ptr(true),
		FirewallMode:   to.Ptr(network.ApplicationGatewayWebApplicationFirewallConfiguration_FirewallMode(network.ApplicationGatewayWebApplicationFirewallConfiguration_FirewallMode_Detection)),
		RuleSetType:    to.Ptr("OWASP"),
		RuleSetVersion: to.Ptr("3.2"),
	})
	subResARMID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "applicationGatewayWebApplicationFirewallConfiguration", subResname)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewayFrontendIPConfiguration(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string, subnet *genruntime.ResourceReference, publicIP *genruntime.ResourceReference) ([]network.ApplicationGatewayFrontendIPConfiguration, string) {
	appGtwFeIpName := "app-gw-fip-config-1"
	appGtwFeIpName2 := "app-gw-fip-config-2"
	appGtwFeIpConfig := []network.ApplicationGatewayFrontendIPConfiguration{
		{
			Name:                      to.Ptr(appGtwFeIpName),
			PrivateIPAddress:          to.Ptr("10.0.0.10"),
			PrivateIPAllocationMethod: to.Ptr(network.IPAllocationMethod("Static")),
			Subnet: &network.ApplicationGatewaySubResource{
				Reference: subnet,
			},
		},
		{
			Name: to.Ptr(appGtwFeIpName2),
			PublicIPAddress: &network.ApplicationGatewaySubResource{
				Reference: publicIP,
			},
		},
	}
	appGtwARMID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "frontendIPConfigurations", appGtwFeIpName)
	tc.Expect(err).To(BeNil())
	return appGtwFeIpConfig, appGtwARMID
}

func defineApplicationGatewayFrontendPort(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayFrontendPort, string) {
	appGtwFePortName := "app-gw-fe-port-1"
	AppGtwFEPorts := []network.ApplicationGatewayFrontendPort{
		{
			Name: to.Ptr(appGtwFePortName),
			Port: to.Ptr(443),
		},
	}
	appGtwFePortsID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "frontendPorts", appGtwFePortName)
	tc.Expect(err).To(BeNil())
	return AppGtwFEPorts, appGtwFePortsID
}

func defineApplicationGatewaySslCertificate(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewaySslCertificate, string) {
	subResname := "app-gw-sslcert-1"
	subRes := []network.ApplicationGatewaySslCertificate{
		{
			//KeyVaultSecretId: to.Ptr("https://keyvaultname.vault.azure.net/secrets/secretname"),
			Name: to.Ptr(subResname),
			Data: 
		},
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "sslCertificates", subResname)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewaySslProfile(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewaySslProfile, string) {
	subResname := "app-gw-ssl-profile-1"
	subRes := []network.ApplicationGatewaySslProfile{
		{
			Name: to.Ptr(subResname),
			SslPolicy: &network.ApplicationGatewaySslPolicy{
				//MinProtocolVersion: to.Ptr(network.ProtocolsEnum(network.ProtocolsEnum_STATUS_TLSv1_2)),
				/*CipherSuites: []network.CipherSuitesEnum{
					network.CipherSuitesEnum(network.CipherSuitesEnum_TLS_RSA_WITH_AES_256_CBC_SHA256),
					network.CipherSuitesEnum(network.CipherSuitesEnum_TLS_RSA_WITH_AES_256_CBC_SHA),
				},*/
				PolicyName: to.Ptr(network.PolicyNameEnum(network.PolicyNameEnum_AppGwSslPolicy20220101S)),
				PolicyType: to.Ptr(network.ApplicationGatewaySslPolicy_PolicyType(network.ApplicationGatewaySslPolicy_PolicyType_Predefined)),
			},
		},
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "sslProfiles", subResname)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewayBackendAddressPool(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayBackendAddressPool, string) {
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
	subResARMID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "backendAddressPools", subResname)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func defineApplicationGatewayBackendHttpSettings(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGatewayName string) ([]network.ApplicationGatewayBackendHttpSettings, string) {
	subResname := "app-gw-be-http-setting-1"
	subRes := []network.ApplicationGatewayBackendHttpSettings{
		{
			Name:                           to.Ptr(subResname),
			Port:                           to.Ptr(8443),
			Protocol:                       to.Ptr(network.ApplicationGatewayProtocol(network.ApplicationGatewayProtocol_Https)),
			CookieBasedAffinity:            to.Ptr(network.ApplicationGatewayBackendHttpSettingsPropertiesFormat_CookieBasedAffinity_Disabled),
			PickHostNameFromBackendAddress: to.Ptr(false),
		},
	}
	subResARMID, err := getFrontendPortsARMID(tc, rg, "applicationGateways", appGatewayName, "backendHttpSettingsCollection", subResname)
	tc.Expect(err).To(BeNil())
	return subRes, subResARMID
}

func getFrontendPortsARMID(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, appGtwTypeName string, appGatewayName string, appGtwSubResourceTypeName string, subResName string) (string, error) {
	return genericarmclient.MakeResourceGroupScopeARMID(
		tc.AzureSubscription,
		rg.Name,
		"Microsoft.Network",
		appGtwTypeName,
		appGatewayName,
		appGtwSubResourceTypeName,
		subResName,
	)
}
