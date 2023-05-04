// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlaction

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type AzureSqlActionManager struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlActionManager(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlActionManager {
	return &AzureSqlActionManager{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

func (s *AzureSqlActionManager) UpdateUserPassword(
	ctx context.Context,
	groupName string,
	serverName string,
	dbUser string,
	dbName string,
	adminSecretKey secrets.SecretKey,
	adminSecretClient secrets.SecretClient,
	userSecretClient secrets.SecretClient) error {

	adminSecret, err := adminSecretClient.Get(ctx, adminSecretKey)
	if err != nil {
		return err
	}

	azuresqluserManager := azuresqluser.NewAzureSqlUserManager(s.Creds, userSecretClient, s.Scheme)
	db, err := azuresqluserManager.ConnectToSqlDb(ctx, "sqlserver", serverName, dbName, 1433, string(adminSecret["username"]), string(adminSecret["password"]))
	if err != nil {
		return err
	}
	defer db.Close()

	instance := &azurev1alpha1.AzureSQLUser{
		TypeMeta: metav1.TypeMeta{
			Kind: reflect.TypeOf(azurev1alpha1.AzureSQLUser{}).Name(),
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      dbUser,
			Namespace: adminSecretKey.Namespace,
		},
		Spec: azurev1alpha1.AzureSQLUserSpec{
			Server: serverName,
			DbName: dbName,
		},
	}

	userSecret := azuresqluserManager.GetOrPrepareSecret(ctx, instance, userSecretClient)
	// reset user from secret in case it was loaded
	userExists, err := azuresqluserManager.UserExists(ctx, db, string(userSecret["username"]))
	if err != nil {
		return errors.Wrap(err, "failed checking for user")
	}

	if !userExists {
		return fmt.Errorf("user does not exist")
	}

	password := helpers.NewPassword()
	userSecret["password"] = []byte(password)

	err = azuresqluserManager.UpdateUser(ctx, userSecret, db)
	if err != nil {
		return errors.Wrap(err, "error updating user credentials")
	}

	secretKey := azuresqluser.MakeSecretKey(userSecretClient, instance)
	err = userSecretClient.Upsert(
		ctx,
		secretKey,
		userSecret,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
	)
	if err != nil {
		return errors.Wrap(err, "failed to update secret")
	}

	return nil
}

// UpdateAdminPassword gets the server instance from Azure, updates the admin password
// for the server and stores the new password in the secret
func (s *AzureSqlActionManager) UpdateAdminPassword(
	ctx context.Context,
	subscriptionID string,
	groupName string,
	serverName string,
	secretKey secrets.SecretKey,
	secretClient secrets.SecretClient,
) error {

	azuresqlserverManager := azuresqlserver.NewAzureSqlServerManager(s.Creds, secretClient, s.Scheme)
	// Get the SQL server instance
	server, err := azuresqlserverManager.GetServer(ctx, subscriptionID, groupName, serverName)
	if err != nil {
		return err
	}

	// We were able to get the server instance from Azure, so we proceed to update the admin password
	azureSqlServerProperties := azuresqlshared.SQLServerProperties{
		AdministratorLogin:         server.ServerProperties.AdministratorLogin,
		AdministratorLoginPassword: server.ServerProperties.AdministratorLoginPassword,
	}

	// Get the secret from the secretclient. If we cannot get this we return err right away.
	data, err := secretClient.Get(ctx, secretKey)
	if err != nil {
		return err
	}

	// Generate a new password
	newPassword := helpers.NewPassword()
	azureSqlServerProperties.AdministratorLoginPassword = to.StringPtr(newPassword)

	// Update the SQL server with the newly generated password
	_, _, err = azuresqlserverManager.CreateOrUpdateSQLServer(
		ctx,
		subscriptionID,
		groupName,
		*server.Location,
		serverName,
		server.Tags,
		azureSqlServerProperties,
		true)

	if err != nil {
		azerr := errhelp.NewAzureError(err)
		if !strings.Contains(azerr.Type, errhelp.AsyncOpIncompleteError) {
			return err
		}
	}

	// Update the secret with the new password
	data["password"] = []byte(*azureSqlServerProperties.AdministratorLoginPassword)
	err = secretClient.Upsert(
		ctx,
		secretKey,
		data,
	)
	if err != nil {
		return err
	}

	return nil
}
