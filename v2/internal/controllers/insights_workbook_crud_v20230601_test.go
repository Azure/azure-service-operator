/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	insights "github.com/Azure/azure-service-operator/v2/api/insights/v1api20230601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Insights_Workbook_v20230601_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a workbook with the required GUID-based azureName
	workbookName := tc.Namer.GenerateName("workbook")
	uuid, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	kind := insights.Workbook_Kind_Spec_Shared
	workbook := &insights.Workbook{
		ObjectMeta: tc.MakeObjectMetaWithName(workbookName),
		Spec: insights.Workbook_Spec{
			AzureName:   uuid.String(),
			Location:    tc.AzureRegion,
			Owner:       testcommon.AsOwner(rg),
			Category:    to.Ptr("workbook"),
			Kind:        &kind,
			DisplayName: to.Ptr("Test Workbook"),
			Description: to.Ptr("A test workbook created by ASO integration tests"),
			SerializedData: to.Ptr(`{
				"version": "Notebook/1.0",
				"items": [
					{
						"type": 1,
						"content": {
							"json": "## Test Workbook\n\nThis is a test workbook created by Azure Service Operator integration tests.\n\n"
						},
						"name": "text - 0"
					},
					{
						"type": 3,
						"content": {
							"version": "KqlItem/1.0",
							"query": "// Test query\nresources\n| limit 10",
							"size": 0,
							"title": "Sample Query",
							"queryType": 1,
							"resourceType": "microsoft.resourcegraph/resources"
						},
						"name": "query - 1"
					}
				],
				"isLocked": false,
				"fallbackResourceIds": []
			}`),
			Version: to.Ptr("Notebook/1.0"),
			Tags: map[string]string{
				"purpose": "testing",
				"source":  "azure-service-operator",
			},
		},
	}

	tc.CreateResourceAndWait(workbook)

	// Perform some assertions on the workbook we just created
	tc.Expect(workbook.Status.Id).ToNot(BeNil())
	tc.Expect(workbook.Status.Kind).ToNot(BeNil())
	tc.Expect(*workbook.Status.Kind).To(Equal(insights.Workbook_Kind_STATUS_Shared))

	// Test updating the workbook
	tc.LogSectionf("Updating workbook description and tags")
	old := workbook.DeepCopy()
	workbook.Spec.Description = to.Ptr("Updated test workbook description")
	workbook.Spec.Tags["updated"] = "true"
	tc.PatchResourceAndWait(old, workbook)

	tc.Expect(workbook.Status.Id).ToNot(BeNil())
	armId := *workbook.Status.Id

	// Verify the update
	tc.Expect(workbook.Status.Description).ToNot(BeNil())
	tc.Expect(*workbook.Status.Description).To(Equal("Updated test workbook description"))
	tc.Expect(workbook.Status.Tags).To(HaveKey("updated"))
	tc.Expect(workbook.Status.Tags["updated"]).To(Equal("true"))

	tc.DeleteResourceAndWait(workbook)

	// Ensure delete by checking the resource no longer exists
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(insights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
