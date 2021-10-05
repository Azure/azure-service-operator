/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/kr/pretty"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	documentdb "github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515"
	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"
	"github.com/Azure/go-autorest/autorest/to"
)

func Test_CosmosDB_DatabaseAccount_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because storage accounts have strict names
	namer := tc.Namer.WithSeparator("")

	// Create a Cosmos DB account
	kind := documentdb.DatabaseAccountsSpecKindMongoDB
	acct := documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("db")),
		Spec: documentdb.DatabaseAccounts_Spec{
			Location: &tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Capabilities: []documentdb.Capability{{
				Name: to.StringPtr("EnableMongo"),
			}},
			DatabaseAccountOfferType: documentdb.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferTypeStandard,
			Locations: []documentdb.Location{
				{
					LocationName: &tc.AzureRegion,
				},
			},
		},
	}

	tc.CreateResourceAndWait(&acct)

	expectedKind := documentdb.DatabaseAccountGetResultsStatusKindMongoDB
	tc.Expect(*acct.Status.Kind).To(Equal(expectedKind))

	tc.Expect(acct.Status.Id).ToNot(BeNil())
	armId := *acct.Status.Id

	tc.T.Log("updating tags on account")
	old := acct.DeepCopy()
	acct.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	tc.Patch(old, &acct)

	objectKey := client.ObjectKeyFromObject(&acct)

	tc.T.Log("waiting for new tag in status")
	tc.Eventually(func() map[string]string {
		var updated documentdb.DatabaseAccount
		tc.GetResource(objectKey, &updated)
		tc.T.Log(pretty.Sprint("current tags:", updated.Status.Tags))
		return updated.Status.Tags
	}).Should(HaveKey("scratchcard"))

	// Run sub-tests
	tc.RunParallelSubtests(testcommon.Subtest{
		Name: "CosmosDB MongoDB Database CRUD",
		Test: func(testContext testcommon.KubePerTestContext) {
			CosmosDB_MongoDB_Database_CRUD(testContext, &acct)
		},
	})

	tc.DeleteResourceAndWait(&acct)

	// Ensure that the resource group was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(tc.Ctx, armId, string(documentdb.DatabaseAccountsSpecAPIVersion20210515))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func CosmosDB_MongoDB_Database_CRUD(tc testcommon.KubePerTestContext, acct client.Object) {
	// Need to make the name match the id.
	name := tc.Namer.GenerateName("mongo")
	db := documentdb.DatabaseAccountsMongodbDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccountsMongodbDatabases_Spec{
			Location: &tc.AzureRegion,
			Options: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{
					MaxThroughput: to.IntPtr(4000),
				},
			},
			Owner: testcommon.AsOwner(acct),
			Resource: documentdb.MongoDBDatabaseResource{
				Id: name,
			},
		},
	}

	tc.T.Log("creating mongo database")
	tc.CreateResourceAndWait(&db)
	defer tc.DeleteResourceAndWait(&db)

	// TODO: for reasons I don't understand tags set on the database
	// never come back on the status. The tags are in the API
	// reference and I can't see anything wrong with the ARM-to-k8s
	// conversion for the status type. :(

	// tc.T.Log("updating tags on database")
	// old := db.DeepCopy()
	// db.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	// tc.Patch(old, &db)

	// objectKey := client.ObjectKeyFromObject(&db)

	// tc.T.Log("waiting for new tag in status")
	// tc.Eventually(func() map[string]string {
	// 	var updated documentdb.DatabaseAccountsMongodbDatabase
	// 	tc.GetResource(objectKey, &updated)
	// 	tc.T.Log(pretty.Sprint("current tags:", updated.Status.Tags))
	// 	return updated.Status.Tags
	// }).Should(HaveKey("scratchcard"))

	tc.RunParallelSubtests(testcommon.Subtest{
		Name: "CosmosDB MongoDB Collection CRUD",
		Test: func(testContext testcommon.KubePerTestContext) {
			CosmosDB_MongoDB_Collection_CRUD(testContext, &db)
		},
	})

	tc.T.Log("cleaning up database")
}

func CosmosDB_MongoDB_Collection_CRUD(tc testcommon.KubePerTestContext, db client.Object) {
	name := tc.Namer.GenerateName("collection")
	collection := documentdb.DatabaseAccountsMongodbDatabasesCollection{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccountsMongodbDatabasesCollections_Spec{
			Location: &tc.AzureRegion,
			Options: &documentdb.CreateUpdateOptions{
				Throughput: to.IntPtr(400),
			},
			Owner: testcommon.AsOwner(db),
			Resource: documentdb.MongoDBCollectionResource{
				Id: name,
				Indexes: []documentdb.MongoIndex{{
					Key: &documentdb.MongoIndexKeys{
						Keys: []string{"_id"},
					},
					Options: &documentdb.MongoIndexOptions{
						Unique: to.BoolPtr(true),
					},
				}},
			},
		},
	}

	tc.T.Log("creating mongo collection")
	tc.CreateResourceAndWait(&collection)
	defer tc.DeleteResourceAndWait(&collection)

	tc.T.Log("updating indexes on collection")
	// Add another index to the collection.
	old := collection.DeepCopy()
	collection.Spec.Resource.Indexes = append(
		collection.Spec.Resource.Indexes,
		documentdb.MongoIndex{
			Key: &documentdb.MongoIndexKeys{
				Keys: []string{"col1"},
			},
			// No need to specify false for uniqueness, that's the
			// default.
		},
	)
	tc.Patch(old, &collection)

	objectKey := client.ObjectKeyFromObject(&collection)

	tc.T.Log("waiting for new index in status")
	tc.Eventually(func() []documentdb.MongoIndex_Status {
		var updated documentdb.DatabaseAccountsMongodbDatabasesCollection
		tc.GetResource(objectKey, &updated)
		tc.T.Log(pretty.Sprint("current indexes:", updated.Status.Resource.Indexes))
		return updated.Status.Resource.Indexes
	}).Should(ContainElement(documentdb.MongoIndex_Status{
		Key: &documentdb.MongoIndexKeys_Status{
			Keys: []string{"col1"},
		},
	}))

	tc.T.Log("cleaning up collection")
}
