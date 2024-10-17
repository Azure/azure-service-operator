/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_DocumentDB_MongoDatabase_v20231115_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// The normal default region has capacity constraints, so we use australiaeast instead
	tc.AzureRegion = to.Ptr("australiaeast")

	// Create our resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Declare a Cosmos DB account for our MongoDB database
	kind := documentdb.DatabaseAccount_Kind_Spec_MongoDB
	offerType := documentdb.DatabaseAccountOfferType_Standard
	acct := documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("db")),
		Spec: documentdb.DatabaseAccount_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Kind:     &kind,
			Capabilities: []documentdb.Capability{
				{
					Name: to.Ptr("EnableMongo"),
				},
			},
			DatabaseAccountOfferType: &offerType,
			Locations: []documentdb.Location{
				{
					LocationName: tc.AzureRegion,
				},
			},
		},
	}

	// Declare a MongoDB database
	name := tc.Namer.GenerateName("mongo")
	db := documentdb.MongodbDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.MongodbDatabase_Spec{
			Location: tc.AzureRegion,
			Options: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{
					MaxThroughput: to.Ptr(4000),
				},
			},
			Owner: testcommon.AsOwner(&acct),
			Resource: &documentdb.MongoDBDatabaseResource{
				Id: &name,
			},
		},
	}

	// Create both resources in Azure
	tc.CreateResourcesAndWait(&acct, &db)

	// Perform some assertions on the resources we just created
	expectedKind := documentdb.DatabaseAccount_Kind_STATUS_MongoDB
	tc.Expect(acct.Status.Kind).ToNot(BeNil())
	tc.Expect(*acct.Status.Kind).To(Equal(expectedKind))
	tc.Expect(acct.Status.Id).ToNot(BeNil())
	tc.Expect(db.Status.Id).ToNot(BeNil())

	// Update the account to ensure that works
	tc.LogSectionf("Updating tags on account")
	old := acct.DeepCopy()
	acct.Spec.Tags = map[string]string{"scratchcard": "lanyard"}
	tc.PatchResourceAndWait(old, &acct)
	tc.Expect(acct.Status.Tags).To(HaveKey("scratchcard"))

	// Run sub-tests
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Collection CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DocumentDB_MongoDB_Collection_v20231115_CRUD(tc, &db)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Database throughput settings CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DocumentDB_MongoDB_Database_ThroughputSettings_v20231115_CRUD(tc, &db)
			},
		})

	// Delete the database and make sure it goes away
	armId := *db.Status.Id
	tc.DeleteResourceAndWait(&db)

	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(documentdb.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())

	// Delete the account and make sure it goes away
	armId = *acct.Status.Id
	tc.DeleteResourceAndWait(&acct)

	exists, _, err = tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(documentdb.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func DocumentDB_MongoDB_Collection_v20231115_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
	name := tc.Namer.GenerateName("collection")
	collection := documentdb.MongodbDatabaseCollection{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.MongodbDatabaseCollection_Spec{
			Location: tc.AzureRegion,
			Options: &documentdb.CreateUpdateOptions{
				Throughput: to.Ptr(400),
			},
			Owner: testcommon.AsOwner(db),
			Resource: &documentdb.MongoDBCollectionResource{
				Id: &name,
				Indexes: []documentdb.MongoIndex{{
					Key: &documentdb.MongoIndexKeys{
						Keys: []string{"_id"},
					},
					Options: &documentdb.MongoIndexOptions{
						Unique: to.Ptr(true),
					},
				}},
			},
		},
	}

	tc.T.Log("creating mongo collection")
	tc.CreateResourceAndWait(&collection)
	defer func() {
		tc.LogSectionf("Cleaning up collection")
		tc.DeleteResourceAndWait(&collection)
	}()

	tc.LogSectionf("Updating indexes on collection")
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
	tc.PatchResourceAndWait(old, &collection)
	tc.Expect(collection.Status.Resource).ToNot(BeNil())
	tc.Expect(collection.Status.Resource.Indexes).To(ContainElement(documentdb.MongoIndex_STATUS{
		Key: &documentdb.MongoIndexKeys_STATUS{
			Keys: []string{"col1"},
		},
	}))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Database Collection throughput settings CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DocumentDB_MongoDB_Database_Collections_ThroughputSettings_v20231515_CRUD(tc, &collection)
			},
		})
}

func DocumentDB_MongoDB_Database_ThroughputSettings_v20231115_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
	throughputSettings := documentdb.MongodbDatabaseThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.MongodbDatabaseThroughputSetting_Spec{
			Owner: testcommon.AsOwner(db),
			Resource: &documentdb.ThroughputSettingsResource{
				// We cannot change this to be a fixed throughput as we already created the database using
				// autoscale and they do not allow switching back to fixed from that.
				AutoscaleSettings: &documentdb.AutoscaleSettingsResource{
					MaxThroughput: to.Ptr(5000),
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
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings.MaxThroughput).To(Equal(to.Ptr(5000)))

	tc.LogSubsectionf("Increase max throughput to 6000")
	old := throughputSettings.DeepCopy()
	throughputSettings.Spec.Resource.AutoscaleSettings.MaxThroughput = to.Ptr(6000)
	tc.PatchResourceAndWait(old, &throughputSettings)
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings.MaxThroughput).To(Equal(to.Ptr(6000)))
}

func DocumentDB_MongoDB_Database_Collections_ThroughputSettings_v20231515_CRUD(tc *testcommon.KubePerTestContext, collection client.Object) {
	throughputSettings := documentdb.MongodbDatabaseCollectionThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.MongodbDatabaseCollectionThroughputSetting_Spec{
			Owner: testcommon.AsOwner(collection),
			Resource: &documentdb.ThroughputSettingsResource{
				Throughput: to.Ptr(500),
			},
		},
	}

	tc.LogSectionf("creating mongo database collections throughput")
	tc.CreateResourceAndWait(&throughputSettings)
	// no DELETE, this is not a real resource - to delete it you must delete its parent

	// Ensure that the status is what we expect
	tc.Expect(throughputSettings.Status.Id).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.Throughput).To(Equal(to.Ptr(500)))

	tc.T.Log("increase throughput to 600")
	old := throughputSettings.DeepCopy()
	throughputSettings.Spec.Resource.Throughput = to.Ptr(600)
	tc.PatchResourceAndWait(old, &throughputSettings)
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.Throughput).To(Equal(to.Ptr(600)))
}
