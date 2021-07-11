/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	network "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.network/v1alpha1api20201101"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_PublicIP_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateNewTestResourceGroupAndWait()

	// Public IP Address
	// TODO: Note the microsoft.networking package also defines a PublicIPAddress type, so
	// TODO: depluralization of this resource doesn't work because of the collision.
	sku := network.PublicIPAddressSkuNameStandard
	publicIPAddress := &network.PublicIPAddress{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("publicip")),
		Spec: network.PublicIPAddresses_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Sku: &network.PublicIPAddressSku{
				Name: &sku,
			},
			PublicIPAllocationMethod: network.PublicIPAddressPropertiesFormatPublicIPAllocationMethodStatic,
		},
	}

	tc.CreateResourceAndWait(publicIPAddress)

	tc.Expect(publicIPAddress.Status.Id).ToNot(BeNil())
	armId := *publicIPAddress.Status.Id

	// Perform a simple patch
	patcher := tc.NewResourcePatcher(publicIPAddress)

	idleTimeoutInMinutes := 7
	publicIPAddress.Spec.IdleTimeoutInMinutes = &idleTimeoutInMinutes
	patcher.Patch(publicIPAddress)

	objectKey, err := client.ObjectKeyFromObject(publicIPAddress)
	tc.Expect(err).ToNot(HaveOccurred())

	// ensure state got updated in Azure
	tc.Eventually(func() *int {
		updatedIP := &network.PublicIPAddress{}
		tc.GetResource(objectKey, updatedIP)
		return updatedIP.Status.IdleTimeoutInMinutes
	}).Should(Equal(&idleTimeoutInMinutes))

	tc.DeleteResourceAndWait(publicIPAddress)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(network.PublicIPAddressesSpecAPIVersion20201101))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
