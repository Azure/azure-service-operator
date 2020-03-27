// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
)

// NewAzureCosmosDBManager creates a new cosmos db client
func NewAzureCosmosDBManager() *AzureCosmosDBManager {
	return &AzureCosmosDBManager{}
}

// CosmosDBManager client functions
type CosmosDBManager interface {
	// CreateCosmosDB creates a new cosmos database account
	CreateCosmosDB(ctx context.Context, groupName string, cosmosDBName string, location string, kind azurev1alpha1.CosmosDBKind, dbType azurev1alpha1.CosmosDBDatabaseAccountOfferType, tags map[string]*string) (*documentdb.DatabaseAccount, error)

	// GetCosmosDB gets a cosmos database account
	GetCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (result documentdb.DatabaseAccount, err error)

	// DeleteCosmosDB removes the cosmos database account
	DeleteCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (result documentdb.DatabaseAccountsDeleteFuture, err error)
}
