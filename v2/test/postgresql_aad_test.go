/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"database/sql"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	_ "github.com/jackc/pgx/v5/stdlib" // the pgx lib
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	postgresqlv1 "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1"
	postgresql "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v20250801"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	sqlv1 "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	postgresqlutil "github.com/Azure/azure-service-operator/v2/internal/util/postgresql"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
)

// Test_PostgreSQL_AADUser makes the identity running the tests (which is also the operator's identity) the server's
// Microsoft Entra administrator, then manages roles bound to a user-assigned managed identity. Roles are bound by object
// ID (pgaadauth_create_principal_with_oid), which needs no extra directory permissions, so this runs in the CI
// subscription. In CI, AZURE_TEST_IDENTITY_NAME names the managed identity the tests run as; locally, sign in as a user.
func Test_PostgreSQL_AADUser(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)
	tc.AzureRegion = to.Ptr(postgresqlTestRegion)

	// Working out who the tests run as is the same as for Azure SQL; only the token scope differs.
	adminIdentity, err := azureSQLAdminForTest(tc)
	tc.Expect(err).ToNot(HaveOccurred())

	adminPrincipalType := postgresql.PrincipalType_ServicePrincipal
	if adminIdentity.principalType == sqlv1.ServerExternalAdministrator_PrincipalType_User {
		adminPrincipalType = postgresql.PrincipalType_User
	}

	rg := tc.CreateTestResourceGroupAndWait()

	adminPasswordKey := "adminPassword"
	adminSecret := newPostgresSQLSecret(tc, adminPasswordKey, tc.Namer.GeneratePassword())
	tc.CreateResource(adminSecret)

	server := newPostgreSQLServer(tc, rg, "myadmin", adminPasswordKey, adminSecret.Name)
	server.Spec.AuthConfig = &postgresql.AuthConfig{
		ActiveDirectoryAuth: to.Ptr(postgresql.MicrosoftEntraAuth_Enabled),
		PasswordAuth:        to.Ptr(postgresql.AuthConfig_PasswordAuth_Enabled),
		TenantId:            to.Ptr(adminIdentity.tenantID),
	}
	firewallRule := newPostgreSQLServerOpenFirewallRule(tc, server)

	// The managed identity the AAD users are bound to. Nothing signs in as it; the tests check the binding.
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}

	tc.CreateResourcesAndWait(server, firewallRule, mi)
	tc.Expect(server.Status.FullyQualifiedDomainName).ToNot(BeNil())
	tc.Expect(mi.Status.PrincipalId).ToNot(BeNil())

	administrator := &postgresql.FlexibleServersAdministrator{
		ObjectMeta: tc.MakeObjectMeta("aadadmin"),
		Spec: postgresql.FlexibleServersAdministrator_Spec{
			Owner:         testcommon.AsOwner(server),
			AzureName:     adminIdentity.objectID,
			PrincipalName: to.Ptr(adminIdentity.login),
			PrincipalType: to.Ptr(adminPrincipalType),
			TenantId:      to.Ptr(adminIdentity.tenantID),
		},
	}
	tc.CreateResourceAndWait(administrator)

	admin := postgreSQLAADAdmin{
		fqdn:  to.Value(server.Status.FullyQualifiedDomainName),
		login: adminIdentity.login,
		tc:    tc,
	}

	// Ensure that the firewall rule and the administrator have taken effect
	tc.G.Eventually(
		func() error {
			db, err := admin.connect()
			if err != nil {
				return err
			}

			return db.Close()
		},
		5*time.Minute,
	).Should(Succeed())

	principalID := to.Value(mi.Status.PrincipalId)

	// These must run sequentially as they're mutating SQL state
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "PostgreSQL AAD User CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				PostgreSQL_AADUser_CRUD(testContext, server, admin, principalID)
			},
		},
		testcommon.Subtest{
			Name: "PostgreSQL AAD User neither adopts nor drops a local role",
			Test: func(testContext *testcommon.KubePerTestContext) {
				PostgreSQL_AADUser_LocalRoleCollision(testContext, server, admin, principalID)
			},
		},
	)
}

// PostgreSQL_AADUser_CRUD creates a role bound to a managed identity, updates its roles (which reconciles the existing
// role, checking its binding), and deletes it.
func PostgreSQL_AADUser_CRUD(
	tc *testcommon.KubePerTestContext,
	server *postgresql.FlexibleServer,
	admin postgreSQLAADAdmin,
	principalID string,
) {
	username := tc.NoSpaceNamer.GenerateName("aaduser")
	user := &postgresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(username),
		Spec: postgresqlv1.UserSpec{
			Owner: testcommon.AsKubernetesOwner(server),
			Roles: []string{"pg_read_all_stats"},
			AADUser: &postgresqlv1.AADUserSpec{
				ServerAdminUsername: admin.login,
				ObjectId:            principalID,
				PrincipalType:       "service",
			},
		},
	}
	tc.CreateResourceAndWait(user)

	db, err := admin.connect()
	tc.Expect(err).ToNot(HaveOccurred())
	defer db.Close()

	// The role is bound to the managed identity
	principal, err := postgresqlutil.GetAADPrincipal(tc.Ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(principal).ToNot(BeNil())
	tc.Expect(principal.ObjectID).To(testcommon.EqualsIgnoreCase(principalID))
	tc.Expect(principal.PrincipalType).To(testcommon.EqualsIgnoreCase("service"))

	roles, err := postgresqlutil.GetUserServerRoles(tc.Ctx, db, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make(user.Spec.Roles...)))

	// Update the roles: the role now exists, so this reconcile verifies its binding before granting them
	old := user.DeepCopy()
	user.Spec.Roles = []string{"pg_read_all_stats", "pg_read_all_settings"}
	tc.PatchResourceAndWait(old, user)

	roles, err = postgresqlutil.GetUserServerRoles(tc.Ctx, db, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(Equal(set.Make(user.Spec.Roles...)))

	tc.DeleteResourceAndWait(user)

	exists, err := postgresqlutil.DoesUserExist(tc.Ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())
}

// PostgreSQL_AADUser_LocalRoleCollision checks that an AAD user whose name is taken by a local (password) role refuses
// to adopt it, and that deleting the failed resource leaves the local role alone.
func PostgreSQL_AADUser_LocalRoleCollision(
	tc *testcommon.KubePerTestContext,
	server *postgresql.FlexibleServer,
	admin postgreSQLAADAdmin,
	principalID string,
) {
	db, err := admin.connect()
	tc.Expect(err).ToNot(HaveOccurred())
	defer db.Close()

	username := tc.NoSpaceNamer.GenerateName("localrole")
	_, err = postgresqlutil.CreateUser(tc.Ctx, db, username, tc.Namer.GeneratePassword())
	tc.Expect(err).ToNot(HaveOccurred())
	defer func() {
		tc.Expect(postgresqlutil.DropUser(tc.Ctx, db, username)).To(Succeed())
	}()

	user := &postgresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(username),
		Spec: postgresqlv1.UserSpec{
			Owner: testcommon.AsKubernetesOwner(server),
			Roles: []string{"pg_read_all_stats"},
			AADUser: &postgresqlv1.AADUserSpec{
				ServerAdminUsername: admin.login,
				ObjectId:            principalID,
				PrincipalType:       "service",
			},
		},
	}
	tc.CreateResourceAndWaitForState(user, metav1.ConditionFalse, conditions.ConditionSeverityError)

	// The local role was neither bound nor granted anything
	principal, err := postgresqlutil.GetAADPrincipal(tc.Ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(principal).To(BeNil())

	roles, err := postgresqlutil.GetUserServerRoles(tc.Ctx, db, postgresqlutil.SQLUser{Name: username})
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(roles).To(BeEmpty())

	// Deleting the resource doesn't drop the local role
	tc.DeleteResourceAndWait(user)

	exists, err := postgresqlutil.DoesUserExist(tc.Ctx, db, username)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeTrue())
}

// postgreSQLAADAdmin signs in to a server as its Microsoft Entra administrator: the identity running the tests.
type postgreSQLAADAdmin struct {
	fqdn  string
	login string
	tc    *testcommon.KubePerTestContext
}

func (a postgreSQLAADAdmin) connect() (*sql.DB, error) {
	token, err := a.tc.AzureClient.Creds().GetToken(a.tc.Ctx, policy.TokenRequestOptions{Scopes: []string{postgresqlutil.Scope}})
	if err != nil {
		return nil, err
	}

	return postgresqlutil.ConnectToDBAAD(
		a.tc.Ctx,
		a.fqdn,
		postgresqlutil.DefaultMaintanenceDatabase,
		postgresqlutil.PSqlServerPort,
		a.login,
		token.Token)
}
