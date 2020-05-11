// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"strings"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
)

type MySQLServerClient struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewMySQLServerClient(secretclient secrets.SecretClient, scheme *runtime.Scheme) *MySQLServerClient {
	return &MySQLServerClient{
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getMySQLServersClient() mysql.ServersClient {
	serversClient := mysql.NewServersClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

func getMySQLCheckNameAvailabilityClient() mysql.CheckNameAvailabilityClient {
	nameavailabilityClient := mysql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient
}

func (m *MySQLServerClient) CheckServerNameAvailability(ctx context.Context, servername string) (bool, error) {

	client := getMySQLCheckNameAvailabilityClient()

	resourceType := "Microsoft.DBforMySQL/servers"

	nameAvailabilityRequest := mysql.NameAvailabilityRequest{
		Name: &servername,
		Type: &resourceType,
	}
	_, err := client.Execute(ctx, nameAvailabilityRequest)
	if err == nil { // Name available
		return true, nil
	}
	return false, err

}

func (m *MySQLServerClient) CreateServerIfValid(ctx context.Context, servername string, resourcegroup string, location string, tags map[string]*string, serverversion mysql.ServerVersion, sslenforcement mysql.SslEnforcementEnum, skuInfo mysql.Sku, adminlogin string, adminpassword string, createmode string, sourceserver string) (pollingURL string, server mysql.Server, err error) {

	client := getMySQLServersClient()

	// Check if name is valid if this is the first create call
	valid, err := m.CheckServerNameAvailability(ctx, servername)
	if !valid {
		return "", server, err
	}

	var result mysql.ServersCreateFuture
	var serverProperties mysql.BasicServerPropertiesForCreate
	var skuData *mysql.Sku

	if strings.EqualFold(createmode, "replica") {
		serverProperties = &mysql.ServerPropertiesForReplica{
			SourceServerID: to.StringPtr(sourceserver),
			CreateMode:     mysql.CreateModeReplica,
		}

	} else {
		serverProperties = &mysql.ServerPropertiesForDefaultCreate{
			AdministratorLogin:         &adminlogin,
			AdministratorLoginPassword: &adminpassword,
			Version:                    serverversion,
			SslEnforcement:             sslenforcement,
			//StorageProfile: &mysql.StorageProfile{},
			CreateMode: mysql.CreateModeServerPropertiesForCreate,
		}
		skuData = &skuInfo

	}

	result, _ = client.Create(
		ctx,
		resourcegroup,
		servername,
		mysql.ServerForCreate{
			Location:   &location,
			Tags:       tags,
			Properties: serverProperties,
			Sku:        skuData,
		},
	)

	res, err := result.Result(client)
	return result.PollingURL(), res, err

}

func (m *MySQLServerClient) DeleteServer(ctx context.Context, resourcegroup string, servername string) (status string, err error) {

	client := getMySQLServersClient()

	_, err = client.Get(ctx, resourcegroup, servername)
	if err == nil { // Server present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername)
		return future.Status(), err
	}
	// Server not present so return success anyway
	return "Server not present", nil

}

func (m *MySQLServerClient) GetServer(ctx context.Context, resourcegroup string, servername string) (server mysql.Server, err error) {

	client := getMySQLServersClient()
	return client.Get(ctx, resourcegroup, servername)
}
