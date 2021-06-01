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

	"github.com/Azure/azure-service-operator/api/v1alpha2"
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

// PSecretUsernameKey is the username key in secret
const PSecretFullyQualifiedUsernameKey = "fullyQualifiedUsername"

// PSecretPasswordKey is the password key in secret
const PSecretPasswordKey = "password"

//PostgreSqlUserManager for psqluser manager
type PostgreSqlUserManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

//NewPostgreSqlUserManager creates a new PostgreSqlUserManager
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

func (m *PostgreSqlUserManager) GetDbOwnerRole(dbName string) string {
	return dbName
}

// EnsureUserOwnsDatabases adds the user to the database owner role (and create a dedicated owner role if non exists)
func (m *PostgreSqlUserManager) EnsureUserOwnsDatabases(ctx context.Context, user string, ownedDatabases []string, adminUser string, db *sql.DB) error {
	var errorStrings []string

	// remove host from username, if specified
	if strings.Contains(adminUser, "@") {
		adminUser = strings.Split(adminUser, "@")[0]
	}

	for _, dbName := range ownedDatabases {

		databaseExists, err := m.DatabaseExists(ctx, db, dbName)
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
			continue
		}
		if !databaseExists {
			errorStrings = append(errorStrings, fmt.Sprintf("Database %s does not exist.", dbName))
			continue
		}

		dbOwnerRole := m.GetDbOwnerRole(dbName)

		err = m.EnsureDatabaseHasOwner(ctx, dbName, dbOwnerRole, adminUser, db)
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
			continue
		}

		// add dbOwnerRole to users roles
		err = m.GrantUserRoles(ctx, user, []string{dbOwnerRole}, db)
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
			continue
		}

		// add execution context to users role for db
		if err = helpers.FindBadChars(dbOwnerRole); err != nil {
			errorStrings = append(errorStrings, fmt.Sprintf("Problem found with dbOwnerRole: %s", err.Error()))
			continue
		}

		tsql := fmt.Sprintf("ALTER ROLE %q IN DATABASE %q SET role TO '%s' ", user, dbName, dbOwnerRole)
		_, err = db.ExecContext(ctx, tsql)
		if err != nil {
			errorStrings = append(errorStrings, err.Error())
			continue
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// EnsureDatabaseHasOwnOwnerROle
func (m *PostgreSqlUserManager) EnsureDatabaseHasOwner(ctx context.Context, dbName string, dbOwner string, adminUser string, db *sql.DB) error {
	var errorStrings []string

	if err := helpers.FindBadChars(dbOwner); err != nil {
		return fmt.Errorf("Problem found with dbOwner: %w", err)
	}

	if err := helpers.FindBadChars(dbName); err != nil {
		return fmt.Errorf("Problem found with dbName: %w", err)
	}

	// check if database has the owner role dbOwner
	tsql := fmt.Sprintf(`SELECT pg_catalog.pg_get_userbyid(d.datdba) as "OWNER" 
						FROM pg_catalog.pg_database d 
						WHERE d.datname = '%s' 
						AND pg_catalog.pg_get_userbyid(d.datdba) = '%s';`,
		dbName, dbOwner)

	res, err := db.ExecContext(ctx, tsql)
	if err != nil {
		return err
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		// the role dbOwner is not owner of the database dbName

		// create owner role if not exists
		roleExists, err := m.RoleExists(ctx, db, dbOwner)
		if err != nil {
			return err
		}
		if !roleExists {
			tsql = fmt.Sprintf("CREATE ROLE %q", dbOwner)
			_, err = db.ExecContext(ctx, tsql)
			if err != nil {
				return err
			}
		}

		// add to admin user, otherwise the admin user can not
		// assign the role as owner to the database
		tsql = fmt.Sprintf("GRANT %q TO %q", dbOwner, adminUser)
		_, err = db.ExecContext(ctx, tsql)
		if err != nil {
			return err
		}

		// update the owner of the database to the new role
		tsql := fmt.Sprintf("ALTER DATABASE %q OWNER TO %q", dbName, dbOwner)
		_, err = db.ExecContext(ctx, tsql)
		if err != nil {
			return err
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// GrantUserRoles grants roles to a user for a given database
func (m *PostgreSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string

	if err := helpers.FindBadChars(user); err != nil {
		return errors.Wrap(err, "problem found with username")
	}

	for _, role := range roles {
		tsql := fmt.Sprintf("GRANT %q TO %q", role, user)
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

// DatabaseExists checks if a database exists
func (m *PostgreSqlUserManager) DatabaseExists(ctx context.Context, db *sql.DB, dbName string) (bool, error) {

	res, err := db.ExecContext(ctx, "SELECT datname FROM pg_catalog.pg_database WHERE datname = $1", dbName)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err

}

// RoleExists checks if db contains role
func (m *PostgreSqlUserManager) RoleExists(ctx context.Context, db *sql.DB, rolname string) (bool, error) {

	res, err := db.ExecContext(ctx, "SELECT * FROM pg_roles WHERE rolname = $1", rolname)
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
func (m *PostgreSqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha2.PostgreSQLUser, secretClient secrets.SecretClient) (bool, error) {

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
func (m *PostgreSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha2.PostgreSQLUser, secretClient secrets.SecretClient) map[string][]byte {
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
			PSecretUsernameKey:               []byte(""),
			PSecretPasswordKey:               []byte(pw),
			"PSqlServerNamespace":            []byte(instance.Namespace),
			"PSqlServerName":                 []byte(instance.Spec.Server),
			"fullyQualifiedServerName":       []byte(instance.Spec.Server + "." + psqldbdnssuffix),
			PSecretFullyQualifiedUsernameKey: []byte(""),
		}
	}

	return secret
}
