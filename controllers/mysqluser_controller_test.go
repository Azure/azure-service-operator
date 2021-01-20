// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all mysqluser

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	assert := assert.New(t)
	var err error
	var mysqlServerName string
	var mysqlDatabaseName string
	var mysqlUser *v1alpha2.MySQLUser

	mysqlServerName = GenerateTestResourceNameWithRandom("psqlserver-test", 10)
	mysqlDatabaseName = GenerateTestResourceNameWithRandom("psqldb-test", 10)
	mysqlUsername := "mysql-test-user" + helpers.RandomString(10)
	roles := []string{"select on *.*"}

	secret := &v1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      mysqlServerName,
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
