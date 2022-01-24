// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlserver || azuresqlservercombined
// +build all azuresqlserver azuresqlservercombined

package controllers

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	resourcemanagerkeyvaults "github.com/Azure/azure-service-operator/pkg/resourcemanager/keyvaults"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	testcommon "github.com/Azure/azure-service-operator/test/common"
)

func TestAzureSqlServerCombinedHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)
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
	RequireInstance(ctx, t, tc, sqlServerInstance2)

	sqlDatabaseName1 := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	sqlDatabaseName2 := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	sqlDatabaseName3 := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	var sqlDatabaseInstance1 *v1beta1.AzureSqlDatabase
	var sqlDatabaseInstance2 *v1beta1.AzureSqlDatabase
	var sqlDatabaseInstance3 *v1beta1.AzureSqlDatabase

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

	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance1 = &v1beta1.AzureSqlDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlDatabaseName1,
			Namespace: "default",
		},
		Spec: v1beta1.AzureSqlDatabaseSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Server:        sqlServerName,
			Edition:       0,
		},
	}

	EnsureInstance(ctx, t, tc, sqlDatabaseInstance1)

	// run sub tests that require 1 sql server ----------------------------------
	t.Run("group1", func(t *testing.T) {
		t.Run("set up second database in primary server using sku with maxsizebytes, then update it to use a different SKU", func(t *testing.T) {
			t.Parallel()

			maxSize := resource.MustParse("500Mi")
			// Create the SqlDatabase object and expect the Reconcile to be created
			sqlDatabaseInstance2 = &v1beta1.AzureSqlDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlDatabaseName2,
					Namespace: "default",
				},
				Spec: v1beta1.AzureSqlDatabaseSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
					Server:        sqlServerName,
					Sku: &v1beta1.SqlDatabaseSku{
						Name: "S0",
						Tier: "Standard",
					},
					MaxSize: &maxSize,
				},
			}

			EnsureInstance(ctx, t, tc, sqlDatabaseInstance2)

			namespacedName := types.NamespacedName{Name: sqlDatabaseName2, Namespace: "default"}
			err = tc.k8sClient.Get(ctx, namespacedName, sqlDatabaseInstance2)
			require.Equal(nil, err, "get sql database in k8s")

			originalHash := sqlDatabaseInstance2.Status.SpecHash

			sqlDatabaseInstance2.Spec.Sku = &v1beta1.SqlDatabaseSku{
				Name: "Basic",
				Tier: "Basic",
			}
			maxSizeMb := 100
			maxSize = resource.MustParse(fmt.Sprintf("%dMi", maxSizeMb))
			sqlDatabaseInstance2.Spec.MaxSize = &maxSize

			err = tc.k8sClient.Update(ctx, sqlDatabaseInstance2)
			require.Equal(nil, err, "updating sql database in k8s")

			require.Eventually(func() bool {
				db := &v1beta1.AzureSqlDatabase{}
				err = tc.k8sClient.Get(ctx, namespacedName, db)
				require.Equal(nil, err, "err getting DB from k8s")
				return originalHash != db.Status.SpecHash
			}, tc.timeout, tc.retry, "wait for sql database to be updated in k8s")

			require.Eventually(func() bool {
				db, err := tc.sqlDbManager.GetDB(ctx, rgName, sqlServerName, sqlDatabaseName2)
				require.Equal(nil, err, "err getting DB fromAzure")
				return db.Sku != nil && db.Sku.Name != nil && *db.Sku.Name == "Basic"
			}, tc.timeout, tc.retry, "wait for sql database Sku.Name to be updated in azure")

			require.Eventually(func() bool {
				db, err := tc.sqlDbManager.GetDB(ctx, rgName, sqlServerName, sqlDatabaseName2)
				require.Equal(nil, err, "err getting DB fromAzure")
				return db.Sku != nil && db.Sku.Tier != nil && *db.Sku.Tier == "Basic"
			}, tc.timeout, tc.retry, "wait for sql database Sku.Tier to be updated in azure")

			require.Eventually(func() bool {
				db, err := tc.sqlDbManager.GetDB(ctx, rgName, sqlServerName, sqlDatabaseName2)
				require.Equal(nil, err, "err getting DB fromAzure")
				return db.MaxSizeBytes != nil && *db.MaxSizeBytes == int64(maxSizeMb)*int64(1024)*int64(1024)
			}, tc.timeout, tc.retry, "wait for sql database MaxSizeBytes to be updated in azure")

			// At this point the DB should be in a stable state, ensure that the right status is set
			db := &v1beta1.AzureSqlDatabase{}
			err = tc.k8sClient.Get(ctx, namespacedName, db)
			require.Equal(nil, err, "err getting DB from k8s")

			require.Equal(false, db.Status.Provisioning)
			require.Equal(false, db.Status.FailedProvisioning)
			require.Equal(true, db.Status.Provisioned)
		})

		// Create a database in the new server
		t.Run("set up database with short and long term retention", func(t *testing.T) {
			t.Parallel()

			// Create the SqlDatabase object and expect the Reconcile to be created
			sqlDatabaseInstance3 = &v1beta1.AzureSqlDatabase{
				ObjectMeta: metav1.ObjectMeta{
					Name:      sqlDatabaseName3,
					Namespace: "default",
				},
				Spec: v1beta1.AzureSqlDatabaseSpec{
					Location:      rgLocation,
					ResourceGroup: rgName,
					Server:        sqlServerName,
					Sku: &v1beta1.SqlDatabaseSku{
						Name: "S0",
						Tier: "Standard",
					},
					WeeklyRetention: "P3W",
					ShortTermRetentionPolicy: &v1beta1.SQLDatabaseShortTermRetentionPolicy{
						RetentionDays: 3,
					},
				},
			}

			EnsureInstance(ctx, t, tc, sqlDatabaseInstance3)

			// Now update with an invalid retention policy
			sqlDatabaseInstance3.Spec.ShortTermRetentionPolicy.RetentionDays = -1
			err = tc.k8sClient.Update(ctx, sqlDatabaseInstance3)
			require.Equal(nil, err, "updating sql database in k8s")

			namespacedName := types.NamespacedName{Name: sqlDatabaseName3, Namespace: "default"}
			require.Eventually(func() bool {
				db := &v1beta1.AzureSqlDatabase{}
				err = tc.k8sClient.Get(ctx, namespacedName, db)
				require.Equal(nil, err, "err getting DB from k8s")
				return db.Status.Provisioned == false && strings.Contains(db.Status.Message, errhelp.BackupRetentionPolicyInvalid)
			}, tc.timeout, tc.retry, "wait for sql database to be updated in k8s")
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

	var sqlFailoverGroupInstance *v1beta1.AzureSqlFailoverGroup
	sqlFailoverGroupName := GenerateTestResourceNameWithRandom("sqlfog-dev", 10)

	sqlFailoverGroupNamespacedName := types.NamespacedName{Name: sqlFailoverGroupName, Namespace: "default"}

	t.Run("group3", func(t *testing.T) {
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
					FailoverPolicy:               v1beta1.FailoverPolicyAutomatic,
					FailoverGracePeriod:          60,
					SecondaryServer:              sqlServerTwoName,
					SecondaryServerResourceGroup: rgName,
					DatabaseList:                 []string{sqlDatabaseName1},
				},
			}

			EnsureInstance(ctx, t, tc, sqlFailoverGroupInstance)

			// verify secret has been created
			require.Eventually(func() bool {
				key := secrets.SecretKey{Name: sqlFailoverGroupInstance.Name, Namespace: sqlFailoverGroupInstance.Namespace, Kind: "AzureSqlFailoverGroup"}
				var secrets, err = tc.secretClient.Get(ctx, key)

				return err == nil && strings.Contains(string(secrets["azureSqlPrimaryServer"]), sqlServerName)
			}, tc.timeout, tc.retry, "wait for secret store to show failovergroup server names  ")

			sqlFailoverGroupInstance.Spec.FailoverPolicy = v1beta1.FailoverPolicyManual
			sqlFailoverGroupInstance.Spec.FailoverGracePeriod = 0 // GracePeriod cannot be set when policy is manual

			err = tc.k8sClient.Update(ctx, sqlFailoverGroupInstance)
			require.Equal(nil, err, "updating sql failover group in k8s")

			failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(config.GlobalCredentials())
			require.Equal(nil, err, "getting failovergroup client")

			require.Eventually(func() bool {
				fog, err := failoverGroupsClient.Get(ctx, rgName, sqlServerName, sqlFailoverGroupName)
				require.Equal(nil, err, "err getting failover group from Azure")
				return fog.ReadWriteEndpoint.FailoverPolicy == sql.Manual
			}, tc.timeout, tc.retry, "wait for sql failover group failover policy to be updated in Azure")
		})

	})

	t.Run("group4", func(t *testing.T) {

		t.Run("delete failovergroup", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlFailoverGroupInstance)
		})
		t.Run("delete dbs", func(t *testing.T) {
			t.Parallel()
			EnsureDelete(ctx, t, tc, sqlDatabaseInstance1)
			EnsureDelete(ctx, t, tc, sqlDatabaseInstance2)
			EnsureDelete(ctx, t, tc, sqlDatabaseInstance3)
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

func TestAzureSqlServer_KeyVaultSoftDelete_CreateDeleteCreateAgain(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)

	rgLocation := "westus2"

	// Create a KeyVault with soft delete enabled that we can use to perform our tests
	keyVaultName := GenerateAlphaNumTestResourceNameWithRandom("kvsoftdel", 5)
	objID, err := resourcemanagerkeyvaults.GetObjectID(
		context.Background(),
		config.GlobalCredentials(),
		config.GlobalCredentials().TenantID(),
		config.GlobalCredentials().ClientID())
	require.NoError(err)

	err = testcommon.CreateKeyVaultSoftDeleteEnabled(
		context.Background(),
		config.GlobalCredentials(),
		tc.resourceGroupName,
		keyVaultName,
		rgLocation,
		objID)
	require.NoError(err)

	sqlServerName := GenerateTestResourceNameWithRandom("sqlserver", 10)
	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}
	sqlServerInstance := v1beta1.NewAzureSQLServer(sqlServerNamespacedName, tc.resourceGroupName, rgLocation)
	sqlServerInstance.Spec.KeyVaultToStoreSecrets = keyVaultName

	// create and wait
	RequireInstance(ctx, t, tc, sqlServerInstance)

	EnsureDelete(ctx, t, tc, sqlServerInstance)

	// Recreate with the same name
	sqlServerInstance = v1beta1.NewAzureSQLServer(sqlServerNamespacedName, tc.resourceGroupName, rgLocation)
	sqlServerInstance.Spec.KeyVaultToStoreSecrets = keyVaultName
	RequireInstance(ctx, t, tc, sqlServerInstance)

	EnsureDelete(ctx, t, tc, sqlServerInstance)
}
