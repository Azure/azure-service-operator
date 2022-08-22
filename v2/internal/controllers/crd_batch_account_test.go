/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1beta20210101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Batch_Account_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	account := &batch.BatchAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("batchacc")),
		Spec: batch.BatchAccounts_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(account)

	tc.Expect(account.Status.Id).ToNot(BeNil())
	armId := *account.Status.Id

	tc.DeleteResourceAndWait(account)

	// Ensure that the account was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(batch.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
