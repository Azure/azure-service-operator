/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql" //sql drive link
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	mysqlv1 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1"
	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	mysqlutil "github.com/Azure/azure-service-operator/v2/internal/util/mysql"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_MySQL_Combined(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePassword()
	secret := newSecret(tc, adminPasswordKey, adminPassword)

	tc.CreateResource(secret)

	flexibleServer := newMySQLServer(tc, rg, adminUsername, adminPasswordKey, secret.Name)
	tc.CreateResourceAndWait(flexibleServer)

	firewallRule := newMySQLServerOpenFirewallRule(tc, flexibleServer)
	tc.CreateResourceAndWait(firewallRule)

	tc.Expect(flexibleServer.Status.FullyQualifiedDomainName).ToNot(BeNil())
	fqdn := *flexibleServer.Status.FullyQualifiedDomainName

	// These must run sequentially as they're mutating SQL state
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "MySQL User Helpers",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQL_User_Helpers(testContext, fqdn, adminUsername, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "MySQL User CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQL_User_CRUD(testContext, flexibleServer, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "MySQL Secret Rollover",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQL_AdminSecret_Rollvoer(testContext, fqdn, adminUsername, adminPasswordKey, adminPassword, secret)
			},
		},
	)
}

// MySQL_AdminSecret_Rollover ensures that when a secret is modified, the modified value
// is sent to Azure. This cannot be tested in the recording tests because they do not use
// a cached client. The index functionality used to check if a secret is being used by an
// ASO resource requires the cached client (the indexes are local to the cache).
func MySQL_AdminSecret_Rollvoer(tc *testcommon.KubePerTestContext, fqdn string, adminUsername string, adminPasswordKey string, adminPassword string, secret *v1.Secret) {
	// Connect to the DB
	conn, err := mysqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		mysqlutil.SystemDatabase,
		mysqlutil.ServerPort,
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

	// Connect to the DB - this may fail initially as reconcile runs and MySQL
	// performs the update
	tc.G.Eventually(
		func() error {
			conn, err = mysqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				mysqlutil.SystemDatabase,
				mysqlutil.ServerPort,
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

// We could also test this with https://hub.docker.com/_/mysql, but since we're provisioning a real SQL server anyway we might
// as well use it
func MySQL_User_Helpers(tc *testcommon.KubePerTestContext, fqdn string, adminUsername string, adminPassword string) {
	// Connect to the DB
	ctx := tc.Ctx
	db, err := mysqlutil.ConnectToDB(
		ctx,
		fqdn,
		mysqlutil.SystemDatabase,
		mysqlutil.ServerPort,
		adminUsername,
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer db.Close()

	username := "testuser"
	hostname := ""
	userPassword := tc.Namer.GeneratePassword()
	tc.Expect(mysqlutil.CreateOrUpdateUser(ctx, db, username, hostname, userPassword)).To(Succeed())

	exists, err := mysqlutil.DoesUserExist(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	serverPrivs, err := mysqlutil.GetUserServerPrivileges(ctx, db, username, hostname)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(BeEmpty())

	dbPrivs, err := mysqlutil.GetUserDatabasePrivileges(ctx, db, username, hostname)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(dbPrivs).To(BeEmpty())

	// Test setting some privs
	expectedServerPrivs := []string{"CREATE USER", "PROCESS"}
	tc.Expect(mysqlutil.ReconcileUserServerPrivileges(ctx, db, username, hostname, expectedServerPrivs)).To(Succeed())

	serverPrivs, err = mysqlutil.GetUserServerPrivileges(ctx, db, username, hostname)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](expectedServerPrivs...)))

	// Update privs to add some and remove some
	expectedServerPrivs = []string{"CREATE USER", "SHOW DATABASES"}
	tc.Expect(mysqlutil.ReconcileUserServerPrivileges(ctx, db, username, hostname, expectedServerPrivs)).To(Succeed())

	serverPrivs, err = mysqlutil.GetUserServerPrivileges(ctx, db, username, hostname)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](expectedServerPrivs...)))

	// Delete the user
	tc.Expect(mysqlutil.DropUser(ctx, db, username)).To(Succeed())

	exists, err = mysqlutil.DoesUserExist(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func MySQL_User_CRUD(tc *testcommon.KubePerTestContext, server *mysql.FlexibleServer, adminPassword string) {
	passwordKey := "password"
	password := tc.Namer.GeneratePassword()
	userSecret := newSecret(tc, passwordKey, password)

	tc.CreateResource(userSecret)

	username := tc.NoSpaceNamer.GenerateName("user")
	user := &mysqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(username),
		Spec: mysqlv1.UserSpec{
			Owner: testcommon.AsOwner(server),
			Privileges: []string{
				"CREATE USER",
				"PROCESS",
			},
			LocalUser: &mysqlv1.LocalUserSpec{
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
	conn, err := mysqlutil.ConnectToDB(
		ctx,
		fqdn,
		mysqlutil.SystemDatabase,
		mysqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	// Confirm that we have the right privs on the actual server
	serverPrivs, err := mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, username, "")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](user.Spec.Privileges...)))

	// Update the user
	old := user.DeepCopy()
	user.Spec.Privileges = []string{
		"CREATE USER",
		"PROCESS",
		"SHOW DATABASES",
	}
	tc.PatchResourceAndWait(old, user)

	// Confirm that we have the right privs on the actual server
	serverPrivs, err = mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, username, "")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](user.Spec.Privileges...)))

	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Confirm we can connect as the user
	conn, err = mysqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		"",
		mysqlutil.ServerPort,
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
			conn, err = mysqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				"",
				mysqlutil.ServerPort,
				user.Spec.AzureName,
				newPassword)
			if err != nil {
				return err
			}

			return conn.Close()
		},
		2*time.Minute, // We expect this to pass pretty quickly
	).Should(Succeed())

	originalUser := user.DeepCopy()

	// Confirm that we cannot change the user owner
	old = user.DeepCopy()
	user.Spec.Owner.Name = "adifferentowner"
	err = tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("updating 'Owner.Name' is not allowed"))

	// Confirm that we cannot change the user AzureName
	user = originalUser.DeepCopy()
	old = user.DeepCopy()
	user.Spec.AzureName = "adifferentname"
	err = tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("updating 'AzureName' is not allowed"))

	// Confirm that we cannot change the user type from local to AAD
	user = originalUser.DeepCopy()
	old = user.DeepCopy()
	user.Spec.LocalUser = nil
	user.Spec.AADUser = &mysqlv1.AADUserSpec{
		ServerAdminUsername: "someadminuser",
	}
	err = tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("cannot change from local user to AAD user"))

	user = originalUser.DeepCopy()
	tc.DeleteResourceAndWait(user)

	conn, err = mysqlutil.ConnectToDB(
		ctx,
		fqdn,
		mysqlutil.SystemDatabase,
		mysqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())

	exists, err := mysqlutil.DoesUserExist(ctx, conn, user.Name)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

//func Test_MySQL_Helpers(t *testing.T) {
//	t.Parallel()
//	tc := globalTestContext.ForTest(t)
//
//	rg := tc.CreateTestResourceGroupAndWait()
//
//	adminUsername := "myadmin"
//	adminPasswordKey := "adminPassword"
//	adminPassword := tc.Namer.GeneratePassword()
//	secret := newSecret(tc, adminPasswordKey, adminPassword)
//
//	tc.CreateResource(secret)
//
//	flexibleServer := newMySQLServer(tc, rg, adminUsername, adminPasswordKey, secret.Name)
//	tc.CreateResourceAndWait(flexibleServer)
//
//	firewallRule := newMySQLServerOpenFirewallRule(tc, flexibleServer)
//	tc.CreateResourceAndWait(firewallRule)
//
//	tc.Expect(flexibleServer.Status.FullyQualifiedDomainName).ToNot(BeNil())
//	fqdn := *flexibleServer.Status.FullyQualifiedDomainName
//
//	// Connect to the DB
//	ctx := context.Background()
//	db, err := mysqlutil.ConnectToDB(
//		ctx,
//		fqdn,
//		mysqlutil.SystemDatabase,
//		mysqlutil.ServerPort,
//		adminUsername,
//		adminPassword)
//	tc.Expect(err).ToNot(HaveOccurred())
//	defer db.Close()
//
//	// TODO: These should be subtests, maybe?
//	username := "testuser"
//	hostname := ""
//	userPassword := tc.Namer.GeneratePassword()
//	tc.Expect(mysqlutil.CreateUser(ctx, db, username, hostname, userPassword)).To(Succeed())
//
//	exists, err := mysqlutil.DoesUserExist(ctx, db, username)
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(exists).To(BeTrue())
//
//	serverPrivs, err := mysqlutil.ExtractUserServerPrivileges(ctx, db, username, hostname)
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(serverPrivs).To(BeEmpty())
//
//	dbPrivs, err := mysqlutil.ExtractUserDatabasePrivileges(ctx, db, username, hostname)
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(dbPrivs).To(BeEmpty())
//
//	// Test setting some privs
//	expectedServerPrivs := []string{"CREATE USER", "PROCESS"}
//	tc.Expect(mysqlutil.EnsureUserServerPrivileges(ctx, db, username, hostname, expectedServerPrivs)).To(Succeed())
//
//	serverPrivs, err = mysqlutil.ExtractUserServerPrivileges(ctx, db, username, hostname)
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(serverPrivs).To(Equal(set.Make[string](expectedServerPrivs...)))
//
//	// Update privs to add some and remove some
//	expectedServerPrivs = []string{"CREATE USER", "SHOW DATABASES"}
//	tc.Expect(mysqlutil.EnsureUserServerPrivileges(ctx, db, username, hostname, expectedServerPrivs)).To(Succeed())
//
//	serverPrivs, err = mysqlutil.ExtractUserServerPrivileges(ctx, db, username, hostname)
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(serverPrivs).To(Equal(set.Make[string](expectedServerPrivs...)))
//
//	// Delete the user
//	tc.Expect(mysqlutil.DropUser(ctx, db, username)).To(Succeed())
//
//	exists, err = mysqlutil.DoesUserExist(ctx, db, username)
//	tc.Expect(err).ToNot(HaveOccurred())
//	tc.Expect(exists).To(BeFalse())
//}

func Test_MySQL_User(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePassword()
	adminSecret := newSecret(tc, adminPasswordKey, adminPassword)

	passwordKey := "password"
	password := tc.Namer.GeneratePassword()
	userSecret := newSecret(tc, passwordKey, password)

	tc.CreateResource(adminSecret)
	tc.CreateResource(userSecret)

	flexibleServer := newMySQLServer(tc, rg, adminUsername, adminPasswordKey, adminSecret.Name)
	firewallRule := newMySQLServerOpenFirewallRule(tc, flexibleServer)

	user := &mysqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("user")),
		Spec: mysqlv1.UserSpec{
			Owner: testcommon.AsOwner(flexibleServer),
			Privileges: []string{
				"CREATE USER",
				"PROCESS",
			},
			LocalUser: &mysqlv1.LocalUserSpec{
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

	// TODO: Test other stuff?
	// TODO: Password rollover?

	tc.DeleteResourceAndWait(user)
}

func newSecret(tc *testcommon.KubePerTestContext, key string, password string) *v1.Secret {
	secret := &v1.Secret{
		ObjectMeta: tc.MakeObjectMeta("mysqlsecret"),
		StringData: map[string]string{
			key: password,
		},
	}

	return secret
}

func newMySQLServer(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, adminUsername string, adminKey string, adminSecretName string) *mysql.FlexibleServer {
	// Force this test to run in a region that is not capacity constrained.
	// location := tc.AzureRegion TODO: Uncomment this line when West US 2 is no longer constrained
	location := to.Ptr("australiaeast")

	version := mysql.ServerVersion_8021
	secretRef := genruntime.SecretReference{
		Name: adminSecretName,
		Key:  adminKey,
	}
	tier := mysql.Sku_Tier_GeneralPurpose
	flexibleServer := &mysql.FlexibleServer{
		ObjectMeta: tc.MakeObjectMeta("mysql"),
		Spec: mysql.FlexibleServer_Spec{
			Location: location,
			Owner:    testcommon.AsOwner(rg),
			Version:  &version,
			Sku: &mysql.Sku{
				Name: to.Ptr("Standard_D4ds_v4"),
				Tier: &tier,
			},
			AdministratorLogin:         to.Ptr(adminUsername),
			AdministratorLoginPassword: &secretRef,
			Storage: &mysql.Storage{
				StorageSizeGB: to.Ptr(128),
			},
		},
	}

	return flexibleServer
}

func newMySQLServerOpenFirewallRule(tc *testcommon.KubePerTestContext, flexibleServer *mysql.FlexibleServer) *mysql.FlexibleServersFirewallRule {
	// This rule opens access to the public internet. Safe in this case
	// because there's no data in the database anyway
	firewallRule := &mysql.FlexibleServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: mysql.FlexibleServers_FirewallRule_Spec{
			Owner:          testcommon.AsOwner(flexibleServer),
			StartIpAddress: to.Ptr("0.0.0.0"),
			EndIpAddress:   to.Ptr("255.255.255.255"),
		},
	}

	return firewallRule
}
