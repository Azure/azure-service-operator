/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/kr/pretty"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	documentdb "github.com/Azure/azure-service-operator/v2/api/microsoft.documentdb/v1alpha1api20210515"
	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"
)

func Test_CosmosDB_MongoDatabase_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because cosmos DB accounts have strict names
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

	// Create a mongo database
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
			Owner: testcommon.AsOwner(&acct),
			Resource: documentdb.MongoDBDatabaseResource{
				Id: name,
			},
		},
	}

	tc.CreateResourcesAndWait(&acct, &db)
	defer tc.DeleteResourcesAndWait(&acct, &db)

	// Perform some assertions on the resources we just created
	expectedKind := documentdb.DatabaseAccountGetResultsStatusKindMongoDB
	tc.Expect(*acct.Status.Kind).To(Equal(expectedKind))
	tc.Expect(acct.Status.Id).ToNot(BeNil())

	tc.Expect(db.Status.Id).ToNot(BeNil())

	// Update the account to ensure that works
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
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Collection CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				CosmosDB_MongoDB_Collection_CRUD(testContext, &db)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Database throughput settings CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				CosmosDB_MongoDB_Database_ThroughputSettings_CRUD(testContext, &db)
			},
		})
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

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Database Collection throughput settings CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				CosmosDB_MongoDB_Database_Collections_ThroughputSettings_CRUD(testContext, &collection)
			},
		})

	tc.T.Log("cleaning up collection")
}

func CosmosDB_MongoDB_Database_ThroughputSettings_CRUD(tc testcommon.KubePerTestContext, db client.Object) {
	throughputSettings := documentdb.MongodbDatabaseThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.DatabaseAccountsMongodbDatabasesThroughputSettings_Spec{
			Owner: testcommon.AsOwner(db),
			Resource: documentdb.ThroughputSettingsResource{
				// We cannot change this to be a fixed throughput as we already created the database using
				// autoscale and they do not allow switching back to fixed from that.
				AutoscaleSettings: &documentdb.AutoscaleSettingsResource{
					MaxThroughput: 5000,
				},
			},
		},
	}

	tc.T.Log("creating mongo database throughput")
	tc.CreateResourceAndWait(&throughputSettings)
	// no DELETE, this is not a real resource - to delete it you must delete its parent

	// Ensure that the status is what we expect
	tc.Expect(throughputSettings.Status.Id).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings.MaxThroughput).To(Equal(5000))

	tc.T.Log("increase max throughput to 6000")
	old := throughputSettings.DeepCopy()
	throughputSettings.Spec.Resource.AutoscaleSettings.MaxThroughput = 6000
	tc.Patch(old, &throughputSettings)

	objectKey := client.ObjectKeyFromObject(&throughputSettings)

	tc.T.Log("waiting for new throughput in status")
	tc.Eventually(func() int {
		var updated documentdb.MongodbDatabaseThroughputSetting
		tc.GetResource(objectKey, &updated)
		return updated.Status.Resource.AutoscaleSettings.MaxThroughput
	}).Should(Equal(6000))
	tc.T.Log("throughput successfully updated in status")
}

func CosmosDB_MongoDB_Database_Collections_ThroughputSettings_CRUD(tc testcommon.KubePerTestContext, collection client.Object) {
	throughputSettings := documentdb.MongodbDatabaseCollectionThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.DatabaseAccountsMongodbDatabasesCollectionsThroughputSettings_Spec{
			Owner: testcommon.AsOwner(collection),
			Resource: documentdb.ThroughputSettingsResource{
				Throughput: to.IntPtr(500),
			},
		},
	}

	tc.T.Log("creating mongo database collections throughput")
	tc.CreateResourceAndWait(&throughputSettings)
	// no DELETE, this is not a real resource - to delete it you must delete its parent

	// Ensure that the status is what we expect
	tc.Expect(throughputSettings.Status.Id).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.Throughput).To(Equal(to.IntPtr(500)))

	tc.T.Log("increase throughput to 600")
	old := throughputSettings.DeepCopy()
	throughputSettings.Spec.Resource.Throughput = to.IntPtr(600)
	tc.Patch(old, &throughputSettings)

	objectKey := client.ObjectKeyFromObject(&throughputSettings)

	tc.T.Log("waiting for new throughput in status")
	tc.Eventually(func() *int {
		var updated documentdb.MongodbDatabaseCollectionThroughputSetting
		tc.GetResource(objectKey, &updated)
		return updated.Status.Resource.Throughput
	}).Should(Equal(to.IntPtr(600)))
	tc.T.Log("throughput successfully updated in status")
}
