/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"
	"net/http"
	"reflect"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

const (
	// #nosec
	namespacedSecretName = "aso-credential"
)

type armClient struct {
	genericClient  *genericarmclient.GenericClient
	secretData     map[string][]byte
	credentialFrom types.NamespacedName
}

func (c *armClient) GenericClient() *genericarmclient.GenericClient {
	return c.genericClient
}

func (c *armClient) CredentialFrom() string {
	return c.credentialFrom.String()
}

type armClientCache struct {
	lock sync.Mutex
	// clients map hold namespaced name for a secret to all the registered ARM clients
	clients      map[string]*armClient
	globalClient *armClient
	kubeClient   kubeclient.Client
	cloudConfig  cloud.Configuration
	httpClient   *http.Client
}

func NewARMClientCache(
	client *genericarmclient.GenericClient,
	podNamespace string,
	kubeClient kubeclient.Client,
	configuration cloud.Configuration,
	httpClient *http.Client) *armClientCache {

	globalClient := &armClient{
		genericClient:  client,
		credentialFrom: types.NamespacedName{Name: "aso-controller-settings", Namespace: podNamespace},
	}

	return &armClientCache{
		lock:         sync.Mutex{},
		clients:      make(map[string]*armClient),
		globalClient: globalClient,
		kubeClient:   kubeClient,
		cloudConfig:  configuration,
		httpClient:   httpClient,
	}
}

func newARMClient(client *genericarmclient.GenericClient, secretData map[string][]byte, credentialFrom types.NamespacedName) *armClient {
	return &armClient{
		genericClient:  client,
		secretData:     secretData,
		credentialFrom: credentialFrom,
	}
}

func (c *armClientCache) Register(client *armClient) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.clients[client.CredentialFrom()] = client
}

func (c *armClientCache) Lookup(key string) (*armClient, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	client, ok := c.clients[key]
	return client, ok
}

// GetClient function is a helper method to find and return client and credential used for the given resource
func (c *armClientCache) GetClient(ctx context.Context, obj genruntime.ARMMetaObject) (*genericarmclient.GenericClient, string, error) {
	// Namespaced secret
	secret, err := c.getSecret(ctx, obj.GetNamespace(), namespacedSecretName)
	if err != nil {
		return nil, "", err
	}

	if secret != nil {
		nsName := types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}
		client, ok := c.Lookup(nsName.String())
		// if client does not already exist, or secret data is changed, we need to get new token.
		if !ok || !matchSecretData(client.secretData, secret.Data) {
			credential, subscriptionID, err := c.newCredentialFromSecret(secret, nsName)
			if err != nil {
				return nil, "", err
			}

			newClient, err := genericarmclient.NewGenericClientFromHTTPClient(c.cloudConfig, credential, c.httpClient, subscriptionID, metrics.NewARMClientMetrics())
			if err != nil {
				return nil, "", err
			}

			armClient := newARMClient(newClient, secret.Data, nsName)
			c.Register(armClient)
			return armClient.GenericClient(), armClient.CredentialFrom(), nil
		} else {
			return client.GenericClient(), client.CredentialFrom(), nil
		}
	}
	// If not found, default would be global
	return c.globalClient.GenericClient(), c.globalClient.CredentialFrom(), nil
}

func (c *armClientCache) newCredentialFromSecret(secret *v1.Secret, nsName types.NamespacedName) (azcore.TokenCredential, string, error) {
	subscriptionID, ok := secret.Data[config.SubscriptionIDVar]
	if !ok {
		return nil, "", core.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.SubscriptionIDVar))
	}
	tenantID, ok := secret.Data[config.TenantIDVar]
	if !ok {
		return nil, "", core.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.TenantIDVar))
	}
	clientID, ok := secret.Data[config.AzureClientID]
	if !ok {
		return nil, "", core.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.AzureClientID))
	}
	clientSecret, ok := secret.Data[config.AzureClientSecret]
	// TODO: We check this for now until we support Workload Identity. When we support workload identity for multitenancy, !ok would mean to check for workload identity.
	if !ok {
		return nil, "", core.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.AzureClientSecret))
	}

	credential, err := azidentity.NewClientSecretCredential(string(tenantID), string(clientID), string(clientSecret), nil)
	if err != nil {
		return nil, "", errors.Wrap(err, errors.Errorf("Invalid Client Secret Credential for %q encountered", nsName.String()).Error())
	}
	return credential, string(subscriptionID), nil
}

func (c *armClientCache) getSecret(ctx context.Context, namespace string, secretName string) (*v1.Secret, error) {
	secret := &v1.Secret{}

	err := c.kubeClient.Get(
		ctx,
		types.NamespacedName{Namespace: namespace, Name: secretName},
		secret)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return secret, nil
}

func matchSecretData(old, new map[string][]byte) bool {
	return reflect.DeepEqual(old, new)
}
