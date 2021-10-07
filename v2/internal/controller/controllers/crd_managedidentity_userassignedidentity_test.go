/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	managedidentity "github.com/Azure/azure-service-operator/v2/api/microsoft.managedidentity/v1alpha1api20181130"
	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"
)

func Test_ManagedIdentity_UserAssignedIdentity_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentities_Spec{
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
	tc.Patch(old, mi)

	objectKey := client.ObjectKeyFromObject(mi)

	// ensure state got updated in Azure
	tc.Eventually(func() map[string]string {
		updated := &managedidentity.UserAssignedIdentity{}
		tc.GetResource(objectKey, updated)
		return updated.Status.Tags
	}).Should(HaveKey("foo"))

	tc.DeleteResourceAndWait(mi)

	// Ensure that the resource group was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(managedidentity.UserAssignedIdentitiesSpecAPIVersion20181130))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
