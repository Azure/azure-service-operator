// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlserver || azuresqluser
// +build all azuresqlserver azuresqluser

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAzureSQLUserControllerNoAdminSecret(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var sqlUser *azurev1alpha1.AzureSQLUser

	sqlServerName := GenerateTestResourceNameWithRandom("sqlusr-test", 10)
	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldb-test", 10)
	resourceGroup := GenerateTestResourceNameWithRandom("myrg", 10)

	username := "sql-test-user" + helpers.RandomString(10)
	roles := []string{"db_owner"}

	sqlUser = &azurev1alpha1.AzureSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSQLUserSpec{
			ResourceGroup: resourceGroup,
			Server:        sqlServerName,
			DbName:        sqlDatabaseName,
			AdminSecret:   "",
			Roles:         roles,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, sqlUser, "admin secret", false)

	EnsureDelete(ctx, t, tc, sqlUser)
}

func TestAzureSQLUserControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error
	var sqlUser *azurev1alpha1.AzureSQLUser

	sqlServerName := GenerateTestResourceNameWithRandom("sqlusr-test", 10)
	sqlDatabaseName := GenerateTestResourceNameWithRandom("sqldb-test", 10)

	username := "sql-test-user" + helpers.RandomString(10)
	roles := []string{"db_owner"}

	adminSecretKey := secrets.SecretKey{Name: sqlServerName, Namespace: "default", Kind: "AzureSQLServer"}
	data := map[string][]byte{
		"username": []byte("username"),
		"password": []byte("password"),
	}
	err = tc.secretClient.Upsert(ctx, adminSecretKey, data)
	assert.NoError(err)

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

	EnsureInstanceWithResult(ctx, t, tc, sqlUser, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlUser)
}

func TestAzureSQLUserValidatesDatabaseName(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	var sqlUser *azurev1alpha1.AzureSQLUser

	sqlServerName := GenerateTestResourceNameWithRandom("sqlusr-test", 10)
	sqlDatabaseName := "master"

	username := "sql-test-user" + helpers.RandomString(10)
	roles := []string{"db_owner"}

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

	assert := assert.New(t)

	err := tc.k8sClient.Create(ctx, sqlUser)
	assert.Error(err)
	assert.Contains(err.Error(), "'master' is a reserved database name and cannot be used")
}
