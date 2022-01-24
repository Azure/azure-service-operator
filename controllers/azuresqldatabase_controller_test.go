// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlserver || azuresqldatabase
// +build all azuresqlserver azuresqldatabase

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAzureSqlDatabaseControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgLocation := tc.resourceGroupLocation
	sqlServerName := GenerateTestResourceNameWithRandom("sqldb-test-srv", 10)
	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldatabase-dev", 10)

	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &v1beta1.AzureSqlDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlDatabaseName,
			Namespace: "default",
		},
		Spec: v1beta1.AzureSqlDatabaseSpec{
			Location:      rgLocation,
			ResourceGroup: GenerateTestResourceNameWithRandom("rg-test-srv", 10),
			Server:        sqlServerName,
			Edition:       0,
		},
	}
	EnsureInstanceWithResult(ctx, t, tc, sqlDatabaseInstance, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlDatabaseInstance)
}

func TestAzureSqlDatabaseControllerNoServer(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	sqlServerName := GenerateTestResourceNameWithRandom("sqldb-test-srv", 10)
	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldatabase-dev", 10)

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

	EnsureInstanceWithResult(ctx, t, tc, sqlDatabaseInstance, errhelp.ParentNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlDatabaseInstance)

}
