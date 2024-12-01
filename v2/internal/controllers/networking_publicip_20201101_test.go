/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Networking_PublicIP_20201101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	publicIPAddress := newPublicIP20201101(tc, testcommon.AsOwner(rg))

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
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newPublicIP20201101(tc *testcommon.KubePerTestContext, owner *genruntime.KnownResourceReference) *network.PublicIPAddress {
	// Public IP Address
	// TODO: Note the microsoft.networking package also defines a PublicIPAddress type, so
	// TODO: depluralization of this resource doesn't work because of the collision.
	sku := network.PublicIPAddressSku_Name_Standard
	allocationMethod := network.IPAllocationMethod_Static
	return &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddress_Spec{
			Location: tc.AzureRegion,
			Owner:    owner,
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			PublicIPAllocationMethod: &allocationMethod,
		},
	}
}
