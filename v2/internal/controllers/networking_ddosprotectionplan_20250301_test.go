/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v20250301"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Networking_DdosProtectionPlan_20250301_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	ddosPlan := &network.DdosProtectionPlan{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("ddos")),
		Spec: network.DdosProtectionPlan_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(ddosPlan)

	tc.Expect(ddosPlan.Status.Id).ToNot(BeNil())
	armId := *ddosPlan.Status.Id

	tc.Expect(ddosPlan.Status.ProvisioningState).ToNot(BeNil())

	// Perform a simple patch by adding tags
	old := ddosPlan.DeepCopy()
	key := "foo"
	ddosPlan.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, ddosPlan)
	tc.Expect(ddosPlan.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(ddosPlan)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(network.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
