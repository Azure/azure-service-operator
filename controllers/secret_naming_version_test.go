// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all
// +build all

package controllers

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/go-autorest/autorest/to"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	kvsecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
)

// NOTE: Tests in this file are intended to be run twice, once in the v1 secret naming mode
// and once in the v2 secret naming mode. In some instances they parallel tests
// which don't assert on secret name or shape.
// The naming of the tests is important! It's how the makefile knows to run just these tests
// in the old V1 mode.
// Not every resource is included here, only resources which for legacy reasons in the old v1
// secret naming mode had secret naming schemes which were different than the standard.

// This test intentionally mirrors TestAppInsightsController in appinsights_controller_test.go but is
// focused on ensuring the shape of the saved secret is correct
func TestAppInsights_SecretNamedCorrectly(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	assert := assert.New(t)

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	appInsightsName := GenerateTestResourceNameWithRandom("appinsights", 6)

	// Create an instance of Azure AppInsights
	instance := &v1alpha1.AppInsights{
		ObjectMeta: metav1.ObjectMeta{
			Name:      appInsightsName,
			Namespace: "default",
		},
		Spec: v1alpha1.AppInsightsSpec{
			Kind:            "web",
			Location:        rgLocation,
			ResourceGroup:   rgName,
			ApplicationType: "other",
		},
	}

	EnsureInstance(ctx, t, tc, instance)

	// Make sure the secret is created
	var keyName string
	if tc.secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		keyName = fmt.Sprintf("appinsights-%s-%s", instance.Spec.ResourceGroup, instance.Name)
	} else {
		keyName = instance.Name
	}

	secretKey := secrets.SecretKey{Name: keyName, Namespace: instance.Namespace, Kind: "appinsights"}

	// Secret is created after reconciliation is in provisioned state
	var secret map[string][]byte
	assert.Eventually(func() bool {
		var err error
		secret, err = tc.secretClient.Get(ctx, secretKey)
		return err == nil
	}, tc.timeoutFast, tc.retry, "should be able to get secret")

	assert.NotEmpty(string(secret["instrumentationKey"]))

	EnsureDelete(ctx, t, tc, instance)
}

// This test intentionally mirrors TestStorageAccountController in storage_controller_test.go but is
// focused on ensuring the shape of the saved secret is correct
func TestStorageAccount_SecretNamedCorrectly(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	assert := assert.New(t)

	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	storageName := "storageacct" + helpers.RandomString(6)

	instance := &v1alpha1.StorageAccount{
		ObjectMeta: metav1.ObjectMeta{
			Name:      storageName,
			Namespace: "default",
		},
		Spec: v1alpha1.StorageAccountSpec{
			Kind:          "BlobStorage",
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: v1alpha1.StorageAccountSku{
				Name: "Standard_LRS",
			},
			AccessTier:             "Hot",
			EnableHTTPSTrafficOnly: to.BoolPtr(true),
		},
	}

	EnsureInstance(ctx, t, tc, instance)

	// Make sure the secret is created
	var keyName string
	if tc.secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		keyName = fmt.Sprintf("storageaccount-%s-%s", instance.Spec.ResourceGroup, instance.Name)
	} else {
		keyName = instance.Name
	}
	secretKey := secrets.SecretKey{Name: keyName, Namespace: instance.Namespace, Kind: "storageaccount"}
	secret, err := tc.secretClient.Get(ctx, secretKey)
	assert.NoError(err)

	assert.NotEmpty(string(secret["StorageAccountName"]))
	assert.NotEmpty(string(secret["connectionString0"]))
	assert.NotEmpty(string(secret["key0"]))
	assert.NotEmpty(string(secret["connectionString1"]))
	assert.NotEmpty(string(secret["key1"]))

	EnsureDelete(ctx, t, tc, instance)
}

func assertSQLServerAdminSecretCreated(ctx context.Context, t *testing.T, sqlServerInstance *v1beta1.AzureSqlServer) {
	if len(sqlServerInstance.Spec.KeyVaultToStoreSecrets) == 0 {
		secret := &v1.Secret{}
		assert.Eventually(t, func() bool {
			var expectedServerSecretName string
			if tc.secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
				expectedServerSecretName = sqlServerInstance.Name
			} else {
				expectedServerSecretName = fmt.Sprintf("%s-%s", "azuresqlserver", sqlServerInstance.Name)
			}
			err := tc.k8sClient.Get(ctx, types.NamespacedName{Namespace: sqlServerInstance.Namespace, Name: expectedServerSecretName}, secret)
			if err != nil {
				return false
			}
			return secret.Name == expectedServerSecretName && secret.Namespace == sqlServerInstance.Namespace
		}, tc.timeoutFast, tc.retry, "wait for server to have secret")
	} else {
		// Check that the user's secret is in the keyvault
		keyVaultSecretClient := kvsecrets.New(
			sqlServerInstance.Spec.KeyVaultToStoreSecrets,
			config.GlobalCredentials(),
			config.SecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())

		assert.Eventually(t, func() bool {
			expectedSecretName := makeSQLServerKeyVaultSecretName(sqlServerInstance)
			result, err := keyVaultSecretClient.KeyVaultClient.GetSecret(ctx, kvsecrets.GetVaultsURL(sqlServerInstance.Spec.KeyVaultToStoreSecrets), expectedSecretName, "")
			if err != nil {
				return false
			}
			secrets := *result.Value

			return len(string(secrets)) > 0
		}, tc.timeoutFast, tc.retry, "wait for keyvault to have secret for server")
	}
}

func TestAzureSqlServerAndUser_SecretNamedCorrectly(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	sqlServerName := GenerateTestResourceNameWithRandom("sqlserver1", 10)
	rgLocation := "westus2"

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}
	sqlServerInstance := v1beta1.NewAzureSQLServer(sqlServerNamespacedName, rgName, rgLocation)

	// create and wait
	RequireInstance(ctx, t, tc, sqlServerInstance)

	//verify secret exists in k8s for server 1
	assertSQLServerAdminSecretCreated(ctx, t, sqlServerInstance)

	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &v1beta1.AzureSqlDatabase{
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

	// Create firewall rule
	sqlFirewallRuleNamespacedName := types.NamespacedName{
		Name:      GenerateTestResourceNameWithRandom("sqlfwr-local1", 10),
		Namespace: "default",
	}
	sqlFirewallRule := v1beta1.NewAzureSQLFirewallRule(
		sqlFirewallRuleNamespacedName,
		rgName,
		sqlServerName,
		"1.1.1.1",
		"255.255.255.255",
	)
	EnsureInstance(ctx, t, tc, sqlFirewallRule)

	t.Run("sub test for actions", func(t *testing.T) {
		RunSQLActionHappy(t, sqlServerName)
	})

	var sqlUser *v1alpha1.AzureSQLUser
	var kvSqlUser1 *v1alpha1.AzureSQLUser
	var kvSqlUser2 *v1alpha1.AzureSQLUser

	// Run user subtests
	t.Run("group2", func(t *testing.T) {

		t.Run("set up user in db", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "user1" + helpers.RandomString(10)
			roles := []string{"db_owner"}
			keyVaultSecretFormats := []string{"adonet"}

			sqlUser = &v1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: v1alpha1.AzureSQLUserSpec{
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
				key := secrets.SecretKey{Name: sqlUser.Name, Namespace: sqlUser.Namespace, Kind: "azuresqluser"}
				var secrets, err = tc.secretClient.Get(ctx, key)

				return err == nil && strings.Contains(string(secrets["azureSqlDatabaseName"]), sqlDatabaseName)
			}, tc.timeoutFast, tc.retry, "wait for secret store to show azure sql user credentials")

			t.Log(sqlUser.Status)
		})

		t.Run("set up user in db with custom keyvault", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "user2" + helpers.RandomString(10)
			roles := []string{"db_owner"}

			// This test will attempt to persist secrets to the KV that was instantiated as part of the test suite
			kvSqlUser1 = &v1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: v1alpha1.AzureSQLUserSpec{
					Server:                 sqlServerName,
					DbName:                 sqlDatabaseName,
					ResourceGroup:          rgName,
					Roles:                  roles,
					KeyVaultToStoreSecrets: tc.keyvaultName,
				},
			}

			EnsureInstance(ctx, t, tc, kvSqlUser1)

			// Check that the user's secret is in the keyvault
			keyVaultSecretClient := kvsecrets.New(
				tc.keyvaultName,
				config.GlobalCredentials(),
				config.SecretNamingVersion(),
				config.PurgeDeletedKeyVaultSecrets(),
				config.RecoverSoftDeletedKeyVaultSecrets())

			assert.Eventually(func() bool {
				key := makeSQLUserSecretKey(keyVaultSecretClient, kvSqlUser1)
				var secrets, _ = keyVaultSecretClient.Get(ctx, key)

				return strings.Contains(string(secrets["azureSqlDatabaseName"]), sqlDatabaseName)
			}, tc.timeoutFast, tc.retry, "wait for keyvault to show azure sql user credentials")

			t.Log(kvSqlUser1.Status)
		})

		t.Run("set up user in db with custom keyvault and custom formatting", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "user3" + helpers.RandomString(10)
			roles := []string{"db_owner"}
			formats := []string{"adonet"}

			// This test will attempt to persist secrets to the KV that was instantiated as part of the test suite
			kvSqlUser2 = &v1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: v1alpha1.AzureSQLUserSpec{
					Server:                 sqlServerName,
					DbName:                 sqlDatabaseName,
					ResourceGroup:          rgName,
					Roles:                  roles,
					KeyVaultToStoreSecrets: tc.keyvaultName,
					KeyVaultSecretFormats:  formats,
				},
			}

			EnsureInstance(ctx, t, tc, kvSqlUser2)

			// Check that the user's secret is in the keyvault
			keyVaultSecretClient := kvsecrets.New(
				tc.keyvaultName,
				config.GlobalCredentials(),
				config.SecretNamingVersion(),
				config.PurgeDeletedKeyVaultSecrets(),
				config.RecoverSoftDeletedKeyVaultSecrets())

			assert.Eventually(func() bool {
				key := makeSQLUserSecretKey(keyVaultSecretClient, kvSqlUser2)
				key.Name = key.Name + "-adonet"
				var secrets, err = keyVaultSecretClient.Get(ctx, key, secrets.Flatten(true))
				assert.NoError(err)

				return len(string(secrets["secret"])) > 0
			}, tc.timeoutFast, tc.retry, "wait for keyvault to show azure sql user credentials with custom formats")

			t.Log(kvSqlUser2.Status)
		})
	})

	t.Run("deploy sql action and roll user credentials", func(t *testing.T) {
		keyVaultName := tc.keyvaultName
		keyVaultSecretClient := kvsecrets.New(
			keyVaultName,
			config.GlobalCredentials(),
			config.SecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())

		key := makeSQLUserSecretKey(keyVaultSecretClient, kvSqlUser1)
		oldSecret, err := keyVaultSecretClient.Get(ctx, key)
		assert.NoError(err)

		sqlActionName := GenerateTestResourceNameWithRandom("azuresqlaction-dev", 10)
		sqlActionInstance := &v1alpha1.AzureSqlAction{
			ObjectMeta: metav1.ObjectMeta{
				Name:      sqlActionName,
				Namespace: "default",
			},
			Spec: v1alpha1.AzureSqlActionSpec{
				ResourceGroup:      rgName,
				ServerName:         sqlServerName,
				ActionName:         "rollusercreds",
				DbName:             sqlDatabaseName,
				DbUser:             kvSqlUser1.ObjectMeta.Name,
				UserSecretKeyVault: keyVaultName,
			},
		}

		err = tc.k8sClient.Create(ctx, sqlActionInstance)
		assert.Equal(nil, err, "create sqlaction in k8s")

		sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}

		assert.Eventually(func() bool {
			_ = tc.k8sClient.Get(ctx, sqlActionInstanceNamespacedName, sqlActionInstance)
			return sqlActionInstance.Status.Provisioned
		}, tc.timeout, tc.retry, "wait for sql action to be submitted")

		newSecret, err := keyVaultSecretClient.Get(ctx, key)
		assert.NoError(err)

		assert.NotEqual(oldSecret["password"], newSecret["password"], "password should have been updated")
		assert.Equal(oldSecret["username"], newSecret["username"], "usernames should be the same")
	})

	t.Run("delete db users and ensure that their secrets have been cleaned up", func(t *testing.T) {
		EnsureDelete(ctx, t, tc, sqlUser)
		EnsureDelete(ctx, t, tc, kvSqlUser1)
		EnsureDelete(ctx, t, tc, kvSqlUser2)

		// Check that the user's secret is in the keyvault
		keyVaultSecretClient := kvsecrets.New(
			tc.keyvaultName,
			config.GlobalCredentials(),
			config.SecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())

		assert.Eventually(func() bool {
			key := secrets.SecretKey{Name: sqlUser.ObjectMeta.Name, Namespace: sqlUser.ObjectMeta.Namespace, Kind: "azuresqluser"}
			var _, err = tc.secretClient.Get(ctx, key)

			// Once the secret is gone, the Kube secret client will return an error
			return err != nil && strings.Contains(err.Error(), "not found")
		}, tc.timeoutFast, tc.retry, "wait for the azuresqluser kube secret to be deleted")

		assert.Eventually(func() bool {
			key := makeSQLUserSecretKey(keyVaultSecretClient, kvSqlUser1)

			var _, err = keyVaultSecretClient.Get(ctx, key)

			// Once the secret is gone, the KV secret client will return an err
			return err != nil && strings.Contains(err.Error(), "could not be found")
		}, tc.timeoutFast, tc.retry, "wait for the azuresqluser keyvault secret to be deleted")

		assert.Eventually(func() bool {
			key := makeSQLUserSecretKey(keyVaultSecretClient, kvSqlUser2)
			key.Name = key.Name + "-adonet"
			var _, err = keyVaultSecretClient.Get(ctx, key, secrets.Flatten(true))

			// Once the secret is gone, the KV secret client will return an err
			return err != nil && strings.Contains(err.Error(), "could not be found")
		}, tc.timeoutFast, tc.retry, "wait for the azuresqluser custom formatted keyvault secret to be deleted")
	})

	// Delete the SQL server
	EnsureDelete(ctx, t, tc, sqlServerInstance)
}

func TestAzureSqlServerKVSecretAndUser_SecretNamedCorrectly(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	sqlServerName := GenerateTestResourceNameWithRandom("kvsqlserv", 10)
	rgLocation := "westus2"

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}
	sqlServerInstance := v1beta1.NewAzureSQLServer(sqlServerNamespacedName, rgName, rgLocation)
	sqlServerInstance.Spec.KeyVaultToStoreSecrets = tc.keyvaultName

	// create and wait
	RequireInstance(ctx, t, tc, sqlServerInstance)

	//verify secret exists in k8s for server 1
	assertSQLServerAdminSecretCreated(ctx, t, sqlServerInstance)

	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &v1beta1.AzureSqlDatabase{
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

	// Create firewall rule
	sqlFirewallRuleNamespacedName := types.NamespacedName{
		Name:      GenerateTestResourceNameWithRandom("sqlfwr-local2", 10),
		Namespace: "default",
	}
	sqlFirewallRule := v1beta1.NewAzureSQLFirewallRule(
		sqlFirewallRuleNamespacedName,
		rgName,
		sqlServerName,
		"1.1.1.1",
		"255.255.255.255",
	)
	EnsureInstance(ctx, t, tc, sqlFirewallRule)

	var kvSqlUser1 *v1alpha1.AzureSQLUser

	// Run user subtests
	t.Run("group2", func(t *testing.T) {

		t.Run("set up user in db, specifying adminSecretName", func(t *testing.T) {
			t.Parallel()

			// create a sql user and verify it provisions
			username := "user4" + helpers.RandomString(10)
			roles := []string{"db_owner"}

			var adminSecret string
			if tc.secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
				adminSecret = fmt.Sprintf("%s-%s", sqlServerInstance.Namespace, sqlServerInstance.Name)
			} else {
				adminSecret = sqlServerInstance.Name
			}

			// This test will attempt to persist secrets to the KV that was instantiated as part of the test suite
			kvSqlUser1 = &v1alpha1.AzureSQLUser{
				ObjectMeta: metav1.ObjectMeta{
					Name:      username,
					Namespace: "default",
				},
				Spec: v1alpha1.AzureSQLUserSpec{
					Server:                 sqlServerName,
					DbName:                 sqlDatabaseName,
					ResourceGroup:          rgName,
					Roles:                  roles,
					KeyVaultToStoreSecrets: tc.keyvaultName,
					AdminSecretKeyVault:    tc.keyvaultName,
					AdminSecret:            adminSecret,
				},
			}

			EnsureInstance(ctx, t, tc, kvSqlUser1)

			// Check that the user's secret is in the keyvault
			keyVaultSecretClient := kvsecrets.New(
				tc.keyvaultName,
				config.GlobalCredentials(),
				config.SecretNamingVersion(),
				config.PurgeDeletedKeyVaultSecrets(),
				config.RecoverSoftDeletedKeyVaultSecrets())

			assert.Eventually(func() bool {
				key := makeSQLUserSecretKey(keyVaultSecretClient, kvSqlUser1)
				key.Name = key.Name
				var secrets, err = keyVaultSecretClient.Get(ctx, key, secrets.Flatten(true))
				assert.NoError(err)

				return len(string(secrets["secret"])) > 0
			}, tc.timeoutFast, tc.retry, "wait for keyvault to show azure sql user credentials with custom formats")

			t.Log(kvSqlUser1.Status)
		})
	})

	// Delete the SQL server
	EnsureDelete(ctx, t, tc, sqlServerInstance)
}

func makeSQLUserSecretKey(secretClient *kvsecrets.SecretClient, instance *v1alpha1.AzureSQLUser) secrets.SecretKey {
	var keyNamespace string
	if secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		keyNamespace = "azuresqluser-" + instance.Spec.Server + "-" + instance.Spec.DbName
	} else {
		keyNamespace = instance.Namespace
	}
	return secrets.SecretKey{Name: instance.Name, Namespace: keyNamespace, Kind: "azuresqluser"}
}

func makeSQLServerKeyVaultSecretName(instance *v1beta1.AzureSqlServer) string {
	if tc.secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		return fmt.Sprintf("%s-%s", instance.Namespace, instance.Name)
	} else {
		return fmt.Sprintf("%s-%s-%s", "azuresqlserver", instance.Namespace, instance.Name)
	}
}
