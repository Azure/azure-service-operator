// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package azuresqlserver

import (
	"context"
	"errors"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/services/preview/sql/mgmt/2015-05-01-preview/sql"
	azuresqlshared "github.com/Azure/azure-service-operator/pkg/resourcemanager/azuresql/azuresqlshared"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
)

const typeOfService = "Microsoft.Sql/servers"

type AzureSqlServerManager struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewAzureSqlServerManager(secretClient secrets.SecretClient, scheme *runtime.Scheme) *AzureSqlServerManager {
	return &AzureSqlServerManager{
		SecretClient: secretClient,
		Scheme:       scheme,
	}
}

// DeleteSQLServer deletes a SQL server
func (sdk *AzureSqlServerManager) DeleteSQLServer(ctx context.Context, resourceGroupName string, serverName string) (result autorest.Response, err error) {
	result = autorest.Response{
		Response: &http.Response{
			StatusCode: 200,
		},
	}

	// check to see if the server exists, if it doesn't then short-circuit
	_, err = sdk.GetServer(ctx, resourceGroupName, serverName)
	if err != nil {
		return result, nil
	}

	serversClient := azuresqlshared.GetGoServersClient()
	future, err := serversClient.Delete(
		ctx,
		resourceGroupName,
		serverName,
	)
	if err != nil {
		return result, err
	}

	return future.Result(serversClient)
}

// GetServer returns a SQL server
func (_ *AzureSqlServerManager) GetServer(ctx context.Context, resourceGroupName string, serverName string) (result sql.Server, err error) {
	serversClient := azuresqlshared.GetGoServersClient()

	return serversClient.Get(
		ctx,
		resourceGroupName,
		serverName,
	)
}

// CreateOrUpdateSQLServer creates a SQL server in Azure
func (_ *AzureSqlServerManager) CreateOrUpdateSQLServer(ctx context.Context, resourceGroupName string, location string, serverName string, tags map[string]*string, properties azuresqlshared.SQLServerProperties, forceUpdate bool) (*sql.Server, error) {
	serversClient := azuresqlshared.GetGoServersClient()
	serverProp := azuresqlshared.SQLServerPropertiesToServer(properties)

	if forceUpdate == false {
		checkNameResult, _ := CheckNameAvailability(ctx, serverName)
		if checkNameResult.Reason == sql.AlreadyExists {
			return nil, errors.New("AlreadyExists")
		} else if checkNameResult.Reason == sql.Invalid {
			return nil, errors.New("InvalidServerName")
		}
	}

	// issue the creation
	future, err := serversClient.CreateOrUpdate(
		ctx,
		resourceGroupName,
		serverName,
		sql.Server{
			Location:         to.StringPtr(location),
			ServerProperties: &serverProp,
			Tags:             tags,
		})

	if err != nil {
		return nil, err
	}
	err = future.WaitForCompletionRef(ctx, serversClient.Client)
	if err != nil {
		return nil, err
	}

	result, err := future.Result(serversClient)
	return &result, err
}

func CheckNameAvailability(ctx context.Context, serverName string) (result sql.CheckNameAvailabilityResponse, err error) {
	serversClient := azuresqlshared.GetGoServersClient()

	response, err := serversClient.CheckNameAvailability(
		ctx,
		sql.CheckNameAvailabilityRequest{
			Name: to.StringPtr(serverName),
			Type: to.StringPtr(typeOfService),
		},
	)
	if err != nil {
		return result, err
	}

	return response, err
}
