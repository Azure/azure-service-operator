// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.
package azuresqlshared

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
)

// getGoDbClient retrieves a DatabasesClient
func GetGoDbClient() sql.DatabasesClient {
	dbClient := sql.NewDatabasesClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	dbClient.Authorizer = a
	dbClient.AddToUserAgent(config.UserAgent())
	return dbClient
}

// getGoServersClient retrieves a ServersClient
func GetGoServersClient() sql.ServersClient {
	serversClient := sql.NewServersClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

// getGoFailoverGroupsClient retrieves a FailoverGroupsClient
func GetGoFailoverGroupsClient() sql.FailoverGroupsClient {
	failoverGroupsClient := sql.NewFailoverGroupsClient(config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	failoverGroupsClient.Authorizer = a
	failoverGroupsClient.AddToUserAgent(config.UserAgent())
	return failoverGroupsClient
}
