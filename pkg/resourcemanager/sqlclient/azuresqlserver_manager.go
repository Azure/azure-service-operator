// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
)

type AzureSqlServerManager interface {
	CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, properties SQLServerProperties) (result sql.Server, err error)
	DeleteSQLServer(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error)
	GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
}
