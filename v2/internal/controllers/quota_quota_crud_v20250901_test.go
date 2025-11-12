/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	quota "github.com/Azure/azure-service-operator/v2/api/quota/v1api20250901"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Quota_Quota_v20250901_CRUD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tc := globalTestContext.ForTest(t)

	// Compute quotas must be scoped to provider-location, not ResourceGroup
	providerLocationScope := fmt.Sprintf("/subscriptions/%s/providers/Microsoft.Compute/locations/%s",
		tc.AzureSubscription, *tc.AzureRegion)

	quotaResource := &quota.Quota{
		ObjectMeta: tc.MakeObjectMeta("asotest-quota"),
		Spec: quota.Quota_Spec{
			AzureName: "cores", // Must match the quota name in properties
			Owner: &genruntime.ArbitraryOwnerReference{
				ARMID: providerLocationScope,
			},
			Properties: &quota.QuotaProperties{
				Name: &quota.ResourceName{
					Value: to.Ptr("cores"),
				},
				Limit: &quota.LimitJsonObject{
					LimitValue: &quota.LimitObject{
						Value:           to.Ptr(100),
						LimitObjectType: to.Ptr(quota.LimitType_LimitValue),
						LimitType:       to.Ptr(quota.QuotaLimitTypes_Independent), // Add missing limitType
					},
				},
			},
		},
	}

	// Create the quota resource without registering it for cleanup
	// since Azure Quota API doesn't support deletion
	tc.CreateResourceAndWaitWithoutCleanup(quotaResource)

	g.Expect(quotaResource.Status.Id).ToNot(BeNil())
	g.Expect(quotaResource.Status.Name).ToNot(BeNil())
	g.Expect(quotaResource.Status.Properties).ToNot(BeNil())
	if quotaResource.Status.Properties != nil {
		g.Expect(quotaResource.Status.Properties.Name).ToNot(BeNil())
		if quotaResource.Status.Properties.Name != nil {
			g.Expect(quotaResource.Status.Properties.Name.Value).To(Equal(to.Ptr("cores")))
		}
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

	// Remove owner reference and finalizers to prevent deletion attempts
	// since Azure Quota API doesn't support deletion
	old = quotaResource.DeepCopy()
	quotaResource.SetOwnerReferences([]metav1.OwnerReference{}) // Remove owner reference
	quotaResource.SetFinalizers([]string{})                     // Remove finalizers
	tc.Patch(old, quotaResource)
}
