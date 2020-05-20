// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"
	"strings"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	azurev1alpha2 "github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"
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

	createmode := "Default"
	if len(instance.Spec.CreateMode) != 0 {
		createmode = instance.Spec.CreateMode
	}

	// If a replica is requested, ensure that source server is specified
	if strings.EqualFold(createmode, "replica") {
		if len(instance.Spec.ReplicaProperties.SourceServerId) == 0 {
			instance.Status.Message = "Replica requested but source server unspecified"
			return true, nil
		}
	}

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := m.GetOrPrepareSecret(ctx, instance)
	if err != nil {
		return false, err
	}

	err = m.AddServerCredsToSecrets(ctx, instance.Name, secret, instance)
	if err != nil {
		return false, err
	}

	// convert kube labels to expected tag format
	labels := helpers.LabelsToTags(instance.GetLabels())

	// Check if this server already exists and its state if it does. This is required
	// to overcome the issue with the lack of idempotence of the Create call
	hash := ""
	server, err := m.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
	if err != nil {
		// handle failures in the async operation
		if instance.Status.PollingURL != "" {
			pClient := pollclient.NewPollClient()
			res, err := pClient.Get(ctx, instance.Status.PollingURL)
			if err != nil {
				instance.Status.Provisioning = false
				return false, err
			}

			if res.Status == "Failed" {
				instance.Status.Provisioning = false
				instance.Status.RequestedAt = nil
				ignore := []string{
					errhelp.SubscriptionDoesNotHaveServer,
					errhelp.ServiceBusy,
				}
				if !helpers.ContainsString(ignore, res.Error.Code) {
					instance.Status.Message = res.Error.Error()
					return true, nil
				}
			}
		}
	} else {
		instance.Status.State = string(server.UserVisibleState)

		hash = helpers.Hash256(instance.Spec)
		if instance.Status.SpecHash == hash && (instance.Status.Provisioned || instance.Status.FailedProvisioning) {
			instance.Status.RequestedAt = nil
			return true, nil
		}
		if server.UserVisibleState == mysql.ServerStateReady {
			// Update secret with FQ name of the server. We ignore the error.
			m.UpdateServerNameInSecret(ctx, instance.Name, secret, *server.FullyQualifiedDomainName, instance)

			instance.Status.Provisioned = true
			instance.Status.Provisioning = false
			instance.Status.Message = resourcemanager.SuccessMsg
			instance.Status.ResourceId = *server.ID
			instance.Status.State = string(server.UserVisibleState)
			instance.Status.SpecHash = hash
			return true, nil
		}
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

	pollURL, server, err := m.CreateServerIfValid(
		ctx,
		*instance,
		labels,
		skuInfo,
		adminlogin,
		adminpassword,
		createmode,
	)
	if err != nil {
		// let the user know what happened
		instance.Status.Message = errhelp.StripErrorIDs(err)
		instance.Status.Provisioning = false
		azerr := errhelp.NewAzureErrorAzureError(err)

		catchRequeue := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.AsyncOpIncompleteError,
			errhelp.SubscriptionDoesNotHaveServer,
			errhelp.ServiceBusy,
		}
		catchUnrecoverable := []string{
			errhelp.ProvisioningDisabled,
			errhelp.LocationNotAvailableForResourceType,
			errhelp.InvalidRequestContent,
			errhelp.InternalServerError,
		}

		// handle the errors
		if helpers.ContainsString(catchRequeue, azerr.Type) {
			if azerr.Type == errhelp.AsyncOpIncompleteError {
				instance.Status.Provisioning = true
				instance.Status.PollingURL = pollURL
			}
			return false, nil
		}

		if helpers.ContainsString(catchUnrecoverable, azerr.Type) {
			// Unrecoverable error, so stop reconcilation
			instance.Status.Message = "Reconcilation hit unrecoverable error: " + errhelp.StripErrorIDs(err)
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
		catch := []string{
			errhelp.AsyncOpIncompleteError,
		}
		gone := []string{
			errhelp.ResourceGroupNotFoundErrorCode,
			errhelp.ParentNotFoundErrorCode,
			errhelp.NotFoundErrorCode,
			errhelp.ResourceNotFound,
		}
		azerr := errhelp.NewAzureErrorAzureError(err)
		if helpers.ContainsString(catch, azerr.Type) {
			return true, nil
		} else if helpers.ContainsString(gone, azerr.Type) {
			return false, nil
		}
		return true, err
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
	st := v1alpha1.ASOStatus(instance.Status)
	return &st, nil
}

// convert concerts a runtime.Object to a MySQLServer object
func (m *MySQLServerClient) convert(obj runtime.Object) (*v1alpha2.MySQLServer, error) {
	local, ok := obj.(*v1alpha2.MySQLServer)
	if !ok {
		return nil, fmt.Errorf("failed type assertion on kind: %s", obj.GetObjectKind().GroupVersionKind().String())
	}
	return local, nil
}

// AddServerCredsToSecrets saves the server's admin credentials in the secret store
func (m *MySQLServerClient) AddServerCredsToSecrets(ctx context.Context, secretName string, data map[string][]byte, instance *azurev1alpha2.MySQLServer) error {
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

// UpdateSecretWithFullServerName updates the secret with the fully qualified server name
func (m *MySQLServerClient) UpdateServerNameInSecret(ctx context.Context, secretName string, data map[string][]byte, fullservername string, instance *azurev1alpha2.MySQLServer) error {
	key := types.NamespacedName{
		Name:      secretName,
		Namespace: instance.Namespace,
	}

	data["fullyQualifiedServerName"] = []byte(fullservername)

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
func (m *MySQLServerClient) GetOrPrepareSecret(ctx context.Context, instance *azurev1alpha2.MySQLServer) (map[string][]byte, error) {
	name := instance.Name
	createmode := instance.Spec.CreateMode

	// If createmode == default, then this is a new server creation, so generate username/password
	// If createmode == replica, then get the credentials from the source server secret and use that

	secret := map[string][]byte{}
	var key types.NamespacedName
	var Username string
	var Password string

	if strings.EqualFold(createmode, "default") { // new Mysql server creation
		// See if secret already exists and return if it does
		key = types.NamespacedName{Name: name, Namespace: instance.Namespace}
		if stored, err := m.SecretClient.Get(ctx, key); err == nil {
			return stored, nil
		}
		// Generate random username password if secret does not exist already
		Username = helpers.GenerateRandomUsername(10)
		Password = helpers.NewPassword()
	} else { // replica
		sourceServerId := instance.Spec.ReplicaProperties.SourceServerId
		if len(sourceServerId) != 0 {
			// Parse to get source server name
			sourceServerIdSplit := strings.Split(sourceServerId, "/")
			sourceserver := sourceServerIdSplit[len(sourceServerIdSplit)-1]

			// Get the username and password from the source server's secret
			key = types.NamespacedName{Name: sourceserver, Namespace: instance.Namespace}
			if sourcesecret, err := m.SecretClient.Get(ctx, key); err == nil {
				Username = string(sourcesecret["username"])
				Password = string(sourcesecret["password"])
			}
		}
	}

	// Populate secret fields
	secret["username"] = []byte(Username)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", Username, name))
	secret["password"] = []byte(Password)
	secret["mySqlServerName"] = []byte(name)
	return secret, nil
}
