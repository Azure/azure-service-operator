/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	insightswebtest "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20180501preview"
	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20200202"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1alpha1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Insights_Component_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a component
	component := &insights.Component{
		ObjectMeta: tc.MakeObjectMeta("component"),
		Spec: insights.Components_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			// According to their documentation you can set anything here, it's ignored.
			ApplicationType: insights.ApplicationInsightsComponentPropertiesApplicationTypeOther,
			Kind:            "web",
		},
	}

	tc.CreateResourceAndWait(component)

	tc.Expect(component.Status.Location).To(Equal(&tc.AzureRegion))
	tc.Expect(component.Status.Kind).To(Equal(to.StringPtr("web")))
	tc.Expect(component.Status.Id).ToNot(BeNil())
	armId := *component.Status.Id

	// Perform a simple patch.
	old := component.DeepCopy()
	component.Spec.RetentionInDays = to.IntPtr(60)
	tc.Patch(old, component)

	objectKey := client.ObjectKeyFromObject(component)

	// Ensure state eventually gets updated in k8s from change in Azure.
	tc.Eventually(func() *int {
		var updated insights.Component
		tc.GetResource(objectKey, &updated)
		return updated.Status.RetentionInDays
	}).Should(BeEquivalentTo(to.IntPtr(60)))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Insights WebTest CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				Insights_WebTest_CRUD(tc, rg, component)
			},
		})

	tc.DeleteResourceAndWait(component)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(insights.ComponentsSpecAPIVersion20200202))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Insights_WebTest_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, component *insights.Component) {
	horribleHiddenLink := fmt.Sprintf("hidden-link:%s", to.String(component.Status.Id))

	// Create a webtest
	webtest := &insightswebtest.Webtest{
		ObjectMeta: tc.MakeObjectMeta("webtest"),
		Spec: insightswebtest.Webtests_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Tags: map[string]string{
				horribleHiddenLink: "Resource",
			},
			Name:      "mywebtest",
			Enabled:   to.BoolPtr(true),
			Frequency: to.IntPtr(300),
			Kind:      insightswebtest.WebTestPropertiesKindStandard,
			Locations: []insightswebtest.WebTestGeolocation{
				{
					Id: to.StringPtr("us-ca-sjc-azr"), // This is US west...
				},
			},
			Request: &insightswebtest.WebTestPropertiesRequest{
				HttpVerb:   to.StringPtr("GET"),
				RequestUrl: to.StringPtr("https://github.com/Azure/azure-service-operator"),
			},
			ValidationRules: &insightswebtest.WebTestPropertiesValidationRules{
				ExpectedHttpStatusCode:        to.IntPtr(200),
				SSLCheck:                      to.BoolPtr(true),
				SSLCertRemainingLifetimeCheck: to.IntPtr(7),
			},
		},
	}

	tc.CreateResourceAndWait(webtest)

	expectedKind := insightswebtest.WebTestPropertiesStatusKindStandard
	tc.Expect(webtest.Status.Location).To(Equal(&tc.AzureRegion))
	tc.Expect(webtest.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(webtest.Status.Id).ToNot(BeNil())
	armId := *webtest.Status.Id

	// Perform a simple patch.
	old := webtest.DeepCopy()
	webtest.Spec.Enabled = to.BoolPtr(false)
	tc.Patch(old, webtest)

	objectKey := client.ObjectKeyFromObject(webtest)

	// Ensure state eventually gets updated in k8s from change in Azure.
	tc.Eventually(func() *bool {
		var updated insightswebtest.Webtest
		tc.GetResource(objectKey, &updated)
		return updated.Status.Enabled
	}).Should(BeEquivalentTo(to.BoolPtr(false)))

	tc.DeleteResourceAndWait(webtest)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(insightswebtest.WebtestsSpecAPIVersion20180501Preview))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())

}
