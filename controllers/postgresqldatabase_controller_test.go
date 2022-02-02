// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || psqldatabase
// +build all psqldatabase

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//Postgresql database controller unhappy test cases

func TestPSQLDatabaseControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("psqlsrv-rg", 10)

	postgreSQLServerName := GenerateTestResourceNameWithRandom("psql-srv", 10)
	postgreSQLDatabaseName := GenerateTestResourceNameWithRandom("psql-db", 10)

	// Create the PostgreSQLDatabase object and expect the Reconcile to be created
	postgreSQLDatabaseInstance1 := &azurev1alpha1.PostgreSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLDatabaseSpec{
			ResourceGroup: rgName,
			Server:        postgreSQLServerName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, postgreSQLDatabaseInstance1, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, postgreSQLDatabaseInstance1)

}

func TestPSQLDatabaseControllerNoSever(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName

	postgreSQLServerName := GenerateTestResourceNameWithRandom("psql-srv", 10)
	postgreSQLDatabaseName := GenerateTestResourceNameWithRandom("psql-db", 10)

	// Create the PostgreSQLDatabase object and expect the Reconcile to be created
	postgreSQLDatabaseInstance2 := &azurev1alpha1.PostgreSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      postgreSQLDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLDatabaseSpec{
			ResourceGroup: rgName,
			Server:        postgreSQLServerName,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, postgreSQLDatabaseInstance2, errhelp.ResourceNotFound, false)
	EnsureDelete(ctx, t, tc, postgreSQLDatabaseInstance2)

}
