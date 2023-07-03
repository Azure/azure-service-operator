/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	network "github.com/Azure/azure-service-operator/v2/api/network/v1api20201101"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	storage "github.com/Azure/azure-service-operator/v2/api/storage/v1api20210401"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-logical-server/azuredeploy.json
// for SQL Server inspiration
func Test_SQL_Server_CRUD(t *testing.T) {
	t.Parallel()

	tc := globalTestContext.ForTest(t)
	// Use a different region where we have quota
	tc.AzureRegion = to.Ptr("eastus")

	secretName := "sqlsecret"
	adminPasswordKey := "adminPassword"
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	rg := tc.CreateTestResourceGroupAndWait()

	configMap := "myconfig"
	configMapKey := "fqdn"

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
						Name: configMap,
						Key:  configMapKey,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(server)
	tc.Expect(server.Status.Id).ToNot(BeNil())
	tc.Expect(server.Status.Location).To(Equal(tc.AzureRegion))

	tc.ExpectConfigMapHasKeysAndValues(
		configMap,
		configMapKey,
		*server.Status.FullyQualifiedDomainName)

	storageDetails := makeStorageAccountForSQLVulnerabilityAssessment(tc, rg)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "SQL Database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Databases_CRUD(tc, server, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL ConnectionPolicy CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_ConnectionPolicy_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL AdvancedThreatProtection CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_AdvancedThreatProtection_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL VulnerabilityAssessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_VulnerabilityAssessments_CRUD(tc, server, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL FirewallRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_FirewallRules_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL IPV6 FirewallRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_IPV6_FirewallRules_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL Outbound FirewallRules CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_OutboundFirewallRule_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL ElasticPool CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_ElasticPool_CRUD(tc, server)
			},
		},
		testcommon.Subtest{
			Name: "SQL VirtualNetworkRule CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_VirtualNetworkRule_CRUD(tc, server, rg)
			},
		},
		testcommon.Subtest{
			Name: "SQL AuditingSetting CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Server_AuditingSetting_CRUD(tc, server, storageDetails)
			},
		})

	armId := *server.Status.Id
	tc.DeleteResourceAndWait(server)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

func SQL_Server_ConnectionPolicy_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	connectionType := sql.ServerConnectionPolicyProperties_ConnectionType_Default
	policy := &sql.ServersConnectionPolicy{
		ObjectMeta: tc.MakeObjectMeta("connpolicy"),
		Spec: sql.Servers_ConnectionPolicy_Spec{
			Owner:          testcommon.AsOwner(server),
			ConnectionType: &connectionType,
		},
	}

	tc.CreateResourceAndWait(policy)

	tc.Expect(policy.Status.Id).ToNot(BeNil())

	// TODO: Delete is not allowed for this resource
	//tc.DeleteResourceAndWait(policy)
}

func SQL_Server_AdvancedThreatProtection_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	advancedThreatProtection := sql.AdvancedThreatProtectionProperties_State_Enabled
	policy := &sql.ServersAdvancedThreatProtectionSetting{
		ObjectMeta: tc.MakeObjectMeta("atp"),
		Spec: sql.Servers_AdvancedThreatProtectionSetting_Spec{
			Owner: testcommon.AsOwner(server),
			State: &advancedThreatProtection,
		},
	}

	tc.CreateResourceAndWait(policy)

	tc.Expect(policy.Status.Id).ToNot(BeNil())

	// TODO: Delete is not allowed for this resource
	// tc.DeleteResourceAndWait(policy)
}

func SQL_Server_VulnerabilityAssessments_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, storageDetails vulnStorageAccountDetails) {
	enabled := sql.ServerSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled
	alertPolicy := &sql.ServersSecurityAlertPolicy{
		ObjectMeta: tc.MakeObjectMeta("alertpolicy"),
		Spec: sql.Servers_SecurityAlertPolicy_Spec{
			Owner: testcommon.AsOwner(server),
			State: &enabled,
		},
	}

	tc.CreateResourceAndWait(alertPolicy)

	tc.Expect(alertPolicy.Status.Id).ToNot(BeNil())

	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	vulnerabilityAssessment := &sql.ServersVulnerabilityAssessment{
		ObjectMeta: tc.MakeObjectMeta("vulnassessment"),
		Spec: sql.Servers_VulnerabilityAssessment_Spec{
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

	tc.CreateResourceAndWait(vulnerabilityAssessment)

	tc.Expect(vulnerabilityAssessment.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(vulnerabilityAssessment)
	// TODO: It seems like delete of this resource isn't actually honored by the service - it's accepted but doesn't remove the resource
	//exists, _, err := tc.AzureClient.HeadByID(
	//	tc.Ctx,
	//	armId,
	//	string(sql.APIVersion_Value))
	//tc.Expect(err).ToNot(HaveOccurred())
	//tc.Expect(exists).To(BeFalse())
}

func SQL_Server_FirewallRules_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	firewall := &sql.ServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: sql.Servers_FirewallRule_Spec{
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

func SQL_Server_IPV6_FirewallRules_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	firewall := &sql.ServersIPV6FirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: sql.Servers_Ipv6FirewallRule_Spec{
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

func SQL_Server_AuditingSetting_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, storageDetails vulnStorageAccountDetails) {
	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	enabled := sql.ServerBlobAuditingPolicyProperties_State_Enabled
	auditingSetting := &sql.ServersAuditingSetting{
		ObjectMeta: tc.MakeObjectMeta("audit"),
		Spec: sql.Servers_AuditingSetting_Spec{
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

	tc.CreateResourceAndWait(auditingSetting)

	tc.Expect(auditingSetting.Status.Id).ToNot(BeNil())

	// Resource doesn't support delete
	// tc.DeleteResourceAndWait(auditingSetting)
}

func SQL_Server_OutboundFirewallRule_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	outboundRule := &sql.ServersOutboundFirewallRule{
		ObjectMeta: tc.MakeObjectMetaWithName("outboundrule"),
		Spec: sql.Servers_OutboundFirewallRule_Spec{
			Owner:     testcommon.AsOwner(server),
			AzureName: "server.database.windows.net",
		},
	}

	tc.CreateResourceAndWait(outboundRule)

	tc.Expect(outboundRule.Status.Id).ToNot(BeNil())
	armId := *outboundRule.Status.Id

	tc.DeleteResourceAndWait(outboundRule)
	exists, _, err := tc.AzureClient.HeadByID(
		tc.Ctx,
		armId,
		string(sql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

// TODO: This test can't be recorded with service principal credentials, which is what we run EnvTest with currently
//func SQL_Server_Administrator_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup) {
//	mi := &managedidentity2018.UserAssignedIdentity{
//		ObjectMeta: tc.MakeObjectMeta("mi"),
//		Spec: managedidentity2018.UserAssignedIdentity_Spec{
//			Location: tc.AzureRegion,
//			Owner:    testcommon.AsOwner(rg),
//			OperatorSpec: &managedidentity2018.UserAssignedIdentityOperatorSpec{
//				ConfigMaps: &managedidentity2018.UserAssignedIdentityOperatorConfigMaps{
//					PrincipalId: &genruntime.ConfigMapDestination{
//						Name: "mimap",
//						Key:  "oid",
//					},
//					TenantId: &genruntime.ConfigMapDestination{
//						Name: "mimap",
//						Key:  "tenant",
//					},
//				},
//			},
//		},
//	}
//
//	tc.CreateResourceAndWait(mi)
//
//	adminType := sql.AdministratorProperties_AdministratorType_ActiveDirectory
//	firewall := &sql.ServersAdministrator{
//		ObjectMeta: tc.MakeObjectMeta("atp"),
//		Spec: sql.Servers_Administrator_Spec{
//			Owner:          testcommon.AsOwner(server),
//			SidFromConfig: &genruntime.ConfigMapReference{
//				Name: "mimap",
//				Key:  "oid",
//			},
//			TenantIdFromConfig: &genruntime.ConfigMapReference{
//				Name: "mimap",
//				Key:  "tenant",
//			},
//			AdministratorType: &adminType,
//			Login:
//		},
//	}
//
//	tc.CreateResourceAndWait(firewall)
//
//	tc.Expect(firewall.Status.Id).ToNot(BeNil())
//	armId := *firewall.Status.Id
//
//	tc.DeleteResourceAndWait(firewall)
//	exists, _, err := tc.AzureClient.HeadByID(
//		tc.Ctx,
//		armId,
//		string(sql.APIVersion_Value))
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(exists).To(BeFalse())
//}

// See https://github.com/Azure/azure-quickstart-templates/blob/master/quickstarts/microsoft.sql/sql-elastic-pool-create/azuredeploy.json
// for SQL Server Elastic Pool examples
func SQL_Server_ElasticPool_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server) {
	elasticPool := &sql.ServersElasticPool{
		ObjectMeta: tc.MakeObjectMeta("pool"),
		Spec: sql.Servers_ElasticPool_Spec{
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
func SQL_Server_VirtualNetworkRule_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, rg *resources.ResourceGroup) {
	subnet := makeSubnetForSQLServer(tc, rg)

	vnetRule := &sql.ServersVirtualNetworkRule{
		ObjectMeta: tc.MakeObjectMeta("pool"),
		Spec: sql.Servers_VirtualNetworkRule_Spec{
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
func SQL_Databases_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, storageDetails vulnStorageAccountDetails) {
	db := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.Servers_Database_Spec{
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
				SQL_BackupLongTermRetention_CRUD(tc, db)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Vulnerability Assessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_VulnerabilityAssessment_CRUD(tc, db, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Auditing Setting CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_AuditingSetting_CRUD(tc, db, storageDetails)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Transparent Data Encryption Assessment CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_TransparentDataEncryption_CRUD(tc, db)
			},
		},
		testcommon.Subtest{
			Name: "SQL Database Advanced Threat Protection CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				SQL_Database_AdvancedThreatProtection_CRUD(tc, db)
			},
		},
		// TODO: Not actually supported in ARM at the moment
		//testcommon.Subtest{
		//	Name: "SQL Database GeoBackupPolicy CRUD",
		//	Test: func(tc *testcommon.KubePerTestContext) {
		//		SQL_Database_GeoBackupPolicy_CRUD(tc, db)
		//	},
		//}
	)

	tc.DeleteResourceAndWait(db)
	tc.ExpectResourceIsDeletedInAzure(armId, string(sql.APIVersion_Value))
}

func SQL_BackupLongTermRetention_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
	policy := &sql.ServersDatabasesBackupLongTermRetentionPolicy{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.Servers_Databases_BackupLongTermRetentionPolicy_Spec{
			Owner:           testcommon.AsOwner(db),
			WeeklyRetention: to.Ptr("P30D"),
		},
	}

	tc.CreateResourceAndWait(policy)
	tc.Expect(policy.Status.Id).ToNot(BeNil())
}

func SQL_Database_VulnerabilityAssessment_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase, storageDetails vulnStorageAccountDetails) {
	enabled := sql.DatabaseSecurityAlertPoliciesSecurityAlertsPolicyProperties_State_Enabled
	securityAlertPolicy := &sql.ServersDatabasesSecurityAlertPolicy{
		ObjectMeta: tc.MakeObjectMeta("alertpolicy"),
		Spec: sql.Servers_Databases_SecurityAlertPolicy_Spec{
			Owner: testcommon.AsOwner(db),
			State: &enabled,
		},
	}

	tc.CreateResourceAndWait(securityAlertPolicy)

	tc.Expect(securityAlertPolicy.Status.Id).ToNot(BeNil())

	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	vulnerabilityAssessment := &sql.ServersDatabasesVulnerabilityAssessment{
		ObjectMeta: tc.MakeObjectMeta("vulnassessment"),
		Spec: sql.Servers_Databases_VulnerabilityAssessment_Spec{
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

	tc.CreateResourceAndWait(vulnerabilityAssessment)

	tc.Expect(vulnerabilityAssessment.Status.Id).ToNot(BeNil())

	tc.DeleteResourceAndWait(vulnerabilityAssessment)
	// TODO: It seems like delete of this resource isn't actually honored by the service - it's accepted but doesn't remove the resource
	//exists, _, err := tc.AzureClient.HeadByID(
	//	tc.Ctx,
	//	armId,
	//	string(sql.APIVersion_Value))
	//tc.Expect(err).ToNot(HaveOccurred())
	//tc.Expect(exists).To(BeFalse())
}

func SQL_Database_AuditingSetting_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase, storageDetails vulnStorageAccountDetails) {
	secret := tc.GetSecret(storageDetails.secretName)
	blobEndpoint := string(secret.Data[storageDetails.blobEndpointSecretKey])

	enabled := sql.DatabaseBlobAuditingPolicyProperties_State_Enabled
	auditingSetting := &sql.ServersDatabasesAuditingSetting{
		ObjectMeta: tc.MakeObjectMeta("audit"),
		Spec: sql.Servers_Databases_AuditingSetting_Spec{
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

	tc.CreateResourceAndWait(auditingSetting)

	tc.Expect(auditingSetting.Status.Id).ToNot(BeNil())

	// Resource doesn't support delete
	// tc.DeleteResourceAndWait(auditingSetting)
}

func SQL_Database_TransparentDataEncryption_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
	enabled := sql.TransparentDataEncryptionProperties_State_Enabled
	transparentDataEncryption := &sql.ServersDatabasesTransparentDataEncryption{
		ObjectMeta: tc.MakeObjectMeta("encrypt"),
		Spec: sql.Servers_Databases_TransparentDataEncryption_Spec{
			Owner: testcommon.AsOwner(db),
			State: &enabled,
		},
	}

	tc.CreateResourceAndWait(transparentDataEncryption)

	tc.Expect(transparentDataEncryption.Status.Id).ToNot(BeNil())

	// Delete is not supported for this resource
	// tc.DeleteResourceAndWait(transparentDataEncryption)
}

func SQL_Database_AdvancedThreatProtection_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
	enabled := sql.AdvancedThreatProtectionProperties_State_Enabled
	advancedProtection := &sql.ServersDatabasesAdvancedThreatProtectionSetting{
		ObjectMeta: tc.MakeObjectMeta("atp"),
		Spec: sql.Servers_Databases_AdvancedThreatProtectionSetting_Spec{
			Owner: testcommon.AsOwner(db),
			State: &enabled,
		},
	}

	tc.CreateResourceAndWait(advancedProtection)

	tc.Expect(advancedProtection.Status.Id).ToNot(BeNil())

	// TODO: Delete is not allowed for this resource
	// tc.DeleteResourceAndWait(policy)
}

// Not actually supported in ARM at the moment
//func SQL_Database_GeoBackupPolicy_CRUD(tc *testcommon.KubePerTestContext, db *sql.ServersDatabase) {
//	enabled := sql.GeoBackupPolicyProperties_State_Enabled
//	geoBackupPolicy := &sql.ServersDatabasesGeoBackupPolicy{
//		ObjectMeta: tc.MakeObjectMeta("atp"),
//		Spec: sql.Servers_Databases_GeoBackupPolicy_Spec{
//			Owner: testcommon.AsOwner(db),
//			State: &enabled,
//		},
//	}
//
//	tc.CreateResourceAndWait(geoBackupPolicy)
//
//	tc.Expect(geoBackupPolicy.Status.Id).ToNot(BeNil())
//
//	// TODO: Delete is not allowed for this resource
//	// tc.DeleteResourceAndWait(policy)
//}

type vulnStorageAccountDetails struct {
	container             string
	secretName            string
	blobEndpointSecretKey string
	keySecretKey          string
}

func makeStorageAccountForSQLVulnerabilityAssessment(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) vulnStorageAccountDetails {
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
		Spec: storage.StorageAccounts_BlobService_Spec{
			Owner: testcommon.AsOwner(acct),
		},
	}
	blobContainer := &storage.StorageAccountsBlobServicesContainer{
		ObjectMeta: tc.MakeObjectMetaWithName(vulnerabilityAssessmentsContainerName),
		Spec: storage.StorageAccounts_BlobServices_Container_Spec{
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

func makeSubnetForSQLServer(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *network.VirtualNetworksSubnet {
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
		Spec: network.VirtualNetworks_Subnet_Spec{
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
