// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"

	psql "github.com/Azure/azure-sdk-for-go/services/postgresql/mgmt/2017-12-01/postgresql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
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

func getPSQLServersClient() psql.ServersClient {
	serversClient := psql.NewServersClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	serversClient.Authorizer = a
	serversClient.AddToUserAgent(config.UserAgent())
	return serversClient
}

func getPSQLCheckNameAvailabilityClient() psql.CheckNameAvailabilityClient {
	nameavailabilityClient := psql.NewCheckNameAvailabilityClientWithBaseURI(config.BaseURI(), config.SubscriptionID())
	a, _ := iam.GetResourceManagementAuthorizer()
	nameavailabilityClient.Authorizer = a
	nameavailabilityClient.AddToUserAgent(config.UserAgent())
	return nameavailabilityClient
}

func (p *PSQLServerClient) CheckServerNameAvailability(ctx context.Context, servername string) (bool, error) {

	client := getPSQLCheckNameAvailabilityClient()

	resourceType := "server"

	nameAvailabilityRequest := psql.NameAvailabilityRequest{
		Name: &servername,
		Type: &resourceType,
	}
	_, err := client.Execute(ctx, nameAvailabilityRequest)
	if err == nil { // Name available
		return true, nil
	}
	return false, err

}
func (p *PSQLServerClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		p.SecretClient = options.SecretClient
	}

	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := p.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}
	// Update secret
	err = p.AddServerCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	client := getPSQLServersClient()

	// convert kube labels to expected tag format
	labels := map[string]*string{}
	for k, v := range instance.GetLabels() {
		labels[k] = &v
	}
	instance.Status.Provisioning = true
	// Check if this server already exists and its state if it does. This is required
	// to overcome the issue with the lack of idempotence of the Create call

	server, err := p.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		instance.Status.State = string(server.UserVisibleState)
		if server.UserVisibleState == "Ready" {
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.Message = resourcemanager.SuccessMsg
			return true, nil
		}
		return false, nil
	}

	adminlogin := string(secret["username"])
	adminpassword := string(secret["password"])

	skuInfo := psql.Sku{
		Name:     to.StringPtr(instance.Spec.Sku.Name),
		Tier:     psql.SkuTier(instance.Spec.Sku.Tier),
		Capacity: to.Int32Ptr(instance.Spec.Sku.Capacity),
		Size:     to.StringPtr(instance.Spec.Sku.Size),
		Family:   to.StringPtr(instance.Spec.Sku.Family),
	}
	future, err := p.CreateServerIfValid(
		ctx,
		instance.Name,
		instance.Spec.ResourceGroup,
		instance.Spec.Location,
		labels,
		psql.ServerVersion(instance.Spec.ServerVersion),
		psql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
		skuInfo,
		adminlogin,
		adminpassword,
	)

	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		catchUnrecoverableErrors := []string{
			errhelp.ProvisioningDisabled,
			errhelp.LocationNotAvailableForResourceType,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		if helpers.ContainsString(catchUnrecoverableErrors, azerr.Type) {
			// Unrecoverable error, so stop reconcilation
			instance.Status.Message = "Reconcilation hit unrecoverable error: " + err.Error()
			return true, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	instance.Status.State = string(server.UserVisibleState)

	server, err = future.Result(client)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
		}

		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			// most of these error technically mean the resource is actually not provisioning
			switch azerr.Type {
			case errhelp.AsyncOpIncompleteError:
				instance.Status.Provisioning = true
			}
			// reconciliation is not done but error is acceptable
			return false, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	instance.Status.State = string(server.UserVisibleState)

	if instance.Status.Provisioning {
		instance.Status.Provisioned = true
		instance.Status.Provisioning = false
		instance.Status.Message = resourcemanager.SuccessMsg
	} else {
		instance.Status.Provisioned = false
		instance.Status.Provisioning = true
	}

	return true, nil
}

func (p *PSQLServerClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		p.SecretClient = options.SecretClient
	}

	instance, err := p.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := p.DeleteServer(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err != nil {
		if !errhelp.IsAsynchronousOperationNotComplete(err) {
			return true, err
		}
	}
	instance.Status.State = status

	if err == nil {
		if status != "InProgress" {
			// Best case deletion of secrets
			key := types.NamespacedName{Name: instance.Name, Namespace: instance.Namespace}
			p.SecretClient.Delete(ctx, key)
			return false, nil
		}
	}

	return true, nil
}

func (p *PSQLServerClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := p.convert(obj)
	if err != nil {
		return nil, err
	}

	return []resourcemanager.KubeParent{
		{
			Key: types.NamespacedName{
				Namespace: instance.Namespace,
				Name:      instance.Spec.ResourceGroup,
			},
			Target: &azurev1alpha1.ResourceGroup{},
		},
	}, nil
}

func (g *PSQLServerClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := g.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (p *PSQLServerClient) convert(obj runtime.Object) (*v1alpha1.PostgreSQLServer, error) {
	local, ok := obj.(*v1alpha1.PostgreSQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (p *PSQLServerClient) CreateServerIfValid(ctx context.Context, servername string, resourcegroup string, location string, tags map[string]*string, serverversion psql.ServerVersion, sslenforcement psql.SslEnforcementEnum, skuInfo psql.Sku, adminlogin string, adminpassword string) (future psql.ServersCreateFuture, err error) {

	client := getPSQLServersClient()

	// Check if name is valid if this is the first create call
	valid, err := p.CheckServerNameAvailability(ctx, servername)
	if valid == false {
		return future, err
	}

	future, err = client.Create(
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
				//StorageProfile: &psql.StorageProfile{},
				CreateMode: psql.CreateModeServerPropertiesForCreate,
			},
			Sku: &skuInfo,
		},
	)

	return future, err
}

func (p *PSQLServerClient) DeleteServer(ctx context.Context, resourcegroup string, servername string) (status string, err error) {

	client := getPSQLServersClient()

	_, err = client.Get(ctx, resourcegroup, servername)
	if err == nil { // Server present, so go ahead and delete
		future, err := client.Delete(ctx, resourcegroup, servername)
		return future.Status(), err
	}
	// Server not present so return success anyway
	return "Server not present", nil

}

func (p *PSQLServerClient) GetServer(ctx context.Context, resourcegroup string, servername string) (server psql.Server, err error) {

	client := getPSQLServersClient()
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

func (p *PSQLServerClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.PostgreSQLServer) (map[string][]byte, error) {
	name := instance.Name

	usernameLength := 8
	passwordLength := 16

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := p.SecretClient.Get(ctx, key); err == nil {
		return stored, nil
	}

	randomUsername, err := helpers.GenerateRandomUsername(usernameLength, 0)
	if err != nil {
		return secret, err
	}

	randomPassword, err := helpers.GenerateRandomPassword(passwordLength)
	if err != nil {
		return secret, err
	}

	secret["username"] = []byte(randomUsername)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, name))
	secret["password"] = []byte(randomPassword)
	secret["postgreSqlServerName"] = []byte(name)
	// TODO: The below may not be right for non Azure public cloud.
	secret["fullyQualifiedServerName"] = []byte(name + ".postgres.database.azure.com")

	return secret, nil
}
