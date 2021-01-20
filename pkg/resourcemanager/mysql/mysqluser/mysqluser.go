// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package mysqluser

import (
	"context"
	"database/sql"
	"fmt"

	mysqlmgmt "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	_ "github.com/go-sql-driver/mysql" //mysql drive link
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql"
	mysqldatabase "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/database"
	mysqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/server"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

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
func (m *MySqlUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (db mysqlmgmt.Database, err error) {
	dbClient := mysqldatabase.GetMySQLDatabasesClient(m.Creds)
	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// GetServer retrieves a server
func (m *MySqlUserManager) GetServer(ctx context.Context, resourceGroupName, serverName string) (mysqlmgmt.Server, error) {
	client := mysqlserver.NewMySQLServerClient(m.Creds, m.SecretClient, m.Scheme)
	return client.GetServer(ctx, resourceGroupName, serverName)
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

// DeleteSecrets deletes the secrets associated with a SQLUser
func (m *MySqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha2.MySQLUser, secretClient secrets.SecretClient) (bool, error) {
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
func (m *MySqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha2.MySQLUser, secretClient secrets.SecretClient) map[string][]byte {
	key := GetNamespacedName(instance, secretClient)

	secret, err := secretClient.Get(ctx, key)
	if err != nil {
		pw := helpers.NewPassword()
		return map[string][]byte{
			"username":                 []byte(""),
			"password":                 []byte(pw),
			"MySqlServerNamespace":     []byte(instance.Namespace),
			"MySqlServerName":          []byte(instance.Spec.Server),
			"fullyQualifiedServerName": []byte(mysql.GetFullSQLServerName(instance.Spec.Server)),
		}
	}
	return secret
}

// GetNamespacedName gets the namespaced-name
func GetNamespacedName(instance *v1alpha2.MySQLUser, secretClient secrets.SecretClient) types.NamespacedName {

	return types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
}
