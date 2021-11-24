/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1alpha1api20200202"
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

	tc.DeleteResourceAndWait(component)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(insights.ComponentsSpecAPIVersion20200202))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
