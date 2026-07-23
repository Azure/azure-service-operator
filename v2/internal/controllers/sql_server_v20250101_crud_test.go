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

	"github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/keyvault/armkeyvault"

	keyvault "github.com/Azure/azure-service-operator/v2/api/keyvault/v1api20230701"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20230131"
	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v20250101"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-logical-server-aad-only-auth/azuredeploy.json
// for SQL Server inspiration
func Test_SQL_Server_v20250101_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Use a different region where we have quota
	tc.AzureRegion = to.Ptr("australiaeast")

	rg := tc.CreateTestResourceGroupAndWait()

	connectionConfigMap := "myconfig"
	connectionConfigMapKey := "fqdn"

	identityConfigMap := "adminconfig"

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					PrincipalId: &genruntime.ConfigMapDestination{
						Name: identityConfigMap,
						Key:  "oid",
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: identityConfigMap,
						Key:  "tenant",
					},
				},
				ConfigMapExpressions: []*core.DestinationExpression{
					{
						Name:  identityConfigMap,
						Key:   "name",
						Value: "self.spec.azureName",
					},
				},
			},
		},
	}

	server := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserver"),
		Spec: sql.Server_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Identity: &sql.ResourceIdentity{
				Type: to.Ptr(sql.IdentityType_UserAssigned),
				UserAssignedIdentities: []sql.UserAssignedIdentityDetails{
					{
						Reference: *tc.MakeReferenceFromResource(mi),
					},
				},
			},
			PrimaryUserAssignedIdentityReference: tc.MakeReferenceFromResource(mi),
			Administrators: &sql.ServerExternalAdministrator{
				AdministratorType:         to.Ptr(sql.AdministratorType_ActiveDirectory),
				PrincipalType:             to.Ptr(sql.PrincipalType_Application),
				AzureADOnlyAuthentication: to.Ptr(true),
				LoginFromConfig: &genruntime.ConfigMapReference{
					Name: identityConfigMap,
					Key:  "name",
				},
				Sid:      to.Ptr(sql.AzureCoreUuid("00000000-0000-0000-0000-000000000000")), // placeholder, would come from MI
				TenantId: to.Ptr(sql.AzureCoreUuid("00000000-0000-0000-0000-000000000000")), // placeholder, would come from MI
			},
			Version: to.Ptr("12.0"),
			OperatorSpec: &sql.ServerOperatorSpec{
				ConfigMaps: &sql.ServerOperatorConfigMaps{
					FullyQualifiedDomainName: &genruntime.ConfigMapDestination{
						Name: connectionConfigMap,
						Key:  connectionConfigMapKey,
					},
				},
			},
		},
	}

	tc.CreateResourcesAndWait(server, mi)
	tc.Expect(server.Status.Id).ToNot(BeNil())
	tc.Expect(server.Status.Location).To(Equal(tc.AzureRegion))

	tc.ExpectConfigMapHasKeysAndValues(
		connectionConfigMap,
		connectionConfigMapKey,
		*server.Status.FullyQualifiedDomainName,
	)

	storageDetails := makeStorageAccountForSQLVulnerabilityAssessment_v20250101(tc, rg)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SQL Database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Databases_v20250101_CRUD(tc, server, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL ConnectionPolicy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_ConnectionPolicy_v20250101_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL AdvancedThreatProtection CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_AdvancedThreatProtection_v20250101_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL VulnerabilityAssessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_VulnerabilityAssessments_v20250101_CRUD(tc, server, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL FirewallRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_FirewallRules_v20250101_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL IPV6 FirewallRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_IPV6_FirewallRules_v20250101_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL Outbound FirewallRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_OutboundFirewallRule_v20250101_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL ElasticPool CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_ElasticPool_v20250101_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL VirtualNetworkRule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_VirtualNetworkRule_v20250101_CRUD(tc, server, rg)
			},
		},
		testcommon.Subtest{
			Name: "SQL AuditingSetting CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_AuditingSetting_v20250101_CRUD(tc, server, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL ServersKey and EncryptionProtector CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_Key_And_EncryptionProtector_v20250101_CRUD(tc, server, rg, mi)
			},
		},
	)

	armId := *server.Status.Id
	tc.DeleteResourceAndWait(server)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

func SQL_Server_ConnectionPolicy_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	connectionType := sql.ServerConnectionType_Default
	policy := &sql.ServersConnectionPolicy{
		ObjectMeta: tc.MakeObjectMeta("connpolicy"),
		Spec: sql.ServersConnectionPolicy_Spec{
			Owner:          testcommon.AsOwner(server),
			ConnectionType: &connectionType,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&policy.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(policy)

	tc.Expect(policy.Status.Id).ToNot(BeNil())
}

func SQL_Server_AdvancedThreatProtection_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	advancedThreatProtection := sql.AdvancedThreatProtectionState_Enabled
	policy := &sql.ServersAdvancedThreatProtectionSetting{
		ObjectMeta: tc.MakeObjectMeta("atp"),
		Spec: sql.ServersAdvancedThreatProtectionSetting_Spec{
			Owner: testcommon.AsOwner(server),
			State: &advancedThreatProtection,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&policy.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(policy)

	tc.Expect(policy.Status.Id).ToNot(BeNil())
}

func SQL_Server_VulnerabilityAssessments_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, storageDetails vulnStorageAccountDetails) {
	enabled := sql.SecurityAlertsPolicyState_Enabled
	alertPolicy := &sql.ServersSecurityAlertPolicy{
		ObjectMeta: tc.MakeObjectMeta("alertpolicy"),
		Spec: sql.ServersSecurityAlertPolicy_Spec{
			Owner: testcommon.AsOwner(server),
			State: &enabled,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&alertPolicy.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	vulnerabilityAssessment := &sql.ServersVulnerabilityAssessment{
		ObjectMeta: tc.MakeObjectMeta("vulnassessment"),
		Spec: sql.ServersVulnerabilityAssessment_Spec{
			Owner: testcommon.AsOwner(server),
			RecurringScans: &sql.VulnerabilityAssessmentRecurringScansProperties{
				IsEnabled: to.Ptr(false),
			},
			StorageAccountAccessKey: &genruntime.SecretReference{
				Name: storageDetails.secretName,
				Key:  storageDetails.keySecretKey,
			},
			// TODO: Make this easier to build a combined path in ASO itself
			StorageContainerPath: to.Ptr(fmt.Sprintf("%s%s", blobEndpoint, storageDetails.container)),
		},
	}

	tc.CreateResourcesAndWait(alertPolicy, vulnerabilityAssessment)

	tc.Expect(alertPolicy.Status.Id).ToNot(BeNil())
	tc.Expect(vulnerabilityAssessment.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(vulnerabilityAssessment)
}

func SQL_Server_FirewallRules_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	firewall := &sql.ServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: sql.ServersFirewallRule_Spec{
			Owner:          testcommon.AsOwner(server),
			StartIpAddress: to.Ptr("0.0.0.0"),
			EndIpAddress:   to.Ptr("0.0.0.0"),
		},
	}

	tc.CreateResourceAndWait(firewall)

	tc.Expect(firewall.Status.Id).ToNot(BeNil())
	armId := *firewall.Status.Id

	tc.DeleteResourceAndWait(firewall)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

func SQL_Server_IPV6_FirewallRules_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	firewall := &sql.ServersIPV6FirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: sql.ServersIPV6FirewallRule_Spec{
			Owner:            testcommon.AsOwner(server),
			StartIPv6Address: to.Ptr("2001:db8::"),
			EndIPv6Address:   to.Ptr("2001:db8:0000:0000:0000:0000:00ff:ffff"),
		},
	}

	tc.CreateResourceAndWait(firewall)

	tc.Expect(firewall.Status.Id).ToNot(BeNil())
	armId := *firewall.Status.Id

	tc.DeleteResourceAndWait(firewall)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

func SQL_Server_AuditingSetting_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, storageDetails vulnStorageAccountDetails) {
	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	enabled := sql.BlobAuditingPolicyState_Enabled
	auditingSetting := &sql.ServersAuditingSetting{
		ObjectMeta: tc.MakeObjectMeta("audit"),
		Spec: sql.ServersAuditingSetting_Spec{
			Owner:                        testcommon.AsOwner(server),
			State:                        &enabled,
			StorageAccountSubscriptionId: &tc.AzureSubscription, // TODO: Make this easier for users to set? Via configmap?
			StorageAccountAccessKey: &genruntime.SecretReference{
				Name: storageDetails.secretName,
				Key:  storageDetails.keySecretKey,
			},
			StorageEndpoint: &blobEndpoint,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&auditingSetting.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(auditingSetting)

	tc.Expect(auditingSetting.Status.Id).ToNot(BeNil())
}

func SQL_Server_OutboundFirewallRule_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	outboundRule := &sql.ServersOutboundFirewallRule{
		ObjectMeta: tc.MakeObjectMetaWithName("outboundrule"),
		Spec: sql.ServersOutboundFirewallRule_Spec{
			Owner:     testcommon.AsOwner(server),
			AzureName: "server.database.windows.net",
		},
	}

	tc.CreateResourceAndWait(outboundRule)

	tc.Expect(outboundRule.Status.Id).ToNot(BeNil())
	armId := *outboundRule.Status.Id

	tc.DeleteResourceAndWait(outboundRule)
	exists, _, err := tc.AzureClient.CheckExistenceWithGetByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value),
	)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-elastic-pool-create/azuredeploy.json
// for SQL Server Elastic Pool examples
func SQL_Server_ElasticPool_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	elasticPool := &sql.ServersElasticPool{
		ObjectMeta: tc.MakeObjectMeta("pool"),
		Spec: sql.ServersElasticPool_Spec{
			Owner:    testcommon.AsOwner(server),
			Location: tc.AzureRegion,
			Sku: &sql.Sku{
				Name:     to.Ptr("StandardPool"),
				Tier:     to.Ptr("Standard"),
				Capacity: to.Ptr(100),
			},
			PerDatabaseSettings: &sql.ElasticPoolPerDatabaseSettings{
				MinCapacity: to.Ptr(float64(0)),
				MaxCapacity: to.Ptr(float64(100)),
			},
		},
	}

	tc.CreateResourceAndWait(elasticPool)

	tc.Expect(elasticPool.Status.Id).ToNot(BeNil())
	armId := *elasticPool.Status.Id

	tc.DeleteResourceAndWait(elasticPool)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.web/asev2-appservice-sql-vpngw/azuredeploy.json
func SQL_Server_VirtualNetworkRule_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup) {
	subnet := makeSubnetForSQLServer_v20250101(tc, rg)

	vnetRule := &sql.ServersVirtualNetworkRule{
		ObjectMeta: tc.MakeObjectMeta("pool"),
		Spec: sql.ServersVirtualNetworkRule_Spec{
			Owner:                         testcommon.AsOwner(server),
			VirtualNetworkSubnetReference: tc.MakeReferenceFromResource(subnet),
		},
	}

	tc.CreateResourceAndWait(vnetRule)

	tc.Expect(vnetRule.Status.Id).ToNot(BeNil())
	armId := *vnetRule.Status.Id

	tc.DeleteResourceAndWait(vnetRule)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-database-transparent-encryption-create/azuredeploy.json
// for a SQL Database example
func SQL_Databases_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, storageDetails vulnStorageAccountDetails) {
	db := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.ServersDatabase_Spec{
			Owner:     testcommon.AsOwner(server),
			Location:  tc.AzureRegion,
			Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		},
	}

	tc.CreateResourceAndWait(db)

	tc.Expect(db.Status.Id).ToNot(BeNil())
	armId := *db.Status.Id

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SQL Database BackupLongTermRetentionPolicy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_BackupLongTermRetention_v20250101_CRUD(tc, db)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Vulnerability Assessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_VulnerabilityAssessment_v20250101_CRUD(tc, db, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Auditing Setting CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_AuditingSetting_v20250101_CRUD(tc, db, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Transparent Data Encryption Assessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_TransparentDataEncryption_v20250101_CRUD(tc, db)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Advanced Threat Protection CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_AdvancedThreatProtection_v20250101_CRUD(tc, db)
			},
		},
	)

	tc.DeleteResourceAndWait(db)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

func SQL_BackupLongTermRetention_v20250101_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
	policy := &sql.ServersDatabasesBackupLongTermRetentionPolicy{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.ServersDatabasesBackupLongTermRetentionPolicy_Spec{
			Owner:           testcommon.AsOwner(db),
			WeeklyRetention: to.Ptr("P30D"),
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&policy.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(policy)
	tc.Expect(policy.Status.Id).ToNot(BeNil())
}

func SQL_Database_VulnerabilityAssessment_v20250101_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase, storageDetails vulnStorageAccountDetails) {
	enabled := sql.SecurityAlertsPolicyState_Enabled
	securityAlertPolicy := &sql.ServersDatabasesSecurityAlertPolicy{
		ObjectMeta: tc.MakeObjectMeta("alertpolicy"),
		Spec: sql.ServersDatabasesSecurityAlertPolicy_Spec{
			Owner: testcommon.AsOwner(db),
			State: &enabled,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&securityAlertPolicy.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	vulnerabilityAssessment := &sql.ServersDatabasesVulnerabilityAssessment{
		ObjectMeta: tc.MakeObjectMeta("vulnassessment"),
		Spec: sql.ServersDatabasesVulnerabilityAssessment_Spec{
			Owner: testcommon.AsOwner(db),
			RecurringScans: &sql.VulnerabilityAssessmentRecurringScansProperties{
				IsEnabled: to.Ptr(false),
			},
			StorageAccountAccessKey: &genruntime.SecretReference{
				Name: storageDetails.secretName,
				Key:  storageDetails.keySecretKey,
			},
			// TODO: Make this easier to build a combined path in ASO itself
			StorageContainerPath: to.Ptr(fmt.Sprintf("%s%s", blobEndpoint, storageDetails.container)),
		},
	}

	tc.CreateResourcesAndWait(securityAlertPolicy, vulnerabilityAssessment)

	tc.Expect(securityAlertPolicy.Status.Id).ToNot(BeNil())
	tc.Expect(vulnerabilityAssessment.Status.Id).ToNot(BeNil())

	// This is an odd case - the DELETE is accepted by the service, but isn't actually honoured: the resource still exists after deletion
	// So we'll skip our usual verification to make sure it's gone.
	tc.DeleteResourceAndWait(vulnerabilityAssessment)
}

func SQL_Database_AuditingSetting_v20250101_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase, storageDetails vulnStorageAccountDetails) {
	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	enabled := sql.BlobAuditingPolicyState_Enabled
	auditingSetting := &sql.ServersDatabasesAuditingSetting{
		ObjectMeta: tc.MakeObjectMeta("audit"),
		Spec: sql.ServersDatabasesAuditingSetting_Spec{
			Owner:                        testcommon.AsOwner(db),
			State:                        &enabled,
			StorageAccountSubscriptionId: &tc.AzureSubscription, // TODO: Make this easier for users to set? Via configmap?
			StorageAccountAccessKey: &genruntime.SecretReference{
				Name: storageDetails.secretName,
				Key:  storageDetails.keySecretKey,
			},
			StorageEndpoint: &blobEndpoint,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&auditingSetting.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(auditingSetting)

	tc.Expect(auditingSetting.Status.Id).ToNot(BeNil())
}

func SQL_Database_TransparentDataEncryption_v20250101_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
	enabled := sql.TransparentDataEncryptionState_Enabled
	transparentDataEncryption := &sql.ServersDatabasesTransparentDataEncryption{
		ObjectMeta: tc.MakeObjectMeta("encrypt"),
		Spec: sql.ServersDatabasesTransparentDataEncryption_Spec{
			Owner: testcommon.AsOwner(db),
			State: &enabled,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&transparentDataEncryption.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(transparentDataEncryption)

	tc.Expect(transparentDataEncryption.Status.Id).ToNot(BeNil())
}

func SQL_Database_AdvancedThreatProtection_v20250101_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
	enabled := sql.AdvancedThreatProtectionState_Enabled
	advancedProtection := &sql.ServersDatabasesAdvancedThreatProtectionSetting{
		ObjectMeta: tc.MakeObjectMeta("atp"),
		Spec: sql.ServersDatabasesAdvancedThreatProtectionSetting_Spec{
			Owner: testcommon.AsOwner(db),
			State: &enabled,
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&advancedProtection.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(advancedProtection)

	tc.Expect(advancedProtection.Status.Id).ToNot(BeNil())
}

func makeStorageAccountForSQLVulnerabilityAssessment_v20250101(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) vulnStorageAccountDetails {
	const vulnerabilityAssessmentsContainerName = "vulnerabilityassessments"
	const secretName = "storagesecret"

	acct := newStorageAccount(tc, rg)
	acct.Spec.OperatorSpec = &storage.StorageAccountOperatorSpec{
		Secrets: &storage.StorageAccountOperatorSecrets{
			BlobEndpoint: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "blobEndpoint",
			},
			Key1: &genruntime.SecretDestination{
				Name: secretName,
				Key:  "key1",
			},
		},
	}

	blobService := &storage.StorageAccountsBlobService{
		ObjectMeta: tc.MakeObjectMeta("blobservice"),
		Spec: storage.StorageAccountsBlobService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	// We can delete it from the cluster by applying this annotation, but this won't change anything in Azure.
	tc.AddAnnotation(&blobService.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMetaWithName(vulnerabilityAssessmentsContainerName),
		Spec: storage.StorageAccountsBlobServicesContainer_Spec{
			Owner: testcommon.AsOwner(blobService),
		},
	}

	tc.CreateResourcesAndWait(acct, blobService, blobContainer)

	return vulnStorageAccountDetails{
		container:             vulnerabilityAssessmentsContainerName,
		secretName:            secretName,
		blobEndpointSecretKey: "blobEndpoint",
		keySecretKey:          "key1",
	}
}

func makeSubnetForSQLServer_v20250101(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *network.VirtualNetworksSubnet {
	vnet := &network.VirtualNetwork{
		ObjectMeta: tc.MakeObjectMeta("vn"),
		Spec: network.VirtualNetwork_Spec{
			Owner:    testcommon.AsOwner(rg),
			Location: tc.AzureRegion,
			AddressSpace: &network.AddressSpace{
				AddressPrefixes: []string{"10.0.0.0/16"},
			},
		},
	}
	privateEndpointNetworkPolicyDisabled := network.SubnetPropertiesFormat_PrivateEndpointNetworkPolicies_Disabled
	privateLinkNetworkPolicyDisabled := network.SubnetPropertiesFormat_PrivateLinkServiceNetworkPolicies_Disabled

	subnet := &network.VirtualNetworksSubnet{
		ObjectMeta: tc.MakeObjectMeta("subnet"),
		Spec: network.VirtualNetworksSubnet_Spec{
			Owner:         testcommon.AsOwner(vnet),
			AddressPrefix: to.Ptr("10.0.0.0/24"),
			ServiceEndpoints: []network.ServiceEndpointPropertiesFormat{
				{
					Service: to.Ptr("Microsoft.Sql"),
				},
			},
			PrivateEndpointNetworkPolicies:    &privateEndpointNetworkPolicyDisabled,
			PrivateLinkServiceNetworkPolicies: &privateLinkNetworkPolicyDisabled,
		},
	}
	tc.CreateResourcesAndWait(vnet, subnet)

	return subnet
}

func SQL_Server_Key_And_EncryptionProtector_v20250101_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup, mi *managedidentity.UserAssignedIdentity) {
	// Create a KeyVault for the SQL Server TDE key
	// The KeyVault needs access policies for both:
	// 1. The current user/SP (to create keys via ARM SDK)
	// 2. The SQL Server's managed identity (so SQL can use the key for TDE)
	kv := &keyvault.Vault{
		ObjectMeta: tc.MakeObjectMeta("sqlkv"),
		Spec: keyvault.Vault_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Properties: &keyvault.VaultProperties{
				CreateMode: to.Ptr(keyvault.VaultProperties_CreateMode_CreateOrRecover),
				Sku: &keyvault.Sku{
					Family: to.Ptr(keyvault.Sku_Family_A),
					Name:   to.Ptr(keyvault.Sku_Name_Standard),
				},
				TenantId:                     to.Ptr(tc.AzureTenant),
				EnableSoftDelete:             to.Ptr(true),
				SoftDeleteRetentionInDays:    to.Ptr(7),
				EnablePurgeProtection:        to.Ptr(true),
				EnabledForDiskEncryption:     to.Ptr(true),
				EnableRbacAuthorization:      to.Ptr(false),
				EnabledForDeployment:         to.Ptr(true),
				EnabledForTemplateDeployment: to.Ptr(true),
				AccessPolicies: []keyvault.AccessPolicyEntry{
					{
						// Access policy for the current user/SP running the test
						TenantId: to.Ptr(tc.AzureTenant),
						ObjectId: to.Ptr(tc.AzureTenant),
						Permissions: &keyvault.Permissions{
							Keys: []keyvault.Permissions_Keys{
								keyvault.Permissions_Keys_Get,
								keyvault.Permissions_Keys_List,
								keyvault.Permissions_Keys_Create,
								keyvault.Permissions_Keys_UnwrapKey,
								keyvault.Permissions_Keys_WrapKey,
							},
						},
					},
					{
						// Access policy for the SQL Server's managed identity
						TenantId: to.Ptr(tc.AzureTenant),
						ObjectId: to.Ptr(*mi.Status.PrincipalId),
						Permissions: &keyvault.Permissions{
							Keys: []keyvault.Permissions_Keys{
								keyvault.Permissions_Keys_Get,
								keyvault.Permissions_Keys_List,
								keyvault.Permissions_Keys_UnwrapKey,
								keyvault.Permissions_Keys_WrapKey,
							},
						},
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(kv)

	// Create a key in the KeyVault using the ARM SDK
	key := createKeyVaultKeyForSQL_v20250101(tc, kv, rg)
	tc.Expect(key.Properties.KeyURIWithVersion).ToNot(BeNil())

	keyURI := *key.Properties.KeyURIWithVersion

	// The Azure name for a ServersKey must follow the pattern: <vaultname>_<keyname>_<keyversion>
	// Extract parts from the key URI: https://<vaultname>.vault.azure.net/keys/<keyname>/<keyversion>
	vaultName := kv.AzureName()
	keyName := "sqlenckey"
	// Extract the key version from the URI
	parts := strings.Split(keyURI, "/")
	keyVersion := parts[len(parts)-1]
	serverKeyName := fmt.Sprintf("%s_%s_%s", vaultName, keyName, keyVersion)

	// Create the ServersKey resource
	// The Azure name must follow the pattern <vaultname>_<keyname>_<keyversion> (with underscores),
	// but the Kubernetes metadata.name must be a valid RFC 1123 subdomain (no underscores).
	serversKey := &sql.ServersKey{
		ObjectMeta: tc.MakeObjectMeta("serverkey"),
		Spec: sql.ServersKey_Spec{
			AzureName:     serverKeyName,
			Owner:         testcommon.AsOwner(server),
			ServerKeyType: to.Ptr(sql.ServerKeyType_AzureKeyVault),
			Uri:           to.Ptr(keyURI),
		},
	}

	tc.CreateResourceAndWait(serversKey)
	tc.Expect(serversKey.Status.Id).ToNot(BeNil())

	// Create the ServersEncryptionProtector resource pointing to the key
	encryptionProtector := &sql.ServersEncryptionProtector{
		ObjectMeta: tc.MakeObjectMeta("current"),
		Spec: sql.ServersEncryptionProtector_Spec{
			Owner:               testcommon.AsOwner(server),
			ServerKeyType:       to.Ptr(sql.ServerKeyType_AzureKeyVault),
			ServerKeyName:       to.Ptr(serverKeyName),
			AutoRotationEnabled: to.Ptr(true),
		},
	}

	// Don't try to delete directly, this is not a real resource - to delete it in Azure you must delete its parent.
	tc.AddAnnotation(&encryptionProtector.ObjectMeta, "serviceoperator.azure.com/reconcile-policy", "detach-on-delete")

	tc.CreateResourceAndWait(encryptionProtector)
	tc.Expect(encryptionProtector.Status.Id).ToNot(BeNil())
}

func createKeyVaultKeyForSQL_v20250101(tc *testcommon.KubePerTestContext, kv *keyvault.Vault, rg *resources.ResourceGroup) armkeyvault.Key {
	client, err := armkeyvault.NewKeysClient(tc.AzureSubscription, tc.AzureClient.Creds(), tc.AzureClient.ClientOptions())
	tc.Expect(err).To(BeNil())

	keyProperties := armkeyvault.KeyCreateParameters{
		Properties: &armkeyvault.KeyProperties{
			Attributes: &armkeyvault.KeyAttributes{
				Enabled: to.Ptr(true),
			},
			KeySize: to.Ptr(int32(2048)),
			Kty:     to.Ptr(armkeyvault.JSONWebKeyTypeRSA),
		},
	}

	keyResp, err := client.CreateIfNotExist(tc.Ctx, rg.Name, kv.AzureName(), "sqlenckey", keyProperties, nil)
	tc.Expect(err).ToNot(HaveOccurred())

	return keyResp.Key
}
