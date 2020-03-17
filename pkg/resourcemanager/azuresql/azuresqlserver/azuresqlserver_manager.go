// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlserver

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"

	"github.com/Azure/go-autorest/autorest"
)

type SqlServerManager interface {
	CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, tags map[string]*string, properties azuresqlshared.SQLServerProperties, forceUpdate bool) (result sql.Server, err error)
	DeleteSQLServer(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error)
	GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error)
	resourcemanager.ARMClient
}
