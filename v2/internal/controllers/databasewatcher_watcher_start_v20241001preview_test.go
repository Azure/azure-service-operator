/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	databasewatcher "github.com/Azure/azure-service-operator/v2/api/databasewatcher/v20241001preview"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// ARM creates every watcher stopped and refuses to start one without a data store or a target, so this
// supplies both and expects the operator to start it
func Test_DatabaseWatcher_WatcherStart_v20241001preview(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)

	// Database Watcher isn't available in the default test region, but the resources it reads from and
	// writes to can stay there
	watcherRegion := to.Ptr("eastus")

	const kustoClusterName = "asotest-watcher-datastore"

	adminPasswordSecretRef := createPasswordSecret("sqlsecret", "adminPassword", tc)

	rg := tc.CreateTestResourceGroupAndWait()

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

	sqlDatabase := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("sqldb"),
		Spec: sql.ServersDatabase_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(server),
		},
	}

	tc.CreateResourcesAndWait(server, sqlDatabase)

	watcher := &databasewatcher.Watcher{
		ObjectMeta: tc.MakeObjectMeta("watcher"),
		Spec: databasewatcher.Watcher_Spec{
			Location: watcherRegion,
			Owner:    testcommon.AsOwner(rg),
			Identity: &databasewatcher.ManagedServiceIdentity{
				Type: to.Ptr(databasewatcher.ManagedServiceIdentityType_SystemAssigned),
			},
			Datastore: databaseWatcherDatastore(tc.AzureSubscription, rg.Name, kustoClusterName),
		},
	}

	// The watcher is ready before it runs: ARM won't start it until it has a target, and a target can't
	// be created until its owner is ready
	tc.CreateResourceAndWait(watcher)
	tc.Expect(watcher.Status.Status).To(Equal(to.Ptr(databasewatcher.WatcherStatus_STATUS_Stopped)))

	target := &databasewatcher.Target{
		ObjectMeta: tc.MakeObjectMeta("target"),
		Spec: databasewatcher.Target_Spec{
			Owner: testcommon.AsOwner(watcher),
			Properties: &databasewatcher.TargetProperties{
				SqlDb: &databasewatcher.SqlDbSingleDatabaseTargetProperties{
					ConnectionServerName:     server.Status.FullyQualifiedDomainName,
					SqlDbResourceReference:   tc.MakeReferenceFromResource(sqlDatabase),
					TargetAuthenticationType: to.Ptr(databasewatcher.TargetAuthenticationType_Aad),
					TargetType:               to.Ptr(databasewatcher.SqlDbSingleDatabaseTargetProperties_TargetType_SqlDb),
				},
			},
		},
	}

	// The target isn't ready until the watcher it belongs to is running, so this waits for the start
	tc.CreateResourceAndWait(target)

	// The status on the watcher itself is only as fresh as the watcher's last reconcile, so ask Azure
	tc.Expect(watcherStatusFromAzure(tc, watcher)).To(Equal("Running"))

	tc.DeleteResourceAndWait(target)
	tc.DeleteResourceAndWait(watcher)
}

// databaseWatcherDatastore builds the data store a watcher writes to. Azure records it without checking
// that it leads anywhere, so the cluster named here need not exist.
func databaseWatcherDatastore(subscription string, resourceGroup string, clusterName string) *databasewatcher.Datastore {
	kustoClusterID := fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Kusto/clusters/%s",
		subscription,
		resourceGroup,
		clusterName,
	)

	// KustoClusterDisplayName is optional in the ARM schema, but ARM rejects a data store without one for
	// the adx and free offerings
	return &databasewatcher.Datastore{
		AdxClusterResourceReference: &genruntime.ResourceReference{ARMID: kustoClusterID},
		KustoClusterDisplayName:     to.Ptr(clusterName),
		KustoClusterUri:             to.Ptr(fmt.Sprintf("https://%s.eastus.kusto.windows.net", clusterName)),
		KustoDataIngestionUri:       to.Ptr(fmt.Sprintf("https://ingest-%s.eastus.kusto.windows.net", clusterName)),
		KustoDatabaseName:           to.Ptr("watcher-data"),
		KustoManagementUrl:          to.Ptr("https://portal.azure.com/"),
		KustoOfferingType:           to.Ptr(databasewatcher.KustoOfferingType_Adx),
	}
}

func watcherStatusFromAzure(tc *testcommon.KubePerTestContext, watcher *databasewatcher.Watcher) string {
	var state struct {
		Properties struct {
			Status string `json:"status"`
		} `json:"properties"`
	}

	_, err := tc.AzureClient.GetByID(tc.Ctx, *watcher.Status.Id, watcher.GetAPIVersion(), &state)
	tc.Expect(err).ToNot(HaveOccurred())

	return state.Properties.Status
}
