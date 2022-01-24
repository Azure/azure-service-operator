// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || cosmos
// +build all cosmos

package controllers

import (
	"context"
	"testing"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/stretchr/testify/assert"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
			Location:      tc.resourceGroupLocation,
			ResourceGroup: tc.resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	EnsureInstance(ctx, t, tc, dbInstance)

	key := secrets.SecretKey{Name: name, Namespace: namespace, Kind: "CosmosDB"}
	assert.Eventually(func() bool {
		secret, err := tc.secretClient.Get(ctx, key)
		return err == nil && len(secret) > 0
	}, tc.timeoutFast, tc.retry, "wait for cosmosdb to have secret")

	t.Run("group1", func(t *testing.T) {
		t.Run("Test Cosmos DB SQL Database happy path", func(t *testing.T) {
			CosmosDBSQLDatabaseHappyPath(t, name)
		})
		t.Run("Test Cosmos DB SQL Database failed throughput update", func(t *testing.T) {
			CosmosDBSQLDatabase_FailedThroughputUpdate(t, name)
		})
	})

	EnsureDelete(ctx, t, tc, dbInstance)

	assert.Eventually(func() bool {
		_, err := tc.secretClient.Get(ctx, key)
		return err != nil
	}, tc.timeoutFast, tc.retry, "wait for cosmosdb to delete secret")

}

func TestCosmosDBControllerNoResourceGroup(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	rgLocation := tc.resourceGroupLocation
	//wrong resource group name
	resourceGroupName := "gone"

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmosdb", 8)
	cosmosDBNamespace := "default"

	dbInstance1 := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, dbInstance1, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, dbInstance1)
}

func TestCosmosDBControllerInvalidLocation(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	resourceGroupName := tc.resourceGroupName
	//rglocation doesnot exist
	rgLocation := GenerateTestResourceNameWithRandom("cosmos-lo", 10)

	cosmosDBAccountName := GenerateTestResourceNameWithRandom("cosmos-db", 8)
	cosmosDBNamespace := "default"

	dbInstance2 := &v1alpha1.CosmosDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cosmosDBAccountName,
			Namespace: cosmosDBNamespace,
		},
		Spec: v1alpha1.CosmosDBSpec{
			Location:      rgLocation,
			ResourceGroup: resourceGroupName,
			Kind:          v1alpha1.CosmosDBKindGlobalDocumentDB,
			Properties: v1alpha1.CosmosDBProperties{
				DatabaseAccountOfferType: v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard,
			},
		},
	}

	//error meessage to be expected
	errMessage := "The specified location '" + rgLocation + "' is invalid"

	EnsureInstanceWithResult(ctx, t, tc, dbInstance2, errMessage, false)
	EnsureDelete(ctx, t, tc, dbInstance2)
}
