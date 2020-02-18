// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package azuresqluser

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqldb"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	k8sSecrets "github.com/Azure/azure-service-operator/pkg/secrets/kube"
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

	if options.SecretClient != nil {
		s.SecretClient = options.SecretClient
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	var adminSecretClient secrets.SecretClient
	// if the admin credentials haven't been set, default admin credentials to servername
	if len(instance.Spec.AdminSecret) == 0 {
		instance.Spec.AdminSecret = instance.Spec.Server
	}

	// if the admin secret keyvault is not specified, assume it is a kube secret
	if len(instance.Spec.AdminSecretKeyVault) == 0 {
		if options.KubeClient != nil {
			adminSecretClient = k8sSecrets.New(options.KubeClient)
		} else {
			return false, err
		}
	} else {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault)
	}

	// need this to detect missing databases
	dbClient := azuresqldb.NewAzureSqlDbManager(s.Log)

	// get admin creds for server
	key := types.NamespacedName{Name: instance.Spec.AdminSecret, Namespace: instance.Namespace}
	adminSecret, err := adminSecretClient.Get(ctx, key)
	if err != nil {
		instance.Status.Provisioning = false
		instance.Status.Message = fmt.Sprintf("admin secret : %s, not found", instance.Spec.AdminSecret)
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
	DBSecret := s.GetOrPrepareSecret(ctx, instance)
	// reset user from secret in case it was loaded
	user := string(DBSecret[SecretUsernameKey])
	if user == "" {
		user = fmt.Sprintf("%s-%s", instance.Name, uuid.New())
		DBSecret[SecretUsernameKey] = []byte(user)
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

		// apply roles to user
		if len(instance.Spec.Roles) == 0 {
			instance.Status.Message = "No roles specified for user"
			return false, fmt.Errorf("No roles specified for database user")
		}
		err = s.GrantUserRoles(ctx, user, instance.Spec.Roles, db)
		if err != nil {
			instance.Status.Message = "GrantUserRoles failed"
			return false, fmt.Errorf("GrantUserRoles failed")
		}

		// publish user secret
		key = types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
		err = s.SecretClient.Upsert(
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

	if options.SecretClient != nil {
		s.SecretClient = options.SecretClient
	}

	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	var adminSecretClient secrets.SecretClient
	// get admin credentials to connect to db
	key := types.NamespacedName{Name: instance.Spec.AdminSecret, Namespace: instance.Namespace}

	// if the admin secret keyvault is not specified, assume it is a kube secret
	if len(instance.Spec.AdminSecretKeyVault) == 0 {
		if options.KubeClient != nil {
			adminSecretClient = k8sSecrets.New(options.KubeClient)
		} else {
			return false, err
		}
	} else {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault)
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

	DBSecret := s.GetOrPrepareSecret(ctx, instance)
	if string(DBSecret[SecretUsernameKey]) == "" {
		// maybe this could be an error...
		return false, nil // fmt.Errorf("no username in SQLUser secret")
	}

	exists, err := s.UserExists(ctx, db, user)
	if err != nil {
		return true, err
	}
	if !exists {
		// Best case deletion of secrets
		key = types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
		s.SecretClient.Delete(ctx, key)
		return false, nil
	}

	err = s.DropUser(ctx, db, user)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete AzureSqlUser failed with %s", err.Error())
		return false, err
	}

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

// GetOrPrepareSecret gets or creates a secret
func (s *AzureSqlUserManager) GetOrPrepareSecret(ctx context.Context, instance *v1alpha1.AzureSQLUser) map[string][]byte {
	pw, _ := generateRandomPassword(16)
	key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}

	secret, err := s.SecretClient.Get(ctx, key)
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
