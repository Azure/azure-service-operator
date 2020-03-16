// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver azuresqldatabase

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/stretchr/testify/assert"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlDatabaseControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgLocation := tc.resourceGroupLocation
	sqlServerName := GenerateTestResourceNameWithRandom("sqldb-test-srv", 10)
	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldatabase-dev", 10)

	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &azurev1alpha1.AzureSqlDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlDatabaseSpec{
			Location:      rgLocation,
			ResourceGroup: GenerateTestResourceNameWithRandom("rg-test-srv", 10),
			Server:        sqlServerName,
			Edition:       0,
		},
	}

	err := tc.k8sClient.Create(ctx, sqlDatabaseInstance)
	assert.Equal(nil, err, "create db in k8s")

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return strings.Contains(sqlDatabaseInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for rg not found error")

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	assert.Equal(nil, err, "delete db in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for resource not found error")

}

func TestAzureSqlDatabaseControllerNoServer(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName
	rgLocation := tc.resourceGroupLocation
	sqlServerName := GenerateTestResourceNameWithRandom("sqldb-test-srv", 10)
	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldatabase-dev", 10)

	// Create the SqlDatabase object and expect the Reconcile to be created
	sqlDatabaseInstance := &azurev1alpha1.AzureSqlDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlDatabaseName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlDatabaseSpec{
			Location:      rgLocation,
			ResourceGroup: rgName,
			Server:        sqlServerName,
			Edition:       0,
		},
	}

	err := tc.k8sClient.Create(ctx, sqlDatabaseInstance)
	assert.Equal(false, apierrors.IsInvalid(err), "create db resource")
	assert.Equal(nil, err, "create db in k8s")

	sqlDatabaseNamespacedName := types.NamespacedName{Name: sqlDatabaseName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return strings.Contains(sqlDatabaseInstance.Status.Message, errhelp.ParentNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for rg not found error")

	err = tc.k8sClient.Delete(ctx, sqlDatabaseInstance)
	assert.Equal(nil, err, "delete db in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlDatabaseNamespacedName, sqlDatabaseInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for resource not found error")

}
