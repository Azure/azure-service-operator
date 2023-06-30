// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqluser

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1beta1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
)

func GetAdminSecretKey(adminSecretName string, namespace string) secrets.SecretKey {
	return secrets.SecretKey{Name: adminSecretName, Namespace: namespace, Kind: reflect.TypeOf(v1beta1.AzureSqlServer{}).Name()}
}

func (s *AzureSqlUserManager) getAdminSecret(ctx context.Context, instance *v1alpha1.AzureSQLUser) (map[string][]byte, error) {
	adminSecretClient := s.SecretClient
	adminSecretName := instance.Spec.AdminSecret
	if len(adminSecretName) == 0 {
		adminSecretName = instance.Spec.Server
	}
	adminSecretKey := GetAdminSecretKey(adminSecretName, instance.Namespace)

	// if the admin secret keyvault is not specified, fall back to global secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(
			instance.Spec.AdminSecretKeyVault,
			s.Creds,
			s.SecretClient.GetSecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())

		// This is here for legacy reasons
		if len(instance.Spec.AdminSecret) != 0 && s.SecretClient.GetSecretNamingVersion() == secrets.SecretNamingV1 {
			adminSecretKey.Namespace = ""
		}
	}

	// get admin creds for server
	secret, err := adminSecretClient.Get(ctx, adminSecretKey)
	if err != nil {
		return nil, errors.Wrap(err, "AzureSqlServer admin secret not found")
	}

	return secret, nil
}

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

	adminSecret, err := s.getAdminSecret(ctx, instance)
	if err != nil {
		instance.Status.Message = err.Error()
		return false, nil
	}

	adminUser := string(adminSecret[SecretUsernameKey])
	adminPassword := string(adminSecret[SecretPasswordKey])

	var userSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		userSecretClient = options.SecretClient
	} else {
		userSecretClient = s.SecretClient
	}

	_, err = s.GetDB(ctx, instance.Spec.SubscriptionID, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		requeueErrors := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(requeueErrors, azerr.Type) {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		// if this is an unmarshall error - ignore and continue, otherwise report error and requeue
		if !strings.Contains(errorString, "cannot unmarshal array into Go struct field serviceError2.details") {
			return false, err
		}
	}

	db, err := s.ConnectToSqlDb(ctx, DriverName, instance.Spec.Server, instance.Spec.DbName, SqlServerPort, adminUser, adminPassword)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		// catch firewall issue - keep cycling until it clears up
		if strings.Contains(err.Error(), "create a firewall rule for this IP address") {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		return false, err
	}
	defer db.Close()

	userSecretKey := MakeSecretKey(userSecretClient, instance)

	// create or get new user secret
	userSecret := s.GetOrPrepareSecret(ctx, instance, userSecretClient)
	// reset user from secret in case it was loaded
	user := string(userSecret[SecretUsernameKey])
	if user == "" {
		user = fmt.Sprintf("%s-%s", requestedUsername, uuid.New())
		userSecret[SecretUsernameKey] = []byte(user)
	}

	// Publishing the user secret:
	// We do this first so if the keyvault does not have right permissions we will not proceed to creating the user
	err = userSecretClient.Upsert(
		ctx,
		userSecretKey,
		userSecret,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
	)
	if err != nil {
		instance.Status.Message = "failed to update secret, err: " + err.Error()
		return false, err
	}

	// Preformatted special formats are only available through keyvault as they require separated secrets
	if userSecretClient.IsKeyVault() {
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
		var toDelete []string

		for formatName, requested := range requestedFormats {
			// Add the format to the output map if it has been requested otherwise call for its deletion from the secret store
			if requested {
				switch formatName {
				case "adonet":
					formattedSecrets["adonet"] = []byte(fmt.Sprintf(
						"Server=tcp:%s,1433;Initial Catalog=%s;Persist Security Info=False;User ID=%s;Password=%s;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;",
						string(userSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
						user,
						string(userSecret["password"]),
					))

				case "adonet-urlonly":
					formattedSecrets["adonet-urlonly"] = []byte(fmt.Sprintf(
						"Server=tcp:%s,1433;Initial Catalog=%s;Persist Security Info=False; MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout",
						string(userSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
					))

				case "jdbc":
					formattedSecrets["jdbc"] = []byte(fmt.Sprintf(
						"jdbc:sqlserver://%s:1433;database=%s;user=%s@%s;password=%s;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*."+config.Environment().SQLDatabaseDNSSuffix+";loginTimeout=30;",
						string(userSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
						user,
						instance.Spec.Server,
						string(userSecret["password"]),
					))
				case "jdbc-urlonly":
					formattedSecrets["jdbc-urlonly"] = []byte(fmt.Sprintf(
						"jdbc:sqlserver://%s:1433;database=%s;encrypt=true;trustServerCertificate=false;hostNameInCertificate=*."+config.Environment().SQLDatabaseDNSSuffix+";loginTimeout=30;",
						string(userSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
					))

				case "odbc":
					formattedSecrets["odbc"] = []byte(fmt.Sprintf(
						"Server=tcp:%s,1433;Initial Catalog=%s;Persist Security Info=False;User ID=%s;Password=%s;MultipleActiveResultSets=False;Encrypt=True;TrustServerCertificate=False;Connection Timeout=30;",
						string(userSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
						user,
						string(userSecret["password"]),
					))
				case "odbc-urlonly":
					formattedSecrets["odbc-urlonly"] = []byte(fmt.Sprintf(
						"Driver={ODBC Driver 13 for SQL Server};Server=tcp:%s,1433;Database=%s; Encrypt=yes;TrustServerCertificate=no;Connection Timeout=30;",
						string(userSecret["fullyQualifiedServerName"]),
						instance.Spec.DbName,
					))
				case "server":
					formattedSecrets["server"] = userSecret["fullyQualifiedServerName"]

				case "database":
					formattedSecrets["database"] = []byte(instance.Spec.DbName)

				case "username":
					formattedSecrets["username"] = []byte(user)

				case "password":
					formattedSecrets["password"] = userSecret["password"]
				}
			} else {
				toDelete = append(toDelete, formatName)
			}
		}

		// Purposefully ignoring error here for now
		userSecretClient.Delete(
			ctx,
			userSecretKey,
			secrets.Flatten(true, toDelete...))

		err = userSecretClient.Upsert(
			ctx,
			userSecretKey,
			formattedSecrets,
			secrets.WithOwner(instance),
			secrets.WithScheme(s.Scheme),
			secrets.Flatten(true),
		)
		if err != nil {
			return false, err
		}
	}

	userExists, err := s.UserExists(ctx, db, string(userSecret[SecretUsernameKey]))
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %s", err)
		return false, nil
	}

	if !userExists {
		user, err = s.CreateUser(ctx, userSecret, db)
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

	adminSecret, err := s.getAdminSecret(ctx, instance)
	if err != nil {
		// assuming if the admin secret is gone the sql server is too
		return false, nil
	}

	// short circuit connection if database doesn't exist
	_, err = s.GetDB(ctx, instance.Spec.SubscriptionID, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = err.Error()

		catch := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, err
	}

	var adminUser = string(adminSecret[SecretUsernameKey])
	var adminPassword = string(adminSecret[SecretPasswordKey])

	db, err := s.ConnectToSqlDb(ctx, DriverName, instance.Spec.Server, instance.Spec.DbName, SqlServerPort, adminUser, adminPassword)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		if strings.Contains(err.Error(), "create a firewall rule for this IP address") {

			// Stop the reconcile
			return false, nil
		}
		return false, err
	}
	defer db.Close()

	var sqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		sqlUserSecretClient = options.SecretClient
	} else {
		sqlUserSecretClient = s.SecretClient
	}

	userSecretKey := MakeSecretKey(sqlUserSecretClient, instance)
	userSecret, err := sqlUserSecretClient.Get(ctx, userSecretKey)
	if err != nil {
		//user secret is gone
		return false, nil
	}

	user := string(userSecret[SecretUsernameKey])

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

	instance.Status.Message = "Delete AzureSqlUser succeeded"

	//successfully delete
	return false, nil
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
