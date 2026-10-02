/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/onsi/gomega"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/api/web/v20260715"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Web_ServerFarm_v20260715_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Flex Consumption is only available in a subset of regions.
	tc.AzureRegion = to.Ptr("northeurope")
	serverFarm := newServerFarmV20260715(tc, rg, *tc.AzureRegion)

	tc.CreateResourceAndWait(serverFarm)

	armId := *serverFarm.Status.Id
	old := serverFarm.DeepCopy()
	serverFarm.Spec.Tags = map[string]string{"cost-center": "12345"}
	tc.PatchResourceAndWait(old, serverFarm)
	tc.Expect(serverFarm.Status.Tags).To(gomega.HaveKeyWithValue("cost-center", "12345"))

	tc.DeleteResourcesAndWait(serverFarm)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(v20260715.APIVersion_Value),
	)
	tc.Expect(err).ToNot(gomega.HaveOccurred())
	tc.Expect(exists).To(gomega.BeFalse())
}

// newServerFarmV20260715 builds a Flex Consumption (FC1) plan used by the Function App CRUD test.
func newServerFarmV20260715(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, location string) *v20260715.ServerFarm {
	serverFarm := &v20260715.ServerFarm{
		ObjectMeta: tc.MakeObjectMeta("appservice"),
		Spec: v20260715.ServerFarm_Spec{
			Location: &location,
			Owner:    testcommon.AsOwner(rg),
			Kind:     to.Ptr("functionapp"),
			Reserved: to.Ptr(true),
			Sku: &v20260715.SkuDescription{
				Name: to.Ptr("FC1"),
				Tier: to.Ptr("FlexConsumption"),
			},
		},
	}
	return serverFarm
}
