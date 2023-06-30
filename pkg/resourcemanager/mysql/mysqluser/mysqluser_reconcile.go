// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package mysqluser

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	_ "github.com/go-sql-driver/mysql" //sql drive link
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	mysqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql/server"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/mysql"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"
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
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(
			instance.Spec.AdminSecretKeyVault,
			s.Creds,
			s.SecretClient.GetSecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())
	}

	adminSecretKey := secrets.SecretKey{Name: instance.Spec.GetAdminSecretName(), Namespace: instance.Namespace, Kind: reflect.TypeOf(v1alpha2.MySQLServer{}).Name()}

	// get admin creds for server
	adminSecret, err := adminSecretClient.Get(ctx, adminSecretKey)
	if err != nil {
		err = errors.Wrap(err, "MySQLServer admin secret not found")
		instance.Status.Provisioning = false
		instance.Status.Message = err.Error()
		return false, nil
	}

	var mysqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		mysqlUserSecretClient = options.SecretClient
	} else {
		mysqlUserSecretClient = s.SecretClient
	}

	_, err = s.GetServer(ctx, instance.Spec.ResourceGroup, instance.Spec.Server)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		if mysql.IsErrorResourceNotFound(err) {
			return false, nil
		}

		return false, err
	}

	adminUser := string(adminSecret[mysqlserver.FullyQualifiedUsernameSecretKey])
	adminPassword := string(adminSecret[mysqlserver.PasswordSecretKey])
	fullServerName := string(adminSecret[mysqlserver.FullyQualifiedServerNameSecretKey])

	db, err := mysql.ConnectToSqlDB(
		ctx,
		mysql.DriverName,
		fullServerName,
		mysql.SystemDatabase,
		mysql.ServerPort,
		adminUser,
		adminPassword)
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
	defer db.Close()

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	// create or get new user secret
	userSecret := s.GetOrPrepareSecret(ctx, instance, mysqlUserSecretClient)
	// reset user from secret in case it was loaded
	user := string(userSecret[MSecretUsernameKey])
	if user == "" {
		user = fmt.Sprintf(requestedUsername)
		userSecret[MSecretUsernameKey] = []byte(user)
	}

	// Publishing the user secret:
	// We do this first so if the keyvault does not have right permissions we will not proceed to creating the user
	err = mysqlUserSecretClient.Upsert(
		ctx,
		secretKey,
		userSecret,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
	)
	if err != nil {
		instance.Status.Message = "failed to update secret, err: " + err.Error()
		return false, err
	}

	user, err = s.CreateUser(ctx, userSecret, db)
	if err != nil {
		instance.Status.Message = "failed creating user, err: " + err.Error()
		return false, err
	}

	err = mysql.EnsureUserServerRoles(ctx, db, user, instance.Spec.Roles)
	if err != nil {
		err = errors.Wrap(err, "ensuring server roles")
		instance.Status.Message = err.Error()
		return false, err
	}

	err = mysql.EnsureUserDatabaseRoles(ctx, db, user, instance.Spec.DatabaseRoles)
	if err != nil {
		err = errors.Wrap(err, "ensuring database roles")
		instance.Status.Message = err.Error()
		return false, err
	}

	instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
	instance.Status.State = "Succeeded"

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

	adminSecretName := instance.Spec.AdminSecret

	if len(instance.Spec.AdminSecret) == 0 {
		adminSecretName = instance.Spec.Server
	}
	adminSecretKey := secrets.SecretKey{Name: adminSecretName, Namespace: instance.Namespace, Kind: reflect.TypeOf(v1alpha2.MySQLServer{}).Name()}

	var mysqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		mysqlUserSecretClient = options.SecretClient
	} else {
		mysqlUserSecretClient = s.SecretClient
	}

	// if the admin secret keyvault is not specified, fall back to configured secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(
			instance.Spec.AdminSecretKeyVault,
			s.Creds,
			s.SecretClient.GetSecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())
	}

	adminSecret, err := adminSecretClient.Get(ctx, adminSecretKey)
	if err != nil {
		// assuming if the admin secret is gone the sql server is too
		return false, nil
	}

	// short circuit connection if database doesn't exist
	_, err = s.GetServer(ctx, instance.Spec.ResourceGroup, instance.Spec.Server)
	if err != nil {
		instance.Status.Message = err.Error()
		if mysql.IsErrorResourceNotFound(err) {
			return false, nil
		}

		return false, err
	}

	adminUser := string(adminSecret["fullyQualifiedUsername"])
	adminPassword := string(adminSecret[MSecretPasswordKey])
	fullServerName := string(adminSecret["fullyQualifiedServerName"])

	db, err := mysql.ConnectToSqlDB(
		ctx,
		mysql.DriverName,
		fullServerName,
		mysql.SystemDatabase,
		mysql.ServerPort,
		adminUser,
		adminPassword,
	)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		if strings.Contains(err.Error(), "is not allowed to connect to this MySQL server") {

			//for the ip address has no access to server, stop the reconcile and delete the user from controller
			return false, nil
		}
		return false, err
	}
	defer db.Close()

	var userSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		userSecretClient = options.SecretClient
	} else {
		userSecretClient = s.SecretClient
	}

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	userSecret, err := userSecretClient.Get(ctx, secretKey)
	if err != nil {
		return false, nil
	}

	user := string(userSecret[MSecretUsernameKey])

	err = mysql.DropUser(ctx, db, user)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete MySqlUser failed with %s", err.Error())
		return false, err
	}

	// Once the user has been dropped, also delete their secrets.
	s.DeleteSecrets(ctx, instance, mysqlUserSecretClient)

	instance.Status.Message = "Delete MySqlUser succeeded"

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
	st := v1alpha1.ASOStatus(instance.Status)
	return &st, nil
}

func (s *MySqlUserManager) convert(obj runtime.Object) (*v1alpha2.MySQLUser, error) {
	local, ok := obj.(*v1alpha2.MySQLUser)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
