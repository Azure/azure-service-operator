/*
Copyright (c) Microsoft Corporation.
Licensed under the MIT license.
*/

package test

import (
	"testing"

	. "github.com/onsi/gomega"

	_ "github.com/microsoft/go-mssqldb"

	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1api20181130"
	resources "github.com/Azure/azure-service-operator/v2/api/resources/v1api20200601"
	azuresqlv1 "github.com/Azure/azure-service-operator/v2/api/sql/v1"
	sql "github.com/Azure/azure-service-operator/v2/api/sql/v1api20211101"
	"github.com/Azure/azure-service-operator/v2/internal/testcommon"
	"github.com/Azure/azure-service-operator/v2/internal/util/to"
)

func Test_AzureSQL_AADUser(t *testing.T) {
	t.Parallel()
	t.Skip("Cannot run in CI currently due to AAD tenant restrictions on CI subscription")

	tc := globalTestContext.ForTest(t)
	tc.AzureRegion = to.Ptr("eastus")

	rg := tc.CreateTestResourceGroupAndWait()

	server := newAzureSQLServerWithAADAdmin(tc, rg)
	database := newAzureSQLServerDatabase(tc, server)
	firewallRule := newSQLServerOpenFirewallRule(tc, server)

	tc.CreateResourcesAndWait(server, database, firewallRule)
	tc.Expect(server.Status.FullyQualifiedDomainName).ToNot(BeNil())

	tc.RunSubtests(
		testcommon.Subtest{
			Name: "AzureSQL AAD User CRUD",
			Test: func(testContext *testcommon.KubePerTestContext) {
				AzureSQL_AADUser_CRUD(testContext, server, database, rg)
			},
		},
		testcommon.Subtest{
			Name: "AzureSQL AAD User With Alias",
			Test: func(testContext *testcommon.KubePerTestContext) {
				AzureSQL_AADUser_WithAlias(testContext, server, database, rg)
			},
		},
	)
}

func AzureSQL_AADUser_CRUD(
	tc *testcommon.KubePerTestContext,
	server *sql.Server,
	database *sql.ServersDatabase,
	rg *resources.ResourceGroup,
) {
	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}
	tc.CreateResourceAndWait(mi)

	user := &azuresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName(mi.Name),
		Spec: azuresqlv1.UserSpec{
			Owner: testcommon.AsOwner(database),
			Roles: []string{"db_datareader"},
			AADUser: &azuresqlv1.AADUserSpec{
				ServerAdminUsername: "yourAADAdminName",
			},
		},
	}
	tc.CreateResourceAndWait(user)
	tc.Expect(user.Status.Conditions).ToNot(BeEmpty())

	old := user.DeepCopy()
	user.Spec.Roles = []string{"db_datareader", "db_datawriter"}
	tc.PatchResourceAndWait(old, user)

	originalUser := user.DeepCopy()

	old = user.DeepCopy()
	user.Spec.AADUser = nil
	user.Spec.LocalUser = &azuresqlv1.LocalUserSpec{ServerAdminUsername: "someadmin"}
	err := tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("cannot change from AAD user to local user"))

	user = originalUser.DeepCopy()
	tc.DeleteResourceAndWait(user)
}

func AzureSQL_AADUser_WithAlias(
	tc *testcommon.KubePerTestContext,
	server *sql.Server,
	database *sql.ServersDatabase,
	rg *resources.ResourceGroup,
) {
	alias := "myaliaseduser"

	mi := &managedidentity.UserAssignedIdentity{
		ObjectMeta: tc.MakeObjectMeta("mi-long-name"),
		Spec: managedidentity.UserAssignedIdentity_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
		},
	}
	tc.CreateResourceAndWait(mi)

	user := &azuresqlv1.User{
		ObjectMeta: tc.MakeObjectMetaWithName("aaduser-alias"),
		Spec: azuresqlv1.UserSpec{
			AzureName: mi.Name,
			Owner:     testcommon.AsOwner(database),
			Roles:     []string{"db_datareader"},
			AADUser: &azuresqlv1.AADUserSpec{
				Alias:               alias,
				ServerAdminUsername: "yourAADAdminName",
			},
		},
	}
	tc.CreateResourceAndWait(user)

	originalUser := user.DeepCopy()

	old := user.DeepCopy()
	user.Spec.AADUser.Alias = "adifferentalias"
	err := tc.PatchAndExpectError(old, user)
	tc.Expect(err).To(HaveOccurred())
	tc.Expect(err.Error()).To(ContainSubstring("cannot change AAD user 'alias'"))

	user = originalUser.DeepCopy()
	tc.DeleteResourceAndWait(user)
}

func newAzureSQLServerWithAADAdmin(tc *testcommon.KubePerTestContext, rg *resources.ResourceGroup) *sql.Server {
	server := &sql.Server{
		ObjectMeta: tc.MakeObjectMeta("sqlserver"),
		Spec: sql.Server_Spec{
			Location: tc.AzureRegion,
			Owner:    testcommon.AsOwner(rg),
			Version:  to.Ptr("12.0"),
		},
	}
	return server
}
