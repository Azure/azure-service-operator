// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package psqluser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	_ "github.com/lib/pq" //the pg lib
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	psdatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

// PSqlServerPort is the default server port for sql server
const PSqlServerPort = 5432

// PDriverName is driver name for psqldb connection
const PDriverName = "postgres"

// PSecretUsernameKey is the username key in secret
const PSecretUsernameKey = "username"

// PSecretPasswordKey is the password key in secret
const PSecretPasswordKey = "password"

// PostgreSqlUserManager for psqluser manager
type PostgreSqlUserManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

// NewPostgreSqlUserManager creates a new PostgreSqlUserManager
func NewPostgreSqlUserManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *PostgreSqlUserManager {
	return &PostgreSqlUserManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (m *PostgreSqlUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (db psql.Database, err error) {
	dbClient, err := psdatabase.GetPSQLDatabasesClient(m.Creds)
	if err != nil {
		return psql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// ConnectToSqlDb connects to the PostgreSQL db using the given credentials
func (m *PostgreSqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, fullservername string, database string, port int, user string, password string) (*sql.DB, error) {

	connString := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=require connect_timeout=30", fullservername, user, password, port, database)

	db, err := sql.Open(drivername, connString)
	if err != nil {
		return db, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		return db, err
	}

	return db, err
}

// GrantUserRoles grants roles to a user for a given database
func (m *PostgreSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string

	if err := helpers.FindBadChars(user); err != nil {
		return errors.Wrap(err, "problem found with username")
	}

	for _, role := range roles {
		tsql := fmt.Sprintf("GRANT %s TO %q", role, user)

		if err := helpers.FindBadChars(role); err != nil {
			return errors.Wrap(err, "problem found with role")
		}

		_, err := db.ExecContext(ctx, tsql)
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// CreateUser creates user with secret credentials
func (m *PostgreSqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[PSecretUsernameKey])
	newPassword := string(secret[PSecretPasswordKey])

	// make an effort to prevent sql injection
	if err := helpers.FindBadChars(newUser); err != nil {
		return "", errors.Wrap(err, "problem found with username")
	}
	if err := helpers.FindBadChars(newPassword); err != nil {
		return "", errors.Wrap(err, "problem found with password")
	}

	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD '%s'", newUser, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// UpdateUser - Updates user password
func (m *PostgreSqlUserManager) UpdateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) error {
	user := string(secret[PSecretUsernameKey])
	newPassword := helpers.NewPassword()

	// make an effort to prevent sql injection
	if err := helpers.FindBadChars(user); err != nil {
		return errors.Wrap(err, "problem found with username")
	}
	if err := helpers.FindBadChars(newPassword); err != nil {
		return errors.Wrap(err, "problem found with password")
	}

	tsql := fmt.Sprintf("ALTER USER '%s' WITH PASSWORD '%s'", user, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	return err
}

// UserExists checks if db contains user
func (m *PostgreSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {

	res, err := db.ExecContext(ctx, "SELECT * FROM pg_user WHERE usename = $1", username)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err

}

// DropUser drops a user from db
func (m *PostgreSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	if err := helpers.FindBadChars(user); err != nil {
		return errors.Wrap(err, "problem found with username")
	}

	tsql := fmt.Sprintf("DROP USER IF EXISTS %q", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// DeleteSecrets deletes the secrets associated with a SQLUser
func (m *PostgreSqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) (bool, error) {

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	// delete standard user secret
	err := secretClient.Delete(
		ctx,
		secretKey,
	)
	if err != nil {
		instance.Status.Message = "failed to delete secret, err: " + err.Error()
		return false, err
	}

	return false, nil
}

// GetOrPrepareSecret gets or creates a secret
func (m *PostgreSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) map[string][]byte {
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	secret, err := secretClient.Get(ctx, secretKey)

	psqldbdnssuffix := "postgres.database.azure.com"
	if config.Environment().Name != "AzurePublicCloud" {
		psqldbdnssuffix = "postgres." + config.Environment().SQLDatabaseDNSSuffix
	}

	if err != nil {
		// @todo: find out whether this is an error due to non existing key or failed conn
		pw := helpers.NewPassword()
		return map[string][]byte{
			"username":                 []byte(""),
			"password":                 []byte(pw),
			"PSqlServerNamespace":      []byte(instance.Namespace),
			"PSqlServerName":           []byte(instance.Spec.Server),
			"fullyQualifiedServerName": []byte(instance.Spec.Server + "." + psqldbdnssuffix),
			"PSqlDatabaseName":         []byte(instance.Spec.DbName),
		}
	}

	return secret
}
