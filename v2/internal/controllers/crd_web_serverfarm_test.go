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

func newServerFarm(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *v1beta20220301.ServerFarm {
	serverfarm := &v1beta20220301.ServerFarm{
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
