// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"

	"github.com/Azure/go-autorest/autorest"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

// NewAzureCosmosDBManager creates a new cosmos db client
func NewAzureCosmosDBManager(creds config.Credentials, secretClient secrets.SecretClient) *AzureCosmosDBManager {
	return &AzureCosmosDBManager{
		Creds:        creds,
		SecretClient: secretClient,
	}
}

// NewARMClient returns a new manager (but as an ARMClient).
func NewARMClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) resourcemanager.ARMClient {
	return NewAzureCosmosDBManager(creds, secretClient)
}

// CosmosDBManager client functions
type CosmosDBManager interface {
	// CreateOrUpdateCosmosDB creates a new cosmos database account
	CreateOrUpdateCosmosDB(ctx context.Context, cosmosDBName string, spec v1alpha1.CosmosDBSpec, tags map[string]*string) (*documentdb.DatabaseAccount, string, error)

	// GetCosmosDB gets a cosmos database account
	GetCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (*documentdb.DatabaseAccount, error)

	// DeleteCosmosDB removes the cosmos database account
	DeleteCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (*autorest.Response, error)

	// CheckNameExistsCosmosDB check if the account name already exists globally
	CheckNameExistsCosmosDB(ctx context.Context, accountName string) (bool, error)

	// ListKeys lists the read & write keys for a database account
	ListKeys(ctx context.Context, groupName string, accountName string) (*documentdb.DatabaseAccountListKeysResult, error)

	// ListConnectionStrings lists the connection strings for a database account
	ListConnectionStrings(ctx context.Context, groupName string, accountName string) (*documentdb.DatabaseAccountListConnectionStringsResult, error)

	resourcemanager.ARMClient
}
