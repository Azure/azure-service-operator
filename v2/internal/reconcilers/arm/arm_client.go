/*
 * Copyright (c) Microsoft Corporation.
 * Licensed under the MIT license.
 */

package arm

import (
	"k8s.io/apimachinery/pkg/types"

	"github.com/Azure/azure-service-operator/v2/internal/genericarmclient"
)

type armClient struct {
	genericClient  *genericarmclient.GenericClient
	secretData     map[string][]byte
	credentialFrom types.NamespacedName
}

func newARMClient(client *genericarmclient.GenericClient, secretData map[string][]byte, credentialFrom types.NamespacedName) *armClient {
	return &armClient{
		genericClient:  client,
		secretData:     secretData,
		credentialFrom: credentialFrom,
	}
}

func (c *armClient) GenericClient() *genericarmclient.GenericClient {
	return c.genericClient
}

func (c *armClient) CredentialFrom() string {
	return c.credentialFrom.String()
}
