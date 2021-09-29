/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	. "github.com/onsi/gomega"
	"sigs.k8s.io/controller-runtime/pkg/client"

	postgresql "github.com/Azure/azure-service-operator/v2/api/microsoft.dbforpostgresql/v1alpha1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/controller/testcommon"
)

func Test_DBForPostgreSQL_FlexibleServer_CRUD(t *testing.T) {
	t.Parallel()

	g := NewGomegaWithT(t)
	ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	version := postgresql.ServerPropertiesVersion13
	flexibleServer := &postgresql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("postgresql"),
		Spec: postgresql.FlexibleServers_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg.ObjectMeta),
			Version:  &version,
			Sku: &postgresql.Sku{
				Name: "Standard_D4s_v3",
				Tier: postgresql.SkuTierGeneralPurpose,
			},
			AdministratorLogin:         to.StringPtr("myadmin"),
			AdministratorLoginPassword: to.StringPtr(tc.Namer.GeneratePassword()),
			Storage: &postgresql.Storage{
				StorageSizeGB: to.IntPtr(128),
			},
		},
	}

	tc.CreateResourceAndWait(flexibleServer)

	// It should be created in Kubernetes
	g.Expect(flexibleServer.Status.Id).ToNot(BeNil())
	armId := *flexibleServer.Status.Id

	// Perform a simple patch
	old := flexibleServer.DeepCopy()
	flexibleServer.Spec.MaintenanceWindow = &postgresql.MaintenanceWindow{
		CustomWindow: to.StringPtr("enabled"),
		DayOfWeek:    to.IntPtr(5),
	}
	tc.Patch(old, flexibleServer)

	objectKey := client.ObjectKeyFromObject(flexibleServer)

	// ensure state got updated in Azure
	tc.Eventually(func() *int {
		updatedServer := &postgresql.FlexibleServer{}
		tc.GetResource(objectKey, updatedServer)
		if updatedServer.Status.MaintenanceWindow == nil {
			return nil
		}

		return updatedServer.Status.MaintenanceWindow.DayOfWeek
	}).Should(Equal(to.IntPtr(5)))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Flexible servers database CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				FlexibleServer_Database_CRUD(testContext, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers firewall CRUD",
			Test: func(testContext testcommon.KubePerTestContext) {
				FlexibleServer_FirewallRule_CRUD(testContext, flexibleServer)
			},
		},
	)

	tc.DeleteResourceAndWait(flexibleServer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadResource(ctx, armId, string(postgresql.FlexibleServersDatabasesSpecAPIVersion20210601))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}

func FlexibleServer_Database_CRUD(tc testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	database := &postgresql.FlexibleServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: postgresql.FlexibleServersDatabases_Spec{
			Owner:   testcommon.AsOwner(flexibleServer.ObjectMeta),
			Charset: to.StringPtr("utf8"),
		},
	}
	tc.CreateResourceAndWait(database)
	defer tc.DeleteResourceAndWait(database)

	tc.Expect(database.Status.Id).ToNot(BeNil())
}

func FlexibleServer_FirewallRule_CRUD(tc testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	firewall := &postgresql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("fwrule"),
		Spec: postgresql.FlexibleServersFirewallRules_Spec{
			Owner: testcommon.AsOwner(flexibleServer.ObjectMeta),
			// I think that these rules are allow rules - somebody with this IP can access the server.
			StartIpAddress: "1.2.3.4",
			EndIpAddress:   "1.2.3.4",
		},
	}

	tc.CreateResourceAndWait(firewall)
	defer tc.DeleteResourceAndWait(firewall)

	tc.Expect(firewall.Status.Id).ToNot(BeNil())
}
