/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20241101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Compute_AvailabilitySet_v20241101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create an availability set.
	platformFaultDomainCount := 2
	platformUpdateDomainCount := 3
	availabilitySet := &compute.AvailabilitySet{
		ObjectMeta: tc.MakeObjectMeta("availabilityset"),
		Spec: compute.AvailabilitySet_Spec{
			Location:                  tc.AzureRegion,
			Owner:                     testcommon.AsOwner(rg),
			PlatformFaultDomainCount:  &platformFaultDomainCount,
			PlatformUpdateDomainCount: &platformUpdateDomainCount,
		},
	}
	tc.CreateResourceAndWait(availabilitySet)

	tc.Expect(availabilitySet.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(*availabilitySet.Status.PlatformFaultDomainCount).To(Equal(platformFaultDomainCount))
	tc.Expect(*availabilitySet.Status.PlatformUpdateDomainCount).To(Equal(platformUpdateDomainCount))
	tc.Expect(availabilitySet.Status.Id).ToNot(BeNil())
	armId := *availabilitySet.Status.Id

	// Perform a simple patch - update tags
	old := availabilitySet.DeepCopy()
	newTags := map[string]string{"env": "test"}
	availabilitySet.Spec.Tags = newTags
	tc.PatchResourceAndWait(old, availabilitySet)
	tc.Expect(availabilitySet.Status.Tags).To(Equal(newTags))

	// Now check we can delete it
	tc.DeleteResourceAndWait(availabilitySet)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(compute.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
