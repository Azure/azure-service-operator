// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Package config manages loading configuration from environment and command-line params
package config

import (
	"fmt"

	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/marstr/randname"

	"github.com/Azure/azure-service-operator/pkg/secrets"
)

var (
	// these are our *global* config settings, to be shared by all packages.
	// each has corresponding public accessors below.
	// if anything requires a `Set` accessor, that indicates it perhaps
	// shouldn't be set here, because mutable vars shouldn't be global.

	// TODO: eliminate this!
	creds                             credentials
	locationDefault                   string
	authorizationServerURL            string
	cloudName                         string
	useDeviceFlow                     bool
	buildID                           string
	keepResources                     bool
	userAgent                         string
	baseURI                           string
	environment                       *azure.Environment
	podNamespace                      string
	targetNamespaces                  []string
	secretNamingVersion               secrets.SecretNamingVersion
	operatorMode                      OperatorMode
	purgeDeletedKeyVaultSecrets       bool
	recoverSoftDeletedKeyVaultSecrets bool
	testResourcePrefix                string // used to generate resource names in tests, should probably exist in a test only package
)

// GlobalCredentials returns the configured credentials.
// TODO: get rid of all uses of this.
func GlobalCredentials() Credentials {
	return creds
}

// deprecated: use DefaultLocation() instead
// Location returns the Azure location to be utilized.
func Location() string {
	return locationDefault
}

// DefaultLocation returns the default location wherein to create new resources.
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

// UseDeviceFlow specifies if interactive auth should be used. Interactive
// auth uses the OAuth Device Flow grant type.
func UseDeviceFlow() bool {
	return useDeviceFlow
}

// KeepResources specifies whether to keep resources created by samples.
func KeepResources() bool {
	return keepResources
}

// UserAgent specifies a string to append to the agent identifier.
func UserAgent() string {
	if len(userAgent) > 0 {
		return userAgent
	}
	return "azure-service-operator-v1"
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

// PodNamespace returns the namespace the manager pod is running in
func PodNamespace() string {
	return podNamespace
}

// TargetNamespaces returns the namespaces the operator should watch for resources.
func TargetNamespaces() []string {
	return targetNamespaces
}

// SelectedMode returns the mode configuration for the operator.
func SelectedMode() OperatorMode {
	return operatorMode
}

// AppendRandomSuffix will append a suffix of five random characters to the specified prefix.
func AppendRandomSuffix(prefix string) string {
	return randname.GenerateWithPrefix(prefix, 5)
}

// Base URI for provisioning
func BaseURI() string {
	return baseURI
}

// SecretNamingVersion is the version of secret naming that the operator uses
func SecretNamingVersion() secrets.SecretNamingVersion {
	return secretNamingVersion
}

// PurgeDeletedKeyVaultSecrets determines if the operator should issue a secret Purge request in addition
// to Delete when deleting secrets in Azure Key Vault. This only applies to secrets that are stored in Azure Key Vault.
// It does nothing if the secret is stored in Kubernetes.
func PurgeDeletedKeyVaultSecrets() bool {
	return purgeDeletedKeyVaultSecrets
}

// RecoverSoftDeletedKeyVaultSecrets determines if the operator should issue a secret Recover request when it
// encounters an "ObjectIsDeletedButRecoverable" error from Azure Key Vault during secret creation. This error
// can occur when a Key Vault has soft delete enabled and an ASO resource was deleted and recreated with the same name.
// This only applies to secrets that are stored in Azure Key Vault.
// It does nothing if the secret is stored in Kubernetes.
func RecoverSoftDeletedKeyVaultSecrets() bool {
	return recoverSoftDeletedKeyVaultSecrets
}

// ConfigString returns the parts of the configuration file with are not secrets as a string for easy logging
func ConfigString() string {
	creds := GlobalCredentials()
	return fmt.Sprintf(
		"clientID: %q, tenantID: %q, subscriptionID: %q, cloudName: %q, useDeviceFlow: %t, useManagedIdentity: %t, operatorMode: %s, targetNamespaces: %s,"+
			" podNamespace: %q, secretNamingVersion: %q, purgeDeletedKeyVaultSecrets: %t, recoverSoftDeletedkeyVaultSecrets: %t",
		creds.ClientID(),
		creds.TenantID(),
		creds.SubscriptionID(),
		cloudName,
		UseDeviceFlow(),
		creds.UseManagedIdentity(),
		operatorMode,
		targetNamespaces,
		podNamespace,
		SecretNamingVersion(),
		PurgeDeletedKeyVaultSecrets(),
		RecoverSoftDeletedKeyVaultSecrets())
}
