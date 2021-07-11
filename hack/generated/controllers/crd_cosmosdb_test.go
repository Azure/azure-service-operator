/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	documentdb "github.com/Azure/azure-service-operator/hack/generated/_apis/microsoft.documentdb/v1alpha1api20210515"
	"github.com/Azure/azure-service-operator/hack/generated/pkg/testcommon"
)

func Test_CosmosDB_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateNewTestResourceGroupAndWait()

	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a Cosmos DB account
	kind := documentdb.DatabaseAccountsSpecKindGlobalDocumentDB
	acct := &documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("db")),
		Spec: documentdb.DatabaseAccounts_Spec{
			Location:                 &tc.AzureRegion,
			Owner:                    testcommon.AsOwner(rg.ObjectMeta),
			Kind:                     &kind,
			DatabaseAccountOfferType: documentdb.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferTypeStandard,
			Locations: []documentdb.Location{
				{
					LocationName: &tc.AzureRegion,
				},
			},
		},
	}

	tc.CreateResourceAndWait(acct)

	expectedKind := documentdb.DatabaseAccountGetResultsStatusKindGlobalDocumentDB
	tc.Expect(*acct.Status.Kind).To(Equal(expectedKind))

	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Run sub-tests
	/*
		t.Run("Blob Services CRUD", func(t *testing.T) {
			StorageAccount_BlobServices_CRUD(t, testContext, acct.ObjectMeta)
		})
	*/

	tc.DeleteResourceAndWait(acct)

	// Ensure that the resource group was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(documentdb.DatabaseAccountsSpecAPIVersion20210515))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
