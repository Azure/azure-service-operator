// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all cosmos

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestCosmosDBHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmosdb", 8)
	cosmosDBNamespace := "default"

	dbInstance := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	err := tc.k8sClient.Create(ctx, dbInstance)
	assert.Equal(nil, err, "create cosmos db in k8s")

	dbNamespacedName := types.NamespacedName{
		Name:      cosmosDBAccountName,
		Namespace: cosmosDBNamespace,
	}
	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, dbNamespacedName, dbInstance)
		return HasFinalizer(dbInstance, finalizerName)
	}, tc.timeout, tc.retry, "wait for finalizer")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, dbNamespacedName, dbInstance)
		return dbInstance.Status.State == "Creating"
	}, tc.timeout, tc.retry, "wait for creating state")

	assert.Eventually(func() bool {
		_ = tc.k8sClient.Get(ctx, dbNamespacedName, dbInstance)
		return dbInstance.Status.State == "Succeeded"
	}, tc.timeout, tc.retry, "wait for succeeded state")

	err = tc.k8sClient.Delete(ctx, dbInstance)
	assert.Equal(nil, err, "delete cosmos db in k8s")

	assert.Eventually(func() bool {
		err = tc.k8sClient.Get(ctx, dbNamespacedName, dbInstance)
		return apierrors.IsNotFound(err)
	}, tc.timeout, tc.retry, "wait for cosmos db to be gone")

}
