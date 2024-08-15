/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20200930"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Compute_Snapshot_20200930_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.LogSectionf("Create Resource Group")
	rg := tc.CreateTestResourceGroupAndWait()

	tc.LogSectionf("Create Snapshot")
	createOption := compute.CreationData_CreateOption_Empty
	snapshot := &compute.Snapshot{
		ObjectMeta: tc.MakeObjectMeta("snapshot"),
		Spec: compute.Snapshot_Spec{
			CreationData: &compute.CreationData{
				CreateOption: &createOption,
			},
			DiskSizeGB: to.Ptr(32),
			Location:   tc.AzureRegion,
			Owner:      testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(snapshot)
	tc.Expect(snapshot.Status.Id).ToNot(BeNil())
	tc.Expect(snapshot.Status.DiskSizeGB).To(Equal(snapshot.Spec.DiskSizeGB))
	tc.Expect(snapshot.Status.Location).To(Equal(snapshot.Spec.Location))
	armId := *snapshot.Status.Id

	// Perform a simple patch to resize the disk
	tc.LogSectionf("Patch Snapshot")
	old := snapshot.DeepCopy()
	snapshot.Spec.DiskSizeGB = to.Ptr(64)
	tc.PatchResourceAndWait(old, snapshot)

	tc.Expect(snapshot.Status.DiskSizeGB).ToNot(BeNil())
	tc.Expect(*snapshot.Status.DiskSizeGB).To(Equal(64))

	// Delete VM and resources.
	tc.LogSectionf("Clean up")
	tc.DeleteResourceAndWait(snapshot)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(compute.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
