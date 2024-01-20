/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"
	"time"

	_ "github.com/microsoft/go-mssqldb"
	. "github.com/onsi/gomega"
	v1 "k8s.io/api/core/v1"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	azuresqlv1 "github.com/Azure/azure-service-operator/v2/api/sql/v1"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	azuresqlutil "github.com/Azure/azure-service-operator/v2/internal/util/azuresql"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

func Test_AzureSQL_Combined(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Use a different region where we have quota
	tc.AzureRegion = to.Ptr("eastus")

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePasswordOfLength(60) // Use a long password to ensure we meet complexity requirements

	secret := newSecret(tc, adminPasswordKey, adminPassword)
	tc.CreateResource(secret)

	server := newAzureSQLServer(tc, rg, adminUsername, adminPasswordKey, secret.Name)
	database := newAzureSQLServerDatabase(tc, server)
	firewallRule := newSQLServerOpenFirewallRule(tc, server)

	tc.CreateResourcesAndWait(server, database, firewallRule)

	tc.Expect(server.Status.FullyQualifiedDomainName).ToNot(BeNil())
	fqdn := *server.Status.FullyQualifiedDomainName

	// Ensure that firewall rule access has worked. It can take up to 5 minutes to take effect
	tc.G.Eventually(
		func() error {
			db, err := azuresqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				database.AzureName(),
				azuresqlutil.ServerPort,
				adminUsername,
				adminPassword)
			if err != nil {
				return err
			}

			return db.Close()
		},
		7*time.Minute,
	).Should(Succeed())

	// These must run sequentially as they're mutating SQL state
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "AzureSQL User Helpers",
			Test: func(testContext *testcommon.KubePerTestContext) {
				AzureSQL_User_Helpers(testContext, fqdn, database.AzureName(), adminUsername, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "AzureSQL User CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				AzureSQL_User_CRUD(testContext, server, database, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "AzureSQL Secret Rollover",
			Test: func(testContext *testcommon.KubePerTestContext) {
				AzureSQL_AdminSecret_Rollover(testContext, fqdn, database.AzureName(), adminUsername, adminPasswordKey, adminPassword, secret)
			},
		},
	)
}

// AzureSQL_AdminSecret_Rollover ensures that when a secret is modified, the modified value
// is sent to Azure. This cannot be tested in the recording tests because it's not possible
// to attempt to connect to the server in replay mode (there is no server).
func AzureSQL_AdminSecret_Rollover(tc *testcommon.KubePerTestContext, fqdn string, database string, adminUsername string, adminPasswordKey string, adminPassword string, secret *v1.Secret) {
	ctx := tc.Ctx

	// Connect to the DB
	db, err := azuresqlutil.ConnectToDB(
		ctx,
		fqdn,
		database,
		azuresqlutil.ServerPort,
		adminUsername,
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	// Close the connection
	tc.Expect(db.Close()).To(Succeed())

	// Update the secret
	newAdminPassword := tc.Namer.GeneratePasswordOfLength(60)

	newSecret := &v1.Secret{
		ObjectMeta: secret.ObjectMeta,
		StringData: map[string]string{
			adminPasswordKey: newAdminPassword,
		},
	}
	tc.UpdateResource(newSecret)

	// Connect to the DB - this may fail initially as reconcile runs and the server updates
	tc.G.Eventually(
		func() error {
			db, err = azuresqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				database,
				azuresqlutil.ServerPort,
				adminUsername,
				newAdminPassword)
			if err != nil {
				return err
			}

			return db.Close()
		},
		2*time.Minute, // We expect this to pass pretty quickly
	).Should(Succeed())
}

func AzureSQL_User_Helpers(tc *testcommon.KubePerTestContext, fqdn string, database string, adminUsername string, adminPassword string) {
	// Connect to the DB
	ctx := tc.Ctx

	db, err := azuresqlutil.ConnectToDB(
		ctx,
		fqdn,
		database,
		azuresqlutil.ServerPort,
		adminUsername,
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer db.Close()

	username := "testuser"
	userPassword := tc.Namer.GeneratePasswordOfLength(60) // Use a long password to ensure we meet complexity requirements
	tc.Expect(azuresqlutil.CreateOrUpdateUser(ctx, db, username, userPassword)).To(Succeed())

	exists, err := azuresqlutil.DoesUserExist(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())

	roles, err := azuresqlutil.GetUserRoles(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(BeEmpty())

	// Test setting some roles
	expectedRoles := []string{"db_datareader", "db_datawriter"}
	tc.Expect(azuresqlutil.ReconcileUserRoles(ctx, db, username, expectedRoles)).To(Succeed())

	roles, err = azuresqlutil.GetUserRoles(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make[string](expectedRoles...)))

	// Update roles to add some and remove some
	expectedRoles = []string{"db_securityadmin", "db_datawriter"}
	tc.Expect(azuresqlutil.ReconcileUserRoles(ctx, db, username, expectedRoles)).To(Succeed())

	roles, err = azuresqlutil.GetUserRoles(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make[string](expectedRoles...)))

	// Delete the user
	tc.Expect(azuresqlutil.DropUser(ctx, db, username)).To(Succeed())

	exists, err = azuresqlutil.DoesUserExist(ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func AzureSQL_User_CRUD(tc *testcommon.KubePerTestContext, server *sql.Server, database *sql.ServersDatabase, adminPassword string) {
	passwordKey := "password"
	password := tc.Namer.GeneratePasswordOfLength(60) // Use a long password to ensure we meet complexity requirements
	userSecret := newSecret(tc, passwordKey, password)

	tc.CreateResource(userSecret)

	username := tc.NoSpaceNamer.GenerateName("user")
	user := &azuresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(username),
		Spec: azuresqlv1.UserSpec{
			Owner: testcommon.AsOwner(database),
			Roles: []string{
				"db_datareader",
			},
			LocalUser: &azuresqlv1.LocalUserSpec{
				// TODO: Should adminusername be able to be sourced from a configmap?
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
	fqdn := to.Value(server.Status.FullyQualifiedDomainName)
	conn, err := azuresqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		database.AzureName(),
		azuresqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	// Confirm that we have the right roles
	roles, err := azuresqlutil.GetUserRoles(tc.Ctx, conn, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make[string](user.Spec.Roles...)))

	// Update the user
	old := user.DeepCopy()
	user.Spec.Roles = []string{
		"db_datareader",
		"db_datawriter",
		"db_securityadmin",
	}
	tc.PatchResourceAndWait(old, user)

	// Confirm that we have the right roles
	roles, err = azuresqlutil.GetUserRoles(tc.Ctx, conn, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make[string](user.Spec.Roles...)))

	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Confirm we can connect as the user
	conn, err = azuresqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		database.AzureName(),
		azuresqlutil.ServerPort,
		user.Spec.AzureName,
		password)
	tc.Expect(err).ToNot(HaveOccurred())
	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())

	// Update the secret
	newPassword := tc.Namer.GeneratePasswordOfLength(60)
	updatedSecret := &v1.Secret{
		ObjectMeta: userSecret.ObjectMeta,
		StringData: map[string]string{
			passwordKey: newPassword,
		},
	}
	tc.UpdateResource(updatedSecret)

	// Connect to the DB as the user, using the new secret
	tc.G.Eventually(
		func() error {
			conn, err = azuresqlutil.ConnectToDB(
				tc.Ctx,
				fqdn,
				database.AzureName(),
				azuresqlutil.ServerPort,
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

	user = originalUser.DeepCopy()
	tc.DeleteResourceAndWait(user)

	conn, err = azuresqlutil.ConnectToDB(
		tc.Ctx,
		fqdn,
		database.AzureName(),
		azuresqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		adminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	exists, err := azuresqlutil.DoesUserExist(tc.Ctx, conn, user.Name)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

func newAzureSQLServer(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, adminUsername string, adminKey string, adminSecretName string) *sql.Server {
	secretRef := genruntime.SecretReference{
		Name: adminSecretName,
		Key:  adminKey,
	}
	server := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserver"),
		Spec: sql.Server_Spec{
			Location:                   tc.AzureRegion,
			Owner:                      testcommon.AsOwner(rg),
			AdministratorLogin:         to.Ptr(adminUsername),
			AdministratorLoginPassword: &secretRef,
			Version:                    to.Ptr("12.0"),
		},
	}

	return server
}

func newAzureSQLServerDatabase(tc *testcommon.KubePerTestContext, server *sql.Server) *sql.ServersDatabase {
	db := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.Servers_Database_Spec{
			Owner:     testcommon.AsOwner(server),
			Location:  tc.AzureRegion,
			Collation: to.Ptr("SQL_Latin1_General_CP1_CI_AS"),
		},
	}

	return db
}

func newSQLServerOpenFirewallRule(tc *testcommon.KubePerTestContext, server *sql.Server) *sql.ServersFirewallRule {
	firewall := &sql.ServersFirewallRule{
		ObjectMeta: tc.MakeObjectMeta("firewall"),
		Spec: sql.Servers_FirewallRule_Spec{
			Owner:          testcommon.AsOwner(server),
			StartIpAddress: to.Ptr("0.0.0.0"),
			EndIpAddress:   to.Ptr("255.255.255.255"),
		},
	}

	return firewall
}
