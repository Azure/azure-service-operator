// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package mysqluser

import (
	"context"
	"database/sql"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

// MySQLUser interface provides the functions of mySQLUser
type MySQLUser interface {
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (mysql.Database, error)
	ConnectToMySqlDb(ctx context.Context, drivername string, server string, dbname string, port int, username string, password string) (*sql.DB, error)
	GrantUserRoles(ctx context.Context, user string, database string, roles []string, db *sql.DB) error
	CreateUser(ctx context.Context, server string, secret map[string][]byte, db *sql.DB) (string, error)
	UserExists(ctx context.Context, db *sql.DB, username string) (bool, error)
	DropUser(ctx context.Context, db *sql.DB, user string) error
	DeleteSecrets(ctx context.Context, instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) (bool, error)
	GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) map[string][]byte

	// also embed methods from AsyncClient
	resourcemanager.ARMClient
}
