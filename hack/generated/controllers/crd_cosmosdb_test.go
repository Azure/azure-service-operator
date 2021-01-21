/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"

	documentdb "github.com/Azure/k8s-infra/hack/generated/_apis/microsoft.documentdb/v20150408"
	"github.com/Azure/k8s-infra/hack/generated/pkg/testcommon"
)

func Test_CosmosDB_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	testContext, err := testContext.ForTest(t)
	g.Expect(err).ToNot(HaveOccurred())

	rg, err := testContext.CreateNewTestResourceGroup(testcommon.WaitForCreation)
	g.Expect(err).ToNot(HaveOccurred())

	// Custom namer because storage accounts have strict names
	namer := testContext.Namer.WithSeparator("")

	// Create a Cosmos DB account
	kind := documentdb.DatabaseAccountsSpecKindGlobalDocumentDB
	acct := &documentdb.DatabaseAccount{
		ObjectMeta: testContext.MakeObjectMetaWithName(namer.GenerateName("db")),
		Spec: documentdb.DatabaseAccounts_Spec{
			ApiVersion: "2015-04-08", // TODO [apiversion]: This should be removed from the storage type eventually
			Location:   &testContext.AzureRegion,
			Owner:      testcommon.AsOwner(rg.ObjectMeta),
			Kind:       &kind,
			Properties: documentdb.DatabaseAccountCreateUpdateProperties{
				DatabaseAccountOfferType: documentdb.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferTypeStandard,
				Locations: []documentdb.Location{
					{
						LocationName: &testContext.AzureRegion,
					},
				},
			},
		},
	}
	err = testContext.KubeClient.Create(ctx, acct)
	g.Expect(err).ToNot(HaveOccurred())

	// It should be created in Kubernetes
	g.Eventually(acct).Should(testContext.Match.BeProvisioned(ctx))

	expectedKind := documentdb.DatabaseAccountStatusKindGlobalDocumentDB
	g.Expect(*acct.Status.Kind).To(Equal(expectedKind))

	g.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	// Run sub-tests
	/*
		t.Run("Blob Services CRUD", func(t *testing.T) {
			StorageAccount_BlobServices_CRUD(t, testContext, acct.ObjectMeta)
		})
	*/

	// Delete
	err = testContext.KubeClient.Delete(ctx, acct)
	g.Expect(err).ToNot(HaveOccurred())
	g.Eventually(acct).Should(testContext.Match.BeDeleted(ctx))

	// Ensure that the resource group was really deleted in Azure
	exists, err := testContext.AzureClient.HeadResource(ctx, armId, "2015-04-08")
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(exists).To(BeFalse())
}
