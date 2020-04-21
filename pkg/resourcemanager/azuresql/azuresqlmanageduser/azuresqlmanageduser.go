// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqluser

import (
	"context"
	"database/sql"
	"fmt"

	azuresql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"k8s.io/apimachinery/pkg/runtime"

	_ "github.com/denisenkom/go-mssqldb"
	mssql "github.com/denisenkom/go-mssqldb"
)

// SqlServerPort is the default server port for sql server
const SqlServerPort = 1433

// DriverName is driver name for db connection
const DriverName = "sqlserver"

type AzureSqlManagedUserManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlManagedUserManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlManagedUserManager {
	return &AzureSqlManagedUserManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (s *AzureSqlManagedUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (azuresql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient()
	if err != nil {
		return azuresql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
		"serviceTierAdvisors, transparentDataEncryption",
	)
}

// ConnectToSqlDb connects to the SQL db using the current identity of operator (should be MI)
func (s *AzureSqlManagedUserManager) ConnectToSqlDbAsCurrentUser(ctx context.Context, drivername string, server string, database string, port int) (*sql.DB, error) {

	fullServerAddress := fmt.Sprintf("%s."+config.Environment().SQLDatabaseDNSSuffix, server)
	connString := fmt.Sprintf("server=%s;port=%d;database=%s;Persist Security Info=False;Pooling=False;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30", fullServerAddress, port, database)

	tokenProvider, err := getMSITokenProvider()
	if err != nil {
		return db, err
	}

	connector, err := mssql.NewAccessTokenConnector(connString, tokenProvider)
	if err != nil {
		return db, err
	}

	db := sql.OpenDB(connector)
	defer db.Close()

	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}

	return db, err
}

// EnableUserAndRoles creates user with secret credentials
func (s *AzureSqlManagedUserManager) EnableUserAndRoles(ctx context.Context, MIUserClientId string, roles []string, db *sql.DB) (string, error) {

	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD='%s'", newUser, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// UserExists checks if db contains user
func (s *AzureSqlManagedUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {
	res, err := db.ExecContext(
		ctx,
		"SELECT * FROM sysusers WHERE NAME=@user",
		sql.Named("user", username),
	)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// DropUser drops a user from db
func (s *AzureSqlManagedUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	tsql := "DROP USER @user"
	_, err := db.ExecContext(ctx, tsql, sql.Named("user", user))
	return err
}

func getMSITokenProvider() (func() (string, error), error) {
	msi, err := iam.GetMSITokenForResource("https://database.windows.net/")
	if err != nil {
		return nil, err
	}

	return func() (string, error) {
		msi.EnsureFresh()
		token := msi.OAuthToken()
		return token, nil
	}, nil
}
