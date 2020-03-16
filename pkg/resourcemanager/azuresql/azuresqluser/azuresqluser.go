// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqluser

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	"github.com/go-logr/logr"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/runtime"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/sethvargo/go-password/password"
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
	Log          logr.Logger
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlUserManager(log logr.Logger, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlUserManager {
	return &AzureSqlUserManager{
		Log:          log,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// ConnectToSqlDb connects to the SQL db using the given credentials
func (m *AzureSqlUserManager) ConnectToSqlDb(ctx context.Context, drivername string, server string, database string, port int, user string, password string) (*sql.DB, error) {

	fullServerAddress := fmt.Sprintf("%s.database.windows.net", server)
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;Persist Security Info=False;Pooling=False;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30", fullServerAddress, user, password, port, database)

	db, err := sql.Open(drivername, connString)
	if err != nil {
		m.Log.Info("ConnectToSqlDb", "error from sql.Open is:", err.Error())
		return db, err
	}

	err = db.PingContext(ctx)
	if err != nil {
		m.Log.Info("ConnectToSqlDb", "error from db.Ping is:", err.Error())
		return db, err
	}

	return db, err
}

// Grants roles to a user for a given database
func (m *AzureSqlUserManager) GrantUserRoles(ctx context.Context, user string, roles []string, db *sql.DB) error {
	var errorStrings []string
	for _, role := range roles {
		tsql := fmt.Sprintf("sp_addrolemember \"%s\", \"%s\"", role, user)
		_, err := db.ExecContext(ctx, tsql)
		if err != nil {
			m.Log.Info("GrantUserRoles:", "Error executing add role:", err.Error())
			errorStrings = append(errorStrings, err.Error())
		}
	}

	if len(errorStrings) != 0 {
		return fmt.Errorf(strings.Join(errorStrings, "\n"))
	}
	return nil
}

// Creates user with secret credentials
func (m *AzureSqlUserManager) CreateUser(ctx context.Context, secret map[string][]byte, db *sql.DB) (string, error) {
	newUser := string(secret[SecretUsernameKey])
	newPassword := string(secret[SecretPasswordKey])
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

// UserExists checks if db contains user
func (m *AzureSqlUserManager) UserExists(ctx context.Context, db *sql.DB, username string) (bool, error) {
	res, err := db.ExecContext(ctx, fmt.Sprintf("SELECT * FROM sysusers WHERE NAME='%s'", username))
	if err != nil {
		return false, err
	}
	rows, err := res.RowsAffected()
	return rows > 0, err
}

// Drops user from db
func (m *AzureSqlUserManager) DropUser(ctx context.Context, db *sql.DB, user string) error {
	tsql := fmt.Sprintf("DROP USER \"%s\"", user)
	_, err := db.ExecContext(ctx, tsql)
	return err
}

func (s *AzureSqlUserManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	adminSecretClient := s.SecretClient

	adminsecretName := instance.Spec.AdminSecret
	if len(instance.Spec.AdminSecret) == 0 {
		adminsecretName = instance.Spec.Server
	}
	key := types.NamespacedName{Name: adminsecretName, Namespace: instance.Namespace}

	var sqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		sqlUserSecretClient = options.SecretClient
	} else {
		sqlUserSecretClient = s.SecretClient
	}

	// if the admin secret keyvault is not specified, fall back to global secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault)
		if len(instance.Spec.AdminSecret) != 0 {
			key = types.NamespacedName{Name: instance.Spec.AdminSecret}
		}
	}

	// need this to detect missing databases
	dbClient := azuresqldb.NewAzureSqlDbManager(s.Log)

	// get admin creds for server
	adminSecret, err := adminSecretClient.Get(ctx, key)
	if err != nil {
		instance.Status.Provisioning = false
		instance.Status.Message = fmt.Sprintf("admin secret : %s, not found in %s", key.String(), reflect.TypeOf(adminSecretClient).Elem().Name())
		return false, nil
	}

	adminUser := string(adminSecret[SecretUsernameKey])
	adminPassword := string(adminSecret[SecretPasswordKey])

	_, err = dbClient.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = err.Error()

		catch := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, err
	}

	db, err := s.ConnectToSqlDb(ctx, DriverName, instance.Spec.Server, instance.Spec.DbName, SqlServerPort, adminUser, adminPassword)
	if err != nil {
		instance.Status.Message = err.Error()
		if strings.Contains(err.Error(), "create a firewall rule for this IP address") {
			return false, nil
		}
		return false, err
	}

	// create or get new user secret
	DBSecret := s.GetOrPrepareSecret(ctx, instance, sqlUserSecretClient)
	// reset user from secret in case it was loaded
	user := string(DBSecret[SecretUsernameKey])
	if user == "" {
		user = fmt.Sprintf("%s-%s", instance.Name, uuid.New())
		DBSecret[SecretUsernameKey] = []byte(user)
	}

	// publish user secret
	// We do this first so if the keyvault does not have right permissions we will not proceed to creating the user

	// determine our key namespace - if we're persisting to kube, we should use the actual instance namespace.
	// In keyvault we have some creative freedom to allow more flexibility
	key = types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
	var dbUserCustomNamespace string

	keyVaultEnabled := reflect.TypeOf(sqlUserSecretClient).Elem().Name() == "KeyvaultSecretClient"

	if keyVaultEnabled {
		// For a keyvault secret store, check for supplied namespace parameters
		if instance.Spec.KeyVaultSecretPrefix != "" {
			dbUserCustomNamespace = instance.Spec.KeyVaultSecretPrefix
		} else {
			dbUserCustomNamespace = "azuresqluser-" + string(DBSecret["azureSqlServerName"]) + "-" + string(DBSecret["azureSqlDatabaseName"])
		}

		key = types.NamespacedName{Namespace: dbUserCustomNamespace, Name: instance.Name}
	}

	// publish standard user secret
	err = sqlUserSecretClient.Upsert(
		ctx,
		key,
		DBSecret,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
	)
	if err != nil {
		instance.Status.Message = "failed to update secret, err: " + err.Error()
		return false, err
	}

	// Preformatted special formats are only available through keyvault as they require separated secrets
	if keyVaultEnabled {
		// Instantiate a map of all formats and flip the bool to true for any that have been requested in the spec.
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
						string(DBSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
						user,
						string(DBSecret["password"]),
					))

				case "adonet-urlonly":
					formattedSecrets["adonet-urlonly"] = []byte(fmt.Sprintf(
						"Server=tcp:%v,1433;Initial Catalog=%v;Persist Security Info=False; MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout",
						string(DBSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
					))

				case "jdbc":
					formattedSecrets["jdbc"] = []byte(fmt.Sprintf(
						"jdbc:sqlserver://%v:1433;database=%v;user=%v@%v;password=%v;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.database.windows.net;loginTimeout=30;",
						string(DBSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
						user,
						instance.Spec.Server,
						string(DBSecret["password"]),
					))
				case "jdbc-urlonly":
					formattedSecrets["jdbc-urlonly"] = []byte(fmt.Sprintf(
						"jdbc:sqlserver://%v:1433;database=%v;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*.database.windows.net;loginTimeout=30;",
						string(DBSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
					))

				case "odbc":
					formattedSecrets["odbc"] = []byte(fmt.Sprintf(
						"Server=tcp:%v,1433;Initial Catalog=%v;Persist Security Info=False;User ID=%v;Password=%v;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;",
						string(DBSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
						user,
						string(DBSecret["password"]),
					))
				case "odbc-urlonly":
					formattedSecrets["odbc-urlonly"] = []byte(fmt.Sprintf(
						"Driver={ODBC Driver 13 for SQL Server};Server=tcp:%v,1433;Database=%v; Encrypt=yes;TrustServerCertificate=no;Connection Timeout=30;",
						string(DBSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
					))
				case "server":
					formattedSecrets["server"] = DBSecret["fullyQualifiedServerName"]

				case "database":
					formattedSecrets["database"] = []byte(instance.Spec.DbName)

				case "username":
					formattedSecrets["username"] = []byte(user)

				case "password":
					formattedSecrets["password"] = DBSecret["password"]
				}
			} else {
				err = sqlUserSecretClient.Delete(
					ctx,
					types.NamespacedName{Namespace: dbUserCustomNamespace, Name: instance.Name + "-" + formatName},
				)
			}
		}

		err = sqlUserSecretClient.Upsert(
			ctx,
			types.NamespacedName{Namespace: dbUserCustomNamespace, Name: instance.Name},
			formattedSecrets,
			secrets.WithOwner(instance),
			secrets.WithScheme(s.Scheme),
			secrets.Flatten(true),
		)
		if err != nil {
			return false, err
		}
	}

	userExists, err := s.UserExists(ctx, db, string(DBSecret[SecretUsernameKey]))
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %v", err)
		return false, nil
	}

	if !userExists {
		user, err = s.CreateUser(ctx, DBSecret, db)
		if err != nil {
			instance.Status.Message = "failed creating user, err: " + err.Error()
			return false, err
		}
	}

	// apply roles to user
	if len(instance.Spec.Roles) == 0 {
		instance.Status.Message = "No roles specified for user"
		return false, fmt.Errorf("No roles specified for database user")
	}

	err = s.GrantUserRoles(ctx, user, instance.Spec.Roles, db)
	if err != nil {
		fmt.Println(err)
		instance.Status.Message = "GrantUserRoles failed"
		return false, fmt.Errorf("GrantUserRoles failed")
	}

	instance.Status.Provisioned = true
	instance.Status.State = "Succeeded"
	instance.Status.Message = resourcemanager.SuccessMsg

	return true, nil
}

func (s *AzureSqlUserManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	adminSecretClient := s.SecretClient

	adminsecretName := instance.Spec.AdminSecret
	if len(instance.Spec.AdminSecret) == 0 {
		adminsecretName = instance.Spec.Server
	}
	key := types.NamespacedName{Name: adminsecretName, Namespace: instance.Namespace}

	var sqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		sqlUserSecretClient = options.SecretClient
	} else {
		sqlUserSecretClient = s.SecretClient
	}

	// if the admin secret keyvault is not specified, fall back to global secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault)
		if len(instance.Spec.AdminSecret) != 0 {
			key = types.NamespacedName{Name: instance.Spec.AdminSecret}
		}
	}

	adminSecret, err := adminSecretClient.Get(ctx, key)
	if err != nil {
		// assuming if the admin secret is gone the sql server is too
		return false, nil
	}

	// short circuit connection if database doesn't exist
	dbClient := azuresqldb.NewAzureSqlDbManager(s.Log)
	_, err = dbClient.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = err.Error()

		catch := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, err
	}

	var user = string(adminSecret[SecretUsernameKey])
	var password = string(adminSecret[SecretPasswordKey])

	db, err := s.ConnectToSqlDb(ctx, DriverName, instance.Spec.Server, instance.Spec.DbName, SqlServerPort, user, password)
	if err != nil {
		return false, err
	}

	exists, err := s.UserExists(ctx, db, user)
	if err != nil {
		return true, err
	}
	if !exists {
		s.DeleteSecrets(ctx, instance, sqlUserSecretClient)
		return false, nil
	}

	err = s.DropUser(ctx, db, user)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete AzureSqlUser failed with %s", err.Error())
		return false, err
	}

	// Once the user has been dropped, also delete their secrets.
	s.DeleteSecrets(ctx, instance, sqlUserSecretClient)

	instance.Status.Message = fmt.Sprintf("Delete AzureSqlUser succeeded")

	return true, nil
}

func (s *AzureSqlUserManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.DbName,
			},
			Target: &v1alpha1.AzureSqlDatabase{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &v1alpha1.AzureSqlServer{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &v1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *AzureSqlUserManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (s *AzureSqlUserManager) convert(obj runtime.Object) (*v1alpha1.AzureSQLUser, error) {
	local, ok := obj.(*v1alpha1.AzureSQLUser)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

// Deletes the secrets associated with a SQLUser
func (s *AzureSqlUserManager) DeleteSecrets(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) (bool, error) {
	// determine our key namespace - if we're persisting to kube, we should use the actual instance namespace.
	// In keyvault we have some creative freedom to allow more flexibility
	DBSecret := s.GetOrPrepareSecret(ctx, instance, s.SecretClient)
	secretKey := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}

	var dbUserCustomNamespace string

	keyVaultEnabled := reflect.TypeOf(s.SecretClient).Elem().Name() == "KeyvaultSecretClient"

	if keyVaultEnabled {
		// For a keyvault secret store, check for supplied namespace parameters
		if instance.Spec.KeyVaultSecretPrefix != "" {
			dbUserCustomNamespace = instance.Spec.KeyVaultSecretPrefix
		} else {
			dbUserCustomNamespace = "azuresqluser-" + string(DBSecret["azureSqlServerName"]) + "-" + string(DBSecret["azureSqlDatabaseName"])
		}

		secretKey = types.NamespacedName{Namespace: dbUserCustomNamespace, Name: instance.Name}
	}

	// delete standard user secret
	err := s.SecretClient.Delete(
		ctx,
		secretKey,
	)
	if err != nil {
		instance.Status.Message = "failed to delete secret, err: " + err.Error()
		return false, err
	}

	// delete all the custom formatted secrets if keyvault is in use
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
			err = s.SecretClient.Delete(
				ctx,
				types.NamespacedName{Namespace: dbUserCustomNamespace, Name: instance.Name + "-" + formatName},
			)
		}
	}

	return false, nil
}

// GetOrPrepareSecret gets or creates a secret
func (s *AzureSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.AzureSQLUser, secretClient secrets.SecretClient) map[string][]byte {
	pw, _ := generateRandomPassword(16)
	key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}

	secret, err := secretClient.Get(ctx, key)
	if err != nil {
		// @todo: find out whether this is an error due to non existing key or failed conn
		return map[string][]byte{
			"username":                 []byte(""),
			"password":                 []byte(pw),
			"azureSqlServerNamespace":  []byte(instance.Namespace),
			"azureSqlServerName":       []byte(instance.Spec.Server),
			"fullyQualifiedServerName": []byte(instance.Spec.Server + ".database.windows.net"),
			"azureSqlDatabaseName":     []byte(instance.Spec.DbName),
		}
	}
	return secret
}

// helper function to generate random password for sql server
func generateRandomPassword(n int) (string, error) {

	// Math - Generate a password where: 1/3 of the # of chars are digits, 1/3 of the # of chars are symbols,
	// and the remaining 1/3 is a mix of upper- and lower-case letters
	digits := n / 3
	symbols := n / 3

	// Generate a password that is n characters long, with # of digits and symbols described above,
	// allowing upper and lower case letters, and disallowing repeat characters.
	res, err := password.Generate(n, digits, symbols, false, false)
	if err != nil {
		return "", err
	}

	return res, nil
}
