/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	app "github.com/Azure/azure-service-operator/v2/api/app/v1api20240301"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_App_ManagedEnvironment_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	managedEnv := newManagedEnvironment(tc, rg)
	tc.CreateResourcesAndWait(managedEnv)

	tc.Expect(managedEnv.Status.Id).ToNot(BeNil())
	armId := *managedEnv.Status.Id

	old := managedEnv.DeepCopy()
	managedEnv.Spec.Tags = map[string]string{"foo": "bar"}
	tc.PatchResourceAndWait(old, managedEnv)

	tc.DeleteResourceAndWait(managedEnv)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(app.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func newManagedEnvironment(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *app.ManagedEnvironment {
	managedEnv := &app.ManagedEnvironment{
		ObjectMeta: tc.MakeObjectMeta("env"),
		Spec: app.ManagedEnvironment_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			WorkloadProfiles: []app.WorkloadProfile{
				{
					MaximumCount:        to.Ptr(2),
					MinimumCount:        to.Ptr(1),
					Name:                to.Ptr("profile1"),
					WorkloadProfileType: to.Ptr("D4"),
				},
			},
		},
	}
	return managedEnv
}
