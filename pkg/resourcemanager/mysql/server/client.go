// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
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

	resourceType := "server"

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
func (m *MySQLServerClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		m.SecretClient = options.SecretClient
	}

	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := m.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}
	// Update secret
	err = m.AddServerCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	//client := getMySQLServersClient()

	// convert kube labels to expected tag format
	labels := map[string]*string{}
	for k, v := range instance.GetLabels() {
		labels[k] = &v
	}

	// Check if this server already exists and its state if it does. This is required
	// to overcome the issue with the lack of idempotence of the Create call

	server, err := m.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err == nil {
		instance.Status.State = string(server.UserVisibleState)
		if server.UserVisibleState == mysql.ServerStateReady {
			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.Message = resourcemanager.SuccessMsg
			return true, nil
		}
		return false, nil
	}

	if instance.Status.Provisioning {
		return false, nil
	}

	adminlogin := string(secret["username"])
	adminpassword := string(secret["password"])
	skuInfo := mysql.Sku{
		Name:     to.StringPtr(instance.Spec.Sku.Name),
		Tier:     mysql.SkuTier(instance.Spec.Sku.Tier),
		Capacity: to.Int32Ptr(instance.Spec.Sku.Capacity),
		Size:     to.StringPtr(instance.Spec.Sku.Size),
		Family:   to.StringPtr(instance.Spec.Sku.Family),
	}

	_, err = m.CreateServerIfValid(
		ctx,
		instance.Name,
		instance.Spec.ResourceGroup,
		instance.Spec.Location,
		labels,
		mysql.ServerVersion(instance.Spec.ServerVersion),
		mysql.SslEnforcementEnum(instance.Spec.SSLEnforcement),
		skuInfo,
		adminlogin,
		adminpassword,
	)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = err.Error()
		instance.Status.Provisioning = false
		azerr := errhelp.NewAzureErrorAzureError(err)

		// errors we expect might happen that we are ok with waiting for
		catch := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
		}

		if helpers.ContainsString(catch, azerr.Type) {
			// reconciliation is not done but error is acceptable
			return false, nil
		}

		catchUnrecoverableErrors := []string{
			errhelp.ProvisioningDisabled,
			errhelp.LocationNotAvailableForResourceType,
		}

		if helpers.ContainsString(catchUnrecoverableErrors, azerr.Type) {
			// Unrecoverable error, so stop reconcilation
			instance.Status.Message = "Reconcilation hit unrecoverable error: " + err.Error()
			return true, nil
		}
		// reconciliation not done and we don't know what happened
		return false, err
	}

	instance.Status.Provisioning = true
	instance.Status.Message = "Server request submitted to Azure"

	return false, nil
}

func (m *MySQLServerClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	if options.SecretClient != nil {
		m.SecretClient = options.SecretClient
	}

	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	status, err := m.DeleteServer(ctx, instance.Spec.ResourceGroup, instance.Name)
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
			m.SecretClient.Delete(ctx, key)
			return false, nil
		}
	}

	return true, nil
}

func (m *MySQLServerClient) GetParents(obj runtime.Object) ([]resourcemanager.KubeParent, error) {

	instance, err := m.convert(obj)
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

func (m *MySQLServerClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

func (m *MySQLServerClient) convert(obj runtime.Object) (*v1alpha1.MySQLServer, error) {
	local, ok := obj.(*v1alpha1.MySQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

func (m *MySQLServerClient) CreateServerIfValid(ctx context.Context, servername string, resourcegroup string, location string, tags map[string]*string, serverversion mysql.ServerVersion, sslenforcement mysql.SslEnforcementEnum, skuInfo mysql.Sku, adminlogin string, adminpassword string) (future mysql.ServersCreateFuture, err error) {

	client := getMySQLServersClient()

	// Check if name is valid if this is the first create call
	valid, err := m.CheckServerNameAvailability(ctx, servername)
	if valid == false {
		return future, err
	}

	return client.Create(
		ctx,
		resourcegroup,
		servername,
		mysql.ServerForCreate{
			Location: &location,
			Tags:     tags,
			Properties: &mysql.ServerPropertiesForDefaultCreate{
				AdministratorLogin:         &adminlogin,
				AdministratorLoginPassword: &adminpassword,
				Version:                    serverversion,
				SslEnforcement:             sslenforcement,
				//StorageProfile: &mysql.StorageProfile{},
				CreateMode: mysql.CreateModeServerPropertiesForCreate,
			},
			Sku: &skuInfo,
		},
	)
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

func (m *MySQLServerClient) AddServerCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha1.MySQLServer) error {
	key := types.NamespacedName{
		Name:      secretName,
		Namespace: instance.Namespace,
	}

	err := m.SecretClient.Upsert(ctx,
		key,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(m.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

func (m *MySQLServerClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.MySQLServer) (map[string][]byte, error) {
	name := instance.Name

	usernameLength := 8
	passwordLength := 16

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := m.SecretClient.Get(ctx, key); err == nil {
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
	secret["fullyQualifiedServerName"] = []byte(name + ".mysql.database.azure.com")

	return secret, nil
}
