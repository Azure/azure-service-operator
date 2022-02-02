// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || mysql || mysqldatabase
// +build all mysql mysqldatabase

package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
)

func TestMySQLDatabaseControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	server := GenerateTestResourceNameWithRandom("mysqlsrv-", 10)
	mySQLDBName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	// Create the mySQLDB object and expect the Reconcile to be created
	mySQLDBInstance := &azurev1alpha1.MySQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mySQLDBName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLDatabaseSpec{
			Server:        server,
			ResourceGroup: rgName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, mySQLDBInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, mySQLDBInstance)
}

func TestMySQLDatabaseControllerNoServer(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	server := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	rgName := tc.resourceGroupName
	mySQLDBName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	// Create the mySQLDB object and expect the Reconcile to be created
	mySQLDBInstance := &azurev1alpha1.MySQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mySQLDBName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLDatabaseSpec{
			Server:        server,
			ResourceGroup: rgName,
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, mySQLDBInstance, errhelp.ResourceNotFound, false)
	EnsureDelete(ctx, t, tc, mySQLDBInstance)
}
