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

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1alpha1api20210515"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_CosmosDB_SQLDatabase_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because cosmosdb accounts have stricter name
	// requirements - no hyphens allowed.
	namer := tc.Namer.WithSeparator("")

	// Create a Cosmos DB account
	kind := documentdb.DatabaseAccountsSpecKindGlobalDocumentDB
	acct := documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("sqlacct")),
		Spec: documentdb.DatabaseAccounts_Spec{
			Location:                 &tc.AzureRegion,
			Owner:                    testcommon.AsOwner(rg),
			Kind:                     &kind,
			DatabaseAccountOfferType: documentdb.DatabaseAccountCreateUpdatePropertiesDatabaseAccountOfferTypeStandard,
			Locations: []documentdb.Location{
				{
					LocationName: &tc.AzureRegion,
				},
			},
		},
	}

	dbName := tc.Namer.GenerateName("sqldb")
	db := documentdb.SqlDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(dbName),
		Spec: documentdb.DatabaseAccountsSqlDatabases_Spec{
			Location: &tc.AzureRegion,
			Owner:    testcommon.AsOwner(&acct),
			Options: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{
					MaxThroughput: to.IntPtr(4000),
				},
			},
			Resource: documentdb.SqlDatabaseResource{
				Id: dbName,
			},
		},
	}
	tc.T.Logf("Creating SQL account and database %q", dbName)
	tc.CreateResourcesAndWait(&acct, &db)
	defer tc.DeleteResourcesAndWait(&acct, &db)

	tc.T.Logf("SQL account and database successfully created")
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB SQL Container CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Container_CRUD(tc, &db)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL Database throughputsettings CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Database_ThroughputSettings_CRUD(tc, &db)
			},
		})

	// There aren't any attributes to update for databases, other than
	// throughput settings once they're available.
}

func CosmosDB_SQL_Container_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
	name := tc.Namer.GenerateName("container")
	lastWriterWins := documentdb.ConflictResolutionPolicyModeLastWriterWins
	consistent := documentdb.IndexingPolicyIndexingModeConsistent
	hash := documentdb.ContainerPartitionKeyKindHash
	container := documentdb.SqlDatabaseContainer{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccountsSqlDatabasesContainers_Spec{
			Location: &tc.AzureRegion,
			Options: &documentdb.CreateUpdateOptions{
				Throughput: to.IntPtr(400),
			},
			Owner: testcommon.AsOwner(db),
			Resource: documentdb.SqlContainerResource{
				Id: name,
				ConflictResolutionPolicy: &documentdb.ConflictResolutionPolicy{
					Mode: &lastWriterWins,
				},
				DefaultTtl: to.IntPtr(200),
				IndexingPolicy: &documentdb.IndexingPolicy{
					IndexingMode: &consistent,
					IncludedPaths: []documentdb.IncludedPath{{
						Path: to.StringPtr("/*"),
					}},
					ExcludedPaths: []documentdb.ExcludedPath{{
						Path: to.StringPtr("/myPathToNotIndex/*"),
					}},
				},
				PartitionKey: &documentdb.ContainerPartitionKey{
					Kind:  &hash,
					Paths: []string{"/myPartitionKey"},
				},
			},
		},
	}

	tc.T.Logf("Creating SQL container %q", name)
	tc.CreateResourceAndWait(&container)
	defer tc.DeleteResourceAndWait(&container)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB SQL Trigger CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Trigger_CRUD(tc, &container)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL Stored Procedure CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_StoredProcedure_CRUD(tc, &container)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL User-defined Function CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_UserDefinedFunction_CRUD(tc, &container)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL Container ThroughputSettings CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Database_Container_ThroughputSettings_CRUD(tc, &container)
			},
		})

	tc.T.Logf("Updating the default TTL on container %q", name)
	old := container.DeepCopy()
	container.Spec.Resource.DefaultTtl = to.IntPtr(400)
	tc.Patch(old, &container)

	objectKey := client.ObjectKeyFromObject(&container)

	tc.T.Log("Waiting for new TTL in status")
	tc.Eventually(func() int {
		var updated documentdb.SqlDatabaseContainer
		tc.GetResource(objectKey, &updated)
		resource := updated.Status.Resource
		if resource == nil {
			tc.T.Log("resource is nil")
			return 0
		}
		tc.T.Log(pretty.Sprint("current default TTL:", resource.DefaultTtl))
		if resource.DefaultTtl == nil {
			return 0
		}
		return *resource.DefaultTtl
	}).Should(Equal(400))

	tc.T.Logf("Cleaning up container %q", name)
}

func CosmosDB_SQL_Trigger_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	name := tc.Namer.GenerateName("trigger")
	pre := documentdb.SqlTriggerResourceTriggerTypePre
	create := documentdb.SqlTriggerResourceTriggerOperationCreate
	trigger := documentdb.SqlDatabaseContainerTrigger{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccountsSqlDatabasesContainersTriggers_Spec{
			Location: &tc.AzureRegion,
			Owner:    testcommon.AsOwner(container),
			Resource: documentdb.SqlTriggerResource{
				Id:               name,
				TriggerType:      &pre,
				TriggerOperation: &create,
				Body:             to.StringPtr(triggerBody),
			},
		},
	}

	tc.CreateResourceAndWait(&trigger)
	defer tc.DeleteResourceAndWait(&trigger)

	tc.T.Logf("Updating the trigger type on trigger %q", name)
	post := documentdb.SqlTriggerResourceTriggerTypePost
	old := trigger.DeepCopy()
	trigger.Spec.Resource.TriggerType = &post
	tc.Patch(old, &trigger)

	objectKey := client.ObjectKeyFromObject(&trigger)

	tc.T.Log("Waiting for new type in status")
	tc.Eventually(func() string {
		var updated documentdb.SqlDatabaseContainerTrigger
		tc.GetResource(objectKey, &updated)
		resource := updated.Status.Resource
		if resource == nil {
			tc.T.Log("resource is nil")
			return ""
		}
		tc.T.Log(pretty.Sprint("current trigger type:", resource.TriggerType))
		if resource.TriggerType == nil {
			return ""
		}
		return string(*resource.TriggerType)
	}).Should(Equal("Post"))

	tc.T.Logf("Cleaning up trigger %q", name)
}

const triggerBody = `
function validateToDoItemTimestamp(){
    var context=getContext();
    var request=context.getRequest();
    var itemToCreate=request.getBody();
    if(!('timestamp' in itemToCreate)) {
         var ts=new Date();
         itemToCreate['timestamp']=ts.getTime();
    }
    request.setBody(itemToCreate);
}`

func CosmosDB_SQL_StoredProcedure_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	name := tc.Namer.GenerateName("storedproc")
	storedProcedure := documentdb.SqlDatabaseContainerStoredProcedure{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccountsSqlDatabasesContainersStoredProcedures_Spec{
			Location: &tc.AzureRegion,
			Owner:    testcommon.AsOwner(container),
			Resource: documentdb.SqlStoredProcedureResource{
				Id:   name,
				Body: to.StringPtr(storedProcedureBody),
			},
		},
	}
	tc.CreateResourceAndWait(&storedProcedure)
	defer tc.DeleteResourceAndWait(&storedProcedure)

	tc.T.Logf("Updating the body on stored procedure %q", name)
	old := storedProcedure.DeepCopy()
	newBody := "your deodorant doesn't work!"
	storedProcedure.Spec.Resource.Body = &newBody
	tc.Patch(old, &storedProcedure)

	objectKey := client.ObjectKeyFromObject(&storedProcedure)

	tc.T.Log("Waiting for new body in status")
	tc.Eventually(func() string {
		var updated documentdb.SqlDatabaseContainerStoredProcedure
		tc.GetResource(objectKey, &updated)
		resource := updated.Status.Resource
		if resource == nil {
			tc.T.Log("resource is nil")
			return ""
		}
		tc.T.Log(pretty.Sprint("current stored procedure body:", resource.Body))
		if resource.Body == nil {
			return ""
		}
		return *resource.Body
	}).Should(Equal(newBody))

	tc.T.Logf("Cleaning up stored procedure %q", name)
}

const storedProcedureBody = `
function () {
    var context = getContext();
    var response = context.getResponse();
    response.setBody('Hello, World');
}`

func CosmosDB_SQL_UserDefinedFunction_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	name := tc.Namer.GenerateName("udf")
	userDefinedFunction := documentdb.SqlDatabaseContainerUserDefinedFunction{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccountsSqlDatabasesContainersUserDefinedFunctions_Spec{
			AzureName: name,
			Location:  &tc.AzureRegion,
			Owner:     testcommon.AsOwner(container),
			Resource: documentdb.SqlUserDefinedFunctionResource{
				Id:   name,
				Body: to.StringPtr(userDefinedFunctionBody),
			},
		},
	}
	tc.CreateResourceAndWait(&userDefinedFunction)
	defer tc.DeleteResourceAndWait(&userDefinedFunction)

	tc.T.Logf("Updating the body on user-defined function %q", name)
	old := userDefinedFunction.DeepCopy()
	newBody := "wonder what Jacinda would do?"
	userDefinedFunction.Spec.Resource.Body = &newBody
	tc.Patch(old, &userDefinedFunction)

	objectKey := client.ObjectKeyFromObject(&userDefinedFunction)

	tc.T.Log("Waiting for new body in status")
	tc.Eventually(func() string {
		var updated documentdb.SqlDatabaseContainerUserDefinedFunction
		tc.GetResource(objectKey, &updated)
		resource := updated.Status.Resource
		if resource == nil {
			tc.T.Log("resource is nil")
			return ""
		}
		tc.T.Log(pretty.Sprint("current function body:", resource.Body))
		if resource.Body == nil {
			return ""
		}
		return *resource.Body
	}).Should(Equal(newBody))

	tc.T.Logf("Cleaning up user-defined function %q", name)
}

const userDefinedFunctionBody = `
function tax(income) {
    if (income == undefined)
        throw 'no input';
    if (income < 1000)
        return income*0.1;
    else if(income < 10000)
        return income*0.2;
    else
        return income*0.4;
}`

func CosmosDB_SQL_Database_ThroughputSettings_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
	throughputSettings := documentdb.SqlDatabaseThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.DatabaseAccountsSqlDatabasesThroughputSettings_Spec{
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

	tc.T.Log("creating SQL database throughput")
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
		var updated documentdb.SqlDatabaseThroughputSetting
		tc.GetResource(objectKey, &updated)
		return updated.Status.Resource.AutoscaleSettings.MaxThroughput
	}).Should(Equal(6000))
	tc.T.Log("throughput successfully updated in status")
}

func CosmosDB_SQL_Database_Container_ThroughputSettings_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	throughputSettings := documentdb.SqlDatabaseContainerThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.DatabaseAccountsSqlDatabasesContainersThroughputSettings_Spec{
			Owner: testcommon.AsOwner(container),
			Resource: documentdb.ThroughputSettingsResource{
				Throughput: to.IntPtr(500),
			},
		},
	}

	tc.T.Log("creating SQL database container throughput")
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
		var updated documentdb.SqlDatabaseContainerThroughputSetting
		tc.GetResource(objectKey, &updated)
		return updated.Status.Resource.Throughput
	}).Should(Equal(to.IntPtr(600)))
	tc.T.Log("throughput successfully updated in status")
}
