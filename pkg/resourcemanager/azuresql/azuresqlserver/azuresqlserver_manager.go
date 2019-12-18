// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresqlserver

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
)

func NewAzureSqlServerManager(log logr.Logger) *AzureSqlServerManager {
	return &AzureSqlServerManager{Log: log}
}

type SqlServerManager interface {
	CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, properties azuresqlshared.SQLServerProperties) (result sql.Server, err error)
	DeleteSQLServer(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error)
	GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
}
