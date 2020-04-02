// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package database

import (
	"context"
	"net/http"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PostgreSQLDatabaseManager interface {
	CheckDatabaseNameAvailability(ctx context.Context,
		databasename string) (bool, error)

	CreateDatabaseIfValid(ctx context.Context,
		databasename string,
		servername string,
		resourcegroup string) (*http.Response, error)

	DeleteDatabase(ctx context.Context,
		databasename string,
		servername string,
		resourcegroup string) (string, error)

	GetDatabase(ctx context.Context,
		resourcegroup string,
		servername string,
		database string) (psql.Database, error)

	// also embed async client methods
	resourcemanager.ARMClient
}
