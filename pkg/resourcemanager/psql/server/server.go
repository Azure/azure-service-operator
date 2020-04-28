// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"
	"strings"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/config"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/iam"
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

type PSQLServerClient struct {
	SecretClient secrets.SecretClient
	Scheme       *runtime.Scheme
}

func NewPSQLServerClient(secretclient secrets.SecretClient, scheme *runtime.Scheme) *PSQLServerClient {
	return &PSQLServerClient{
		SecretClient: secretclient,
		Scheme:       scheme,
	}
}

func getPSQLServersClient() (psql.ServersClient, error) {
	serversClient := psql.NewServersClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return psql.ServersClient{}, err
	}
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient, nil
}

func getPSQLCheckNameAvailabilityClient() (psql.CheckNameAvailabilityClient, error) {
	nameavailabilityClient := psql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, err := iam.GetResourceManagementAuthorizer()
	if err != nil {
		return psql.CheckNameAvailabilityClient{}, err
	}
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient, nil
}

func (p *PSQLServerClient) CheckServerNameAvailability(ctx context.Context, servername string) (bool, error) {

	client, err := getPSQLCheckNameAvailabilityClient()
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

func (p *PSQLServerClient) CreateServerIfValid(ctx context.Context,
	servername string,
	resourcegroup string,
	location string,
	tags map[string]*string,
	serverversion psql.ServerVersion,
	sslenforcement psql.SslEnforcementEnum,
	skuInfo psql.Sku, adminlogin string,
	adminpassword string,
	createmode string,
	sourceserver string) (pollingURL string, server psql.Server, err error) {

	client, err := getPSQLServersClient()
	if err != nil {
		return "", psql.Server{}, err
	}

	// Check if name is valid if this is the first create call
	valid, err := p.CheckServerNameAvailability(ctx, servername)
	if !valid {
		return "", psql.Server{}, err
	}

	var result pqsl.ServersCreateFuture

	if strings.EqualFold(createmode, string(psql.CreateModeReplica)) {
		result, err = client.Create(
			ctx,
			resourcegroup,
			servername,
			psql.ServerForCreate{
				Location: &location,
				Tags:     tags,
				Properties: &psql.ServerPropertiesForReplica{
					SourceServerID: to.StringPtr(sourceserver),
					CreateMode:     psql.CreateModeReplica,
				},
			},
		)
	} else {
		result, err = client.Create(
			ctx,
			resourcegroup,
			servername,
			psql.ServerForCreate{
				Location: &location,
				Tags:     tags,
				Properties: &psql.ServerPropertiesForDefaultCreate{
					AdministratorLogin:         &adminlogin,
					AdministratorLoginPassword: &adminpassword,
					Version:                    serverversion,
					SslEnforcement:             sslenforcement,
					CreateMode:                 psql.CreateModeServerPropertiesForCreate,
				},
				Sku: &skuInfo,
			},
		)

	}

	if err != nil {
		return "", psql.Server{}, err
	}
	res, err := result.Result(client)
	return result.PollingURL(), res, err
}

func (p *PSQLServerClient) DeleteServer(ctx context.Context, resourcegroup string, servername string) (status string, err error) {

	client, err := getPSQLServersClient()
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

func (p *PSQLServerClient) GetServer(ctx context.Context, resourcegroup string, servername string) (server psql.Server, err error) {

	client, err := getPSQLServersClient()
	if err != nil {
		return psql.Server{}, err
	}

	return client.Get(ctx, resourcegroup, servername)
}

func (p *PSQLServerClient) AddServerCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.PostgreSQLServer) error {
	key := types.NamespacedName{
		Name:      secretName,
		Namespace: instance.Namespace,
	}

	err := p.SecretClient.Upsert(ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(p.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *PSQLServerClient) UpdateSecretWithFullServerName(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.PostgreSQLServer, fullservername string) error {
	key := types.NamespacedName{
		Name:      secretName,
		Namespace: instance.Namespace,
	}

	data["fullyQualifiedServerName"] = []byte(fullservername)

	err := p.SecretClient.Upsert(ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(p.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *PSQLServerClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.PostgreSQLServer) (map[string][]byte, error) {
	name := instance.Name

	usernameLength := 8

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := p.SecretClient.Get(ctx, key); err == nil {
		return stored, nil
	}

	randomUsername := helpers.GenerateRandomUsername(usernameLength)
	randomPassword := helpers.NewPassword()

	secret["username"] = []byte(randomUsername)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, name))
	secret["password"] = []byte(randomPassword)
	secret["postgreSqlServerName"] = []byte(name)

	return secret, nil
}
