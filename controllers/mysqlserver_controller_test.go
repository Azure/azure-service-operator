// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all mysql

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
)

func TestMySQLServerControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	rgLocation := tc.resourceGroupLocation
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	// Create the mySQLServer object and expect the Reconcile to be created
	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, rgName, rgLocation)

	EnsureInstanceWithResult(ctx, t, tc, mySQLServerInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}

func TestMySQLServerControllerBadLocation(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgLocation := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	rgName := tc.resourceGroupName
	mySQLServerName := GenerateTestResourceNameWithRandom("mysql-srv", 10)

	// Create the mySQLServer object and expect the Reconcile to be created
	mySQLServerInstance := v1alpha2.NewDefaultMySQLServer(mySQLServerName, rgName, rgLocation)

	EnsureInstanceWithResult(ctx, t, tc, mySQLServerInstance, errhelp.InvalidResourceLocation, false)
	EnsureDelete(ctx, t, tc, mySQLServerInstance)
}
