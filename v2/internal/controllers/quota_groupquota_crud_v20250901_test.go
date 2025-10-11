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

func Test_Quota_GroupQuota_CRUD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tc := globalTestContext.ForTest(t)

	// Create a Resource Group to own the GroupQuota (since it's an extension resource)
	rg := tc.CreateTestResourceGroupAndWait()

	groupQuota := &quota.GroupQuota{
		ObjectMeta: tc.MakeObjectMeta("groupquota"),
		Spec: quota.GroupQuota_Spec{
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

	tc.CreateResourceAndWait(groupQuota)
	defer tc.DeleteResourceAndWait(groupQuota)

	g.Expect(groupQuota.Status.Id).ToNot(BeNil())
	g.Expect(groupQuota.Status.Name).ToNot(BeNil())
	g.Expect(groupQuota.Status.Properties).ToNot(BeNil())
	if groupQuota.Status.Properties != nil {
		g.Expect(groupQuota.Status.Properties.Name).ToNot(BeNil())
		g.Expect(groupQuota.Status.Properties.ResourceType).To(Equal(to.Ptr("cores")))
	}

	// Update the GroupQuota limit
	old := groupQuota.DeepCopy()
	groupQuota.Spec.Properties.Limit.LimitValue.Value = to.Ptr(200)
	tc.PatchResourceAndWait(old, groupQuota)

	if groupQuota.Status.Properties != nil && 
		groupQuota.Status.Properties.Limit != nil && 
		groupQuota.Status.Properties.Limit.LimitValue != nil {
		g.Expect(groupQuota.Status.Properties.Limit.LimitValue.Value).To(Equal(to.Ptr(200)))
	}
}
