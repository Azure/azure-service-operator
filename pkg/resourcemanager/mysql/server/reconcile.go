// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"
	"strings"

	mysql "github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/go-autorest/autorest/to"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/api/v1alpha1"
	azurev1alpha1 "github.com/Azure/azure-service-operator/api/v1alpha1"
	"github.com/Azure/azure-service-operator/api/v1alpha2"
	azurev1alpha2 "github.com/Azure/azure-service-operator/api/v1alpha2"
	"github.com/Azure/azure-service-operator/pkg/errhelp"
	"github.com/Azure/azure-service-operator/pkg/helpers"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager"
	"github.com/Azure/azure-service-operator/pkg/resourcemanager/pollclient"
	"github.com/Azure/azure-service-operator/pkg/secrets"
)

// Ensure idempotently instantiates the requested server (if possible) in Azure
func (m *MySQLServerClient) Ensure(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {
	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := m.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
	}

	instance, err := m.convert(obj)
	if err != nil {
		return true, err
	}

	createmode := mysql.CreateModeDefault
	if len(instance.Spec.CreateMode) != 0 {
		createmode = mysql.CreateMode(instance.Spec.CreateMode)
	}

	// If a replica is requested, ensure that source server is specified
	if createmode == mysql.CreateModeReplica {
		if len(instance.Spec.ReplicaProperties.SourceServerId) == 0 {
			instance.Status.Message = "Replica requested but source server unspecified"
			return true, nil
		}
	}

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := m.GetOrPrepareSecret(ctx, secretClient, instance)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Failed to get or prepare secret: %s", err.Error())
		return false, err
	}

	err = m.AddServerCredsToSecrets(ctx, secretClient, secret, instance)
	if err != nil {
		return false, err
	}

	// convert kube labels to expected tag format
	labels := helpers.LabelsToTags(instance.GetLabels())

	hash := helpers.Hash256(instance.Spec)
	if instance.Status.SpecHash == hash && (instance.Status.Provisioned || instance.Status.FailedProvisioning) {
		instance.Status.RequestedAt = nil
		return true, nil
	} else if instance.Status.SpecHash != hash && !instance.Status.Provisioning {
		instance.Status.Provisioned = false
	}

	if instance.Status.Provisioning {

		// Check if this server already exists and its state if it does. This is required
		// to overcome the issue with the lack of idempotence of the Create call
		server, err := m.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
		if err != nil {
			// handle failures in the async operation
			if instance.Status.PollingURL != "" {
				pClient := pollclient.NewPollClient(m.Creds)
				res, err := pClient.Get(ctx, instance.Status.PollingURL)
				if err != nil {
					instance.Status.Provisioning = false
					return false, err
				}

				if res.Status == pollclient.LongRunningOperationPollStatusFailed {
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
				m.UpdateServerNameInSecret(ctx, secretClient, secret, *server.FullyQualifiedDomainName, instance)

				instance.Status.Provisioned = true
				instance.Status.Provisioning = false
				instance.Status.Message = resourcemanager.SuccessMsg
				instance.Status.ResourceId = *server.ID
				instance.Status.State = string(server.UserVisibleState)
				instance.Status.SpecHash = hash
				instance.Status.PollingURL = ""
				return true, nil
			}
		}

	}

	if !instance.Status.Provisioning && !instance.Status.Provisioned {
		instance.Status.Provisioning = true
		instance.Status.FailedProvisioning = false

		adminlogin := string(secret["username"])
		adminpassword := string(secret["password"])
		skuInfo := mysql.Sku{
			Name:     to.StringPtr(instance.Spec.Sku.Name),
			Tier:     mysql.SkuTier(instance.Spec.Sku.Tier),
			Capacity: to.Int32Ptr(instance.Spec.Sku.Capacity),
			Size:     to.StringPtr(instance.Spec.Sku.Size),
			Family:   to.StringPtr(instance.Spec.Sku.Family),
		}

		pollURL, _, err := m.CreateServerIfValid(
			ctx,
			*instance,
			labels,
			skuInfo,
			adminlogin,
			adminpassword,
			createmode,
			hash,
		)
		if err != nil {
			instance.Status.Message = errhelp.StripErrorIDs(err)
			instance.Status.Provisioning = false

			azerr := errhelp.NewAzureError(err)

			catchInProgress := []string{
				errhelp.AsyncOpIncompleteError,
				errhelp.AlreadyExists,
			}
			catchKnownError := []string{
				errhelp.ResourceGroupNotFoundErrorCode,
				errhelp.ParentNotFoundErrorCode,
				errhelp.NotFoundErrorCode,
				errhelp.ServiceBusy,
				errhelp.InternalServerError,
			}

			// handle the errors
			if helpers.ContainsString(catchInProgress, azerr.Type) {
				if azerr.Type == errhelp.AsyncOpIncompleteError {
					instance.Status.PollingURL = pollURL
				}
				instance.Status.Message = "MySQL server exists but may not be ready"
				instance.Status.Provisioning = true
				return false, nil
			}

			if helpers.ContainsString(catchKnownError, azerr.Type) {
				return false, nil
			}

			// serious error occurred, end reconciliation and mark it as failed
			instance.Status.Message = errhelp.StripErrorIDs(err)
			instance.Status.Provisioned = false
			instance.Status.FailedProvisioning = true
			return true, nil
		}

		instance.Status.PollingURL = pollURL
		instance.Status.Message = "request submitted to Azure"
	}
	return false, nil
}

// Delete idempotently ensures the server is gone from Azure
func (m *MySQLServerClient) Delete(ctx context.Context, obj runtime.Object, opts ...resourcemanager.ConfigOption) (bool, error) {

	options := &resourcemanager.Options{}
	for _, opt := range opts {
		opt(options)
	}

	secretClient := m.SecretClient
	if options.SecretClient != nil {
		secretClient = options.SecretClient
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
		azerr := errhelp.NewAzureError(err)
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
			secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
			secretClient.Delete(ctx, secretKey)
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
func (m *MySQLServerClient) AddServerCredsToSecrets(ctx context.Context, secretClient secrets.SecretClient, data map[string][]byte, instance *azurev1alpha2.MySQLServer) error {
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	err := secretClient.Upsert(ctx,
		secretKey,
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
func (m *MySQLServerClient) UpdateServerNameInSecret(ctx context.Context, secretClient secrets.SecretClient, data map[string][]byte, fullservername string, instance *azurev1alpha2.MySQLServer) error {
	secretKey := secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}

	data["fullyQualifiedServerName"] = []byte(fullservername)

	err := secretClient.Upsert(ctx,
		secretKey,
		data,
		secrets.WithOwner(instance),
		secrets.WithScheme(m.Scheme),
	)
	if err != nil {
		return err
	}

	return nil
}

// GetOrPrepareSecret gets the admin credentials if they are stored or generates some if not
func (m *MySQLServerClient) GetOrPrepareSecret(ctx context.Context, secretClient secrets.SecretClient, instance *azurev1alpha2.MySQLServer) (map[string][]byte, error) {
	createmode := instance.Spec.CreateMode

	// If createmode == default, then this is a new server creation, so generate username/password
	// If createmode == replica, then get the credentials from the source server secret and use that

	secret := map[string][]byte{}
	var key secrets.SecretKey
	var username string
	var password string

	if strings.EqualFold(createmode, "default") { // new Mysql server creation
		// See if secret already exists and return if it does
		key = secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
		if stored, err := secretClient.Get(ctx, key); err == nil {
			return stored, nil
		}
		// Generate random username password if secret does not exist already
		username = helpers.GenerateRandomUsername(10)
		password = helpers.NewPassword()
	} else { // replica
		sourceServerId := instance.Spec.ReplicaProperties.SourceServerId
		if len(sourceServerId) != 0 {
			// Parse to get source server name
			sourceServerIdSplit := strings.Split(sourceServerId, "/")
			sourceServer := sourceServerIdSplit[len(sourceServerIdSplit)-1]

			// Get the username and password from the source server's secret
			key = secrets.SecretKey{Name: sourceServer, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
			if sourceSecret, err := secretClient.Get(ctx, key); err == nil {
				username = string(sourceSecret["username"])
				password = string(sourceSecret["password"])
			}
		}
	}

	// Populate secret fields
	secret["username"] = []byte(username)
	secret["fullyQualifiedUsername"] = []byte(fmt.Sprintf("%s@%s", username, instance.Name))
	secret["password"] = []byte(password)
	secret["mySqlServerName"] = []byte(instance.Name)
	return secret, nil
}
