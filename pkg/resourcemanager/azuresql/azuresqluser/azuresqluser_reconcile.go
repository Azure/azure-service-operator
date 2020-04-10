// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqluser

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
	"github.com/google/uuid"
	"k8s.io/apimachinery/pkg/runtime"

	_ "github.com/denisenkom/go-mssqldb"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure that user exists
func (s *AzureSqlUserManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return false, err
	}

	requestedUsername := instance.Spec.Username
	if len(requestedUsername) == 0 {
		requestedUsername = instance.Name
	}

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
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

	// get admin creds for server
	adminSecret, err := adminSecretClient.Get(ctx, key)
	if err != nil {
		instance.Status.Provisioning = false
		instance.Status.Message = fmt.Sprintf("admin secret : %s, not found in %s", key.String(), reflect.TypeOf(adminSecretClient).Elem().Name())
		return false, nil
	}

	adminUser := string(adminSecret[SecretUsernameKey])
	adminPassword := string(adminSecret[SecretPasswordKey])

	_, err = s.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = err.Error()
		azerr := errhelp.NewAzureErrorAzureError(err)

		catch := []string{
			errhelp.ResourceNotFound,
		}
		if helpers.ContainsString(catch, azerr.Type) {
			instance.Status.Message = fmt.Sprintf("Waiting for Azure SQL server %s to provision", instance.Spec.Server)
			instance.Status.Provisioning = false
			return false, nil
		}

		ignore := []string{
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		if helpers.ContainsString(ignore, azerr.Type) {
			return false, nil
		}

		return false, err
	}

	db, err := s.ConnectToSqlDb(ctx, DriverName, instance.Spec.Server, instance.Spec.DbName, SqlServerPort, adminUser, adminPassword)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)

		// catch firewall issue - keep cycling until it clears up
		if strings.Contains(err.Error(), "create a firewall rule for this IP address") {
			instance.Status.Provisioned = false
			instance.Status.Provisioning = false
			return false, nil
		}

		return false, err
	}

	// determine our key namespace - if we're persisting to kube, we should use the actual instance namespace.
	// In keyvault we have to avoid collisions with other secrets so we create a custom namespace with the user's parameters
	key = GetNamespacedName(instance, sqlUserSecretClient)

	// create or get new user secret
	DBSecret := s.GetOrPrepareSecret(ctx, instance, sqlUserSecretClient)
	// reset user from secret in case it was loaded
	user := string(DBSecret[SecretUsernameKey])
	if user == "" {
		user = fmt.Sprintf("%s-%s", requestedUsername, uuid.New())
		DBSecret[SecretUsernameKey] = []byte(user)
	}

	// Publishing the user secret:
	// We do this first so if the keyvault does not have right permissions we will not proceed to creating the user
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
	keyVaultEnabled := reflect.TypeOf(sqlUserSecretClient).Elem().Name() == "KeyvaultSecretClient"
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
					types.NamespacedName{Namespace: key.Namespace, Name: instance.Name + "-" + formatName},
				)
			}
		}

		err = sqlUserSecretClient.Upsert(
			ctx,
			types.NamespacedName{Namespace: key.Namespace, Name: instance.Name},
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

// Delete deletes a user
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
	_, err = s.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
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
		instance.Status.Message = errhelp.StripErrorIDs(err)
		if strings.Contains(err.Error(), "create a firewall rule for this IP address") {

			// there is nothing much we can do here - cycle forever
			return true, err
		}
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

// GetParents gets the parents of the user
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

// GetStatus gets the status
func (s *AzureSqlUserManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := s.convert(obj)
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
