// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

//go:build all || cosmos
// +build all cosmos

package controllers

import (
	"context"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2021-03-15/documentdb"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/cosmosdb"
)

func CosmosDBSQLDatabaseHappyPath(t *testing.T, cosmosDBAccountName string) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := require.New(t)

	cosmosDBSQLDatabaseClient, err := cosmosdb.GetCosmosDBSQLDatabaseClient(config.GlobalCredentials())
	assert.Equal(nil, err, "failed to get cosmos db SQL database client")

	name := GenerateTestResourceNameWithRandom("cosmosdbsql", 8)
	namespace := "default"

	var throughPutRUs int32 = 400
	sqlDBInstance := &v1alpha1.CosmosDBSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.CosmosDBSQLDatabaseSpec{
			ResourceGroup: tc.resourceGroupName,
			Account:       cosmosDBAccountName,
			Throughput:    &throughPutRUs,
		},
	}

	// Create the Cosmos SQL DB
	EnsureInstance(ctx, t, tc, sqlDBInstance)
	assert.Eventually(func() bool {
		cosmosSQLDB, err := cosmosDBSQLDatabaseClient.GetSQLDatabase(ctx, tc.resourceGroupName, cosmosDBAccountName, name)
		assert.Equal(nil, err, "err getting DB from Azure")

		throughputMatches, err := doesThroughputMatch(
			ctx,
			cosmosDBSQLDatabaseClient,
			tc.resourceGroupName,
			cosmosDBAccountName,
			name,
			throughPutRUs)
		assert.Equal(nil, err, "err ensuring throughput matches")

		return cosmosSQLDB.Name != nil && *cosmosSQLDB.Name == name && throughputMatches
	}, tc.timeout, tc.retry, "wait for cosmos SQL DB to exist in Azure")

	// Get the updated Cosmos SQL DB from k8s
	namespacedName := types.NamespacedName{Name: name, Namespace: namespace}
	err = tc.k8sClient.Get(ctx, namespacedName, sqlDBInstance)
	assert.Equal(nil, err, "getting cosmos SQL DB from k8s")

	// Update the Cosmos SQL DB throughput
	throughPutRUs = 1000
	sqlDBInstance.Spec.Throughput = &throughPutRUs
	err = tc.k8sClient.Update(ctx, sqlDBInstance)
	assert.Equal(nil, err, "updating cosmos sql DB in k8s")

	assert.Eventually(func() bool {
		throughputMatches, err := doesThroughputMatch(
			ctx,
			cosmosDBSQLDatabaseClient,
			tc.resourceGroupName,
			cosmosDBAccountName,
			name,
			throughPutRUs)
		assert.Equal(nil, err, "err ensuring throughput matches")

		return throughputMatches
	}, tc.timeout, tc.retry, "wait for cosmos SQL DB throughput to be updated in Azure")

	EnsureDelete(ctx, t, tc, sqlDBInstance)
}

func CosmosDBSQLDatabase_FailedThroughputUpdate(t *testing.T, cosmosDBAccountName string) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()
	assert := require.New(t)

	name := GenerateTestResourceNameWithRandom("cosmosdbsql", 8)
	namespace := "default"

	sqlDBInstance := &v1alpha1.CosmosDBSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.CosmosDBSQLDatabaseSpec{
			ResourceGroup: tc.resourceGroupName,
			Account:       cosmosDBAccountName,
		},
	}

	// Create the Cosmos SQL DB
	EnsureInstance(ctx, t, tc, sqlDBInstance)

	// Get the updated Cosmos SQL DB from k8s
	namespacedName := types.NamespacedName{Name: name, Namespace: namespace}
	err := tc.k8sClient.Get(ctx, namespacedName, sqlDBInstance)
	assert.Equal(nil, err, "getting cosmos SQL DB from k8s")

	// Update the Cosmos SQL DB throughput (we expect that this will be rejected by Cosmos DB)
	var throughPutRUs int32 = 1000
	sqlDBInstance.Spec.Throughput = &throughPutRUs
	err = tc.k8sClient.Update(ctx, sqlDBInstance)
	assert.Equal(nil, err, "updating cosmos sql DB in k8s")

	// Ensure we get an error eventually
	assert.Eventually(func() bool {
		updatedInstance := &v1alpha1.CosmosDBSQLDatabase{}
		err = tc.k8sClient.Get(ctx, namespacedName, updatedInstance)
		assert.Equal(nil, err, "err getting Cosmos SQL DB from k8s")

		return updatedInstance.Status.Provisioned == false &&
			updatedInstance.Status.FailedProvisioning == true &&
			strings.Contains(updatedInstance.Status.Message, "Throughput update is not supported for resources created without dedicated throughput")
	}, tc.timeout, tc.retry, "wait for sql database to be updated in k8s")

	EnsureDelete(ctx, t, tc, sqlDBInstance)
}

func doesThroughputMatch(
	ctx context.Context,
	client documentdb.SQLResourcesClient,
	rgName string,
	cosmosDBAccountName string,
	dbName string,
	expectedThroughput int32) (bool, error) {

	throughputSettings, err := client.GetSQLDatabaseThroughput(ctx, rgName, cosmosDBAccountName, dbName)
	if err != nil {
		return false, errors.Wrap(err, "err getting cosmos sql DB throughput from Azure")
	}

	if throughputSettings.ThroughputSettingsGetProperties == nil {
		return false, nil
	}

	if throughputSettings.ThroughputSettingsGetProperties.Resource == nil {
		return false, nil
	}

	if throughputSettings.ThroughputSettingsGetProperties.Resource.Throughput == nil {
		return false, nil
	}

	return *throughputSettings.ThroughputSettingsGetProperties.Resource.Throughput == expectedThroughput, nil
}

func TestCosmosDBSQLDatabaseNoCosmosDBAccount(t *testing.T) {
	t.Parallel()
	defer PanicRecover(t)
	ctx := context.Background()

	//wrong resource group name
	resourceGroupName := "gone"

	name := GenerateTestResourceNameWithRandom("cosmosdbsql", 8)
	namespace := "default"

	sqlDBInstance := &v1alpha1.CosmosDBSQLDatabase{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
		Spec: v1alpha1.CosmosDBSQLDatabaseSpec{
			ResourceGroup: resourceGroupName,
			Account:       name,
		},
	}

	EnsureInstanceWithResult(ctx, t, tc, sqlDBInstance, errhelp.ResourceGroupNotFoundErrorCode, false)
	EnsureDelete(ctx, t, tc, sqlDBInstance)
}
