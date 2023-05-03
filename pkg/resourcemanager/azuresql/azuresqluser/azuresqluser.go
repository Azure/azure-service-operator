// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqluser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	azuresql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/v3.0/sql"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

// SqlServerPort is the default server port for sql server
const SqlServerPort = 1433

// DriverName is driver name for db connection
const DriverName = "sqlserver"

// SecretUsernameKey is the username key in secret
const SecretUsernameKey = "username"

// SecretPasswordKey is the password key in secret
const SecretPasswordKey = "password"

type AzureSqlUserManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlUserManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlUserManager {
	return &AzureSqlUserManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (m *AzureSqlUserManager) GetDB(ctx context.Context, subscriptionID string, resourceGroupName string, serverName string, databaseName string) (azuresql.Database, error) {
	dbClient, err := azuresqlshared.GetGoDbClient(azuresqlshared.GetSubscriptionCredentials(m.Creds, subscriptionID))
	if err != nil {
		return azuresql.Database{}, err
	}

	return dbClient.Get(
		ctx,
		resourceGroupName,
		serverName,
		databaseName,
	)
}

// ConnectToSqlDb connects to the SQL db using the given credentials
func (m *AzureSqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, server string, database string, port int, user string, password string) (*sql.DB, error) {

	fullServerAddress := fmt.Sprintf("%s."+config.Environment().SQLDatabaseDNSSuffix, server)
	// Make sure to set connection timeout: https://github.com/denisenkom/go-mssqldb/issues/609
	connString := fmt.Sprintf(
		"server=%s;user id=%s;password=%s;port=%d;database=%s;Persist Security Info=False;Pooling=False;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30",
		fullServerAddress,
		user,
		password,
		port,
		database)

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
func (m *AzureSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string
	for _, role := range roles {
		tsql := "sp_addrolemember @role, @user"

		_, err := db.ExecContext(
			ctx, tsql,
			sql.Named("role", role),
			sql.Named("user", user),
		)
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
func (m *AzureSqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[SecretUsernameKey])
	newPassword := string(secret[SecretPasswordKey])

	// make an effort to prevent sql injection
	if err := helpers.FindBadChars(newUser); err != nil {
		return "", errors.Wrap(err, "Problem found with username")
	}
	if err := helpers.FindBadChars(newPassword); err != nil {
		return "", errors.Wrap(err, "Problem found with password")
	}

	tsql := fmt.Sprintf("CREATE USER \"%s\" WITH PASSWORD='%s'", newUser, newPassword)
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
func (m *AzureSqlUserManager) UpdateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) error {
	user := string(secret[SecretUsernameKey])
	newPassword := helpers.NewPassword()

	// make an effort to prevent sql injection
	if err := helpers.FindBadChars(user); err != nil {
		return errors.Wrap(err, "problem found with username")
	}
	if err := helpers.FindBadChars(newPassword); err != nil {
		return errors.Wrap(err, "problem found with password")
	}

	tsql := fmt.Sprintf("ALTER USER \"%s\" WITH PASSWORD='%s'", user, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	return err
}

// UserExists checks if db contains user
func (m *AzureSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {
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
func (m *AzureSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	tsql := fmt.Sprintf("DROP USER %q", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

// DeleteSecrets deletes the secrets associated with a SQLUser
func (m *AzureSqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) (bool, error) {
	secretKey := MakeSecretKey(secretClient, instance)

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
	if secretClient.IsKeyVault() {
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

		err = secretClient.Delete(
			ctx,
			secretKey,
			secrets.Flatten(true, customFormatNames...),
		)
		if err != nil {
			instance.Status.Message = "failed to delete secret, err: " + err.Error()
			return false, err
		}
	}

	return false, nil
}

// GetOrPrepareSecret gets or creates a secret
func (m *AzureSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) map[string][]byte {
	secretKey := MakeSecretKey(secretClient, instance)

	secret, err := secretClient.Get(ctx, secretKey)
	if err != nil {
		// @todo: find out whether this is an error due to non existing key or failed conn
		pw := helpers.NewPassword()
		return map[string][]byte{
			"username":                 []byte(""),
			"password":                 []byte(pw),
			"azureSqlServerNamespace":  []byte(instance.Namespace),
			"azureSqlServerName":       []byte(instance.Spec.Server),
			"fullyQualifiedServerName": []byte(instance.Spec.Server + "." + config.Environment().SQLDatabaseDNSSuffix),
			"azureSqlDatabaseName":     []byte(instance.Spec.DbName),
		}
	}
	return secret
}

func MakeSecretKey(secretClient secrets.SecretClient, instance *v1alpha1.AzureSQLUser) secrets.SecretKey {
	if secretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
		// TODO: It's awkward that we have to have IsKeyVault on our interface which should be abstracting
		// TODO: away KeyVault vs Kubernetes... but we need it for legacy reasons (ick).
		if !secretClient.IsKeyVault() {
			// The fact that this path doesn't take into account secret prefix is for legacy reasons... (although it feels like a bug)
			return secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
		}

		var namespace string
		if instance.Spec.KeyVaultSecretPrefix != "" {
			namespace = instance.Spec.KeyVaultSecretPrefix
		} else {
			namespace = "azuresqluser-" + instance.Spec.Server + "-" + instance.Spec.DbName
		}
		return secrets.SecretKey{Name: instance.Name, Namespace: namespace, Kind: instance.TypeMeta.Kind}
	}

	// TODO: Ignoring prefix here... we ok with that?
	return secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
}
