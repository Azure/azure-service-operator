// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// +build all cosmos

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCosmosDBHappyPath(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmosdb", 8)
	cosmosDBNamespace := "default"

	dbInstance := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
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

	EnsureInstance(ctx, t, tc, dbInstance)

	EnsureDelete(ctx, t, tc, dbInstance)

}
