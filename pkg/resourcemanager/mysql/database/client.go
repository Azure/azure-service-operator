// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"context"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// MySQLDatabaseClient struct
type MySQLDatabaseClient struct {
	creds config.Credentials
}

// NewMySQLDatabaseClient create a new MySQLDatabaseClient
func NewMySQLDatabaseClient(creds config.Credentials) *MySQLDatabaseClient {
	return &MySQLDatabaseClient{creds: creds}
}

// GetMySQLDatabasesClient return the mysqldatabaseclient
func GetMySQLDatabasesClient(creds config.Credentials) mysql.DatabasesClient {
	databasesClient := mysql.NewDatabasesClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	databasesClient.Authorizer = a
	databasesClient.AddToUserAgent(config.UserAgent())
	return databasesClient
}

func getMySQLCheckNameAvailabilityClient(creds config.Credentials) mysql.CheckNameAvailabilityClient {
	nameavailabilityClient := mysql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient
}

func (m *MySQLDatabaseClient) CheckDatabaseNameAvailability(ctx context.Context, databasename string) (bool, error) {

	client := getMySQLCheckNameAvailabilityClient(m.creds)

	resourceType := "Microsoft.DBforMySQL/servers/databases"

	nameAvailabilityRequest := mysql.NameAvailabilityRequest{
		Name: &databasename,
		Type: &resourceType,
	}
	_, err := client.Execute(ctx, nameAvailabilityRequest)
	if err == nil { // Name available
		return true, nil
	}
	return false, err

}

func (m *MySQLDatabaseClient) CreateDatabaseIfValid(ctx context.Context, databasename string, servername string, resourcegroup string) (future mysql.DatabasesCreateOrUpdateFuture, err error) {

	client := GetMySQLDatabasesClient(m.creds)

	// Check if name is valid if this is the first create call
	valid, err := m.CheckDatabaseNameAvailability(ctx, databasename)
	if valid == false {
		return future, err
	}

	dbParameters := mysql.Database{}

	future, err = client.CreateOrUpdate(
		ctx,
		resourcegroup,
		servername,
		databasename,
		dbParameters,
	)

	return future, err
}

func (m *MySQLDatabaseClient) DeleteDatabase(ctx context.Context, databasename string, servername string, resourcegroup string) (status string, err error) {

	client := GetMySQLDatabasesClient(m.creds)

	_, err = client.Get(ctx, resourcegroup, servername, databasename)
	if err == nil { // db present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername, databasename)
		return future.Status(), err
	}
	// db not present so return success anyway
	return "db not present", nil

}

func (m *MySQLDatabaseClient) GetDatabase(ctx context.Context, resourcegroup string, servername string, databasename string) (db mysql.Database, err error) {

	client := GetMySQLDatabasesClient(m.creds)

	return client.Get(ctx, resourcegroup, servername, databasename)
}
