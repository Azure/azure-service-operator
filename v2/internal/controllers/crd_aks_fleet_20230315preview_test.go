/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	fleet "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230315preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	. "github.com/onsi/gomega"
)

func Test_AKS_Fleet_20230315_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()
	region := to.Ptr("westus3")
	flt := &fleet.Fleet{
		ObjectMeta: tc.MakeObjectMeta("fleet"),
		Spec: fleet.Fleet_Spec{
			Location: region,
			Owner:    testcommon.AsOwner(rg),
			HubProfile: &fleet.FleetHubProfile{
				DnsPrefix: to.Ptr("aso"),
			},
			Tags: map[string]string{
				"name": "test-tag",
			},
		},
	}
	// creating a fleet
	tc.CreateResourceAndWait(flt)
	tc.Expect(flt.Status.Id).ToNot(BeNil())
	tc.Expect(flt.Status.Tags).ToNot(BeNil())
	tc.Expect(flt.Spec.Tags["name"]).To(Equal("test-tag"))
	// armId := *flt.Status.Id

	// patching a fleet
	old := flt.DeepCopy()
	flt.Spec.Tags = map[string]string{
		"name": "test-tag2",
	}
	tc.PatchResourceAndWait(old, flt)
	tc.Expect(flt.Spec.Tags["name"]).To(Equal("test-tag2"))

	// Run sub tests
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Fleet FleetMember CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				AKS_Fleet_FleetMember_20230315Preview_CRUD(tc, flt)
			},
		})
	// tc.RunParallelSubtests(
	// 	testcommon.Subtest{
	// 		Name: "Fleet UpdateRun CRUD",
	// 		Test: func(tc *testcommon.KubePerTestContext) {
	// 			AKS_Fleet_UpdateRun_20230315Preview_CRUD(tc, flt)
	// 		},
	// 	},
	// )

	// delete a fleet
	// tc.DeleteResourceAndWait(flt)

	// // Ensure that fleet was really deleted in Azure
	// exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(fleet.APIVersion_Value))
	// tc.Expect(err).ToNot(HaveOccurred())
	// tc.Expect(retryAfter).To(BeZero())
	// tc.Expect(exists).To(BeFalse())
}

func AKS_Fleet_FleetMember_20230315Preview_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet) {
	flt_member := &fleet.FleetsMember{
		ObjectMeta: tc.MakeObjectMeta("fleetmember"),
		Spec: fleet.Fleets_Member_Spec{
			Owner: testcommon.AsOwner(flt),
			// ClusterResourceReference: &genruntime.ResourceReference{
			// 	Group: "containerservice",
			// 	Kind:  "fleetmember",
			// 	Name:  "test-fleet",
			// },
			Group: to.Ptr("containerservice"),
		},
	}

	//create fleet member
	tc.CreateResourceAndWait(flt_member)
	tc.Expect(flt_member.Status.Id).ToNot(BeNil())
	tc.Expect(flt_member.Status.ClusterResourceId).ToNot(BeNil())

	// delete
	defer tc.DeleteResourceAndWait(flt_member)

}

func AKS_Fleet_UpdateRun_20230315Preview_CRUD(tc *testcommon.KubePerTestContext, flt *fleet.Fleet) {
	updateRun := &fleet.FleetsUpdateRun{
		Spec: fleet.Fleets_UpdateRun_Spec{
			ManagedClusterUpdate: &fleet.ManagedClusterUpdate{
				Upgrade: &fleet.ManagedClusterUpgradeSpec{
					Type: to.Ptr(fleet.ManagedClusterUpgradeType_Full),
				},
			},
		},
	}

	tc.CreateResourceAndWait(updateRun)

	defer tc.DeleteResourceAndWait(updateRun)

	tc.Expect(updateRun.Status.Id).ToNot(BeNil())

	// a basic assertion on a few properties
	tc.Expect(updateRun.Status.ManagedClusterUpdate.Upgrade.Type).To(Equal(to.Ptr(fleet.ManagedClusterUpgradeType_Full)))

	// Perform a simple patch
	old := updateRun.DeepCopy()
	updateRun.Spec.ManagedClusterUpdate.Upgrade.Type = to.Ptr(fleet.ManagedClusterUpgradeType_NodeImageOnly)

	tc.PatchResourceAndWait(old, updateRun)

	tc.Expect(updateRun.Status.ManagedClusterUpdate.Upgrade.Type).To(Equal(to.Ptr(fleet.ManagedClusterUpgradeType_NodeImageOnly)))
}
