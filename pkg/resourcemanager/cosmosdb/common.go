// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdb

import (
	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2021-03-15/documentdb"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

func GetCosmosDBAccountClient(creds config.Credentials) (documentdb.DatabaseAccountsClient, error) {
	client := documentdb.NewDatabaseAccountsClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return documentdb.DatabaseAccountsClient{}, err
	}

	client.Authorizer = a

	err = client.AddToUserAgent(config.UserAgent())
	return client, err
}

func GetCosmosDBSQLDatabaseClient(creds config.Credentials) (documentdb.SQLResourcesClient, error) {
	client := documentdb.NewSQLResourcesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return documentdb.SQLResourcesClient{}, err
	}

	client.Authorizer = a

	err = client.AddToUserAgent(config.UserAgent())
	return client, err
}
