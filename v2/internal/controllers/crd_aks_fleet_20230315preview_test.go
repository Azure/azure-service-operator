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
	armId := *flt.Status.Id

	// patching a fleet
	old := flt.DeepCopy()
	flt.Spec.Tags = map[string]string{
		"name": "test-tag2",
	}
	tc.PatchResourceAndWait(old, flt)
	tc.Expect(flt.Spec.Tags["name"]).To(Equal("test-tag2"))

	// delete a fleet
	tc.DeleteResourceAndWait(flt)

	// Ensure that fleet was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(fleet.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
