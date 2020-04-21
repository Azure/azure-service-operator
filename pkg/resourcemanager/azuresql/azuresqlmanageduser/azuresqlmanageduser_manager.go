// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqluser

import (
	"context"
	"database/sql"

	azuresql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
)

type SqlManagedUserManager interface {
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (azuresql.Database, error)
	ConnectToSqlDbAsCurrentUser(ctx context.Context, drivername string, server string, dbname string, port int) (*sql.DB, error)

	//DeleteSecrets(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) (bool, error)
	//GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) map[string][]byte

	EnableUserAndRoles(ctx context.Context, MIUserClientId string, roles []string, db *sql.DB) (string, error)
	DropUser(ctx context.Context, db *sql.DB, user string) error
	UserExists(ctx context.Context, db *sql.DB, username string) (bool, error)

	// also embed methods from AsyncClient
	resourcemanager.ARMClient
}
