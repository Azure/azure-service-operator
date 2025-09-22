/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	aks "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20230201"
	fleet "github.com/Azure/azure-service-operator/v2/api/containerservice/v1api20250301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AKS_Fleet_20250301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("westus3") // TODO: the default test region of westus2 doesn't allow ds2_v2 at the moment

	rg := tc.CreateTestResourceGroupAndWait()
	region := tc.AzureRegion
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
	armId := *flt.Status.Id

	// Test the new UpdateStrategy resource from 2025-03-01 API
	updateStrategy := &fleet.FleetsUpdateStrategy{
		ObjectMeta: tc.MakeObjectMeta("updatestrategy"),
		Spec: fleet.FleetsUpdateStrategy_Spec{
			Owner: testcommon.AsOwner(flt),
			Strategy: &fleet.UpdateRunStrategy{
				Stages: []fleet.UpdateStage{
					{
						Name: to.Ptr("stage1"),
						Groups: []fleet.UpdateGroup{
							{
								Name: to.Ptr("group1"),
							},
						},
					},
				},
			},
		},
	}

	// Create and verify UpdateStrategy
	tc.CreateResourceAndWait(updateStrategy)
	tc.Expect(updateStrategy.Status.Id).ToNot(BeNil())
	tc.Expect(updateStrategy.Status.Strategy).ToNot(BeNil())
	tc.Expect(updateStrategy.Status.Strategy.Stages).To(HaveLen(1))
	tc.Expect(updateStrategy.Status.Strategy.Stages[0].Name).ToNot(BeNil())
	tc.Expect(*updateStrategy.Status.Strategy.Stages[0].Name).To(Equal("stage1"))

	// Test the new AutoUpgradeProfile resource from 2025-03-01 API
	autoUpgradeProfile := &fleet.FleetsAutoUpgradeProfile{
		ObjectMeta: tc.MakeObjectMeta("autoupgradeprofile"),
		Spec: fleet.FleetsAutoUpgradeProfile_Spec{
			Owner:   testcommon.AsOwner(flt),
			Channel: to.Ptr(fleet.UpgradeChannel_Stable),
		},
	}

	// Create and verify AutoUpgradeProfile
	tc.CreateResourceAndWait(autoUpgradeProfile)
	tc.Expect(autoUpgradeProfile.Status.Id).ToNot(BeNil())
	tc.Expect(autoUpgradeProfile.Status.Channel).ToNot(BeNil())
	tc.Expect(*autoUpgradeProfile.Status.Channel).To(Equal(fleet.UpgradeChannel_STATUS_Stable))

	// Clean up AutoUpgradeProfile
	tc.DeleteResourceAndWait(autoUpgradeProfile)

	// Clean up UpdateStrategy
	tc.DeleteResourceAndWait(updateStrategy)

	// Clean up Fleet
	tc.DeleteResourceAndWait(flt)

	// Ensure the fleet was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(fleet.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
