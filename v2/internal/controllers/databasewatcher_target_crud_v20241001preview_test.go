/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

// A target points at a SQL database, so this creates a server and a database for it, along with the
// shared private link resource that reaches that server privately.
func Test_DatabaseWatcher_Target_v20241001preview_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Database Watcher isn't available in the default test region
	tc.AzureRegion = to.Ptr("eastus")

	adminPasswordSecretRef := createPasswordSecret("sqlsecret", "adminPassword", tc)

	rg := tc.CreateTestResourceGroupAndWait()

	// A target isn't ready until the watcher it belongs to runs, and a watcher can't run without
	// somewhere to write what it collects
	watcher := &databasewatcher.Watcher{
		ObjectMeta: tc.MakeObjectMeta("watcher"),
		Spec: databasewatcher.Watcher_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Identity: &databasewatcher.ManagedServiceIdentity{
				Type: to.Ptr(databasewatcher.ManagedServiceIdentityType_SystemAssigned),
			},
			Datastore: databaseWatcherDatastore(tc.AzureSubscription, rg.Name, "asotest-target-datastore"),
		},
	}

	server := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserver"),
		Spec: sql.Server_Spec{
			Location:                   tc.AzureRegion,
			Owner:                      testcommon.AsOwner(rg),
			AdministratorLogin:         to.Ptr("myadmin"),
			AdministratorLoginPassword: &adminPasswordSecretRef,
			Version:                    to.Ptr("12.0"),
		},
	}

	db := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.ServersDatabase_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(server),
		},
	}

	tc.CreateResourcesAndWait(watcher, server, db)

	target := &databasewatcher.Target{
		ObjectMeta: tc.MakeObjectMeta("target"),
		Spec: databasewatcher.Target_Spec{
			Owner: testcommon.AsOwner(watcher),
			Properties: &databasewatcher.TargetProperties{
				SqlDb: &databasewatcher.SqlDbSingleDatabaseTargetProperties{
					ConnectionServerName:     server.Status.FullyQualifiedDomainName,
					SqlDbResourceReference:   tc.MakeReferenceFromResource(db),
					TargetAuthenticationType: to.Ptr(databasewatcher.TargetAuthenticationType_Aad),
					TargetType:               to.Ptr(databasewatcher.SqlDbSingleDatabaseTargetProperties_TargetType_SqlDb),
				},
			},
		},
	}

	// DnsZone must be omitted for SQL logical servers
	sharedPrivateLink := &databasewatcher.SharedPrivateLink{
		ObjectMeta: tc.MakeObjectMeta("spl"),
		Spec: databasewatcher.SharedPrivateLink_Spec{
			Owner:                        testcommon.AsOwner(watcher),
			GroupId:                      to.Ptr("sqlServer"),
			PrivateLinkResourceReference: tc.MakeReferenceFromResource(server),
			RequestMessage:               to.Ptr("Please approve the connection from the database watcher"),
		},
	}

	tc.CreateResourcesAndWait(target, sharedPrivateLink)

	tc.Expect(target.Status.Id).ToNot(BeNil())
	tc.Expect(target.Status.Properties).ToNot(BeNil())
	tc.Expect(target.Status.Properties.SqlDb).ToNot(BeNil())

	// Status.Status reports the connection state, which the SQL server side fills in asynchronously, so
	// it isn't populated yet
	tc.Expect(sharedPrivateLink.Status.Id).ToNot(BeNil())
	tc.Expect(sharedPrivateLink.Status.GroupId).To(Equal(to.Ptr("sqlServer")))

	old := target.DeepCopy()
	target.Spec.Properties.SqlDb.ReadIntent = to.Ptr(true)
	tc.PatchResourceAndWait(old, target)
	tc.Expect(target.Status.Properties.SqlDb.ReadIntent).To(Equal(to.Ptr(true)))

	// The shared private link must be torn down before the server it points at
	splARMID := *sharedPrivateLink.Status.Id
	tc.DeleteResourceAndWait(sharedPrivateLink)

	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, splARMID, string(databasewatcher.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())

	targetARMID := *target.Status.Id
	tc.DeleteResourceAndWait(target)

	exists, retryAfter, err = tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, targetARMID, string(databasewatcher.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}
