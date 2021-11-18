/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/kr/pretty"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	mysql "github.com/Azure/azure-service-operator/v2/api/microsoft.dbformysql/v1alpha1api20210501"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
)

func Test_DBForMySQL_FlexibleServer_CRUD(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	version := mysql.ServerPropertiesVersion8021
	flexibleServer := &mysql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("mysql"),
		Spec: mysql.FlexibleServers_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &mysql.Sku{
				Name: "Standard_D4ds_v4",
				Tier: mysql.SkuTierGeneralPurpose,
			},
			AdministratorLogin:         to.StringPtr("myadmin"),
			AdministratorLoginPassword: to.StringPtr(tc.Namer.GeneratePassword()),
			Storage: &mysql.Storage{
				StorageSizeGB: to.IntPtr(128),
			},
		},
	}

	tc.CreateResourceAndWait(flexibleServer)

	// It should be created in Kubernetes
	tc.Expect(flexibleServer.Status.Id).ToNot(BeNil())
	armId := *flexibleServer.Status.Id

	// Perform a simple patch
	old := flexibleServer.DeepCopy()
	disabled := mysql.BackupGeoRedundantBackupDisabled
	flexibleServer.Spec.Backup = &mysql.Backup{
		BackupRetentionDays: to.IntPtr(5),
		GeoRedundantBackup:  &disabled,
	}
	tc.Patch(old, flexibleServer)

	objectKey := client.ObjectKeyFromObject(flexibleServer)

	// ensure state got updated in Azure
	tc.Eventually(func() *int {
		var updatedServer mysql.FlexibleServer
		tc.GetResource(objectKey, &updatedServer)
		if updatedServer.Status.Backup == nil {
			return nil
		}

		tc.T.Log(pretty.Sprint(updatedServer.Status.Backup))
		return updatedServer.Status.Backup.BackupRetentionDays
	}).Should(Equal(to.IntPtr(5)))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "MySQL Flexible servers database CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_Database_CRUD(testContext, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "MySQL Flexible servers firewall CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQLFlexibleServer_FirewallRule_CRUD(testContext, flexibleServer)
			},
		},
	)

	tc.DeleteResourceAndWait(flexibleServer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(tc.Ctx, armId, string(mysql.FlexibleServersDatabasesSpecAPIVersion20210501))
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(retryAfter).To(BeZero())
	tc.Expect(exists).To(BeFalse())
}

func MySQLFlexibleServer_Database_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) {
	// The RP doesn't like databases with hyphens in the name,
	// although it doesn't give nice errors to point this out.
	namer := tc.Namer.WithSeparator("")
	database := &mysql.FlexibleServersDatabase{
		ObjectMeta: tc.MakeObjectMetaWithName(namer.GenerateName("db")),
		Spec: mysql.FlexibleServersDatabases_Spec{
			Owner:   testcommon.AsOwner(flexibleServer),
			Charset: to.StringPtr("utf8mb4"),
		},
	}
	tc.CreateResourceAndWait(database)
	defer tc.DeleteResourceAndWait(database)

	tc.Expect(database.Status.Id).ToNot(BeNil())
}

func MySQLFlexibleServer_FirewallRule_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) {
	rule := &mysql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("fwrule"),
		Spec: mysql.FlexibleServersFirewallRules_Spec{
			Owner:          testcommon.AsOwner(flexibleServer),
			StartIpAddress: "1.2.3.4",
			EndIpAddress:   "1.2.3.4",
		},
	}

	tc.CreateResourceAndWait(rule)
	defer tc.DeleteResourceAndWait(rule)

	old := rule.DeepCopy()
	rule.Spec.EndIpAddress = "1.2.3.5"
	tc.Patch(old, rule)

	objectKey := client.ObjectKeyFromObject(rule)
	tc.Eventually(func() *string {
		var updatedRule mysql.FlexibleServersFirewallRule
		tc.GetResource(objectKey, &updatedRule)
		return updatedRule.Status.EndIpAddress
	}).Should(Equal(to.StringPtr("1.2.3.5")))

	// The GET responses are coming back from the RP with no ARM ID -
	// this seems invalid per the ARM spec.
	// https://github.com/Azure/azure-resource-manager-rpc/blob/master/v1.0/resource-api-reference.md#get-resource
	// tc.Expect(rule.Status.Id).ToNot(BeNil())
}
