/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" //the pgx lib
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	postgresqlv1 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1"
	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20210601"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	postgresqlutil "github.com/Azure/azure-service-operator/v2/internal/util/postgresql"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_PostgreSql_Combined(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePassword()
	secret := newPostgresSqlSecret(tc, adminPasswordKey, adminPassword)

	tc.CreateResource(secret)

	flexibleServer := newPostgreSqlServer(tc, rg, adminUsername, adminPasswordKey, secret.Name)
	tc.CreateResourceAndWait(flexibleServer)

	firewallRule := newPostgreSqlServerOpenFirewallRule(tc, flexibleServer)
	tc.CreateResourceAndWait(firewallRule)

	tc.Expect(flexibleServer.Status.FullyQualifiedDomainName).ToNot(BeNil())
	fqdn := *flexibleServer.Status.FullyQualifiedDomainName

	// These must run sequentially as they're mutating SQL state
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "PostgreSql User Helpers",
			Test: func(testContext *testcommon.KubePerTestContext) {
				PostgreSql_User_Helpers(testContext, fqdn, adminUsername, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "PostgreSql User CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				PostgreSql_User_CRUD(testContext, flexibleServer, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "PostgreSql Secret Rollover",
			Test: func(testContext *testcommon.KubePerTestContext) {
				PostgreSql_AdminSecret_Rollover(testContext, fqdn, adminUsername, adminPasswordKey, adminPassword, secret)
			},
		},
	)
}

// PostgreSql_AdminSecret_Rollover ensures that when a secret is modified, the modified value
// is sent to Azure. This cannot be tested in the recording tests because they do not use
// a cached client. The index functionality used to check if a secret is being used by an
// ASO resource requires the cached client (the indexes are local to the cache).
func PostgreSql_AdminSecret_Rollover(tc *testcommon.KubePerTestContext, fqdn string, adminUsername string, adminPasswordKey string, adminPassword string, secret *v1.Secret) {
	// Connect to the DB
	conn, err := postgresqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		postgresqlutil.DefaultMaintanenceDatabase,
		postgresqlutil.PSqlServerPort,
		adminUsername,
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Update the secret
	newAdminPassword := tc.Namer.GeneratePasswordOfLength(40)

	newSecret := &v1.Secret{
		ObjectMeta: secret.ObjectMeta,
		StringData: map[string]string{
			adminPasswordKey: newAdminPassword,
		},
	}
	tc.UpdateResource(newSecret)

	// Connect to the DB - this may fail initially as reconcile runs and PostgreSql
	// performs the update
	tc.G.Eventually(
		func() error {
			conn, err = postgresqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				postgresqlutil.DefaultMaintanenceDatabase,
				postgresqlutil.PSqlServerPort,
				adminUsername,
				newAdminPassword)
			if err != nil {
				return err
			}

			return conn.Close()
		},
		2*time.Minute, // We expect this to pass pretty quickly
	).Should(Succeed())
}

// We could also test this with https://hub.docker.com/_/postgresql, but since we're provisioning a real SQL server anyway we might
// as well use it
func PostgreSql_User_Helpers(tc *testcommon.KubePerTestContext, fqdn string, adminUsername string, adminPassword string) {
	// Connect to the DB
	ctx := tc.Ctx
	db, err := postgresqlutil.ConnectToDB(
		ctx,
		fqdn,
		postgresqlutil.DefaultMaintanenceDatabase,
		postgresqlutil.PSqlServerPort,
		adminUsername,
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer db.Close()

	username := "testuser"
	userPassword := tc.Namer.GeneratePassword()
	sqlUser, err := postgresqlutil.CreateUser(ctx, db, username, userPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(sqlUser).ToNot(BeNil())

	exists, err := postgresqlutil.DoesUserExist(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	userRoles, err := postgresqlutil.GetUserServerRoles(ctx, db, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(userRoles).To(BeEmpty())

	// Test setting some user roles
	expectedUserRoles := []string{"azure_pg_admin"}
	tc.Expect(postgresqlutil.ReconcileUserServerRoles(ctx, db, postgresqlutil.SQLUser{Name: username}, expectedUserRoles)).To(Succeed())

	userRoles, err = postgresqlutil.GetUserServerRoles(ctx, db, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(userRoles).To(Equal(set.Make[string](expectedUserRoles...)))

	// Update user roles to add some and remove some
	expectedUserRoles = []string{"pg_read_all_stats"}
	tc.Expect(postgresqlutil.ReconcileUserServerRoles(ctx, db, postgresqlutil.SQLUser{Name: username}, expectedUserRoles)).To(Succeed())

	userRoles, err = postgresqlutil.GetUserServerRoles(ctx, db, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(userRoles).To(Equal(set.Make[string](expectedUserRoles...)))

	// Delete the user
	tc.Expect(postgresqlutil.DropUser(ctx, db, username)).To(Succeed())

	exists, err = postgresqlutil.DoesUserExist(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func PostgreSql_User_CRUD(tc *testcommon.KubePerTestContext, server *postgresql.FlexibleServer, adminPassword string) {
	passwordKey := "password"
	password := tc.Namer.GeneratePassword()
	userSecret := newPostgresSqlSecret(tc, passwordKey, password)

	tc.CreateResource(userSecret)

	username := tc.NoSpaceNamer.GenerateName("user")
	user := &postgresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(username),
		Spec: postgresqlv1.UserSpec{
			Owner: testcommon.AsOwner(server),
			Roles: []string{
				"azure_pg_admin",
			},
			LocalUser: &postgresqlv1.LocalUserSpec{
				ServerAdminUsername: to.Value(server.Spec.AdministratorLogin),
				ServerAdminPassword: server.Spec.AdministratorLoginPassword,
				Password: &genruntime.SecretReference{
					Name: userSecret.Name,
					Key:  passwordKey,
				},
			},
		},
	}
	tc.CreateResourcesAndWait(user)

	// Connect to the DB
	ctx := tc.Ctx
	fqdn := to.Value(server.Status.FullyQualifiedDomainName)
	conn, err := postgresqlutil.ConnectToDB(
		ctx,
		fqdn,
		postgresqlutil.DefaultMaintanenceDatabase,
		postgresqlutil.PSqlServerPort,
		to.Value(server.Spec.AdministratorLogin),
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	// Confirm that we have the right user roles on the actual server
	userRoles, err := postgresqlutil.GetUserServerRoles(tc.Ctx, conn, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(userRoles).To(Equal(set.Make[string](user.Spec.Roles...)))

	// Update the user
	old := user.DeepCopy()
	user.Spec.Roles = []string{
		"azure_pg_admin",
		"pg_read_all_stats",
	}
	tc.PatchResourceAndWait(old, user)

	// Confirm that we have the right user roles on the actual server
	userRoles, err = postgresqlutil.GetUserServerRoles(tc.Ctx, conn, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(userRoles).To(Equal(set.Make[string](user.Spec.Roles...)))

	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Confirm we can connect as the user
	conn, err = postgresqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		"",
		postgresqlutil.PSqlServerPort,
		user.Spec.AzureName,
		password)
	tc.Expect(err).ToNot(HaveOccurred())
	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Update the secret
	newPassword := tc.Namer.GeneratePassword()
	newSecret := &v1.Secret{
		ObjectMeta: userSecret.ObjectMeta,
		StringData: map[string]string{
			passwordKey: newPassword,
		},
	}
	tc.UpdateResource(newSecret)

	// Connect to the DB as the user, using the new secret
	tc.G.Eventually(
		func() error {
			conn, err = postgresqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				"",
				postgresqlutil.PSqlServerPort,
				user.Spec.AzureName,
				newPassword)
			if err != nil {
				return err
			}

			return conn.Close()
		},
		2*time.Minute, // We expect this to pass pretty quickly
	).Should(Succeed())

	tc.DeleteResourceAndWait(user)
}

func Test_PostgreSql_User(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePassword()
	adminSecret := newPostgresSqlSecret(tc, adminPasswordKey, adminPassword)

	passwordKey := "password"
	password := tc.Namer.GeneratePassword()
	userSecret := newPostgresSqlSecret(tc, passwordKey, password)

	tc.CreateResource(adminSecret)
	tc.CreateResource(userSecret)

	flexibleServer := newPostgreSqlServer(tc, rg, adminUsername, adminPasswordKey, adminSecret.Name)
	firewallRule := newPostgreSqlServerOpenFirewallRule(tc, flexibleServer)

	user := &postgresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("user")),
		Spec: postgresqlv1.UserSpec{
			Owner: testcommon.AsOwner(flexibleServer),
			Roles: []string{
				"azure_pg_admin",
			},
			LocalUser: &postgresqlv1.LocalUserSpec{
				ServerAdminUsername: adminUsername,
				ServerAdminPassword: flexibleServer.Spec.AdministratorLoginPassword,
				Password: &genruntime.SecretReference{
					Name: userSecret.Name,
					Key:  passwordKey,
				},
			},
		},
	}
	tc.CreateResourcesAndWait(flexibleServer, firewallRule, user)

	tc.DeleteResourceAndWait(user)
}

func newPostgresSqlSecret(tc *testcommon.KubePerTestContext, key string, password string) *v1.Secret {
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("postgresqlsecret"),
		StringData: map[string]string{
			key: password,
		},
	}

	return secret
}

func newPostgreSqlServer(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, adminUsername string, adminKey string, adminSecretName string) *postgresql.FlexibleServer {
	// Force this test to run in a region that is not capacity constrained.
	// location := tc.AzureRegion TODO: Uncomment this line when West US 2 is no longer constrained
	location := to.Ptr("australiaeast")

	version := postgresql.ServerVersion_13
	secretRef := genruntime.SecretReference{
		Name: adminSecretName,
		Key:  adminKey,
	}
	tier := postgresql.Sku_Tier_Burstable
	flexibleServer := &postgresql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("postgresql"),
		Spec: postgresql.FlexibleServer_Spec{
			Location: location,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &postgresql.Sku{
				Name: to.Ptr("Standard_B1ms"),
				Tier: &tier,
			},
			AdministratorLogin:         to.Ptr(adminUsername),
			AdministratorLoginPassword: &secretRef,
			Storage: &postgresql.Storage{
				StorageSizeGB: to.Ptr(32),
			},
		},
	}

	return flexibleServer
}

func newPostgreSqlServerOpenFirewallRule(tc *testcommon.KubePerTestContext, flexibleServer *postgresql.FlexibleServer) *postgresql.FlexibleServersFirewallRule {
	// This rule opens access to the public internet. Safe in this case
	// because there's no data in the database anyway
	firewallRule := &postgresql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: postgresql.FlexibleServers_FirewallRule_Spec{
			Owner:          testcommon.AsOwner(flexibleServer),
			StartIpAddress: to.Ptr("0.0.0.0"),
			EndIpAddress:   to.Ptr("255.255.255.255"),
		},
	}

	return firewallRule
}
