/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package entra

import (
	"context"
	"net/http"
	"sync"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/cloud"
	msgraphsdk "github.com/microsoftgraph/msgraph-sdk-go"
	msgraphauth "github.com/microsoftgraph/msgraph-sdk-go-core/authentication"

	"github.com/Azure/azure-service-operator/v2/internal/identity"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

// EntraClientCache is a cache for entraClients to hold multiple credential clients and global credential client.
type EntraClientCache struct {
	lock sync.Mutex

	clients            map[string]*entraClient // clients allows quick lookup of an entraClient for each namespace
	cloudConfig        cloud.Configuration     //!! What's the entra equivalent?
	credentialProvider identity.CredentialProvider
	httpClient         *http.Client
	//!! TODO entraMetrics         *metrics.EntraClientMetrics
}

func NewEntraClientCache(
	credentialProvider identity.CredentialProvider,
	configuration cloud.Configuration,
	httpClient *http.Client,
	//!! armMetrics *metrics.ARMClientMetrics,
) *EntraClientCache {
	return &EntraClientCache{
		lock:               sync.Mutex{},
		clients:            make(map[string]*entraClient),
		cloudConfig:        configuration,
		credentialProvider: credentialProvider,
		httpClient:         httpClient,
	}
}

// register adds a new client to the cache.
// Any existing client with the same key will be replaced.
func (c *EntraClientCache) register(client *entraClient) {
	c.lock.Lock()
	defer c.lock.Unlock()

	key := client.credential.CredentialFrom().String()
	c.clients[key] = client
}

// lookup retrieves a client from the cache.
// Returns the client and true if found, or nil and false if not found.
func (c *EntraClientCache) lookup(key string) (*entraClient, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()

	client, ok := c.clients[key]
	return client, ok
}

// GetConnection finds and returns connection details to be used for a given resource
func (c *EntraClientCache) GetConnection(
	ctx context.Context,
	obj genruntime.EntraMetaObject,
) (Connection, error) {
	cred, err := c.credentialProvider.GetCredential(ctx, obj)
	if err != nil {
		return nil, err
	}

	client, err := c.getEntraClientFromCredential(cred)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func (c *EntraClientCache) getEntraClientFromCredential(
	cred *identity.Credential,
) (*entraClient, error) {
	// First try to find the client in the cache, in case we've already created it
	client, ok := c.lookup(cred.CredentialFrom().String())
	if ok && cred.SecretsEqual(client.credential) {
		return client, nil
	}

	// Create an AzureIdentityAuthenticationProvider to reuse the Azure credential we already have
	authProvider, err := msgraphauth.NewAzureIdentityAuthenticationProvider(cred.TokenCredential())
	if err != nil {
		return nil, err
	}

	// Create a GraphRequestAdapter with the authentication provider
	// We need to use this factory method in order to set the httpClient, allowing for test recordings
	requestAdapter, err := msgraphsdk.NewGraphRequestAdapterWithParseNodeFactoryAndSerializationWriterFactoryAndHttpClient(
		authProvider,
		nil, // ParseNodeFactory
		nil, // SerializationWriterFactory
		c.httpClient)
	if err != nil {
		return nil, err
	}

	newClient := msgraphsdk.NewGraphServiceClient(requestAdapter)

	entraClient := newEntraClient(newClient, cred)
	c.register(entraClient)

	return entraClient, nil
}
