// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
)

// AzureCosmosDBManager is the struct which contains helper functions for resource groups
type AzureCosmosDBManager struct{}

func getCosmosDBClient() documentdb.DatabaseAccountsClient {
	cosmosDBClient := documentdb.NewDatabaseAccountsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		// should we have Fatalf's in our code?
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	cosmosDBClient.Authorizer = a
	cosmosDBClient.AddToUserAgent(config.UserAgent())
	return cosmosDBClient
}

// CreateCosmosDB creates a new CosmosDB
func (*AzureCosmosDBManager) CreateCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string,
	location string,
	kind v1alpha1.CosmosDBKind,
	dbType v1alpha1.CosmosDBDatabaseAccountOfferType,
	tags map[string]*string) (*documentdb.DatabaseAccount, error) {
	cosmosDBClient := getCosmosDBClient()

	dbKind := documentdb.DatabaseAccountKind(kind)
	sDBType := string(dbType)

	/*
	*   Current state of Locations and CosmosDB properties:
	*   Creating a Database account with CosmosDB requires
	*   that DatabaseAccountCreateUpdateProperties be sent over
	*   and currently we are not reading most of these values in
	*   as part of the Spec for CosmosDB.  We are currently
	*   specifying a single Location as part of a location array
	*   which matches the location set for the overall CosmosDB
	*   instance.  This matches the general behavior of creating
	*   a CosmosDB instance in the portal where the only
	*   geo-relicated region is the sole region the CosmosDB
	*   is created in.
	 */
	locationObj := documentdb.Location{
		ID:               to.StringPtr(fmt.Sprintf("%s-%s", cosmosDBName, location)),
		FailoverPriority: to.Int32Ptr(0),
		LocationName:     to.StringPtr(location),
	}

	locationsArray := []documentdb.Location{
		locationObj,
	}
	createUpdateParams := documentdb.DatabaseAccountCreateUpdateParameters{
		Location: to.StringPtr(location),
		Tags:     tags,
		Name:     &cosmosDBName,
		Kind:     dbKind,
		Type:     to.StringPtr("Microsoft.DocumentDb/databaseAccounts"),
		ID:       &cosmosDBName,
		DatabaseAccountCreateUpdateProperties: &documentdb.DatabaseAccountCreateUpdateProperties{
			DatabaseAccountOfferType:      &sDBType,
			EnableMultipleWriteLocations:  to.BoolPtr(false),
			IsVirtualNetworkFilterEnabled: to.BoolPtr(false),
			Locations:                     &locationsArray,
		},
	}
	future, err := cosmosDBClient.CreateOrUpdate(
		ctx, groupName, cosmosDBName, createUpdateParams)

	if err != nil {
		return nil, err
	}

	//TODO: this seems synchronous?
	err = future.WaitForCompletionRef(ctx, cosmosDBClient.Client)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(cosmosDBClient)
	return &result, err
}

// GetCosmosDB gets the cosmos db account
func (*AzureCosmosDBManager) GetCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string) (result documentdb.DatabaseAccount, err error) {
	cosmosDBClient := getCosmosDBClient()

	return cosmosDBClient.Get(ctx, groupName, cosmosDBName)
}

// DeleteCosmosDB removes the resource group named by env var
func (*AzureCosmosDBManager) DeleteCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string) (result documentdb.DatabaseAccountsDeleteFuture, err error) {
	cosmosDBClient := getCosmosDBClient()
	return cosmosDBClient.Delete(ctx, groupName, cosmosDBName)
}
