// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package cosmosdbs

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
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
	accountName string,
	spec v1alpha1.CosmosDBSpec,
	tags map[string]*string) (*documentdb.DatabaseAccount, string, error) {
	cosmosDBClient, err := getCosmosDBClient()
	if err != nil {
		return nil, "", err
	}

	createUpdateParams := documentdb.DatabaseAccountCreateUpdateParameters{
		Location: &spec.Location,
		Tags:     tags,
		Name:     &accountName,
		Kind:     documentdb.DatabaseAccountKind(spec.Kind),
		DatabaseAccountCreateUpdateProperties: &documentdb.DatabaseAccountCreateUpdateProperties{
			DatabaseAccountOfferType:      getAccountOfferType(spec),
			IsVirtualNetworkFilterEnabled: &spec.Properties.IsVirtualNetworkFilterEnabled,
			VirtualNetworkRules:           getVirtualNetworkRules(spec),
			EnableMultipleWriteLocations:  &spec.Properties.EnableMultipleWriteLocations,
			Locations:                     getLocations(spec),
			Capabilities:                  getCapabilities(spec),
			IPRangeFilter:                 getIPRangeFilter(spec),
		},
	}
	createUpdateFuture, err := cosmosDBClient.CreateOrUpdate(
		ctx, spec.ResourceGroup, accountName, createUpdateParams)
	if err != nil {
		// initial create request failed, wrap error
		return nil, "", err
	}

	result, err := createUpdateFuture.Result(cosmosDBClient)
	if err != nil {
		// there is no immediate result, wrap error
		return &result, createUpdateFuture.PollingURL(), err
	}
	return &result, createUpdateFuture.PollingURL(), nil
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

// ListConnectionStrings lists the connection strings for a database account
func (*AzureCosmosDBManager) ListConnectionStrings(
	ctx context.Context,
	groupName string,
	accountName string) (*documentdb.DatabaseAccountListConnectionStringsResult, error) {
	client, err := getCosmosDBClient()
	if err != nil {
		return nil, err
	}

	result, err := client.ListConnectionStrings(ctx, groupName, accountName)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func getAccountOfferType(spec v1alpha1.CosmosDBSpec) *string {
	kind := string(spec.Properties.DatabaseAccountOfferType)
	if kind == "" {
		kind = string(documentdb.Standard)
	}
	return &kind
}

func getLocations(spec v1alpha1.CosmosDBSpec) *[]documentdb.Location {
	if spec.Locations == nil || len(*spec.Locations) <= 1 {
		return &[]documentdb.Location{
			{
				LocationName:     to.StringPtr(spec.Location),
				FailoverPriority: to.Int32Ptr(0),
				IsZoneRedundant:  to.BoolPtr(false),
			},
		}
	}

	locations := make([]documentdb.Location, len(*spec.Locations))
	for i, l := range *spec.Locations {
		locations[i] = documentdb.Location{
			LocationName:     to.StringPtr(l.LocationName),
			FailoverPriority: to.Int32Ptr(l.FailoverPriority),
			IsZoneRedundant:  to.BoolPtr(l.IsZoneRedundant),
		}
	}
	return &locations
}

func getVirtualNetworkRules(spec v1alpha1.CosmosDBSpec) *[]documentdb.VirtualNetworkRule {
	if spec.VirtualNetworkRules == nil {
		return nil
	}

	vNetRules := make([]documentdb.VirtualNetworkRule, len(*spec.VirtualNetworkRules))
	for i, r := range *spec.VirtualNetworkRules {
		vNetRules[i] = documentdb.VirtualNetworkRule{
			ID:                               r.SubnetID,
			IgnoreMissingVNetServiceEndpoint: r.IgnoreMissingVNetServiceEndpoint,
		}
	}
	return &vNetRules
}

func getCapabilities(spec v1alpha1.CosmosDBSpec) *[]documentdb.Capability {
	capabilities := []documentdb.Capability{}
	if spec.Kind == v1alpha1.CosmosDBKindMongoDB && spec.Properties.MongoDBVersion == "3.6" {
		capabilities = []documentdb.Capability{
			{Name: to.StringPtr("EnableMongo")},
		}
	}
	if spec.Properties.Capabilities != nil {
		for _, i := range *spec.Properties.Capabilities {
			name := i.Name
			capabilities = append(capabilities, documentdb.Capability{
				Name: name,
			})
		}
	}
	return &capabilities
}

func getIPRangeFilter(spec v1alpha1.CosmosDBSpec) *string {
	sIPRules := ""
	if spec.IPRules != nil {
		sIPRules = strings.Join(*spec.IPRules, ",")
	}
	return &sIPRules
}
