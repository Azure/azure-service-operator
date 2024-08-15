/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20240302"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20220701"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Compute_DiskAccess_20240302_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create DiskAccess
	diskAccess := &compute.DiskAccess{
		ObjectMeta: tc.MakeObjectMeta("diskaccess"),
		Spec: compute.DiskAccess_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	// Create VirtualNetwork
	vnet := newVMVirtualNetwork(tc, testcommon.AsOwner(rg))
	subnet := newVMSubnet(tc, testcommon.AsOwner(vnet))

	// Create PrivateEndpointConnection... wait do we actually need this, what is this?
	//diskPEC := &compute.DiskAccessesPrivateEndpointConnection{
	//	ObjectMeta: tc.MakeObjectMeta("pec"),
	//	Spec: compute.DiskAccesses_PrivateEndpointConnection_Spec{
	//		Owner: testcommon.AsOwner(diskAccess),
	//		// TODO: Do we need to remvoe ProvState here?
	//	},
	//}

	// Create PrivateEndpoint
	privateEndpoint := &network.PrivateEndpoint{
		ObjectMeta: tc.MakeObjectMeta("privateendpoint"),
		Spec: network.PrivateEndpoint_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			PrivateLinkServiceConnections: []network.PrivateLinkServiceConnection{
				{
					Name:                        to.Ptr("diskEndpoint"),
					PrivateLinkServiceReference: tc.MakeReferenceFromResource(diskAccess),
					GroupIds:                    []string{"disks"},
				},
			},
			Subnet: &network.Subnet_PrivateEndpoint_SubResourceEmbedded{
				Reference: tc.MakeReferenceFromResource(subnet),
			},
		},
	}

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
			DiskSizeGB:          &sizeInGb,
			NetworkAccessPolicy: to.Ptr(compute.NetworkAccessPolicy_AllowPrivate),
			DiskAccessReference: tc.MakeReferenceFromResource(diskAccess),
		},
	}

	tc.CreateResourcesAndWait(diskAccess, vnet, subnet, privateEndpoint, disk)

	tc.Expect(diskAccess.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(diskAccess.Status.Id).ToNot(BeNil())
	armId := *diskAccess.Status.Id

	// Perform a simple patch.
	old := diskAccess.DeepCopy()
	diskAccess.Spec.Tags = map[string]string{
		"foo": "bar",
	}
	tc.PatchResourceAndWait(old, diskAccess)
	tc.Expect(diskAccess.Status.Tags).To(BeEquivalentTo(diskAccess.Spec.Tags))

	// First delete the disk, as that will block individual disk access deletion
	tc.DeleteResourceAndWait(disk)

	tc.DeleteResourceAndWait(diskAccess)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(compute.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
