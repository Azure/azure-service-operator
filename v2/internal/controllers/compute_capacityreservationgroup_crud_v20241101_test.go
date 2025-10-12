/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	compute "github.com/Azure/azure-service-operator/v2/api/compute/v1api20241101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Compute_CapacityReservationGroup_CRUD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Create a CapacityReservationGroup
	capacityReservationGroup := &compute.CapacityReservationGroup{
		ObjectMeta: tc.MakeObjectMeta("capacityreservationgroup"),
		Spec: compute.CapacityReservationGroup_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Tags: map[string]string{
				"environment": "test",
				"purpose":     "capacity-reservation",
			},
			Zones: []string{"1", "2"},
		},
	}

	tc.CreateResourceAndWait(capacityReservationGroup)

	// Ensure that the status is what we expect
	tc.Expect(capacityReservationGroup.Status.Id).ToNot(BeNil())
	tc.Expect(capacityReservationGroup.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(capacityReservationGroup.Status.Name).ToNot(BeNil())
	tc.Expect(capacityReservationGroup.Status.Type).To(Equal(to.Ptr("Microsoft.Compute/capacityReservationGroups")))
	tc.Expect(capacityReservationGroup.Status.Zones).To(Equal([]string{"1", "2"}))
	tc.Expect(capacityReservationGroup.Status.Tags).To(HaveKeyWithValue("environment", "test"))
	tc.Expect(capacityReservationGroup.Status.Tags).To(HaveKeyWithValue("purpose", "capacity-reservation"))

	armId := *capacityReservationGroup.Status.Id

	// Perform a simple patch to add another tag
	old := capacityReservationGroup.DeepCopy()
	capacityReservationGroup.Spec.Tags["updated"] = "true"
	tc.PatchResourceAndWait(old, capacityReservationGroup)

	tc.Expect(capacityReservationGroup.Status.Tags).To(HaveKeyWithValue("updated", "true"))

	tc.DeleteResourceAndWait(capacityReservationGroup)

	// Ensure that the resource was really deleted in Azure
	ctx := context.Background()
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(ctx, armId, string(compute.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
