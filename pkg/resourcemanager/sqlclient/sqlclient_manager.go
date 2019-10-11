// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
)

var AzureSQLManager SQLManager = &azureSqlManager{}

// SQLManager interface
type SQLManager interface {
	CreateOrUpdateSQLServer(sdkClient GoSDKClient, properties SQLServerProperties) (result sql.Server, err error)
	DeleteSQLServer(sdkClient GoSDKClient) (result autorest.Response, err error)
	GetServer(sdkClient GoSDKClient) (result sql.Server, err error)
	DeleteSQLFirewallRule(sdkClient GoSDKClient, ruleName string) (err error)
	DeleteDB(sdkClient GoSDKClient, databaseName string) (result autorest.Response, err error)
	GetSQLFirewallRule(sdkClient GoSDKClient, ruleName string) (result sql.FirewallRule, err error)
	GetDB(sdkClient GoSDKClient, databaseName string) (sql.Database, error)
	CreateOrUpdateDB(sdkClient GoSDKClient, properties SQLDatabaseProperties) (sql.DatabasesCreateOrUpdateFuture, error)
	CreateOrUpdateSQLFirewallRule(sdkClient GoSDKClient, ruleName string, startIP string, endIP string) (result bool, err error)
}
