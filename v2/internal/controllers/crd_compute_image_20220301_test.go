/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	compute2020 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20200930"
	compute2022 "github.com/Azure/azure-service-operator/v2/api/compute/v1api20220301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Compute_Image_20220301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	tc.LogSection("Create Resource Group")
	rg := tc.CreateTestResourceGroupAndWait()

	tc.LogSection("Create Snapshot")
	createOption := compute2020.CreationData_CreateOption_Empty
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
	v2 := compute2022.HyperVGenerationType_V2
	linuxOS := compute2022.ImageOSDisk_OsType_Linux
	linuxOSState := compute2022.ImageOSDisk_OsState_Generalized
	image := &compute2022.Image{
		ObjectMeta: tc.MakeObjectMeta("image"),
		Spec: compute2022.Image_Spec{
			HyperVGeneration: &v2,
			Location:         tc.AzureRegion,
			Owner:            testcommon.AsOwner(rg),
			StorageProfile: &compute2022.ImageStorageProfile{
				OsDisk: &compute2022.ImageOSDisk{
					DiskSizeGB: to.IntPtr(32),
					OsType:     &linuxOS,
					OsState:    &linuxOSState,
					Snapshot: &compute2022.SubResource{
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
	// Delete image.
	tc.DeleteResourceAndWait(image)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, imageARMId, string(compute2022.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
