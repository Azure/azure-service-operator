// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver azuresqlservercombined

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/stretchr/testify/assert"

	helpers "github.com/Azure/azure-service-operator/pkg/helpers"
	kvsecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
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
	sqlServerInstance := v1beta1.NewAzureSQLServer(sqlServerNamespacedName, rgName, rgLocation)

	// Send request for 2nd server (failovergroup test) before waiting on first server
	sqlServerInstance2 := v1beta1.NewAzureSQLServer(sqlServerNamespacedName2, rgName, rgLocation2)

	// create and wait
	RequireInstance(ctx, t, tc, sqlServerInstance)

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
	var sqlDatabaseInstance *v1beta1.AzureSqlDatabase

	sqlFirewallRuleNamespacedNameLocal := types.NamespacedName{
		Name:      GenerateTestResourceNameWithRandom("sqlfwr-local", 10),
		Namespace: "default",
	}
	sqlFirewallRuleNamespacedNameRemote := types.NamespacedName{
		Name:      GenerateTestResourceNameWithRandom("sqlfwr-remote", 10),
		Namespace: "default",
	}

	var sqlFirewallRuleInstanceLocal *v1beta1.AzureSqlFirewallRule
	var sqlFirewallRuleInstanceRemote *v1beta1.AzureSqlFirewallRule

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
			sqlDatabaseInstance = &v1beta1.AzureSqlDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlDatabaseName,
					Namespace: "default",
				},
				Spec: v1beta1.AzureSqlDatabaseSpec{
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
			sqlFirewallRuleInstanceLocal = v1beta1.NewAzureSQLFirewallRule(
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
			sqlFirewallRuleInstanceRemote = v1beta1.NewAzureSQLFirewallRule(
				sqlFirewallRuleNamespacedNameRemote,
				rgName,
				sqlServerName,
				"0.0.0.0",
				"0.0.0.0",
			)

			EnsureInstance(ctx, t, tc, sqlFirewallRuleInstanceRemote)
		})

		// Create VNet and VNetRules -----
		t.Run("run subtest to test VNet Rule in primary server", func(t *testing.T) {
			t.Parallel()
			RunAzureSqlVNetRuleHappyPath(t, sqlServerName, rgLocation)
		})

	})

	var sqlUser *azurev1alpha1.AzureSQLUser
	var kvSqlUser1 *azurev1alpha1.AzureSQLUser
	var kvSqlUser2 *azurev1alpha1.AzureSQLUser

	// run sub tests that require 2 servers or have to be run after rolladmincreds test ------------------
	t.Run("group2", func(t *testing.T) {

		t.Run("set up user in first db", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "sql-test-user" + helpers.RandomString(10)
			roles := []string{"db_owner"}
			keyVaultSecretFormats := []string{"adonet"}

			sqlUser = &azurev1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSQLUserSpec{
					Server:                sqlServerName,
					DbName:                sqlDatabaseName,
					ResourceGroup:         rgName,
					Roles:                 roles,
					KeyVaultSecretFormats: keyVaultSecretFormats,
				},
			}

			EnsureInstance(ctx, t, tc, sqlUser)

			// verify user's secret has been created
			// this test suite defaults to Kube Secrets. They do not support keyvault-specific config but the spec is passed anyway
			// to verify that passing them does not break the service
			assert.Eventually(func() bool {
				key := types.NamespacedName{Name: sqlUser.ObjectMeta.Name, Namespace: sqlUser.ObjectMeta.Namespace}
				var secrets, _ = tc.secretClient.Get(ctx, key)

				return strings.Contains(string(secrets["azureSqlDatabaseName"]), sqlDatabaseName)
			}, tc.timeoutFast, tc.retry, "wait for secret store to show azure sql user credentials")

			t.Log(sqlUser.Status)
		})

		t.Run("set up user in first db with custom keyvault", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "sql-test-user" + helpers.RandomString(10)
			roles := []string{"db_owner"}

			// This test will attempt to persist secrets to the KV that was instantiated as part of the test suite
			keyVaultName := tc.keyvaultName

			kvSqlUser1 = &azurev1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSQLUserSpec{
					Server:                 sqlServerName,
					DbName:                 sqlDatabaseName,
					ResourceGroup:          rgName,
					Roles:                  roles,
					KeyVaultToStoreSecrets: keyVaultName,
				},
			}

			EnsureInstance(ctx, t, tc, kvSqlUser1)

			// Check that the user's secret is in the keyvault
			keyVaultSecretClient := kvsecrets.New(keyVaultName)

			assert.Eventually(func() bool {
				keyNamespace := "azuresqluser-" + sqlServerName + "-" + sqlDatabaseName
				key := types.NamespacedName{Name: kvSqlUser1.ObjectMeta.Name, Namespace: keyNamespace}
				var secrets, _ = keyVaultSecretClient.Get(ctx, key)

				return strings.Contains(string(secrets["azureSqlDatabaseName"]), sqlDatabaseName)
			}, tc.timeoutFast, tc.retry, "wait for keyvault to show azure sql user credentials")

			t.Log(kvSqlUser1.Status)
		})

		t.Run("set up user in first db with custom keyvault and custom formatting", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "sql-test-user" + helpers.RandomString(10)
			roles := []string{"db_owner"}
			formats := []string{"adonet"}

			// This test will attempt to persist secrets to the KV that was instantiated as part of the test suite
			keyVaultName := tc.keyvaultName

			kvSqlUser2 = &azurev1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: azurev1alpha1.AzureSQLUserSpec{
					Server:                 sqlServerName,
					DbName:                 sqlDatabaseName,
					ResourceGroup:          rgName,
					Roles:                  roles,
					KeyVaultToStoreSecrets: keyVaultName,
					KeyVaultSecretFormats:  formats,
				},
			}

			EnsureInstance(ctx, t, tc, kvSqlUser2)

			// Check that the user's secret is in the keyvault
			keyVaultSecretClient := kvsecrets.New(keyVaultName)

			assert.Eventually(func() bool {
				keyNamespace := "azuresqluser-" + sqlServerName + "-" + sqlDatabaseName
				keyName := kvSqlUser2.ObjectMeta.Name + "-adonet"
				key := types.NamespacedName{Name: keyName, Namespace: keyNamespace}
				var secrets, _ = keyVaultSecretClient.Get(ctx, key)

				return len(string(secrets[keyNamespace+"-"+keyName])) > 0
			}, tc.timeoutFast, tc.retry, "wait for keyvault to show azure sql user credentials with custom formats")

			t.Log(kvSqlUser2.Status)
		})
	})

	t.Run("deploy sql action and roll user credentials", func(t *testing.T) {
		keyNamespace := "azuresqluser-" + sqlServerName + "-" + sqlDatabaseName
		key := types.NamespacedName{Name: kvSqlUser1.ObjectMeta.Name, Namespace: keyNamespace}

		keyVaultName := tc.keyvaultName
		keyVaultSecretClient := kvsecrets.New(keyVaultName)
		var oldSecret, _ = keyVaultSecretClient.Get(ctx, key)

		sqlActionName := GenerateTestResourceNameWithRandom("azuresqlaction-dev", 10)
		sqlActionInstance := &azurev1alpha1.AzureSqlAction{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlActionName,
				Namespace: "default",
			},
			Spec: azurev1alpha1.AzureSqlActionSpec{
				ResourceGroup:      rgName,
				ServerName:         sqlServerName,
				ActionName:         "rollusercreds",
				DbName:             sqlDatabaseName,
				DbUser:             kvSqlUser1.ObjectMeta.Name,
				UserSecretKeyVault: keyVaultName,
			},
		}

		err := tc.k8sClient.Create(ctx, sqlActionInstance)
		assert.Equal(nil, err, "create sqlaction in k8s")

		sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}

		assert.Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlActionInstanceNamespacedName, sqlActionInstance)
			return sqlActionInstance.Status.Provisioned
		}, tc.timeout, tc.retry, "wait for sql action to be submitted")

		var newSecret, _ = keyVaultSecretClient.Get(ctx, key)

		assert.NotEqual(oldSecret["password"], newSecret["password"], "password should have been updated")
		assert.Equal(oldSecret["username"], newSecret["username"], "usernames should be the same")
	})

	var sqlFailoverGroupInstance *v1beta1.AzureSqlFailoverGroup
	sqlFailoverGroupName := GenerateTestResourceNameWithRandom("sqlfog-dev", 10)

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	t.Run("group3", func(t *testing.T) {
		t.Run("delete db users and ensure that their secrets have been cleaned up", func(t *testing.T) {
			EnsureDelete(ctx, t, tc, sqlUser)
			EnsureDelete(ctx, t, tc, kvSqlUser1)
			EnsureDelete(ctx, t, tc, kvSqlUser2)

			// Check that the user's secret is in the keyvault
			keyVaultSecretClient := kvsecrets.New(tc.keyvaultName)

			assert.Eventually(func() bool {
				key := types.NamespacedName{Name: sqlUser.ObjectMeta.Name, Namespace: sqlUser.ObjectMeta.Namespace}
				var _, err = tc.secretClient.Get(ctx, key)

				// Once the secret is gone, the Kube secret client will return an error
				return err != nil && strings.Contains(err.Error(), "not found")
			}, tc.timeoutFast, tc.retry, "wait for the azuresqluser kube secret to be deleted")

			assert.Eventually(func() bool {
				keyNamespace := "azuresqluser-" + sqlServerName + "-" + sqlDatabaseName
				key := types.NamespacedName{Name: kvSqlUser1.ObjectMeta.Name, Namespace: keyNamespace}
				var _, err = keyVaultSecretClient.Get(ctx, key)

				// Once the secret is gone, the KV secret client will return an err
				return err != nil && strings.Contains(err.Error(), "secret does not exist")
			}, tc.timeoutFast, tc.retry, "wait for the azuresqluser keyvault secret to be deleted")

			assert.Eventually(func() bool {
				keyNamespace := "azuresqluser-" + sqlServerName + "-" + sqlDatabaseName
				keyName := kvSqlUser2.ObjectMeta.Name + "-adonet"
				key := types.NamespacedName{Name: keyName, Namespace: keyNamespace}
				var _, err = keyVaultSecretClient.Get(ctx, key)

				// Once the secret is gone, the KV secret client will return an err
				return err != nil && strings.Contains(err.Error(), "secret does not exist")
			}, tc.timeoutFast, tc.retry, "wait for the azuresqluser custom formatted keyvault secret to be deleted")
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
			sqlFailoverGroupInstance = &v1beta1.AzureSqlFailoverGroup{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlFailoverGroupNamespacedName.Name,
					Namespace: sqlFailoverGroupNamespacedName.Namespace,
				},
				Spec: v1beta1.AzureSqlFailoverGroupSpec{
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
