/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

// armClient is a wrapper around generic client to keep a track of secretData used to create it and credentialFrom which
// that secret was retrieved.
type armClient struct {
	genericClient  *genericarmclient.GenericClient
	secretData     map[string][]byte
	credentialFrom types.NamespacedName
	subscriptionID string
}

func newARMClient(
	client *genericarmclient.GenericClient,
	secretData map[string][]byte,
	credentialFrom types.NamespacedName,
	subscriptionID string) *armClient {
	return &armClient{
		genericClient:  client,
		secretData:     secretData,
		credentialFrom: credentialFrom,
		subscriptionID: subscriptionID,
	}
}

func (c *armClient) GenericClient() *genericarmclient.GenericClient {
	return c.genericClient
}

func (c *armClient) CredentialFrom() string {
	return c.credentialFrom.String()
}

// Connection describes a client + credentials set used to connect to Azure.
type Connection struct {
	Client         *genericarmclient.GenericClient
	CredentialFrom types.NamespacedName
	SubscriptionID string
}

func newConnection(client *armClient) *Connection {
	return &Connection{
		Client:         client.genericClient,
		CredentialFrom: client.credentialFrom,
		SubscriptionID: client.subscriptionID,
	}
}
