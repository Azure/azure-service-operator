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
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

// AzureCosmosDBManager is the struct which contains helper functions for resource groups
type AzureCosmosDBManager struct{}

func getCosmosDBClient() (documentdb.DatabaseAccountsClient, error) {
	cosmosDBClient := documentdb.NewDatabaseAccountsClientWithBaseURI(config.BaseURI(), config.SubscriptionID())

	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		cosmosDBClient = documentdb.DatabaseAccountsClient{}
	} else {
		cosmosDBClient.Authorizer = a
		cosmosDBClient.AddToUserAgent(config.UserAgent())
	}

	return cosmosDBClient, err
}

// CreateOrUpdateCosmosDB creates a new CosmosDB
func (*AzureCosmosDBManager) CreateOrUpdateCosmosDB(
	ctx context.Context,
	accountName string,
	spec v1alpha1.CosmosDBSpec,
	tags map[string]*string) (*documentdb.DatabaseAccount, *errhelp.AzureError) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return nil, errhelp.NewAzureErrorAzureError(err)
	}

	groupName := spec.ResourceGroup
	kind := documentdb.DatabaseAccountKind(spec.Kind)
	primaryLocation := spec.Location

	// set optional defaults
	var accountOfferType *string
	if spec.Properties.DatabaseAccountOfferType == "" {
		accountOfferType = to.StringPtr(string(v1alpha1.CosmosDBDatabaseAccountOfferTypeStandard))
	} else {
		accountOfferType = to.StringPtr(string(spec.Properties.DatabaseAccountOfferType))
	}

	// capabilities are used to enable different API types
	var capabilities *[]documentdb.Capability

	// when MongoDBVersion is set to 3.6 we need to add the 'EnableMongo' capability, otherwise default to 3.2 behavior
	if kind == documentdb.MongoDB && spec.Properties.MongoDBVersion == "3.6" {
		capabilities = &[]documentdb.Capability{
			{Name: to.StringPtr("EnableMongo")},
		}
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
		ID:               to.StringPtr(fmt.Sprintf("%s-%s", accountName, primaryLocation)),
		FailoverPriority: to.Int32Ptr(0),
		LocationName:     to.StringPtr(primaryLocation),
	}
	locations := &[]documentdb.Location{
		locationObj,
	}

	createUpdateParams := documentdb.DatabaseAccountCreateUpdateParameters{
		Location: to.StringPtr(primaryLocation),
		Tags:     tags,
		Name:     to.StringPtr(accountName),
		Kind:     kind,
		Type:     to.StringPtr("Microsoft.DocumentDb/databaseAccounts"),
		DatabaseAccountCreateUpdateProperties: &documentdb.DatabaseAccountCreateUpdateProperties{
			DatabaseAccountOfferType:      accountOfferType,
			EnableMultipleWriteLocations:  to.BoolPtr(false),
			IsVirtualNetworkFilterEnabled: to.BoolPtr(false),
			Locations:                     locations,
			Capabilities:                  capabilities,
		},
	}
	createUpdateFuture, err := cosmosDBClient.CreateOrUpdate(
		ctx, groupName, accountName, createUpdateParams)

	if err != nil {
		// initial create request failed, wrap error
		return nil, errhelp.NewAzureErrorAzureError(err)
	}

	result, err := createUpdateFuture.Result(cosmosDBClient)
	if err != nil {
		// there is no immediate result, wrap error
		return &result, errhelp.NewAzureErrorAzureError(err)
	}
	return &result, nil
}

// GetCosmosDB gets the cosmos db account
func (*AzureCosmosDBManager) GetCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string) (*documentdb.DatabaseAccount, *errhelp.AzureError) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return nil, errhelp.NewAzureErrorAzureError(err)
	}

	result, err := cosmosDBClient.Get(ctx, groupName, cosmosDBName)
	if err != nil {
		return &result, errhelp.NewAzureErrorAzureError(err)
	}
	return &result, nil
}

// CheckNameExistsCosmosDB checks if the global account name already exists
func (*AzureCosmosDBManager) CheckNameExistsCosmosDB(
	ctx context.Context,
	accountName string) (bool, *errhelp.AzureError) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return false, errhelp.NewAzureErrorAzureError(err)
	}

	response, err := cosmosDBClient.CheckNameExists(ctx, accountName)
	if err != nil {
		return false, errhelp.NewAzureErrorAzureError(err)
	}

	switch response.StatusCode {
	case http.StatusNotFound:
		return false, nil
	case http.StatusOK:
		return true, nil
	default:
		return false, errhelp.NewAzureErrorAzureError(fmt.Errorf("unhandled status code for CheckNameExists"))
	}
}

// DeleteCosmosDB removes the resource group named by env var
func (*AzureCosmosDBManager) DeleteCosmosDB(
	ctx context.Context,
	groupName string,
	cosmosDBName string) (*autorest.Response, *errhelp.AzureError) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return nil, errhelp.NewAzureErrorAzureError(err)
	}

	deleteFuture, err := cosmosDBClient.Delete(ctx, groupName, cosmosDBName)
	if err != nil {
		return nil, errhelp.NewAzureErrorAzureError(err)
	}

	ar, err := deleteFuture.Result(cosmosDBClient)
	if err != nil {
		return nil, errhelp.NewAzureErrorAzureError(err)
	}
	return &ar, nil
}
