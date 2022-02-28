/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	batch "github.com/Azure/azure-service-operator/v2/api/batch/v1alpha1api20210101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_Batch_Account_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because batch accounts have strict names
	namer := tc.Namer.WithSeparator("")

	account := &batch.BatchAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("batchacc")),
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
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(batch.BatchAccountsSpecAPIVersion20210101))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
