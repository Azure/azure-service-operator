// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver azuresqluser

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSQLUserControllerNoAdminSecret(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error
	var sqlServerName string
	var sqlDatabaseName string
	var sqlUser *azurev1alpha1.AzureSQLUser

	sqlServerName = GenerateTestResourceNameWithRandom("sqlusr-test", 10)

	username := "sql-test-user" + helpers.RandomString(10)
	roles := []string{"db_owner"}

	sqlUser = &azurev1alpha1.AzureSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSQLUserSpec{
			Server:      sqlServerName,
			DbName:      sqlDatabaseName,
			AdminSecret: "",
			Roles:       roles,
		},
	}

	// Create the sqlUser
	err = tc.k8sClient.Create(ctx, sqlUser)
	assert.Equal(nil, err, "create db user in k8s")

	sqlUserNamespacedName := types.NamespacedName{Name: username, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return HasFinalizer(sqlUser, finalizerName)
	}, tc.timeout, tc.retry, "wait for finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return strings.Contains(sqlUser.Status.Message, "admin secret")
	}, tc.timeout, tc.retry, "wait for missing admin secret message")

	err = tc.k8sClient.Delete(ctx, sqlUser)
	assert.Equal(nil, err, "delete db user in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for user to be gone from k8s")

}

func TestAzureSQLUserControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error
	var sqlServerName string
	var sqlDatabaseName string
	var sqlUser *azurev1alpha1.AzureSQLUser

	sqlServerName = GenerateTestResourceNameWithRandom("sqlusr-test", 10)

	username := "sql-test-user" + helpers.RandomString(10)
	roles := []string{"db_owner"}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlServerName,
			Namespace: "default",
		},
		// Needed to avoid nil map error
		Data: map[string][]byte{
			"username": []byte("username"),
			"password": []byte("password"),
		},
		Type: "Opaque",
	}

	// Create the sqlUser
	err = tc.k8sClient.Create(ctx, secret)
	assert.Equal(nil, err, "create admin secret in k8s")

	sqlUser = &azurev1alpha1.AzureSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSQLUserSpec{
			Server:        sqlServerName,
			DbName:        sqlDatabaseName,
			AdminSecret:   "",
			Roles:         roles,
			ResourceGroup: "fakerg" + helpers.RandomString(10),
		},
	}

	// Create the sqlUser
	err = tc.k8sClient.Create(ctx, sqlUser)
	assert.Equal(nil, err, "create db user in k8s")

	sqlUserNamespacedName := types.NamespacedName{Name: username, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return HasFinalizer(sqlUser, finalizerName)
	}, tc.timeout, tc.retry, "wait for finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return strings.Contains(sqlUser.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for rg fail message")

	err = tc.k8sClient.Delete(ctx, sqlUser)
	assert.Equal(nil, err, "delete db user in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlUserNamespacedName, sqlUser)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for user to be gone from k8s")

}
