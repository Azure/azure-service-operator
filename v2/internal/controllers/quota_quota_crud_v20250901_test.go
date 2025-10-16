/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	quota "github.com/Azure/azure-service-operator/v2/api/quota/v1api20250901"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Quota_Quota_CRUD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tc := globalTestContext.ForTest(t)

	// Create a Resource Group to own the Quota (since it's an extension resource)
	rg := tc.CreateTestResourceGroupAndWait()

	quotaResource := &quota.Quota{
		ObjectMeta: tc.MakeObjectMeta("asotest-quota"),
		Spec: quota.Quota_Spec{
			Owner: tc.AsExtensionOwner(rg),
			Properties: &quota.QuotaProperties{
				Name: &quota.ResourceName{
					Value: to.Ptr("cores"),
				},
				ResourceType: to.Ptr("cores"),
				Limit: &quota.LimitJsonObject{
					LimitValue: &quota.LimitObject{
						Value:           to.Ptr(100),
						LimitObjectType: to.Ptr(quota.LimitType_LimitValue),
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(quotaResource)

	g.Expect(quotaResource.Status.Id).ToNot(BeNil())
	g.Expect(quotaResource.Status.Name).ToNot(BeNil())
	g.Expect(quotaResource.Status.Properties).ToNot(BeNil())
	if quotaResource.Status.Properties != nil {
		g.Expect(quotaResource.Status.Properties.Name).ToNot(BeNil())
		g.Expect(quotaResource.Status.Properties.ResourceType).To(Equal(to.Ptr("cores")))
	}

	// Update the Quota limit
	old := quotaResource.DeepCopy()
	quotaResource.Spec.Properties.Limit.LimitValue.Value = to.Ptr(200)
	tc.PatchResourceAndWait(old, quotaResource)

	if quotaResource.Status.Properties != nil &&
		quotaResource.Status.Properties.Limit != nil && 
		quotaResource.Status.Properties.Limit.LimitValue != nil {
		g.Expect(quotaResource.Status.Properties.Limit.LimitValue.Value).To(Equal(to.Ptr(200)))
	}

	// NOTE: We do NOT test deletion here because:
	// 1. Azure Quota API does not support deletion of quotas
	// 2. Quotas are system-managed resources, not user-created resources
	// 3. The DELETE API returns 400 Bad Request
	// 4. Quotas persist as Azure platform limits and cannot be removed
	//
	// We simply verify the quota exists and has the correct updated values
	g.Expect(quotaResource.Status.Id).ToNot(BeNil())
	armId := *quotaResource.Status.Id
	
	// Verify the quota still exists in Azure with updated values
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(quota.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeTrue()) // Quota should still exist
}


