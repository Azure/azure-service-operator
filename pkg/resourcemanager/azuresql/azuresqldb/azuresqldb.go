// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqldb

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
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
func (m *AzureSqlDbManager) GetServer(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient, err := azuresqlshared.GetGoServersClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
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
func (m *AzureSqlDbManager) GetDB(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
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
	subscriptionID string,
	resourceGroupName string,
	serverName string,
	databaseName string) (future *sql.DatabasesDeleteFuture, err error) {

	// check to see if the server exists, if it doesn't then short-circuit
	server, err := m.GetServer(ctx, subscriptionID, resourceGroupName, serverName)
	if err != nil || *server.State != "Ready" {
		return nil, ignoreNotFound(err)
	}

	// check to see if the db exists, if it doesn't then short-circuit
	db, err := m.GetDB(ctx, subscriptionID, resourceGroupName, serverName, databaseName)
	if err != nil {
		return nil, ignoreNotFound(err)
	}
	if db.Name == nil {
		// The database doesn't exist, we don't need to delete it.
		return nil, nil
	}

	dbClient, err := azuresqlshared.GetGoDbClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
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
	subscriptionID string,
	resourceGroupName string,
	location string,
	serverName string,
	tags map[string]*string,
	properties azuresqlshared.SQLDatabaseProperties) (string, *sql.Database, error) {

	dbClient, err := azuresqlshared.GetGoDbClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
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
func (m *AzureSqlDbManager) AddLongTermRetention(
	ctx context.Context,
	subscriptionID string,
	resourceGroupName string,
	serverName string,
	databaseName string,
	policy azuresqlshared.SQLDatabaseBackupLongTermRetentionPolicy) (*sql.BackupLongTermRetentionPoliciesCreateOrUpdateFuture, error) {

	longTermClient, err := azuresqlshared.GetBackupLongTermRetentionPoliciesClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
	if err != nil {
		return nil, err
	}

	// validate the input and exit if nothing needs to happen - this is ok!
	if policy.WeeklyRetention == "" && policy.MonthlyRetention == "" && policy.YearlyRetention == "" {
		return nil, nil
	}

	// validate the pairing of yearly retention and week of year
	if policy.YearlyRetention != "" && (policy.WeekOfYear <= 0 || policy.WeekOfYear > 52) {
		return nil, fmt.Errorf("weekOfYear must be greater than 0 and less or equal to 52 when yearlyRetention is used")
	}

	// create pointers so that we can pass nils if needed
	pWeeklyRetention := &policy.WeeklyRetention
	if policy.WeeklyRetention == "" {
		pWeeklyRetention = nil
	}
	pMonthlyRetention := &policy.MonthlyRetention
	if policy.MonthlyRetention == "" {
		pMonthlyRetention = nil
	}
	pYearlyRetention := &policy.YearlyRetention
	pWeekOfYear := &policy.WeekOfYear
	if policy.YearlyRetention == "" {
		pYearlyRetention = nil
		pWeekOfYear = nil
	}

	future, err := longTermClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		sql.BackupLongTermRetentionPolicy{
			LongTermRetentionPolicyProperties: &sql.LongTermRetentionPolicyProperties{
				WeeklyRetention:  pWeeklyRetention,
				MonthlyRetention: pMonthlyRetention,
				YearlyRetention:  pYearlyRetention,
				WeekOfYear:       pWeekOfYear,
			},
		},
	)

	if err != nil {
		return nil, err
	}

	return &future, err
}

func (m *AzureSqlDbManager) AddShortTermRetention(
	ctx context.Context,
	subscriptionID string,
	resourceGroupName string,
	serverName string,
	databaseName string,
	policy *v1beta1.SQLDatabaseShortTermRetentionPolicy) (*sql.BackupShortTermRetentionPoliciesCreateOrUpdateFuture, error) {

	client, err := azuresqlshared.GetBackupShortTermRetentionPoliciesClient(azuresqlshared.GetSubscriptionCredentials(m.creds, subscriptionID))
	if err != nil {
		return nil, errors.Wrapf(err, "couldn't create BackupShortTermRetentionPoliciesClient")
	}

	var policyProperties *sql.BackupShortTermRetentionPolicyProperties
	if policy == nil {
		// If policy is nil we're in a bit of an awkward situation since we cannot know if the customer has mutated
		// the retention policy in a previous reconciliation loop and then subsequently removed it. If they have,
		// "doing nothing" here is wrong because that leaves them in the previous modified state (but with no reflection
		// of that fact in the Spec).
		// Unfortunately you cannot update the retention policy to nil, nor can you delete it, so we must awkwardly
		// set it back to its default configuration.
		// Note: There are risks here, such as if the default on the server and the default in our code drift apart
		// at some point in the future.
		policyProperties = &sql.BackupShortTermRetentionPolicyProperties{
			RetentionDays: to.Int32Ptr(7), // 7 is the magical default as of Jan 2021
		}
	} else {
		policyProperties = &sql.BackupShortTermRetentionPolicyProperties{
			RetentionDays: to.Int32Ptr(policy.RetentionDays),
		}
	}

	future, err := client.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		sql.BackupShortTermRetentionPolicy{
			BackupShortTermRetentionPolicyProperties: policyProperties,
		})

	if err != nil {
		return nil, err
	}

	return &future, err
}

var goneCodes = []string{
	errhelp.ResourceGroupNotFoundErrorCode,
	errhelp.ParentNotFoundErrorCode,
	errhelp.NotFoundErrorCode,
	errhelp.ResourceNotFound,
}

func isNotFound(azerr *errhelp.AzureError) bool {
	return helpers.ContainsString(goneCodes, azerr.Type)
}

var incompleteCodes = []string{
	errhelp.AsyncOpIncompleteError,
}

func isIncompleteOp(azerr *errhelp.AzureError) bool {
	return helpers.ContainsString(incompleteCodes, azerr.Type)
}

func ignoreNotFound(err error) error {
	if isNotFound(errhelp.NewAzureError(err)) {
		return nil
	}
	return err
}
