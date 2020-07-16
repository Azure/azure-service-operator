// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlfailovergroup

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/go-autorest/autorest"
)

type SqlFailoverGroupManager interface {
	CreateOrUpdateFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string, properties azuresqlshared.SQLFailoverGroupProperties) (result sql.FailoverGroupsCreateOrUpdateFuture, err error)
	DeleteFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failoverGroupName string) (result autorest.Response, err error)
	GetFailoverGroup(ctx context.Context, resourceGroupName string, serverName string, failovergroupname string) (sql.FailoverGroup, error)
	GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (sql.Database, error)
	resourcemanager.ARMClient
}
