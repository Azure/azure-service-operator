/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	monitor "github.com/Azure/azure-service-operator/v2/api/monitor/v1api20230403"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Monitor_Account_CRUD(t *testing.T) {
	t.Parallel()
	g := NewGomegaWithT(t)

	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// This was adapted from https://github.com/Azure/azure-rest-api-specs/blob/main/specification/monitor/resource-manager/Microsoft.Monitor/stable/2023-04-03/examples/AzureMonitorWorkspacesCreate.json
	acct := &monitor.Account{
		ObjectMeta: tc.MakeObjectMeta("acct"),
		Spec: monitor.Account_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}
	tc.CreateResourcesAndWait(acct)

	// Ensure that the status is what we expect
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	tc.DeleteResourceAndWait(acct)

	// Ensure that the resource was really deleted in Azure
	ctx := context.Background()
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(ctx, armId, string(monitor.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}
