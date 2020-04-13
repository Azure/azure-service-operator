// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all cosmos

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestCosmosDBHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := assert.New(t)

	name := GenerateTestResourceNameWithRandom("cosmosdb", 8)
	namespace := "default"

	dbInstance := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      "westus",
			ResourceGroup: tc.resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	key := types.NamespacedName{Name: name, Namespace: namespace}
	secret := &v1.Secret{}

	EnsureInstance(ctx, t, tc, dbInstance)

	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, key, secret)
		return err == nil && secret.ObjectMeta.Name == name && secret.ObjectMeta.Namespace == namespace
	}, tc.timeoutFast, tc.retry, "wait for cosmosdb to have secret")

	EnsureDelete(ctx, t, tc, dbInstance)

	assert.Eventually(func() bool {
		err := tc.k8sClient.Get(ctx, key, secret)
		return err != nil
	}, tc.timeoutFast, tc.retry, "wait for cosmosdb to delete secret")

}
