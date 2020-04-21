// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all psql

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPSQLDatabaseController(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var rgName string
	var rgLocation string
	var postgreSQLServerName string
	var postgreSQLServerInstance *azurev1alpha1.PostgreSQLServer

	// Add any setup steps that needs to be executed before each test
	rgName = tc.resourceGroupName
	rgLocation = tc.resourceGroupLocation

	postgreSQLServerName = GenerateTestResourceNameWithRandom("psql-srv", 10)

	// Create the PostgreSQLServer object and expect the Reconcile to be created
	postgreSQLServerInstance = &azurev1alpha1.PostgreSQLServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLServerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLServerSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Sku: azurev1alpha1.AzureDBsSQLSku{
				Name:     "B_Gen5_2",
				Tier:     azurev1alpha1.SkuTier("Basic"),
				Family:   "Gen5",
				Size:     "51200",
				Capacity: 2,
			},
			ServerVersion:  azurev1alpha1.ServerVersion("10"),
			SSLEnforcement: azurev1alpha1.SslEnforcementEnumEnabled,
		},
	}

	EnsureInstance(ctx, t, tc, postgreSQLServerInstance)

	postgreSQLDatabaseName := GenerateTestResourceNameWithRandom("psql-db", 10)

	// Create the PostgreSQLDatabase object and expect the Reconcile to be created
	postgreSQLDatabaseInstance := &azurev1alpha1.PostgreSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLDatabaseSpec{
			ResourceGroup: rgName,
			Server:        postgreSQLServerName,
		},
	}

	EnsureInstance(ctx, t, tc, postgreSQLDatabaseInstance)

	EnsureDelete(ctx, t, tc, postgreSQLDatabaseInstance)

	// Test firewall rule -------------------------------

	postgreSQLFirewallRuleName := GenerateTestResourceNameWithRandom("psql-fwrule", 10)

	// Create the PostgreSQLFirewallRule object and expect the Reconcile to be created
	postgreSQLFirewallRuleInstance := &azurev1alpha1.PostgreSQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLFirewallRuleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLFirewallRuleSpec{
			ResourceGroup:  rgName,
			Server:         postgreSQLServerName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	EnsureInstance(ctx, t, tc, postgreSQLFirewallRuleInstance)

	EnsureDelete(ctx, t, tc, postgreSQLFirewallRuleInstance)

	// Add any teardown steps that needs to be executed after each test
	EnsureDelete(ctx, t, tc, postgreSQLServerInstance)

}
