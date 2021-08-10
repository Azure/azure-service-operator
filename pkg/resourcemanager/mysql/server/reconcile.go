// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

package server

import (
	"context"
	"fmt"
	"strings"

	"github.com/Azure/azure-sdk-for-go/services/mysql/mgmt/2017-12-01/mysql"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/Azure/azure-service-operator/api"
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

const (
	UsernameSecretKey                 = "username"
	PasswordSecretKey                 = "password"
	FullyQualifiedServerNameSecretKey = "fullyQualifiedServerName"
	MySQLServerNameSecretKey          = "mySqlServerName"
	FullyQualifiedUsernameSecretKey   = "fullyQualifiedUsername"
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

	pClient := pollclient.NewPollClient(m.Creds)
	lroPollResult, err := pClient.PollLongRunningOperationIfNeededV1Alpha2(ctx, &instance.Status, api.PollingURLKindCreateOrUpdate)
	if err != nil {
		instance.Status.Message = err.Error()
		return false, err
	}

	if lroPollResult == pollclient.PollResultTryAgainLater {
		// Need to wait a bit before trying again
		return false, nil
	}
	if lroPollResult == pollclient.PollResultBadRequest {
		// Reached a terminal state
		instance.Status.SetFailedProvisioning(instance.Status.Message)
		return true, nil
	}

	createMode := mysql.CreateModeDefault
	if len(instance.Spec.CreateMode) != 0 {
		createMode = mysql.CreateMode(instance.Spec.CreateMode)
	}

	// If a replica is requested, ensure that source server is specified
	if createMode == mysql.CreateModeReplica {
		if len(instance.Spec.ReplicaProperties.SourceServerId) == 0 {
			instance.Status.SetFailedProvisioning("Replica requested but source server unspecified")
			return true, nil
		}
	}

	adminCreds, err := m.GetUserProvidedAdminCredentials(ctx, instance)
	if err != nil {
		// The error already has the details we need
		instance.Status.Message = err.Error()
		return false, err
	}

	// Check to see if secret exists and if yes retrieve the admin login and password
	secret, err := m.GetOrPrepareSecret(ctx, secretClient, instance, adminCreds)
	if err != nil {
		instance.Status.Message = fmt.Sprintf("Failed to get or prepare secret: %s", err.Error())
		return false, err
	}

	// If the user didn't provide administrator credentials, get them from the secret
	if adminCreds == nil {
		adminCreds = &MySQLCredentials{
			username: string(secret[UsernameSecretKey]),
			password: string(secret[PasswordSecretKey]),
		}
	}

	// convert kube labels to expected tag format
	labels := helpers.LabelsToTags(instance.GetLabels()) // TODO: Where is this used

	// Bail out early if we've already provisioned or failed provisioning and the spec hash hasn't changed
	hash := m.calculateHash(instance.Spec, secret)
	if instance.Status.SpecHash == hash && (instance.Status.Provisioned || instance.Status.FailedProvisioning) {
		instance.Status.RequestedAt = nil
		return true, nil
	}

	err = m.UpsertSecrets(ctx, secretClient, secret, instance)
	if err != nil {
		return false, err
	}

	if instance.Status.Provisioning {
		// Get the existing resource - the LRO is done so the expectation is that this exists
		server, err := m.GetServer(ctx, instance.Spec.ResourceGroup, instance.Name)
		if err == nil {
			instance.Status.State = string(server.UserVisibleState)
			if server.UserVisibleState == mysql.ServerStateReady {
				// Update secret with FQ name of the server. We ignore the error.
				m.UpdateServerNameInSecret(ctx, secretClient, secret, *server.FullyQualifiedDomainName, instance)

				instance.Status.SetProvisioned(resourcemanager.SuccessMsg)
				instance.Status.ResourceId = *server.ID
				instance.Status.State = string(server.UserVisibleState)
				instance.Status.SpecHash = hash
				instance.Status.PollingURL = ""
				return true, nil
			}
		}
	}

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
		adminCreds.username,
		adminCreds.password,
		createMode,
	)
	if err != nil {
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
			instance.Status.SetProvisioning("MySQL server exists but may not be ready")
			return false, nil
		}

		if helpers.ContainsString(catchKnownError, azerr.Type) {
			instance.Status.SetProvisioning(errhelp.StripErrorIDs(err))
			return false, nil
		}

		// serious error occurred, end reconciliation and mark it as failed
		instance.Status.SetFailedProvisioning(errhelp.StripErrorIDs(err))
		return true, nil
	}

	instance.Status.SetPollingURL(pollURL, api.PollingURLKindCreateOrUpdate)
	instance.Status.SetProvisioning("request submitted to Azure")

	return false, nil
}

func (m *MySQLServerClient) calculateHash(spec azurev1alpha2.MySQLServerSpec, secret map[string][]byte) string {
	return helpers.Hash256(struct {
		Spec   azurev1alpha2.MySQLServerSpec
		Secret map[string][]byte
	}{
		Spec:   spec,
		Secret: secret,
	})
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

// UpsertSecrets saves the server's admin credentials in the secret store
func (m *MySQLServerClient) UpsertSecrets(ctx context.Context, secretClient secrets.SecretClient, data map[string][]byte, instance *azurev1alpha2.MySQLServer) error {
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

	data[FullyQualifiedServerNameSecretKey] = []byte(fullservername)

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
func (m *MySQLServerClient) GetOrPrepareSecret(
	ctx context.Context,
	secretClient secrets.SecretClient,
	instance *azurev1alpha2.MySQLServer,
	adminCredentials *MySQLCredentials) (map[string][]byte, error) {

	secret := map[string][]byte{}
	var key secrets.SecretKey
	var username string
	var password string

	// If createmode == default, then this is a new server creation, so generate username/password
	// If createmode == replica, then get the credentials from the source server secret and use that
	if strings.EqualFold(instance.Spec.CreateMode, "default") { // new Mysql server creation
		// See if secret already exists and return if it does
		key = secrets.SecretKey{Name: instance.Name, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
		if stored, err := secretClient.Get(ctx, key); err == nil {
			// Ensure that we use the most up to date secret information if the user has provided it
			if adminCredentials != nil {
				stored[UsernameSecretKey] = []byte(adminCredentials.username)
				stored[PasswordSecretKey] = []byte(adminCredentials.password)
			}
			return stored, nil
		}
		// Generate random username and password if secret does not exist already and it
		// wasn't provided in the spec
		if adminCredentials == nil {
			username = helpers.GenerateRandomUsername(10)
			password = helpers.NewPassword()
		} else {
			username = adminCredentials.username
			password = adminCredentials.password
		}
	} else {
		// We only attempt to get the username and password from the primary if there's not
		// a user specified secret. If there IS a user specified secret then just use that.
		if adminCredentials == nil {
			sourceServerId := instance.Spec.ReplicaProperties.SourceServerId
			if len(sourceServerId) != 0 {
				// Parse to get source server name
				sourceServerIdSplit := strings.Split(sourceServerId, "/")
				sourceServer := sourceServerIdSplit[len(sourceServerIdSplit)-1]

				// Get the username and password from the source server's secret
				key = secrets.SecretKey{Name: sourceServer, Namespace: instance.Namespace, Kind: instance.TypeMeta.Kind}
				if sourceSecret, err := secretClient.Get(ctx, key); err == nil {
					username = string(sourceSecret[UsernameSecretKey])
					password = string(sourceSecret[PasswordSecretKey])
				}
			}
		} else {
			username = adminCredentials.username
			password = adminCredentials.password
		}
	}

	// Populate secret fields derived from username and password
	secret[UsernameSecretKey] = []byte(username)
	secret[PasswordSecretKey] = []byte(password)
	secret[FullyQualifiedUsernameSecretKey] = []byte(makeFullyQualifiedUsername(username, instance.Name))
	secret[MySQLServerNameSecretKey] = []byte(instance.Name)
	return secret, nil
}

// MySQLCredentials is a username/password pair for a MySQL account
type MySQLCredentials struct {
	username string
	password string
}

// GetUserProvidedAdminCredentials gets the user provided MySQLCredentials, or nil if none was
// specified by the user.
func (m *MySQLServerClient) GetUserProvidedAdminCredentials(
	ctx context.Context,
	instance *azurev1alpha2.MySQLServer) (*MySQLCredentials, error) {

	if instance.Spec.AdminSecret == "" {
		return nil, nil
	}

	key := client.ObjectKey{Namespace: instance.Namespace, Name: instance.Spec.AdminSecret}
	secret := &v1.Secret{}
	if err := m.KubeReader.Get(ctx, key, secret); err != nil {
		return nil, errors.Wrapf(err, "Failed to get AdminSecret %q", key)
	}

	adminUsernameBytes, ok := secret.Data[UsernameSecretKey]
	if !ok {
		return nil, errors.Errorf("AdminSecret %s is missing required %q field", instance.Spec.AdminSecret, UsernameSecretKey)
	}
	adminPasswordBytes, ok := secret.Data[PasswordSecretKey]
	if !ok {
		return nil, errors.Errorf("AdminSecret %s is missing required %q field", instance.Spec.AdminSecret, PasswordSecretKey)
	}

	return &MySQLCredentials{
		username: string(adminUsernameBytes),
		password: string(adminPasswordBytes),
	}, nil
}

func makeFullyQualifiedUsername(username string, instanceName string) string {
	return fmt.Sprintf("%s@%s", username, instanceName)
}
