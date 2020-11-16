// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"

	"github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

type PSQLServerClient struct {
	Creds        config.Credentials
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewPSQLServerClient(creds config.Credentials, secretclient secrets.SecretClient, scheme *runtime.Scheme) *PSQLServerClient {
	return &PSQLServerClient{
		Creds:        creds,
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

// NewARMClient returns a new manager (but as an ARMClient).
func NewARMClient(creds config.Credentials, secretClient secrets.SecretClient, scheme *runtime.Scheme) resourcemanager.ARMClient {
	return NewPSQLServerClient(creds, secretClient, scheme)
}

func getPSQLServersClient(creds config.Credentials) (psql.ServersClient, error) {
	serversClient := psql.NewServersClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return psql.ServersClient{}, err
	}
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient, nil
}

func getPSQLCheckNameAvailabilityClient(creds config.Credentials) (psql.CheckNameAvailabilityClient, error) {
	nameavailabilityClient := psql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), creds.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer(creds)
	if err != nil {
		return psql.CheckNameAvailabilityClient{}, err
	}
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient, nil
}

func (c *PSQLServerClient) CheckServerNameAvailability(ctx context.Context, servername string) (bool, error) {

	client, err := getPSQLCheckNameAvailabilityClient(c.Creds)
	if err != nil {
		return false, err
	}

	resourceType := "Microsoft.DBforPostgreSQL/servers"

	nameAvailabilityRequest := psql.NameAvailabilityRequest{
		Name: &servername,
		Type: &resourceType,
	}
	_, err = client.Execute(ctx, nameAvailabilityRequest)
	if err == nil { // Name available
		return true, nil
	}
	return false, err

}

func (c *PSQLServerClient) CreateServerIfValid(ctx context.Context,
	instance v1alpha2.PostgreSQLServer,
	tags map[string]*string,
	skuInfo psql.Sku, adminlogin string,
	adminpassword string,
	createmode psql.CreateMode,
	hash string) (pollingURL string, server psql.Server, err error) {

	client, err := getPSQLServersClient(c.Creds)
	if err != nil {
		return "", psql.Server{}, err
	}

	// Check if name is valid if this is the first create call
	valid, err := c.CheckServerNameAvailability(ctx, instance.Name)
	if !valid {
		return "", psql.Server{}, err
	}

	var serverProperties psql.BasicServerPropertiesForCreate
	var skuData *psql.Sku
	var storageProfile *psql.StorageProfile
	if instance.Spec.StorageProfile != nil {
		obj := psql.StorageProfile(*instance.Spec.StorageProfile)
		storageProfile = &obj
	}

	if createmode == psql.CreateModeReplica {
		serverProperties = &psql.ServerPropertiesForReplica{
			SourceServerID: to.StringPtr(instance.Spec.ReplicaProperties.SourceServerId),
			CreateMode:     psql.CreateModeReplica,
			StorageProfile: storageProfile,
		}

	} else {
		serverProperties = &psql.ServerPropertiesForDefaultCreate{
			AdministratorLogin:         &adminlogin,
			AdministratorLoginPassword: &adminpassword,
			Version:                    psql.ServerVersion(instance.Spec.ServerVersion),
			SslEnforcement:             psql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
			CreateMode:                 psql.CreateModeServerPropertiesForCreate,
			StorageProfile:             storageProfile,
		}
		skuData = &skuInfo
	}

	if hash != instance.Status.SpecHash && instance.Status.SpecHash != "" {
		res, err := client.Update(
			ctx,
			instance.Spec.ResourceGroup,
			instance.Name,
			psql.ServerUpdateParameters{
				Tags: tags,
				Sku:  skuData,
				ServerUpdateParametersProperties: &psql.ServerUpdateParametersProperties{
					StorageProfile: storageProfile,
					Version:        psql.ServerVersion(instance.Spec.ServerVersion),
					SslEnforcement: psql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
				},
			},
		)
		if err != nil {
			return "", psql.Server{}, err
		}

		pollingURL = res.PollingURL()
		server, err = res.Result(client)

	} else {
		res, err := client.Create(
			ctx,
			instance.Spec.ResourceGroup,
			instance.Name,
			psql.ServerForCreate{
				Location:   &instance.Spec.Location,
				Tags:       tags,
				Properties: serverProperties,
				Sku:        skuData,
			},
		)
		if err != nil {
			return "", psql.Server{}, err
		}

		pollingURL = res.PollingURL()
		server, err = res.Result(client)
	}

	return pollingURL, server, err
}

func (c *PSQLServerClient) DeleteServer(ctx context.Context, resourcegroup string, servername string) (status string, err error) {

	client, err := getPSQLServersClient(c.Creds)
	if err != nil {
		return "", err
	}

	_, err = client.Get(ctx, resourcegroup, servername)
	if err == nil { // Server present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername)
		return future.Status(), err
	}
	// Server not present so return success anyway
	return "Server not present", nil

}

func (c *PSQLServerClient) GetServer(ctx context.Context, resourcegroup string, servername string) (server psql.Server, err error) {

	client, err := getPSQLServersClient(c.Creds)
	if err != nil {
		return psql.Server{}, err
	}

	return client.Get(ctx, resourcegroup, servername)
}

func (c *PSQLServerClient) AddServerCredsToSecrets(ctx context.Context, data map[string][]byte, instance *v1alpha2.PostgreSQLServer) error {
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	err := c.SecretClient.Upsert(ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(c.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *PSQLServerClient) UpdateSecretWithFullServerName(ctx context.Context, data map[string][]byte, instance *v1alpha2.PostgreSQLServer, fullServerName string) error {
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	data["fullyQualifiedServerName"] = []byte(fullServerName)

	err := c.SecretClient.Upsert(ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(c.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (c *PSQLServerClient) GetOrPrepareSecret(ctx context.Context, instance *v1alpha2.PostgreSQLServer) (map[string][]byte, error) {
	usernameLength := 8

	secret := map[string][]byte{}

	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
	if stored, err := c.SecretClient.Get(ctx, secretKey); err == nil {
		return stored, nil
	}

	randomUsername := helpers.GenerateRandomUsername(usernameLength)
	randomPassword := helpers.NewPassword()

	secret["username"] = []byte(randomUsername)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, instance.Name))
	secret["password"] = []byte(randomPassword)
	secret["postgreSqlServerName"] = []byte(instance.Name)

	return secret, nil
}
