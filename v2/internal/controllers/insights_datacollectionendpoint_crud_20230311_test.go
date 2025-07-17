/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230311"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Insights_DataCollectionEndpoint_v20230311_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create the data collection endpoint
	endpoint := &insights.DataCollectionEndpoint{
		ObjectMeta: tc.MakeObjectMeta("dce"),
		Spec: insights.DataCollectionEndpoint_Spec{
			Description: to.Ptr("Data collection endpoint for testing"),
			Kind:        to.Ptr(insights.DataCollectionEndpoint_Kind_Spec_Linux),
			Location:    tc.AzureRegion,
			Owner:       testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(endpoint)

	// Verify status fields are populated correctly
	tc.Expect(endpoint.Status.Id).ToNot(BeNil())
	tc.Expect(endpoint.Status.Name).ToNot(BeNil())
	tc.Expect(endpoint.Status.Type).ToNot(BeNil())
	tc.Expect(endpoint.Status.Location).ToNot(BeNil())

	armId := *endpoint.Status.Id

	// Test update
	old := endpoint.DeepCopy()
	key := "environment"
	endpoint.Spec.Tags = map[string]string{key: "test"}

	tc.PatchResourceAndWait(old, endpoint)
	tc.Expect(endpoint.Status.Tags).To(HaveKey(key))

	tc.DeleteResourceAndWait(endpoint)

	// Ensure delete
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
