// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
)

// NewAzureCosmosDBManager creates a new cosmos db client
func NewAzureCosmosDBManager() *AzureCosmosDBManager {
	return &AzureCosmosDBManager{}
}

// CosmosDBManager client functions
type CosmosDBManager interface {
	// CreateOrUpdateCosmosDB creates a new cosmos database account
	CreateOrUpdateCosmosDB(ctx context.Context, groupName string, cosmosDBName string, location string, kind v1alpha1.CosmosDBKind, dbType v1alpha1.CosmosDBDatabaseAccountOfferType, tags map[string]*string) (*documentdb.DatabaseAccount, *errhelp.AzureError)

	// GetCosmosDB gets a cosmos database account
	GetCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (*documentdb.DatabaseAccount, *errhelp.AzureError)

	// DeleteCosmosDB removes the cosmos database account
	DeleteCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (*documentdb.DatabaseAccountsDeleteFuture, *errhelp.AzureError)

	// CheckNameExistsCosmosDB check if the account name already exists globally
	CheckNameExistsCosmosDB(ctx context.Context, accountName string) (bool, error)
}
