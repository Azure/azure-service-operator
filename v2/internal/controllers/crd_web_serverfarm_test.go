/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/onsi/gomega"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/api/web/v1api20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Web_ServerFarm_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Our default region (West US 2) is capacity constrained for web at the moment.
	// location := tc.AzureRegion
	location := "westus"
	serverFarm := newServerFarm(tc, rg, location)

	tc.CreateResourceAndWait(serverFarm)

	armId := *serverFarm.Status.Id
	old := serverFarm.DeepCopy()
	serverFarm.Spec.PerSiteScaling = to.BoolPtr(true)
	tc.PatchResourceAndWait(old, serverFarm)
	tc.Expect(serverFarm.Status.PerSiteScaling).To(gomega.Equal(to.BoolPtr(true)))

	tc.DeleteResourcesAndWait(serverFarm)

	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(v1api20220301.APIVersion_Value))
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())

}

func newServerFarm(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, location string) *v1api20220301.ServerFarm {
	serverFarm := &v1api20220301.ServerFarm{
		ObjectMeta: tc.MakeObjectMeta("appservice"),
		Spec: v1api20220301.Serverfarm_Spec{
			Location: &location,
			Owner:    testcommon.AsOwner(rg),
			Sku: &v1api20220301.SkuDescription{
				Name: to.StringPtr("P1v2"),
			},
		},
	}
	return serverFarm
}
