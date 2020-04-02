// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"context"
	"net/http"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

type PSQLDatabaseClient struct {
}

func NewPSQLDatabaseClient() *PSQLDatabaseClient {
	return &PSQLDatabaseClient{}
}

func getPSQLDatabasesClient() (psql.DatabasesClient, error) {
	databasesClient := psql.NewDatabasesClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return psql.DatabasesClient{}, err
	}
	databasesClient.Authorizer = a
	databasesClient.AddToUserAgent(config.UserAgent())
	return databasesClient, err
}

func getPSQLCheckNameAvailabilityClient() (psql.CheckNameAvailabilityClient, error) {
	nameavailabilityClient := psql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return psql.CheckNameAvailabilityClient{}, err
	}
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient, err
}

func (p *PSQLDatabaseClient) CheckDatabaseNameAvailability(ctx context.Context, databasename string) (bool, error) {

	client, err := getPSQLCheckNameAvailabilityClient()
	if err != nil {
		return false, err
	}

	resourceType := "database"

	nameAvailabilityRequest := psql.NameAvailabilityRequest{
		Name: &databasename,
		Type: &resourceType,
	}
	_, err = client.Execute(ctx, nameAvailabilityRequest)
	if err == nil { // Name available
		return true, nil
	}
	return false, err

}

func (p *PSQLDatabaseClient) CreateDatabaseIfValid(ctx context.Context, databasename string, servername string, resourcegroup string) (*http.Response, error) {

	client, err := getPSQLDatabasesClient()
	if err != nil {
		return &http.Response{
			StatusCode: 500,
		}, err
	}

	// Check if name is valid if this is the first create call
	valid, err := p.CheckDatabaseNameAvailability(ctx, databasename)
	if valid == false {
		return &http.Response{
			StatusCode: 500,
		}, err
	}

	dbParameters := psql.Database{}

	future, err := client.CreateOrUpdate(
		ctx,
		resourcegroup,
		servername,
		databasename,
		dbParameters,
	)
	if err != nil {
		return &http.Response{
			StatusCode: 500,
		}, err
	}

	return future.GetResult(client)
}

func (p *PSQLDatabaseClient) DeleteDatabase(ctx context.Context, databasename string, servername string, resourcegroup string) (status string, err error) {

	client, err := getPSQLDatabasesClient()
	if err != nil {
		return "", err
	}

	_, err = client.Get(ctx, resourcegroup, servername, databasename)
	if err == nil { // db present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername, databasename)
		return future.Status(), err
	}

	// db not present so return success anyway
	return "db not present", nil
}

func (p *PSQLDatabaseClient) GetDatabase(ctx context.Context, resourcegroup string, servername string, databasename string) (db psql.Database, err error) {

	client, err := getPSQLDatabasesClient()
	if err != nil {
		return psql.Database{}, err
	}

	return client.Get(ctx, resourcegroup, servername, databasename)
}
