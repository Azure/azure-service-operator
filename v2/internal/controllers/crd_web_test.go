/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/api/web/v1beta20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/onsi/gomega"
)

func Test_Web_ServerFarm_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	serverfarm := newServerFarm(tc, rg)

	tc.CreateResourceAndWait(serverfarm)

	armId := *serverfarm.Status.Id
	old := serverfarm.DeepCopy()
	serverfarm.Spec.PerSiteScaling = to.BoolPtr(true)
	tc.PatchResourceAndWait(old, serverfarm)
	tc.Expect(serverfarm.Status.PerSiteScaling).To(gomega.Equal(to.BoolPtr(true)))

	tc.DeleteResourcesAndWait(serverfarm)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(v1beta20220301.APIVersion_Value))
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())

}

func Test_Web_Site_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	serverfarm := newServerFarm(tc, rg)

	site := &v1beta20220301.Site{
		ObjectMeta: tc.MakeObjectMeta("function"),
		Spec: v1beta20220301.Sites_Spec{
			Enabled:             to.BoolPtr(true),
			Owner:               testcommon.AsOwner(rg),
			Location:            tc.AzureRegion,
			ServerFarmReference: tc.MakeReferenceFromResource(serverfarm),
		},
	}
	tc.ExportAsSample(serverfarm)

	tc.CreateResourceAndWait(serverfarm)
	tc.CreateResourceAndWait(site)

	armId := *site.Status.Id
	old := site.DeepCopy()
	site.Spec.Enabled = to.BoolPtr(false)
	tc.PatchResourceAndWait(old, site)
	tc.Expect(site.Status.Enabled).To(gomega.Equal(to.BoolPtr(false)))

	tc.DeleteResourceAndWait(site)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(v1beta20220301.APIVersion_Value))
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())
}

func newServerFarm(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *v1beta20220301.Serverfarm {
	serverfarm := &v1beta20220301.Serverfarm{
		ObjectMeta: tc.MakeObjectMeta("appservice"),
		Spec: v1beta20220301.Serverfarms_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &v1beta20220301.SkuDescription{
				Name: to.StringPtr("P1v2"),
			},
		},
	}
	return serverfarm
}
