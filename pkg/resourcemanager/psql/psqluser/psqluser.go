// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package psqluser

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
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

// ConnectToSqlDb connects to the SQL db using the given credentials
func (s *PostgreSqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, server string, database string, port int, user string, password string) (*sql.DB, error) {

	psqldbdnssuffix := "database.azure.com"
	if config.Environment().Name != "AzurePublicCloud" {
		psqldbdnssuffix = config.Environment().SQLDatabaseDNSSuffix
	}
	//the host or fullserveraddress should be:
	//for public cloud <server>.mysql.database.azure.com
	//for China cloud <server>.mysql.database.chinacloudapi.cn
	//for German cloud <server>.mysql.database.cloudapi.de
	//for US government <server>.mysql.database.usgovcloudapi.net
	fullServerAddress := fmt.Sprintf("%s.mysql."+psqldbdnssuffix, server)

	connString := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=require connect_timeout=30", fullServerAddress, user, password, port, database)

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
	for _, role := range roles {
		tsql := fmt.Sprintf("GRANT %s TO %q", role, user)

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
	if err := findBadChars(newUser); err != nil {
		return "", fmt.Errorf("Problem found with username: %v", err)
	}
	if err := findBadChars(newPassword); err != nil {
		return "", fmt.Errorf("Problem found with password: %v", err)
	}

	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD '%s'", newUser, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	// TODO: Have db lib do string interpolation
	//tsql := fmt.Sprintf(`CREATE USER @User WITH PASSWORD='@Password'`)
	//_, err := db.ExecContext(ctx, tsql, sql.Named("User", newUser), sql.Named("Password", newPassword))

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
	if err := findBadChars(user); err != nil {
		return fmt.Errorf("Problem found with username: %v", err)
	}
	if err := findBadChars(newPassword); err != nil {
		return fmt.Errorf("Problem found with password: %v", err)
	}

	tsql := fmt.Sprintf("ALTER USER '%s' WITH PASSWORD '%s'", user, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	return err
}

// UserExists checks if db contains user
func (s *PostgreSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {

	tsql := fmt.Sprintf("SELECT * FROM pg_user WHERE usename = '%s'", username)
	res, err := db.ExecContext(ctx, tsql)
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err

}

// DropUser drops a user from db
func (s *PostgreSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	tsql := fmt.Sprintf("DROP USER %s", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// DeleteSecrets deletes the secrets associated with a SQLUser
func (s *PostgreSqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) (bool, error) {
	// determine our key namespace - if we're persisting to kube, we should use the actual instance namespace.
	// In keyvault we have some creative freedom to allow more flexibility
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

	// delete all the custom formatted secrets if keyvault is in use
	keyVaultEnabled := reflect.TypeOf(secretClient).Elem().Name() == "KeyvaultSecretClient"
	if keyVaultEnabled {
		customFormatNames := []string{
			"adonet",
			"adonet-urlonly",
			"jdbc",
			"jdbc-urlonly",
			"odbc",
			"odbc-urlonly",
			"server",
			"database",
			"username",
			"password",
		}

		for _, formatName := range customFormatNames {
			key := types.NamespacedName{Namespace: secretKey.Namespace, Name: instance.Name + "-" + formatName}

			err = secretClient.Delete(
				ctx,
				key,
			)
			if err != nil {
				instance.Status.Message = "failed to delete secret, err: " + err.Error()
				return false, err
			}
		}
	}

	return false, nil
}

// GetOrPrepareSecret gets or creates a secret
func (s *PostgreSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) map[string][]byte {
	key := GetNamespacedName(instance, secretClient)

	secret, err := secretClient.Get(ctx, key)
	if err != nil {
		// @todo: find out whether this is an error due to non existing key or failed conn
		pw := helpers.NewPassword()
		return map[string][]byte{
			"username":                 []byte(""),
			"password":                 []byte(pw),
			"PSqlServerNamespace":      []byte(instance.Namespace),
			"PSqlServerName":           []byte(instance.Spec.Server),
			"fullyQualifiedServerName": []byte(instance.Spec.Server + "." + config.Environment().SQLDatabaseDNSSuffix),
			"PSqlDatabaseName":         []byte(instance.Spec.DbName),
		}
	}
	return secret
}

// GetNamespacedName gets the namespaced-name
func GetNamespacedName(instance *v1alpha1.PostgreSQLUser, secretClient secrets.SecretClient) types.NamespacedName {
	var namespacedName types.NamespacedName
	keyVaultEnabled := reflect.TypeOf(secretClient).Elem().Name() == "KeyvaultSecretClient"

	if keyVaultEnabled {
		// For a keyvault secret store, check for supplied namespace parameters
		var dbUserCustomNamespace string
		if instance.Spec.KeyVaultSecretPrefix != "" {
			dbUserCustomNamespace = instance.Spec.KeyVaultSecretPrefix
		} else {
			dbUserCustomNamespace = "psqluser-" + instance.Spec.Server + "-" + instance.Spec.DbName
		}

		namespacedName = types.NamespacedName{Namespace: dbUserCustomNamespace, Name: instance.Name}
	} else {
		namespacedName = types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
	}

	return namespacedName
}

func findBadChars(stack string) error {
	badChars := []string{
		"'",
		"\"",
		";",
		"--",
		"/*",
	}

	for _, s := range badChars {
		if idx := strings.Index(stack, s); idx > -1 {
			return fmt.Errorf("potentially dangerous character seqience found: '%s' at pos: %d", s, idx)
		}
	}
	return nil
}
