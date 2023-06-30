// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package psqluser

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	keyvaultSecrets "github.com/Azure/azure-service-operator/pkg/secrets/keyvault"

	_ "github.com/lib/pq" //pglib
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure that user exists
func (m *PostgreSqlUserManager) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	instance, err := m.convert(obj)
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

	adminSecretClient := m.SecretClient

	adminSecretName := instance.Spec.AdminSecret
	if len(instance.Spec.AdminSecret) == 0 {
		adminSecretName = instance.Spec.Server
	}

	// TODO: Is there a better way to get TypeMeta here?
	adminSecretKey := secrets.SecretKey{Name: adminSecretName, Namespace: instance.Namespace, Kind: reflect.TypeOf(v1alpha2.PostgreSQLServer{}).Name()}

	var sqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		sqlUserSecretClient = options.SecretClient
	} else {
		sqlUserSecretClient = m.SecretClient
	}

	// if the admin secret keyvault is not specified, fall back to configured secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(
			instance.Spec.AdminSecretKeyVault,
			m.Creds,
			m.SecretClient.GetSecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())
	}

	// get admin creds for server
	adminSecret, err := adminSecretClient.Get(ctx, adminSecretKey)
	if err != nil {
		err = errors.Wrap(err, "PostgreSQLServer admin secret not found")
		instance.Status.Provisioning = false
		instance.Status.Message = err.Error()
		return false, nil
	}

	adminUser := string(adminSecret["fullyQualifiedUsername"])
	adminPassword := string(adminSecret[PSecretPasswordKey])

	_, err = m.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
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
			return false, nil
		}
	}

	fullServerName := string(adminSecret["fullyQualifiedServerName"])
	db, err := m.ConnectToSqlDb(ctx, PDriverName, fullServerName, instance.Spec.DbName, PSqlServerPort, adminUser, adminPassword)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false

		// catch firewall issue - keep cycling until it clears up
		if strings.Contains(err.Error(), "no pg_hba.conf entry for host") {
			instance.Status.Message = errhelp.StripErrorIDs(err) + "\nThe IP address is not allowed to access the server. Modify the firewall rule to include the IP address."
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
	userSecret := m.GetOrPrepareSecret(ctx, instance, sqlUserSecretClient)
	// reset user from secret in case it was loaded
	user := string(userSecret[PSecretUsernameKey])
	if user == "" {
		user = fmt.Sprintf(requestedUsername)
		userSecret[PSecretUsernameKey] = []byte(user)
	}

	// Publishing the user secret:
	// We do this first so if the keyvault does not have right permissions we will not proceed to creating the user
	err = sqlUserSecretClient.Upsert(
		ctx,
		secretKey,
		userSecret,
		secrets.WithOwner(instance),
		secrets.WithScheme(m.Scheme),
	)
	if err != nil {
		instance.Status.Message = "failed to update secret, err: " + err.Error()
		return false, err
	}

	userExists, err := m.UserExists(ctx, db, string(userSecret[PSecretUsernameKey]))
	if err != nil {
		instance.Status.Message = fmt.Sprintf("failed checking for user, err: %s", err)
		return false, nil
	}

	if !userExists {
		user, err = m.CreateUser(ctx, userSecret, db)
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

	err = m.GrantUserRoles(ctx, user, instance.Spec.Roles, db)
	if err != nil {
		instance.Status.Message = "GrantUserRoles failed"
		return false, fmt.Errorf("GrantUserRoles failed")
	}

	instance.Status.Provisioned = true
	instance.Status.State = "Succeeded"
	instance.Status.Message = resourcemanager.SuccessMsg
	// reconcile done
	return true, nil
}

// Delete deletes a user
func (m *PostgreSqlUserManager) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	instance, err := m.convert(obj)
	if err != nil {
		return false, err
	}

	adminSecretClient := m.SecretClient

	adminSecretName := instance.Spec.AdminSecret

	if len(instance.Spec.AdminSecret) == 0 {
		adminSecretName = instance.Spec.Server
	}

	adminSecretKey := secrets.SecretKey{Name: adminSecretName, Namespace: instance.Namespace, Kind: reflect.TypeOf(v1alpha2.PostgreSQLServer{}).Name()}

	// if the admin secret keyvault is not specified, fall back to configured secretclient
	if len(instance.Spec.AdminSecretKeyVault) != 0 {
		adminSecretClient = keyvaultSecrets.New(
			instance.Spec.AdminSecretKeyVault,
			m.Creds,
			m.SecretClient.GetSecretNamingVersion(),
			config.PurgeDeletedKeyVaultSecrets(),
			config.RecoverSoftDeletedKeyVaultSecrets())
	}

	adminSecret, err := adminSecretClient.Get(ctx, adminSecretKey)
	if err != nil {
		// assuming if the admin secret is gone the sql server is too
		return false, nil
	}

	// short circuit connection if database doesn't exist
	_, err = m.GetDB(ctx, instance.Spec.ResourceGroup, instance.Spec.Server, instance.Spec.DbName)
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

	adminUser := string(adminSecret["fullyQualifiedUsername"])
	adminPassword := string(adminSecret[PSecretPasswordKey])
	fullServerName := string(adminSecret["fullyQualifiedServerName"])

	db, err := m.ConnectToSqlDb(ctx, PDriverName, fullServerName, instance.Spec.DbName, PSqlServerPort, adminUser, adminPassword)
	if err != nil {
		instance.Status.Message = errhelp.StripErrorIDs(err)
		if strings.Contains(err.Error(), "no pg_hba.conf entry for host") {
			// The error indicates the client IP has no access to server.
			instance.Status.Message = errhelp.StripErrorIDs(err) + "\nThe IP address is not allowed to access the server. Modify the firewall rule to include the IP address."
			//Stop the reconcile and delete the user from service operator side.
			return false, nil
		}
		//stop the reconcile with unkown error
		return false, err
	}
	defer db.Close()

	var psqlUserSecretClient secrets.SecretClient
	if options.SecretClient != nil {
		psqlUserSecretClient = options.SecretClient
	} else {
		psqlUserSecretClient = m.SecretClient
	}

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	userSecret, err := psqlUserSecretClient.Get(ctx, secretKey)
	if err != nil {

		return false, nil
	}

	user := string(userSecret[PSecretUsernameKey])

	exists, err := m.UserExists(ctx, db, user)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete PostgreSqlUser failed with %s", err.Error())
		return true, err
	}
	if !exists {

		m.DeleteSecrets(ctx, instance, psqlUserSecretClient)
		instance.Status.Message = fmt.Sprintf("The user %s doesn't exist", user)
		//User doesn't exist. Stop the reconcile.
		return false, nil
	}

	err = m.DropUser(ctx, db, user)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Delete PostgreSqlUser failed with %s", err.Error())
		//stop the reconcile with err
		return false, err
	}

	// Once the user has been dropped, also delete their secrets.
	m.DeleteSecrets(ctx, instance, psqlUserSecretClient)

	instance.Status.Message = "Delete PostgreSqlUser succeeded"

	// no err, no requeue, reconcile will stop
	return false, nil
}

// GetParents gets the parents of the user
func (m *PostgreSqlUserManager) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.DbName,
			},
			Target: &v1alpha1.PostgreSQLDatabase{},
		},
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.Server,
			},
			Target: &v1alpha2.PostgreSQLServer{},
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
func (m *PostgreSqlUserManager) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *PostgreSqlUserManager) convert(obj runtime.Object) (*v1alpha1.PostgreSQLUser, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLUser)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}
