/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	v1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	insightswebtest "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20180501preview"
	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1beta20200202"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Insights_Component_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a component
	applicationType := insights.ApplicationInsightsComponentProperties_Application_Type_Other
	component := &insights.Component{
		ObjectMeta: tc.MakeObjectMeta("component"),
		Spec: insights.Component_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			// According to their documentation you can set anything here, it's ignored.
			Application_Type: &applicationType,
			Kind:             to.StringPtr("web"),
		},
	}

	tc.CreateResourceAndWait(component)

	tc.Expect(component.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(component.Status.Kind).To(Equal(to.StringPtr("web")))
	tc.Expect(component.Status.Id).ToNot(BeNil())
	armId := *component.Status.Id

	// Perform a simple patch.
	old := component.DeepCopy()
	component.Spec.RetentionInDays = to.IntPtr(60)
	tc.PatchResourceAndWait(old, component)
	tc.Expect(component.Status.RetentionInDays).To(Equal(to.IntPtr(60)))

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
		string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func Insights_WebTest_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, component *insights.Component) {
	horribleHiddenLink := fmt.Sprintf("hidden-link:%s", to.String(component.Status.Id))

	horribleTags := map[string]string{
		horribleHiddenLink: "Resource",
	}

	horribleJSON, err := json.Marshal(horribleTags)
	if err != nil {
		panic(err)
	}

	// Create a webtest
	om := tc.MakeObjectMeta("webtest")

	kind := insightswebtest.WebTestProperties_Kind_Multistep //!! verify this is the correct value
	webtest := &insightswebtest.Webtest{
		ObjectMeta: om,
		Spec: insightswebtest.Webtest_Spec{
			Location:           tc.AzureRegion,
			Owner:              testcommon.AsOwner(rg),
			SyntheticMonitorId: &om.Name,
			Tags:               &v1.JSON{horribleJSON},
			Name:               to.StringPtr("mywebtest"),
			Enabled:            to.BoolPtr(true),
			Frequency:          to.IntPtr(300),
			Kind:               &kind,
			Locations: []insightswebtest.WebTestGeolocation{
				{
					Id: to.StringPtr("us-ca-sjc-azr"), // This is US west...
				},
			},
			Request: &insightswebtest.WebTestProperties_Request{
				HttpVerb:   to.StringPtr("GET"),
				RequestUrl: to.StringPtr("https://github.com/Azure/azure-service-operator"),
			},
			ValidationRules: &insightswebtest.WebTestProperties_ValidationRules{
				ExpectedHttpStatusCode:        to.IntPtr(200),
				SSLCheck:                      to.BoolPtr(true),
				SSLCertRemainingLifetimeCheck: to.IntPtr(7),
			},
		},
	}

	tc.CreateResourceAndWait(webtest)

	expectedKind := insightswebtest.WebTestProperties_Kind_STATUS_Standard
	tc.Expect(webtest.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(webtest.Status.Kind).To(Equal(&expectedKind))
	tc.Expect(webtest.Status.Id).ToNot(BeNil())
	armId := *webtest.Status.Id

	// Perform a simple patch.
	old := webtest.DeepCopy()
	webtest.Spec.Enabled = to.BoolPtr(false)
	tc.PatchResourceAndWait(old, webtest)
	tc.Expect(webtest.Status.Enabled).To(Equal(to.BoolPtr(false)))

	tc.DeleteResourceAndWait(webtest)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(insightswebtest.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
