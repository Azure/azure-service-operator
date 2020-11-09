// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package mysqluser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	mysqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/database"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	_ "github.com/go-sql-driver/mysql" //mysql drive link
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

// MySqlUserManager for mysqluser manager
type MySqlUserManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

// NewMySqlUserManager creates a new MySqlUserManager
func NewMySqlUserManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *MySqlUserManager {
	return &MySqlUserManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (m *MySqlUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (db mysql.Database, err error) {
	dbClient := mysqldatabase.GetMySQLDatabasesClient(m.Creds)
	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// ConnectToSqlDb connects to the SQL db using the given credentials
func (m *MySqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, fullserver string, database string, port int, user string, password string) (*sql.DB, error) {

	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?tls=skip-verify&interpolateParams=true", user, password, fullserver, port, database)

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
func (m *MySqlUserManager) GrantUserRoles(ctx context.Context, user string, database string, roles []string, db *sql.DB) error {
	var errorStrings []string
	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("Problem found with username: %v", err)
	}

	for _, role := range roles {

		if err := helpers.FindBadChars(role); err != nil {
			return fmt.Errorf("Problem found with role: %v", err)
		}
		//TODO: how to use SQL parameters for grant command, like CreateUser and DropUser
		tsql := fmt.Sprintf("GRANT %s ON `%s`.* TO '%s'", role, database, user)
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
func (m *MySqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[MSecretUsernameKey])
	newPassword := string(secret[MSecretPasswordKey])

	if err := helpers.FindBadChars(newUser); err != nil {
		return "", fmt.Errorf("Problem found with username: %v", err)
	}
	if err := helpers.FindBadChars(newPassword); err != nil {
		return "", fmt.Errorf("Problem found with password: %v", err)
	}

	tsql := "CREATE USER IF NOT EXISTS ? IDENTIFIED BY ? "
	_, err := db.ExecContext(ctx, tsql, newUser, newPassword)

	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// UserExists checks if db contains user
func (m *MySqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {

	err := db.QueryRowContext(ctx, "SELECT * FROM mysql.user WHERE User = $1", username)
	//err := db.ExecContext(ctx, tsql)

	if err != nil {
		return false, nil
	}
	return true, nil

}

// DropUser drops a user from db
func (m *MySqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {

	if err := helpers.FindBadChars(user); err != nil {
		return fmt.Errorf("Problem found with username: %v", err)
	}
	_, err := db.ExecContext(ctx, "DROP USER IF EXISTS ?", user)
	return err
}

// DeleteSecrets deletes the secrets associated with a SQLUser
func (m *MySqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) (bool, error) {
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

	return false, nil
}

// GetOrPrepareSecret gets or creates a secret
func (m *MySqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) map[string][]byte {
	key := GetNamespacedName(instance, secretClient)

	mysqldbdnssuffix := "mysql.database.azure.com"
	if config.Environment().Name != "AzurePublicCloud" {
		mysqldbdnssuffix = "mysql." + string(config.Environment().SQLDatabaseDNSSuffix)
	}

	secret, err := secretClient.Get(ctx, key)
	if err != nil {
		pw := helpers.NewPassword()
		return map[string][]byte{
			"username":                 []byte(""),
			"password":                 []byte(pw),
			"MySqlServerNamespace":     []byte(instance.Namespace),
			"MySqlServerName":          []byte(instance.Spec.Server),
			"fullyQualifiedServerName": []byte(instance.Spec.Server + "." + mysqldbdnssuffix),
			"MySqlDatabaseName":        []byte(instance.Spec.DbName),
		}
	}
	return secret
}

// GetNamespacedName gets the namespaced-name
func GetNamespacedName(instance *v1alpha1.MySQLUser, secretClient secrets.SecretClient) types.NamespacedName {

	return types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
}
