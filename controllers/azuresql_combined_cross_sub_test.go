// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlserver || azuresqlservercombined
// +build all azuresqlserver azuresqlservercombined

package controllers

import (
	"context"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

func TestAzureSqlServerCombinedHappyPathCrossSub(t *testing.T) {
	t.Skip("Skipping as we don't currently have 2 Subscriptions for CI")
	// For this test to pass, remove the t.Skip above and ensure that the identity you are testing with has contributor
	// access to both the primary CI sub and the secondary subscription below.
	// You must also manually create the "test-aso-rg" resourceGroup
	secondSubscription := "82acd5bb-4206-47d4-9c12-a65db028483d"
	secondSubscriptionResourceGroupName := "test-aso-rg"

	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)
	var err error

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	sqlServerName := GenerateTestResourceNameWithRandom("sqlserver", 10)
	rgLocation := "westus3"
	rgLocation2 := "southcentralus"
	sqlServerTwoName := GenerateTestResourceNameWithRandom("sqlserver-two", 10)

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}
	sqlServerNamespacedName2 := types.NamespacedName{Name: sqlServerTwoName, Namespace: "default"}

	// Create the SqlServer object and expect the Reconcile to be created
	sqlServerInstance := v1beta1.NewAzureSQLServer(sqlServerNamespacedName, secondSubscriptionResourceGroupName, rgLocation)
	sqlServerInstance.Spec.SubscriptionID = secondSubscription

	// Send request for 2nd server (failovergroup test) before waiting on first server
	sqlServerInstance2 := v1beta1.NewAzureSQLServer(sqlServerNamespacedName2, rgName, rgLocation2)

	// create and wait
	RequireInstance(ctx, t, tc, sqlServerInstance)
	RequireInstance(ctx, t, tc, sqlServerInstance2)

	sqlDatabaseName1 := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	sqlDatabaseName3 := GenerateTestResourceNameWithRandom("sqldatabase", 10)
	var sqlDatabaseInstance1 *v1beta1.AzureSqlDatabase
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
			Location:       rgLocation,
			ResourceGroup:  secondSubscriptionResourceGroupName,
			Server:         sqlServerName,
			SubscriptionID: secondSubscription,
			Edition:        0,
		},
	}

	EnsureInstance(ctx, t, tc, sqlDatabaseInstance1)

	// run sub tests that require 1 sql server ----------------------------------
	t.Run("group1", func(t *testing.T) {
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
					Location:       rgLocation,
					ResourceGroup:  secondSubscriptionResourceGroupName,
					Server:         sqlServerName,
					SubscriptionID: secondSubscription,
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
				secondSubscriptionResourceGroupName,
				sqlServerName,
				"1.1.1.1",
				"255.255.255.255",
			)
			sqlFirewallRuleInstanceLocal.Spec.SubscriptionID = secondSubscription

			EnsureInstance(ctx, t, tc, sqlFirewallRuleInstanceLocal)
		})

		t.Run("set up azure only firewall rule in primary server", func(t *testing.T) {
			t.Parallel()

			// Create the SqlFirewallRule object and expect the Reconcile to be created
			sqlFirewallRuleInstanceRemote = v1beta1.NewAzureSQLFirewallRule(
				sqlFirewallRuleNamespacedNameRemote,
				secondSubscriptionResourceGroupName,
				sqlServerName,
				"0.0.0.0",
				"0.0.0.0",
			)
			sqlFirewallRuleInstanceRemote.Spec.SubscriptionID = secondSubscription

			EnsureInstance(ctx, t, tc, sqlFirewallRuleInstanceRemote)
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
					ResourceGroup:                secondSubscriptionResourceGroupName,
					Server:                       sqlServerName,
					SubscriptionID:               secondSubscription,
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
				secrets, err := tc.secretClient.Get(ctx, key)

				return err == nil && strings.Contains(string(secrets["azureSqlPrimaryServer"]), sqlServerName)
			}, tc.timeout, tc.retry, "wait for secret store to show failovergroup server names  ")

			sqlFailoverGroupInstance.Spec.FailoverPolicy = v1beta1.FailoverPolicyManual
			sqlFailoverGroupInstance.Spec.FailoverGracePeriod = 0 // GracePeriod cannot be set when policy is manual

			err = tc.k8sClient.Update(ctx, sqlFailoverGroupInstance)
			require.Equal(nil, err, "updating sql failover group in k8s")

			failoverGroupsClient, err := azuresqlshared.GetGoFailoverGroupsClient(config.GlobalCredentials().WithSubscriptionID(secondSubscription))
			require.Equal(nil, err, "getting failovergroup client")

			require.Eventually(func() bool {
				fog, err := failoverGroupsClient.Get(ctx, secondSubscriptionResourceGroupName, sqlServerName, sqlFailoverGroupName)
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
