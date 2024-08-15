/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20240302"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Compute_Disk_20240302_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a disk.
	standardSkuName := compute.DiskSku_Name_Premium_ZRS
	sizeInGb := 500
	createOption := compute.CreationData_CreateOption_Empty
	disk := &compute.Disk{
		ObjectMeta: tc.MakeObjectMeta("disk"),
		Spec: compute.Disk_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &compute.DiskSku{
				Name: &standardSkuName,
			},
			CreationData: &compute.CreationData{
				CreateOption: &createOption,
			},
			DiskSizeGB: &sizeInGb,
		},
	}
	tc.CreateResourceAndWait(disk)

	tc.Expect(disk.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(disk.Status.Sku.Name).To(BeEquivalentTo(&standardSkuName))
	tc.Expect(*disk.Status.DiskSizeGB).To(BeNumerically(">=", 500))
	tc.Expect(disk.Status.Id).ToNot(BeNil())
	armId := *disk.Status.Id

	// Perform a simple patch.
	old := disk.DeepCopy()
	networkAccessPolicy := compute.NetworkAccessPolicy_DenyAll
	disk.Spec.NetworkAccessPolicy = &networkAccessPolicy
	tc.PatchResourceAndWait(old, disk)
	tc.Expect(disk.Status.NetworkAccessPolicy).To(BeEquivalentTo(&networkAccessPolicy))

	tc.DeleteResourceAndWait(disk)

	// Ensure that the resource group was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(compute.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
