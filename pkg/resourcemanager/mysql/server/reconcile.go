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
	"github.com/Azure/azure-service-operator/pkg/secrets"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// Ensure idempotently instantiates the requested server (ig possible) in Azure
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
			instance.Status.ResourceId = *server.ID
			return true, nil
		}
		return false, nil
	}

	// if the create has been sent with no error we need to wait before calling it again
	// @todo set an appropriate time since create has been called to retry
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

		switch azerr.Type {
		case errhelp.ResourceGroupNotFoundErrorCode, errhelp.ParentNotFoundErrorCode:
			// errors we expect might happen that we are ok with waiting for
			return false, nil
		case errhelp.ProvisioningDisabled, errhelp.LocationNotAvailableForResourceType:
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

// Delete idempotently ensures the server is gone from Azure
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

// GetParents returns all the potential kube parents of the server
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

// GetStatus returns a pointer to the server resources' status sub-resource
func (m *MySQLServerClient) GetStatus(obj runtime.Object) (*v1alpha1.ASOStatus, error) {
	instance, err := m.convert(obj)
	if err != nil {
		return nil, err
	}
	return &instance.Status, nil
}

// convert concerts a runtime.Object to a MySQLServer object
func (m *MySQLServerClient) convert(obj runtime.Object) (*v1alpha1.MySQLServer, error) {
	local, ok := obj.(*v1alpha1.MySQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

// AddServerCredsToSecrets saves the server's admin credentials in the secret store
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

// GetOrPrepareSecret gets tje admin credentials if they are stored or generates some if not
func (m *MySQLServerClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha1.MySQLServer) (map[string][]byte, error) {
	name := instance.Name

	secret := map[string][]byte{}

	key := types.NamespacedName{Name: name, Namespace: instance.Namespace}
	if stored, err := m.SecretClient.Get(ctx, key); err == nil {
		return stored, nil
	}

	randomUsername := helpers.GenerateRandomUsername(10)
	randomPassword := helpers.NewPassword()

	secret["username"] = []byte(randomUsername)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", randomUsername, name))
	secret["password"] = []byte(randomPassword)
	secret["postgreSqlServerName"] = []byte(name)
	// TODO: The below may not be right for non Azure public cloud.
	secret["fullyQualifiedServerName"] = []byte(name + ".mysql.database.azure.com")

	return secret, nil
}
