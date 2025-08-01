/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Insights_ActionGroup_v20230101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	actionGroup := &insights.ActionGroup{
		ObjectMeta: tc.MakeObjectMeta("group"),
		Spec: insights.ActionGroup_Spec{
			Enabled:        to.Ptr(false),
			Location:       to.Ptr("global"),
			Owner:          testcommon.AsOwner(rg),
			GroupShortName: to.Ptr("aso"),
		},
	}

	tc.CreateResourceAndWait(actionGroup)

	tc.Expect(actionGroup.Status.Id).ToNot(BeNil())
	armId := *actionGroup.Status.Id

	old := actionGroup.DeepCopy()
	key := "foo"
	actionGroup.Spec.Tags = map[string]string{key: "bar"}

	tc.PatchResourceAndWait(old, actionGroup)
	tc.Expect(actionGroup.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(actionGroup)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
