// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package sqlclient

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/go-autorest/autorest"
	"github.com/go-logr/logr"
)

func NewAzureSqlFailoverGroupManager(log logr.Logger) *AzureSqlFailoverGroupManager {
	return &AzureSqlFailoverGroupManager{Log: log}
}

type AzureSqlFailoverGroupManager interface {
	CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error)
	DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error)
	GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (sql.FailoverGroup, error)
}
