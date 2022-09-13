/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1beta20210515"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1beta20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1beta20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_CosmosDB_SQLDatabase_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because cosmosdb accounts have stricter name
	// requirements - no hyphens allowed.

	// Create a Cosmos DB account
<<<<<<< HEAD
	offerType := documentdb.DatabaseAccountOfferType_Standard
	kind := documentdb.DatabaseAccount_Spec_Kind_GlobalDocumentDB
=======
	offerType := documentdb.DatabaseAccountCreateUpdateProperties_DatabaseAccountOfferType_Standard
	kind := documentdb.DatabaseAccount_Kind_Spec_GlobalDocumentDB
>>>>>>> main
	acct := documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("sqlacct")),
		Spec: documentdb.DatabaseAccount_Spec{
			Location:                 tc.AzureRegion,
			Owner:                    testcommon.AsOwner(rg),
			Kind:                     &kind,
			DatabaseAccountOfferType: &offerType,
			Locations: []documentdb.Location{
				{
					LocationName: tc.AzureRegion,
				},
			},
		},
	}

	dbName := tc.Namer.GenerateName("sqldb")
	db := documentdb.SqlDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(dbName),
		Spec: documentdb.DatabaseAccounts_SqlDatabase_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(&acct),
			Options: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{
					MaxThroughput: to.IntPtr(4000),
				},
			},
			Resource: &documentdb.SqlDatabaseResource{
				Id: &dbName,
			},
		},
	}
	tc.T.Logf("Creating SQL account and database %q", dbName)
	tc.CreateResourcesAndWait(&acct, &db)

	tc.T.Logf("SQL account and database successfully created")
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB SQL RoleAssignment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_RoleAssignment_CRUD(tc, rg, &acct)
			},
		},
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
	lastWriterWins := documentdb.ConflictResolutionPolicy_Mode_LastWriterWins
	consistent := documentdb.IndexingPolicy_IndexingMode_Consistent
	hash := documentdb.ContainerPartitionKey_Kind_Hash
	container := documentdb.SqlDatabaseContainer{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccounts_SqlDatabases_Container_Spec{
			Location: tc.AzureRegion,
			Options: &documentdb.CreateUpdateOptions{
				Throughput: to.IntPtr(400),
			},
			Owner: testcommon.AsOwner(db),
			Resource: &documentdb.SqlContainerResource{
				Id: &name,
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
	tc.PatchResourceAndWait(old, &container)
	tc.Expect(container.Status.Resource).ToNot(BeNil())
	tc.Expect(container.Status.Resource.DefaultTtl).ToNot(BeNil())
	tc.Expect(*container.Status.Resource.DefaultTtl).To(Equal(400))

	tc.T.Logf("Cleaning up container %q", name)
}

func CosmosDB_SQL_Trigger_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	name := tc.Namer.GenerateName("trigger")
	pre := documentdb.SqlTriggerResource_TriggerType_Pre
	create := documentdb.SqlTriggerResource_TriggerOperation_Create
	trigger := documentdb.SqlDatabaseContainerTrigger{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.DatabaseAccounts_SqlDatabases_Containers_Trigger_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(container),
			Resource: &documentdb.SqlTriggerResource{
				Id:               &name,
				TriggerType:      &pre,
				TriggerOperation: &create,
				Body:             to.StringPtr(triggerBody),
			},
		},
	}

	tc.CreateResourceAndWait(&trigger)
	defer tc.DeleteResourceAndWait(&trigger)

	tc.T.Logf("Updating the trigger type on trigger %q", name)
	post := documentdb.SqlTriggerResource_TriggerType_Post
	old := trigger.DeepCopy()
	trigger.Spec.Resource.TriggerType = &post
	tc.PatchResourceAndWait(old, &trigger)
	tc.Expect(trigger.Status.Resource).ToNot(BeNil())
	tc.Expect(trigger.Status.Resource.TriggerType).ToNot(BeNil())
	tc.Expect(string(*trigger.Status.Resource.TriggerType)).To(Equal("Post"))

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
		Spec: documentdb.DatabaseAccounts_SqlDatabases_Containers_StoredProcedure_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(container),
			Resource: &documentdb.SqlStoredProcedureResource{
				Id:   &name,
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
	tc.PatchResourceAndWait(old, &storedProcedure)
	tc.Expect(storedProcedure.Status.Resource).ToNot(BeNil())
	tc.Expect(storedProcedure.Status.Resource.Body).ToNot(BeNil())
	tc.Expect(*storedProcedure.Status.Resource.Body).To(Equal(newBody))

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
		Spec: documentdb.DatabaseAccounts_SqlDatabases_Containers_UserDefinedFunction_Spec{
			AzureName: name,
			Location:  tc.AzureRegion,
			Owner:     testcommon.AsOwner(container),
			Resource: &documentdb.SqlUserDefinedFunctionResource{
				Id:   &name,
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
	tc.PatchResourceAndWait(old, &userDefinedFunction)
	tc.Expect(userDefinedFunction.Status.Resource).ToNot(BeNil())
	tc.Expect(userDefinedFunction.Status.Resource.Body).ToNot(BeNil())
	tc.Expect(*userDefinedFunction.Status.Resource.Body).To(Equal(newBody))

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
		Spec: documentdb.DatabaseAccounts_SqlDatabases_ThroughputSetting_Spec{
			Owner: testcommon.AsOwner(db),
			Resource: &documentdb.ThroughputSettingsResource{
				// We cannot change this to be a fixed throughput as we already created the database using
				// autoscale and they do not allow switching back to fixed from that.
				AutoscaleSettings: &documentdb.AutoscaleSettingsResource{
					MaxThroughput: to.IntPtr(5000),
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
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings.MaxThroughput).To(Equal(to.IntPtr(5000)))

	tc.T.Log("increase max throughput to 6000")
	old := throughputSettings.DeepCopy()
	throughputSettings.Spec.Resource.AutoscaleSettings.MaxThroughput = to.IntPtr(6000)
	tc.PatchResourceAndWait(old, &throughputSettings)
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings.MaxThroughput).To(Equal(to.IntPtr(6000)))
	tc.T.Log("throughput successfully updated in status")
}

func CosmosDB_SQL_Database_Container_ThroughputSettings_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	throughputSettings := documentdb.SqlDatabaseContainerThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.DatabaseAccounts_SqlDatabases_Containers_ThroughputSetting_Spec{
			Owner: testcommon.AsOwner(container),
			Resource: &documentdb.ThroughputSettingsResource{
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
	tc.PatchResourceAndWait(old, &throughputSettings)
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.Throughput).To(Equal(to.IntPtr(600)))
	tc.T.Log("throughput successfully updated in status")
}

func CosmosDB_SQL_RoleAssignment_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, acct *documentdb.DatabaseAccount) {
	// Create a managed identity
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourceAndWait(mi)

	// TODO: See https://github.com/Azure/azure-service-operator/issues/2435 for making referencing this easier in the YAML
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())
	principalId := mi.Status.PrincipalId

	// TODO: It's not easy to generate a GUID... we should make that easier for users
	// Now assign that managed identity to a new role
	roleAssignmentGUID, err := tc.Namer.GenerateUUID()
	tc.Expect(err).ToNot(HaveOccurred())

	// TODO: Making this is very painful. We should make this easier for users too
	roleDefinitionId := fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s/sqlRoleDefinitions/00000000-0000-0000-0000-000000000002",
		tc.AzureSubscription,
		rg.AzureName(),
		acct.AzureName())

	scope := fmt.Sprintf("/subscriptions/%s/resourceGroups/%s/providers/Microsoft.DocumentDB/databaseAccounts/%s",
		tc.AzureSubscription,
		rg.AzureName(),
		acct.AzureName())

	roleAssignment := &documentdb.SqlRoleAssignment{
		ObjectMeta: tc.MakeObjectMetaWithName(roleAssignmentGUID.String()),
		Spec: documentdb.DatabaseAccounts_SqlRoleAssignment_Spec{
			Owner:            testcommon.AsOwner(acct),
			PrincipalId:      principalId,
			RoleDefinitionId: &roleDefinitionId,
			Scope:            &scope,
		},
	}

	tc.CreateResourceAndWait(roleAssignment)

	// Ensure that the status is what we expect
	tc.Expect(roleAssignment.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(roleAssignment)
}
