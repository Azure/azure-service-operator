/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"context"
	"reflect"
	"sync"

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

type ARMClientFactory func(genruntime.ARMMetaObject, context.Context) (*genericarmclient.GenericClient, error)

type ARMClient struct {
	genericClient *genericarmclient.GenericClient
	secretData    map[string][]byte
}

type ARMClients struct {
	lock         sync.Mutex
	clients      map[string]ARMClient
	globalClient *genericarmclient.GenericClient
}

func NewARMClients(globalClient *genericarmclient.GenericClient) ARMClients {
	return ARMClients{
		lock:         sync.Mutex{},
		clients:      make(map[string]ARMClient),
		globalClient: globalClient,
	}
}

func newARMClient(client *genericarmclient.GenericClient, secretData map[string][]byte) ARMClient {
	return ARMClient{
		genericClient: client,
		secretData:    secretData,
	}
}

func (c ARMClients) Register(key string, client *genericarmclient.GenericClient, secretData map[string][]byte) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.clients[key] = newARMClient(client, secretData)
}

func (c ARMClients) Lookup(key string) (ARMClient, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	client, ok := c.clients[key]
	return client, ok
}

func (c ARMClients) GetClient(obj genruntime.ARMMetaObject, kubeClient kubeclient.Client, ctx context.Context, cloud cloud.Configuration) (*genericarmclient.GenericClient, error) {
	// Namespaced secret
	secret, found, err := getSecret(obj, kubeClient, ctx, namespacedSecretName)
	if err != nil {
		return nil, err
	}
	if found {
		nsName := types.NamespacedName{Namespace: secret.Namespace, Name: secret.Name}
		client, ok := c.Lookup(nsName.String())
		// if client does not already exist, or secret data is changed, we need to get new token.
		if !ok || !matchSecretData(client.secretData, secret.Data) {
			client, err := c.newClientFromSecret(secret, nsName, cloud)
			if err != nil {
				return nil, err
			}

			c.Register(nsName.String(), client, secret.Data)
			return client, nil
		} else {
			return client.genericClient, nil
		}
	}
	// If not found, default would be global
	return c.globalClient, nil
}

func (c ARMClients) newClientFromSecret(secret *v1.Secret, nsName types.NamespacedName, cloud cloud.Configuration) (*genericarmclient.GenericClient, error) {
	subscriptionID, ok := secret.Data[config.SubscriptionIDVar]
	if !ok {
		return nil, resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.SubscriptionIDVar))
	}
	tenantID, ok := secret.Data[config.TenantIDVar]
	if !ok {
		return nil, resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.TenantIDVar))
	}
	clientID, ok := secret.Data[config.AzureClientID]
	if !ok {
		return nil, resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.AzureClientID))
	}
	clientSecret, ok := secret.Data[config.AzureClientSecret]
	// TODO: We check this for now until we support Workload Identity. When we support workload identity for multitenancy, !ok would mean to check for workload identity.
	if !ok {
		return nil, resolver.NewSecretNotFoundError(nsName, errors.Errorf("Credential Secret %q does not contain key %q", nsName.String(), config.AzureClientSecret))
	}

	credential, err := azidentity.NewClientSecretCredential(string(tenantID), string(clientID), string(clientSecret), nil)
	if err != nil {
		return nil, errors.Wrap(err, errors.Errorf("Invalid Client Secret Credential for %q encountered", nsName.String()).Error())
	}

	client, err := genericarmclient.NewGenericClient(cloud, credential, string(subscriptionID), metrics.NewARMClientMetrics(), nsName)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getSecret(obj genruntime.ARMMetaObject, kubeClient kubeclient.Client, ctx context.Context, secretName string) (*v1.Secret, bool, error) {
	secret := &v1.Secret{}

	err := kubeClient.Get(
		ctx,
		types.NamespacedName{Namespace: obj.GetNamespace(), Name: secretName},
		secret)
	if err != nil {
		if apierrors.IsNotFound(err) {
			return nil, false, nil
		} else {
			return nil, false, err
		}
	}
	return secret, true, nil
}

func matchSecretData(old, new map[string][]byte) bool {
	return reflect.DeepEqual(old, new)
}
