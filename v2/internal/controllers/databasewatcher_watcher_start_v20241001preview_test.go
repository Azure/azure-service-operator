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

	// The target names its server by fully qualified domain name, which Azure only assigns once the
	// server exists, so the server publishes it and the target reads it back
	const serverConfigMap = "sqlserver-config"
	const serverConfigMapKey = "fullyQualifiedDomainName"

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
			OperatorSpec: &sql.ServerOperatorSpec{
				ConfigMaps: &sql.ServerOperatorConfigMaps{
					FullyQualifiedDomainName: &genruntime.ConfigMapDestination{
						Name: serverConfigMap,
						Key:  serverConfigMapKey,
					},
				},
			},
		},
	}

	sqlDatabase := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("sqldb"),
		Spec: sql.ServersDatabase_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(server),
		},
	}

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

	target := &databasewatcher.Target{
		ObjectMeta: tc.MakeObjectMeta("target"),
		Spec: databasewatcher.Target_Spec{
			Owner: testcommon.AsOwner(watcher),
			Properties: &databasewatcher.TargetProperties{
				SqlDb: &databasewatcher.SqlDbSingleDatabaseTargetProperties{
					ConnectionServerNameFromConfig: &genruntime.ConfigMapReference{
						Name: serverConfigMap,
						Key:  serverConfigMapKey,
					},
					SqlDbResourceReference:   tc.MakeReferenceFromResource(sqlDatabase),
					TargetAuthenticationType: to.Ptr(databasewatcher.TargetAuthenticationType_Aad),
					TargetType:               to.Ptr(databasewatcher.SqlDbSingleDatabaseTargetProperties_TargetType_SqlDb),
				},
			},
		},
	}

	// Everything goes in together, as it would from a single kubectl apply, leaving the sequencing to the
	// operator: the target waits for its server's domain name and for the watcher, and the watcher can't
	// start until the target exists
	tc.CreateResourcesAndWait(server, sqlDatabase, watcher, target)

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
