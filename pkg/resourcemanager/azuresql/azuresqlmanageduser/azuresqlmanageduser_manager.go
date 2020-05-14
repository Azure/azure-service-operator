// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlmanageduser

import (
	"context"
	"database/sql"

	azuresql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type SqlManagedUserManager interface {
	GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (azuresql.Database, error)
	ConnectToSqlDbAsCurrentUser(ctx context.Context, drivername string, server string, dbname string) (*sql.DB, error)

	DeleteSecrets(ctx context.Context, instance *v1alpha1.AzureSQLManagedUser, secretClient secrets.SecretClient) error
	UpdateSecret(ctx context.Context, instance *v1alpha1.AzureSQLManagedUser, secretClient secrets.SecretClient) error

	EnableUser(ctx context.Context, MIName string, MIUserClientId string, db *sql.DB) error
	GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error
	DropUser(ctx context.Context, db *sql.DB, user string) error
	UserExists(ctx context.Context, db *sql.DB, username string) (bool, error)
	// also embed methods from AsyncClient
	resourcemanager.ARMClient
}
