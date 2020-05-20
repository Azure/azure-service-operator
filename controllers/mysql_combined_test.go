// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all mysql

package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
)

func TestMySQLHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgLocation := "eastus2"
	rgName := tc.resourceGroupName
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)
	mySQLReplicaName := GenerateTestResourceNameWithRandom("mysql-rep", 10)

	// Create the mySQLServer object and expect the Reconcile to be created
	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, rgName, rgLocation)

	RequireInstance(ctx, t, tc, mySQLServerInstance)

	// Create a mySQL replica
	mySQLReplicaInstance := v1alpha2.NewReplicaMySQLServer(mySQLReplicaName, rgName, rgLocation, mySQLServerInstance.Status.ResourceId)
	mySQLReplicaInstance.Spec.StorageProfile = nil

	EnsureInstance(ctx, t, tc, mySQLReplicaInstance)

	mySQLDBName := GenerateTestResourceNameWithRandom("mysql-db", 10)

	// Create the mySQLDB object and expect the Reconcile to be created
	mySQLDBInstance := &azurev1alpha1.MySQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mySQLDBName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLDatabaseSpec{
			Server:        mySQLServerName,
			ResourceGroup: rgName,
		},
	}

	EnsureInstance(ctx, t, tc, mySQLDBInstance)

	ruleName := GenerateTestResourceNameWithRandom("mysql-fw", 10)

	ruleInstance := &azurev1alpha1.MySQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ruleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLFirewallRuleSpec{
			Server:         mySQLServerName,
			ResourceGroup:  rgName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "0.0.0.0",
		},
	}

	EnsureInstance(ctx, t, tc, ruleInstance)

	// Create VNet and VNetRules -----
	RunMySqlVNetRuleHappyPath(t, mySQLServerName, rgLocation)

	EnsureDelete(ctx, t, tc, ruleInstance)
	EnsureDelete(ctx, t, tc, mySQLDBInstance)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
	EnsureDelete(ctx, t, tc, mySQLReplicaInstance)
}
