// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// AzureCosmosDBManager is the struct which contains helper functions for resource groups
type AzureCosmosDBManager struct {
	SecretClient secrets.SecretClient
}

func getCosmosDBClient() (documentdb.DatabaseAccountsClient, error) {
	cosmosDBClient := documentdb.NewDatabaseAccountsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		cosmosDBClient = documentdb.DatabaseAccountsClient{}
	} else {
		cosmosDBClient.Authorizer = a
	}

	err = cosmosDBClient.AddToUserAgent(config.UserAgent())
	return cosmosDBClient, err
}

// CreateOrUpdateCosmosDB creates a new CosmosDB
func (*AzureCosmosDBManager) CreateOrUpdateCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string,
	location string,
	kind v1alpha1.CosmosDBKind,
	properties v1alpha1.CosmosDBProperties,
	tags map[string]*string) (*documentdb.DatabaseAccount, error) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return nil, err
	}

	dbKind := documentdb.DatabaseAccountKind(kind)
	sDBType := string(properties.DatabaseAccountOfferType)
	bWriteLocal := bool(properties.EnableMultipleWriteLocations)

	var capabilities []documentdb.Capability
	if dbKind == documentdb.MongoDB && properties.MongoDBVersion == "3.6" {
		capabilities = []documentdb.Capability{
			{Name: to.StringPtr("EnableMongo")},
		}
	} else {
		capabilities = make([]documentdb.Capability, 0)
	}

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
			EnableMultipleWriteLocations:  &bWriteLocal,
			IsVirtualNetworkFilterEnabled: to.BoolPtr(false),
			Locations:                     &locationsArray,
			Capabilities:                  &capabilities,
		},
	}
	createUpdateFuture, err := cosmosDBClient.CreateOrUpdate(
		ctx, groupName, cosmosDBName, createUpdateParams)

	if err != nil {
		// initial create request failed, wrap error
		return nil, err
	}

	result, err := createUpdateFuture.Result(cosmosDBClient)
	if err != nil {
		// there is no immediate result, wrap error
		return &result, err
	}
	return &result, nil
}

// GetCosmosDB gets the cosmos db account
func (*AzureCosmosDBManager) GetCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string) (*documentdb.DatabaseAccount, error) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return nil, err
	}

	result, err := cosmosDBClient.Get(ctx, groupName, cosmosDBName)
	if err != nil {
		return &result, err
	}
	return &result, nil
}

// CheckNameExistsCosmosDB checks if the global account name already exists
func (*AzureCosmosDBManager) CheckNameExistsCosmosDB(
	ctx context.Context,
	accountName string) (bool, error) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return false, err
	}

	response, err := cosmosDBClient.CheckNameExists(ctx, accountName)
	if err != nil {
		return false, err
	}

	switch response.StatusCode {
	case http.StatusNotFound:
		return false, nil
	case http.StatusOK:
		return true, nil
	default:
		return false, fmt.Errorf("unhandled status code for CheckNameExists")
	}
}

// DeleteCosmosDB removes the resource group named by env var
func (*AzureCosmosDBManager) DeleteCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string) (*autorest.Response, error) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return nil, err
	}

	deleteFuture, err := cosmosDBClient.Delete(ctx, groupName, cosmosDBName)
	if err != nil {
		return nil, err
	}

	ar, err := deleteFuture.Result(cosmosDBClient)
	if err != nil {
		return nil, err
	}
	return &ar, nil
}

// ListKeys lists the read & write keys for a database account
func (*AzureCosmosDBManager) ListKeys(
	ctx context.Context,
	groupName string,
	accountName string) (*documentdb.DatabaseAccountListKeysResult, error) {
	client, err := getCosmosDBClient()
	if err != nil {
		return nil, err
	}

	result, err := client.ListKeys(ctx, groupName, accountName)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
