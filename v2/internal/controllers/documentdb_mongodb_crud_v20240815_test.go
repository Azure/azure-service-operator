/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/onsi/gomega"

	v1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20240815"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DocumentDB_MongoDatabase_v20240815_CRUD(t *testing.T) {
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
				{
					Name: to.Ptr("EnableMongoRoleBasedAccessControl"),
				},
			},
			DatabaseAccountOfferType: &offerType,
			Locations: []documentdb.Location{
				{
					LocationName: tc.AzureRegion,
				},
			},
			ApiProperties: &documentdb.ApiProperties{
				ServerVersion: to.Ptr(documentdb.ApiProperties_ServerVersion_70),
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
				DocumentDB_MongoDB_Collection_v20240815_CRUD(tc, &db)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Database throughput settings CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DocumentDB_MongoDB_Database_ThroughputSettings_v20240815_CRUD(tc, &db)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB MongoDB User Definition CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DocumentDB_MongoDB_MongodbUserDefintion_v20240815_CRUD(tc, name, &acct)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB MongoDB Role Definition CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				DocumentDB_MongoDB_MongodbRoleDefinition_v20240815_CRUD(tc, name, &acct)
			},
		},
	)

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

func DocumentDB_MongoDB_Collection_v20240815_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
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
				DocumentDB_MongoDB_Database_Collections_ThroughputSettings_v20240815_CRUD(tc, &collection)
			},
		})
}

func DocumentDB_MongoDB_Database_ThroughputSettings_v20240815_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
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

func DocumentDB_MongoDB_Database_Collections_ThroughputSettings_v20240815_CRUD(tc *testcommon.KubePerTestContext, collection client.Object) {
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

func DocumentDB_MongoDB_MongodbUserDefintion_v20240815_CRUD(tc *testcommon.KubePerTestContext, dbName string, acct client.Object) {
	name := tc.Namer.GenerateName("user")
	userName := tc.Namer.GenerateName("user")
	password := tc.Namer.GeneratePassword()

	secret := v1.Secret{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		StringData: map[string]string{
			"password": password,
		},
	}

	err := tc.CheckIfResourceExists(&secret)
	if err != nil {
		tc.CreateResource(&secret)
	}

	user := documentdb.MongodbUserDefinition{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.MongodbUserDefinition_Spec{
			Owner:        testcommon.AsOwner(acct),
			AzureName:    fmt.Sprintf("%s.%s", dbName, userName),
			UserName:     &userName,
			DatabaseName: to.Ptr(dbName),
			Password: &genruntime.SecretReference{
				Name: name,
				Key:  "password",
			},
			Roles: []documentdb.Role{
				{
					Db:   to.Ptr(dbName),
					Role: to.Ptr("read"),
				},
			},
			Mechanisms: to.Ptr("SCRAM-SHA-256"),
		},
	}

	tc.T.Log("creating mongo user")
	tc.CreateResourceAndWait(&user)
	defer func() {
		tc.LogSectionf("Cleaning up user")
		tc.DeleteResourceAndWait(&user)

		tc.LogSectionf("Cleaning up secret")
		tc.DeleteResourceAndWait(&secret)
	}()
	tc.Expect(user.Status).ToNot(BeNil())
	tc.Expect(user.Status.Id).ToNot(BeNil())
	tc.Expect(user.Status.Roles).To(ContainElement(documentdb.Role_STATUS{
		Db:   to.Ptr(dbName),
		Role: to.Ptr("read"),
	}))

	tc.LogSectionf("Updating roles on user")
	// Add another role to the user.
	old := user.DeepCopy()
	user.Spec.Roles = append(
		user.Spec.Roles,
		documentdb.Role{
			Db:   to.Ptr(dbName),
			Role: to.Ptr("dbAdmin"),
		},
	)
	tc.PatchResourceAndWait(old, &user)
	tc.Expect(user.Status.Roles).To(ContainElement(documentdb.Role_STATUS{
		Db:   to.Ptr(dbName),
		Role: to.Ptr("dbAdmin"),
	}))
}

func DocumentDB_MongoDB_MongodbRoleDefinition_v20240815_CRUD(tc *testcommon.KubePerTestContext, dbName string, acct client.Object) {
	roleName := "dboperator"
	name := fmt.Sprintf("%s.%s", dbName, roleName)

	role := documentdb.MongodbRoleDefinition{
		ObjectMeta: tc.MakeObjectMetaWithName(strings.ReplaceAll(name, ".", "-")),
		Spec: documentdb.MongodbRoleDefinition_Spec{
			Owner:        testcommon.AsOwner(acct),
			AzureName:    name,
			RoleName:     &roleName,
			DatabaseName: to.Ptr(dbName),
			Privileges: []documentdb.Privilege{
				{
					Actions: []string{"dbStats"},
					Resource: &documentdb.Privilege_Resource{
						Db: to.Ptr(dbName),
					},
				},
			},
			Type: to.Ptr(documentdb.MongoRoleDefinitionResource_Type_CustomRole),
		},
	}

	tc.T.Log("creating mongo role")
	tc.CreateResourceAndWait(&role)
	defer func() {
		tc.LogSectionf("Cleaning up role")
		tc.DeleteResourceAndWait(&role)
	}()

	tc.Expect(role.Status).ToNot(BeNil())
	tc.Expect(role.Status.Privileges).To(ContainElement(documentdb.Privilege_STATUS{
		Actions: []string{"dbStats"},
		Resource: &documentdb.Privilege_Resource_STATUS{
			Db: to.Ptr(dbName),
		},
	}))

	tc.LogSectionf("Updating roles on role")
	// Add another privilege.
	old := role.DeepCopy()
	role.Spec.Privileges = append(
		role.Spec.Privileges,
		documentdb.Privilege{
			Actions: []string{"dropDatabase"},
			Resource: &documentdb.Privilege_Resource{
				Db: to.Ptr(dbName),
			},
		},
	)
	tc.PatchResourceAndWait(old, &role)
	tc.Expect(role.Status).ToNot(BeNil())
	tc.Expect(role.Status.Privileges).To(ContainElements([]documentdb.Privilege_STATUS{
		{
			Actions: []string{"dropDatabase"},
			Resource: &documentdb.Privilege_Resource_STATUS{
				Db: to.Ptr(dbName),
			},
		},
		{
			Actions: []string{"dbStats"},
			Resource: &documentdb.Privilege_Resource_STATUS{
				Db: to.Ptr(dbName),
			},
		},
	}))
}
