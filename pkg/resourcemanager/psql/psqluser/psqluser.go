// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package psqluser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	psdatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/psql/database"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	_ "github.com/lib/pq" //the pg lib
	"k8s.io/apimachinery/pkg/types"
)

// PSqlServerPort is the default server port for sql server
const PSqlServerPort = 5432

// PDriverName is driver name for psqldb connection
const PDriverName = "postgres"

// PSecretUsernameKey is the username key in secret
const PSecretUsernameKey = "username"

// PSecretPasswordKey is the password key in secret
const PSecretPasswordKey = "password"

//PostgreSqlUserManager for psqluser manager
type PostgreSqlUserManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

//NewPostgreSqlUserManager creates a new PostgreSqlUserManager
func NewPostgreSqlUserManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *PostgreSqlUserManager {
	return &PostgreSqlUserManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (s *PostgreSqlUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (db psql.Database, err error) {
	dbClient, err := psdatabase.GetPSQLDatabasesClient()
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
func (s *PostgreSqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, fullservername string, database string, port int, user string, password string) (*sql.DB, error) {

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
func (s *PostgreSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string

	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("Problem found with username: %v", err)
	}

	for _, role := range roles {
		tsql := fmt.Sprintf("GRANT %s TO %q", role, user)

		if err := helpers.FindBadChars(role); err != nil {
			return fmt.Errorf("Problem found with role: %v", err)
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
func (s *PostgreSqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[PSecretUsernameKey])
	newPassword := string(secret[PSecretPasswordKey])

	// make an effort to prevent sql injection
	if err := helpers.FindBadChars(newUser); err != nil {
		return "", fmt.Errorf("Problem found with username: %v", err)
	}
	if err := helpers.FindBadChars(newPassword); err != nil {
		return "", fmt.Errorf("Problem found with password: %v", err)
	}

	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD '%s'", newUser, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// UpdateUser - Updates user password
func (s *PostgreSqlUserManager) UpdateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) error {
	user := string(secret[PSecretUsernameKey])
	newPassword := helpers.NewPassword()

	// make an effort to prevent sql injection
	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("Problem found with username: %v", err)
	}
	if err := helpers.FindBadChars(newPassword); err != nil {
		return fmt.Errorf("Problem found with password: %v", err)
	}

	tsql := fmt.Sprintf("ALTER USER '%s' WITH PASSWORD '%s'", user, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	return err
}

// UserExists checks if db contains user
func (s *PostgreSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {

	res, err := db.ExecContext(ctx, "SELECT * FROM pg_user WHERE usename = $1", username)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err

}

// DropUser drops a user from db
func (s *PostgreSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("Problem found with username: %v", err)
	}

	tsql := fmt.Sprintf("DROP USER IF EXISTS %q", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// DeleteSecrets deletes the secrets associated with a SQLUser
func (s *PostgreSqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) (bool, error) {

	secretKey := GetNamespacedName(instance, secretClient)

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
func (s *PostgreSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) map[string][]byte {
	key := GetNamespacedName(instance, secretClient)

	secret, err := secretClient.Get(ctx, key)

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

// GetNamespacedName gets the namespaced-name
func GetNamespacedName(instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) types.NamespacedName {
	return types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
}
