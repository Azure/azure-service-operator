// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver azuresqlservercombined

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlServerCombinedHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	sqlServerName := GenerateTestResourceNameWithRandom("sqlserver", 10)
	rgLocation := "westus2"
	rgLocation2 := "southcentralus"
	sqlServerTwoName := GenerateTestResourceNameWithRandom("sqlserver-two", 10)

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}
	sqlServerNamespacedName2 := types.NamespacedName{Name: sqlServerTwoName, Namespace: "default"}

	// Create the SqlServer object and expect the Reconcile to be created
	sqlServerInstance := azurev1alpha1.NewAzureSQLServer(sqlServerNamespacedName, rgName, rgLocation)

	// Send request for 2nd server (failovergroup test) before waiting on first server
	sqlServerInstance2 := azurev1alpha1.NewAzureSQLServer(sqlServerNamespacedName2, rgName, rgLocation2)

	// create and wait
	EnsureInstance(ctx, t, tc, sqlServerInstance)

	//verify secret exists in k8s for server 1 ---------------------------------
	secret := &v1.Secret{}
	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, types.NamespacedName{Name: sqlServerName, Namespace: sqlServerInstance.Namespace}, secret)

		if err == nil {
			if (secret.ObjectMeta.Name == sqlServerName) && (secret.ObjectMeta.Namespace == sqlServerInstance.Namespace) {
				return true
			}
		}
		return false
	}, tc.timeoutFast, tc.retry, "wait for server to have secret")

	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	var sqlDatabaseInstance *azurev1alpha1.AzureSqlDatabase

	sqlFirewallRuleNamespacedNameLocal := types.NamespacedName{
		Name:      GenerateTestResourceNameWithRandom("sqlfwr-local", 10),
		Namespace: "default",
	}
	sqlFirewallRuleNamespacedNameRemote := types.NamespacedName{
		Name:      GenerateTestResourceNameWithRandom("sqlfwr-remote", 10),
		Namespace: "default",
	}

	var sqlFirewallRuleInstanceLocal *azurev1alpha1.AzureSqlFirewallRule
	var sqlFirewallRuleInstanceRemote *azurev1alpha1.AzureSqlFirewallRule

	// run sub tests that require 1 sql server ----------------------------------
	t.Run("group1", func(t *testing.T) {
		t.Run("sub test for actions", func(t *testing.T) {
			t.Parallel()
			RunSQLActionHappy(t, sqlServerName)
		})

		// Wait for 2nd sql server to resolve
		t.Run("set up secondary server", func(t *testing.T) {
			t.Parallel()
			EnsureInstance(ctx, t, tc, sqlServerInstance2)
		})

		// Create a database in the new server
		t.Run("set up database in primary server", func(t *testing.T) {
			t.Parallel()

			// Create the SqlDatabase object and expect the Reconcile to be created
			sqlDatabaseInstance = &azurev1alpha1.AzureSqlDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlDatabaseName,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSqlDatabaseSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
					Server:        sqlServerName,
					Edition:       0,
				},
			}

			EnsureInstance(ctx, t, tc, sqlDatabaseInstance)
		})

		// Create FirewallRules ---------------------------------------

		t.Run("set up wide range firewall rule in primary server", func(t *testing.T) {
			t.Parallel()

			// Create the SqlFirewallRule object and expect the Reconcile to be created
			sqlFirewallRuleInstanceLocal = azurev1alpha1.NewAzureSQLFirewallRule(
				sqlFirewallRuleNamespacedNameLocal,
				rgName,
				sqlServerName,
				"1.1.1.1",
				"255.255.255.255",
			)

			EnsureInstance(ctx, t, tc, sqlFirewallRuleInstanceLocal)
		})

		t.Run("set up azure only firewall rule in primary server", func(t *testing.T) {
			t.Parallel()

			// Create the SqlFirewallRule object and expect the Reconcile to be created
			sqlFirewallRuleInstanceRemote = azurev1alpha1.NewAzureSQLFirewallRule(
				sqlFirewallRuleNamespacedNameRemote,
				rgName,
				sqlServerName,
				"0.0.0.0",
				"0.0.0.0",
			)

			EnsureInstance(ctx, t, tc, sqlFirewallRuleInstanceRemote)
		})

	})

	var sqlUser *azurev1alpha1.AzureSQLUser

	// run sub tests that require 2 servers or have to be run after rollcreds test ------------------
	t.Run("group2", func(t *testing.T) {

		t.Run("set up user in first db", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "sql-test-user" + helpers.RandomString(10)
			roles := []string{"db_owner"}

			sqlUser = &azurev1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSQLUserSpec{
					Server:        sqlServerName,
					DbName:        sqlDatabaseName,
					ResourceGroup: rgName,
					Roles:         roles,
				},
			}

			EnsureInstance(ctx, t, tc, sqlUser)
			t.Log(sqlUser.Status)
		})
	})

	var sqlFailoverGroupInstance *azurev1alpha1.AzureSqlFailoverGroup
	sqlFailoverGroupName := GenerateTestResourceNameWithRandom("sqlfog-dev", 10)

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	t.Run("group3", func(t *testing.T) {
		t.Run("delete db user", func(t *testing.T) {
			EnsureDelete(ctx, t, tc, sqlUser)
		})

		t.Run("delete local firewallrule", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlFirewallRuleInstanceLocal)
		})

		t.Run("delete remote firewallrule", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlFirewallRuleInstanceRemote)
		})

		t.Run("create failovergroup", func(t *testing.T) {
			t.Parallel()

			// Create the SqlFailoverGroup object and expect the Reconcile to be created
			sqlFailoverGroupInstance = &azurev1alpha1.AzureSqlFailoverGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlFailoverGroupNamespacedName.Name,
					Namespace: sqlFailoverGroupNamespacedName.Namespace,
				},
				Spec: azurev1alpha1.AzureSqlFailoverGroupSpec{
					Location:                     rgLocation,
					ResourceGroup:                rgName,
					Server:                       sqlServerName,
					FailoverPolicy:               "automatic",
					FailoverGracePeriod:          30,
					SecondaryServer:              sqlServerTwoName,
					SecondaryServerResourceGroup: rgName,
					DatabaseList:                 []string{sqlDatabaseName},
				},
			}

			EnsureInstance(ctx, t, tc, sqlFailoverGroupInstance)

			// verify secret has been created
			assert.Eventually(func() bool {
				var secrets, _ = tc.secretClient.Get(ctx, sqlFailoverGroupNamespacedName)

				return strings.Contains(string(secrets["azureSqlPrimaryServer"]), sqlServerName)
			}, tc.timeout, tc.retry, "wait for secret store to show failovergroup server names  ")

		})

	})

	t.Run("group4", func(t *testing.T) {

		t.Run("delete failovergroup", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlFailoverGroupInstance)
		})
		t.Run("delete db", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlDatabaseInstance)
		})

	})

	t.Run("group5", func(t *testing.T) {

		t.Run("delete sqlServerInstance", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlServerInstance)
		})
		t.Run("delete sqlServerInstance2", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlServerInstance2)
		})

	})

}
