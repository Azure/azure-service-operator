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

	// Get the admin secret client and secret for the sql server
	adminSecretClient := s.SecretClient

	adminsecretName := instance.Spec.AdminSecret
	if len(instance.Spec.AdminSecret) == 0 {
		adminsecretName = instance.Spec.Server
	}

	adminKey := types.NamespacedName{Name: adminsecretName, Namespace: instance.Namespace}
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault)
		if len(instance.Spec.AdminSecret) != 0 {
			adminKey = types.NamespacedName{Name: instance.Spec.AdminSecret}
		}
	}

	adminSecret, err := adminSecretClient.Get(ctx, adminKey)
	if err != nil {
		instance.Status.Provisioning = false
		instance.Status.Message = fmt.Sprintf("admin secret : %s, not found in %s", adminKey.String(), reflect.TypeOf(adminSecretClient).Elem().Name())
		return false, nil
	}

	adminUser := string(adminSecret[SecretUsernameKey])
	adminPassword := string(adminSecret[SecretPasswordKey])

	// Get the user secret client for the sql server
	var sqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		sqlUserSecretClient = options.SecretClient
	} else {
		sqlUserSecretClient = s.SecretClient
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
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(requeuErrors, azerr.Type) {
			return false, nil
		}

		// if the database is busy, requeue
		errorString := err.Error()
		if strings.Contains(errorString, "Please retry the connection later") {
			return false, nil
		}

		// if this is an unmarshall error - igmore and continue, otherwise report error and requeue
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

	// Retrieve a secret that exists for this user spec or generate a new one for the user
	userSecret := s.GetOrPrepareSecret(ctx, instance, sqlUserSecretClient)

	// If this is a new user, GetOrPrepareSecet won't assign a username for use on the server-side so we generate one here
	userName := string(userSecret[SecretUsernameKey])
	if userName == "" {
		userName = fmt.Sprintf("%s-%s", requestedUsername, uuid.New())
		userSecret[SecretUsernameKey] = []byte(userName)
	}

	// Publish the user's credentials to the user secret store
	userKey := s.GetSecretNamespacedName(instance, sqlUserSecretClient)
	err = sqlUserSecretClient.Upsert(
		ctx,
		userKey,
		userSecret,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
	)
	if err != nil {
		instance.Status.Message = "failed to update secret, err: " + err.Error()
		return false, err
	}

	// Reconcile any requested custom secrets
	err = s.ReconcileCustomSecrets(ctx, instance, sqlUserSecretClient, userSecret)
	if err != nil {
		instance.Status.Message = "failed to update custom secrets, err: " + err.Error()
		return false, err
	}

	userExists, err := s.UserExists(ctx, db, string(userSecret[SecretUsernameKey]))
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %v", err)
		return false, nil
	}

	if !userExists {
		userName, err = s.CreateUser(ctx, userSecret, db)
		if err != nil {
			instance.Status.Message = "failed creating user, err: " + err.Error()
			return false, err
		}
	}

	// Apply roles to user
	if len(instance.Spec.Roles) == 0 {
		instance.Status.Message = "No roles specified for user"
		return false, fmt.Errorf("No roles specified for database user")
	}

	err = s.GrantUserRoles(ctx, userName, instance.Spec.Roles, db)
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

	adminSecretClient := s.SecretClient

	adminsecretName := instance.Spec.AdminSecret
	if len(instance.Spec.AdminSecret) == 0 {
		adminsecretName = instance.Spec.Server
	}
	adminKey := types.NamespacedName{Name: adminsecretName, Namespace: instance.Namespace}

	// if the admin secret keyvault is not specified, fall back to global secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(instance.Spec.AdminSecretKeyVault)
		if len(instance.Spec.AdminSecret) != 0 {
			adminKey = types.NamespacedName{Name: instance.Spec.AdminSecret}
		}
	}

	adminSecret, err := adminSecretClient.Get(ctx, adminKey)
	if err != nil {
		// assuming if the admin secret is gone the sql server is too
		return false, nil
	}

	// prepare the user secret client
	var sqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		sqlUserSecretClient = options.SecretClient
	} else {
		sqlUserSecretClient = s.SecretClient
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
