// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlaction

import (
	"context"
	"fmt"
	"strings"

	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	azuresqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	azuresqluser "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqluser"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type AzureSqlActionManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlActionManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlActionManager {
	return &AzureSqlActionManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

func (s *AzureSqlActionManager) UpdateUserPassword(ctx context.Context, groupName string, serverName string, dbUser string, dbName string,
	secretKey types.NamespacedName, secretClient secrets.SecretClient, userSecretClient secrets.SecretClient) error {
	data, err := secretClient.Get(ctx, secretKey)
	if err != nil {
		return err
	}

	azuresqluserManager := azuresqluser.NewAzureSqlUserManager(userSecretClient, s.Scheme)
	db, err := azuresqluserManager.ConnectToSqlDb(ctx, "sqlserver", serverName, dbName, 1433, string(data["username"]), string(data["password"]))
	if err != nil {
		return err
	}

	instance := &azurev1alpha1.AzureSQLUser{
		ObjectMeta: metav1.ObjectMeta{
			Name:      dbUser,
			Namespace: secretKey.Namespace,
		},
		Spec: azurev1alpha1.AzureSQLUserSpec{
			Server: serverName,
			DbName: dbName,
		},
	}

	dbUserCustomNamespace := "azuresqluser-" + serverName + "-" + dbName

	DBSecret := azuresqluserManager.GetOrPrepareSecret(ctx, instance, userSecretClient, dbUserCustomNamespace)
	// reset user from secret in case it was loaded
	userExists, err := azuresqluserManager.UserExists(ctx, db, string(DBSecret["username"]))
	if err != nil {
		return fmt.Errorf("failed checking for user, err: %v", err)
	}

	if !userExists {
		return fmt.Errorf("user does not exist")
	}

	password := helpers.NewPassword()
	DBSecret["password"] = []byte(password)

	err = azuresqluserManager.UpdateUser(ctx, DBSecret, db)
	if err != nil {
		return fmt.Errorf("error updating user credentials: %v", err)
	}

	key := types.NamespacedName{Namespace: dbUserCustomNamespace, Name: dbUser}
	err = userSecretClient.Upsert(
		ctx,
		key,
		DBSecret,
		secrets.WithOwner(instance),
		secrets.WithScheme(s.Scheme),
	)
	if err != nil {
		return fmt.Errorf("failed to update secret: %v", err)
	}

	return nil
}

// UpdateAdminPassword gets the server instance from Azure, updates the admin password
// for the server and stores the new password in the secret
func (s *AzureSqlActionManager) UpdateAdminPassword(ctx context.Context, groupName string, serverName string, secretKey types.NamespacedName, secretClient secrets.SecretClient) error {

	azuresqlserverManager := azuresqlserver.NewAzureSqlServerManager(secretClient, s.Scheme)
	// Get the SQL server instance
	server, err := azuresqlserverManager.GetServer(ctx, groupName, serverName)
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
	_, _, err = azuresqlserverManager.CreateOrUpdateSQLServer(ctx, groupName, *server.Location, serverName, server.Tags, azureSqlServerProperties, true)

	if err != nil {
		azerr := errhelp.NewAzureErrorAzureError(err)
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
