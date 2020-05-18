// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package psqluser

import (
	"context"
	"database/sql"

	psql "github.com/Azure/azure-sdk-for-go/services/preview/postgresql/mgmt/2017-12-01-preview/postgresql"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type PSqlUserManager interface {
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (psql.Database, error)
	ConnectToPostgreSqlDb(ctx context.Context, drivername string, server string, dbname string, port int, username string, password string) (*sql.DB, error)
	GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error
	CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error)
	UserExists(ctx context.Context, db *sql.DB, username string) (bool, error)
	DropUser(ctx context.Context, db *sql.DB, user string) error
	DeleteSecrets(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) (bool, error)
	GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) map[string][]byte

	// also embed methods from AsyncClient
	resourcemanager.ARMClient
}
