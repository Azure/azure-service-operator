/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_ManagedIdentity_UserAssignedIdentity_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(mi)

	tc.Expect(mi.Status.TenantId).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())
	tc.Expect(mi.Status.Id).ToNot(BeNil())
	armId := *mi.Status.Id

	// Perform a simple patch
	old := mi.DeepCopy()
	mi.Spec.Tags = map[string]string{
		"foo": "bar",
	}
	tc.PatchResourceAndWait(old, mi)
	tc.Expect(mi.Status.Tags).To(HaveKey("foo"))

	tc.DeleteResourceAndWait(mi)

	// Ensure that the resource group was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(managedidentity.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
