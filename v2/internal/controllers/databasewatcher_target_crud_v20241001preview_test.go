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
	managedIdentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// A target points at a SQL database, so this creates a server and a database for it, along with the
// shared private link resource that reaches that server privately.
func Test_DatabaseWatcher_Target_v20241001preview_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Database Watcher isn't available in the default test region
	tc.AzureRegion = to.Ptr("australiaeast")

	rg := tc.CreateTestResourceGroupAndWait()

	// Managed Identity used for admin access to the SQL Server
	const (
		identityConfigMap         = "sqlidentity-config"
		identityConfigClientId    = "name"
		identityConfigPrincipalID = "principalId"
		identityConfigTenantID    = "tenantId"
	)

	identity := &managedIdentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("sqlidentity"),
		Spec: managedIdentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedIdentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedIdentity.UserAssignedIdentityOperatorConfigMaps{
					ClientId: &genruntime.ConfigMapDestination{
						Name: identityConfigMap,
						Key:  identityConfigClientId,
					},
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: identityConfigMap,
						Key:  identityConfigPrincipalID,
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: identityConfigMap,
						Key:  identityConfigTenantID,
					},
				},
			},
		},
	}

	const kustoClusterName = "asotest-watcher-datastore"

	// The target names its server by fully qualified domain name, which Azure only assigns once the
	// server exists, so the server publishes it and the target reads it back
	const (
		serverConfigMap    = "sqlserver-config"
		serverConfigMapKey = "fullyQualifiedDomainName"
	)

	server := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserver"),
		Spec: sql.Server_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Administrators: &sql.ServerExternalAdministrator{
				AdministratorType:         to.Ptr(sql.ServerExternalAdministrator_AdministratorType_ActiveDirectory),
				AzureADOnlyAuthentication: to.Ptr(true),
				PrincipalType:             to.Ptr(sql.ServerExternalAdministrator_PrincipalType_Application),
				LoginFromConfig: &genruntime.ConfigMapReference{
					Name: identityConfigMap,
					Key:  identityConfigClientId,
				},
				SidFromConfig: &genruntime.ConfigMapReference{
					Name: identityConfigMap,
					Key:  identityConfigPrincipalID,
				},
				TenantIdFromConfig: &genruntime.ConfigMapReference{
					Name: identityConfigMap,
					Key:  identityConfigTenantID,
				},
			},
			Version: to.Ptr("12.0"),
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
			Location: tc.AzureRegion,
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

	// Everything goes in together, as it would from a single kubectl apply, leaving the sequencing to the
	// operator: the target waits for its server's domain name and for the watcher, and the watcher can't
	// start until the target exists
	tc.CreateResourcesAndWait(identity, server, sqlDatabase, watcher, target, sharedPrivateLink)

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

	// The status on the watcher itself is only as fresh as the watcher's last reconcile, so ask Azure
	tc.Expect(watcherStatusFromAzure(tc, watcher)).To(Equal("Running"))

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
