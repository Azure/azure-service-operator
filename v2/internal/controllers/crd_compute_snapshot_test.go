/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Compute_Snapshot_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	tc.LogSection("Create Resource Group")
	rg := tc.CreateTestResourceGroupAndWait()

	tc.LogSection("Create Snapshot")
	createOption := compute.CreationData_CreateOption_Empty
	snapshot := &compute.Snapshot{
		ObjectMeta: tc.MakeObjectMeta("snapshot"),
		Spec: compute.Snapshots_Spec{
			CreationData: &compute.CreationData{
				CreateOption: &createOption,
			},
			DiskSizeGB: to.IntPtr(32),
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
	tc.LogSection("Patch Snapshot")
	old := snapshot.DeepCopy()
	snapshot.Spec.DiskSizeGB = to.IntPtr(64)
	tc.PatchResourceAndWait(old, snapshot)

	tc.Expect(snapshot.Status.DiskSizeGB).ToNot(BeNil())
	tc.Expect(*snapshot.Status.DiskSizeGB).To(Equal(64))

	// Delete VM and resources.
	tc.LogSection("Clean up")
	tc.DeleteResourceAndWait(snapshot)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(compute.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
