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
	"sigs.k8s.io/controller-runtime/pkg/client"

	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20250801"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

// Name slightly different here because running into name collision issues
func Test_DBForPostgreSQL_FlexibleServer_20250801_CRUD_2(t *testing.T) {
	t.Parallel()

	if *isLive {
		t.Skip("can't run in live mode, postresql flexible server takes too long to be provisioned and deleted")
	}

	g := NewGomegaWithT(t)
	ctx := context.Background()
	tc := globalTestContext.ForTest(t)

	// Capacity crunch in West US 2 makes this not work when live
	tc.AzureRegion = to.Ptr("uksouth")

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
	version := postgresql.ServerVersion_17
	tier := postgresql.Sku_Tier_GeneralPurpose
	fqdnConfig := "fqdnconfig"
	flexibleServer := &postgresql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("postgresql"),
		Spec: postgresql.FlexibleServer_Spec{
			Location: tc.AzureRegion,
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
			AuthConfig: &postgresql.AuthConfig{
				ActiveDirectoryAuth: to.Ptr(postgresql.AuthConfig_ActiveDirectoryAuth_Enabled),
				PasswordAuth:        to.Ptr(postgresql.AuthConfig_PasswordAuth_Enabled),
				TenantId:            &tc.AzureTenant,
			},
			OperatorSpec: &postgresql.FlexibleServerOperatorSpec{
				ConfigMaps: &postgresql.FlexibleServerOperatorConfigMaps{
					FullyQualifiedDomainName: &genruntime.ConfigMapDestination{
						Name: fqdnConfig,
						Key:  "fqdn",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(flexibleServer)

	// It should be created in Kubernetes
	g.Expect(flexibleServer.Status.Id).ToNot(BeNil())
	g.Expect(flexibleServer.Status.FullyQualifiedDomainName).ToNot(BeNil())
	armId := *flexibleServer.Status.Id

	// It should have the expected config data written
	tc.ExpectConfigMapHasKeysAndValues(fqdnConfig, "fqdn", *flexibleServer.Status.FullyQualifiedDomainName)

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
			Name: "ConfigMapValuesWrittenToSameConfigMap",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_20250801_ConfigValuesWrittenToSameConfigMap(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers database CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_Database_20250801_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers firewall CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_FirewallRule_20250801_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers configuration CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_Configuration_20250801_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers backup CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_Backup_20250801_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers advanced threat protection CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_AdvancedThreatProtection_20250801_CRUD(tc, flexibleServer)
			},
		},
		testcommon.Subtest{
			Name: "Flexible servers administrator CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_Administrator_20250801_CRUD(tc, rg, flexibleServer)
			},
		},
	)

	// For some reason PostgreSQL didn't like these happening in parallel, so doing them in sequence at the end
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "Flexible servers virtual endpoint CRUD",
			Test: func(tc *testcommon.KubePerTestContext) {
				FlexibleServer_VirtualEndpoint_20250801_CRUD(tc, flexibleServer)
			},
		},
	)

	tc.DeleteResourceAndWait(flexibleServer)

	// Ensure that the resource was really deleted in Azure
	exists, retryAfter, err := tc.AzureClient.CheckExistenceWithGetByID(ctx, armId, string(postgresql.APIVersion_Value))
	g.Expect(err).ToNot(HaveOccurred())
	g.Expect(retryAfter).To(BeZero())
	g.Expect(exists).To(BeFalse())
}

func FlexibleServer_20250801_ConfigValuesWrittenToSameConfigMap(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	old := flexibleServer.DeepCopy()
	flexibleServerConfigMap := "serverconfig"
	flexibleServerConfigMapKey := "fqdn"

	flexibleServer.Spec.OperatorSpec = &postgresql.FlexibleServerOperatorSpec{
		ConfigMaps: &postgresql.FlexibleServerOperatorConfigMaps{
			FullyQualifiedDomainName: &genruntime.ConfigMapDestination{
				Name: flexibleServerConfigMap,
				Key:  flexibleServerConfigMapKey,
			},
		},
	}

	tc.PatchResourceAndWait(old, flexibleServer)

	// There should be at least some config maps at this point
	list := &v1.ConfigMapList{}
	tc.ListResources(list, client.InNamespace(tc.Namespace))
	tc.Expect(list.Items).NotTo(HaveLen(0))

	tc.ExpectConfigMapHasKeysAndValues(
		flexibleServerConfigMap,
		flexibleServerConfigMapKey,
		*flexibleServer.Status.FullyQualifiedDomainName)
}

func FlexibleServer_Database_20250801_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	database := &postgresql.FlexibleServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: postgresql.FlexibleServersDatabase_Spec{
			Owner:   testcommon.AsOwner(flexibleServer),
			Charset: to.Ptr("utf8"),
		},
	}
	tc.CreateResourceAndWait(database)
	defer tc.DeleteResourceAndWait(database)

	tc.Expect(database.Status.Id).ToNot(BeNil())
}

func FlexibleServer_FirewallRule_20250801_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	firewall := &postgresql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("fwrule"),
		Spec: postgresql.FlexibleServersFirewallRule_Spec{
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

func FlexibleServer_Configuration_20250801_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	configuration := &postgresql.FlexibleServersConfiguration{
		ObjectMeta: tc.MakeObjectMeta("pgaudit"),
		Spec: postgresql.FlexibleServersConfiguration_Spec{
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

func FlexibleServer_VirtualEndpoint_20250801_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	virtualEndpoint := &postgresql.FlexibleServersVirtualEndpoint{
		ObjectMeta: tc.MakeObjectMeta("virtualendpoint"),
		Spec: postgresql.FlexibleServersVirtualEndpoint_Spec{
			Owner:        testcommon.AsOwner(flexibleServer),
			EndpointType: to.Ptr(postgresql.VirtualEndpointResourceProperties_EndpointType_ReadWrite),
			Members: []string{
				flexibleServer.Spec.AzureName,
			},
		},
	}

	tc.CreateResourceAndWait(virtualEndpoint)
	defer tc.DeleteResourceAndWait(virtualEndpoint)

	tc.Expect(virtualEndpoint.Status.Id).ToNot(BeNil())
}

func FlexibleServer_AdvancedThreatProtection_20250801_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	threatProtection := &postgresql.FlexibleServersAdvancedThreatProtectionSettings{
		ObjectMeta: tc.MakeObjectMeta("advthreat"),
		Spec: postgresql.FlexibleServersAdvancedThreatProtectionSettings_Spec{
			Owner: testcommon.AsOwner(flexibleServer),
			State: to.Ptr(postgresql.AdvancedThreatProtectionSettingsProperties_State_Enabled),
		},
	}

	tc.CreateResourceAndWait(threatProtection)
	// This isn't a "real" resource so it cannot be deleted directly
	// defer tc.DeleteResourceAndWait(threatProtection)

	tc.Expect(threatProtection.Status.Id).ToNot(BeNil())
}

func FlexibleServer_Backup_20250801_CRUD(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) {
	backup := &postgresql.FlexibleServersBackup{
		ObjectMeta: tc.MakeObjectMeta("backup"),
		Spec: postgresql.FlexibleServersBackup_Spec{
			Owner: testcommon.AsOwner(flexibleServer),
		},
	}

	tc.CreateResourceAndWait(backup)
	defer tc.DeleteResourceAndWait(backup)

	tc.Expect(backup.Status.Id).ToNot(BeNil())
}

// TODO: https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/security-entra-concepts
// TODO: https://github.com/Azure/azure-quickstart-templates/tree/master/quickstarts/microsoft.dbforpostgresql/flexible-postgresql-with-aad
// TODO: https://learn.microsoft.com/en-us/azure/postgresql/flexible-server/security-entra-configure
func FlexibleServer_Administrator_20250801_CRUD(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, flexibleServer *postgresql.FlexibleServer) {
	// Create a managed identity to serve as aad admin
	configMapName := "my-configmap"
	// clientIDKey := "clientId"
	tenantIDKey := "tenantId"

	// Create a managed identity to use as the AAD administrator
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			OperatorSpec: &managedidentity.UserAssignedIdentityOperatorSpec{
				ConfigMapExpressions: []*core.DestinationExpression{
					// {
					// 	Name:  configMapName,
					// 	Key:   clientIDKey,
					// 	Value: "self.status.clientId",
					// },
					{
						Name:  configMapName,
						Key:   tenantIDKey,
						Value: "self.status.tenantId",
					},
				},
			},
		},
	}

	tc.CreateResourceAndWait(mi)

	administrator := &postgresql.FlexibleServersAdministrator{
		ObjectMeta: tc.MakeObjectMeta("admin"),
		Spec: postgresql.FlexibleServersAdministrator_Spec{
			Owner:         testcommon.AsOwner(flexibleServer),
			AzureName:     to.Value(mi.Status.PrincipalId),
			PrincipalName: mi.Status.Name,
			TenantIdFromConfig: &genruntime.ConfigMapReference{
				Name: configMapName,
				Key:  tenantIDKey,
			},
			PrincipalType: to.Ptr(postgresql.AdministratorMicrosoftEntraPropertiesForAdd_PrincipalType_ServicePrincipal),
		},
	}

	tc.CreateResourceAndWait(administrator)
	defer tc.DeleteResourceAndWait(administrator)

	tc.Expect(administrator.Status.Id).ToNot(BeNil())
}
