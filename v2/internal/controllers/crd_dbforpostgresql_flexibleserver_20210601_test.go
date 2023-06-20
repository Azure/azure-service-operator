/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package controllers_test

import (
	"context"
	"testing"

	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_DBForPostgreSQL_FlexibleServer_20210601_CRUD(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("can't run in live mode, postresql flexible server takes too long to be provisioned and deletion")
	}

	g := NewGomegaWithT(t)
	ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	//location := tc.AzureRegion Capacity crunch in West US 2 makes this not work when live
	location := "eastus"

	rg := tc.CreateTestResourceGroupAndWait()

	adminPasswordKey := "adminPassword"
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("postgresqlsecret"),
		StringData: map[string]string{
			adminPasswordKey: tc.Namer.GeneratePassword(),
		},
	}

	tc.CreateResource(secret)

	secretRef := genruntime.SecretReference{
		Name: secret.Name,
		Key:  adminPasswordKey,
	}
	version := postgresql.ServerVersion_13
	tier := postgresql.Sku_Tier_GeneralPurpose
	fqdnSecret := "fqdnsecret"
	flexibleServer := &postgresql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("postgresql"),
		Spec: postgresql.FlexibleServer_Spec{
			Location: &location,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &postgresql.Sku{
				Name: to.Ptr("Standard_D4s_v3"),
				Tier: &tier,
			},
			AdministratorLogin:         to.Ptr("myadmin"),
			AdministratorLoginPassword: &secretRef,
			Storage: &postgresql.Storage{
				StorageSizeGB: to.Ptr(128),
			},
			OperatorSpec: &postgresql.FlexibleServerOperatorSpec{
				Secrets: &postgresql.FlexibleServerOperatorSecrets{
					FullyQualifiedDomainName: &genruntime.SecretDestination{Name: fqdnSecret, Key: "fqdn"},
				},
			},
		},
	}

	// TODO: Create the secret here instead?
	tc.CreateResourceAndWait(flexibleServer)

	// It should be created in Kubernetes
	g.Expect(flexibleServer.Status.Id).ToNot(BeNil())
	armId := *flexibleServer.Status.Id

	// It should have the expected secret data written
	tc.ExpectSecretHasKeys(fqdnSecret, "fqdn")

	// Perform a simple patch
	old := flexibleServer.DeepCopy()
	flexibleServer.Spec.MaintenanceWindow = &postgresql.MaintenanceWindow{
		CustomWindow: to.Ptr("enabled"),
		DayOfWeek:    to.Ptr(5),
	}
	tc.PatchResourceAndWait(old, flexibleServer)
	tc.Expect(flexibleServer.Status.MaintenanceWindow).ToNot(BeNil())
	tc.Expect(flexibleServer.Status.MaintenanceWindow.DayOfWeek).To(Equal(to.Ptr(5)))

	tc.RunParallelSubtests(
		testcommon.Subtest{
			Name: "Flexible servers database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_Database_20210601_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers firewall CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_FirewallRule_20210601_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers configuration CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_Configuration_20210601_CRUD(tc, flexibleServer)
			},
		},
	)

	tc.DeleteResourceAndWait(flexibleServer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.HeadByID(ctx, armId, string(postgresql.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}

func FlexibleServer_Database_20210601_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	database := &postgresql.FlexibleServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: postgresql.FlexibleServers_Database_Spec{
			Owner:   testcommon.AsOwner(flexibleServer),
			Charset: to.Ptr("utf8"),
		},
	}
	tc.CreateResourceAndWait(database)
	defer tc.DeleteResourceAndWait(database)

	tc.Expect(database.Status.Id).ToNot(BeNil())
}

func FlexibleServer_FirewallRule_20210601_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	firewall := &postgresql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("fwrule"),
		Spec: postgresql.FlexibleServers_FirewallRule_Spec{
			Owner: testcommon.AsOwner(flexibleServer),
			// I think that these rules are allow rules - somebody with this IP can access the server.
			StartIpAddress: to.Ptr("1.2.3.4"),
			EndIpAddress:   to.Ptr("1.2.3.4"),
		},
	}

	tc.CreateResourceAndWait(firewall)
	defer tc.DeleteResourceAndWait(firewall)

	tc.Expect(firewall.Status.Id).ToNot(BeNil())
}

func FlexibleServer_Configuration_20210601_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	configuration := &postgresql.FlexibleServersConfiguration{
		ObjectMeta: tc.MakeObjectMeta("pgaudit"),
		Spec: postgresql.FlexibleServers_Configuration_Spec{
			Owner:     testcommon.AsOwner(flexibleServer),
			AzureName: "pgaudit.log",
			Source:    to.Ptr("user-override"),
			Value:     to.Ptr("READ"),
		},
	}

	tc.CreateResourceAndWait(configuration)
	// This isn't a "real" resource so it cannot be deleted directly
	// defer tc.DeleteResourceAndWait(configuration)

	tc.Expect(configuration.Status.Id).ToNot(BeNil())
	tc.Expect(configuration.Status.Value).To(Equal(to.Ptr("READ")))
}
