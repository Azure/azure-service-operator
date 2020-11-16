// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	sql3 "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"

	"github.com/Azure/go-autorest/autorest/to"
)

type AzureSqlDbManager struct {
	creds config.Credentials
}

// Ensure we implement the interface we expect
var _ SqlDbManager = &AzureSqlDbManager{}

func NewAzureSqlDbManager(creds config.Credentials) *AzureSqlDbManager {
	return &AzureSqlDbManager{creds: creds}
}

// GetServer returns a SQL server
func (m *AzureSqlDbManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient, err := azuresqlshared.GetGoServersClient(m.creds)
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
func (m *AzureSqlDbManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient(m.creds)
	if err != nil {
		return sql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// DeleteDB deletes a DB
func (m *AzureSqlDbManager) DeleteDB(
	ctx context.Context,
	resourceGroupName string,
	serverName string,
	databaseName string) (future *sql.DatabasesDeleteFuture, err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := m.GetServer(ctx, resourceGroupName, serverName)
	if err != nil || *server.State != "Ready" {
		return nil, nil
	}

	// check to see if the db exists, if it doesn't then short-circuit
	_, err = m.GetDB(ctx, resourceGroupName, serverName, databaseName)
	if err != nil {
		return nil, nil
	}

	dbClient, err := azuresqlshared.GetGoDbClient(m.creds)
	if err != nil {
		return nil, err
	}

	result, err := dbClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)

	if err != nil {
		return nil, err
	}

	return &result, err
}

// CreateOrUpdateDB creates or updates a DB in Azure
func (m *AzureSqlDbManager) CreateOrUpdateDB(
	ctx context.Context,
	resourceGroupName string,
	location string,
	serverName string,
	tags map[string]*string,
	properties azuresqlshared.SQLDatabaseProperties) (string, *sql.Database, error) {

	dbClient, err := azuresqlshared.GetGoDbClient(m.creds)
	if err != nil {
		return "", nil, err
	}

	dbProp := azuresqlshared.SQLDatabasePropertiesToDatabase(properties)
	dbSku := azuresqlshared.SQLDatabasePropertiesToSku(properties)

	future, err := dbClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		properties.DatabaseName,
		sql.Database{
			Location:           to.StringPtr(location),
			DatabaseProperties: &dbProp,
			Tags:               tags,
			Sku:                dbSku,
		})

	if err != nil {
		return "", nil, err
	}

	result, err := future.Result(dbClient)

	return future.PollingURL(), &result, err
}

// AddLongTermRetention enables / disables long term retention
func (m *AzureSqlDbManager) AddLongTermRetention(ctx context.Context, resourceGroupName string, serverName string, databaseName string, weeklyRetention string, monthlyRetention string, yearlyRetention string, weekOfYear int32) (*http.Response, error) {

	longTermClient, err := azuresqlshared.GetBackupLongTermRetentionPoliciesClient(m.creds)
	// TODO: Probably shouldn't return a response at all in the err case here (all through this function)
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

	if err != nil {
		return &http.Response{
			StatusCode: 500,
		}, nil
	}

	return future.Response(), err
}
