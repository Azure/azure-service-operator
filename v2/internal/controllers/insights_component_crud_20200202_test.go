/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	insights20171001 "github.com/Azure/azure-service-operator/v2/api/insights/v1api20171001"
	insightswebtest20180501preview "github.com/Azure/azure-service-operator/v2/api/insights/v1api20180501preview"
	insights20200202 "github.com/Azure/azure-service-operator/v2/api/insights/v1api20200202"
	insightswebtest20220615 "github.com/Azure/azure-service-operator/v2/api/insights/v1api20220615"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Insights_Component_v20200202_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	component := newAppInsightsComponent(tc, rg)
	tc.CreateResourceAndWait(component)

	tc.Expect(component.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(component.Status.Kind).To(Equal(to.Ptr("web")))
	tc.Expect(component.Status.Id).ToNot(BeNil())
	armId := *component.Status.Id

	// Perform a simple patch.
	old := component.DeepCopy()
	component.Spec.RetentionInDays = to.Ptr(60)
	tc.PatchResourceAndWait(old, component)
	tc.Expect(component.Status.RetentionInDays).To(Equal(to.Ptr(60)))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Insights WebTest 20180501preview CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Insights_WebTest_20180501preview_CRUD(tc, rg, component)
			},
		},
		testcommon.Subtest{
			Name: "Insights WebTest 20220615 CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Insights_WebTest_20220615_CRUD(tc, rg, component)
			},
		},
		testcommon.Subtest{
			Name: "Insights PricingPlan 20171001 CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Insights_PricingPlan_20171001_CRUD(tc, rg, component)
			},
		},
	)

	tc.DeleteResourceAndWait(component)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(insights20200202.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Insights_WebTest_20220615_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, component *insights20200202.Component) {
	horribleHiddenLink := fmt.Sprintf("hidden-link:%s", to.Value(component.Status.Id))

	horribleTags := map[string]string{
		horribleHiddenLink: "Resource",
	}

	// Create a webtest
	om := tc.MakeObjectMeta("webtest")

	kind := insightswebtest20220615.WebTestProperties_Kind_Standard
	webtest := &insightswebtest20220615.Webtest{
		ObjectMeta: om,
		Spec: insightswebtest20220615.Webtest_Spec{
			Location:           tc.AzureRegion,
			Owner:              testcommon.AsOwner(rg),
			SyntheticMonitorId: &om.Name,
			Tags:               horribleTags,
			Name:               to.Ptr("mywebtest"),
			Enabled:            to.Ptr(true),
			Frequency:          to.Ptr(300),
			Kind:               &kind,
			Locations: []insightswebtest20220615.WebTestGeolocation{
				{
					Id: to.Ptr("us-ca-sjc-azr"), // This is US west...
				},
			},
			Request: &insightswebtest20220615.WebTestProperties_Request{
				HttpVerb:   to.Ptr("GET"),
				RequestUrl: to.Ptr("https://github.com/Azure/azure-service-operator"),
			},
			ValidationRules: &insightswebtest20220615.WebTestProperties_ValidationRules{
				ExpectedHttpStatusCode:        to.Ptr(200),
				SSLCheck:                      to.Ptr(true),
				SSLCertRemainingLifetimeCheck: to.Ptr(7),
			},
		},
	}

	tc.CreateResourceAndWait(webtest)

	expectedKind := insightswebtest20220615.WebTestProperties_Kind_STATUS_Standard
	tc.Expect(webtest.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(webtest.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(webtest.Status.Id).ToNot(BeNil())

	// Perform a simple patch.
	old := webtest.DeepCopy()
	webtest.Spec.Enabled = to.Ptr(false)
	tc.PatchResourceAndWait(old, webtest)
	tc.Expect(webtest.Status.Enabled).To(Equal(to.Ptr(false)))

	armId := *webtest.Status.Id
	tc.DeleteResourceAndWait(webtest)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(insightswebtest20220615.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Insights_PricingPlan_20171001_CRUD(
	tc *testcommon.KubePerTestContext,
	rg *resources.ResourceGroup,
	component *insights20200202.Component,
) {
	plan := &insights20171001.PricingPlan{
		ObjectMeta: tc.MakeObjectMeta("plan"),
		Spec: insights20171001.PricingPlan_Spec{
			Owner:                          testcommon.AsOwner(component),
			PlanType:                       to.Ptr("Basic"),
			StopSendNotificationWhenHitCap: to.Ptr(true),
		},
	}

	tc.CreateResourceAndWait(plan)

	tc.Expect(plan.Status.PlanType).NotTo(BeNil())
	tc.Expect(*plan.Status.PlanType).To(Equal("Basic"))

	// Perform a simple patch.
	old := plan.DeepCopy()
	plan.Spec.StopSendNotificationWhenHitCap = to.Ptr(false)
	tc.PatchResourceAndWait(old, plan)

	tc.Expect(plan.Status.StopSendNotificationWhenHitCap).NotTo(BeNil())
	tc.Expect(*plan.Status.StopSendNotificationWhenHitCap).To(BeFalse())

	// Pricing plans can't be deleted
}

func Insights_WebTest_20180501preview_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, component *insights20200202.Component) {
	horribleHiddenLink := fmt.Sprintf("hidden-link:%s", to.Value(component.Status.Id))

	horribleTags := map[string]string{
		horribleHiddenLink: "Resource",
	}

	// Create a webtest
	om := tc.MakeObjectMeta("webtest")

	kind := insightswebtest20180501preview.WebTestProperties_Kind_Standard
	webtest := &insightswebtest20180501preview.Webtest{
		ObjectMeta: om,
		Spec: insightswebtest20180501preview.Webtest_Spec{
			Location:           tc.AzureRegion,
			Owner:              testcommon.AsOwner(rg),
			SyntheticMonitorId: &om.Name,
			Tags:               horribleTags,
			Name:               to.Ptr("mywebtest"),
			Enabled:            to.Ptr(true),
			Frequency:          to.Ptr(300),
			Kind:               &kind,
			Locations: []insightswebtest20180501preview.WebTestGeolocation{
				{
					Id: to.Ptr("us-ca-sjc-azr"), // This is US west...
				},
			},
			Request: &insightswebtest20180501preview.WebTestProperties_Request{
				HttpVerb:   to.Ptr("GET"),
				RequestUrl: to.Ptr("https://github.com/Azure/azure-service-operator"),
			},
			ValidationRules: &insightswebtest20180501preview.WebTestProperties_ValidationRules{
				ExpectedHttpStatusCode:        to.Ptr(200),
				SSLCheck:                      to.Ptr(true),
				SSLCertRemainingLifetimeCheck: to.Ptr(7),
			},
		},
	}

	tc.CreateResourceAndWait(webtest)

	expectedKind := insightswebtest20180501preview.WebTestProperties_Kind_STATUS_Standard
	tc.Expect(webtest.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(webtest.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(webtest.Status.Id).ToNot(BeNil())
	armId := *webtest.Status.Id

	// Perform a simple patch.
	old := webtest.DeepCopy()
	webtest.Spec.Enabled = to.Ptr(false)
	tc.PatchResourceAndWait(old, webtest)
	tc.Expect(webtest.Status.Enabled).To(Equal(to.Ptr(false)))

	tc.DeleteResourceAndWait(webtest)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(insightswebtest20180501preview.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Test_Insights_Component_ExportConfigMap(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	component := newAppInsightsComponent(tc, rg)
	tc.CreateResourceAndWait(component)

	tc.Expect(component.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(component.Status.Kind).To(Equal(to.Ptr("web")))
	tc.Expect(component.Status.Id).ToNot(BeNil())
	armId := *component.Status.Id

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "ConfigValuesWrittenToSameConfigMap",
			Test: func(tc *testcommon.KubePerTestContext) {
				Component_ConfigValuesWrittenToSameConfigMap(tc, component)
			},
		},
		testcommon.Subtest{
			Name: "ConfigValuesWrittenToDifferentConfigMap",
			Test: func(tc *testcommon.KubePerTestContext) {
				Component_ConfigValuesWrittenToDifferentConfigMap(tc, component)
			},
		})

	tc.DeleteResourceAndWait(component)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(insights20200202.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Component_ConfigValuesWrittenToSameConfigMap(tc *testcommon.KubePerTestContext, component *insights20200202.Component) {
	old := component.DeepCopy()
	componentConfigMap := "component-config"

	component.Spec.OperatorSpec = &insights20200202.ComponentOperatorSpec{
		ConfigMaps: &insights20200202.ComponentOperatorConfigMaps{
			ConnectionString:   &genruntime.ConfigMapDestination{Name: componentConfigMap, Key: "connectionString"},
			InstrumentationKey: &genruntime.ConfigMapDestination{Name: componentConfigMap, Key: "instrumentationKey"},
		},
	}
	tc.PatchResourceAndWait(old, component)

	tc.ExpectConfigMapHasKeysAndValues(
		componentConfigMap,
		"connectionString", *component.Status.ConnectionString,
		"instrumentationKey", *component.Status.InstrumentationKey)
}

func Component_ConfigValuesWrittenToDifferentConfigMap(tc *testcommon.KubePerTestContext, component *insights20200202.Component) {
	old := component.DeepCopy()
	componentConfigMap1 := "component-config1"
	componentConfigMap2 := "component-config2"

	component.Spec.OperatorSpec = &insights20200202.ComponentOperatorSpec{
		ConfigMaps: &insights20200202.ComponentOperatorConfigMaps{
			ConnectionString:   &genruntime.ConfigMapDestination{Name: componentConfigMap1, Key: "connectionString"},
			InstrumentationKey: &genruntime.ConfigMapDestination{Name: componentConfigMap2, Key: "instrumentationKey"},
		},
	}
	tc.PatchResourceAndWait(old, component)

	tc.ExpectConfigMapHasKeysAndValues(componentConfigMap1, "connectionString", *component.Status.ConnectionString)
	tc.ExpectConfigMapHasKeysAndValues(componentConfigMap2, "instrumentationKey", *component.Status.InstrumentationKey)
}

func newAppInsightsComponent(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *insights20200202.Component {
	// Create a component
	applicationType := insights20200202.ApplicationInsightsComponentProperties_Application_Type_Other
	component := &insights20200202.Component{
		ObjectMeta: tc.MakeObjectMeta("component"),
		Spec: insights20200202.Component_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			// According to their documentation you can set anything here, it's ignored.
			Application_Type: &applicationType,
			Kind:             to.Ptr("web"),
		},
	}
	return component
}
