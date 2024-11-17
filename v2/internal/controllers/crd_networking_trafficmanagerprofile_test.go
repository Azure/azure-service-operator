/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network2020 "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220401"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_TrafficManagerProfile(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	configName := "profile-config"
	configKey := "fqdn"
	tmp := &network.TrafficManagerProfile{
		ObjectMeta: tc.MakeObjectMeta("profile"),
		Spec: network.TrafficManagerProfile_Spec{
			DnsConfig: &network.DnsConfig{
				RelativeName: to.Ptr("aso-test"),
			},
			Location: to.Ptr("global"),
			MonitorConfig: &network.MonitorConfig{
				Protocol: to.Ptr(network.MonitorConfig_Protocol_TCP),
				Port:     to.Ptr(443),
			},
			OperatorSpec: &network.TrafficManagerProfileOperatorSpec{
				ConfigMaps: &network.TrafficManagerProfileOperatorConfigMaps{
					DnsConfigFqdn: &genruntime.ConfigMapDestination{
						Name: configName,
						Key:  configKey,
					},
				},
			},
			Owner:                testcommon.AsOwner(rg),
			TrafficRoutingMethod: to.Ptr(network.ProfileProperties_TrafficRoutingMethod_Performance),
		},
	}

	tc.CreateResourceAndWait(tmp)

	tc.Expect(tmp.Status.Id).ToNot(BeNil())
	tc.ExpectConfigMapHasKeysAndValues(configName, configKey, *tmp.Status.DnsConfig.Fqdn)
	armId := *tmp.Status.Id

	// Perform a simple patch
	old := tmp.DeepCopy()
	tmp.Spec.Tags = map[string]string{
		"foo": "bar",
	}
	tc.PatchResourceAndWait(old, tmp)
	tc.Expect(tmp.Status.Tags).To(HaveKey("foo"))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Test_TrafficManagerProfiles_AzureEndpoint",
			Test: func(tc *testcommon.KubePerTestContext) {
				Networking_TrafficManagerProfiles_AzureEndpoint(tc, rg, tmp)
			},
		},
		testcommon.Subtest{
			Name: "Test_TrafficManagerProfiles_ExternalEndpoint",
			Test: func(tc *testcommon.KubePerTestContext) {
				Networking_TrafficManagerProfiles_ExternalEndpoint(tc, tmp)
			},
		},
		testcommon.Subtest{
			Name: "Test_TrafficManagerProfiles_NestedEndpoint",
			Test: func(tc *testcommon.KubePerTestContext) {
				Networking_TrafficManagerProfiles_NestedEndpoint(tc, rg, tmp)
			},
		})

	tc.DeleteResourceAndWait(tmp)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func Networking_TrafficManagerProfiles_NestedEndpoint(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, tmp *network.TrafficManagerProfile) {
	nestedTmp := &network.TrafficManagerProfile{
		ObjectMeta: tc.MakeObjectMeta("profile"),
		Spec: network.TrafficManagerProfile_Spec{
			DnsConfig: &network.DnsConfig{
				RelativeName: to.Ptr("nested-test"),
			},
			Location: to.Ptr("global"),
			MonitorConfig: &network.MonitorConfig{
				Protocol: to.Ptr(network.MonitorConfig_Protocol_TCP),
				Port:     to.Ptr(443),
			},
			Owner:                testcommon.AsOwner(rg),
			TrafficRoutingMethod: to.Ptr(network.ProfileProperties_TrafficRoutingMethod_Performance),
		},
	}

	tc.CreateResourceAndWait(nestedTmp)
	defer tc.DeleteResourceAndWait(nestedTmp)

	endpoint := &network.TrafficManagerProfilesNestedEndpoint{
		ObjectMeta: tc.MakeObjectMeta("nested-ep"),
		Spec: network.TrafficManagerProfilesNestedEndpoint_Spec{
			EndpointLocation:        tc.AzureRegion,
			Owner:                   testcommon.AsOwner(tmp),
			TargetResourceReference: tc.MakeReferenceFromResource(nestedTmp),
		},
	}

	tc.CreateResourceAndWait(endpoint)
	tc.DeleteResourcesAndWait(endpoint)
}

func Networking_TrafficManagerProfiles_ExternalEndpoint(tc *testcommon.KubePerTestContext, tmp *network.TrafficManagerProfile) {
	endpoint := &network.TrafficManagerProfilesExternalEndpoint{
		ObjectMeta: tc.MakeObjectMeta("external-ep"),
		Spec: network.TrafficManagerProfilesExternalEndpoint_Spec{
			AlwaysServe:      to.Ptr(network.EndpointProperties_AlwaysServe_Enabled),
			EndpointLocation: tc.AzureRegion,
			Owner:            testcommon.AsOwner(tmp),
			Target:           to.Ptr("contoso.com"),
		},
	}

	tc.CreateResourceAndWait(endpoint)
	tc.DeleteResourcesAndWait(endpoint)
}

func Networking_TrafficManagerProfiles_AzureEndpoint(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, tmp *network.TrafficManagerProfile) {
	publicIp := newPublicIP20201101(tc, testcommon.AsOwner(rg))
	publicIp.Spec.DnsSettings = &network2020.PublicIPAddressDnsSettings{
		DomainNameLabel: to.Ptr("mydemoapp"),
	}
	tc.CreateResourceAndWait(publicIp)

	endpoint := &network.TrafficManagerProfilesAzureEndpoint{
		ObjectMeta: tc.MakeObjectMeta("azure-ep"),
		Spec: network.TrafficManagerProfilesAzureEndpoint_Spec{
			AlwaysServe:             to.Ptr(network.EndpointProperties_AlwaysServe_Enabled),
			Owner:                   testcommon.AsOwner(tmp),
			TargetResourceReference: tc.MakeReferenceFromResource(publicIp),
		},
	}

	tc.CreateResourceAndWait(endpoint)
	tc.DeleteResourcesAndWait(endpoint)
}
