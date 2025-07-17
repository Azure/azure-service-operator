/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20201001"
	insights2023 "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_Insights_ActivityLogAlert_v20201001_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a storage account to use as scope for the activity log alert
	sa := newStorageAccount(tc, rg)

	// Create an action group to use in the activity log alert
	actionGroup := &insights2023.ActionGroup{
		ObjectMeta: tc.MakeObjectMeta("actiongroup"),
		Spec: insights2023.ActionGroup_Spec{
			Enabled:        to.Ptr(false),
			Location:       to.Ptr("global"),
			Owner:          testcommon.AsOwner(rg),
			GroupShortName: to.Ptr("aso"),
		},
	}

	// Create the activity log alert
	alert := &insights.ActivityLogAlert{
		ObjectMeta: tc.MakeObjectMeta("alert"),
		Spec: insights.ActivityLogAlert_Spec{
			Actions: &insights.ActionList{
				ActionGroups: []insights.ActionGroupReference{
					{
						ActionGroupReference: tc.MakeReferenceFromResource(actionGroup),
					},
				},
			},
			Condition: &insights.AlertRuleAllOfCondition{
				AllOf: []insights.AlertRuleAnyOfOrLeafCondition{
					{
						Field:  to.Ptr("category"),
						Equals: to.Ptr("Administrative"),
					},
					{
						Field:  to.Ptr("operationName"),
						Equals: to.Ptr("Microsoft.Storage/storageAccounts/write"),
					},
				},
			},
			Description: to.Ptr("Activity log alert for testing"),
			Enabled:     to.Ptr(false),
			Location:    to.Ptr("global"),
			Owner:       testcommon.AsOwner(rg),
			ScopesReferences: []genruntime.ResourceReference{
				*tc.MakeReferenceFromResource(sa),
			},
		},
	}

	tc.CreateResourcesAndWait(sa, actionGroup, alert)

	// Verify status fields are populated correctly
	tc.Expect(alert.Status.Id).ToNot(BeNil())
	tc.Expect(alert.Status.Name).ToNot(BeNil())
	tc.Expect(alert.Status.Type).ToNot(BeNil())

	armId := *alert.Status.Id

	// Test update
	old := alert.DeepCopy()
	key := "environment"
	alert.Spec.Tags = map[string]string{key: "test"}

	tc.PatchResourceAndWait(old, alert)
	tc.Expect(alert.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(alert)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
