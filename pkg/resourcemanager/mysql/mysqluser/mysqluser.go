// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package mysqluser

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	mysqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/database"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	_ "github.com/go-sql-driver/mysql" //sql drive link
	"k8s.io/apimachinery/pkg/types"
)

// MSqlServerPort is the default server port for sql server
const MSqlServerPort = 3306

// MDriverName is driver name for psqldb connection
const MDriverName = "mysql"

// MSecretUsernameKey is the username key in secret
const MSecretUsernameKey = "username"

// MSecretPasswordKey is the password key in secret
const MSecretPasswordKey = "password"

//MySqlUserManager for mysqluser manager
type MySqlUserManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

//NewMySqlUserManager creates a new NewMySqlUserManager
func NewMySqlUserManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *MySqlUserManager {
	return &MySqlUserManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (s *MySqlUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (db mysql.Database, err error) {
	dbClient := mysqldatabase.GetMySQLDatabasesClient()
	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// ConnectToSqlDb connects to the SQL db using the given credentials
func (s *MySqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, server string, database string, port int, user string, password string) (*sql.DB, error) {

	mysqldbdnssuffix := "database.azure.com"
	if config.Environment().Name != "AzurePublicCloud" {
		mysqldbdnssuffix = config.Environment().SQLDatabaseDNSSuffix
	}
	//the host or fullserveraddress should be:
	//for public cloud <server>.mysql.database.azure.com
	//for China cloud <server>.mysql.database.chinacloudapi.cn
	//for German cloud <server>.mysql.database.cloudapi.de
	//for US government <server>.mysql.database.usgovcloudapi.net
	fullServerAddress := fmt.Sprintf("%s.mysql."+mysqldbdnssuffix, server)

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=skip-verify", user, password, fullServerAddress, port, database)

	db, err := sql.Open(drivername, connString)
	if err != nil {
		return db, err
	}

	err = db.Ping()
	if err != nil {
		return db, fmt.Errorf("error ping the mysql db:  %v", err)
	}

	return db, err
}

// GrantUserRoles grants roles to a user for a given database
func (s *MySqlUserManager) GrantUserRoles(ctx context.Context, user string, server string, roles []string, db *sql.DB) error {
	var errorStrings []string
	for _, role := range roles {
		tsql := fmt.Sprintf("GRANT %s TO '%s'@'%s'", role, user, server)

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
func (s *MySqlUserManager) CreateUser(ctx context.Context, server string, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[MSecretUsernameKey])
	newPassword := string(secret[MSecretPasswordKey])

	// make an effort to prevent sql injection
	if err := findBadChars(newUser); err != nil {
		return "", fmt.Errorf("Problem found with username: %v", err)
	}
	if err := findBadChars(newPassword); err != nil {
		return "", fmt.Errorf("Problem found with password: %v", err)
	}

	tsql := fmt.Sprintf("CREATE USER IF NOT EXISTS '%s'@'%s' IDENTIFIED BY '%s'", newUser, server, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// UserExists checks if db contains user
func (s *MySqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string, server string) (bool, error) {

	tsql := fmt.Sprintf("SELECT * FROM mysql.user WHERE (Host = '%s') and (User = '%s')", server, username)
	fmt.Printf(tsql)
	err := db.QueryRowContext(ctx, tsql)
	//err := db.ExecContext(ctx, tsql)

	if err != nil {
		return false, nil
	}
	return true, nil

	//rows, err := res.RowsAffected()
	//return rows > 0, err

}

// DropUser drops a user from db
func (s *MySqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string, server string) error {
	tsql := fmt.Sprintf("DROP USER IF EXISTS '%s'@'%s'", user, server)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// DeleteSecrets deletes the secrets associated with a SQLUser
func (s *MySqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) (bool, error) {
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
func (s *MySqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) map[string][]byte {
	key := GetNamespacedName(instance, secretClient)

	secret, err := secretClient.Get(ctx, key)
	if err != nil {
		// @todo: find out whether this is an error due to non existing key or failed conn
		pw := helpers.NewPassword()
		return map[string][]byte{
			"username":                 []byte(""),
			"password":                 []byte(pw),
			"MySqlServerNamespace":     []byte(instance.Namespace),
			"MySqlServerName":          []byte(instance.Spec.Server),
			"fullyQualifiedServerName": []byte(instance.Spec.Server + "." + config.Environment().SQLDatabaseDNSSuffix),
			"MySqlDatabaseName":        []byte(instance.Spec.DbName),
		}
	}
	return secret
}

// GetNamespacedName gets the namespaced-name
func GetNamespacedName(instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) types.NamespacedName {
	var namespacedName types.NamespacedName
	keyVaultEnabled := reflect.TypeOf(secretClient).Elem().Name() == "KeyvaultSecretClient"

	if keyVaultEnabled {
		// For a keyvault secret store, check for supplied namespace parameters
		var dbUserCustomNamespace string
		if instance.Spec.KeyVaultSecretPrefix != "" {
			dbUserCustomNamespace = instance.Spec.KeyVaultSecretPrefix
		} else {
			dbUserCustomNamespace = "Mysqluser-" + instance.Spec.Server + "-" + instance.Spec.DbName
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
