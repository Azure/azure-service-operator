/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_DatabaseWatcher_Watcher_v20241001preview_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Database Watcher isn't available in the default test region
	tc.AzureRegion = to.Ptr("eastus")

	rg := tc.CreateTestResourceGroupAndWait()

	watcher := &databasewatcher.Watcher{
		ObjectMeta: tc.MakeObjectMeta("watcher"),
		Spec: databasewatcher.Watcher_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Identity: &databasewatcher.ManagedServiceIdentity{
				Type: to.Ptr(databasewatcher.ManagedServiceIdentityType_SystemAssigned),
			},
			Tags: map[string]string{
				"purpose": "testing",
			},
		},
	}

	tc.CreateResourceAndWait(watcher)

	tc.Expect(watcher.Status.Id).ToNot(BeNil())
	tc.Expect(watcher.Status.Name).ToNot(BeNil())
	tc.Expect(watcher.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(watcher.Status.Tags).To(HaveKeyWithValue("purpose", "testing"))
	tc.Expect(watcher.Status.Identity).ToNot(BeNil())
	tc.Expect(watcher.Status.Identity.PrincipalId).ToNot(BeNil())
	tc.Expect(watcher.Status.Status).ToNot(BeNil())

	old := watcher.DeepCopy()
	watcher.Spec.Tags["environment"] = "test"
	tc.PatchResourceAndWait(old, watcher)
	tc.Expect(watcher.Status.Tags).To(HaveKeyWithValue("environment", "test"))

	armID := *watcher.Status.Id
	tc.DeleteResourceAndWait(watcher)

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armID, string(databasewatcher.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
