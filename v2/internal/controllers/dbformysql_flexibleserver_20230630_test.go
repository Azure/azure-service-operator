/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/kr/pretty"

	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20230630"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DBForMySQL_FlexibleServer_20230630_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	tc.AzureRegion = to.Ptr("eastus")

	rg := tc.CreateTestResourceGroupAndWait()
	secretName := "mysqlsecret"
	adminPasswordKey := "adminPassword"
	// Hack here to maintain the consistency of the seed for name generation.
	// TODO: We need to remove this redundant call to `GenerateNameOfLength` and re-record the test
	_ = tc.Namer.GenerateNameOfLength(40)
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	flexibleServer, fqdnSecret := newFlexibleServer20230630(tc, rg, adminPasswordSecretRef)

	tc.CreateResourceAndWait(flexibleServer)

	// It should be created in Kubernetes
	tc.Expect(flexibleServer.Status.Id).ToNot(BeNil())
	armId := *flexibleServer.Status.Id

	// It should have the expected secret data written
	tc.ExpectSecretHasKeys(fqdnSecret, "fqdn")

	// Perform a simple patch
	old := flexibleServer.DeepCopy()
	disabled := mysql.EnableStatusEnum_Disabled
	flexibleServer.Spec.Backup = &mysql.Backup{
		BackupRetentionDays: to.Ptr(5),
		GeoRedundantBackup:  &disabled,
	}
	tc.PatchResourceAndWait(old, flexibleServer)
	tc.Expect(flexibleServer.Status.Backup).ToNot(BeNil())
	tc.T.Log(pretty.Sprint(flexibleServer.Status.Backup))
	tc.Expect(flexibleServer.Status.Backup.BackupRetentionDays).To(Equal(to.Ptr(5)))

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "MySQL Flexible servers AAD Administrators CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_AADAdmin_20230630_CRUD(tc, rg, flexibleServer)
			},
		},
	)

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "MySQL Flexible servers configuration CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_Configuration_20230630_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "MySQL Flexible servers database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_Database_20230630_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "MySQL Flexible servers firewall CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_FirewallRule_20230630_CRUD(tc, flexibleServer)
			},
		},
	)

	tc.DeleteResourceAndWait(flexibleServer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(tc.Ctx, armId, string(mysql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newFlexibleServer20230630(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, adminPasswordSecretRef genruntime.SecretReference) (*mysql.FlexibleServer, string) {
	version := mysql.ServerVersion_8021
	tier := mysql.MySQLServerSku_Tier_GeneralPurpose
	fqdnSecret := "fqdnsecret"
	flexibleServer := &mysql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("mysql"),
		Spec: mysql.FlexibleServer_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &mysql.MySQLServerSku{
				Name: to.Ptr("Standard_D4ds_v4"),
				Tier: &tier,
			},
			AdministratorLogin:         to.Ptr("myadmin"),
			AdministratorLoginPassword: &adminPasswordSecretRef,
			Storage: &mysql.Storage{
				StorageSizeGB: to.Ptr(128),
			},
			OperatorSpec: &mysql.FlexibleServerOperatorSpec{
				Secrets: &mysql.FlexibleServerOperatorSecrets{
					FullyQualifiedDomainName: &genruntime.SecretDestination{Name: fqdnSecret, Key: "fqdn"},
				},
			},
		},
	}

	return flexibleServer, fqdnSecret
}

func MySQLFlexibleServer_Database_20230630_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) {
	// The RP doesn't like databases with hyphens in the name,
	// although it doesn't give nice errors to point this out
	database := &mysql.FlexibleServersDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("db")),
		Spec: mysql.FlexibleServersDatabase_Spec{
			Owner:   testcommon.AsOwner(flexibleServer),
			Charset: to.Ptr("utf8mb4"),
		},
	}
	tc.CreateResourceAndWait(database)
	defer tc.DeleteResourceAndWait(database)

	tc.Expect(database.Status.Id).ToNot(BeNil())
}

func MySQLFlexibleServer_FirewallRule_20230630_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) {
	rule := &mysql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("fwrule"),
		Spec: mysql.FlexibleServersFirewallRule_Spec{
			Owner:          testcommon.AsOwner(flexibleServer),
			StartIpAddress: to.Ptr("1.2.3.4"),
			EndIpAddress:   to.Ptr("1.2.3.4"),
		},
	}

	tc.CreateResourceAndWait(rule)
	defer tc.DeleteResourceAndWait(rule)

	old := rule.DeepCopy()
	rule.Spec.EndIpAddress = to.Ptr("1.2.3.5")
	tc.PatchResourceAndWait(old, rule)
	tc.Expect(rule.Status.EndIpAddress).To(Equal(to.Ptr("1.2.3.5")))

	// The GET responses are coming back from the RP with no ARM ID -
	// this seems invalid per the ARM spec.
	// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md#get-resource
	// tc.Expect(rule.Status.Id).ToNot(BeNil())
}

func MySQLFlexibleServer_AADAdmin_20230630_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, server *mysql.FlexibleServer) {
	// Create a managed identity to serve as aad admin
	configMapName := "my-configmap"
	clientIDKey := "clientId"
	tenantIDKey := "tenantId"

	// Create a managed identity to use as the AAD administrator
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMaps: &managedidentity.UserAssignedIdentityOperatorConfigMaps{
					ClientId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  clientIDKey,
					},
					TenantId: &genruntime.ConfigMapDestination{
						Name: configMapName,
						Key:  tenantIDKey,
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(mi)

	// Update the server to use the managed identity as its UMI
	old := server.DeepCopy()
	server.Spec.Identity = &mysql.MySQLServerIdentity{
		Type: to.Ptr(mysql.MySQLServerIdentity_Type_UserAssigned),
		UserAssignedIdentities: []mysql.UserAssignedIdentityDetails{
			{
				Reference: *tc.MakeReferenceFromResource(mi),
			},
		},
	}

	tc.PatchResourceAndWait(old, server)

	aadAdmin := mysql.AdministratorProperties_AdministratorType_ActiveDirectory
	admin := &mysql.FlexibleServersAdministrator{
		ObjectMeta: tc.MakeObjectMeta("aadadmin"),
		Spec: mysql.FlexibleServersAdministrator_Spec{
			Owner:             testcommon.AsOwner(server),
			AdministratorType: &aadAdmin,
			Login:             &mi.Name,
			TenantIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  tenantIDKey,
			},
			SidFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  clientIDKey,
			},
			IdentityResourceReference: tc.MakeReferenceFromResource(mi),
		},
	}

	tc.CreateResourceAndWait(admin)
	defer tc.DeleteResourceAndWait(admin)

	tc.Expect(admin.Status.Id).ToNot(BeNil())
}

func MySQLFlexibleServer_Configuration_20230630_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) {
	configuration := &mysql.FlexibleServersConfiguration{
		ObjectMeta: tc.MakeObjectMetaWithName("maxconnections"),
		Spec: mysql.FlexibleServersConfiguration_Spec{
			AzureName: "max_connections",
			Owner:     testcommon.AsOwner(flexibleServer),
			Source:    to.Ptr(mysql.ConfigurationProperties_Source_UserOverride),
			Value:     to.Ptr("20"),
		},
	}
	tc.CreateResourceAndWait(configuration)
	tc.Expect(configuration.Status.Id).ToNot(BeNil())
}
