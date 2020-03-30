// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlaction

import (
	"context"
	"strings"

	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	azuresqlserver "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlserver"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
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
	_, err = azuresqlserverManager.CreateOrUpdateSQLServer(ctx, groupName, *server.Location, serverName, server.Tags, azureSqlServerProperties, true)

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
