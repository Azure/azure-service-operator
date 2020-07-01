// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	sql3 "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
)

type AzureSqlDbManager struct {
}

func NewAzureSqlDbManager() *AzureSqlDbManager {
	return &AzureSqlDbManager{}
}

// GetServer returns a SQL server
func (_ *AzureSqlDbManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient, err := azuresqlshared.GetGoServersClient()
	if err != nil {
		return sql.Server{}, err
	}

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

func (_ *AzureSqlDbManager) GetServerWithCreds(ctx context.Context, resourceGroupName string, serverName string, creds map[string]string) (result sql.Server, err error) {
	serversClient, err := azuresqlshared.GetGoServersClientWithCreds(creds)
	if err != nil {
		return sql.Server{}, err
	}

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

// GetDB retrieves a database
func (_ *AzureSqlDbManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient()
	if err != nil {
		return sql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// GetDBWithCreds retrieves a database
func (_ *AzureSqlDbManager) GetDBWithCreds(ctx context.Context, resourceGroupName string, serverName string, databaseName string, creds map[string]string) (sql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClientWithCreds(creds)
	if err != nil {
		return sql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// DeleteDB deletes a DB
func (sdk *AzureSqlDbManager) DeleteDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServer(ctx, resourceGroupName, serverName)
	if err != nil || *server.State != "Ready" {
		return result, nil
	}

	// check to see if the db exists, if it doesn't then short-circuit
	_, err = sdk.GetDB(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		return result, nil
	}

	dbClient, err := azuresqlshared.GetGoDbClient()
	if err != nil {
		return result, err
	}

	result, err = dbClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)

	return result, err
}

// DeleteDB deletes a DB
func (sdk *AzureSqlDbManager) DeleteDBWithCreds(ctx context.Context, resourceGroupName string, serverName string, databaseName string, creds map[string]string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := sdk.GetServerWithCreds(ctx, resourceGroupName, serverName, creds)
	if err != nil || *server.State != "Ready" {
		return result, nil
	}

	// check to see if the db exists, if it doesn't then short-circuit
	_, err = sdk.GetDBWithCreds(ctx, resourceGroupName, serverName, databaseName, creds)
	if err != nil {
		return result, nil
	}

	dbClient, err := azuresqlshared.GetGoDbClientWithCreds(creds)
	if err != nil {
		return result, err
	}

	result, err = dbClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)

	return result, err
}

// CreateOrUpdateDB creates or updates a DB in Azure
func (_ *AzureSqlDbManager) CreateOrUpdateDB(ctx context.Context, resourceGroupName string, location string, serverName string, tags map[string]*string, properties azuresqlshared.SQLDatabaseProperties) (*http.Response, error) {
	dbClient, err := azuresqlshared.GetGoDbClient()
	if err != nil {
		return &http.Response{
			StatusCode: 0,
		}, err
	}

	dbProp := azuresqlshared.SQLDatabasePropertiesToDatabase(properties)

	future, err := dbClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(location),
			DatabaseProperties: &dbProp,
			Tags:               tags,
		})

	return future.Response(), err
}

// CreateOrUpdateDBWithCreds creates or updates a DB in Azure
func (*AzureSqlDbManager) CreateOrUpdateDBWithCreds(ctx context.Context, resourceGroupName string, location string, serverName string, tags map[string]*string, properties azuresqlshared.SQLDatabaseProperties, creds map[string]string) (*http.Response, error) {
	dbClient, err := azuresqlshared.GetGoDbClientWithCreds(creds)
	if err != nil {
		return &http.Response{
			StatusCode: 0,
		}, err
	}

	dbProp := azuresqlshared.SQLDatabasePropertiesToDatabase(properties)

	future, err := dbClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(location),
			DatabaseProperties: &dbProp,
			Tags:               tags,
		},
	)

	return future.Response(), err
}

// AddLongTermRetention enables / disables long term retention
func (_ *AzureSqlDbManager) AddLongTermRetention(ctx context.Context, resourceGroupName string, serverName string, databaseName string, weeklyRetention string, monthlyRetention string, yearlyRetention string, weekOfYear int32, creds map[string]string) (*http.Response, error) {

	longTermClient, err := azuresqlshared.GetBackupLongTermRetentionPoliciesClientWithCreds(creds)
	if err != nil {
		return &http.Response{
			StatusCode: 0,
		}, err
	}

	// validate the input and exit if nothing needs to happen - this is ok!
	if weeklyRetention == "" && monthlyRetention == "" && yearlyRetention == "" {
		return &http.Response{
			StatusCode: 200,
		}, nil
	}

	// validate the pairing of yearly retention and week of year
	if yearlyRetention != "" && (weekOfYear <= 0 || weekOfYear > 52) {
		return &http.Response{
			StatusCode: 500,
		}, fmt.Errorf("weekOfYear must be greater than 0 and less or equal to 52 when yearlyRetention is used")
	}

	// create pointers so that we can pass nils if needed
	pWeeklyRetention := &weeklyRetention
	if weeklyRetention == "" {
		pWeeklyRetention = nil
	}
	pMonthlyRetention := &monthlyRetention
	if monthlyRetention == "" {
		pMonthlyRetention = nil
	}
	pYearlyRetention := &yearlyRetention
	pWeekOfYear := &weekOfYear
	if yearlyRetention == "" {
		pYearlyRetention = nil
		pWeekOfYear = nil
	}

	future, err := longTermClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		sql3.BackupLongTermRetentionPolicy{
			LongTermRetentionPolicyProperties: &sql3.LongTermRetentionPolicyProperties{
				WeeklyRetention:  pWeeklyRetention,
				MonthlyRetention: pMonthlyRetention,
				YearlyRetention:  pYearlyRetention,
				WeekOfYear:       pWeekOfYear,
			},
		},
	)

	return future.Response(), err
}
