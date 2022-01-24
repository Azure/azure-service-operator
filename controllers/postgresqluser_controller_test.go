// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || postgresqluser
// +build all postgresqluser

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

func TestPostgreSQLUserControllerNoAdminSecret(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var postgresqlServerName string
	var postgresqlDatabaseName string
	var postgresqlUser *azurev1alpha1.PostgreSQLUser

	postgresqlServerName = GenerateTestResourceNameWithRandom("psqlserver-test", 10)
	postgresqlDatabaseName = GenerateTestResourceNameWithRandom("psqldb-test", 10)
	resourceGroup := GenerateTestResourceNameWithRandom("myrg", 10)
	pusername := "psql-test-user" + helpers.RandomString(10)
	roles := []string{"azure_pg_admin"}

	postgresqlUser = &azurev1alpha1.PostgreSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pusername,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLUserSpec{
			ResourceGroup: resourceGroup,
			Server:        postgresqlServerName,
			DbName:        postgresqlDatabaseName,
			AdminSecret:   "",
			Roles:         roles,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, postgresqlUser, "admin secret", false)

	EnsureDelete(ctx, t, tc, postgresqlUser)
}

func TestPostgreSQLUserControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)
	var err error
	var psqlServerName string
	var psqlDatabaseName string
	var psqlUser *azurev1alpha1.PostgreSQLUser

	psqlServerName = GenerateTestResourceNameWithRandom("psqlserver-test", 10)
	psqlDatabaseName = GenerateTestResourceNameWithRandom("psqldb-test", 10)
	pusername := "psql-test-user" + helpers.RandomString(10)
	roles := []string{"azure_pg_admin"}

	adminSecretKey := secrets.SecretKey{Name: psqlServerName, Namespace: "default", Kind: "PostgreSQLServer"}
	data := map[string][]byte{
		"username": []byte("username"),
		"password": []byte("password"),
	}
	err = tc.secretClient.Upsert(ctx, adminSecretKey, data)
	assert.NoError(err)

	psqlUser = &azurev1alpha1.PostgreSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pusername,
			Namespace: "default",
		},
		Spec: azurev1alpha1.PostgreSQLUserSpec{
			Server:        psqlServerName,
			DbName:        psqlDatabaseName,
			AdminSecret:   "",
			Roles:         roles,
			ResourceGroup: "fakerg" + helpers.RandomString(10),
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, psqlUser, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, psqlUser)

}
