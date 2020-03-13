// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver azuresqlservercombined testaction

package controllers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	v1 "k8s.io/api/core/v1"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func RunSQLActionHappy(t *testing.T, server string) {
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	// Add any setup steps that needs to be executed before each test
	rgName := tc.resourceGroupName

	//Get SQL credentials to compare after rollover
	secret := &v1.Secret{}
	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, types.NamespacedName{Name: server, Namespace: "default"}, secret)
		if err != nil {
			return false
		}
		return true
	}, tc.timeoutFast, tc.retry, "wait for server to return secret")

	sqlActionName := GenerateTestResourceNameWithRandom("azuresqlaction-dev", 10)

	// Create the Sql Action object and expect the Reconcile to be created
	sqlActionInstance := &azurev1alpha1.AzureSqlAction{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlActionName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlActionSpec{
			ActionName:    "rollcreds",
			ServerName:    server,
			ResourceGroup: rgName,
		},
	}

	err := tc.k8sClient.Create(ctx, sqlActionInstance)
	assert.Equal(nil, err, "create sqlaction in k8s")

	sqlActionInstanceNamespacedName := types.NamespacedName{Name: sqlActionName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlActionInstanceNamespacedName, sqlActionInstance)
		return sqlActionInstance.Status.Provisioned
	}, tc.timeout, tc.retry, "wait for sql action to be submitted")

	// TODO Check SQL Database credentials

	// makre sure credentials are not the same as previous
	secretAfter := &v1.Secret{}
	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, types.NamespacedName{Name: server, Namespace: "default"}, secretAfter)
		if err != nil {
			return false
		}
		return true
	}, tc.timeoutFast, tc.retry, "wait for server to return secret")

	assert.Equal(secret.Data["username"], secretAfter.Data["username"], "username should still be the same")
	assert.NotEqual(string(secret.Data["password"]), string(secretAfter.Data["password"]), "password should have changed")
}
