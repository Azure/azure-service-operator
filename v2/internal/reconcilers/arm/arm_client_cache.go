/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"
	"net/http"
	"reflect"
	"strings"
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
	globalCredentialSecretName = "aso-controller-settings"
	// #nosec
	namespacedSecretName        = "aso-credential"
	perResourceSecretAnnotation = "serviceoperator.azure.com/credential-from"
	namespacedNameSeparator     = "/"
)

// armClientCache is a cache for armClients to hold multiple credential clients and global credential client.
type armClientCache struct {
	lock sync.Mutex
	// clients allows quick lookup of an armClient for each namespace
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
		credentialFrom: types.NamespacedName{Name: globalCredentialSecretName, Namespace: podNamespace},
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

func (c *armClientCache) register(client *armClient) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.clients[client.CredentialFrom()] = client
}

func (c *armClientCache) lookup(key string) (*armClient, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	client, ok := c.clients[key]
	return client, ok
}

// GetClient finds and returns a client and credential to be used for a given resource
func (c *armClientCache) GetClient(ctx context.Context, obj genruntime.ARMMetaObject) (*genericarmclient.GenericClient, string, error) {

	client, err := c.getPerResourceCredential(ctx, obj)
	if err != nil {
		return nil, "", err
	} else if client != nil {
		return client.GenericClient(), client.CredentialFrom(), nil
	}
	// Namespaced secret
	client, err = c.getNamespacedCredential(ctx, obj)
	if err != nil {
		return nil, "", err
	} else if client != nil {
		return client.GenericClient(), client.CredentialFrom(), nil
	}

	// If not found, default is the global client
	return c.globalClient.GenericClient(), c.globalClient.CredentialFrom(), nil
}

func (c *armClientCache) getPerResourceCredential(ctx context.Context, obj genruntime.ARMMetaObject) (*armClient, error) {
	return c.getCredentialFromAnnotation(ctx, obj, perResourceSecretAnnotation)
}

func (c *armClientCache) getNamespacedCredential(ctx context.Context, obj genruntime.ARMMetaObject) (*armClient, error) {
	secret, err := c.getSecret(ctx, obj.GetNamespace(), namespacedSecretName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil // Not finding this secret is allowed, allow caller to proceed to higher scope secret
		}
		return nil, err
	}

	armClient, err := c.getARMClientFromSecret(secret)
	if err != nil {
		return nil, err
	}

	return armClient, nil
}

func (c *armClientCache) getCredentialFromAnnotation(ctx context.Context, obj genruntime.ARMMetaObject, annotation string) (*armClient, error) {
	credentialFrom, ok := obj.GetAnnotations()[annotation]
	if !ok {
		return nil, nil
	}

	// annotation exists, use specified secret
	secretNamespacedName := getSecretNameFromAnnotation(credentialFrom, obj.GetNamespace())

	secret, err := c.getSecret(ctx, secretNamespacedName.Namespace, secretNamespacedName.Name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, core.NewSecretNotFoundError(secretNamespacedName, errors.Errorf("credential Secret %q not found for '%s/%s'", secretNamespacedName.String(), obj.GetNamespace(), obj.GetName()))
		}
		return nil, err
	}

	armClient, err := c.getARMClientFromSecret(secret)
	if err != nil {
		return nil, err
	}
	return armClient, nil
}

func getSecretNameFromAnnotation(credentialFrom string, resourceNamespace string) types.NamespacedName {
	if strings.Contains(credentialFrom, namespacedNameSeparator) {
		slice := strings.Split(credentialFrom, namespacedNameSeparator)
		return types.NamespacedName{Namespace: slice[0], Name: slice[1]}
	} else {
		// If namespace is not specified, look into the resource namespace
		return types.NamespacedName{Namespace: resourceNamespace, Name: credentialFrom}
	}
}

func (c *armClientCache) getARMClientFromSecret(secret *v1.Secret) (*armClient, error) {
	nsName := types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}
	client, ok := c.lookup(nsName.String())

	if ok && matchSecretData(client.secretData, secret.Data) {
		return client, nil
	}

	// if client does not already exist, or secret data is changed, we need to get new token.
	credential, subscriptionID, err := c.newCredentialFromSecret(secret, nsName)
	if err != nil {
		return nil, err
	}

	newClient, err := genericarmclient.NewGenericClientFromHTTPClient(c.cloudConfig, credential, c.httpClient, subscriptionID, metrics.NewARMClientMetrics())
	if err != nil {
		return nil, err
	}

	armClient := newARMClient(newClient, secret.Data, nsName)
	c.register(armClient)
	return armClient, nil
}

func (c *armClientCache) newCredentialFromSecret(secret *v1.Secret, nsName types.NamespacedName) (azcore.TokenCredential, string, error) {
	subscriptionID, ok := secret.Data[config.SubscriptionIDVar]
	if !ok {
		err := core.NewSecretNotFoundError(nsName, errors.Errorf("credential Secret %q does not contain key %q", nsName.String(), config.SubscriptionIDVar))
		return nil, "", err
	}
	tenantID, ok := secret.Data[config.TenantIDVar]
	if !ok {
		err := core.NewSecretNotFoundError(nsName, errors.Errorf("credential Secret %q does not contain key %q", nsName.String(), config.TenantIDVar))
		return nil, "", err
	}
	clientID, ok := secret.Data[config.AzureClientIDVar]
	if !ok {
		err := core.NewSecretNotFoundError(nsName, errors.Errorf("credential Secret %q does not contain key %q", nsName.String(), config.AzureClientIDVar))
		return nil, "", err
	}
	clientSecret, ok := secret.Data[config.AzureClientSecretVar]
	// TODO: We check this for now until we support Workload Identity. When we support workload identity for multitenancy, !ok would mean to check for workload identity.
	if !ok {
		err := core.NewSecretNotFoundError(nsName, errors.Errorf("credential Secret %q does not contain key %q", nsName.String(), config.AzureClientSecretVar))
		return nil, "", err
	}

	credential, err := azidentity.NewClientSecretCredential(string(tenantID), string(clientID), string(clientSecret), nil)
	if err != nil {
		return nil, "", errors.Wrap(err, errors.Errorf("invalid Client Secret Credential for %q encountered", nsName.String()).Error())
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
		return nil, err
	}

	return secret, nil
}

func matchSecretData(old, new map[string][]byte) bool {
	return reflect.DeepEqual(old, new)
}
