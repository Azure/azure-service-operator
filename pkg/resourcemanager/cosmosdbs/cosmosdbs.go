<<<<<<< HEAD
/*
MIT License

Copyright (c) Microsoft Corporation. All rights reserved.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
*/

=======
>>>>>>> Pull and merge changes from master into Azure-sql (#182)
package cosmosdbs

import (
	"context"
	"fmt"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	azurev1 "github.com/Azure/azure-service-operator/api/v1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/go-autorest/autorest/to"
)

func getCosmosDBClient() documentdb.DatabaseAccountsClient {
	cosmosDBClient := documentdb.NewDatabaseAccountsClient(config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		log.Fatalf("failed to initialize authorizer: %v\n", err)
	}
	cosmosDBClient.Authorizer = a
	cosmosDBClient.AddToUserAgent(config.UserAgent())
	return cosmosDBClient
}

// CreateCosmosDB creates a new CosmosDB
func CreateCosmosDB(ctx context.Context, groupName string,
	cosmosDBName string,
	location string,
	kind azurev1.CosmosDBKind,
	dbType azurev1.CosmosDBDatabaseAccountOfferType,
<<<<<<< HEAD
	tags map[string]*string) (*documentdb.DatabaseAccount, error) {
	cosmosDBClient := getCosmosDBClient()
=======
	tags map[string]*string) (documentdb.DatabaseAccount, error) {
	cosmosDBClient := getCosmosDBClient()

	log.Println("CosmosDB:CosmosDBName" + cosmosDBName)

	/* Uncomment and update if we should be checking for name exists first
	result, err = cosmosDBClient.CheckNameExists(ctx, cosmosDBName)
	if err != nil {
		return documentdb.DatabaseAccount.{}, err
	}
	result.
	if *result.NameAvailable == false {
		log.Fatalf("storage account not available: %v\n", result.Reason)
		return storage.Account{}, errors.New("storage account not available")
	}*/

>>>>>>> Pull and merge changes from master into Azure-sql (#182)
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
<<<<<<< HEAD
=======

>>>>>>> Pull and merge changes from master into Azure-sql (#182)
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
<<<<<<< HEAD
	future, err := cosmosDBClient.CreateOrUpdate(
		ctx, groupName, cosmosDBName, createUpdateParams)

	if err != nil {
		return nil, err
	}

	err = future.WaitForCompletionRef(ctx, cosmosDBClient.Client)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(cosmosDBClient)
	return &result, err
=======

	log.Println(fmt.Sprintf("creating cosmosDB '%s' in resource group '%s' and location: %v", cosmosDBName, groupName, location))

	future, err := cosmosDBClient.CreateOrUpdate(
		ctx, groupName, cosmosDBName, createUpdateParams)
	if err != nil {
		log.Println(fmt.Sprintf("ERROR creating cosmosDB '%s' in resource group '%s' and location: %v", cosmosDBName, groupName, location))
		log.Println(fmt.Printf("failed to initialize cosmosdb: %v\n", err))
	}
	return future.Result(cosmosDBClient)
>>>>>>> Pull and merge changes from master into Azure-sql (#182)
}

// DeleteCosmosDB removes the resource group named by env var
func DeleteCosmosDB(ctx context.Context, groupName string, cosmosDBName string) (result documentdb.DatabaseAccountsDeleteFuture, err error) {
	cosmosDBClient := getCosmosDBClient()
	return cosmosDBClient.Delete(ctx, groupName, cosmosDBName)
}
