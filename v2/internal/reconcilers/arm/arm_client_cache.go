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
	kerrors "k8s.io/apimachinery/pkg/util/errors"

	"github.com/Azure/azure-service-operator/v2/identity"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/reconcilers"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/core"
)

const (
	// #nosec
	globalCredentialSecretName = "aso-controller-settings"
	// #nosec
	NamespacedSecretName    = "aso-credential"
	namespacedNameSeparator = "/"
)

// ARMClientCache is a cache for armClients to hold multiple credential clients and global credential client.
type ARMClientCache struct {
	lock sync.Mutex
	// clients allows quick lookup of an armClient for each namespace
	clients      map[string]*armClient
	cloudConfig  cloud.Configuration
	globalClient *armClient
	kubeClient   kubeclient.Client
	httpClient   *http.Client
	armMetrics   *metrics.ARMClientMetrics
}

func NewARMClientCache(
	defaultClient *genericarmclient.GenericClient,
	podNamespace string,
	kubeClient kubeclient.Client,
	configuration cloud.Configuration,
	httpClient *http.Client,
	armMetrics *metrics.ARMClientMetrics) *ARMClientCache {

	globalClient := &armClient{
		genericClient:  defaultClient,
		credentialFrom: types.NamespacedName{Name: globalCredentialSecretName, Namespace: podNamespace},
	}

	return &ARMClientCache{
		lock:         sync.Mutex{},
		clients:      make(map[string]*armClient),
		cloudConfig:  configuration,
		kubeClient:   kubeClient,
		globalClient: globalClient,
		httpClient:   httpClient,
		armMetrics:   armMetrics,
	}
}

func (c *ARMClientCache) SetKubeClient(client kubeclient.Client) {
	c.kubeClient = client
}

func (c *ARMClientCache) register(client *armClient) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.clients[client.CredentialFrom()] = client
}

func (c *ARMClientCache) lookup(key string) (*armClient, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	client, ok := c.clients[key]
	return client, ok
}

// GetClient finds and returns a client and credential to be used for a given resource
func (c *ARMClientCache) GetClient(ctx context.Context, obj genruntime.ARMMetaObject) (*genericarmclient.GenericClient, string, error) {

	client, err := c.getPerResourceCredential(ctx, obj)
	if err != nil {
		return nil, "", err
	} else if client != nil {
		return client.GenericClient(), client.CredentialFrom(), nil
	}

	// Namespaced secret
	client, err = c.getNamespacedCredential(ctx, obj.GetNamespace())
	if err != nil {
		return nil, "", err
	} else if client != nil {
		return client.GenericClient(), client.CredentialFrom(), nil
	}

	if c.globalClient.GenericClient() == nil {
		return nil, "", errors.New("Default global credential is not configured. Use either namespaced or per-resource credential")
	}
	// If not found, default is the global client
	return c.globalClient.GenericClient(), c.globalClient.CredentialFrom(), nil
}

func (c *ARMClientCache) getPerResourceCredential(ctx context.Context, obj genruntime.ARMMetaObject) (*armClient, error) {
	return c.getCredentialFromAnnotation(ctx, obj, reconcilers.PerResourceSecretAnnotation)
}

func (c *ARMClientCache) getNamespacedCredential(ctx context.Context, namespace string) (*armClient, error) {
	secret, err := c.getSecret(ctx, namespace, NamespacedSecretName)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, nil // Not finding this secret is allowed, allow caller to proceed to higher scope secret
		}
		return nil, err
	}

	armClient, err := c.getARMClientFromSecret(secret)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get armClient from namespaced secret for %q/%q", secret.Namespace, secret.Name)
	}

	return armClient, nil
}

func (c *ARMClientCache) getCredentialFromAnnotation(ctx context.Context, obj genruntime.ARMMetaObject, annotation string) (*armClient, error) {
	credentialFrom, ok := obj.GetAnnotations()[annotation]
	if !ok {
		return nil, nil
	}

	// annotation exists, use specified secret
	secretNamespacedName := getSecretNameFromAnnotation(credentialFrom, obj.GetNamespace())

	secret, err := c.getSecret(ctx, secretNamespacedName.Namespace, secretNamespacedName.Name)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, core.NewSecretNotFoundError(secretNamespacedName, errors.Wrapf(err, "credential secret not found"))
		}
		return nil, err
	}

	armClient, err := c.getARMClientFromSecret(secret)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to get armClient from annotation for %q", secretNamespacedName.String())
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

func (c *ARMClientCache) getARMClientFromSecret(secret *v1.Secret) (*armClient, error) {
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

	newClient, err := genericarmclient.NewGenericClientFromHTTPClient(c.cloudConfig, credential, c.httpClient, subscriptionID, c.armMetrics)
	if err != nil {
		return nil, err
	}

	armClient := newARMClient(newClient, secret.Data, nsName)
	c.register(armClient)
	return armClient, nil
}

func (c *ARMClientCache) newCredentialFromSecret(secret *v1.Secret, nsName types.NamespacedName) (azcore.TokenCredential, string, error) {
	var errs []error
	var err error
	var credential azcore.TokenCredential

	subscriptionID, ok := secret.Data[config.SubscriptionIDVar]
	if !ok {
		err = core.NewSecretNotFoundError(nsName, errors.Errorf("credential Secret %q does not contain key %q", nsName.String(), config.SubscriptionIDVar))
		errs = append(errs, err)
	}

	tenantID, ok := secret.Data[config.TenantIDVar]
	if !ok {
		err = core.NewSecretNotFoundError(nsName, errors.Errorf("credential Secret %q does not contain key %q", nsName.String(), config.TenantIDVar))
		errs = append(errs, err)
	}

	clientID, ok := secret.Data[config.ClientIDVar]
	if !ok {
		err = core.NewSecretNotFoundError(nsName, errors.Errorf("credential Secret %q does not contain key %q", nsName.String(), config.ClientIDVar))
		errs = append(errs, err)
	}

	// Missing required properties, fail fast
	if len(errs) > 0 {
		return nil, "", kerrors.NewAggregate(errs)
	}

	clientSecret, hasClientSecret := secret.Data[config.ClientSecretVar]

	if hasClientSecret {
		credential, err = azidentity.NewClientSecretCredential(string(tenantID), string(clientID), string(clientSecret), nil)
		if err != nil {
			return nil, "", errors.Wrap(err, errors.Errorf("invalid Client Secret Credential for %q encountered", nsName.String()).Error())
		}
	} else {
		// Here we check for workload identity if client secret is not provided.
		credential, err = identity.NewWorkloadIdentityCredential(string(tenantID), string(clientID))
		if err != nil {
			err = errors.Wrapf(
				err,
				"credential secret %q does not contain key %q and failed to get workload identity credential for clientID %q from %q ",
				nsName.String(),
				config.ClientSecretVar,
				string(clientID),
				identity.TokenFile)
			return nil, "", err
		}
	}

	return credential, string(subscriptionID), nil
}

func (c *ARMClientCache) getSecret(ctx context.Context, namespace string, secretName string) (*v1.Secret, error) {
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
