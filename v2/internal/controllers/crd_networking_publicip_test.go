/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1beta20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_PublicIP_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Public IP Address
	sku := network.PublicIPAddressSku_Name_Standard
	allocationMethod := network.IPAllocationMethod_Static
	publicIPAddress := &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			PublicIPAllocationMethod: &allocationMethod,
		},
	}

	tc.CreateResourceAndWait(publicIPAddress)

	tc.Expect(publicIPAddress.Status.Id).ToNot(BeNil())
	armId := *publicIPAddress.Status.Id

	// Perform a simple patch
	old := publicIPAddress.DeepCopy()
	idleTimeoutInMinutes := 7
	publicIPAddress.Spec.IdleTimeoutInMinutes = &idleTimeoutInMinutes
	tc.PatchResourceAndWait(old, publicIPAddress)
	tc.Expect(publicIPAddress.Status.IdleTimeoutInMinutes).To(Equal(&idleTimeoutInMinutes))

	tc.DeleteResourceAndWait(publicIPAddress)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
