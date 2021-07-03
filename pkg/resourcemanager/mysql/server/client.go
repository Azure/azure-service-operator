// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type MySQLServerClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
	// KubeReader is used to read secrets in the case the customer has specified a secret containing their
	// MySQLServer admin username/password
	KubeReader   client.Reader
}

func NewMySQLServerClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme, kubeReader client.Reader) *MySQLServerClient {
	return &MySQLServerClient{
		Creds:        creds,
		SecretClient: secretClient,
		Scheme:       scheme,
		KubeReader:   kubeReader,
	}
}

func getMySQLServersClient(creds config.Credentials) mysql.ServersClient {
	serversClient := mysql.NewServersClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

func getMySQLCheckNameAvailabilityClient(creds config.Credentials) mysql.CheckNameAvailabilityClient {
	nameavailabilityClient := mysql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer(creds)
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient
}

func (m *MySQLServerClient) CheckServerNameAvailability(ctx context.Context, servername string) (bool, error) {

	client := getMySQLCheckNameAvailabilityClient(m.Creds)

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

func (m *MySQLServerClient) CreateServerIfValid(ctx context.Context, instance v1alpha2.MySQLServer, tags map[string]*string, skuInfo mysql.Sku, adminlogin string, adminpassword string, createmode mysql.CreateMode, hash string) (pollingURL string, server mysql.Server, err error) {

	client := getMySQLServersClient(m.Creds)

	// Check if name is valid if this is the first create call
	valid, err := m.CheckServerNameAvailability(ctx, instance.Name)
	if !valid {
		return "", server, err
	}

	var serverProperties mysql.BasicServerPropertiesForCreate
	var skuData *mysql.Sku
	var storageProfile *mysql.StorageProfile
	if instance.Spec.StorageProfile != nil {
		obj := mysql.StorageProfile(*instance.Spec.StorageProfile)
		storageProfile = &obj
	}

	if createmode == mysql.CreateModeReplica {
		serverProperties = &mysql.ServerPropertiesForReplica{
			SourceServerID: to.StringPtr(instance.Spec.ReplicaProperties.SourceServerId),
			CreateMode:     mysql.CreateModeReplica,
			StorageProfile: storageProfile,
		}

	} else {
		serverProperties = &mysql.ServerPropertiesForDefaultCreate{
			AdministratorLogin:         &adminlogin,
			AdministratorLoginPassword: &adminpassword,
			Version:                    mysql.ServerVersion(instance.Spec.ServerVersion),
			SslEnforcement:             mysql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
			CreateMode:                 mysql.CreateModeServerPropertiesForCreate,
			StorageProfile:             storageProfile,
		}
		skuData = &skuInfo
	}

	if hash != instance.Status.SpecHash && instance.Status.SpecHash != "" {
		res, err := client.Update(
			ctx,
			instance.Spec.ResourceGroup,
			instance.Name,
			mysql.ServerUpdateParameters{
				Tags: tags,
				Sku:  skuData,
				ServerUpdateParametersProperties: &mysql.ServerUpdateParametersProperties{
					StorageProfile: storageProfile,
					Version:        mysql.ServerVersion(instance.Spec.ServerVersion),
					SslEnforcement: mysql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
				},
			},
		)
		if err != nil {
			return "", mysql.Server{}, err
		}

		pollingURL = res.PollingURL()
		server, err = res.Result(client)

	} else {
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
	}

	return pollingURL, server, err

}

func (m *MySQLServerClient) DeleteServer(ctx context.Context, resourcegroup string, servername string) (status string, err error) {

	client := getMySQLServersClient(m.Creds)

	_, err = client.Get(ctx, resourcegroup, servername)
	if err == nil { // Server present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername)
		return future.Status(), err
	}
	// Server not present so return success anyway
	return "Server not present", nil

}

func (m *MySQLServerClient) GetServer(ctx context.Context, resourcegroup string, servername string) (server mysql.Server, err error) {

	client := getMySQLServersClient(m.Creds)
	return client.Get(ctx, resourcegroup, servername)
}
