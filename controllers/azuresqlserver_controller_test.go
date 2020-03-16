// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all azuresqlserver

package controllers

import (
	"context"
	"strings"
	"testing"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestAzureSqlServerControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	sqlServerName := GenerateTestResourceNameWithRandom("sqlserver-dev", 10)

	// Create the SqlServer object and expect the Reconcile to be created
	sqlServerInstance := &azurev1alpha1.AzureSqlServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      sqlServerName,
			Namespace: "default",
		},
		Spec: azurev1alpha1.AzureSqlServerSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: GenerateTestResourceNameWithRandom("rg-fake-dev", 10),
		},
	}

	err := tc.k8sClient.Create(ctx, sqlServerInstance)
	assert.Equal(nil, err, "create sql server in k8s")

	sqlServerNamespacedName := types.NamespacedName{Name: sqlServerName, Namespace: "default"}

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return HasFinalizer(sqlServerInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return strings.Contains(sqlServerInstance.Status.Message, errhelp.ResourceGroupNotFoundErrorCode)
	}, tc.timeout, tc.retry, "wait for rg error")

	err = tc.k8sClient.Delete(ctx, sqlServerInstance)
	assert.Equal(nil, err, "delete sql server in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, sqlServerNamespacedName, sqlServerInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for server to be gone")

}
