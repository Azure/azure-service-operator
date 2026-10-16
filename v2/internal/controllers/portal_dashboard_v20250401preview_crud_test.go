/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	portal "github.com/Azure/azure-service-operator/v2/api/portal/v20250401preview"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_Portal_Dashboard_v20250401preview_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	dashboard := &portal.Dashboard{
		ObjectMeta: tc.MakeObjectMeta("dashboard"),
		Spec: portal.Dashboard_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &portal.DashboardPropertiesWithProvisioningState{
				Lenses: []portal.DashboardLens{
					{
						Order: to.Ptr(0),
						Parts: []portal.DashboardParts{
							{
								Position: &portal.DashboardPartsPosition{
									ColSpan: to.Ptr(6),
									RowSpan: to.Ptr(4),
									X:       to.Ptr(0),
									Y:       to.Ptr(0),
								},
								Metadata: &portal.DashboardPartMetadata{
									ExtensionHubsExtensionPartTypeMarkdownPart: &portal.MarkdownPartMetadata{
										Type: to.Ptr(portal.DashboardPartMetadataType_ExtensionHubsExtensionPartTypeMarkdownPart),
										Settings: &portal.MarkdownPartMetadataSettings{
											Content: &portal.MarkdownPartMetadataSettingsContent{
												Content:  to.Ptr("# Azure Service Operator"),
												Subtitle: to.Ptr("Managed by ASO"),
												Title:    to.Ptr("Dashboard"),
											},
										},
									},
								},
							},
						},
					},
				},
			},
			Tags: map[string]string{
				"purpose": "testing",
			},
		},
	}

	tc.CreateResourceAndWait(dashboard)

	tc.Expect(dashboard.Status.Id).ToNot(BeNil())
	tc.Expect(dashboard.Status.Properties).ToNot(BeNil())
	tc.Expect(dashboard.Status.Properties.Lenses).To(HaveLen(1))
	tc.Expect(dashboard.Status.Properties.Lenses[0].Parts).To(HaveLen(1))
	tc.Expect(dashboard.Status.Tags).To(HaveKeyWithValue("purpose", "testing"))

	old := dashboard.DeepCopy()
	dashboard.Spec.Tags["environment"] = "test"
	tc.PatchResourceAndWait(old, dashboard)
	tc.Expect(dashboard.Status.Tags).To(HaveKeyWithValue("environment", "test"))

	armID := *dashboard.Status.Id
	tc.DeleteResourceAndWait(dashboard)

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armID, string(portal.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
