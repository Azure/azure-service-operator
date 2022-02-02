// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || mysql || mysqlfirewallrule
// +build all mysql mysqlfirewallrule

package controllers

import (
	"context"
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
)

func TestMySQLAdministratorNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	server := GenerateTestResourceNameWithRandom("mysqlsrv", 10)
	adminName := GenerateTestResourceNameWithRandom("mysql-admin", 10)

	// Create the admin object and expect the Reconcile to be created
	adminInstance := &azurev1alpha1.MySQLServerAdministrator{
		ObjectMeta: metav1.ObjectMeta{
			Name:      adminName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLServerAdministratorSpec{
			ResourceGroup:     rgName,
			Server:            server,
			Login:             "my-mi",
			AdministratorType: azurev1alpha1.MySQLServerAdministratorTypeActiveDirectory,
			Sid:               "00000000-0000-0000-0000-000000000000",
			TenantId:          "00000000-0000-0000-0000-000000000000",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, adminInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, adminInstance)
}

func TestMySQLAdministratorNoServer(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	server := GenerateTestResourceNameWithRandom("mysqlsrv-rg", 10)
	adminName := GenerateTestResourceNameWithRandom("mysql-admin", 10)

	// Create the admin object and expect the Reconcile to be created
	adminInstance := &azurev1alpha1.MySQLServerAdministrator{
		ObjectMeta: metav1.ObjectMeta{
			Name:      adminName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.MySQLServerAdministratorSpec{
			ResourceGroup:     rgName,
			Server:            server,
			Login:             "my-mi",
			AdministratorType: azurev1alpha1.MySQLServerAdministratorTypeActiveDirectory,
			Sid:               "00000000-0000-0000-0000-000000000000",
			TenantId:          "00000000-0000-0000-0000-000000000000",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, adminInstance, errhelp.ResourceNotFound, false)
	EnsureDelete(ctx, t, tc, adminInstance)
}
