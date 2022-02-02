// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || mysqluser
// +build all mysqluser

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

func TestMySQLUserControllerNoAdminSecret(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	var mysqlServerName string
	var mysqlDatabaseName string
	var mysqlUser *v1alpha2.MySQLUser

	mysqlServerName = GenerateTestResourceNameWithRandom("mysqlserver-test", 10)
	mysqlDatabaseName = GenerateTestResourceNameWithRandom("mysqldb-test", 10)
	resourceGroup := GenerateTestResourceNameWithRandom("myrg", 10)
	mysqlusername := "mysql-test-user" + helpers.RandomString(10)
	roles := []string{"select on *.* "}

	mysqlUser = &v1alpha2.MySQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mysqlusername,
			Namespace: "default",
		},
		Spec: v1alpha2.MySQLUserSpec{
			ResourceGroup: resourceGroup,
			Server:        mysqlServerName,
			DatabaseRoles: map[string][]string{
				mysqlDatabaseName: roles,
			},
			AdminSecret: "",
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, mysqlUser, "admin secret", false)

	EnsureDelete(ctx, t, tc, mysqlUser)
}

func TestMySQLUserControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)
	var err error
	var mysqlServerName string
	var mysqlDatabaseName string
	var mysqlUser *v1alpha2.MySQLUser

	mysqlServerName = GenerateTestResourceNameWithRandom("psqlserver-test", 10)
	mysqlDatabaseName = GenerateTestResourceNameWithRandom("psqldb-test", 10)
	mysqlUsername := "mysql-test-user" + helpers.RandomString(10)
	roles := []string{"select on *.*"}

	adminSecretKey := secrets.SecretKey{Name: mysqlServerName, Namespace: "default", Kind: "MySQLServer"}
	data := map[string][]byte{
		"username": []byte("username"),
		"password": []byte("password"),
	}
	err = tc.secretClient.Upsert(ctx, adminSecretKey, data)
	require.NoError(err)

	mysqlUser = &v1alpha2.MySQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mysqlUsername,
			Namespace: "default",
		},
		Spec: v1alpha2.MySQLUserSpec{
			Server:      mysqlServerName,
			AdminSecret: "",
			DatabaseRoles: map[string][]string{
				mysqlDatabaseName: roles,
			},
			ResourceGroup: "fakerg" + helpers.RandomString(10),
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, mysqlUser, errhelp.ResourceGroupNotFoundErrorCode, false)

	EnsureDelete(ctx, t, tc, mysqlUser)
}

func TestMySQLUserWebhook(t *testing.T) {
	// The webhook prevents a user from being created with ALL in roles.
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	require := require.New(t)

	user := v1alpha2.MySQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      GenerateTestResourceNameWithRandom("mysqluser", 5),
			Namespace: "default",
		},
		Spec: v1alpha2.MySQLUserSpec{
			Server:        "does-not-matter",
			Roles:         []string{"PROCESS", "ALL"},
			ResourceGroup: "also-does-not-matter",
		},
	}
	err := tc.k8sClient.Create(ctx, &user)
	require.NotNil(err)
	require.Contains(err.Error(), "ASO admin user doesn't have privileges to grant ALL at server level")
}
