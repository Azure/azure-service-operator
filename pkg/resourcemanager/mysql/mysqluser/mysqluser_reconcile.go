// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package mysqluser

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"

	_ "github.com/go-sql-driver/mysql" //sql drive link
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure that user exists
func (s *MySqlUserManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
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

	var mysqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		mysqlUserSecretClient = options.SecretClient
	} else {
		mysqlUserSecretClient = s.SecretClient
	}

	// if the admin secret keyvault is not specified, fall back to configured secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault, s.Creds)
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

	_, err = s.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		requeuErrors := []string{
			errhelp.ResourceNotFound,
			errhelp.ParentNotFoundErrorCode,
			errhelp.ResourceGroupNotFoundErrorCode,
		}
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(requeuErrors, azerr.Type) {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		// if this is an unmarshall error - ignore and continue, otherwise report error and requeue
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			return false, err
		}
	}

	adminUser := string(adminSecret["fullyQualifiedUsername"])
	adminPassword := string(adminSecret[MSecretPasswordKey])
	fullServerName := string(adminSecret["fullyQualifiedServerName"])

	db, err := s.ConnectToSqlDb(ctx, MDriverName, fullServerName, instance.Spec.DbName, MSqlServerPort, adminUser, adminPassword)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		// catch firewall issue - keep cycling until it clears up
		if strings.Contains(err.Error(), "is not allowed to connect to this MySQL server") {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		return false, err
	}

	// determine our key namespace - if we're persisting to kube, we should use the actual instance namespace.
	// In keyvault we have to avoid collisions with other secrets so we create a custom namespace with the user's parameters
	key = GetNamespacedName(instance, mysqlUserSecretClient)

	// create or get new user secret
	DBSecret := s.GetOrPrepareSecret(ctx, instance, mysqlUserSecretClient)
	// reset user from secret in case it was loaded
	user := string(DBSecret[MSecretUsernameKey])
	if user == "" {
		user = fmt.Sprintf(requestedUsername)
		DBSecret[MSecretUsernameKey] = []byte(user)
	}

	// Publishing the user secret:
	// We do this first so if the keyvault does not have right permissions we will not proceed to creating the user
	err = mysqlUserSecretClient.Upsert(
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

	err = s.GrantUserRoles(ctx, user, string(instance.Spec.DbName), instance.Spec.Roles, db)
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
func (s *MySqlUserManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

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

	var mysqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		mysqlUserSecretClient = options.SecretClient
	} else {
		mysqlUserSecretClient = s.SecretClient
	}

	// if the admin secret keyvault is not specified, fall back to configured secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault, s.Creds)
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
		azerr := errhelp.NewAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return false, nil
		}
		return false, err
	}

	adminuser := string(adminSecret["fullyQualifiedUsername"])
	adminpassword := string(adminSecret[MSecretPasswordKey])
	fullServerName := string(adminSecret["fullyQualifiedServerName"])
	fmt.Println("full server is :" + fullServerName)

	db, err := s.ConnectToSqlDb(ctx, MDriverName, fullServerName, instance.Spec.DbName, MSqlServerPort, adminuser, adminpassword)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		if strings.Contains(err.Error(), "is not allowed to connect to this MySQL server") {

			//for the ip address has no access to server, stop the reconcile and delete the user from controller
			return false, nil
		}
		if strings.Contains(err.Error(), "An internal error has occurred") {
			// there is nothing much we can do here - cycle forever
			return true, nil
		}
		return false, err
	}

	var userSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		userSecretClient = options.SecretClient
	} else {
		userSecretClient = s.SecretClient
	}

	userkey := GetNamespacedName(instance, userSecretClient)
	userSecret, err := userSecretClient.Get(ctx, userkey)
	if err != nil {

		return false, nil
	}

	user := string(userSecret[MSecretUsernameKey])

	err = s.DropUser(ctx, db, user)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete MySqlUser failed with %s", err.Error())
		return false, err
	}

	// Once the user has been dropped, also delete their secrets.
	s.DeleteSecrets(ctx, instance, mysqlUserSecretClient)

	instance.Status.Message = fmt.Sprintf("Delete MySqlUser succeeded")

	return false, nil
}

// GetParents gets the parents of the user
func (s *MySqlUserManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
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
			Target: &v1alpha1.MySQLDatabase{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &v1alpha2.MySQLServer{},
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
func (s *MySqlUserManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := s.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (s *MySqlUserManager) convert(obj runtime.Object) (*v1alpha1.MySQLUser, error) {
	local, ok := obj.(*v1alpha1.MySQLUser)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
