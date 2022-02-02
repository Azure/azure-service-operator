// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || azuresqlserver || azuresqluser
// +build all azuresqlserver azuresqluser

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
)

func TestAzureSQLManagedUserValidatesDatabaseName(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	var sqlUser *azurev1alpha1.AzureSQLManagedUser

	sqlServerName := GenerateTestResourceNameWithRandom("sqlusr-test", 10)
	sqlDatabaseName := "master"

	username := "sql-test-user" + helpers.RandomString(10)
	roles := []string{"db_owner"}

	sqlUser = &azurev1alpha1.AzureSQLManagedUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      username,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSQLManagedUserSpec{
			Server:        sqlServerName,
			DbName:        sqlDatabaseName,
			Roles:         roles,
			ResourceGroup: "fakerg" + helpers.RandomString(10),
		},
	}

	assert := assert.New(t)

	err := tc.k8sClient.Create(ctx, sqlUser)
	assert.Error(err)
	assert.Contains(err.Error(), "'master' is a reserved database name and cannot be used")
}
