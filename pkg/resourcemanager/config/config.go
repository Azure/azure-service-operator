// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Package config manages loading configuration from environment and command-line params
package config

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/marstr/randname"
)

var (
	// these are our *global* config settings, to be shared by all packages.
	// each has corresponding public accessors below.
	// if anything requires a `Set` accessor, that indicates it perhaps
	// shouldn't be set here, because mutable vars shouldn't be global.
	clientID               string
	clientSecret           string
	tenantID               string
	subscriptionID         string
	locationDefault        string
	authorizationServerURL string
	cloudName              string
	useDeviceFlow          bool
	useMI                  bool
	buildID                string
	keepResources          bool
	operatorKeyvault       string
	userAgent              string
	environment            *azure.Environment


	keepResources    bool
	operatorKeyvault string
	groupName        string // deprecated, use baseGroupName instead
	baseGroupName    string
	userAgent        string
	baseURI          string
	environment      *azure.Environment
	testResourcePrefix string // used to generate resource names in tests, should probably exist in a test only package

)

// ClientID is the OAuth client ID.
func ClientID() string {
	return clientID
}

// ClientSecret is the OAuth client secret.
func ClientSecret() string {
	return clientSecret
}

// TenantID is the AAD tenant to which this client belongs.
func TenantID() string {
	return tenantID
}

// SubscriptionID is a target subscription for Azure resources.
func SubscriptionID() string {
	return subscriptionID
}

// deprecated: use DefaultLocation() instead
// Location returns the Azure location to be utilized.
func Location() string {
	return locationDefault
}

// DefaultLocation() returns the default location wherein to create new resources.
// Some resource types are not available in all locations so another location might need
// to be chosen.
func DefaultLocation() string {
	return locationDefault
}

// AuthorizationServerURL is the OAuth authorization server URL.
// Q: Can this be gotten from the `azure.Environment` in `Environment()`?
func AuthorizationServerURL() string {
	return authorizationServerURL
}

// OperatorKeyvault() specifies the keyvault the operator should use to store secrets
func OperatorKeyvault() string {
	return operatorKeyvault
}

// UseDeviceFlow() specifies if interactive auth should be used. Interactive
// auth uses the OAuth Device Flow grant type.
func UseDeviceFlow() bool {
	return useDeviceFlow
}

// UseMI() specifies if managed service identity auth should be used. Used for
// aad-pod-identity
func UseMI() bool {
	return useMI
}

// KeepResources() specifies whether to keep resources created by samples.
func KeepResources() bool {
	return keepResources
}

// UserAgent() specifies a string to append to the agent identifier.
func UserAgent() string {
	if len(userAgent) > 0 {
		return userAgent
	}
	return "sdk-samples"
}

// TestResourcePrefix specifies a string to prefix test resources with
func TestResourcePrefix() string {
	return testResourcePrefix
}

// Environment returns an `azure.Environment{...}` for the current cloud.
func Environment() *azure.Environment {
	if environment != nil {
		return environment
	}

	env, err := azure.EnvironmentFromName(cloudName)
	if err != nil {
		// TODO: move to initialization of var
		panic(fmt.Sprintf(
			"invalid cloud name '%s' specified, cannot continue\n", cloudName))
	}
	environment = &env
	return environment
}

// AppendRandomSuffix will append a suffix of five random characters to the specified prefix.
func AppendRandomSuffix(prefix string) string {
	return randname.GenerateWithPrefix(prefix, 5)
}

// Base URI for provisioning
func BaseURI() string {
	return baseURI
}
