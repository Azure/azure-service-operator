// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqluser

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	azuresql "github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime"

	_ "github.com/denisenkom/go-mssqldb"
	"k8s.io/apimachinery/pkg/types"
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
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlUserManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlUserManager {
	return &AzureSqlUserManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// GetDB retrieves a database
func (s *AzureSqlUserManager) GetDB(ctx context.Context, resourceGroupName string, serverName string, databaseName string) (azuresql.Database, error) {
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

// ConnectToSqlDb connects to the SQL db using the given credentials
func (s *AzureSqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, server string, database string, port int, user string, password string) (*sql.DB, error) {

	fullServerAddress := fmt.Sprintf("%s."+config.Environment().SQLDatabaseDNSSuffix, server)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;Persist Security Info=False;Pooling=False;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30", fullServerAddress, user, password, port, database)

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
func (s *AzureSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
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
func (s *AzureSqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[SecretUsernameKey])
	newPassword := string(secret[SecretPasswordKey])

	// make an effort to prevent sql injection
	if err := findBadChars(newUser); err != nil {
		return "", fmt.Errorf("Problem found with username: %v", err)
	}
	if err := findBadChars(newPassword); err != nil {
		return "", fmt.Errorf("Problem found with password: %v", err)
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

// RotateUserPassword - Generate and save a new password for the user
func (s *AzureSqlUserManager) RotateUserPassword(ctx context.Context, secret map[string][]byte, db *sql.DB) error {
	user := string(secret[SecretUsernameKey])
	newPassword := helpers.NewPassword()

	// make an effort to prevent sql injection
	if err := findBadChars(user); err != nil {
		return fmt.Errorf("Problem found with username: %v", err)
	}
	if err := findBadChars(newPassword); err != nil {
		return fmt.Errorf("Problem found with password: %v", err)
	}

	tsql := fmt.Sprintf("ALTER USER \"%s\" WITH PASSWORD='%s'", user, newPassword)
	_, err := db.ExecContext(ctx, tsql)

	return err
}

// UserExists checks if db contains user
func (s *AzureSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {
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
func (s *AzureSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	tsql := "DROP USER @user"
	_, err := db.ExecContext(ctx, tsql, sql.Named("user", user))
	return err
}

// ReconcileCustomSecrets - Create or update any custom secret formats that were declared for the user based on the standard userSecret
func (s *AzureSqlUserManager) ReconcileCustomSecrets(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient, userSecret map[string][]byte) error {

	// Custom secret formats are only available through keyvault as they require separated secrets
	keyVaultEnabled := reflect.TypeOf(secretClient).Elem().Name() == "KeyvaultSecretClient"
	if !keyVaultEnabled {
		return nil
	}

	userSecretNamespacedName := s.GetSecretNamespacedName(instance, secretClient)

	// Instantiate a map of all possible formats and flip the bool to true for any that have been requested in the spec.
	// Formats that were not requested will be explicitly deleted.
	requestedFormats := map[string]bool{
		"adonet":         false,
		"adonet-urlonly": false,
		"jdbc":           false,
		"jdbc-urlonly":   false,
		"odbc":           false,
		"odbc-urlonly":   false,
		"server":         false,
		"database":       false,
		"username":       false,
		"password":       false,
	}
	for _, format := range instance.Spec.KeyVaultSecretFormats {
		requestedFormats[format] = true
	}

	// Deleted items will be processed immediately but secrets that need to be added will be created in this array and persisted in one pass at the end
	formattedSecrets := make(map[string][]byte)

	for formatName, requested := range requestedFormats {
		// Add the format to the output map if it has been requested otherwise call for its deletion from the secret store
		if requested {
			switch formatName {
			case "adonet":
				formattedSecrets["adonet"] = []byte(fmt.Sprintf(
					"Server=tcp:%v,1433;Initial Catalog=%v;Persist Security Info=False;User ID=%v;Password=%v;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;",
					string(userSecret["fullyQualifiedServerName"]),
					instance.Spec.DbName,
					string(userSecret[SecretUsernameKey]),
					string(userSecret[SecretPasswordKey]),
				))

			case "adonet-urlonly":
				formattedSecrets["adonet-urlonly"] = []byte(fmt.Sprintf(
					"Server=tcp:%v,1433;Initial Catalog=%v;Persist Security Info=False; MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout",
					string(userSecret["fullyQualifiedServerName"]),
					instance.Spec.DbName,
				))

			case "jdbc":
				formattedSecrets["jdbc"] = []byte(fmt.Sprintf(
					"jdbc:sqlserver://%v:1433;database=%v;user=%v@%v;password=%v;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*."+config.Environment().SQLDatabaseDNSSuffix+";loginTimeout=30;",
					string(userSecret["fullyQualifiedServerName"]),
					instance.Spec.DbName,
					string(userSecret[SecretUsernameKey]),
					instance.Spec.Server,
					string(userSecret[SecretPasswordKey]),
				))
			case "jdbc-urlonly":
				formattedSecrets["jdbc-urlonly"] = []byte(fmt.Sprintf(
					"jdbc:sqlserver://%v:1433;database=%v;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*."+config.Environment().SQLDatabaseDNSSuffix+";loginTimeout=30;",
					string(userSecret["fullyQualifiedServerName"]),
					instance.Spec.DbName,
				))

			case "odbc":
				formattedSecrets["odbc"] = []byte(fmt.Sprintf(
					"Server=tcp:%v,1433;Initial Catalog=%v;Persist Security Info=False;User ID=%v;Password=%v;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;",
					string(userSecret["fullyQualifiedServerName"]),
					instance.Spec.DbName,
					string(userSecret[SecretUsernameKey]),
					string(userSecret[SecretPasswordKey]),
				))
			case "odbc-urlonly":
				formattedSecrets["odbc-urlonly"] = []byte(fmt.Sprintf(
					"Driver={ODBC Driver 13 for SQL Server};Server=tcp:%v,1433;Database=%v; Encrypt=yes;TrustServerCertificate=no;Connection Timeout=30;",
					string(userSecret["fullyQualifiedServerName"]),
					instance.Spec.DbName,
				))
			case "server":
				formattedSecrets["server"] = userSecret["fullyQualifiedServerName"]

			case "database":
				formattedSecrets["database"] = []byte(instance.Spec.DbName)

			case "username":
				formattedSecrets["username"] = userSecret[SecretUsernameKey]

			case "password":
				formattedSecrets["password"] = userSecret[SecretPasswordKey]
			}
		} else {
			// delete any formats that were not explicitly requested
			_ = secretClient.Delete(
				ctx,
				types.NamespacedName{Namespace: userSecretNamespacedName.Namespace, Name: instance.Name + "-" + formatName},
			)
		}
	}

	err := secretClient.Upsert(
		ctx,
		types.NamespacedName{Namespace: userSecretNamespacedName.Namespace, Name: instance.Name},
		formattedSecrets,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
		secrets.Flatten(true),
	)
	if err != nil {
		return err
	}

	return nil
}

// DeleteSecrets deletes the secrets associated with a SQLUser
func (s *AzureSqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) (bool, error) {
	userSecretNamespacedName := s.GetSecretNamespacedName(instance, secretClient)

	// delete standard user secret
	err := secretClient.Delete(
		ctx,
		userSecretNamespacedName,
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
			formatKey := types.NamespacedName{Namespace: userSecretNamespacedName.Namespace, Name: instance.Name + "-" + formatName}

			err = secretClient.Delete(
				ctx,
				formatKey,
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
func (s *AzureSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) map[string][]byte {
	key := s.GetSecretNamespacedName(instance, secretClient)

	secret, err := secretClient.Get(ctx, key)
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

// GetSecretNamespacedName - Generate a more unique namespacednames for keys being stored in KeyVault
// Kube ties the default namespacedname to the unique resource object to prevent collisions but KeyVault requires more specificity in the name
func (s *AzureSqlUserManager) GetSecretNamespacedName(instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) types.NamespacedName {
	var namespacedName types.NamespacedName
	keyVaultEnabled := reflect.TypeOf(secretClient).Elem().Name() == "KeyvaultSecretClient"

	if keyVaultEnabled {
		// Generate a custom namespacedname based on the sql server, database, and username
		var dbUserCustomNamespace string
		if instance.Spec.KeyVaultSecretPrefix != "" {
			dbUserCustomNamespace = instance.Spec.KeyVaultSecretPrefix
		} else {
			dbUserCustomNamespace = "azuresqluser-" + instance.Spec.Server + "-" + instance.Spec.DbName
		}

		namespacedName = types.NamespacedName{Namespace: dbUserCustomNamespace, Name: instance.Name}
	} else {
		// If we're not using keyvault, Kube secrets can uniquely identify user objects just from the name and namespace.
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
