/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"fmt"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql" //sql drive link
	. "github.com/onsi/gomega"
	"github.com/pkg/errors"

	mysqlv1 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1"
	mysql "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20210501"
	mysql20220101 "github.com/Azure/azure-service-operator/v2/api/dbformysql/v1api20220101"
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	mysqlutil "github.com/Azure/azure-service-operator/v2/internal/util/mysql"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
	"github.com/Azure/azure-service-operator/v2/pkg/common/annotations"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

const (
	AzureManagedIdentityNameVar          = "AZURE_IDENTITY_NAME"
	AzureManagedIdentityResourceGroupVar = "AZURE_IDENTITY_RG"
)

// This test current cannot be run in the CI subscription.
// In order to run it you must:
//  1. Create a user managed identity to server as the MySQL admin and grant it the required graph permissions as documented at
//     https://learn.microsoft.com/en-us/azure/mysql/flexible-server/how-to-azure-ad#grant-permissions-to-user-assigned-managed-identity.
//     Note that this must be done in an AAD tenant where you have admin permissions, as they are required to assign the
//     required roles to the identity.
//  2. Assign the AZURE_IDENTITY_NAME env variable to the name of the identity.
//  3. Assign the AZURE_IDENTITY_RG to the RG the identity resides in.
//  4. Extract the identityClientID and save it in the cost below.
//  5. Create an AAD user, for example "foo@bar.onmicrosoft.com", and save its name in the aadUserName const below
//  6. Create an AAD group, for example "sqltest" and save its name in the aadGroupName const below
const (
	identityClientID = "6bfc287b-7af9-4c10-9a3f-52cf152070dd"
	aadUserName      = "myidentity@domain.onmicrosoft.com"
	aadGroupName     = "mysqlgroup"
)

func Test_MySQL_AADUser(t *testing.T) {
	t.Parallel()
	t.Skip("Cannot run in CI currently due to AAD tenant restrictions on CI subscription")

	tc := globalTestContext.ForTest(t)

	identityName, identityResourceGroup, err := readEnv()
	tc.Expect(err).ToNot(HaveOccurred())

	rg := tc.CreateTestResourceGroupAndWait()

	adminUsername := "myadmin"
	adminPasswordKey := "adminPassword"
	adminPassword := tc.Namer.GeneratePassword()
	secret := newSecret(tc, adminPasswordKey, adminPassword)

	tc.CreateResource(secret)

	identityARMID := fmt.Sprintf(
		"/subscriptions/%s/resourceGroups/%s/providers/Microsoft.ManagedIdentity/userAssignedIdentities/%s",
		tc.AzureSubscription,
		identityResourceGroup,
		identityName)

	server := newMySQLServer(tc, rg, adminUsername, adminPasswordKey, secret.Name)
	server.Spec.Identity = &mysql.Identity{
		Type: to.Ptr(mysql.Identity_Type_UserAssigned),
		UserAssignedIdentities: []mysql.UserAssignedIdentityDetails{
			{
				Reference: genruntime.ResourceReference{
					ARMID: identityARMID,
				},
			},
		},
	}

	admin := &mysql20220101.FlexibleServersAdministrator{
		ObjectMeta: tc.MakeObjectMeta("aadadmin"),
		Spec: mysql20220101.FlexibleServers_Administrator_Spec{
			Owner:             testcommon.AsOwner(server),
			AdministratorType: to.Ptr(mysql20220101.AdministratorProperties_AdministratorType_ActiveDirectory),
			Login:             to.Ptr(identityName),
			TenantId:          to.Ptr(tc.AzureTenant),
			Sid:               to.Ptr(identityClientID),
			IdentityResourceReference: &genruntime.ResourceReference{
				ARMID: identityARMID,
			},
		},
	}

	firewallRule := newMySQLServerOpenFirewallRule(tc, server)

	tc.CreateResourcesAndWait(server, admin, firewallRule)

	tc.Expect(server.Status.FullyQualifiedDomainName).ToNot(BeNil())

	// These must run sequentially as they're mutating SQL state
	tc.RunSubtests(
		testcommon.Subtest{
			Name: "MySQL AAD User CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQL_AADUser_CRUD(testContext, server, rg, admin, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "MySQL Local User via AAD Admin CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQL_LocalUser_AADAdmin_CRUD(testContext, server, admin, adminPassword)
			},
		},
		testcommon.Subtest{
			Name: "MySQL AAD Group CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				MySQL_AADUserAndGroup_CRUD(testContext, server, admin, adminPassword)
			},
		},
	)
}

func MySQL_AADUser_CRUD(
	tc *testcommon.KubePerTestContext,
	server *mysql.FlexibleServer,
	rg *resources.ResourceGroup,
	admin *mysql20220101.FlexibleServersAdministrator,
	standardAdminPassword string) {

	configMapName := "my-configmap"
	clientIDKey := "clientId"
	tenantIDKey := "tenantId"

	// Create a managed identity to use as the AAD user
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

	user := &mysqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(mi.Name),
		Spec: mysqlv1.UserSpec{
			Owner: testcommon.AsOwner(server),
			Privileges: []string{
				"CREATE USER",
				"PROCESS",
			},
			AADUser: &mysqlv1.AADUserSpec{
				// No Alias here, to ensure that path works
				ServerAdminUsername: to.Value(admin.Spec.Login),
			},
		},
	}
	tc.CreateResourcesAndWait(user)

	// Connect to the DB
	ctx := tc.Ctx
	fqdn := to.Value(server.Status.FullyQualifiedDomainName)
	conn, err := mysqlutil.ConnectToDBAAD(
		ctx,
		fqdn,
		mysqlutil.SystemDatabase,
		mysqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		standardAdminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	// Confirm that we have the right privs on the actual server
	serverPrivs, err := mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, mi.Name, "")
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
	serverPrivs, err = mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, mi.Name, "")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](user.Spec.Privileges...)))

	// Update the user once again, this time to remove privs and also use a resource scoped identity
	secret := newManagedIdentityCredential(tc.AzureSubscription, tc.AzureTenant, to.Value(admin.Spec.Sid), "credential", tc.Namespace)
	tc.CreateResource(secret)

	old = user.DeepCopy()
	user.Annotations[annotations.PerResourceSecret] = secret.Name
	user.Spec.Privileges = []string{
		"SHOW DATABASES",
	}

	tc.PatchResourceAndWait(old, user)

	// Confirm that we have the right privs on the actual server
	serverPrivs, err = mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, mi.Name, "")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](user.Spec.Privileges...)))

	originalUser := user.DeepCopy()

	// Confirm that we cannot change the user AAD Alias
	old = user.DeepCopy()
	user.Spec.AADUser.Alias = "adifferentalias"
	err = tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("cannot change AAD user 'alias'"))

	// Confirm that we cannot change the user type from local to AAD
	user = originalUser.DeepCopy()
	old = user.DeepCopy()
	user.Spec.LocalUser = &mysqlv1.LocalUserSpec{
		ServerAdminUsername: "someadminuser",
	}
	err = tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("cannot change from AAD User to local user"))

	user = originalUser.DeepCopy()
	tc.DeleteResourceAndWait(user)

	exists, err := mysqlutil.DoesUserExist(ctx, conn, user.Name)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())

	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())
}

func MySQL_LocalUser_AADAdmin_CRUD(
	tc *testcommon.KubePerTestContext,
	server *mysql.FlexibleServer,
	admin *mysql20220101.FlexibleServersAdministrator,
	standardAdminPassword string) {

	passwordKey := "password"
	password := tc.Namer.GeneratePassword()
	userSecret := newSecret(tc, passwordKey, password)

	tc.CreateResource(userSecret)

	user := &mysqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(tc.NoSpaceNamer.GenerateName("local")),
		Spec: mysqlv1.UserSpec{
			Owner: testcommon.AsOwner(server),
			Privileges: []string{
				"CREATE USER",
				"PROCESS",
			},
			LocalUser: &mysqlv1.LocalUserSpec{
				ServerAdminUsername: to.Value(admin.Spec.Login),
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
	conn, err := mysqlutil.ConnectToDBAAD(
		ctx,
		fqdn,
		mysqlutil.SystemDatabase,
		mysqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		standardAdminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	// Confirm that we have the right privs on the actual server
	serverPrivs, err := mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, user.Name, "")
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
	serverPrivs, err = mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, user.Name, "")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](user.Spec.Privileges...)))

	tc.DeleteResourceAndWait(user)

	exists, err := mysqlutil.DoesUserExist(ctx, conn, user.Name)
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(exists).To(BeFalse())

	// Close the connection
	tc.Expect(conn.Close()).To(Succeed())
	tc.T.Logf("Done")
}

func MySQL_AADUserAndGroup_CRUD(
	tc *testcommon.KubePerTestContext,
	server *mysql.FlexibleServer,
	admin *mysql20220101.FlexibleServersAdministrator,
	standardAdminPassword string) {

	// Note: when logging in to the DB you still log in with the actual username not the alias.
	// Alias is just for management of it in SQL.
	alias := "myaliaseduser"

	user := &mysqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName("aaduser"),
		Spec: mysqlv1.UserSpec{
			Owner:     testcommon.AsOwner(server),
			AzureName: aadUserName,
			Privileges: []string{
				"CREATE USER",
				"PROCESS",
			},
			AADUser: &mysqlv1.AADUserSpec{
				Alias:               alias,
				ServerAdminUsername: to.Value(admin.Spec.Login),
			},
		},
	}
	tc.CreateResourcesAndWait(user)

	// Connect to the DB
	ctx := tc.Ctx
	fqdn := to.Value(server.Status.FullyQualifiedDomainName)
	conn, err := mysqlutil.ConnectToDBAAD(
		ctx,
		fqdn,
		mysqlutil.SystemDatabase,
		mysqlutil.ServerPort,
		to.Value(server.Spec.AdministratorLogin),
		standardAdminPassword)
	tc.Expect(err).ToNot(HaveOccurred())
	defer conn.Close()

	// Confirm that we have the right privs on the actual server
	serverPrivs, err := mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, alias, "")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](user.Spec.Privileges...)))

	user2 := &mysqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName("groupuser"),
		Spec: mysqlv1.UserSpec{
			Owner:     testcommon.AsOwner(server),
			AzureName: aadGroupName,
			Privileges: []string{
				"CREATE USER",
				"PROCESS",
			},
			AADUser: &mysqlv1.AADUserSpec{
				Alias:               aadGroupName,
				ServerAdminUsername: to.Value(admin.Spec.Login),
			},
		},
	}
	tc.CreateResourcesAndWait(user2)

	// Confirm that we have the right privs on the actual server
	serverPrivs, err = mysqlutil.GetUserServerPrivileges(tc.Ctx, conn, aadGroupName, "")
	tc.Expect(err).ToNot(HaveOccurred())
	tc.Expect(serverPrivs).To(Equal(set.Make[string](user.Spec.Privileges...)))

	tc.DeleteResourcesAndWait(user, user2)
}

func readEnv() (string, string, error) {
	identityName := os.Getenv(AzureManagedIdentityNameVar)
	if identityName == "" {
		return "", "", errors.Errorf("required environment variable %q was not supplied", AzureManagedIdentityNameVar)
	}

	identityResourceGroup := os.Getenv(AzureManagedIdentityResourceGroupVar)
	if identityResourceGroup == "" {
		return "", "", errors.Errorf("required environment variable %q was not supplied", AzureManagedIdentityResourceGroupVar)
	}

	return identityName, identityResourceGroup, nil
}
