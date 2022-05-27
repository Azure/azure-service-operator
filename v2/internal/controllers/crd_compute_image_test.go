/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	compute2020 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20200930"
	compute2021 "github.com/Azure/azure-service-operator/v2/api/compute/v1beta20210701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Compute_Image_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	tc.LogSection("Create Resource Group")
	rg := tc.CreateTestResourceGroupAndWait()

	tc.LogSection("Create Snapshot")
	createOption := compute2020.CreationData_CreateOptionEmpty
	snapshot := &compute2020.Snapshot{
		ObjectMeta: tc.MakeObjectMeta("snapshot"),
		Spec: compute2020.Snapshot_Spec{
			CreationData: &compute2020.CreationData{
				CreateOption: &createOption,
			},
			DiskSizeGB: to.IntPtr(32),
			Location:   tc.AzureRegion,
			Owner:      testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(snapshot)
	tc.Expect(snapshot.Status.Id).ToNot(BeNil())
	snapshotARMId := *snapshot.Status.Id

	tc.LogSection("Create Image")
	v2 := compute2021.HyperVGenerationTypeV2
	linuxOS := compute2021.ImageOSDisk_OsTypeLinux
	linuxOSState := compute2021.ImageOSDisk_OsStateGeneralized
	image := &compute2021.Image{
		ObjectMeta: tc.MakeObjectMeta("image"),
		Spec: compute2021.Image_Spec{
			HyperVGeneration: &v2,
			Location:         tc.AzureRegion,
			Owner:            testcommon.AsOwner(rg),
			StorageProfile: &compute2021.ImageStorageProfile{
				OsDisk: &compute2021.ImageOSDisk{
					DiskSizeGB: to.IntPtr(32),
					OsType:     &linuxOS,
					OsState:    &linuxOSState,
					Snapshot: &compute2021.SubResource{
						Reference: &genruntime.ResourceReference{
							ARMID: snapshotARMId,
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(image)
	tc.Expect(image.Status.Id).ToNot(BeNil())
	tc.Expect(image.Status.StorageProfile).ToNot(BeNil())
	tc.Expect(image.Status.StorageProfile.OsDisk).ToNot(BeNil())
	tc.Expect(*image.Status.StorageProfile.OsDisk.DiskSizeGB).To(Equal(*image.Spec.StorageProfile.OsDisk.DiskSizeGB))
	tc.Expect(string(*image.Status.StorageProfile.OsDisk.OsType)).To(Equal(string(*image.Spec.StorageProfile.OsDisk.OsType)))
	tc.Expect(string(*image.Status.StorageProfile.OsDisk.OsState)).To(Equal(string(*image.Spec.StorageProfile.OsDisk.OsState)))
	imageARMId := *image.Status.Id

	tc.LogSection("Clean up")
	// Delete VM and resources.
	tc.DeleteResourcesAndWait(image, rg)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, imageARMId, string(compute2021.APIVersionValue))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
