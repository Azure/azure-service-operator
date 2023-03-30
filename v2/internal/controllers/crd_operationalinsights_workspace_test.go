/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"

	operationalinsights "github.com/Azure/azure-service-operator/v2/api/operationalinsights/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_OperationalInsights_Workspace_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Create a workspace
	sku := operationalinsights.WorkspaceSku_Name_Standalone
	workspace := &operationalinsights.Workspace{
		ObjectMeta: tc.MakeObjectMeta("workspace"),
		Spec: operationalinsights.Workspace_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Sku: &operationalinsights.WorkspaceSku{
				Name: &sku,
			},
		},
	}

	tc.CreateResourceAndWait(workspace)

	tc.Expect(workspace.Status.Location).To(Equal(tc.AzureRegion))
	tc.Expect(workspace.Status.Id).ToNot(BeNil())
	armId := *workspace.Status.Id

	// Perform a simple patch.
	old := workspace.DeepCopy()
	workspace.Spec.RetentionInDays = to.IntPtr(36)
	tc.PatchResourceAndWait(old, workspace)
	tc.Expect(workspace.Status.RetentionInDays).To(Equal(to.IntPtr(36)))

	tc.DeleteResourceAndWait(workspace)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(operationalinsights.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}
