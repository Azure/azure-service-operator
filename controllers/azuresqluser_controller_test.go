// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver azuresqluser

package controllers

import (
	"context"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestAzureSQLUserControllerNoAdminSecret(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

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

	EnsureInstanceWithResult(ctx, t, tc, sqlUser, "admin secret", false)

	EnsureDelete(ctx, t, tc, sqlUser)
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

	EnsureInstanceWithResult(ctx, t, tc, sqlUser, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, sqlUser)

}
