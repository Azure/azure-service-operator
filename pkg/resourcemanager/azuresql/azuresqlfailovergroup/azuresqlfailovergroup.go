// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresqlfailovergroup

import (
	"context"
	"net/http"

	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	sql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
)

type AzureSqlFailoverGroupManager struct {
	Log logr.Logger
}

func NewAzureSqlFailoverGroupManager(log logr.Logger) *AzureSqlFailoverGroupManager {
	return &AzureSqlFailoverGroupManager{Log: log}
}

// GetServer returns a SQL server
func (_ *AzureSqlFailoverGroupManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient := azuresqlshared.GetGoServersClient()

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

// GetDB retrieves a database
func (_ *AzureSqlFailoverGroupManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error) {
	dbClient := azuresqlshared.GetGoDbClient()

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// GetFailoverGroup retrieves a failover group
func (_ *AzureSqlFailoverGroupManager) GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (sql.FailoverGroup, error) {
	failoverGroupsClient := azuresqlshared.GetGoFailoverGroupsClient()

	return failoverGroupsClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		failovergroupname,
	)
}

// DeleteFailoverGroup deletes a failover group
func (sdk *AzureSqlFailoverGroupManager) DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error) {

	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}
	// check to see if the server exists, if it doesn't then short-circuit
	_, err = sdk.GetServer(ctx, resourceGroupName, serverName)
	if err != nil {
		return result, nil
	}
	// check to see if the failover group exists, if it doesn't then short-circuit
	_, err = sdk.GetFailoverGroup(ctx, resourceGroupName, serverName, failoverGroupName)
	if err != nil {
		return result, nil
	}
	failoverGroupsClient := azuresqlshared.GetGoFailoverGroupsClient()
	future, err := failoverGroupsClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
		failoverGroupName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(failoverGroupsClient)
}

// CreateOrUpdateFailoverGroup creates a failover group
func (sdk *AzureSqlFailoverGroupManager) CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties azuresqlshared.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error) {
	failoverGroupsClient := azuresqlshared.GetGoFailoverGroupsClient()

	// Construct a PartnerInfo object from the server name
	// Get resource ID from the servername to use

	server, err := sdk.GetServer(ctx, properties.SecondaryServerResourceGroup, properties.SecondaryServerName)
	if err != nil {
		return result, nil
	}

	secServerResourceID := server.ID
	partnerServerInfo := sql.PartnerInfo{
		ID:              secServerResourceID,
		ReplicationRole: sql.Secondary,
	}

	partnerServerInfoArray := []sql.PartnerInfo{partnerServerInfo}

	var databaseIDArray []string

	// Parse the Databases in the Databaselist and form array of Resource IDs
	for _, each := range properties.DatabaseList {
		database, err := sdk.GetDB(ctx, resourceGroupName, serverName, each)
		if err != nil {
			return result, err
		}
		databaseIDArray = append(databaseIDArray, *database.ID)
	}

	// Construct FailoverGroupProperties struct
	failoverGroupProperties := sql.FailoverGroupProperties{
		ReadWriteEndpoint: &sql.FailoverGroupReadWriteEndpoint{
			FailoverPolicy:                         azuresqlshared.TranslateFailoverPolicy(properties.FailoverPolicy),
			FailoverWithDataLossGracePeriodMinutes: &properties.FailoverGracePeriod,
		},
		PartnerServers: &partnerServerInfoArray,
		Databases:      &databaseIDArray,
	}

	failoverGroup := sql.FailoverGroup{
		FailoverGroupProperties: &failoverGroupProperties,
	}

	return failoverGroupsClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		failovergroupname,
		failoverGroup)
}
