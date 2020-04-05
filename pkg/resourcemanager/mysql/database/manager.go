// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"context"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type MySQLDatabaseManager interface {
	//convert(obj runtime.Object) (*v1alpha1.PostgreSQLDatabase, error)

	//CheckDatabaseNameAvailability(ctx context.Context, databasename string) (bool, error)
	CreateDatabaseIfValid(ctx context.Context, databasename string, servername string, resourcegroup string) (mysql.DatabasesCreateOrUpdateFuture, error)
	DeleteDatabase(ctx context.Context, databasename string, servername string, resourcegroup string) (string, error)
	GetDatabase(ctx context.Context, resourcegroup string, servername string, database string) (mysql.Database, error)
	// also embed async client methods
	resourcemanager.ARMClient
}
