// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all mysql

package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

func TestMySQLHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgLocation := "eastus2"
	rgName := tc.resourceGroupName
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	// Create the mySQLServer object and expect the Reconcile to be created
	mySQLServerInstance := azurev1alpha1.NewDefaultMySQLServer(mySQLServerName, rgName, rgLocation)

	EnsureInstance(ctx, t, tc, mySQLServerInstance)

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
	EnsureDelete(ctx, t, tc, mySQLDBInstance)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}
