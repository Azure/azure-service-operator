/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	. "github.com/onsi/gomega"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	_ "github.com/microsoft/go-mssqldb"

	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/set"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	azuresqlutil "github.com/Azure/azure-service-operator/v2/internal/util/azuresql"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

const (
	azureSQLTokenScope       = "https://database.windows.net/.default" // #nosec G101
	azureTestIdentityNameVar = "AZURE_TEST_IDENTITY_NAME"
)

func Test_AzureSQL_Combined(t *testing.T) {
	t.Parallel()
	tc := globalTestContext.ForTest(t)

	// Use a different region where we have quota
	tc.AzureRegion = to.Ptr("australiaeast")

	rg := tc.CreateTestResourceGroupAndWait()

	admin, err := azureSQLAdminForTest(tc)
	tc.Expect(err).ToNot(HaveOccurred())

	server := newAzureSQLAADServer(tc, rg, admin)
	database := newAzureSQLServerDatabase(tc, server)
	firewallRule := newSQLServerOpenFirewallRule(tc, server)

	tc.CreateResourcesAndWait(server, database, firewallRule)

	tc.Expect(server.Status.FullyQualifiedDomainName).ToNot(BeNil())
	fqdn := *server.Status.FullyQualifiedDomainName

	// Ensure that firewall rule access has worked. It can take up to 5 minutes to take effect
	tc.G.Eventually(
		func() error {
			db, err := azuresqlutil.ConnectToDBUsingAAD(
				tc.Ctx,
				fqdn,
				database.AzureName(),
				azuresqlutil.ServerPort,
				admin.tokenProvider,
			)
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
				AzureSQL_User_Helpers(testContext, fqdn, database.AzureName(), admin.tokenProvider)
			},
		},
	)
}

func AzureSQL_User_Helpers(tc *testcommon.KubePerTestContext, fqdn string, database string, tokenProvider func() (string, error)) {
	// Connect to the DB
	ctx := tc.Ctx

	db, err := azuresqlutil.ConnectToDBUsingAAD(
		ctx,
		fqdn,
		database,
		azuresqlutil.ServerPort,
		tokenProvider,
	)
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

type azureSQLAdminIdentity struct {
	login         string
	objectID      string
	tenantID      string
	principalType sql.ServerExternalAdministrator_PrincipalType
	tokenProvider func() (string, error)
}

type azureSQLTokenClaims struct {
	ObjectID          string `json:"oid"`
	TenantID          string `json:"tid"`
	IdentityType      string `json:"idtyp"`
	Scopes            string `json:"scp"`
	Name              string `json:"name"`
	PreferredUsername string `json:"preferred_username"`
	UPN               string `json:"upn"`
	UniqueName        string `json:"unique_name"`
}

func azureSQLAdminForTest(tc *testcommon.KubePerTestContext) (azureSQLAdminIdentity, error) {
	credential := tc.AzureClient.Creds()
	tokenProvider := func() (string, error) {
		token, err := credential.GetToken(tc.Ctx, policy.TokenRequestOptions{Scopes: []string{azureSQLTokenScope}})
		if err != nil {
			return "", err
		}

		return token.Token, nil
	}

	token, err := tokenProvider()
	if err != nil {
		return azureSQLAdminIdentity{}, fmt.Errorf("getting Azure SQL access token: %w", err)
	}

	claims, err := parseAzureSQLTokenClaims(token)
	if err != nil {
		return azureSQLAdminIdentity{}, err
	}

	if claims.ObjectID == "" || claims.TenantID == "" {
		return azureSQLAdminIdentity{}, fmt.Errorf("Azure SQL access token is missing oid or tid claim")
	}

	// If AZURE_TEST_IDENTITY_NAME is set, we're in CI and we have a managed identity
	if login := os.Getenv(azureTestIdentityNameVar); login != "" {
		return azureSQLAdminIdentity{
			login:         login,
			objectID:      claims.ObjectID,
			tenantID:      claims.TenantID,
			principalType: sql.ServerExternalAdministrator_PrincipalType_Application,
			tokenProvider: tokenProvider,
		}, nil
	}

	userLogin := firstNonEmpty(claims.PreferredUsername, claims.UPN, claims.UniqueName)
	// User tokens can also contain appid/azp, identifying the client application that acquired the token.
	// Prefer idtyp, with scp and username claims as fallbacks; scp is emitted only for user tokens.
	isUser := strings.EqualFold(claims.IdentityType, "user") || claims.Scopes != "" || userLogin != ""
	if !isUser {
		return azureSQLAdminIdentity{}, fmt.Errorf("Azure SQL application identity requires environment variable %q", azureTestIdentityNameVar)
	}

	login := firstNonEmpty(userLogin, claims.Name)
	if login == "" {
		return azureSQLAdminIdentity{}, fmt.Errorf("Azure SQL user access token is missing a login name claim")
	}

	return azureSQLAdminIdentity{
		login:         login,
		objectID:      claims.ObjectID,
		tenantID:      claims.TenantID,
		principalType: sql.ServerExternalAdministrator_PrincipalType_User,
		tokenProvider: tokenProvider,
	}, nil
}

func parseAzureSQLTokenClaims(token string) (azureSQLTokenClaims, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return azureSQLTokenClaims{}, fmt.Errorf("Azure SQL access token is not a JWT")
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return azureSQLTokenClaims{}, fmt.Errorf("decoding Azure SQL access token claims: %w", err)
	}

	var claims azureSQLTokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return azureSQLTokenClaims{}, fmt.Errorf("parsing Azure SQL access token claims: %w", err)
	}

	return claims, nil
}

func firstNonEmpty(values ...string) string {
	for _, value := range values {
		if value != "" {
			return value
		}
	}

	return ""
}

func newAzureSQLAADServer(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup, admin azureSQLAdminIdentity) *sql.Server {
	server := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserver"),
		Spec: sql.Server_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Administrators: &sql.ServerExternalAdministrator{
				AdministratorType:         to.Ptr(sql.ServerExternalAdministrator_AdministratorType_ActiveDirectory),
				PrincipalType:             to.Ptr(admin.principalType),
				AzureADOnlyAuthentication: to.Ptr(true),
				Login:                     to.Ptr(admin.login),
				Sid:                       to.Ptr(admin.objectID),
				TenantId:                  to.Ptr(admin.tenantID),
			},
			Version: to.Ptr("12.0"),
		},
	}

	return server
}

func newAzureSQLServerDatabase(tc *testcommon.KubePerTestContext, server *sql.Server) *sql.ServersDatabase {
	db := &sql.ServersDatabase{
		ObjectMeta: tc.MakeObjectMeta("db"),
		Spec: sql.ServersDatabase_Spec{
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
		Spec: sql.ServersFirewallRule_Spec{
			Owner:          testcommon.AsOwner(server),
			StartIpAddress: to.Ptr("0.0.0.0"),
			EndIpAddress:   to.Ptr("255.255.255.255"),
		},
	}

	return firewall
}
