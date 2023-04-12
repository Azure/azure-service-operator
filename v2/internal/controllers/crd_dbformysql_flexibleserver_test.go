/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/kr/pretty"
	. "github.com/onsi/gomega"

	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DBForMySQL_FlexibleServer_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()
	secretName := "mysqlsecret"
	adminPasswordKey := "adminPassword"
	adminPasswordSecretRef := createPasswordSecret(secretName, adminPasswordKey, tc)

	flexibleServer, fqdnSecret := newFlexibleServer(tc, rg, adminPasswordSecretRef)

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

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "MySQL Flexible servers database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_Database_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "MySQL Flexible servers firewall CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_FirewallRule_CRUD(tc, flexibleServer)
			},
		},
	)

	tc.DeleteResourceAndWait(flexibleServer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(mysql.APIVersion_Value))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func newFlexibleServer(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, adminPasswordSecretRef genruntime.SecretReference) (*mysql.FlexibleServer, string) {
	//location := tc.AzureRegion Capacity crunch in West US 2 makes this not work when live
	location := "eastus"
	version := mysql.ServerVersion_8021
	tier := mysql.Sku_Tier_GeneralPurpose
	fqdnSecret := "fqdnsecret"
	flexibleServer := &mysql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("mysql"),
		Spec: mysql.FlexibleServer_Spec{
			Location: &location,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &mysql.Sku{
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

func MySQLFlexibleServer_Database_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) {
	// The RP doesn't like databases with hyphens in the name,
	// although it doesn't give nice errors to point this out
	database := &mysql.FlexibleServersDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("db")),
		Spec: mysql.FlexibleServers_Database_Spec{
			Owner:   testcommon.AsOwner(flexibleServer),
			Charset: to.Ptr("utf8mb4"),
		},
	}
	tc.CreateResourceAndWait(database)
	defer tc.DeleteResourceAndWait(database)

	tc.Expect(database.Status.Id).ToNot(BeNil())
}

func MySQLFlexibleServer_FirewallRule_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) {
	rule := &mysql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("fwrule"),
		Spec: mysql.FlexibleServers_FirewallRule_Spec{
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
