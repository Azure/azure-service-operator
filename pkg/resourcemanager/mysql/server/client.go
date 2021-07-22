// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type MySQLServerClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
	// KubeReader is used to read secrets in the case the customer has specified a secret containing their
	// MySQLServer admin username/password
	KubeReader client.Reader
}

func NewMySQLServerClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme, kubeReader client.Reader) *MySQLServerClient {
	return &MySQLServerClient{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
		KubeReader:   kubeReader,
	}
}

func MakeMySQLServerAzureClient(creds config.Credentials) mysql.ServersClient {
	serversClient := mysql.NewServersClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

func MakeMySQLCheckNameAvailabilityAzureClient(creds config.Credentials) mysql.CheckNameAvailabilityClient {
	nameAvailabilityClient := mysql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	nameAvailabilityClient.Authorizer = a
	nameAvailabilityClient.AddToUserAgent(config.UserAgent())
	return nameAvailabilityClient
}

func (m *MySQLServerClient) CreateServerIfValid(
	ctx context.Context,
	instance v1alpha2.MySQLServer,
	tags map[string]*string,
	skuInfo mysql.Sku,
	adminUser string,
	adminPassword string,
	createMode mysql.CreateMode) (pollingURL string, server mysql.Server, err error) {

	client := MakeMySQLServerAzureClient(m.Creds)

	var serverProperties mysql.BasicServerPropertiesForCreate
	var skuData *mysql.Sku
	var storageProfile *mysql.StorageProfile
	if instance.Spec.StorageProfile != nil {
		obj := mysql.StorageProfile(*instance.Spec.StorageProfile)
		storageProfile = &obj
	}

	if createMode == mysql.CreateModeReplica {
		serverProperties = &mysql.ServerPropertiesForReplica{
			SourceServerID: to.StringPtr(instance.Spec.ReplicaProperties.SourceServerId),
			CreateMode:     mysql.CreateModeReplica,
			StorageProfile: storageProfile,
		}
	} else {
		serverProperties = &mysql.ServerPropertiesForDefaultCreate{
			AdministratorLogin:         &adminUser,
			AdministratorLoginPassword: &adminPassword,
			Version:                    mysql.ServerVersion(instance.Spec.ServerVersion),
			SslEnforcement:             mysql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
			CreateMode:                 mysql.CreateModeServerPropertiesForCreate,
			StorageProfile:             storageProfile,
		}
		skuData = &skuInfo
	}

	res, err := client.Create(
		ctx,
		instance.Spec.ResourceGroup,
		instance.Name,
		mysql.ServerForCreate{
			Location:   &instance.Spec.Location,
			Tags:       tags,
			Properties: serverProperties,
			Sku:        skuData,
		},
	)
	if err != nil {
		return "", mysql.Server{}, err
	}

	pollingURL = res.PollingURL()
	server, err = res.Result(client)

	return pollingURL, server, err

}

func (m *MySQLServerClient) DeleteServer(ctx context.Context, resourcegroup string, servername string) (status string, err error) {
	client := MakeMySQLServerAzureClient(m.Creds)

	_, err = client.Get(ctx, resourcegroup, servername)
	if err == nil { // Server present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername)
		return future.Status(), err
	}
	// Server not present so return success anyway
	return "Server not present", nil

}

func (m *MySQLServerClient) GetServer(ctx context.Context, resourcegroup string, servername string) (server mysql.Server, err error) {
	client := MakeMySQLServerAzureClient(m.Creds)
	return client.Get(ctx, resourcegroup, servername)
}
