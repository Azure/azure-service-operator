/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"sigs.k8s.io/controller-runtime/pkg/client"

	documentdb "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DocumentDB_SQLDatabase_v20231115_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Capacity constraints
	tc.AzureRegion = to.Ptr("australiaeast")

	// Create our resource group
	rg := tc.CreateTestResourceGroupAndWait()

	// Custom namer because cosmosdb accounts have stricter name
	// requirements - no hyphens allowed.

	// Declare a Cosmos DB account
	offerType := documentdb.DatabaseAccountOfferType_Standard
	kind := documentdb.DatabaseAccount_Kind_Spec_GlobalDocumentDB
	acct := &documentdb.DatabaseAccount{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("sqlacct")),
		Spec: documentdb.DatabaseAccount_Spec{
			Location:                 tc.AzureRegion,
			Owner:                    testcommon.AsOwner(rg),
			Kind:                     &kind,
			DatabaseAccountOfferType: &offerType,
			DisableLocalAuth:         to.Ptr(true),
			Locations: []documentdb.Location{
				{
					LocationName: to.Ptr("australiaeast"), // Capacity constraints // tc.AzureRegion
				},
			},
		},
	}

	// Declare a SQL database
	dbName := tc.Namer.GenerateName("sqldb")
	db := &documentdb.SqlDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(dbName),
		Spec: documentdb.SqlDatabase_Spec{
			Location: to.Ptr("australiaeast"), // Capacity constraints // tc.AzureRegion
			Owner:    testcommon.AsOwner(acct),
			Options: &documentdb.CreateUpdateOptions{
				AutoscaleSettings: &documentdb.AutoscaleSettings{
					MaxThroughput: to.Ptr(4000),
				},
			},
			Resource: &documentdb.SqlDatabaseResource{
				Id: &dbName,
			},
		},
	}
	tc.LogSectionf("Creating SQL account and database %q", dbName)
	tc.CreateResourcesAndWait(acct, db)

	acctId := *acct.Status.Id

	tc.T.Logf("SQL account and database successfully created")
	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB SQL RoleAssignment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_RoleAssignment_v20231115_CRUD(tc, rg, acct)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL Container CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Container_v20231115_CRUD(tc, db)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL Database throughputsettings CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Database_ThroughputSettings_v20231115_CRUD(tc, db)
			},
		})

	// There aren't any attributes to update for databases, other than
	// throughput settings once they're available.

	tc.DeleteResourceAndWait(acct)

	// Ensure that the resource was really deleted in Azure
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		acctId,
		string(documentdb.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func CosmosDB_SQL_Container_v20231115_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
	tc.LogSectionf("Creating SQL container")

	// Declare a SQL container
	name := tc.Namer.GenerateName("container")
	lastWriterWins := documentdb.ConflictResolutionPolicy_Mode_LastWriterWins
	consistent := documentdb.IndexingPolicy_IndexingMode_Consistent
	hash := documentdb.ContainerPartitionKey_Kind_Hash
	container := &documentdb.SqlDatabaseContainer{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.SqlDatabaseContainer_Spec{
			Location: to.Ptr("australiaeast"), // Capacity constraints // tc.AzureRegion
			Options: &documentdb.CreateUpdateOptions{
				Throughput: to.Ptr(400),
			},
			Owner: testcommon.AsOwner(db),
			Resource: &documentdb.SqlContainerResource{
				Id: &name,
				ConflictResolutionPolicy: &documentdb.ConflictResolutionPolicy{
					Mode: &lastWriterWins,
				},
				DefaultTtl: to.Ptr(200),
				IndexingPolicy: &documentdb.IndexingPolicy{
					IndexingMode: &consistent,
					IncludedPaths: []documentdb.IncludedPath{{
						Path: to.Ptr("/*"),
					}},
					ExcludedPaths: []documentdb.ExcludedPath{{
						Path: to.Ptr("/myPathToNotIndex/*"),
					}},
				},
				PartitionKey: &documentdb.ContainerPartitionKey{
					Kind:  &hash,
					Paths: []string{"/myPartitionKey"},
				},
			},
		},
	}

	tc.CreateResourceAndWait(container)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "CosmosDB SQL Trigger CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Trigger_v20231115_CRUD(tc, container)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL Stored Procedure CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_StoredProcedure_v20231115_CRUD(tc, container)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL User-defined Function CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_UserDefinedFunction_v20231115_CRUD(tc, container)
			},
		},
		testcommon.Subtest{
			Name: "CosmosDB SQL Container ThroughputSettings CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				CosmosDB_SQL_Database_Container_ThroughputSettings_v20231115_CRUD(tc, container)
			},
		})

	tc.LogSubsectionf("Updating the default TTL on container %q", name)
	old := container.DeepCopy()
	container.Spec.Resource.DefaultTtl = to.Ptr(400)
	tc.PatchResourceAndWait(old, container)
	tc.Expect(container.Status.Resource).ToNot(BeNil())
	tc.Expect(container.Status.Resource.DefaultTtl).ToNot(BeNil())
	tc.Expect(*container.Status.Resource.DefaultTtl).To(Equal(400))
}

func CosmosDB_SQL_Trigger_v20231115_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	tc.LogSectionf("Creating SQL Trigger")

	// Declare a trigger
	name := tc.Namer.GenerateName("trigger")
	pre := documentdb.SqlTriggerResource_TriggerType_Pre
	create := documentdb.SqlTriggerResource_TriggerOperation_Create
	trigger := documentdb.SqlDatabaseContainerTrigger{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.SqlDatabaseContainerTrigger_Spec{
			Location: to.Ptr("australiaeast"), // Capacity constraints // tc.AzureRegion
			Owner:    testcommon.AsOwner(container),
			Resource: &documentdb.SqlTriggerResource{
				Id:               &name,
				TriggerType:      &pre,
				TriggerOperation: &create,
				Body:             to.Ptr(triggerBody_v20231115),
			},
		},
	}

	tc.CreateResourceAndWait(&trigger)

	tc.LogSubsectionf("Updating the trigger type on trigger %q", name)
	post := documentdb.SqlTriggerResource_TriggerType_Post
	old := trigger.DeepCopy()
	trigger.Spec.Resource.TriggerType = &post
	tc.PatchResourceAndWait(old, &trigger)
	tc.Expect(trigger.Status.Resource).ToNot(BeNil())
	tc.Expect(trigger.Status.Resource.TriggerType).ToNot(BeNil())
	tc.Expect(string(*trigger.Status.Resource.TriggerType)).To(Equal("Post"))
}

const triggerBody_v20231115 = `
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

func CosmosDB_SQL_StoredProcedure_v20231115_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	name := tc.Namer.GenerateName("storedproc")
	tc.LogSectionf("Updating the body on stored procedure %q", name)

	// Declare a stored procedure
	storedProcedure := documentdb.SqlDatabaseContainerStoredProcedure{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.SqlDatabaseContainerStoredProcedure_Spec{
			Location: to.Ptr("australiaeast"), // Capacity constraints // tc.AzureRegion
			Owner:    testcommon.AsOwner(container),
			Resource: &documentdb.SqlStoredProcedureResource{
				Id:   &name,
				Body: to.Ptr(storedProcedureBody_v20231115),
			},
		},
	}

	tc.CreateResourceAndWait(&storedProcedure)

	old := storedProcedure.DeepCopy()
	newBody := "your code doesn't work!"
	storedProcedure.Spec.Resource.Body = &newBody
	tc.PatchResourceAndWait(old, &storedProcedure)
	tc.Expect(storedProcedure.Status.Resource).ToNot(BeNil())
	tc.Expect(storedProcedure.Status.Resource.Body).ToNot(BeNil())
	tc.Expect(*storedProcedure.Status.Resource.Body).To(Equal(newBody))
}

const storedProcedureBody_v20231115 = `
function () {
    var context = getContext();
    var response = context.getResponse();
    response.setBody('Hello, World');
}`

func CosmosDB_SQL_UserDefinedFunction_v20231115_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	name := tc.Namer.GenerateName("udf")
	tc.LogSectionf("Updating the body on user-defined function %q", name)

	// Declare a user defined function
	userDefinedFunction := documentdb.SqlDatabaseContainerUserDefinedFunction{
		ObjectMeta: tc.MakeObjectMetaWithName(name),
		Spec: documentdb.SqlDatabaseContainerUserDefinedFunction_Spec{
			AzureName: name,
			Location:  to.Ptr("australiaeast"), // Capacity constraints // tc.AzureRegion
			Owner:     testcommon.AsOwner(container),
			Resource: &documentdb.SqlUserDefinedFunctionResource{
				Id:   &name,
				Body: to.Ptr(userDefinedFunctionBody_v20231115),
			},
		},
	}

	// Create the resource
	tc.CreateResourceAndWait(&userDefinedFunction)

	old := userDefinedFunction.DeepCopy()
	newBody := "wonder what to do?"
	userDefinedFunction.Spec.Resource.Body = &newBody
	tc.PatchResourceAndWait(old, &userDefinedFunction)
	tc.Expect(userDefinedFunction.Status.Resource).ToNot(BeNil())
	tc.Expect(userDefinedFunction.Status.Resource.Body).ToNot(BeNil())
	tc.Expect(*userDefinedFunction.Status.Resource.Body).To(Equal(newBody))
}

const userDefinedFunctionBody_v20231115 = `
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

func CosmosDB_SQL_Database_ThroughputSettings_v20231115_CRUD(tc *testcommon.KubePerTestContext, db client.Object) {
	tc.LogSectionf("creating SQL database throughput")

	// Declare a throughput setting
	throughputSettings := documentdb.SqlDatabaseThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.SqlDatabaseThroughputSetting_Spec{
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

	// Create the resource
	tc.CreateResourceAndWait(&throughputSettings)
	// no DELETE, this is not a real resource - to delete it you must delete its parent

	// Ensure that the status is what we expect
	tc.Expect(throughputSettings.Status.Id).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings.MaxThroughput).To(Equal(to.Ptr(5000)))

	tc.LogSubsectionf("increase max throughput to 6000")
	old := throughputSettings.DeepCopy()
	throughputSettings.Spec.Resource.AutoscaleSettings.MaxThroughput = to.Ptr(6000)
	tc.PatchResourceAndWait(old, &throughputSettings)
	tc.Expect(throughputSettings.Status.Resource).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings).ToNot(BeNil())
	tc.Expect(throughputSettings.Status.Resource.AutoscaleSettings.MaxThroughput).To(Equal(to.Ptr(6000)))
	tc.T.Log("throughput successfully updated in status")
}

func CosmosDB_SQL_Database_Container_ThroughputSettings_v20231115_CRUD(tc *testcommon.KubePerTestContext, container client.Object) {
	tc.LogSectionf("creating SQL database container throughput")

	// Declare a throughput setting
	throughputSettings := documentdb.SqlDatabaseContainerThroughputSetting{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.Namer.GenerateName("throughput")),
		Spec: documentdb.SqlDatabaseContainerThroughputSetting_Spec{
			Owner: testcommon.AsOwner(container),
			Resource: &documentdb.ThroughputSettingsResource{
				Throughput: to.Ptr(500),
			},
		},
	}

	// Create the resource
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
	tc.T.Log("throughput successfully updated in status")
}

func CosmosDB_SQL_RoleAssignment_v20231115_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, acct *documentdb.DatabaseAccount) {
	tc.T.Logf("Creating a RoleAssignment")

	configMapName := "my-configmap"
	principalIdKey := "principalId"

	// Declare a managed identity
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: to.Ptr("australiaeast"), // Capacity constraints // tc.AzureRegion
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  principalIdKey,
					},
				},
			},
		},
	}

	// Create the resource
	tc.CreateResourceAndWait(mi)

	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

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

	// Declare a role assignment
	roleAssignment := &documentdb.SqlRoleAssignment{
		ObjectMeta: tc.MakeObjectMeta("roleassignment"),
		Spec: documentdb.SqlRoleAssignment_Spec{
			// Do not set AzureName here, it should be automatically set by webhook
			Owner: testcommon.AsOwner(acct),
			PrincipalIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  principalIdKey,
			},
			RoleDefinitionId: &roleDefinitionId,
			Scope:            &scope,
		},
	}

	// Create the resource
	tc.CreateResourceAndWait(roleAssignment)

	// Ensure that the status is what we expect
	tc.Expect(roleAssignment.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(roleAssignment)
}
