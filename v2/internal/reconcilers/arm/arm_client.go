/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"
	"reflect"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-service-operator/v2/internal/config"
	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
	"github.com/Azure/azure-service-operator/v2/internal/metrics"
	"github.com/Azure/azure-service-operator/v2/internal/resolver"
	"github.com/Azure/azure-service-operator/v2/internal/util/kubeclient"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/pkg/errors"
	v1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
)

const (
	namespacedSecretName = "aso-credential"
)

type armClient struct {
	genericClient  *genericarmclient.GenericClient
	secretData     map[string][]byte
	credentialFrom types.NamespacedName
}

func (c armClient) GenericClient() *genericarmclient.GenericClient {
	return c.genericClient
}

func (c armClient) CredentialFrom() types.NamespacedName {
	return c.credentialFrom
}

type armClients struct {
	lock sync.Mutex
	// clients map hold namespaced name for a secret to all the registered ARM clients
	clients      map[string]*armClient
	globalClient *armClient
}

func NewARMClients(client *genericarmclient.GenericClient, podNamespace string) *armClients {
	globalClient := &armClient{
		genericClient:  client,
		credentialFrom: types.NamespacedName{Name: "aso-controller-settings", Namespace: podNamespace},
	}

	return &armClients{
		lock:         sync.Mutex{},
		clients:      make(map[string]*armClient),
		globalClient: globalClient,
	}
}

func newARMClient(client *genericarmclient.GenericClient, secretData map[string][]byte, credentialFrom types.NamespacedName) *armClient {
	return &armClient{
		genericClient:  client,
		secretData:     secretData,
		credentialFrom: credentialFrom,
	}

}

func (c armClients) Register(client *armClient) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.clients[client.CredentialFrom().String()] = client
}

func (c armClients) Lookup(key string) (*armClient, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	client, ok := c.clients[key]
	return client, ok
}

func (c armClients) GetClient(ctx context.Context, obj genruntime.ARMMetaObject, kubeClient kubeclient.Client, cloud cloud.Configuration) (*armClient, error) {
	// Namespaced secret
	secret, err := getSecret(ctx, kubeClient, obj.GetNamespace(), namespacedSecretName)
	if err != nil {
		return nil, err
	}

	if secret != nil {
		nsName := types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}
		client, ok := c.Lookup(nsName.String())
		// if client does not already exist, or secret data is changed, we need to get new token.
		if !ok || !matchSecretData(client.secretData, secret.Data) {
			credential, subscriptionID, err := c.newCredentialFromSecret(secret, nsName)
			if err != nil {
				return nil, err
			}

			newClient, err := genericarmclient.NewGenericClient(cloud, credential, subscriptionID, metrics.NewARMClientMetrics())
			if err != nil {
				return nil, err
			}

			armClient := newARMClient(newClient, secret.Data, nsName)
			c.Register(armClient)
			return armClient, nil
		} else {
			return client, nil
		}
	}
	// If not found, default would be global
	return c.globalClient, nil
}

func (c armClients) newCredentialFromSecret(secret *v1.Secret, nsName types.NamespacedName) (azcore.TokenCredential, string, error) {
	subscriptionID, ok := secret.Data[config.SubscriptionIDVar]
	if !ok {
		return nil, "", resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.SubscriptionIDVar))
	}
	tenantID, ok := secret.Data[config.TenantIDVar]
	if !ok {
		return nil, "", resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.TenantIDVar))
	}
	clientID, ok := secret.Data[config.AzureClientID]
	if !ok {
		return nil, "", resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.AzureClientID))
	}
	clientSecret, ok := secret.Data[config.AzureClientSecret]
	// TODO: We check this for now until we support Workload Identity. When we support workload identity for multitenancy, !ok would mean to check for workload identity.
	if !ok {
		return nil, "", resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.AzureClientSecret))
	}

	credential, err := azidentity.NewClientSecretCredential(string(tenantID), string(clientID), string(clientSecret), nil)
	if err != nil {
		return nil, "", errors.Wrap(err, errors.Errorf("Invalid Client Secret Credential for %q encountered", nsName.String()).Error())
	}
	return credential, string(subscriptionID), nil
}

func getSecret(ctx context.Context, kubeClient kubeclient.Client, namespace string, secretName string) (*v1.Secret, error) {
	secret := &v1.Secret{}

	err := kubeClient.Get(
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
