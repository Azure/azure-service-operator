// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all mysql

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/mysqluser"
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

	mySQLAdmin := &azurev1alpha1.MySQLServerAdministrator{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "admin",
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLServerAdministratorSpec{
			ResourceGroup:     rgName,
			Server:            mySQLServerName,
			AdministratorType: azurev1alpha1.MySQLServerAdministratorTypeActiveDirectory,
			// The below fields are for a user managed identity which exists in the ASO-CI subscription
			// TODO: That means this test won't pass locally if not run against that sub...
			TenantId: "72f988bf-86f1-41af-91ab-2d7cd011db47",
			Login:    "azureserviceoperator-test-mi",
			Sid:      "ad84ef71-31fc-4797-b970-48bd515edf5c",
		},
	}
	EnsureInstance(ctx, t, tc, mySQLAdmin)

	// Delete the admin
	EnsureDelete(ctx, t, tc, mySQLAdmin)

	ruleName := GenerateTestResourceNameWithRandom("mysql-fw", 10)

	// This rule opens access to the public internet, but in this case
	// there's literally no data in the database
	ruleInstance := &azurev1alpha1.MySQLFirewallRule{
		ObjectMeta: metav1.ObjectMeta{
			Name:      ruleName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLFirewallRuleSpec{
			Server:         mySQLServerName,
			ResourceGroup:  rgName,
			StartIPAddress: "0.0.0.0",
			EndIPAddress:   "255.255.255.255",
		},
	}

	EnsureInstance(ctx, t, tc, ruleInstance)

	// Create user and ensure it can be updated
	RunMySQLUserHappyPath(ctx, t, mySQLServerName,mySQLDBName, rgName)

	// Create VNet and VNetRules -----
	RunMySqlVNetRuleHappyPath(t, mySQLServerName, rgLocation)

	EnsureDelete(ctx, t, tc, ruleInstance)
	EnsureDelete(ctx, t, tc, mySQLDBInstance)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
	EnsureDelete(ctx, t, tc, mySQLReplicaInstance)
}

func RunMySQLUserHappyPath(ctx context.Context, t *testing.T, mySQLServerName string, mySQLDBName string, rgName string) {
	assert := assert.New(t)

	// Create a user in the DB
	username := GenerateTestResourceNameWithRandom("user", 10)
	user := &azurev1alpha1.MySQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLUserSpec{
			ResourceGroup: rgName,
			Server:        mySQLServerName,
			DbName:        mySQLDBName,
			Username:      username,
			Roles:         []string{"SELECT"},
		},
	}
	EnsureInstance(ctx, t, tc, user)

	// Update user
	namespacedName := types.NamespacedName{Name: username, Namespace: "default"}
	err := tc.k8sClient.Get(ctx, namespacedName, user)
	assert.NoError(err)

	updatedRoles := []string{"UPDATE", "DELETE", "CREATE", "DROP"}
	user.Spec.Roles = updatedRoles
	err = tc.k8sClient.Update(ctx, user)
	assert.NoError(err)

	// TODO: Ugh this is fragile, the path to the secret should probably be set on the status?
	// See issue here: https://github.com/Azure/azure-service-operator/issues/1318
	secretNamespacedName := types.NamespacedName{Name: mySQLServerName, Namespace: "default"}
	adminSecret, err := tc.secretClient.Get(ctx, secretNamespacedName)
	adminUser := string(adminSecret["fullyQualifiedUsername"])
	adminPassword := string(adminSecret[mysqluser.MSecretPasswordKey])
	fullServerName := string(adminSecret["fullyQualifiedServerName"])

	db, err := mysql.ConnectToSqlDB(
		ctx,
		mysql.MySQLDriverName,
		fullServerName,
		mySQLDBName,
		mysql.MySQLServerPort,
		adminUser,
		adminPassword)
	assert.NoError(err)

	assert.Eventually(func() bool {
		roles, err := mysql.ExtractUserRoles(ctx, db, username, mySQLDBName)
		assert.NoError(err)

		if len(roles) != len(updatedRoles) {
			return false
		}

		for _, role := range updatedRoles {
			if _, ok := roles[role]; !ok {
				return false
			}
		}

		return true
	}, tc.timeout, tc.retry, "waiting for DB user to be updated")
}
