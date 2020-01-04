package database

import (
	"context"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type PostgreSQLDatabaseManager interface {
	//convert(obj runtime.Object) (*v1alpha1.PostgreSQLDatabase, error)

	//CheckDatabaseNameAvailability(ctx context.Context, databasename string) (bool, error)
	CreateDatabaseIfValid(ctx context.Context, databasename string, servername string, resourcegroup string) (psql.DatabasesCreateOrUpdateFuture, error)
	DeleteDatabase(ctx context.Context, databasename string, servername string, resourcegroup string) (string, error)
	GetDatabase(ctx context.Context, resourcegroup string, servername string, database string) (psql.Database, error)
	// also embed async client methods
	resourcemanager.ARMClient
}
