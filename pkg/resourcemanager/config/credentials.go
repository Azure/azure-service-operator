// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Package config manages loading configuration from environment and command-line params
package config

// Credentials is a read-only holder for information needed to
// authenticate to ARM.
type Credentials interface {
	ClientID() string
	ClientSecret() string
	TenantID() string
	SubscriptionID() string
	UseManagedIdentity() bool
	OperatorKeyvault() string

	WithSubscriptionID(subscriptionID string) Credentials
}

type credentials struct {
	clientID           string
	clientSecret       string
	tenantID           string
	subscriptionID     string
	useManagedIdentity bool
	operatorKeyvault   string
}

var _ Credentials = credentials{}

// ClientID is the OAuth client ID.
func (c credentials) ClientID() string {
	return c.clientID
}

// ClientSecret is the OAuth client secret.
func (c credentials) ClientSecret() string {
	return c.clientSecret
}

// TenantID is the AAD tenant to which this client belongs.
func (c credentials) TenantID() string {
	return c.tenantID
}

// SubscriptionID is a target subscription for Azure resources.
func (c credentials) SubscriptionID() string {
	return c.subscriptionID
}

// UseMI() specifies if managed service identity auth should be used. Used for
// aad-pod-identity
func (c credentials) UseManagedIdentity() bool {
	return c.useManagedIdentity
}

// OperatorKeyvault() specifies the keyvault the operator should use to store secrets
func (c credentials) OperatorKeyvault() string {
	return c.operatorKeyvault
}

// WithSubscriptionID returns the credentials object with a different subscription ID.
// The original credentials object is not modified
func (c credentials) WithSubscriptionID(subscriptionID string) Credentials {
	c.subscriptionID = subscriptionID

	return c
}
