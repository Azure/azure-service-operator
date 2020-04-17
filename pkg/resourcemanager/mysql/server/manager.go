// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type MySQLServerManager interface {
	CreateServerIfValid(ctx context.Context, servername string, resourcegroup string, location string, tags map[string]*string, serverversion mysql.ServerVersion, sslenforcement mysql.SslEnforcementEnum, skuInfo mysql.Sku, adminlogin string, adminpassword string, createmode string, sourceserver string) (pollingURL string, server mysql.Server, err error)
	DeleteServer(ctx context.Context, resourcegroup string, servername string) (string, error)
	GetServer(ctx context.Context, resourcegroup string, servername string) (mysql.Server, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
